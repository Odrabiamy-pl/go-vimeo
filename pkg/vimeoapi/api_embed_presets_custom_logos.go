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


type EmbedPresetsCustomLogosAPI interface {

	/*
	CreateCustomLogo Add a custom user logo

	This method adds a custom logo representing the authenticated user for display in the embedded player. Be sure to use this method in the context of the multi-step upload procedure described in our [Working with Thumbnail Uploads](https://developer.vimeo.com/api/upload/thumbnails#uploading-a-thumbnail) guide. This method represents Step 2 of the procedure.

	@param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
	@param userId The ID of the user.
	@return ApiCreateCustomLogoRequest
	*/
	CreateCustomLogo(ctx context.Context, userId int32) ApiCreateCustomLogoRequest

	// CreateCustomLogoExecute executes the request
	//  @return Picture
	CreateCustomLogoExecute(r ApiCreateCustomLogoRequest) (*Picture, *http.Response, error)

	/*
	CreateCustomLogoAlt1 Add a custom user logo

	This method adds a custom logo representing the authenticated user for display in the embedded player. Be sure to use this method in the context of the multi-step upload procedure described in our [Working with Thumbnail Uploads](https://developer.vimeo.com/api/upload/thumbnails#uploading-a-thumbnail) guide. This method represents Step 2 of the procedure.

	@param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
	@return ApiCreateCustomLogoAlt1Request
	*/
	CreateCustomLogoAlt1(ctx context.Context) ApiCreateCustomLogoAlt1Request

	// CreateCustomLogoAlt1Execute executes the request
	//  @return Picture
	CreateCustomLogoAlt1Execute(r ApiCreateCustomLogoAlt1Request) (*Picture, *http.Response, error)

	/*
	DeleteCustomLogo Delete a custom user logo

	This method deletes the specified custom logo belonging to the authenticated user.

	@param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
	@param logoId The ID of the custom logo.
	@param userId The ID of the user.
	@return ApiDeleteCustomLogoRequest
	*/
	DeleteCustomLogo(ctx context.Context, logoId float32, userId int32) ApiDeleteCustomLogoRequest

	// DeleteCustomLogoExecute executes the request
	DeleteCustomLogoExecute(r ApiDeleteCustomLogoRequest) (*http.Response, error)

	/*
	DeleteCustomLogoAlt1 Delete a custom user logo

	This method deletes the specified custom logo belonging to the authenticated user.

	@param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
	@param logoId The ID of the custom logo.
	@return ApiDeleteCustomLogoAlt1Request
	*/
	DeleteCustomLogoAlt1(ctx context.Context, logoId float32) ApiDeleteCustomLogoAlt1Request

	// DeleteCustomLogoAlt1Execute executes the request
	DeleteCustomLogoAlt1Execute(r ApiDeleteCustomLogoAlt1Request) (*http.Response, error)

	/*
	GetCustomLogo Get a specific custom user logo

	This method returns a single custom logo belonging to the authenticated user.

	@param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
	@param logoId The ID of the custom logo.
	@param userId The ID of the user.
	@return ApiGetCustomLogoRequest
	*/
	GetCustomLogo(ctx context.Context, logoId float32, userId int32) ApiGetCustomLogoRequest

	// GetCustomLogoExecute executes the request
	//  @return Picture
	GetCustomLogoExecute(r ApiGetCustomLogoRequest) (*Picture, *http.Response, error)

	/*
	GetCustomLogoAlt1 Get a specific custom user logo

	This method returns a single custom logo belonging to the authenticated user.

	@param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
	@param logoId The ID of the custom logo.
	@return ApiGetCustomLogoAlt1Request
	*/
	GetCustomLogoAlt1(ctx context.Context, logoId float32) ApiGetCustomLogoAlt1Request

	// GetCustomLogoAlt1Execute executes the request
	//  @return Picture
	GetCustomLogoAlt1Execute(r ApiGetCustomLogoAlt1Request) (*Picture, *http.Response, error)

	/*
	GetCustomLogos Get all the custom logos that belong to the user

	This method returns every custom logo that belongs to the authenticated user or team owner.

	@param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
	@param userId The ID of the user.
	@return ApiGetCustomLogosRequest
	*/
	GetCustomLogos(ctx context.Context, userId int32) ApiGetCustomLogosRequest

	// GetCustomLogosExecute executes the request
	//  @return []Picture
	GetCustomLogosExecute(r ApiGetCustomLogosRequest) ([]Picture, *http.Response, error)

	/*
	GetCustomLogosAlt1 Get all the custom logos that belong to the user

	This method returns every custom logo that belongs to the authenticated user or team owner.

	@param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
	@return ApiGetCustomLogosAlt1Request
	*/
	GetCustomLogosAlt1(ctx context.Context) ApiGetCustomLogosAlt1Request

	// GetCustomLogosAlt1Execute executes the request
	//  @return []Picture
	GetCustomLogosAlt1Execute(r ApiGetCustomLogosAlt1Request) ([]Picture, *http.Response, error)
}

