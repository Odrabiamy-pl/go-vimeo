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


type OnDemandVideosAPI interface {

	/*
	AddVideoToVod Add a video to an On Demand page

	This method adds a video to the specified On Demand page. The authenticated user must be the owner of the page.

	@param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
	@param ondemandId The ID of the On Demand page.
	@param videoId The ID of the video.
	@return ApiAddVideoToVodRequest
	*/
	AddVideoToVod(ctx context.Context, ondemandId float32, videoId float32) ApiAddVideoToVodRequest

	// AddVideoToVodExecute executes the request
	//  @return OnDemandVideo
	AddVideoToVodExecute(r ApiAddVideoToVodRequest) (*OnDemandVideo, *http.Response, error)

	/*
	DeleteVideoFromVod Remove a video from an On Demand page

	This method removes a video from the specified On Demand page. The authenticated user must be the owner of the page.

	@param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
	@param ondemandId The ID of the On Demand page.
	@param videoId The ID of the video.
	@return ApiDeleteVideoFromVodRequest
	*/
	DeleteVideoFromVod(ctx context.Context, ondemandId float32, videoId float32) ApiDeleteVideoFromVodRequest

	// DeleteVideoFromVodExecute executes the request
	DeleteVideoFromVodExecute(r ApiDeleteVideoFromVodRequest) (*http.Response, error)

	/*
	GetVodVideo Get a specific video on an On Demand page

	This method returns a single video on the specified On Demand page. Use this method to determine whether the video is on the page.

	@param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
	@param ondemandId The ID of the On Demand page.
	@param videoId The ID of the video.
	@return ApiGetVodVideoRequest
	*/
	GetVodVideo(ctx context.Context, ondemandId float32, videoId float32) ApiGetVodVideoRequest

	// GetVodVideoExecute executes the request
	//  @return Video
	GetVodVideoExecute(r ApiGetVodVideoRequest) (*Video, *http.Response, error)

	/*
	GetVodVideos Get all the videos on an On Demand page

	This method returns every video on the specified On Demand page.

	@param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
	@param ondemandId The ID of the On Demand page.
	@return ApiGetVodVideosRequest
	*/
	GetVodVideos(ctx context.Context, ondemandId float32) ApiGetVodVideosRequest

	// GetVodVideosExecute executes the request
	//  @return []OnDemandVideo
	GetVodVideosExecute(r ApiGetVodVideosRequest) ([]OnDemandVideo, *http.Response, error)
}

// OnDemandVideosAPIService OnDemandVideosAPI service
type OnDemandVideosAPIService service

type ApiAddVideoToVodRequest struct {
	ctx context.Context
	ApiService OnDemandVideosAPI
	ondemandId float32
	videoId float32
	addVideoToVodRequest *AddVideoToVodRequest
}

func (r ApiAddVideoToVodRequest) AddVideoToVodRequest(addVideoToVodRequest AddVideoToVodRequest) ApiAddVideoToVodRequest {
	r.addVideoToVodRequest = &addVideoToVodRequest
	return r
}

func (r ApiAddVideoToVodRequest) Execute() (*OnDemandVideo, *http.Response, error) {
	return r.ApiService.AddVideoToVodExecute(r)
}

/*
AddVideoToVod Add a video to an On Demand page

This method adds a video to the specified On Demand page. The authenticated user must be the owner of the page.

 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @param ondemandId The ID of the On Demand page.
 @param videoId The ID of the video.
 @return ApiAddVideoToVodRequest
*/
func (a *OnDemandVideosAPIService) AddVideoToVod(ctx context.Context, ondemandId float32, videoId float32) ApiAddVideoToVodRequest {
	return ApiAddVideoToVodRequest{
		ApiService: a,
		ctx: ctx,
		ondemandId: ondemandId,
		videoId: videoId,
	}
}

