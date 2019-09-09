package logic

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

// CertificatesClient is the REST API for Azure Logic Apps.
type CertificatesClient struct {
	BaseClient
}

// NewCertificatesClient creates an instance of the CertificatesClient client.
func NewCertificatesClient(subscriptionID string) CertificatesClient {
	return NewCertificatesClientWithBaseURI(DefaultBaseURI, subscriptionID)
}

// NewCertificatesClientWithBaseURI creates an instance of the CertificatesClient client.
func NewCertificatesClientWithBaseURI(baseURI string, subscriptionID string) CertificatesClient {
	return CertificatesClient{NewWithBaseURI(baseURI, subscriptionID)}
}

// CreateOrUpdate creates or updates an integration account certificate.
// Parameters:
// resourceGroupName - the resource group name.
// integrationAccountName - the integration account name.
// certificateName - the integration account certificate name.
// certificate - the integration account certificate.
func (client CertificatesClient) CreateOrUpdate(ctx context.Context, resourceGroupName string, integrationAccountName string, certificateName string, certificate IntegrationAccountCertificate) (result IntegrationAccountCertificate, err error) {
	if tracing.IsEnabled() {
		ctx = tracing.StartSpan(ctx, fqdn+"/CertificatesClient.CreateOrUpdate")
		defer func() {
			sc := -1
			if result.Response.Response != nil {
				sc = result.Response.Response.StatusCode
			}
			tracing.EndSpan(ctx, sc, err)
		}()
	}
	if err := validation.Validate([]validation.Validation{
		{TargetValue: certificate,
			Constraints: []validation.Constraint{{Target: "certificate.IntegrationAccountCertificateProperties", Name: validation.Null, Rule: true,
				Chain: []validation.Constraint{{Target: "certificate.IntegrationAccountCertificateProperties.Key", Name: validation.Null, Rule: false,
					Chain: []validation.Constraint{{Target: "certificate.IntegrationAccountCertificateProperties.Key.KeyVault", Name: validation.Null, Rule: true, Chain: nil},
						{Target: "certificate.IntegrationAccountCertificateProperties.Key.KeyName", Name: validation.Null, Rule: true, Chain: nil},
					}},
				}}}}}); err != nil {
		return result, validation.NewError("logic.CertificatesClient", "CreateOrUpdate", err.Error())
	}

	req, err := client.CreateOrUpdatePreparer(ctx, resourceGroupName, integrationAccountName, certificateName, certificate)
	if err != nil {
		err = autorest.NewErrorWithError(err, "logic.CertificatesClient", "CreateOrUpdate", nil, "Failure preparing request")
		return
	}

	resp, err := client.CreateOrUpdateSender(req)
	if err != nil {
		result.Response = autorest.Response{Response: resp}
		err = autorest.NewErrorWithError(err, "logic.CertificatesClient", "CreateOrUpdate", resp, "Failure sending request")
		return
	}

	result, err = client.CreateOrUpdateResponder(resp)
	if err != nil {
		err = autorest.NewErrorWithError(err, "logic.CertificatesClient", "CreateOrUpdate", resp, "Failure responding to request")
	}

	return
}

// CreateOrUpdatePreparer prepares the CreateOrUpdate request.
func (client CertificatesClient) CreateOrUpdatePreparer(ctx context.Context, resourceGroupName string, integrationAccountName string, certificateName string, certificate IntegrationAccountCertificate) (*http.Request, error) {
	pathParameters := map[string]interface{}{
		"certificateName":        autorest.Encode("path", certificateName),
		"integrationAccountName": autorest.Encode("path", integrationAccountName),
		"resourceGroupName":      autorest.Encode("path", resourceGroupName),
		"subscriptionId":         autorest.Encode("path", client.SubscriptionID),
	}

	const APIVersion = "2016-06-01"
	queryParameters := map[string]interface{}{
		"api-version": APIVersion,
	}

	preparer := autorest.CreatePreparer(
		autorest.AsContentType("application/json; charset=utf-8"),
		autorest.AsPut(),
		autorest.WithBaseURL(client.BaseURI),
		autorest.WithPathParameters("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Logic/integrationAccounts/{integrationAccountName}/certificates/{certificateName}", pathParameters),
		autorest.WithJSON(certificate),
		autorest.WithQueryParameters(queryParameters))
	return preparer.Prepare((&http.Request{}).WithContext(ctx))
}

