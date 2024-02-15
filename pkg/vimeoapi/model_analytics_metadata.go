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

// checks if the AnalyticsMetadata type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &AnalyticsMetadata{}

// AnalyticsMetadata struct for AnalyticsMetadata
type AnalyticsMetadata struct {
	Connections AnalyticsMetadataConnections `json:"connections"`
}

type _AnalyticsMetadata AnalyticsMetadata

// NewAnalyticsMetadata instantiates a new AnalyticsMetadata object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewAnalyticsMetadata(connections AnalyticsMetadataConnections) *AnalyticsMetadata {
	this := AnalyticsMetadata{}
	this.Connections = connections
	return &this
}

// NewAnalyticsMetadataWithDefaults instantiates a new AnalyticsMetadata object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewAnalyticsMetadataWithDefaults() *AnalyticsMetadata {
	this := AnalyticsMetadata{}
	return &this
}

// GetConnections returns the Connections field value
func (o *AnalyticsMetadata) GetConnections() AnalyticsMetadataConnections {
	if o == nil {
		var ret AnalyticsMetadataConnections
		return ret
	}

	return o.Connections
}

// GetConnectionsOk returns a tuple with the Connections field value
// and a boolean to check if the value has been set.
func (o *AnalyticsMetadata) GetConnectionsOk() (*AnalyticsMetadataConnections, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Connections, true
}

// SetConnections sets field value
func (o *AnalyticsMetadata) SetConnections(v AnalyticsMetadataConnections) {
	o.Connections = v
}

func (o AnalyticsMetadata) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o AnalyticsMetadata) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["connections"] = o.Connections
	return toSerialize, nil
}

func (o *AnalyticsMetadata) UnmarshalJSON(data []byte) (err error) {
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

	varAnalyticsMetadata := _AnalyticsMetadata{}

	decoder := json.NewDecoder(bytes.NewReader(data))
	decoder.DisallowUnknownFields()
	err = decoder.Decode(&varAnalyticsMetadata)

	if err != nil {
		return err
	}

	*o = AnalyticsMetadata(varAnalyticsMetadata)

	return err
}

type NullableAnalyticsMetadata struct {
	value *AnalyticsMetadata
	isSet bool
}

func (v NullableAnalyticsMetadata) Get() *AnalyticsMetadata {
	return v.value
}

func (v *NullableAnalyticsMetadata) Set(val *AnalyticsMetadata) {
	v.value = val
	v.isSet = true
}

func (v NullableAnalyticsMetadata) IsSet() bool {
	return v.isSet
}

func (v *NullableAnalyticsMetadata) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableAnalyticsMetadata(val *AnalyticsMetadata) *NullableAnalyticsMetadata {
	return &NullableAnalyticsMetadata{value: val, isSet: true}
}

func (v NullableAnalyticsMetadata) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableAnalyticsMetadata) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

