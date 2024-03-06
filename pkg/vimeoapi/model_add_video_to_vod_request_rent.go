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

// checks if the AddVideoToVodRequestRent type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &AddVideoToVodRequestRent{}

// AddVideoToVodRequestRent struct for AddVideoToVodRequestRent
type AddVideoToVodRequestRent struct {
	Price *AddVideoToVodRequestRentPrice `json:"price,omitempty"`
}

// NewAddVideoToVodRequestRent instantiates a new AddVideoToVodRequestRent object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewAddVideoToVodRequestRent() *AddVideoToVodRequestRent {
	this := AddVideoToVodRequestRent{}
	return &this
}

// NewAddVideoToVodRequestRentWithDefaults instantiates a new AddVideoToVodRequestRent object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewAddVideoToVodRequestRentWithDefaults() *AddVideoToVodRequestRent {
	this := AddVideoToVodRequestRent{}
	return &this
}

// GetPrice returns the Price field value if set, zero value otherwise.
func (o *AddVideoToVodRequestRent) GetPrice() AddVideoToVodRequestRentPrice {
	if o == nil || IsNil(o.Price) {
		var ret AddVideoToVodRequestRentPrice
		return ret
	}
	return *o.Price
}

// GetPriceOk returns a tuple with the Price field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AddVideoToVodRequestRent) GetPriceOk() (*AddVideoToVodRequestRentPrice, bool) {
	if o == nil || IsNil(o.Price) {
		return nil, false
	}
	return o.Price, true
}

// HasPrice returns a boolean if a field has been set.
func (o *AddVideoToVodRequestRent) HasPrice() bool {
	if o != nil && !IsNil(o.Price) {
		return true
	}

	return false
}

// SetPrice gets a reference to the given AddVideoToVodRequestRentPrice and assigns it to the Price field.
func (o *AddVideoToVodRequestRent) SetPrice(v AddVideoToVodRequestRentPrice) {
	o.Price = &v
}

func (o AddVideoToVodRequestRent) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o AddVideoToVodRequestRent) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Price) {
		toSerialize["price"] = o.Price
	}
	return toSerialize, nil
}

type NullableAddVideoToVodRequestRent struct {
	value *AddVideoToVodRequestRent
	isSet bool
}

func (v NullableAddVideoToVodRequestRent) Get() *AddVideoToVodRequestRent {
	return v.value
}

func (v *NullableAddVideoToVodRequestRent) Set(val *AddVideoToVodRequestRent) {
	v.value = val
	v.isSet = true
}

func (v NullableAddVideoToVodRequestRent) IsSet() bool {
	return v.isSet
}

func (v *NullableAddVideoToVodRequestRent) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableAddVideoToVodRequestRent(val *AddVideoToVodRequestRent) *NullableAddVideoToVodRequestRent {
	return &NullableAddVideoToVodRequestRent{value: val, isSet: true}
}

func (v NullableAddVideoToVodRequestRent) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableAddVideoToVodRequestRent) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
