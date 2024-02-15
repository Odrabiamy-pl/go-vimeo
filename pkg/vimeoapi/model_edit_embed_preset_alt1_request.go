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

// checks if the EditEmbedPresetAlt1Request type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &EditEmbedPresetAlt1Request{}

// EditEmbedPresetAlt1Request struct for EditEmbedPresetAlt1Request
type EditEmbedPresetAlt1Request struct {
	// What to do with the outro.  Option descriptions:  * `nothing` - Disable the outro. 
	Outro *string `json:"outro,omitempty"`
}

// NewEditEmbedPresetAlt1Request instantiates a new EditEmbedPresetAlt1Request object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewEditEmbedPresetAlt1Request() *EditEmbedPresetAlt1Request {
	this := EditEmbedPresetAlt1Request{}
	return &this
}

// NewEditEmbedPresetAlt1RequestWithDefaults instantiates a new EditEmbedPresetAlt1Request object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewEditEmbedPresetAlt1RequestWithDefaults() *EditEmbedPresetAlt1Request {
	this := EditEmbedPresetAlt1Request{}
	return &this
}

// GetOutro returns the Outro field value if set, zero value otherwise.
func (o *EditEmbedPresetAlt1Request) GetOutro() string {
	if o == nil || IsNil(o.Outro) {
		var ret string
		return ret
	}
	return *o.Outro
}

// GetOutroOk returns a tuple with the Outro field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *EditEmbedPresetAlt1Request) GetOutroOk() (*string, bool) {
	if o == nil || IsNil(o.Outro) {
		return nil, false
	}
	return o.Outro, true
}

// HasOutro returns a boolean if a field has been set.
func (o *EditEmbedPresetAlt1Request) HasOutro() bool {
	if o != nil && !IsNil(o.Outro) {
		return true
	}

	return false
}

// SetOutro gets a reference to the given string and assigns it to the Outro field.
func (o *EditEmbedPresetAlt1Request) SetOutro(v string) {
	o.Outro = &v
}

func (o EditEmbedPresetAlt1Request) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o EditEmbedPresetAlt1Request) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Outro) {
		toSerialize["outro"] = o.Outro
	}
	return toSerialize, nil
}

type NullableEditEmbedPresetAlt1Request struct {
	value *EditEmbedPresetAlt1Request
	isSet bool
}

func (v NullableEditEmbedPresetAlt1Request) Get() *EditEmbedPresetAlt1Request {
	return v.value
}

func (v *NullableEditEmbedPresetAlt1Request) Set(val *EditEmbedPresetAlt1Request) {
	v.value = val
	v.isSet = true
}

func (v NullableEditEmbedPresetAlt1Request) IsSet() bool {
	return v.isSet
}

func (v *NullableEditEmbedPresetAlt1Request) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableEditEmbedPresetAlt1Request(val *EditEmbedPresetAlt1Request) *NullableEditEmbedPresetAlt1Request {
	return &NullableEditEmbedPresetAlt1Request{value: val, isSet: true}
}

func (v NullableEditEmbedPresetAlt1Request) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableEditEmbedPresetAlt1Request) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


