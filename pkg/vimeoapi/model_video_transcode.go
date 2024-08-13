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

// checks if the VideoTranscode type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &VideoTranscode{}

// VideoTranscode The transcode information of the video upload.
type VideoTranscode struct {
	// The video's availability status.  Option descriptions:  * `complete` - Transcoding is complete. The video is available.  * `error` - There was a transcoding error. The video isn't available.  * `in_progress` - Transcoding is currently underway. The video isn't available yet. 
	Status *string `json:"status,omitempty"`
}

// NewVideoTranscode instantiates a new VideoTranscode object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewVideoTranscode() *VideoTranscode {
	this := VideoTranscode{}
	return &this
}

// NewVideoTranscodeWithDefaults instantiates a new VideoTranscode object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewVideoTranscodeWithDefaults() *VideoTranscode {
	this := VideoTranscode{}
	return &this
}

// GetStatus returns the Status field value if set, zero value otherwise.
func (o *VideoTranscode) GetStatus() string {
	if o == nil || IsNil(o.Status) {
		var ret string
		return ret
	}
	return *o.Status
}

// GetStatusOk returns a tuple with the Status field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *VideoTranscode) GetStatusOk() (*string, bool) {
	if o == nil || IsNil(o.Status) {
		return nil, false
	}
	return o.Status, true
}

// HasStatus returns a boolean if a field has been set.
func (o *VideoTranscode) HasStatus() bool {
	if o != nil && !IsNil(o.Status) {
		return true
	}

	return false
}

// SetStatus gets a reference to the given string and assigns it to the Status field.
func (o *VideoTranscode) SetStatus(v string) {
	o.Status = &v
}

func (o VideoTranscode) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o VideoTranscode) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Status) {
		toSerialize["status"] = o.Status
	}
	return toSerialize, nil
}

type NullableVideoTranscode struct {
	value *VideoTranscode
	isSet bool
}

func (v NullableVideoTranscode) Get() *VideoTranscode {
	return v.value
}

func (v *NullableVideoTranscode) Set(val *VideoTranscode) {
	v.value = val
	v.isSet = true
}

func (v NullableVideoTranscode) IsSet() bool {
	return v.isSet
}

func (v *NullableVideoTranscode) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableVideoTranscode(val *VideoTranscode) *NullableVideoTranscode {
	return &NullableVideoTranscode{value: val, isSet: true}
}

func (v NullableVideoTranscode) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableVideoTranscode) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


