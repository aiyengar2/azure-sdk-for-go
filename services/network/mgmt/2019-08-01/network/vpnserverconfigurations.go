package network

// Copyright (c) Microsoft and contributors.  All rights reserved.
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
// http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
//
// See the License for the specific language governing permissions and
// limitations under the License.
//
// Code generated by Microsoft (R) AutoRest Code Generator.
// Changes may cause incorrect behavior and will be lost if the code is regenerated.

import (
	"context"
	"github.com/Azure/go-autorest/autorest"
	"github.com/Azure/go-autorest/autorest/azure"
	"github.com/Azure/go-autorest/tracing"
	"net/http"
)

// VpnServerConfigurationsClient is the network Client
type VpnServerConfigurationsClient struct {
	BaseClient
}

// NewVpnServerConfigurationsClient creates an instance of the VpnServerConfigurationsClient client.
func NewVpnServerConfigurationsClient(subscriptionID string) VpnServerConfigurationsClient {
	return NewVpnServerConfigurationsClientWithBaseURI(DefaultBaseURI, subscriptionID)
}

// NewVpnServerConfigurationsClientWithBaseURI creates an instance of the VpnServerConfigurationsClient client.
func NewVpnServerConfigurationsClientWithBaseURI(baseURI string, subscriptionID string) VpnServerConfigurationsClient {
	return VpnServerConfigurationsClient{NewWithBaseURI(baseURI, subscriptionID)}
}

// CreateOrUpdate creates a VpnServerConfiguration resource if it doesn't exist else updates the existing
// VpnServerConfiguration.
// Parameters:
// resourceGroupName - the resource group name of the VpnServerConfiguration.
// vpnServerConfigurationName - the name of the VpnServerConfiguration being created or updated.
// vpnServerConfigurationParameters - parameters supplied to create or update VpnServerConfiguration.
func (client VpnServerConfigurationsClient) CreateOrUpdate(ctx context.Context, resourceGroupName string, vpnServerConfigurationName string, vpnServerConfigurationParameters VpnServerConfiguration) (result VpnServerConfigurationsCreateOrUpdateFuture, err error) {
	if tracing.IsEnabled() {
		ctx = tracing.StartSpan(ctx, fqdn+"/VpnServerConfigurationsClient.CreateOrUpdate")
		defer func() {
			sc := -1
			if result.Response() != nil {
				sc = result.Response().StatusCode
			}
			tracing.EndSpan(ctx, sc, err)
		}()
	}
	req, err := client.CreateOrUpdatePreparer(ctx, resourceGroupName, vpnServerConfigurationName, vpnServerConfigurationParameters)
	if err != nil {
		err = autorest.NewErrorWithError(err, "network.VpnServerConfigurationsClient", "CreateOrUpdate", nil, "Failure preparing request")
		return
	}

	result, err = client.CreateOrUpdateSender(req)
	if err != nil {
		err = autorest.NewErrorWithError(err, "network.VpnServerConfigurationsClient", "CreateOrUpdate", result.Response(), "Failure sending request")
		return
	}

	return
}

// CreateOrUpdatePreparer prepares the CreateOrUpdate request.
func (client VpnServerConfigurationsClient) CreateOrUpdatePreparer(ctx context.Context, resourceGroupName string, vpnServerConfigurationName string, vpnServerConfigurationParameters VpnServerConfiguration) (*http.Request, error) {
	pathParameters := map[string]interface{}{
		"resourceGroupName":          autorest.Encode("path", resourceGroupName),
		"subscriptionId":             autorest.Encode("path", client.SubscriptionID),
		"vpnServerConfigurationName": autorest.Encode("path", vpnServerConfigurationName),
	}

	const APIVersion = "2019-08-01"
	queryParameters := map[string]interface{}{
		"api-version": APIVersion,
	}

	vpnServerConfigurationParameters.Etag = nil
	preparer := autorest.CreatePreparer(
		autorest.AsContentType("application/json; charset=utf-8"),
		autorest.AsPut(),
		autorest.WithBaseURL(client.BaseURI),
		autorest.WithPathParameters("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Network/vpnServerConfigurations/{vpnServerConfigurationName}", pathParameters),
		autorest.WithJSON(vpnServerConfigurationParameters),
		autorest.WithQueryParameters(queryParameters))
	return preparer.Prepare((&http.Request{}).WithContext(ctx))
}

