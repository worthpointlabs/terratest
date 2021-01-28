package azure

import (
	"context"
	"testing"

	"github.com/Azure/azure-sdk-for-go/services/servicebus/mgmt/2017-04-01/servicebus"
	"github.com/stretchr/testify/require"
)

func serviceBusNamespaceClientE(subscriptionID string) (*servicebus.NamespacesClient, error) {
	authorizer, err := NewAuthorizer()
	if err != nil {
		return nil, err
	}

	nsClient := servicebus.NewNamespacesClient(subscriptionID)
	nsClient.Authorizer = *authorizer
	return &nsClient, nil
}

func serviceBusTopicClientE(subscriptionID string) (*servicebus.TopicsClient, error) {
	authorizer, err := NewAuthorizer()
	if err != nil {
		return nil, err
	}

	tClient := servicebus.NewTopicsClient(subscriptionID)
	tClient.Authorizer = *authorizer
	return &tClient, nil
}

func serviceBusSubscriptionsClientE(subscriptionID string) (*servicebus.SubscriptionsClient, error) {
	authorizer, err := NewAuthorizer()
	if err != nil {
		return nil, err
	}

	sClient := servicebus.NewSubscriptionsClient(subscriptionID)
	sClient.Authorizer = *authorizer
	return &sClient, nil
}

// ListServiceBusNamespaceE list all SB namespaces in all resource groups in the given subscription ID. This function would fail the test if there is an error.
func ListServiceBusNamespaceE(subscriptionID string) (*[]servicebus.SBNamespace, error) {
	nsClient, err := serviceBusNamespaceClientE(subscriptionID)
	if err != nil {
		return nil, err
	}

	iteratorSBNamespace, err := nsClient.ListComplete(context.Background())
	if err != nil {
		return nil, err
	}

	results := make([]servicebus.SBNamespace, 0)
	for iteratorSBNamespace.NotDone() {
		results = append(results, iteratorSBNamespace.Value())
		err = iteratorSBNamespace.Next()
		if err != nil {
			return nil, err
		}
	}

	return &results, nil
}

// ListServiceBusNamespace - list all SB namespaces in all resource groups in the given subscription ID.
func ListServiceBusNamespace(t *testing.T, subscriptionID string) *[]servicebus.SBNamespace {
	results, err := ListServiceBusNamespaceE(subscriptionID)
	
	require.NoError(t, err)

	return results
}

// ListServiceBusNamespaceNamesE list names of all SB namespaces in all resource groups in the given subscription ID. This function would fail the test if there is an error.
func ListServiceBusNamespaceNamesE(subscriptionID string) (*[]string, error) {
	sbNamespace, err := ListServiceBusNamespaceE(subscriptionID)

	if err != nil {
		return nil, err
	}

	results := make([]string, 0)
	for _, namespace := range *sbNamespace {
		results = append(results, *namespace.Name)
		if err != nil {
			return nil, err
		}
	}

	return &results, nil
}

// ListServiceBusNamespaceNames list names of all SB namespaces in all resource groups in the given subscription ID.
func ListServiceBusNamespaceNames(t *testing.T, subscriptionID string) *[]string {
	results, err := ListServiceBusNamespaceNamesE(subscriptionID)
	
	require.NoError(t, err)

	return results
}

// ListServiceBusNamespaceIDsE list IDs of all SB namespaces in all resource groups in the given subscription ID. This function would fail the test if there is an error.
func ListServiceBusNamespaceIDsE(subscriptionID string) (*[]string, error) {
	sbNamespace, err := ListServiceBusNamespaceE(subscriptionID)

	if err != nil {
		return nil, err
	}

	results := make([]string, 0)
	for _, namespace := range *sbNamespace {
		results = append(results, *namespace.ID)
		if err != nil {
			return nil, err
		}
	}

	return &results, nil
}

// ListServiceBusNamespaceIDs list IDs of all SB namespaces in all resource groups in the given subscription ID.
func ListServiceBusNamespaceIDs(t *testing.T, subscriptionID string) *[]string {
	results, err := ListServiceBusNamespaceIDsE(subscriptionID)
	require.NoError(t, err)

	return results
}

// ListServiceBusNamespaceByResourceGroupE list all SB namespaces in the given resource group. This function would fail the test if there is an error.
func ListServiceBusNamespaceByResourceGroupE(subscriptionID string, resourceGroup string) (*[]servicebus.SBNamespace, error) {
	nsClient, err := serviceBusNamespaceClientE(subscriptionID)
	if err != nil {
		return nil, err
	}

	iteratorSBNamespace, err := nsClient.ListByResourceGroupComplete(context.Background(), resourceGroup)
	if err != nil {
		return nil, err
	}

	results := make([]servicebus.SBNamespace, 0)

	for iteratorSBNamespace.NotDone() {
		results = append(results, iteratorSBNamespace.Value())
		err = iteratorSBNamespace.Next()
		if err != nil {
			return nil, err
		}
	}

	return &results, nil
}

