// Integration tests that test cross-package functionality in AWS.
package terratest

import (
	"testing"
)

func TestCreateRandomResourceCollectionOptionsForbiddenRegionsWorks(t *testing.T) {
	t.Parallel()

	ro := CreateRandomResourceCollectionOptions()

	// Specify every region but us-east-1
	ro.ForbiddenRegions = []string{
		"us-west-1",
		"us-west-2",
		"eu-west-1",
		"eu-central-1",
		"ap-northeast-1",
		"ap-northeast-2",
		"ap-southeast-1",
		"ap-southeast-2",
		"sa-east-1"}

	rand, err := CreateRandomResourceCollection(ro)
	if err != nil {
		t.Fatalf("Failed to create RandomResourceCollection: %s", err.Error())
	}

	if rand.AwsRegion != "us-east-1" {
		t.Fatalf("Failed to correctly forbid AWS regions. Only valid response should have been us-east-1, but was: %s", rand.AwsRegion)
	}
}
