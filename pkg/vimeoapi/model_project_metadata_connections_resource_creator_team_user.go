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

// checks if the ProjectMetadataConnectionsResourceCreatorTeamUser type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &ProjectMetadataConnectionsResourceCreatorTeamUser{}

// ProjectMetadataConnectionsResourceCreatorTeamUser Information about the team user who created the folder. This data requires a bearer token with the `private` scope.
type ProjectMetadataConnectionsResourceCreatorTeamUser struct {
	// The URI for the team user who created the folder. This data requires a bearer token with the `private` scope.
	Uri string `json:"uri"`
}

// NewProjectMetadataConnectionsResourceCreatorTeamUser instantiates a new ProjectMetadataConnectionsResourceCreatorTeamUser object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewProjectMetadataConnectionsResourceCreatorTeamUser(uri string) *ProjectMetadataConnectionsResourceCreatorTeamUser {
	this := ProjectMetadataConnectionsResourceCreatorTeamUser{}
	this.Uri = uri
	return &this
}

// NewProjectMetadataConnectionsResourceCreatorTeamUserWithDefaults instantiates a new ProjectMetadataConnectionsResourceCreatorTeamUser object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewProjectMetadataConnectionsResourceCreatorTeamUserWithDefaults() *ProjectMetadataConnectionsResourceCreatorTeamUser {
	this := ProjectMetadataConnectionsResourceCreatorTeamUser{}
	return &this
}

// GetUri returns the Uri field value
func (o *ProjectMetadataConnectionsResourceCreatorTeamUser) GetUri() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Uri
}

// GetUriOk returns a tuple with the Uri field value
// and a boolean to check if the value has been set.
func (o *ProjectMetadataConnectionsResourceCreatorTeamUser) GetUriOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Uri, true
}

// SetUri sets field value
func (o *ProjectMetadataConnectionsResourceCreatorTeamUser) SetUri(v string) {
	o.Uri = v
}

func (o ProjectMetadataConnectionsResourceCreatorTeamUser) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o ProjectMetadataConnectionsResourceCreatorTeamUser) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["uri"] = o.Uri
	return toSerialize, nil
}

type NullableProjectMetadataConnectionsResourceCreatorTeamUser struct {
	value *ProjectMetadataConnectionsResourceCreatorTeamUser
	isSet bool
}

func (v NullableProjectMetadataConnectionsResourceCreatorTeamUser) Get() *ProjectMetadataConnectionsResourceCreatorTeamUser {
	return v.value
}

func (v *NullableProjectMetadataConnectionsResourceCreatorTeamUser) Set(val *ProjectMetadataConnectionsResourceCreatorTeamUser) {
	v.value = val
	v.isSet = true
}

func (v NullableProjectMetadataConnectionsResourceCreatorTeamUser) IsSet() bool {
	return v.isSet
}

func (v *NullableProjectMetadataConnectionsResourceCreatorTeamUser) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableProjectMetadataConnectionsResourceCreatorTeamUser(val *ProjectMetadataConnectionsResourceCreatorTeamUser) *NullableProjectMetadataConnectionsResourceCreatorTeamUser {
	return &NullableProjectMetadataConnectionsResourceCreatorTeamUser{value: val, isSet: true}
}

func (v NullableProjectMetadataConnectionsResourceCreatorTeamUser) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableProjectMetadataConnectionsResourceCreatorTeamUser) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