// ListServiceBusNamespaceByResourceGroup list all SB namespaces in the given resource group.
func ListServiceBusNamespaceByResourceGroup(t *testing.T, subscriptionID string, resourceGroup string) *[]servicebus.SBNamespace {
	results, err := ListServiceBusNamespaceByResourceGroupE(subscriptionID, resourceGroup)
	require.NoError(t, err)

	return results
}

// ListServiceBusNamespaceNamesByResourceGroupE list names of all SB namespaces in the given resource group. This function would fail the test if there is an error.
func ListServiceBusNamespaceNamesByResourceGroupE(subscriptionID string, resourceGroup string) (*[]string, error) {
	sbNamespace, err := ListServiceBusNamespaceByResourceGroupE(subscriptionID, resourceGroup)

	if err != nil {
		return nil, err
	}

	results := make([]string, 0)
	for _, namespace := range *sbNamespace {
		results = append(results, *namespace.Name)
		if err != nil {
			return nil, err
		}
	}

	return &results, nil
}

// ListServiceBusNamespaceNamesByResourceGroup list names of all SB namespaces in the given resource group.
func ListServiceBusNamespaceNamesByResourceGroup(t *testing.T, subscriptionID string, resourceGroup string) *[]string {
	results, err := ListServiceBusNamespaceNamesByResourceGroupE(subscriptionID, resourceGroup)
	require.NoError(t, err)

	return results
}

// ListServiceBusNamespaceIDsByResourceGroupE list IDs of all SB namespaces in the given resource group. This function would fail the test if there is an error.
func ListServiceBusNamespaceIDsByResourceGroupE(subscriptionID string, resourceGroup string) (*[]string, error) {
	sbNamespace, err := ListServiceBusNamespaceByResourceGroupE(subscriptionID, resourceGroup)

	if err != nil {
		return nil, err
	}

	results := make([]string, 0)
	for _, namespace := range *sbNamespace {
		results = append(results, *namespace.ID)
		if err != nil {
			return nil, err
		}
	}

	return &results, nil
}

// ListServiceBusNamespaceIDsByResourceGroup list IDs of all SB namespaces in the given resource group.
func ListServiceBusNamespaceIDsByResourceGroup(t *testing.T, subscriptionID string, resourceGroup string) *[]string {
	results, err := ListServiceBusNamespaceIDsByResourceGroupE(subscriptionID, resourceGroup)
	require.NoError(t, err)

	return results
}

// ListNamespaceAuthRulesE - authenticate namespace client and enumerates all values to get list of authorization rules for the given namespace name,
// automatically crossing page boundaries as required. This function would fail the test if there is an error.
func ListNamespaceAuthRulesE(subscriptionID string, namespace string, resourceGroup string) (*[]string, error) {
	nsClient, err := serviceBusNamespaceClientE(subscriptionID)
	if err != nil {
		return nil, err
	}
	iteratorNamespaceRules, err := nsClient.ListAuthorizationRulesComplete(
		context.Background(), resourceGroup, namespace)

	if err != nil {
		return nil, err
	}

	results := make([]string, 0)
	for iteratorNamespaceRules.NotDone() {
		results = append(results, *(iteratorNamespaceRules.Value()).Name)
		err = iteratorNamespaceRules.Next()
		if err != nil {
			return nil, err
		}
	}
	return &results, nil
}

// ListNamespaceAuthRules - authenticate namespace client and enumerates all values to get list of authorization rules for the given namespace name,
// automatically crossing page boundaries as required.
func ListNamespaceAuthRules(t *testing.T, subscriptionID string, namespace string, resourceGroup string) *[]string {
	results, err := ListNamespaceAuthRulesE(subscriptionID, namespace, resourceGroup)
	require.NoError(t, err)

	return results
}

// ListNamespaceTopicsE - authenticate topic client and enumerates all values, automatically crossing page boundaries as required. This function would fail the test if there is an error.
func ListNamespaceTopicsE(subscriptionID string, namespace string, resourceGroup string) (*[]servicebus.SBTopic, error) {
	tClient, err := serviceBusTopicClientE(subscriptionID)
	if err != nil {
		return nil, err
	}

	iteratorTopics, err := tClient.ListByNamespaceComplete(context.Background(), resourceGroup, namespace, nil, nil)
	if err != nil {
		return nil, err
	}

	results := make([]servicebus.SBTopic, 0)

	for iteratorTopics.NotDone() {
		results = append(results, iteratorTopics.Value())
		err = iteratorTopics.Next()
		if err != nil {
			return nil, err
		}
	}

	return &results, nil
}

