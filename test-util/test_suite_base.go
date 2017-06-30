package test_util

import (
	"github.com/gruntwork-io/terratest"
	"log"
)

type TestSuiteBase struct {
	logger             *log.Logger
	resourceCollection *terratest.RandomResourceCollection
	terratestOptions   *terratest.TerratestOptions
	terraformOutput    string
}

