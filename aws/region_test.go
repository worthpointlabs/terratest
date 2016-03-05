package aws

import (
	"testing"
)

func TestGetRegionExcludesForbiddenRegions(t *testing.T) {

	for i := 0; i < 1000; i++ {
		randomRegion := GetRegion()
		for _, forbiddenRegion := range GetForbiddenRegions() {
			if forbiddenRegion == randomRegion {
				t.Fatalf("Returned a forbidden AWS Region: %s", forbiddenRegion)
			}
		}
	}
}