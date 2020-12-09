package main

import (
	"fmt"
	"os"
	"path/filepath"

	"github.com/gruntwork-io/gruntwork-cli/entrypoint"
	"github.com/gruntwork-io/gruntwork-cli/errors"
	"github.com/gruntwork-io/gruntwork-cli/logging"
	"github.com/gruntwork-io/terratest/modules/tspec"
	"github.com/gruntwork-io/terratest/modules/tspec/colors"
	"github.com/sirupsen/logrus"
	"github.com/urfave/cli"
)

var logger = logging.GetLogger("tspec")

const CUSTOM_USAGE_TEXT = `Usage: tspec [--help] [--log-level=info] [--outputdir=OUTPUT_DIR] <PATH_TO_FEATURES>

A BDD testing tool that uses Terratest under the hood.

Options:
   --log-level LEVEL  Set the log level to LEVEL. Must be one of: [panic fatal error warning info debug]
                      (default: "info")
   --outputdir value  Path to directory to output test output to. If unset will use the current directory.
   --help, -h         show help
`

var opts = tspec.Options{
	Output:      colors.Colored(os.Stdout),
	Format:      "pretty", // can define default values
	Concurrency: 1,
	Randomize:   0,
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
	opts.Paths = []string{cliContext.Args().Get(0)}

	status := tspec.TestSuite{
		Name:                 "tspec",
		TestSuiteInitializer: tspec.InitializeTestSuite,
		ScenarioInitializer:  tspec.InitializeScenario,
		Options:              &opts,
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
	app.Description = `A BDD testing tool that uses Terratest under the hood.`
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
