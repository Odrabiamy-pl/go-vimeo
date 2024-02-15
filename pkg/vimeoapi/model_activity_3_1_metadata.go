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

// checks if the Activity31Metadata type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &Activity31Metadata{}

// Activity31Metadata Information about the activity's metadata.
type Activity31Metadata struct {
	Connections Activity31MetadataConnections `json:"connections"`
}

// NewActivity31Metadata instantiates a new Activity31Metadata object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewActivity31Metadata(connections Activity31MetadataConnections) *Activity31Metadata {
	this := Activity31Metadata{}
	this.Connections = connections
	return &this
}

// NewActivity31MetadataWithDefaults instantiates a new Activity31Metadata object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewActivity31MetadataWithDefaults() *Activity31Metadata {
	this := Activity31Metadata{}
	return &this
}

// GetConnections returns the Connections field value
func (o *Activity31Metadata) GetConnections() Activity31MetadataConnections {
	if o == nil {
		var ret Activity31MetadataConnections
		return ret
	}

	return o.Connections
}

// GetConnectionsOk returns a tuple with the Connections field value
// and a boolean to check if the value has been set.
func (o *Activity31Metadata) GetConnectionsOk() (*Activity31MetadataConnections, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Connections, true
}

// SetConnections sets field value
func (o *Activity31Metadata) SetConnections(v Activity31MetadataConnections) {
	o.Connections = v
}

func (o Activity31Metadata) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o Activity31Metadata) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["connections"] = o.Connections
	return toSerialize, nil
}

type NullableActivity31Metadata struct {
	value *Activity31Metadata
	isSet bool
}

func (v NullableActivity31Metadata) Get() *Activity31Metadata {
	return v.value
}

func (v *NullableActivity31Metadata) Set(val *Activity31Metadata) {
	v.value = val
	v.isSet = true
}

func (v NullableActivity31Metadata) IsSet() bool {
	return v.isSet
}

func (v *NullableActivity31Metadata) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableActivity31Metadata(val *Activity31Metadata) *NullableActivity31Metadata {
	return &NullableActivity31Metadata{value: val, isSet: true}
}

func (v NullableActivity31Metadata) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableActivity31Metadata) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


