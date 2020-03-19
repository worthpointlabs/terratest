package nestedrun

import (
	"testing"

	"github.com/gruntwork-io/terratest/modules/logger"
	test_structure "github.com/gruntwork-io/terratest/modules/test-structure"
)

func TestWithStagesAndNestedTests(t *testing.T) {
	t.Parallel()

	test_structure.RunTestStage(t, "setup", func() {
		logger.Logf(t, "setup")
	})

	defer test_structure.RunTestStage(t, "cleanup", func() {
		logger.Logf(t, "cleanup")
	})

	test_structure.RunTestStage(t, "deploy", func() {
		logger.Logf(t, "deploy")
	})

	t.Run("group", func(t *testing.T) {
		test_structure.RunTestStage(t, "validate", func() {
			logger.Logf(t, "validate")
		})
	})
}

func TestWithStagesAndMultiLayerNestedTests(t *testing.T) {
	t.Parallel()

	test_structure.RunTestStage(t, "setup", func() {
		logger.Logf(t, "setup")
	})

	defer test_structure.RunTestStage(t, "cleanup", func() {
		logger.Logf(t, "cleanup")
	})

	test_structure.RunTestStage(t, "deploy", func() {
		logger.Logf(t, "deploy")
	})

	t.Run("group", func(t *testing.T) {
		t.Run("subtest", func(t *testing.T) {
			t.Run("subsubtest", func(t *testing.T) {
				test_structure.RunTestStage(t, "validate", func() {
					logger.Logf(t, "validate")
				})
			})
		})
	})
}

func TestWithStagesAndDifferentNestedStages(t *testing.T) {
	t.Parallel()

	test_structure.RunTestStage(t, "setup", func() {
		logger.Logf(t, "setup")
	})

	defer test_structure.RunTestStage(t, "cleanup", func() {
		logger.Logf(t, "cleanup")
	})

	test_structure.RunTestStage(t, "deploy", func() {
		logger.Logf(t, "deploy")
	})

	t.Run("foogroup", func(t *testing.T) {
		test_structure.RunTestStage(t, "validate_foo", func() {
			logger.Logf(t, "validate_foo")
		})
	})

	t.Run("bargroup", func(t *testing.T) {
		test_structure.RunTestStage(t, "validate_bar", func() {
			logger.Logf(t, "validate_bar")
		})
	})
}
