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

// checks if the EditPictureAlt1Request type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &EditPictureAlt1Request{}

// EditPictureAlt1Request struct for EditPictureAlt1Request
type EditPictureAlt1Request struct {
	// Whether the picture is the authenticated user's active portrait.
	Active *bool `json:"active,omitempty"`
}

// NewEditPictureAlt1Request instantiates a new EditPictureAlt1Request object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewEditPictureAlt1Request() *EditPictureAlt1Request {
	this := EditPictureAlt1Request{}
	return &this
}

// NewEditPictureAlt1RequestWithDefaults instantiates a new EditPictureAlt1Request object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewEditPictureAlt1RequestWithDefaults() *EditPictureAlt1Request {
	this := EditPictureAlt1Request{}
	return &this
}

// GetActive returns the Active field value if set, zero value otherwise.
func (o *EditPictureAlt1Request) GetActive() bool {
	if o == nil || IsNil(o.Active) {
		var ret bool
		return ret
	}
	return *o.Active
}

// GetActiveOk returns a tuple with the Active field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *EditPictureAlt1Request) GetActiveOk() (*bool, bool) {
	if o == nil || IsNil(o.Active) {
		return nil, false
	}
	return o.Active, true
}

// HasActive returns a boolean if a field has been set.
func (o *EditPictureAlt1Request) HasActive() bool {
	if o != nil && !IsNil(o.Active) {
		return true
	}

	return false
}

// SetActive gets a reference to the given bool and assigns it to the Active field.
func (o *EditPictureAlt1Request) SetActive(v bool) {
	o.Active = &v
}

func (o EditPictureAlt1Request) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o EditPictureAlt1Request) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Active) {
		toSerialize["active"] = o.Active
	}
	return toSerialize, nil
}

type NullableEditPictureAlt1Request struct {
	value *EditPictureAlt1Request
	isSet bool
}

func (v NullableEditPictureAlt1Request) Get() *EditPictureAlt1Request {
	return v.value
}

func (v *NullableEditPictureAlt1Request) Set(val *EditPictureAlt1Request) {
	v.value = val
	v.isSet = true
}

func (v NullableEditPictureAlt1Request) IsSet() bool {
	return v.isSet
}

func (v *NullableEditPictureAlt1Request) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableEditPictureAlt1Request(val *EditPictureAlt1Request) *NullableEditPictureAlt1Request {
	return &NullableEditPictureAlt1Request{value: val, isSet: true}
}

func (v NullableEditPictureAlt1Request) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableEditPictureAlt1Request) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


