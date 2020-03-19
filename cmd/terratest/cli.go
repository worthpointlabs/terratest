package main

import (
	"os"

	"github.com/gruntwork-io/gruntwork-cli/entrypoint"
	"github.com/gruntwork-io/gruntwork-cli/errors"
	"github.com/gruntwork-io/gruntwork-cli/logging"
	"github.com/sirupsen/logrus"
	"github.com/urfave/cli"
)

const customUsageText = `Usage: terratest [--help] [--log-level=info] [--path=TEST_PATH] [--package=TEST_PACKAGE_NAME]

A CLI frontend for driving common terratest workflows.

Options:
   --log-level LEVEL  Set the log level to LEVEL. Must be one of: [panic fatal error warning info debug]
                      (default: "info")
   --path value       Path to directory containing go test package that uses terratest. (default: current directory)
   --package value    Name of go test package to collect terratest Test functions from. (default: test)
   --help, -h         show help
`

var (
	logLevelFlag = cli.StringFlag{
		Name:  "loglevel",
		Value: logrus.InfoLevel.String(),
	}

	testPackagePathFlag = cli.StringFlag{
		Name:  "path",
		Value: ".",
	}
	testPackageNameFlag = cli.StringFlag{
		Name:  "package",
		Value: "test",
	}
)

// initCli initializes the CLI app before any command is actually executed. This function will handle all the setup
// code, such as setting up the logger with the appropriate log level.
func initCli(cliContext *cli.Context) error {
	// Set logging level
	logLevel := cliContext.String(logLevelFlag.Name)
	level, err := logrus.ParseLevel(logLevel)
	if err != nil {
		return errors.WithStackTrace(err)
	}
	logging.SetGlobalLogLevel(level)

	// If logging level is for debugging (debug or trace), enable stacktrace debugging
	if level == logrus.DebugLevel || level == logrus.TraceLevel {
		os.Setenv("GRUNTWORK_DEBUG", "true")
	}
	return nil
}

func cliAction(ctx *cli.Context) error {
	projectLogger := getProjectLogger()

	testPackagePath := ctx.String(testPackagePathFlag.Name)
	testPackageName := ctx.String(testPackageNameFlag.Name)

	parser := terratestParser{logger: projectLogger}
	testStagesMap, err := parser.parseTestPackage(testPackagePath, testPackageName)
	if err != nil {
		return err
	}

	return runTestInteractively(projectLogger, testPackagePath, testStagesMap)
}

func terratestCli(version string) *cli.App {
	app := entrypoint.NewApp()
	cli.AppHelpTemplate = customUsageText

	app.Name = "terratest"
	app.Author = "Gruntwork <www.gruntwork.io>"
	app.Description = "A CLI frontend for driving common terratest workflows."
	app.EnableBashCompletion = true
	// Set the version number from your app from the VERSION variable that is passed in at build time
	app.Version = version

	app.Before = initCli
	app.Action = cliAction

	app.Flags = []cli.Flag{
		logLevelFlag,
		testPackagePathFlag,
		testPackageNameFlag,
	}
	return app
}
