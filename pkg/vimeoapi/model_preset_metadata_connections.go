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

// checks if the PresetMetadataConnections type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &PresetMetadataConnections{}

// PresetMetadataConnections A list of resource URIs related to the preset.
type PresetMetadataConnections struct {
	Videos PresetMetadataConnectionsVideos `json:"videos"`
}

type _PresetMetadataConnections PresetMetadataConnections

// NewPresetMetadataConnections instantiates a new PresetMetadataConnections object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewPresetMetadataConnections(videos PresetMetadataConnectionsVideos) *PresetMetadataConnections {
	this := PresetMetadataConnections{}
	this.Videos = videos
	return &this
}

// NewPresetMetadataConnectionsWithDefaults instantiates a new PresetMetadataConnections object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewPresetMetadataConnectionsWithDefaults() *PresetMetadataConnections {
	this := PresetMetadataConnections{}
	return &this
}

// GetVideos returns the Videos field value
func (o *PresetMetadataConnections) GetVideos() PresetMetadataConnectionsVideos {
	if o == nil {
		var ret PresetMetadataConnectionsVideos
		return ret
	}

	return o.Videos
}

// GetVideosOk returns a tuple with the Videos field value
// and a boolean to check if the value has been set.
func (o *PresetMetadataConnections) GetVideosOk() (*PresetMetadataConnectionsVideos, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Videos, true
}

// SetVideos sets field value
func (o *PresetMetadataConnections) SetVideos(v PresetMetadataConnectionsVideos) {
	o.Videos = v
}

func (o PresetMetadataConnections) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o PresetMetadataConnections) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["videos"] = o.Videos
	return toSerialize, nil
}

func (o *PresetMetadataConnections) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"videos",
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

	varPresetMetadataConnections := _PresetMetadataConnections{}

	decoder := json.NewDecoder(bytes.NewReader(data))
	decoder.DisallowUnknownFields()
	err = decoder.Decode(&varPresetMetadataConnections)

	if err != nil {
		return err
	}

	*o = PresetMetadataConnections(varPresetMetadataConnections)

	return err
}

type NullablePresetMetadataConnections struct {
	value *PresetMetadataConnections
	isSet bool
}

func (v NullablePresetMetadataConnections) Get() *PresetMetadataConnections {
	return v.value
}

func (v *NullablePresetMetadataConnections) Set(val *PresetMetadataConnections) {
	v.value = val
	v.isSet = true
}

func (v NullablePresetMetadataConnections) IsSet() bool {
	return v.isSet
}

func (v *NullablePresetMetadataConnections) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullablePresetMetadataConnections(val *PresetMetadataConnections) *NullablePresetMetadataConnections {
	return &NullablePresetMetadataConnections{value: val, isSet: true}
}

func (v NullablePresetMetadataConnections) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullablePresetMetadataConnections) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

