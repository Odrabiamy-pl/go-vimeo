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


type VideosThumbnailsAPI interface {

	/*
	CreateVideoThumbnail Add a video thumbnail

	This method adds a thumbnail image to the specified video. The authenticated user must have team permissions for the video.

	@param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
	@param videoId The ID of the video.
	@return ApiCreateVideoThumbnailRequest
	*/
	CreateVideoThumbnail(ctx context.Context, videoId float32) ApiCreateVideoThumbnailRequest

	// CreateVideoThumbnailExecute executes the request
	//  @return Picture
	CreateVideoThumbnailExecute(r ApiCreateVideoThumbnailRequest) (*Picture, *http.Response, error)

	/*
	CreateVideoThumbnailAlt1 Add a video thumbnail

	This method adds a thumbnail image to the specified video. The authenticated user must have team permissions for the video.

	@param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
	@param channelId The ID of the channel.
	@param videoId The ID of the video.
	@return ApiCreateVideoThumbnailAlt1Request
	*/
	CreateVideoThumbnailAlt1(ctx context.Context, channelId float32, videoId float32) ApiCreateVideoThumbnailAlt1Request

	// CreateVideoThumbnailAlt1Execute executes the request
	//  @return Picture
	CreateVideoThumbnailAlt1Execute(r ApiCreateVideoThumbnailAlt1Request) (*Picture, *http.Response, error)

	/*
	DeleteVideoThumbnail Delete a video thumbnail

	This method deletes the specified thumbnail image from a video. The authenticated user must have team permissions for the video.

	@param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
	@param pictureId The ID of the thumbnail.
	@param videoId The ID of the video.
	@return ApiDeleteVideoThumbnailRequest
	*/
	DeleteVideoThumbnail(ctx context.Context, pictureId float32, videoId float32) ApiDeleteVideoThumbnailRequest

	// DeleteVideoThumbnailExecute executes the request
	DeleteVideoThumbnailExecute(r ApiDeleteVideoThumbnailRequest) (*http.Response, error)

	/*
	EditVideoThumbnail Edit a video thumbnail

	This method edits the specified video thumbnail image. The authenticated user must be the owner of the thumbnail.

	@param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
	@param pictureId The ID of the thumbnail.
	@param videoId The ID of the video.
	@return ApiEditVideoThumbnailRequest
	*/
	EditVideoThumbnail(ctx context.Context, pictureId float32, videoId float32) ApiEditVideoThumbnailRequest

	// EditVideoThumbnailExecute executes the request
	//  @return Picture
	EditVideoThumbnailExecute(r ApiEditVideoThumbnailRequest) (*Picture, *http.Response, error)

	/*
	GetVideoThumbnail Get a specific video thumbnail

	This method returns a single thumbnail image from the specified video. The authenticated user must have team permissions for the video.

	@param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
	@param pictureId The ID of the thumbnail.
	@param videoId The ID of the video.
	@return ApiGetVideoThumbnailRequest
	*/
	GetVideoThumbnail(ctx context.Context, pictureId float32, videoId float32) ApiGetVideoThumbnailRequest

	// GetVideoThumbnailExecute executes the request
	//  @return Picture
	GetVideoThumbnailExecute(r ApiGetVideoThumbnailRequest) (*Picture, *http.Response, error)

	/*
	GetVideoThumbnails Get all the thumbnails of a video

	This method returns all thumbnail images of the specified video. The authenticated user must have team permissions for the video.

	@param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
	@param videoId The ID of the video.
	@return ApiGetVideoThumbnailsRequest
	*/
	GetVideoThumbnails(ctx context.Context, videoId float32) ApiGetVideoThumbnailsRequest

	// GetVideoThumbnailsExecute executes the request
	//  @return []Picture
	GetVideoThumbnailsExecute(r ApiGetVideoThumbnailsRequest) ([]Picture, *http.Response, error)

	/*
	GetVideoThumbnailsAlt1 Get all the thumbnails of a video

	This method returns all thumbnail images of the specified video. The authenticated user must have team permissions for the video.

	@param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
	@param channelId The ID of the channel.
	@param videoId The ID of the video.
	@return ApiGetVideoThumbnailsAlt1Request
	*/
	GetVideoThumbnailsAlt1(ctx context.Context, channelId float32, videoId float32) ApiGetVideoThumbnailsAlt1Request

	// GetVideoThumbnailsAlt1Execute executes the request
	//  @return []Picture
	GetVideoThumbnailsAlt1Execute(r ApiGetVideoThumbnailsAlt1Request) ([]Picture, *http.Response, error)
}

