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


type AuthenticationExtrasExchangeAPI interface {

	/*
	ExchangeAuthCode Exchange an authorization code for an access token

	This method exchanges an OAuth authorization code for an OAuth access token.

	@param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
	@return ApiExchangeAuthCodeRequest
	*/
	ExchangeAuthCode(ctx context.Context) ApiExchangeAuthCodeRequest

	// ExchangeAuthCodeExecute executes the request
	//  @return Auth
	ExchangeAuthCodeExecute(r ApiExchangeAuthCodeRequest) (*Auth, *http.Response, error)
}

// AuthenticationExtrasExchangeAPIService AuthenticationExtrasExchangeAPI service
type AuthenticationExtrasExchangeAPIService service

type ApiExchangeAuthCodeRequest struct {
	ctx context.Context
	ApiService AuthenticationExtrasExchangeAPI
	exchangeAuthCodeRequest *ExchangeAuthCodeRequest
}

func (r ApiExchangeAuthCodeRequest) ExchangeAuthCodeRequest(exchangeAuthCodeRequest ExchangeAuthCodeRequest) ApiExchangeAuthCodeRequest {
	r.exchangeAuthCodeRequest = &exchangeAuthCodeRequest
	return r
}

func (r ApiExchangeAuthCodeRequest) Execute() (*Auth, *http.Response, error) {
	return r.ApiService.ExchangeAuthCodeExecute(r)
}

/*
ExchangeAuthCode Exchange an authorization code for an access token

This method exchanges an OAuth authorization code for an OAuth access token.

 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @return ApiExchangeAuthCodeRequest
*/
func (a *AuthenticationExtrasExchangeAPIService) ExchangeAuthCode(ctx context.Context) ApiExchangeAuthCodeRequest {
	return ApiExchangeAuthCodeRequest{
		ApiService: a,
		ctx: ctx,
	}
}

// Execute executes the request
//  @return Auth
func (a *AuthenticationExtrasExchangeAPIService) ExchangeAuthCodeExecute(r ApiExchangeAuthCodeRequest) (*Auth, *http.Response, error) {
	var (
		localVarHTTPMethod   = http.MethodPost
		localVarPostBody     interface{}
		formFiles            []formFile
		localVarReturnValue  *Auth
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "AuthenticationExtrasExchangeAPIService.ExchangeAuthCode")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/oauth/access_token"

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}
	if r.exchangeAuthCodeRequest == nil {
		return localVarReturnValue, nil, reportError("exchangeAuthCodeRequest is required and must be specified")
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
	localVarPostBody = r.exchangeAuthCodeRequest
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
