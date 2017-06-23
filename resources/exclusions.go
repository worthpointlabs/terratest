package resources

// As of 4/6/16, these AWS regions do not support the NAT Gateway service.
var REGIONS_WITHOUT_NAT_GATEWAY_SUPPORT = []string{
	"eu-central-1",
	"ap-northeast-2",
	"sa-east-1",
}

// As of 6/9/16, these AWS regions do not support t2.nano instances
var REGIONS_WITHOUT_T2_NANO = []string{
	"ap-southeast-2",
}

var REGIONS_WITHOUT_ECS_SUPPORT = []string{
	"ap-northeast-2",
	"sa-east-1",
}