// Execute executes the request
//  @return OnDemandVideo
func (a *OnDemandVideosAPIService) AddVideoToVodExecute(r ApiAddVideoToVodRequest) (*OnDemandVideo, *http.Response, error) {
	var (
		localVarHTTPMethod   = http.MethodPut
		localVarPostBody     interface{}
		formFiles            []formFile
		localVarReturnValue  *OnDemandVideo
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "OnDemandVideosAPIService.AddVideoToVod")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/ondemand/pages/{ondemand_id}/videos/{video_id}"
	localVarPath = strings.Replace(localVarPath, "{"+"ondemand_id"+"}", url.PathEscape(parameterValueToString(r.ondemandId, "ondemandId")), -1)
	localVarPath = strings.Replace(localVarPath, "{"+"video_id"+"}", url.PathEscape(parameterValueToString(r.videoId, "videoId")), -1)

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}
	if r.addVideoToVodRequest == nil {
		return localVarReturnValue, nil, reportError("addVideoToVodRequest is required and must be specified")
	}

	// to determine the Content-Type header
	localVarHTTPContentTypes := []string{"application/vnd.vimeo.ondemand.video+json"}

	// set Content-Type header
	localVarHTTPContentType := selectHeaderContentType(localVarHTTPContentTypes)
	if localVarHTTPContentType != "" {
		localVarHeaderParams["Content-Type"] = localVarHTTPContentType
	}

	// to determine the Accept header
	localVarHTTPHeaderAccepts := []string{"application/vnd.vimeo.ondemand.video+json"}

	// set Accept header
	localVarHTTPHeaderAccept := selectHeaderAccept(localVarHTTPHeaderAccepts)
	if localVarHTTPHeaderAccept != "" {
		localVarHeaderParams["Accept"] = localVarHTTPHeaderAccept
	}
	// body params
	localVarPostBody = r.addVideoToVodRequest
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

type ApiDeleteVideoFromVodRequest struct {
	ctx context.Context
	ApiService OnDemandVideosAPI
	ondemandId float32
	videoId float32
}

func (r ApiDeleteVideoFromVodRequest) Execute() (*http.Response, error) {
	return r.ApiService.DeleteVideoFromVodExecute(r)
}

/*
DeleteVideoFromVod Remove a video from an On Demand page

This method removes a video from the specified On Demand page. The authenticated user must be the owner of the page.

 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @param ondemandId The ID of the On Demand page.
 @param videoId The ID of the video.
 @return ApiDeleteVideoFromVodRequest
*/
func (a *OnDemandVideosAPIService) DeleteVideoFromVod(ctx context.Context, ondemandId float32, videoId float32) ApiDeleteVideoFromVodRequest {
	return ApiDeleteVideoFromVodRequest{
		ApiService: a,
		ctx: ctx,
		ondemandId: ondemandId,
		videoId: videoId,
	}
}

