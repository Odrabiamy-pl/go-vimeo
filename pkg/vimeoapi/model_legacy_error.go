/*
Vimeo API

Build something great. Vimeo's API supports flexible, high-quality video integration with your custom apps.

API version: 3.4
*/

package vimeoapi

import (
	"encoding/json"
)

// checks if the LegacyError type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &LegacyError{}

// LegacyError struct for LegacyError
type LegacyError struct {
	// The error message.
	Error string `json:"error"`
}

// NewLegacyError instantiates a new LegacyError object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewLegacyError(error_ string) *LegacyError {
	this := LegacyError{}
	this.Error = error_
	return &this
}

// NewLegacyErrorWithDefaults instantiates a new LegacyError object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewLegacyErrorWithDefaults() *LegacyError {
	this := LegacyError{}
	return &this
}

// GetError returns the Error field value
func (o *LegacyError) GetError() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Error
}

// GetErrorOk returns a tuple with the Error field value
// and a boolean to check if the value has been set.
func (o *LegacyError) GetErrorOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Error, true
}

// SetError sets field value
func (o *LegacyError) SetError(v string) {
	o.Error = v
}

func (o LegacyError) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o LegacyError) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["error"] = o.Error
	return toSerialize, nil
}

type NullableLegacyError struct {
	value *LegacyError
	isSet bool
}

func (v NullableLegacyError) Get() *LegacyError {
	return v.value
}

func (v *NullableLegacyError) Set(val *LegacyError) {
	v.value = val
	v.isSet = true
}

func (v NullableLegacyError) IsSet() bool {
	return v.isSet
}

func (v *NullableLegacyError) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableLegacyError(val *LegacyError) *NullableLegacyError {
	return &NullableLegacyError{value: val, isSet: true}
}

func (v NullableLegacyError) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableLegacyError) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
