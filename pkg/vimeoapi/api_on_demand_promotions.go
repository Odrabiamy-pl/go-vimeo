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


type OnDemandPromotionsAPI interface {

	/*
	CreateVodPromotion Add a promotion to an On Demand page

	This method adds a promotion to the specified On Demand page. The authenticated user must be the owner of the page.

	@param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
	@param ondemandId The ID of the On Demand page.
	@return ApiCreateVodPromotionRequest
	*/
	CreateVodPromotion(ctx context.Context, ondemandId float32) ApiCreateVodPromotionRequest

	// CreateVodPromotionExecute executes the request
	//  @return OnDemandPromotion
	CreateVodPromotionExecute(r ApiCreateVodPromotionRequest) (*OnDemandPromotion, *http.Response, error)

	/*
	DeleteVodPromotion Delete a promotion on an On Demand page

	This method deletes a promotion on the specified On Demand page. The authenticated user must be the owner of the page.

	@param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
	@param ondemandId The ID of the On Demand page.
	@param promotionId The ID of the promotion.
	@return ApiDeleteVodPromotionRequest
	*/
	DeleteVodPromotion(ctx context.Context, ondemandId float32, promotionId float32) ApiDeleteVodPromotionRequest

	// DeleteVodPromotionExecute executes the request
	DeleteVodPromotionExecute(r ApiDeleteVodPromotionRequest) (*http.Response, error)

	/*
	GetVodPromotion Get a specific promotion on an On Demand page

	This method returns a single promotion on the specified On Demand page. The authenticated user must be the owner of the page.

	@param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
	@param ondemandId The ID of the On Demand page.
	@param promotionId The ID of the promotion.
	@return ApiGetVodPromotionRequest
	*/
	GetVodPromotion(ctx context.Context, ondemandId float32, promotionId float32) ApiGetVodPromotionRequest

	// GetVodPromotionExecute executes the request
	//  @return OnDemandPromotion
	GetVodPromotionExecute(r ApiGetVodPromotionRequest) (*OnDemandPromotion, *http.Response, error)

	/*
	GetVodPromotionCodes Get all the codes of a promotion on an On Demand page

	This method returns every code of the specified promotion on an On Demand page. The authenticated user must be the owner of the page.

	@param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
	@param ondemandId The ID of the On Demand page.
	@param promotionId The ID of the promotion.
	@return ApiGetVodPromotionCodesRequest
	*/
	GetVodPromotionCodes(ctx context.Context, ondemandId float32, promotionId float32) ApiGetVodPromotionCodesRequest

	// GetVodPromotionCodesExecute executes the request
	//  @return OnDemandPromotionCode
	GetVodPromotionCodesExecute(r ApiGetVodPromotionCodesRequest) (*OnDemandPromotionCode, *http.Response, error)

	/*
	GetVodPromotions Get all the promotions on an On Demand page

	This method returns every promotion on the specified On Demand page. The authenticated user must be the owner of the page.

	@param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
	@param ondemandId The ID of the On Demand page.
	@return ApiGetVodPromotionsRequest
	*/
	GetVodPromotions(ctx context.Context, ondemandId float32) ApiGetVodPromotionsRequest

	// GetVodPromotionsExecute executes the request
	//  @return OnDemandPromotion
	GetVodPromotionsExecute(r ApiGetVodPromotionsRequest) (*OnDemandPromotion, *http.Response, error)
}

// OnDemandPromotionsAPIService OnDemandPromotionsAPI service
type OnDemandPromotionsAPIService service

type ApiCreateVodPromotionRequest struct {
	ctx context.Context
	ApiService OnDemandPromotionsAPI
	ondemandId float32
	createVodPromotionRequest *CreateVodPromotionRequest
}

func (r ApiCreateVodPromotionRequest) CreateVodPromotionRequest(createVodPromotionRequest CreateVodPromotionRequest) ApiCreateVodPromotionRequest {
	r.createVodPromotionRequest = &createVodPromotionRequest
	return r
}