// CreateOrUpdateSender sends the CreateOrUpdate request. The method will close the
// http.Response Body if it receives an error.
func (client VpnServerConfigurationsClient) CreateOrUpdateSender(req *http.Request) (future VpnServerConfigurationsCreateOrUpdateFuture, err error) {
	sd := autorest.GetSendDecorators(req.Context(), azure.DoRetryWithRegistration(client.Client))
	var resp *http.Response
	resp, err = autorest.SendWithSender(client, req, sd...)
	if err != nil {
		return
	}
	future.Future, err = azure.NewFutureFromResponse(resp)
	return
}

// CreateOrUpdateResponder handles the response to the CreateOrUpdate request. The method always
// closes the http.Response Body.
func (client VpnServerConfigurationsClient) CreateOrUpdateResponder(resp *http.Response) (result VpnServerConfiguration, err error) {
	err = autorest.Respond(
		resp,
		client.ByInspecting(),
		azure.WithErrorUnlessStatusCode(http.StatusOK, http.StatusCreated),
		autorest.ByUnmarshallingJSON(&result),
		autorest.ByClosing())
	result.Response = autorest.Response{Response: resp}
	return
}

// Delete deletes a VpnServerConfiguration.
// Parameters:
// resourceGroupName - the resource group name of the VpnServerConfiguration.
// vpnServerConfigurationName - the name of the VpnServerConfiguration being deleted.
func (client VpnServerConfigurationsClient) Delete(ctx context.Context, resourceGroupName string, vpnServerConfigurationName string) (result VpnServerConfigurationsDeleteFuture, err error) {
	if tracing.IsEnabled() {
		ctx = tracing.StartSpan(ctx, fqdn+"/VpnServerConfigurationsClient.Delete")
		defer func() {
			sc := -1
			if result.Response() != nil {
				sc = result.Response().StatusCode
			}
			tracing.EndSpan(ctx, sc, err)
		}()
	}
	req, err := client.DeletePreparer(ctx, resourceGroupName, vpnServerConfigurationName)
	if err != nil {
		err = autorest.NewErrorWithError(err, "network.VpnServerConfigurationsClient", "Delete", nil, "Failure preparing request")
		return
	}

	result, err = client.DeleteSender(req)
	if err != nil {
		err = autorest.NewErrorWithError(err, "network.VpnServerConfigurationsClient", "Delete", result.Response(), "Failure sending request")
		return
	}

	return
}

// DeletePreparer prepares the Delete request.
func (client VpnServerConfigurationsClient) DeletePreparer(ctx context.Context, resourceGroupName string, vpnServerConfigurationName string) (*http.Request, error) {
	pathParameters := map[string]interface{}{
		"resourceGroupName":          autorest.Encode("path", resourceGroupName),
		"subscriptionId":             autorest.Encode("path", client.SubscriptionID),
		"vpnServerConfigurationName": autorest.Encode("path", vpnServerConfigurationName),
	}

	const APIVersion = "2019-08-01"
	queryParameters := map[string]interface{}{
		"api-version": APIVersion,
	}

	preparer := autorest.CreatePreparer(
		autorest.AsDelete(),
		autorest.WithBaseURL(client.BaseURI),
		autorest.WithPathParameters("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Network/vpnServerConfigurations/{vpnServerConfigurationName}", pathParameters),
		autorest.WithQueryParameters(queryParameters))
	return preparer.Prepare((&http.Request{}).WithContext(ctx))
}

// DeleteSender sends the Delete request. The method will close the
// http.Response Body if it receives an error.
func (client VpnServerConfigurationsClient) DeleteSender(req *http.Request) (future VpnServerConfigurationsDeleteFuture, err error) {
	sd := autorest.GetSendDecorators(req.Context(), azure.DoRetryWithRegistration(client.Client))
	var resp *http.Response
	resp, err = autorest.SendWithSender(client, req, sd...)
	if err != nil {
		return
	}
	future.Future, err = azure.NewFutureFromResponse(resp)
	return
}

// DeleteResponder handles the response to the Delete request. The method always
// closes the http.Response Body.
func (client VpnServerConfigurationsClient) DeleteResponder(resp *http.Response) (result autorest.Response, err error) {
	err = autorest.Respond(
		resp,
		client.ByInspecting(),
		azure.WithErrorUnlessStatusCode(http.StatusOK, http.StatusAccepted, http.StatusNoContent),
		autorest.ByClosing())
	result.Response = resp
	return
}

