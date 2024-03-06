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

// checks if the SubscriptionPlansDiscount type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &SubscriptionPlansDiscount{}

// SubscriptionPlansDiscount Information about the plan discount.
type SubscriptionPlansDiscount struct {
	// The annual discount.
	Annual float32 `json:"annual"`
}

// NewSubscriptionPlansDiscount instantiates a new SubscriptionPlansDiscount object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewSubscriptionPlansDiscount(annual float32) *SubscriptionPlansDiscount {
	this := SubscriptionPlansDiscount{}
	this.Annual = annual
	return &this
}

// NewSubscriptionPlansDiscountWithDefaults instantiates a new SubscriptionPlansDiscount object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewSubscriptionPlansDiscountWithDefaults() *SubscriptionPlansDiscount {
	this := SubscriptionPlansDiscount{}
	return &this
}

// GetAnnual returns the Annual field value
func (o *SubscriptionPlansDiscount) GetAnnual() float32 {
	if o == nil {
		var ret float32
		return ret
	}

	return o.Annual
}

// GetAnnualOk returns a tuple with the Annual field value
// and a boolean to check if the value has been set.
func (o *SubscriptionPlansDiscount) GetAnnualOk() (*float32, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Annual, true
}

// SetAnnual sets field value
func (o *SubscriptionPlansDiscount) SetAnnual(v float32) {
	o.Annual = v
}

func (o SubscriptionPlansDiscount) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o SubscriptionPlansDiscount) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["annual"] = o.Annual
	return toSerialize, nil
}

type NullableSubscriptionPlansDiscount struct {
	value *SubscriptionPlansDiscount
	isSet bool
}

func (v NullableSubscriptionPlansDiscount) Get() *SubscriptionPlansDiscount {
	return v.value
}

func (v *NullableSubscriptionPlansDiscount) Set(val *SubscriptionPlansDiscount) {
	v.value = val
	v.isSet = true
}

func (v NullableSubscriptionPlansDiscount) IsSet() bool {
	return v.isSet
}

func (v *NullableSubscriptionPlansDiscount) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableSubscriptionPlansDiscount(val *SubscriptionPlansDiscount) *NullableSubscriptionPlansDiscount {
	return &NullableSubscriptionPlansDiscount{value: val, isSet: true}
}

func (v NullableSubscriptionPlansDiscount) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableSubscriptionPlansDiscount) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
