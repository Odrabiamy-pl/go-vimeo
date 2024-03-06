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

type WebinarEmailsAPI interface {

	/*
		GetWebinarEmailSettings Get customization email data for a webinar

		This method returns customized email data for the specified webinar. The authenticated user must have administrative access to the webinar.

		@param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
		@param userId The ID of the user.
		@param webinarId The ID of the webinar.
		@return ApiGetWebinarEmailSettingsRequest
	*/
	GetWebinarEmailSettings(ctx context.Context, userId int32, webinarId string) ApiGetWebinarEmailSettingsRequest

	// GetWebinarEmailSettingsExecute executes the request
	//  @return WebinarEmailSettings
	GetWebinarEmailSettingsExecute(r ApiGetWebinarEmailSettingsRequest) (*WebinarEmailSettings, *http.Response, error)

	/*
		GetWebinarEmailSettingsAlt1 Get customization email data for a webinar

		This method returns customized email data for the specified webinar. The authenticated user must have administrative access to the webinar.

		@param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
		@param webinarId The ID of the webinar.
		@return ApiGetWebinarEmailSettingsAlt1Request
	*/
	GetWebinarEmailSettingsAlt1(ctx context.Context, webinarId string) ApiGetWebinarEmailSettingsAlt1Request

	// GetWebinarEmailSettingsAlt1Execute executes the request
	//  @return WebinarEmailSettings
	GetWebinarEmailSettingsAlt1Execute(r ApiGetWebinarEmailSettingsAlt1Request) (*WebinarEmailSettings, *http.Response, error)

	/*
		UpdateWebinarEmailSettings Customize the email preferences of a webinar

		This method causes the authenticated user to customize the email preferences of the specified webinar. The user must have administrative access to the webinar.

		@param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
		@param userId The ID of the user.
		@param webinarId The ID of the webinar.
		@return ApiUpdateWebinarEmailSettingsRequest
	*/
	UpdateWebinarEmailSettings(ctx context.Context, userId int32, webinarId string) ApiUpdateWebinarEmailSettingsRequest

	// UpdateWebinarEmailSettingsExecute executes the request
	//  @return WebinarEmailSettings
	UpdateWebinarEmailSettingsExecute(r ApiUpdateWebinarEmailSettingsRequest) (*WebinarEmailSettings, *http.Response, error)

	/*
		UpdateWebinarEmailSettingsAlt1 Customize the email preferences of a webinar

		This method causes the authenticated user to customize the email preferences of the specified webinar. The user must have administrative access to the webinar.

		@param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
		@param webinarId The ID of the webinar.
		@return ApiUpdateWebinarEmailSettingsAlt1Request
	*/
	UpdateWebinarEmailSettingsAlt1(ctx context.Context, webinarId string) ApiUpdateWebinarEmailSettingsAlt1Request

	// UpdateWebinarEmailSettingsAlt1Execute executes the request
	//  @return WebinarEmailSettings
	UpdateWebinarEmailSettingsAlt1Execute(r ApiUpdateWebinarEmailSettingsAlt1Request) (*WebinarEmailSettings, *http.Response, error)
}

// WebinarEmailsAPIService WebinarEmailsAPI service
type WebinarEmailsAPIService service

type ApiGetWebinarEmailSettingsRequest struct {
	ctx        context.Context
	ApiService WebinarEmailsAPI
	userId     int32
	webinarId  string
}

func (r ApiGetWebinarEmailSettingsRequest) Execute() (*WebinarEmailSettings, *http.Response, error) {
	return r.ApiService.GetWebinarEmailSettingsExecute(r)
}

/*
GetWebinarEmailSettings Get customization email data for a webinar

This method returns customized email data for the specified webinar. The authenticated user must have administrative access to the webinar.

	@param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
	@param userId The ID of the user.
	@param webinarId The ID of the webinar.
	@return ApiGetWebinarEmailSettingsRequest
*/
func (a *WebinarEmailsAPIService) GetWebinarEmailSettings(ctx context.Context, userId int32, webinarId string) ApiGetWebinarEmailSettingsRequest {
	return ApiGetWebinarEmailSettingsRequest{
		ApiService: a,
		ctx:        ctx,
		userId:     userId,
		webinarId:  webinarId,
	}
}

