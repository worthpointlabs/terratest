package aws

// Return an Ubuntu 14.04 LTS public AMI from the given region.
// The choice of Ubuntu is somewhat arbitrary. It's expected that this function wil be used to populate launch configs
// and Bastion Host AMI choices, so this seemed like a sensible default.
func GetUbuntuAmi(region string) string {

	amis := make(map[string]string)
	amis["us-east-1"] 	= "ami-fce3c696"
	amis["us-west-1"] 	= "ami-06116566"
	amis["us-west-2"] 	= "ami-9abea4fb"
	amis["eu-west-1"] 	= "ami-f95ef58a"
	amis["eu-central-1"] 	= "ami-87564feb"
	amis["ap-northeast-1"] 	= "ami-a21529cc"
	amis["ap-northeast-2"] 	= "ami-09dc1267"
	amis["ap-southeast-1"] 	= "ami-25c00c46"
	amis["ap-southeast-2"] 	= "ami-6c14310f"
	amis["sa-east-1"] 	= "ami-0fb83963"

	return amis[region]

}
