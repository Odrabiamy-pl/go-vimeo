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

// checks if the DisabledVideoPropertiesEmbedPresetsReasonsInner type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &DisabledVideoPropertiesEmbedPresetsReasonsInner{}

// DisabledVideoPropertiesEmbedPresetsReasonsInner struct for DisabledVideoPropertiesEmbedPresetsReasonsInner
type DisabledVideoPropertiesEmbedPresetsReasonsInner struct {
	// The icon that represents the reason why embed presets are disabled.  Option descriptions:  * `clock` - The reason is represented by a clock icon.  * `create` - The reason is represented by a create icon.  * `image` - The reason is represented by an image icon.  * `theme` - The reason is represented by a theme icon. 
	Icon string `json:"icon"`
	// A user-deliverable message of why embed presets are disabled.
	Message string `json:"message"`
}

// NewDisabledVideoPropertiesEmbedPresetsReasonsInner instantiates a new DisabledVideoPropertiesEmbedPresetsReasonsInner object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewDisabledVideoPropertiesEmbedPresetsReasonsInner(icon string, message string) *DisabledVideoPropertiesEmbedPresetsReasonsInner {
	this := DisabledVideoPropertiesEmbedPresetsReasonsInner{}
	this.Icon = icon
	this.Message = message
	return &this
}

// NewDisabledVideoPropertiesEmbedPresetsReasonsInnerWithDefaults instantiates a new DisabledVideoPropertiesEmbedPresetsReasonsInner object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewDisabledVideoPropertiesEmbedPresetsReasonsInnerWithDefaults() *DisabledVideoPropertiesEmbedPresetsReasonsInner {
	this := DisabledVideoPropertiesEmbedPresetsReasonsInner{}
	return &this
}

// GetIcon returns the Icon field value
func (o *DisabledVideoPropertiesEmbedPresetsReasonsInner) GetIcon() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Icon
}

// GetIconOk returns a tuple with the Icon field value
// and a boolean to check if the value has been set.
func (o *DisabledVideoPropertiesEmbedPresetsReasonsInner) GetIconOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Icon, true
}

// SetIcon sets field value
func (o *DisabledVideoPropertiesEmbedPresetsReasonsInner) SetIcon(v string) {
	o.Icon = v
}

// GetMessage returns the Message field value
func (o *DisabledVideoPropertiesEmbedPresetsReasonsInner) GetMessage() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Message
}

// GetMessageOk returns a tuple with the Message field value
// and a boolean to check if the value has been set.
func (o *DisabledVideoPropertiesEmbedPresetsReasonsInner) GetMessageOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Message, true
}

// SetMessage sets field value
func (o *DisabledVideoPropertiesEmbedPresetsReasonsInner) SetMessage(v string) {
	o.Message = v
}

func (o DisabledVideoPropertiesEmbedPresetsReasonsInner) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o DisabledVideoPropertiesEmbedPresetsReasonsInner) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["icon"] = o.Icon
	toSerialize["message"] = o.Message
	return toSerialize, nil
}

type NullableDisabledVideoPropertiesEmbedPresetsReasonsInner struct {
	value *DisabledVideoPropertiesEmbedPresetsReasonsInner
	isSet bool
}

func (v NullableDisabledVideoPropertiesEmbedPresetsReasonsInner) Get() *DisabledVideoPropertiesEmbedPresetsReasonsInner {
	return v.value
}

func (v *NullableDisabledVideoPropertiesEmbedPresetsReasonsInner) Set(val *DisabledVideoPropertiesEmbedPresetsReasonsInner) {
	v.value = val
	v.isSet = true
}

func (v NullableDisabledVideoPropertiesEmbedPresetsReasonsInner) IsSet() bool {
	return v.isSet
}

func (v *NullableDisabledVideoPropertiesEmbedPresetsReasonsInner) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableDisabledVideoPropertiesEmbedPresetsReasonsInner(val *DisabledVideoPropertiesEmbedPresetsReasonsInner) *NullableDisabledVideoPropertiesEmbedPresetsReasonsInner {
	return &NullableDisabledVideoPropertiesEmbedPresetsReasonsInner{value: val, isSet: true}
}

func (v NullableDisabledVideoPropertiesEmbedPresetsReasonsInner) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableDisabledVideoPropertiesEmbedPresetsReasonsInner) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


