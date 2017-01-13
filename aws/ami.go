package aws

// Return an Ubuntu 14.04 LTS public AMI from the given region.
// The choice of Ubuntu is somewhat arbitrary. It's expected that this function wil be used to populate launch configs
// and Bastion Host AMI choices, so this seemed like a sensible default.
func GetUbuntu1404Ami(region string) string {
	amis := map[string]string{
		"ap-northeast-1": "ami-6fccbe08",
		"ap-northeast-2": "<ami-not-available-for-region-ap-northeast-2>",
		"ap-south-1":     "<ami-not-available-for-region-ap-south-1>",
		"ap-southeast-1": "ami-50e64d33",
		"ap-southeast-2": "<ami-not-available-for-region-ap-southeast-2>",
		"ca-central-1":   "<ami-not-available-for-region-ca-central-1>",
		"eu-central-1":   "ami-78559817",
		"eu-west-1":      "ami-a192bad2",
		"eu-west-2":      "<ami-not-available-for-region-us-west-2>",
		"sa-east-1":      "ami-ff861c93",
		"us-east-1":      "ami-49c9295f",
		"us-east-2":      "<ami-not-available-for-region-us-east-2>",
		"us-west-1":      "ami-3e21725e",
		"us-west-2":      "ami-5e63d13e",
	}
	return amis[region]

}

// Return an Ubuntu 16.04 LTS - Xenial (HVM) public AMI from the given region.
// https://cloud-images.ubuntu.com/locator/ec2/
func GetUbuntu1604Ami(region string) string {
	amis := map[string]string{
		"ap-northeast-1": "ami-18afc47f",
		"ap-northeast-2": "ami-93d600fd",
		"ap-south-1":     "ami-dd3442b2",
		"ap-southeast-1": "ami-87b917e4",
		"ap-southeast-2": "ami-e6b58e85",
		"ca-central-1":   "ami-7112a015",
		"eu-central-1":   "ami-fe408091",
		"eu-west-1":      "ami-ca80a0b9",
		"eu-west-2":      "ami-ede2e889",
		"sa-east-1":      "ami-e075ed8c",
		"us-east-1":      "ami-9dcfdb8a",
		"us-east-2":      "ami-fcc19b99",
		"us-west-1":      "ami-b05203d0",
		"us-west-2":      "ami-b2d463d2",
	}
	return amis[region]

}

// Return a CentOS 7 public AMI from the given region.
// https://aws.amazon.com/marketplace/pp/B00O7WM7QW
// WARNING: I believe you have to accept the terms & conditions of this AMI in AWS MarketPlace for your AWS Account before
// you can successfully launch the AMI.
func GetCentos7Ami(region string) string {
	amis := map[string]string{
		"ap-northeast-1": "ami-eec1c380",
		"ap-northeast-2": "ami-c74789a9",
		"ap-south-1":	  "ami-95cda6fa",
		"ap-southeast-1": "ami-f068a193",
		"ap-southeast-2": "ami-fedafc9d",
		"ca-central-1":   "ami-af62d0cb",
		"eu-central-1":   "ami-9bf712f4",
		"eu-west-1":      "ami-7abd0209",
		"eu-west-2":      "ami-bb373ddf",
		"sa-east-1":      "ami-26b93b4a",
		"us-east-1":      "ami-6d1c2007",
		"us-east-2":      "ami-6a2d760f",
		"us-west-1":      "ami-af4333cf",
		"us-west-2":      "ami-d2c924b2",
	}
	return amis[region]

}

// Return an Amazon Linux AMI HVM, SSD Volume Type public AMI for the given region. This AMI is useful
// when you want to test with an AMI that has AWS utilities pre-installed, such as the awscli or cfn-signal.
func GetAmazonLinuxAmi(region string) string {
	amis := map[string]string{
		"ap-northeast-1": "ami-9f0c67f8",
		"ap-northeast-2": "ami-94bb6dfa",
		"ap-south-1":     "ami-9fc7b0f0",
		"ap-southeast-1": "ami-4dd6782e",
		"ap-southeast-2": "ami-28cff44b",
		"ca-central-1":   "ami-eb20928f",
		"eu-central-1":   "ami-211ada4e",
		"eu-west-1":      "ami-c51e3eb6",
		"eu-west-2":      "ami-bfe0eadb",
		"sa-east-1":      "ami-bb40d8d7",
		"us-east-1":      "ami-9be6f38c",
		"us-east-2":      "ami-38cd975d",
		"us-west-1":      "ami-b73d6cd7",
		"us-west-2":      "ami-1e299d7e",
	}
	return amis[region]
}

// Return an Amazon ECS-Optimized Amazon Linux AMI for the given region. This AMI is useful for running an ECS cluster.
// http://docs.aws.amazon.com/AmazonECS/latest/developerguide/ecs-optimized_AMI.html
func GetEcsOptimizedAmazonLinuxAmi(region string) string {
	// Regions not supported by ECS: ap-northeast-2, sa-east-1
	amis := map[string]string{
		"ap-northeast-1": "ami-30bdce57",
		"ap-northeast-2": "<ami-not-available-for-region-ap-northeast-2>",
		"ap-south-1":     "<ami-not-available-for-region-ap-south-1>",
		"ap-southeast-1": "ami-9f75ddfc",
		"ap-southeast-2": "ami-cf393cac",
		"ca-central-1":   "ami-1b01b37f",
		"eu-central-1":   "ami-38dc1157",
		"eu-west-1":      "ami-e3fbd290",
		"eu-west-2":      "ami-77f6fc13",
		"sa-east-1":      "<ami-not-available-for-region-sa-east-1>",
		"us-east-1":      "ami-a58760b3",
		"us-east-2":      "ami-a6e4bec3",
		"us-west-1":      "ami-74cb9b14",
		"us-west-2":      "ami-5b6dde3b",
	}

	return amis[region]
}
