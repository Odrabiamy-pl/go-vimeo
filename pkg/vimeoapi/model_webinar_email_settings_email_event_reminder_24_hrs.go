/*
Vimeo API

Build something great. Vimeo's API supports flexible, high-quality video integration with your custom apps.

API version: 3.4
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package vimeoapi

import (
	"encoding/json"
	"bytes"
	"fmt"
)

// checks if the WebinarEmailSettingsEmailEventReminder24Hrs type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &WebinarEmailSettingsEmailEventReminder24Hrs{}

// WebinarEmailSettingsEmailEventReminder24Hrs The email customization details for the webinar reminder email, which goes out 24 hours before the event.
type WebinarEmailSettingsEmailEventReminder24Hrs struct {
	Custom NullableWebinarEmailSettingsEmailEventReminder24HrsCustom `json:"custom"`
	// The email default details for the webinar reminder email, which goes out 24 hours before the event.
	Default WebinarEmailContent `json:"default"`
}

type _WebinarEmailSettingsEmailEventReminder24Hrs WebinarEmailSettingsEmailEventReminder24Hrs

// NewWebinarEmailSettingsEmailEventReminder24Hrs instantiates a new WebinarEmailSettingsEmailEventReminder24Hrs object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewWebinarEmailSettingsEmailEventReminder24Hrs(custom NullableWebinarEmailSettingsEmailEventReminder24HrsCustom, default_ WebinarEmailContent) *WebinarEmailSettingsEmailEventReminder24Hrs {
	this := WebinarEmailSettingsEmailEventReminder24Hrs{}
	this.Custom = custom
	this.Default = default_
	return &this
}

// NewWebinarEmailSettingsEmailEventReminder24HrsWithDefaults instantiates a new WebinarEmailSettingsEmailEventReminder24Hrs object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewWebinarEmailSettingsEmailEventReminder24HrsWithDefaults() *WebinarEmailSettingsEmailEventReminder24Hrs {
	this := WebinarEmailSettingsEmailEventReminder24Hrs{}
	return &this
}

// GetCustom returns the Custom field value
// If the value is explicit nil, the zero value for WebinarEmailSettingsEmailEventReminder24HrsCustom will be returned
func (o *WebinarEmailSettingsEmailEventReminder24Hrs) GetCustom() WebinarEmailSettingsEmailEventReminder24HrsCustom {
	if o == nil || o.Custom.Get() == nil {
		var ret WebinarEmailSettingsEmailEventReminder24HrsCustom
		return ret
	}

	return *o.Custom.Get()
}

// GetCustomOk returns a tuple with the Custom field value
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *WebinarEmailSettingsEmailEventReminder24Hrs) GetCustomOk() (*WebinarEmailSettingsEmailEventReminder24HrsCustom, bool) {
	if o == nil {
		return nil, false
	}
	return o.Custom.Get(), o.Custom.IsSet()
}

// SetCustom sets field value
func (o *WebinarEmailSettingsEmailEventReminder24Hrs) SetCustom(v WebinarEmailSettingsEmailEventReminder24HrsCustom) {
	o.Custom.Set(&v)
}

// GetDefault returns the Default field value
func (o *WebinarEmailSettingsEmailEventReminder24Hrs) GetDefault() WebinarEmailContent {
	if o == nil {
		var ret WebinarEmailContent
		return ret
	}

	return o.Default
}

// GetDefaultOk returns a tuple with the Default field value
// and a boolean to check if the value has been set.
func (o *WebinarEmailSettingsEmailEventReminder24Hrs) GetDefaultOk() (*WebinarEmailContent, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Default, true
}

// SetDefault sets field value
func (o *WebinarEmailSettingsEmailEventReminder24Hrs) SetDefault(v WebinarEmailContent) {
	o.Default = v
}

func (o WebinarEmailSettingsEmailEventReminder24Hrs) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o WebinarEmailSettingsEmailEventReminder24Hrs) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["custom"] = o.Custom.Get()
	toSerialize["default"] = o.Default
	return toSerialize, nil
}

func (o *WebinarEmailSettingsEmailEventReminder24Hrs) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"custom",
		"default",
	}

	allProperties := make(map[string]interface{})

	err = json.Unmarshal(data, &allProperties)

	if err != nil {
		return err;
	}

	for _, requiredProperty := range(requiredProperties) {
		if _, exists := allProperties[requiredProperty]; !exists {
			return fmt.Errorf("no value given for required property %v", requiredProperty)
		}
	}

	varWebinarEmailSettingsEmailEventReminder24Hrs := _WebinarEmailSettingsEmailEventReminder24Hrs{}

	decoder := json.NewDecoder(bytes.NewReader(data))
	decoder.DisallowUnknownFields()
	err = decoder.Decode(&varWebinarEmailSettingsEmailEventReminder24Hrs)

	if err != nil {
		return err
	}

	*o = WebinarEmailSettingsEmailEventReminder24Hrs(varWebinarEmailSettingsEmailEventReminder24Hrs)

	return err
}

type NullableWebinarEmailSettingsEmailEventReminder24Hrs struct {
	value *WebinarEmailSettingsEmailEventReminder24Hrs
	isSet bool
}

func (v NullableWebinarEmailSettingsEmailEventReminder24Hrs) Get() *WebinarEmailSettingsEmailEventReminder24Hrs {
	return v.value
}

func (v *NullableWebinarEmailSettingsEmailEventReminder24Hrs) Set(val *WebinarEmailSettingsEmailEventReminder24Hrs) {
	v.value = val
	v.isSet = true
}

func (v NullableWebinarEmailSettingsEmailEventReminder24Hrs) IsSet() bool {
	return v.isSet
}

func (v *NullableWebinarEmailSettingsEmailEventReminder24Hrs) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableWebinarEmailSettingsEmailEventReminder24Hrs(val *WebinarEmailSettingsEmailEventReminder24Hrs) *NullableWebinarEmailSettingsEmailEventReminder24Hrs {
	return &NullableWebinarEmailSettingsEmailEventReminder24Hrs{value: val, isSet: true}
}

func (v NullableWebinarEmailSettingsEmailEventReminder24Hrs) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableWebinarEmailSettingsEmailEventReminder24Hrs) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