// VideosThumbnailsAPIService VideosThumbnailsAPI service
type VideosThumbnailsAPIService service

type ApiCreateVideoThumbnailRequest struct {
	ctx context.Context
	ApiService VideosThumbnailsAPI
	videoId float32
	createVideoThumbnailAlt1Request *CreateVideoThumbnailAlt1Request
}

func (r ApiCreateVideoThumbnailRequest) CreateVideoThumbnailAlt1Request(createVideoThumbnailAlt1Request CreateVideoThumbnailAlt1Request) ApiCreateVideoThumbnailRequest {
	r.createVideoThumbnailAlt1Request = &createVideoThumbnailAlt1Request
	return r
}

func (r ApiCreateVideoThumbnailRequest) Execute() (*Picture, *http.Response, error) {
	return r.ApiService.CreateVideoThumbnailExecute(r)
}

/*
CreateVideoThumbnail Add a video thumbnail

This method adds a thumbnail image to the specified video. The authenticated user must have team permissions for the video.

 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @param videoId The ID of the video.
 @return ApiCreateVideoThumbnailRequest
*/
func (a *VideosThumbnailsAPIService) CreateVideoThumbnail(ctx context.Context, videoId float32) ApiCreateVideoThumbnailRequest {
	return ApiCreateVideoThumbnailRequest{
		ApiService: a,
		ctx: ctx,
		videoId: videoId,
	}
}

// Execute executes the request
//  @return Picture
func (a *VideosThumbnailsAPIService) CreateVideoThumbnailExecute(r ApiCreateVideoThumbnailRequest) (*Picture, *http.Response, error) {
	var (
		localVarHTTPMethod   = http.MethodPost
		localVarPostBody     interface{}
		formFiles            []formFile
		localVarReturnValue  *Picture
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "VideosThumbnailsAPIService.CreateVideoThumbnail")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/videos/{video_id}/pictures"
	localVarPath = strings.Replace(localVarPath, "{"+"video_id"+"}", url.PathEscape(parameterValueToString(r.videoId, "videoId")), -1)

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}

	// to determine the Content-Type header
	localVarHTTPContentTypes := []string{"application/vnd.vimeo.picture+json"}

	// set Content-Type header
	localVarHTTPContentType := selectHeaderContentType(localVarHTTPContentTypes)
	if localVarHTTPContentType != "" {
		localVarHeaderParams["Content-Type"] = localVarHTTPContentType
	}

	// to determine the Accept header
	localVarHTTPHeaderAccepts := []string{"application/vnd.vimeo.picture+json"}

	// set Accept header
	localVarHTTPHeaderAccept := selectHeaderAccept(localVarHTTPHeaderAccepts)
	if localVarHTTPHeaderAccept != "" {
		localVarHeaderParams["Accept"] = localVarHTTPHeaderAccept
	}
	// body params
	localVarPostBody = r.createVideoThumbnailAlt1Request
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

type ApiCreateVideoThumbnailAlt1Request struct {
	ctx context.Context
	ApiService VideosThumbnailsAPI
	channelId float32
	videoId float32
	createVideoThumbnailAlt1Request *CreateVideoThumbnailAlt1Request
}

func (r ApiCreateVideoThumbnailAlt1Request) CreateVideoThumbnailAlt1Request(createVideoThumbnailAlt1Request CreateVideoThumbnailAlt1Request) ApiCreateVideoThumbnailAlt1Request {
	r.createVideoThumbnailAlt1Request = &createVideoThumbnailAlt1Request
	return r
}

func (r ApiCreateVideoThumbnailAlt1Request) Execute() (*Picture, *http.Response, error) {
	return r.ApiService.CreateVideoThumbnailAlt1Execute(r)
}

/*
CreateVideoThumbnailAlt1 Add a video thumbnail

This method adds a thumbnail image to the specified video. The authenticated user must have team permissions for the video.

 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @param channelId The ID of the channel.
 @param videoId The ID of the video.
 @return ApiCreateVideoThumbnailAlt1Request
*/
func (a *VideosThumbnailsAPIService) CreateVideoThumbnailAlt1(ctx context.Context, channelId float32, videoId float32) ApiCreateVideoThumbnailAlt1Request {
	return ApiCreateVideoThumbnailAlt1Request{
		ApiService: a,
		ctx: ctx,
		channelId: channelId,
		videoId: videoId,
	}
}

