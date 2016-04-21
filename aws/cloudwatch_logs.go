package aws

import (
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/cloudwatchlogs"
	"github.com/aws/aws-sdk-go/aws"
)

// Return the CloudWatch log messages in the given region for the given log stream and log group
func GetCloudWatchLogEntries(awsRegion string, logStreamName string, logGroupName string) ([]string, error) {
	entries := []string{}

	svc := cloudwatchlogs.New(session.New(), aws.NewConfig().WithRegion(awsRegion))
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
