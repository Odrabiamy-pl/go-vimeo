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

// checks if the SetLiveEventWhitelistAlt1Request type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &SetLiveEventWhitelistAlt1Request{}

// SetLiveEventWhitelistAlt1Request struct for SetLiveEventWhitelistAlt1Request
type SetLiveEventWhitelistAlt1Request struct {
	// An array of the domains on which the embedded event can appear.
	AllowedDomains []string `json:"allowed_domains,omitempty"`
}

// NewSetLiveEventWhitelistAlt1Request instantiates a new SetLiveEventWhitelistAlt1Request object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewSetLiveEventWhitelistAlt1Request() *SetLiveEventWhitelistAlt1Request {
	this := SetLiveEventWhitelistAlt1Request{}
	return &this
}

// NewSetLiveEventWhitelistAlt1RequestWithDefaults instantiates a new SetLiveEventWhitelistAlt1Request object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewSetLiveEventWhitelistAlt1RequestWithDefaults() *SetLiveEventWhitelistAlt1Request {
	this := SetLiveEventWhitelistAlt1Request{}
	return &this
}

// GetAllowedDomains returns the AllowedDomains field value if set, zero value otherwise.
func (o *SetLiveEventWhitelistAlt1Request) GetAllowedDomains() []string {
	if o == nil || IsNil(o.AllowedDomains) {
		var ret []string
		return ret
	}
	return o.AllowedDomains
}

// GetAllowedDomainsOk returns a tuple with the AllowedDomains field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SetLiveEventWhitelistAlt1Request) GetAllowedDomainsOk() ([]string, bool) {
	if o == nil || IsNil(o.AllowedDomains) {
		return nil, false
	}
	return o.AllowedDomains, true
}

// HasAllowedDomains returns a boolean if a field has been set.
func (o *SetLiveEventWhitelistAlt1Request) HasAllowedDomains() bool {
	if o != nil && !IsNil(o.AllowedDomains) {
		return true
	}

	return false
}

// SetAllowedDomains gets a reference to the given []string and assigns it to the AllowedDomains field.
func (o *SetLiveEventWhitelistAlt1Request) SetAllowedDomains(v []string) {
	o.AllowedDomains = v
}

func (o SetLiveEventWhitelistAlt1Request) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o SetLiveEventWhitelistAlt1Request) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.AllowedDomains) {
		toSerialize["allowed_domains"] = o.AllowedDomains
	}
	return toSerialize, nil
}

type NullableSetLiveEventWhitelistAlt1Request struct {
	value *SetLiveEventWhitelistAlt1Request
	isSet bool
}

func (v NullableSetLiveEventWhitelistAlt1Request) Get() *SetLiveEventWhitelistAlt1Request {
	return v.value
}

func (v *NullableSetLiveEventWhitelistAlt1Request) Set(val *SetLiveEventWhitelistAlt1Request) {
	v.value = val
	v.isSet = true
}

func (v NullableSetLiveEventWhitelistAlt1Request) IsSet() bool {
	return v.isSet
}

func (v *NullableSetLiveEventWhitelistAlt1Request) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableSetLiveEventWhitelistAlt1Request(val *SetLiveEventWhitelistAlt1Request) *NullableSetLiveEventWhitelistAlt1Request {
	return &NullableSetLiveEventWhitelistAlt1Request{value: val, isSet: true}
}

func (v NullableSetLiveEventWhitelistAlt1Request) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableSetLiveEventWhitelistAlt1Request) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
