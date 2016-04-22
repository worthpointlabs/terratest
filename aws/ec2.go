package aws

import (
	"github.com/aws/aws-sdk-go/service/ec2"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/aws"
	"fmt"
)

// Return all the ids of EC2 instances in the given region with the given tag
func GetEc2InstanceIdsByTag(awsRegion string, tagName string, tagValue string) ([]string, error) {
	instanceIds := []string{}
	svc := ec2.New(session.New(), aws.NewConfig().WithRegion(awsRegion))

	tagFilter := &ec2.Filter{
		Name: aws.String(fmt.Sprintf("tag:%s", tagName)),
		Values: []*string{aws.String(tagValue)},
	}
	output, err := svc.DescribeInstances(&ec2.DescribeInstancesInput{Filters: []*ec2.Filter{tagFilter}})
	if err != nil {
		return instanceIds, err
	}

	for _, reservation := range output.Reservations {
		for _, instance := range reservation.Instances {
			instanceIds = append(instanceIds, *instance.InstanceId)
		}
	}

	return instanceIds, err
}
