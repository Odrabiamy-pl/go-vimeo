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

// checks if the CreateLiveEventAlt1RequestEmbedLogos type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &CreateLiveEventAlt1RequestEmbedLogos{}

// CreateLiveEventAlt1RequestEmbedLogos struct for CreateLiveEventAlt1RequestEmbedLogos
type CreateLiveEventAlt1RequestEmbedLogos struct {
	Custom *CreateLiveEventAlt1RequestEmbedLogosCustom `json:"custom,omitempty"`
	// Whether to show the Vimeo logo on the embed player.
	Vimeo *bool `json:"vimeo,omitempty"`
}

// NewCreateLiveEventAlt1RequestEmbedLogos instantiates a new CreateLiveEventAlt1RequestEmbedLogos object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewCreateLiveEventAlt1RequestEmbedLogos() *CreateLiveEventAlt1RequestEmbedLogos {
	this := CreateLiveEventAlt1RequestEmbedLogos{}
	return &this
}

// NewCreateLiveEventAlt1RequestEmbedLogosWithDefaults instantiates a new CreateLiveEventAlt1RequestEmbedLogos object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewCreateLiveEventAlt1RequestEmbedLogosWithDefaults() *CreateLiveEventAlt1RequestEmbedLogos {
	this := CreateLiveEventAlt1RequestEmbedLogos{}
	return &this
}

// GetCustom returns the Custom field value if set, zero value otherwise.
func (o *CreateLiveEventAlt1RequestEmbedLogos) GetCustom() CreateLiveEventAlt1RequestEmbedLogosCustom {
	if o == nil || IsNil(o.Custom) {
		var ret CreateLiveEventAlt1RequestEmbedLogosCustom
		return ret
	}
	return *o.Custom
}

// GetCustomOk returns a tuple with the Custom field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CreateLiveEventAlt1RequestEmbedLogos) GetCustomOk() (*CreateLiveEventAlt1RequestEmbedLogosCustom, bool) {
	if o == nil || IsNil(o.Custom) {
		return nil, false
	}
	return o.Custom, true
}

// HasCustom returns a boolean if a field has been set.
func (o *CreateLiveEventAlt1RequestEmbedLogos) HasCustom() bool {
	if o != nil && !IsNil(o.Custom) {
		return true
	}

	return false
}

// SetCustom gets a reference to the given CreateLiveEventAlt1RequestEmbedLogosCustom and assigns it to the Custom field.
func (o *CreateLiveEventAlt1RequestEmbedLogos) SetCustom(v CreateLiveEventAlt1RequestEmbedLogosCustom) {
	o.Custom = &v
}

// GetVimeo returns the Vimeo field value if set, zero value otherwise.
func (o *CreateLiveEventAlt1RequestEmbedLogos) GetVimeo() bool {
	if o == nil || IsNil(o.Vimeo) {
		var ret bool
		return ret
	}
	return *o.Vimeo
}

// GetVimeoOk returns a tuple with the Vimeo field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CreateLiveEventAlt1RequestEmbedLogos) GetVimeoOk() (*bool, bool) {
	if o == nil || IsNil(o.Vimeo) {
		return nil, false
	}
	return o.Vimeo, true
}

// HasVimeo returns a boolean if a field has been set.
func (o *CreateLiveEventAlt1RequestEmbedLogos) HasVimeo() bool {
	if o != nil && !IsNil(o.Vimeo) {
		return true
	}

	return false
}

// SetVimeo gets a reference to the given bool and assigns it to the Vimeo field.
func (o *CreateLiveEventAlt1RequestEmbedLogos) SetVimeo(v bool) {
	o.Vimeo = &v
}

func (o CreateLiveEventAlt1RequestEmbedLogos) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o CreateLiveEventAlt1RequestEmbedLogos) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Custom) {
		toSerialize["custom"] = o.Custom
	}
	if !IsNil(o.Vimeo) {
		toSerialize["vimeo"] = o.Vimeo
	}
	return toSerialize, nil
}

type NullableCreateLiveEventAlt1RequestEmbedLogos struct {
	value *CreateLiveEventAlt1RequestEmbedLogos
	isSet bool
}

func (v NullableCreateLiveEventAlt1RequestEmbedLogos) Get() *CreateLiveEventAlt1RequestEmbedLogos {
	return v.value
}

func (v *NullableCreateLiveEventAlt1RequestEmbedLogos) Set(val *CreateLiveEventAlt1RequestEmbedLogos) {
	v.value = val
	v.isSet = true
}

func (v NullableCreateLiveEventAlt1RequestEmbedLogos) IsSet() bool {
	return v.isSet
}

func (v *NullableCreateLiveEventAlt1RequestEmbedLogos) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableCreateLiveEventAlt1RequestEmbedLogos(val *CreateLiveEventAlt1RequestEmbedLogos) *NullableCreateLiveEventAlt1RequestEmbedLogos {
	return &NullableCreateLiveEventAlt1RequestEmbedLogos{value: val, isSet: true}
}

func (v NullableCreateLiveEventAlt1RequestEmbedLogos) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableCreateLiveEventAlt1RequestEmbedLogos) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
