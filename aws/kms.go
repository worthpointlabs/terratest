package aws

import (
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/kms"
	"github.com/aws/aws-sdk-go/aws"
)

func CreateKmsClient(awsRegion string) (*kms.KMS, error) {
	awsConfig, err := CreateAwsConfig(awsRegion)
	if err != nil {
		return nil, err
	}

	return kms.New(session.New(), awsConfig), nil
}

// This exists because KMS keys cost $1/mo for each one created. In our automated tests, we often run 100s times in,
// especially during inital development when debugging. Rather than create a new key each time we run a test and incurring
// the $1 charge, a dedicated key with the alias 'dedicated-test-key' has been created in each region. This method allows
// this key to be retrieved and used for testing purposes.
func GetDedicatedTestKeyArn(awsRegion string) (string, error) {
	kmsClient, err := CreateKmsClient(awsRegion)
	if err != nil {
		return "", err
	}

	result, err := kmsClient.DescribeKey(&kms.DescribeKeyInput{
		KeyId: aws.String("alias/dedicated-test-key"),
	})

	if err != nil {
		return "", err
	}

	return *result.KeyMetadata.Arn, nil
}
