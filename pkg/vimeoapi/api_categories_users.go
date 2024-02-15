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


type CategoriesUsersAPI interface {

	/*
	CheckIfUserSubscribedToCategory Check if the user follows a category

	This method determines whether the authenticated user follows the specified category.

	@param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
	@param category The name of the category.
	@param userId The ID of the user.
	@return ApiCheckIfUserSubscribedToCategoryRequest
	*/
	CheckIfUserSubscribedToCategory(ctx context.Context, category string, userId float32) ApiCheckIfUserSubscribedToCategoryRequest

	// CheckIfUserSubscribedToCategoryExecute executes the request
	CheckIfUserSubscribedToCategoryExecute(r ApiCheckIfUserSubscribedToCategoryRequest) (*http.Response, error)

	/*
	CheckIfUserSubscribedToCategoryAlt1 Check if the user follows a category

	This method determines whether the authenticated user follows the specified category.

	@param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
	@param category The name of the category.
	@return ApiCheckIfUserSubscribedToCategoryAlt1Request
	*/
	CheckIfUserSubscribedToCategoryAlt1(ctx context.Context, category string) ApiCheckIfUserSubscribedToCategoryAlt1Request

	// CheckIfUserSubscribedToCategoryAlt1Execute executes the request
	CheckIfUserSubscribedToCategoryAlt1Execute(r ApiCheckIfUserSubscribedToCategoryAlt1Request) (*http.Response, error)

	/*
	GetCategorySubscriptions Get all the categories that the user follows

	This method returns every category that the authenticated user follows.

	@param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
	@param userId The ID of the user.
	@return ApiGetCategorySubscriptionsRequest
	*/
	GetCategorySubscriptions(ctx context.Context, userId float32) ApiGetCategorySubscriptionsRequest

	// GetCategorySubscriptionsExecute executes the request
	//  @return []Category
	GetCategorySubscriptionsExecute(r ApiGetCategorySubscriptionsRequest) ([]Category, *http.Response, error)

	/*
	GetCategorySubscriptionsAlt1 Get all the categories that the user follows

	This method returns every category that the authenticated user follows.

	@param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
	@return ApiGetCategorySubscriptionsAlt1Request
	*/
	GetCategorySubscriptionsAlt1(ctx context.Context) ApiGetCategorySubscriptionsAlt1Request

	// GetCategorySubscriptionsAlt1Execute executes the request
	//  @return []Category
	GetCategorySubscriptionsAlt1Execute(r ApiGetCategorySubscriptionsAlt1Request) ([]Category, *http.Response, error)

	/*
	SubscribeToCategory Cause the user to follow a specific category

	This method causes the authenticated user to follow the specified category.

	@param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
	@param category The name of the category.
	@param userId The ID of the user.
	@return ApiSubscribeToCategoryRequest
	*/
	SubscribeToCategory(ctx context.Context, category float32, userId float32) ApiSubscribeToCategoryRequest

	// SubscribeToCategoryExecute executes the request
	SubscribeToCategoryExecute(r ApiSubscribeToCategoryRequest) (*http.Response, error)

	/*
	SubscribeToCategoryAlt1 Cause the user to follow a specific category

	This method causes the authenticated user to follow the specified category.

	@param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
	@param category The name of the category.
	@return ApiSubscribeToCategoryAlt1Request
	*/
	SubscribeToCategoryAlt1(ctx context.Context, category float32) ApiSubscribeToCategoryAlt1Request

	// SubscribeToCategoryAlt1Execute executes the request
	SubscribeToCategoryAlt1Execute(r ApiSubscribeToCategoryAlt1Request) (*http.Response, error)

	/*
	UnsubscribeFromCategory Cause the user to stop following a category

	This method causes the authenticated user to stop following the specified category.

	@param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
	@param category The name of the category.
	@param userId The ID of the user.
	@return ApiUnsubscribeFromCategoryRequest
	*/
	UnsubscribeFromCategory(ctx context.Context, category string, userId float32) ApiUnsubscribeFromCategoryRequest

	// UnsubscribeFromCategoryExecute executes the request
	UnsubscribeFromCategoryExecute(r ApiUnsubscribeFromCategoryRequest) (*http.Response, error)

	/*
	UnsubscribeFromCategoryAlt1 Cause the user to stop following a category

	This method causes the authenticated user to stop following the specified category.

	@param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
	@param category The name of the category.
	@return ApiUnsubscribeFromCategoryAlt1Request
	*/
	UnsubscribeFromCategoryAlt1(ctx context.Context, category string) ApiUnsubscribeFromCategoryAlt1Request

	// UnsubscribeFromCategoryAlt1Execute executes the request
	UnsubscribeFromCategoryAlt1Execute(r ApiUnsubscribeFromCategoryAlt1Request) (*http.Response, error)
}

