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


type ShowcasesCustomShowcaseThumbnailsAPI interface {

	/*
	CreateShowcaseCustomThumb Add a custom thumbnail to a showcase

	This method adds an uploaded image file as a custom thumbnail for the specified showcase. The image doesn't need to be a still from a showcase video, unlike with the [standard thumbnail method](https://developer.vimeo.com/api/reference/albums#set_video_as_album_thumbnail). The authenticated user must be the owner of the showcase.

For information on how to upload the thumbnail, see our [Working with Thumbnail Uploads](https://developer.vimeo.com/api/upload/thumbnails) guide, and follow the same steps.

	@param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
	@param albumId The ID of the showcase.
	@param userId The ID of the user.
	@return ApiCreateShowcaseCustomThumbRequest
	*/
	CreateShowcaseCustomThumb(ctx context.Context, albumId float32, userId float32) ApiCreateShowcaseCustomThumbRequest

	// CreateShowcaseCustomThumbExecute executes the request
	//  @return Picture
	CreateShowcaseCustomThumbExecute(r ApiCreateShowcaseCustomThumbRequest) (*Picture, *http.Response, error)

	/*
	DeleteShowcaseCustomThumbnail Delete a custom showcase thumbnail

	This method deletes the specified custom thumbnail from its showcase. The authenticated user must be the owner of the showcase.

	@param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
	@param albumId The ID of the showcase.
	@param thumbnailId The ID of the custom thumbnail.
	@param userId The ID of the user.
	@return ApiDeleteShowcaseCustomThumbnailRequest
	*/
	DeleteShowcaseCustomThumbnail(ctx context.Context, albumId float32, thumbnailId float32, userId float32) ApiDeleteShowcaseCustomThumbnailRequest

	// DeleteShowcaseCustomThumbnailExecute executes the request
	DeleteShowcaseCustomThumbnailExecute(r ApiDeleteShowcaseCustomThumbnailRequest) (*http.Response, error)

	/*
	GetShowcaseCustomThumbnail Get a specific custom showcase thumbnail

	This method returns a single custom thumbnail of the specified showcase. The authenticated user must be the owner of the showcase.

	@param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
	@param albumId The ID of the showcase.
	@param thumbnailId The ID of the custom thumbnail.
	@param userId The ID of the user.
	@return ApiGetShowcaseCustomThumbnailRequest
	*/
	GetShowcaseCustomThumbnail(ctx context.Context, albumId float32, thumbnailId float32, userId float32) ApiGetShowcaseCustomThumbnailRequest

	// GetShowcaseCustomThumbnailExecute executes the request
	//  @return Picture
	GetShowcaseCustomThumbnailExecute(r ApiGetShowcaseCustomThumbnailRequest) (*Picture, *http.Response, error)

	/*
	GetShowcaseCustomThumbs Get all the custom thumbnails of a showcase

	This method returns every custom thumbnail of the specified showcase. The authenticated user must be the owner of the showcase.

	@param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
	@param albumId The ID of the showcase.
	@param userId The ID of the user.
	@return ApiGetShowcaseCustomThumbsRequest
	*/
	GetShowcaseCustomThumbs(ctx context.Context, albumId float32, userId float32) ApiGetShowcaseCustomThumbsRequest

	// GetShowcaseCustomThumbsExecute executes the request
	//  @return []Picture
	GetShowcaseCustomThumbsExecute(r ApiGetShowcaseCustomThumbsRequest) ([]Picture, *http.Response, error)

	/*
	ReplaceShowcaseCustomThumb Replace a custom showcase thumbnail

	This method replaces the specified custom showcase thumbnail with a new image file. The authenticated user must be the owner of the showcase.

For information on how to upload the thumbnail, see our [Working with Thumbnail Uploads](https://developer.vimeo.com/api/upload/thumbnails) guide.

	@param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
	@param albumId The ID of the showcase.
	@param thumbnailId The ID of the custom thumbnail.
	@param userId The ID of the user.
	@return ApiReplaceShowcaseCustomThumbRequest
	*/
	ReplaceShowcaseCustomThumb(ctx context.Context, albumId float32, thumbnailId float32, userId float32) ApiReplaceShowcaseCustomThumbRequest

	// ReplaceShowcaseCustomThumbExecute executes the request
	//  @return Picture
	ReplaceShowcaseCustomThumbExecute(r ApiReplaceShowcaseCustomThumbRequest) (*Picture, *http.Response, error)
}

