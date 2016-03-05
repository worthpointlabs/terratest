package aws

import (
	"testing"
)

func TestGetRandomRegionExcludesForbiddenRegions(t *testing.T) {

	for i := 0; i < 1000; i++ {
		randomRegion, _ := GetRandomRegion()
		for _, forbiddenRegion := range GetForbiddenRegions() {
			if forbiddenRegion == randomRegion {
				t.Fatalf("Returned a forbidden AWS Region: %s", forbiddenRegion)
			}
		}
	}

}

func TestGetRandomRegionReturnsAZs(t *testing.T) {

	for i := 0; i < 1000; i++ {
		randomRegion, randomRegionAZs := GetRandomRegion()

		switch randomRegion {
		case "us-west-1":
			regionAZs := "us-west-1a,us-west-1b"
			if randomRegionAZs != regionAZs {
				t.Fatalf("AZs for %s do not match '%s", randomRegion, regionAZs)
			}
		default:
		}
	}

}