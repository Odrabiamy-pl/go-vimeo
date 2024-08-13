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
)


type AuthenticationExtrasConvertAPI interface {

	/*
	ConvertAccessToken Convert an OAuth 1 access token to an OAuth 2 access token

	This method exchanges a legacy Advanced API OAuth 1 token for an API v3 OAuth 2 token.

	@param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
	@return ApiConvertAccessTokenRequest
	*/
	ConvertAccessToken(ctx context.Context) ApiConvertAccessTokenRequest

	// ConvertAccessTokenExecute executes the request
	//  @return Auth
	ConvertAccessTokenExecute(r ApiConvertAccessTokenRequest) (*Auth, *http.Response, error)
}

// AuthenticationExtrasConvertAPIService AuthenticationExtrasConvertAPI service
type AuthenticationExtrasConvertAPIService service

type ApiConvertAccessTokenRequest struct {
	ctx context.Context
	ApiService AuthenticationExtrasConvertAPI
	convertAccessTokenRequest *ConvertAccessTokenRequest
}

func (r ApiConvertAccessTokenRequest) ConvertAccessTokenRequest(convertAccessTokenRequest ConvertAccessTokenRequest) ApiConvertAccessTokenRequest {
	r.convertAccessTokenRequest = &convertAccessTokenRequest
	return r
}

func (r ApiConvertAccessTokenRequest) Execute() (*Auth, *http.Response, error) {
	return r.ApiService.ConvertAccessTokenExecute(r)
}

/*
ConvertAccessToken Convert an OAuth 1 access token to an OAuth 2 access token

This method exchanges a legacy Advanced API OAuth 1 token for an API v3 OAuth 2 token.

 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @return ApiConvertAccessTokenRequest
*/
func (a *AuthenticationExtrasConvertAPIService) ConvertAccessToken(ctx context.Context) ApiConvertAccessTokenRequest {
	return ApiConvertAccessTokenRequest{
		ApiService: a,
		ctx: ctx,
	}
}

// Execute executes the request
//  @return Auth
func (a *AuthenticationExtrasConvertAPIService) ConvertAccessTokenExecute(r ApiConvertAccessTokenRequest) (*Auth, *http.Response, error) {
	var (
		localVarHTTPMethod   = http.MethodPost
		localVarPostBody     interface{}
		formFiles            []formFile
		localVarReturnValue  *Auth
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "AuthenticationExtrasConvertAPIService.ConvertAccessToken")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/oauth/authorize/vimeo_oauth1"

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}
	if r.convertAccessTokenRequest == nil {
		return localVarReturnValue, nil, reportError("convertAccessTokenRequest is required and must be specified")
	}

	// to determine the Content-Type header
	localVarHTTPContentTypes := []string{"application/vnd.vimeo.auth+json"}

	// set Content-Type header
	localVarHTTPContentType := selectHeaderContentType(localVarHTTPContentTypes)
	if localVarHTTPContentType != "" {
		localVarHeaderParams["Content-Type"] = localVarHTTPContentType
	}

	// to determine the Accept header
	localVarHTTPHeaderAccepts := []string{"application/vnd.vimeo.auth+json"}

	// set Accept header
	localVarHTTPHeaderAccept := selectHeaderAccept(localVarHTTPHeaderAccepts)
	if localVarHTTPHeaderAccept != "" {
		localVarHeaderParams["Accept"] = localVarHTTPHeaderAccept
	}
	// body params
	localVarPostBody = r.convertAccessTokenRequest
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
			var v AuthError
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
