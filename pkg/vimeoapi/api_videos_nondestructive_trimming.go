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


type VideosNondestructiveTrimmingAPI interface {

	/*
	ClipTrim Start a trim operation for a video

	This method starts a trim operation for the specified video.

	@param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
	@param videoId The ID of the video.
	@return ApiClipTrimRequest
	*/
	ClipTrim(ctx context.Context, videoId int32) ApiClipTrimRequest

	// ClipTrimExecute executes the request
	//  @return TrimmedVideo
	ClipTrimExecute(r ApiClipTrimRequest) (*TrimmedVideo, *http.Response, error)

	/*
	GetClipTrim Get the status of a video's trim operation

	This method returns the status of the trim operation for the specified video.

	@param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
	@param videoId The ID of the video.
	@return ApiGetClipTrimRequest
	*/
	GetClipTrim(ctx context.Context, videoId int32) ApiGetClipTrimRequest

	// GetClipTrimExecute executes the request
	//  @return TrimmedVideo
	GetClipTrimExecute(r ApiGetClipTrimRequest) (*TrimmedVideo, *http.Response, error)
}

// VideosNondestructiveTrimmingAPIService VideosNondestructiveTrimmingAPI service
type VideosNondestructiveTrimmingAPIService service

type ApiClipTrimRequest struct {
	ctx context.Context
	ApiService VideosNondestructiveTrimmingAPI
	videoId int32
	clipTrimRequest *ClipTrimRequest
}

func (r ApiClipTrimRequest) ClipTrimRequest(clipTrimRequest ClipTrimRequest) ApiClipTrimRequest {
	r.clipTrimRequest = &clipTrimRequest
	return r
}

func (r ApiClipTrimRequest) Execute() (*TrimmedVideo, *http.Response, error) {
	return r.ApiService.ClipTrimExecute(r)
}

/*
ClipTrim Start a trim operation for a video

This method starts a trim operation for the specified video.

 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @param videoId The ID of the video.
 @return ApiClipTrimRequest
*/
func (a *VideosNondestructiveTrimmingAPIService) ClipTrim(ctx context.Context, videoId int32) ApiClipTrimRequest {
	return ApiClipTrimRequest{
		ApiService: a,
		ctx: ctx,
		videoId: videoId,
	}
}

// Execute executes the request
//  @return TrimmedVideo
func (a *VideosNondestructiveTrimmingAPIService) ClipTrimExecute(r ApiClipTrimRequest) (*TrimmedVideo, *http.Response, error) {
	var (
		localVarHTTPMethod   = http.MethodPost
		localVarPostBody     interface{}
		formFiles            []formFile
		localVarReturnValue  *TrimmedVideo
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "VideosNondestructiveTrimmingAPIService.ClipTrim")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/videos/{video_id}/trim"
	localVarPath = strings.Replace(localVarPath, "{"+"video_id"+"}", url.PathEscape(parameterValueToString(r.videoId, "videoId")), -1)

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}

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
	localVarPostBody = r.clipTrimRequest
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

type ApiGetClipTrimRequest struct {
	ctx context.Context
	ApiService VideosNondestructiveTrimmingAPI
	videoId int32
}

func (r ApiGetClipTrimRequest) Execute() (*TrimmedVideo, *http.Response, error) {
	return r.ApiService.GetClipTrimExecute(r)
}

/*
GetClipTrim Get the status of a video's trim operation

This method returns the status of the trim operation for the specified video.

 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @param videoId The ID of the video.
 @return ApiGetClipTrimRequest
*/
func (a *VideosNondestructiveTrimmingAPIService) GetClipTrim(ctx context.Context, videoId int32) ApiGetClipTrimRequest {
	return ApiGetClipTrimRequest{
		ApiService: a,
		ctx: ctx,
		videoId: videoId,
	}
}

// Execute executes the request
//  @return TrimmedVideo
func (a *VideosNondestructiveTrimmingAPIService) GetClipTrimExecute(r ApiGetClipTrimRequest) (*TrimmedVideo, *http.Response, error) {
	var (
		localVarHTTPMethod   = http.MethodGet
		localVarPostBody     interface{}
		formFiles            []formFile
		localVarReturnValue  *TrimmedVideo
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "VideosNondestructiveTrimmingAPIService.GetClipTrim")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/videos/{video_id}/trim"
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