// Get retrieves the details of a VpnServerConfiguration.
// Parameters:
// resourceGroupName - the resource group name of the VpnServerConfiguration.
// vpnServerConfigurationName - the name of the VpnServerConfiguration being retrieved.
func (client VpnServerConfigurationsClient) Get(ctx context.Context, resourceGroupName string, vpnServerConfigurationName string) (result VpnServerConfiguration, err error) {
	if tracing.IsEnabled() {
		ctx = tracing.StartSpan(ctx, fqdn+"/VpnServerConfigurationsClient.Get")
		defer func() {
			sc := -1
			if result.Response.Response != nil {
				sc = result.Response.Response.StatusCode
			}
			tracing.EndSpan(ctx, sc, err)
		}()
	}
	req, err := client.GetPreparer(ctx, resourceGroupName, vpnServerConfigurationName)
	if err != nil {
		err = autorest.NewErrorWithError(err, "network.VpnServerConfigurationsClient", "Get", nil, "Failure preparing request")
		return
	}

	resp, err := client.GetSender(req)
	if err != nil {
		result.Response = autorest.Response{Response: resp}
		err = autorest.NewErrorWithError(err, "network.VpnServerConfigurationsClient", "Get", resp, "Failure sending request")
		return
	}

	result, err = client.GetResponder(resp)
	if err != nil {
		err = autorest.NewErrorWithError(err, "network.VpnServerConfigurationsClient", "Get", resp, "Failure responding to request")
	}

	return
}

// GetPreparer prepares the Get request.
func (client VpnServerConfigurationsClient) GetPreparer(ctx context.Context, resourceGroupName string, vpnServerConfigurationName string) (*http.Request, error) {
	pathParameters := map[string]interface{}{
		"resourceGroupName":          autorest.Encode("path", resourceGroupName),
		"subscriptionId":             autorest.Encode("path", client.SubscriptionID),
		"vpnServerConfigurationName": autorest.Encode("path", vpnServerConfigurationName),
	}

	const APIVersion = "2019-08-01"
	queryParameters := map[string]interface{}{
		"api-version": APIVersion,
	}

	preparer := autorest.CreatePreparer(
		autorest.AsGet(),
		autorest.WithBaseURL(client.BaseURI),
		autorest.WithPathParameters("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Network/vpnServerConfigurations/{vpnServerConfigurationName}", pathParameters),
		autorest.WithQueryParameters(queryParameters))
	return preparer.Prepare((&http.Request{}).WithContext(ctx))
}

// GetSender sends the Get request. The method will close the
// http.Response Body if it receives an error.
func (client VpnServerConfigurationsClient) GetSender(req *http.Request) (*http.Response, error) {
	sd := autorest.GetSendDecorators(req.Context(), azure.DoRetryWithRegistration(client.Client))
	return autorest.SendWithSender(client, req, sd...)
}

// GetResponder handles the response to the Get request. The method always
// closes the http.Response Body.
func (client VpnServerConfigurationsClient) GetResponder(resp *http.Response) (result VpnServerConfiguration, err error) {
	err = autorest.Respond(
		resp,
		client.ByInspecting(),
		azure.WithErrorUnlessStatusCode(http.StatusOK),
		autorest.ByUnmarshallingJSON(&result),
		autorest.ByClosing())
	result.Response = autorest.Response{Response: resp}
	return
}

// List lists all the VpnServerConfigurations in a subscription.
func (client VpnServerConfigurationsClient) List(ctx context.Context) (result ListVpnServerConfigurationsResultPage, err error) {
	if tracing.IsEnabled() {
		ctx = tracing.StartSpan(ctx, fqdn+"/VpnServerConfigurationsClient.List")
		defer func() {
			sc := -1
			if result.lvscr.Response.Response != nil {
				sc = result.lvscr.Response.Response.StatusCode
			}
			tracing.EndSpan(ctx, sc, err)
		}()
	}
	result.fn = client.listNextResults
	req, err := client.ListPreparer(ctx)
	if err != nil {
		err = autorest.NewErrorWithError(err, "network.VpnServerConfigurationsClient", "List", nil, "Failure preparing request")
		return
	}

	resp, err := client.ListSender(req)
	if err != nil {
		result.lvscr.Response = autorest.Response{Response: resp}
		err = autorest.NewErrorWithError(err, "network.VpnServerConfigurationsClient", "List", resp, "Failure sending request")
		return
	}

	result.lvscr, err = client.ListResponder(resp)
	if err != nil {
		err = autorest.NewErrorWithError(err, "network.VpnServerConfigurationsClient", "List", resp, "Failure responding to request")
	}

	return
}

