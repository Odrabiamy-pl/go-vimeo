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

// checks if the EditVideoRequestEmbedLogos type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &EditVideoRequestEmbedLogos{}

// EditVideoRequestEmbedLogos struct for EditVideoRequestEmbedLogos
type EditVideoRequestEmbedLogos struct {
	Custom *EditVideoRequestEmbedLogosCustom `json:"custom,omitempty"`
	// Whether to show the Vimeo logo on the embeddable player.
	Vimeo *bool `json:"vimeo,omitempty"`
}

// NewEditVideoRequestEmbedLogos instantiates a new EditVideoRequestEmbedLogos object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewEditVideoRequestEmbedLogos() *EditVideoRequestEmbedLogos {
	this := EditVideoRequestEmbedLogos{}
	return &this
}

// NewEditVideoRequestEmbedLogosWithDefaults instantiates a new EditVideoRequestEmbedLogos object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewEditVideoRequestEmbedLogosWithDefaults() *EditVideoRequestEmbedLogos {
	this := EditVideoRequestEmbedLogos{}
	return &this
}

// GetCustom returns the Custom field value if set, zero value otherwise.
func (o *EditVideoRequestEmbedLogos) GetCustom() EditVideoRequestEmbedLogosCustom {
	if o == nil || IsNil(o.Custom) {
		var ret EditVideoRequestEmbedLogosCustom
		return ret
	}
	return *o.Custom
}

// GetCustomOk returns a tuple with the Custom field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *EditVideoRequestEmbedLogos) GetCustomOk() (*EditVideoRequestEmbedLogosCustom, bool) {
	if o == nil || IsNil(o.Custom) {
		return nil, false
	}
	return o.Custom, true
}

// HasCustom returns a boolean if a field has been set.
func (o *EditVideoRequestEmbedLogos) HasCustom() bool {
	if o != nil && !IsNil(o.Custom) {
		return true
	}

	return false
}

// SetCustom gets a reference to the given EditVideoRequestEmbedLogosCustom and assigns it to the Custom field.
func (o *EditVideoRequestEmbedLogos) SetCustom(v EditVideoRequestEmbedLogosCustom) {
	o.Custom = &v
}

// GetVimeo returns the Vimeo field value if set, zero value otherwise.
func (o *EditVideoRequestEmbedLogos) GetVimeo() bool {
	if o == nil || IsNil(o.Vimeo) {
		var ret bool
		return ret
	}
	return *o.Vimeo
}

// GetVimeoOk returns a tuple with the Vimeo field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *EditVideoRequestEmbedLogos) GetVimeoOk() (*bool, bool) {
	if o == nil || IsNil(o.Vimeo) {
		return nil, false
	}
	return o.Vimeo, true
}

// HasVimeo returns a boolean if a field has been set.
func (o *EditVideoRequestEmbedLogos) HasVimeo() bool {
	if o != nil && !IsNil(o.Vimeo) {
		return true
	}

	return false
}

// SetVimeo gets a reference to the given bool and assigns it to the Vimeo field.
func (o *EditVideoRequestEmbedLogos) SetVimeo(v bool) {
	o.Vimeo = &v
}

func (o EditVideoRequestEmbedLogos) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o EditVideoRequestEmbedLogos) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Custom) {
		toSerialize["custom"] = o.Custom
	}
	if !IsNil(o.Vimeo) {
		toSerialize["vimeo"] = o.Vimeo
	}
	return toSerialize, nil
}

type NullableEditVideoRequestEmbedLogos struct {
	value *EditVideoRequestEmbedLogos
	isSet bool
}

func (v NullableEditVideoRequestEmbedLogos) Get() *EditVideoRequestEmbedLogos {
	return v.value
}

func (v *NullableEditVideoRequestEmbedLogos) Set(val *EditVideoRequestEmbedLogos) {
	v.value = val
	v.isSet = true
}

func (v NullableEditVideoRequestEmbedLogos) IsSet() bool {
	return v.isSet
}

func (v *NullableEditVideoRequestEmbedLogos) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableEditVideoRequestEmbedLogos(val *EditVideoRequestEmbedLogos) *NullableEditVideoRequestEmbedLogos {
	return &NullableEditVideoRequestEmbedLogos{value: val, isSet: true}
}

func (v NullableEditVideoRequestEmbedLogos) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableEditVideoRequestEmbedLogos) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
