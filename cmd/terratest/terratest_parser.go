package main

import (
	"fmt"
	"go/ast"
	goparser "go/parser"
	"go/token"
	"strings"
)

type Parser struct {
	curLevel int
	allFuncs map[string]*ast.BlockStmt
}

type terratestStage struct {
	level int
	name  string
}

type terratestFuncCallType int

const (
	terratestRunTestStage terratestFuncCallType = iota
	terratestTRun
)

type terratestFuncCall struct {
	callType terratestFuncCallType
	expr     *ast.CallExpr
}

func (parser *Parser) parseTestPackage(testPackagePath string, testPackageName string) (map[string][]terratestStage, error) {
	pkg, err := parser.getTestPackage(testPackagePath, testPackageName)
	if err != nil {
		return nil, err
	}

	parser.allFuncs = parser.getAllTopLevelFunctions(pkg)
	testFuncs := parser.getTopLevelTestFunctions()

	testFuncToStages := map[string][]terratestStage{}
	for testFuncName, testFunc := range testFuncs {
		terratestFuncCalls, err := parser.getAllTerratestCalls(testFunc)
		if err != nil {
			return nil, err
		}
		allStages, nestedStages, err := parser.getStagesFromTerratestFuncCalls(terratestFuncCalls)
		if err != nil {
			return nil, err
		}
		testFuncToStages[testFuncName] = allStages
		for nestedTestName, stages := range nestedStages {
			testFuncToStages[fmt.Sprintf("%s/%s", testFuncName, nestedTestName)] = stages
		}
	}
	return testFuncToStages, nil
}

func (parser *Parser) getTestPackage(testPackagePath string, testPackageName string) (*ast.Package, error) {
	fset := token.NewFileSet()
	pkgs, err := goparser.ParseDir(fset, testPackagePath, nil, goparser.AllErrors)
	if err != nil {
		return nil, err
	}
	pkg, hasPkg := pkgs[testPackageName]
	if !hasPkg {
		// TODO: return concrete error
		return nil, fmt.Errorf("%s does not have go package %s", testPackagePath, testPackageName)
	}
	return pkg, nil
}

func (parser *Parser) getAllTopLevelFunctions(pkg *ast.Package) map[string]*ast.BlockStmt {
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

func (parser *Parser) getTopLevelTestFunctions() map[string]*ast.BlockStmt {
	testFuncs := map[string]*ast.BlockStmt{}
	for funcName, funcBody := range parser.allFuncs {
		if strings.HasPrefix(funcName, "Test") {
			testFuncs[funcName] = funcBody
		}
	}
	return testFuncs
}

func (parser *Parser) getStagesFromTerratestFuncCalls(calls []terratestFuncCall) ([]terratestStage, map[string][]terratestStage, error) {
	stages := []terratestStage{}
	nestedStages := map[string][]terratestStage{}
	for _, call := range calls {
		switch call.callType {
		case terratestRunTestStage:
			stageName, err := parser.getStageFromRunTestStageCall(call.expr)
			if err != nil {
				return nil, nil, err
			}
			// We need to append to both the current stage list and the nested list, so the nested tests get all the
			// deferred stages.
			stage := terratestStage{level: parser.curLevel, name: stageName}
			stages = append(stages, stage)
			for key, val := range nestedStages {
				nestedStages[key] = append(val, stage)
			}
		case terratestTRun:
			foo, bar, err := parser.handleNestedTest(call)
			if err != nil {
				return nil, nil, err
			}

			stages = append(stages, foo...)

			stagesToInclude := []terratestStage{}
			for _, stage := range stages {
				if stage.level <= parser.curLevel {
					stagesToInclude = append(stagesToInclude, stage)
				}
			}

			for key, val := range bar {
				nestedStages[key] = append(stagesToInclude, val...)
			}
		}
	}
	return stages, nestedStages, nil
}

func (parser *Parser) handleNestedTest(call terratestFuncCall) ([]terratestStage, map[string][]terratestStage, error) {
	defer func() {
		parser.curLevel--
	}()
	parser.curLevel++

	nestedStages := map[string][]terratestStage{}

	// Get the nested stages, as well as the stages from additional t.Run calls within the subtest, and merge to
	// the top level list.
	nestedTestName, err := parser.getTestNameFromTRunCall(call.expr)
	if err != nil {
		return nil, nil, err
	}
	testFunc, err := parser.getTestFuncFromTRunCall(call.expr)
	if err != nil {
		return nil, nil, err
	}
	tRunCalls, err := parser.getAllTerratestCalls(testFunc)
	if err != nil {
		return nil, nil, err
	}
	tRunStages, tRunNestedStages, err := parser.getStagesFromTerratestFuncCalls(tRunCalls)
	if err != nil {
		return nil, nil, err
	}

	// Make sure to capture the stages for nested test runs, to be merged into the top level. Note that we
	// import all the stages we have so far so that they are included in the nested tests
	nestedStages[nestedTestName] = tRunStages

	// Merge stages for tests spawned from additional t.Run calls within this subtest.
	for nestedNestedTestName, nestedNestedStages := range tRunNestedStages {
		// NOTE: we don't need to append the tRunStages here because they are already included.
		nestedStages[fmt.Sprintf("%s/%s", nestedTestName, nestedNestedTestName)] = nestedNestedStages
	}

	return tRunStages, nestedStages, nil

}

func (parser *Parser) getAllTerratestCalls(funcBody *ast.BlockStmt) ([]terratestFuncCall, error) {
	stmts := funcBody.List
	allCalls := []terratestFuncCall{}
	deferredCalls := []terratestFuncCall{}
	for _, stmt := range stmts {
		switch typedStmt := stmt.(type) {
		case *ast.ExprStmt:
			expr := typedStmt.X
			if isFunctionCall(expr) {
				callExpr := expr.(*ast.CallExpr)
				calls, err := parser.terratestFuncCallsFromCallExpr(callExpr)
				if err != nil {
					return allCalls, err
				}
				allCalls = append(allCalls, calls...)
			}
		case *ast.DeferStmt:
			calls, err := parser.terratestFuncCallsFromCallExpr(typedStmt.Call)
			if err != nil {
				return allCalls, err
			}
			deferredCalls = append(calls, deferredCalls...)
		}
	}
	for _, call := range deferredCalls {
		allCalls = append(allCalls, call)
	}
	return allCalls, nil
}

func (parser *Parser) terratestFuncCallsFromCallExpr(callExpr *ast.CallExpr) ([]terratestFuncCall, error) {
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
	} else if packageFuncBody := parser.getPackageFuncBodyFromCall(callExpr); packageFuncBody != nil {
		calls, err := parser.getAllTerratestCalls(packageFuncBody)
		if err != nil {
			return allCalls, err
		}
		allCalls = append(allCalls, calls...)
	}
	return allCalls, nil
}

