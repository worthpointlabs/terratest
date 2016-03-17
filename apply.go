package terratest

import (
	"fmt"

	"github.com/gruntwork-io/terratest/log"
	"github.com/gruntwork-io/terratest/aws"
	"github.com/gruntwork-io/terratest/terraform"
)

// Apply handles all setup required for a Terraform Apply operation but does not perform a destroy operation or do any cleanup.
// The caller of this function is expected to call Destroy to clean up the Terraform template when done.
func Apply(ao *ApplyOptions) (string, *RandomResourceCollection, *ApplyOptions, error) {
	logger := log.NewLogger(ao.TestName)
	var output string

	// SETUP

	// Generate random values to allow tests to run in parallel
	logger.Println("Generating random values needed for terraform apply...")
	var rand *RandomResourceCollection

	defer rand.DestroyResources()
	rand, err := CreateRandomResourceCollection()
	if err != nil {
		return output, rand, ao, fmt.Errorf("Failed to create random resource collection: %s", err.Error())
	}

	// Configure terraform to use Remote State.
	err = aws.AssertS3BucketExists(ao.TfRemoteStateS3BucketRegion, ao.TfRemoteStateS3BucketName)
	if err != nil {
		return output, rand, ao, fmt.Errorf("Test failed because the S3 Bucket '%s' does not exist in the '%s' region.\n", ao.TfRemoteStateS3BucketName, ao.TfRemoteStateS3BucketRegion)
	}

	terraform.ConfigureRemoteState(ao.TemplatePath, ao.TfRemoteStateS3BucketName, ao.generateTfStateFileName(rand), ao.TfRemoteStateS3BucketRegion, logger)

	// TERRAFORM APPLY
	// Download all the Terraform modules
	logger.Println("Running terraform get...")
	err = terraform.Get(ao.TemplatePath, logger)
	if err != nil {
		return output, rand, ao, fmt.Errorf("Failed to call terraform get successfully: %s\n", err.Error())
	}

	// Apply the Terraform template
	logger.Println("Running terraform apply...")
	if ao.AttemptTerraformRetry {
		output, err = terraform.ApplyAndGetOutputWithRetry(ao.TemplatePath, ao.Vars, ao.RetryableTerraformErrors, logger)
	} else {
		output, err = terraform.ApplyAndGetOutput(ao.TemplatePath, ao.Vars, logger)
	}
	if err != nil {
		return output, rand, ao, fmt.Errorf("Failed to terraform apply: %s\n", err.Error())
	}

	return output, rand, ao, nil
}