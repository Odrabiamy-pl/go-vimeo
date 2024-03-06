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

// checks if the EditVideoVersionRequest type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &EditVideoVersionRequest{}

// EditVideoVersionRequest struct for EditVideoVersionRequest
type EditVideoVersionRequest struct {
	// A description of the video version. This description can make use of the full unicode character set.
	Description *string `json:"description,omitempty"`
	// Whether the video version is active.
	IsCurrent *bool `json:"is_current,omitempty"`
}

// NewEditVideoVersionRequest instantiates a new EditVideoVersionRequest object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewEditVideoVersionRequest() *EditVideoVersionRequest {
	this := EditVideoVersionRequest{}
	return &this
}

// NewEditVideoVersionRequestWithDefaults instantiates a new EditVideoVersionRequest object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewEditVideoVersionRequestWithDefaults() *EditVideoVersionRequest {
	this := EditVideoVersionRequest{}
	return &this
}

// GetDescription returns the Description field value if set, zero value otherwise.
func (o *EditVideoVersionRequest) GetDescription() string {
	if o == nil || IsNil(o.Description) {
		var ret string
		return ret
	}
	return *o.Description
}

// GetDescriptionOk returns a tuple with the Description field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *EditVideoVersionRequest) GetDescriptionOk() (*string, bool) {
	if o == nil || IsNil(o.Description) {
		return nil, false
	}
	return o.Description, true
}

// HasDescription returns a boolean if a field has been set.
func (o *EditVideoVersionRequest) HasDescription() bool {
	if o != nil && !IsNil(o.Description) {
		return true
	}

	return false
}

// SetDescription gets a reference to the given string and assigns it to the Description field.
func (o *EditVideoVersionRequest) SetDescription(v string) {
	o.Description = &v
}

// GetIsCurrent returns the IsCurrent field value if set, zero value otherwise.
func (o *EditVideoVersionRequest) GetIsCurrent() bool {
	if o == nil || IsNil(o.IsCurrent) {
		var ret bool
		return ret
	}
	return *o.IsCurrent
}

// GetIsCurrentOk returns a tuple with the IsCurrent field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *EditVideoVersionRequest) GetIsCurrentOk() (*bool, bool) {
	if o == nil || IsNil(o.IsCurrent) {
		return nil, false
	}
	return o.IsCurrent, true
}

// HasIsCurrent returns a boolean if a field has been set.
func (o *EditVideoVersionRequest) HasIsCurrent() bool {
	if o != nil && !IsNil(o.IsCurrent) {
		return true
	}

	return false
}

// SetIsCurrent gets a reference to the given bool and assigns it to the IsCurrent field.
func (o *EditVideoVersionRequest) SetIsCurrent(v bool) {
	o.IsCurrent = &v
}

func (o EditVideoVersionRequest) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o EditVideoVersionRequest) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Description) {
		toSerialize["description"] = o.Description
	}
	if !IsNil(o.IsCurrent) {
		toSerialize["is_current"] = o.IsCurrent
	}
	return toSerialize, nil
}

type NullableEditVideoVersionRequest struct {
	value *EditVideoVersionRequest
	isSet bool
}

func (v NullableEditVideoVersionRequest) Get() *EditVideoVersionRequest {
	return v.value
}

func (v *NullableEditVideoVersionRequest) Set(val *EditVideoVersionRequest) {
	v.value = val
	v.isSet = true
}

func (v NullableEditVideoVersionRequest) IsSet() bool {
	return v.isSet
}

func (v *NullableEditVideoVersionRequest) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableEditVideoVersionRequest(val *EditVideoVersionRequest) *NullableEditVideoVersionRequest {
	return &NullableEditVideoVersionRequest{value: val, isSet: true}
}

func (v NullableEditVideoVersionRequest) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableEditVideoVersionRequest) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
