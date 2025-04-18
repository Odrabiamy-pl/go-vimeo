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

// checks if the ReplaceShowcaseLogoRequest type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &ReplaceShowcaseLogoRequest{}

// ReplaceShowcaseLogoRequest struct for ReplaceShowcaseLogoRequest
type ReplaceShowcaseLogoRequest struct {
	// Whether to make this image the active showcase logo.
	Active *bool `json:"active,omitempty"`
}

// NewReplaceShowcaseLogoRequest instantiates a new ReplaceShowcaseLogoRequest object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewReplaceShowcaseLogoRequest() *ReplaceShowcaseLogoRequest {
	this := ReplaceShowcaseLogoRequest{}
	return &this
}

// NewReplaceShowcaseLogoRequestWithDefaults instantiates a new ReplaceShowcaseLogoRequest object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewReplaceShowcaseLogoRequestWithDefaults() *ReplaceShowcaseLogoRequest {
	this := ReplaceShowcaseLogoRequest{}
	return &this
}

// GetActive returns the Active field value if set, zero value otherwise.
func (o *ReplaceShowcaseLogoRequest) GetActive() bool {
	if o == nil || IsNil(o.Active) {
		var ret bool
		return ret
	}
	return *o.Active
}

// GetActiveOk returns a tuple with the Active field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ReplaceShowcaseLogoRequest) GetActiveOk() (*bool, bool) {
	if o == nil || IsNil(o.Active) {
		return nil, false
	}
	return o.Active, true
}

// HasActive returns a boolean if a field has been set.
func (o *ReplaceShowcaseLogoRequest) HasActive() bool {
	if o != nil && !IsNil(o.Active) {
		return true
	}

	return false
}

// SetActive gets a reference to the given bool and assigns it to the Active field.
func (o *ReplaceShowcaseLogoRequest) SetActive(v bool) {
	o.Active = &v
}

func (o ReplaceShowcaseLogoRequest) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o ReplaceShowcaseLogoRequest) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Active) {
		toSerialize["active"] = o.Active
	}
	return toSerialize, nil
}

type NullableReplaceShowcaseLogoRequest struct {
	value *ReplaceShowcaseLogoRequest
	isSet bool
}

func (v NullableReplaceShowcaseLogoRequest) Get() *ReplaceShowcaseLogoRequest {
	return v.value
}

func (v *NullableReplaceShowcaseLogoRequest) Set(val *ReplaceShowcaseLogoRequest) {
	v.value = val
	v.isSet = true
}

func (v NullableReplaceShowcaseLogoRequest) IsSet() bool {
	return v.isSet
}

func (v *NullableReplaceShowcaseLogoRequest) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableReplaceShowcaseLogoRequest(val *ReplaceShowcaseLogoRequest) *NullableReplaceShowcaseLogoRequest {
	return &NullableReplaceShowcaseLogoRequest{value: val, isSet: true}
}

func (v NullableReplaceShowcaseLogoRequest) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableReplaceShowcaseLogoRequest) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


