package aws

// Return an Ubuntu 14.04 LTS public AMI from the given region.
// The choice of Ubuntu is somewhat arbitrary. It's expected that this function wil be used to populate launch configs
// and Bastion Host AMI choices, so this seemed like a sensible default.
func GetUbuntuAmi(region string) string {
	amis := map[string]string{
		"us-east-1":      "ami-fce3c696",
		"us-west-1":      "ami-06116566",
		"us-west-2":      "ami-9abea4fb",
		"eu-west-1":      "ami-f95ef58a",
		"eu-central-1":   "ami-87564feb",
		"ap-northeast-1": "ami-a21529cc",
		"ap-northeast-2": "ami-09dc1267",
		"ap-southeast-1": "ami-25c00c46",
		"ap-southeast-2": "ami-6c14310f",
		"sa-east-1":      "ami-0fb83963",
	}
	return amis[region]

}

// Return an Ubuntu 16.04 LTS - Xenial (HVM) public AMI from the given region.
// https://aws.amazon.com/marketplace/ordering?productId=d83d0782-cb94-46d7-8993-f4ce15d1a484&ref_=dtl_psb_continue&region=us-east-1
func GetUbuntu1604Ami(region string) string {
	amis := map[string]string{
		"us-east-1":      "ami-29f96d3e",
		"us-west-1":      "ami-26155546",
		"us-west-2":      "ami-114b8471",
		"eu-west-1":      "ami-be3559cd",
		"eu-central-1":   "ami-9e6a9ef1",
		"ap-northeast-1": "ami-e95da788",
		"ap-northeast-2": "ami-a9f63cc7",
		"ap-southeast-1": "ami-041cc367",
		"ap-southeast-2": "ami-c52114a6",
		"sa-east-1":      "ami-d0bb2cbc",
	}
	return amis[region]

}

// Return a CentOS 7 public AMI from the given region.
// https://aws.amazon.com/marketplace/pp/B00O7WM7QW
// WARNING: I believe you have to accept the terms & conditions of this AMI in AWS MarketPlace for your AWS Account before
// you can successfully launch the AMI.
func GetCentos7Ami(region string) string {
	amis := map[string]string{
		"us-east-1":      "ami-6d1c2007",
		"us-west-1":      "ami-af4333cf",
		"us-west-2":      "ami-d2c924b2",
		"eu-west-1":      "ami-7abd0209",
		"eu-central-1":   "ami-9bf712f4",
		"ap-northeast-1": "ami-eec1c380",
		"ap-northeast-2": "ami-c74789a9",
		"ap-southeast-1": "ami-f068a193",
		"ap-southeast-2": "ami-fedafc9d",
		"sa-east-1":      "ami-26b93b4a",
	}
	return amis[region]

}

// Return an Amazon Linux AMI 2015.09.2 (HVM), SSD Volume Type public AMI for the given region. This AMI is useful
// when you want to test with an AMI that has AWS utilities pre-installed, such as the awscli or cfn-signal.
func GetAmazonLinuxAmi(region string) string {
	amis := map[string]string{
		"us-east-1":      "ami-08111162",
		"us-west-1":      "ami-1b0f7d7b",
		"us-west-2":      "ami-c229c0a2",
		"eu-west-1":      "ami-31328842",
		"eu-central-1":   "ami-e2df388d",
		"ap-northeast-1": "ami-f80e0596",
		"ap-northeast-2": "	ami-6598510b",
		"ap-southeast-1": "ami-e90dc68a",
		"ap-southeast-2": "ami-f2210191",
		"sa-east-1":      "ami-1e159872",
	}
	return amis[region]
}

// Return an Amazon ECS-Optimized Amazon Linux AMI 2016.03.c AMI for the given region. This AMI is useful for running
// an ECS cluster.
func GetEcsOptimizedAmazonLinuxAmi(region string) string {
	// Regions not supported by ECS: ap-northeast-2, sa-east-1
	amis := map[string]string{
		"us-east-1":      "ami-719e631c",
		"us-west-1":      "ami-87d6ade7",
		"us-west-2":      "ami-beba42de",
		"eu-west-1":      "ami-949704e7",
		"eu-central-1":   "ami-7d5ab512",
		"ap-northeast-1": "ami-92638ff3",
		"ap-southeast-1": "ami-489a4a2b",
		"ap-southeast-2": "ami-f01d3393",
	}

	return amis[region]
}
