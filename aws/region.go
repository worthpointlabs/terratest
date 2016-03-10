package aws

import (
	"github.com/gruntwork-io/terraform-test/util"

	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/ec2"
	"github.com/gruntwork-io/terraform-test/log"
)

func GetForbiddenRegions() []string {
	return []string{
		"us-west-2",
	}
}

// Get a randomly chosen AWS region
func GetRandomRegion() string {

	allRegions := []string{
		"us-east-1",
		"us-west-1",
		"us-west-2",
		"eu-west-1",
		"eu-central-1",
		"ap-northeast-1",
		"ap-northeast-2",
		"ap-southeast-1",
		"ap-southeast-2",
		"sa-east-1",
	}

	// Select a random region
	// If our randomIndex gave us a region that's forbidden, keep iterating until we get a valid one.
	var randomIndex int
	randomIndexIsValid := false

	for !randomIndexIsValid {
		randomIndex = util.Random(0,len(allRegions))
		randomIndexIsValid = true

		for _, forbiddenRegion := range GetForbiddenRegions() {
			if forbiddenRegion == allRegions[randomIndex] {
				randomIndexIsValid = false
			}
		}
	}

	return allRegions[randomIndex]
}

// Get the Availability Zones for a given AWS region. Note that for certain regions (e.g. us-east-1), different AWS
// accounts have access to different availability zones.
func GetAvailabilityZones(region string) []string {
	log := log.NewLogger("GetAvailabilityZones")

	svc := ec2.New(session.New(), aws.NewConfig().WithRegion(region))
	_, err := svc.Config.Credentials.Get()
	if err != nil {
		log.Fatalf("Failed to open EC2 session: %s\n", err.Error())
	}

	params := &ec2.DescribeAvailabilityZonesInput{
		DryRun: aws.Bool(false),
	}
	resp, err := svc.DescribeAvailabilityZones(params)
	if err != nil {
		log.Fatalf("Failed to fetch AWS Availability Zones: %s\n", err.Error())
	}

	var azs []string
	for _, availabilityZone := range resp.AvailabilityZones {
		azs = append(azs, *availabilityZone.ZoneName)
	}

	return azs
}