// Execute executes the request
//  @return Picture
func (a *VideosThumbnailsAPIService) CreateVideoThumbnailAlt1Execute(r ApiCreateVideoThumbnailAlt1Request) (*Picture, *http.Response, error) {
	var (
		localVarHTTPMethod   = http.MethodPost
		localVarPostBody     interface{}
		formFiles            []formFile
		localVarReturnValue  *Picture
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "VideosThumbnailsAPIService.CreateVideoThumbnailAlt1")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/channels/{channel_id}/videos/{video_id}/pictures"
	localVarPath = strings.Replace(localVarPath, "{"+"channel_id"+"}", url.PathEscape(parameterValueToString(r.channelId, "channelId")), -1)
	localVarPath = strings.Replace(localVarPath, "{"+"video_id"+"}", url.PathEscape(parameterValueToString(r.videoId, "videoId")), -1)

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}

	// to determine the Content-Type header
	localVarHTTPContentTypes := []string{"application/vnd.vimeo.picture+json"}

	// set Content-Type header
	localVarHTTPContentType := selectHeaderContentType(localVarHTTPContentTypes)
	if localVarHTTPContentType != "" {
		localVarHeaderParams["Content-Type"] = localVarHTTPContentType
	}

	// to determine the Accept header
	localVarHTTPHeaderAccepts := []string{"application/vnd.vimeo.picture+json"}

	// set Accept header
	localVarHTTPHeaderAccept := selectHeaderAccept(localVarHTTPHeaderAccepts)
	if localVarHTTPHeaderAccept != "" {
		localVarHeaderParams["Accept"] = localVarHTTPHeaderAccept
	}
	// body params
	localVarPostBody = r.createVideoThumbnailAlt1Request
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

type ApiDeleteVideoThumbnailRequest struct {
	ctx context.Context
	ApiService VideosThumbnailsAPI
	pictureId float32
	videoId float32
}

func (r ApiDeleteVideoThumbnailRequest) Execute() (*http.Response, error) {
	return r.ApiService.DeleteVideoThumbnailExecute(r)
}

/*
DeleteVideoThumbnail Delete a video thumbnail

This method deletes the specified thumbnail image from a video. The authenticated user must have team permissions for the video.

 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @param pictureId The ID of the thumbnail.
 @param videoId The ID of the video.
 @return ApiDeleteVideoThumbnailRequest
*/
func (a *VideosThumbnailsAPIService) DeleteVideoThumbnail(ctx context.Context, pictureId float32, videoId float32) ApiDeleteVideoThumbnailRequest {
	return ApiDeleteVideoThumbnailRequest{
		ApiService: a,
		ctx: ctx,
		pictureId: pictureId,
		videoId: videoId,
	}
}

// Execute executes the request
func (a *VideosThumbnailsAPIService) DeleteVideoThumbnailExecute(r ApiDeleteVideoThumbnailRequest) (*http.Response, error) {
	var (
		localVarHTTPMethod   = http.MethodDelete
		localVarPostBody     interface{}
		formFiles            []formFile
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "VideosThumbnailsAPIService.DeleteVideoThumbnail")
	if err != nil {
		return nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/videos/{video_id}/pictures/{picture_id}"
	localVarPath = strings.Replace(localVarPath, "{"+"picture_id"+"}", url.PathEscape(parameterValueToString(r.pictureId, "pictureId")), -1)
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
	localVarHTTPHeaderAccepts := []string{}

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
		return localVarHTTPResponse, newErr
	}

	return localVarHTTPResponse, nil
}

type ApiEditVideoThumbnailRequest struct {
	ctx context.Context
	ApiService VideosThumbnailsAPI
	pictureId float32
	videoId float32
	editVideoThumbnailRequest *EditVideoThumbnailRequest
}

func (r ApiEditVideoThumbnailRequest) EditVideoThumbnailRequest(editVideoThumbnailRequest EditVideoThumbnailRequest) ApiEditVideoThumbnailRequest {
	r.editVideoThumbnailRequest = &editVideoThumbnailRequest
	return r
}

func (r ApiEditVideoThumbnailRequest) Execute() (*Picture, *http.Response, error) {
	return r.ApiService.EditVideoThumbnailExecute(r)
}

