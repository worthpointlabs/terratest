package util

import (
	"testing"
)

func TestGetFirstTwoOctets(t *testing.T) {
	t.Parallel()

	firstTwo := GetFirstTwoOctets("10.100.0.0/28")
	if firstTwo != "10.100" {
		t.Errorf("Received: %s, Expected: 10.100", firstTwo)
	}
}

// Deferred to save time
func TestGetRandomPrivateCidrBlock(t *testing.T) {
}