// ListPreparer prepares the List request.
func (client VpnServerConfigurationsClient) ListPreparer(ctx context.Context) (*http.Request, error) {
	pathParameters := map[string]interface{}{
		"subscriptionId": autorest.Encode("path", client.SubscriptionID),
	}

	const APIVersion = "2019-08-01"
	queryParameters := map[string]interface{}{
		"api-version": APIVersion,
	}

	preparer := autorest.CreatePreparer(
		autorest.AsGet(),
		autorest.WithBaseURL(client.BaseURI),
		autorest.WithPathParameters("/subscriptions/{subscriptionId}/providers/Microsoft.Network/vpnServerConfigurations", pathParameters),
		autorest.WithQueryParameters(queryParameters))
	return preparer.Prepare((&http.Request{}).WithContext(ctx))
}

// ListSender sends the List request. The method will close the
// http.Response Body if it receives an error.
func (client VpnServerConfigurationsClient) ListSender(req *http.Request) (*http.Response, error) {
	sd := autorest.GetSendDecorators(req.Context(), azure.DoRetryWithRegistration(client.Client))
	return autorest.SendWithSender(client, req, sd...)
}

// ListResponder handles the response to the List request. The method always
// closes the http.Response Body.
func (client VpnServerConfigurationsClient) ListResponder(resp *http.Response) (result ListVpnServerConfigurationsResult, err error) {
	err = autorest.Respond(
		resp,
		client.ByInspecting(),
		azure.WithErrorUnlessStatusCode(http.StatusOK),
		autorest.ByUnmarshallingJSON(&result),
		autorest.ByClosing())
	result.Response = autorest.Response{Response: resp}
	return
}

// listNextResults retrieves the next set of results, if any.
func (client VpnServerConfigurationsClient) listNextResults(ctx context.Context, lastResults ListVpnServerConfigurationsResult) (result ListVpnServerConfigurationsResult, err error) {
	req, err := lastResults.listVpnServerConfigurationsResultPreparer(ctx)
	if err != nil {
		return result, autorest.NewErrorWithError(err, "network.VpnServerConfigurationsClient", "listNextResults", nil, "Failure preparing next results request")
	}
	if req == nil {
		return
	}
	resp, err := client.ListSender(req)
	if err != nil {
		result.Response = autorest.Response{Response: resp}
		return result, autorest.NewErrorWithError(err, "network.VpnServerConfigurationsClient", "listNextResults", resp, "Failure sending next results request")
	}
	result, err = client.ListResponder(resp)
	if err != nil {
		err = autorest.NewErrorWithError(err, "network.VpnServerConfigurationsClient", "listNextResults", resp, "Failure responding to next results request")
	}
	return
}

// ListComplete enumerates all values, automatically crossing page boundaries as required.
func (client VpnServerConfigurationsClient) ListComplete(ctx context.Context) (result ListVpnServerConfigurationsResultIterator, err error) {
	if tracing.IsEnabled() {
		ctx = tracing.StartSpan(ctx, fqdn+"/VpnServerConfigurationsClient.List")
		defer func() {
			sc := -1
			if result.Response().Response.Response != nil {
				sc = result.page.Response().Response.Response.StatusCode
			}
			tracing.EndSpan(ctx, sc, err)
		}()
	}
	result.page, err = client.List(ctx)
	return
}