// Execute executes the request
func (a *OnDemandVideosAPIService) DeleteVideoFromVodExecute(r ApiDeleteVideoFromVodRequest) (*http.Response, error) {
	var (
		localVarHTTPMethod   = http.MethodDelete
		localVarPostBody     interface{}
		formFiles            []formFile
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "OnDemandVideosAPIService.DeleteVideoFromVod")
	if err != nil {
		return nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/ondemand/pages/{ondemand_id}/videos/{video_id}"
	localVarPath = strings.Replace(localVarPath, "{"+"ondemand_id"+"}", url.PathEscape(parameterValueToString(r.ondemandId, "ondemandId")), -1)
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
	localVarHTTPHeaderAccepts := []string{"application/vnd.vimeo.ondemand.video+json"}

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

type ApiGetVodVideoRequest struct {
	ctx context.Context
	ApiService OnDemandVideosAPI
	ondemandId float32
	videoId float32
}

func (r ApiGetVodVideoRequest) Execute() (*Video, *http.Response, error) {
	return r.ApiService.GetVodVideoExecute(r)
}

/*
GetVodVideo Get a specific video on an On Demand page

This method returns a single video on the specified On Demand page. Use this method to determine whether the video is on the page.

 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @param ondemandId The ID of the On Demand page.
 @param videoId The ID of the video.
 @return ApiGetVodVideoRequest
*/
func (a *OnDemandVideosAPIService) GetVodVideo(ctx context.Context, ondemandId float32, videoId float32) ApiGetVodVideoRequest {
	return ApiGetVodVideoRequest{
		ApiService: a,
		ctx: ctx,
		ondemandId: ondemandId,
		videoId: videoId,
	}
}

// Execute executes the request
//  @return Video
func (a *OnDemandVideosAPIService) GetVodVideoExecute(r ApiGetVodVideoRequest) (*Video, *http.Response, error) {
	var (
		localVarHTTPMethod   = http.MethodGet
		localVarPostBody     interface{}
		formFiles            []formFile
		localVarReturnValue  *Video
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "OnDemandVideosAPIService.GetVodVideo")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/ondemand/pages/{ondemand_id}/videos/{video_id}"
	localVarPath = strings.Replace(localVarPath, "{"+"ondemand_id"+"}", url.PathEscape(parameterValueToString(r.ondemandId, "ondemandId")), -1)
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
	localVarHTTPHeaderAccepts := []string{"application/vnd.vimeo.ondemand.video+json"}

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

type ApiGetVodVideosRequest struct {
	ctx context.Context
	ApiService OnDemandVideosAPI
	ondemandId float32
	direction *string
	filter *string
	page *float32
	perPage *float32
	sort *string
}

// The sort direction of the results.  Option descriptions:  * &#x60;asc&#x60; - Sort the results in ascending order.  * &#x60;desc&#x60; - Sort the results in descending order. 
func (r ApiGetVodVideosRequest) Direction(direction string) ApiGetVodVideosRequest {
	r.direction = &direction
	return r
}

// The attribute by which to filter the results.  Option descriptions:  * &#x60;all&#x60; - Filter for all videos.  * &#x60;buy&#x60; - Filter for purchased videos.  * &#x60;expiring_soon&#x60; - Filter for videos that expire soon.  * &#x60;extra&#x60; - Filter for extra footage videos.  * &#x60;main&#x60; - Filter for main videos.  * &#x60;main.viewable&#x60; - Filter for videos that are both the main video and are viewable.  * &#x60;rent&#x60; - Filter for rented videos.  * &#x60;trailer&#x60; - Filter for trailer videos.  * &#x60;unwatched&#x60; - Filter for unwatched videos.  * &#x60;viewable&#x60; - Filter for videos that are viewable.  * &#x60;watched&#x60; - Filter for watched videos. 
func (r ApiGetVodVideosRequest) Filter(filter string) ApiGetVodVideosRequest {
	r.filter = &filter
	return r
}

// The page number of the results to show.
func (r ApiGetVodVideosRequest) Page(page float32) ApiGetVodVideosRequest {
	r.page = &page
	return r
}

// The number of items to show on each page of results, up to a maximum of 100.
func (r ApiGetVodVideosRequest) PerPage(perPage float32) ApiGetVodVideosRequest {
	r.perPage = &perPage
	return r
}

// The way to sort the results.  Option descriptions:  * &#x60;date&#x60; - Sort the results by date.  * &#x60;default&#x60; - Use the default sorting method.  * &#x60;episode&#x60; - Sort the results by episode.  * &#x60;manual&#x60; - Sort the results manually.  * &#x60;name&#x60; - Sort the results by name.  * &#x60;purchase_time&#x60; - Sort the results by time of purchase.  * &#x60;release_date&#x60; - Sort the results by release date. 
func (r ApiGetVodVideosRequest) Sort(sort string) ApiGetVodVideosRequest {
	r.sort = &sort
	return r
}

func (r ApiGetVodVideosRequest) Execute() ([]OnDemandVideo, *http.Response, error) {
	return r.ApiService.GetVodVideosExecute(r)
}

/*
GetVodVideos Get all the videos on an On Demand page

This method returns every video on the specified On Demand page.

 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @param ondemandId The ID of the On Demand page.
 @return ApiGetVodVideosRequest
*/
func (a *OnDemandVideosAPIService) GetVodVideos(ctx context.Context, ondemandId float32) ApiGetVodVideosRequest {
	return ApiGetVodVideosRequest{
		ApiService: a,
		ctx: ctx,
		ondemandId: ondemandId,
	}
}

// Execute executes the request
//  @return []OnDemandVideo
func (a *OnDemandVideosAPIService) GetVodVideosExecute(r ApiGetVodVideosRequest) ([]OnDemandVideo, *http.Response, error) {
	var (
		localVarHTTPMethod   = http.MethodGet
		localVarPostBody     interface{}
		formFiles            []formFile
		localVarReturnValue  []OnDemandVideo
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "OnDemandVideosAPIService.GetVodVideos")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/ondemand/pages/{ondemand_id}/videos"
	localVarPath = strings.Replace(localVarPath, "{"+"ondemand_id"+"}", url.PathEscape(parameterValueToString(r.ondemandId, "ondemandId")), -1)

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
	localVarHTTPHeaderAccepts := []string{"application/vnd.vimeo.ondemand.video+json"}

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