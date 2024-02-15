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


type GroupsEssentialsAPI interface {

	/*
	CreateGroup Create a group

	This method creates a new group.

	@param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
	@return ApiCreateGroupRequest
	*/
	CreateGroup(ctx context.Context) ApiCreateGroupRequest

	// CreateGroupExecute executes the request
	//  @return Group
	CreateGroupExecute(r ApiCreateGroupRequest) (*Group, *http.Response, error)

	/*
	DeleteGroup Delete a group

	This method deletes the specified group. The authenticated user must be the owner of the group.

	@param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
	@param groupId The ID of the group.
	@return ApiDeleteGroupRequest
	*/
	DeleteGroup(ctx context.Context, groupId float32) ApiDeleteGroupRequest

	// DeleteGroupExecute executes the request
	DeleteGroupExecute(r ApiDeleteGroupRequest) (*http.Response, error)

	/*
	GetGroup Get a specific group

	This method returns the specified group.

	@param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
	@param groupId The ID of the group.
	@return ApiGetGroupRequest
	*/
	GetGroup(ctx context.Context, groupId float32) ApiGetGroupRequest

	// GetGroupExecute executes the request
	//  @return Group
	GetGroupExecute(r ApiGetGroupRequest) (*Group, *http.Response, error)

	/*
	GetGroups Get all groups

	This method returns every available group.

	@param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
	@return ApiGetGroupsRequest
	*/
	GetGroups(ctx context.Context) ApiGetGroupsRequest

	// GetGroupsExecute executes the request
	//  @return []Group
	GetGroupsExecute(r ApiGetGroupsRequest) ([]Group, *http.Response, error)
}

// GroupsEssentialsAPIService GroupsEssentialsAPI service
type GroupsEssentialsAPIService service

type ApiCreateGroupRequest struct {
	ctx context.Context
	ApiService GroupsEssentialsAPI
	createGroupRequest *CreateGroupRequest
}

func (r ApiCreateGroupRequest) CreateGroupRequest(createGroupRequest CreateGroupRequest) ApiCreateGroupRequest {
	r.createGroupRequest = &createGroupRequest
	return r
}

func (r ApiCreateGroupRequest) Execute() (*Group, *http.Response, error) {
	return r.ApiService.CreateGroupExecute(r)
}

/*
CreateGroup Create a group

This method creates a new group.

 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @return ApiCreateGroupRequest
*/
func (a *GroupsEssentialsAPIService) CreateGroup(ctx context.Context) ApiCreateGroupRequest {
	return ApiCreateGroupRequest{
		ApiService: a,
		ctx: ctx,
	}
}