// ListByResourceGroup lists all the vpnServerConfigurations in a resource group.
// Parameters:
// resourceGroupName - the resource group name of the VpnServerConfiguration.
func (client VpnServerConfigurationsClient) ListByResourceGroup(ctx context.Context, resourceGroupName string) (result ListVpnServerConfigurationsResultPage, err error) {
	if tracing.IsEnabled() {
		ctx = tracing.StartSpan(ctx, fqdn+"/VpnServerConfigurationsClient.ListByResourceGroup")
		defer func() {
			sc := -1
			if result.lvscr.Response.Response != nil {
				sc = result.lvscr.Response.Response.StatusCode
			}
			tracing.EndSpan(ctx, sc, err)
		}()
	}
	result.fn = client.listByResourceGroupNextResults
	req, err := client.ListByResourceGroupPreparer(ctx, resourceGroupName)
	if err != nil {
		err = autorest.NewErrorWithError(err, "network.VpnServerConfigurationsClient", "ListByResourceGroup", nil, "Failure preparing request")
		return
	}

	resp, err := client.ListByResourceGroupSender(req)
	if err != nil {
		result.lvscr.Response = autorest.Response{Response: resp}
		err = autorest.NewErrorWithError(err, "network.VpnServerConfigurationsClient", "ListByResourceGroup", resp, "Failure sending request")
		return
	}

	result.lvscr, err = client.ListByResourceGroupResponder(resp)
	if err != nil {
		err = autorest.NewErrorWithError(err, "network.VpnServerConfigurationsClient", "ListByResourceGroup", resp, "Failure responding to request")
	}

	return
}

// ListByResourceGroupPreparer prepares the ListByResourceGroup request.
func (client VpnServerConfigurationsClient) ListByResourceGroupPreparer(ctx context.Context, resourceGroupName string) (*http.Request, error) {
	pathParameters := map[string]interface{}{
		"resourceGroupName": autorest.Encode("path", resourceGroupName),
		"subscriptionId":    autorest.Encode("path", client.SubscriptionID),
	}

	const APIVersion = "2019-08-01"
	queryParameters := map[string]interface{}{
		"api-version": APIVersion,
	}

	preparer := autorest.CreatePreparer(
		autorest.AsGet(),
		autorest.WithBaseURL(client.BaseURI),
		autorest.WithPathParameters("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Network/vpnServerConfigurations", pathParameters),
		autorest.WithQueryParameters(queryParameters))
	return preparer.Prepare((&http.Request{}).WithContext(ctx))
}

// ListByResourceGroupSender sends the ListByResourceGroup request. The method will close the
// http.Response Body if it receives an error.
func (client VpnServerConfigurationsClient) ListByResourceGroupSender(req *http.Request) (*http.Response, error) {
	sd := autorest.GetSendDecorators(req.Context(), azure.DoRetryWithRegistration(client.Client))
	return autorest.SendWithSender(client, req, sd...)
}

// ListByResourceGroupResponder handles the response to the ListByResourceGroup request. The method always
// closes the http.Response Body.
func (client VpnServerConfigurationsClient) ListByResourceGroupResponder(resp *http.Response) (result ListVpnServerConfigurationsResult, err error) {
	err = autorest.Respond(
		resp,
		client.ByInspecting(),
		azure.WithErrorUnlessStatusCode(http.StatusOK),
		autorest.ByUnmarshallingJSON(&result),
		autorest.ByClosing())
	result.Response = autorest.Response{Response: resp}
	return
}

// listByResourceGroupNextResults retrieves the next set of results, if any.
func (client VpnServerConfigurationsClient) listByResourceGroupNextResults(ctx context.Context, lastResults ListVpnServerConfigurationsResult) (result ListVpnServerConfigurationsResult, err error) {
	req, err := lastResults.listVpnServerConfigurationsResultPreparer(ctx)
	if err != nil {
		return result, autorest.NewErrorWithError(err, "network.VpnServerConfigurationsClient", "listByResourceGroupNextResults", nil, "Failure preparing next results request")
	}
	if req == nil {
		return
	}
	resp, err := client.ListByResourceGroupSender(req)
	if err != nil {
		result.Response = autorest.Response{Response: resp}
		return result, autorest.NewErrorWithError(err, "network.VpnServerConfigurationsClient", "listByResourceGroupNextResults", resp, "Failure sending next results request")
	}
	result, err = client.ListByResourceGroupResponder(resp)
	if err != nil {
		err = autorest.NewErrorWithError(err, "network.VpnServerConfigurationsClient", "listByResourceGroupNextResults", resp, "Failure responding to next results request")
	}
	return
}

