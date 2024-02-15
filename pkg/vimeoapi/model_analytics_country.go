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

// checks if the AnalyticsCountry type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &AnalyticsCountry{}

// AnalyticsCountry struct for AnalyticsCountry
type AnalyticsCountry struct {
	// The country code in ISO-3166 format.
	Code string `json:"code"`
	// The name of the country.
	Name string `json:"name"`
}

type _AnalyticsCountry AnalyticsCountry

// NewAnalyticsCountry instantiates a new AnalyticsCountry object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewAnalyticsCountry(code string, name string) *AnalyticsCountry {
	this := AnalyticsCountry{}
	this.Code = code
	this.Name = name
	return &this
}

// NewAnalyticsCountryWithDefaults instantiates a new AnalyticsCountry object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewAnalyticsCountryWithDefaults() *AnalyticsCountry {
	this := AnalyticsCountry{}
	return &this
}

// GetCode returns the Code field value
func (o *AnalyticsCountry) GetCode() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Code
}

// GetCodeOk returns a tuple with the Code field value
// and a boolean to check if the value has been set.
func (o *AnalyticsCountry) GetCodeOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Code, true
}

// SetCode sets field value
func (o *AnalyticsCountry) SetCode(v string) {
	o.Code = v
}

// GetName returns the Name field value
func (o *AnalyticsCountry) GetName() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Name
}

// GetNameOk returns a tuple with the Name field value
// and a boolean to check if the value has been set.
func (o *AnalyticsCountry) GetNameOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Name, true
}

// SetName sets field value
func (o *AnalyticsCountry) SetName(v string) {
	o.Name = v
}

func (o AnalyticsCountry) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o AnalyticsCountry) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["code"] = o.Code
	toSerialize["name"] = o.Name
	return toSerialize, nil
}

func (o *AnalyticsCountry) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"code",
		"name",
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

	varAnalyticsCountry := _AnalyticsCountry{}

	decoder := json.NewDecoder(bytes.NewReader(data))
	decoder.DisallowUnknownFields()
	err = decoder.Decode(&varAnalyticsCountry)

	if err != nil {
		return err
	}

	*o = AnalyticsCountry(varAnalyticsCountry)

	return err
}

type NullableAnalyticsCountry struct {
	value *AnalyticsCountry
	isSet bool
}

func (v NullableAnalyticsCountry) Get() *AnalyticsCountry {
	return v.value
}

func (v *NullableAnalyticsCountry) Set(val *AnalyticsCountry) {
	v.value = val
	v.isSet = true
}

func (v NullableAnalyticsCountry) IsSet() bool {
	return v.isSet
}

func (v *NullableAnalyticsCountry) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableAnalyticsCountry(val *AnalyticsCountry) *NullableAnalyticsCountry {
	return &NullableAnalyticsCountry{value: val, isSet: true}
}

func (v NullableAnalyticsCountry) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableAnalyticsCountry) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


