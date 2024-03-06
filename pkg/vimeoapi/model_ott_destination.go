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

// checks if the OttDestination type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &OttDestination{}

// OttDestination struct for OttDestination
type OttDestination struct {
	// The OTT destination's canonical relative URI.
	Id string `json:"id"`
	// The ID of the OTT channel.
	OttChannelId float32 `json:"ott_channel_id"`
	// The name of the OTT channel.
	OttChannelName string `json:"ott_channel_name"`
	// The subdomain of the OTT channel.
	OttChannelSubdomain string `json:"ott_channel_subdomain"`
	// The ID of the OTT event.
	OttEventId float32 `json:"ott_event_id"`
	// The ID of the current recurring live event.
	RecurringLiveEventId float32 `json:"recurring_live_event_id"`
}

// NewOttDestination instantiates a new OttDestination object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewOttDestination(id string, ottChannelId float32, ottChannelName string, ottChannelSubdomain string, ottEventId float32, recurringLiveEventId float32) *OttDestination {
	this := OttDestination{}
	this.Id = id
	this.OttChannelId = ottChannelId
	this.OttChannelName = ottChannelName
	this.OttChannelSubdomain = ottChannelSubdomain
	this.OttEventId = ottEventId
	this.RecurringLiveEventId = recurringLiveEventId
	return &this
}

// NewOttDestinationWithDefaults instantiates a new OttDestination object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewOttDestinationWithDefaults() *OttDestination {
	this := OttDestination{}
	return &this
}

// GetId returns the Id field value
func (o *OttDestination) GetId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Id
}

// GetIdOk returns a tuple with the Id field value
// and a boolean to check if the value has been set.
func (o *OttDestination) GetIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Id, true
}

// SetId sets field value
func (o *OttDestination) SetId(v string) {
	o.Id = v
}

// GetOttChannelId returns the OttChannelId field value
func (o *OttDestination) GetOttChannelId() float32 {
	if o == nil {
		var ret float32
		return ret
	}

	return o.OttChannelId
}

// GetOttChannelIdOk returns a tuple with the OttChannelId field value
// and a boolean to check if the value has been set.
func (o *OttDestination) GetOttChannelIdOk() (*float32, bool) {
	if o == nil {
		return nil, false
	}
	return &o.OttChannelId, true
}

// SetOttChannelId sets field value
func (o *OttDestination) SetOttChannelId(v float32) {
	o.OttChannelId = v
}

// GetOttChannelName returns the OttChannelName field value
func (o *OttDestination) GetOttChannelName() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.OttChannelName
}

// GetOttChannelNameOk returns a tuple with the OttChannelName field value
// and a boolean to check if the value has been set.
func (o *OttDestination) GetOttChannelNameOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.OttChannelName, true
}

// SetOttChannelName sets field value
func (o *OttDestination) SetOttChannelName(v string) {
	o.OttChannelName = v
}

// GetOttChannelSubdomain returns the OttChannelSubdomain field value
func (o *OttDestination) GetOttChannelSubdomain() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.OttChannelSubdomain
}

// GetOttChannelSubdomainOk returns a tuple with the OttChannelSubdomain field value
// and a boolean to check if the value has been set.
func (o *OttDestination) GetOttChannelSubdomainOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.OttChannelSubdomain, true
}

// SetOttChannelSubdomain sets field value
func (o *OttDestination) SetOttChannelSubdomain(v string) {
	o.OttChannelSubdomain = v
}

// GetOttEventId returns the OttEventId field value
func (o *OttDestination) GetOttEventId() float32 {
	if o == nil {
		var ret float32
		return ret
	}

	return o.OttEventId
}

// GetOttEventIdOk returns a tuple with the OttEventId field value
// and a boolean to check if the value has been set.
func (o *OttDestination) GetOttEventIdOk() (*float32, bool) {
	if o == nil {
		return nil, false
	}
	return &o.OttEventId, true
}

// SetOttEventId sets field value
func (o *OttDestination) SetOttEventId(v float32) {
	o.OttEventId = v
}

// GetRecurringLiveEventId returns the RecurringLiveEventId field value
func (o *OttDestination) GetRecurringLiveEventId() float32 {
	if o == nil {
		var ret float32
		return ret
	}

	return o.RecurringLiveEventId
}

// GetRecurringLiveEventIdOk returns a tuple with the RecurringLiveEventId field value
// and a boolean to check if the value has been set.
func (o *OttDestination) GetRecurringLiveEventIdOk() (*float32, bool) {
	if o == nil {
		return nil, false
	}
	return &o.RecurringLiveEventId, true
}

// SetRecurringLiveEventId sets field value
func (o *OttDestination) SetRecurringLiveEventId(v float32) {
	o.RecurringLiveEventId = v
}

func (o OttDestination) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o OttDestination) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["id"] = o.Id
	toSerialize["ott_channel_id"] = o.OttChannelId
	toSerialize["ott_channel_name"] = o.OttChannelName
	toSerialize["ott_channel_subdomain"] = o.OttChannelSubdomain
	toSerialize["ott_event_id"] = o.OttEventId
	toSerialize["recurring_live_event_id"] = o.RecurringLiveEventId
	return toSerialize, nil
}

type NullableOttDestination struct {
	value *OttDestination
	isSet bool
}

func (v NullableOttDestination) Get() *OttDestination {
	return v.value
}

func (v *NullableOttDestination) Set(val *OttDestination) {
	v.value = val
	v.isSet = true
}

func (v NullableOttDestination) IsSet() bool {
	return v.isSet
}

func (v *NullableOttDestination) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableOttDestination(val *OttDestination) *NullableOttDestination {
	return &NullableOttDestination{value: val, isSet: true}
}

func (v NullableOttDestination) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableOttDestination) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