func isFunctionCall(expr ast.Expr) bool {
	_, isCallExpr := expr.(*ast.CallExpr)
	return isCallExpr
}

// getStageFromRunTestStageCall takes the AST representation of a test_structure.RunTestStage call and returns the
// value of the second argument to the function.
func (parser *Parser) getStageFromRunTestStageCall(runTestStageCallExpr *ast.CallExpr) (string, error) {
	stageNameArgExpr := runTestStageCallExpr.Args[1]
	switch typedStageNameArgExpr := stageNameArgExpr.(type) {
	case *ast.BasicLit:
		if typedStageNameArgExpr.Kind == token.STRING {
			return strings.Trim(typedStageNameArgExpr.Value, "\""), nil
		}
	}
	return "", fmt.Errorf("Could not find stage name from RunTestStage call")
}

// TODO: This does not handle complex t.Run calls, like using a range expression. We may have to execute?
func (parser *Parser) getTestNameFromTRunCall(tRunCallExpr *ast.CallExpr) (string, error) {
	testNameArgExpr := tRunCallExpr.Args[0]
	switch testNameArg := testNameArgExpr.(type) {
	case *ast.BasicLit:
		if testNameArg.Kind == token.STRING {
			return strings.Trim(testNameArg.Value, "\""), nil
		}
	}
	return "", fmt.Errorf("Could not find test name from tRun call")
}

func (parser *Parser) getTestFuncFromTRunCall(tRunCallExpr *ast.CallExpr) (*ast.BlockStmt, error) {
	testFuncArgExpr := tRunCallExpr.Args[1]
	switch testFuncArg := testFuncArgExpr.(type) {
	case *ast.Ident:
		return parser.allFuncs[testFuncArg.Name], nil
	case *ast.FuncLit:
		return testFuncArg.Body, nil
	}
	return nil, fmt.Errorf("Could not find test func body from tRun call")
}

func isTestStageCall(callExpr *ast.CallExpr) bool {
	selector, isSelector := callExpr.Fun.(*ast.SelectorExpr)
	if !isSelector {
		return false
	}
	return selector.Sel.Name == "RunTestStage"
}

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

func (parser *Parser) getPackageFuncBodyFromCall(callExpr *ast.CallExpr) *ast.BlockStmt {
	switch fun := callExpr.Fun.(type) {
	case *ast.Ident:
		return parser.allFuncs[fun.Name]
	case *ast.FuncLit:
		return fun.Body
	}
	return nil
}
