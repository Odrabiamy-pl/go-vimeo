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

// checks if the CategoryMetadata type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &CategoryMetadata{}

// CategoryMetadata Metadata about the category.
type CategoryMetadata struct {
	Connections CategoryMetadataConnections `json:"connections"`
	Interactions CategoryMetadataInteractions `json:"interactions"`
}

// NewCategoryMetadata instantiates a new CategoryMetadata object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewCategoryMetadata(connections CategoryMetadataConnections, interactions CategoryMetadataInteractions) *CategoryMetadata {
	this := CategoryMetadata{}
	this.Connections = connections
	this.Interactions = interactions
	return &this
}

// NewCategoryMetadataWithDefaults instantiates a new CategoryMetadata object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewCategoryMetadataWithDefaults() *CategoryMetadata {
	this := CategoryMetadata{}
	return &this
}

// GetConnections returns the Connections field value
func (o *CategoryMetadata) GetConnections() CategoryMetadataConnections {
	if o == nil {
		var ret CategoryMetadataConnections
		return ret
	}

	return o.Connections
}

// GetConnectionsOk returns a tuple with the Connections field value
// and a boolean to check if the value has been set.
func (o *CategoryMetadata) GetConnectionsOk() (*CategoryMetadataConnections, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Connections, true
}

// SetConnections sets field value
func (o *CategoryMetadata) SetConnections(v CategoryMetadataConnections) {
	o.Connections = v
}

// GetInteractions returns the Interactions field value
func (o *CategoryMetadata) GetInteractions() CategoryMetadataInteractions {
	if o == nil {
		var ret CategoryMetadataInteractions
		return ret
	}

	return o.Interactions
}

// GetInteractionsOk returns a tuple with the Interactions field value
// and a boolean to check if the value has been set.
func (o *CategoryMetadata) GetInteractionsOk() (*CategoryMetadataInteractions, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Interactions, true
}

// SetInteractions sets field value
func (o *CategoryMetadata) SetInteractions(v CategoryMetadataInteractions) {
	o.Interactions = v
}

func (o CategoryMetadata) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o CategoryMetadata) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["connections"] = o.Connections
	toSerialize["interactions"] = o.Interactions
	return toSerialize, nil
}

type NullableCategoryMetadata struct {
	value *CategoryMetadata
	isSet bool
}

func (v NullableCategoryMetadata) Get() *CategoryMetadata {
	return v.value
}

func (v *NullableCategoryMetadata) Set(val *CategoryMetadata) {
	v.value = val
	v.isSet = true
}

func (v NullableCategoryMetadata) IsSet() bool {
	return v.isSet
}

func (v *NullableCategoryMetadata) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableCategoryMetadata(val *CategoryMetadata) *NullableCategoryMetadata {
	return &NullableCategoryMetadata{value: val, isSet: true}
}

func (v NullableCategoryMetadata) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableCategoryMetadata) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


