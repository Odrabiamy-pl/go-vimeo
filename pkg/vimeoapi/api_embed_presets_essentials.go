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


type EmbedPresetsEssentialsAPI interface {

	/*
	CreateEmbedPresets Create an embed preset

	This method creates an embed preset.

	@param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
	@param userId The ID of the user.
	@return ApiCreateEmbedPresetsRequest
	*/
	CreateEmbedPresets(ctx context.Context, userId int32) ApiCreateEmbedPresetsRequest

	// CreateEmbedPresetsExecute executes the request
	//  @return Preset
	CreateEmbedPresetsExecute(r ApiCreateEmbedPresetsRequest) (*Preset, *http.Response, error)

	/*
	CreateEmbedPresetsAlt1 Create an embed preset

	This method creates an embed preset.

	@param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
	@return ApiCreateEmbedPresetsAlt1Request
	*/
	CreateEmbedPresetsAlt1(ctx context.Context) ApiCreateEmbedPresetsAlt1Request

	// CreateEmbedPresetsAlt1Execute executes the request
	//  @return Preset
	CreateEmbedPresetsAlt1Execute(r ApiCreateEmbedPresetsAlt1Request) (*Preset, *http.Response, error)

	/*
	EditEmbedPreset Edit an embed preset

	This method edits the specified embed preset. The authenticated user must be the owner of the preset.

	@param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
	@param presetId The ID of the preset.
	@param userId The ID of the user.
	@return ApiEditEmbedPresetRequest
	*/
	EditEmbedPreset(ctx context.Context, presetId float32, userId int32) ApiEditEmbedPresetRequest

	// EditEmbedPresetExecute executes the request
	//  @return Preset
	EditEmbedPresetExecute(r ApiEditEmbedPresetRequest) (*Preset, *http.Response, error)

	/*
	EditEmbedPresetAlt1 Edit an embed preset

	This method edits the specified embed preset. The authenticated user must be the owner of the preset.

	@param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
	@param presetId The ID of the preset.
	@return ApiEditEmbedPresetAlt1Request
	*/
	EditEmbedPresetAlt1(ctx context.Context, presetId float32) ApiEditEmbedPresetAlt1Request

	// EditEmbedPresetAlt1Execute executes the request
	//  @return Preset
	EditEmbedPresetAlt1Execute(r ApiEditEmbedPresetAlt1Request) (*Preset, *http.Response, error)

	/*
	GetEmbedPreset Get a specific embed preset

	This method returns a single embed preset. The authenticated user must be the owner of the preset.

	@param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
	@param presetId The ID of the preset.
	@param userId The ID of the user.
	@return ApiGetEmbedPresetRequest
	*/
	GetEmbedPreset(ctx context.Context, presetId float32, userId int32) ApiGetEmbedPresetRequest

	// GetEmbedPresetExecute executes the request
	//  @return Preset
	GetEmbedPresetExecute(r ApiGetEmbedPresetRequest) (*Preset, *http.Response, error)

	/*
	GetEmbedPresetAlt1 Get a specific embed preset

	This method returns a single embed preset. The authenticated user must be the owner of the preset.

	@param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
	@param presetId The ID of the preset.
	@return ApiGetEmbedPresetAlt1Request
	*/
	GetEmbedPresetAlt1(ctx context.Context, presetId float32) ApiGetEmbedPresetAlt1Request

	// GetEmbedPresetAlt1Execute executes the request
	//  @return Preset
	GetEmbedPresetAlt1Execute(r ApiGetEmbedPresetAlt1Request) (*Preset, *http.Response, error)

	/*
	GetEmbedPresets Get all the embed presets that a user has created

	This method returns every embed preset that belongs to the authenticated user.

	@param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
	@param userId The ID of the user.
	@return ApiGetEmbedPresetsRequest
	*/
	GetEmbedPresets(ctx context.Context, userId int32) ApiGetEmbedPresetsRequest

	// GetEmbedPresetsExecute executes the request
	//  @return []Preset
	GetEmbedPresetsExecute(r ApiGetEmbedPresetsRequest) ([]Preset, *http.Response, error)

	/*
	GetEmbedPresetsAlt1 Get all the embed presets that a user has created

	This method returns every embed preset that belongs to the authenticated user.

	@param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
	@return ApiGetEmbedPresetsAlt1Request
	*/
	GetEmbedPresetsAlt1(ctx context.Context) ApiGetEmbedPresetsAlt1Request

	// GetEmbedPresetsAlt1Execute executes the request
	//  @return []Preset
	GetEmbedPresetsAlt1Execute(r ApiGetEmbedPresetsAlt1Request) ([]Preset, *http.Response, error)
}

