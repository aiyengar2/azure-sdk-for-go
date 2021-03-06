package servicefabricmesh

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
	"github.com/Azure/go-autorest/autorest/validation"
	"github.com/Azure/go-autorest/tracing"
	"net/http"
)

// SecretValueClient is the service Fabric Mesh Management Client
type SecretValueClient struct {
	BaseClient
}

// NewSecretValueClient creates an instance of the SecretValueClient client.
func NewSecretValueClient(subscriptionID string) SecretValueClient {
	return NewSecretValueClientWithBaseURI(DefaultBaseURI, subscriptionID)
}

// NewSecretValueClientWithBaseURI creates an instance of the SecretValueClient client using a custom endpoint.  Use
// this when interacting with an Azure cloud that uses a non-standard base URI (sovereign clouds, Azure stack).
func NewSecretValueClientWithBaseURI(baseURI string, subscriptionID string) SecretValueClient {
	return SecretValueClient{NewWithBaseURI(baseURI, subscriptionID)}
}

// Create creates a new value of the specified secret resource. The name of the value is typically the version
// identifier. Once created the value cannot be changed.
// Parameters:
// resourceGroupName - azure resource group name
// secretResourceName - the name of the secret resource.
// secretValueResourceName - the name of the secret resource value which is typically the version identifier
// for the value.
// secretValueResourceDescription - description for creating a value of a secret resource.
func (client SecretValueClient) Create(ctx context.Context, resourceGroupName string, secretResourceName string, secretValueResourceName string, secretValueResourceDescription SecretValueResourceDescription) (result SecretValueResourceDescription, err error) {
	if tracing.IsEnabled() {
		ctx = tracing.StartSpan(ctx, fqdn+"/SecretValueClient.Create")
		defer func() {
			sc := -1
			if result.Response.Response != nil {
				sc = result.Response.Response.StatusCode
			}
			tracing.EndSpan(ctx, sc, err)
		}()
	}
	if err := validation.Validate([]validation.Validation{
		{TargetValue: secretValueResourceDescription,
			Constraints: []validation.Constraint{{Target: "secretValueResourceDescription.SecretValueResourceProperties", Name: validation.Null, Rule: true, Chain: nil}}}}); err != nil {
		return result, validation.NewError("servicefabricmesh.SecretValueClient", "Create", err.Error())
	}

	req, err := client.CreatePreparer(ctx, resourceGroupName, secretResourceName, secretValueResourceName, secretValueResourceDescription)
	if err != nil {
		err = autorest.NewErrorWithError(err, "servicefabricmesh.SecretValueClient", "Create", nil, "Failure preparing request")
		return
	}

	resp, err := client.CreateSender(req)
	if err != nil {
		result.Response = autorest.Response{Response: resp}
		err = autorest.NewErrorWithError(err, "servicefabricmesh.SecretValueClient", "Create", resp, "Failure sending request")
		return
	}

	result, err = client.CreateResponder(resp)
	if err != nil {
		err = autorest.NewErrorWithError(err, "servicefabricmesh.SecretValueClient", "Create", resp, "Failure responding to request")
	}

	return
}

// CreatePreparer prepares the Create request.
func (client SecretValueClient) CreatePreparer(ctx context.Context, resourceGroupName string, secretResourceName string, secretValueResourceName string, secretValueResourceDescription SecretValueResourceDescription) (*http.Request, error) {
	pathParameters := map[string]interface{}{
		"resourceGroupName":       autorest.Encode("path", resourceGroupName),
		"secretResourceName":      secretResourceName,
		"secretValueResourceName": secretValueResourceName,
		"subscriptionId":          autorest.Encode("path", client.SubscriptionID),
	}

	const APIVersion = "2018-09-01-preview"
	queryParameters := map[string]interface{}{
		"api-version": APIVersion,
	}

	preparer := autorest.CreatePreparer(
		autorest.AsContentType("application/json; charset=utf-8"),
		autorest.AsPut(),
		autorest.WithBaseURL(client.BaseURI),
		autorest.WithPathParameters("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.ServiceFabricMesh/secrets/{secretResourceName}/values/{secretValueResourceName}", pathParameters),
		autorest.WithJSON(secretValueResourceDescription),
		autorest.WithQueryParameters(queryParameters))
	return preparer.Prepare((&http.Request{}).WithContext(ctx))
}

