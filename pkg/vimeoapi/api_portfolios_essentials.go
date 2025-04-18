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


type PortfoliosEssentialsAPI interface {

	/*
	GetPortfolio Get a specific portfolio

	This method returns a single portfolio belonging to the authenticated user.

	@param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
	@param portfolioId The ID of the portfolio.
	@param userId The ID of the user.
	@return ApiGetPortfolioRequest
	*/
	GetPortfolio(ctx context.Context, portfolioId float32, userId int32) ApiGetPortfolioRequest

	// GetPortfolioExecute executes the request
	//  @return Portfolio
	GetPortfolioExecute(r ApiGetPortfolioRequest) (*Portfolio, *http.Response, error)

	/*
	GetPortfolioAlt1 Get a specific portfolio

	This method returns a single portfolio belonging to the authenticated user.

	@param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
	@param portfolioId The ID of the portfolio.
	@return ApiGetPortfolioAlt1Request
	*/
	GetPortfolioAlt1(ctx context.Context, portfolioId float32) ApiGetPortfolioAlt1Request

	// GetPortfolioAlt1Execute executes the request
	//  @return Portfolio
	GetPortfolioAlt1Execute(r ApiGetPortfolioAlt1Request) (*Portfolio, *http.Response, error)

	/*
	GetPortfolios Get all the portfolios that belong to the user

	This method returns every portfolio belonging to the authenticated user.

	@param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
	@param userId The ID of the user.
	@return ApiGetPortfoliosRequest
	*/
	GetPortfolios(ctx context.Context, userId int32) ApiGetPortfoliosRequest

	// GetPortfoliosExecute executes the request
	//  @return []Portfolio
	GetPortfoliosExecute(r ApiGetPortfoliosRequest) ([]Portfolio, *http.Response, error)

	/*
	GetPortfoliosAlt1 Get all the portfolios that belong to the user

	This method returns every portfolio belonging to the authenticated user.

	@param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
	@return ApiGetPortfoliosAlt1Request
	*/
	GetPortfoliosAlt1(ctx context.Context) ApiGetPortfoliosAlt1Request

	// GetPortfoliosAlt1Execute executes the request
	//  @return []Portfolio
	GetPortfoliosAlt1Execute(r ApiGetPortfoliosAlt1Request) ([]Portfolio, *http.Response, error)
}

// PortfoliosEssentialsAPIService PortfoliosEssentialsAPI service
type PortfoliosEssentialsAPIService service

type ApiGetPortfolioRequest struct {
	ctx context.Context
	ApiService PortfoliosEssentialsAPI
	portfolioId float32
	userId int32
}

func (r ApiGetPortfolioRequest) Execute() (*Portfolio, *http.Response, error) {
	return r.ApiService.GetPortfolioExecute(r)
}

/*
GetPortfolio Get a specific portfolio

This method returns a single portfolio belonging to the authenticated user.

 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @param portfolioId The ID of the portfolio.
 @param userId The ID of the user.
 @return ApiGetPortfolioRequest
*/
func (a *PortfoliosEssentialsAPIService) GetPortfolio(ctx context.Context, portfolioId float32, userId int32) ApiGetPortfolioRequest {
	return ApiGetPortfolioRequest{
		ApiService: a,
		ctx: ctx,
		portfolioId: portfolioId,
		userId: userId,
	}
}