// CreateOrUpdateSender sends the CreateOrUpdate request. The method will close the
// http.Response Body if it receives an error.
func (client CertificatesClient) CreateOrUpdateSender(req *http.Request) (*http.Response, error) {
	sd := autorest.GetSendDecorators(req.Context(), azure.DoRetryWithRegistration(client.Client))
	return autorest.SendWithSender(client, req, sd...)
}

// CreateOrUpdateResponder handles the response to the CreateOrUpdate request. The method always
// closes the http.Response Body.
func (client CertificatesClient) CreateOrUpdateResponder(resp *http.Response) (result IntegrationAccountCertificate, err error) {
	err = autorest.Respond(
		resp,
		client.ByInspecting(),
		azure.WithErrorUnlessStatusCode(http.StatusOK, http.StatusCreated),
		autorest.ByUnmarshallingJSON(&result),
		autorest.ByClosing())
	result.Response = autorest.Response{Response: resp}
	return
}

// Delete deletes an integration account certificate.
// Parameters:
// resourceGroupName - the resource group name.
// integrationAccountName - the integration account name.
// certificateName - the integration account certificate name.
func (client CertificatesClient) Delete(ctx context.Context, resourceGroupName string, integrationAccountName string, certificateName string) (result autorest.Response, err error) {
	if tracing.IsEnabled() {
		ctx = tracing.StartSpan(ctx, fqdn+"/CertificatesClient.Delete")
		defer func() {
			sc := -1
			if result.Response != nil {
				sc = result.Response.StatusCode
			}
			tracing.EndSpan(ctx, sc, err)
		}()
	}
	req, err := client.DeletePreparer(ctx, resourceGroupName, integrationAccountName, certificateName)
	if err != nil {
		err = autorest.NewErrorWithError(err, "logic.CertificatesClient", "Delete", nil, "Failure preparing request")
		return
	}

	resp, err := client.DeleteSender(req)
	if err != nil {
		result.Response = resp
		err = autorest.NewErrorWithError(err, "logic.CertificatesClient", "Delete", resp, "Failure sending request")
		return
	}

	result, err = client.DeleteResponder(resp)
	if err != nil {
		err = autorest.NewErrorWithError(err, "logic.CertificatesClient", "Delete", resp, "Failure responding to request")
	}

	return
}

// DeletePreparer prepares the Delete request.
func (client CertificatesClient) DeletePreparer(ctx context.Context, resourceGroupName string, integrationAccountName string, certificateName string) (*http.Request, error) {
	pathParameters := map[string]interface{}{
		"certificateName":        autorest.Encode("path", certificateName),
		"integrationAccountName": autorest.Encode("path", integrationAccountName),
		"resourceGroupName":      autorest.Encode("path", resourceGroupName),
		"subscriptionId":         autorest.Encode("path", client.SubscriptionID),
	}

	const APIVersion = "2016-06-01"
	queryParameters := map[string]interface{}{
		"api-version": APIVersion,
	}

	preparer := autorest.CreatePreparer(
		autorest.AsDelete(),
		autorest.WithBaseURL(client.BaseURI),
		autorest.WithPathParameters("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Logic/integrationAccounts/{integrationAccountName}/certificates/{certificateName}", pathParameters),
		autorest.WithQueryParameters(queryParameters))
	return preparer.Prepare((&http.Request{}).WithContext(ctx))
}

// DeleteSender sends the Delete request. The method will close the
// http.Response Body if it receives an error.
func (client CertificatesClient) DeleteSender(req *http.Request) (*http.Response, error) {
	sd := autorest.GetSendDecorators(req.Context(), azure.DoRetryWithRegistration(client.Client))
	return autorest.SendWithSender(client, req, sd...)
}

// DeleteResponder handles the response to the Delete request. The method always
// closes the http.Response Body.
func (client CertificatesClient) DeleteResponder(resp *http.Response) (result autorest.Response, err error) {
	err = autorest.Respond(
		resp,
		client.ByInspecting(),
		azure.WithErrorUnlessStatusCode(http.StatusOK, http.StatusNoContent),
		autorest.ByClosing())
	result.Response = resp
	return
}

