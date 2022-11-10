package compute

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See License.txt in the project root for license information.
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

// GalleryApplicationVersionsClient is the compute Client
type GalleryApplicationVersionsClient struct {
	BaseClient
}

// NewGalleryApplicationVersionsClient creates an instance of the GalleryApplicationVersionsClient client.
func NewGalleryApplicationVersionsClient(subscriptionID string) GalleryApplicationVersionsClient {
	return NewGalleryApplicationVersionsClientWithBaseURI(DefaultBaseURI, subscriptionID)
}

// NewGalleryApplicationVersionsClientWithBaseURI creates an instance of the GalleryApplicationVersionsClient client
// using a custom endpoint.  Use this when interacting with an Azure cloud that uses a non-standard base URI (sovereign
// clouds, Azure stack).
func NewGalleryApplicationVersionsClientWithBaseURI(baseURI string, subscriptionID string) GalleryApplicationVersionsClient {
	return GalleryApplicationVersionsClient{NewWithBaseURI(baseURI, subscriptionID)}
}

// CreateOrUpdate create or update a gallery Application Version.
// Parameters:
// resourceGroupName - the name of the resource group.
// galleryName - the name of the Shared Application Gallery in which the Application Definition resides.
// galleryApplicationName - the name of the gallery Application Definition in which the Application Version is
// to be created.
// galleryApplicationVersionName - the name of the gallery Application Version to be created. Needs to follow
// semantic version name pattern: The allowed characters are digit and period. Digits must be within the range
// of a 32-bit integer. Format: <MajorVersion>.<MinorVersion>.<Patch>
// galleryApplicationVersion - parameters supplied to the create or update gallery Application Version
// operation.
func (client GalleryApplicationVersionsClient) CreateOrUpdate(ctx context.Context, resourceGroupName string, galleryName string, galleryApplicationName string, galleryApplicationVersionName string, galleryApplicationVersion GalleryApplicationVersion) (result GalleryApplicationVersionsCreateOrUpdateFuture, err error) {
	if tracing.IsEnabled() {
		ctx = tracing.StartSpan(ctx, fqdn+"/GalleryApplicationVersionsClient.CreateOrUpdate")
		defer func() {
			sc := -1
			if result.FutureAPI != nil && result.FutureAPI.Response() != nil {
				sc = result.FutureAPI.Response().StatusCode
			}
			tracing.EndSpan(ctx, sc, err)
		}()
	}
	if err := validation.Validate([]validation.Validation{
		{TargetValue: galleryApplicationVersion,
			Constraints: []validation.Constraint{{Target: "galleryApplicationVersion.GalleryApplicationVersionProperties", Name: validation.Null, Rule: false,
				Chain: []validation.Constraint{{Target: "galleryApplicationVersion.GalleryApplicationVersionProperties.PublishingProfile", Name: validation.Null, Rule: true,
					Chain: []validation.Constraint{{Target: "galleryApplicationVersion.GalleryApplicationVersionProperties.PublishingProfile.Source", Name: validation.Null, Rule: true,
						Chain: []validation.Constraint{{Target: "galleryApplicationVersion.GalleryApplicationVersionProperties.PublishingProfile.Source.MediaLink", Name: validation.Null, Rule: true, Chain: nil}}},
						{Target: "galleryApplicationVersion.GalleryApplicationVersionProperties.PublishingProfile.ManageActions", Name: validation.Null, Rule: false,
							Chain: []validation.Constraint{{Target: "galleryApplicationVersion.GalleryApplicationVersionProperties.PublishingProfile.ManageActions.Install", Name: validation.Null, Rule: true, Chain: nil},
								{Target: "galleryApplicationVersion.GalleryApplicationVersionProperties.PublishingProfile.ManageActions.Remove", Name: validation.Null, Rule: true, Chain: nil},
							}},
					}},
				}}}}}); err != nil {
		return result, validation.NewError("compute.GalleryApplicationVersionsClient", "CreateOrUpdate", err.Error())
	}

	req, err := client.CreateOrUpdatePreparer(ctx, resourceGroupName, galleryName, galleryApplicationName, galleryApplicationVersionName, galleryApplicationVersion)
	if err != nil {
		err = autorest.NewErrorWithError(err, "compute.GalleryApplicationVersionsClient", "CreateOrUpdate", nil, "Failure preparing request")
		return
	}

	result, err = client.CreateOrUpdateSender(req)
	if err != nil {
		err = autorest.NewErrorWithError(err, "compute.GalleryApplicationVersionsClient", "CreateOrUpdate", result.Response(), "Failure sending request")
		return
	}

	return
}

