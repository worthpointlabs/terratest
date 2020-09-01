// +build azure

// NOTE: We use build tags to differentiate azure testing because we currently do not have azure access setup for
// CircleCI.

package test

import (
	"fmt"
	"strings"

	"testing"

	"github.com/gruntwork-io/terratest/modules/azure"
	"github.com/gruntwork-io/terratest/modules/random"
	"github.com/gruntwork-io/terratest/modules/terraform"
	"github.com/stretchr/testify/assert"
)

func TestTerraformAzureACRExample(t *testing.T) {
	t.Parallel()

	_random := strings.ToLower(random.UniqueId())

	expectedResourceName := fmt.Sprintf("tmpacr%s", _random)
	expectedResourceGroupName := fmt.Sprintf("tmp-rg-%s", _random)

	terraformOptions := &terraform.Options{
		TerraformDir: "../../examples/terraform-azure-acr-example",
		Vars: map[string]interface{}{
			"acr_name":            expectedResourceName,
			"resource_group_name": expectedResourceGroupName,
		},
	}
	// At the end of the test, run `terraform destroy` to clean up any resources that were created
	defer terraform.Destroy(t, terraformOptions)

	// This will run `terraform init` and `terraform apply` and fail the test if there are any errors
	terraform.InitAndApply(t, terraformOptions)

	// TODO

	client := azure.GetACRClient(t, expectedResourceName, expectedResourceGroupName, "")

	assert := assert.New(t)

	assert.NotEmpty(*client.Name)

	assert.Equal(fmt.Sprintf("%s.azurecr.io", *client.Name), client.LoginServer)
}
