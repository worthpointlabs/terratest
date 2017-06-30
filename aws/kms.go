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
