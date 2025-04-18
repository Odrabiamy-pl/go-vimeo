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

// checks if the GroupMetadata type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &GroupMetadata{}

// GroupMetadata Metadata about the group.
type GroupMetadata struct {
	Connections GroupMetadataConnections `json:"connections"`
	Interactions GroupMetadataInteractions `json:"interactions"`
}

// NewGroupMetadata instantiates a new GroupMetadata object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewGroupMetadata(connections GroupMetadataConnections, interactions GroupMetadataInteractions) *GroupMetadata {
	this := GroupMetadata{}
	this.Connections = connections
	this.Interactions = interactions
	return &this
}

// NewGroupMetadataWithDefaults instantiates a new GroupMetadata object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewGroupMetadataWithDefaults() *GroupMetadata {
	this := GroupMetadata{}
	return &this
}

// GetConnections returns the Connections field value
func (o *GroupMetadata) GetConnections() GroupMetadataConnections {
	if o == nil {
		var ret GroupMetadataConnections
		return ret
	}

	return o.Connections
}

// GetConnectionsOk returns a tuple with the Connections field value
// and a boolean to check if the value has been set.
func (o *GroupMetadata) GetConnectionsOk() (*GroupMetadataConnections, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Connections, true
}

// SetConnections sets field value
func (o *GroupMetadata) SetConnections(v GroupMetadataConnections) {
	o.Connections = v
}

// GetInteractions returns the Interactions field value
func (o *GroupMetadata) GetInteractions() GroupMetadataInteractions {
	if o == nil {
		var ret GroupMetadataInteractions
		return ret
	}

	return o.Interactions
}

// GetInteractionsOk returns a tuple with the Interactions field value
// and a boolean to check if the value has been set.
func (o *GroupMetadata) GetInteractionsOk() (*GroupMetadataInteractions, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Interactions, true
}

// SetInteractions sets field value
func (o *GroupMetadata) SetInteractions(v GroupMetadataInteractions) {
	o.Interactions = v
}

func (o GroupMetadata) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o GroupMetadata) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["connections"] = o.Connections
	toSerialize["interactions"] = o.Interactions
	return toSerialize, nil
}

type NullableGroupMetadata struct {
	value *GroupMetadata
	isSet bool
}

func (v NullableGroupMetadata) Get() *GroupMetadata {
	return v.value
}

func (v *NullableGroupMetadata) Set(val *GroupMetadata) {
	v.value = val
	v.isSet = true
}

func (v NullableGroupMetadata) IsSet() bool {
	return v.isSet
}

func (v *NullableGroupMetadata) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableGroupMetadata(val *GroupMetadata) *NullableGroupMetadata {
	return &NullableGroupMetadata{value: val, isSet: true}
}

func (v NullableGroupMetadata) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableGroupMetadata) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


