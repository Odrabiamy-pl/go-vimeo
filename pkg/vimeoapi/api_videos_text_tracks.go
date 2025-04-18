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


type VideosTextTracksAPI interface {

	/*
	CreateTextTrack Add a text track to a video

	This method adds a text track to the specified video. For more information, see [Working with Text Track Uploads](https://developer.vimeo.com/api/upload/texttracks).

	@param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
	@param videoId The ID of the video.
	@return ApiCreateTextTrackRequest
	*/
	CreateTextTrack(ctx context.Context, videoId int32) ApiCreateTextTrackRequest

	// CreateTextTrackExecute executes the request
	//  @return TextTrack
	CreateTextTrackExecute(r ApiCreateTextTrackRequest) (*TextTrack, *http.Response, error)

	/*
	CreateTextTrackAlt1 Add a text track to a video

	This method adds a text track to the specified video. For more information, see [Working with Text Track Uploads](https://developer.vimeo.com/api/upload/texttracks).

	@param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
	@param channelId The ID of the channel.
	@param videoId The ID of the video.
	@return ApiCreateTextTrackAlt1Request
	*/
	CreateTextTrackAlt1(ctx context.Context, channelId float32, videoId int32) ApiCreateTextTrackAlt1Request

	// CreateTextTrackAlt1Execute executes the request
	//  @return TextTrack
	CreateTextTrackAlt1Execute(r ApiCreateTextTrackAlt1Request) (*TextTrack, *http.Response, error)

	/*
	DeleteTextTrack Delete a text track

	This method deletes the specified text track from a video. The authenticated user must be the owner of the video.

	@param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
	@param texttrackId The ID of the text track.
	@param videoId The ID of the video.
	@return ApiDeleteTextTrackRequest
	*/
	DeleteTextTrack(ctx context.Context, texttrackId float32, videoId int32) ApiDeleteTextTrackRequest

	// DeleteTextTrackExecute executes the request
	DeleteTextTrackExecute(r ApiDeleteTextTrackRequest) (*http.Response, error)

	/*
	EditTextTrack Edit a text track

	This method edits the specified text track of a video. The authenticated user must be the owner of the video.

	@param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
	@param texttrackId The ID of the text track.
	@param videoId The ID of the video.
	@return ApiEditTextTrackRequest
	*/
	EditTextTrack(ctx context.Context, texttrackId float32, videoId int32) ApiEditTextTrackRequest

	// EditTextTrackExecute executes the request
	//  @return TextTrack
	EditTextTrackExecute(r ApiEditTextTrackRequest) (*TextTrack, *http.Response, error)

	/*
	GetTextTrack Get a specific text track

	This method returns a single text track of the specified video. The authenticated user must be the owner of the video.

	@param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
	@param texttrackId The ID of the text track.
	@param videoId The ID of the video.
	@return ApiGetTextTrackRequest
	*/
	GetTextTrack(ctx context.Context, texttrackId float32, videoId int32) ApiGetTextTrackRequest

	// GetTextTrackExecute executes the request
	//  @return TextTrack
	GetTextTrackExecute(r ApiGetTextTrackRequest) (*TextTrack, *http.Response, error)

	/*
	GetTextTracks Get all the text tracks of a video

	This method returns every text track of the specified video. The authenticated user must be the owner of the video.

	@param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
	@param videoId The ID of the video.
	@return ApiGetTextTracksRequest
	*/
	GetTextTracks(ctx context.Context, videoId int32) ApiGetTextTracksRequest

	// GetTextTracksExecute executes the request
	//  @return []TextTrack
	GetTextTracksExecute(r ApiGetTextTracksRequest) ([]TextTrack, *http.Response, error)

	/*
	GetTextTracksAlt1 Get all the text tracks of a video

	This method returns every text track of the specified video. The authenticated user must be the owner of the video.

	@param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
	@param channelId The ID of the channel.
	@param videoId The ID of the video.
	@return ApiGetTextTracksAlt1Request
	*/
	GetTextTracksAlt1(ctx context.Context, channelId float32, videoId int32) ApiGetTextTracksAlt1Request

	// GetTextTracksAlt1Execute executes the request
	//  @return []TextTrack
	GetTextTracksAlt1Execute(r ApiGetTextTracksAlt1Request) ([]TextTrack, *http.Response, error)
}

