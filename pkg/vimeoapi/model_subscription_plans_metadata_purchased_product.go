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

// checks if the SubscriptionPlansMetadataPurchasedProduct type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &SubscriptionPlansMetadataPurchasedProduct{}

// SubscriptionPlansMetadataPurchasedProduct Information about the purchased product.
type SubscriptionPlansMetadataPurchasedProduct struct {
	// The display price of the purchased product.
	DisplayPrice float32 `json:"display_price"`
	// Whether the purchased product is billed as a monthly subscription.
	IsMonthly bool `json:"is_monthly"`
}

// NewSubscriptionPlansMetadataPurchasedProduct instantiates a new SubscriptionPlansMetadataPurchasedProduct object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewSubscriptionPlansMetadataPurchasedProduct(displayPrice float32, isMonthly bool) *SubscriptionPlansMetadataPurchasedProduct {
	this := SubscriptionPlansMetadataPurchasedProduct{}
	this.DisplayPrice = displayPrice
	this.IsMonthly = isMonthly
	return &this
}

// NewSubscriptionPlansMetadataPurchasedProductWithDefaults instantiates a new SubscriptionPlansMetadataPurchasedProduct object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewSubscriptionPlansMetadataPurchasedProductWithDefaults() *SubscriptionPlansMetadataPurchasedProduct {
	this := SubscriptionPlansMetadataPurchasedProduct{}
	return &this
}

// GetDisplayPrice returns the DisplayPrice field value
func (o *SubscriptionPlansMetadataPurchasedProduct) GetDisplayPrice() float32 {
	if o == nil {
		var ret float32
		return ret
	}

	return o.DisplayPrice
}

// GetDisplayPriceOk returns a tuple with the DisplayPrice field value
// and a boolean to check if the value has been set.
func (o *SubscriptionPlansMetadataPurchasedProduct) GetDisplayPriceOk() (*float32, bool) {
	if o == nil {
		return nil, false
	}
	return &o.DisplayPrice, true
}

// SetDisplayPrice sets field value
func (o *SubscriptionPlansMetadataPurchasedProduct) SetDisplayPrice(v float32) {
	o.DisplayPrice = v
}

// GetIsMonthly returns the IsMonthly field value
func (o *SubscriptionPlansMetadataPurchasedProduct) GetIsMonthly() bool {
	if o == nil {
		var ret bool
		return ret
	}

	return o.IsMonthly
}

// GetIsMonthlyOk returns a tuple with the IsMonthly field value
// and a boolean to check if the value has been set.
func (o *SubscriptionPlansMetadataPurchasedProduct) GetIsMonthlyOk() (*bool, bool) {
	if o == nil {
		return nil, false
	}
	return &o.IsMonthly, true
}

// SetIsMonthly sets field value
func (o *SubscriptionPlansMetadataPurchasedProduct) SetIsMonthly(v bool) {
	o.IsMonthly = v
}

func (o SubscriptionPlansMetadataPurchasedProduct) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o SubscriptionPlansMetadataPurchasedProduct) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["display_price"] = o.DisplayPrice
	toSerialize["is_monthly"] = o.IsMonthly
	return toSerialize, nil
}

type NullableSubscriptionPlansMetadataPurchasedProduct struct {
	value *SubscriptionPlansMetadataPurchasedProduct
	isSet bool
}

func (v NullableSubscriptionPlansMetadataPurchasedProduct) Get() *SubscriptionPlansMetadataPurchasedProduct {
	return v.value
}

func (v *NullableSubscriptionPlansMetadataPurchasedProduct) Set(val *SubscriptionPlansMetadataPurchasedProduct) {
	v.value = val
	v.isSet = true
}

func (v NullableSubscriptionPlansMetadataPurchasedProduct) IsSet() bool {
	return v.isSet
}

func (v *NullableSubscriptionPlansMetadataPurchasedProduct) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableSubscriptionPlansMetadataPurchasedProduct(val *SubscriptionPlansMetadataPurchasedProduct) *NullableSubscriptionPlansMetadataPurchasedProduct {
	return &NullableSubscriptionPlansMetadataPurchasedProduct{value: val, isSet: true}
}

func (v NullableSubscriptionPlansMetadataPurchasedProduct) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableSubscriptionPlansMetadataPurchasedProduct) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