// CreateOrUpdatePreparer prepares the CreateOrUpdate request.
func (client GalleryApplicationVersionsClient) CreateOrUpdatePreparer(ctx context.Context, resourceGroupName string, galleryName string, galleryApplicationName string, galleryApplicationVersionName string, galleryApplicationVersion GalleryApplicationVersion) (*http.Request, error) {
	pathParameters := map[string]interface{}{
		"galleryApplicationName":        autorest.Encode("path", galleryApplicationName),
		"galleryApplicationVersionName": autorest.Encode("path", galleryApplicationVersionName),
		"galleryName":                   autorest.Encode("path", galleryName),
		"resourceGroupName":             autorest.Encode("path", resourceGroupName),
		"subscriptionId":                autorest.Encode("path", client.SubscriptionID),
	}

	const APIVersion = "2021-10-01"
	queryParameters := map[string]interface{}{
		"api-version": APIVersion,
	}

	preparer := autorest.CreatePreparer(
		autorest.AsContentType("application/json; charset=utf-8"),
		autorest.AsPut(),
		autorest.WithBaseURL(client.BaseURI),
		autorest.WithPathParameters("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Compute/galleries/{galleryName}/applications/{galleryApplicationName}/versions/{galleryApplicationVersionName}", pathParameters),
		autorest.WithJSON(galleryApplicationVersion),
		autorest.WithQueryParameters(queryParameters))
	return preparer.Prepare((&http.Request{}).WithContext(ctx))
}

// CreateOrUpdateSender sends the CreateOrUpdate request. The method will close the
// http.Response Body if it receives an error.
func (client GalleryApplicationVersionsClient) CreateOrUpdateSender(req *http.Request) (future GalleryApplicationVersionsCreateOrUpdateFuture, err error) {
	var resp *http.Response
	future.FutureAPI = &azure.Future{}
	resp, err = client.Send(req, azure.DoRetryWithRegistration(client.Client))
	if err != nil {
		return
	}
	var azf azure.Future
	azf, err = azure.NewFutureFromResponse(resp)
	future.FutureAPI = &azf
	future.Result = future.result
	return
}

// CreateOrUpdateResponder handles the response to the CreateOrUpdate request. The method always
// closes the http.Response Body.
func (client GalleryApplicationVersionsClient) CreateOrUpdateResponder(resp *http.Response) (result GalleryApplicationVersion, err error) {
	err = autorest.Respond(
		resp,
		azure.WithErrorUnlessStatusCode(http.StatusOK, http.StatusCreated, http.StatusAccepted),
		autorest.ByUnmarshallingJSON(&result),
		autorest.ByClosing())
	result.Response = autorest.Response{Response: resp}
	return
}

// Delete delete a gallery Application Version.
// Parameters:
// resourceGroupName - the name of the resource group.
// galleryName - the name of the Shared Application Gallery in which the Application Definition resides.
// galleryApplicationName - the name of the gallery Application Definition in which the Application Version
// resides.
// galleryApplicationVersionName - the name of the gallery Application Version to be deleted.
func (client GalleryApplicationVersionsClient) Delete(ctx context.Context, resourceGroupName string, galleryName string, galleryApplicationName string, galleryApplicationVersionName string) (result GalleryApplicationVersionsDeleteFuture, err error) {
	if tracing.IsEnabled() {
		ctx = tracing.StartSpan(ctx, fqdn+"/GalleryApplicationVersionsClient.Delete")
		defer func() {
			sc := -1
			if result.FutureAPI != nil && result.FutureAPI.Response() != nil {
				sc = result.FutureAPI.Response().StatusCode
			}
			tracing.EndSpan(ctx, sc, err)
		}()
	}
	req, err := client.DeletePreparer(ctx, resourceGroupName, galleryName, galleryApplicationName, galleryApplicationVersionName)
	if err != nil {
		err = autorest.NewErrorWithError(err, "compute.GalleryApplicationVersionsClient", "Delete", nil, "Failure preparing request")
		return
	}

	result, err = client.DeleteSender(req)
	if err != nil {
		err = autorest.NewErrorWithError(err, "compute.GalleryApplicationVersionsClient", "Delete", result.Response(), "Failure sending request")
		return
	}

	return
}

