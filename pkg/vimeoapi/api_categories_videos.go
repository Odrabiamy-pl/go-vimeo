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


type CategoriesVideosAPI interface {

	/*
	CheckCategoryForVideo Get a specific video in a category

	This method returns a single video in the specified category. You can use this method to determine whether the video belongs to the category.

	@param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
	@param category The name of the category.
	@param videoId The ID of the video.
	@return ApiCheckCategoryForVideoRequest
	*/
	CheckCategoryForVideo(ctx context.Context, category string, videoId int32) ApiCheckCategoryForVideoRequest

	// CheckCategoryForVideoExecute executes the request
	//  @return Video
	CheckCategoryForVideoExecute(r ApiCheckCategoryForVideoRequest) (*Video, *http.Response, error)

	/*
	GetCategoryVideos Get all the videos in a category

	This method returns every video that belongs to the specified category.

	@param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
	@param category The name of the category.
	@return ApiGetCategoryVideosRequest
	*/
	GetCategoryVideos(ctx context.Context, category string) ApiGetCategoryVideosRequest

	// GetCategoryVideosExecute executes the request
	//  @return []Video
	GetCategoryVideosExecute(r ApiGetCategoryVideosRequest) ([]Video, *http.Response, error)

	/*
	GetVideoCategories Get all the categories to which a video belongs

	This method returns every category that contains the specified video.

	@param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
	@param videoId The ID of the video.
	@return ApiGetVideoCategoriesRequest
	*/
	GetVideoCategories(ctx context.Context, videoId int32) ApiGetVideoCategoriesRequest

	// GetVideoCategoriesExecute executes the request
	//  @return []Category
	GetVideoCategoriesExecute(r ApiGetVideoCategoriesRequest) ([]Category, *http.Response, error)

	/*
	SuggestVideoCategory Suggest categories for a video

	This method sets multiple categories and subcategories for the specified video. Include the categories as a JSON block in the body of the request using the **category** field, like this: `[{ "category": "Tech" }, { "category": "Music" }]`. The authenticated user must have edit access to the video. For more information on batch requests like this one, see [Using Common Formats and Parameters](https://developer.vimeo.com/api/common-formats#working-with-batch-requests).

	@param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
	@param videoId The ID of the video.
	@return ApiSuggestVideoCategoryRequest
	*/
	SuggestVideoCategory(ctx context.Context, videoId int32) ApiSuggestVideoCategoryRequest

	// SuggestVideoCategoryExecute executes the request
	//  @return Category
	SuggestVideoCategoryExecute(r ApiSuggestVideoCategoryRequest) (*Category, *http.Response, error)
}

// CategoriesVideosAPIService CategoriesVideosAPI service
type CategoriesVideosAPIService service

type ApiCheckCategoryForVideoRequest struct {
	ctx context.Context
	ApiService CategoriesVideosAPI
	category string
	videoId int32
}

func (r ApiCheckCategoryForVideoRequest) Execute() (*Video, *http.Response, error) {
	return r.ApiService.CheckCategoryForVideoExecute(r)
}

/*
CheckCategoryForVideo Get a specific video in a category

This method returns a single video in the specified category. You can use this method to determine whether the video belongs to the category.

 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @param category The name of the category.
 @param videoId The ID of the video.
 @return ApiCheckCategoryForVideoRequest
*/
func (a *CategoriesVideosAPIService) CheckCategoryForVideo(ctx context.Context, category string, videoId int32) ApiCheckCategoryForVideoRequest {
	return ApiCheckCategoryForVideoRequest{
		ApiService: a,
		ctx: ctx,
		category: category,
		videoId: videoId,
	}
}

