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

// checks if the DisabledVideoPropertiesEmbedReasonsInner type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &DisabledVideoPropertiesEmbedReasonsInner{}

// DisabledVideoPropertiesEmbedReasonsInner struct for DisabledVideoPropertiesEmbedReasonsInner
type DisabledVideoPropertiesEmbedReasonsInner struct {
	// The icon that represents the reason why embed is disabled.  Option descriptions:  * `clock` - The reason is represented by a clock icon.  * `create` - The reason is represented by a create icon.  * `image` - The reason is represented by an image icon.  * `theme` - The reason is represented by a theme icon. 
	Icon string `json:"icon"`
	// A user-deliverable message of why embed is disabled.
	Message string `json:"message"`
}

// NewDisabledVideoPropertiesEmbedReasonsInner instantiates a new DisabledVideoPropertiesEmbedReasonsInner object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewDisabledVideoPropertiesEmbedReasonsInner(icon string, message string) *DisabledVideoPropertiesEmbedReasonsInner {
	this := DisabledVideoPropertiesEmbedReasonsInner{}
	this.Icon = icon
	this.Message = message
	return &this
}

// NewDisabledVideoPropertiesEmbedReasonsInnerWithDefaults instantiates a new DisabledVideoPropertiesEmbedReasonsInner object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewDisabledVideoPropertiesEmbedReasonsInnerWithDefaults() *DisabledVideoPropertiesEmbedReasonsInner {
	this := DisabledVideoPropertiesEmbedReasonsInner{}
	return &this
}

// GetIcon returns the Icon field value
func (o *DisabledVideoPropertiesEmbedReasonsInner) GetIcon() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Icon
}

// GetIconOk returns a tuple with the Icon field value
// and a boolean to check if the value has been set.
func (o *DisabledVideoPropertiesEmbedReasonsInner) GetIconOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Icon, true
}

// SetIcon sets field value
func (o *DisabledVideoPropertiesEmbedReasonsInner) SetIcon(v string) {
	o.Icon = v
}

// GetMessage returns the Message field value
func (o *DisabledVideoPropertiesEmbedReasonsInner) GetMessage() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Message
}

// GetMessageOk returns a tuple with the Message field value
// and a boolean to check if the value has been set.
func (o *DisabledVideoPropertiesEmbedReasonsInner) GetMessageOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Message, true
}

// SetMessage sets field value
func (o *DisabledVideoPropertiesEmbedReasonsInner) SetMessage(v string) {
	o.Message = v
}

func (o DisabledVideoPropertiesEmbedReasonsInner) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o DisabledVideoPropertiesEmbedReasonsInner) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["icon"] = o.Icon
	toSerialize["message"] = o.Message
	return toSerialize, nil
}

type NullableDisabledVideoPropertiesEmbedReasonsInner struct {
	value *DisabledVideoPropertiesEmbedReasonsInner
	isSet bool
}

func (v NullableDisabledVideoPropertiesEmbedReasonsInner) Get() *DisabledVideoPropertiesEmbedReasonsInner {
	return v.value
}

func (v *NullableDisabledVideoPropertiesEmbedReasonsInner) Set(val *DisabledVideoPropertiesEmbedReasonsInner) {
	v.value = val
	v.isSet = true
}

func (v NullableDisabledVideoPropertiesEmbedReasonsInner) IsSet() bool {
	return v.isSet
}

func (v *NullableDisabledVideoPropertiesEmbedReasonsInner) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableDisabledVideoPropertiesEmbedReasonsInner(val *DisabledVideoPropertiesEmbedReasonsInner) *NullableDisabledVideoPropertiesEmbedReasonsInner {
	return &NullableDisabledVideoPropertiesEmbedReasonsInner{value: val, isSet: true}
}

func (v NullableDisabledVideoPropertiesEmbedReasonsInner) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableDisabledVideoPropertiesEmbedReasonsInner) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


