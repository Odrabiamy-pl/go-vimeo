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

// checks if the WebinarEmailSettingsEmailPostEventThankYou type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &WebinarEmailSettingsEmailPostEventThankYou{}

// WebinarEmailSettingsEmailPostEventThankYou The email customization details for the webinar post-event thank-you email.
type WebinarEmailSettingsEmailPostEventThankYou struct {
	Custom NullableWebinarEmailSettingsEmailPostEventThankYouCustom `json:"custom"`
	Default WebinarEmailContent `json:"default"`
}

// NewWebinarEmailSettingsEmailPostEventThankYou instantiates a new WebinarEmailSettingsEmailPostEventThankYou object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewWebinarEmailSettingsEmailPostEventThankYou(custom NullableWebinarEmailSettingsEmailPostEventThankYouCustom, default_ WebinarEmailContent) *WebinarEmailSettingsEmailPostEventThankYou {
	this := WebinarEmailSettingsEmailPostEventThankYou{}
	this.Custom = custom
	this.Default = default_
	return &this
}

// NewWebinarEmailSettingsEmailPostEventThankYouWithDefaults instantiates a new WebinarEmailSettingsEmailPostEventThankYou object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewWebinarEmailSettingsEmailPostEventThankYouWithDefaults() *WebinarEmailSettingsEmailPostEventThankYou {
	this := WebinarEmailSettingsEmailPostEventThankYou{}
	return &this
}

// GetCustom returns the Custom field value
// If the value is explicit nil, the zero value for WebinarEmailSettingsEmailPostEventThankYouCustom will be returned
func (o *WebinarEmailSettingsEmailPostEventThankYou) GetCustom() WebinarEmailSettingsEmailPostEventThankYouCustom {
	if o == nil || o.Custom.Get() == nil {
		var ret WebinarEmailSettingsEmailPostEventThankYouCustom
		return ret
	}

	return *o.Custom.Get()
}

// GetCustomOk returns a tuple with the Custom field value
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *WebinarEmailSettingsEmailPostEventThankYou) GetCustomOk() (*WebinarEmailSettingsEmailPostEventThankYouCustom, bool) {
	if o == nil {
		return nil, false
	}
	return o.Custom.Get(), o.Custom.IsSet()
}

// SetCustom sets field value
func (o *WebinarEmailSettingsEmailPostEventThankYou) SetCustom(v WebinarEmailSettingsEmailPostEventThankYouCustom) {
	o.Custom.Set(&v)
}

// GetDefault returns the Default field value
func (o *WebinarEmailSettingsEmailPostEventThankYou) GetDefault() WebinarEmailContent {
	if o == nil {
		var ret WebinarEmailContent
		return ret
	}

	return o.Default
}

// GetDefaultOk returns a tuple with the Default field value
// and a boolean to check if the value has been set.
func (o *WebinarEmailSettingsEmailPostEventThankYou) GetDefaultOk() (*WebinarEmailContent, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Default, true
}

// SetDefault sets field value
func (o *WebinarEmailSettingsEmailPostEventThankYou) SetDefault(v WebinarEmailContent) {
	o.Default = v
}

func (o WebinarEmailSettingsEmailPostEventThankYou) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o WebinarEmailSettingsEmailPostEventThankYou) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["custom"] = o.Custom.Get()
	toSerialize["default"] = o.Default
	return toSerialize, nil
}

type NullableWebinarEmailSettingsEmailPostEventThankYou struct {
	value *WebinarEmailSettingsEmailPostEventThankYou
	isSet bool
}

func (v NullableWebinarEmailSettingsEmailPostEventThankYou) Get() *WebinarEmailSettingsEmailPostEventThankYou {
	return v.value
}

func (v *NullableWebinarEmailSettingsEmailPostEventThankYou) Set(val *WebinarEmailSettingsEmailPostEventThankYou) {
	v.value = val
	v.isSet = true
}

func (v NullableWebinarEmailSettingsEmailPostEventThankYou) IsSet() bool {
	return v.isSet
}

func (v *NullableWebinarEmailSettingsEmailPostEventThankYou) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableWebinarEmailSettingsEmailPostEventThankYou(val *WebinarEmailSettingsEmailPostEventThankYou) *NullableWebinarEmailSettingsEmailPostEventThankYou {
	return &NullableWebinarEmailSettingsEmailPostEventThankYou{value: val, isSet: true}
}

func (v NullableWebinarEmailSettingsEmailPostEventThankYou) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableWebinarEmailSettingsEmailPostEventThankYou) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


