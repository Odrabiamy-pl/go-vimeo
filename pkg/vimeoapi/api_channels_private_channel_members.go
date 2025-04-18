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


type ChannelsPrivateChannelMembersAPI interface {

	/*
	DeleteChannelPrivacyUser Restrict a user from accessing a private channel

	This method prevents a single user from being able to access the specified private channel. The authenticated user must be the owner of the channel.

	@param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
	@param channelId The ID of the channel.
	@param userId The ID of the user.
	@return ApiDeleteChannelPrivacyUserRequest
	*/
	DeleteChannelPrivacyUser(ctx context.Context, channelId float32, userId int32) ApiDeleteChannelPrivacyUserRequest

	// DeleteChannelPrivacyUserExecute executes the request
	DeleteChannelPrivacyUserExecute(r ApiDeleteChannelPrivacyUserRequest) (*http.Response, error)

	/*
	GetChannelPrivacyUsers Get all the users who can access a private channel

	This method returns all the users who have access to the specified private channel. The authenticated user must be the owner of the channel.

	@param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
	@param channelId The ID of the channel.
	@return ApiGetChannelPrivacyUsersRequest
	*/
	GetChannelPrivacyUsers(ctx context.Context, channelId float32) ApiGetChannelPrivacyUsersRequest

	// GetChannelPrivacyUsersExecute executes the request
	//  @return []User
	GetChannelPrivacyUsersExecute(r ApiGetChannelPrivacyUsersRequest) ([]User, *http.Response, error)

	/*
	SetChannelPrivacyUser Permit a specific user to access a private channel

	This method gives a single user access to the specified private channel. The authenticated user must be the owner of the channel.

	@param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
	@param channelId The ID of the channel.
	@param userId The ID of the user.
	@return ApiSetChannelPrivacyUserRequest
	*/
	SetChannelPrivacyUser(ctx context.Context, channelId float32, userId int32) ApiSetChannelPrivacyUserRequest

	// SetChannelPrivacyUserExecute executes the request
	SetChannelPrivacyUserExecute(r ApiSetChannelPrivacyUserRequest) (*http.Response, error)

	/*
	SetChannelPrivacyUsers Permit a list of users to access a private channel

	This method gives multiple users access to the specified private channel. The authenticated user must be the owner of the channel.

	@param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
	@param channelId The ID of the channel.
	@return ApiSetChannelPrivacyUsersRequest
	*/
	SetChannelPrivacyUsers(ctx context.Context, channelId float32) ApiSetChannelPrivacyUsersRequest

	// SetChannelPrivacyUsersExecute executes the request
	//  @return []User
	SetChannelPrivacyUsersExecute(r ApiSetChannelPrivacyUsersRequest) ([]User, *http.Response, error)
}

// ChannelsPrivateChannelMembersAPIService ChannelsPrivateChannelMembersAPI service
type ChannelsPrivateChannelMembersAPIService service

type ApiDeleteChannelPrivacyUserRequest struct {
	ctx context.Context
	ApiService ChannelsPrivateChannelMembersAPI
	channelId float32
	userId int32
}

func (r ApiDeleteChannelPrivacyUserRequest) Execute() (*http.Response, error) {
	return r.ApiService.DeleteChannelPrivacyUserExecute(r)
}

/*
DeleteChannelPrivacyUser Restrict a user from accessing a private channel

This method prevents a single user from being able to access the specified private channel. The authenticated user must be the owner of the channel.

 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @param channelId The ID of the channel.
 @param userId The ID of the user.
 @return ApiDeleteChannelPrivacyUserRequest
*/
func (a *ChannelsPrivateChannelMembersAPIService) DeleteChannelPrivacyUser(ctx context.Context, channelId float32, userId int32) ApiDeleteChannelPrivacyUserRequest {
	return ApiDeleteChannelPrivacyUserRequest{
		ApiService: a,
		ctx: ctx,
		channelId: channelId,
		userId: userId,
	}
}