// VideosTextTracksAPIService VideosTextTracksAPI service
type VideosTextTracksAPIService service

type ApiCreateTextTrackRequest struct {
	ctx context.Context
	ApiService VideosTextTracksAPI
	videoId int32
	createTextTrackAlt1Request *CreateTextTrackAlt1Request
}

func (r ApiCreateTextTrackRequest) CreateTextTrackAlt1Request(createTextTrackAlt1Request CreateTextTrackAlt1Request) ApiCreateTextTrackRequest {
	r.createTextTrackAlt1Request = &createTextTrackAlt1Request
	return r
}

func (r ApiCreateTextTrackRequest) Execute() (*TextTrack, *http.Response, error) {
	return r.ApiService.CreateTextTrackExecute(r)
}

/*
CreateTextTrack Add a text track to a video

This method adds a text track to the specified video. For more information, see [Working with Text Track Uploads](https://developer.vimeo.com/api/upload/texttracks).

 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @param videoId The ID of the video.
 @return ApiCreateTextTrackRequest
*/
func (a *VideosTextTracksAPIService) CreateTextTrack(ctx context.Context, videoId int32) ApiCreateTextTrackRequest {
	return ApiCreateTextTrackRequest{
		ApiService: a,
		ctx: ctx,
		videoId: videoId,
	}
}

// Execute executes the request
//  @return TextTrack
func (a *VideosTextTracksAPIService) CreateTextTrackExecute(r ApiCreateTextTrackRequest) (*TextTrack, *http.Response, error) {
	var (
		localVarHTTPMethod   = http.MethodPost
		localVarPostBody     interface{}
		formFiles            []formFile
		localVarReturnValue  *TextTrack
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "VideosTextTracksAPIService.CreateTextTrack")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/videos/{video_id}/texttracks"
	localVarPath = strings.Replace(localVarPath, "{"+"video_id"+"}", url.PathEscape(parameterValueToString(r.videoId, "videoId")), -1)

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}
	if r.createTextTrackAlt1Request == nil {
		return localVarReturnValue, nil, reportError("createTextTrackAlt1Request is required and must be specified")
	}

	// to determine the Content-Type header
	localVarHTTPContentTypes := []string{"application/vnd.vimeo.video.texttrack+json"}

	// set Content-Type header
	localVarHTTPContentType := selectHeaderContentType(localVarHTTPContentTypes)
	if localVarHTTPContentType != "" {
		localVarHeaderParams["Content-Type"] = localVarHTTPContentType
	}

	// to determine the Accept header
	localVarHTTPHeaderAccepts := []string{"application/vnd.vimeo.video.texttrack+json"}

	// set Accept header
	localVarHTTPHeaderAccept := selectHeaderAccept(localVarHTTPHeaderAccepts)
	if localVarHTTPHeaderAccept != "" {
		localVarHeaderParams["Accept"] = localVarHTTPHeaderAccept
	}
	// body params
	localVarPostBody = r.createTextTrackAlt1Request
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

type ApiCreateTextTrackAlt1Request struct {
	ctx context.Context
	ApiService VideosTextTracksAPI
	channelId float32
	videoId int32
	createTextTrackAlt1Request *CreateTextTrackAlt1Request
}

func (r ApiCreateTextTrackAlt1Request) CreateTextTrackAlt1Request(createTextTrackAlt1Request CreateTextTrackAlt1Request) ApiCreateTextTrackAlt1Request {
	r.createTextTrackAlt1Request = &createTextTrackAlt1Request
	return r
}

func (r ApiCreateTextTrackAlt1Request) Execute() (*TextTrack, *http.Response, error) {
	return r.ApiService.CreateTextTrackAlt1Execute(r)
}

