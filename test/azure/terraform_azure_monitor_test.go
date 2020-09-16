// +build azure

// NOTE: We use build tags to differentiate azure testing because we currently do not have azure access setup for
// CircleCI.

package test

import (
	"testing"

	"github.com/gruntwork-io/terratest/modules/azure"
	"github.com/gruntwork-io/terratest/modules/terraform"
	"github.com/stretchr/testify/assert"
)

func TestTerraformAzureMonitorExample(t *testing.T) {
	t.Parallel()

	subscriptionID := ""
	prefix := "terratest-lb"

	// website::tag::1:: Configure Terraform setting up a path to Terraform code.
	terraformOptions := &terraform.Options{
		// The path to where our Terraform code is located
		TerraformDir: "../../examples/azure/terraform-azure-loadbalancer-example",
		Vars: map[string]interface{}{
			"prefix": prefix,
		},
	}

	// website::tag::4:: At the end of the test, run `terraform destroy` to clean up any resources that were created
	defer terraform.Destroy(t, terraformOptions)

	// website::tag::2:: Run `terraform init` and `terraform apply`. Fail the test if there are any errors.
	terraform.InitAndApply(t, terraformOptions)

	resourceGroupName := terraform.Output(t, terraformOptions, "resource_group_name")
	loadBalancerName := terraform.Output(t, terraformOptions, "lb_name")

	lbExists := azure.LoadBalancerExists(t, loadBalancerName, resourceGroupName, subscriptionID)

	assert.Equal(t, lbExists, true, "Load balancer should exist")

	// // website::tag::3:: Run `terraform output` to get the values of output variables
	// vmName := terraform.Output(t, terraformOptions, "vm_name")
	// resourceGroupName := terraform.Output(t, terraformOptions, "resource_group_name")

	// // website::tag::4:: Look up the size of the given Virtual Machine and ensure it matches the output.
	// actualVMSize := azure.GetSizeOfVirtualMachine(t, vmName, resourceGroupName, "")
	// expectedVMSize := compute.VirtualMachineSizeTypes("Standard_B1s")
	// assert.Equal(t, expectedVMSize, actualVMSize)
}
