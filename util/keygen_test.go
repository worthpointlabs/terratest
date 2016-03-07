package util

import (
	"testing"
)

// Basic test to ensure we can successfully generate keypairs (no explicit validation for now)
func TestGenerateKeypairWorks(t *testing.T) {
	_, _, err := GenerateRSAKeyPair(2048)
	if err != nil {
		t.Error(err)
	}
}