// CreateSender sends the Create request. The method will close the
// http.Response Body if it receives an error.
func (client SecretValueClient) CreateSender(req *http.Request) (*http.Response, error) {
	sd := autorest.GetSendDecorators(req.Context(), azure.DoRetryWithRegistration(client.Client))
	return autorest.SendWithSender(client, req, sd...)
}

// CreateResponder handles the response to the Create request. The method always
// closes the http.Response Body.
func (client SecretValueClient) CreateResponder(resp *http.Response) (result SecretValueResourceDescription, err error) {
	err = autorest.Respond(
		resp,
		client.ByInspecting(),
		azure.WithErrorUnlessStatusCode(http.StatusOK, http.StatusCreated, http.StatusAccepted),
		autorest.ByUnmarshallingJSON(&result),
		autorest.ByClosing())
	result.Response = autorest.Response{Response: resp}
	return
}

// Delete deletes the secret value resource identified by the name. The name of the resource is typically the version
// associated with that value. Deletion will fail if the specified value is in use.
// Parameters:
// resourceGroupName - azure resource group name
// secretResourceName - the name of the secret resource.
// secretValueResourceName - the name of the secret resource value which is typically the version identifier
// for the value.
func (client SecretValueClient) Delete(ctx context.Context, resourceGroupName string, secretResourceName string, secretValueResourceName string) (result autorest.Response, err error) {
	if tracing.IsEnabled() {
		ctx = tracing.StartSpan(ctx, fqdn+"/SecretValueClient.Delete")
		defer func() {
			sc := -1
			if result.Response != nil {
				sc = result.Response.StatusCode
			}
			tracing.EndSpan(ctx, sc, err)
		}()
	}
	req, err := client.DeletePreparer(ctx, resourceGroupName, secretResourceName, secretValueResourceName)
	if err != nil {
		err = autorest.NewErrorWithError(err, "servicefabricmesh.SecretValueClient", "Delete", nil, "Failure preparing request")
		return
	}

	resp, err := client.DeleteSender(req)
	if err != nil {
		result.Response = resp
		err = autorest.NewErrorWithError(err, "servicefabricmesh.SecretValueClient", "Delete", resp, "Failure sending request")
		return
	}

	result, err = client.DeleteResponder(resp)
	if err != nil {
		err = autorest.NewErrorWithError(err, "servicefabricmesh.SecretValueClient", "Delete", resp, "Failure responding to request")
	}

	return
}

// DeletePreparer prepares the Delete request.
func (client SecretValueClient) DeletePreparer(ctx context.Context, resourceGroupName string, secretResourceName string, secretValueResourceName string) (*http.Request, error) {
	pathParameters := map[string]interface{}{
		"resourceGroupName":       autorest.Encode("path", resourceGroupName),
		"secretResourceName":      secretResourceName,
		"secretValueResourceName": secretValueResourceName,
		"subscriptionId":          autorest.Encode("path", client.SubscriptionID),
	}

	const APIVersion = "2018-09-01-preview"
	queryParameters := map[string]interface{}{
		"api-version": APIVersion,
	}

	preparer := autorest.CreatePreparer(
		autorest.AsDelete(),
		autorest.WithBaseURL(client.BaseURI),
		autorest.WithPathParameters("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.ServiceFabricMesh/secrets/{secretResourceName}/values/{secretValueResourceName}", pathParameters),
		autorest.WithQueryParameters(queryParameters))
	return preparer.Prepare((&http.Request{}).WithContext(ctx))
}

// DeleteSender sends the Delete request. The method will close the
// http.Response Body if it receives an error.
func (client SecretValueClient) DeleteSender(req *http.Request) (*http.Response, error) {
	sd := autorest.GetSendDecorators(req.Context(), azure.DoRetryWithRegistration(client.Client))
	return autorest.SendWithSender(client, req, sd...)
}