// EmbedPresetsEssentialsAPIService EmbedPresetsEssentialsAPI service
type EmbedPresetsEssentialsAPIService service

type ApiCreateEmbedPresetsRequest struct {
	ctx context.Context
	ApiService EmbedPresetsEssentialsAPI
	userId int32
	createEmbedPresetsAlt1Request *CreateEmbedPresetsAlt1Request
}

func (r ApiCreateEmbedPresetsRequest) CreateEmbedPresetsAlt1Request(createEmbedPresetsAlt1Request CreateEmbedPresetsAlt1Request) ApiCreateEmbedPresetsRequest {
	r.createEmbedPresetsAlt1Request = &createEmbedPresetsAlt1Request
	return r
}

func (r ApiCreateEmbedPresetsRequest) Execute() (*Preset, *http.Response, error) {
	return r.ApiService.CreateEmbedPresetsExecute(r)
}

/*
CreateEmbedPresets Create an embed preset

This method creates an embed preset.

 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @param userId The ID of the user.
 @return ApiCreateEmbedPresetsRequest
*/
func (a *EmbedPresetsEssentialsAPIService) CreateEmbedPresets(ctx context.Context, userId int32) ApiCreateEmbedPresetsRequest {
	return ApiCreateEmbedPresetsRequest{
		ApiService: a,
		ctx: ctx,
		userId: userId,
	}
}

// Execute executes the request
//  @return Preset
func (a *EmbedPresetsEssentialsAPIService) CreateEmbedPresetsExecute(r ApiCreateEmbedPresetsRequest) (*Preset, *http.Response, error) {
	var (
		localVarHTTPMethod   = http.MethodPost
		localVarPostBody     interface{}
		formFiles            []formFile
		localVarReturnValue  *Preset
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "EmbedPresetsEssentialsAPIService.CreateEmbedPresets")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/users/{user_id}/presets"
	localVarPath = strings.Replace(localVarPath, "{"+"user_id"+"}", url.PathEscape(parameterValueToString(r.userId, "userId")), -1)

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}

	// to determine the Content-Type header
	localVarHTTPContentTypes := []string{"application/vnd.vimeo.preset+json"}

	// set Content-Type header
	localVarHTTPContentType := selectHeaderContentType(localVarHTTPContentTypes)
	if localVarHTTPContentType != "" {
		localVarHeaderParams["Content-Type"] = localVarHTTPContentType
	}

	// to determine the Accept header
	localVarHTTPHeaderAccepts := []string{"application/vnd.vimeo.preset+json"}

	// set Accept header
	localVarHTTPHeaderAccept := selectHeaderAccept(localVarHTTPHeaderAccepts)
	if localVarHTTPHeaderAccept != "" {
		localVarHeaderParams["Accept"] = localVarHTTPHeaderAccept
	}
	// body params
	localVarPostBody = r.createEmbedPresetsAlt1Request
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

type ApiCreateEmbedPresetsAlt1Request struct {
	ctx context.Context
	ApiService EmbedPresetsEssentialsAPI
	createEmbedPresetsAlt1Request *CreateEmbedPresetsAlt1Request
}

func (r ApiCreateEmbedPresetsAlt1Request) CreateEmbedPresetsAlt1Request(createEmbedPresetsAlt1Request CreateEmbedPresetsAlt1Request) ApiCreateEmbedPresetsAlt1Request {
	r.createEmbedPresetsAlt1Request = &createEmbedPresetsAlt1Request
	return r
}

func (r ApiCreateEmbedPresetsAlt1Request) Execute() (*Preset, *http.Response, error) {
	return r.ApiService.CreateEmbedPresetsAlt1Execute(r)
}

/*
CreateEmbedPresetsAlt1 Create an embed preset

This method creates an embed preset.

 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @return ApiCreateEmbedPresetsAlt1Request
*/
func (a *EmbedPresetsEssentialsAPIService) CreateEmbedPresetsAlt1(ctx context.Context) ApiCreateEmbedPresetsAlt1Request {
	return ApiCreateEmbedPresetsAlt1Request{
		ApiService: a,
		ctx: ctx,
	}
}

