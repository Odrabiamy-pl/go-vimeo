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

// checks if the CategoryMetadataInteractions type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &CategoryMetadataInteractions{}

// CategoryMetadataInteractions The permissible actions related to the category.
type CategoryMetadataInteractions struct {
	Follow CategoryMetadataInteractionsFollow `json:"follow"`
}

// NewCategoryMetadataInteractions instantiates a new CategoryMetadataInteractions object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewCategoryMetadataInteractions(follow CategoryMetadataInteractionsFollow) *CategoryMetadataInteractions {
	this := CategoryMetadataInteractions{}
	this.Follow = follow
	return &this
}

// NewCategoryMetadataInteractionsWithDefaults instantiates a new CategoryMetadataInteractions object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewCategoryMetadataInteractionsWithDefaults() *CategoryMetadataInteractions {
	this := CategoryMetadataInteractions{}
	return &this
}

// GetFollow returns the Follow field value
func (o *CategoryMetadataInteractions) GetFollow() CategoryMetadataInteractionsFollow {
	if o == nil {
		var ret CategoryMetadataInteractionsFollow
		return ret
	}

	return o.Follow
}

// GetFollowOk returns a tuple with the Follow field value
// and a boolean to check if the value has been set.
func (o *CategoryMetadataInteractions) GetFollowOk() (*CategoryMetadataInteractionsFollow, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Follow, true
}

// SetFollow sets field value
func (o *CategoryMetadataInteractions) SetFollow(v CategoryMetadataInteractionsFollow) {
	o.Follow = v
}

func (o CategoryMetadataInteractions) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o CategoryMetadataInteractions) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["follow"] = o.Follow
	return toSerialize, nil
}

type NullableCategoryMetadataInteractions struct {
	value *CategoryMetadataInteractions
	isSet bool
}

func (v NullableCategoryMetadataInteractions) Get() *CategoryMetadataInteractions {
	return v.value
}

func (v *NullableCategoryMetadataInteractions) Set(val *CategoryMetadataInteractions) {
	v.value = val
	v.isSet = true
}

func (v NullableCategoryMetadataInteractions) IsSet() bool {
	return v.isSet
}

func (v *NullableCategoryMetadataInteractions) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableCategoryMetadataInteractions(val *CategoryMetadataInteractions) *NullableCategoryMetadataInteractions {
	return &NullableCategoryMetadataInteractions{value: val, isSet: true}
}

func (v NullableCategoryMetadataInteractions) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableCategoryMetadataInteractions) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


