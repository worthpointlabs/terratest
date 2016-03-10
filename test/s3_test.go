// Integration tests that validate S3-related code in AWS.
package test

import (
	"strings"
	"testing"

	"github.com/gruntwork-io/terraform-test/aws"
	"github.com/gruntwork-io/terraform-test/log"
	"github.com/gruntwork-io/terraform-test/util"
)

func TestCreateAndDestroyS3Bucket(t *testing.T) {
	logger := log.NewLogger("TestCreateAndDestroyS3Bucket")

	// SETUP
	region := aws.GetRandomRegion()
	id := util.UniqueId()
	logger.Printf("Random values selected. Region = %s, Id = %s\n", region, id)

	// TEST
	s3BucketName := "gruntwork-terraform-test-" + strings.ToLower(id)

	aws.CreateS3Bucket(region, s3BucketName)
	aws.DeleteS3Bucket(region, s3BucketName)
}

func TestAssertS3BucketExistsNoFalseNegative(t *testing.T) {
	logger := log.NewLogger("TestAssertS3BucketExists")

	// SETUP
	region := aws.GetRandomRegion()
	s3BucketName := "gruntwork-terraform-test-" + strings.ToLower(util.UniqueId())
	logger.Printf("Random values selected. Region = %s, s3BucketName = %s\n", region, s3BucketName)

	aws.CreateS3Bucket(region, s3BucketName)

	// TEST
	err := aws.AssertS3BucketExists(region, s3BucketName)
	if err != nil {
		t.Fatalf("Function claimed that S3 Bucket '%s' does not exist, but in fact it does.", s3BucketName)
	}

	// TEARDOWN
	aws.DeleteS3Bucket(region, s3BucketName)
}

func TestAssertS3BucketExistsNoFalsePositive(t *testing.T) {
	logger := log.NewLogger("TestAssertS3BucketExists")

	// SETUP
	region := aws.GetRandomRegion()
	s3BucketName := "gruntwork-terraform-test-" + strings.ToLower(util.UniqueId())
	logger.Printf("Random values selected. Region = %s, s3BucketName = %s\n", region, s3BucketName)

	// We elect not to create the S3 bucket to confirm that our function correctly reports it doesn't exist.
	//aws.CreateS3Bucket(region, s3BucketName)

	// TEST
	err := aws.AssertS3BucketExists(region, s3BucketName)
	if err == nil {
		t.Fatalf("Function claimed that S3 Bucket '%s' exists, but in fact it does not.", s3BucketName)
	}
}