// Execute executes the request
//  @return Portfolio
func (a *PortfoliosEssentialsAPIService) GetPortfolioExecute(r ApiGetPortfolioRequest) (*Portfolio, *http.Response, error) {
	var (
		localVarHTTPMethod   = http.MethodGet
		localVarPostBody     interface{}
		formFiles            []formFile
		localVarReturnValue  *Portfolio
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "PortfoliosEssentialsAPIService.GetPortfolio")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/users/{user_id}/portfolios/{portfolio_id}"
	localVarPath = strings.Replace(localVarPath, "{"+"portfolio_id"+"}", url.PathEscape(parameterValueToString(r.portfolioId, "portfolioId")), -1)
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
	localVarHTTPHeaderAccepts := []string{"application/vnd.vimeo.portfolio+json"}

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

type ApiGetPortfolioAlt1Request struct {
	ctx context.Context
	ApiService PortfoliosEssentialsAPI
	portfolioId float32
}

func (r ApiGetPortfolioAlt1Request) Execute() (*Portfolio, *http.Response, error) {
	return r.ApiService.GetPortfolioAlt1Execute(r)
}

/*
GetPortfolioAlt1 Get a specific portfolio

This method returns a single portfolio belonging to the authenticated user.

 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @param portfolioId The ID of the portfolio.
 @return ApiGetPortfolioAlt1Request
*/
func (a *PortfoliosEssentialsAPIService) GetPortfolioAlt1(ctx context.Context, portfolioId float32) ApiGetPortfolioAlt1Request {
	return ApiGetPortfolioAlt1Request{
		ApiService: a,
		ctx: ctx,
		portfolioId: portfolioId,
	}
}

// Execute executes the request
//  @return Portfolio
func (a *PortfoliosEssentialsAPIService) GetPortfolioAlt1Execute(r ApiGetPortfolioAlt1Request) (*Portfolio, *http.Response, error) {
	var (
		localVarHTTPMethod   = http.MethodGet
		localVarPostBody     interface{}
		formFiles            []formFile
		localVarReturnValue  *Portfolio
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "PortfoliosEssentialsAPIService.GetPortfolioAlt1")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/me/portfolios/{portfolio_id}"
	localVarPath = strings.Replace(localVarPath, "{"+"portfolio_id"+"}", url.PathEscape(parameterValueToString(r.portfolioId, "portfolioId")), -1)

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
	localVarHTTPHeaderAccepts := []string{"application/vnd.vimeo.portfolio+json"}

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

type ApiGetPortfoliosRequest struct {
	ctx context.Context
	ApiService PortfoliosEssentialsAPI
	userId int32
	direction *string
	page *float32
	perPage *float32
	query *string
	sort *string
}

// The sort direction of the results.  Option descriptions:  * &#x60;asc&#x60; - Sort the results in ascending order.  * &#x60;desc&#x60; - Sort the results in descending order. 
func (r ApiGetPortfoliosRequest) Direction(direction string) ApiGetPortfoliosRequest {
	r.direction = &direction
	return r
}

// The page number of the results to show.
func (r ApiGetPortfoliosRequest) Page(page float32) ApiGetPortfoliosRequest {
	r.page = &page
	return r
}

// The number of items to show on each page of results, up to a maximum of 100.
func (r ApiGetPortfoliosRequest) PerPage(perPage float32) ApiGetPortfoliosRequest {
	r.perPage = &perPage
	return r
}

// The search query to use to filter the results.
func (r ApiGetPortfoliosRequest) Query(query string) ApiGetPortfoliosRequest {
	r.query = &query
	return r
}

// The way to sort the results.  Option descriptions:  * &#x60;alphabetical&#x60; - Sort the results alphabetically.  * &#x60;date&#x60; - Sort the results by creation date. 
func (r ApiGetPortfoliosRequest) Sort(sort string) ApiGetPortfoliosRequest {
	r.sort = &sort
	return r
}

func (r ApiGetPortfoliosRequest) Execute() ([]Portfolio, *http.Response, error) {
	return r.ApiService.GetPortfoliosExecute(r)
}

/*
GetPortfolios Get all the portfolios that belong to the user

This method returns every portfolio belonging to the authenticated user.

 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @param userId The ID of the user.
 @return ApiGetPortfoliosRequest
*/
func (a *PortfoliosEssentialsAPIService) GetPortfolios(ctx context.Context, userId int32) ApiGetPortfoliosRequest {
	return ApiGetPortfoliosRequest{
		ApiService: a,
		ctx: ctx,
		userId: userId,
	}
}

// Execute executes the request
//  @return []Portfolio
func (a *PortfoliosEssentialsAPIService) GetPortfoliosExecute(r ApiGetPortfoliosRequest) ([]Portfolio, *http.Response, error) {
	var (
		localVarHTTPMethod   = http.MethodGet
		localVarPostBody     interface{}
		formFiles            []formFile
		localVarReturnValue  []Portfolio
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "PortfoliosEssentialsAPIService.GetPortfolios")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/users/{user_id}/portfolios"
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
	localVarHTTPHeaderAccepts := []string{"application/vnd.vimeo.portfolio+json"}

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

type ApiGetPortfoliosAlt1Request struct {
	ctx context.Context
	ApiService PortfoliosEssentialsAPI
	direction *string
	page *float32
	perPage *float32
	query *string
	sort *string
}

// The sort direction of the results.  Option descriptions:  * &#x60;asc&#x60; - Sort the results in ascending order.  * &#x60;desc&#x60; - Sort the results in descending order. 
func (r ApiGetPortfoliosAlt1Request) Direction(direction string) ApiGetPortfoliosAlt1Request {
	r.direction = &direction
	return r
}

// The page number of the results to show.
func (r ApiGetPortfoliosAlt1Request) Page(page float32) ApiGetPortfoliosAlt1Request {
	r.page = &page
	return r
}

// The number of items to show on each page of results, up to a maximum of 100.
func (r ApiGetPortfoliosAlt1Request) PerPage(perPage float32) ApiGetPortfoliosAlt1Request {
	r.perPage = &perPage
	return r
}

// The search query to use to filter the results.
func (r ApiGetPortfoliosAlt1Request) Query(query string) ApiGetPortfoliosAlt1Request {
	r.query = &query
	return r
}

// The way to sort the results.  Option descriptions:  * &#x60;alphabetical&#x60; - Sort the results alphabetically.  * &#x60;date&#x60; - Sort the results by creation date. 
func (r ApiGetPortfoliosAlt1Request) Sort(sort string) ApiGetPortfoliosAlt1Request {
	r.sort = &sort
	return r
}

func (r ApiGetPortfoliosAlt1Request) Execute() ([]Portfolio, *http.Response, error) {
	return r.ApiService.GetPortfoliosAlt1Execute(r)
}

/*
GetPortfoliosAlt1 Get all the portfolios that belong to the user

This method returns every portfolio belonging to the authenticated user.

 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @return ApiGetPortfoliosAlt1Request
*/
func (a *PortfoliosEssentialsAPIService) GetPortfoliosAlt1(ctx context.Context) ApiGetPortfoliosAlt1Request {
	return ApiGetPortfoliosAlt1Request{
		ApiService: a,
		ctx: ctx,
	}
}

// Execute executes the request
//  @return []Portfolio
func (a *PortfoliosEssentialsAPIService) GetPortfoliosAlt1Execute(r ApiGetPortfoliosAlt1Request) ([]Portfolio, *http.Response, error) {
	var (
		localVarHTTPMethod   = http.MethodGet
		localVarPostBody     interface{}
		formFiles            []formFile
		localVarReturnValue  []Portfolio
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "PortfoliosEssentialsAPIService.GetPortfoliosAlt1")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/me/portfolios"

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
	localVarHTTPHeaderAccepts := []string{"application/vnd.vimeo.portfolio+json"}

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
