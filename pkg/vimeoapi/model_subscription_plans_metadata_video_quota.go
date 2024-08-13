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

// checks if the SubscriptionPlansMetadataVideoQuota type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &SubscriptionPlansMetadataVideoQuota{}

// SubscriptionPlansMetadataVideoQuota The video upload quotas associated with the product.
type SubscriptionPlansMetadataVideoQuota struct {
	// The total video upload quota associated with the product.
	Lifetime NullableFloat32 `json:"lifetime"`
	// The monthly video upload quota associated with the product.
	Monthly float32 `json:"monthly"`
	// The yearly video upload quota associated with the product.
	Yearly float32 `json:"yearly"`
}

// NewSubscriptionPlansMetadataVideoQuota instantiates a new SubscriptionPlansMetadataVideoQuota object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewSubscriptionPlansMetadataVideoQuota(lifetime NullableFloat32, monthly float32, yearly float32) *SubscriptionPlansMetadataVideoQuota {
	this := SubscriptionPlansMetadataVideoQuota{}
	this.Lifetime = lifetime
	this.Monthly = monthly
	this.Yearly = yearly
	return &this
}

// NewSubscriptionPlansMetadataVideoQuotaWithDefaults instantiates a new SubscriptionPlansMetadataVideoQuota object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewSubscriptionPlansMetadataVideoQuotaWithDefaults() *SubscriptionPlansMetadataVideoQuota {
	this := SubscriptionPlansMetadataVideoQuota{}
	return &this
}

// GetLifetime returns the Lifetime field value
// If the value is explicit nil, the zero value for float32 will be returned
func (o *SubscriptionPlansMetadataVideoQuota) GetLifetime() float32 {
	if o == nil || o.Lifetime.Get() == nil {
		var ret float32
		return ret
	}

	return *o.Lifetime.Get()
}

// GetLifetimeOk returns a tuple with the Lifetime field value
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *SubscriptionPlansMetadataVideoQuota) GetLifetimeOk() (*float32, bool) {
	if o == nil {
		return nil, false
	}
	return o.Lifetime.Get(), o.Lifetime.IsSet()
}

// SetLifetime sets field value
func (o *SubscriptionPlansMetadataVideoQuota) SetLifetime(v float32) {
	o.Lifetime.Set(&v)
}

// GetMonthly returns the Monthly field value
func (o *SubscriptionPlansMetadataVideoQuota) GetMonthly() float32 {
	if o == nil {
		var ret float32
		return ret
	}

	return o.Monthly
}

// GetMonthlyOk returns a tuple with the Monthly field value
// and a boolean to check if the value has been set.
func (o *SubscriptionPlansMetadataVideoQuota) GetMonthlyOk() (*float32, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Monthly, true
}

// SetMonthly sets field value
func (o *SubscriptionPlansMetadataVideoQuota) SetMonthly(v float32) {
	o.Monthly = v
}

// GetYearly returns the Yearly field value
func (o *SubscriptionPlansMetadataVideoQuota) GetYearly() float32 {
	if o == nil {
		var ret float32
		return ret
	}

	return o.Yearly
}

// GetYearlyOk returns a tuple with the Yearly field value
// and a boolean to check if the value has been set.
func (o *SubscriptionPlansMetadataVideoQuota) GetYearlyOk() (*float32, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Yearly, true
}

// SetYearly sets field value
func (o *SubscriptionPlansMetadataVideoQuota) SetYearly(v float32) {
	o.Yearly = v
}

func (o SubscriptionPlansMetadataVideoQuota) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o SubscriptionPlansMetadataVideoQuota) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["lifetime"] = o.Lifetime.Get()
	toSerialize["monthly"] = o.Monthly
	toSerialize["yearly"] = o.Yearly
	return toSerialize, nil
}

type NullableSubscriptionPlansMetadataVideoQuota struct {
	value *SubscriptionPlansMetadataVideoQuota
	isSet bool
}

func (v NullableSubscriptionPlansMetadataVideoQuota) Get() *SubscriptionPlansMetadataVideoQuota {
	return v.value
}

func (v *NullableSubscriptionPlansMetadataVideoQuota) Set(val *SubscriptionPlansMetadataVideoQuota) {
	v.value = val
	v.isSet = true
}

func (v NullableSubscriptionPlansMetadataVideoQuota) IsSet() bool {
	return v.isSet
}

func (v *NullableSubscriptionPlansMetadataVideoQuota) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableSubscriptionPlansMetadataVideoQuota(val *SubscriptionPlansMetadataVideoQuota) *NullableSubscriptionPlansMetadataVideoQuota {
	return &NullableSubscriptionPlansMetadataVideoQuota{value: val, isSet: true}
}

func (v NullableSubscriptionPlansMetadataVideoQuota) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableSubscriptionPlansMetadataVideoQuota) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


