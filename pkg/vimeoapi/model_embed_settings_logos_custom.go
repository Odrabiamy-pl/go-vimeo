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

// checks if the EmbedSettingsLogosCustom type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &EmbedSettingsLogosCustom{}

// EmbedSettingsLogosCustom A collection of information relating to custom logos in the embeddable player.
type EmbedSettingsLogosCustom struct {
	// Whether the custom logo appears in the embeddable player.
	Active bool `json:"active"`
	// The URL that loads upon clicking the custom logo.
	Link NullableString `json:"link"`
	// Whether the custom logo appears even when the player interface is hidden.
	Sticky bool `json:"sticky"`
	// The URL of the selected custom logo.
	Url NullableString `json:"url"`
	// Whether the custom logo should use the URL link.
	UseLink bool `json:"use_link"`
}

// NewEmbedSettingsLogosCustom instantiates a new EmbedSettingsLogosCustom object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewEmbedSettingsLogosCustom(active bool, link NullableString, sticky bool, url NullableString, useLink bool) *EmbedSettingsLogosCustom {
	this := EmbedSettingsLogosCustom{}
	this.Active = active
	this.Link = link
	this.Sticky = sticky
	this.Url = url
	this.UseLink = useLink
	return &this
}

// NewEmbedSettingsLogosCustomWithDefaults instantiates a new EmbedSettingsLogosCustom object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewEmbedSettingsLogosCustomWithDefaults() *EmbedSettingsLogosCustom {
	this := EmbedSettingsLogosCustom{}
	return &this
}

// GetActive returns the Active field value
func (o *EmbedSettingsLogosCustom) GetActive() bool {
	if o == nil {
		var ret bool
		return ret
	}

	return o.Active
}

// GetActiveOk returns a tuple with the Active field value
// and a boolean to check if the value has been set.
func (o *EmbedSettingsLogosCustom) GetActiveOk() (*bool, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Active, true
}

// SetActive sets field value
func (o *EmbedSettingsLogosCustom) SetActive(v bool) {
	o.Active = v
}

// GetLink returns the Link field value
// If the value is explicit nil, the zero value for string will be returned
func (o *EmbedSettingsLogosCustom) GetLink() string {
	if o == nil || o.Link.Get() == nil {
		var ret string
		return ret
	}

	return *o.Link.Get()
}

// GetLinkOk returns a tuple with the Link field value
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *EmbedSettingsLogosCustom) GetLinkOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.Link.Get(), o.Link.IsSet()
}

// SetLink sets field value
func (o *EmbedSettingsLogosCustom) SetLink(v string) {
	o.Link.Set(&v)
}

// GetSticky returns the Sticky field value
func (o *EmbedSettingsLogosCustom) GetSticky() bool {
	if o == nil {
		var ret bool
		return ret
	}

	return o.Sticky
}

// GetStickyOk returns a tuple with the Sticky field value
// and a boolean to check if the value has been set.
func (o *EmbedSettingsLogosCustom) GetStickyOk() (*bool, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Sticky, true
}

// SetSticky sets field value
func (o *EmbedSettingsLogosCustom) SetSticky(v bool) {
	o.Sticky = v
}

// GetUrl returns the Url field value
// If the value is explicit nil, the zero value for string will be returned
func (o *EmbedSettingsLogosCustom) GetUrl() string {
	if o == nil || o.Url.Get() == nil {
		var ret string
		return ret
	}

	return *o.Url.Get()
}

// GetUrlOk returns a tuple with the Url field value
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *EmbedSettingsLogosCustom) GetUrlOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.Url.Get(), o.Url.IsSet()
}

// SetUrl sets field value
func (o *EmbedSettingsLogosCustom) SetUrl(v string) {
	o.Url.Set(&v)
}

// GetUseLink returns the UseLink field value
func (o *EmbedSettingsLogosCustom) GetUseLink() bool {
	if o == nil {
		var ret bool
		return ret
	}

	return o.UseLink
}

// GetUseLinkOk returns a tuple with the UseLink field value
// and a boolean to check if the value has been set.
func (o *EmbedSettingsLogosCustom) GetUseLinkOk() (*bool, bool) {
	if o == nil {
		return nil, false
	}
	return &o.UseLink, true
}

// SetUseLink sets field value
func (o *EmbedSettingsLogosCustom) SetUseLink(v bool) {
	o.UseLink = v
}

func (o EmbedSettingsLogosCustom) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o EmbedSettingsLogosCustom) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["active"] = o.Active
	toSerialize["link"] = o.Link.Get()
	toSerialize["sticky"] = o.Sticky
	toSerialize["url"] = o.Url.Get()
	toSerialize["use_link"] = o.UseLink
	return toSerialize, nil
}

type NullableEmbedSettingsLogosCustom struct {
	value *EmbedSettingsLogosCustom
	isSet bool
}

func (v NullableEmbedSettingsLogosCustom) Get() *EmbedSettingsLogosCustom {
	return v.value
}

func (v *NullableEmbedSettingsLogosCustom) Set(val *EmbedSettingsLogosCustom) {
	v.value = val
	v.isSet = true
}

func (v NullableEmbedSettingsLogosCustom) IsSet() bool {
	return v.isSet
}

func (v *NullableEmbedSettingsLogosCustom) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableEmbedSettingsLogosCustom(val *EmbedSettingsLogosCustom) *NullableEmbedSettingsLogosCustom {
	return &NullableEmbedSettingsLogosCustom{value: val, isSet: true}
}

func (v NullableEmbedSettingsLogosCustom) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableEmbedSettingsLogosCustom) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


