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

// checks if the OnDemandPageSubscription type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &OnDemandPageSubscription{}

// OnDemandPageSubscription Information about subscribing to the On Demand page, if subscription is enabled.
type OnDemandPageSubscription struct {
	// Whether the On Demand product is active.
	Active bool `json:"active"`
	// The link to the On Demand product.
	Link NullableString `json:"link"`
	// The On Demand product's rental period.
	Period *string `json:"period,omitempty"`
	// The accepted currencies and respective pricing for the On Demand product.
	Price map[string]interface{} `json:"price"`
}

// NewOnDemandPageSubscription instantiates a new OnDemandPageSubscription object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewOnDemandPageSubscription(active bool, link NullableString, price map[string]interface{}) *OnDemandPageSubscription {
	this := OnDemandPageSubscription{}
	this.Active = active
	this.Link = link
	this.Price = price
	return &this
}

// NewOnDemandPageSubscriptionWithDefaults instantiates a new OnDemandPageSubscription object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewOnDemandPageSubscriptionWithDefaults() *OnDemandPageSubscription {
	this := OnDemandPageSubscription{}
	return &this
}

// GetActive returns the Active field value
func (o *OnDemandPageSubscription) GetActive() bool {
	if o == nil {
		var ret bool
		return ret
	}

	return o.Active
}

// GetActiveOk returns a tuple with the Active field value
// and a boolean to check if the value has been set.
func (o *OnDemandPageSubscription) GetActiveOk() (*bool, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Active, true
}

// SetActive sets field value
func (o *OnDemandPageSubscription) SetActive(v bool) {
	o.Active = v
}

// GetLink returns the Link field value
// If the value is explicit nil, the zero value for string will be returned
func (o *OnDemandPageSubscription) GetLink() string {
	if o == nil || o.Link.Get() == nil {
		var ret string
		return ret
	}

	return *o.Link.Get()
}

// GetLinkOk returns a tuple with the Link field value
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *OnDemandPageSubscription) GetLinkOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.Link.Get(), o.Link.IsSet()
}

// SetLink sets field value
func (o *OnDemandPageSubscription) SetLink(v string) {
	o.Link.Set(&v)
}

// GetPeriod returns the Period field value if set, zero value otherwise.
func (o *OnDemandPageSubscription) GetPeriod() string {
	if o == nil || IsNil(o.Period) {
		var ret string
		return ret
	}
	return *o.Period
}

// GetPeriodOk returns a tuple with the Period field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *OnDemandPageSubscription) GetPeriodOk() (*string, bool) {
	if o == nil || IsNil(o.Period) {
		return nil, false
	}
	return o.Period, true
}

// HasPeriod returns a boolean if a field has been set.
func (o *OnDemandPageSubscription) HasPeriod() bool {
	if o != nil && !IsNil(o.Period) {
		return true
	}

	return false
}

// SetPeriod gets a reference to the given string and assigns it to the Period field.
func (o *OnDemandPageSubscription) SetPeriod(v string) {
	o.Period = &v
}

// GetPrice returns the Price field value
func (o *OnDemandPageSubscription) GetPrice() map[string]interface{} {
	if o == nil {
		var ret map[string]interface{}
		return ret
	}

	return o.Price
}

// GetPriceOk returns a tuple with the Price field value
// and a boolean to check if the value has been set.
func (o *OnDemandPageSubscription) GetPriceOk() (map[string]interface{}, bool) {
	if o == nil {
		return map[string]interface{}{}, false
	}
	return o.Price, true
}

// SetPrice sets field value
func (o *OnDemandPageSubscription) SetPrice(v map[string]interface{}) {
	o.Price = v
}

func (o OnDemandPageSubscription) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o OnDemandPageSubscription) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["active"] = o.Active
	toSerialize["link"] = o.Link.Get()
	if !IsNil(o.Period) {
		toSerialize["period"] = o.Period
	}
	toSerialize["price"] = o.Price
	return toSerialize, nil
}

type NullableOnDemandPageSubscription struct {
	value *OnDemandPageSubscription
	isSet bool
}

func (v NullableOnDemandPageSubscription) Get() *OnDemandPageSubscription {
	return v.value
}

func (v *NullableOnDemandPageSubscription) Set(val *OnDemandPageSubscription) {
	v.value = val
	v.isSet = true
}

func (v NullableOnDemandPageSubscription) IsSet() bool {
	return v.isSet
}

func (v *NullableOnDemandPageSubscription) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableOnDemandPageSubscription(val *OnDemandPageSubscription) *NullableOnDemandPageSubscription {
	return &NullableOnDemandPageSubscription{value: val, isSet: true}
}

func (v NullableOnDemandPageSubscription) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableOnDemandPageSubscription) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