// EmbedPresetsCustomLogosAPIService EmbedPresetsCustomLogosAPI service
type EmbedPresetsCustomLogosAPIService service

type ApiCreateCustomLogoRequest struct {
	ctx context.Context
	ApiService EmbedPresetsCustomLogosAPI
	userId int32
}

func (r ApiCreateCustomLogoRequest) Execute() (*Picture, *http.Response, error) {
	return r.ApiService.CreateCustomLogoExecute(r)
}

/*
CreateCustomLogo Add a custom user logo

This method adds a custom logo representing the authenticated user for display in the embedded player. Be sure to use this method in the context of the multi-step upload procedure described in our [Working with Thumbnail Uploads](https://developer.vimeo.com/api/upload/thumbnails#uploading-a-thumbnail) guide. This method represents Step 2 of the procedure.

 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @param userId The ID of the user.
 @return ApiCreateCustomLogoRequest
*/
func (a *EmbedPresetsCustomLogosAPIService) CreateCustomLogo(ctx context.Context, userId int32) ApiCreateCustomLogoRequest {
	return ApiCreateCustomLogoRequest{
		ApiService: a,
		ctx: ctx,
		userId: userId,
	}
}

// Execute executes the request
//  @return Picture
func (a *EmbedPresetsCustomLogosAPIService) CreateCustomLogoExecute(r ApiCreateCustomLogoRequest) (*Picture, *http.Response, error) {
	var (
		localVarHTTPMethod   = http.MethodPost
		localVarPostBody     interface{}
		formFiles            []formFile
		localVarReturnValue  *Picture
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "EmbedPresetsCustomLogosAPIService.CreateCustomLogo")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/users/{user_id}/customlogos"
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

type ApiCreateCustomLogoAlt1Request struct {
	ctx context.Context
	ApiService EmbedPresetsCustomLogosAPI
}

func (r ApiCreateCustomLogoAlt1Request) Execute() (*Picture, *http.Response, error) {
	return r.ApiService.CreateCustomLogoAlt1Execute(r)
}

/*
CreateCustomLogoAlt1 Add a custom user logo

This method adds a custom logo representing the authenticated user for display in the embedded player. Be sure to use this method in the context of the multi-step upload procedure described in our [Working with Thumbnail Uploads](https://developer.vimeo.com/api/upload/thumbnails#uploading-a-thumbnail) guide. This method represents Step 2 of the procedure.

 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @return ApiCreateCustomLogoAlt1Request
*/
func (a *EmbedPresetsCustomLogosAPIService) CreateCustomLogoAlt1(ctx context.Context) ApiCreateCustomLogoAlt1Request {
	return ApiCreateCustomLogoAlt1Request{
		ApiService: a,
		ctx: ctx,
	}
}

// Execute executes the request
//  @return Picture
func (a *EmbedPresetsCustomLogosAPIService) CreateCustomLogoAlt1Execute(r ApiCreateCustomLogoAlt1Request) (*Picture, *http.Response, error) {
	var (
		localVarHTTPMethod   = http.MethodPost
		localVarPostBody     interface{}
		formFiles            []formFile
		localVarReturnValue  *Picture
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "EmbedPresetsCustomLogosAPIService.CreateCustomLogoAlt1")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/me/customlogos"

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

type ApiDeleteCustomLogoRequest struct {
	ctx context.Context
	ApiService EmbedPresetsCustomLogosAPI
	logoId float32
	userId int32
}

func (r ApiDeleteCustomLogoRequest) Execute() (*http.Response, error) {
	return r.ApiService.DeleteCustomLogoExecute(r)
}

/*
DeleteCustomLogo Delete a custom user logo

This method deletes the specified custom logo belonging to the authenticated user.

 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @param logoId The ID of the custom logo.
 @param userId The ID of the user.
 @return ApiDeleteCustomLogoRequest
*/
func (a *EmbedPresetsCustomLogosAPIService) DeleteCustomLogo(ctx context.Context, logoId float32, userId int32) ApiDeleteCustomLogoRequest {
	return ApiDeleteCustomLogoRequest{
		ApiService: a,
		ctx: ctx,
		logoId: logoId,
		userId: userId,
	}
}

// Execute executes the request
func (a *EmbedPresetsCustomLogosAPIService) DeleteCustomLogoExecute(r ApiDeleteCustomLogoRequest) (*http.Response, error) {
	var (
		localVarHTTPMethod   = http.MethodDelete
		localVarPostBody     interface{}
		formFiles            []formFile
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "EmbedPresetsCustomLogosAPIService.DeleteCustomLogo")
	if err != nil {
		return nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/users/{user_id}/customlogos/{logo_id}"
	localVarPath = strings.Replace(localVarPath, "{"+"logo_id"+"}", url.PathEscape(parameterValueToString(r.logoId, "logoId")), -1)
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
		}
		return localVarHTTPResponse, newErr
	}

	return localVarHTTPResponse, nil
}

type ApiDeleteCustomLogoAlt1Request struct {
	ctx context.Context
	ApiService EmbedPresetsCustomLogosAPI
	logoId float32
}

func (r ApiDeleteCustomLogoAlt1Request) Execute() (*http.Response, error) {
	return r.ApiService.DeleteCustomLogoAlt1Execute(r)
}

/*
DeleteCustomLogoAlt1 Delete a custom user logo

This method deletes the specified custom logo belonging to the authenticated user.

 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @param logoId The ID of the custom logo.
 @return ApiDeleteCustomLogoAlt1Request
*/
func (a *EmbedPresetsCustomLogosAPIService) DeleteCustomLogoAlt1(ctx context.Context, logoId float32) ApiDeleteCustomLogoAlt1Request {
	return ApiDeleteCustomLogoAlt1Request{
		ApiService: a,
		ctx: ctx,
		logoId: logoId,
	}
}

// Execute executes the request
func (a *EmbedPresetsCustomLogosAPIService) DeleteCustomLogoAlt1Execute(r ApiDeleteCustomLogoAlt1Request) (*http.Response, error) {
	var (
		localVarHTTPMethod   = http.MethodDelete
		localVarPostBody     interface{}
		formFiles            []formFile
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "EmbedPresetsCustomLogosAPIService.DeleteCustomLogoAlt1")
	if err != nil {
		return nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/me/customlogos/{logo_id}"
	localVarPath = strings.Replace(localVarPath, "{"+"logo_id"+"}", url.PathEscape(parameterValueToString(r.logoId, "logoId")), -1)

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
		}
		return localVarHTTPResponse, newErr
	}

	return localVarHTTPResponse, nil
}

