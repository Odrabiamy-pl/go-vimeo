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


type VideosEmbedPrivacyAPI interface {

	/*
	AddVideoPrivacyDomain Add a domain to a video's whitelist

	This method adds the specified domain to a video's whitelist.

	@param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
	@param domain The domain name.
	@param videoId The ID of the video.
	@return ApiAddVideoPrivacyDomainRequest
	*/
	AddVideoPrivacyDomain(ctx context.Context, domain string, videoId int32) ApiAddVideoPrivacyDomainRequest

	// AddVideoPrivacyDomainExecute executes the request
	AddVideoPrivacyDomainExecute(r ApiAddVideoPrivacyDomainRequest) (*http.Response, error)

	/*
	DeleteVideoPrivacyDomain Remove a domain from a video's whitelist

	This method removes the specified domain from a video's whitelist.

	@param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
	@param domain The domain name.
	@param videoId The ID of the video.
	@return ApiDeleteVideoPrivacyDomainRequest
	*/
	DeleteVideoPrivacyDomain(ctx context.Context, domain string, videoId int32) ApiDeleteVideoPrivacyDomainRequest

	// DeleteVideoPrivacyDomainExecute executes the request
	DeleteVideoPrivacyDomainExecute(r ApiDeleteVideoPrivacyDomainRequest) (*http.Response, error)

	/*
	GetVideoPrivacyDomains Get all the domains on a video's whitelist

	This method returns every domain on the specified video's whitelist.

	@param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
	@param videoId The ID of the video.
	@return ApiGetVideoPrivacyDomainsRequest
	*/
	GetVideoPrivacyDomains(ctx context.Context, videoId int32) ApiGetVideoPrivacyDomainsRequest

	// GetVideoPrivacyDomainsExecute executes the request
	//  @return []Domain
	GetVideoPrivacyDomainsExecute(r ApiGetVideoPrivacyDomainsRequest) ([]Domain, *http.Response, error)
}

// VideosEmbedPrivacyAPIService VideosEmbedPrivacyAPI service
type VideosEmbedPrivacyAPIService service

type ApiAddVideoPrivacyDomainRequest struct {
	ctx context.Context
	ApiService VideosEmbedPrivacyAPI
	domain string
	videoId int32
}

func (r ApiAddVideoPrivacyDomainRequest) Execute() (*http.Response, error) {
	return r.ApiService.AddVideoPrivacyDomainExecute(r)
}

/*
AddVideoPrivacyDomain Add a domain to a video's whitelist

This method adds the specified domain to a video's whitelist.

 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @param domain The domain name.
 @param videoId The ID of the video.
 @return ApiAddVideoPrivacyDomainRequest
*/
func (a *VideosEmbedPrivacyAPIService) AddVideoPrivacyDomain(ctx context.Context, domain string, videoId int32) ApiAddVideoPrivacyDomainRequest {
	return ApiAddVideoPrivacyDomainRequest{
		ApiService: a,
		ctx: ctx,
		domain: domain,
		videoId: videoId,
	}
}

// Execute executes the request
func (a *VideosEmbedPrivacyAPIService) AddVideoPrivacyDomainExecute(r ApiAddVideoPrivacyDomainRequest) (*http.Response, error) {
	var (
		localVarHTTPMethod   = http.MethodPut
		localVarPostBody     interface{}
		formFiles            []formFile
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "VideosEmbedPrivacyAPIService.AddVideoPrivacyDomain")
	if err != nil {
		return nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/videos/{video_id}/privacy/domains/{domain}"
	localVarPath = strings.Replace(localVarPath, "{"+"domain"+"}", url.PathEscape(parameterValueToString(r.domain, "domain")), -1)
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
		}
		return localVarHTTPResponse, newErr
	}

	return localVarHTTPResponse, nil
}

type ApiDeleteVideoPrivacyDomainRequest struct {
	ctx context.Context
	ApiService VideosEmbedPrivacyAPI
	domain string
	videoId int32
}

func (r ApiDeleteVideoPrivacyDomainRequest) Execute() (*http.Response, error) {
	return r.ApiService.DeleteVideoPrivacyDomainExecute(r)
}

/*
DeleteVideoPrivacyDomain Remove a domain from a video's whitelist

This method removes the specified domain from a video's whitelist.

 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @param domain The domain name.
 @param videoId The ID of the video.
 @return ApiDeleteVideoPrivacyDomainRequest
*/
func (a *VideosEmbedPrivacyAPIService) DeleteVideoPrivacyDomain(ctx context.Context, domain string, videoId int32) ApiDeleteVideoPrivacyDomainRequest {
	return ApiDeleteVideoPrivacyDomainRequest{
		ApiService: a,
		ctx: ctx,
		domain: domain,
		videoId: videoId,
	}
}

// Execute executes the request
func (a *VideosEmbedPrivacyAPIService) DeleteVideoPrivacyDomainExecute(r ApiDeleteVideoPrivacyDomainRequest) (*http.Response, error) {
	var (
		localVarHTTPMethod   = http.MethodDelete
		localVarPostBody     interface{}
		formFiles            []formFile
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "VideosEmbedPrivacyAPIService.DeleteVideoPrivacyDomain")
	if err != nil {
		return nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/videos/{video_id}/privacy/domains/{domain}"
	localVarPath = strings.Replace(localVarPath, "{"+"domain"+"}", url.PathEscape(parameterValueToString(r.domain, "domain")), -1)
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

type ApiGetVideoPrivacyDomainsRequest struct {
	ctx context.Context
	ApiService VideosEmbedPrivacyAPI
	videoId int32
	page *float32
	perPage *float32
}

// The page number of the results to show.
func (r ApiGetVideoPrivacyDomainsRequest) Page(page float32) ApiGetVideoPrivacyDomainsRequest {
	r.page = &page
	return r
}

// The number of items to show on each page of results, up to a maximum of 100.
func (r ApiGetVideoPrivacyDomainsRequest) PerPage(perPage float32) ApiGetVideoPrivacyDomainsRequest {
	r.perPage = &perPage
	return r
}

func (r ApiGetVideoPrivacyDomainsRequest) Execute() ([]Domain, *http.Response, error) {
	return r.ApiService.GetVideoPrivacyDomainsExecute(r)
}

/*
GetVideoPrivacyDomains Get all the domains on a video's whitelist

This method returns every domain on the specified video's whitelist.

 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @param videoId The ID of the video.
 @return ApiGetVideoPrivacyDomainsRequest
*/
func (a *VideosEmbedPrivacyAPIService) GetVideoPrivacyDomains(ctx context.Context, videoId int32) ApiGetVideoPrivacyDomainsRequest {
	return ApiGetVideoPrivacyDomainsRequest{
		ApiService: a,
		ctx: ctx,
		videoId: videoId,
	}
}

// Execute executes the request
//  @return []Domain
func (a *VideosEmbedPrivacyAPIService) GetVideoPrivacyDomainsExecute(r ApiGetVideoPrivacyDomainsRequest) ([]Domain, *http.Response, error) {
	var (
		localVarHTTPMethod   = http.MethodGet
		localVarPostBody     interface{}
		formFiles            []formFile
		localVarReturnValue  []Domain
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "VideosEmbedPrivacyAPIService.GetVideoPrivacyDomains")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/videos/{video_id}/privacy/domains"
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
	localVarHTTPHeaderAccepts := []string{"application/vnd.vimeo.domain+json"}

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
