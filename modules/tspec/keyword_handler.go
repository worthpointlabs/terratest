package tspec

import (
	"fmt"
	"testing"

	"github.com/gruntwork-io/terratest/modules/terraform"
	"github.com/stretchr/testify/assert"
)

// Global variable to reuse between steps
var TerraformModulePath string
var TerraformOptions *terraform.Options

func InitializeTestSuite(ctx *TestSuiteContext) {
	ctx.BeforeSuite(func() {
    // Do nothing for now
  })
}

func InitializeScenario(ctx *ScenarioContext) {
	ctx.Step(`^the Terraform module at "([^"]*)"$`, terraformModulePathHandler)
	ctx.Step(`^an input variable named "([^"]*)" with the value "([^"]*)"$`, terraformInputVarHandler)
	ctx.Step(`^an environment variable named "([^"]*)" with the value "([^"]*)"$`, terraformEnvVarHandler)
	ctx.Step(`^I run "([^"]*)"$`, whenIRunHandler)
	ctx.Step(`^the "([^"]*)" output is "([^"]*)"$`, assertOutputEqualsHandler)
	ctx.Step(`^the "([^"]*)" output should match "([^"]*)"$`, assertOutputMatchesHandler)

	// Initialize new Terraform Options
	// TODO - this may or may not have parallelism issues. We need to test that scenarios don't share the global TerraformOptions.
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
		RetryableTerraformErrors: retryableErrors,

		// Variables to pass to our Terraform code using -var options
		Vars: map[string]interface{}{},

		// Environment variables to set when running Terraform
		EnvVars: map[string]string{},
	}

	// TODO - run Terraform destroy automatically
}

func terraformModulePathHandler(path string) error {
  TerraformOptions.TerraformDir = path
	return nil
}

func terraformInputVarHandler(inputVar string, value string) error {
	TerraformOptions.Vars[inputVar] = value
	return nil
}

func terraformEnvVarHandler(envVar string, value string) error {
	TerraformOptions.EnvVars[envVar] = value
	return nil
}

func whenIRunHandler(cmd string) error {
  // TODO - write logs using the context somehow
  //logger.Infof("Executing command %s", cmd)

	switch cmd {
	case "terraform apply":
		innerT := &testing.T{}
		out, err := terraform.InitAndApplyE(innerT, TerraformOptions)
		if err != nil {
			return fmt.Errorf("There was an error running Terraform apply: %s", out)
		}
		return nil
	default:
		return ErrPending
	}
}

func assertOutputEqualsHandler(outputVar, expected string) error {
	innerT := &testing.T{}
	output := terraform.Output(innerT, TerraformOptions, outputVar)
	return assertExpectedAndActual(
		assert.Equal, expected, output,
		"Expected %s output to be %s", outputVar, expected,
	)
}

func assertOutputMatchesHandler(outputVar, expectedMatchRegex string) error {
	innerT := &testing.T{}
	output := terraform.Output(innerT, TerraformOptions, outputVar)
	return assertExpectedAndActual(
		assert.Regexp, expectedMatchRegex, output,
		"Expected %s output to match the regex %s", outputVar, expectedMatchRegex,
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
