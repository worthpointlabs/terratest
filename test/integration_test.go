// Integration tests that actually test functions with AWS.  These should be run periodically, but not on every commit.
package test

import (
	"testing"
	"github.com/gruntwork-io/terraform-test/util"
	"github.com/gruntwork-io/terraform-test/aws"
)

func TestUploadKeyPair(t *testing.T) {
	// Create the keypair
	keyPair, err := util.GenerateRSAKeyPair(2048)
	if err != nil {
		t.Errorf("Failed to generate keypair: %s\n", err.Error())
	}

	// Assign randomly generated values
	region := aws.GetRandomRegion()
	id := util.UniqueId()

	// Create key in EC2
	t.Logf("Creating EC2 Keypair %s in %s...", id, region)
	aws.CreateEC2KeyPair(region, id, keyPair.PublicKey)

	// If destroy succeeds, then we assume key was there to destroy in the first place
	t.Logf("Destroying EC2 Keypair %s in %s...", id, region)
	aws.DeleteEC2KeyPair(region, id)
}
