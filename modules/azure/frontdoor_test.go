// +build azure

// NOTE: We use build tags to differentiate azure testing because we currently do not have azure access setup for
// CircleCI.

package azure

import (
	"testing"

	"github.com/stretchr/testify/require"
)

/*
The below tests are currently stubbed out, with the expectation that they will throw errors.
If/when methods to create and delete front door are added, these tests can be extended.
*/

func TestFrontDoorExists(t *testing.T) {
	t.Parallel()

	instanceName := "TestFrontDoorInstance"
	resourceGroupName := "TestResourceGroup"
	subscriptionID := ""

	exists, err := FrontDoorExistsE(instanceName, resourceGroupName, subscriptionID)

	require.False(t, exists)
	require.Error(t, err)
}

func TestGetFrontDoor(t *testing.T) {
	t.Parallel()

	instanceName := "TestFrontDoorInstance"
	resourceGroupName := "TestResourceGroup"
	subscriptionID := ""

	instance, err := GetFrontDoorE(instanceName, resourceGroupName, subscriptionID)

	require.Nil(t, instance)
	require.Error(t, err)
}

func TestFrontDoorFrontendEndpointExists(t *testing.T) {
	t.Parallel()

	endpointName := "TestFrontendEndpoint"
	instanceName := "TestFrontDoorInstance"
	resourceGroupName := "TestResourceGroup"
	subscriptionID := ""

	endpoint, err := FrontDoorFrontendEndpointExistsE(endpointName, instanceName, resourceGroupName, subscriptionID)

	require.False(t, endpoint)
	require.Error(t, err)
}

func TestGetFrontDoorFrontendEndpoint(t *testing.T) {
	t.Parallel()

	endpointName := "TestFrontendEndpoint"
	instanceName := "TestFrontDoorInstance"
	resourceGroupName := "TestResourceGroup"
	subscriptionID := ""

	endpoint, err := GetFrontDoorFrontendEndpointE(endpointName, instanceName, resourceGroupName, subscriptionID)

	require.Nil(t, endpoint)
	require.Error(t, err)
}