// DeletePreparer prepares the Delete request.
func (client GalleryApplicationVersionsClient) DeletePreparer(ctx context.Context, resourceGroupName string, galleryName string, galleryApplicationName string, galleryApplicationVersionName string) (*http.Request, error) {
	pathParameters := map[string]interface{}{
		"galleryApplicationName":        autorest.Encode("path", galleryApplicationName),
		"galleryApplicationVersionName": autorest.Encode("path", galleryApplicationVersionName),
		"galleryName":                   autorest.Encode("path", galleryName),
		"resourceGroupName":             autorest.Encode("path", resourceGroupName),
		"subscriptionId":                autorest.Encode("path", client.SubscriptionID),
	}

	const APIVersion = "2021-10-01"
	queryParameters := map[string]interface{}{
		"api-version": APIVersion,
	}

	preparer := autorest.CreatePreparer(
		autorest.AsDelete(),
		autorest.WithBaseURL(client.BaseURI),
		autorest.WithPathParameters("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Compute/galleries/{galleryName}/applications/{galleryApplicationName}/versions/{galleryApplicationVersionName}", pathParameters),
		autorest.WithQueryParameters(queryParameters))
	return preparer.Prepare((&http.Request{}).WithContext(ctx))
}

// DeleteSender sends the Delete request. The method will close the
// http.Response Body if it receives an error.
func (client GalleryApplicationVersionsClient) DeleteSender(req *http.Request) (future GalleryApplicationVersionsDeleteFuture, err error) {
	var resp *http.Response
	future.FutureAPI = &azure.Future{}
	resp, err = client.Send(req, azure.DoRetryWithRegistration(client.Client))
	if err != nil {
		return
	}
	var azf azure.Future
	azf, err = azure.NewFutureFromResponse(resp)
	future.FutureAPI = &azf
	future.Result = future.result
	return
}

// DeleteResponder handles the response to the Delete request. The method always
// closes the http.Response Body.
func (client GalleryApplicationVersionsClient) DeleteResponder(resp *http.Response) (result autorest.Response, err error) {
	err = autorest.Respond(
		resp,
		azure.WithErrorUnlessStatusCode(http.StatusOK, http.StatusAccepted, http.StatusNoContent),
		autorest.ByClosing())
	result.Response = resp
	return
}

// Get retrieves information about a gallery Application Version.
// Parameters:
// resourceGroupName - the name of the resource group.
// galleryName - the name of the Shared Application Gallery in which the Application Definition resides.
// galleryApplicationName - the name of the gallery Application Definition in which the Application Version
// resides.
// galleryApplicationVersionName - the name of the gallery Application Version to be retrieved.
// expand - the expand expression to apply on the operation.
func (client GalleryApplicationVersionsClient) Get(ctx context.Context, resourceGroupName string, galleryName string, galleryApplicationName string, galleryApplicationVersionName string, expand ReplicationStatusTypes) (result GalleryApplicationVersion, err error) {
	if tracing.IsEnabled() {
		ctx = tracing.StartSpan(ctx, fqdn+"/GalleryApplicationVersionsClient.Get")
		defer func() {
			sc := -1
			if result.Response.Response != nil {
				sc = result.Response.Response.StatusCode
			}
			tracing.EndSpan(ctx, sc, err)
		}()
	}
	req, err := client.GetPreparer(ctx, resourceGroupName, galleryName, galleryApplicationName, galleryApplicationVersionName, expand)
	if err != nil {
		err = autorest.NewErrorWithError(err, "compute.GalleryApplicationVersionsClient", "Get", nil, "Failure preparing request")
		return
	}

	resp, err := client.GetSender(req)
	if err != nil {
		result.Response = autorest.Response{Response: resp}
		err = autorest.NewErrorWithError(err, "compute.GalleryApplicationVersionsClient", "Get", resp, "Failure sending request")
		return
	}

	result, err = client.GetResponder(resp)
	if err != nil {
		err = autorest.NewErrorWithError(err, "compute.GalleryApplicationVersionsClient", "Get", resp, "Failure responding to request")
		return
	}

	return
}

