package azure

import (
	"context"
	"testing"

	"github.com/Azure/azure-sdk-for-go/services/containerregistry/mgmt/2019-05-01/containerregistry"

	"github.com/stretchr/testify/require"
)

// GetACRClient is a helper function that will setup an Azure Container Instances client on your behalf
// resourceName - required to find the Function App
// resGroupName - use an empty string if you have the AZURE_RES_GROUP_NAME environment variable set
// subscriptionId - use an empty string if you have the ARM_SUBSCRIPTION_ID environment variable set
func GetACRClient(t *testing.T, resourceName string, resGroupName string, subscriptionID string) *containerregistry.Registry {
	resource, err := getACRClientE(resourceName, resGroupName, subscriptionID)

	require.NoError(t, err)

	return resource
}

func getACRClientE(resourceName string, resGroupName string, subscriptionID string) (*containerregistry.Registry, error) {
	// Validate Azure subscription ID
	subscriptionID, err := getTargetAzureSubscription(subscriptionID)
	if err != nil {
		return nil, err
	}

	managedServicesClient := containerregistry.NewRegistriesClient(subscriptionID)
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
