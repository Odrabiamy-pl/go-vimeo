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

// checks if the CreateEmbedPresetsAlt1Request type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &CreateEmbedPresetsAlt1Request{}

// CreateEmbedPresetsAlt1Request struct for CreateEmbedPresetsAlt1Request
type CreateEmbedPresetsAlt1Request struct {
	Embed *CreateEmbedPresetsAlt1RequestEmbed `json:"embed,omitempty"`
	// The name of the embed preset.
	Name *string `json:"name,omitempty"`
}

// NewCreateEmbedPresetsAlt1Request instantiates a new CreateEmbedPresetsAlt1Request object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewCreateEmbedPresetsAlt1Request() *CreateEmbedPresetsAlt1Request {
	this := CreateEmbedPresetsAlt1Request{}
	return &this
}

// NewCreateEmbedPresetsAlt1RequestWithDefaults instantiates a new CreateEmbedPresetsAlt1Request object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewCreateEmbedPresetsAlt1RequestWithDefaults() *CreateEmbedPresetsAlt1Request {
	this := CreateEmbedPresetsAlt1Request{}
	return &this
}

// GetEmbed returns the Embed field value if set, zero value otherwise.
func (o *CreateEmbedPresetsAlt1Request) GetEmbed() CreateEmbedPresetsAlt1RequestEmbed {
	if o == nil || IsNil(o.Embed) {
		var ret CreateEmbedPresetsAlt1RequestEmbed
		return ret
	}
	return *o.Embed
}

// GetEmbedOk returns a tuple with the Embed field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CreateEmbedPresetsAlt1Request) GetEmbedOk() (*CreateEmbedPresetsAlt1RequestEmbed, bool) {
	if o == nil || IsNil(o.Embed) {
		return nil, false
	}
	return o.Embed, true
}

// HasEmbed returns a boolean if a field has been set.
func (o *CreateEmbedPresetsAlt1Request) HasEmbed() bool {
	if o != nil && !IsNil(o.Embed) {
		return true
	}

	return false
}

// SetEmbed gets a reference to the given CreateEmbedPresetsAlt1RequestEmbed and assigns it to the Embed field.
func (o *CreateEmbedPresetsAlt1Request) SetEmbed(v CreateEmbedPresetsAlt1RequestEmbed) {
	o.Embed = &v
}

// GetName returns the Name field value if set, zero value otherwise.
func (o *CreateEmbedPresetsAlt1Request) GetName() string {
	if o == nil || IsNil(o.Name) {
		var ret string
		return ret
	}
	return *o.Name
}

// GetNameOk returns a tuple with the Name field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CreateEmbedPresetsAlt1Request) GetNameOk() (*string, bool) {
	if o == nil || IsNil(o.Name) {
		return nil, false
	}
	return o.Name, true
}

// HasName returns a boolean if a field has been set.
func (o *CreateEmbedPresetsAlt1Request) HasName() bool {
	if o != nil && !IsNil(o.Name) {
		return true
	}

	return false
}

// SetName gets a reference to the given string and assigns it to the Name field.
func (o *CreateEmbedPresetsAlt1Request) SetName(v string) {
	o.Name = &v
}

func (o CreateEmbedPresetsAlt1Request) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o CreateEmbedPresetsAlt1Request) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Embed) {
		toSerialize["embed"] = o.Embed
	}
	if !IsNil(o.Name) {
		toSerialize["name"] = o.Name
	}
	return toSerialize, nil
}

type NullableCreateEmbedPresetsAlt1Request struct {
	value *CreateEmbedPresetsAlt1Request
	isSet bool
}

func (v NullableCreateEmbedPresetsAlt1Request) Get() *CreateEmbedPresetsAlt1Request {
	return v.value
}

func (v *NullableCreateEmbedPresetsAlt1Request) Set(val *CreateEmbedPresetsAlt1Request) {
	v.value = val
	v.isSet = true
}

func (v NullableCreateEmbedPresetsAlt1Request) IsSet() bool {
	return v.isSet
}

func (v *NullableCreateEmbedPresetsAlt1Request) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableCreateEmbedPresetsAlt1Request(val *CreateEmbedPresetsAlt1Request) *NullableCreateEmbedPresetsAlt1Request {
	return &NullableCreateEmbedPresetsAlt1Request{value: val, isSet: true}
}

func (v NullableCreateEmbedPresetsAlt1Request) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableCreateEmbedPresetsAlt1Request) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