// GetPreparer prepares the Get request.
func (client GalleryApplicationVersionsClient) GetPreparer(ctx context.Context, resourceGroupName string, galleryName string, galleryApplicationName string, galleryApplicationVersionName string, expand ReplicationStatusTypes) (*http.Request, error) {
	pathParameters := map[string]interface{}{
		"galleryApplicationName":        autorest.Encode("path", galleryApplicationName),
		"galleryApplicationVersionName": autorest.Encode("path", galleryApplicationVersionName),
		"galleryName":                   autorest.Encode("path", galleryName),
		"resourceGroupName":             autorest.Encode("path", resourceGroupName),
		"subscriptionId":                autorest.Encode("path", client.SubscriptionID),
	}

	const APIVersion = "2021-10-01"
	queryParameters := map[string]interface{}{
		"api-version": APIVersion,
	}
	if len(string(expand)) > 0 {
		queryParameters["$expand"] = autorest.Encode("query", expand)
	}

	preparer := autorest.CreatePreparer(
		autorest.AsGet(),
		autorest.WithBaseURL(client.BaseURI),
		autorest.WithPathParameters("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Compute/galleries/{galleryName}/applications/{galleryApplicationName}/versions/{galleryApplicationVersionName}", pathParameters),
		autorest.WithQueryParameters(queryParameters))
	return preparer.Prepare((&http.Request{}).WithContext(ctx))
}

// GetSender sends the Get request. The method will close the
// http.Response Body if it receives an error.
func (client GalleryApplicationVersionsClient) GetSender(req *http.Request) (*http.Response, error) {
	return client.Send(req, azure.DoRetryWithRegistration(client.Client))
}

// GetResponder handles the response to the Get request. The method always
// closes the http.Response Body.
func (client GalleryApplicationVersionsClient) GetResponder(resp *http.Response) (result GalleryApplicationVersion, err error) {
	err = autorest.Respond(
		resp,
		azure.WithErrorUnlessStatusCode(http.StatusOK),
		autorest.ByUnmarshallingJSON(&result),
		autorest.ByClosing())
	result.Response = autorest.Response{Response: resp}
	return
}

// ListByGalleryApplication list gallery Application Versions in a gallery Application Definition.
// Parameters:
// resourceGroupName - the name of the resource group.
// galleryName - the name of the Shared Application Gallery in which the Application Definition resides.
// galleryApplicationName - the name of the Shared Application Gallery Application Definition from which the
// Application Versions are to be listed.
func (client GalleryApplicationVersionsClient) ListByGalleryApplication(ctx context.Context, resourceGroupName string, galleryName string, galleryApplicationName string) (result GalleryApplicationVersionListPage, err error) {
	if tracing.IsEnabled() {
		ctx = tracing.StartSpan(ctx, fqdn+"/GalleryApplicationVersionsClient.ListByGalleryApplication")
		defer func() {
			sc := -1
			if result.gavl.Response.Response != nil {
				sc = result.gavl.Response.Response.StatusCode
			}
			tracing.EndSpan(ctx, sc, err)
		}()
	}
	result.fn = client.listByGalleryApplicationNextResults
	req, err := client.ListByGalleryApplicationPreparer(ctx, resourceGroupName, galleryName, galleryApplicationName)
	if err != nil {
		err = autorest.NewErrorWithError(err, "compute.GalleryApplicationVersionsClient", "ListByGalleryApplication", nil, "Failure preparing request")
		return
	}

	resp, err := client.ListByGalleryApplicationSender(req)
	if err != nil {
		result.gavl.Response = autorest.Response{Response: resp}
		err = autorest.NewErrorWithError(err, "compute.GalleryApplicationVersionsClient", "ListByGalleryApplication", resp, "Failure sending request")
		return
	}

	result.gavl, err = client.ListByGalleryApplicationResponder(resp)
	if err != nil {
		err = autorest.NewErrorWithError(err, "compute.GalleryApplicationVersionsClient", "ListByGalleryApplication", resp, "Failure responding to request")
		return
	}
	if result.gavl.hasNextLink() && result.gavl.IsEmpty() {
		err = result.NextWithContext(ctx)
		return
	}

	return
}

