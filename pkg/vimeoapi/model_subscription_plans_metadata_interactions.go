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

// checks if the SubscriptionPlansMetadataInteractions type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &SubscriptionPlansMetadataInteractions{}

// SubscriptionPlansMetadataInteractions struct for SubscriptionPlansMetadataInteractions
type SubscriptionPlansMetadataInteractions struct {
	Purchase SubscriptionPlansMetadataInteractionsPurchase `json:"purchase"`
}

// NewSubscriptionPlansMetadataInteractions instantiates a new SubscriptionPlansMetadataInteractions object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewSubscriptionPlansMetadataInteractions(purchase SubscriptionPlansMetadataInteractionsPurchase) *SubscriptionPlansMetadataInteractions {
	this := SubscriptionPlansMetadataInteractions{}
	this.Purchase = purchase
	return &this
}

// NewSubscriptionPlansMetadataInteractionsWithDefaults instantiates a new SubscriptionPlansMetadataInteractions object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewSubscriptionPlansMetadataInteractionsWithDefaults() *SubscriptionPlansMetadataInteractions {
	this := SubscriptionPlansMetadataInteractions{}
	return &this
}

// GetPurchase returns the Purchase field value
func (o *SubscriptionPlansMetadataInteractions) GetPurchase() SubscriptionPlansMetadataInteractionsPurchase {
	if o == nil {
		var ret SubscriptionPlansMetadataInteractionsPurchase
		return ret
	}

	return o.Purchase
}

// GetPurchaseOk returns a tuple with the Purchase field value
// and a boolean to check if the value has been set.
func (o *SubscriptionPlansMetadataInteractions) GetPurchaseOk() (*SubscriptionPlansMetadataInteractionsPurchase, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Purchase, true
}

// SetPurchase sets field value
func (o *SubscriptionPlansMetadataInteractions) SetPurchase(v SubscriptionPlansMetadataInteractionsPurchase) {
	o.Purchase = v
}

func (o SubscriptionPlansMetadataInteractions) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o SubscriptionPlansMetadataInteractions) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["purchase"] = o.Purchase
	return toSerialize, nil
}

type NullableSubscriptionPlansMetadataInteractions struct {
	value *SubscriptionPlansMetadataInteractions
	isSet bool
}

func (v NullableSubscriptionPlansMetadataInteractions) Get() *SubscriptionPlansMetadataInteractions {
	return v.value
}

func (v *NullableSubscriptionPlansMetadataInteractions) Set(val *SubscriptionPlansMetadataInteractions) {
	v.value = val
	v.isSet = true
}

func (v NullableSubscriptionPlansMetadataInteractions) IsSet() bool {
	return v.isSet
}

func (v *NullableSubscriptionPlansMetadataInteractions) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableSubscriptionPlansMetadataInteractions(val *SubscriptionPlansMetadataInteractions) *NullableSubscriptionPlansMetadataInteractions {
	return &NullableSubscriptionPlansMetadataInteractions{value: val, isSet: true}
}

func (v NullableSubscriptionPlansMetadataInteractions) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableSubscriptionPlansMetadataInteractions) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


