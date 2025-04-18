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

// checks if the PresetMetadata type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &PresetMetadata{}

// PresetMetadata Metadata about the preset.
type PresetMetadata struct {
	Connections PresetMetadataConnections `json:"connections"`
}

// NewPresetMetadata instantiates a new PresetMetadata object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewPresetMetadata(connections PresetMetadataConnections) *PresetMetadata {
	this := PresetMetadata{}
	this.Connections = connections
	return &this
}

// NewPresetMetadataWithDefaults instantiates a new PresetMetadata object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewPresetMetadataWithDefaults() *PresetMetadata {
	this := PresetMetadata{}
	return &this
}

// GetConnections returns the Connections field value
func (o *PresetMetadata) GetConnections() PresetMetadataConnections {
	if o == nil {
		var ret PresetMetadataConnections
		return ret
	}

	return o.Connections
}

// GetConnectionsOk returns a tuple with the Connections field value
// and a boolean to check if the value has been set.
func (o *PresetMetadata) GetConnectionsOk() (*PresetMetadataConnections, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Connections, true
}

// SetConnections sets field value
func (o *PresetMetadata) SetConnections(v PresetMetadataConnections) {
	o.Connections = v
}

func (o PresetMetadata) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o PresetMetadata) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["connections"] = o.Connections
	return toSerialize, nil
}

type NullablePresetMetadata struct {
	value *PresetMetadata
	isSet bool
}

func (v NullablePresetMetadata) Get() *PresetMetadata {
	return v.value
}

func (v *NullablePresetMetadata) Set(val *PresetMetadata) {
	v.value = val
	v.isSet = true
}

func (v NullablePresetMetadata) IsSet() bool {
	return v.isSet
}

func (v *NullablePresetMetadata) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullablePresetMetadata(val *PresetMetadata) *NullablePresetMetadata {
	return &NullablePresetMetadata{value: val, isSet: true}
}

func (v NullablePresetMetadata) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullablePresetMetadata) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


