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

// checks if the LiveEventMetadata type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &LiveEventMetadata{}

// LiveEventMetadata Metadata about the event.
type LiveEventMetadata struct {
	Connections  LiveEventMetadataConnections  `json:"connections"`
	Interactions LiveEventMetadataInteractions `json:"interactions"`
}

// NewLiveEventMetadata instantiates a new LiveEventMetadata object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewLiveEventMetadata(connections LiveEventMetadataConnections, interactions LiveEventMetadataInteractions) *LiveEventMetadata {
	this := LiveEventMetadata{}
	this.Connections = connections
	this.Interactions = interactions
	return &this
}

// NewLiveEventMetadataWithDefaults instantiates a new LiveEventMetadata object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewLiveEventMetadataWithDefaults() *LiveEventMetadata {
	this := LiveEventMetadata{}
	return &this
}

// GetConnections returns the Connections field value
func (o *LiveEventMetadata) GetConnections() LiveEventMetadataConnections {
	if o == nil {
		var ret LiveEventMetadataConnections
		return ret
	}

	return o.Connections
}

// GetConnectionsOk returns a tuple with the Connections field value
// and a boolean to check if the value has been set.
func (o *LiveEventMetadata) GetConnectionsOk() (*LiveEventMetadataConnections, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Connections, true
}

// SetConnections sets field value
func (o *LiveEventMetadata) SetConnections(v LiveEventMetadataConnections) {
	o.Connections = v
}

// GetInteractions returns the Interactions field value
func (o *LiveEventMetadata) GetInteractions() LiveEventMetadataInteractions {
	if o == nil {
		var ret LiveEventMetadataInteractions
		return ret
	}

	return o.Interactions
}

// GetInteractionsOk returns a tuple with the Interactions field value
// and a boolean to check if the value has been set.
func (o *LiveEventMetadata) GetInteractionsOk() (*LiveEventMetadataInteractions, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Interactions, true
}

// SetInteractions sets field value
func (o *LiveEventMetadata) SetInteractions(v LiveEventMetadataInteractions) {
	o.Interactions = v
}

func (o LiveEventMetadata) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o LiveEventMetadata) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["connections"] = o.Connections
	toSerialize["interactions"] = o.Interactions
	return toSerialize, nil
}

type NullableLiveEventMetadata struct {
	value *LiveEventMetadata
	isSet bool
}

func (v NullableLiveEventMetadata) Get() *LiveEventMetadata {
	return v.value
}

func (v *NullableLiveEventMetadata) Set(val *LiveEventMetadata) {
	v.value = val
	v.isSet = true
}

func (v NullableLiveEventMetadata) IsSet() bool {
	return v.isSet
}

func (v *NullableLiveEventMetadata) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableLiveEventMetadata(val *LiveEventMetadata) *NullableLiveEventMetadata {
	return &NullableLiveEventMetadata{value: val, isSet: true}
}

func (v NullableLiveEventMetadata) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableLiveEventMetadata) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