// DeleteResponder handles the response to the Delete request. The method always
// closes the http.Response Body.
func (client SecretValueClient) DeleteResponder(resp *http.Response) (result autorest.Response, err error) {
	err = autorest.Respond(
		resp,
		client.ByInspecting(),
		azure.WithErrorUnlessStatusCode(http.StatusOK, http.StatusAccepted, http.StatusNoContent),
		autorest.ByClosing())
	result.Response = resp
	return
}

// Get get the information about the specified named secret value resources. The information does not include the
// actual value of the secret.
// Parameters:
// resourceGroupName - azure resource group name
// secretResourceName - the name of the secret resource.
// secretValueResourceName - the name of the secret resource value which is typically the version identifier
// for the value.
func (client SecretValueClient) Get(ctx context.Context, resourceGroupName string, secretResourceName string, secretValueResourceName string) (result SecretValueResourceDescription, err error) {
	if tracing.IsEnabled() {
		ctx = tracing.StartSpan(ctx, fqdn+"/SecretValueClient.Get")
		defer func() {
			sc := -1
			if result.Response.Response != nil {
				sc = result.Response.Response.StatusCode
			}
			tracing.EndSpan(ctx, sc, err)
		}()
	}
	req, err := client.GetPreparer(ctx, resourceGroupName, secretResourceName, secretValueResourceName)
	if err != nil {
		err = autorest.NewErrorWithError(err, "servicefabricmesh.SecretValueClient", "Get", nil, "Failure preparing request")
		return
	}

	resp, err := client.GetSender(req)
	if err != nil {
		result.Response = autorest.Response{Response: resp}
		err = autorest.NewErrorWithError(err, "servicefabricmesh.SecretValueClient", "Get", resp, "Failure sending request")
		return
	}

	result, err = client.GetResponder(resp)
	if err != nil {
		err = autorest.NewErrorWithError(err, "servicefabricmesh.SecretValueClient", "Get", resp, "Failure responding to request")
	}

	return
}

// GetPreparer prepares the Get request.
func (client SecretValueClient) GetPreparer(ctx context.Context, resourceGroupName string, secretResourceName string, secretValueResourceName string) (*http.Request, error) {
	pathParameters := map[string]interface{}{
		"resourceGroupName":       autorest.Encode("path", resourceGroupName),
		"secretResourceName":      secretResourceName,
		"secretValueResourceName": secretValueResourceName,
		"subscriptionId":          autorest.Encode("path", client.SubscriptionID),
	}

	const APIVersion = "2018-09-01-preview"
	queryParameters := map[string]interface{}{
		"api-version": APIVersion,
	}

	preparer := autorest.CreatePreparer(
		autorest.AsGet(),
		autorest.WithBaseURL(client.BaseURI),
		autorest.WithPathParameters("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.ServiceFabricMesh/secrets/{secretResourceName}/values/{secretValueResourceName}", pathParameters),
		autorest.WithQueryParameters(queryParameters))
	return preparer.Prepare((&http.Request{}).WithContext(ctx))
}

// GetSender sends the Get request. The method will close the
// http.Response Body if it receives an error.
func (client SecretValueClient) GetSender(req *http.Request) (*http.Response, error) {
	sd := autorest.GetSendDecorators(req.Context(), azure.DoRetryWithRegistration(client.Client))
	return autorest.SendWithSender(client, req, sd...)
}

// GetResponder handles the response to the Get request. The method always
// closes the http.Response Body.
func (client SecretValueClient) GetResponder(resp *http.Response) (result SecretValueResourceDescription, err error) {
	err = autorest.Respond(
		resp,
		client.ByInspecting(),
		azure.WithErrorUnlessStatusCode(http.StatusOK),
		autorest.ByUnmarshallingJSON(&result),
		autorest.ByClosing())
	result.Response = autorest.Response{Response: resp}
	return
}

