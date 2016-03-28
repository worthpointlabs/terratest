// Integration tests that test cross-package functionality in AWS.
package terratest

import (
	"testing"
)

func TestCreateRandomResourceCollectionOptionsForbiddenRegionsWorks(t *testing.T) {
	t.Parallel()

	ro := NewRandomResourceCollectionOptions()

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

func TestFetchAwsAvailabilityZones(t *testing.T) {
	t.Parallel()

	ro := NewRandomResourceCollectionOptions()
	rand, err := CreateRandomResourceCollection(ro)
	if err != nil {
		t.Fatalf("Failed to create RandomResourceCollection: %s", err.Error())
	}

	// Manually set the AWS Region to us-west-2 for testing purposes
	rand.AwsRegion = "us-west-2"
	actual := rand.FetchAwsAvailabilityZones()
	expected := []string{"us-west-2a","us-west-2b","us-west-2c"}
	//expected := []string{"us-west-2a,us-west-2b,us-west-2c"}

	for index,_ := range expected {
		if actual[index] != expected[index] {
			t.Fatalf("Expected: %s, but received %s", expected[index], actual[index])
		}
	}
}

func TestFetchAwsAvailabilityZonesAsString(t *testing.T) {
	t.Parallel()

	ro := NewRandomResourceCollectionOptions()
	rand, err := CreateRandomResourceCollection(ro)
	if err != nil {
		t.Fatalf("Failed to create RandomResourceCollection: %s", err.Error())
	}

	// Manually set the AWS Region to us-west-2 for testing purposes
	rand.AwsRegion = "us-west-2"
	actual := rand.FetchAwsAvailabilityZonesAsString()
	expected := "us-west-2a,us-west-2b,us-west-2c"

	if actual != expected {
		t.Fatalf("Expected: %s, but received %s", expected, actual)
	}
}

func TestGetRandomPrivateCidrBlock(t *testing.T) {
	t.Parallel()

	ro := NewRandomResourceCollectionOptions()
	rand, err := CreateRandomResourceCollection(ro)
	if err != nil {
		t.Fatalf("Failed to create RandomResourceCollection: %s", err.Error())
	}

	actual := rand.GetRandomPrivateCidrBlock(18)
	actualPrefix := string(actual[len(actual)-3:])
	expPrefix := "/18"

	if actualPrefix != expPrefix {
		t.Fatalf("Expected: %s, but received: %s", expPrefix, actualPrefix)
	}
}

func TestAllParametersSet(t *testing.T) {
	t.Parallel()

	ro := NewRandomResourceCollectionOptions()
	rand, err := CreateRandomResourceCollection(ro)
	if err != nil {
		t.Fatalf("Failed to create RandomResourceCollection: %s", err.Error())
	}

	if len(rand.AccountId) == 0 {
		t.Fatalf("CreateRandomResourceCollection has an empty AccountId: %s", rand)
	}

	if len(rand.AmiId) == 0 {
		t.Fatalf("CreateRandomResourceCollection has an empty AMI ID: %s", rand)
	}

	if len(rand.AwsRegion) == 0 {
		t.Fatalf("CreateRandomResourceCollection has an empty region: %s", rand)
	}

	if len(rand.UniqueId) == 0 {
		t.Fatalf("CreateRandomResourceCollection has an empty Unique Id: %s", rand)
	}

	if rand.KeyPair == nil {
		t.Fatalf("CreateRandomResourceCollection has a nil Key Pair: %s", rand)
	}
}

func TestGetRandomVpc(t *testing.T) {
	t.Parallel()

	ro := NewRandomResourceCollectionOptions()
	rand, err := CreateRandomResourceCollection(ro)
	if err != nil {
		t.Fatalf("Failed to create RandomResourceCollection: %s", err.Error())
	}

	vpc, err := rand.GetRandomVpc()
	if err != nil {
		t.Fatalf("Failed to get random VPC: %s", err.Error())
	}

	if vpc.Id == "" {
		t.Fatalf("GetRandomVpc returned a VPC without an ID: %s", vpc)
	}

	if vpc.Name == "" {
		t.Fatalf("GetRandomVpc returned a VPC without a name: %s", vpc)
	}

	if len(vpc.SubnetIds) == 0 {
		t.Fatalf("GetRandomVpc returned a VPC with no subnets: %s", vpc)
	}
}