package azure

import (
	"testing"

	"github.com/stretchr/testify/require"
)

/*
The below tests are currently stubbed out, with the expectation that they will throw errors.
If/when methods to create and delete front door are added, these tests can be extended.
*/

func TestFrontDoorInstanceExists(t *testing.T) {
	t.Parallel()

	instanceName := "TestFrontDoorInstance"
	resourceGroupName := "TestResourceGroup"
	subscriptionID := ""

	exists, err := FrontDoorInstanceExistsE(instanceName, resourceGroupName, subscriptionID)

	require.False(t, exists)
	require.Error(t, err)
}

func TestGetFrontDoorInstance(t *testing.T) {
	t.Parallel()

	instanceName := "TestFrontDoorInstance"
	resourceGroupName := "TestResourceGroup"
	subscriptionID := ""

	instance, err := GetFrontDoorInstanceE(instanceName, resourceGroupName, subscriptionID)

	require.Nil(t, instance)
	require.Error(t, err)
}

func TestFrontendEndpointExistsForFrontDoorInstance(t *testing.T) {
	t.Parallel()

	endpointName := "TestFrontendEndpoint"
	instanceName := "TestFrontDoorInstance"
	resourceGroupName := "TestResourceGroup"
	subscriptionID := ""

	endpoint, err := FrontendEndpointExistsE(endpointName, instanceName, resourceGroupName, subscriptionID)

	require.False(t, endpoint)
	require.Error(t, err)
}

func TestGetFrontendEndpointForFrontDoorInstance(t *testing.T) {
	t.Parallel()

	endpointName := "TestFrontendEndpoint"
	instanceName := "TestFrontDoorInstance"
	resourceGroupName := "TestResourceGroup"
	subscriptionID := ""

	endpoint, err := GetFrontendEndpointForFrontDoorInstanceE(endpointName, instanceName, resourceGroupName, subscriptionID)

	require.Nil(t, endpoint)
	require.Error(t, err)
}