// ShowcasesCustomShowcaseThumbnailsAPIService ShowcasesCustomShowcaseThumbnailsAPI service
type ShowcasesCustomShowcaseThumbnailsAPIService service

type ApiCreateShowcaseCustomThumbRequest struct {
	ctx context.Context
	ApiService ShowcasesCustomShowcaseThumbnailsAPI
	albumId float32
	userId float32
}

func (r ApiCreateShowcaseCustomThumbRequest) Execute() (*Picture, *http.Response, error) {
	return r.ApiService.CreateShowcaseCustomThumbExecute(r)
}

/*
CreateShowcaseCustomThumb Add a custom thumbnail to a showcase

This method adds an uploaded image file as a custom thumbnail for the specified showcase. The image doesn't need to be a still from a showcase video, unlike with the [standard thumbnail method](https://developer.vimeo.com/api/reference/albums#set_video_as_album_thumbnail). The authenticated user must be the owner of the showcase.

For information on how to upload the thumbnail, see our [Working with Thumbnail Uploads](https://developer.vimeo.com/api/upload/thumbnails) guide, and follow the same steps.

 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @param albumId The ID of the showcase.
 @param userId The ID of the user.
 @return ApiCreateShowcaseCustomThumbRequest
*/
func (a *ShowcasesCustomShowcaseThumbnailsAPIService) CreateShowcaseCustomThumb(ctx context.Context, albumId float32, userId float32) ApiCreateShowcaseCustomThumbRequest {
	return ApiCreateShowcaseCustomThumbRequest{
		ApiService: a,
		ctx: ctx,
		albumId: albumId,
		userId: userId,
	}
}