// ListByResourceGroupComplete enumerates all values, automatically crossing page boundaries as required.
func (client VpnServerConfigurationsClient) ListByResourceGroupComplete(ctx context.Context, resourceGroupName string) (result ListVpnServerConfigurationsResultIterator, err error) {
	if tracing.IsEnabled() {
		ctx = tracing.StartSpan(ctx, fqdn+"/VpnServerConfigurationsClient.ListByResourceGroup")
		defer func() {
			sc := -1
			if result.Response().Response.Response != nil {
				sc = result.page.Response().Response.Response.StatusCode
			}
			tracing.EndSpan(ctx, sc, err)
		}()
	}
	result.page, err = client.ListByResourceGroup(ctx, resourceGroupName)
	return
}

// UpdateTags updates VpnServerConfiguration tags.
// Parameters:
// resourceGroupName - the resource group name of the VpnServerConfiguration.
// vpnServerConfigurationName - the name of the VpnServerConfiguration being updated.
// vpnServerConfigurationParameters - parameters supplied to update VpnServerConfiguration tags.
func (client VpnServerConfigurationsClient) UpdateTags(ctx context.Context, resourceGroupName string, vpnServerConfigurationName string, vpnServerConfigurationParameters TagsObject) (result VpnServerConfigurationsUpdateTagsFuture, err error) {
	if tracing.IsEnabled() {
		ctx = tracing.StartSpan(ctx, fqdn+"/VpnServerConfigurationsClient.UpdateTags")
		defer func() {
			sc := -1
			if result.Response() != nil {
				sc = result.Response().StatusCode
			}
			tracing.EndSpan(ctx, sc, err)
		}()
	}
	req, err := client.UpdateTagsPreparer(ctx, resourceGroupName, vpnServerConfigurationName, vpnServerConfigurationParameters)
	if err != nil {
		err = autorest.NewErrorWithError(err, "network.VpnServerConfigurationsClient", "UpdateTags", nil, "Failure preparing request")
		return
	}

	result, err = client.UpdateTagsSender(req)
	if err != nil {
		err = autorest.NewErrorWithError(err, "network.VpnServerConfigurationsClient", "UpdateTags", result.Response(), "Failure sending request")
		return
	}

	return
}

// UpdateTagsPreparer prepares the UpdateTags request.
func (client VpnServerConfigurationsClient) UpdateTagsPreparer(ctx context.Context, resourceGroupName string, vpnServerConfigurationName string, vpnServerConfigurationParameters TagsObject) (*http.Request, error) {
	pathParameters := map[string]interface{}{
		"resourceGroupName":          autorest.Encode("path", resourceGroupName),
		"subscriptionId":             autorest.Encode("path", client.SubscriptionID),
		"vpnServerConfigurationName": autorest.Encode("path", vpnServerConfigurationName),
	}

	const APIVersion = "2019-08-01"
	queryParameters := map[string]interface{}{
		"api-version": APIVersion,
	}

	preparer := autorest.CreatePreparer(
		autorest.AsContentType("application/json; charset=utf-8"),
		autorest.AsPatch(),
		autorest.WithBaseURL(client.BaseURI),
		autorest.WithPathParameters("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Network/vpnServerConfigurations/{vpnServerConfigurationName}", pathParameters),
		autorest.WithJSON(vpnServerConfigurationParameters),
		autorest.WithQueryParameters(queryParameters))
	return preparer.Prepare((&http.Request{}).WithContext(ctx))
}

// UpdateTagsSender sends the UpdateTags request. The method will close the
// http.Response Body if it receives an error.
func (client VpnServerConfigurationsClient) UpdateTagsSender(req *http.Request) (future VpnServerConfigurationsUpdateTagsFuture, err error) {
	sd := autorest.GetSendDecorators(req.Context(), azure.DoRetryWithRegistration(client.Client))
	var resp *http.Response
	resp, err = autorest.SendWithSender(client, req, sd...)
	if err != nil {
		return
	}
	future.Future, err = azure.NewFutureFromResponse(resp)
	return
}

// UpdateTagsResponder handles the response to the UpdateTags request. The method always
// closes the http.Response Body.
func (client VpnServerConfigurationsClient) UpdateTagsResponder(resp *http.Response) (result VpnServerConfiguration, err error) {
	err = autorest.Respond(
		resp,
		client.ByInspecting(),
		azure.WithErrorUnlessStatusCode(http.StatusOK, http.StatusCreated),
		autorest.ByUnmarshallingJSON(&result),
		autorest.ByClosing())
	result.Response = autorest.Response{Response: resp}
	return
}