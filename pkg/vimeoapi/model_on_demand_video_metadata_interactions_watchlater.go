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

// checks if the OnDemandVideoMetadataInteractionsWatchlater type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &OnDemandVideoMetadataInteractionsWatchlater{}

// OnDemandVideoMetadataInteractionsWatchlater Information about the authenticated user's interaction to watch the video later.
type OnDemandVideoMetadataInteractionsWatchlater struct {
	// Whether the authenticated user has added the video to their Watch Later queue.
	Added bool `json:"added"`
	// The time in ISO 8601 format when the authenticated user added the video to their Watch Later queue.
	AddedTime string `json:"added_time"`
	// The URI for the authenticated user to add the video to their Watch Later queue.
	Uri string `json:"uri"`
}

type _OnDemandVideoMetadataInteractionsWatchlater OnDemandVideoMetadataInteractionsWatchlater

// NewOnDemandVideoMetadataInteractionsWatchlater instantiates a new OnDemandVideoMetadataInteractionsWatchlater object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewOnDemandVideoMetadataInteractionsWatchlater(added bool, addedTime string, uri string) *OnDemandVideoMetadataInteractionsWatchlater {
	this := OnDemandVideoMetadataInteractionsWatchlater{}
	this.Added = added
	this.AddedTime = addedTime
	this.Uri = uri
	return &this
}

// NewOnDemandVideoMetadataInteractionsWatchlaterWithDefaults instantiates a new OnDemandVideoMetadataInteractionsWatchlater object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewOnDemandVideoMetadataInteractionsWatchlaterWithDefaults() *OnDemandVideoMetadataInteractionsWatchlater {
	this := OnDemandVideoMetadataInteractionsWatchlater{}
	return &this
}

// GetAdded returns the Added field value
func (o *OnDemandVideoMetadataInteractionsWatchlater) GetAdded() bool {
	if o == nil {
		var ret bool
		return ret
	}

	return o.Added
}

// GetAddedOk returns a tuple with the Added field value
// and a boolean to check if the value has been set.
func (o *OnDemandVideoMetadataInteractionsWatchlater) GetAddedOk() (*bool, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Added, true
}

// SetAdded sets field value
func (o *OnDemandVideoMetadataInteractionsWatchlater) SetAdded(v bool) {
	o.Added = v
}

// GetAddedTime returns the AddedTime field value
func (o *OnDemandVideoMetadataInteractionsWatchlater) GetAddedTime() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.AddedTime
}

// GetAddedTimeOk returns a tuple with the AddedTime field value
// and a boolean to check if the value has been set.
func (o *OnDemandVideoMetadataInteractionsWatchlater) GetAddedTimeOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.AddedTime, true
}

// SetAddedTime sets field value
func (o *OnDemandVideoMetadataInteractionsWatchlater) SetAddedTime(v string) {
	o.AddedTime = v
}

// GetUri returns the Uri field value
func (o *OnDemandVideoMetadataInteractionsWatchlater) GetUri() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Uri
}

// GetUriOk returns a tuple with the Uri field value
// and a boolean to check if the value has been set.
func (o *OnDemandVideoMetadataInteractionsWatchlater) GetUriOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Uri, true
}

// SetUri sets field value
func (o *OnDemandVideoMetadataInteractionsWatchlater) SetUri(v string) {
	o.Uri = v
}

func (o OnDemandVideoMetadataInteractionsWatchlater) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o OnDemandVideoMetadataInteractionsWatchlater) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["added"] = o.Added
	toSerialize["added_time"] = o.AddedTime
	toSerialize["uri"] = o.Uri
	return toSerialize, nil
}

func (o *OnDemandVideoMetadataInteractionsWatchlater) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"added",
		"added_time",
		"uri",
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

	varOnDemandVideoMetadataInteractionsWatchlater := _OnDemandVideoMetadataInteractionsWatchlater{}

	decoder := json.NewDecoder(bytes.NewReader(data))
	decoder.DisallowUnknownFields()
	err = decoder.Decode(&varOnDemandVideoMetadataInteractionsWatchlater)

	if err != nil {
		return err
	}

	*o = OnDemandVideoMetadataInteractionsWatchlater(varOnDemandVideoMetadataInteractionsWatchlater)

	return err
}

type NullableOnDemandVideoMetadataInteractionsWatchlater struct {
	value *OnDemandVideoMetadataInteractionsWatchlater
	isSet bool
}

func (v NullableOnDemandVideoMetadataInteractionsWatchlater) Get() *OnDemandVideoMetadataInteractionsWatchlater {
	return v.value
}

func (v *NullableOnDemandVideoMetadataInteractionsWatchlater) Set(val *OnDemandVideoMetadataInteractionsWatchlater) {
	v.value = val
	v.isSet = true
}

func (v NullableOnDemandVideoMetadataInteractionsWatchlater) IsSet() bool {
	return v.isSet
}

func (v *NullableOnDemandVideoMetadataInteractionsWatchlater) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableOnDemandVideoMetadataInteractionsWatchlater(val *OnDemandVideoMetadataInteractionsWatchlater) *NullableOnDemandVideoMetadataInteractionsWatchlater {
	return &NullableOnDemandVideoMetadataInteractionsWatchlater{value: val, isSet: true}
}

func (v NullableOnDemandVideoMetadataInteractionsWatchlater) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableOnDemandVideoMetadataInteractionsWatchlater) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


