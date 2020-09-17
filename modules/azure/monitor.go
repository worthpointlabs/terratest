package azure

import (
	"context"

	"github.com/Azure/azure-sdk-for-go/services/preview/monitor/mgmt/2019-06-01/insights"
	"github.com/gruntwork-io/terratest/modules/testing"
	"github.com/stretchr/testify/require"
)

// DiagnosticSettingsResourceExists indicates whether the speficied Azure Availability Set exists
func DiagnosticSettingsResourceExists(t testing.TestingT, actionGroupName string, resGroupName string, subscriptionID string) bool {
	exists, err := DiagnosticSettingsResourceExistsE(t, actionGroupName, resGroupName, subscriptionID)
	require.NoError(t, err)
	return exists
}

// DiagnosticSettingsResourceExistsE indicates whether the speficied Azure Availability Set exists
func DiagnosticSettingsResourceExistsE(t testing.TestingT, actionGroupName string, resGroupName string, subscriptionID string) (bool, error) {
	_, err := GetDiagnosticsSettingsE(t, actionGroupName, resGroupName, subscriptionID)
	if err != nil {
		return false, err
	}
	return true, nil
}

// GetDiagnosticsSettingsE gets the diagnostics settings for a specified resource
func GetDiagnosticsSettingsE(t testing.TestingT, name string, resourceURI string, subscriptionID string) (*insights.DiagnosticSettingsResource, error) {
	client, err := GetDiagnosticsSettingsClient(t, subscriptionID)

	if err != nil {
		return nil, err
	}

	settings, err := client.Get(context.Background(), resourceURI, name)

	if err != nil {
		return nil, err
	}

	return &settings, nil
}

// GetDiagnosticsSettingsClient returns diagnostics settings client
func GetDiagnosticsSettingsClient(t testing.TestingT, subscriptionID string) (*insights.DiagnosticSettingsClient, error) {
	client := insights.NewDiagnosticSettingsClient(subscriptionID)
	authorizer, err := NewAuthorizer()

	if err != nil {
		return nil, err
	}

	client.Authorizer = *authorizer
	return &client, nil
}

// GetVMInsights get diagnostics VM onboarding status
func GetVMInsights(t testing.TestingT, resourceURI string, subscriptionID string) (*insights.VMInsightsOnboardingStatus, error) {
	client, err := GetVMInsightsClient(t, subscriptionID)

	if err != nil {
		return nil, err
	}

	status, err := client.GetOnboardingStatus(context.Background(), resourceURI)

	if err != nil {
		return nil, err
	}

	return &status, nil

}

// GetVMInsightsClient gets a diagnostics operations client
func GetVMInsightsClient(t testing.TestingT, subscriptionID string) (*insights.VMInsightsClient, error) {
	client := insights.NewVMInsightsClient(subscriptionID)

	authorizer, err := NewAuthorizer()

	if err != nil {
		return nil, err
	}

	client.Authorizer = *authorizer
	return &client, nil
}

// ActionGroupExists indicates whether the speficied Azure Availability Set exists
func ActionGroupExists(t testing.TestingT, actionGroupName string, resGroupName string, subscriptionID string) bool {
	exists, err := ActionGroupExistsE(t, actionGroupName, resGroupName, subscriptionID)
	require.NoError(t, err)
	return exists
}

// ActionGroupExistsE indicates whether the speficied Azure Availability Set exists
func ActionGroupExistsE(t testing.TestingT, actionGroupName string, resGroupName string, subscriptionID string) (bool, error) {
	_, err := GetActionGroupE(t, actionGroupName, resGroupName, subscriptionID)
	if err != nil {
		return false, err
	}
	return true, nil
}

// GetActionGroupE gets a Action Group in the specified Azure Resource Group
func GetActionGroupE(t testing.TestingT, actionGroupName string, resGroupName string, subscriptionID string) (*insights.ActionGroupResource, error) {
	// Validate resource group name and subscription ID
	_, err := getTargetAzureResourceGroupName(resGroupName)
	if err != nil {
		return nil, err
	}

	// Get the client refrence
	client, err := GetActionGroupsClientE(subscriptionID)
	if err != nil {
		return nil, err
	}

	// Get the Action Group
	actionGroup, err := client.Get(context.Background(), resGroupName, actionGroupName)
	if err != nil {
		return nil, err
	}

	return &actionGroup, nil
}

// GetActionGroupsClientE gets an Action Groups client in the specified Azure Subscription
func GetActionGroupsClientE(subscriptionID string) (*insights.ActionGroupsClient, error) {
	// Validate Azure subscription ID
	subscriptionID, err := getTargetAzureSubscription(subscriptionID)
	if err != nil {
		return nil, err
	}

	// Get the Action Groups client
	client := insights.NewActionGroupsClient(subscriptionID)

	// Create an authorizer
	authorizer, err := NewAuthorizer()
	if err != nil {
		return nil, err
	}
	client.Authorizer = *authorizer

	return &client, nil
}

// GetActivityLogAlertResourceE gets a Action Group in the specified Azure Resource Group
func GetActivityLogAlertResourceE(t testing.TestingT, activityLogAlertName string, resGroupName string, subscriptionID string) (*insights.ActivityLogAlertResource, error) {
	// Validate resource group name and subscription ID
	_, err := getTargetAzureResourceGroupName(resGroupName)
	if err != nil {
		return nil, err
	}

	// Get the client refrence
	client, err := GetActivityLogAlertsClientE(subscriptionID)
	if err != nil {
		return nil, err
	}

	// Get the Action Group
	activityLogAlertResource, err := client.Get(context.Background(), resGroupName, activityLogAlertName)
	if err != nil {
		return nil, err
	}

	return &activityLogAlertResource, nil
}

// GetActivityLogAlertsClientE gets an Action Groups client in the specified Azure Subscription
func GetActivityLogAlertsClientE(subscriptionID string) (*insights.ActivityLogAlertsClient, error) {
	// Validate Azure subscription ID
	subscriptionID, err := getTargetAzureSubscription(subscriptionID)
	if err != nil {
		return nil, err
	}

	// Get the Action Groups client
	client := insights.NewActivityLogAlertsClient(subscriptionID)

	// Create an authorizer
	authorizer, err := NewAuthorizer()
	if err != nil {
		return nil, err
	}
	client.Authorizer = *authorizer

	return &client, nil
}