// List gets information about all secret value resources of the specified secret resource. The information includes
// the names of the secret value resources, but not the actual values.
// Parameters:
// resourceGroupName - azure resource group name
// secretResourceName - the name of the secret resource.
func (client SecretValueClient) List(ctx context.Context, resourceGroupName string, secretResourceName string) (result SecretValueResourceDescriptionListPage, err error) {
	if tracing.IsEnabled() {
		ctx = tracing.StartSpan(ctx, fqdn+"/SecretValueClient.List")
		defer func() {
			sc := -1
			if result.svrdl.Response.Response != nil {
				sc = result.svrdl.Response.Response.StatusCode
			}
			tracing.EndSpan(ctx, sc, err)
		}()
	}
	result.fn = client.listNextResults
	req, err := client.ListPreparer(ctx, resourceGroupName, secretResourceName)
	if err != nil {
		err = autorest.NewErrorWithError(err, "servicefabricmesh.SecretValueClient", "List", nil, "Failure preparing request")
		return
	}

	resp, err := client.ListSender(req)
	if err != nil {
		result.svrdl.Response = autorest.Response{Response: resp}
		err = autorest.NewErrorWithError(err, "servicefabricmesh.SecretValueClient", "List", resp, "Failure sending request")
		return
	}

	result.svrdl, err = client.ListResponder(resp)
	if err != nil {
		err = autorest.NewErrorWithError(err, "servicefabricmesh.SecretValueClient", "List", resp, "Failure responding to request")
	}

	return
}

// ListPreparer prepares the List request.
func (client SecretValueClient) ListPreparer(ctx context.Context, resourceGroupName string, secretResourceName string) (*http.Request, error) {
	pathParameters := map[string]interface{}{
		"resourceGroupName":  autorest.Encode("path", resourceGroupName),
		"secretResourceName": secretResourceName,
		"subscriptionId":     autorest.Encode("path", client.SubscriptionID),
	}

	const APIVersion = "2018-09-01-preview"
	queryParameters := map[string]interface{}{
		"api-version": APIVersion,
	}

	preparer := autorest.CreatePreparer(
		autorest.AsGet(),
		autorest.WithBaseURL(client.BaseURI),
		autorest.WithPathParameters("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.ServiceFabricMesh/secrets/{secretResourceName}/values", pathParameters),
		autorest.WithQueryParameters(queryParameters))
	return preparer.Prepare((&http.Request{}).WithContext(ctx))
}

// ListSender sends the List request. The method will close the
// http.Response Body if it receives an error.
func (client SecretValueClient) ListSender(req *http.Request) (*http.Response, error) {
	sd := autorest.GetSendDecorators(req.Context(), azure.DoRetryWithRegistration(client.Client))
	return autorest.SendWithSender(client, req, sd...)
}

