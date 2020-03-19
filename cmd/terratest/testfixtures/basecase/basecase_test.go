package basecase

import (
	"testing"

	"github.com/gruntwork-io/terratest/modules/logger"
	test_structure "github.com/gruntwork-io/terratest/modules/test-structure"
)

func TestWithStages(t *testing.T) {
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

	test_structure.RunTestStage(t, "validate", func() {
		logger.Logf(t, "validate")
	})
}
