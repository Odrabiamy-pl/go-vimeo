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


type WebinarEssentialsAPI interface {

	/*
	CreateWebinar Create a webinar

	This method creates a new webinar for the authenticated user.

	@param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
	@param userId The ID of the user.
	@return ApiCreateWebinarRequest
	*/
	CreateWebinar(ctx context.Context, userId float32) ApiCreateWebinarRequest

	// CreateWebinarExecute executes the request
	//  @return Webinar
	CreateWebinarExecute(r ApiCreateWebinarRequest) (*Webinar, *http.Response, error)

	/*
	CreateWebinarAlt1 Create a webinar

	This method creates a new webinar for the authenticated user.

	@param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
	@return ApiCreateWebinarAlt1Request
	*/
	CreateWebinarAlt1(ctx context.Context) ApiCreateWebinarAlt1Request

	// CreateWebinarAlt1Execute executes the request
	//  @return Webinar
	CreateWebinarAlt1Execute(r ApiCreateWebinarAlt1Request) (*Webinar, *http.Response, error)

	/*
	GetWebinar Get a specific webinar

	This method returns a single webinar belonging to the authenticated user.

	@param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
	@param userId The ID of the user.
	@param webinarId The ID of the webinar.
	@return ApiGetWebinarRequest
	*/
	GetWebinar(ctx context.Context, userId float32, webinarId string) ApiGetWebinarRequest

	// GetWebinarExecute executes the request
	//  @return Webinar
	GetWebinarExecute(r ApiGetWebinarRequest) (*Webinar, *http.Response, error)

	/*
	GetWebinarAlt1 Get a specific webinar

	This method returns a single webinar belonging to the authenticated user.

	@param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
	@param webinarId The ID of the webinar.
	@return ApiGetWebinarAlt1Request
	*/
	GetWebinarAlt1(ctx context.Context, webinarId string) ApiGetWebinarAlt1Request

	// GetWebinarAlt1Execute executes the request
	//  @return Webinar
	GetWebinarAlt1Execute(r ApiGetWebinarAlt1Request) (*Webinar, *http.Response, error)

	/*
	GetWebinarAlt2 Get a specific webinar

	This method returns a single webinar belonging to the authenticated user.

	@param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
	@param webinarId The ID of the webinar.
	@return ApiGetWebinarAlt2Request
	*/
	GetWebinarAlt2(ctx context.Context, webinarId string) ApiGetWebinarAlt2Request

	// GetWebinarAlt2Execute executes the request
	//  @return Webinar
	GetWebinarAlt2Execute(r ApiGetWebinarAlt2Request) (*Webinar, *http.Response, error)

	/*
	UpdateWebinar Update a webinar

	This method updates a webinar belonging to the authenticated user.

	@param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
	@param userId The ID of the user.
	@param webinarId The ID of the webinar.
	@return ApiUpdateWebinarRequest
	*/
	UpdateWebinar(ctx context.Context, userId float32, webinarId string) ApiUpdateWebinarRequest

	// UpdateWebinarExecute executes the request
	//  @return Webinar
	UpdateWebinarExecute(r ApiUpdateWebinarRequest) (*Webinar, *http.Response, error)

	/*
	UpdateWebinarAlt1 Update a webinar

	This method updates a webinar belonging to the authenticated user.

	@param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
	@param webinarId The ID of the webinar.
	@return ApiUpdateWebinarAlt1Request
	*/
	UpdateWebinarAlt1(ctx context.Context, webinarId string) ApiUpdateWebinarAlt1Request

	// UpdateWebinarAlt1Execute executes the request
	//  @return Webinar
	UpdateWebinarAlt1Execute(r ApiUpdateWebinarAlt1Request) (*Webinar, *http.Response, error)
}

// WebinarEssentialsAPIService WebinarEssentialsAPI service
type WebinarEssentialsAPIService service

type ApiCreateWebinarRequest struct {
	ctx context.Context
	ApiService WebinarEssentialsAPI
	userId float32
	createWebinarAlt1Request *CreateWebinarAlt1Request
}