// ListByGalleryApplicationPreparer prepares the ListByGalleryApplication request.
func (client GalleryApplicationVersionsClient) ListByGalleryApplicationPreparer(ctx context.Context, resourceGroupName string, galleryName string, galleryApplicationName string) (*http.Request, error) {
	pathParameters := map[string]interface{}{
		"galleryApplicationName": autorest.Encode("path", galleryApplicationName),
		"galleryName":            autorest.Encode("path", galleryName),
		"resourceGroupName":      autorest.Encode("path", resourceGroupName),
		"subscriptionId":         autorest.Encode("path", client.SubscriptionID),
	}

	const APIVersion = "2021-10-01"
	queryParameters := map[string]interface{}{
		"api-version": APIVersion,
	}

	preparer := autorest.CreatePreparer(
		autorest.AsGet(),
		autorest.WithBaseURL(client.BaseURI),
		autorest.WithPathParameters("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Compute/galleries/{galleryName}/applications/{galleryApplicationName}/versions", pathParameters),
		autorest.WithQueryParameters(queryParameters))
	return preparer.Prepare((&http.Request{}).WithContext(ctx))
}

// ListByGalleryApplicationSender sends the ListByGalleryApplication request. The method will close the
// http.Response Body if it receives an error.
func (client GalleryApplicationVersionsClient) ListByGalleryApplicationSender(req *http.Request) (*http.Response, error) {
	return client.Send(req, azure.DoRetryWithRegistration(client.Client))
}

// ListByGalleryApplicationResponder handles the response to the ListByGalleryApplication request. The method always
// closes the http.Response Body.
func (client GalleryApplicationVersionsClient) ListByGalleryApplicationResponder(resp *http.Response) (result GalleryApplicationVersionList, err error) {
	err = autorest.Respond(
		resp,
		azure.WithErrorUnlessStatusCode(http.StatusOK),
		autorest.ByUnmarshallingJSON(&result),
		autorest.ByClosing())
	result.Response = autorest.Response{Response: resp}
	return
}

// listByGalleryApplicationNextResults retrieves the next set of results, if any.
func (client GalleryApplicationVersionsClient) listByGalleryApplicationNextResults(ctx context.Context, lastResults GalleryApplicationVersionList) (result GalleryApplicationVersionList, err error) {
	req, err := lastResults.galleryApplicationVersionListPreparer(ctx)
	if err != nil {
		return result, autorest.NewErrorWithError(err, "compute.GalleryApplicationVersionsClient", "listByGalleryApplicationNextResults", nil, "Failure preparing next results request")
	}
	if req == nil {
		return
	}
	resp, err := client.ListByGalleryApplicationSender(req)
	if err != nil {
		result.Response = autorest.Response{Response: resp}
		return result, autorest.NewErrorWithError(err, "compute.GalleryApplicationVersionsClient", "listByGalleryApplicationNextResults", resp, "Failure sending next results request")
	}
	result, err = client.ListByGalleryApplicationResponder(resp)
	if err != nil {
		err = autorest.NewErrorWithError(err, "compute.GalleryApplicationVersionsClient", "listByGalleryApplicationNextResults", resp, "Failure responding to next results request")
	}
	return
}

// ListByGalleryApplicationComplete enumerates all values, automatically crossing page boundaries as required.
func (client GalleryApplicationVersionsClient) ListByGalleryApplicationComplete(ctx context.Context, resourceGroupName string, galleryName string, galleryApplicationName string) (result GalleryApplicationVersionListIterator, err error) {
	if tracing.IsEnabled() {
		ctx = tracing.StartSpan(ctx, fqdn+"/GalleryApplicationVersionsClient.ListByGalleryApplication")
		defer func() {
			sc := -1
			if result.Response().Response.Response != nil {
				sc = result.page.Response().Response.Response.StatusCode
			}
			tracing.EndSpan(ctx, sc, err)
		}()
	}
	result.page, err = client.ListByGalleryApplication(ctx, resourceGroupName, galleryName, galleryApplicationName)
	return
}

