/*
Vimeo API

Build something great. Vimeo's API supports flexible, high-quality video integration with your custom apps.

API version: 3.4
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package vimeoapi

import (
	"encoding/json"
	"bytes"
	"fmt"
)

// checks if the CreateLiveEventAlt1Request type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &CreateLiveEventAlt1Request{}

// CreateLiveEventAlt1Request struct for CreateLiveEventAlt1Request
type CreateLiveEventAlt1Request struct {
	// Whether the share link is usable.
	AllowShareLink *bool `json:"allow_share_link,omitempty"`
	// Whether automated closed captions are enabled for the event.
	AutoCcEnabled *bool `json:"auto_cc_enabled,omitempty"`
	// A comma-separated list of keywords that improve the quality of the automated closed captions.
	AutoCcKeywords *string `json:"auto_cc_keywords,omitempty"`
	// The language in which the automated closed captions appear.  Option descriptions:  * `de-DE` - The language is German.  * `en-US` - The language is English.  * `es-ES` - The language is Spanish.  * `fr-FR` - The language is French.  * `pt-BR` - The language is Portuguese. 
	AutoCcLang *string `json:"auto_cc_lang,omitempty"`
	// Whether the title for the next video in the event is generated based on the time of the stream and the **title** field of the event.
	AutomaticallyTitleStream *bool `json:"automatically_title_stream,omitempty"`
	// Whether to display the live chat client on the Vimeo event page.
	ChatEnabled *bool `json:"chat_enabled,omitempty"`
	// A list of values describing the content in this event. To return the list of all possible content rating values, use the [`/contentratings`](https://developer.vimeo.com/api/reference/videos#get_content_ratings) endpoint.
	ContentRating []string `json:"content_rating,omitempty"`
	// Whether the DVR feature is enabled.
	Dvr *bool `json:"dvr,omitempty"`
	Embed *CreateLiveEventAlt1RequestEmbed `json:"embed,omitempty"`
	// The URI of the event's folder.
	FolderUri *string `json:"folder_uri,omitempty"`
	InteractionToolsSettings *CreateLiveEventAlt1RequestInteractionToolsSettings `json:"interaction_tools_settings,omitempty"`
	// Whether the event has low-latency streaming enabled.
	LowLatency *bool `json:"low_latency,omitempty"`
	// The order in which the videos of the event appear within the event's playlist.  Option descriptions:  * `added_first` - The most recently added videos appear first.  * `added_last` - The most recently added videos appear last.  * `alphabetical` - The videos appear in alphabetical order.  * `arranged` - The videos appear in the order in which the user has arranged them.  * `comments` - The videos appear in order of number of comments.  * `duration` - The videos appear in order of duration.  * `likes` - The videos appear in order of number of likes.  * `newest` - The newest videos appear first.  * `oldest` - The oldest videos appear first.  * `plays` - The videos appear in order of number of plays. 
	PlaylistSort *string `json:"playlist_sort,omitempty"`
	// Whether the event has RTMP preview enabled.
	RtmpPreview *bool `json:"rtmp_preview,omitempty"`
	Schedule *CreateLiveEventAlt1RequestSchedule `json:"schedule,omitempty"`
	// Whether the scheduled playback feature is enabled.
	ScheduledPlayback *bool `json:"scheduled_playback,omitempty"`
	// The description of the next video to be streamed to the event.
	StreamDescription *string `json:"stream_description,omitempty"`
	StreamEmbed *CreateLiveEventAlt1RequestStreamEmbed `json:"stream_embed,omitempty"`
	// The password when **stream_privacy.view** is `password`. Anyone with the password can view the videos generated by streaming to the event.
	StreamPassword *string `json:"stream_password,omitempty"`
	StreamPrivacy *CreateLiveEventAlt1RequestStreamPrivacy `json:"stream_privacy,omitempty"`
	// The title of the next video to be streamed to the event. This parameter is required when **automatically_title_stream** is `false`.
	StreamTitle *string `json:"stream_title,omitempty"`
	// The time zone used in resolving the timestamps that are included in automatically generated video titles.
	TimeZone *string `json:"time_zone,omitempty"`
	// The title of the event. If **automatically_title_stream** is `true`, this value is the base title for videos created by streaming to this event.
	Title string `json:"title"`
}

type _CreateLiveEventAlt1Request CreateLiveEventAlt1Request

// NewCreateLiveEventAlt1Request instantiates a new CreateLiveEventAlt1Request object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewCreateLiveEventAlt1Request(title string) *CreateLiveEventAlt1Request {
	this := CreateLiveEventAlt1Request{}
	this.Title = title
	return &this
}

// NewCreateLiveEventAlt1RequestWithDefaults instantiates a new CreateLiveEventAlt1Request object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewCreateLiveEventAlt1RequestWithDefaults() *CreateLiveEventAlt1Request {
	this := CreateLiveEventAlt1Request{}
	return &this
}

// GetAllowShareLink returns the AllowShareLink field value if set, zero value otherwise.
func (o *CreateLiveEventAlt1Request) GetAllowShareLink() bool {
	if o == nil || IsNil(o.AllowShareLink) {
		var ret bool
		return ret
	}
	return *o.AllowShareLink
}

// GetAllowShareLinkOk returns a tuple with the AllowShareLink field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CreateLiveEventAlt1Request) GetAllowShareLinkOk() (*bool, bool) {
	if o == nil || IsNil(o.AllowShareLink) {
		return nil, false
	}
	return o.AllowShareLink, true
}

// HasAllowShareLink returns a boolean if a field has been set.
func (o *CreateLiveEventAlt1Request) HasAllowShareLink() bool {
	if o != nil && !IsNil(o.AllowShareLink) {
		return true
	}

	return false
}

// SetAllowShareLink gets a reference to the given bool and assigns it to the AllowShareLink field.
func (o *CreateLiveEventAlt1Request) SetAllowShareLink(v bool) {
	o.AllowShareLink = &v
}

// GetAutoCcEnabled returns the AutoCcEnabled field value if set, zero value otherwise.
func (o *CreateLiveEventAlt1Request) GetAutoCcEnabled() bool {
	if o == nil || IsNil(o.AutoCcEnabled) {
		var ret bool
		return ret
	}
	return *o.AutoCcEnabled
}

// GetAutoCcEnabledOk returns a tuple with the AutoCcEnabled field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CreateLiveEventAlt1Request) GetAutoCcEnabledOk() (*bool, bool) {
	if o == nil || IsNil(o.AutoCcEnabled) {
		return nil, false
	}
	return o.AutoCcEnabled, true
}

// HasAutoCcEnabled returns a boolean if a field has been set.
func (o *CreateLiveEventAlt1Request) HasAutoCcEnabled() bool {
	if o != nil && !IsNil(o.AutoCcEnabled) {
		return true
	}

	return false
}

// SetAutoCcEnabled gets a reference to the given bool and assigns it to the AutoCcEnabled field.
func (o *CreateLiveEventAlt1Request) SetAutoCcEnabled(v bool) {
	o.AutoCcEnabled = &v
}

// GetAutoCcKeywords returns the AutoCcKeywords field value if set, zero value otherwise.
func (o *CreateLiveEventAlt1Request) GetAutoCcKeywords() string {
	if o == nil || IsNil(o.AutoCcKeywords) {
		var ret string
		return ret
	}
	return *o.AutoCcKeywords
}

// GetAutoCcKeywordsOk returns a tuple with the AutoCcKeywords field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CreateLiveEventAlt1Request) GetAutoCcKeywordsOk() (*string, bool) {
	if o == nil || IsNil(o.AutoCcKeywords) {
		return nil, false
	}
	return o.AutoCcKeywords, true
}

// HasAutoCcKeywords returns a boolean if a field has been set.
func (o *CreateLiveEventAlt1Request) HasAutoCcKeywords() bool {
	if o != nil && !IsNil(o.AutoCcKeywords) {
		return true
	}

	return false
}

// SetAutoCcKeywords gets a reference to the given string and assigns it to the AutoCcKeywords field.
func (o *CreateLiveEventAlt1Request) SetAutoCcKeywords(v string) {
	o.AutoCcKeywords = &v
}

// GetAutoCcLang returns the AutoCcLang field value if set, zero value otherwise.
func (o *CreateLiveEventAlt1Request) GetAutoCcLang() string {
	if o == nil || IsNil(o.AutoCcLang) {
		var ret string
		return ret
	}
	return *o.AutoCcLang
}

// GetAutoCcLangOk returns a tuple with the AutoCcLang field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CreateLiveEventAlt1Request) GetAutoCcLangOk() (*string, bool) {
	if o == nil || IsNil(o.AutoCcLang) {
		return nil, false
	}
	return o.AutoCcLang, true
}

// HasAutoCcLang returns a boolean if a field has been set.
func (o *CreateLiveEventAlt1Request) HasAutoCcLang() bool {
	if o != nil && !IsNil(o.AutoCcLang) {
		return true
	}

	return false
}

// SetAutoCcLang gets a reference to the given string and assigns it to the AutoCcLang field.
func (o *CreateLiveEventAlt1Request) SetAutoCcLang(v string) {
	o.AutoCcLang = &v
}

// GetAutomaticallyTitleStream returns the AutomaticallyTitleStream field value if set, zero value otherwise.
func (o *CreateLiveEventAlt1Request) GetAutomaticallyTitleStream() bool {
	if o == nil || IsNil(o.AutomaticallyTitleStream) {
		var ret bool
		return ret
	}
	return *o.AutomaticallyTitleStream
}

// GetAutomaticallyTitleStreamOk returns a tuple with the AutomaticallyTitleStream field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CreateLiveEventAlt1Request) GetAutomaticallyTitleStreamOk() (*bool, bool) {
	if o == nil || IsNil(o.AutomaticallyTitleStream) {
		return nil, false
	}
	return o.AutomaticallyTitleStream, true
}

// HasAutomaticallyTitleStream returns a boolean if a field has been set.
func (o *CreateLiveEventAlt1Request) HasAutomaticallyTitleStream() bool {
	if o != nil && !IsNil(o.AutomaticallyTitleStream) {
		return true
	}

	return false
}

// SetAutomaticallyTitleStream gets a reference to the given bool and assigns it to the AutomaticallyTitleStream field.
func (o *CreateLiveEventAlt1Request) SetAutomaticallyTitleStream(v bool) {
	o.AutomaticallyTitleStream = &v
}

// GetChatEnabled returns the ChatEnabled field value if set, zero value otherwise.
func (o *CreateLiveEventAlt1Request) GetChatEnabled() bool {
	if o == nil || IsNil(o.ChatEnabled) {
		var ret bool
		return ret
	}
	return *o.ChatEnabled
}

// GetChatEnabledOk returns a tuple with the ChatEnabled field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CreateLiveEventAlt1Request) GetChatEnabledOk() (*bool, bool) {
	if o == nil || IsNil(o.ChatEnabled) {
		return nil, false
	}
	return o.ChatEnabled, true
}

// HasChatEnabled returns a boolean if a field has been set.
func (o *CreateLiveEventAlt1Request) HasChatEnabled() bool {
	if o != nil && !IsNil(o.ChatEnabled) {
		return true
	}

	return false
}

// SetChatEnabled gets a reference to the given bool and assigns it to the ChatEnabled field.
func (o *CreateLiveEventAlt1Request) SetChatEnabled(v bool) {
	o.ChatEnabled = &v
}

// GetContentRating returns the ContentRating field value if set, zero value otherwise.
func (o *CreateLiveEventAlt1Request) GetContentRating() []string {
	if o == nil || IsNil(o.ContentRating) {
		var ret []string
		return ret
	}
	return o.ContentRating
}

// GetContentRatingOk returns a tuple with the ContentRating field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CreateLiveEventAlt1Request) GetContentRatingOk() ([]string, bool) {
	if o == nil || IsNil(o.ContentRating) {
		return nil, false
	}
	return o.ContentRating, true
}

// HasContentRating returns a boolean if a field has been set.
func (o *CreateLiveEventAlt1Request) HasContentRating() bool {
	if o != nil && !IsNil(o.ContentRating) {
		return true
	}

	return false
}

// SetContentRating gets a reference to the given []string and assigns it to the ContentRating field.
func (o *CreateLiveEventAlt1Request) SetContentRating(v []string) {
	o.ContentRating = v
}

// GetDvr returns the Dvr field value if set, zero value otherwise.
func (o *CreateLiveEventAlt1Request) GetDvr() bool {
	if o == nil || IsNil(o.Dvr) {
		var ret bool
		return ret
	}
	return *o.Dvr
}

// GetDvrOk returns a tuple with the Dvr field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CreateLiveEventAlt1Request) GetDvrOk() (*bool, bool) {
	if o == nil || IsNil(o.Dvr) {
		return nil, false
	}
	return o.Dvr, true
}

// HasDvr returns a boolean if a field has been set.
func (o *CreateLiveEventAlt1Request) HasDvr() bool {
	if o != nil && !IsNil(o.Dvr) {
		return true
	}

	return false
}

// SetDvr gets a reference to the given bool and assigns it to the Dvr field.
func (o *CreateLiveEventAlt1Request) SetDvr(v bool) {
	o.Dvr = &v
}

// GetEmbed returns the Embed field value if set, zero value otherwise.
func (o *CreateLiveEventAlt1Request) GetEmbed() CreateLiveEventAlt1RequestEmbed {
	if o == nil || IsNil(o.Embed) {
		var ret CreateLiveEventAlt1RequestEmbed
		return ret
	}
	return *o.Embed
}

// GetEmbedOk returns a tuple with the Embed field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CreateLiveEventAlt1Request) GetEmbedOk() (*CreateLiveEventAlt1RequestEmbed, bool) {
	if o == nil || IsNil(o.Embed) {
		return nil, false
	}
	return o.Embed, true
}

// HasEmbed returns a boolean if a field has been set.
func (o *CreateLiveEventAlt1Request) HasEmbed() bool {
	if o != nil && !IsNil(o.Embed) {
		return true
	}

	return false
}

// SetEmbed gets a reference to the given CreateLiveEventAlt1RequestEmbed and assigns it to the Embed field.
func (o *CreateLiveEventAlt1Request) SetEmbed(v CreateLiveEventAlt1RequestEmbed) {
	o.Embed = &v
}

// GetFolderUri returns the FolderUri field value if set, zero value otherwise.
func (o *CreateLiveEventAlt1Request) GetFolderUri() string {
	if o == nil || IsNil(o.FolderUri) {
		var ret string
		return ret
	}
	return *o.FolderUri
}

// GetFolderUriOk returns a tuple with the FolderUri field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CreateLiveEventAlt1Request) GetFolderUriOk() (*string, bool) {
	if o == nil || IsNil(o.FolderUri) {
		return nil, false
	}
	return o.FolderUri, true
}

// HasFolderUri returns a boolean if a field has been set.
func (o *CreateLiveEventAlt1Request) HasFolderUri() bool {
	if o != nil && !IsNil(o.FolderUri) {
		return true
	}

	return false
}

// SetFolderUri gets a reference to the given string and assigns it to the FolderUri field.
func (o *CreateLiveEventAlt1Request) SetFolderUri(v string) {
	o.FolderUri = &v
}

// GetInteractionToolsSettings returns the InteractionToolsSettings field value if set, zero value otherwise.
func (o *CreateLiveEventAlt1Request) GetInteractionToolsSettings() CreateLiveEventAlt1RequestInteractionToolsSettings {
	if o == nil || IsNil(o.InteractionToolsSettings) {
		var ret CreateLiveEventAlt1RequestInteractionToolsSettings
		return ret
	}
	return *o.InteractionToolsSettings
}

// GetInteractionToolsSettingsOk returns a tuple with the InteractionToolsSettings field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CreateLiveEventAlt1Request) GetInteractionToolsSettingsOk() (*CreateLiveEventAlt1RequestInteractionToolsSettings, bool) {
	if o == nil || IsNil(o.InteractionToolsSettings) {
		return nil, false
	}
	return o.InteractionToolsSettings, true
}

// HasInteractionToolsSettings returns a boolean if a field has been set.
func (o *CreateLiveEventAlt1Request) HasInteractionToolsSettings() bool {
	if o != nil && !IsNil(o.InteractionToolsSettings) {
		return true
	}

	return false
}

// SetInteractionToolsSettings gets a reference to the given CreateLiveEventAlt1RequestInteractionToolsSettings and assigns it to the InteractionToolsSettings field.
func (o *CreateLiveEventAlt1Request) SetInteractionToolsSettings(v CreateLiveEventAlt1RequestInteractionToolsSettings) {
	o.InteractionToolsSettings = &v
}

// GetLowLatency returns the LowLatency field value if set, zero value otherwise.
func (o *CreateLiveEventAlt1Request) GetLowLatency() bool {
	if o == nil || IsNil(o.LowLatency) {
		var ret bool
		return ret
	}
	return *o.LowLatency
}

// GetLowLatencyOk returns a tuple with the LowLatency field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CreateLiveEventAlt1Request) GetLowLatencyOk() (*bool, bool) {
	if o == nil || IsNil(o.LowLatency) {
		return nil, false
	}
	return o.LowLatency, true
}

// HasLowLatency returns a boolean if a field has been set.
func (o *CreateLiveEventAlt1Request) HasLowLatency() bool {
	if o != nil && !IsNil(o.LowLatency) {
		return true
	}

	return false
}

// SetLowLatency gets a reference to the given bool and assigns it to the LowLatency field.
func (o *CreateLiveEventAlt1Request) SetLowLatency(v bool) {
	o.LowLatency = &v
}

// GetPlaylistSort returns the PlaylistSort field value if set, zero value otherwise.
func (o *CreateLiveEventAlt1Request) GetPlaylistSort() string {
	if o == nil || IsNil(o.PlaylistSort) {
		var ret string
		return ret
	}
	return *o.PlaylistSort
}

// GetPlaylistSortOk returns a tuple with the PlaylistSort field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CreateLiveEventAlt1Request) GetPlaylistSortOk() (*string, bool) {
	if o == nil || IsNil(o.PlaylistSort) {
		return nil, false
	}
	return o.PlaylistSort, true
}

// HasPlaylistSort returns a boolean if a field has been set.
func (o *CreateLiveEventAlt1Request) HasPlaylistSort() bool {
	if o != nil && !IsNil(o.PlaylistSort) {
		return true
	}

	return false
}

// SetPlaylistSort gets a reference to the given string and assigns it to the PlaylistSort field.
func (o *CreateLiveEventAlt1Request) SetPlaylistSort(v string) {
	o.PlaylistSort = &v
}

// GetRtmpPreview returns the RtmpPreview field value if set, zero value otherwise.
func (o *CreateLiveEventAlt1Request) GetRtmpPreview() bool {
	if o == nil || IsNil(o.RtmpPreview) {
		var ret bool
		return ret
	}
	return *o.RtmpPreview
}

// GetRtmpPreviewOk returns a tuple with the RtmpPreview field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CreateLiveEventAlt1Request) GetRtmpPreviewOk() (*bool, bool) {
	if o == nil || IsNil(o.RtmpPreview) {
		return nil, false
	}
	return o.RtmpPreview, true
}

// HasRtmpPreview returns a boolean if a field has been set.
func (o *CreateLiveEventAlt1Request) HasRtmpPreview() bool {
	if o != nil && !IsNil(o.RtmpPreview) {
		return true
	}

	return false
}

// SetRtmpPreview gets a reference to the given bool and assigns it to the RtmpPreview field.
func (o *CreateLiveEventAlt1Request) SetRtmpPreview(v bool) {
	o.RtmpPreview = &v
}

// GetSchedule returns the Schedule field value if set, zero value otherwise.
func (o *CreateLiveEventAlt1Request) GetSchedule() CreateLiveEventAlt1RequestSchedule {
	if o == nil || IsNil(o.Schedule) {
		var ret CreateLiveEventAlt1RequestSchedule
		return ret
	}
	return *o.Schedule
}

// GetScheduleOk returns a tuple with the Schedule field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CreateLiveEventAlt1Request) GetScheduleOk() (*CreateLiveEventAlt1RequestSchedule, bool) {
	if o == nil || IsNil(o.Schedule) {
		return nil, false
	}
	return o.Schedule, true
}

// HasSchedule returns a boolean if a field has been set.
func (o *CreateLiveEventAlt1Request) HasSchedule() bool {
	if o != nil && !IsNil(o.Schedule) {
		return true
	}

	return false
}

// SetSchedule gets a reference to the given CreateLiveEventAlt1RequestSchedule and assigns it to the Schedule field.
func (o *CreateLiveEventAlt1Request) SetSchedule(v CreateLiveEventAlt1RequestSchedule) {
	o.Schedule = &v
}

// GetScheduledPlayback returns the ScheduledPlayback field value if set, zero value otherwise.
func (o *CreateLiveEventAlt1Request) GetScheduledPlayback() bool {
	if o == nil || IsNil(o.ScheduledPlayback) {
		var ret bool
		return ret
	}
	return *o.ScheduledPlayback
}

// GetScheduledPlaybackOk returns a tuple with the ScheduledPlayback field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CreateLiveEventAlt1Request) GetScheduledPlaybackOk() (*bool, bool) {
	if o == nil || IsNil(o.ScheduledPlayback) {
		return nil, false
	}
	return o.ScheduledPlayback, true
}

// HasScheduledPlayback returns a boolean if a field has been set.
func (o *CreateLiveEventAlt1Request) HasScheduledPlayback() bool {
	if o != nil && !IsNil(o.ScheduledPlayback) {
		return true
	}

	return false
}

// SetScheduledPlayback gets a reference to the given bool and assigns it to the ScheduledPlayback field.
func (o *CreateLiveEventAlt1Request) SetScheduledPlayback(v bool) {
	o.ScheduledPlayback = &v
}

// GetStreamDescription returns the StreamDescription field value if set, zero value otherwise.
func (o *CreateLiveEventAlt1Request) GetStreamDescription() string {
	if o == nil || IsNil(o.StreamDescription) {
		var ret string
		return ret
	}
	return *o.StreamDescription
}

// GetStreamDescriptionOk returns a tuple with the StreamDescription field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CreateLiveEventAlt1Request) GetStreamDescriptionOk() (*string, bool) {
	if o == nil || IsNil(o.StreamDescription) {
		return nil, false
	}
	return o.StreamDescription, true
}

// HasStreamDescription returns a boolean if a field has been set.
func (o *CreateLiveEventAlt1Request) HasStreamDescription() bool {
	if o != nil && !IsNil(o.StreamDescription) {
		return true
	}

	return false
}

// SetStreamDescription gets a reference to the given string and assigns it to the StreamDescription field.
func (o *CreateLiveEventAlt1Request) SetStreamDescription(v string) {
	o.StreamDescription = &v
}

// GetStreamEmbed returns the StreamEmbed field value if set, zero value otherwise.
func (o *CreateLiveEventAlt1Request) GetStreamEmbed() CreateLiveEventAlt1RequestStreamEmbed {
	if o == nil || IsNil(o.StreamEmbed) {
		var ret CreateLiveEventAlt1RequestStreamEmbed
		return ret
	}
	return *o.StreamEmbed
}

// GetStreamEmbedOk returns a tuple with the StreamEmbed field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CreateLiveEventAlt1Request) GetStreamEmbedOk() (*CreateLiveEventAlt1RequestStreamEmbed, bool) {
	if o == nil || IsNil(o.StreamEmbed) {
		return nil, false
	}
	return o.StreamEmbed, true
}

// HasStreamEmbed returns a boolean if a field has been set.
func (o *CreateLiveEventAlt1Request) HasStreamEmbed() bool {
	if o != nil && !IsNil(o.StreamEmbed) {
		return true
	}

	return false
}

// SetStreamEmbed gets a reference to the given CreateLiveEventAlt1RequestStreamEmbed and assigns it to the StreamEmbed field.
func (o *CreateLiveEventAlt1Request) SetStreamEmbed(v CreateLiveEventAlt1RequestStreamEmbed) {
	o.StreamEmbed = &v
}

// GetStreamPassword returns the StreamPassword field value if set, zero value otherwise.
func (o *CreateLiveEventAlt1Request) GetStreamPassword() string {
	if o == nil || IsNil(o.StreamPassword) {
		var ret string
		return ret
	}
	return *o.StreamPassword
}

// GetStreamPasswordOk returns a tuple with the StreamPassword field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CreateLiveEventAlt1Request) GetStreamPasswordOk() (*string, bool) {
	if o == nil || IsNil(o.StreamPassword) {
		return nil, false
	}
	return o.StreamPassword, true
}

// HasStreamPassword returns a boolean if a field has been set.
func (o *CreateLiveEventAlt1Request) HasStreamPassword() bool {
	if o != nil && !IsNil(o.StreamPassword) {
		return true
	}

	return false
}

// SetStreamPassword gets a reference to the given string and assigns it to the StreamPassword field.
func (o *CreateLiveEventAlt1Request) SetStreamPassword(v string) {
	o.StreamPassword = &v
}

// GetStreamPrivacy returns the StreamPrivacy field value if set, zero value otherwise.
func (o *CreateLiveEventAlt1Request) GetStreamPrivacy() CreateLiveEventAlt1RequestStreamPrivacy {
	if o == nil || IsNil(o.StreamPrivacy) {
		var ret CreateLiveEventAlt1RequestStreamPrivacy
		return ret
	}
	return *o.StreamPrivacy
}

// GetStreamPrivacyOk returns a tuple with the StreamPrivacy field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CreateLiveEventAlt1Request) GetStreamPrivacyOk() (*CreateLiveEventAlt1RequestStreamPrivacy, bool) {
	if o == nil || IsNil(o.StreamPrivacy) {
		return nil, false
	}
	return o.StreamPrivacy, true
}

// HasStreamPrivacy returns a boolean if a field has been set.
func (o *CreateLiveEventAlt1Request) HasStreamPrivacy() bool {
	if o != nil && !IsNil(o.StreamPrivacy) {
		return true
	}

	return false
}

// SetStreamPrivacy gets a reference to the given CreateLiveEventAlt1RequestStreamPrivacy and assigns it to the StreamPrivacy field.
func (o *CreateLiveEventAlt1Request) SetStreamPrivacy(v CreateLiveEventAlt1RequestStreamPrivacy) {
	o.StreamPrivacy = &v
}

// GetStreamTitle returns the StreamTitle field value if set, zero value otherwise.
func (o *CreateLiveEventAlt1Request) GetStreamTitle() string {
	if o == nil || IsNil(o.StreamTitle) {
		var ret string
		return ret
	}
	return *o.StreamTitle
}

// GetStreamTitleOk returns a tuple with the StreamTitle field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CreateLiveEventAlt1Request) GetStreamTitleOk() (*string, bool) {
	if o == nil || IsNil(o.StreamTitle) {
		return nil, false
	}
	return o.StreamTitle, true
}

// HasStreamTitle returns a boolean if a field has been set.
func (o *CreateLiveEventAlt1Request) HasStreamTitle() bool {
	if o != nil && !IsNil(o.StreamTitle) {
		return true
	}

	return false
}

// SetStreamTitle gets a reference to the given string and assigns it to the StreamTitle field.
func (o *CreateLiveEventAlt1Request) SetStreamTitle(v string) {
	o.StreamTitle = &v
}

// GetTimeZone returns the TimeZone field value if set, zero value otherwise.
func (o *CreateLiveEventAlt1Request) GetTimeZone() string {
	if o == nil || IsNil(o.TimeZone) {
		var ret string
		return ret
	}
	return *o.TimeZone
}

// GetTimeZoneOk returns a tuple with the TimeZone field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CreateLiveEventAlt1Request) GetTimeZoneOk() (*string, bool) {
	if o == nil || IsNil(o.TimeZone) {
		return nil, false
	}
	return o.TimeZone, true
}

// HasTimeZone returns a boolean if a field has been set.
func (o *CreateLiveEventAlt1Request) HasTimeZone() bool {
	if o != nil && !IsNil(o.TimeZone) {
		return true
	}

	return false
}

// SetTimeZone gets a reference to the given string and assigns it to the TimeZone field.
func (o *CreateLiveEventAlt1Request) SetTimeZone(v string) {
	o.TimeZone = &v
}

// GetTitle returns the Title field value
func (o *CreateLiveEventAlt1Request) GetTitle() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Title
}

// GetTitleOk returns a tuple with the Title field value
// and a boolean to check if the value has been set.
func (o *CreateLiveEventAlt1Request) GetTitleOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Title, true
}

// SetTitle sets field value
func (o *CreateLiveEventAlt1Request) SetTitle(v string) {
	o.Title = v
}

func (o CreateLiveEventAlt1Request) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o CreateLiveEventAlt1Request) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.AllowShareLink) {
		toSerialize["allow_share_link"] = o.AllowShareLink
	}
	if !IsNil(o.AutoCcEnabled) {
		toSerialize["auto_cc_enabled"] = o.AutoCcEnabled
	}
	if !IsNil(o.AutoCcKeywords) {
		toSerialize["auto_cc_keywords"] = o.AutoCcKeywords
	}
	if !IsNil(o.AutoCcLang) {
		toSerialize["auto_cc_lang"] = o.AutoCcLang
	}
	if !IsNil(o.AutomaticallyTitleStream) {
		toSerialize["automatically_title_stream"] = o.AutomaticallyTitleStream
	}
	if !IsNil(o.ChatEnabled) {
		toSerialize["chat_enabled"] = o.ChatEnabled
	}
	if !IsNil(o.ContentRating) {
		toSerialize["content_rating"] = o.ContentRating
	}
	if !IsNil(o.Dvr) {
		toSerialize["dvr"] = o.Dvr
	}
	if !IsNil(o.Embed) {
		toSerialize["embed"] = o.Embed
	}
	if !IsNil(o.FolderUri) {
		toSerialize["folder_uri"] = o.FolderUri
	}
	if !IsNil(o.InteractionToolsSettings) {
		toSerialize["interaction_tools_settings"] = o.InteractionToolsSettings
	}
	if !IsNil(o.LowLatency) {
		toSerialize["low_latency"] = o.LowLatency
	}
	if !IsNil(o.PlaylistSort) {
		toSerialize["playlist_sort"] = o.PlaylistSort
	}
	if !IsNil(o.RtmpPreview) {
		toSerialize["rtmp_preview"] = o.RtmpPreview
	}
	if !IsNil(o.Schedule) {
		toSerialize["schedule"] = o.Schedule
	}
	if !IsNil(o.ScheduledPlayback) {
		toSerialize["scheduled_playback"] = o.ScheduledPlayback
	}
	if !IsNil(o.StreamDescription) {
		toSerialize["stream_description"] = o.StreamDescription
	}
	if !IsNil(o.StreamEmbed) {
		toSerialize["stream_embed"] = o.StreamEmbed
	}
	if !IsNil(o.StreamPassword) {
		toSerialize["stream_password"] = o.StreamPassword
	}
	if !IsNil(o.StreamPrivacy) {
		toSerialize["stream_privacy"] = o.StreamPrivacy
	}
	if !IsNil(o.StreamTitle) {
		toSerialize["stream_title"] = o.StreamTitle
	}
	if !IsNil(o.TimeZone) {
		toSerialize["time_zone"] = o.TimeZone
	}
	toSerialize["title"] = o.Title
	return toSerialize, nil
}

func (o *CreateLiveEventAlt1Request) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"title",
	}

	allProperties := make(map[string]interface{})

	err = json.Unmarshal(data, &allProperties)

	if err != nil {
		return err;
	}

	for _, requiredProperty := range(requiredProperties) {
		if _, exists := allProperties[requiredProperty]; !exists {
			return fmt.Errorf("no value given for required property %v", requiredProperty)
		}
	}

	varCreateLiveEventAlt1Request := _CreateLiveEventAlt1Request{}

	decoder := json.NewDecoder(bytes.NewReader(data))
	decoder.DisallowUnknownFields()
	err = decoder.Decode(&varCreateLiveEventAlt1Request)

	if err != nil {
		return err
	}

	*o = CreateLiveEventAlt1Request(varCreateLiveEventAlt1Request)

	return err
}

type NullableCreateLiveEventAlt1Request struct {
	value *CreateLiveEventAlt1Request
	isSet bool
}

func (v NullableCreateLiveEventAlt1Request) Get() *CreateLiveEventAlt1Request {
	return v.value
}

func (v *NullableCreateLiveEventAlt1Request) Set(val *CreateLiveEventAlt1Request) {
	v.value = val
	v.isSet = true
}

func (v NullableCreateLiveEventAlt1Request) IsSet() bool {
	return v.isSet
}

func (v *NullableCreateLiveEventAlt1Request) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableCreateLiveEventAlt1Request(val *CreateLiveEventAlt1Request) *NullableCreateLiveEventAlt1Request {
	return &NullableCreateLiveEventAlt1Request{value: val, isSet: true}
}

func (v NullableCreateLiveEventAlt1Request) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableCreateLiveEventAlt1Request) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

