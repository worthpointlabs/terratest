package main

import (
	"go/ast"
	"go/parser"
	"go/token"
	"strings"
)

// getAllTopLevelFunctions takes in an AST representation of a golang package and returns a map that maps the function
// names to their function bodies. This only handles functions that are declared at the top level within the package;
// does not handle inline function declarations or functions declared outside the package.
func getAllTopLevelFunctions(pkg *ast.Package) map[string]*ast.BlockStmt {
	funcs := map[string]*ast.BlockStmt{}
	for _, file := range pkg.Files {
		for _, decl := range file.Decls {
			switch typedDecl := decl.(type) {
			case *ast.FuncDecl:
				funcs[typedDecl.Name.Name] = typedDecl.Body
			}
		}
	}
	return funcs
}

// getTopLevelTestFunctions takes a map of function names to function bodies and filters it down to only those functions
// that have names that start with the string "Test".
func getTopLevelTestFunctions(allFuncs map[string]*ast.BlockStmt) map[string]*ast.BlockStmt {
	testFuncs := map[string]*ast.BlockStmt{}
	for funcName, funcBody := range allFuncs {
		if strings.HasPrefix(funcName, "Test") {
			testFuncs[funcName] = funcBody
		}
	}
	return testFuncs
}

func getPackageFuncBodyFromCall(allFuncs map[string]*ast.BlockStmt, callExpr *ast.CallExpr) *ast.BlockStmt {
	switch fun := callExpr.Fun.(type) {
	case *ast.Ident:
		return allFuncs[fun.Name]
	case *ast.FuncLit:
		return fun.Body
	}
	return nil
}

// getTestPackage parses the given directory path and returns the golang AST corresponding to the requested package
// name.
func getTestPackage(testPackagePath string, testPackageName string) (*ast.Package, error) {
	fset := token.NewFileSet()
	pkgs, err := parser.ParseDir(fset, testPackagePath, nil, parser.AllErrors)
	if err != nil {
		return nil, err
	}
	pkg, hasPkg := pkgs[testPackageName]
	if !hasPkg {
		return nil, NoPackageError{path: testPackagePath, name: testPackageName}
	}
	return pkg, nil
}

// getAllTerratestCalls walks the AST of a function body to scan for all terratest function calls (including nested and
// deferred), where a terratest function call is one of the function calls defined in terratestFuncCallType. This
// returns the list of terratest function calls in execution order, and not definition order. Meaning, deferred calls
// will be at the end of the list.
func getAllTerratestCalls(allFuncs map[string]*ast.BlockStmt, funcBody *ast.BlockStmt) ([]terratestFuncCall, error) {
	stmts := funcBody.List
	allCalls := []terratestFuncCall{}
	// We track deferred calls separately so that they can be added to the end of the list, after all the terratest
	// calls have been already logged.
	deferredCalls := []terratestFuncCall{}
	for _, stmt := range stmts {
		// TODO: Handle for loops and if blocks
		switch typedStmt := stmt.(type) {
		case *ast.ExprStmt:
			expr := typedStmt.X
			if isFunctionCall(expr) {
				callExpr := expr.(*ast.CallExpr)
				calls, err := terratestFuncCallsFromCallExpr(allFuncs, callExpr)
				if err != nil {
					return allCalls, err
				}
				allCalls = append(allCalls, calls...)
			}
			// NOTE: We ignore non function call expressions
		case *ast.DeferStmt:
			calls, err := terratestFuncCallsFromCallExpr(allFuncs, typedStmt.Call)
			if err != nil {
				return allCalls, err
			}
			deferredCalls = append(calls, deferredCalls...)
		}
	}
	// Take all the deferred calls and add to the end of the list.
	for _, call := range deferredCalls {
		allCalls = append(allCalls, call)
	}
	return allCalls, nil
}