// Execute executes the request
//  @return Preset
func (a *EmbedPresetsEssentialsAPIService) CreateEmbedPresetsAlt1Execute(r ApiCreateEmbedPresetsAlt1Request) (*Preset, *http.Response, error) {
	var (
		localVarHTTPMethod   = http.MethodPost
		localVarPostBody     interface{}
		formFiles            []formFile
		localVarReturnValue  *Preset
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "EmbedPresetsEssentialsAPIService.CreateEmbedPresetsAlt1")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/me/presets"

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}

	// to determine the Content-Type header
	localVarHTTPContentTypes := []string{"application/vnd.vimeo.preset+json"}

	// set Content-Type header
	localVarHTTPContentType := selectHeaderContentType(localVarHTTPContentTypes)
	if localVarHTTPContentType != "" {
		localVarHeaderParams["Content-Type"] = localVarHTTPContentType
	}

	// to determine the Accept header
	localVarHTTPHeaderAccepts := []string{"application/vnd.vimeo.preset+json"}

	// set Accept header
	localVarHTTPHeaderAccept := selectHeaderAccept(localVarHTTPHeaderAccepts)
	if localVarHTTPHeaderAccept != "" {
		localVarHeaderParams["Accept"] = localVarHTTPHeaderAccept
	}
	// body params
	localVarPostBody = r.createEmbedPresetsAlt1Request
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

type ApiEditEmbedPresetRequest struct {
	ctx context.Context
	ApiService EmbedPresetsEssentialsAPI
	presetId float32
	userId int32
	editEmbedPresetAlt1Request *EditEmbedPresetAlt1Request
}

func (r ApiEditEmbedPresetRequest) EditEmbedPresetAlt1Request(editEmbedPresetAlt1Request EditEmbedPresetAlt1Request) ApiEditEmbedPresetRequest {
	r.editEmbedPresetAlt1Request = &editEmbedPresetAlt1Request
	return r
}

func (r ApiEditEmbedPresetRequest) Execute() (*Preset, *http.Response, error) {
	return r.ApiService.EditEmbedPresetExecute(r)
}

/*
EditEmbedPreset Edit an embed preset

This method edits the specified embed preset. The authenticated user must be the owner of the preset.

 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @param presetId The ID of the preset.
 @param userId The ID of the user.
 @return ApiEditEmbedPresetRequest
*/
func (a *EmbedPresetsEssentialsAPIService) EditEmbedPreset(ctx context.Context, presetId float32, userId int32) ApiEditEmbedPresetRequest {
	return ApiEditEmbedPresetRequest{
		ApiService: a,
		ctx: ctx,
		presetId: presetId,
		userId: userId,
	}
}

// Execute executes the request
//  @return Preset
func (a *EmbedPresetsEssentialsAPIService) EditEmbedPresetExecute(r ApiEditEmbedPresetRequest) (*Preset, *http.Response, error) {
	var (
		localVarHTTPMethod   = http.MethodPatch
		localVarPostBody     interface{}
		formFiles            []formFile
		localVarReturnValue  *Preset
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "EmbedPresetsEssentialsAPIService.EditEmbedPreset")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/users/{user_id}/presets/{preset_id}"
	localVarPath = strings.Replace(localVarPath, "{"+"preset_id"+"}", url.PathEscape(parameterValueToString(r.presetId, "presetId")), -1)
	localVarPath = strings.Replace(localVarPath, "{"+"user_id"+"}", url.PathEscape(parameterValueToString(r.userId, "userId")), -1)

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}

	// to determine the Content-Type header
	localVarHTTPContentTypes := []string{"application/vnd.vimeo.preset+json"}

	// set Content-Type header
	localVarHTTPContentType := selectHeaderContentType(localVarHTTPContentTypes)
	if localVarHTTPContentType != "" {
		localVarHeaderParams["Content-Type"] = localVarHTTPContentType
	}

	// to determine the Accept header
	localVarHTTPHeaderAccepts := []string{"application/vnd.vimeo.preset+json"}

	// set Accept header
	localVarHTTPHeaderAccept := selectHeaderAccept(localVarHTTPHeaderAccepts)
	if localVarHTTPHeaderAccept != "" {
		localVarHeaderParams["Accept"] = localVarHTTPHeaderAccept
	}
	// body params
	localVarPostBody = r.editEmbedPresetAlt1Request
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

type ApiEditEmbedPresetAlt1Request struct {
	ctx context.Context
	ApiService EmbedPresetsEssentialsAPI
	presetId float32
	editEmbedPresetAlt1Request *EditEmbedPresetAlt1Request
}

func (r ApiEditEmbedPresetAlt1Request) EditEmbedPresetAlt1Request(editEmbedPresetAlt1Request EditEmbedPresetAlt1Request) ApiEditEmbedPresetAlt1Request {
	r.editEmbedPresetAlt1Request = &editEmbedPresetAlt1Request
	return r
}

func (r ApiEditEmbedPresetAlt1Request) Execute() (*Preset, *http.Response, error) {
	return r.ApiService.EditEmbedPresetAlt1Execute(r)
}

