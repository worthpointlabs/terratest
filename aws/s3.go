package aws

import (
	"os"
	"strings"

	"github.com/gruntwork-io/terraform-test/log"

	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/service/s3"
	"github.com/aws/aws-sdk-go/aws/session"
)

// Create an S3 bucket.
func CreateS3Bucket(region string, name string) {
	log := log.NewLogger("CreateS3Bucket")

	svc := s3.New(session.New(), aws.NewConfig().WithRegion(region))
	_, err := svc.Config.Credentials.Get()
	if err != nil {
		log.Fatalf("Failed to open S3 session: %s\n", err.Error())
	}

	params := &s3.CreateBucketInput{
		Bucket: aws.String(name),
	}
	_, err = svc.CreateBucket(params)
	if err != nil {
		log.Printf("Failed to create S3 bucket: %s", err.Error())
		os.Exit(1)
	}
}

// Destroy an S3 bucket.
func DeleteS3Bucket(region string, name string) {
	log := log.NewLogger("DestroyS3Bucket")

	svc := s3.New(session.New(), aws.NewConfig().WithRegion(region))
	_, err := svc.Config.Credentials.Get()
	if err != nil {
		log.Fatalf("Failed to open S3 session: %s\n", err.Error())
	}

	params := &s3.DeleteBucketInput{
		Bucket: aws.String(name),
	}
	_, err = svc.DeleteBucket(params)
	if err != nil {
		log.Printf("Failed to delete S3 bucket: %s", err.Error())
		os.Exit(1)
	}
}

// Returns true if the given S3 bucket exists.
func AssertS3BucketExists(region string, name string) bool {
	log := log.NewLogger("AssertS3BucketExists")

	svc := s3.New(session.New(), aws.NewConfig().WithRegion(region))
	_, err := svc.Config.Credentials.Get()
	if err != nil {
		log.Fatalf("Failed to open S3 session: %s\n", err.Error())
	}

	params := &s3.HeadBucketInput{
		Bucket: aws.String(name),
	}
	_, err = svc.HeadBucket(params)

	bucketExists := false
	if err != nil {
		// We expect a missing bucket to return this error code.  Otherwise, fail because we can't be sure what
		// the AWS response means.
		if ! strings.Contains(err.Error(), "status code: 404") {
			log.Printf("Failed to assert whether bucket exists: %s", err.Error())
			os.Exit(1)
		}
	} else {
		bucketExists = true
	}

	return bucketExists
}