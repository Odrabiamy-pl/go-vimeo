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

// checks if the OnDemandPagePreorder type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &OnDemandPagePreorder{}

// OnDemandPagePreorder struct for OnDemandPagePreorder
type OnDemandPagePreorder struct {
	// Whether the On Demand page is available for preorder.
	Active bool `json:"active"`
	// The time in ISO 8601 format when the preorder was cancelled.
	CancelTime string `json:"cancel_time"`
	// The time in ISO 8601 format when the preorder was released to the public.
	PublishTime string `json:"publish_time"`
	// The time in ISO 8601 format when the preorder started.
	Time string `json:"time"`
}

// NewOnDemandPagePreorder instantiates a new OnDemandPagePreorder object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewOnDemandPagePreorder(active bool, cancelTime string, publishTime string, time string) *OnDemandPagePreorder {
	this := OnDemandPagePreorder{}
	this.Active = active
	this.CancelTime = cancelTime
	this.PublishTime = publishTime
	this.Time = time
	return &this
}

// NewOnDemandPagePreorderWithDefaults instantiates a new OnDemandPagePreorder object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewOnDemandPagePreorderWithDefaults() *OnDemandPagePreorder {
	this := OnDemandPagePreorder{}
	return &this
}

// GetActive returns the Active field value
func (o *OnDemandPagePreorder) GetActive() bool {
	if o == nil {
		var ret bool
		return ret
	}

	return o.Active
}

// GetActiveOk returns a tuple with the Active field value
// and a boolean to check if the value has been set.
func (o *OnDemandPagePreorder) GetActiveOk() (*bool, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Active, true
}

// SetActive sets field value
func (o *OnDemandPagePreorder) SetActive(v bool) {
	o.Active = v
}

// GetCancelTime returns the CancelTime field value
func (o *OnDemandPagePreorder) GetCancelTime() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.CancelTime
}

// GetCancelTimeOk returns a tuple with the CancelTime field value
// and a boolean to check if the value has been set.
func (o *OnDemandPagePreorder) GetCancelTimeOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.CancelTime, true
}

// SetCancelTime sets field value
func (o *OnDemandPagePreorder) SetCancelTime(v string) {
	o.CancelTime = v
}

// GetPublishTime returns the PublishTime field value
func (o *OnDemandPagePreorder) GetPublishTime() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.PublishTime
}

// GetPublishTimeOk returns a tuple with the PublishTime field value
// and a boolean to check if the value has been set.
func (o *OnDemandPagePreorder) GetPublishTimeOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.PublishTime, true
}

// SetPublishTime sets field value
func (o *OnDemandPagePreorder) SetPublishTime(v string) {
	o.PublishTime = v
}

// GetTime returns the Time field value
func (o *OnDemandPagePreorder) GetTime() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Time
}

// GetTimeOk returns a tuple with the Time field value
// and a boolean to check if the value has been set.
func (o *OnDemandPagePreorder) GetTimeOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Time, true
}

// SetTime sets field value
func (o *OnDemandPagePreorder) SetTime(v string) {
	o.Time = v
}

func (o OnDemandPagePreorder) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o OnDemandPagePreorder) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["active"] = o.Active
	toSerialize["cancel_time"] = o.CancelTime
	toSerialize["publish_time"] = o.PublishTime
	toSerialize["time"] = o.Time
	return toSerialize, nil
}

type NullableOnDemandPagePreorder struct {
	value *OnDemandPagePreorder
	isSet bool
}

func (v NullableOnDemandPagePreorder) Get() *OnDemandPagePreorder {
	return v.value
}

func (v *NullableOnDemandPagePreorder) Set(val *OnDemandPagePreorder) {
	v.value = val
	v.isSet = true
}

func (v NullableOnDemandPagePreorder) IsSet() bool {
	return v.isSet
}

func (v *NullableOnDemandPagePreorder) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableOnDemandPagePreorder(val *OnDemandPagePreorder) *NullableOnDemandPagePreorder {
	return &NullableOnDemandPagePreorder{value: val, isSet: true}
}

func (v NullableOnDemandPagePreorder) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableOnDemandPagePreorder) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