func (r ApiCreateVodPromotionRequest) Execute() (*OnDemandPromotion, *http.Response, error) {
	return r.ApiService.CreateVodPromotionExecute(r)
}

/*
CreateVodPromotion Add a promotion to an On Demand page

This method adds a promotion to the specified On Demand page. The authenticated user must be the owner of the page.

 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @param ondemandId The ID of the On Demand page.
 @return ApiCreateVodPromotionRequest
*/
func (a *OnDemandPromotionsAPIService) CreateVodPromotion(ctx context.Context, ondemandId float32) ApiCreateVodPromotionRequest {
	return ApiCreateVodPromotionRequest{
		ApiService: a,
		ctx: ctx,
		ondemandId: ondemandId,
	}
}

// Execute executes the request
//  @return OnDemandPromotion
func (a *OnDemandPromotionsAPIService) CreateVodPromotionExecute(r ApiCreateVodPromotionRequest) (*OnDemandPromotion, *http.Response, error) {
	var (
		localVarHTTPMethod   = http.MethodPost
		localVarPostBody     interface{}
		formFiles            []formFile
		localVarReturnValue  *OnDemandPromotion
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "OnDemandPromotionsAPIService.CreateVodPromotion")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/ondemand/pages/{ondemand_id}/promotions"
	localVarPath = strings.Replace(localVarPath, "{"+"ondemand_id"+"}", url.PathEscape(parameterValueToString(r.ondemandId, "ondemandId")), -1)

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}
	if r.createVodPromotionRequest == nil {
		return localVarReturnValue, nil, reportError("createVodPromotionRequest is required and must be specified")
	}

	// to determine the Content-Type header
	localVarHTTPContentTypes := []string{"application/vnd.vimeo.ondemand.promotion+json"}

	// set Content-Type header
	localVarHTTPContentType := selectHeaderContentType(localVarHTTPContentTypes)
	if localVarHTTPContentType != "" {
		localVarHeaderParams["Content-Type"] = localVarHTTPContentType
	}

	// to determine the Accept header
	localVarHTTPHeaderAccepts := []string{"application/vnd.vimeo.ondemand.promotion+json"}

	// set Accept header
	localVarHTTPHeaderAccept := selectHeaderAccept(localVarHTTPHeaderAccepts)
	if localVarHTTPHeaderAccept != "" {
		localVarHeaderParams["Accept"] = localVarHTTPHeaderAccept
	}
	// body params
	localVarPostBody = r.createVodPromotionRequest
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

type ApiDeleteVodPromotionRequest struct {
	ctx context.Context
	ApiService OnDemandPromotionsAPI
	ondemandId float32
	promotionId float32
}

func (r ApiDeleteVodPromotionRequest) Execute() (*http.Response, error) {
	return r.ApiService.DeleteVodPromotionExecute(r)
}

/*
DeleteVodPromotion Delete a promotion on an On Demand page

This method deletes a promotion on the specified On Demand page. The authenticated user must be the owner of the page.

 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @param ondemandId The ID of the On Demand page.
 @param promotionId The ID of the promotion.
 @return ApiDeleteVodPromotionRequest
*/
func (a *OnDemandPromotionsAPIService) DeleteVodPromotion(ctx context.Context, ondemandId float32, promotionId float32) ApiDeleteVodPromotionRequest {
	return ApiDeleteVodPromotionRequest{
		ApiService: a,
		ctx: ctx,
		ondemandId: ondemandId,
		promotionId: promotionId,
	}
}

