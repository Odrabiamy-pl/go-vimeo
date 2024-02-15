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


type VideosLanguagesAPI interface {

	/*
	GetLanguages Get all languages

	This method returns all available video languages.

	@param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
	@return ApiGetLanguagesRequest
	*/
	GetLanguages(ctx context.Context) ApiGetLanguagesRequest

	// GetLanguagesExecute executes the request
	//  @return []Language
	GetLanguagesExecute(r ApiGetLanguagesRequest) ([]Language, *http.Response, error)
}

// VideosLanguagesAPIService VideosLanguagesAPI service
type VideosLanguagesAPIService service

type ApiGetLanguagesRequest struct {
	ctx context.Context
	ApiService VideosLanguagesAPI
	filter *string
	page *float32
	perPage *float32
}

// The attribute by which to filter the results.  Option descriptions:  * &#x60;audiotracks&#x60; - Return languages that can be used for audio tracks.  * &#x60;texttracks&#x60; - Return languages that can be used for text tracks. 
func (r ApiGetLanguagesRequest) Filter(filter string) ApiGetLanguagesRequest {
	r.filter = &filter
	return r
}

// The page number of the results to show.
func (r ApiGetLanguagesRequest) Page(page float32) ApiGetLanguagesRequest {
	r.page = &page
	return r
}

// The number of items to show on each page of results, up to a maximum of 100.
func (r ApiGetLanguagesRequest) PerPage(perPage float32) ApiGetLanguagesRequest {
	r.perPage = &perPage
	return r
}

func (r ApiGetLanguagesRequest) Execute() ([]Language, *http.Response, error) {
	return r.ApiService.GetLanguagesExecute(r)
}

/*
GetLanguages Get all languages

This method returns all available video languages.

 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @return ApiGetLanguagesRequest
*/
func (a *VideosLanguagesAPIService) GetLanguages(ctx context.Context) ApiGetLanguagesRequest {
	return ApiGetLanguagesRequest{
		ApiService: a,
		ctx: ctx,
	}
}

// Execute executes the request
//  @return []Language
func (a *VideosLanguagesAPIService) GetLanguagesExecute(r ApiGetLanguagesRequest) ([]Language, *http.Response, error) {
	var (
		localVarHTTPMethod   = http.MethodGet
		localVarPostBody     interface{}
		formFiles            []formFile
		localVarReturnValue  []Language
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "VideosLanguagesAPIService.GetLanguages")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/languages"

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}

	if r.filter != nil {
		parameterAddToHeaderOrQuery(localVarQueryParams, "filter", r.filter, "")
	}
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
	localVarHTTPHeaderAccepts := []string{"application/vnd.vimeo.language+json"}

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