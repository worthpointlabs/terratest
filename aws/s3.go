package aws

import (
	"os"

	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/service/s3"
	"github.com/aws/aws-sdk-go/aws/session"

	"github.com/gruntwork-io/terraform-test/log"
)

// Create an S3 bucket
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

// Destroy an S3 bucket
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