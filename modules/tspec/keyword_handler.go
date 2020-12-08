package tspec

import (
	"testing"

	"github.com/gruntwork-io/terratest/modules/terraform"
	"github.com/stretchr/testify/assert"
)

// Global variable to reuse between steps
var TerraformModulePath string
var TerraformOptions *terraform.Options

func InitializeTestSuite(ctx *TestSuiteContext) {
	ctx.BeforeSuite(func() {
    // Initialize new Terraform Options
    retryableErrors := map[string]string{
  		"diffs didn't match during apply": "This usually indicates a minor Terraform timing bug (https://github.com/hashicorp/terraform/issues/5200) that goes away when you reapply. Retrying terraform apply.",

  		// `terraform init` frequently fails in CI due to network issues accessing plugins. The reason is unknown, but
  		// eventually these succeed after a few retries.
  		".*unable to verify signature.*":             "Failed to retrieve plugin due to transient network error.",
  		".*unable to verify checksum.*":              "Failed to retrieve plugin due to transient network error.",
  		".*no provider exists with the given name.*": "Failed to retrieve plugin due to transient network error.",
  		".*registry service is unreachable.*":        "Failed to retrieve plugin due to transient network error.",
  	}

    TerraformOptions = &terraform.Options{
      TerraformDir:             "",
  		Vars:                     nil,
  		RetryableTerraformErrors: retryableErrors,
    }
  })
}

func InitializeScenario(ctx *ScenarioContext) {
	ctx.Step(`^the Terraform module at "([^"]*)"$`, terraformModulePathHandler)
	ctx.Step(`^I run "([^"]*)"$`, whenIRunHandler)
	ctx.Step(`^the "([^"]*)" output is "([^"]*)"$`, assertOutputIsHandler)
	// TODO - run Terraform destroy automatically
}

func terraformModulePathHandler(path string) error {
  TerraformOptions.TerraformDir = path
	return nil
}

func whenIRunHandler(cmd string) error {
  // TODO - write logs using the context somehow
  //logger.Infof("Executing command %s", cmd)

	switch cmd {
	case "terraform apply":
		innerT := &testing.T{}
    // TODO - Pick a Random Region
    terraformVars := map[string]interface{}{

		}

    TerraformOptions.Vars = terraformVars
		terraform.InitAndApply(innerT, TerraformOptions)
		return nil
	default:
		return ErrPending
	}
}

func assertOutputIsHandler(outputVar, expected string) error {
	innerT := &testing.T{}
	output := terraform.Output(innerT, TerraformOptions, outputVar)
	return assertExpectedAndActual(
		assert.Equal, expected, output,
		"Expected %s output to be %s", outputVar, expected,
	)
}

func createBaseTerratestOptions(templatePath string, awsRegion string) *terraform.Options {
	/*
		terraformVars := map[string]interface{}{
			"aws_region": awsRegion,
			//"name":       uniqueID,
		}
	*/

	retryableErrors := map[string]string{
		"diffs didn't match during apply": "This usually indicates a minor Terraform timing bug (https://github.com/hashicorp/terraform/issues/5200) that goes away when you reapply. Retrying terraform apply.",

		// `terraform init` frequently fails in CI due to network issues accessing plugins. The reason is unknown, but
		// eventually these succeed after a few retries.
		".*unable to verify signature.*":             "Failed to retrieve plugin due to transient network error.",
		".*unable to verify checksum.*":              "Failed to retrieve plugin due to transient network error.",
		".*no provider exists with the given name.*": "Failed to retrieve plugin due to transient network error.",
		".*registry service is unreachable.*":        "Failed to retrieve plugin due to transient network error.",
	}

	terratestOptions := terraform.Options{
		TerraformDir:             templatePath,
		Vars:                     nil,
		RetryableTerraformErrors: retryableErrors,
	}
	return &terratestOptions
}
