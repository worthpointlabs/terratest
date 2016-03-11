// Functions defined in main.go represent the "interface" for this test tool.
// They generally handle all setup, test execution, and teardown.
// Consumers of this library should only use other functions if you need to deviate from the setup/teardown in these "main" functions.
package main

import (
	"fmt"

	"github.com/gruntwork-io/terraform-test/log"
	"github.com/gruntwork-io/terraform-test/aws"
	"github.com/gruntwork-io/terraform-test/util"
	"github.com/gruntwork-io/terraform-test/terraform"
)

func main() {
}

// This function wraps all setup and teardown required for a Terraform Apply operation. It returns the output of the terraform operations.
func TerraformApply(testName string, templatePath string, vars map[string]string, attemptTerraformRetry bool) (string, error) {
	logger := log.NewLogger(testName)
	var output string

	// SETUP

	// Generate random values to allow tests to run in parallel
	// - Note that if two tests run at the same time we don't expect a conflict, but by randomly selecting a region
	//   we can reduce the likelihood of hitting standard AWS limits.
	// - The "id" is used to namespace all terraform resource names. In fact, Terraform templates should be written
	//   so that all resources that terraform creates have namespaced names to enable parallel test runs of the same test.
	region := aws.GetRandomRegion()
	id := util.UniqueId()
	logger.Printf("Random values selected. Region = %s, Id = %s\n", region, id)

	// Generate a random RSA Keypair and upload it to AWS to create a new EC2 Keypair.
	keyPair, err := util.GenerateRSAKeyPair(2048)
	if err != nil {
		return output, fmt.Errorf("Failed to generate keypair: %s\n", err.Error())
	}
	logger.Println("Generated keypair. Printing out private key...")
	logger.Printf("%s", keyPair.PrivateKey)

	logger.Println("Creating EC2 KeyPair...")
	defer aws.DeleteEC2KeyPair(region, id)
	err = aws.CreateEC2KeyPair(region, id, keyPair.PublicKey)
	if err != nil {
		return output, fmt.Errorf("Failed to create EC2 KeyPair: %s\n", err.Error())
	}

	// Configure terraform to use Remote State.
	err = aws.AssertS3BucketExists(TF_REMOTE_STATE_S3_BUCKET_REGION, TF_REMOTE_STATE_S3_BUCKET_NAME)
	if err != nil {
		return output, fmt.Errorf("Test failed because the S3 Bucket '%s' does not exist in the '%s' region.\n", TF_REMOTE_STATE_S3_BUCKET_NAME, TF_REMOTE_STATE_S3_BUCKET_REGION)
	}

	terraform.ConfigureRemoteState(templatePath, TF_REMOTE_STATE_S3_BUCKET_NAME, id + "/terraform.tfstate", TF_REMOTE_STATE_S3_BUCKET_REGION, logger)

	// TEST

	// Download all the Terraform modules
	logger.Println("Running terraform get...")

	err = terraform.Get(templatePath, logger)
	if err != nil {
		return output, fmt.Errorf("Failed to call terraform get successfully: %s\n", err.Error())
	}

	// Apply the Terraform template
	logger.Println("Running terraform apply...")

	defer TerraformDestroyHelper(testName, templatePath, vars, TF_REMOTE_STATE_S3_BUCKET_NAME, id + "/terraform.tfstate")
	if attemptTerraformRetry {
		output, err = terraform.ApplyAndGetOutputWithRetry(templatePath, vars, logger)
	} else {
		output, err = terraform.ApplyAndGetOutput(templatePath, vars, logger)
	}
	if err != nil {
		return output, fmt.Errorf("Failed to terraform apply: %s\n", err.Error())
	}

	return output, nil
}

// Helper function that allows Terraform Destroy to be called after Terraform Apply returns
func TerraformDestroyHelper(testName string, templatePath string, vars map[string]string, remoteStateS3BucketName string, remoteStateS3ObjectName string) {
	logger := log.NewLogger(testName)
	err := terraform.Destroy(templatePath, vars, logger)
	if err != nil {
		logger.Printf(`Failed to terraform destroy.
** WARNING ** Terraform destroy has failed which means you must manually delete any resources created by the "terraform apply" run.
Terraform Template Path: %s
Test Name: %s
AWS region: <scroll up to see it>
Official Error Message: %s
`, templatePath, testName, err.Error())
	}
}