/*
CreateTextTrackAlt1 Add a text track to a video

This method adds a text track to the specified video. For more information, see [Working with Text Track Uploads](https://developer.vimeo.com/api/upload/texttracks).

 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @param channelId The ID of the channel.
 @param videoId The ID of the video.
 @return ApiCreateTextTrackAlt1Request
*/
func (a *VideosTextTracksAPIService) CreateTextTrackAlt1(ctx context.Context, channelId float32, videoId int32) ApiCreateTextTrackAlt1Request {
	return ApiCreateTextTrackAlt1Request{
		ApiService: a,
		ctx: ctx,
		channelId: channelId,
		videoId: videoId,
	}
}

// Execute executes the request
//  @return TextTrack
func (a *VideosTextTracksAPIService) CreateTextTrackAlt1Execute(r ApiCreateTextTrackAlt1Request) (*TextTrack, *http.Response, error) {
	var (
		localVarHTTPMethod   = http.MethodPost
		localVarPostBody     interface{}
		formFiles            []formFile
		localVarReturnValue  *TextTrack
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "VideosTextTracksAPIService.CreateTextTrackAlt1")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/channels/{channel_id}/videos/{video_id}/texttracks"
	localVarPath = strings.Replace(localVarPath, "{"+"channel_id"+"}", url.PathEscape(parameterValueToString(r.channelId, "channelId")), -1)
	localVarPath = strings.Replace(localVarPath, "{"+"video_id"+"}", url.PathEscape(parameterValueToString(r.videoId, "videoId")), -1)

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}
	if r.createTextTrackAlt1Request == nil {
		return localVarReturnValue, nil, reportError("createTextTrackAlt1Request is required and must be specified")
	}

	// to determine the Content-Type header
	localVarHTTPContentTypes := []string{"application/vnd.vimeo.video.texttrack+json"}

	// set Content-Type header
	localVarHTTPContentType := selectHeaderContentType(localVarHTTPContentTypes)
	if localVarHTTPContentType != "" {
		localVarHeaderParams["Content-Type"] = localVarHTTPContentType
	}

	// to determine the Accept header
	localVarHTTPHeaderAccepts := []string{"application/vnd.vimeo.video.texttrack+json"}

	// set Accept header
	localVarHTTPHeaderAccept := selectHeaderAccept(localVarHTTPHeaderAccepts)
	if localVarHTTPHeaderAccept != "" {
		localVarHeaderParams["Accept"] = localVarHTTPHeaderAccept
	}
	// body params
	localVarPostBody = r.createTextTrackAlt1Request
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

type ApiDeleteTextTrackRequest struct {
	ctx context.Context
	ApiService VideosTextTracksAPI
	texttrackId float32
	videoId int32
}

func (r ApiDeleteTextTrackRequest) Execute() (*http.Response, error) {
	return r.ApiService.DeleteTextTrackExecute(r)
}

/*
DeleteTextTrack Delete a text track

This method deletes the specified text track from a video. The authenticated user must be the owner of the video.

 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @param texttrackId The ID of the text track.
 @param videoId The ID of the video.
 @return ApiDeleteTextTrackRequest
*/
func (a *VideosTextTracksAPIService) DeleteTextTrack(ctx context.Context, texttrackId float32, videoId int32) ApiDeleteTextTrackRequest {
	return ApiDeleteTextTrackRequest{
		ApiService: a,
		ctx: ctx,
		texttrackId: texttrackId,
		videoId: videoId,
	}
}

