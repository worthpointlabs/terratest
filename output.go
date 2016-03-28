package terratest

import (
	"github.com/gruntwork-io/terratest/terraform"
	"github.com/gruntwork-io/terratest/log"
)

func Output(ao *ApplyOptions, key string) (string, error) {
	logger := log.NewLogger(ao.TestName)
	return terraform.Output(ao.TemplatePath, key, logger)
}
