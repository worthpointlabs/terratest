// Package terratest is a test framework for terraform templates.
//
// terratest exposes a basic interface for creating resources needed for testing, loading a template's terraform variables
// and doing a full-cycle terraform apply and destroy.
//
// It is meant to be used by authors of Terraform templates to systematically validate that their terraform templates
// work as expected across a range of inputs such as a randomly selected AWS Region.
package terratest

import (
	"log"
	"testing"
)

// This is intended to be used as part of setting up a suite of tests for a particular module. This is particularly
// useful when you want to run apply/destroy once, but want to run many assert statements to verify state is as expected.
//
// Example:
//
//  func TestFooSuite(t *testing.T) {
//      testSuite := TestSuiteBase{}
//
//      testSuite.logger = terralog.NewLogger("TestFooSuiteSuite")
//      testSuite.resourceCollection = resources.CreateBaseRandomResourceCollection(t, "")
//
//      defer tearDownSuite(t, &testSuite)
//
//      _, err = terratest.Apply(testSuite.terratestOptions)
//      if err != nil {
//          t.Fatalf("Unexpected error when applying terraform templates: %v", err)
//      }
//
//      t.Run("foo tests", func(t *testing.T) {
//          t.Run("fooTest1", WrapTestCase(testOne, &testSuite))
//          t.Run("fooTest2", WrapTestCase(testTwo, &testSuite))
//          t.Run("fooTest3", WrapTestCase(testThree, &testSuite))
//       })
//  }
//
type TestSuiteBase struct {
	Logger             *log.Logger
	ResourceCollection *RandomResourceCollection
	TerratestOptions   *TerratestOptions
	TerraformOutput    string
}

func WrapTestCase(testCase func(t *testing.T, testSuite *TestSuiteBase), testSuite *TestSuiteBase) func(t *testing.T) {
	return func(t *testing.T) {
		testCase(t, testSuite)
	}
}