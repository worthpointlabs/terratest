package aws

import (
	"github.com/aws/aws-sdk-go/service/sns"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/aws"
	"fmt"
)

// Create an SNS Topic and return the ARN
func CreateSnsTopic(awsRegion string, snsTopicName string) (string, error) {
	var snsTopicArn string

	svc := sns.New(session.New(), aws.NewConfig().WithRegion(awsRegion))

	createTopicInput := &sns.CreateTopicInput{
		Name: &snsTopicName,
	}

	output, err := svc.CreateTopic(createTopicInput)
	if err != nil {
		return snsTopicArn, fmt.Errorf("Failed to create SNS topic: %s\n", err.Error())
	}

	snsTopicArn = *output.TopicArn

	return snsTopicArn, err
}

// Delete an SNS Topic
func DeleteSNSTopic(awsRegion string, snsTopicArn string) error {
	svc := sns.New(session.New(), aws.NewConfig().WithRegion(awsRegion))

	deleteTopicInput := &sns.DeleteTopicInput{
		TopicArn: aws.String(snsTopicArn),
	}

	_, err := svc.DeleteTopic(deleteTopicInput)
	if err != nil {
		return fmt.Errorf("Failed to delete the SNS Topic '%s': %s\n", snsTopicArn, err.Error())
	}

	return nil
}
