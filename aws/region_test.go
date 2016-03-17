package aws

import (
	"testing"
)

func TestGetRandomRegionExcludesForbiddenRegions(t *testing.T) {
	t.Parallel()

	for i := 0; i < 1000; i++ {
		randomRegion := GetRandomRegion()
		for _, forbiddenRegion := range GetForbiddenRegions() {
			if forbiddenRegion == randomRegion {
				t.Fatalf("Returned a forbidden AWS Region: %s", forbiddenRegion)
			}
		}
	}

}

func TestGetAvailabilityZones(t *testing.T) {
	t.Parallel()

	azs := GetAvailabilityZones("us-west-2")

	if azs[0] != "us-west-2a"  {
		t.Fatalf("Expected us-west-2a, received: %s", azs[0])
	}

	if azs[1] != "us-west-2b"  {
		t.Fatalf("Expected us-west-2b, received: %s", azs[1])
	}

	if azs[2] != "us-west-2c"  {
		t.Fatalf("Expected us-west-2c, received: %s", azs[2])
	}
}