func (r ApiCreateWebinarRequest) CreateWebinarAlt1Request(createWebinarAlt1Request CreateWebinarAlt1Request) ApiCreateWebinarRequest {
	r.createWebinarAlt1Request = &createWebinarAlt1Request
	return r
}

func (r ApiCreateWebinarRequest) Execute() (*Webinar, *http.Response, error) {
	return r.ApiService.CreateWebinarExecute(r)
}

/*
CreateWebinar Create a webinar

This method creates a new webinar for the authenticated user.

 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @param userId The ID of the user.
 @return ApiCreateWebinarRequest
*/
func (a *WebinarEssentialsAPIService) CreateWebinar(ctx context.Context, userId float32) ApiCreateWebinarRequest {
	return ApiCreateWebinarRequest{
		ApiService: a,
		ctx: ctx,
		userId: userId,
	}
}

// Execute executes the request
//  @return Webinar
func (a *WebinarEssentialsAPIService) CreateWebinarExecute(r ApiCreateWebinarRequest) (*Webinar, *http.Response, error) {
	var (
		localVarHTTPMethod   = http.MethodPost
		localVarPostBody     interface{}
		formFiles            []formFile
		localVarReturnValue  *Webinar
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "WebinarEssentialsAPIService.CreateWebinar")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/users/{user_id}/webinars"
	localVarPath = strings.Replace(localVarPath, "{"+"user_id"+"}", url.PathEscape(parameterValueToString(r.userId, "userId")), -1)

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}
	if r.createWebinarAlt1Request == nil {
		return localVarReturnValue, nil, reportError("createWebinarAlt1Request is required and must be specified")
	}

	// to determine the Content-Type header
	localVarHTTPContentTypes := []string{"application/vnd.vimeo.webinars+json"}

	// set Content-Type header
	localVarHTTPContentType := selectHeaderContentType(localVarHTTPContentTypes)
	if localVarHTTPContentType != "" {
		localVarHeaderParams["Content-Type"] = localVarHTTPContentType
	}

	// to determine the Accept header
	localVarHTTPHeaderAccepts := []string{"application/vnd.vimeo.webinars+json"}

	// set Accept header
	localVarHTTPHeaderAccept := selectHeaderAccept(localVarHTTPHeaderAccepts)
	if localVarHTTPHeaderAccept != "" {
		localVarHeaderParams["Accept"] = localVarHTTPHeaderAccept
	}
	// body params
	localVarPostBody = r.createWebinarAlt1Request
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
		if localVarHTTPResponse.StatusCode == 429 {
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

type ApiCreateWebinarAlt1Request struct {
	ctx context.Context
	ApiService WebinarEssentialsAPI
	createWebinarAlt1Request *CreateWebinarAlt1Request
}

func (r ApiCreateWebinarAlt1Request) CreateWebinarAlt1Request(createWebinarAlt1Request CreateWebinarAlt1Request) ApiCreateWebinarAlt1Request {
	r.createWebinarAlt1Request = &createWebinarAlt1Request
	return r
}

func (r ApiCreateWebinarAlt1Request) Execute() (*Webinar, *http.Response, error) {
	return r.ApiService.CreateWebinarAlt1Execute(r)
}

/*
CreateWebinarAlt1 Create a webinar

This method creates a new webinar for the authenticated user.

 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @return ApiCreateWebinarAlt1Request
*/
func (a *WebinarEssentialsAPIService) CreateWebinarAlt1(ctx context.Context) ApiCreateWebinarAlt1Request {
	return ApiCreateWebinarAlt1Request{
		ApiService: a,
		ctx: ctx,
	}
}

// Execute executes the request
//  @return Webinar
func (a *WebinarEssentialsAPIService) CreateWebinarAlt1Execute(r ApiCreateWebinarAlt1Request) (*Webinar, *http.Response, error) {
	var (
		localVarHTTPMethod   = http.MethodPost
		localVarPostBody     interface{}
		formFiles            []formFile
		localVarReturnValue  *Webinar
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "WebinarEssentialsAPIService.CreateWebinarAlt1")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/me/webinars"

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}
	if r.createWebinarAlt1Request == nil {
		return localVarReturnValue, nil, reportError("createWebinarAlt1Request is required and must be specified")
	}

	// to determine the Content-Type header
	localVarHTTPContentTypes := []string{"application/vnd.vimeo.webinars+json"}

	// set Content-Type header
	localVarHTTPContentType := selectHeaderContentType(localVarHTTPContentTypes)
	if localVarHTTPContentType != "" {
		localVarHeaderParams["Content-Type"] = localVarHTTPContentType
	}

	// to determine the Accept header
	localVarHTTPHeaderAccepts := []string{"application/vnd.vimeo.webinars+json"}

	// set Accept header
	localVarHTTPHeaderAccept := selectHeaderAccept(localVarHTTPHeaderAccepts)
	if localVarHTTPHeaderAccept != "" {
		localVarHeaderParams["Accept"] = localVarHTTPHeaderAccept
	}
	// body params
	localVarPostBody = r.createWebinarAlt1Request
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
		if localVarHTTPResponse.StatusCode == 429 {
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

type ApiGetWebinarRequest struct {
	ctx context.Context
	ApiService WebinarEssentialsAPI
	userId float32
	webinarId string
}

func (r ApiGetWebinarRequest) Execute() (*Webinar, *http.Response, error) {
	return r.ApiService.GetWebinarExecute(r)
}

/*
GetWebinar Get a specific webinar

This method returns a single webinar belonging to the authenticated user.

 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @param userId The ID of the user.
 @param webinarId The ID of the webinar.
 @return ApiGetWebinarRequest
*/
func (a *WebinarEssentialsAPIService) GetWebinar(ctx context.Context, userId float32, webinarId string) ApiGetWebinarRequest {
	return ApiGetWebinarRequest{
		ApiService: a,
		ctx: ctx,
		userId: userId,
		webinarId: webinarId,
	}
}

// Execute executes the request
//  @return Webinar
func (a *WebinarEssentialsAPIService) GetWebinarExecute(r ApiGetWebinarRequest) (*Webinar, *http.Response, error) {
	var (
		localVarHTTPMethod   = http.MethodGet
		localVarPostBody     interface{}
		formFiles            []formFile
		localVarReturnValue  *Webinar
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "WebinarEssentialsAPIService.GetWebinar")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/users/{user_id}/webinars/{webinar_id}"
	localVarPath = strings.Replace(localVarPath, "{"+"user_id"+"}", url.PathEscape(parameterValueToString(r.userId, "userId")), -1)
	localVarPath = strings.Replace(localVarPath, "{"+"webinar_id"+"}", url.PathEscape(parameterValueToString(r.webinarId, "webinarId")), -1)

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
	localVarHTTPHeaderAccepts := []string{"application/vnd.vimeo.webinars+json"}

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

type ApiGetWebinarAlt1Request struct {
	ctx context.Context
	ApiService WebinarEssentialsAPI
	webinarId string
}

func (r ApiGetWebinarAlt1Request) Execute() (*Webinar, *http.Response, error) {
	return r.ApiService.GetWebinarAlt1Execute(r)
}

/*
GetWebinarAlt1 Get a specific webinar

This method returns a single webinar belonging to the authenticated user.

 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @param webinarId The ID of the webinar.
 @return ApiGetWebinarAlt1Request
*/
func (a *WebinarEssentialsAPIService) GetWebinarAlt1(ctx context.Context, webinarId string) ApiGetWebinarAlt1Request {
	return ApiGetWebinarAlt1Request{
		ApiService: a,
		ctx: ctx,
		webinarId: webinarId,
	}
}

// Execute executes the request
//  @return Webinar
func (a *WebinarEssentialsAPIService) GetWebinarAlt1Execute(r ApiGetWebinarAlt1Request) (*Webinar, *http.Response, error) {
	var (
		localVarHTTPMethod   = http.MethodGet
		localVarPostBody     interface{}
		formFiles            []formFile
		localVarReturnValue  *Webinar
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "WebinarEssentialsAPIService.GetWebinarAlt1")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/webinars/{webinar_id}"
	localVarPath = strings.Replace(localVarPath, "{"+"webinar_id"+"}", url.PathEscape(parameterValueToString(r.webinarId, "webinarId")), -1)

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
	localVarHTTPHeaderAccepts := []string{"application/vnd.vimeo.webinars+json"}

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

type ApiGetWebinarAlt2Request struct {
	ctx context.Context
	ApiService WebinarEssentialsAPI
	webinarId string
}

func (r ApiGetWebinarAlt2Request) Execute() (*Webinar, *http.Response, error) {
	return r.ApiService.GetWebinarAlt2Execute(r)
}

/*
GetWebinarAlt2 Get a specific webinar

This method returns a single webinar belonging to the authenticated user.

 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @param webinarId The ID of the webinar.
 @return ApiGetWebinarAlt2Request
*/
func (a *WebinarEssentialsAPIService) GetWebinarAlt2(ctx context.Context, webinarId string) ApiGetWebinarAlt2Request {
	return ApiGetWebinarAlt2Request{
		ApiService: a,
		ctx: ctx,
		webinarId: webinarId,
	}
}

// Execute executes the request
//  @return Webinar
func (a *WebinarEssentialsAPIService) GetWebinarAlt2Execute(r ApiGetWebinarAlt2Request) (*Webinar, *http.Response, error) {
	var (
		localVarHTTPMethod   = http.MethodGet
		localVarPostBody     interface{}
		formFiles            []formFile
		localVarReturnValue  *Webinar
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "WebinarEssentialsAPIService.GetWebinarAlt2")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/me/webinars/{webinar_id}"
	localVarPath = strings.Replace(localVarPath, "{"+"webinar_id"+"}", url.PathEscape(parameterValueToString(r.webinarId, "webinarId")), -1)

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
	localVarHTTPHeaderAccepts := []string{"application/vnd.vimeo.webinars+json"}

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

type ApiUpdateWebinarRequest struct {
	ctx context.Context
	ApiService WebinarEssentialsAPI
	userId float32
	webinarId string
	updateWebinarAlt1Request *UpdateWebinarAlt1Request
}

func (r ApiUpdateWebinarRequest) UpdateWebinarAlt1Request(updateWebinarAlt1Request UpdateWebinarAlt1Request) ApiUpdateWebinarRequest {
	r.updateWebinarAlt1Request = &updateWebinarAlt1Request
	return r
}

func (r ApiUpdateWebinarRequest) Execute() (*Webinar, *http.Response, error) {
	return r.ApiService.UpdateWebinarExecute(r)
}

/*
UpdateWebinar Update a webinar

This method updates a webinar belonging to the authenticated user.

 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @param userId The ID of the user.
 @param webinarId The ID of the webinar.
 @return ApiUpdateWebinarRequest
*/
func (a *WebinarEssentialsAPIService) UpdateWebinar(ctx context.Context, userId float32, webinarId string) ApiUpdateWebinarRequest {
	return ApiUpdateWebinarRequest{
		ApiService: a,
		ctx: ctx,
		userId: userId,
		webinarId: webinarId,
	}
}

// Execute executes the request
//  @return Webinar
func (a *WebinarEssentialsAPIService) UpdateWebinarExecute(r ApiUpdateWebinarRequest) (*Webinar, *http.Response, error) {
	var (
		localVarHTTPMethod   = http.MethodPatch
		localVarPostBody     interface{}
		formFiles            []formFile
		localVarReturnValue  *Webinar
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "WebinarEssentialsAPIService.UpdateWebinar")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/users/{user_id}/webinars/{webinar_id}"
	localVarPath = strings.Replace(localVarPath, "{"+"user_id"+"}", url.PathEscape(parameterValueToString(r.userId, "userId")), -1)
	localVarPath = strings.Replace(localVarPath, "{"+"webinar_id"+"}", url.PathEscape(parameterValueToString(r.webinarId, "webinarId")), -1)

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}

	// to determine the Content-Type header
	localVarHTTPContentTypes := []string{"application/vnd.vimeo.webinars+json"}

	// set Content-Type header
	localVarHTTPContentType := selectHeaderContentType(localVarHTTPContentTypes)
	if localVarHTTPContentType != "" {
		localVarHeaderParams["Content-Type"] = localVarHTTPContentType
	}

	// to determine the Accept header
	localVarHTTPHeaderAccepts := []string{"application/vnd.vimeo.webinars+json"}

	// set Accept header
	localVarHTTPHeaderAccept := selectHeaderAccept(localVarHTTPHeaderAccepts)
	if localVarHTTPHeaderAccept != "" {
		localVarHeaderParams["Accept"] = localVarHTTPHeaderAccept
	}
	// body params
	localVarPostBody = r.updateWebinarAlt1Request
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

type ApiUpdateWebinarAlt1Request struct {
	ctx context.Context
	ApiService WebinarEssentialsAPI
	webinarId string
	updateWebinarAlt1Request *UpdateWebinarAlt1Request
}

func (r ApiUpdateWebinarAlt1Request) UpdateWebinarAlt1Request(updateWebinarAlt1Request UpdateWebinarAlt1Request) ApiUpdateWebinarAlt1Request {
	r.updateWebinarAlt1Request = &updateWebinarAlt1Request
	return r
}

func (r ApiUpdateWebinarAlt1Request) Execute() (*Webinar, *http.Response, error) {
	return r.ApiService.UpdateWebinarAlt1Execute(r)
}

/*
UpdateWebinarAlt1 Update a webinar

This method updates a webinar belonging to the authenticated user.

 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @param webinarId The ID of the webinar.
 @return ApiUpdateWebinarAlt1Request
*/
func (a *WebinarEssentialsAPIService) UpdateWebinarAlt1(ctx context.Context, webinarId string) ApiUpdateWebinarAlt1Request {
	return ApiUpdateWebinarAlt1Request{
		ApiService: a,
		ctx: ctx,
		webinarId: webinarId,
	}
}

// Execute executes the request
//  @return Webinar
func (a *WebinarEssentialsAPIService) UpdateWebinarAlt1Execute(r ApiUpdateWebinarAlt1Request) (*Webinar, *http.Response, error) {
	var (
		localVarHTTPMethod   = http.MethodPatch
		localVarPostBody     interface{}
		formFiles            []formFile
		localVarReturnValue  *Webinar
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "WebinarEssentialsAPIService.UpdateWebinarAlt1")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/me/webinars/{webinar_id}"
	localVarPath = strings.Replace(localVarPath, "{"+"webinar_id"+"}", url.PathEscape(parameterValueToString(r.webinarId, "webinarId")), -1)

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}

	// to determine the Content-Type header
	localVarHTTPContentTypes := []string{"application/vnd.vimeo.webinars+json"}

	// set Content-Type header
	localVarHTTPContentType := selectHeaderContentType(localVarHTTPContentTypes)
	if localVarHTTPContentType != "" {
		localVarHeaderParams["Content-Type"] = localVarHTTPContentType
	}

	// to determine the Accept header
	localVarHTTPHeaderAccepts := []string{"application/vnd.vimeo.webinars+json"}

	// set Accept header
	localVarHTTPHeaderAccept := selectHeaderAccept(localVarHTTPHeaderAccepts)
	if localVarHTTPHeaderAccept != "" {
		localVarHeaderParams["Accept"] = localVarHTTPHeaderAccept
	}
	// body params
	localVarPostBody = r.updateWebinarAlt1Request
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