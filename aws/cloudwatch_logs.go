package aws

import (
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/cloudwatchlogs"
	"github.com/aws/aws-sdk-go/aws"
)

// Return the CloudWatch log messages for the given log group and log stream
func GetCloudWatchLogEntries(logGroupName string, logStreamName string) ([]string, error) {
	entries := []string{}

	svc := cloudwatchlogs.New(session.New())
	output, err := svc.GetLogEvents(&cloudwatchlogs.GetLogEventsInput{
		LogGroupName: aws.String(logGroupName),
		LogStreamName: aws.String(logStreamName),
	})

	if err != nil {
		return entries, err
	}

	for _, event := range output.Events {
		entries = append(entries, *event.Message)
	}

	return entries, nil
}
