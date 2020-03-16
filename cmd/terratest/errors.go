package main

import (
	"fmt"
)

// NoPackageError is returned when there is no go package with the given name in the directory path.
type NoPackageError struct {
	path string
	name string
}

func (err NoPackageError) Error() string {
	return fmt.Sprintf("%s does not have go package %s", err.path, err.name)
}

// NoStageNameError is returned when no name for the test stage can be extracted.
type NoStageNameError struct{}

func (err NoStageNameError) Error() string {
	return "Could not find test stage name from RunTestStage call"
}

// NoTestNameError is returned when no name for the test can be extracted.
type NoTestNameError struct{}

func (err NoTestNameError) Error() string {
	return "Could not find test name from t.Run call"
}

// NoTestBodyError is returned when no function body for the test can be extracted.
type NoTestBodyError struct{}

func (err NoTestBodyError) Error() string {
	return "Could not find test function body from t.Run call"
}