// Execute executes the request
//  @return Picture
func (a *ShowcasesCustomShowcaseThumbnailsAPIService) CreateShowcaseCustomThumbExecute(r ApiCreateShowcaseCustomThumbRequest) (*Picture, *http.Response, error) {
	var (
		localVarHTTPMethod   = http.MethodPost
		localVarPostBody     interface{}
		formFiles            []formFile
		localVarReturnValue  *Picture
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "ShowcasesCustomShowcaseThumbnailsAPIService.CreateShowcaseCustomThumb")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/users/{user_id}/albums/{album_id}/custom_thumbnails"
	localVarPath = strings.Replace(localVarPath, "{"+"album_id"+"}", url.PathEscape(parameterValueToString(r.albumId, "albumId")), -1)
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

type ApiDeleteShowcaseCustomThumbnailRequest struct {
	ctx context.Context
	ApiService ShowcasesCustomShowcaseThumbnailsAPI
	albumId float32
	thumbnailId float32
	userId float32
}

func (r ApiDeleteShowcaseCustomThumbnailRequest) Execute() (*http.Response, error) {
	return r.ApiService.DeleteShowcaseCustomThumbnailExecute(r)
}

/*
DeleteShowcaseCustomThumbnail Delete a custom showcase thumbnail

This method deletes the specified custom thumbnail from its showcase. The authenticated user must be the owner of the showcase.

 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @param albumId The ID of the showcase.
 @param thumbnailId The ID of the custom thumbnail.
 @param userId The ID of the user.
 @return ApiDeleteShowcaseCustomThumbnailRequest
*/
func (a *ShowcasesCustomShowcaseThumbnailsAPIService) DeleteShowcaseCustomThumbnail(ctx context.Context, albumId float32, thumbnailId float32, userId float32) ApiDeleteShowcaseCustomThumbnailRequest {
	return ApiDeleteShowcaseCustomThumbnailRequest{
		ApiService: a,
		ctx: ctx,
		albumId: albumId,
		thumbnailId: thumbnailId,
		userId: userId,
	}
}

// Execute executes the request
func (a *ShowcasesCustomShowcaseThumbnailsAPIService) DeleteShowcaseCustomThumbnailExecute(r ApiDeleteShowcaseCustomThumbnailRequest) (*http.Response, error) {
	var (
		localVarHTTPMethod   = http.MethodDelete
		localVarPostBody     interface{}
		formFiles            []formFile
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "ShowcasesCustomShowcaseThumbnailsAPIService.DeleteShowcaseCustomThumbnail")
	if err != nil {
		return nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/users/{user_id}/albums/{album_id}/custom_thumbnails/{thumbnail_id}"
	localVarPath = strings.Replace(localVarPath, "{"+"album_id"+"}", url.PathEscape(parameterValueToString(r.albumId, "albumId")), -1)
	localVarPath = strings.Replace(localVarPath, "{"+"thumbnail_id"+"}", url.PathEscape(parameterValueToString(r.thumbnailId, "thumbnailId")), -1)
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
	localVarHTTPHeaderAccepts := []string{"application/vnd.vimeo.picture+json"}

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

type ApiGetShowcaseCustomThumbnailRequest struct {
	ctx context.Context
	ApiService ShowcasesCustomShowcaseThumbnailsAPI
	albumId float32
	thumbnailId float32
	userId float32
}

func (r ApiGetShowcaseCustomThumbnailRequest) Execute() (*Picture, *http.Response, error) {
	return r.ApiService.GetShowcaseCustomThumbnailExecute(r)
}

/*
GetShowcaseCustomThumbnail Get a specific custom showcase thumbnail

This method returns a single custom thumbnail of the specified showcase. The authenticated user must be the owner of the showcase.

 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @param albumId The ID of the showcase.
 @param thumbnailId The ID of the custom thumbnail.
 @param userId The ID of the user.
 @return ApiGetShowcaseCustomThumbnailRequest
*/
func (a *ShowcasesCustomShowcaseThumbnailsAPIService) GetShowcaseCustomThumbnail(ctx context.Context, albumId float32, thumbnailId float32, userId float32) ApiGetShowcaseCustomThumbnailRequest {
	return ApiGetShowcaseCustomThumbnailRequest{
		ApiService: a,
		ctx: ctx,
		albumId: albumId,
		thumbnailId: thumbnailId,
		userId: userId,
	}
}

// Execute executes the request
//  @return Picture
func (a *ShowcasesCustomShowcaseThumbnailsAPIService) GetShowcaseCustomThumbnailExecute(r ApiGetShowcaseCustomThumbnailRequest) (*Picture, *http.Response, error) {
	var (
		localVarHTTPMethod   = http.MethodGet
		localVarPostBody     interface{}
		formFiles            []formFile
		localVarReturnValue  *Picture
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "ShowcasesCustomShowcaseThumbnailsAPIService.GetShowcaseCustomThumbnail")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/users/{user_id}/albums/{album_id}/custom_thumbnails/{thumbnail_id}"
	localVarPath = strings.Replace(localVarPath, "{"+"album_id"+"}", url.PathEscape(parameterValueToString(r.albumId, "albumId")), -1)
	localVarPath = strings.Replace(localVarPath, "{"+"thumbnail_id"+"}", url.PathEscape(parameterValueToString(r.thumbnailId, "thumbnailId")), -1)
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

type ApiGetShowcaseCustomThumbsRequest struct {
	ctx context.Context
	ApiService ShowcasesCustomShowcaseThumbnailsAPI
	albumId float32
	userId float32
	page *float32
	perPage *float32
}

// The page number of the results to show.
func (r ApiGetShowcaseCustomThumbsRequest) Page(page float32) ApiGetShowcaseCustomThumbsRequest {
	r.page = &page
	return r
}

// The number of items to show on each page of results, up to a maximum of 100.
func (r ApiGetShowcaseCustomThumbsRequest) PerPage(perPage float32) ApiGetShowcaseCustomThumbsRequest {
	r.perPage = &perPage
	return r
}

func (r ApiGetShowcaseCustomThumbsRequest) Execute() ([]Picture, *http.Response, error) {
	return r.ApiService.GetShowcaseCustomThumbsExecute(r)
}

/*
GetShowcaseCustomThumbs Get all the custom thumbnails of a showcase

This method returns every custom thumbnail of the specified showcase. The authenticated user must be the owner of the showcase.

 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @param albumId The ID of the showcase.
 @param userId The ID of the user.
 @return ApiGetShowcaseCustomThumbsRequest
*/
func (a *ShowcasesCustomShowcaseThumbnailsAPIService) GetShowcaseCustomThumbs(ctx context.Context, albumId float32, userId float32) ApiGetShowcaseCustomThumbsRequest {
	return ApiGetShowcaseCustomThumbsRequest{
		ApiService: a,
		ctx: ctx,
		albumId: albumId,
		userId: userId,
	}
}

// Execute executes the request
//  @return []Picture
func (a *ShowcasesCustomShowcaseThumbnailsAPIService) GetShowcaseCustomThumbsExecute(r ApiGetShowcaseCustomThumbsRequest) ([]Picture, *http.Response, error) {
	var (
		localVarHTTPMethod   = http.MethodGet
		localVarPostBody     interface{}
		formFiles            []formFile
		localVarReturnValue  []Picture
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "ShowcasesCustomShowcaseThumbnailsAPIService.GetShowcaseCustomThumbs")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/users/{user_id}/albums/{album_id}/custom_thumbnails"
	localVarPath = strings.Replace(localVarPath, "{"+"album_id"+"}", url.PathEscape(parameterValueToString(r.albumId, "albumId")), -1)
	localVarPath = strings.Replace(localVarPath, "{"+"user_id"+"}", url.PathEscape(parameterValueToString(r.userId, "userId")), -1)

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

type ApiReplaceShowcaseCustomThumbRequest struct {
	ctx context.Context
	ApiService ShowcasesCustomShowcaseThumbnailsAPI
	albumId float32
	thumbnailId float32
	userId float32
	replaceShowcaseCustomThumbRequest *ReplaceShowcaseCustomThumbRequest
}

func (r ApiReplaceShowcaseCustomThumbRequest) ReplaceShowcaseCustomThumbRequest(replaceShowcaseCustomThumbRequest ReplaceShowcaseCustomThumbRequest) ApiReplaceShowcaseCustomThumbRequest {
	r.replaceShowcaseCustomThumbRequest = &replaceShowcaseCustomThumbRequest
	return r
}

func (r ApiReplaceShowcaseCustomThumbRequest) Execute() (*Picture, *http.Response, error) {
	return r.ApiService.ReplaceShowcaseCustomThumbExecute(r)
}

/*
ReplaceShowcaseCustomThumb Replace a custom showcase thumbnail

This method replaces the specified custom showcase thumbnail with a new image file. The authenticated user must be the owner of the showcase.

For information on how to upload the thumbnail, see our [Working with Thumbnail Uploads](https://developer.vimeo.com/api/upload/thumbnails) guide.

 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @param albumId The ID of the showcase.
 @param thumbnailId The ID of the custom thumbnail.
 @param userId The ID of the user.
 @return ApiReplaceShowcaseCustomThumbRequest
*/
func (a *ShowcasesCustomShowcaseThumbnailsAPIService) ReplaceShowcaseCustomThumb(ctx context.Context, albumId float32, thumbnailId float32, userId float32) ApiReplaceShowcaseCustomThumbRequest {
	return ApiReplaceShowcaseCustomThumbRequest{
		ApiService: a,
		ctx: ctx,
		albumId: albumId,
		thumbnailId: thumbnailId,
		userId: userId,
	}
}

// Execute executes the request
//  @return Picture
func (a *ShowcasesCustomShowcaseThumbnailsAPIService) ReplaceShowcaseCustomThumbExecute(r ApiReplaceShowcaseCustomThumbRequest) (*Picture, *http.Response, error) {
	var (
		localVarHTTPMethod   = http.MethodPatch
		localVarPostBody     interface{}
		formFiles            []formFile
		localVarReturnValue  *Picture
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "ShowcasesCustomShowcaseThumbnailsAPIService.ReplaceShowcaseCustomThumb")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/users/{user_id}/albums/{album_id}/custom_thumbnails/{thumbnail_id}"
	localVarPath = strings.Replace(localVarPath, "{"+"album_id"+"}", url.PathEscape(parameterValueToString(r.albumId, "albumId")), -1)
	localVarPath = strings.Replace(localVarPath, "{"+"thumbnail_id"+"}", url.PathEscape(parameterValueToString(r.thumbnailId, "thumbnailId")), -1)
	localVarPath = strings.Replace(localVarPath, "{"+"user_id"+"}", url.PathEscape(parameterValueToString(r.userId, "userId")), -1)

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
	localVarPostBody = r.replaceShowcaseCustomThumbRequest
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