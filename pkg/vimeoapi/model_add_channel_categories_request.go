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

// checks if the AddChannelCategoriesRequest type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &AddChannelCategoriesRequest{}

// AddChannelCategoriesRequest struct for AddChannelCategoriesRequest
type AddChannelCategoriesRequest struct {
	// The array of category URIs to add.
	Channels []string `json:"channels"`
}

type _AddChannelCategoriesRequest AddChannelCategoriesRequest

// NewAddChannelCategoriesRequest instantiates a new AddChannelCategoriesRequest object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewAddChannelCategoriesRequest(channels []string) *AddChannelCategoriesRequest {
	this := AddChannelCategoriesRequest{}
	this.Channels = channels
	return &this
}

// NewAddChannelCategoriesRequestWithDefaults instantiates a new AddChannelCategoriesRequest object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewAddChannelCategoriesRequestWithDefaults() *AddChannelCategoriesRequest {
	this := AddChannelCategoriesRequest{}
	return &this
}

// GetChannels returns the Channels field value
func (o *AddChannelCategoriesRequest) GetChannels() []string {
	if o == nil {
		var ret []string
		return ret
	}

	return o.Channels
}

// GetChannelsOk returns a tuple with the Channels field value
// and a boolean to check if the value has been set.
func (o *AddChannelCategoriesRequest) GetChannelsOk() ([]string, bool) {
	if o == nil {
		return nil, false
	}
	return o.Channels, true
}

// SetChannels sets field value
func (o *AddChannelCategoriesRequest) SetChannels(v []string) {
	o.Channels = v
}

func (o AddChannelCategoriesRequest) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o AddChannelCategoriesRequest) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["channels"] = o.Channels
	return toSerialize, nil
}

func (o *AddChannelCategoriesRequest) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"channels",
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

	varAddChannelCategoriesRequest := _AddChannelCategoriesRequest{}

	decoder := json.NewDecoder(bytes.NewReader(data))
	decoder.DisallowUnknownFields()
	err = decoder.Decode(&varAddChannelCategoriesRequest)

	if err != nil {
		return err
	}

	*o = AddChannelCategoriesRequest(varAddChannelCategoriesRequest)

	return err
}

type NullableAddChannelCategoriesRequest struct {
	value *AddChannelCategoriesRequest
	isSet bool
}

func (v NullableAddChannelCategoriesRequest) Get() *AddChannelCategoriesRequest {
	return v.value
}

func (v *NullableAddChannelCategoriesRequest) Set(val *AddChannelCategoriesRequest) {
	v.value = val
	v.isSet = true
}

func (v NullableAddChannelCategoriesRequest) IsSet() bool {
	return v.isSet
}

func (v *NullableAddChannelCategoriesRequest) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableAddChannelCategoriesRequest(val *AddChannelCategoriesRequest) *NullableAddChannelCategoriesRequest {
	return &NullableAddChannelCategoriesRequest{value: val, isSet: true}
}

func (v NullableAddChannelCategoriesRequest) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableAddChannelCategoriesRequest) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