/*
EditEmbedPresetAlt1 Edit an embed preset

This method edits the specified embed preset. The authenticated user must be the owner of the preset.

 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @param presetId The ID of the preset.
 @return ApiEditEmbedPresetAlt1Request
*/
func (a *EmbedPresetsEssentialsAPIService) EditEmbedPresetAlt1(ctx context.Context, presetId float32) ApiEditEmbedPresetAlt1Request {
	return ApiEditEmbedPresetAlt1Request{
		ApiService: a,
		ctx: ctx,
		presetId: presetId,
	}
}

// Execute executes the request
//  @return Preset
func (a *EmbedPresetsEssentialsAPIService) EditEmbedPresetAlt1Execute(r ApiEditEmbedPresetAlt1Request) (*Preset, *http.Response, error) {
	var (
		localVarHTTPMethod   = http.MethodPatch
		localVarPostBody     interface{}
		formFiles            []formFile
		localVarReturnValue  *Preset
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "EmbedPresetsEssentialsAPIService.EditEmbedPresetAlt1")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/me/presets/{preset_id}"
	localVarPath = strings.Replace(localVarPath, "{"+"preset_id"+"}", url.PathEscape(parameterValueToString(r.presetId, "presetId")), -1)

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}

	// to determine the Content-Type header
	localVarHTTPContentTypes := []string{"application/vnd.vimeo.preset+json"}

	// set Content-Type header
	localVarHTTPContentType := selectHeaderContentType(localVarHTTPContentTypes)
	if localVarHTTPContentType != "" {
		localVarHeaderParams["Content-Type"] = localVarHTTPContentType
	}

	// to determine the Accept header
	localVarHTTPHeaderAccepts := []string{"application/vnd.vimeo.preset+json"}

	// set Accept header
	localVarHTTPHeaderAccept := selectHeaderAccept(localVarHTTPHeaderAccepts)
	if localVarHTTPHeaderAccept != "" {
		localVarHeaderParams["Accept"] = localVarHTTPHeaderAccept
	}
	// body params
	localVarPostBody = r.editEmbedPresetAlt1Request
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

type ApiGetEmbedPresetRequest struct {
	ctx context.Context
	ApiService EmbedPresetsEssentialsAPI
	presetId float32
	userId int32
}

func (r ApiGetEmbedPresetRequest) Execute() (*Preset, *http.Response, error) {
	return r.ApiService.GetEmbedPresetExecute(r)
}

/*
GetEmbedPreset Get a specific embed preset

This method returns a single embed preset. The authenticated user must be the owner of the preset.

 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @param presetId The ID of the preset.
 @param userId The ID of the user.
 @return ApiGetEmbedPresetRequest
*/
func (a *EmbedPresetsEssentialsAPIService) GetEmbedPreset(ctx context.Context, presetId float32, userId int32) ApiGetEmbedPresetRequest {
	return ApiGetEmbedPresetRequest{
		ApiService: a,
		ctx: ctx,
		presetId: presetId,
		userId: userId,
	}
}