// terratestFuncCallsFromCallExpr takes a function call expression and extracts all the terratest function calls, where
// a terratest function call is one of the function calls defined in terratestFuncCallType. If and only if the given
// function call is a call to a function that is defined within the package, this function will recurse into the defined
// function body to look for terratest function calls from there.
func terratestFuncCallsFromCallExpr(allFuncs map[string]*ast.BlockStmt, callExpr *ast.CallExpr) ([]terratestFuncCall, error) {
	allCalls := []terratestFuncCall{}
	if isTestStageCall(callExpr) {
		allCalls = append(
			allCalls,
			terratestFuncCall{
				callType: terratestRunTestStage,
				expr:     callExpr,
			},
		)
	} else if isTRunCall(callExpr) {
		allCalls = append(
			allCalls,
			terratestFuncCall{
				callType: terratestTRun,
				expr:     callExpr,
			},
		)
	} else if packageFuncBody := getPackageFuncBodyFromCall(allFuncs, callExpr); packageFuncBody != nil {
		// Enter package function by recursing on package function body
		calls, err := getAllTerratestCalls(allFuncs, packageFuncBody)
		if err != nil {
			return allCalls, err
		}
		allCalls = append(allCalls, calls...)
	}
	return allCalls, nil
}

// getStageFromRunTestStageCall takes the AST representation of a test_structure.RunTestStage call and returns the
// value of the second argument to the function.
func getStageFromRunTestStageCall(runTestStageCallExpr *ast.CallExpr) (string, error) {
	stageNameArgExpr := runTestStageCallExpr.Args[1]
	switch typedStageNameArgExpr := stageNameArgExpr.(type) {
	case *ast.BasicLit:
		if typedStageNameArgExpr.Kind == token.STRING {
			return strings.Trim(typedStageNameArgExpr.Value, "\""), nil
		}
	}
	return "", NoStageNameError{}
}

// getTestNameFromTRunCall takes a call expression representing a function call to `t.Run` and extracts out the test
// name, which is the first argument to the function.
// NOTE: This does not handle complex t.Run calls, like using a range expression.
func getTestNameFromTRunCall(tRunCallExpr *ast.CallExpr) (string, error) {
	testNameArgExpr := tRunCallExpr.Args[0]
	switch testNameArg := testNameArgExpr.(type) {
	case *ast.BasicLit:
		if testNameArg.Kind == token.STRING {
			return strings.Trim(testNameArg.Value, "\""), nil
		}
	}
	return "", NoTestNameError{}
}

// getTestFuncFromTRunCall extracts the function body of the test from the `t.Run` function call, which is the second
// argument to the function. This handles both function literals, and test functions that call a function defined in the
// package.
// NOTE: This does not handle test function calls not defined within the package.
func getTestFuncFromTRunCall(allFuncs map[string]*ast.BlockStmt, tRunCallExpr *ast.CallExpr) (*ast.BlockStmt, error) {
	testFuncArgExpr := tRunCallExpr.Args[1]
	switch testFuncArg := testFuncArgExpr.(type) {
	case *ast.Ident:
		return allFuncs[testFuncArg.Name], nil
	case *ast.FuncLit:
		return testFuncArg.Body, nil
	}
	return nil, NoTestBodyError{}
}

// isTestStageCall returns true if the function call is a terratest call to define a test stage
// (test_structure.RunTestStage)
func isTestStageCall(callExpr *ast.CallExpr) bool {
	selector, isSelector := callExpr.Fun.(*ast.SelectorExpr)
	if !isSelector {
		return false
	}

	ident, isIdent := selector.X.(*ast.Ident)
	if !isIdent {
		return false
	}

	return ident.Name == "test_structure" && selector.Sel.Name == "RunTestStage"
}

// isTRunCall returns true if the function call is a call to t.Run.
func isTRunCall(callExpr *ast.CallExpr) bool {
	selector, isSelector := callExpr.Fun.(*ast.SelectorExpr)
	if !isSelector {
		return false
	}

	ident, isIdent := selector.X.(*ast.Ident)
	if !isIdent {
		return false
	}

	return ident.Name == "t" && selector.Sel.Name == "Run"
}

// isFunctionCall returns true if the given expression is a function call expression
func isFunctionCall(expr ast.Expr) bool {
	_, isCallExpr := expr.(*ast.CallExpr)
	return isCallExpr
}
