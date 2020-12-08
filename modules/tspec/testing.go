package tspec

import (
	"github.com/gruntwork-io/terratest/modules/testing"
)

func GetT() testing.TestingT {
	return &tspecTestingT{}
}

type tspecTestingT struct {
	testing.TestingT
}

// Extends tspecTestingT to have #Name() method, that is compatible with testing.TestingT
func (t *tspecTestingT) Name() string {
	return "[tspec]"
}
