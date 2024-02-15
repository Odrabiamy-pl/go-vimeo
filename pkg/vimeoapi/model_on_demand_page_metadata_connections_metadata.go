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

// checks if the OnDemandPageMetadataConnectionsMetadata type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &OnDemandPageMetadataConnectionsMetadata{}

// OnDemandPageMetadataConnectionsMetadata struct for OnDemandPageMetadataConnectionsMetadata
type OnDemandPageMetadataConnectionsMetadata struct {
	Connections OnDemandPageMetadataConnectionsMetadataConnections `json:"connections"`
}

type _OnDemandPageMetadataConnectionsMetadata OnDemandPageMetadataConnectionsMetadata

// NewOnDemandPageMetadataConnectionsMetadata instantiates a new OnDemandPageMetadataConnectionsMetadata object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewOnDemandPageMetadataConnectionsMetadata(connections OnDemandPageMetadataConnectionsMetadataConnections) *OnDemandPageMetadataConnectionsMetadata {
	this := OnDemandPageMetadataConnectionsMetadata{}
	this.Connections = connections
	return &this
}

// NewOnDemandPageMetadataConnectionsMetadataWithDefaults instantiates a new OnDemandPageMetadataConnectionsMetadata object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewOnDemandPageMetadataConnectionsMetadataWithDefaults() *OnDemandPageMetadataConnectionsMetadata {
	this := OnDemandPageMetadataConnectionsMetadata{}
	return &this
}

// GetConnections returns the Connections field value
func (o *OnDemandPageMetadataConnectionsMetadata) GetConnections() OnDemandPageMetadataConnectionsMetadataConnections {
	if o == nil {
		var ret OnDemandPageMetadataConnectionsMetadataConnections
		return ret
	}

	return o.Connections
}

// GetConnectionsOk returns a tuple with the Connections field value
// and a boolean to check if the value has been set.
func (o *OnDemandPageMetadataConnectionsMetadata) GetConnectionsOk() (*OnDemandPageMetadataConnectionsMetadataConnections, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Connections, true
}

// SetConnections sets field value
func (o *OnDemandPageMetadataConnectionsMetadata) SetConnections(v OnDemandPageMetadataConnectionsMetadataConnections) {
	o.Connections = v
}

func (o OnDemandPageMetadataConnectionsMetadata) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o OnDemandPageMetadataConnectionsMetadata) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["connections"] = o.Connections
	return toSerialize, nil
}

func (o *OnDemandPageMetadataConnectionsMetadata) UnmarshalJSON(data []byte) (err error) {
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

	varOnDemandPageMetadataConnectionsMetadata := _OnDemandPageMetadataConnectionsMetadata{}

	decoder := json.NewDecoder(bytes.NewReader(data))
	decoder.DisallowUnknownFields()
	err = decoder.Decode(&varOnDemandPageMetadataConnectionsMetadata)

	if err != nil {
		return err
	}

	*o = OnDemandPageMetadataConnectionsMetadata(varOnDemandPageMetadataConnectionsMetadata)

	return err
}

type NullableOnDemandPageMetadataConnectionsMetadata struct {
	value *OnDemandPageMetadataConnectionsMetadata
	isSet bool
}

func (v NullableOnDemandPageMetadataConnectionsMetadata) Get() *OnDemandPageMetadataConnectionsMetadata {
	return v.value
}

func (v *NullableOnDemandPageMetadataConnectionsMetadata) Set(val *OnDemandPageMetadataConnectionsMetadata) {
	v.value = val
	v.isSet = true
}

func (v NullableOnDemandPageMetadataConnectionsMetadata) IsSet() bool {
	return v.isSet
}

func (v *NullableOnDemandPageMetadataConnectionsMetadata) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableOnDemandPageMetadataConnectionsMetadata(val *OnDemandPageMetadataConnectionsMetadata) *NullableOnDemandPageMetadataConnectionsMetadata {
	return &NullableOnDemandPageMetadataConnectionsMetadata{value: val, isSet: true}
}

func (v NullableOnDemandPageMetadataConnectionsMetadata) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableOnDemandPageMetadataConnectionsMetadata) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

