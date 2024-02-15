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

// checks if the OnDemandVideoRent type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &OnDemandVideoRent{}

// OnDemandVideoRent Information about renting the video.
type OnDemandVideoRent struct {
	// Whether the video can be rented.
	Active bool `json:"active"`
	// A map of currency type to price.
	Price map[string]interface{} `json:"price"`
	// Whether the video has been rented.
	Purchased *bool `json:"purchased,omitempty"`
}

type _OnDemandVideoRent OnDemandVideoRent

// NewOnDemandVideoRent instantiates a new OnDemandVideoRent object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewOnDemandVideoRent(active bool, price map[string]interface{}) *OnDemandVideoRent {
	this := OnDemandVideoRent{}
	this.Active = active
	this.Price = price
	return &this
}

// NewOnDemandVideoRentWithDefaults instantiates a new OnDemandVideoRent object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewOnDemandVideoRentWithDefaults() *OnDemandVideoRent {
	this := OnDemandVideoRent{}
	return &this
}

// GetActive returns the Active field value
func (o *OnDemandVideoRent) GetActive() bool {
	if o == nil {
		var ret bool
		return ret
	}

	return o.Active
}

// GetActiveOk returns a tuple with the Active field value
// and a boolean to check if the value has been set.
func (o *OnDemandVideoRent) GetActiveOk() (*bool, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Active, true
}

// SetActive sets field value
func (o *OnDemandVideoRent) SetActive(v bool) {
	o.Active = v
}

// GetPrice returns the Price field value
func (o *OnDemandVideoRent) GetPrice() map[string]interface{} {
	if o == nil {
		var ret map[string]interface{}
		return ret
	}

	return o.Price
}

// GetPriceOk returns a tuple with the Price field value
// and a boolean to check if the value has been set.
func (o *OnDemandVideoRent) GetPriceOk() (map[string]interface{}, bool) {
	if o == nil {
		return map[string]interface{}{}, false
	}
	return o.Price, true
}

// SetPrice sets field value
func (o *OnDemandVideoRent) SetPrice(v map[string]interface{}) {
	o.Price = v
}

// GetPurchased returns the Purchased field value if set, zero value otherwise.
func (o *OnDemandVideoRent) GetPurchased() bool {
	if o == nil || IsNil(o.Purchased) {
		var ret bool
		return ret
	}
	return *o.Purchased
}

// GetPurchasedOk returns a tuple with the Purchased field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *OnDemandVideoRent) GetPurchasedOk() (*bool, bool) {
	if o == nil || IsNil(o.Purchased) {
		return nil, false
	}
	return o.Purchased, true
}

// HasPurchased returns a boolean if a field has been set.
func (o *OnDemandVideoRent) HasPurchased() bool {
	if o != nil && !IsNil(o.Purchased) {
		return true
	}

	return false
}

// SetPurchased gets a reference to the given bool and assigns it to the Purchased field.
func (o *OnDemandVideoRent) SetPurchased(v bool) {
	o.Purchased = &v
}

func (o OnDemandVideoRent) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o OnDemandVideoRent) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["active"] = o.Active
	toSerialize["price"] = o.Price
	if !IsNil(o.Purchased) {
		toSerialize["purchased"] = o.Purchased
	}
	return toSerialize, nil
}

func (o *OnDemandVideoRent) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"active",
		"price",
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

	varOnDemandVideoRent := _OnDemandVideoRent{}

	decoder := json.NewDecoder(bytes.NewReader(data))
	decoder.DisallowUnknownFields()
	err = decoder.Decode(&varOnDemandVideoRent)

	if err != nil {
		return err
	}

	*o = OnDemandVideoRent(varOnDemandVideoRent)

	return err
}

type NullableOnDemandVideoRent struct {
	value *OnDemandVideoRent
	isSet bool
}

func (v NullableOnDemandVideoRent) Get() *OnDemandVideoRent {
	return v.value
}

func (v *NullableOnDemandVideoRent) Set(val *OnDemandVideoRent) {
	v.value = val
	v.isSet = true
}

func (v NullableOnDemandVideoRent) IsSet() bool {
	return v.isSet
}

func (v *NullableOnDemandVideoRent) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableOnDemandVideoRent(val *OnDemandVideoRent) *NullableOnDemandVideoRent {
	return &NullableOnDemandVideoRent{value: val, isSet: true}
}

func (v NullableOnDemandVideoRent) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableOnDemandVideoRent) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

