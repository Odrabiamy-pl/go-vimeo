/*
Vimeo API

Build something great. Vimeo's API supports flexible, high-quality video integration with your custom apps.

API version: 3.4
*/

package vimeoapi

import (
	"bytes"
	"context"
	"io"
	"net/http"
	"net/url"
	"reflect"
	"strings"
)

type VideosEssentialsAPI interface {
	/*
		GetVideo Get a specific video

		This method returns a single video.

		@param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
		@param videoId The ID of the video.
		@return ApiGetVideoRequest
	*/
	GetVideo(ctx context.Context, videoId int32) ApiGetVideoRequest

	// GetVideoExecute executes the request
	//  @return Video
	GetVideoExecute(r ApiGetVideoRequest) (*Video, *http.Response, error)

	/*
		GetVideos Get all the videos that the user has uploaded

		This method returns all the videos that the authenticated user has uploaded.

		@param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
		@param userId The ID of the user.
		@return ApiGetVideosRequest
	*/
	GetVideos(ctx context.Context, userId int32) ApiGetVideosRequest

	// GetVideosExecute executes the request
	//  @return []Video
	GetVideosExecute(r ApiGetVideosRequest) ([]Video, *http.Response, error)
}

// VideosEssentialsAPIService VideosEssentialsAPI service
type VideosEssentialsAPIService service

type ApiGetVideoRequest struct {
	ctx        context.Context
	ApiService VideosEssentialsAPI
	videoId    int32
}

func (r ApiGetVideoRequest) Execute() (*Video, *http.Response, error) {
	return r.ApiService.GetVideoExecute(r)
}

/*
GetVideo Get a specific video

This method returns a single video.

	@param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
	@param videoId The ID of the video.
	@return ApiGetVideoRequest
*/
func (a *VideosEssentialsAPIService) GetVideo(ctx context.Context, videoId int32) ApiGetVideoRequest {
	return ApiGetVideoRequest{
		ApiService: a,
		ctx:        ctx,
		videoId:    videoId,
	}
}