// CategoriesUsersAPIService CategoriesUsersAPI service
type CategoriesUsersAPIService service

type ApiCheckIfUserSubscribedToCategoryRequest struct {
	ctx context.Context
	ApiService CategoriesUsersAPI
	category string
	userId float32
}

func (r ApiCheckIfUserSubscribedToCategoryRequest) Execute() (*http.Response, error) {
	return r.ApiService.CheckIfUserSubscribedToCategoryExecute(r)
}

/*
CheckIfUserSubscribedToCategory Check if the user follows a category

This method determines whether the authenticated user follows the specified category.

 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @param category The name of the category.
 @param userId The ID of the user.
 @return ApiCheckIfUserSubscribedToCategoryRequest
*/
func (a *CategoriesUsersAPIService) CheckIfUserSubscribedToCategory(ctx context.Context, category string, userId float32) ApiCheckIfUserSubscribedToCategoryRequest {
	return ApiCheckIfUserSubscribedToCategoryRequest{
		ApiService: a,
		ctx: ctx,
		category: category,
		userId: userId,
	}
}

// Execute executes the request
func (a *CategoriesUsersAPIService) CheckIfUserSubscribedToCategoryExecute(r ApiCheckIfUserSubscribedToCategoryRequest) (*http.Response, error) {
	var (
		localVarHTTPMethod   = http.MethodGet
		localVarPostBody     interface{}
		formFiles            []formFile
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "CategoriesUsersAPIService.CheckIfUserSubscribedToCategory")
	if err != nil {
		return nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/users/{user_id}/categories/{category}"
	localVarPath = strings.Replace(localVarPath, "{"+"category"+"}", url.PathEscape(parameterValueToString(r.category, "category")), -1)
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
	localVarHTTPHeaderAccepts := []string{}

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
		return localVarHTTPResponse, newErr
	}

	return localVarHTTPResponse, nil
}

type ApiCheckIfUserSubscribedToCategoryAlt1Request struct {
	ctx context.Context
	ApiService CategoriesUsersAPI
	category string
}

func (r ApiCheckIfUserSubscribedToCategoryAlt1Request) Execute() (*http.Response, error) {
	return r.ApiService.CheckIfUserSubscribedToCategoryAlt1Execute(r)
}

/*
CheckIfUserSubscribedToCategoryAlt1 Check if the user follows a category

This method determines whether the authenticated user follows the specified category.

 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @param category The name of the category.
 @return ApiCheckIfUserSubscribedToCategoryAlt1Request
*/
func (a *CategoriesUsersAPIService) CheckIfUserSubscribedToCategoryAlt1(ctx context.Context, category string) ApiCheckIfUserSubscribedToCategoryAlt1Request {
	return ApiCheckIfUserSubscribedToCategoryAlt1Request{
		ApiService: a,
		ctx: ctx,
		category: category,
	}
}

// Execute executes the request
func (a *CategoriesUsersAPIService) CheckIfUserSubscribedToCategoryAlt1Execute(r ApiCheckIfUserSubscribedToCategoryAlt1Request) (*http.Response, error) {
	var (
		localVarHTTPMethod   = http.MethodGet
		localVarPostBody     interface{}
		formFiles            []formFile
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "CategoriesUsersAPIService.CheckIfUserSubscribedToCategoryAlt1")
	if err != nil {
		return nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/me/categories/{category}"
	localVarPath = strings.Replace(localVarPath, "{"+"category"+"}", url.PathEscape(parameterValueToString(r.category, "category")), -1)

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
	localVarHTTPHeaderAccepts := []string{}

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
		return localVarHTTPResponse, newErr
	}

	return localVarHTTPResponse, nil
}

