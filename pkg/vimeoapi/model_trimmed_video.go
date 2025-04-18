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

// checks if the TrimmedVideo type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &TrimmedVideo{}

// TrimmedVideo struct for TrimmedVideo
type TrimmedVideo struct {
	// The ID of the video. _This field is deprecated._
	ClipId float32 `json:"clip_id"`
	// The time in ISO 8601 format when the trim was created.
	CreatedOn string `json:"created_on"`
	// The most recent version of the trimmed video. _This field is deprecated._
	CreatedVersionId string `json:"created_version_id"`
	// The end of the trim from the last trim, in seconds.
	End string `json:"end"`
	// Whether the transcoding jobs for the video file have finished. _This field is deprecated._
	IsClipFinishedTranscoding bool `json:"is_clip_finished_transcoding"`
	Metadata TrimmedVideoMetadata `json:"metadata"`
	// The time in ISO 8601 format when the trim policy was last modified.
	ModifiedOn NullableString `json:"modified_on"`
	// The video version that is the source of the trimmed video. _This field is deprecated._
	RootVersionId string `json:"root_version_id"`
	// The start of the trim from the last trim, in seconds.
	Start string `json:"start"`
	// The URI of the player or the trim service. _This field is deprecated._
	Uri string `json:"uri"`
	// The quality of the root version video file.
	VersionQuality string `json:"version_quality"`
}

// NewTrimmedVideo instantiates a new TrimmedVideo object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewTrimmedVideo(clipId float32, createdOn string, createdVersionId string, end string, isClipFinishedTranscoding bool, metadata TrimmedVideoMetadata, modifiedOn NullableString, rootVersionId string, start string, uri string, versionQuality string) *TrimmedVideo {
	this := TrimmedVideo{}
	this.ClipId = clipId
	this.CreatedOn = createdOn
	this.CreatedVersionId = createdVersionId
	this.End = end
	this.IsClipFinishedTranscoding = isClipFinishedTranscoding
	this.Metadata = metadata
	this.ModifiedOn = modifiedOn
	this.RootVersionId = rootVersionId
	this.Start = start
	this.Uri = uri
	this.VersionQuality = versionQuality
	return &this
}

// NewTrimmedVideoWithDefaults instantiates a new TrimmedVideo object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewTrimmedVideoWithDefaults() *TrimmedVideo {
	this := TrimmedVideo{}
	return &this
}

// GetClipId returns the ClipId field value
func (o *TrimmedVideo) GetClipId() float32 {
	if o == nil {
		var ret float32
		return ret
	}

	return o.ClipId
}

// GetClipIdOk returns a tuple with the ClipId field value
// and a boolean to check if the value has been set.
func (o *TrimmedVideo) GetClipIdOk() (*float32, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ClipId, true
}

// SetClipId sets field value
func (o *TrimmedVideo) SetClipId(v float32) {
	o.ClipId = v
}

// GetCreatedOn returns the CreatedOn field value
func (o *TrimmedVideo) GetCreatedOn() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.CreatedOn
}

// GetCreatedOnOk returns a tuple with the CreatedOn field value
// and a boolean to check if the value has been set.
func (o *TrimmedVideo) GetCreatedOnOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.CreatedOn, true
}

// SetCreatedOn sets field value
func (o *TrimmedVideo) SetCreatedOn(v string) {
	o.CreatedOn = v
}

// GetCreatedVersionId returns the CreatedVersionId field value
func (o *TrimmedVideo) GetCreatedVersionId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.CreatedVersionId
}

// GetCreatedVersionIdOk returns a tuple with the CreatedVersionId field value
// and a boolean to check if the value has been set.
func (o *TrimmedVideo) GetCreatedVersionIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.CreatedVersionId, true
}

// SetCreatedVersionId sets field value
func (o *TrimmedVideo) SetCreatedVersionId(v string) {
	o.CreatedVersionId = v
}

// GetEnd returns the End field value
func (o *TrimmedVideo) GetEnd() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.End
}

// GetEndOk returns a tuple with the End field value
// and a boolean to check if the value has been set.
func (o *TrimmedVideo) GetEndOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.End, true
}

// SetEnd sets field value
func (o *TrimmedVideo) SetEnd(v string) {
	o.End = v
}

// GetIsClipFinishedTranscoding returns the IsClipFinishedTranscoding field value
func (o *TrimmedVideo) GetIsClipFinishedTranscoding() bool {
	if o == nil {
		var ret bool
		return ret
	}

	return o.IsClipFinishedTranscoding
}

// GetIsClipFinishedTranscodingOk returns a tuple with the IsClipFinishedTranscoding field value
// and a boolean to check if the value has been set.
func (o *TrimmedVideo) GetIsClipFinishedTranscodingOk() (*bool, bool) {
	if o == nil {
		return nil, false
	}
	return &o.IsClipFinishedTranscoding, true
}

