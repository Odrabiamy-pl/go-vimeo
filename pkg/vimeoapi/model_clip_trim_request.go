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

// checks if the ClipTrimRequest type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &ClipTrimRequest{}

// ClipTrimRequest struct for ClipTrimRequest
type ClipTrimRequest struct {
	// The end position in seconds of the trim in the video.
	TrimEnd *string `json:"trim_end,omitempty"`
	// The start position in seconds of the trim in the video.
	TrimStart *string `json:"trim_start,omitempty"`
}

// NewClipTrimRequest instantiates a new ClipTrimRequest object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewClipTrimRequest() *ClipTrimRequest {
	this := ClipTrimRequest{}
	return &this
}

// NewClipTrimRequestWithDefaults instantiates a new ClipTrimRequest object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewClipTrimRequestWithDefaults() *ClipTrimRequest {
	this := ClipTrimRequest{}
	return &this
}

// GetTrimEnd returns the TrimEnd field value if set, zero value otherwise.
func (o *ClipTrimRequest) GetTrimEnd() string {
	if o == nil || IsNil(o.TrimEnd) {
		var ret string
		return ret
	}
	return *o.TrimEnd
}

// GetTrimEndOk returns a tuple with the TrimEnd field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ClipTrimRequest) GetTrimEndOk() (*string, bool) {
	if o == nil || IsNil(o.TrimEnd) {
		return nil, false
	}
	return o.TrimEnd, true
}

// HasTrimEnd returns a boolean if a field has been set.
func (o *ClipTrimRequest) HasTrimEnd() bool {
	if o != nil && !IsNil(o.TrimEnd) {
		return true
	}

	return false
}

// SetTrimEnd gets a reference to the given string and assigns it to the TrimEnd field.
func (o *ClipTrimRequest) SetTrimEnd(v string) {
	o.TrimEnd = &v
}

// GetTrimStart returns the TrimStart field value if set, zero value otherwise.
func (o *ClipTrimRequest) GetTrimStart() string {
	if o == nil || IsNil(o.TrimStart) {
		var ret string
		return ret
	}
	return *o.TrimStart
}

// GetTrimStartOk returns a tuple with the TrimStart field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ClipTrimRequest) GetTrimStartOk() (*string, bool) {
	if o == nil || IsNil(o.TrimStart) {
		return nil, false
	}
	return o.TrimStart, true
}

// HasTrimStart returns a boolean if a field has been set.
func (o *ClipTrimRequest) HasTrimStart() bool {
	if o != nil && !IsNil(o.TrimStart) {
		return true
	}

	return false
}

// SetTrimStart gets a reference to the given string and assigns it to the TrimStart field.
func (o *ClipTrimRequest) SetTrimStart(v string) {
	o.TrimStart = &v
}

func (o ClipTrimRequest) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o ClipTrimRequest) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.TrimEnd) {
		toSerialize["trim_end"] = o.TrimEnd
	}
	if !IsNil(o.TrimStart) {
		toSerialize["trim_start"] = o.TrimStart
	}
	return toSerialize, nil
}

type NullableClipTrimRequest struct {
	value *ClipTrimRequest
	isSet bool
}

func (v NullableClipTrimRequest) Get() *ClipTrimRequest {
	return v.value
}

func (v *NullableClipTrimRequest) Set(val *ClipTrimRequest) {
	v.value = val
	v.isSet = true
}

func (v NullableClipTrimRequest) IsSet() bool {
	return v.isSet
}

func (v *NullableClipTrimRequest) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableClipTrimRequest(val *ClipTrimRequest) *NullableClipTrimRequest {
	return &NullableClipTrimRequest{value: val, isSet: true}
}

func (v NullableClipTrimRequest) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableClipTrimRequest) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