type ApiGetCategorySubscriptionsRequest struct {
	ctx context.Context
	ApiService CategoriesUsersAPI
	userId float32
	direction *string
	page *float32
	perPage *float32
	sort *string
}

// The sort direction of the results.  Option descriptions:  * &#x60;asc&#x60; - Sort the results in ascending order.  * &#x60;desc&#x60; - Sort the results in descending order. 
func (r ApiGetCategorySubscriptionsRequest) Direction(direction string) ApiGetCategorySubscriptionsRequest {
	r.direction = &direction
	return r
}

// The page number of the results to show.
func (r ApiGetCategorySubscriptionsRequest) Page(page float32) ApiGetCategorySubscriptionsRequest {
	r.page = &page
	return r
}

// The number of items to show on each page of results, up to a maximum of 100.
func (r ApiGetCategorySubscriptionsRequest) PerPage(perPage float32) ApiGetCategorySubscriptionsRequest {
	r.perPage = &perPage
	return r
}

// The way to sort the results.  Option descriptions:  * &#x60;alphabetical&#x60; - Sort the results alphabetically.  * &#x60;date&#x60; - Sort the results by date.  * &#x60;name&#x60; - Sort the results by name. 
func (r ApiGetCategorySubscriptionsRequest) Sort(sort string) ApiGetCategorySubscriptionsRequest {
	r.sort = &sort
	return r
}

func (r ApiGetCategorySubscriptionsRequest) Execute() ([]Category, *http.Response, error) {
	return r.ApiService.GetCategorySubscriptionsExecute(r)
}

/*
GetCategorySubscriptions Get all the categories that the user follows

This method returns every category that the authenticated user follows.

 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @param userId The ID of the user.
 @return ApiGetCategorySubscriptionsRequest
*/
func (a *CategoriesUsersAPIService) GetCategorySubscriptions(ctx context.Context, userId float32) ApiGetCategorySubscriptionsRequest {
	return ApiGetCategorySubscriptionsRequest{
		ApiService: a,
		ctx: ctx,
		userId: userId,
	}
}