/*
EditVideoThumbnail Edit a video thumbnail

This method edits the specified video thumbnail image. The authenticated user must be the owner of the thumbnail.

 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @param pictureId The ID of the thumbnail.
 @param videoId The ID of the video.
 @return ApiEditVideoThumbnailRequest
*/
func (a *VideosThumbnailsAPIService) EditVideoThumbnail(ctx context.Context, pictureId float32, videoId float32) ApiEditVideoThumbnailRequest {
	return ApiEditVideoThumbnailRequest{
		ApiService: a,
		ctx: ctx,
		pictureId: pictureId,
		videoId: videoId,
	}
}

// Execute executes the request
//  @return Picture
func (a *VideosThumbnailsAPIService) EditVideoThumbnailExecute(r ApiEditVideoThumbnailRequest) (*Picture, *http.Response, error) {
	var (
		localVarHTTPMethod   = http.MethodPatch
		localVarPostBody     interface{}
		formFiles            []formFile
		localVarReturnValue  *Picture
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "VideosThumbnailsAPIService.EditVideoThumbnail")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/videos/{video_id}/pictures/{picture_id}"
	localVarPath = strings.Replace(localVarPath, "{"+"picture_id"+"}", url.PathEscape(parameterValueToString(r.pictureId, "pictureId")), -1)
	localVarPath = strings.Replace(localVarPath, "{"+"video_id"+"}", url.PathEscape(parameterValueToString(r.videoId, "videoId")), -1)

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}

	// to determine the Content-Type header
	localVarHTTPContentTypes := []string{"application/vnd.vimeo.picture+json"}

	// set Content-Type header
	localVarHTTPContentType := selectHeaderContentType(localVarHTTPContentTypes)
	if localVarHTTPContentType != "" {
		localVarHeaderParams["Content-Type"] = localVarHTTPContentType
	}

	// to determine the Accept header
	localVarHTTPHeaderAccepts := []string{"application/vnd.vimeo.picture+json"}

	// set Accept header
	localVarHTTPHeaderAccept := selectHeaderAccept(localVarHTTPHeaderAccepts)
	if localVarHTTPHeaderAccept != "" {
		localVarHeaderParams["Accept"] = localVarHTTPHeaderAccept
	}
	// body params
	localVarPostBody = r.editVideoThumbnailRequest
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

type ApiGetVideoThumbnailRequest struct {
	ctx context.Context
	ApiService VideosThumbnailsAPI
	pictureId float32
	videoId float32
}

func (r ApiGetVideoThumbnailRequest) Execute() (*Picture, *http.Response, error) {
	return r.ApiService.GetVideoThumbnailExecute(r)
}

/*
GetVideoThumbnail Get a specific video thumbnail

This method returns a single thumbnail image from the specified video. The authenticated user must have team permissions for the video.

 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @param pictureId The ID of the thumbnail.
 @param videoId The ID of the video.
 @return ApiGetVideoThumbnailRequest
*/
func (a *VideosThumbnailsAPIService) GetVideoThumbnail(ctx context.Context, pictureId float32, videoId float32) ApiGetVideoThumbnailRequest {
	return ApiGetVideoThumbnailRequest{
		ApiService: a,
		ctx: ctx,
		pictureId: pictureId,
		videoId: videoId,
	}
}

