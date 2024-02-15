/*
Vimeo API

Build something great. Vimeo's API supports flexible, high-quality video integration with your custom apps.

API version: 3.4
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package vimeoapi

import (
	"encoding/json"
	"bytes"
	"fmt"
)

// checks if the GroupMetadataConnections type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &GroupMetadataConnections{}

// GroupMetadataConnections A collection of information that is connected to this resource.
type GroupMetadataConnections struct {
	Users GroupMetadataConnectionsUsers `json:"users"`
	Videos GroupMetadataConnectionsVideos `json:"videos"`
}

type _GroupMetadataConnections GroupMetadataConnections

// NewGroupMetadataConnections instantiates a new GroupMetadataConnections object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewGroupMetadataConnections(users GroupMetadataConnectionsUsers, videos GroupMetadataConnectionsVideos) *GroupMetadataConnections {
	this := GroupMetadataConnections{}
	this.Users = users
	this.Videos = videos
	return &this
}

// NewGroupMetadataConnectionsWithDefaults instantiates a new GroupMetadataConnections object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewGroupMetadataConnectionsWithDefaults() *GroupMetadataConnections {
	this := GroupMetadataConnections{}
	return &this
}

// GetUsers returns the Users field value
func (o *GroupMetadataConnections) GetUsers() GroupMetadataConnectionsUsers {
	if o == nil {
		var ret GroupMetadataConnectionsUsers
		return ret
	}

	return o.Users
}

// GetUsersOk returns a tuple with the Users field value
// and a boolean to check if the value has been set.
func (o *GroupMetadataConnections) GetUsersOk() (*GroupMetadataConnectionsUsers, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Users, true
}

// SetUsers sets field value
func (o *GroupMetadataConnections) SetUsers(v GroupMetadataConnectionsUsers) {
	o.Users = v
}

// GetVideos returns the Videos field value
func (o *GroupMetadataConnections) GetVideos() GroupMetadataConnectionsVideos {
	if o == nil {
		var ret GroupMetadataConnectionsVideos
		return ret
	}

	return o.Videos
}

// GetVideosOk returns a tuple with the Videos field value
// and a boolean to check if the value has been set.
func (o *GroupMetadataConnections) GetVideosOk() (*GroupMetadataConnectionsVideos, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Videos, true
}

// SetVideos sets field value
func (o *GroupMetadataConnections) SetVideos(v GroupMetadataConnectionsVideos) {
	o.Videos = v
}

func (o GroupMetadataConnections) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o GroupMetadataConnections) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["users"] = o.Users
	toSerialize["videos"] = o.Videos
	return toSerialize, nil
}

func (o *GroupMetadataConnections) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"users",
		"videos",
	}

	allProperties := make(map[string]interface{})

	err = json.Unmarshal(data, &allProperties)

	if err != nil {
		return err;
	}

	for _, requiredProperty := range(requiredProperties) {
		if _, exists := allProperties[requiredProperty]; !exists {
			return fmt.Errorf("no value given for required property %v", requiredProperty)
		}
	}

	varGroupMetadataConnections := _GroupMetadataConnections{}

	decoder := json.NewDecoder(bytes.NewReader(data))
	decoder.DisallowUnknownFields()
	err = decoder.Decode(&varGroupMetadataConnections)

	if err != nil {
		return err
	}

	*o = GroupMetadataConnections(varGroupMetadataConnections)

	return err
}

type NullableGroupMetadataConnections struct {
	value *GroupMetadataConnections
	isSet bool
}

func (v NullableGroupMetadataConnections) Get() *GroupMetadataConnections {
	return v.value
}

func (v *NullableGroupMetadataConnections) Set(val *GroupMetadataConnections) {
	v.value = val
	v.isSet = true
}

func (v NullableGroupMetadataConnections) IsSet() bool {
	return v.isSet
}

func (v *NullableGroupMetadataConnections) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableGroupMetadataConnections(val *GroupMetadataConnections) *NullableGroupMetadataConnections {
	return &NullableGroupMetadataConnections{value: val, isSet: true}
}

func (v NullableGroupMetadataConnections) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableGroupMetadataConnections) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


