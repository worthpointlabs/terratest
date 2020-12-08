package tspec

import (
	"fmt"

	"github.com/gruntwork-io/terratest/modules/terraform"
	"github.com/stretchr/testify/assert"
)

// Globals to reuse between scenario steps
type terraformFeature struct {
	path             string
	options          *terraform.Options
	terraformIsDirty bool // Whether or not we need to run `terraform destroy` after a scenario
}

func (t *terraformFeature) initialize(*Scenario) {
	retryableErrors := map[string]string{
		"diffs didn't match during apply": "This usually indicates a minor Terraform timing bug (https://github.com/hashicorp/terraform/issues/5200) that goes away when you reapply. Retrying terraform apply.",

		// `terraform init` frequently fails in CI due to network issues accessing plugins. The reason is unknown, but
		// eventually these succeed after a few retries.
		".*unable to verify signature.*":             "Failed to retrieve plugin due to transient network error.",
		".*unable to verify checksum.*":              "Failed to retrieve plugin due to transient network error.",
		".*no provider exists with the given name.*": "Failed to retrieve plugin due to transient network error.",
		".*registry service is unreachable.*":        "Failed to retrieve plugin due to transient network error.",
	}

	t.options = &terraform.Options{
		TerraformDir:             "",
		RetryableTerraformErrors: retryableErrors,

		// Variables to pass to our Terraform code using -var options
		Vars: map[string]interface{}{},

		// Environment variables to set when running Terraform
		EnvVars: map[string]string{},
	}
}

func (t *terraformFeature) GetOptions() *terraform.Options {
	return t.options
}

func (t *terraformFeature) SetDirty(value bool) {
	t.terraformIsDirty = value
}

func (t *terraformFeature) IsDirty() bool {
	return t.terraformIsDirty
}

func (t *terraformFeature) pathHandler(path string) error {
	t.options.TerraformDir = path
	return nil
}

func (t *terraformFeature) inputVarHandler(inputVar string, value string) error {
	t.options.Vars[inputVar] = value
	return nil
}

func (t *terraformFeature) envVarHandler(envVar string, value string) error {
	t.options.EnvVars[envVar] = value
	return nil
}

func (t *terraformFeature) whenIRunHandler(cmd string) error {
	switch cmd {
	case "terraform apply":
		out, err := terraform.InitAndApplyE(GetT(), t.options)
		if err != nil {
			return fmt.Errorf("There was an error running Terraform apply: %s", out)
		}
		t.SetDirty(true)
		return nil
	default:
		return ErrPending
	}
}

func (t *terraformFeature) assertOutputEqualsHandler(outputVar, expected string) error {
	output := terraform.Output(GetT(), t.options, outputVar)
	return assertExpectedAndActual(
		assert.Equal, expected, output,
		"Expected %s output to be %s", outputVar, expected,
	)
}

func (t *terraformFeature) assertOutputMatchesHandler(outputVar, expectedMatchRegex string) error {
	output := terraform.Output(GetT(), t.options, outputVar)
	return assertExpectedAndActual(
		assert.Regexp, expectedMatchRegex, output,
		"Expected %s output to match the regex %s", outputVar, expectedMatchRegex,
	)
}

func InitializeTestSuite(ctx *TestSuiteContext) {
	ctx.BeforeSuite(func() {
		// Do nothing for now
	})
}

func InitializeScenario(ctx *ScenarioContext) {
	// Initialize new Terraform Options
	// TODO - this may or may not have parallelism issues. We need to test that scenarios don't share the global TerraformOptions.
	terraformFeature := &terraformFeature{}
	ctx.BeforeScenario(terraformFeature.initialize)

	// Possibly Run Terraform Destroy
	ctx.AfterScenario(func(sc *Scenario, err error) {
		if terraformFeature.IsDirty() {
			terraform.DestroyE(GetT(), terraformFeature.GetOptions())
			terraformFeature.SetDirty(true)
		}
	})

	// Define Steps
	ctx.Step(`^the Terraform module at "([^"]*)"$`, terraformFeature.pathHandler)
	ctx.Step(`^an input variable named "([^"]*)" with the value "([^"]*)"$`, terraformFeature.inputVarHandler)
	ctx.Step(`^an environment variable named "([^"]*)" with the value "([^"]*)"$`, terraformFeature.envVarHandler)
	ctx.Step(`^I run "([^"]*)"$`, terraformFeature.whenIRunHandler)
	ctx.Step(`^the "([^"]*)" output is "([^"]*)"$`, terraformFeature.assertOutputEqualsHandler)
	ctx.Step(`^the "([^"]*)" output should match "([^"]*)"$`, terraformFeature.assertOutputMatchesHandler)
}