// Get gets an integration account certificate.
// Parameters:
// resourceGroupName - the resource group name.
// integrationAccountName - the integration account name.
// certificateName - the integration account certificate name.
func (client CertificatesClient) Get(ctx context.Context, resourceGroupName string, integrationAccountName string, certificateName string) (result IntegrationAccountCertificate, err error) {
	if tracing.IsEnabled() {
		ctx = tracing.StartSpan(ctx, fqdn+"/CertificatesClient.Get")
		defer func() {
			sc := -1
			if result.Response.Response != nil {
				sc = result.Response.Response.StatusCode
			}
			tracing.EndSpan(ctx, sc, err)
		}()
	}
	req, err := client.GetPreparer(ctx, resourceGroupName, integrationAccountName, certificateName)
	if err != nil {
		err = autorest.NewErrorWithError(err, "logic.CertificatesClient", "Get", nil, "Failure preparing request")
		return
	}

	resp, err := client.GetSender(req)
	if err != nil {
		result.Response = autorest.Response{Response: resp}
		err = autorest.NewErrorWithError(err, "logic.CertificatesClient", "Get", resp, "Failure sending request")
		return
	}

	result, err = client.GetResponder(resp)
	if err != nil {
		err = autorest.NewErrorWithError(err, "logic.CertificatesClient", "Get", resp, "Failure responding to request")
	}

	return
}

// GetPreparer prepares the Get request.
func (client CertificatesClient) GetPreparer(ctx context.Context, resourceGroupName string, integrationAccountName string, certificateName string) (*http.Request, error) {
	pathParameters := map[string]interface{}{
		"certificateName":        autorest.Encode("path", certificateName),
		"integrationAccountName": autorest.Encode("path", integrationAccountName),
		"resourceGroupName":      autorest.Encode("path", resourceGroupName),
		"subscriptionId":         autorest.Encode("path", client.SubscriptionID),
	}

	const APIVersion = "2016-06-01"
	queryParameters := map[string]interface{}{
		"api-version": APIVersion,
	}

	preparer := autorest.CreatePreparer(
		autorest.AsGet(),
		autorest.WithBaseURL(client.BaseURI),
		autorest.WithPathParameters("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Logic/integrationAccounts/{integrationAccountName}/certificates/{certificateName}", pathParameters),
		autorest.WithQueryParameters(queryParameters))
	return preparer.Prepare((&http.Request{}).WithContext(ctx))
}

// GetSender sends the Get request. The method will close the
// http.Response Body if it receives an error.
func (client CertificatesClient) GetSender(req *http.Request) (*http.Response, error) {
	sd := autorest.GetSendDecorators(req.Context(), azure.DoRetryWithRegistration(client.Client))
	return autorest.SendWithSender(client, req, sd...)
}

// GetResponder handles the response to the Get request. The method always
// closes the http.Response Body.
func (client CertificatesClient) GetResponder(resp *http.Response) (result IntegrationAccountCertificate, err error) {
	err = autorest.Respond(
		resp,
		client.ByInspecting(),
		azure.WithErrorUnlessStatusCode(http.StatusOK),
		autorest.ByUnmarshallingJSON(&result),
		autorest.ByClosing())
	result.Response = autorest.Response{Response: resp}
	return
}

// ListByIntegrationAccounts gets a list of integration account certificates.
// Parameters:
// resourceGroupName - the resource group name.
// integrationAccountName - the integration account name.
// top - the number of items to be included in the result.
func (client CertificatesClient) ListByIntegrationAccounts(ctx context.Context, resourceGroupName string, integrationAccountName string, top *int32) (result IntegrationAccountCertificateListResultPage, err error) {
	if tracing.IsEnabled() {
		ctx = tracing.StartSpan(ctx, fqdn+"/CertificatesClient.ListByIntegrationAccounts")
		defer func() {
			sc := -1
			if result.iaclr.Response.Response != nil {
				sc = result.iaclr.Response.Response.StatusCode
			}
			tracing.EndSpan(ctx, sc, err)
		}()
	}
	result.fn = client.listByIntegrationAccountsNextResults
	req, err := client.ListByIntegrationAccountsPreparer(ctx, resourceGroupName, integrationAccountName, top)
	if err != nil {
		err = autorest.NewErrorWithError(err, "logic.CertificatesClient", "ListByIntegrationAccounts", nil, "Failure preparing request")
		return
	}

	resp, err := client.ListByIntegrationAccountsSender(req)
	if err != nil {
		result.iaclr.Response = autorest.Response{Response: resp}
		err = autorest.NewErrorWithError(err, "logic.CertificatesClient", "ListByIntegrationAccounts", resp, "Failure sending request")
		return
	}

	result.iaclr, err = client.ListByIntegrationAccountsResponder(resp)
	if err != nil {
		err = autorest.NewErrorWithError(err, "logic.CertificatesClient", "ListByIntegrationAccounts", resp, "Failure responding to request")
	}

	return
}

