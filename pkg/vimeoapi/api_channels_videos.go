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


type ChannelsVideosAPI interface {

	/*
	AddVideoToChannel Add a specific video to a channel

	This method adds a single video to the specified channel. The authenticated user must be a moderator of the channel.

	@param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
	@param channelId The ID of the channel.
	@param videoId The ID of the video.
	@return ApiAddVideoToChannelRequest
	*/
	AddVideoToChannel(ctx context.Context, channelId float32, videoId float32) ApiAddVideoToChannelRequest

	// AddVideoToChannelExecute executes the request
	AddVideoToChannelExecute(r ApiAddVideoToChannelRequest) (*http.Response, error)

	/*
	AddVideosToChannel Add a list of videos to a channel

	This method adds multiple videos to the specified channel. The authenticated user must be a moderator of the channel.

	@param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
	@param channelId The ID of the channel.
	@return ApiAddVideosToChannelRequest
	*/
	AddVideosToChannel(ctx context.Context, channelId float32) ApiAddVideosToChannelRequest

	// AddVideosToChannelExecute executes the request
	AddVideosToChannelExecute(r ApiAddVideosToChannelRequest) (*http.Response, error)

	/*
	DeleteVideoFromChannel Remove a specific video from a channel

	This method removes a single video from the specified channel. The authenticated user must be a moderator of the channel.

	@param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
	@param channelId The ID of the channel.
	@param videoId The ID of the video.
	@return ApiDeleteVideoFromChannelRequest
	*/
	DeleteVideoFromChannel(ctx context.Context, channelId float32, videoId float32) ApiDeleteVideoFromChannelRequest

	// DeleteVideoFromChannelExecute executes the request
	DeleteVideoFromChannelExecute(r ApiDeleteVideoFromChannelRequest) (*http.Response, error)

	/*
	GetAvailableVideoChannels Get all the channels to which the user can add or remove a specific video

	This method returns every channel to which the authenticated user can add or remove the specified video. The authenticated user must be a moderator of the channel.

	@param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
	@param videoId The ID of the video.
	@return ApiGetAvailableVideoChannelsRequest
	*/
	GetAvailableVideoChannels(ctx context.Context, videoId float32) ApiGetAvailableVideoChannelsRequest

	// GetAvailableVideoChannelsExecute executes the request
	//  @return []Channel
	GetAvailableVideoChannelsExecute(r ApiGetAvailableVideoChannelsRequest) ([]Channel, *http.Response, error)

	/*
	GetChannelVideo Get a specific video in a channel

	This method returns a single video in the specified channel. You can use it to determine whether the video is in the channel.

	@param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
	@param channelId The ID of the channel.
	@param videoId The ID of the video.
	@return ApiGetChannelVideoRequest
	*/
	GetChannelVideo(ctx context.Context, channelId float32, videoId float32) ApiGetChannelVideoRequest

	// GetChannelVideoExecute executes the request
	//  @return Video
	GetChannelVideoExecute(r ApiGetChannelVideoRequest) (*Video, *http.Response, error)

	/*
	GetChannelVideos Get all the videos in a channel

	This method returns every video in the specified channel.

	@param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
	@param channelId The ID of the channel.
	@return ApiGetChannelVideosRequest
	*/
	GetChannelVideos(ctx context.Context, channelId float32) ApiGetChannelVideosRequest

	// GetChannelVideosExecute executes the request
	//  @return []Video
	GetChannelVideosExecute(r ApiGetChannelVideosRequest) ([]Video, *http.Response, error)

	/*
	RemoveVideosFromChannel Remove a list of videos from a channel

	This method removes multiple videos from the specified channel. Include the videos by their URI as a JSON block in the body of the request using the **video_uri** field, like this: `[{ "video_uri": "/videos/1234" }, { "video_uri": "/videos/1235" }]`. The authenticated user must be a moderator of the channel. For more information on batch requests like this one, see [Using Common Formats and Parameters](https://developer.vimeo.com/api/common-formats#working-with-batch-requests).

	@param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
	@param channelId The ID of the channel.
	@return ApiRemoveVideosFromChannelRequest
	*/
	RemoveVideosFromChannel(ctx context.Context, channelId float32) ApiRemoveVideosFromChannelRequest

	// RemoveVideosFromChannelExecute executes the request
	RemoveVideosFromChannelExecute(r ApiRemoveVideosFromChannelRequest) (*http.Response, error)
}

