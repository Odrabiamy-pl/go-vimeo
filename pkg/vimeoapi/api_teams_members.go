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


type TeamsMembersAPI interface {

	/*
	GetTeamInformation Get membership information about a team

	This method returns information about the membership of the specified team. Usage is currently limited to the team join forms.

	@param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
	@param teamUserId The ID of the team user.
	@param userId The ID of the team owner.
	@return ApiGetTeamInformationRequest
	*/
	GetTeamInformation(ctx context.Context, teamUserId float32, userId float32) ApiGetTeamInformationRequest

	// GetTeamInformationExecute executes the request
	GetTeamInformationExecute(r ApiGetTeamInformationRequest) (*http.Response, error)

	/*
	GetTeamRoleInformation Get information about the user's role on a team

	This method returns information about the authenticated user's role on the specified team owner's team.

	@param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
	@param userId The ID of the team owner.
	@return ApiGetTeamRoleInformationRequest
	*/
	GetTeamRoleInformation(ctx context.Context, userId float32) ApiGetTeamRoleInformationRequest

	// GetTeamRoleInformationExecute executes the request
	GetTeamRoleInformationExecute(r ApiGetTeamRoleInformationRequest) (*http.Response, error)

	/*
	GetTeammembersInformation Get membership information about a team

	This method returns information about the membership of the specified team. Usage is currently limited to the team join forms.

	@param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
	@param code The code corresponding to the desired team. This value appears under `TeamUser` > `code`.
	@return ApiGetTeammembersInformationRequest
	*/
	GetTeammembersInformation(ctx context.Context, code string) ApiGetTeammembersInformationRequest

	// GetTeammembersInformationExecute executes the request
	GetTeammembersInformationExecute(r ApiGetTeammembersInformationRequest) (*http.Response, error)
}

// TeamsMembersAPIService TeamsMembersAPI service
type TeamsMembersAPIService service

type ApiGetTeamInformationRequest struct {
	ctx context.Context
	ApiService TeamsMembersAPI
	teamUserId float32
	userId float32
}

func (r ApiGetTeamInformationRequest) Execute() (*http.Response, error) {
	return r.ApiService.GetTeamInformationExecute(r)
}

/*
GetTeamInformation Get membership information about a team

This method returns information about the membership of the specified team. Usage is currently limited to the team join forms.

 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @param teamUserId The ID of the team user.
 @param userId The ID of the team owner.
 @return ApiGetTeamInformationRequest
*/
func (a *TeamsMembersAPIService) GetTeamInformation(ctx context.Context, teamUserId float32, userId float32) ApiGetTeamInformationRequest {
	return ApiGetTeamInformationRequest{
		ApiService: a,
		ctx: ctx,
		teamUserId: teamUserId,
		userId: userId,
	}
}

// Execute executes the request
func (a *TeamsMembersAPIService) GetTeamInformationExecute(r ApiGetTeamInformationRequest) (*http.Response, error) {
	var (
		localVarHTTPMethod   = http.MethodGet
		localVarPostBody     interface{}
		formFiles            []formFile
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "TeamsMembersAPIService.GetTeamInformation")
	if err != nil {
		return nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/users/{user_id}/team_users/{team_user_id}"
	localVarPath = strings.Replace(localVarPath, "{"+"team_user_id"+"}", url.PathEscape(parameterValueToString(r.teamUserId, "teamUserId")), -1)
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

type ApiGetTeamRoleInformationRequest struct {
	ctx context.Context
	ApiService TeamsMembersAPI
	userId float32
}

func (r ApiGetTeamRoleInformationRequest) Execute() (*http.Response, error) {
	return r.ApiService.GetTeamRoleInformationExecute(r)
}

/*
GetTeamRoleInformation Get information about the user's role on a team

This method returns information about the authenticated user's role on the specified team owner's team.

 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @param userId The ID of the team owner.
 @return ApiGetTeamRoleInformationRequest
*/
func (a *TeamsMembersAPIService) GetTeamRoleInformation(ctx context.Context, userId float32) ApiGetTeamRoleInformationRequest {
	return ApiGetTeamRoleInformationRequest{
		ApiService: a,
		ctx: ctx,
		userId: userId,
	}
}

// Execute executes the request
func (a *TeamsMembersAPIService) GetTeamRoleInformationExecute(r ApiGetTeamRoleInformationRequest) (*http.Response, error) {
	var (
		localVarHTTPMethod   = http.MethodGet
		localVarPostBody     interface{}
		formFiles            []formFile
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "TeamsMembersAPIService.GetTeamRoleInformation")
	if err != nil {
		return nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/users/{user_id}/team/role"
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

type ApiGetTeammembersInformationRequest struct {
	ctx context.Context
	ApiService TeamsMembersAPI
	code string
}

func (r ApiGetTeammembersInformationRequest) Execute() (*http.Response, error) {
	return r.ApiService.GetTeammembersInformationExecute(r)
}

/*
GetTeammembersInformation Get membership information about a team

This method returns information about the membership of the specified team. Usage is currently limited to the team join forms.

 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @param code The code corresponding to the desired team. This value appears under `TeamUser` > `code`.
 @return ApiGetTeammembersInformationRequest
*/
func (a *TeamsMembersAPIService) GetTeammembersInformation(ctx context.Context, code string) ApiGetTeammembersInformationRequest {
	return ApiGetTeammembersInformationRequest{
		ApiService: a,
		ctx: ctx,
		code: code,
	}
}

// Execute executes the request
func (a *TeamsMembersAPIService) GetTeammembersInformationExecute(r ApiGetTeammembersInformationRequest) (*http.Response, error) {
	var (
		localVarHTTPMethod   = http.MethodGet
		localVarPostBody     interface{}
		formFiles            []formFile
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "TeamsMembersAPIService.GetTeammembersInformation")
	if err != nil {
		return nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/teammembers/{code}"
	localVarPath = strings.Replace(localVarPath, "{"+"code"+"}", url.PathEscape(parameterValueToString(r.code, "code")), -1)

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