// Execute executes the request
//  @return Video
func (a *CategoriesVideosAPIService) CheckCategoryForVideoExecute(r ApiCheckCategoryForVideoRequest) (*Video, *http.Response, error) {
	var (
		localVarHTTPMethod   = http.MethodGet
		localVarPostBody     interface{}
		formFiles            []formFile
		localVarReturnValue  *Video
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "CategoriesVideosAPIService.CheckCategoryForVideo")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/categories/{category}/videos/{video_id}"
	localVarPath = strings.Replace(localVarPath, "{"+"category"+"}", url.PathEscape(parameterValueToString(r.category, "category")), -1)
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

type ApiGetCategoryVideosRequest struct {
	ctx context.Context
	ApiService CategoriesVideosAPI
	category string
	direction *string
	filter *string
	filterEmbeddable *bool
	page *float32
	perPage *float32
	query *string
	sort *string
}

// The sort direction of the results.  Option descriptions:  * &#x60;asc&#x60; - Sort the results in ascending order.  * &#x60;desc&#x60; - Sort the results in descending order. 
func (r ApiGetCategoryVideosRequest) Direction(direction string) ApiGetCategoryVideosRequest {
	r.direction = &direction
	return r
}

// The attribute by which to filter the results.  Option descriptions:  * &#x60;conditional_featured&#x60; - Return featured videos.  * &#x60;embeddable&#x60; - Return embeddable videos. 
func (r ApiGetCategoryVideosRequest) Filter(filter string) ApiGetCategoryVideosRequest {
	r.filter = &filter
	return r
}

// Whether to filter the results by embeddable videos (&#x60;true&#x60;) or non-embeddable videos (&#x60;false&#x60;). This parameter is required only when **filter** is &#x60;embeddable&#x60;.
func (r ApiGetCategoryVideosRequest) FilterEmbeddable(filterEmbeddable bool) ApiGetCategoryVideosRequest {
	r.filterEmbeddable = &filterEmbeddable
	return r
}

// The page number of the results to show.
func (r ApiGetCategoryVideosRequest) Page(page float32) ApiGetCategoryVideosRequest {
	r.page = &page
	return r
}

// The number of items to show on each page of results, up to a maximum of 100.
func (r ApiGetCategoryVideosRequest) PerPage(perPage float32) ApiGetCategoryVideosRequest {
	r.perPage = &perPage
	return r
}

// The search query to use to filter the results.
func (r ApiGetCategoryVideosRequest) Query(query string) ApiGetCategoryVideosRequest {
	r.query = &query
	return r
}

// The way to sort the results.  Option descriptions:  * &#x60;alphabetical&#x60; - Sort the results alphabetically.  * &#x60;comments&#x60; - Sort the results by number of comments.  * &#x60;date&#x60; - Sort the results by date.  * &#x60;duration&#x60; - Sort the results by duration.  * &#x60;featured&#x60; - Sort the results by featured status.  * &#x60;likes&#x60; - Sort the results by number of likes.  * &#x60;plays&#x60; - Sort the results by number of plays.  * &#x60;relevant&#x60; - Sort the results by relevance. 
func (r ApiGetCategoryVideosRequest) Sort(sort string) ApiGetCategoryVideosRequest {
	r.sort = &sort
	return r
}

func (r ApiGetCategoryVideosRequest) Execute() ([]Video, *http.Response, error) {
	return r.ApiService.GetCategoryVideosExecute(r)
}

/*
GetCategoryVideos Get all the videos in a category

This method returns every video that belongs to the specified category.

 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @param category The name of the category.
 @return ApiGetCategoryVideosRequest
*/
func (a *CategoriesVideosAPIService) GetCategoryVideos(ctx context.Context, category string) ApiGetCategoryVideosRequest {
	return ApiGetCategoryVideosRequest{
		ApiService: a,
		ctx: ctx,
		category: category,
	}
}

// Execute executes the request
//  @return []Video
func (a *CategoriesVideosAPIService) GetCategoryVideosExecute(r ApiGetCategoryVideosRequest) ([]Video, *http.Response, error) {
	var (
		localVarHTTPMethod   = http.MethodGet
		localVarPostBody     interface{}
		formFiles            []formFile
		localVarReturnValue  []Video
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "CategoriesVideosAPIService.GetCategoryVideos")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/categories/{category}/videos"
	localVarPath = strings.Replace(localVarPath, "{"+"category"+"}", url.PathEscape(parameterValueToString(r.category, "category")), -1)

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}

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

type ApiGetVideoCategoriesRequest struct {
	ctx context.Context
	ApiService CategoriesVideosAPI
	videoId int32
	page *float32
	perPage *float32
}

// The page number of the results to show.
func (r ApiGetVideoCategoriesRequest) Page(page float32) ApiGetVideoCategoriesRequest {
	r.page = &page
	return r
}

// The number of items to show on each page of results, up to a maximum of 100.
func (r ApiGetVideoCategoriesRequest) PerPage(perPage float32) ApiGetVideoCategoriesRequest {
	r.perPage = &perPage
	return r
}

func (r ApiGetVideoCategoriesRequest) Execute() ([]Category, *http.Response, error) {
	return r.ApiService.GetVideoCategoriesExecute(r)
}

/*
GetVideoCategories Get all the categories to which a video belongs

This method returns every category that contains the specified video.

 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @param videoId The ID of the video.
 @return ApiGetVideoCategoriesRequest
*/
func (a *CategoriesVideosAPIService) GetVideoCategories(ctx context.Context, videoId int32) ApiGetVideoCategoriesRequest {
	return ApiGetVideoCategoriesRequest{
		ApiService: a,
		ctx: ctx,
		videoId: videoId,
	}
}

// Execute executes the request
//  @return []Category
func (a *CategoriesVideosAPIService) GetVideoCategoriesExecute(r ApiGetVideoCategoriesRequest) ([]Category, *http.Response, error) {
	var (
		localVarHTTPMethod   = http.MethodGet
		localVarPostBody     interface{}
		formFiles            []formFile
		localVarReturnValue  []Category
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "CategoriesVideosAPIService.GetVideoCategories")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/videos/{video_id}/categories"
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
	localVarHTTPHeaderAccepts := []string{"application/vnd.vimeo.category+json"}

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

type ApiSuggestVideoCategoryRequest struct {
	ctx context.Context
	ApiService CategoriesVideosAPI
	videoId int32
	suggestVideoCategoryRequest *SuggestVideoCategoryRequest
}

func (r ApiSuggestVideoCategoryRequest) SuggestVideoCategoryRequest(suggestVideoCategoryRequest SuggestVideoCategoryRequest) ApiSuggestVideoCategoryRequest {
	r.suggestVideoCategoryRequest = &suggestVideoCategoryRequest
	return r
}

func (r ApiSuggestVideoCategoryRequest) Execute() (*Category, *http.Response, error) {
	return r.ApiService.SuggestVideoCategoryExecute(r)
}

/*
SuggestVideoCategory Suggest categories for a video

This method sets multiple categories and subcategories for the specified video. Include the categories as a JSON block in the body of the request using the **category** field, like this: `[{ "category": "Tech" }, { "category": "Music" }]`. The authenticated user must have edit access to the video. For more information on batch requests like this one, see [Using Common Formats and Parameters](https://developer.vimeo.com/api/common-formats#working-with-batch-requests).

 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @param videoId The ID of the video.
 @return ApiSuggestVideoCategoryRequest
*/
func (a *CategoriesVideosAPIService) SuggestVideoCategory(ctx context.Context, videoId int32) ApiSuggestVideoCategoryRequest {
	return ApiSuggestVideoCategoryRequest{
		ApiService: a,
		ctx: ctx,
		videoId: videoId,
	}
}

// Execute executes the request
//  @return Category
func (a *CategoriesVideosAPIService) SuggestVideoCategoryExecute(r ApiSuggestVideoCategoryRequest) (*Category, *http.Response, error) {
	var (
		localVarHTTPMethod   = http.MethodPut
		localVarPostBody     interface{}
		formFiles            []formFile
		localVarReturnValue  *Category
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "CategoriesVideosAPIService.SuggestVideoCategory")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/videos/{video_id}/categories"
	localVarPath = strings.Replace(localVarPath, "{"+"video_id"+"}", url.PathEscape(parameterValueToString(r.videoId, "videoId")), -1)

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}
	if r.suggestVideoCategoryRequest == nil {
		return localVarReturnValue, nil, reportError("suggestVideoCategoryRequest is required and must be specified")
	}

	// to determine the Content-Type header
	localVarHTTPContentTypes := []string{"application/vnd.vimeo.category+json"}

	// set Content-Type header
	localVarHTTPContentType := selectHeaderContentType(localVarHTTPContentTypes)
	if localVarHTTPContentType != "" {
		localVarHeaderParams["Content-Type"] = localVarHTTPContentType
	}

	// to determine the Accept header
	localVarHTTPHeaderAccepts := []string{"application/vnd.vimeo.category+json"}

	// set Accept header
	localVarHTTPHeaderAccept := selectHeaderAccept(localVarHTTPHeaderAccepts)
	if localVarHTTPHeaderAccept != "" {
		localVarHeaderParams["Accept"] = localVarHTTPHeaderAccept
	}
	// body params
	localVarPostBody = r.suggestVideoCategoryRequest
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