// ListResponder handles the response to the List request. The method always
// closes the http.Response Body.
func (client SecretValueClient) ListResponder(resp *http.Response) (result SecretValueResourceDescriptionList, err error) {
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
func (client SecretValueClient) listNextResults(ctx context.Context, lastResults SecretValueResourceDescriptionList) (result SecretValueResourceDescriptionList, err error) {
	req, err := lastResults.secretValueResourceDescriptionListPreparer(ctx)
	if err != nil {
		return result, autorest.NewErrorWithError(err, "servicefabricmesh.SecretValueClient", "listNextResults", nil, "Failure preparing next results request")
	}
	if req == nil {
		return
	}
	resp, err := client.ListSender(req)
	if err != nil {
		result.Response = autorest.Response{Response: resp}
		return result, autorest.NewErrorWithError(err, "servicefabricmesh.SecretValueClient", "listNextResults", resp, "Failure sending next results request")
	}
	result, err = client.ListResponder(resp)
	if err != nil {
		err = autorest.NewErrorWithError(err, "servicefabricmesh.SecretValueClient", "listNextResults", resp, "Failure responding to next results request")
	}
	return
}

// ListComplete enumerates all values, automatically crossing page boundaries as required.
func (client SecretValueClient) ListComplete(ctx context.Context, resourceGroupName string, secretResourceName string) (result SecretValueResourceDescriptionListIterator, err error) {
	if tracing.IsEnabled() {
		ctx = tracing.StartSpan(ctx, fqdn+"/SecretValueClient.List")
		defer func() {
			sc := -1
			if result.Response().Response.Response != nil {
				sc = result.page.Response().Response.Response.StatusCode
			}
			tracing.EndSpan(ctx, sc, err)
		}()
	}
	result.page, err = client.List(ctx, resourceGroupName, secretResourceName)
	return
}

// ListValue lists the decrypted value of the specified named value of the secret resource. This is a privileged
// operation.
// Parameters:
// resourceGroupName - azure resource group name
// secretResourceName - the name of the secret resource.
// secretValueResourceName - the name of the secret resource value which is typically the version identifier
// for the value.
func (client SecretValueClient) ListValue(ctx context.Context, resourceGroupName string, secretResourceName string, secretValueResourceName string) (result SecretValue, err error) {
	if tracing.IsEnabled() {
		ctx = tracing.StartSpan(ctx, fqdn+"/SecretValueClient.ListValue")
		defer func() {
			sc := -1
			if result.Response.Response != nil {
				sc = result.Response.Response.StatusCode
			}
			tracing.EndSpan(ctx, sc, err)
		}()
	}
	req, err := client.ListValuePreparer(ctx, resourceGroupName, secretResourceName, secretValueResourceName)
	if err != nil {
		err = autorest.NewErrorWithError(err, "servicefabricmesh.SecretValueClient", "ListValue", nil, "Failure preparing request")
		return
	}

	resp, err := client.ListValueSender(req)
	if err != nil {
		result.Response = autorest.Response{Response: resp}
		err = autorest.NewErrorWithError(err, "servicefabricmesh.SecretValueClient", "ListValue", resp, "Failure sending request")
		return
	}

	result, err = client.ListValueResponder(resp)
	if err != nil {
		err = autorest.NewErrorWithError(err, "servicefabricmesh.SecretValueClient", "ListValue", resp, "Failure responding to request")
	}

	return
}

// ListValuePreparer prepares the ListValue request.
func (client SecretValueClient) ListValuePreparer(ctx context.Context, resourceGroupName string, secretResourceName string, secretValueResourceName string) (*http.Request, error) {
	pathParameters := map[string]interface{}{
		"resourceGroupName":       autorest.Encode("path", resourceGroupName),
		"secretResourceName":      secretResourceName,
		"secretValueResourceName": secretValueResourceName,
		"subscriptionId":          autorest.Encode("path", client.SubscriptionID),
	}

	const APIVersion = "2018-09-01-preview"
	queryParameters := map[string]interface{}{
		"api-version": APIVersion,
	}

	preparer := autorest.CreatePreparer(
		autorest.AsPost(),
		autorest.WithBaseURL(client.BaseURI),
		autorest.WithPathParameters("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.ServiceFabricMesh/secrets/{secretResourceName}/values/{secretValueResourceName}/list_value", pathParameters),
		autorest.WithQueryParameters(queryParameters))
	return preparer.Prepare((&http.Request{}).WithContext(ctx))
}

// ListValueSender sends the ListValue request. The method will close the
// http.Response Body if it receives an error.
func (client SecretValueClient) ListValueSender(req *http.Request) (*http.Response, error) {
	sd := autorest.GetSendDecorators(req.Context(), azure.DoRetryWithRegistration(client.Client))
	return autorest.SendWithSender(client, req, sd...)
}

// ListValueResponder handles the response to the ListValue request. The method always
// closes the http.Response Body.
func (client SecretValueClient) ListValueResponder(resp *http.Response) (result SecretValue, err error) {
	err = autorest.Respond(
		resp,
		client.ByInspecting(),
		azure.WithErrorUnlessStatusCode(http.StatusOK),
		autorest.ByUnmarshallingJSON(&result),
		autorest.ByClosing())
	result.Response = autorest.Response{Response: resp}
	return
}
