package azure

import (
	"context"
	"testing"

	"github.com/Azure/azure-sdk-for-go/services/containerinstance/mgmt/2018-10-01/containerinstance"

	"github.com/stretchr/testify/require"
)

// GetACIClient is a helper function that will setup an Azure Container Instances client on your behalf
// resourceName - required to find the Function App
// resGroupName - use an empty string if you have the AZURE_RES_GROUP_NAME environment variable set
// subscriptionId - use an empty string if you have the ARM_SUBSCRIPTION_ID environment variable set
func GetACIClient(t *testing.T, resourceName string, resGroupName string, subscriptionID string) *containerinstance.ContainerGroup {
	resource, err := getACIClientE(resourceName, resGroupName, subscriptionID)

	require.NoError(t, err)

	return resource
}

func getACIClientE(resourceName string, resGroupName string, subscriptionID string) (*containerinstance.ContainerGroup, error) {
	// Validate Azure subscription ID
	subscriptionID, err := getTargetAzureSubscription(subscriptionID)
	if err != nil {
		return nil, err
	}

	managedServicesClient := containerinstance.NewContainerGroupsClient(subscriptionID)
	authorizer, err := NewAuthorizer()

	if err != nil {
		return nil, err
	}

	managedServicesClient.Authorizer = *authorizer

	resource, err := managedServicesClient.Get(context.Background(), resGroupName, resourceName)
	if err != nil {
		return nil, err
	}

	return &resource, nil
}