// Execute executes the request
//  @return Group
func (a *GroupsEssentialsAPIService) CreateGroupExecute(r ApiCreateGroupRequest) (*Group, *http.Response, error) {
	var (
		localVarHTTPMethod   = http.MethodPost
		localVarPostBody     interface{}
		formFiles            []formFile
		localVarReturnValue  *Group
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "GroupsEssentialsAPIService.CreateGroup")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/groups"

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}
	if r.createGroupRequest == nil {
		return localVarReturnValue, nil, reportError("createGroupRequest is required and must be specified")
	}

	// to determine the Content-Type header
	localVarHTTPContentTypes := []string{"application/vnd.vimeo.group+json"}

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
	// body params
	localVarPostBody = r.createGroupRequest
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
		if localVarHTTPResponse.StatusCode == 403 {
			var v LegacyError
			err = a.client.decode(&v, localVarBody, localVarHTTPResponse.Header.Get("Content-Type"))
			if err != nil {
				newErr.error = err.Error()
				return localVarReturnValue, localVarHTTPResponse, newErr
			}
					newErr.error = formatErrorMessage(localVarHTTPResponse.Status, &v)
					newErr.model = v
			return localVarReturnValue, localVarHTTPResponse, newErr
		}
		if localVarHTTPResponse.StatusCode == 400 {
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

type ApiDeleteGroupRequest struct {
	ctx context.Context
	ApiService GroupsEssentialsAPI
	groupId float32
}

func (r ApiDeleteGroupRequest) Execute() (*http.Response, error) {
	return r.ApiService.DeleteGroupExecute(r)
}

/*
DeleteGroup Delete a group

This method deletes the specified group. The authenticated user must be the owner of the group.

 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @param groupId The ID of the group.
 @return ApiDeleteGroupRequest
*/
func (a *GroupsEssentialsAPIService) DeleteGroup(ctx context.Context, groupId float32) ApiDeleteGroupRequest {
	return ApiDeleteGroupRequest{
		ApiService: a,
		ctx: ctx,
		groupId: groupId,
	}
}

// Execute executes the request
func (a *GroupsEssentialsAPIService) DeleteGroupExecute(r ApiDeleteGroupRequest) (*http.Response, error) {
	var (
		localVarHTTPMethod   = http.MethodDelete
		localVarPostBody     interface{}
		formFiles            []formFile
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "GroupsEssentialsAPIService.DeleteGroup")
	if err != nil {
		return nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/groups/{group_id}"
	localVarPath = strings.Replace(localVarPath, "{"+"group_id"+"}", url.PathEscape(parameterValueToString(r.groupId, "groupId")), -1)

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}

	// to determine the Content-Type header
	localVarHTTPContentTypes := []string{}

	// set Content-Type header
	localVarHTTPContentType := selectHeaderContentType(localVarHTTPContentTypes)
	if localVarHTTPContentType != "" {
		localVarHeaderParams["Content-Type"] = localVarHTTPContentType
	}

	// to determine the Accept header
	localVarHTTPHeaderAccepts := []string{"application/json"}

	// set Accept header
	localVarHTTPHeaderAccept := selectHeaderAccept(localVarHTTPHeaderAccepts)
	if localVarHTTPHeaderAccept != "" {
		localVarHeaderParams["Accept"] = localVarHTTPHeaderAccept
	}
	req, err := a.client.prepareRequest(r.ctx, localVarPath, localVarHTTPMethod, localVarPostBody, localVarHeaderParams, localVarQueryParams, localVarFormParams, formFiles)
	if err != nil {
		return nil, err
	}

	localVarHTTPResponse, err := a.client.callAPI(req)
	if err != nil || localVarHTTPResponse == nil {
		return localVarHTTPResponse, err
	}

	localVarBody, err := io.ReadAll(localVarHTTPResponse.Body)
	localVarHTTPResponse.Body.Close()
	localVarHTTPResponse.Body = io.NopCloser(bytes.NewBuffer(localVarBody))
	if err != nil {
		return localVarHTTPResponse, err
	}

	if localVarHTTPResponse.StatusCode >= 300 {
		newErr := &GenericOpenAPIError{
			body:  localVarBody,
			error: localVarHTTPResponse.Status,
		}
		if localVarHTTPResponse.StatusCode == 403 {
			var v LegacyError
			err = a.client.decode(&v, localVarBody, localVarHTTPResponse.Header.Get("Content-Type"))
			if err != nil {
				newErr.error = err.Error()
				return localVarHTTPResponse, newErr
			}
					newErr.error = formatErrorMessage(localVarHTTPResponse.Status, &v)
					newErr.model = v
		}
		return localVarHTTPResponse, newErr
	}

	return localVarHTTPResponse, nil
}

type ApiGetGroupRequest struct {
	ctx context.Context
	ApiService GroupsEssentialsAPI
	groupId float32
}

func (r ApiGetGroupRequest) Execute() (*Group, *http.Response, error) {
	return r.ApiService.GetGroupExecute(r)
}

/*
GetGroup Get a specific group

This method returns the specified group.

 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @param groupId The ID of the group.
 @return ApiGetGroupRequest
*/
func (a *GroupsEssentialsAPIService) GetGroup(ctx context.Context, groupId float32) ApiGetGroupRequest {
	return ApiGetGroupRequest{
		ApiService: a,
		ctx: ctx,
		groupId: groupId,
	}
}

// Execute executes the request
//  @return Group
func (a *GroupsEssentialsAPIService) GetGroupExecute(r ApiGetGroupRequest) (*Group, *http.Response, error) {
	var (
		localVarHTTPMethod   = http.MethodGet
		localVarPostBody     interface{}
		formFiles            []formFile
		localVarReturnValue  *Group
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "GroupsEssentialsAPIService.GetGroup")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/groups/{group_id}"
	localVarPath = strings.Replace(localVarPath, "{"+"group_id"+"}", url.PathEscape(parameterValueToString(r.groupId, "groupId")), -1)

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}

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

type ApiGetGroupsRequest struct {
	ctx context.Context
	ApiService GroupsEssentialsAPI
	direction *string
	filter *string
	page *float32
	perPage *float32
	query *string
	sort *string
}

// The sort direction of the results.  Option descriptions:  * &#x60;asc&#x60; - Sort the results in ascending order.  * &#x60;desc&#x60; - Sort the results in descending order. 
func (r ApiGetGroupsRequest) Direction(direction string) ApiGetGroupsRequest {
	r.direction = &direction
	return r
}

// The attribute by which to filter the results.  Option descriptions:  * &#x60;featured&#x60; - Return featured groups. 
func (r ApiGetGroupsRequest) Filter(filter string) ApiGetGroupsRequest {
	r.filter = &filter
	return r
}

// The page number of the results to show.
func (r ApiGetGroupsRequest) Page(page float32) ApiGetGroupsRequest {
	r.page = &page
	return r
}

// The number of items to show on each page of results, up to a maximum of 100.
func (r ApiGetGroupsRequest) PerPage(perPage float32) ApiGetGroupsRequest {
	r.perPage = &perPage
	return r
}

// The search query to use to filter the results.
func (r ApiGetGroupsRequest) Query(query string) ApiGetGroupsRequest {
	r.query = &query
	return r
}

// The way to sort the results.  Option descriptions:  * &#x60;alphabetical&#x60; - Sort the results alphabetically.  * &#x60;date&#x60; - Sort the results by creation date.  * &#x60;followers&#x60; - Sort the results by number of followers.  * &#x60;relevant&#x60; - Sort the results by relevance. This option is available for search queries only.  * &#x60;videos&#x60; - Sort the results by number of videos. 
func (r ApiGetGroupsRequest) Sort(sort string) ApiGetGroupsRequest {
	r.sort = &sort
	return r
}

func (r ApiGetGroupsRequest) Execute() ([]Group, *http.Response, error) {
	return r.ApiService.GetGroupsExecute(r)
}

/*
GetGroups Get all groups

This method returns every available group.

 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @return ApiGetGroupsRequest
*/
func (a *GroupsEssentialsAPIService) GetGroups(ctx context.Context) ApiGetGroupsRequest {
	return ApiGetGroupsRequest{
		ApiService: a,
		ctx: ctx,
	}
}

// Execute executes the request
//  @return []Group
func (a *GroupsEssentialsAPIService) GetGroupsExecute(r ApiGetGroupsRequest) ([]Group, *http.Response, error) {
	var (
		localVarHTTPMethod   = http.MethodGet
		localVarPostBody     interface{}
		formFiles            []formFile
		localVarReturnValue  []Group
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "GroupsEssentialsAPIService.GetGroups")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/groups"

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}

	if r.direction != nil {
		parameterAddToHeaderOrQuery(localVarQueryParams, "direction", r.direction, "")
	}
	if r.filter != nil {
		parameterAddToHeaderOrQuery(localVarQueryParams, "filter", r.filter, "")
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