// ListByIntegrationAccountsPreparer prepares the ListByIntegrationAccounts request.
func (client CertificatesClient) ListByIntegrationAccountsPreparer(ctx context.Context, resourceGroupName string, integrationAccountName string, top *int32) (*http.Request, error) {
	pathParameters := map[string]interface{}{
		"integrationAccountName": autorest.Encode("path", integrationAccountName),
		"resourceGroupName":      autorest.Encode("path", resourceGroupName),
		"subscriptionId":         autorest.Encode("path", client.SubscriptionID),
	}

	const APIVersion = "2016-06-01"
	queryParameters := map[string]interface{}{
		"api-version": APIVersion,
	}
	if top != nil {
		queryParameters["$top"] = autorest.Encode("query", *top)
	}

	preparer := autorest.CreatePreparer(
		autorest.AsGet(),
		autorest.WithBaseURL(client.BaseURI),
		autorest.WithPathParameters("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Logic/integrationAccounts/{integrationAccountName}/certificates", pathParameters),
		autorest.WithQueryParameters(queryParameters))
	return preparer.Prepare((&http.Request{}).WithContext(ctx))
}

// ListByIntegrationAccountsSender sends the ListByIntegrationAccounts request. The method will close the
// http.Response Body if it receives an error.
func (client CertificatesClient) ListByIntegrationAccountsSender(req *http.Request) (*http.Response, error) {
	sd := autorest.GetSendDecorators(req.Context(), azure.DoRetryWithRegistration(client.Client))
	return autorest.SendWithSender(client, req, sd...)
}

// ListByIntegrationAccountsResponder handles the response to the ListByIntegrationAccounts request. The method always
// closes the http.Response Body.
func (client CertificatesClient) ListByIntegrationAccountsResponder(resp *http.Response) (result IntegrationAccountCertificateListResult, err error) {
	err = autorest.Respond(
		resp,
		client.ByInspecting(),
		azure.WithErrorUnlessStatusCode(http.StatusOK),
		autorest.ByUnmarshallingJSON(&result),
		autorest.ByClosing())
	result.Response = autorest.Response{Response: resp}
	return
}

// listByIntegrationAccountsNextResults retrieves the next set of results, if any.
func (client CertificatesClient) listByIntegrationAccountsNextResults(ctx context.Context, lastResults IntegrationAccountCertificateListResult) (result IntegrationAccountCertificateListResult, err error) {
	req, err := lastResults.integrationAccountCertificateListResultPreparer(ctx)
	if err != nil {
		return result, autorest.NewErrorWithError(err, "logic.CertificatesClient", "listByIntegrationAccountsNextResults", nil, "Failure preparing next results request")
	}
	if req == nil {
		return
	}
	resp, err := client.ListByIntegrationAccountsSender(req)
	if err != nil {
		result.Response = autorest.Response{Response: resp}
		return result, autorest.NewErrorWithError(err, "logic.CertificatesClient", "listByIntegrationAccountsNextResults", resp, "Failure sending next results request")
	}
	result, err = client.ListByIntegrationAccountsResponder(resp)
	if err != nil {
		err = autorest.NewErrorWithError(err, "logic.CertificatesClient", "listByIntegrationAccountsNextResults", resp, "Failure responding to next results request")
	}
	return
}

// ListByIntegrationAccountsComplete enumerates all values, automatically crossing page boundaries as required.
func (client CertificatesClient) ListByIntegrationAccountsComplete(ctx context.Context, resourceGroupName string, integrationAccountName string, top *int32) (result IntegrationAccountCertificateListResultIterator, err error) {
	if tracing.IsEnabled() {
		ctx = tracing.StartSpan(ctx, fqdn+"/CertificatesClient.ListByIntegrationAccounts")
		defer func() {
			sc := -1
			if result.Response().Response.Response != nil {
				sc = result.page.Response().Response.Response.StatusCode
			}
			tracing.EndSpan(ctx, sc, err)
		}()
	}
	result.page, err = client.ListByIntegrationAccounts(ctx, resourceGroupName, integrationAccountName, top)
	return
}
