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

func TestTerraformAzureFunctionAppExample(t *testing.T) {
	t.Parallel()

	_random := strings.ToLower(random.UniqueId())

	expectedKind := "functionapp"
	expectedResourceName := fmt.Sprintf("%s-func", _random)
	expectedResourceGroupName := fmt.Sprintf("%s-rg", _random)

	terraformOptions := &terraform.Options{
		TerraformDir: "../../examples/azure/terraform-azure-functionapp-example",
		Vars: map[string]interface{}{
			"location":     "Central US",
			"project_name": _random,
		},
	}
	// At the end of the test, run `terraform destroy` to clean up any resources that were created
	defer terraform.Destroy(t, terraformOptions)

	// This will run `terraform init` and `terraform apply` and fail the test if there are any errors
	terraform.InitAndApply(t, terraformOptions)

	site := azure.GetFunctionApp(t, expectedResourceName, expectedResourceGroupName, "")

	assert := assert.New(t)

	assert.NotNil(*site)

	assert.Equal(expectedKind, *site.Kind)

	assert.NotEmpty(*site.OutboundIPAddresses)
}