// SetIsClipFinishedTranscoding sets field value
func (o *TrimmedVideo) SetIsClipFinishedTranscoding(v bool) {
	o.IsClipFinishedTranscoding = v
}

// GetMetadata returns the Metadata field value
func (o *TrimmedVideo) GetMetadata() TrimmedVideoMetadata {
	if o == nil {
		var ret TrimmedVideoMetadata
		return ret
	}

	return o.Metadata
}

// GetMetadataOk returns a tuple with the Metadata field value
// and a boolean to check if the value has been set.
func (o *TrimmedVideo) GetMetadataOk() (*TrimmedVideoMetadata, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Metadata, true
}

// SetMetadata sets field value
func (o *TrimmedVideo) SetMetadata(v TrimmedVideoMetadata) {
	o.Metadata = v
}

// GetModifiedOn returns the ModifiedOn field value
// If the value is explicit nil, the zero value for string will be returned
func (o *TrimmedVideo) GetModifiedOn() string {
	if o == nil || o.ModifiedOn.Get() == nil {
		var ret string
		return ret
	}

	return *o.ModifiedOn.Get()
}

// GetModifiedOnOk returns a tuple with the ModifiedOn field value
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *TrimmedVideo) GetModifiedOnOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.ModifiedOn.Get(), o.ModifiedOn.IsSet()
}

// SetModifiedOn sets field value
func (o *TrimmedVideo) SetModifiedOn(v string) {
	o.ModifiedOn.Set(&v)
}

// GetRootVersionId returns the RootVersionId field value
func (o *TrimmedVideo) GetRootVersionId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.RootVersionId
}

// GetRootVersionIdOk returns a tuple with the RootVersionId field value
// and a boolean to check if the value has been set.
func (o *TrimmedVideo) GetRootVersionIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.RootVersionId, true
}

// SetRootVersionId sets field value
func (o *TrimmedVideo) SetRootVersionId(v string) {
	o.RootVersionId = v
}

// GetStart returns the Start field value
func (o *TrimmedVideo) GetStart() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Start
}

// GetStartOk returns a tuple with the Start field value
// and a boolean to check if the value has been set.
func (o *TrimmedVideo) GetStartOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Start, true
}

// SetStart sets field value
func (o *TrimmedVideo) SetStart(v string) {
	o.Start = v
}

// GetUri returns the Uri field value
func (o *TrimmedVideo) GetUri() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Uri
}

// GetUriOk returns a tuple with the Uri field value
// and a boolean to check if the value has been set.
func (o *TrimmedVideo) GetUriOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Uri, true
}

// SetUri sets field value
func (o *TrimmedVideo) SetUri(v string) {
	o.Uri = v
}

// GetVersionQuality returns the VersionQuality field value
func (o *TrimmedVideo) GetVersionQuality() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.VersionQuality
}

// GetVersionQualityOk returns a tuple with the VersionQuality field value
// and a boolean to check if the value has been set.
func (o *TrimmedVideo) GetVersionQualityOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.VersionQuality, true
}

// SetVersionQuality sets field value
func (o *TrimmedVideo) SetVersionQuality(v string) {
	o.VersionQuality = v
}

func (o TrimmedVideo) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o TrimmedVideo) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["clip_id"] = o.ClipId
	toSerialize["created_on"] = o.CreatedOn
	toSerialize["created_version_id"] = o.CreatedVersionId
	toSerialize["end"] = o.End
	toSerialize["is_clip_finished_transcoding"] = o.IsClipFinishedTranscoding
	toSerialize["metadata"] = o.Metadata
	toSerialize["modified_on"] = o.ModifiedOn.Get()
	toSerialize["root_version_id"] = o.RootVersionId
	toSerialize["start"] = o.Start
	toSerialize["uri"] = o.Uri
	toSerialize["version_quality"] = o.VersionQuality
	return toSerialize, nil
}

type NullableTrimmedVideo struct {
	value *TrimmedVideo
	isSet bool
}

func (v NullableTrimmedVideo) Get() *TrimmedVideo {
	return v.value
}

func (v *NullableTrimmedVideo) Set(val *TrimmedVideo) {
	v.value = val
	v.isSet = true
}

func (v NullableTrimmedVideo) IsSet() bool {
	return v.isSet
}

func (v *NullableTrimmedVideo) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableTrimmedVideo(val *TrimmedVideo) *NullableTrimmedVideo {
	return &NullableTrimmedVideo{value: val, isSet: true}
}

func (v NullableTrimmedVideo) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableTrimmedVideo) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


