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

// checks if the UpdateWebinarRegistrationsAlt1Request type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &UpdateWebinarRegistrationsAlt1Request{}

// UpdateWebinarRegistrationsAlt1Request struct for UpdateWebinarRegistrationsAlt1Request
type UpdateWebinarRegistrationsAlt1Request struct {
	// Whether to block the webinar registrant.
	IsBlocked *bool `json:"is_blocked,omitempty"`
}

// NewUpdateWebinarRegistrationsAlt1Request instantiates a new UpdateWebinarRegistrationsAlt1Request object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewUpdateWebinarRegistrationsAlt1Request() *UpdateWebinarRegistrationsAlt1Request {
	this := UpdateWebinarRegistrationsAlt1Request{}
	return &this
}

// NewUpdateWebinarRegistrationsAlt1RequestWithDefaults instantiates a new UpdateWebinarRegistrationsAlt1Request object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewUpdateWebinarRegistrationsAlt1RequestWithDefaults() *UpdateWebinarRegistrationsAlt1Request {
	this := UpdateWebinarRegistrationsAlt1Request{}
	return &this
}

// GetIsBlocked returns the IsBlocked field value if set, zero value otherwise.
func (o *UpdateWebinarRegistrationsAlt1Request) GetIsBlocked() bool {
	if o == nil || IsNil(o.IsBlocked) {
		var ret bool
		return ret
	}
	return *o.IsBlocked
}

// GetIsBlockedOk returns a tuple with the IsBlocked field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UpdateWebinarRegistrationsAlt1Request) GetIsBlockedOk() (*bool, bool) {
	if o == nil || IsNil(o.IsBlocked) {
		return nil, false
	}
	return o.IsBlocked, true
}

// HasIsBlocked returns a boolean if a field has been set.
func (o *UpdateWebinarRegistrationsAlt1Request) HasIsBlocked() bool {
	if o != nil && !IsNil(o.IsBlocked) {
		return true
	}

	return false
}

// SetIsBlocked gets a reference to the given bool and assigns it to the IsBlocked field.
func (o *UpdateWebinarRegistrationsAlt1Request) SetIsBlocked(v bool) {
	o.IsBlocked = &v
}

func (o UpdateWebinarRegistrationsAlt1Request) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o UpdateWebinarRegistrationsAlt1Request) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.IsBlocked) {
		toSerialize["is_blocked"] = o.IsBlocked
	}
	return toSerialize, nil
}

type NullableUpdateWebinarRegistrationsAlt1Request struct {
	value *UpdateWebinarRegistrationsAlt1Request
	isSet bool
}

func (v NullableUpdateWebinarRegistrationsAlt1Request) Get() *UpdateWebinarRegistrationsAlt1Request {
	return v.value
}

func (v *NullableUpdateWebinarRegistrationsAlt1Request) Set(val *UpdateWebinarRegistrationsAlt1Request) {
	v.value = val
	v.isSet = true
}

func (v NullableUpdateWebinarRegistrationsAlt1Request) IsSet() bool {
	return v.isSet
}

func (v *NullableUpdateWebinarRegistrationsAlt1Request) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableUpdateWebinarRegistrationsAlt1Request(val *UpdateWebinarRegistrationsAlt1Request) *NullableUpdateWebinarRegistrationsAlt1Request {
	return &NullableUpdateWebinarRegistrationsAlt1Request{value: val, isSet: true}
}

func (v NullableUpdateWebinarRegistrationsAlt1Request) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableUpdateWebinarRegistrationsAlt1Request) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


