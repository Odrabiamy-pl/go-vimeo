/*
Vimeo API

Build something great. Vimeo's API supports flexible, high-quality video integration with your custom apps.

API version: 3.4
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package vimeoapi

import (
	"encoding/json"
)

// checks if the VideoMetadata type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &VideoMetadata{}

// VideoMetadata The video's metadata.
type VideoMetadata struct {
	// Whether the video can be replaced.
	CanBeReplaced *bool                    `json:"can_be_replaced,omitempty"`
	Connections   VideoMetadataConnections `json:"connections"`
	// Whether the video has chapter suggestions.
	HasChapterSuggestions bool `json:"has_chapter_suggestions"`
	// Whether the video has the email capture feature.
	HasEmailCapture *bool                     `json:"has_email_capture,omitempty"`
	Interactions    VideoMetadataInteractions `json:"interactions"`
	// Whether the video is a screen recording.
	IsScreenRecord bool `json:"is_screen_record"`
	// Whether the video is a Vimeo Create video.
	IsVimeoCreate bool `json:"is_vimeo_create"`
	// Whether the video is a webinar.
	IsWebinar bool `json:"is_webinar"`
	// Whether the video is a Zoom upload.
	IsZoomUpload bool `json:"is_zoom_upload"`
}

// NewVideoMetadata instantiates a new VideoMetadata object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewVideoMetadata(connections VideoMetadataConnections, hasChapterSuggestions bool, interactions VideoMetadataInteractions, isScreenRecord bool, isVimeoCreate bool, isWebinar bool, isZoomUpload bool) *VideoMetadata {
	this := VideoMetadata{}
	this.Connections = connections
	this.HasChapterSuggestions = hasChapterSuggestions
	this.Interactions = interactions
	this.IsScreenRecord = isScreenRecord
	this.IsVimeoCreate = isVimeoCreate
	this.IsWebinar = isWebinar
	this.IsZoomUpload = isZoomUpload
	return &this
}

// NewVideoMetadataWithDefaults instantiates a new VideoMetadata object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewVideoMetadataWithDefaults() *VideoMetadata {
	this := VideoMetadata{}
	return &this
}

// GetCanBeReplaced returns the CanBeReplaced field value if set, zero value otherwise.
func (o *VideoMetadata) GetCanBeReplaced() bool {
	if o == nil || IsNil(o.CanBeReplaced) {
		var ret bool
		return ret
	}
	return *o.CanBeReplaced
}

// GetCanBeReplacedOk returns a tuple with the CanBeReplaced field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *VideoMetadata) GetCanBeReplacedOk() (*bool, bool) {
	if o == nil || IsNil(o.CanBeReplaced) {
		return nil, false
	}
	return o.CanBeReplaced, true
}

// HasCanBeReplaced returns a boolean if a field has been set.
func (o *VideoMetadata) HasCanBeReplaced() bool {
	if o != nil && !IsNil(o.CanBeReplaced) {
		return true
	}

	return false
}

// SetCanBeReplaced gets a reference to the given bool and assigns it to the CanBeReplaced field.
func (o *VideoMetadata) SetCanBeReplaced(v bool) {
	o.CanBeReplaced = &v
}

// GetConnections returns the Connections field value
func (o *VideoMetadata) GetConnections() VideoMetadataConnections {
	if o == nil {
		var ret VideoMetadataConnections
		return ret
	}

	return o.Connections
}

// GetConnectionsOk returns a tuple with the Connections field value
// and a boolean to check if the value has been set.
func (o *VideoMetadata) GetConnectionsOk() (*VideoMetadataConnections, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Connections, true
}

// SetConnections sets field value
func (o *VideoMetadata) SetConnections(v VideoMetadataConnections) {
	o.Connections = v
}

// GetHasChapterSuggestions returns the HasChapterSuggestions field value
func (o *VideoMetadata) GetHasChapterSuggestions() bool {
	if o == nil {
		var ret bool
		return ret
	}

	return o.HasChapterSuggestions
}

// GetHasChapterSuggestionsOk returns a tuple with the HasChapterSuggestions field value
// and a boolean to check if the value has been set.
func (o *VideoMetadata) GetHasChapterSuggestionsOk() (*bool, bool) {
	if o == nil {
		return nil, false
	}
	return &o.HasChapterSuggestions, true
}

// SetHasChapterSuggestions sets field value
func (o *VideoMetadata) SetHasChapterSuggestions(v bool) {
	o.HasChapterSuggestions = v
}

// GetHasEmailCapture returns the HasEmailCapture field value if set, zero value otherwise.
func (o *VideoMetadata) GetHasEmailCapture() bool {
	if o == nil || IsNil(o.HasEmailCapture) {
		var ret bool
		return ret
	}
	return *o.HasEmailCapture
}

// GetHasEmailCaptureOk returns a tuple with the HasEmailCapture field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *VideoMetadata) GetHasEmailCaptureOk() (*bool, bool) {
	if o == nil || IsNil(o.HasEmailCapture) {
		return nil, false
	}
	return o.HasEmailCapture, true
}

// HasHasEmailCapture returns a boolean if a field has been set.
func (o *VideoMetadata) HasHasEmailCapture() bool {
	if o != nil && !IsNil(o.HasEmailCapture) {
		return true
	}

	return false
}

// SetHasEmailCapture gets a reference to the given bool and assigns it to the HasEmailCapture field.
func (o *VideoMetadata) SetHasEmailCapture(v bool) {
	o.HasEmailCapture = &v
}

// GetInteractions returns the Interactions field value
func (o *VideoMetadata) GetInteractions() VideoMetadataInteractions {
	if o == nil {
		var ret VideoMetadataInteractions
		return ret
	}

	return o.Interactions
}

// GetInteractionsOk returns a tuple with the Interactions field value
// and a boolean to check if the value has been set.
func (o *VideoMetadata) GetInteractionsOk() (*VideoMetadataInteractions, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Interactions, true
}

// SetInteractions sets field value
func (o *VideoMetadata) SetInteractions(v VideoMetadataInteractions) {
	o.Interactions = v
}

// GetIsScreenRecord returns the IsScreenRecord field value
func (o *VideoMetadata) GetIsScreenRecord() bool {
	if o == nil {
		var ret bool
		return ret
	}

	return o.IsScreenRecord
}

// GetIsScreenRecordOk returns a tuple with the IsScreenRecord field value
// and a boolean to check if the value has been set.
func (o *VideoMetadata) GetIsScreenRecordOk() (*bool, bool) {
	if o == nil {
		return nil, false
	}
	return &o.IsScreenRecord, true
}

// SetIsScreenRecord sets field value
func (o *VideoMetadata) SetIsScreenRecord(v bool) {
	o.IsScreenRecord = v
}

// GetIsVimeoCreate returns the IsVimeoCreate field value
func (o *VideoMetadata) GetIsVimeoCreate() bool {
	if o == nil {
		var ret bool
		return ret
	}

	return o.IsVimeoCreate
}

// GetIsVimeoCreateOk returns a tuple with the IsVimeoCreate field value
// and a boolean to check if the value has been set.
func (o *VideoMetadata) GetIsVimeoCreateOk() (*bool, bool) {
	if o == nil {
		return nil, false
	}
	return &o.IsVimeoCreate, true
}

// SetIsVimeoCreate sets field value
func (o *VideoMetadata) SetIsVimeoCreate(v bool) {
	o.IsVimeoCreate = v
}

// GetIsWebinar returns the IsWebinar field value
func (o *VideoMetadata) GetIsWebinar() bool {
	if o == nil {
		var ret bool
		return ret
	}

	return o.IsWebinar
}

// GetIsWebinarOk returns a tuple with the IsWebinar field value
// and a boolean to check if the value has been set.
func (o *VideoMetadata) GetIsWebinarOk() (*bool, bool) {
	if o == nil {
		return nil, false
	}
	return &o.IsWebinar, true
}

// SetIsWebinar sets field value
func (o *VideoMetadata) SetIsWebinar(v bool) {
	o.IsWebinar = v
}

// GetIsZoomUpload returns the IsZoomUpload field value
func (o *VideoMetadata) GetIsZoomUpload() bool {
	if o == nil {
		var ret bool
		return ret
	}

	return o.IsZoomUpload
}

// GetIsZoomUploadOk returns a tuple with the IsZoomUpload field value
// and a boolean to check if the value has been set.
func (o *VideoMetadata) GetIsZoomUploadOk() (*bool, bool) {
	if o == nil {
		return nil, false
	}
	return &o.IsZoomUpload, true
}

// SetIsZoomUpload sets field value
func (o *VideoMetadata) SetIsZoomUpload(v bool) {
	o.IsZoomUpload = v
}

func (o VideoMetadata) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o VideoMetadata) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.CanBeReplaced) {
		toSerialize["can_be_replaced"] = o.CanBeReplaced
	}
	toSerialize["connections"] = o.Connections
	toSerialize["has_chapter_suggestions"] = o.HasChapterSuggestions
	if !IsNil(o.HasEmailCapture) {
		toSerialize["has_email_capture"] = o.HasEmailCapture
	}
	toSerialize["interactions"] = o.Interactions
	toSerialize["is_screen_record"] = o.IsScreenRecord
	toSerialize["is_vimeo_create"] = o.IsVimeoCreate
	toSerialize["is_webinar"] = o.IsWebinar
	toSerialize["is_zoom_upload"] = o.IsZoomUpload
	return toSerialize, nil
}

type NullableVideoMetadata struct {
	value *VideoMetadata
	isSet bool
}

func (v NullableVideoMetadata) Get() *VideoMetadata {
	return v.value
}

func (v *NullableVideoMetadata) Set(val *VideoMetadata) {
	v.value = val
	v.isSet = true
}

func (v NullableVideoMetadata) IsSet() bool {
	return v.isSet
}

func (v *NullableVideoMetadata) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableVideoMetadata(val *VideoMetadata) *NullableVideoMetadata {
	return &NullableVideoMetadata{value: val, isSet: true}
}

func (v NullableVideoMetadata) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableVideoMetadata) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
