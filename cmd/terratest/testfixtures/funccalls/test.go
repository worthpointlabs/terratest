package funccalls

import (
	"testing"

	"github.com/gruntwork-io/terratest/modules/logger"
	test_structure "github.com/gruntwork-io/terratest/modules/test-structure"
)

func TestWithStagesAndOneLevelFuncCall(t *testing.T) {
	t.Parallel()

	setup(t)
	defer cleanup(t)
	deploy(t)
	validate(t)
}

func TestWithStagesAndMultiLevelFuncCall(t *testing.T) {
	t.Parallel()

	setup(t)
	defer cleanup(t)
	nestedNestedDeploy(t)
	validate(t)
}

func TestWithStagesAndNestedRunsWithFuncCall(t *testing.T) {
	t.Parallel()

	setup(t)
	defer cleanup(t)
	deploy(t)
	nestedTestValidate(t)
}

func nestedNestedDeploy(t *testing.T) {
	nestedDeploy(t)
}

func nestedDeploy(t *testing.T) {
	deploy(t)
}

func nestedTestValidate(t *testing.T) {
	t.Run("group", validate)
}

func deploy(t *testing.T) {
	test_structure.RunTestStage(t, "deploy", func() {
		logger.Logf(t, "deploy")
	})
}

func validate(t *testing.T) {
	test_structure.RunTestStage(t, "validate", func() {
		logger.Logf(t, "validate")
	})
}

func setup(t *testing.T) {
	test_structure.RunTestStage(t, "setup", func() {
		logger.Logf(t, "setup")
	})
}

func cleanup(t *testing.T) {
	test_structure.RunTestStage(t, "cleanup", func() {
		logger.Logf(t, "cleanup")
	})
}