// Execute executes the request
func (a *VideosTextTracksAPIService) DeleteTextTrackExecute(r ApiDeleteTextTrackRequest) (*http.Response, error) {
	var (
		localVarHTTPMethod   = http.MethodDelete
		localVarPostBody     interface{}
		formFiles            []formFile
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "VideosTextTracksAPIService.DeleteTextTrack")
	if err != nil {
		return nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/videos/{video_id}/texttracks/{texttrack_id}"
	localVarPath = strings.Replace(localVarPath, "{"+"texttrack_id"+"}", url.PathEscape(parameterValueToString(r.texttrackId, "texttrackId")), -1)
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
	localVarHTTPHeaderAccepts := []string{"application/vnd.vimeo.video.texttrack+json"}

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
		if localVarHTTPResponse.StatusCode == 404 {
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

type ApiEditTextTrackRequest struct {
	ctx context.Context
	ApiService VideosTextTracksAPI
	texttrackId float32
	videoId int32
	editTextTrackRequest *EditTextTrackRequest
}

func (r ApiEditTextTrackRequest) EditTextTrackRequest(editTextTrackRequest EditTextTrackRequest) ApiEditTextTrackRequest {
	r.editTextTrackRequest = &editTextTrackRequest
	return r
}

func (r ApiEditTextTrackRequest) Execute() (*TextTrack, *http.Response, error) {
	return r.ApiService.EditTextTrackExecute(r)
}

/*
EditTextTrack Edit a text track

This method edits the specified text track of a video. The authenticated user must be the owner of the video.

 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @param texttrackId The ID of the text track.
 @param videoId The ID of the video.
 @return ApiEditTextTrackRequest
*/
func (a *VideosTextTracksAPIService) EditTextTrack(ctx context.Context, texttrackId float32, videoId int32) ApiEditTextTrackRequest {
	return ApiEditTextTrackRequest{
		ApiService: a,
		ctx: ctx,
		texttrackId: texttrackId,
		videoId: videoId,
	}
}

// Execute executes the request
//  @return TextTrack
func (a *VideosTextTracksAPIService) EditTextTrackExecute(r ApiEditTextTrackRequest) (*TextTrack, *http.Response, error) {
	var (
		localVarHTTPMethod   = http.MethodPatch
		localVarPostBody     interface{}
		formFiles            []formFile
		localVarReturnValue  *TextTrack
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "VideosTextTracksAPIService.EditTextTrack")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/videos/{video_id}/texttracks/{texttrack_id}"
	localVarPath = strings.Replace(localVarPath, "{"+"texttrack_id"+"}", url.PathEscape(parameterValueToString(r.texttrackId, "texttrackId")), -1)
	localVarPath = strings.Replace(localVarPath, "{"+"video_id"+"}", url.PathEscape(parameterValueToString(r.videoId, "videoId")), -1)

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}

	// to determine the Content-Type header
	localVarHTTPContentTypes := []string{"application/vnd.vimeo.video.texttrack+json"}

	// set Content-Type header
	localVarHTTPContentType := selectHeaderContentType(localVarHTTPContentTypes)
	if localVarHTTPContentType != "" {
		localVarHeaderParams["Content-Type"] = localVarHTTPContentType
	}

	// to determine the Accept header
	localVarHTTPHeaderAccepts := []string{"application/vnd.vimeo.video.texttrack+json"}

	// set Accept header
	localVarHTTPHeaderAccept := selectHeaderAccept(localVarHTTPHeaderAccepts)
	if localVarHTTPHeaderAccept != "" {
		localVarHeaderParams["Accept"] = localVarHTTPHeaderAccept
	}
	// body params
	localVarPostBody = r.editTextTrackRequest
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

type ApiGetTextTrackRequest struct {
	ctx context.Context
	ApiService VideosTextTracksAPI
	texttrackId float32
	videoId int32
}

func (r ApiGetTextTrackRequest) Execute() (*TextTrack, *http.Response, error) {
	return r.ApiService.GetTextTrackExecute(r)
}

/*
GetTextTrack Get a specific text track

This method returns a single text track of the specified video. The authenticated user must be the owner of the video.

 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @param texttrackId The ID of the text track.
 @param videoId The ID of the video.
 @return ApiGetTextTrackRequest
*/
func (a *VideosTextTracksAPIService) GetTextTrack(ctx context.Context, texttrackId float32, videoId int32) ApiGetTextTrackRequest {
	return ApiGetTextTrackRequest{
		ApiService: a,
		ctx: ctx,
		texttrackId: texttrackId,
		videoId: videoId,
	}
}

// Execute executes the request
//  @return TextTrack
func (a *VideosTextTracksAPIService) GetTextTrackExecute(r ApiGetTextTrackRequest) (*TextTrack, *http.Response, error) {
	var (
		localVarHTTPMethod   = http.MethodGet
		localVarPostBody     interface{}
		formFiles            []formFile
		localVarReturnValue  *TextTrack
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "VideosTextTracksAPIService.GetTextTrack")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/videos/{video_id}/texttracks/{texttrack_id}"
	localVarPath = strings.Replace(localVarPath, "{"+"texttrack_id"+"}", url.PathEscape(parameterValueToString(r.texttrackId, "texttrackId")), -1)
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
	localVarHTTPHeaderAccepts := []string{"application/vnd.vimeo.video.texttrack+json"}

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

type ApiGetTextTracksRequest struct {
	ctx context.Context
	ApiService VideosTextTracksAPI
	videoId int32
	page *float32
	perPage *float32
}

// The page number of the results to show.
func (r ApiGetTextTracksRequest) Page(page float32) ApiGetTextTracksRequest {
	r.page = &page
	return r
}

// The number of items to show on each page of results, up to a maximum of 100.
func (r ApiGetTextTracksRequest) PerPage(perPage float32) ApiGetTextTracksRequest {
	r.perPage = &perPage
	return r
}

func (r ApiGetTextTracksRequest) Execute() ([]TextTrack, *http.Response, error) {
	return r.ApiService.GetTextTracksExecute(r)
}

/*
GetTextTracks Get all the text tracks of a video

This method returns every text track of the specified video. The authenticated user must be the owner of the video.

 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @param videoId The ID of the video.
 @return ApiGetTextTracksRequest
*/
func (a *VideosTextTracksAPIService) GetTextTracks(ctx context.Context, videoId int32) ApiGetTextTracksRequest {
	return ApiGetTextTracksRequest{
		ApiService: a,
		ctx: ctx,
		videoId: videoId,
	}
}

// Execute executes the request
//  @return []TextTrack
func (a *VideosTextTracksAPIService) GetTextTracksExecute(r ApiGetTextTracksRequest) ([]TextTrack, *http.Response, error) {
	var (
		localVarHTTPMethod   = http.MethodGet
		localVarPostBody     interface{}
		formFiles            []formFile
		localVarReturnValue  []TextTrack
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "VideosTextTracksAPIService.GetTextTracks")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/videos/{video_id}/texttracks"
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
	localVarHTTPHeaderAccepts := []string{"application/vnd.vimeo.video.texttrack+json"}

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

type ApiGetTextTracksAlt1Request struct {
	ctx context.Context
	ApiService VideosTextTracksAPI
	channelId float32
	videoId int32
	page *float32
	perPage *float32
}

// The page number of the results to show.
func (r ApiGetTextTracksAlt1Request) Page(page float32) ApiGetTextTracksAlt1Request {
	r.page = &page
	return r
}

// The number of items to show on each page of results, up to a maximum of 100.
func (r ApiGetTextTracksAlt1Request) PerPage(perPage float32) ApiGetTextTracksAlt1Request {
	r.perPage = &perPage
	return r
}

func (r ApiGetTextTracksAlt1Request) Execute() ([]TextTrack, *http.Response, error) {
	return r.ApiService.GetTextTracksAlt1Execute(r)
}

/*
GetTextTracksAlt1 Get all the text tracks of a video

This method returns every text track of the specified video. The authenticated user must be the owner of the video.

 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @param channelId The ID of the channel.
 @param videoId The ID of the video.
 @return ApiGetTextTracksAlt1Request
*/
func (a *VideosTextTracksAPIService) GetTextTracksAlt1(ctx context.Context, channelId float32, videoId int32) ApiGetTextTracksAlt1Request {
	return ApiGetTextTracksAlt1Request{
		ApiService: a,
		ctx: ctx,
		channelId: channelId,
		videoId: videoId,
	}
}

// Execute executes the request
//  @return []TextTrack
func (a *VideosTextTracksAPIService) GetTextTracksAlt1Execute(r ApiGetTextTracksAlt1Request) ([]TextTrack, *http.Response, error) {
	var (
		localVarHTTPMethod   = http.MethodGet
		localVarPostBody     interface{}
		formFiles            []formFile
		localVarReturnValue  []TextTrack
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "VideosTextTracksAPIService.GetTextTracksAlt1")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/channels/{channel_id}/videos/{video_id}/texttracks"
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
	localVarHTTPHeaderAccepts := []string{"application/vnd.vimeo.video.texttrack+json"}

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
