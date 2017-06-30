// Package terratest is a test framework for terraform templates.
//
// terratest exposes a basic interface for creating resources needed for testing, loading a template's terraform variables
// and doing a full-cycle terraform apply and destroy.
//
// It is meant to be used by authors of Terraform templates to systematically validate that their terraform templates
// work as expected across a range of inputs such as a randomly selected AWS Region.
package terratest

import "log"

type TestSuiteBase struct {
	logger             *log.Logger
	resourceCollection *RandomResourceCollection
	terratestOptions   *TerratestOptions
	terraformOutput    string
}
