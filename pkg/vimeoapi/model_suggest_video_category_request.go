/*
Vimeo API

Build something great. Vimeo's API supports flexible, high-quality video integration with your custom apps.

API version: 3.4
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package vimeoapi

import (
	"encoding/json"
	"bytes"
	"fmt"
)

// checks if the SuggestVideoCategoryRequest type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &SuggestVideoCategoryRequest{}

// SuggestVideoCategoryRequest struct for SuggestVideoCategoryRequest
type SuggestVideoCategoryRequest struct {
	// An array of the names of the desired categories.
	Category []string `json:"category"`
}

type _SuggestVideoCategoryRequest SuggestVideoCategoryRequest

// NewSuggestVideoCategoryRequest instantiates a new SuggestVideoCategoryRequest object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewSuggestVideoCategoryRequest(category []string) *SuggestVideoCategoryRequest {
	this := SuggestVideoCategoryRequest{}
	this.Category = category
	return &this
}

// NewSuggestVideoCategoryRequestWithDefaults instantiates a new SuggestVideoCategoryRequest object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewSuggestVideoCategoryRequestWithDefaults() *SuggestVideoCategoryRequest {
	this := SuggestVideoCategoryRequest{}
	return &this
}

// GetCategory returns the Category field value
func (o *SuggestVideoCategoryRequest) GetCategory() []string {
	if o == nil {
		var ret []string
		return ret
	}

	return o.Category
}

// GetCategoryOk returns a tuple with the Category field value
// and a boolean to check if the value has been set.
func (o *SuggestVideoCategoryRequest) GetCategoryOk() ([]string, bool) {
	if o == nil {
		return nil, false
	}
	return o.Category, true
}

// SetCategory sets field value
func (o *SuggestVideoCategoryRequest) SetCategory(v []string) {
	o.Category = v
}

func (o SuggestVideoCategoryRequest) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o SuggestVideoCategoryRequest) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["category"] = o.Category
	return toSerialize, nil
}

func (o *SuggestVideoCategoryRequest) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"category",
	}

	allProperties := make(map[string]interface{})

	err = json.Unmarshal(data, &allProperties)

	if err != nil {
		return err;
	}

	for _, requiredProperty := range(requiredProperties) {
		if _, exists := allProperties[requiredProperty]; !exists {
			return fmt.Errorf("no value given for required property %v", requiredProperty)
		}
	}

	varSuggestVideoCategoryRequest := _SuggestVideoCategoryRequest{}

	decoder := json.NewDecoder(bytes.NewReader(data))
	decoder.DisallowUnknownFields()
	err = decoder.Decode(&varSuggestVideoCategoryRequest)

	if err != nil {
		return err
	}

	*o = SuggestVideoCategoryRequest(varSuggestVideoCategoryRequest)

	return err
}

type NullableSuggestVideoCategoryRequest struct {
	value *SuggestVideoCategoryRequest
	isSet bool
}

func (v NullableSuggestVideoCategoryRequest) Get() *SuggestVideoCategoryRequest {
	return v.value
}

func (v *NullableSuggestVideoCategoryRequest) Set(val *SuggestVideoCategoryRequest) {
	v.value = val
	v.isSet = true
}

func (v NullableSuggestVideoCategoryRequest) IsSet() bool {
	return v.isSet
}

func (v *NullableSuggestVideoCategoryRequest) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableSuggestVideoCategoryRequest(val *SuggestVideoCategoryRequest) *NullableSuggestVideoCategoryRequest {
	return &NullableSuggestVideoCategoryRequest{value: val, isSet: true}
}

func (v NullableSuggestVideoCategoryRequest) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableSuggestVideoCategoryRequest) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