// Execute executes the request
//
//	@return Video
func (a *VideosEssentialsAPIService) GetVideoExecute(r ApiGetVideoRequest) (*Video, *http.Response, error) {
	var (
		localVarHTTPMethod  = http.MethodGet
		localVarPostBody    interface{}
		formFiles           []formFile
		localVarReturnValue *Video
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "VideosEssentialsAPIService.GetVideo")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/videos/{video_id}"
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
	localVarHTTPHeaderAccepts := []string{"application/vnd.vimeo.video+json"}

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

type ApiGetVideosRequest struct {
	ctx                  context.Context
	ApiService           VideosEssentialsAPI
	userId               int32
	containingUri        *string
	direction            *string
	filter               *string
	filterEmbeddable     *bool
	filterPlayable       *bool
	filterScreenRecorded *bool
	filterTag            *string
	filterTagAllOf       *string
	filterTagExclude     *string
	includeTeamContent   *string
	page                 *float32
	perPage              *float32
	query                *string
	queryFields          *[]string
	sort                 *string
}

// The page that contains the video URI. The field is available only when not paired with **query**.
func (r ApiGetVideosRequest) ContainingUri(containingUri string) ApiGetVideosRequest {
	r.containingUri = &containingUri
	return r
}

// The sort direction of the results.  Option descriptions:  * &#x60;asc&#x60; - Sort the results in ascending order.  * &#x60;desc&#x60; - Sort the results in descending order.
func (r ApiGetVideosRequest) Direction(direction string) ApiGetVideosRequest {
	r.direction = &direction
	return r
}

// The attribute by which to filter the results.  Option descriptions:  * &#x60;app_only&#x60; - Return app-only videos.  * &#x60;embeddable&#x60; - Return embeddable videos.  * &#x60;featured&#x60; - Return featured videos.  * &#x60;live&#x60; - Return only live videos.  * &#x60;no_placeholder&#x60; - Return no placeholder videos.  * &#x60;nolive&#x60; - Return no live videos.  * &#x60;playable&#x60; - Return playable videos.  * &#x60;screen_recorded&#x60; - Return screen-recorded videos.
func (r ApiGetVideosRequest) Filter(filter string) ApiGetVideosRequest {
	r.filter = &filter
	return r
}

// Whether to filter the results by embeddable videos (&#x60;true&#x60;) or non-embeddable videos (&#x60;false&#x60;). This parameter is required only when **filter** is &#x60;embeddable&#x60;.
func (r ApiGetVideosRequest) FilterEmbeddable(filterEmbeddable bool) ApiGetVideosRequest {
	r.filterEmbeddable = &filterEmbeddable
	return r
}

// Whether to filter the results by playable videos (&#x60;true&#x60;) or non-playable videos (&#x60;false&#x60;).
func (r ApiGetVideosRequest) FilterPlayable(filterPlayable bool) ApiGetVideosRequest {
	r.filterPlayable = &filterPlayable
	return r
}

// Whether to filter the results by screen-recorded videos (&#x60;true&#x60;) or non-screen-recorded videos (&#x60;false&#x60;).
func (r ApiGetVideosRequest) FilterScreenRecorded(filterScreenRecorded bool) ApiGetVideosRequest {
	r.filterScreenRecorded = &filterScreenRecorded
	return r
}

// A comma-separated list of tags to filter on. All results must include at least one of these tags.
func (r ApiGetVideosRequest) FilterTag(filterTag string) ApiGetVideosRequest {
	r.filterTag = &filterTag
	return r
}

// A comma-separated list of tags to filter on. All results must include all of these tags.
func (r ApiGetVideosRequest) FilterTagAllOf(filterTagAllOf string) ApiGetVideosRequest {
	r.filterTagAllOf = &filterTagAllOf
	return r
}

// A comma-separated list of tags to exclude. All results must exclude all of these tags.
func (r ApiGetVideosRequest) FilterTagExclude(filterTagExclude string) ApiGetVideosRequest {
	r.filterTagExclude = &filterTagExclude
	return r
}

// Whether to include content from the user&#39;s teams when searching. _This field is deprecated._
func (r ApiGetVideosRequest) IncludeTeamContent(includeTeamContent string) ApiGetVideosRequest {
	r.includeTeamContent = &includeTeamContent
	return r
}

// The page number of the results to show.
func (r ApiGetVideosRequest) Page(page float32) ApiGetVideosRequest {
	r.page = &page
	return r
}

// The number of items to show on each page of results, up to a maximum of 100.
func (r ApiGetVideosRequest) PerPage(perPage float32) ApiGetVideosRequest {
	r.perPage = &perPage
	return r
}

// The search query to use to filter the results.
func (r ApiGetVideosRequest) Query(query string) ApiGetVideosRequest {
	r.query = &query
	return r
}

// A comma-separated list of fields to query over. The default value is &#x60;title,description,chapters,tags&#x60;.  Option descriptions:  * &#x60;chapters&#x60; - Query by chapter titles that have been added to the video.  * &#x60;description&#x60; - Query by the description of the video.  * &#x60;tags&#x60; - Query by tag names that have been added to the video.  * &#x60;title&#x60; - Query by the title of the video.
func (r ApiGetVideosRequest) QueryFields(queryFields []string) ApiGetVideosRequest {
	r.queryFields = &queryFields
	return r
}

// The way to sort the results.  Option descriptions:  * &#x60;alphabetical&#x60; - Sort the results alphabetically by title.  * &#x60;date&#x60; - Sort the results by date.  * &#x60;default&#x60; - Use the default sorting method.  * &#x60;duration&#x60; - Sort the results by duration.  * &#x60;last_user_action_event_date&#x60; - Sort the results by last user interaction. If a result hasn&#39;t had an interaction, the upload date is used instead.  * &#x60;likes&#x60; - Sort the results by number of likes. To use this option, **direction** must be &#x60;desc&#x60;.  * &#x60;modified_time&#x60; - Sort the results by last modification.  * &#x60;plays&#x60; - Sort the results by number of plays. To use this option, **direction** must be &#x60;desc&#x60;.
func (r ApiGetVideosRequest) Sort(sort string) ApiGetVideosRequest {
	r.sort = &sort
	return r
}

func (r ApiGetVideosRequest) Execute() ([]Video, *http.Response, error) {
	return r.ApiService.GetVideosExecute(r)
}

/*
GetVideos Get all the videos that the user has uploaded

This method returns all the videos that the authenticated user has uploaded.

	@param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
	@param userId The ID of the user.
	@return ApiGetVideosRequest
*/
func (a *VideosEssentialsAPIService) GetVideos(ctx context.Context, userId int32) ApiGetVideosRequest {
	return ApiGetVideosRequest{
		ApiService: a,
		ctx:        ctx,
		userId:     userId,
	}
}

// Execute executes the request
//
//	@return []Video
func (a *VideosEssentialsAPIService) GetVideosExecute(r ApiGetVideosRequest) ([]Video, *http.Response, error) {
	var (
		localVarHTTPMethod  = http.MethodGet
		localVarPostBody    interface{}
		formFiles           []formFile
		localVarReturnValue []Video
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "VideosEssentialsAPIService.GetVideos")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/users/{user_id}/videos"
	localVarPath = strings.Replace(localVarPath, "{"+"user_id"+"}", url.PathEscape(parameterValueToString(r.userId, "userId")), -1)

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}

	if r.containingUri != nil {
		parameterAddToHeaderOrQuery(localVarQueryParams, "containing_uri", r.containingUri, "")
	}
	if r.direction != nil {
		parameterAddToHeaderOrQuery(localVarQueryParams, "direction", r.direction, "")
	}
	if r.filter != nil {
		parameterAddToHeaderOrQuery(localVarQueryParams, "filter", r.filter, "")
	}
	if r.filterEmbeddable != nil {
		parameterAddToHeaderOrQuery(localVarQueryParams, "filter_embeddable", r.filterEmbeddable, "")
	}
	if r.filterPlayable != nil {
		parameterAddToHeaderOrQuery(localVarQueryParams, "filter_playable", r.filterPlayable, "")
	}
	if r.filterScreenRecorded != nil {
		parameterAddToHeaderOrQuery(localVarQueryParams, "filter_screen_recorded", r.filterScreenRecorded, "")
	}
	if r.filterTag != nil {
		parameterAddToHeaderOrQuery(localVarQueryParams, "filter_tag", r.filterTag, "")
	}
	if r.filterTagAllOf != nil {
		parameterAddToHeaderOrQuery(localVarQueryParams, "filter_tag_all_of", r.filterTagAllOf, "")
	}
	if r.filterTagExclude != nil {
		parameterAddToHeaderOrQuery(localVarQueryParams, "filter_tag_exclude", r.filterTagExclude, "")
	}
	if r.includeTeamContent != nil {
		parameterAddToHeaderOrQuery(localVarQueryParams, "include_team_content", r.includeTeamContent, "")
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
	if r.queryFields != nil {
		t := *r.queryFields
		if reflect.TypeOf(t).Kind() == reflect.Slice {
			s := reflect.ValueOf(t)
			for i := 0; i < s.Len(); i++ {
				parameterAddToHeaderOrQuery(localVarQueryParams, "query_fields", s.Index(i).Interface(), "multi")
			}
		} else {
			parameterAddToHeaderOrQuery(localVarQueryParams, "query_fields", t, "multi")
		}
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
	localVarHTTPHeaderAccepts := []string{"application/vnd.vimeo.video+json"}

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
