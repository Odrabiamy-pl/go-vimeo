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

// checks if the PortfolioMetadata type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &PortfolioMetadata{}

// PortfolioMetadata Metadata about the portfolio.
type PortfolioMetadata struct {
	Connections PortfolioMetadataConnections `json:"connections"`
}

type _PortfolioMetadata PortfolioMetadata

// NewPortfolioMetadata instantiates a new PortfolioMetadata object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewPortfolioMetadata(connections PortfolioMetadataConnections) *PortfolioMetadata {
	this := PortfolioMetadata{}
	this.Connections = connections
	return &this
}

// NewPortfolioMetadataWithDefaults instantiates a new PortfolioMetadata object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewPortfolioMetadataWithDefaults() *PortfolioMetadata {
	this := PortfolioMetadata{}
	return &this
}

// GetConnections returns the Connections field value
func (o *PortfolioMetadata) GetConnections() PortfolioMetadataConnections {
	if o == nil {
		var ret PortfolioMetadataConnections
		return ret
	}

	return o.Connections
}

// GetConnectionsOk returns a tuple with the Connections field value
// and a boolean to check if the value has been set.
func (o *PortfolioMetadata) GetConnectionsOk() (*PortfolioMetadataConnections, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Connections, true
}

// SetConnections sets field value
func (o *PortfolioMetadata) SetConnections(v PortfolioMetadataConnections) {
	o.Connections = v
}

func (o PortfolioMetadata) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o PortfolioMetadata) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["connections"] = o.Connections
	return toSerialize, nil
}

func (o *PortfolioMetadata) UnmarshalJSON(data []byte) (err error) {
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

	varPortfolioMetadata := _PortfolioMetadata{}

	decoder := json.NewDecoder(bytes.NewReader(data))
	decoder.DisallowUnknownFields()
	err = decoder.Decode(&varPortfolioMetadata)

	if err != nil {
		return err
	}

	*o = PortfolioMetadata(varPortfolioMetadata)

	return err
}

type NullablePortfolioMetadata struct {
	value *PortfolioMetadata
	isSet bool
}

func (v NullablePortfolioMetadata) Get() *PortfolioMetadata {
	return v.value
}

func (v *NullablePortfolioMetadata) Set(val *PortfolioMetadata) {
	v.value = val
	v.isSet = true
}

func (v NullablePortfolioMetadata) IsSet() bool {
	return v.isSet
}

func (v *NullablePortfolioMetadata) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullablePortfolioMetadata(val *PortfolioMetadata) *NullablePortfolioMetadata {
	return &NullablePortfolioMetadata{value: val, isSet: true}
}

func (v NullablePortfolioMetadata) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullablePortfolioMetadata) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


