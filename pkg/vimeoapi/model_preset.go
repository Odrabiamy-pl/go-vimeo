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

// checks if the Preset type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &Preset{}

// Preset struct for Preset
type Preset struct {
	Metadata PresetMetadata `json:"metadata"`
	// The display name of the preset group.
	Name string `json:"name"`
	Settings PresetSettings `json:"settings"`
	// The canonical relative URI of the preset object.
	Uri string `json:"uri"`
	User NullablePresetUser `json:"user"`
}

// NewPreset instantiates a new Preset object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewPreset(metadata PresetMetadata, name string, settings PresetSettings, uri string, user NullablePresetUser) *Preset {
	this := Preset{}
	this.Metadata = metadata
	this.Name = name
	this.Settings = settings
	this.Uri = uri
	this.User = user
	return &this
}

// NewPresetWithDefaults instantiates a new Preset object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewPresetWithDefaults() *Preset {
	this := Preset{}
	return &this
}

// GetMetadata returns the Metadata field value
func (o *Preset) GetMetadata() PresetMetadata {
	if o == nil {
		var ret PresetMetadata
		return ret
	}

	return o.Metadata
}

// GetMetadataOk returns a tuple with the Metadata field value
// and a boolean to check if the value has been set.
func (o *Preset) GetMetadataOk() (*PresetMetadata, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Metadata, true
}

// SetMetadata sets field value
func (o *Preset) SetMetadata(v PresetMetadata) {
	o.Metadata = v
}

// GetName returns the Name field value
func (o *Preset) GetName() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Name
}

// GetNameOk returns a tuple with the Name field value
// and a boolean to check if the value has been set.
func (o *Preset) GetNameOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Name, true
}

// SetName sets field value
func (o *Preset) SetName(v string) {
	o.Name = v
}

// GetSettings returns the Settings field value
func (o *Preset) GetSettings() PresetSettings {
	if o == nil {
		var ret PresetSettings
		return ret
	}

	return o.Settings
}

// GetSettingsOk returns a tuple with the Settings field value
// and a boolean to check if the value has been set.
func (o *Preset) GetSettingsOk() (*PresetSettings, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Settings, true
}

// SetSettings sets field value
func (o *Preset) SetSettings(v PresetSettings) {
	o.Settings = v
}

// GetUri returns the Uri field value
func (o *Preset) GetUri() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Uri
}

// GetUriOk returns a tuple with the Uri field value
// and a boolean to check if the value has been set.
func (o *Preset) GetUriOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Uri, true
}

// SetUri sets field value
func (o *Preset) SetUri(v string) {
	o.Uri = v
}

// GetUser returns the User field value
// If the value is explicit nil, the zero value for PresetUser will be returned
func (o *Preset) GetUser() PresetUser {
	if o == nil || o.User.Get() == nil {
		var ret PresetUser
		return ret
	}

	return *o.User.Get()
}

// GetUserOk returns a tuple with the User field value
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *Preset) GetUserOk() (*PresetUser, bool) {
	if o == nil {
		return nil, false
	}
	return o.User.Get(), o.User.IsSet()
}

// SetUser sets field value
func (o *Preset) SetUser(v PresetUser) {
	o.User.Set(&v)
}

func (o Preset) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o Preset) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["metadata"] = o.Metadata
	toSerialize["name"] = o.Name
	toSerialize["settings"] = o.Settings
	toSerialize["uri"] = o.Uri
	toSerialize["user"] = o.User.Get()
	return toSerialize, nil
}

type NullablePreset struct {
	value *Preset
	isSet bool
}

func (v NullablePreset) Get() *Preset {
	return v.value
}

func (v *NullablePreset) Set(val *Preset) {
	v.value = val
	v.isSet = true
}

func (v NullablePreset) IsSet() bool {
	return v.isSet
}

func (v *NullablePreset) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullablePreset(val *Preset) *NullablePreset {
	return &NullablePreset{value: val, isSet: true}
}

func (v NullablePreset) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullablePreset) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


