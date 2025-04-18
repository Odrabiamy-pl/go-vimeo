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

// checks if the OnDemandVideoMetadataInteractions type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &OnDemandVideoMetadataInteractions{}

// OnDemandVideoMetadataInteractions struct for OnDemandVideoMetadataInteractions
type OnDemandVideoMetadataInteractions struct {
	Likes OnDemandVideoMetadataInteractionsLikes `json:"likes"`
	Watchlater OnDemandVideoMetadataInteractionsWatchlater `json:"watchlater"`
}

// NewOnDemandVideoMetadataInteractions instantiates a new OnDemandVideoMetadataInteractions object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewOnDemandVideoMetadataInteractions(likes OnDemandVideoMetadataInteractionsLikes, watchlater OnDemandVideoMetadataInteractionsWatchlater) *OnDemandVideoMetadataInteractions {
	this := OnDemandVideoMetadataInteractions{}
	this.Likes = likes
	this.Watchlater = watchlater
	return &this
}

// NewOnDemandVideoMetadataInteractionsWithDefaults instantiates a new OnDemandVideoMetadataInteractions object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewOnDemandVideoMetadataInteractionsWithDefaults() *OnDemandVideoMetadataInteractions {
	this := OnDemandVideoMetadataInteractions{}
	return &this
}

// GetLikes returns the Likes field value
func (o *OnDemandVideoMetadataInteractions) GetLikes() OnDemandVideoMetadataInteractionsLikes {
	if o == nil {
		var ret OnDemandVideoMetadataInteractionsLikes
		return ret
	}

	return o.Likes
}

// GetLikesOk returns a tuple with the Likes field value
// and a boolean to check if the value has been set.
func (o *OnDemandVideoMetadataInteractions) GetLikesOk() (*OnDemandVideoMetadataInteractionsLikes, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Likes, true
}

// SetLikes sets field value
func (o *OnDemandVideoMetadataInteractions) SetLikes(v OnDemandVideoMetadataInteractionsLikes) {
	o.Likes = v
}

// GetWatchlater returns the Watchlater field value
func (o *OnDemandVideoMetadataInteractions) GetWatchlater() OnDemandVideoMetadataInteractionsWatchlater {
	if o == nil {
		var ret OnDemandVideoMetadataInteractionsWatchlater
		return ret
	}

	return o.Watchlater
}

// GetWatchlaterOk returns a tuple with the Watchlater field value
// and a boolean to check if the value has been set.
func (o *OnDemandVideoMetadataInteractions) GetWatchlaterOk() (*OnDemandVideoMetadataInteractionsWatchlater, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Watchlater, true
}

// SetWatchlater sets field value
func (o *OnDemandVideoMetadataInteractions) SetWatchlater(v OnDemandVideoMetadataInteractionsWatchlater) {
	o.Watchlater = v
}

func (o OnDemandVideoMetadataInteractions) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o OnDemandVideoMetadataInteractions) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["likes"] = o.Likes
	toSerialize["watchlater"] = o.Watchlater
	return toSerialize, nil
}

type NullableOnDemandVideoMetadataInteractions struct {
	value *OnDemandVideoMetadataInteractions
	isSet bool
}

func (v NullableOnDemandVideoMetadataInteractions) Get() *OnDemandVideoMetadataInteractions {
	return v.value
}

func (v *NullableOnDemandVideoMetadataInteractions) Set(val *OnDemandVideoMetadataInteractions) {
	v.value = val
	v.isSet = true
}

func (v NullableOnDemandVideoMetadataInteractions) IsSet() bool {
	return v.isSet
}

func (v *NullableOnDemandVideoMetadataInteractions) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableOnDemandVideoMetadataInteractions(val *OnDemandVideoMetadataInteractions) *NullableOnDemandVideoMetadataInteractions {
	return &NullableOnDemandVideoMetadataInteractions{value: val, isSet: true}
}

func (v NullableOnDemandVideoMetadataInteractions) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableOnDemandVideoMetadataInteractions) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


