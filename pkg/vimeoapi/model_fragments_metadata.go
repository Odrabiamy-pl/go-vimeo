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

// checks if the FragmentsMetadata type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &FragmentsMetadata{}

// FragmentsMetadata Metadata about the fragments.
type FragmentsMetadata struct {
	Connections FragmentsMetadataConnections `json:"connections"`
}

type _FragmentsMetadata FragmentsMetadata

// NewFragmentsMetadata instantiates a new FragmentsMetadata object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewFragmentsMetadata(connections FragmentsMetadataConnections) *FragmentsMetadata {
	this := FragmentsMetadata{}
	this.Connections = connections
	return &this
}

// NewFragmentsMetadataWithDefaults instantiates a new FragmentsMetadata object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewFragmentsMetadataWithDefaults() *FragmentsMetadata {
	this := FragmentsMetadata{}
	return &this
}

// GetConnections returns the Connections field value
func (o *FragmentsMetadata) GetConnections() FragmentsMetadataConnections {
	if o == nil {
		var ret FragmentsMetadataConnections
		return ret
	}

	return o.Connections
}

// GetConnectionsOk returns a tuple with the Connections field value
// and a boolean to check if the value has been set.
func (o *FragmentsMetadata) GetConnectionsOk() (*FragmentsMetadataConnections, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Connections, true
}

// SetConnections sets field value
func (o *FragmentsMetadata) SetConnections(v FragmentsMetadataConnections) {
	o.Connections = v
}

func (o FragmentsMetadata) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o FragmentsMetadata) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["connections"] = o.Connections
	return toSerialize, nil
}

func (o *FragmentsMetadata) UnmarshalJSON(data []byte) (err error) {
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

	varFragmentsMetadata := _FragmentsMetadata{}

	decoder := json.NewDecoder(bytes.NewReader(data))
	decoder.DisallowUnknownFields()
	err = decoder.Decode(&varFragmentsMetadata)

	if err != nil {
		return err
	}

	*o = FragmentsMetadata(varFragmentsMetadata)

	return err
}

type NullableFragmentsMetadata struct {
	value *FragmentsMetadata
	isSet bool
}

func (v NullableFragmentsMetadata) Get() *FragmentsMetadata {
	return v.value
}

func (v *NullableFragmentsMetadata) Set(val *FragmentsMetadata) {
	v.value = val
	v.isSet = true
}

func (v NullableFragmentsMetadata) IsSet() bool {
	return v.isSet
}

func (v *NullableFragmentsMetadata) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableFragmentsMetadata(val *FragmentsMetadata) *NullableFragmentsMetadata {
	return &NullableFragmentsMetadata{value: val, isSet: true}
}

func (v NullableFragmentsMetadata) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableFragmentsMetadata) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


