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

// checks if the FragmentsMetadataConnections type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &FragmentsMetadataConnections{}

// FragmentsMetadataConnections A collection of information that is connected to this resource.
type FragmentsMetadataConnections struct {
	Clip FragmentsMetadataConnectionsClip `json:"clip"`
}

// NewFragmentsMetadataConnections instantiates a new FragmentsMetadataConnections object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewFragmentsMetadataConnections(clip FragmentsMetadataConnectionsClip) *FragmentsMetadataConnections {
	this := FragmentsMetadataConnections{}
	this.Clip = clip
	return &this
}

// NewFragmentsMetadataConnectionsWithDefaults instantiates a new FragmentsMetadataConnections object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewFragmentsMetadataConnectionsWithDefaults() *FragmentsMetadataConnections {
	this := FragmentsMetadataConnections{}
	return &this
}

// GetClip returns the Clip field value
func (o *FragmentsMetadataConnections) GetClip() FragmentsMetadataConnectionsClip {
	if o == nil {
		var ret FragmentsMetadataConnectionsClip
		return ret
	}

	return o.Clip
}

// GetClipOk returns a tuple with the Clip field value
// and a boolean to check if the value has been set.
func (o *FragmentsMetadataConnections) GetClipOk() (*FragmentsMetadataConnectionsClip, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Clip, true
}

// SetClip sets field value
func (o *FragmentsMetadataConnections) SetClip(v FragmentsMetadataConnectionsClip) {
	o.Clip = v
}

func (o FragmentsMetadataConnections) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o FragmentsMetadataConnections) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["clip"] = o.Clip
	return toSerialize, nil
}

type NullableFragmentsMetadataConnections struct {
	value *FragmentsMetadataConnections
	isSet bool
}

func (v NullableFragmentsMetadataConnections) Get() *FragmentsMetadataConnections {
	return v.value
}

func (v *NullableFragmentsMetadataConnections) Set(val *FragmentsMetadataConnections) {
	v.value = val
	v.isSet = true
}

func (v NullableFragmentsMetadataConnections) IsSet() bool {
	return v.isSet
}

func (v *NullableFragmentsMetadataConnections) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableFragmentsMetadataConnections(val *FragmentsMetadataConnections) *NullableFragmentsMetadataConnections {
	return &NullableFragmentsMetadataConnections{value: val, isSet: true}
}

func (v NullableFragmentsMetadataConnections) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableFragmentsMetadataConnections) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