// Execute executes the request
//  @return []Category
func (a *CategoriesUsersAPIService) GetCategorySubscriptionsExecute(r ApiGetCategorySubscriptionsRequest) ([]Category, *http.Response, error) {
	var (
		localVarHTTPMethod   = http.MethodGet
		localVarPostBody     interface{}
		formFiles            []formFile
		localVarReturnValue  []Category
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "CategoriesUsersAPIService.GetCategorySubscriptions")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/users/{user_id}/categories"
	localVarPath = strings.Replace(localVarPath, "{"+"user_id"+"}", url.PathEscape(parameterValueToString(r.userId, "userId")), -1)

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}

	if r.direction != nil {
		parameterAddToHeaderOrQuery(localVarQueryParams, "direction", r.direction, "")
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

type ApiGetCategorySubscriptionsAlt1Request struct {
	ctx context.Context
	ApiService CategoriesUsersAPI
	direction *string
	page *float32
	perPage *float32
	sort *string
}

// The sort direction of the results.  Option descriptions:  * &#x60;asc&#x60; - Sort the results in ascending order.  * &#x60;desc&#x60; - Sort the results in descending order. 
func (r ApiGetCategorySubscriptionsAlt1Request) Direction(direction string) ApiGetCategorySubscriptionsAlt1Request {
	r.direction = &direction
	return r
}

// The page number of the results to show.
func (r ApiGetCategorySubscriptionsAlt1Request) Page(page float32) ApiGetCategorySubscriptionsAlt1Request {
	r.page = &page
	return r
}

// The number of items to show on each page of results, up to a maximum of 100.
func (r ApiGetCategorySubscriptionsAlt1Request) PerPage(perPage float32) ApiGetCategorySubscriptionsAlt1Request {
	r.perPage = &perPage
	return r
}

// The way to sort the results.  Option descriptions:  * &#x60;alphabetical&#x60; - Sort the results alphabetically.  * &#x60;date&#x60; - Sort the results by date.  * &#x60;name&#x60; - Sort the results by name. 
func (r ApiGetCategorySubscriptionsAlt1Request) Sort(sort string) ApiGetCategorySubscriptionsAlt1Request {
	r.sort = &sort
	return r
}

func (r ApiGetCategorySubscriptionsAlt1Request) Execute() ([]Category, *http.Response, error) {
	return r.ApiService.GetCategorySubscriptionsAlt1Execute(r)
}

/*
GetCategorySubscriptionsAlt1 Get all the categories that the user follows

This method returns every category that the authenticated user follows.

 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @return ApiGetCategorySubscriptionsAlt1Request
*/
func (a *CategoriesUsersAPIService) GetCategorySubscriptionsAlt1(ctx context.Context) ApiGetCategorySubscriptionsAlt1Request {
	return ApiGetCategorySubscriptionsAlt1Request{
		ApiService: a,
		ctx: ctx,
	}
}

// Execute executes the request
//  @return []Category
func (a *CategoriesUsersAPIService) GetCategorySubscriptionsAlt1Execute(r ApiGetCategorySubscriptionsAlt1Request) ([]Category, *http.Response, error) {
	var (
		localVarHTTPMethod   = http.MethodGet
		localVarPostBody     interface{}
		formFiles            []formFile
		localVarReturnValue  []Category
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "CategoriesUsersAPIService.GetCategorySubscriptionsAlt1")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/me/categories"

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}

	if r.direction != nil {
		parameterAddToHeaderOrQuery(localVarQueryParams, "direction", r.direction, "")
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

type ApiSubscribeToCategoryRequest struct {
	ctx context.Context
	ApiService CategoriesUsersAPI
	category float32
	userId float32
}

func (r ApiSubscribeToCategoryRequest) Execute() (*http.Response, error) {
	return r.ApiService.SubscribeToCategoryExecute(r)
}

/*
SubscribeToCategory Cause the user to follow a specific category

This method causes the authenticated user to follow the specified category.

 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @param category The name of the category.
 @param userId The ID of the user.
 @return ApiSubscribeToCategoryRequest
*/
func (a *CategoriesUsersAPIService) SubscribeToCategory(ctx context.Context, category float32, userId float32) ApiSubscribeToCategoryRequest {
	return ApiSubscribeToCategoryRequest{
		ApiService: a,
		ctx: ctx,
		category: category,
		userId: userId,
	}
}

// Execute executes the request
func (a *CategoriesUsersAPIService) SubscribeToCategoryExecute(r ApiSubscribeToCategoryRequest) (*http.Response, error) {
	var (
		localVarHTTPMethod   = http.MethodPut
		localVarPostBody     interface{}
		formFiles            []formFile
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "CategoriesUsersAPIService.SubscribeToCategory")
	if err != nil {
		return nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/users/{user_id}/categories/{category}"
	localVarPath = strings.Replace(localVarPath, "{"+"category"+"}", url.PathEscape(parameterValueToString(r.category, "category")), -1)
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
	localVarHTTPHeaderAccepts := []string{}

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
		return localVarHTTPResponse, newErr
	}

	return localVarHTTPResponse, nil
}

type ApiSubscribeToCategoryAlt1Request struct {
	ctx context.Context
	ApiService CategoriesUsersAPI
	category float32
}

func (r ApiSubscribeToCategoryAlt1Request) Execute() (*http.Response, error) {
	return r.ApiService.SubscribeToCategoryAlt1Execute(r)
}

/*
SubscribeToCategoryAlt1 Cause the user to follow a specific category

This method causes the authenticated user to follow the specified category.

 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @param category The name of the category.
 @return ApiSubscribeToCategoryAlt1Request
*/
func (a *CategoriesUsersAPIService) SubscribeToCategoryAlt1(ctx context.Context, category float32) ApiSubscribeToCategoryAlt1Request {
	return ApiSubscribeToCategoryAlt1Request{
		ApiService: a,
		ctx: ctx,
		category: category,
	}
}

// Execute executes the request
func (a *CategoriesUsersAPIService) SubscribeToCategoryAlt1Execute(r ApiSubscribeToCategoryAlt1Request) (*http.Response, error) {
	var (
		localVarHTTPMethod   = http.MethodPut
		localVarPostBody     interface{}
		formFiles            []formFile
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "CategoriesUsersAPIService.SubscribeToCategoryAlt1")
	if err != nil {
		return nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/me/categories/{category}"
	localVarPath = strings.Replace(localVarPath, "{"+"category"+"}", url.PathEscape(parameterValueToString(r.category, "category")), -1)

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
	localVarHTTPHeaderAccepts := []string{}

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
		return localVarHTTPResponse, newErr
	}

	return localVarHTTPResponse, nil
}

type ApiUnsubscribeFromCategoryRequest struct {
	ctx context.Context
	ApiService CategoriesUsersAPI
	category string
	userId float32
}

func (r ApiUnsubscribeFromCategoryRequest) Execute() (*http.Response, error) {
	return r.ApiService.UnsubscribeFromCategoryExecute(r)
}

/*
UnsubscribeFromCategory Cause the user to stop following a category

This method causes the authenticated user to stop following the specified category.

 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @param category The name of the category.
 @param userId The ID of the user.
 @return ApiUnsubscribeFromCategoryRequest
*/
func (a *CategoriesUsersAPIService) UnsubscribeFromCategory(ctx context.Context, category string, userId float32) ApiUnsubscribeFromCategoryRequest {
	return ApiUnsubscribeFromCategoryRequest{
		ApiService: a,
		ctx: ctx,
		category: category,
		userId: userId,
	}
}

// Execute executes the request
func (a *CategoriesUsersAPIService) UnsubscribeFromCategoryExecute(r ApiUnsubscribeFromCategoryRequest) (*http.Response, error) {
	var (
		localVarHTTPMethod   = http.MethodDelete
		localVarPostBody     interface{}
		formFiles            []formFile
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "CategoriesUsersAPIService.UnsubscribeFromCategory")
	if err != nil {
		return nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/users/{user_id}/categories/{category}"
	localVarPath = strings.Replace(localVarPath, "{"+"category"+"}", url.PathEscape(parameterValueToString(r.category, "category")), -1)
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
	localVarHTTPHeaderAccepts := []string{}

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
		return localVarHTTPResponse, newErr
	}

	return localVarHTTPResponse, nil
}

type ApiUnsubscribeFromCategoryAlt1Request struct {
	ctx context.Context
	ApiService CategoriesUsersAPI
	category string
}

func (r ApiUnsubscribeFromCategoryAlt1Request) Execute() (*http.Response, error) {
	return r.ApiService.UnsubscribeFromCategoryAlt1Execute(r)
}

/*
UnsubscribeFromCategoryAlt1 Cause the user to stop following a category

This method causes the authenticated user to stop following the specified category.

 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @param category The name of the category.
 @return ApiUnsubscribeFromCategoryAlt1Request
*/
func (a *CategoriesUsersAPIService) UnsubscribeFromCategoryAlt1(ctx context.Context, category string) ApiUnsubscribeFromCategoryAlt1Request {
	return ApiUnsubscribeFromCategoryAlt1Request{
		ApiService: a,
		ctx: ctx,
		category: category,
	}
}

// Execute executes the request
func (a *CategoriesUsersAPIService) UnsubscribeFromCategoryAlt1Execute(r ApiUnsubscribeFromCategoryAlt1Request) (*http.Response, error) {
	var (
		localVarHTTPMethod   = http.MethodDelete
		localVarPostBody     interface{}
		formFiles            []formFile
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "CategoriesUsersAPIService.UnsubscribeFromCategoryAlt1")
	if err != nil {
		return nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/me/categories/{category}"
	localVarPath = strings.Replace(localVarPath, "{"+"category"+"}", url.PathEscape(parameterValueToString(r.category, "category")), -1)

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
	localVarHTTPHeaderAccepts := []string{}

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
		return localVarHTTPResponse, newErr
	}

	return localVarHTTPResponse, nil
}