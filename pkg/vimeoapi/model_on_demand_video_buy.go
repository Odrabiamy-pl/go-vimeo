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

// checks if the OnDemandVideoBuy type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &OnDemandVideoBuy{}

// OnDemandVideoBuy Information about purchasing the video.
type OnDemandVideoBuy struct {
	// Whether the video can be purchased.
	Active bool `json:"active"`
	// A map of currency type to price.
	Price map[string]interface{} `json:"price"`
	// Whether the video has been purchased.
	Purchased *bool `json:"purchased,omitempty"`
}

// NewOnDemandVideoBuy instantiates a new OnDemandVideoBuy object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewOnDemandVideoBuy(active bool, price map[string]interface{}) *OnDemandVideoBuy {
	this := OnDemandVideoBuy{}
	this.Active = active
	this.Price = price
	return &this
}

// NewOnDemandVideoBuyWithDefaults instantiates a new OnDemandVideoBuy object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewOnDemandVideoBuyWithDefaults() *OnDemandVideoBuy {
	this := OnDemandVideoBuy{}
	return &this
}

// GetActive returns the Active field value
func (o *OnDemandVideoBuy) GetActive() bool {
	if o == nil {
		var ret bool
		return ret
	}

	return o.Active
}

// GetActiveOk returns a tuple with the Active field value
// and a boolean to check if the value has been set.
func (o *OnDemandVideoBuy) GetActiveOk() (*bool, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Active, true
}

// SetActive sets field value
func (o *OnDemandVideoBuy) SetActive(v bool) {
	o.Active = v
}

// GetPrice returns the Price field value
func (o *OnDemandVideoBuy) GetPrice() map[string]interface{} {
	if o == nil {
		var ret map[string]interface{}
		return ret
	}

	return o.Price
}

// GetPriceOk returns a tuple with the Price field value
// and a boolean to check if the value has been set.
func (o *OnDemandVideoBuy) GetPriceOk() (map[string]interface{}, bool) {
	if o == nil {
		return map[string]interface{}{}, false
	}
	return o.Price, true
}

// SetPrice sets field value
func (o *OnDemandVideoBuy) SetPrice(v map[string]interface{}) {
	o.Price = v
}

// GetPurchased returns the Purchased field value if set, zero value otherwise.
func (o *OnDemandVideoBuy) GetPurchased() bool {
	if o == nil || IsNil(o.Purchased) {
		var ret bool
		return ret
	}
	return *o.Purchased
}

// GetPurchasedOk returns a tuple with the Purchased field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *OnDemandVideoBuy) GetPurchasedOk() (*bool, bool) {
	if o == nil || IsNil(o.Purchased) {
		return nil, false
	}
	return o.Purchased, true
}

// HasPurchased returns a boolean if a field has been set.
func (o *OnDemandVideoBuy) HasPurchased() bool {
	if o != nil && !IsNil(o.Purchased) {
		return true
	}

	return false
}

// SetPurchased gets a reference to the given bool and assigns it to the Purchased field.
func (o *OnDemandVideoBuy) SetPurchased(v bool) {
	o.Purchased = &v
}

func (o OnDemandVideoBuy) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o OnDemandVideoBuy) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["active"] = o.Active
	toSerialize["price"] = o.Price
	if !IsNil(o.Purchased) {
		toSerialize["purchased"] = o.Purchased
	}
	return toSerialize, nil
}

type NullableOnDemandVideoBuy struct {
	value *OnDemandVideoBuy
	isSet bool
}

func (v NullableOnDemandVideoBuy) Get() *OnDemandVideoBuy {
	return v.value
}

func (v *NullableOnDemandVideoBuy) Set(val *OnDemandVideoBuy) {
	v.value = val
	v.isSet = true
}

func (v NullableOnDemandVideoBuy) IsSet() bool {
	return v.isSet
}

func (v *NullableOnDemandVideoBuy) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableOnDemandVideoBuy(val *OnDemandVideoBuy) *NullableOnDemandVideoBuy {
	return &NullableOnDemandVideoBuy{value: val, isSet: true}
}

func (v NullableOnDemandVideoBuy) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableOnDemandVideoBuy) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


