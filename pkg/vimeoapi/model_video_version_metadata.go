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

// checks if the VideoVersionMetadata type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &VideoVersionMetadata{}

// VideoVersionMetadata The video version's metadata.
type VideoVersionMetadata struct {
	Connections VideoVersionMetadataConnections `json:"connections"`
}

type _VideoVersionMetadata VideoVersionMetadata

// NewVideoVersionMetadata instantiates a new VideoVersionMetadata object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewVideoVersionMetadata(connections VideoVersionMetadataConnections) *VideoVersionMetadata {
	this := VideoVersionMetadata{}
	this.Connections = connections
	return &this
}

// NewVideoVersionMetadataWithDefaults instantiates a new VideoVersionMetadata object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewVideoVersionMetadataWithDefaults() *VideoVersionMetadata {
	this := VideoVersionMetadata{}
	return &this
}

// GetConnections returns the Connections field value
func (o *VideoVersionMetadata) GetConnections() VideoVersionMetadataConnections {
	if o == nil {
		var ret VideoVersionMetadataConnections
		return ret
	}

	return o.Connections
}

// GetConnectionsOk returns a tuple with the Connections field value
// and a boolean to check if the value has been set.
func (o *VideoVersionMetadata) GetConnectionsOk() (*VideoVersionMetadataConnections, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Connections, true
}

// SetConnections sets field value
func (o *VideoVersionMetadata) SetConnections(v VideoVersionMetadataConnections) {
	o.Connections = v
}

func (o VideoVersionMetadata) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o VideoVersionMetadata) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["connections"] = o.Connections
	return toSerialize, nil
}

func (o *VideoVersionMetadata) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"connections",
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

	varVideoVersionMetadata := _VideoVersionMetadata{}

	decoder := json.NewDecoder(bytes.NewReader(data))
	decoder.DisallowUnknownFields()
	err = decoder.Decode(&varVideoVersionMetadata)

	if err != nil {
		return err
	}

	*o = VideoVersionMetadata(varVideoVersionMetadata)

	return err
}

type NullableVideoVersionMetadata struct {
	value *VideoVersionMetadata
	isSet bool
}

func (v NullableVideoVersionMetadata) Get() *VideoVersionMetadata {
	return v.value
}

func (v *NullableVideoVersionMetadata) Set(val *VideoVersionMetadata) {
	v.value = val
	v.isSet = true
}

func (v NullableVideoVersionMetadata) IsSet() bool {
	return v.isSet
}

func (v *NullableVideoVersionMetadata) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableVideoVersionMetadata(val *VideoVersionMetadata) *NullableVideoVersionMetadata {
	return &NullableVideoVersionMetadata{value: val, isSet: true}
}

func (v NullableVideoVersionMetadata) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableVideoVersionMetadata) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