// Execute executes the request
func (a *ChannelsPrivateChannelMembersAPIService) DeleteChannelPrivacyUserExecute(r ApiDeleteChannelPrivacyUserRequest) (*http.Response, error) {
	var (
		localVarHTTPMethod   = http.MethodDelete
		localVarPostBody     interface{}
		formFiles            []formFile
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "ChannelsPrivateChannelMembersAPIService.DeleteChannelPrivacyUser")
	if err != nil {
		return nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/channels/{channel_id}/privacy/users/{user_id}"
	localVarPath = strings.Replace(localVarPath, "{"+"channel_id"+"}", url.PathEscape(parameterValueToString(r.channelId, "channelId")), -1)
	localVarPath = strings.Replace(localVarPath, "{"+"user_id"+"}", url.PathEscape(parameterValueToString(r.userId, "userId")), -1)

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
		if localVarHTTPResponse.StatusCode == 401 {
			var v ModelError
			err = a.client.decode(&v, localVarBody, localVarHTTPResponse.Header.Get("Content-Type"))
			if err != nil {
				newErr.error = err.Error()
				return localVarHTTPResponse, newErr
			}
					newErr.error = formatErrorMessage(localVarHTTPResponse.Status, &v)
					newErr.model = v
			return localVarHTTPResponse, newErr
		}
		if localVarHTTPResponse.StatusCode == 403 {
			var v ModelError
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

type ApiGetChannelPrivacyUsersRequest struct {
	ctx context.Context
	ApiService ChannelsPrivateChannelMembersAPI
	channelId float32
	direction *string
	page *float32
	perPage *float32
}

// The sort direction of the results.  Option descriptions:  * &#x60;asc&#x60; - Sort the results in ascending order.  * &#x60;desc&#x60; - Sort the results in descending order. 
func (r ApiGetChannelPrivacyUsersRequest) Direction(direction string) ApiGetChannelPrivacyUsersRequest {
	r.direction = &direction
	return r
}

// The page number of the results to show.
func (r ApiGetChannelPrivacyUsersRequest) Page(page float32) ApiGetChannelPrivacyUsersRequest {
	r.page = &page
	return r
}

// The number of items to show on each page of results, up to a maximum of 100.
func (r ApiGetChannelPrivacyUsersRequest) PerPage(perPage float32) ApiGetChannelPrivacyUsersRequest {
	r.perPage = &perPage
	return r
}

func (r ApiGetChannelPrivacyUsersRequest) Execute() ([]User, *http.Response, error) {
	return r.ApiService.GetChannelPrivacyUsersExecute(r)
}

/*
GetChannelPrivacyUsers Get all the users who can access a private channel

This method returns all the users who have access to the specified private channel. The authenticated user must be the owner of the channel.

 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @param channelId The ID of the channel.
 @return ApiGetChannelPrivacyUsersRequest
*/
func (a *ChannelsPrivateChannelMembersAPIService) GetChannelPrivacyUsers(ctx context.Context, channelId float32) ApiGetChannelPrivacyUsersRequest {
	return ApiGetChannelPrivacyUsersRequest{
		ApiService: a,
		ctx: ctx,
		channelId: channelId,
	}
}

// Execute executes the request
//  @return []User
func (a *ChannelsPrivateChannelMembersAPIService) GetChannelPrivacyUsersExecute(r ApiGetChannelPrivacyUsersRequest) ([]User, *http.Response, error) {
	var (
		localVarHTTPMethod   = http.MethodGet
		localVarPostBody     interface{}
		formFiles            []formFile
		localVarReturnValue  []User
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "ChannelsPrivateChannelMembersAPIService.GetChannelPrivacyUsers")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/channels/{channel_id}/privacy/users"
	localVarPath = strings.Replace(localVarPath, "{"+"channel_id"+"}", url.PathEscape(parameterValueToString(r.channelId, "channelId")), -1)

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
	// to determine the Content-Type header
	localVarHTTPContentTypes := []string{}

	// set Content-Type header
	localVarHTTPContentType := selectHeaderContentType(localVarHTTPContentTypes)
	if localVarHTTPContentType != "" {
		localVarHeaderParams["Content-Type"] = localVarHTTPContentType
	}

	// to determine the Accept header
	localVarHTTPHeaderAccepts := []string{"application/vnd.vimeo.user+json"}

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
		if localVarHTTPResponse.StatusCode == 401 {
			var v ModelError
			err = a.client.decode(&v, localVarBody, localVarHTTPResponse.Header.Get("Content-Type"))
			if err != nil {
				newErr.error = err.Error()
				return localVarReturnValue, localVarHTTPResponse, newErr
			}
					newErr.error = formatErrorMessage(localVarHTTPResponse.Status, &v)
					newErr.model = v
			return localVarReturnValue, localVarHTTPResponse, newErr
		}
		if localVarHTTPResponse.StatusCode == 403 {
			var v ModelError
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

type ApiSetChannelPrivacyUserRequest struct {
	ctx context.Context
	ApiService ChannelsPrivateChannelMembersAPI
	channelId float32
	userId int32
}

func (r ApiSetChannelPrivacyUserRequest) Execute() (*http.Response, error) {
	return r.ApiService.SetChannelPrivacyUserExecute(r)
}

/*
SetChannelPrivacyUser Permit a specific user to access a private channel

This method gives a single user access to the specified private channel. The authenticated user must be the owner of the channel.

 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @param channelId The ID of the channel.
 @param userId The ID of the user.
 @return ApiSetChannelPrivacyUserRequest
*/
func (a *ChannelsPrivateChannelMembersAPIService) SetChannelPrivacyUser(ctx context.Context, channelId float32, userId int32) ApiSetChannelPrivacyUserRequest {
	return ApiSetChannelPrivacyUserRequest{
		ApiService: a,
		ctx: ctx,
		channelId: channelId,
		userId: userId,
	}
}

// Execute executes the request
func (a *ChannelsPrivateChannelMembersAPIService) SetChannelPrivacyUserExecute(r ApiSetChannelPrivacyUserRequest) (*http.Response, error) {
	var (
		localVarHTTPMethod   = http.MethodPut
		localVarPostBody     interface{}
		formFiles            []formFile
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "ChannelsPrivateChannelMembersAPIService.SetChannelPrivacyUser")
	if err != nil {
		return nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/channels/{channel_id}/privacy/users/{user_id}"
	localVarPath = strings.Replace(localVarPath, "{"+"channel_id"+"}", url.PathEscape(parameterValueToString(r.channelId, "channelId")), -1)
	localVarPath = strings.Replace(localVarPath, "{"+"user_id"+"}", url.PathEscape(parameterValueToString(r.userId, "userId")), -1)

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
		if localVarHTTPResponse.StatusCode == 401 {
			var v ModelError
			err = a.client.decode(&v, localVarBody, localVarHTTPResponse.Header.Get("Content-Type"))
			if err != nil {
				newErr.error = err.Error()
				return localVarHTTPResponse, newErr
			}
					newErr.error = formatErrorMessage(localVarHTTPResponse.Status, &v)
					newErr.model = v
			return localVarHTTPResponse, newErr
		}
		if localVarHTTPResponse.StatusCode == 403 {
			var v ModelError
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

type ApiSetChannelPrivacyUsersRequest struct {
	ctx context.Context
	ApiService ChannelsPrivateChannelMembersAPI
	channelId float32
	setChannelPrivacyUsersRequest *SetChannelPrivacyUsersRequest
}

func (r ApiSetChannelPrivacyUsersRequest) SetChannelPrivacyUsersRequest(setChannelPrivacyUsersRequest SetChannelPrivacyUsersRequest) ApiSetChannelPrivacyUsersRequest {
	r.setChannelPrivacyUsersRequest = &setChannelPrivacyUsersRequest
	return r
}

func (r ApiSetChannelPrivacyUsersRequest) Execute() ([]User, *http.Response, error) {
	return r.ApiService.SetChannelPrivacyUsersExecute(r)
}

/*
SetChannelPrivacyUsers Permit a list of users to access a private channel

This method gives multiple users access to the specified private channel. The authenticated user must be the owner of the channel.

 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @param channelId The ID of the channel.
 @return ApiSetChannelPrivacyUsersRequest
*/
func (a *ChannelsPrivateChannelMembersAPIService) SetChannelPrivacyUsers(ctx context.Context, channelId float32) ApiSetChannelPrivacyUsersRequest {
	return ApiSetChannelPrivacyUsersRequest{
		ApiService: a,
		ctx: ctx,
		channelId: channelId,
	}
}

// Execute executes the request
//  @return []User
func (a *ChannelsPrivateChannelMembersAPIService) SetChannelPrivacyUsersExecute(r ApiSetChannelPrivacyUsersRequest) ([]User, *http.Response, error) {
	var (
		localVarHTTPMethod   = http.MethodPut
		localVarPostBody     interface{}
		formFiles            []formFile
		localVarReturnValue  []User
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "ChannelsPrivateChannelMembersAPIService.SetChannelPrivacyUsers")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/channels/{channel_id}/privacy/users"
	localVarPath = strings.Replace(localVarPath, "{"+"channel_id"+"}", url.PathEscape(parameterValueToString(r.channelId, "channelId")), -1)

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}
	if r.setChannelPrivacyUsersRequest == nil {
		return localVarReturnValue, nil, reportError("setChannelPrivacyUsersRequest is required and must be specified")
	}

	// to determine the Content-Type header
	localVarHTTPContentTypes := []string{"application/vnd.vimeo.user+json"}

	// set Content-Type header
	localVarHTTPContentType := selectHeaderContentType(localVarHTTPContentTypes)
	if localVarHTTPContentType != "" {
		localVarHeaderParams["Content-Type"] = localVarHTTPContentType
	}

	// to determine the Accept header
	localVarHTTPHeaderAccepts := []string{"application/vnd.vimeo.user+json"}

	// set Accept header
	localVarHTTPHeaderAccept := selectHeaderAccept(localVarHTTPHeaderAccepts)
	if localVarHTTPHeaderAccept != "" {
		localVarHeaderParams["Accept"] = localVarHTTPHeaderAccept
	}
	// body params
	localVarPostBody = r.setChannelPrivacyUsersRequest
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
		if localVarHTTPResponse.StatusCode == 400 {
			var v ModelError
			err = a.client.decode(&v, localVarBody, localVarHTTPResponse.Header.Get("Content-Type"))
			if err != nil {
				newErr.error = err.Error()
				return localVarReturnValue, localVarHTTPResponse, newErr
			}
					newErr.error = formatErrorMessage(localVarHTTPResponse.Status, &v)
					newErr.model = v
			return localVarReturnValue, localVarHTTPResponse, newErr
		}
		if localVarHTTPResponse.StatusCode == 401 {
			var v ModelError
			err = a.client.decode(&v, localVarBody, localVarHTTPResponse.Header.Get("Content-Type"))
			if err != nil {
				newErr.error = err.Error()
				return localVarReturnValue, localVarHTTPResponse, newErr
			}
					newErr.error = formatErrorMessage(localVarHTTPResponse.Status, &v)
					newErr.model = v
			return localVarReturnValue, localVarHTTPResponse, newErr
		}
		if localVarHTTPResponse.StatusCode == 403 {
			var v ModelError
			err = a.client.decode(&v, localVarBody, localVarHTTPResponse.Header.Get("Content-Type"))
			if err != nil {
				newErr.error = err.Error()
				return localVarReturnValue, localVarHTTPResponse, newErr
			}
					newErr.error = formatErrorMessage(localVarHTTPResponse.Status, &v)
					newErr.model = v
			return localVarReturnValue, localVarHTTPResponse, newErr
		}
		if localVarHTTPResponse.StatusCode == 404 {
			var v ModelError
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
