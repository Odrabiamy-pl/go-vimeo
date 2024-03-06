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

// checks if the AddTagsToChannelRequest type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &AddTagsToChannelRequest{}

// AddTagsToChannelRequest struct for AddTagsToChannelRequest
type AddTagsToChannelRequest struct {
	// An array of tags to assign.
	Tag []string `json:"tag"`
}

// NewAddTagsToChannelRequest instantiates a new AddTagsToChannelRequest object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewAddTagsToChannelRequest(tag []string) *AddTagsToChannelRequest {
	this := AddTagsToChannelRequest{}
	this.Tag = tag
	return &this
}

// NewAddTagsToChannelRequestWithDefaults instantiates a new AddTagsToChannelRequest object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewAddTagsToChannelRequestWithDefaults() *AddTagsToChannelRequest {
	this := AddTagsToChannelRequest{}
	return &this
}

// GetTag returns the Tag field value
func (o *AddTagsToChannelRequest) GetTag() []string {
	if o == nil {
		var ret []string
		return ret
	}

	return o.Tag
}

// GetTagOk returns a tuple with the Tag field value
// and a boolean to check if the value has been set.
func (o *AddTagsToChannelRequest) GetTagOk() ([]string, bool) {
	if o == nil {
		return nil, false
	}
	return o.Tag, true
}

// SetTag sets field value
func (o *AddTagsToChannelRequest) SetTag(v []string) {
	o.Tag = v
}

func (o AddTagsToChannelRequest) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o AddTagsToChannelRequest) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["tag"] = o.Tag
	return toSerialize, nil
}

type NullableAddTagsToChannelRequest struct {
	value *AddTagsToChannelRequest
	isSet bool
}

func (v NullableAddTagsToChannelRequest) Get() *AddTagsToChannelRequest {
	return v.value
}

func (v *NullableAddTagsToChannelRequest) Set(val *AddTagsToChannelRequest) {
	v.value = val
	v.isSet = true
}

func (v NullableAddTagsToChannelRequest) IsSet() bool {
	return v.isSet
}

func (v *NullableAddTagsToChannelRequest) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableAddTagsToChannelRequest(val *AddTagsToChannelRequest) *NullableAddTagsToChannelRequest {
	return &NullableAddTagsToChannelRequest{value: val, isSet: true}
}

func (v NullableAddTagsToChannelRequest) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableAddTagsToChannelRequest) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