// Update update a gallery Application Version.
// Parameters:
// resourceGroupName - the name of the resource group.
// galleryName - the name of the Shared Application Gallery in which the Application Definition resides.
// galleryApplicationName - the name of the gallery Application Definition in which the Application Version is
// to be updated.
// galleryApplicationVersionName - the name of the gallery Application Version to be updated. Needs to follow
// semantic version name pattern: The allowed characters are digit and period. Digits must be within the range
// of a 32-bit integer. Format: <MajorVersion>.<MinorVersion>.<Patch>
// galleryApplicationVersion - parameters supplied to the update gallery Application Version operation.
func (client GalleryApplicationVersionsClient) Update(ctx context.Context, resourceGroupName string, galleryName string, galleryApplicationName string, galleryApplicationVersionName string, galleryApplicationVersion GalleryApplicationVersionUpdate) (result GalleryApplicationVersionsUpdateFuture, err error) {
	if tracing.IsEnabled() {
		ctx = tracing.StartSpan(ctx, fqdn+"/GalleryApplicationVersionsClient.Update")
		defer func() {
			sc := -1
			if result.FutureAPI != nil && result.FutureAPI.Response() != nil {
				sc = result.FutureAPI.Response().StatusCode
			}
			tracing.EndSpan(ctx, sc, err)
		}()
	}
	req, err := client.UpdatePreparer(ctx, resourceGroupName, galleryName, galleryApplicationName, galleryApplicationVersionName, galleryApplicationVersion)
	if err != nil {
		err = autorest.NewErrorWithError(err, "compute.GalleryApplicationVersionsClient", "Update", nil, "Failure preparing request")
		return
	}

	result, err = client.UpdateSender(req)
	if err != nil {
		err = autorest.NewErrorWithError(err, "compute.GalleryApplicationVersionsClient", "Update", result.Response(), "Failure sending request")
		return
	}

	return
}

// UpdatePreparer prepares the Update request.
func (client GalleryApplicationVersionsClient) UpdatePreparer(ctx context.Context, resourceGroupName string, galleryName string, galleryApplicationName string, galleryApplicationVersionName string, galleryApplicationVersion GalleryApplicationVersionUpdate) (*http.Request, error) {
	pathParameters := map[string]interface{}{
		"galleryApplicationName":        autorest.Encode("path", galleryApplicationName),
		"galleryApplicationVersionName": autorest.Encode("path", galleryApplicationVersionName),
		"galleryName":                   autorest.Encode("path", galleryName),
		"resourceGroupName":             autorest.Encode("path", resourceGroupName),
		"subscriptionId":                autorest.Encode("path", client.SubscriptionID),
	}

	const APIVersion = "2021-10-01"
	queryParameters := map[string]interface{}{
		"api-version": APIVersion,
	}

	preparer := autorest.CreatePreparer(
		autorest.AsContentType("application/json; charset=utf-8"),
		autorest.AsPatch(),
		autorest.WithBaseURL(client.BaseURI),
		autorest.WithPathParameters("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Compute/galleries/{galleryName}/applications/{galleryApplicationName}/versions/{galleryApplicationVersionName}", pathParameters),
		autorest.WithJSON(galleryApplicationVersion),
		autorest.WithQueryParameters(queryParameters))
	return preparer.Prepare((&http.Request{}).WithContext(ctx))
}

// UpdateSender sends the Update request. The method will close the
// http.Response Body if it receives an error.
func (client GalleryApplicationVersionsClient) UpdateSender(req *http.Request) (future GalleryApplicationVersionsUpdateFuture, err error) {
	var resp *http.Response
	future.FutureAPI = &azure.Future{}
	resp, err = client.Send(req, azure.DoRetryWithRegistration(client.Client))
	if err != nil {
		return
	}
	var azf azure.Future
	azf, err = azure.NewFutureFromResponse(resp)
	future.FutureAPI = &azf
	future.Result = future.result
	return
}

// UpdateResponder handles the response to the Update request. The method always
// closes the http.Response Body.
func (client GalleryApplicationVersionsClient) UpdateResponder(resp *http.Response) (result GalleryApplicationVersion, err error) {
	err = autorest.Respond(
		resp,
		azure.WithErrorUnlessStatusCode(http.StatusOK),
		autorest.ByUnmarshallingJSON(&result),
		autorest.ByClosing())
	result.Response = autorest.Response{Response: resp}
	return
}
