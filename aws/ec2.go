package aws

import (
	"github.com/aws/aws-sdk-go/service/ec2"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/aws"
)

// Return all the ids of EC2 instances with the given tag
func getEc2InstanceIdsByTag(tagName string, tagValue string) ([]string, error) {
	instanceIds := []string{}
	svc := ec2.New(session.New())

	// TODO: filter using tags
	asgFilter := &ec2.Filter{
		Name: aws.String("requester-id"),
		Values: []*string{aws.String(asgId)},
	}
	output, err := svc.DescribeInstances(&ec2.DescribeInstancesInput{Filters: []*ec2.Filter{asgFilter}})
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