// ListNamespaceTopics - authenticate topic client and enumerates all values, automatically crossing page boundaries as required.
func ListNamespaceTopics(t *testing.T, subscriptionID string, namespace string, resourceGroup string) *[]servicebus.SBTopic {
	results, err := ListNamespaceTopicsE(subscriptionID, namespace, resourceGroup)
	require.NoError(t, err)

	return results
}

// ListTopicSubscriptionsE - authenticate subscriptions client and enumerates all values, automatically crossing page boundaries as required. This function would fail the test if there is an error.
func ListTopicSubscriptionsE(subscriptionID string, namespace string, resourceGroup string, topicName string) ([]servicebus.SBSubscription, error) {
	sClient, err := serviceBusSubscriptionsClientE(subscriptionID)
	if err != nil {
		return nil, err
	}
	iteratorSubscription, err := sClient.ListByTopicComplete(context.Background(), resourceGroup, namespace, topicName, nil, nil)

	if err != nil {
		return nil, err
	}

	results := make([]servicebus.SBSubscription, 0)

	for iteratorSubscription.NotDone() {
		results = append(results, iteratorSubscription.Value())
		err = iteratorSubscription.Next()
		if err != nil {
			return nil, err
		}
	}
	return results, nil
}

// ListTopicSubscriptions - authenticate subscriptions client and enumerates all values, automatically crossing page boundaries as required.
func ListTopicSubscriptions(t *testing.T, subscriptionID string, namespace string, resourceGroup string, topicName string) *[]servicebus.SBSubscription {
	results, err := ListTopicSubscriptionsE(subscriptionID, namespace, resourceGroup, topicName)
	require.NoError(t, err)

	return &results
}

// ListTopicSubscriptionsNameE - authenticate subscriptions client and enumerates all values to get list of subscriptions for the given topic name,
// automatically crossing page boundaries as required. This function would fail the test if there is an error.
func ListTopicSubscriptionsNameE(subscriptionID string, namespace string, resourceGroup string, topicName string) (*[]string, error) {
	sClient, err := serviceBusSubscriptionsClientE(subscriptionID)
	if err != nil {
		return nil, err
	}
	iteratorSubscription, err := sClient.ListByTopicComplete(context.Background(), resourceGroup, namespace, topicName, nil, nil)

	if err != nil {
		return nil, err
	}

	results := make([]string, 0)
	for iteratorSubscription.NotDone() {
		results = append(results, *(iteratorSubscription.Value()).Name)
		err = iteratorSubscription.Next()
		if err != nil {
			return nil, err
		}
	}
	return &results, nil
}

// ListTopicSubscriptionsName -  authenticate subscriptions client and enumerates all values to get list of subscriptions for the given topic name,
// automatically crossing page boundaries as required.
func ListTopicSubscriptionsName(t *testing.T, subscriptionID string, namespace string, resourceGroup string, topicName string) *[]string {
	results, err := ListTopicSubscriptionsNameE(subscriptionID, namespace, resourceGroup, topicName)
	require.NoError(t, err)

	return results
}

// ListTopicAuthRulesE - authenticate topic client and enumerates all values to get list of authorization rules for the given topic name,
// automatically crossing page boundaries as required. This function would fail the test if there is an error.
func ListTopicAuthRulesE(subscriptionID string, namespace string, resourceGroup string, topicName string) (*[]string, error) {
	tClient, err := serviceBusTopicClientE(subscriptionID)
	if err != nil {
		return nil, err
	}
	iteratorTopicsRules, err := tClient.ListAuthorizationRulesComplete(
		context.Background(), resourceGroup, namespace, topicName)

	if err != nil {
		return nil, err
	}

	results := make([]string, 0)
	for iteratorTopicsRules.NotDone() {
		results = append(results, *(iteratorTopicsRules.Value()).Name)
		err = iteratorTopicsRules.Next()
		if err != nil {
			return nil, err
		}
	}
	return &results, nil
}

// ListTopicAuthRules - authenticate topic client and enumerates all values to get list of authorization rules for the given topic name,
// automatically crossing page boundaries as required.
func ListTopicAuthRules(t *testing.T, subscriptionID string, namespace string, resourceGroup string, topicName string) *[]string {
	results, err := ListTopicAuthRulesE(subscriptionID, namespace, resourceGroup, topicName)
	require.NoError(t, err)

	return results
}
