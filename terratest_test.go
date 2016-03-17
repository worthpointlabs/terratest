// Integration tests that test cross-package functionality in AWS.
package terratest

import (
	"path"
	"testing"

	"github.com/gruntwork-io/terratest/aws"
	"github.com/gruntwork-io/terratest/util"
)

// This is the directory where our test fixtures are.
const fixtureDir = "./test-fixtures"

func TestUploadKeyPair(t *testing.T) {
	t.Parallel()

	// Assign randomly generated values
	region := aws.GetRandomRegion()
	id := util.UniqueId()

	// Create the keypair
	keyPair, err := util.GenerateRSAKeyPair(2048)
	if err != nil {
		t.Errorf("Failed to generate keypair: %s\n", err.Error())
	}

	// Create key in EC2
	t.Logf("Creating EC2 Keypair %s in %s...", id, region)
	defer aws.DeleteEC2KeyPair(region, id)
	aws.CreateEC2KeyPair(region, id, keyPair.PublicKey)
}

func TestTerraformApplyOnMinimalExample(t *testing.T) {
	t.Parallel()

	rand, err := CreateRandomResourceCollection()
	defer rand.DestroyResources()
	if err != nil {
		t.Errorf("Failed to create random resource collection: %s\n", err.Error())
	}

	vars := make(map[string]string)
	vars["aws_region"] = rand.AwsRegion
	vars["ec2_key_name"] = rand.KeyPair.Name
	vars["ec2_instance_name"] = rand.UniqueId
	vars["ec2_image"] = rand.AmiId

	ao := NewApplyOptions()
	ao.TestName = "Test - TestTerraformApplyMainFunction"
	ao.TemplatePath = path.Join(fixtureDir, "minimal-example")
	ao.Vars = vars
	ao.AttemptTerraformRetry = false

	ApplyAndDestroy(ao)
}

func TestTerraformApplyOnMinimalExampleWithRetry(t *testing.T) {
	t.Parallel()

	rand, err := CreateRandomResourceCollection()
	defer rand.DestroyResources()
	if err != nil {
		t.Errorf("Failed to create random resource collection: %s\n", err.Error())
	}

	vars := make(map[string]string)
	vars["aws_region"] = rand.AwsRegion
	vars["ec2_key_name"] = rand.KeyPair.Name
	vars["ec2_instance_name"] = rand.UniqueId
	vars["ec2_image"] = rand.AmiId

	ao := NewApplyOptions()
	ao.TestName = "Test - TestTerraformApplyMainFunction"
	ao.TemplatePath = path.Join(fixtureDir, "minimal-example")
	ao.Vars = vars
	ao.AttemptTerraformRetry = true

	ApplyAndDestroy(ao)
}
