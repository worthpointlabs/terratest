package aws

import (
	"github.com/aws/aws-sdk-go/service/ec2"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/aws"
	"errors"
	"strconv"
)

var VpcIdFilterName = "vpc-id"

type Vpc struct {
	Id         string 		// The ID of the VPC
	Name       string		// The name of the VPC
	Subnets    []ec2.Subnet	// A list of subnets in the VPC
}

var IS_DEFAULT_FILTER_NAME = "isDefault"
var IS_DEFAULT_FILTER_VALUE = "true"

func GetDefaultVpc(awsRegion string) (Vpc, error) {
	vpc := Vpc{}

	svc := ec2.New(session.New(), aws.NewConfig().WithRegion(awsRegion))
	defaultVpcFilter := ec2.Filter{Name: &IS_DEFAULT_FILTER_NAME, Values: []*string{&IS_DEFAULT_FILTER_VALUE}}
	vpcs, err := svc.DescribeVpcs(&ec2.DescribeVpcsInput{Filters: []*ec2.Filter{&defaultVpcFilter}})
	if err != nil {
		return vpc, err
	}

	numVpcs := len(vpcs.Vpcs)
	if numVpcs != 1 {
		return vpc, errors.New("Expected to find one default VPC in region " + awsRegion + " but found " + strconv.Itoa(numVpcs))
	}

	defaultVpc := vpcs.Vpcs[0]

	vpc.Id = *defaultVpc.VpcId
	vpc.Name = FindVpcName(defaultVpc)

	vpc.Subnets, err = GetSubnetsForVpc(vpc.Id, awsRegion)
	return vpc, err
}

func FindVpcName(vpc *ec2.Vpc) string {
	for _, tag := range vpc.Tags {
		if *tag.Key == "Name" {
			return *tag.Value
		}
	}

	if *vpc.IsDefault {
		return "Default"
	}

	return ""
}

func GetSubnetsForVpc(vpcId string, awsRegion string) ([]ec2.Subnet, error) {
	subnets := []ec2.Subnet{}

	svc := ec2.New(session.New(), aws.NewConfig().WithRegion(awsRegion))

	vpcIdFilter := ec2.Filter{Name: &VpcIdFilterName, Values: []*string{&vpcId}}
	subnetOutput, err := svc.DescribeSubnets(&ec2.DescribeSubnetsInput{Filters: []*ec2.Filter{&vpcIdFilter}})
	if err != nil {
		return subnets, err
	}

	for _, subnet := range subnetOutput.Subnets {
		subnets = append(subnets, *subnet)
	}
	return subnets, nil
}