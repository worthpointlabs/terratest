package aws

// Return an Ubuntu 14.04 LTS public AMI from the given region.
// The choice of Ubuntu is somewhat arbitrary. It's expected that this function wil be used to populate launch configs
// and Bastion Host AMI choices, so this seemed like a sensible default.
func GetUbuntuAmi(region string) string {
	amis := map[string]string{
		"us-east-1": "ami-fce3c696",
		"us-west-1": "ami-06116566",
		"us-west-2": "ami-9abea4fb",
		"eu-west-1": "ami-f95ef58a",
		"eu-central-1": "ami-87564feb",
		"ap-northeast-1": "ami-a21529cc",
		"ap-northeast-2": "ami-09dc1267",
		"ap-southeast-1": "ami-25c00c46",
		"ap-southeast-2": "ami-6c14310f",
		"sa-east-1": "ami-0fb83963",
	}
	return amis[region]

}

// Return an Amazon Linux AMI 2015.09.2 (HVM), SSD Volume Type public AMI for the given region. This AMI is useful
// when you want to test with an AMI that has AWS utilities pre-installed, such as the awscli or cfn-signal.
func GetAmazonLinuxAmi(region string) string {
	amis := map[string]string{
		"us-east-1": "ami-08111162",
		"us-west-1": "ami-1b0f7d7b",
		"us-west-2": "ami-c229c0a2",
		"eu-west-1": "ami-31328842",
		"eu-central-1": "ami-e2df388d",
		"ap-northeast-1": "ami-f80e0596",
		"ap-northeast-2": "	ami-6598510b",
		"ap-southeast-1": "ami-e90dc68a",
		"ap-southeast-2": "ami-f2210191",
		"sa-east-1": "ami-1e159872",
	}
	return amis[region]
}

// Return an Amazon ECS-Optimized Amazon Linux AMI 2016.03.c AMI for the given region. This AMI is useful for running
// an ECS cluster.
func GetEcsOptimizedAmazonLinuxAmi(region string) string {
	// Regions not supported by ECS: ap-northeast-2, sa-east-1
	amis := map[string]string {
		"us-east-1": "ami-719e631c",
		"us-west-1": "ami-87d6ade7",
		"us-west-2": "ami-beba42de",
		"eu-west-1": "ami-949704e7",
		"eu-central-1": "ami-7d5ab512",
		"ap-northeast-1": "ami-92638ff3",
		"ap-southeast-1": "ami-489a4a2b",
		"ap-southeast-2": "ami-f01d3393",
	}

	return amis[region]
}