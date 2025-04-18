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

// checks if the SubscriptionPlansPromotionPrice type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &SubscriptionPlansPromotionPrice{}

// SubscriptionPlansPromotionPrice The price map of the promotion.
type SubscriptionPlansPromotionPrice struct {
	// The promotional annual price, charged annually.
	Annual *float32 `json:"annual,omitempty"`
	// The promotional monthly price, charged annually.
	AnnualMonthly *float32 `json:"annual_monthly,omitempty"`
	// The promotional monthly price, charged monthly.
	Monthly *float32 `json:"monthly,omitempty"`
}

// NewSubscriptionPlansPromotionPrice instantiates a new SubscriptionPlansPromotionPrice object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewSubscriptionPlansPromotionPrice() *SubscriptionPlansPromotionPrice {
	this := SubscriptionPlansPromotionPrice{}
	return &this
}

// NewSubscriptionPlansPromotionPriceWithDefaults instantiates a new SubscriptionPlansPromotionPrice object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewSubscriptionPlansPromotionPriceWithDefaults() *SubscriptionPlansPromotionPrice {
	this := SubscriptionPlansPromotionPrice{}
	return &this
}

// GetAnnual returns the Annual field value if set, zero value otherwise.
func (o *SubscriptionPlansPromotionPrice) GetAnnual() float32 {
	if o == nil || IsNil(o.Annual) {
		var ret float32
		return ret
	}
	return *o.Annual
}

// GetAnnualOk returns a tuple with the Annual field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SubscriptionPlansPromotionPrice) GetAnnualOk() (*float32, bool) {
	if o == nil || IsNil(o.Annual) {
		return nil, false
	}
	return o.Annual, true
}

// HasAnnual returns a boolean if a field has been set.
func (o *SubscriptionPlansPromotionPrice) HasAnnual() bool {
	if o != nil && !IsNil(o.Annual) {
		return true
	}

	return false
}

// SetAnnual gets a reference to the given float32 and assigns it to the Annual field.
func (o *SubscriptionPlansPromotionPrice) SetAnnual(v float32) {
	o.Annual = &v
}

// GetAnnualMonthly returns the AnnualMonthly field value if set, zero value otherwise.
func (o *SubscriptionPlansPromotionPrice) GetAnnualMonthly() float32 {
	if o == nil || IsNil(o.AnnualMonthly) {
		var ret float32
		return ret
	}
	return *o.AnnualMonthly
}

// GetAnnualMonthlyOk returns a tuple with the AnnualMonthly field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SubscriptionPlansPromotionPrice) GetAnnualMonthlyOk() (*float32, bool) {
	if o == nil || IsNil(o.AnnualMonthly) {
		return nil, false
	}
	return o.AnnualMonthly, true
}

// HasAnnualMonthly returns a boolean if a field has been set.
func (o *SubscriptionPlansPromotionPrice) HasAnnualMonthly() bool {
	if o != nil && !IsNil(o.AnnualMonthly) {
		return true
	}

	return false
}

// SetAnnualMonthly gets a reference to the given float32 and assigns it to the AnnualMonthly field.
func (o *SubscriptionPlansPromotionPrice) SetAnnualMonthly(v float32) {
	o.AnnualMonthly = &v
}

// GetMonthly returns the Monthly field value if set, zero value otherwise.
func (o *SubscriptionPlansPromotionPrice) GetMonthly() float32 {
	if o == nil || IsNil(o.Monthly) {
		var ret float32
		return ret
	}
	return *o.Monthly
}

// GetMonthlyOk returns a tuple with the Monthly field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SubscriptionPlansPromotionPrice) GetMonthlyOk() (*float32, bool) {
	if o == nil || IsNil(o.Monthly) {
		return nil, false
	}
	return o.Monthly, true
}

// HasMonthly returns a boolean if a field has been set.
func (o *SubscriptionPlansPromotionPrice) HasMonthly() bool {
	if o != nil && !IsNil(o.Monthly) {
		return true
	}

	return false
}

// SetMonthly gets a reference to the given float32 and assigns it to the Monthly field.
func (o *SubscriptionPlansPromotionPrice) SetMonthly(v float32) {
	o.Monthly = &v
}

func (o SubscriptionPlansPromotionPrice) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o SubscriptionPlansPromotionPrice) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Annual) {
		toSerialize["annual"] = o.Annual
	}
	if !IsNil(o.AnnualMonthly) {
		toSerialize["annual_monthly"] = o.AnnualMonthly
	}
	if !IsNil(o.Monthly) {
		toSerialize["monthly"] = o.Monthly
	}
	return toSerialize, nil
}

type NullableSubscriptionPlansPromotionPrice struct {
	value *SubscriptionPlansPromotionPrice
	isSet bool
}

func (v NullableSubscriptionPlansPromotionPrice) Get() *SubscriptionPlansPromotionPrice {
	return v.value
}

func (v *NullableSubscriptionPlansPromotionPrice) Set(val *SubscriptionPlansPromotionPrice) {
	v.value = val
	v.isSet = true
}

func (v NullableSubscriptionPlansPromotionPrice) IsSet() bool {
	return v.isSet
}

func (v *NullableSubscriptionPlansPromotionPrice) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableSubscriptionPlansPromotionPrice(val *SubscriptionPlansPromotionPrice) *NullableSubscriptionPlansPromotionPrice {
	return &NullableSubscriptionPlansPromotionPrice{value: val, isSet: true}
}

func (v NullableSubscriptionPlansPromotionPrice) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableSubscriptionPlansPromotionPrice) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