// Execute executes the request
func (a *OnDemandPromotionsAPIService) DeleteVodPromotionExecute(r ApiDeleteVodPromotionRequest) (*http.Response, error) {
	var (
		localVarHTTPMethod   = http.MethodDelete
		localVarPostBody     interface{}
		formFiles            []formFile
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "OnDemandPromotionsAPIService.DeleteVodPromotion")
	if err != nil {
		return nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/ondemand/pages/{ondemand_id}/promotions/{promotion_id}"
	localVarPath = strings.Replace(localVarPath, "{"+"ondemand_id"+"}", url.PathEscape(parameterValueToString(r.ondemandId, "ondemandId")), -1)
	localVarPath = strings.Replace(localVarPath, "{"+"promotion_id"+"}", url.PathEscape(parameterValueToString(r.promotionId, "promotionId")), -1)

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

type ApiGetVodPromotionRequest struct {
	ctx context.Context
	ApiService OnDemandPromotionsAPI
	ondemandId float32
	promotionId float32
}

func (r ApiGetVodPromotionRequest) Execute() (*OnDemandPromotion, *http.Response, error) {
	return r.ApiService.GetVodPromotionExecute(r)
}

/*
GetVodPromotion Get a specific promotion on an On Demand page

This method returns a single promotion on the specified On Demand page. The authenticated user must be the owner of the page.

 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @param ondemandId The ID of the On Demand page.
 @param promotionId The ID of the promotion.
 @return ApiGetVodPromotionRequest
*/
func (a *OnDemandPromotionsAPIService) GetVodPromotion(ctx context.Context, ondemandId float32, promotionId float32) ApiGetVodPromotionRequest {
	return ApiGetVodPromotionRequest{
		ApiService: a,
		ctx: ctx,
		ondemandId: ondemandId,
		promotionId: promotionId,
	}
}

// Execute executes the request
//  @return OnDemandPromotion
func (a *OnDemandPromotionsAPIService) GetVodPromotionExecute(r ApiGetVodPromotionRequest) (*OnDemandPromotion, *http.Response, error) {
	var (
		localVarHTTPMethod   = http.MethodGet
		localVarPostBody     interface{}
		formFiles            []formFile
		localVarReturnValue  *OnDemandPromotion
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "OnDemandPromotionsAPIService.GetVodPromotion")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/ondemand/pages/{ondemand_id}/promotions/{promotion_id}"
	localVarPath = strings.Replace(localVarPath, "{"+"ondemand_id"+"}", url.PathEscape(parameterValueToString(r.ondemandId, "ondemandId")), -1)
	localVarPath = strings.Replace(localVarPath, "{"+"promotion_id"+"}", url.PathEscape(parameterValueToString(r.promotionId, "promotionId")), -1)

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
	localVarHTTPHeaderAccepts := []string{"application/vnd.vimeo.ondemand.promotion+json"}

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

type ApiGetVodPromotionCodesRequest struct {
	ctx context.Context
	ApiService OnDemandPromotionsAPI
	ondemandId float32
	promotionId float32
	page *float32
	perPage *float32
}

// The page number of the results to show.
func (r ApiGetVodPromotionCodesRequest) Page(page float32) ApiGetVodPromotionCodesRequest {
	r.page = &page
	return r
}

// The number of items to show on each page of results, up to a maximum of 100.
func (r ApiGetVodPromotionCodesRequest) PerPage(perPage float32) ApiGetVodPromotionCodesRequest {
	r.perPage = &perPage
	return r
}

func (r ApiGetVodPromotionCodesRequest) Execute() (*OnDemandPromotionCode, *http.Response, error) {
	return r.ApiService.GetVodPromotionCodesExecute(r)
}

/*
GetVodPromotionCodes Get all the codes of a promotion on an On Demand page

This method returns every code of the specified promotion on an On Demand page. The authenticated user must be the owner of the page.

 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @param ondemandId The ID of the On Demand page.
 @param promotionId The ID of the promotion.
 @return ApiGetVodPromotionCodesRequest
*/
func (a *OnDemandPromotionsAPIService) GetVodPromotionCodes(ctx context.Context, ondemandId float32, promotionId float32) ApiGetVodPromotionCodesRequest {
	return ApiGetVodPromotionCodesRequest{
		ApiService: a,
		ctx: ctx,
		ondemandId: ondemandId,
		promotionId: promotionId,
	}
}

// Execute executes the request
//  @return OnDemandPromotionCode
func (a *OnDemandPromotionsAPIService) GetVodPromotionCodesExecute(r ApiGetVodPromotionCodesRequest) (*OnDemandPromotionCode, *http.Response, error) {
	var (
		localVarHTTPMethod   = http.MethodGet
		localVarPostBody     interface{}
		formFiles            []formFile
		localVarReturnValue  *OnDemandPromotionCode
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "OnDemandPromotionsAPIService.GetVodPromotionCodes")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/ondemand/pages/{ondemand_id}/promotions/{promotion_id}/codes"
	localVarPath = strings.Replace(localVarPath, "{"+"ondemand_id"+"}", url.PathEscape(parameterValueToString(r.ondemandId, "ondemandId")), -1)
	localVarPath = strings.Replace(localVarPath, "{"+"promotion_id"+"}", url.PathEscape(parameterValueToString(r.promotionId, "promotionId")), -1)

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
	localVarHTTPHeaderAccepts := []string{"application/vnd.vimeo.ondemand.promocode+json"}

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

type ApiGetVodPromotionsRequest struct {
	ctx context.Context
	ApiService OnDemandPromotionsAPI
	ondemandId float32
	filter *string
	page *float32
	perPage *float32
}

// The type of filter to apply to the results.  Option descriptions:  * &#x60;batch&#x60; - Filter the results by the &#x60;batch&#x60; promotion.  * &#x60;default&#x60; - Filter the results by the default promotion.  * &#x60;single&#x60; - Filter the results by the &#x60;single&#x60; promotion.  * &#x60;vip&#x60; - Filter the results by the &#x60;vip&#x60; promotion. 
func (r ApiGetVodPromotionsRequest) Filter(filter string) ApiGetVodPromotionsRequest {
	r.filter = &filter
	return r
}

// The page number of the results to show.
func (r ApiGetVodPromotionsRequest) Page(page float32) ApiGetVodPromotionsRequest {
	r.page = &page
	return r
}

// The number of items to show on each page of results, up to a maximum of 100.
func (r ApiGetVodPromotionsRequest) PerPage(perPage float32) ApiGetVodPromotionsRequest {
	r.perPage = &perPage
	return r
}

func (r ApiGetVodPromotionsRequest) Execute() (*OnDemandPromotion, *http.Response, error) {
	return r.ApiService.GetVodPromotionsExecute(r)
}

/*
GetVodPromotions Get all the promotions on an On Demand page

This method returns every promotion on the specified On Demand page. The authenticated user must be the owner of the page.

 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @param ondemandId The ID of the On Demand page.
 @return ApiGetVodPromotionsRequest
*/
func (a *OnDemandPromotionsAPIService) GetVodPromotions(ctx context.Context, ondemandId float32) ApiGetVodPromotionsRequest {
	return ApiGetVodPromotionsRequest{
		ApiService: a,
		ctx: ctx,
		ondemandId: ondemandId,
	}
}

// Execute executes the request
//  @return OnDemandPromotion
func (a *OnDemandPromotionsAPIService) GetVodPromotionsExecute(r ApiGetVodPromotionsRequest) (*OnDemandPromotion, *http.Response, error) {
	var (
		localVarHTTPMethod   = http.MethodGet
		localVarPostBody     interface{}
		formFiles            []formFile
		localVarReturnValue  *OnDemandPromotion
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "OnDemandPromotionsAPIService.GetVodPromotions")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/ondemand/pages/{ondemand_id}/promotions"
	localVarPath = strings.Replace(localVarPath, "{"+"ondemand_id"+"}", url.PathEscape(parameterValueToString(r.ondemandId, "ondemandId")), -1)

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}
	if r.filter == nil {
		return localVarReturnValue, nil, reportError("filter is required and must be specified")
	}

	parameterAddToHeaderOrQuery(localVarQueryParams, "filter", r.filter, "")
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
	localVarHTTPHeaderAccepts := []string{"application/vnd.vimeo.ondemand.promotion+json"}

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
