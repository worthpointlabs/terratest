package azure

import (
	"context"
	"testing"

	"github.com/Azure/azure-sdk-for-go/services/web/mgmt/2019-08-01/web"
	"github.com/stretchr/testify/require"
)

// GetFunctionApp gets the Function App
// resourceName - required to find the Function App
// resGroupName - use an empty string if you have the AZURE_RES_GROUP_NAME environment variable set
// subscriptionId - use an empty string if you have the ARM_SUBSCRIPTION_ID environment variable set
func GetFunctionApp(t *testing.T, resourceName string, resGroupName string, subscriptionID string) *web.Site {
	site, err := getAppServicePlanE(resourceName, resGroupName, subscriptionID)

	require.NoError(t, err)

	return site
}

func getAppServicePlanE(resourceName string, resGroupName string, subscriptionID string) (*web.Site, error) {
	rgName, err := getTargetAzureResourceGroupName(resGroupName)
	if err != nil {
		return nil, err
	}

	client, err := getFunctionAppClientE(subscriptionID)
	if err != nil {
		return nil, err
	}

	resource, err := client.Get(context.Background(), rgName, resourceName)
	if err != nil {
		return nil, err
	}

	return &resource, nil
}

func getFunctionAppClientE(subscriptionID string) (*web.AppsClient, error) {
	// Validate Azure subscription ID
	subscriptionID, err := getTargetAzureSubscription(subscriptionID)
	if err != nil {
		return nil, err
	}

	managedServicesClient := web.NewAppsClient(subscriptionID)
	authorizer, err := NewAuthorizer()

	if err != nil {
		return nil, err
	}

	managedServicesClient.Authorizer = *authorizer

	return &managedServicesClient, nil
}
