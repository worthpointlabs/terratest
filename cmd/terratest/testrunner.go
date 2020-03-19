package main

import (
	"fmt"
	"sort"

	"github.com/AlecAivazis/survey/v2"
	"github.com/gruntwork-io/gruntwork-cli/collections"
	"github.com/gruntwork-io/gruntwork-cli/errors"
	"github.com/gruntwork-io/gruntwork-cli/shell"
	"github.com/sirupsen/logrus"
)

type wizardState int

const (
	stateChooseTest wizardState = iota
	stateRunTest
)

func runTestInteractively(
	logger *logrus.Entry,
	testPackagePath string,
	testPackageName string,
) error {
	parser := terratestParser{logger: logger}
	curState := stateChooseTest
	curTest := ""
	var curTestStagesToRun []string = nil

	for {
		// Parse the test package on each loop so we always have the latest test stages list
		testStagesMap, err := parser.parseTestPackage(testPackagePath, testPackageName)
		if err != nil {
			return err
		}

		switch curState {
		case stateChooseTest:
			// Handle state
			testNames := getTestNamesFromStagesMap(testStagesMap)
			testToRun, err := chooseTestToRun(testNames)
			if err != nil {
				return err
			}

			// Update state
			curTest = testToRun
			curTestStagesToRun = nil // reset so that we initially attempt to run all stages
			curState = stateRunTest

		case stateRunTest:
			// Validation checks: make sure the updated stage map still has the current test we are running
			allStages, hasTest := testStagesMap[curTest]
			if !hasTest {
				logger.Warnf("Test %s no longer exists in package.", curTest)
				curState = stateChooseTest
				// TODO: figure out how not to use continue
				continue
			}
			allStageNames := getStageNamesFromStageSlice(allStages)

			if curTestStagesToRun == nil {
				curTestStagesToRun = allStageNames
			}

			stagesToRun, stagesToSkip, err := chooseStagesToRun(allStageNames, curTestStagesToRun)
			if err != nil {
				return err
			}

			// Update state
			curTestStagesToRun = stagesToRun

			logger.Infof("Selected to run test %s with stages:", curTest)
			for _, stage := range curTestStagesToRun {
				logger.Infof("\t- %s", stage)
			}
			logger.Info("Running test")
			if err := runTest(testPackagePath, curTest, stagesToSkip); err != nil {
				// NOTE: We don't return the error here, as test failures will lead to an error and we don't want to
				// exit the interactive runner everytime the test fails.
				logger.Errorf("Error running test %s: %s", curTest, err)
			}

			// Prompt if we should continue running the current test, and if not, return to choosing which test to run
			shouldContinue, err := askToContinueWithCurrentTest(curTest)
			if err != nil {
				return err
			}
			if !shouldContinue {
				curState = stateChooseTest
			}
		}
	}
}

func runTest(testPackagePath string, testName string, testStagesToSkip []string) error {
	env := map[string]string{}
	for _, stage := range testStagesToSkip {
		env[fmt.Sprintf("SKIP_%s", stage)] = "true"
	}

	options := shell.NewShellOptions()
	options.Env = env
	testRegex := fmt.Sprintf("^%s$", testName)
	return shell.RunShellCommand(
		options, "go", "test", "-count", "1", "-v", "-run", testRegex, "-timeout", "2h", testPackagePath)
}

func getTestNamesFromStagesMap(testStagesMap map[string][]terratestStage) []string {
	testNames := []string{}
	for testName, _ := range testStagesMap {
		testNames = append(testNames, testName)
	}
	sort.Strings(testNames)
	return testNames
}

func getStageNamesFromStageSlice(testStages []terratestStage) []string {
	stageNames := []string{}
	for _, stage := range testStages {
		stageNames = append(stageNames, stage.name)
	}
	return stageNames
}

func chooseTestToRun(testNames []string) (string, error) {
	var testToRun string
	prompt := &survey.Select{
		Message: "Choose test to run:",
		Options: testNames,
	}
	err := survey.AskOne(prompt, &testToRun, survey.WithValidator(survey.Required))
	return testToRun, errors.WithStackTrace(err)
}

func chooseStagesToRun(testStages []string, alreadyChecked []string) ([]string, []string, error) {
	stagesToRun := []string{}
	prompt := &survey.MultiSelect{
		Message: "Choose test stages to run:",
		Default: alreadyChecked,
		Options: testStages,
	}
	err := survey.AskOne(prompt, &stagesToRun)
	if err != nil {
		return nil, nil, errors.WithStackTrace(err)
	}

	stagesToSkip := []string{}
	for _, stage := range testStages {
		if !collections.ListContainsElement(stagesToRun, stage) {
			stagesToSkip = append(stagesToSkip, stage)
		}
	}
	return stagesToRun, stagesToSkip, nil
}

func askToContinueWithCurrentTest(testName string) (bool, error) {
	shouldContinue := false
	prompt := &survey.Confirm{
		Message: fmt.Sprintf("Continue running %s?", testName),
	}
	err := survey.AskOne(prompt, &shouldContinue)
	return shouldContinue, errors.WithStackTrace(err)
}
