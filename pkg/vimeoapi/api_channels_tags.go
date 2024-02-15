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


type ChannelsTagsAPI interface {

	/*
	AddChannelTag Add a specific tag to a channel

	This method adds a single tag to the specified channel. The authenticated user must be the owner of the channel.

	@param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
	@param channelId The ID of the channel.
	@param word The word to use as the tag.
	@return ApiAddChannelTagRequest
	*/
	AddChannelTag(ctx context.Context, channelId float32, word string) ApiAddChannelTagRequest

	// AddChannelTagExecute executes the request
	AddChannelTagExecute(r ApiAddChannelTagRequest) (*http.Response, error)

	/*
	AddTagsToChannel Add a list of tags to a channel

	This method adds multiple tags to the specified channel. Include the tags as a JSON array in the body of the request using the **tag** field, like this: `[{ "tag": "funny" }, { "tag": "concert" }]`. The authenticated user must be the owner of the channel. For more information on batch requests like this one, see [Using Common Formats and Parameters](https://developer.vimeo.com/api/common-formats#working-with-batch-requests).

	@param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
	@param channelId The ID of the channel.
	@return ApiAddTagsToChannelRequest
	*/
	AddTagsToChannel(ctx context.Context, channelId float32) ApiAddTagsToChannelRequest

	// AddTagsToChannelExecute executes the request
	//  @return []Tag
	AddTagsToChannelExecute(r ApiAddTagsToChannelRequest) ([]Tag, *http.Response, error)

	/*
	CheckIfChannelHasTag Check if a tag has been added to a channel

	This method determines whether a tag has been added to the specified channel.

	@param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
	@param channelId The ID of the channel.
	@param word The word to use as the tag.
	@return ApiCheckIfChannelHasTagRequest
	*/
	CheckIfChannelHasTag(ctx context.Context, channelId float32, word string) ApiCheckIfChannelHasTagRequest

	// CheckIfChannelHasTagExecute executes the request
	CheckIfChannelHasTagExecute(r ApiCheckIfChannelHasTagRequest) (*http.Response, error)

	/*
	DeleteTagFromChannel Remove a tag from a channel

	This method removes a single tag from the specified channel. The authenticated user must be the owner of the channel.

	@param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
	@param channelId The ID of the channel.
	@param word The word to use as the tag.
	@return ApiDeleteTagFromChannelRequest
	*/
	DeleteTagFromChannel(ctx context.Context, channelId float32, word string) ApiDeleteTagFromChannelRequest

	// DeleteTagFromChannelExecute executes the request
	DeleteTagFromChannelExecute(r ApiDeleteTagFromChannelRequest) (*http.Response, error)

	/*
	GetChannelTags Get all the tags that have been added to a channel

	This method returns every tag that has been added to the specified channel.

	@param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
	@param channelId The ID of the channel.
	@return ApiGetChannelTagsRequest
	*/
	GetChannelTags(ctx context.Context, channelId float32) ApiGetChannelTagsRequest

	// GetChannelTagsExecute executes the request
	//  @return []Tag
	GetChannelTagsExecute(r ApiGetChannelTagsRequest) ([]Tag, *http.Response, error)
}

// ChannelsTagsAPIService ChannelsTagsAPI service
type ChannelsTagsAPIService service

type ApiAddChannelTagRequest struct {
	ctx context.Context
	ApiService ChannelsTagsAPI
	channelId float32
	word string
}

func (r ApiAddChannelTagRequest) Execute() (*http.Response, error) {
	return r.ApiService.AddChannelTagExecute(r)
}

/*
AddChannelTag Add a specific tag to a channel

This method adds a single tag to the specified channel. The authenticated user must be the owner of the channel.

 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @param channelId The ID of the channel.
 @param word The word to use as the tag.
 @return ApiAddChannelTagRequest
*/
func (a *ChannelsTagsAPIService) AddChannelTag(ctx context.Context, channelId float32, word string) ApiAddChannelTagRequest {
	return ApiAddChannelTagRequest{
		ApiService: a,
		ctx: ctx,
		channelId: channelId,
		word: word,
	}
}

