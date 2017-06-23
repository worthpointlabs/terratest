package resources

import (
	"testing"
	"github.com/gruntwork-io/terratest"
)

func CreateBaseRandomResourceCollection(t *testing.T, requireNatGateway bool, requireEcs bool) *terratest.RandomResourceCollection {
	exludedRegions := REGIONS_WITHOUT_T2_NANO

	if (requireNatGateway) {
		exludedRegions = append(exludedRegions, REGIONS_WITHOUT_NAT_GATEWAY_SUPPORT...)
	}

	if (requireEcs) {
		exludedRegions = append(exludedRegions, REGIONS_WITHOUT_ECS_SUPPORT...)
	}

	resourceCollectionOptions := terratest.NewRandomResourceCollectionOptions()
	resourceCollectionOptions.ForbiddenRegions = exludedRegions

	randomResourceCollection, err := terratest.CreateRandomResourceCollection(resourceCollectionOptions)
	if err != nil {
		t.Fatalf("Failed to create random resource collection: %s\n", err.Error())
	}

	return randomResourceCollection
}
