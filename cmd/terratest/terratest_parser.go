package main

import (
	"fmt"
	"go/ast"
)

// terratestParser can be used for parsing go tests that use terratest. The stuct is used to track the following states
// during parsing:
// - Level of nested tests (e.g., each t.Run calls increment the level). This is used to know which stages belong to
//   nested tests.
// - All top level functions in the test package. This is used for walking function calls.
type terratestParser struct {
	curLevel int
	allFuncs map[string]*ast.BlockStmt
}

// terratestStage represents a single stage in terratest (e.g., test_structure.RunTestStage calls). The struct captures
// which test level it belongs to, and the name of the stage.
type terratestStage struct {
	level int
	name  string
}

// terratestFuncCallType are types of function calls that are relevant to terratest. Currently these are:
// - test_structure.RunTestStage
// - t.Run
type terratestFuncCallType int

const (
	terratestRunTestStage terratestFuncCallType = iota
	terratestTRun
)

// terratestFuncCall represents a function call that is significant to terratest. See terratestFuncCallType for the
// specific calls tracked.
type terratestFuncCall struct {
	callType terratestFuncCallType
	expr     *ast.CallExpr
}

// parseTestPackage will take a path to a directory containing a go package with tests that use terratest and return all
// the test functions and their stages. Since the directory may contain multiple go packages (e.g., nested), this
// function also takes in the specific go package for which to extract the terratest tests and stages.
//
// For nested tests, the stages of the nested tests are available in the top level, and vice versa for the relevant
// stages of the top level. See the unit test `TestParserDifferentNestedRunCase` for an example.
func (parser *terratestParser) parseTestPackage(testPackagePath string, testPackageName string) (map[string][]terratestStage, error) {
	pkg, err := getTestPackage(testPackagePath, testPackageName)
	if err != nil {
		return nil, err
	}

	parser.allFuncs = getAllTopLevelFunctions(pkg)
	testFuncs := getTopLevelTestFunctions(parser.allFuncs)

	testFuncToStages := map[string][]terratestStage{}
	for testFuncName, testFunc := range testFuncs {
		terratestFuncCalls, err := getAllTerratestCalls(parser.allFuncs, testFunc)
		if err != nil {
			return nil, err
		}
		allStages, nestedStagesMap, err := parser.getStagesFromTerratestFuncCalls(terratestFuncCalls)
		if err != nil {
			return nil, err
		}
		testFuncToStages[testFuncName] = allStages

		// For nested tests and their stages, make sure to namespace with the current test name, just like it is done
		// with `go test`.
		for nestedTestName, stages := range nestedStagesMap {
			testFuncToStages[fmt.Sprintf("%s/%s", testFuncName, nestedTestName)] = stages
		}
	}
	return testFuncToStages, nil
}

// getStagesFromTerratestFuncCalls takes a list of all the terratest function calls and returns all the terratest
// stages defined by the calls, expanding out nested tests (t.Run calls) in the process. As an optimization, this
// routine will return a map of the nested tests and their stages that can be merged upstream.
func (parser *terratestParser) getStagesFromTerratestFuncCalls(calls []terratestFuncCall) ([]terratestStage, map[string][]terratestStage, error) {
	stages := []terratestStage{}
	nestedStagesMap := map[string][]terratestStage{}
	for _, call := range calls {
		switch call.callType {
		case terratestRunTestStage:
			stageName, err := getStageFromRunTestStageCall(call.expr)
			if err != nil {
				return nil, nil, err
			}
			stage := terratestStage{level: parser.curLevel, name: stageName}

			// We need to append the new stage to both the current stage list and all the nested lists we have so far,
			// so the nested tests get all the deferred stages.
			stages = append(stages, stage)
			for key, val := range nestedStagesMap {
				nestedStagesMap[key] = append(val, stage)
			}
		case terratestTRun:
			nestedStages, nestedNestedStagesMap, err := parser.getStagesFromNestedTest(call)
			if err != nil {
				return nil, nil, err
			}

			// First, append the found nested stages to the current stages list. Note that nestedStages will include all
			// the stages from any additional nested functions, since handleNestedTest recurses with this function so no
			// special processing to expand out nestedNestedStagesMap is necessary when appending to the top level
			// stages list.
			stages = append(stages, nestedStages...)

			// To ensure the nested tests have the full stage list, we need to prepend the list with all the stages from
			// the parent. Since we don't want to include the stages from other siblings of the nested test, we need to
			// find all the stages that belong to the parent of the nested test.
			stagesToInclude := []terratestStage{}
			for _, stage := range stages {
				if stage.level <= parser.curLevel {
					stagesToInclude = append(stagesToInclude, stage)
				}
			}
			for key, val := range nestedNestedStagesMap {
				nestedStagesMap[key] = append(stagesToInclude, val...)
			}
		}
	}
	return stages, nestedStagesMap, nil
}

// getStagesFromNestedTest will extract all the stages from the nested test (t.Run call). This will return a list of all
// the stages corresponding to the nested test, scoped from the nested test (as in, parents are omitted) and a map of
// test names to stages for any additionally nested tests.
func (parser *terratestParser) getStagesFromNestedTest(call terratestFuncCall) ([]terratestStage, map[string][]terratestStage, error) {
	// Since we entered a nested test, manage the parser depth.
	defer func() {
		parser.curLevel--
	}()
	parser.curLevel++

	nestedStages := map[string][]terratestStage{}

	// Get the nested stages, as well as the stages from additional t.Run calls within the subtest, and merge to
	// the top level list.
	nestedTestName, err := getTestNameFromTRunCall(call.expr)
	if err != nil {
		// Log error as warning and continue. We do this so that we still capture all the stages for the parent, but
		// functionality for linking to the nested stages to the nested test names will be broken.
		// TODO: use a proper logger with logging levels
		// TODO: Figure out how to extract position for better debugging.
		fmt.Println("WARNING: Could not extract nested test name from call expression. Nested stages will be extracted, but not tracked under a name.")
	}
	testFunc, err := getTestFuncFromTRunCall(parser.allFuncs, call.expr)
	if err != nil {
		return nil, nil, err
	}
	tRunCalls, err := getAllTerratestCalls(parser.allFuncs, testFunc)
	if err != nil {
		return nil, nil, err
	}
	tRunStages, tRunNestedStages, err := parser.getStagesFromTerratestFuncCalls(tRunCalls)
	if err != nil {
		return nil, nil, err
	}

	if nestedTestName != "" {
		// Make sure to capture the stages for nested test runs, to be merged into the top level. Note that tRunStages will
		// include all the stages for the nested tests as well (included before being returned by
		// getStagesFromTerratestFuncCalls).
		nestedStages[nestedTestName] = tRunStages

		// Merge stages for tests spawned from additional t.Run calls within this subtest.
		for nestedNestedTestName, nestedNestedStages := range tRunNestedStages {
			// NOTE: we don't need to append the tRunStages here because they are already included before being returned by
			// getStagesFromTerratestFuncCalls).
			nestedStages[fmt.Sprintf("%s/%s", nestedTestName, nestedNestedTestName)] = nestedNestedStages
		}
	}

	return tRunStages, nestedStages, nil
}
