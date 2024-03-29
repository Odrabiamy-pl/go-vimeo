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

// checks if the VideoMetadataConnectionsTeamPermissions type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &VideoMetadataConnectionsTeamPermissions{}

// VideoMetadataConnectionsTeamPermissions Information about the video's team permissions list. This data requires a bearer token with the `private` scope.
type VideoMetadataConnectionsTeamPermissions struct {
	// An array of HTTP methods permitted on this URI. This data requires a bearer token with the `private` scope.
	Options []string `json:"options"`
}

// NewVideoMetadataConnectionsTeamPermissions instantiates a new VideoMetadataConnectionsTeamPermissions object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewVideoMetadataConnectionsTeamPermissions(options []string) *VideoMetadataConnectionsTeamPermissions {
	this := VideoMetadataConnectionsTeamPermissions{}
	this.Options = options
	return &this
}

// NewVideoMetadataConnectionsTeamPermissionsWithDefaults instantiates a new VideoMetadataConnectionsTeamPermissions object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewVideoMetadataConnectionsTeamPermissionsWithDefaults() *VideoMetadataConnectionsTeamPermissions {
	this := VideoMetadataConnectionsTeamPermissions{}
	return &this
}

// GetOptions returns the Options field value
func (o *VideoMetadataConnectionsTeamPermissions) GetOptions() []string {
	if o == nil {
		var ret []string
		return ret
	}

	return o.Options
}

// GetOptionsOk returns a tuple with the Options field value
// and a boolean to check if the value has been set.
func (o *VideoMetadataConnectionsTeamPermissions) GetOptionsOk() ([]string, bool) {
	if o == nil {
		return nil, false
	}
	return o.Options, true
}

// SetOptions sets field value
func (o *VideoMetadataConnectionsTeamPermissions) SetOptions(v []string) {
	o.Options = v
}

func (o VideoMetadataConnectionsTeamPermissions) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o VideoMetadataConnectionsTeamPermissions) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["options"] = o.Options
	return toSerialize, nil
}

type NullableVideoMetadataConnectionsTeamPermissions struct {
	value *VideoMetadataConnectionsTeamPermissions
	isSet bool
}

func (v NullableVideoMetadataConnectionsTeamPermissions) Get() *VideoMetadataConnectionsTeamPermissions {
	return v.value
}

func (v *NullableVideoMetadataConnectionsTeamPermissions) Set(val *VideoMetadataConnectionsTeamPermissions) {
	v.value = val
	v.isSet = true
}

func (v NullableVideoMetadataConnectionsTeamPermissions) IsSet() bool {
	return v.isSet
}

func (v *NullableVideoMetadataConnectionsTeamPermissions) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableVideoMetadataConnectionsTeamPermissions(val *VideoMetadataConnectionsTeamPermissions) *NullableVideoMetadataConnectionsTeamPermissions {
	return &NullableVideoMetadataConnectionsTeamPermissions{value: val, isSet: true}
}

func (v NullableVideoMetadataConnectionsTeamPermissions) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableVideoMetadataConnectionsTeamPermissions) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
