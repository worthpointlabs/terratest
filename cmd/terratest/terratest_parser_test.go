package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

const (
	baseCase  = "./testfixtures/basecase"
	nestedRun = "./testfixtures/nestedrun"
	funcCall  = "./testfixtures/funccalls"
)

func TestParserBaseCase(t *testing.T) {
	t.Parallel()

	testDir := baseCase
	parser := &terratestParser{}
	testFuncsToStages, err := parser.parseTestPackage(testDir, "basecase")
	require.NoError(t, err)

	stages, hasExpectedTestFunc := testFuncsToStages["TestWithStages"]
	require.True(t, hasExpectedTestFunc)

	assert.Equal(
		t,
		stages,
		[]terratestStage{{0, "setup"}, {0, "deploy"}, {0, "validate"}, {0, "cleanup"}},
	)
}

func TestParserNestedRunCase(t *testing.T) {
	t.Parallel()

	testDir := nestedRun
	parser := &terratestParser{}
	testFuncsToStages, err := parser.parseTestPackage(testDir, "nestedrun")
	require.NoError(t, err)

	stages, hasExpectedTestFunc := testFuncsToStages["TestWithStagesAndNestedTests"]
	require.True(t, hasExpectedTestFunc)
	assert.Equal(
		t,
		stages,
		[]terratestStage{{0, "setup"}, {0, "deploy"}, {1, "validate"}, {0, "cleanup"}},
	)

	nestedStages, hasExpectedNestedTestFunc := testFuncsToStages["TestWithStagesAndNestedTests/group"]
	require.True(t, hasExpectedNestedTestFunc)
	assert.Equal(
		t,
		nestedStages,
		[]terratestStage{{0, "setup"}, {0, "deploy"}, {1, "validate"}, {0, "cleanup"}},
	)
}

func TestParserMultiLayerNestedRunCase(t *testing.T) {
	t.Parallel()

	testDir := nestedRun
	parser := &terratestParser{}
	testFuncsToStages, err := parser.parseTestPackage(testDir, "nestedrun")
	require.NoError(t, err)

	allTestNames := []string{
		"TestWithStagesAndMultiLayerNestedTests",
		"TestWithStagesAndMultiLayerNestedTests/group",
		"TestWithStagesAndMultiLayerNestedTests/group/subtest",
		"TestWithStagesAndMultiLayerNestedTests/group/subtest/subsubtest",
	}

	for _, testName := range allTestNames {
		stages, hasExpectedTestFunc := testFuncsToStages[testName]
		require.True(t, hasExpectedTestFunc)
		assert.Equal(
			t,
			stages,
			[]terratestStage{{0, "setup"}, {0, "deploy"}, {3, "validate"}, {0, "cleanup"}},
		)
	}
}

func TestParserDifferentNestedRunCase(t *testing.T) {
	t.Parallel()

	testDir := nestedRun
	parser := &terratestParser{}
	testFuncsToStages, err := parser.parseTestPackage(testDir, "nestedrun")
	require.NoError(t, err)

	stages, hasExpectedTestFunc := testFuncsToStages["TestWithStagesAndDifferentNestedStages"]
	require.True(t, hasExpectedTestFunc)
	assert.Equal(
		t,
		stages,
		[]terratestStage{{0, "setup"}, {0, "deploy"}, {1, "validate_foo"}, {1, "validate_bar"}, {0, "cleanup"}},
	)

	nestedFooStages, hasExpectedNestedFooTestFunc := testFuncsToStages["TestWithStagesAndDifferentNestedStages/foogroup"]
	require.True(t, hasExpectedNestedFooTestFunc)
	assert.Equal(
		t,
		nestedFooStages,
		[]terratestStage{{0, "setup"}, {0, "deploy"}, {1, "validate_foo"}, {0, "cleanup"}},
	)

	nestedBarStages, hasExpectedNestedBarTestFunc := testFuncsToStages["TestWithStagesAndDifferentNestedStages/bargroup"]
	require.True(t, hasExpectedNestedBarTestFunc)
	assert.Equal(
		t,
		nestedBarStages,
		[]terratestStage{{0, "setup"}, {0, "deploy"}, {1, "validate_bar"}, {0, "cleanup"}},
	)
}

func TestParserOneAndMultiLevelFuncCalls(t *testing.T) {
	t.Parallel()

	testDir := funcCall
	parser := &terratestParser{}
	testFuncsToStages, err := parser.parseTestPackage(testDir, "funccalls")
	require.NoError(t, err)

	for _, tFuncName := range []string{"TestWithStagesAndOneLevelFuncCall", "TestWithStagesAndMultiLevelFuncCall"} {
		stages, hasExpectedTestFunc := testFuncsToStages[tFuncName]
		require.True(t, hasExpectedTestFunc)

		assert.Equal(
			t,
			stages,
			[]terratestStage{{0, "setup"}, {0, "deploy"}, {0, "validate"}, {0, "cleanup"}},
		)
	}
}

func TestParserNestedRunsWithFuncCalls(t *testing.T) {
	t.Parallel()

	testDir := funcCall
	parser := &terratestParser{}
	testFuncsToStages, err := parser.parseTestPackage(testDir, "funccalls")
	require.NoError(t, err)

	stages, hasExpectedTestFunc := testFuncsToStages["TestWithStagesAndNestedRunsWithFuncCall"]
	require.True(t, hasExpectedTestFunc)

	assert.Equal(
		t,
		stages,
		[]terratestStage{{0, "setup"}, {0, "deploy"}, {1, "validate"}, {0, "cleanup"}},
	)

	nestedStages, hasExpectedNestedTestFunc := testFuncsToStages["TestWithStagesAndNestedRunsWithFuncCall/group"]
	require.True(t, hasExpectedNestedTestFunc)

	assert.Equal(
		t,
		nestedStages,
		[]terratestStage{{0, "setup"}, {0, "deploy"}, {1, "validate"}, {0, "cleanup"}},
	)
}
