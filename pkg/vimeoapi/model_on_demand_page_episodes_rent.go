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

// checks if the OnDemandPageEpisodesRent type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &OnDemandPageEpisodesRent{}

// OnDemandPageEpisodesRent struct for OnDemandPageEpisodesRent
type OnDemandPageEpisodesRent struct {
	// Whether all the videos on the On Demand page can be rented as a whole.
	Active bool `json:"active"`
	// The rental period for the video.  Option descriptions:  * `1 day` - The rental period is one day.  * `1 month` - The rental period is one month.  * `1 week` - The rental period is one week.  * `1 year` - The rental period is one year.  * `2 day` - The rental period is two days.  * `24 hour` - The rental period is 24 hours.  * `3 day` - The rental period is three days.  * `3 month` - The rental period is three months.  * `30 day` - The rental period is 30 days.  * `48 hour` - The rental period is 48 hours.  * `6 month` - The rental period is six months.  * `60 day` - The rental period is 60 days.  * `7 day` - The rental period is 7 days.  * `72 hour` - The rental period is 72 hours. 
	Period NullableString `json:"period"`
	// The default price to rent an episode.
	Price NullableFloat32 `json:"price"`
}

// NewOnDemandPageEpisodesRent instantiates a new OnDemandPageEpisodesRent object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewOnDemandPageEpisodesRent(active bool, period NullableString, price NullableFloat32) *OnDemandPageEpisodesRent {
	this := OnDemandPageEpisodesRent{}
	this.Active = active
	this.Period = period
	this.Price = price
	return &this
}

// NewOnDemandPageEpisodesRentWithDefaults instantiates a new OnDemandPageEpisodesRent object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewOnDemandPageEpisodesRentWithDefaults() *OnDemandPageEpisodesRent {
	this := OnDemandPageEpisodesRent{}
	return &this
}

// GetActive returns the Active field value
func (o *OnDemandPageEpisodesRent) GetActive() bool {
	if o == nil {
		var ret bool
		return ret
	}

	return o.Active
}

// GetActiveOk returns a tuple with the Active field value
// and a boolean to check if the value has been set.
func (o *OnDemandPageEpisodesRent) GetActiveOk() (*bool, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Active, true
}

// SetActive sets field value
func (o *OnDemandPageEpisodesRent) SetActive(v bool) {
	o.Active = v
}

// GetPeriod returns the Period field value
// If the value is explicit nil, the zero value for string will be returned
func (o *OnDemandPageEpisodesRent) GetPeriod() string {
	if o == nil || o.Period.Get() == nil {
		var ret string
		return ret
	}

	return *o.Period.Get()
}

// GetPeriodOk returns a tuple with the Period field value
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *OnDemandPageEpisodesRent) GetPeriodOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.Period.Get(), o.Period.IsSet()
}

// SetPeriod sets field value
func (o *OnDemandPageEpisodesRent) SetPeriod(v string) {
	o.Period.Set(&v)
}

// GetPrice returns the Price field value
// If the value is explicit nil, the zero value for float32 will be returned
func (o *OnDemandPageEpisodesRent) GetPrice() float32 {
	if o == nil || o.Price.Get() == nil {
		var ret float32
		return ret
	}

	return *o.Price.Get()
}

// GetPriceOk returns a tuple with the Price field value
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *OnDemandPageEpisodesRent) GetPriceOk() (*float32, bool) {
	if o == nil {
		return nil, false
	}
	return o.Price.Get(), o.Price.IsSet()
}

// SetPrice sets field value
func (o *OnDemandPageEpisodesRent) SetPrice(v float32) {
	o.Price.Set(&v)
}

func (o OnDemandPageEpisodesRent) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o OnDemandPageEpisodesRent) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["active"] = o.Active
	toSerialize["period"] = o.Period.Get()
	toSerialize["price"] = o.Price.Get()
	return toSerialize, nil
}

type NullableOnDemandPageEpisodesRent struct {
	value *OnDemandPageEpisodesRent
	isSet bool
}

func (v NullableOnDemandPageEpisodesRent) Get() *OnDemandPageEpisodesRent {
	return v.value
}

func (v *NullableOnDemandPageEpisodesRent) Set(val *OnDemandPageEpisodesRent) {
	v.value = val
	v.isSet = true
}

func (v NullableOnDemandPageEpisodesRent) IsSet() bool {
	return v.isSet
}

func (v *NullableOnDemandPageEpisodesRent) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableOnDemandPageEpisodesRent(val *OnDemandPageEpisodesRent) *NullableOnDemandPageEpisodesRent {
	return &NullableOnDemandPageEpisodesRent{value: val, isSet: true}
}

func (v NullableOnDemandPageEpisodesRent) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableOnDemandPageEpisodesRent) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


