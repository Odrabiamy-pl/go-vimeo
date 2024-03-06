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

// checks if the ChannelMetadataConnections type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &ChannelMetadataConnections{}

// ChannelMetadataConnections A collection of information that is connected to this resource.
type ChannelMetadataConnections struct {
	PrivacyUsers ChannelMetadataConnectionsPrivacyUsers `json:"privacy_users"`
	Users        ChannelMetadataConnectionsUsers        `json:"users"`
	Videos       ChannelMetadataConnectionsVideos       `json:"videos"`
}

// NewChannelMetadataConnections instantiates a new ChannelMetadataConnections object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewChannelMetadataConnections(privacyUsers ChannelMetadataConnectionsPrivacyUsers, users ChannelMetadataConnectionsUsers, videos ChannelMetadataConnectionsVideos) *ChannelMetadataConnections {
	this := ChannelMetadataConnections{}
	this.PrivacyUsers = privacyUsers
	this.Users = users
	this.Videos = videos
	return &this
}

// NewChannelMetadataConnectionsWithDefaults instantiates a new ChannelMetadataConnections object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewChannelMetadataConnectionsWithDefaults() *ChannelMetadataConnections {
	this := ChannelMetadataConnections{}
	return &this
}

// GetPrivacyUsers returns the PrivacyUsers field value
func (o *ChannelMetadataConnections) GetPrivacyUsers() ChannelMetadataConnectionsPrivacyUsers {
	if o == nil {
		var ret ChannelMetadataConnectionsPrivacyUsers
		return ret
	}

	return o.PrivacyUsers
}

// GetPrivacyUsersOk returns a tuple with the PrivacyUsers field value
// and a boolean to check if the value has been set.
func (o *ChannelMetadataConnections) GetPrivacyUsersOk() (*ChannelMetadataConnectionsPrivacyUsers, bool) {
	if o == nil {
		return nil, false
	}
	return &o.PrivacyUsers, true
}

// SetPrivacyUsers sets field value
func (o *ChannelMetadataConnections) SetPrivacyUsers(v ChannelMetadataConnectionsPrivacyUsers) {
	o.PrivacyUsers = v
}

// GetUsers returns the Users field value
func (o *ChannelMetadataConnections) GetUsers() ChannelMetadataConnectionsUsers {
	if o == nil {
		var ret ChannelMetadataConnectionsUsers
		return ret
	}

	return o.Users
}

// GetUsersOk returns a tuple with the Users field value
// and a boolean to check if the value has been set.
func (o *ChannelMetadataConnections) GetUsersOk() (*ChannelMetadataConnectionsUsers, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Users, true
}

// SetUsers sets field value
func (o *ChannelMetadataConnections) SetUsers(v ChannelMetadataConnectionsUsers) {
	o.Users = v
}

// GetVideos returns the Videos field value
func (o *ChannelMetadataConnections) GetVideos() ChannelMetadataConnectionsVideos {
	if o == nil {
		var ret ChannelMetadataConnectionsVideos
		return ret
	}

	return o.Videos
}

// GetVideosOk returns a tuple with the Videos field value
// and a boolean to check if the value has been set.
func (o *ChannelMetadataConnections) GetVideosOk() (*ChannelMetadataConnectionsVideos, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Videos, true
}

// SetVideos sets field value
func (o *ChannelMetadataConnections) SetVideos(v ChannelMetadataConnectionsVideos) {
	o.Videos = v
}

func (o ChannelMetadataConnections) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o ChannelMetadataConnections) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["privacy_users"] = o.PrivacyUsers
	toSerialize["users"] = o.Users
	toSerialize["videos"] = o.Videos
	return toSerialize, nil
}

type NullableChannelMetadataConnections struct {
	value *ChannelMetadataConnections
	isSet bool
}

func (v NullableChannelMetadataConnections) Get() *ChannelMetadataConnections {
	return v.value
}

func (v *NullableChannelMetadataConnections) Set(val *ChannelMetadataConnections) {
	v.value = val
	v.isSet = true
}

func (v NullableChannelMetadataConnections) IsSet() bool {
	return v.isSet
}

func (v *NullableChannelMetadataConnections) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableChannelMetadataConnections(val *ChannelMetadataConnections) *NullableChannelMetadataConnections {
	return &NullableChannelMetadataConnections{value: val, isSet: true}
}

func (v NullableChannelMetadataConnections) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableChannelMetadataConnections) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