// Execute executes the request
//  @return Preset
func (a *EmbedPresetsEssentialsAPIService) GetEmbedPresetExecute(r ApiGetEmbedPresetRequest) (*Preset, *http.Response, error) {
	var (
		localVarHTTPMethod   = http.MethodGet
		localVarPostBody     interface{}
		formFiles            []formFile
		localVarReturnValue  *Preset
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "EmbedPresetsEssentialsAPIService.GetEmbedPreset")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/users/{user_id}/presets/{preset_id}"
	localVarPath = strings.Replace(localVarPath, "{"+"preset_id"+"}", url.PathEscape(parameterValueToString(r.presetId, "presetId")), -1)
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
	localVarHTTPHeaderAccepts := []string{"application/vnd.vimeo.preset+json"}

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

type ApiGetEmbedPresetAlt1Request struct {
	ctx context.Context
	ApiService EmbedPresetsEssentialsAPI
	presetId float32
}

func (r ApiGetEmbedPresetAlt1Request) Execute() (*Preset, *http.Response, error) {
	return r.ApiService.GetEmbedPresetAlt1Execute(r)
}

/*
GetEmbedPresetAlt1 Get a specific embed preset

This method returns a single embed preset. The authenticated user must be the owner of the preset.

 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @param presetId The ID of the preset.
 @return ApiGetEmbedPresetAlt1Request
*/
func (a *EmbedPresetsEssentialsAPIService) GetEmbedPresetAlt1(ctx context.Context, presetId float32) ApiGetEmbedPresetAlt1Request {
	return ApiGetEmbedPresetAlt1Request{
		ApiService: a,
		ctx: ctx,
		presetId: presetId,
	}
}

// Execute executes the request
//  @return Preset
func (a *EmbedPresetsEssentialsAPIService) GetEmbedPresetAlt1Execute(r ApiGetEmbedPresetAlt1Request) (*Preset, *http.Response, error) {
	var (
		localVarHTTPMethod   = http.MethodGet
		localVarPostBody     interface{}
		formFiles            []formFile
		localVarReturnValue  *Preset
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "EmbedPresetsEssentialsAPIService.GetEmbedPresetAlt1")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/me/presets/{preset_id}"
	localVarPath = strings.Replace(localVarPath, "{"+"preset_id"+"}", url.PathEscape(parameterValueToString(r.presetId, "presetId")), -1)

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
	localVarHTTPHeaderAccepts := []string{"application/vnd.vimeo.preset+json"}

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

type ApiGetEmbedPresetsRequest struct {
	ctx context.Context
	ApiService EmbedPresetsEssentialsAPI
	userId int32
	page *float32
	perPage *float32
}

// The page number of the results to show.
func (r ApiGetEmbedPresetsRequest) Page(page float32) ApiGetEmbedPresetsRequest {
	r.page = &page
	return r
}

// The number of items to show on each page of results, up to a maximum of 100.
func (r ApiGetEmbedPresetsRequest) PerPage(perPage float32) ApiGetEmbedPresetsRequest {
	r.perPage = &perPage
	return r
}

func (r ApiGetEmbedPresetsRequest) Execute() ([]Preset, *http.Response, error) {
	return r.ApiService.GetEmbedPresetsExecute(r)
}

/*
GetEmbedPresets Get all the embed presets that a user has created

This method returns every embed preset that belongs to the authenticated user.

 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @param userId The ID of the user.
 @return ApiGetEmbedPresetsRequest
*/
func (a *EmbedPresetsEssentialsAPIService) GetEmbedPresets(ctx context.Context, userId int32) ApiGetEmbedPresetsRequest {
	return ApiGetEmbedPresetsRequest{
		ApiService: a,
		ctx: ctx,
		userId: userId,
	}
}

// Execute executes the request
//  @return []Preset
func (a *EmbedPresetsEssentialsAPIService) GetEmbedPresetsExecute(r ApiGetEmbedPresetsRequest) ([]Preset, *http.Response, error) {
	var (
		localVarHTTPMethod   = http.MethodGet
		localVarPostBody     interface{}
		formFiles            []formFile
		localVarReturnValue  []Preset
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "EmbedPresetsEssentialsAPIService.GetEmbedPresets")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/users/{user_id}/presets"
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
	localVarHTTPHeaderAccepts := []string{"application/vnd.vimeo.preset+json"}

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

type ApiGetEmbedPresetsAlt1Request struct {
	ctx context.Context
	ApiService EmbedPresetsEssentialsAPI
	page *float32
	perPage *float32
}

// The page number of the results to show.
func (r ApiGetEmbedPresetsAlt1Request) Page(page float32) ApiGetEmbedPresetsAlt1Request {
	r.page = &page
	return r
}

// The number of items to show on each page of results, up to a maximum of 100.
func (r ApiGetEmbedPresetsAlt1Request) PerPage(perPage float32) ApiGetEmbedPresetsAlt1Request {
	r.perPage = &perPage
	return r
}

func (r ApiGetEmbedPresetsAlt1Request) Execute() ([]Preset, *http.Response, error) {
	return r.ApiService.GetEmbedPresetsAlt1Execute(r)
}

/*
GetEmbedPresetsAlt1 Get all the embed presets that a user has created

This method returns every embed preset that belongs to the authenticated user.

 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @return ApiGetEmbedPresetsAlt1Request
*/
func (a *EmbedPresetsEssentialsAPIService) GetEmbedPresetsAlt1(ctx context.Context) ApiGetEmbedPresetsAlt1Request {
	return ApiGetEmbedPresetsAlt1Request{
		ApiService: a,
		ctx: ctx,
	}
}

// Execute executes the request
//  @return []Preset
func (a *EmbedPresetsEssentialsAPIService) GetEmbedPresetsAlt1Execute(r ApiGetEmbedPresetsAlt1Request) ([]Preset, *http.Response, error) {
	var (
		localVarHTTPMethod   = http.MethodGet
		localVarPostBody     interface{}
		formFiles            []formFile
		localVarReturnValue  []Preset
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "EmbedPresetsEssentialsAPIService.GetEmbedPresetsAlt1")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/me/presets"

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
	localVarHTTPHeaderAccepts := []string{"application/vnd.vimeo.preset+json"}

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