// Execute executes the request
//
//	@return WebinarEmailSettings
func (a *WebinarEmailsAPIService) GetWebinarEmailSettingsExecute(r ApiGetWebinarEmailSettingsRequest) (*WebinarEmailSettings, *http.Response, error) {
	var (
		localVarHTTPMethod  = http.MethodGet
		localVarPostBody    interface{}
		formFiles           []formFile
		localVarReturnValue *WebinarEmailSettings
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "WebinarEmailsAPIService.GetWebinarEmailSettings")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/users/{user_id}/webinars/{webinar_id}/email_settings"
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
	localVarHTTPHeaderAccepts := []string{"application/vnd.vimeo.webinar.email.settings+json"}

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

type ApiGetWebinarEmailSettingsAlt1Request struct {
	ctx        context.Context
	ApiService WebinarEmailsAPI
	webinarId  string
}

func (r ApiGetWebinarEmailSettingsAlt1Request) Execute() (*WebinarEmailSettings, *http.Response, error) {
	return r.ApiService.GetWebinarEmailSettingsAlt1Execute(r)
}

/*
GetWebinarEmailSettingsAlt1 Get customization email data for a webinar

This method returns customized email data for the specified webinar. The authenticated user must have administrative access to the webinar.

	@param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
	@param webinarId The ID of the webinar.
	@return ApiGetWebinarEmailSettingsAlt1Request
*/
func (a *WebinarEmailsAPIService) GetWebinarEmailSettingsAlt1(ctx context.Context, webinarId string) ApiGetWebinarEmailSettingsAlt1Request {
	return ApiGetWebinarEmailSettingsAlt1Request{
		ApiService: a,
		ctx:        ctx,
		webinarId:  webinarId,
	}
}

// Execute executes the request
//
//	@return WebinarEmailSettings
func (a *WebinarEmailsAPIService) GetWebinarEmailSettingsAlt1Execute(r ApiGetWebinarEmailSettingsAlt1Request) (*WebinarEmailSettings, *http.Response, error) {
	var (
		localVarHTTPMethod  = http.MethodGet
		localVarPostBody    interface{}
		formFiles           []formFile
		localVarReturnValue *WebinarEmailSettings
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "WebinarEmailsAPIService.GetWebinarEmailSettingsAlt1")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/me/webinars/{webinar_id}/email_settings"
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
	localVarHTTPHeaderAccepts := []string{"application/vnd.vimeo.webinar.email.settings+json"}

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

type ApiUpdateWebinarEmailSettingsRequest struct {
	ctx                                   context.Context
	ApiService                            WebinarEmailsAPI
	userId                                int32
	webinarId                             string
	updateWebinarEmailSettingsAlt1Request *UpdateWebinarEmailSettingsAlt1Request
}

func (r ApiUpdateWebinarEmailSettingsRequest) UpdateWebinarEmailSettingsAlt1Request(updateWebinarEmailSettingsAlt1Request UpdateWebinarEmailSettingsAlt1Request) ApiUpdateWebinarEmailSettingsRequest {
	r.updateWebinarEmailSettingsAlt1Request = &updateWebinarEmailSettingsAlt1Request
	return r
}

func (r ApiUpdateWebinarEmailSettingsRequest) Execute() (*WebinarEmailSettings, *http.Response, error) {
	return r.ApiService.UpdateWebinarEmailSettingsExecute(r)
}

/*
UpdateWebinarEmailSettings Customize the email preferences of a webinar

This method causes the authenticated user to customize the email preferences of the specified webinar. The user must have administrative access to the webinar.

	@param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
	@param userId The ID of the user.
	@param webinarId The ID of the webinar.
	@return ApiUpdateWebinarEmailSettingsRequest
*/
func (a *WebinarEmailsAPIService) UpdateWebinarEmailSettings(ctx context.Context, userId int32, webinarId string) ApiUpdateWebinarEmailSettingsRequest {
	return ApiUpdateWebinarEmailSettingsRequest{
		ApiService: a,
		ctx:        ctx,
		userId:     userId,
		webinarId:  webinarId,
	}
}

// Execute executes the request
//
//	@return WebinarEmailSettings
func (a *WebinarEmailsAPIService) UpdateWebinarEmailSettingsExecute(r ApiUpdateWebinarEmailSettingsRequest) (*WebinarEmailSettings, *http.Response, error) {
	var (
		localVarHTTPMethod  = http.MethodPatch
		localVarPostBody    interface{}
		formFiles           []formFile
		localVarReturnValue *WebinarEmailSettings
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "WebinarEmailsAPIService.UpdateWebinarEmailSettings")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/users/{user_id}/webinars/{webinar_id}/email_settings"
	localVarPath = strings.Replace(localVarPath, "{"+"user_id"+"}", url.PathEscape(parameterValueToString(r.userId, "userId")), -1)
	localVarPath = strings.Replace(localVarPath, "{"+"webinar_id"+"}", url.PathEscape(parameterValueToString(r.webinarId, "webinarId")), -1)

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}

	// to determine the Content-Type header
	localVarHTTPContentTypes := []string{"application/vnd.vimeo.webinar.email.settings+json"}

	// set Content-Type header
	localVarHTTPContentType := selectHeaderContentType(localVarHTTPContentTypes)
	if localVarHTTPContentType != "" {
		localVarHeaderParams["Content-Type"] = localVarHTTPContentType
	}

	// to determine the Accept header
	localVarHTTPHeaderAccepts := []string{"application/vnd.vimeo.webinar.email.settings+json"}

	// set Accept header
	localVarHTTPHeaderAccept := selectHeaderAccept(localVarHTTPHeaderAccepts)
	if localVarHTTPHeaderAccept != "" {
		localVarHeaderParams["Accept"] = localVarHTTPHeaderAccept
	}
	// body params
	localVarPostBody = r.updateWebinarEmailSettingsAlt1Request
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

type ApiUpdateWebinarEmailSettingsAlt1Request struct {
	ctx                                   context.Context
	ApiService                            WebinarEmailsAPI
	webinarId                             string
	updateWebinarEmailSettingsAlt1Request *UpdateWebinarEmailSettingsAlt1Request
}

func (r ApiUpdateWebinarEmailSettingsAlt1Request) UpdateWebinarEmailSettingsAlt1Request(updateWebinarEmailSettingsAlt1Request UpdateWebinarEmailSettingsAlt1Request) ApiUpdateWebinarEmailSettingsAlt1Request {
	r.updateWebinarEmailSettingsAlt1Request = &updateWebinarEmailSettingsAlt1Request
	return r
}

func (r ApiUpdateWebinarEmailSettingsAlt1Request) Execute() (*WebinarEmailSettings, *http.Response, error) {
	return r.ApiService.UpdateWebinarEmailSettingsAlt1Execute(r)
}

/*
UpdateWebinarEmailSettingsAlt1 Customize the email preferences of a webinar

This method causes the authenticated user to customize the email preferences of the specified webinar. The user must have administrative access to the webinar.

	@param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
	@param webinarId The ID of the webinar.
	@return ApiUpdateWebinarEmailSettingsAlt1Request
*/
func (a *WebinarEmailsAPIService) UpdateWebinarEmailSettingsAlt1(ctx context.Context, webinarId string) ApiUpdateWebinarEmailSettingsAlt1Request {
	return ApiUpdateWebinarEmailSettingsAlt1Request{
		ApiService: a,
		ctx:        ctx,
		webinarId:  webinarId,
	}
}

// Execute executes the request
//
//	@return WebinarEmailSettings
func (a *WebinarEmailsAPIService) UpdateWebinarEmailSettingsAlt1Execute(r ApiUpdateWebinarEmailSettingsAlt1Request) (*WebinarEmailSettings, *http.Response, error) {
	var (
		localVarHTTPMethod  = http.MethodPatch
		localVarPostBody    interface{}
		formFiles           []formFile
		localVarReturnValue *WebinarEmailSettings
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "WebinarEmailsAPIService.UpdateWebinarEmailSettingsAlt1")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/me/webinars/{webinar_id}/email_settings"
	localVarPath = strings.Replace(localVarPath, "{"+"webinar_id"+"}", url.PathEscape(parameterValueToString(r.webinarId, "webinarId")), -1)

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}

	// to determine the Content-Type header
	localVarHTTPContentTypes := []string{"application/vnd.vimeo.webinar.email.settings+json"}

	// set Content-Type header
	localVarHTTPContentType := selectHeaderContentType(localVarHTTPContentTypes)
	if localVarHTTPContentType != "" {
		localVarHeaderParams["Content-Type"] = localVarHTTPContentType
	}

	// to determine the Accept header
	localVarHTTPHeaderAccepts := []string{"application/vnd.vimeo.webinar.email.settings+json"}

	// set Accept header
	localVarHTTPHeaderAccept := selectHeaderAccept(localVarHTTPHeaderAccepts)
	if localVarHTTPHeaderAccept != "" {
		localVarHeaderParams["Accept"] = localVarHTTPHeaderAccept
	}
	// body params
	localVarPostBody = r.updateWebinarEmailSettingsAlt1Request
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
