package main

import (
	"fmt"
	"os"
	"path/filepath"
	"testing"

	"github.com/gruntwork-io/gruntwork-cli/entrypoint"
	"github.com/gruntwork-io/gruntwork-cli/errors"
	"github.com/gruntwork-io/gruntwork-cli/logging"
	"github.com/gruntwork-io/terratest/modules/terraform"
	"github.com/gruntwork-io/terratest/modules/tspec"
	"github.com/gruntwork-io/terratest/modules/tspec/colors"
	"github.com/stretchr/testify/assert"
	"github.com/sirupsen/logrus"
	"github.com/urfave/cli"
)

var logger = logging.GetLogger("tspec")

const CUSTOM_USAGE_TEXT = `Usage: tspec [--help] [--log-level=info] [--outputdir=OUTPUT_DIR] <PATH_TO_FEATURES>

A tool for parsing parallel terratest output to produce a test summary and to break out the interleaved logs by test for better debuggability.

Options:
   --log-level LEVEL  Set the log level to LEVEL. Must be one of: [panic fatal error warning info debug]
                      (default: "info")
   --outputdir value  Path to directory to output test output to. If unset will use the current directory.
   --help, -h         show help
`

var opts = tspec.Options{
	Output: colors.Colored(os.Stdout),
	Format: "progress", // can define default values
}

func run(cliContext *cli.Context) error {
	outputDir := cliContext.String("outputdir")
	logLevel := cliContext.String("log-level")
	level, err := logrus.ParseLevel(logLevel)
	if err != nil {
		return errors.WithStackTrace(err)
	}
	logger.SetLevel(level)

	outputDir, err = filepath.Abs(outputDir)
	if err != nil {
		logger.Fatalf("Error extracting absolute path of output directory: %s", err)
	}

	// parse args as features path
	cliContext.Args()

	status := tspec.TestSuite{
		Name: "tspec",
		TestSuiteInitializer: InitializeTestSuite,
		ScenarioInitializer:  InitializeScenario,
		Options: &opts,
	}.Run()

	logger.Infof("Exit code is: %d", status)

	return nil
}

func main() {
	app := entrypoint.NewApp()
	cli.AppHelpTemplate = CUSTOM_USAGE_TEXT
	entrypoint.HelpTextLineWidth = 120

	app.Name = "tspec"
	app.Author = "Gruntwork <www.gruntwork.io>"
	app.Description = `A tool for parsing parallel terratest output to produce a test summary and to break out the interleaved logs by test for better debuggability.`
	app.Action = run

	currentDir, err := os.Getwd()
	if err != nil {
		logger.Fatalf("Error finding current directory: %s", err)
	}
	defaultOutputDir := filepath.Join(currentDir, "out")

	outputDirFlag := cli.StringFlag{
		Name:  "outputdir, o",
		Value: defaultOutputDir,
		Usage: "Path to directory to output test output to. If unset will use the current directory.",
	}
	logLevelFlag := cli.StringFlag{
		Name:  "log-level",
		Value: logrus.InfoLevel.String(),
		Usage: fmt.Sprintf("Set the log level to `LEVEL`. Must be one of: %v", logrus.AllLevels),
	}
	app.Flags = []cli.Flag{
		logLevelFlag,
		outputDirFlag,
	}

	entrypoint.RunApp(app)
}

// Global variable to reuse between steps
var TerraformModulePath string

func InitializeTestSuite(ctx *tspec.TestSuiteContext) {
	ctx.BeforeSuite(func() { TerraformModulePath = "" })
}

func InitializeScenario(ctx *tspec.ScenarioContext) {
	ctx.Step(`^the Terraform module at "([^"]*)"$`, theTerraformModuleAt)
	ctx.Step(`^I run "([^"]*)"$`, iRun)
	ctx.Step(`^the "([^"]*)" output is "([^"]*)"$`, checkOutputIs)
	// TODO - run Terraform destroy automatically
}

func theTerraformModuleAt(path string) error {
	TerraformModulePath = path
	return nil
}

func iRun(cmd string) error {
	fmt.Println(fmt.Sprintf("module path is: %s", cmd))
	switch cmd {
	case "terraform apply":
		innerT := &testing.T{}
		options := createBaseTerratestOptions(TerraformModulePath, "us-east-1")
		terraform.InitAndApply(innerT, options)
		return nil
	default:
		return tspec.ErrPending
	}
}

func checkOutputIs(outputVar, expected string) error {
	innerT := &testing.T{}
	options := createBaseTerratestOptions(TerraformModulePath, "us-east-1")
	output := terraform.Output(innerT, options, outputVar)
	return tspec.AssertExpectedAndActual(
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