// Execute executes the request
//  @return Picture
func (a *VideosThumbnailsAPIService) GetVideoThumbnailExecute(r ApiGetVideoThumbnailRequest) (*Picture, *http.Response, error) {
	var (
		localVarHTTPMethod   = http.MethodGet
		localVarPostBody     interface{}
		formFiles            []formFile
		localVarReturnValue  *Picture
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "VideosThumbnailsAPIService.GetVideoThumbnail")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/videos/{video_id}/pictures/{picture_id}"
	localVarPath = strings.Replace(localVarPath, "{"+"picture_id"+"}", url.PathEscape(parameterValueToString(r.pictureId, "pictureId")), -1)
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
	localVarHTTPHeaderAccepts := []string{"application/vnd.vimeo.picture+json"}

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

type ApiGetVideoThumbnailsRequest struct {
	ctx context.Context
	ApiService VideosThumbnailsAPI
	videoId float32
	page *float32
	perPage *float32
}

// The page number of the results to show.
func (r ApiGetVideoThumbnailsRequest) Page(page float32) ApiGetVideoThumbnailsRequest {
	r.page = &page
	return r
}

// The number of items to show on each page of results, up to a maximum of 100.
func (r ApiGetVideoThumbnailsRequest) PerPage(perPage float32) ApiGetVideoThumbnailsRequest {
	r.perPage = &perPage
	return r
}

func (r ApiGetVideoThumbnailsRequest) Execute() ([]Picture, *http.Response, error) {
	return r.ApiService.GetVideoThumbnailsExecute(r)
}

/*
GetVideoThumbnails Get all the thumbnails of a video

This method returns all thumbnail images of the specified video. The authenticated user must have team permissions for the video.

 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @param videoId The ID of the video.
 @return ApiGetVideoThumbnailsRequest
*/
func (a *VideosThumbnailsAPIService) GetVideoThumbnails(ctx context.Context, videoId float32) ApiGetVideoThumbnailsRequest {
	return ApiGetVideoThumbnailsRequest{
		ApiService: a,
		ctx: ctx,
		videoId: videoId,
	}
}

// Execute executes the request
//  @return []Picture
func (a *VideosThumbnailsAPIService) GetVideoThumbnailsExecute(r ApiGetVideoThumbnailsRequest) ([]Picture, *http.Response, error) {
	var (
		localVarHTTPMethod   = http.MethodGet
		localVarPostBody     interface{}
		formFiles            []formFile
		localVarReturnValue  []Picture
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "VideosThumbnailsAPIService.GetVideoThumbnails")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/videos/{video_id}/pictures"
	localVarPath = strings.Replace(localVarPath, "{"+"video_id"+"}", url.PathEscape(parameterValueToString(r.videoId, "videoId")), -1)

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
	localVarHTTPHeaderAccepts := []string{"application/vnd.vimeo.picture+json"}

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

type ApiGetVideoThumbnailsAlt1Request struct {
	ctx context.Context
	ApiService VideosThumbnailsAPI
	channelId float32
	videoId float32
	page *float32
	perPage *float32
}

// The page number of the results to show.
func (r ApiGetVideoThumbnailsAlt1Request) Page(page float32) ApiGetVideoThumbnailsAlt1Request {
	r.page = &page
	return r
}

// The number of items to show on each page of results, up to a maximum of 100.
func (r ApiGetVideoThumbnailsAlt1Request) PerPage(perPage float32) ApiGetVideoThumbnailsAlt1Request {
	r.perPage = &perPage
	return r
}

func (r ApiGetVideoThumbnailsAlt1Request) Execute() ([]Picture, *http.Response, error) {
	return r.ApiService.GetVideoThumbnailsAlt1Execute(r)
}

/*
GetVideoThumbnailsAlt1 Get all the thumbnails of a video

This method returns all thumbnail images of the specified video. The authenticated user must have team permissions for the video.

 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @param channelId The ID of the channel.
 @param videoId The ID of the video.
 @return ApiGetVideoThumbnailsAlt1Request
*/
func (a *VideosThumbnailsAPIService) GetVideoThumbnailsAlt1(ctx context.Context, channelId float32, videoId float32) ApiGetVideoThumbnailsAlt1Request {
	return ApiGetVideoThumbnailsAlt1Request{
		ApiService: a,
		ctx: ctx,
		channelId: channelId,
		videoId: videoId,
	}
}

// Execute executes the request
//  @return []Picture
func (a *VideosThumbnailsAPIService) GetVideoThumbnailsAlt1Execute(r ApiGetVideoThumbnailsAlt1Request) ([]Picture, *http.Response, error) {
	var (
		localVarHTTPMethod   = http.MethodGet
		localVarPostBody     interface{}
		formFiles            []formFile
		localVarReturnValue  []Picture
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "VideosThumbnailsAPIService.GetVideoThumbnailsAlt1")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/channels/{channel_id}/videos/{video_id}/pictures"
	localVarPath = strings.Replace(localVarPath, "{"+"channel_id"+"}", url.PathEscape(parameterValueToString(r.channelId, "channelId")), -1)
	localVarPath = strings.Replace(localVarPath, "{"+"video_id"+"}", url.PathEscape(parameterValueToString(r.videoId, "videoId")), -1)

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
	localVarHTTPHeaderAccepts := []string{"application/vnd.vimeo.picture+json"}

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
