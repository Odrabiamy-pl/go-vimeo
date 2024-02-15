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

// checks if the VideoVersionTranscode type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &VideoVersionTranscode{}

// VideoVersionTranscode The version's transcode information.
type VideoVersionTranscode struct {
	// The status code for the availability of the video version.  Option descriptions:  * `complete` - Transcoding is complete. The video version is available.  * `error` - There was a transcoding error. The video version isn't available.  * `in_progress` - Transcoding is in progress. The video version isn't available yet. 
	Status *string `json:"status,omitempty"`
}

// NewVideoVersionTranscode instantiates a new VideoVersionTranscode object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewVideoVersionTranscode() *VideoVersionTranscode {
	this := VideoVersionTranscode{}
	return &this
}

// NewVideoVersionTranscodeWithDefaults instantiates a new VideoVersionTranscode object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewVideoVersionTranscodeWithDefaults() *VideoVersionTranscode {
	this := VideoVersionTranscode{}
	return &this
}

// GetStatus returns the Status field value if set, zero value otherwise.
func (o *VideoVersionTranscode) GetStatus() string {
	if o == nil || IsNil(o.Status) {
		var ret string
		return ret
	}
	return *o.Status
}

// GetStatusOk returns a tuple with the Status field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *VideoVersionTranscode) GetStatusOk() (*string, bool) {
	if o == nil || IsNil(o.Status) {
		return nil, false
	}
	return o.Status, true
}

// HasStatus returns a boolean if a field has been set.
func (o *VideoVersionTranscode) HasStatus() bool {
	if o != nil && !IsNil(o.Status) {
		return true
	}

	return false
}

// SetStatus gets a reference to the given string and assigns it to the Status field.
func (o *VideoVersionTranscode) SetStatus(v string) {
	o.Status = &v
}

func (o VideoVersionTranscode) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o VideoVersionTranscode) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Status) {
		toSerialize["status"] = o.Status
	}
	return toSerialize, nil
}

type NullableVideoVersionTranscode struct {
	value *VideoVersionTranscode
	isSet bool
}

func (v NullableVideoVersionTranscode) Get() *VideoVersionTranscode {
	return v.value
}

func (v *NullableVideoVersionTranscode) Set(val *VideoVersionTranscode) {
	v.value = val
	v.isSet = true
}

func (v NullableVideoVersionTranscode) IsSet() bool {
	return v.isSet
}

func (v *NullableVideoVersionTranscode) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableVideoVersionTranscode(val *VideoVersionTranscode) *NullableVideoVersionTranscode {
	return &NullableVideoVersionTranscode{value: val, isSet: true}
}

func (v NullableVideoVersionTranscode) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableVideoVersionTranscode) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

