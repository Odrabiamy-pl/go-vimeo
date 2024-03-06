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

// checks if the ProjectMetadataConnectionsTeamPermissions type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &ProjectMetadataConnectionsTeamPermissions{}

// ProjectMetadataConnectionsTeamPermissions Information about the folder's team permissions list. This data requires a bearer token with the `private` scope.
type ProjectMetadataConnectionsTeamPermissions struct {
	// An array of HTTP methods permitted on this URI. This data requires a bearer token with the `private` scope.
	Options []string `json:"options"`
}

// NewProjectMetadataConnectionsTeamPermissions instantiates a new ProjectMetadataConnectionsTeamPermissions object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewProjectMetadataConnectionsTeamPermissions(options []string) *ProjectMetadataConnectionsTeamPermissions {
	this := ProjectMetadataConnectionsTeamPermissions{}
	this.Options = options
	return &this
}

// NewProjectMetadataConnectionsTeamPermissionsWithDefaults instantiates a new ProjectMetadataConnectionsTeamPermissions object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewProjectMetadataConnectionsTeamPermissionsWithDefaults() *ProjectMetadataConnectionsTeamPermissions {
	this := ProjectMetadataConnectionsTeamPermissions{}
	return &this
}

// GetOptions returns the Options field value
func (o *ProjectMetadataConnectionsTeamPermissions) GetOptions() []string {
	if o == nil {
		var ret []string
		return ret
	}

	return o.Options
}

// GetOptionsOk returns a tuple with the Options field value
// and a boolean to check if the value has been set.
func (o *ProjectMetadataConnectionsTeamPermissions) GetOptionsOk() ([]string, bool) {
	if o == nil {
		return nil, false
	}
	return o.Options, true
}

// SetOptions sets field value
func (o *ProjectMetadataConnectionsTeamPermissions) SetOptions(v []string) {
	o.Options = v
}

func (o ProjectMetadataConnectionsTeamPermissions) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o ProjectMetadataConnectionsTeamPermissions) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["options"] = o.Options
	return toSerialize, nil
}

type NullableProjectMetadataConnectionsTeamPermissions struct {
	value *ProjectMetadataConnectionsTeamPermissions
	isSet bool
}

func (v NullableProjectMetadataConnectionsTeamPermissions) Get() *ProjectMetadataConnectionsTeamPermissions {
	return v.value
}

func (v *NullableProjectMetadataConnectionsTeamPermissions) Set(val *ProjectMetadataConnectionsTeamPermissions) {
	v.value = val
	v.isSet = true
}

func (v NullableProjectMetadataConnectionsTeamPermissions) IsSet() bool {
	return v.isSet
}

func (v *NullableProjectMetadataConnectionsTeamPermissions) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableProjectMetadataConnectionsTeamPermissions(val *ProjectMetadataConnectionsTeamPermissions) *NullableProjectMetadataConnectionsTeamPermissions {
	return &NullableProjectMetadataConnectionsTeamPermissions{value: val, isSet: true}
}

func (v NullableProjectMetadataConnectionsTeamPermissions) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableProjectMetadataConnectionsTeamPermissions) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
