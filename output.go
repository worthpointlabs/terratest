package terratest

import (
	"github.com/gruntwork-io/terratest/terraform"
	"github.com/gruntwork-io/terratest/log"
)

func Output(options *TerratestOptions, key string) (string, error) {
	logger := log.NewLogger(options.TestName)
	return terraform.Output(options.TemplatePath, key, logger)
}
