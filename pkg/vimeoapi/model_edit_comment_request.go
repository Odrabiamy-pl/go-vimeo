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

// checks if the EditCommentRequest type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &EditCommentRequest{}

// EditCommentRequest struct for EditCommentRequest
type EditCommentRequest struct {
	// The new comment text.
	Text string `json:"text"`
}

// NewEditCommentRequest instantiates a new EditCommentRequest object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewEditCommentRequest(text string) *EditCommentRequest {
	this := EditCommentRequest{}
	this.Text = text
	return &this
}

// NewEditCommentRequestWithDefaults instantiates a new EditCommentRequest object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewEditCommentRequestWithDefaults() *EditCommentRequest {
	this := EditCommentRequest{}
	return &this
}

// GetText returns the Text field value
func (o *EditCommentRequest) GetText() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Text
}

// GetTextOk returns a tuple with the Text field value
// and a boolean to check if the value has been set.
func (o *EditCommentRequest) GetTextOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Text, true
}

// SetText sets field value
func (o *EditCommentRequest) SetText(v string) {
	o.Text = v
}

func (o EditCommentRequest) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o EditCommentRequest) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["text"] = o.Text
	return toSerialize, nil
}

type NullableEditCommentRequest struct {
	value *EditCommentRequest
	isSet bool
}

func (v NullableEditCommentRequest) Get() *EditCommentRequest {
	return v.value
}

func (v *NullableEditCommentRequest) Set(val *EditCommentRequest) {
	v.value = val
	v.isSet = true
}

func (v NullableEditCommentRequest) IsSet() bool {
	return v.isSet
}

func (v *NullableEditCommentRequest) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableEditCommentRequest(val *EditCommentRequest) *NullableEditCommentRequest {
	return &NullableEditCommentRequest{value: val, isSet: true}
}

func (v NullableEditCommentRequest) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableEditCommentRequest) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