type ApiGetCustomLogoRequest struct {
	ctx context.Context
	ApiService EmbedPresetsCustomLogosAPI
	logoId float32
	userId int32
}

func (r ApiGetCustomLogoRequest) Execute() (*Picture, *http.Response, error) {
	return r.ApiService.GetCustomLogoExecute(r)
}

/*
GetCustomLogo Get a specific custom user logo

This method returns a single custom logo belonging to the authenticated user.

 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @param logoId The ID of the custom logo.
 @param userId The ID of the user.
 @return ApiGetCustomLogoRequest
*/
func (a *EmbedPresetsCustomLogosAPIService) GetCustomLogo(ctx context.Context, logoId float32, userId int32) ApiGetCustomLogoRequest {
	return ApiGetCustomLogoRequest{
		ApiService: a,
		ctx: ctx,
		logoId: logoId,
		userId: userId,
	}
}

// Execute executes the request
//  @return Picture
func (a *EmbedPresetsCustomLogosAPIService) GetCustomLogoExecute(r ApiGetCustomLogoRequest) (*Picture, *http.Response, error) {
	var (
		localVarHTTPMethod   = http.MethodGet
		localVarPostBody     interface{}
		formFiles            []formFile
		localVarReturnValue  *Picture
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "EmbedPresetsCustomLogosAPIService.GetCustomLogo")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/users/{user_id}/customlogos/{logo_id}"
	localVarPath = strings.Replace(localVarPath, "{"+"logo_id"+"}", url.PathEscape(parameterValueToString(r.logoId, "logoId")), -1)
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

type ApiGetCustomLogoAlt1Request struct {
	ctx context.Context
	ApiService EmbedPresetsCustomLogosAPI
	logoId float32
}

func (r ApiGetCustomLogoAlt1Request) Execute() (*Picture, *http.Response, error) {
	return r.ApiService.GetCustomLogoAlt1Execute(r)
}

/*
GetCustomLogoAlt1 Get a specific custom user logo

This method returns a single custom logo belonging to the authenticated user.

 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @param logoId The ID of the custom logo.
 @return ApiGetCustomLogoAlt1Request
*/
func (a *EmbedPresetsCustomLogosAPIService) GetCustomLogoAlt1(ctx context.Context, logoId float32) ApiGetCustomLogoAlt1Request {
	return ApiGetCustomLogoAlt1Request{
		ApiService: a,
		ctx: ctx,
		logoId: logoId,
	}
}

// Execute executes the request
//  @return Picture
func (a *EmbedPresetsCustomLogosAPIService) GetCustomLogoAlt1Execute(r ApiGetCustomLogoAlt1Request) (*Picture, *http.Response, error) {
	var (
		localVarHTTPMethod   = http.MethodGet
		localVarPostBody     interface{}
		formFiles            []formFile
		localVarReturnValue  *Picture
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "EmbedPresetsCustomLogosAPIService.GetCustomLogoAlt1")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/me/customlogos/{logo_id}"
	localVarPath = strings.Replace(localVarPath, "{"+"logo_id"+"}", url.PathEscape(parameterValueToString(r.logoId, "logoId")), -1)

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

type ApiGetCustomLogosRequest struct {
	ctx context.Context
	ApiService EmbedPresetsCustomLogosAPI
	userId int32
}

func (r ApiGetCustomLogosRequest) Execute() ([]Picture, *http.Response, error) {
	return r.ApiService.GetCustomLogosExecute(r)
}

/*
GetCustomLogos Get all the custom logos that belong to the user

This method returns every custom logo that belongs to the authenticated user or team owner.

 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @param userId The ID of the user.
 @return ApiGetCustomLogosRequest
*/
func (a *EmbedPresetsCustomLogosAPIService) GetCustomLogos(ctx context.Context, userId int32) ApiGetCustomLogosRequest {
	return ApiGetCustomLogosRequest{
		ApiService: a,
		ctx: ctx,
		userId: userId,
	}
}

// Execute executes the request
//  @return []Picture
func (a *EmbedPresetsCustomLogosAPIService) GetCustomLogosExecute(r ApiGetCustomLogosRequest) ([]Picture, *http.Response, error) {
	var (
		localVarHTTPMethod   = http.MethodGet
		localVarPostBody     interface{}
		formFiles            []formFile
		localVarReturnValue  []Picture
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "EmbedPresetsCustomLogosAPIService.GetCustomLogos")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/users/{user_id}/customlogos"
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

type ApiGetCustomLogosAlt1Request struct {
	ctx context.Context
	ApiService EmbedPresetsCustomLogosAPI
}

func (r ApiGetCustomLogosAlt1Request) Execute() ([]Picture, *http.Response, error) {
	return r.ApiService.GetCustomLogosAlt1Execute(r)
}

/*
GetCustomLogosAlt1 Get all the custom logos that belong to the user

This method returns every custom logo that belongs to the authenticated user or team owner.

 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @return ApiGetCustomLogosAlt1Request
*/
func (a *EmbedPresetsCustomLogosAPIService) GetCustomLogosAlt1(ctx context.Context) ApiGetCustomLogosAlt1Request {
	return ApiGetCustomLogosAlt1Request{
		ApiService: a,
		ctx: ctx,
	}
}

// Execute executes the request
//  @return []Picture
func (a *EmbedPresetsCustomLogosAPIService) GetCustomLogosAlt1Execute(r ApiGetCustomLogosAlt1Request) ([]Picture, *http.Response, error) {
	var (
		localVarHTTPMethod   = http.MethodGet
		localVarPostBody     interface{}
		formFiles            []formFile
		localVarReturnValue  []Picture
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "EmbedPresetsCustomLogosAPIService.GetCustomLogosAlt1")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/me/customlogos"

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
