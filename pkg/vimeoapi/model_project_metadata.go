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

// checks if the ProjectMetadata type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &ProjectMetadata{}

// ProjectMetadata Information about the folders's metadata.
type ProjectMetadata struct {
	Connections ProjectMetadataConnections `json:"connections"`
	Interactions ProjectMetadataInteractions `json:"interactions"`
}

type _ProjectMetadata ProjectMetadata

// NewProjectMetadata instantiates a new ProjectMetadata object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewProjectMetadata(connections ProjectMetadataConnections, interactions ProjectMetadataInteractions) *ProjectMetadata {
	this := ProjectMetadata{}
	this.Connections = connections
	this.Interactions = interactions
	return &this
}

// NewProjectMetadataWithDefaults instantiates a new ProjectMetadata object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewProjectMetadataWithDefaults() *ProjectMetadata {
	this := ProjectMetadata{}
	return &this
}

// GetConnections returns the Connections field value
func (o *ProjectMetadata) GetConnections() ProjectMetadataConnections {
	if o == nil {
		var ret ProjectMetadataConnections
		return ret
	}

	return o.Connections
}

// GetConnectionsOk returns a tuple with the Connections field value
// and a boolean to check if the value has been set.
func (o *ProjectMetadata) GetConnectionsOk() (*ProjectMetadataConnections, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Connections, true
}

// SetConnections sets field value
func (o *ProjectMetadata) SetConnections(v ProjectMetadataConnections) {
	o.Connections = v
}

// GetInteractions returns the Interactions field value
func (o *ProjectMetadata) GetInteractions() ProjectMetadataInteractions {
	if o == nil {
		var ret ProjectMetadataInteractions
		return ret
	}

	return o.Interactions
}

// GetInteractionsOk returns a tuple with the Interactions field value
// and a boolean to check if the value has been set.
func (o *ProjectMetadata) GetInteractionsOk() (*ProjectMetadataInteractions, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Interactions, true
}

// SetInteractions sets field value
func (o *ProjectMetadata) SetInteractions(v ProjectMetadataInteractions) {
	o.Interactions = v
}

func (o ProjectMetadata) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o ProjectMetadata) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["connections"] = o.Connections
	toSerialize["interactions"] = o.Interactions
	return toSerialize, nil
}

func (o *ProjectMetadata) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"connections",
		"interactions",
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

	varProjectMetadata := _ProjectMetadata{}

	decoder := json.NewDecoder(bytes.NewReader(data))
	decoder.DisallowUnknownFields()
	err = decoder.Decode(&varProjectMetadata)

	if err != nil {
		return err
	}

	*o = ProjectMetadata(varProjectMetadata)

	return err
}

type NullableProjectMetadata struct {
	value *ProjectMetadata
	isSet bool
}

func (v NullableProjectMetadata) Get() *ProjectMetadata {
	return v.value
}

func (v *NullableProjectMetadata) Set(val *ProjectMetadata) {
	v.value = val
	v.isSet = true
}

func (v NullableProjectMetadata) IsSet() bool {
	return v.isSet
}

func (v *NullableProjectMetadata) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableProjectMetadata(val *ProjectMetadata) *NullableProjectMetadata {
	return &NullableProjectMetadata{value: val, isSet: true}
}

func (v NullableProjectMetadata) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableProjectMetadata) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

