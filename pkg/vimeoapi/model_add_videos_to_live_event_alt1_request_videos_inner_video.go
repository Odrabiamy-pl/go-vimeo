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

// checks if the AddVideosToLiveEventAlt1RequestVideosInnerVideo type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &AddVideosToLiveEventAlt1RequestVideosInnerVideo{}

// AddVideosToLiveEventAlt1RequestVideosInnerVideo struct for AddVideosToLiveEventAlt1RequestVideosInnerVideo
type AddVideosToLiveEventAlt1RequestVideosInnerVideo struct {
	// The URI of a video to add.
	Uri *string `json:"uri,omitempty"`
}

// NewAddVideosToLiveEventAlt1RequestVideosInnerVideo instantiates a new AddVideosToLiveEventAlt1RequestVideosInnerVideo object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewAddVideosToLiveEventAlt1RequestVideosInnerVideo() *AddVideosToLiveEventAlt1RequestVideosInnerVideo {
	this := AddVideosToLiveEventAlt1RequestVideosInnerVideo{}
	return &this
}

// NewAddVideosToLiveEventAlt1RequestVideosInnerVideoWithDefaults instantiates a new AddVideosToLiveEventAlt1RequestVideosInnerVideo object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewAddVideosToLiveEventAlt1RequestVideosInnerVideoWithDefaults() *AddVideosToLiveEventAlt1RequestVideosInnerVideo {
	this := AddVideosToLiveEventAlt1RequestVideosInnerVideo{}
	return &this
}

// GetUri returns the Uri field value if set, zero value otherwise.
func (o *AddVideosToLiveEventAlt1RequestVideosInnerVideo) GetUri() string {
	if o == nil || IsNil(o.Uri) {
		var ret string
		return ret
	}
	return *o.Uri
}

// GetUriOk returns a tuple with the Uri field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AddVideosToLiveEventAlt1RequestVideosInnerVideo) GetUriOk() (*string, bool) {
	if o == nil || IsNil(o.Uri) {
		return nil, false
	}
	return o.Uri, true
}

// HasUri returns a boolean if a field has been set.
func (o *AddVideosToLiveEventAlt1RequestVideosInnerVideo) HasUri() bool {
	if o != nil && !IsNil(o.Uri) {
		return true
	}

	return false
}

// SetUri gets a reference to the given string and assigns it to the Uri field.
func (o *AddVideosToLiveEventAlt1RequestVideosInnerVideo) SetUri(v string) {
	o.Uri = &v
}

func (o AddVideosToLiveEventAlt1RequestVideosInnerVideo) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o AddVideosToLiveEventAlt1RequestVideosInnerVideo) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Uri) {
		toSerialize["uri"] = o.Uri
	}
	return toSerialize, nil
}

type NullableAddVideosToLiveEventAlt1RequestVideosInnerVideo struct {
	value *AddVideosToLiveEventAlt1RequestVideosInnerVideo
	isSet bool
}

func (v NullableAddVideosToLiveEventAlt1RequestVideosInnerVideo) Get() *AddVideosToLiveEventAlt1RequestVideosInnerVideo {
	return v.value
}

func (v *NullableAddVideosToLiveEventAlt1RequestVideosInnerVideo) Set(val *AddVideosToLiveEventAlt1RequestVideosInnerVideo) {
	v.value = val
	v.isSet = true
}

func (v NullableAddVideosToLiveEventAlt1RequestVideosInnerVideo) IsSet() bool {
	return v.isSet
}

func (v *NullableAddVideosToLiveEventAlt1RequestVideosInnerVideo) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableAddVideosToLiveEventAlt1RequestVideosInnerVideo(val *AddVideosToLiveEventAlt1RequestVideosInnerVideo) *NullableAddVideosToLiveEventAlt1RequestVideosInnerVideo {
	return &NullableAddVideosToLiveEventAlt1RequestVideosInnerVideo{value: val, isSet: true}
}

func (v NullableAddVideosToLiveEventAlt1RequestVideosInnerVideo) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableAddVideosToLiveEventAlt1RequestVideosInnerVideo) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