// ChannelsVideosAPIService ChannelsVideosAPI service
type ChannelsVideosAPIService service

type ApiAddVideoToChannelRequest struct {
	ctx context.Context
	ApiService ChannelsVideosAPI
	channelId float32
	videoId float32
}

func (r ApiAddVideoToChannelRequest) Execute() (*http.Response, error) {
	return r.ApiService.AddVideoToChannelExecute(r)
}

/*
AddVideoToChannel Add a specific video to a channel

This method adds a single video to the specified channel. The authenticated user must be a moderator of the channel.

 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @param channelId The ID of the channel.
 @param videoId The ID of the video.
 @return ApiAddVideoToChannelRequest
*/
func (a *ChannelsVideosAPIService) AddVideoToChannel(ctx context.Context, channelId float32, videoId float32) ApiAddVideoToChannelRequest {
	return ApiAddVideoToChannelRequest{
		ApiService: a,
		ctx: ctx,
		channelId: channelId,
		videoId: videoId,
	}
}

// Execute executes the request
func (a *ChannelsVideosAPIService) AddVideoToChannelExecute(r ApiAddVideoToChannelRequest) (*http.Response, error) {
	var (
		localVarHTTPMethod   = http.MethodPut
		localVarPostBody     interface{}
		formFiles            []formFile
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "ChannelsVideosAPIService.AddVideoToChannel")
	if err != nil {
		return nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/channels/{channel_id}/videos/{video_id}"
	localVarPath = strings.Replace(localVarPath, "{"+"channel_id"+"}", url.PathEscape(parameterValueToString(r.channelId, "channelId")), -1)
	localVarPath = strings.Replace(localVarPath, "{"+"video_id"+"}", url.PathEscape(parameterValueToString(r.videoId, "videoId")), -1)

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
		if localVarHTTPResponse.StatusCode == 404 {
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

type ApiAddVideosToChannelRequest struct {
	ctx context.Context
	ApiService ChannelsVideosAPI
	channelId float32
	addVideosToChannelRequest *AddVideosToChannelRequest
}

func (r ApiAddVideosToChannelRequest) AddVideosToChannelRequest(addVideosToChannelRequest AddVideosToChannelRequest) ApiAddVideosToChannelRequest {
	r.addVideosToChannelRequest = &addVideosToChannelRequest
	return r
}

func (r ApiAddVideosToChannelRequest) Execute() (*http.Response, error) {
	return r.ApiService.AddVideosToChannelExecute(r)
}

/*
AddVideosToChannel Add a list of videos to a channel

This method adds multiple videos to the specified channel. The authenticated user must be a moderator of the channel.

 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @param channelId The ID of the channel.
 @return ApiAddVideosToChannelRequest
*/
func (a *ChannelsVideosAPIService) AddVideosToChannel(ctx context.Context, channelId float32) ApiAddVideosToChannelRequest {
	return ApiAddVideosToChannelRequest{
		ApiService: a,
		ctx: ctx,
		channelId: channelId,
	}
}

// Execute executes the request
func (a *ChannelsVideosAPIService) AddVideosToChannelExecute(r ApiAddVideosToChannelRequest) (*http.Response, error) {
	var (
		localVarHTTPMethod   = http.MethodPut
		localVarPostBody     interface{}
		formFiles            []formFile
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "ChannelsVideosAPIService.AddVideosToChannel")
	if err != nil {
		return nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/channels/{channel_id}/videos"
	localVarPath = strings.Replace(localVarPath, "{"+"channel_id"+"}", url.PathEscape(parameterValueToString(r.channelId, "channelId")), -1)

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}
	if r.addVideosToChannelRequest == nil {
		return nil, reportError("addVideosToChannelRequest is required and must be specified")
	}

	// to determine the Content-Type header
	localVarHTTPContentTypes := []string{"application/json"}

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
	// body params
	localVarPostBody = r.addVideosToChannelRequest
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
		if localVarHTTPResponse.StatusCode == 404 {
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

type ApiDeleteVideoFromChannelRequest struct {
	ctx context.Context
	ApiService ChannelsVideosAPI
	channelId float32
	videoId float32
}

func (r ApiDeleteVideoFromChannelRequest) Execute() (*http.Response, error) {
	return r.ApiService.DeleteVideoFromChannelExecute(r)
}

/*
DeleteVideoFromChannel Remove a specific video from a channel

This method removes a single video from the specified channel. The authenticated user must be a moderator of the channel.

 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @param channelId The ID of the channel.
 @param videoId The ID of the video.
 @return ApiDeleteVideoFromChannelRequest
*/
func (a *ChannelsVideosAPIService) DeleteVideoFromChannel(ctx context.Context, channelId float32, videoId float32) ApiDeleteVideoFromChannelRequest {
	return ApiDeleteVideoFromChannelRequest{
		ApiService: a,
		ctx: ctx,
		channelId: channelId,
		videoId: videoId,
	}
}

// Execute executes the request
func (a *ChannelsVideosAPIService) DeleteVideoFromChannelExecute(r ApiDeleteVideoFromChannelRequest) (*http.Response, error) {
	var (
		localVarHTTPMethod   = http.MethodDelete
		localVarPostBody     interface{}
		formFiles            []formFile
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "ChannelsVideosAPIService.DeleteVideoFromChannel")
	if err != nil {
		return nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/channels/{channel_id}/videos/{video_id}"
	localVarPath = strings.Replace(localVarPath, "{"+"channel_id"+"}", url.PathEscape(parameterValueToString(r.channelId, "channelId")), -1)
	localVarPath = strings.Replace(localVarPath, "{"+"video_id"+"}", url.PathEscape(parameterValueToString(r.videoId, "videoId")), -1)

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
		if localVarHTTPResponse.StatusCode == 404 {
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

type ApiGetAvailableVideoChannelsRequest struct {
	ctx context.Context
	ApiService ChannelsVideosAPI
	videoId float32
}

func (r ApiGetAvailableVideoChannelsRequest) Execute() ([]Channel, *http.Response, error) {
	return r.ApiService.GetAvailableVideoChannelsExecute(r)
}

/*
GetAvailableVideoChannels Get all the channels to which the user can add or remove a specific video

This method returns every channel to which the authenticated user can add or remove the specified video. The authenticated user must be a moderator of the channel.

 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @param videoId The ID of the video.
 @return ApiGetAvailableVideoChannelsRequest
*/
func (a *ChannelsVideosAPIService) GetAvailableVideoChannels(ctx context.Context, videoId float32) ApiGetAvailableVideoChannelsRequest {
	return ApiGetAvailableVideoChannelsRequest{
		ApiService: a,
		ctx: ctx,
		videoId: videoId,
	}
}

// Execute executes the request
//  @return []Channel
func (a *ChannelsVideosAPIService) GetAvailableVideoChannelsExecute(r ApiGetAvailableVideoChannelsRequest) ([]Channel, *http.Response, error) {
	var (
		localVarHTTPMethod   = http.MethodGet
		localVarPostBody     interface{}
		formFiles            []formFile
		localVarReturnValue  []Channel
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "ChannelsVideosAPIService.GetAvailableVideoChannels")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/videos/{video_id}/available_channels"
	localVarPath = strings.Replace(localVarPath, "{"+"video_id"+"}", url.PathEscape(parameterValueToString(r.videoId, "videoId")), -1)

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
	localVarHTTPHeaderAccepts := []string{"application/vnd.vimeo.channel+json"}

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

type ApiGetChannelVideoRequest struct {
	ctx context.Context
	ApiService ChannelsVideosAPI
	channelId float32
	videoId float32
}

func (r ApiGetChannelVideoRequest) Execute() (*Video, *http.Response, error) {
	return r.ApiService.GetChannelVideoExecute(r)
}

/*
GetChannelVideo Get a specific video in a channel

This method returns a single video in the specified channel. You can use it to determine whether the video is in the channel.

 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @param channelId The ID of the channel.
 @param videoId The ID of the video.
 @return ApiGetChannelVideoRequest
*/
func (a *ChannelsVideosAPIService) GetChannelVideo(ctx context.Context, channelId float32, videoId float32) ApiGetChannelVideoRequest {
	return ApiGetChannelVideoRequest{
		ApiService: a,
		ctx: ctx,
		channelId: channelId,
		videoId: videoId,
	}
}

// Execute executes the request
//  @return Video
func (a *ChannelsVideosAPIService) GetChannelVideoExecute(r ApiGetChannelVideoRequest) (*Video, *http.Response, error) {
	var (
		localVarHTTPMethod   = http.MethodGet
		localVarPostBody     interface{}
		formFiles            []formFile
		localVarReturnValue  *Video
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "ChannelsVideosAPIService.GetChannelVideo")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/channels/{channel_id}/videos/{video_id}"
	localVarPath = strings.Replace(localVarPath, "{"+"channel_id"+"}", url.PathEscape(parameterValueToString(r.channelId, "channelId")), -1)
	localVarPath = strings.Replace(localVarPath, "{"+"video_id"+"}", url.PathEscape(parameterValueToString(r.videoId, "videoId")), -1)

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
	localVarHTTPHeaderAccepts := []string{"application/vnd.vimeo.video+json"}

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

type ApiGetChannelVideosRequest struct {
	ctx context.Context
	ApiService ChannelsVideosAPI
	channelId float32
	containingUri *string
	direction *string
	filter *string
	filterEmbeddable *bool
	page *float32
	perPage *float32
	query *string
	sort *string
}

// The page that contains the video URI.
func (r ApiGetChannelVideosRequest) ContainingUri(containingUri string) ApiGetChannelVideosRequest {
	r.containingUri = &containingUri
	return r
}

// The sort direction of the results.  Option descriptions:  * &#x60;asc&#x60; - Sort the results in ascending order.  * &#x60;desc&#x60; - Sort the results in descending order. 
func (r ApiGetChannelVideosRequest) Direction(direction string) ApiGetChannelVideosRequest {
	r.direction = &direction
	return r
}

// The attribute by which to filter the results.  Option descriptions:  * &#x60;embeddable&#x60; - Return embeddable videos. 
func (r ApiGetChannelVideosRequest) Filter(filter string) ApiGetChannelVideosRequest {
	r.filter = &filter
	return r
}

// Whether to filter the results by embeddable videos (&#x60;true&#x60;) or non-embeddable videos (&#x60;false&#x60;). This parameter is required only when **filter** is &#x60;embeddable&#x60;.
func (r ApiGetChannelVideosRequest) FilterEmbeddable(filterEmbeddable bool) ApiGetChannelVideosRequest {
	r.filterEmbeddable = &filterEmbeddable
	return r
}

// The page number of the results to show.
func (r ApiGetChannelVideosRequest) Page(page float32) ApiGetChannelVideosRequest {
	r.page = &page
	return r
}

// The number of items to show on each page of results, up to a maximum of 100.
func (r ApiGetChannelVideosRequest) PerPage(perPage float32) ApiGetChannelVideosRequest {
	r.perPage = &perPage
	return r
}

// The search query to use to filter the results.
func (r ApiGetChannelVideosRequest) Query(query string) ApiGetChannelVideosRequest {
	r.query = &query
	return r
}

// The way to sort the results.  Option descriptions:  * &#x60;added&#x60; - Sort the results by date added.  * &#x60;alphabetical&#x60; - Sort the results alphabetically.  * &#x60;comments&#x60; - Sort the results by number of comments.  * &#x60;date&#x60; - Sort the results by creation date.  * &#x60;default&#x60; - Use the default sorting method.  * &#x60;duration&#x60; - Sort the results by duration.  * &#x60;likes&#x60; - Sort the results by number of likes.  * &#x60;manual&#x60; - Sort the results as the user has arranged them.  * &#x60;modified_time&#x60; - Sort the results by last modification.  * &#x60;plays&#x60; - Sort the results by number of plays. 
func (r ApiGetChannelVideosRequest) Sort(sort string) ApiGetChannelVideosRequest {
	r.sort = &sort
	return r
}

func (r ApiGetChannelVideosRequest) Execute() ([]Video, *http.Response, error) {
	return r.ApiService.GetChannelVideosExecute(r)
}

/*
GetChannelVideos Get all the videos in a channel

This method returns every video in the specified channel.

 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @param channelId The ID of the channel.
 @return ApiGetChannelVideosRequest
*/
func (a *ChannelsVideosAPIService) GetChannelVideos(ctx context.Context, channelId float32) ApiGetChannelVideosRequest {
	return ApiGetChannelVideosRequest{
		ApiService: a,
		ctx: ctx,
		channelId: channelId,
	}
}

// Execute executes the request
//  @return []Video
func (a *ChannelsVideosAPIService) GetChannelVideosExecute(r ApiGetChannelVideosRequest) ([]Video, *http.Response, error) {
	var (
		localVarHTTPMethod   = http.MethodGet
		localVarPostBody     interface{}
		formFiles            []formFile
		localVarReturnValue  []Video
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "ChannelsVideosAPIService.GetChannelVideos")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/channels/{channel_id}/videos"
	localVarPath = strings.Replace(localVarPath, "{"+"channel_id"+"}", url.PathEscape(parameterValueToString(r.channelId, "channelId")), -1)

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}

	if r.containingUri != nil {
		parameterAddToHeaderOrQuery(localVarQueryParams, "containing_uri", r.containingUri, "")
	}
	if r.direction != nil {
		parameterAddToHeaderOrQuery(localVarQueryParams, "direction", r.direction, "")
	}
	if r.filter != nil {
		parameterAddToHeaderOrQuery(localVarQueryParams, "filter", r.filter, "")
	}
	if r.filterEmbeddable != nil {
		parameterAddToHeaderOrQuery(localVarQueryParams, "filter_embeddable", r.filterEmbeddable, "")
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
	localVarHTTPHeaderAccepts := []string{"application/vnd.vimeo.video+json"}

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

type ApiRemoveVideosFromChannelRequest struct {
	ctx context.Context
	ApiService ChannelsVideosAPI
	channelId float32
	removeVideosFromChannelRequest *RemoveVideosFromChannelRequest
}

func (r ApiRemoveVideosFromChannelRequest) RemoveVideosFromChannelRequest(removeVideosFromChannelRequest RemoveVideosFromChannelRequest) ApiRemoveVideosFromChannelRequest {
	r.removeVideosFromChannelRequest = &removeVideosFromChannelRequest
	return r
}

func (r ApiRemoveVideosFromChannelRequest) Execute() (*http.Response, error) {
	return r.ApiService.RemoveVideosFromChannelExecute(r)
}

/*
RemoveVideosFromChannel Remove a list of videos from a channel

This method removes multiple videos from the specified channel. Include the videos by their URI as a JSON block in the body of the request using the **video_uri** field, like this: `[{ "video_uri": "/videos/1234" }, { "video_uri": "/videos/1235" }]`. The authenticated user must be a moderator of the channel. For more information on batch requests like this one, see [Using Common Formats and Parameters](https://developer.vimeo.com/api/common-formats#working-with-batch-requests).

 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @param channelId The ID of the channel.
 @return ApiRemoveVideosFromChannelRequest
*/
func (a *ChannelsVideosAPIService) RemoveVideosFromChannel(ctx context.Context, channelId float32) ApiRemoveVideosFromChannelRequest {
	return ApiRemoveVideosFromChannelRequest{
		ApiService: a,
		ctx: ctx,
		channelId: channelId,
	}
}

// Execute executes the request
func (a *ChannelsVideosAPIService) RemoveVideosFromChannelExecute(r ApiRemoveVideosFromChannelRequest) (*http.Response, error) {
	var (
		localVarHTTPMethod   = http.MethodDelete
		localVarPostBody     interface{}
		formFiles            []formFile
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "ChannelsVideosAPIService.RemoveVideosFromChannel")
	if err != nil {
		return nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/channels/{channel_id}/videos"
	localVarPath = strings.Replace(localVarPath, "{"+"channel_id"+"}", url.PathEscape(parameterValueToString(r.channelId, "channelId")), -1)

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}
	if r.removeVideosFromChannelRequest == nil {
		return nil, reportError("removeVideosFromChannelRequest is required and must be specified")
	}

	// to determine the Content-Type header
	localVarHTTPContentTypes := []string{"application/json"}

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
	// body params
	localVarPostBody = r.removeVideosFromChannelRequest
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
		if localVarHTTPResponse.StatusCode == 404 {
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