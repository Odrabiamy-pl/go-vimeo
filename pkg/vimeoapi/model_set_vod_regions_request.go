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

// checks if the SetVodRegionsRequest type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &SetVodRegionsRequest{}

// SetVodRegionsRequest struct for SetVodRegionsRequest
type SetVodRegionsRequest struct {
	// An array of country codes for the regions to add.
	Countries []string `json:"countries"`
}

// NewSetVodRegionsRequest instantiates a new SetVodRegionsRequest object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewSetVodRegionsRequest(countries []string) *SetVodRegionsRequest {
	this := SetVodRegionsRequest{}
	this.Countries = countries
	return &this
}

// NewSetVodRegionsRequestWithDefaults instantiates a new SetVodRegionsRequest object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewSetVodRegionsRequestWithDefaults() *SetVodRegionsRequest {
	this := SetVodRegionsRequest{}
	return &this
}

// GetCountries returns the Countries field value
func (o *SetVodRegionsRequest) GetCountries() []string {
	if o == nil {
		var ret []string
		return ret
	}

	return o.Countries
}

// GetCountriesOk returns a tuple with the Countries field value
// and a boolean to check if the value has been set.
func (o *SetVodRegionsRequest) GetCountriesOk() ([]string, bool) {
	if o == nil {
		return nil, false
	}
	return o.Countries, true
}

// SetCountries sets field value
func (o *SetVodRegionsRequest) SetCountries(v []string) {
	o.Countries = v
}

func (o SetVodRegionsRequest) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o SetVodRegionsRequest) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["countries"] = o.Countries
	return toSerialize, nil
}

type NullableSetVodRegionsRequest struct {
	value *SetVodRegionsRequest
	isSet bool
}

func (v NullableSetVodRegionsRequest) Get() *SetVodRegionsRequest {
	return v.value
}

func (v *NullableSetVodRegionsRequest) Set(val *SetVodRegionsRequest) {
	v.value = val
	v.isSet = true
}

func (v NullableSetVodRegionsRequest) IsSet() bool {
	return v.isSet
}

func (v *NullableSetVodRegionsRequest) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableSetVodRegionsRequest(val *SetVodRegionsRequest) *NullableSetVodRegionsRequest {
	return &NullableSetVodRegionsRequest{value: val, isSet: true}
}

func (v NullableSetVodRegionsRequest) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableSetVodRegionsRequest) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


