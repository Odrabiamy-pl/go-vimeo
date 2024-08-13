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

// checks if the EditLiveEventAutoCcAlt1Request type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &EditLiveEventAutoCcAlt1Request{}

// EditLiveEventAutoCcAlt1Request struct for EditLiveEventAutoCcAlt1Request
type EditLiveEventAutoCcAlt1Request struct {
	// Whether automated closed captions are enabled for the event.
	AutoCcEnabled bool `json:"auto_cc_enabled"`
	// A comma-separated list of keywords that improve the quality of the automated closed captions.
	AutoCcKeywords *string `json:"auto_cc_keywords,omitempty"`
	// The language in which the automated closed captions appear.  Option descriptions:  * `de-DE` - The language is German.  * `en-US` - The language is English.  * `es-ES` - The language is Spanish.  * `fr-FR` - The language is French.  * `pt-BR` - The language is Portuguese. 
	AutoCcLang *string `json:"auto_cc_lang,omitempty"`
}

// NewEditLiveEventAutoCcAlt1Request instantiates a new EditLiveEventAutoCcAlt1Request object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewEditLiveEventAutoCcAlt1Request(autoCcEnabled bool) *EditLiveEventAutoCcAlt1Request {
	this := EditLiveEventAutoCcAlt1Request{}
	this.AutoCcEnabled = autoCcEnabled
	return &this
}

// NewEditLiveEventAutoCcAlt1RequestWithDefaults instantiates a new EditLiveEventAutoCcAlt1Request object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewEditLiveEventAutoCcAlt1RequestWithDefaults() *EditLiveEventAutoCcAlt1Request {
	this := EditLiveEventAutoCcAlt1Request{}
	return &this
}

// GetAutoCcEnabled returns the AutoCcEnabled field value
func (o *EditLiveEventAutoCcAlt1Request) GetAutoCcEnabled() bool {
	if o == nil {
		var ret bool
		return ret
	}

	return o.AutoCcEnabled
}

// GetAutoCcEnabledOk returns a tuple with the AutoCcEnabled field value
// and a boolean to check if the value has been set.
func (o *EditLiveEventAutoCcAlt1Request) GetAutoCcEnabledOk() (*bool, bool) {
	if o == nil {
		return nil, false
	}
	return &o.AutoCcEnabled, true
}

// SetAutoCcEnabled sets field value
func (o *EditLiveEventAutoCcAlt1Request) SetAutoCcEnabled(v bool) {
	o.AutoCcEnabled = v
}

// GetAutoCcKeywords returns the AutoCcKeywords field value if set, zero value otherwise.
func (o *EditLiveEventAutoCcAlt1Request) GetAutoCcKeywords() string {
	if o == nil || IsNil(o.AutoCcKeywords) {
		var ret string
		return ret
	}
	return *o.AutoCcKeywords
}

// GetAutoCcKeywordsOk returns a tuple with the AutoCcKeywords field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *EditLiveEventAutoCcAlt1Request) GetAutoCcKeywordsOk() (*string, bool) {
	if o == nil || IsNil(o.AutoCcKeywords) {
		return nil, false
	}
	return o.AutoCcKeywords, true
}

// HasAutoCcKeywords returns a boolean if a field has been set.
func (o *EditLiveEventAutoCcAlt1Request) HasAutoCcKeywords() bool {
	if o != nil && !IsNil(o.AutoCcKeywords) {
		return true
	}

	return false
}

// SetAutoCcKeywords gets a reference to the given string and assigns it to the AutoCcKeywords field.
func (o *EditLiveEventAutoCcAlt1Request) SetAutoCcKeywords(v string) {
	o.AutoCcKeywords = &v
}

// GetAutoCcLang returns the AutoCcLang field value if set, zero value otherwise.
func (o *EditLiveEventAutoCcAlt1Request) GetAutoCcLang() string {
	if o == nil || IsNil(o.AutoCcLang) {
		var ret string
		return ret
	}
	return *o.AutoCcLang
}

// GetAutoCcLangOk returns a tuple with the AutoCcLang field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *EditLiveEventAutoCcAlt1Request) GetAutoCcLangOk() (*string, bool) {
	if o == nil || IsNil(o.AutoCcLang) {
		return nil, false
	}
	return o.AutoCcLang, true
}

// HasAutoCcLang returns a boolean if a field has been set.
func (o *EditLiveEventAutoCcAlt1Request) HasAutoCcLang() bool {
	if o != nil && !IsNil(o.AutoCcLang) {
		return true
	}

	return false
}

// SetAutoCcLang gets a reference to the given string and assigns it to the AutoCcLang field.
func (o *EditLiveEventAutoCcAlt1Request) SetAutoCcLang(v string) {
	o.AutoCcLang = &v
}

func (o EditLiveEventAutoCcAlt1Request) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o EditLiveEventAutoCcAlt1Request) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["auto_cc_enabled"] = o.AutoCcEnabled
	if !IsNil(o.AutoCcKeywords) {
		toSerialize["auto_cc_keywords"] = o.AutoCcKeywords
	}
	if !IsNil(o.AutoCcLang) {
		toSerialize["auto_cc_lang"] = o.AutoCcLang
	}
	return toSerialize, nil
}

type NullableEditLiveEventAutoCcAlt1Request struct {
	value *EditLiveEventAutoCcAlt1Request
	isSet bool
}

func (v NullableEditLiveEventAutoCcAlt1Request) Get() *EditLiveEventAutoCcAlt1Request {
	return v.value
}

func (v *NullableEditLiveEventAutoCcAlt1Request) Set(val *EditLiveEventAutoCcAlt1Request) {
	v.value = val
	v.isSet = true
}

func (v NullableEditLiveEventAutoCcAlt1Request) IsSet() bool {
	return v.isSet
}

func (v *NullableEditLiveEventAutoCcAlt1Request) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableEditLiveEventAutoCcAlt1Request(val *EditLiveEventAutoCcAlt1Request) *NullableEditLiveEventAutoCcAlt1Request {
	return &NullableEditLiveEventAutoCcAlt1Request{value: val, isSet: true}
}

func (v NullableEditLiveEventAutoCcAlt1Request) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableEditLiveEventAutoCcAlt1Request) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


