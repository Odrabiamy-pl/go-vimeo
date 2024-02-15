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
	"strings"
)


type CategoriesGroupsAPI interface {

	/*
	GetCategoryGroups Get all the groups in a category

	This method returns every group that belongs to the specified category.

	@param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
	@param category The name of the category.
	@return ApiGetCategoryGroupsRequest
	*/
	GetCategoryGroups(ctx context.Context, category string) ApiGetCategoryGroupsRequest

	// GetCategoryGroupsExecute executes the request
	//  @return []Group
	GetCategoryGroupsExecute(r ApiGetCategoryGroupsRequest) ([]Group, *http.Response, error)
}

// CategoriesGroupsAPIService CategoriesGroupsAPI service
type CategoriesGroupsAPIService service

type ApiGetCategoryGroupsRequest struct {
	ctx context.Context
	ApiService CategoriesGroupsAPI
	category string
	direction *string
	page *float32
	perPage *float32
	query *string
	sort *string
}

// The sort direction of the results.  Option descriptions:  * &#x60;asc&#x60; - Sort the results in ascending order.  * &#x60;desc&#x60; - Sort the results in descending order. 
func (r ApiGetCategoryGroupsRequest) Direction(direction string) ApiGetCategoryGroupsRequest {
	r.direction = &direction
	return r
}

// The page number of the results to show.
func (r ApiGetCategoryGroupsRequest) Page(page float32) ApiGetCategoryGroupsRequest {
	r.page = &page
	return r
}

// The number of items to show on each page of results, up to a maximum of 100.
func (r ApiGetCategoryGroupsRequest) PerPage(perPage float32) ApiGetCategoryGroupsRequest {
	r.perPage = &perPage
	return r
}

// The search query to use to filter the results.
func (r ApiGetCategoryGroupsRequest) Query(query string) ApiGetCategoryGroupsRequest {
	r.query = &query
	return r
}

// The way to sort the results.  Option descriptions:  * &#x60;alphabetical&#x60; - Sort the results alphabetically.  * &#x60;date&#x60; - Sort the results by date.  * &#x60;members&#x60; - Sort the results by number of members.  * &#x60;videos&#x60; - Sort the results by number of videos. 
func (r ApiGetCategoryGroupsRequest) Sort(sort string) ApiGetCategoryGroupsRequest {
	r.sort = &sort
	return r
}

func (r ApiGetCategoryGroupsRequest) Execute() ([]Group, *http.Response, error) {
	return r.ApiService.GetCategoryGroupsExecute(r)
}

/*
GetCategoryGroups Get all the groups in a category

This method returns every group that belongs to the specified category.

 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @param category The name of the category.
 @return ApiGetCategoryGroupsRequest
*/
func (a *CategoriesGroupsAPIService) GetCategoryGroups(ctx context.Context, category string) ApiGetCategoryGroupsRequest {
	return ApiGetCategoryGroupsRequest{
		ApiService: a,
		ctx: ctx,
		category: category,
	}
}

// Execute executes the request
//  @return []Group
func (a *CategoriesGroupsAPIService) GetCategoryGroupsExecute(r ApiGetCategoryGroupsRequest) ([]Group, *http.Response, error) {
	var (
		localVarHTTPMethod   = http.MethodGet
		localVarPostBody     interface{}
		formFiles            []formFile
		localVarReturnValue  []Group
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "CategoriesGroupsAPIService.GetCategoryGroups")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/categories/{category}/groups"
	localVarPath = strings.Replace(localVarPath, "{"+"category"+"}", url.PathEscape(parameterValueToString(r.category, "category")), -1)

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}

	if r.direction != nil {
		parameterAddToHeaderOrQuery(localVarQueryParams, "direction", r.direction, "")
	}
	if r.page != nil {
		parameterAddToHeaderOrQuery(localVarQueryParams, "page", r.page, "")
	}
	if r.perPage != nil {
		parameterAddToHeaderOrQuery(localVarQueryParams, "per_page", r.perPage, "")
	}
	if r.query != nil {
		parameterAddToHeaderOrQuery(localVarQueryParams, "query", r.query, "")
	}
	if r.sort != nil {
		parameterAddToHeaderOrQuery(localVarQueryParams, "sort", r.sort, "")
	}
	// to determine the Content-Type header
	localVarHTTPContentTypes := []string{}

	// set Content-Type header
	localVarHTTPContentType := selectHeaderContentType(localVarHTTPContentTypes)
	if localVarHTTPContentType != "" {
		localVarHeaderParams["Content-Type"] = localVarHTTPContentType
	}

	// to determine the Accept header
	localVarHTTPHeaderAccepts := []string{"application/vnd.vimeo.group+json"}

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
		if localVarHTTPResponse.StatusCode == 404 {
			var v LegacyError
			err = a.client.decode(&v, localVarBody, localVarHTTPResponse.Header.Get("Content-Type"))
			if err != nil {
				newErr.error = err.Error()
				return localVarReturnValue, localVarHTTPResponse, newErr
			}
					newErr.error = formatErrorMessage(localVarHTTPResponse.Status, &v)
					newErr.model = v
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