// Execute executes the request
func (a *ChannelsTagsAPIService) AddChannelTagExecute(r ApiAddChannelTagRequest) (*http.Response, error) {
	var (
		localVarHTTPMethod   = http.MethodPut
		localVarPostBody     interface{}
		formFiles            []formFile
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "ChannelsTagsAPIService.AddChannelTag")
	if err != nil {
		return nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/channels/{channel_id}/tags/{word}"
	localVarPath = strings.Replace(localVarPath, "{"+"channel_id"+"}", url.PathEscape(parameterValueToString(r.channelId, "channelId")), -1)
	localVarPath = strings.Replace(localVarPath, "{"+"word"+"}", url.PathEscape(parameterValueToString(r.word, "word")), -1)

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
	localVarHTTPHeaderAccepts := []string{"application/vnd.vimeo.tag+json"}

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
		if localVarHTTPResponse.StatusCode == 400 {
			var v LegacyError
			err = a.client.decode(&v, localVarBody, localVarHTTPResponse.Header.Get("Content-Type"))
			if err != nil {
				newErr.error = err.Error()
				return localVarHTTPResponse, newErr
			}
					newErr.error = formatErrorMessage(localVarHTTPResponse.Status, &v)
					newErr.model = v
			return localVarHTTPResponse, newErr
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

type ApiAddTagsToChannelRequest struct {
	ctx context.Context
	ApiService ChannelsTagsAPI
	channelId float32
	addTagsToChannelRequest *AddTagsToChannelRequest
}

func (r ApiAddTagsToChannelRequest) AddTagsToChannelRequest(addTagsToChannelRequest AddTagsToChannelRequest) ApiAddTagsToChannelRequest {
	r.addTagsToChannelRequest = &addTagsToChannelRequest
	return r
}

func (r ApiAddTagsToChannelRequest) Execute() ([]Tag, *http.Response, error) {
	return r.ApiService.AddTagsToChannelExecute(r)
}

/*
AddTagsToChannel Add a list of tags to a channel

This method adds multiple tags to the specified channel. Include the tags as a JSON array in the body of the request using the **tag** field, like this: `[{ "tag": "funny" }, { "tag": "concert" }]`. The authenticated user must be the owner of the channel. For more information on batch requests like this one, see [Using Common Formats and Parameters](https://developer.vimeo.com/api/common-formats#working-with-batch-requests).

 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @param channelId The ID of the channel.
 @return ApiAddTagsToChannelRequest
*/
func (a *ChannelsTagsAPIService) AddTagsToChannel(ctx context.Context, channelId float32) ApiAddTagsToChannelRequest {
	return ApiAddTagsToChannelRequest{
		ApiService: a,
		ctx: ctx,
		channelId: channelId,
	}
}

// Execute executes the request
//  @return []Tag
func (a *ChannelsTagsAPIService) AddTagsToChannelExecute(r ApiAddTagsToChannelRequest) ([]Tag, *http.Response, error) {
	var (
		localVarHTTPMethod   = http.MethodPut
		localVarPostBody     interface{}
		formFiles            []formFile
		localVarReturnValue  []Tag
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "ChannelsTagsAPIService.AddTagsToChannel")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/channels/{channel_id}/tags"
	localVarPath = strings.Replace(localVarPath, "{"+"channel_id"+"}", url.PathEscape(parameterValueToString(r.channelId, "channelId")), -1)

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}
	if r.addTagsToChannelRequest == nil {
		return localVarReturnValue, nil, reportError("addTagsToChannelRequest is required and must be specified")
	}

	// to determine the Content-Type header
	localVarHTTPContentTypes := []string{"application/vnd.vimeo.tag+json"}

	// set Content-Type header
	localVarHTTPContentType := selectHeaderContentType(localVarHTTPContentTypes)
	if localVarHTTPContentType != "" {
		localVarHeaderParams["Content-Type"] = localVarHTTPContentType
	}

	// to determine the Accept header
	localVarHTTPHeaderAccepts := []string{"application/vnd.vimeo.tag+json"}

	// set Accept header
	localVarHTTPHeaderAccept := selectHeaderAccept(localVarHTTPHeaderAccepts)
	if localVarHTTPHeaderAccept != "" {
		localVarHeaderParams["Accept"] = localVarHTTPHeaderAccept
	}
	// body params
	localVarPostBody = r.addTagsToChannelRequest
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

type ApiCheckIfChannelHasTagRequest struct {
	ctx context.Context
	ApiService ChannelsTagsAPI
	channelId float32
	word string
}

func (r ApiCheckIfChannelHasTagRequest) Execute() (*http.Response, error) {
	return r.ApiService.CheckIfChannelHasTagExecute(r)
}

/*
CheckIfChannelHasTag Check if a tag has been added to a channel

This method determines whether a tag has been added to the specified channel.

 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @param channelId The ID of the channel.
 @param word The word to use as the tag.
 @return ApiCheckIfChannelHasTagRequest
*/
func (a *ChannelsTagsAPIService) CheckIfChannelHasTag(ctx context.Context, channelId float32, word string) ApiCheckIfChannelHasTagRequest {
	return ApiCheckIfChannelHasTagRequest{
		ApiService: a,
		ctx: ctx,
		channelId: channelId,
		word: word,
	}
}

// Execute executes the request
func (a *ChannelsTagsAPIService) CheckIfChannelHasTagExecute(r ApiCheckIfChannelHasTagRequest) (*http.Response, error) {
	var (
		localVarHTTPMethod   = http.MethodGet
		localVarPostBody     interface{}
		formFiles            []formFile
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "ChannelsTagsAPIService.CheckIfChannelHasTag")
	if err != nil {
		return nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/channels/{channel_id}/tags/{word}"
	localVarPath = strings.Replace(localVarPath, "{"+"channel_id"+"}", url.PathEscape(parameterValueToString(r.channelId, "channelId")), -1)
	localVarPath = strings.Replace(localVarPath, "{"+"word"+"}", url.PathEscape(parameterValueToString(r.word, "word")), -1)

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
	localVarHTTPHeaderAccepts := []string{"application/vnd.vimeo.tag+json"}

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
		if localVarHTTPResponse.StatusCode == 400 {
			var v LegacyError
			err = a.client.decode(&v, localVarBody, localVarHTTPResponse.Header.Get("Content-Type"))
			if err != nil {
				newErr.error = err.Error()
				return localVarHTTPResponse, newErr
			}
					newErr.error = formatErrorMessage(localVarHTTPResponse.Status, &v)
					newErr.model = v
			return localVarHTTPResponse, newErr
		}
		if localVarHTTPResponse.StatusCode == 404 {
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

type ApiDeleteTagFromChannelRequest struct {
	ctx context.Context
	ApiService ChannelsTagsAPI
	channelId float32
	word string
}

func (r ApiDeleteTagFromChannelRequest) Execute() (*http.Response, error) {
	return r.ApiService.DeleteTagFromChannelExecute(r)
}

/*
DeleteTagFromChannel Remove a tag from a channel

This method removes a single tag from the specified channel. The authenticated user must be the owner of the channel.

 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @param channelId The ID of the channel.
 @param word The word to use as the tag.
 @return ApiDeleteTagFromChannelRequest
*/
func (a *ChannelsTagsAPIService) DeleteTagFromChannel(ctx context.Context, channelId float32, word string) ApiDeleteTagFromChannelRequest {
	return ApiDeleteTagFromChannelRequest{
		ApiService: a,
		ctx: ctx,
		channelId: channelId,
		word: word,
	}
}

// Execute executes the request
func (a *ChannelsTagsAPIService) DeleteTagFromChannelExecute(r ApiDeleteTagFromChannelRequest) (*http.Response, error) {
	var (
		localVarHTTPMethod   = http.MethodDelete
		localVarPostBody     interface{}
		formFiles            []formFile
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "ChannelsTagsAPIService.DeleteTagFromChannel")
	if err != nil {
		return nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/channels/{channel_id}/tags/{word}"
	localVarPath = strings.Replace(localVarPath, "{"+"channel_id"+"}", url.PathEscape(parameterValueToString(r.channelId, "channelId")), -1)
	localVarPath = strings.Replace(localVarPath, "{"+"word"+"}", url.PathEscape(parameterValueToString(r.word, "word")), -1)

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
	localVarHTTPHeaderAccepts := []string{"application/vnd.vimeo.tag+json"}

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
		if localVarHTTPResponse.StatusCode == 400 {
			var v LegacyError
			err = a.client.decode(&v, localVarBody, localVarHTTPResponse.Header.Get("Content-Type"))
			if err != nil {
				newErr.error = err.Error()
				return localVarHTTPResponse, newErr
			}
					newErr.error = formatErrorMessage(localVarHTTPResponse.Status, &v)
					newErr.model = v
			return localVarHTTPResponse, newErr
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

type ApiGetChannelTagsRequest struct {
	ctx context.Context
	ApiService ChannelsTagsAPI
	channelId float32
}

func (r ApiGetChannelTagsRequest) Execute() ([]Tag, *http.Response, error) {
	return r.ApiService.GetChannelTagsExecute(r)
}

/*
GetChannelTags Get all the tags that have been added to a channel

This method returns every tag that has been added to the specified channel.

 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @param channelId The ID of the channel.
 @return ApiGetChannelTagsRequest
*/
func (a *ChannelsTagsAPIService) GetChannelTags(ctx context.Context, channelId float32) ApiGetChannelTagsRequest {
	return ApiGetChannelTagsRequest{
		ApiService: a,
		ctx: ctx,
		channelId: channelId,
	}
}

// Execute executes the request
//  @return []Tag
func (a *ChannelsTagsAPIService) GetChannelTagsExecute(r ApiGetChannelTagsRequest) ([]Tag, *http.Response, error) {
	var (
		localVarHTTPMethod   = http.MethodGet
		localVarPostBody     interface{}
		formFiles            []formFile
		localVarReturnValue  []Tag
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "ChannelsTagsAPIService.GetChannelTags")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/channels/{channel_id}/tags"
	localVarPath = strings.Replace(localVarPath, "{"+"channel_id"+"}", url.PathEscape(parameterValueToString(r.channelId, "channelId")), -1)

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
	localVarHTTPHeaderAccepts := []string{"application/vnd.vimeo.tag+json"}

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