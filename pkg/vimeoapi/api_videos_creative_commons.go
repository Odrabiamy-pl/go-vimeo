/*
Vimeo API

Build something great. Vimeo's API supports flexible, high-quality video integration with your custom apps.

API version: 3.4
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package vimeoapi

import (
	"bytes"
	"context"
	"io"
	"net/http"
	"net/url"
)

type VideosCreativeCommonsAPI interface {

	/*
		GetCcLicenses Get all Creative Commons licenses

		This method returns all available Creative Commons licenses.

		@param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
		@return ApiGetCcLicensesRequest
	*/
	GetCcLicenses(ctx context.Context) ApiGetCcLicensesRequest

	// GetCcLicensesExecute executes the request
	//  @return []CreativeCommons
	GetCcLicensesExecute(r ApiGetCcLicensesRequest) ([]CreativeCommons, *http.Response, error)
}

// VideosCreativeCommonsAPIService VideosCreativeCommonsAPI service
type VideosCreativeCommonsAPIService service

type ApiGetCcLicensesRequest struct {
	ctx        context.Context
	ApiService VideosCreativeCommonsAPI
	page       *float32
	perPage    *float32
}

// The page number of the results to show.
func (r ApiGetCcLicensesRequest) Page(page float32) ApiGetCcLicensesRequest {
	r.page = &page
	return r
}

// The number of items to show on each page of results, up to a maximum of 100.
func (r ApiGetCcLicensesRequest) PerPage(perPage float32) ApiGetCcLicensesRequest {
	r.perPage = &perPage
	return r
}

func (r ApiGetCcLicensesRequest) Execute() ([]CreativeCommons, *http.Response, error) {
	return r.ApiService.GetCcLicensesExecute(r)
}

/*
GetCcLicenses Get all Creative Commons licenses

This method returns all available Creative Commons licenses.

	@param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
	@return ApiGetCcLicensesRequest
*/
func (a *VideosCreativeCommonsAPIService) GetCcLicenses(ctx context.Context) ApiGetCcLicensesRequest {
	return ApiGetCcLicensesRequest{
		ApiService: a,
		ctx:        ctx,
	}
}

// Execute executes the request
//
//	@return []CreativeCommons
func (a *VideosCreativeCommonsAPIService) GetCcLicensesExecute(r ApiGetCcLicensesRequest) ([]CreativeCommons, *http.Response, error) {
	var (
		localVarHTTPMethod  = http.MethodGet
		localVarPostBody    interface{}
		formFiles           []formFile
		localVarReturnValue []CreativeCommons
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "VideosCreativeCommonsAPIService.GetCcLicenses")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/creativecommons"

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}

	if r.page != nil {
		parameterAddToHeaderOrQuery(localVarQueryParams, "page", r.page, "")
	}
	if r.perPage != nil {
		parameterAddToHeaderOrQuery(localVarQueryParams, "per_page", r.perPage, "")
	}
	// to determine the Content-Type header
	localVarHTTPContentTypes := []string{}

	// set Content-Type header
	localVarHTTPContentType := selectHeaderContentType(localVarHTTPContentTypes)
	if localVarHTTPContentType != "" {
		localVarHeaderParams["Content-Type"] = localVarHTTPContentType
	}

	// to determine the Accept header
	localVarHTTPHeaderAccepts := []string{"application/vnd.vimeo.creativecommons+json"}

	// set Accept header
	localVarHTTPHeaderAccept := selectHeaderAccept(localVarHTTPHeaderAccepts)
	if localVarHTTPHeaderAccept != "" {
		localVarHeaderParams["Accept"] = localVarHTTPHeaderAccept
	}
	req, err := a.client.prepareRequest(r.ctx, localVarPath, localVarHTTPMethod, localVarPostBody, localVarHeaderParams, localVarQueryParams, localVarFormParams, formFiles)
	if err != nil {
		return localVarReturnValue, nil, err
	}

	localVarHTTPResponse, err := a.client.callAPI(req)
	if err != nil || localVarHTTPResponse == nil {
		return localVarReturnValue, localVarHTTPResponse, err
	}

	localVarBody, err := io.ReadAll(localVarHTTPResponse.Body)
	localVarHTTPResponse.Body.Close()
	localVarHTTPResponse.Body = io.NopCloser(bytes.NewBuffer(localVarBody))
	if err != nil {
		return localVarReturnValue, localVarHTTPResponse, err
	}

	if localVarHTTPResponse.StatusCode >= 300 {
		newErr := &GenericOpenAPIError{
			body:  localVarBody,
			error: localVarHTTPResponse.Status,
		}
		return localVarReturnValue, localVarHTTPResponse, newErr
	}

	err = a.client.decode(&localVarReturnValue, localVarBody, localVarHTTPResponse.Header.Get("Content-Type"))
	if err != nil {
		newErr := &GenericOpenAPIError{
			body:  localVarBody,
			error: err.Error(),
		}
		return localVarReturnValue, localVarHTTPResponse, newErr
	}

	return localVarReturnValue, localVarHTTPResponse, nil
}
