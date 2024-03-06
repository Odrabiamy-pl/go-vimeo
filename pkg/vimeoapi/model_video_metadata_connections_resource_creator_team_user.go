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

// checks if the VideoMetadataConnectionsResourceCreatorTeamUser type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &VideoMetadataConnectionsResourceCreatorTeamUser{}

// VideoMetadataConnectionsResourceCreatorTeamUser Information about the team user who uploaded the video. This data requires a bearer token with the `private` scope.
type VideoMetadataConnectionsResourceCreatorTeamUser struct {
	// The URI for the team user who uploaded the video. This data requires a bearer token with the `private` scope.
	Uri string `json:"uri"`
}

// NewVideoMetadataConnectionsResourceCreatorTeamUser instantiates a new VideoMetadataConnectionsResourceCreatorTeamUser object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewVideoMetadataConnectionsResourceCreatorTeamUser(uri string) *VideoMetadataConnectionsResourceCreatorTeamUser {
	this := VideoMetadataConnectionsResourceCreatorTeamUser{}
	this.Uri = uri
	return &this
}

// NewVideoMetadataConnectionsResourceCreatorTeamUserWithDefaults instantiates a new VideoMetadataConnectionsResourceCreatorTeamUser object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewVideoMetadataConnectionsResourceCreatorTeamUserWithDefaults() *VideoMetadataConnectionsResourceCreatorTeamUser {
	this := VideoMetadataConnectionsResourceCreatorTeamUser{}
	return &this
}

// GetUri returns the Uri field value
func (o *VideoMetadataConnectionsResourceCreatorTeamUser) GetUri() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Uri
}

// GetUriOk returns a tuple with the Uri field value
// and a boolean to check if the value has been set.
func (o *VideoMetadataConnectionsResourceCreatorTeamUser) GetUriOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Uri, true
}

// SetUri sets field value
func (o *VideoMetadataConnectionsResourceCreatorTeamUser) SetUri(v string) {
	o.Uri = v
}

func (o VideoMetadataConnectionsResourceCreatorTeamUser) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o VideoMetadataConnectionsResourceCreatorTeamUser) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["uri"] = o.Uri
	return toSerialize, nil
}

type NullableVideoMetadataConnectionsResourceCreatorTeamUser struct {
	value *VideoMetadataConnectionsResourceCreatorTeamUser
	isSet bool
}

func (v NullableVideoMetadataConnectionsResourceCreatorTeamUser) Get() *VideoMetadataConnectionsResourceCreatorTeamUser {
	return v.value
}

func (v *NullableVideoMetadataConnectionsResourceCreatorTeamUser) Set(val *VideoMetadataConnectionsResourceCreatorTeamUser) {
	v.value = val
	v.isSet = true
}

func (v NullableVideoMetadataConnectionsResourceCreatorTeamUser) IsSet() bool {
	return v.isSet
}

func (v *NullableVideoMetadataConnectionsResourceCreatorTeamUser) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableVideoMetadataConnectionsResourceCreatorTeamUser(val *VideoMetadataConnectionsResourceCreatorTeamUser) *NullableVideoMetadataConnectionsResourceCreatorTeamUser {
	return &NullableVideoMetadataConnectionsResourceCreatorTeamUser{value: val, isSet: true}
}

func (v NullableVideoMetadataConnectionsResourceCreatorTeamUser) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableVideoMetadataConnectionsResourceCreatorTeamUser) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
