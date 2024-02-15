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

// checks if the CategoryMetadataConnections type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &CategoryMetadataConnections{}

// CategoryMetadataConnections A collection of information that is connected to this resource.
type CategoryMetadataConnections struct {
	Channels CategoryMetadataConnectionsChannels `json:"channels"`
	Groups CategoryMetadataConnectionsGroups `json:"groups"`
	Users CategoryMetadataConnectionsUsers `json:"users"`
	Videos CategoryMetadataConnectionsVideos `json:"videos"`
}

type _CategoryMetadataConnections CategoryMetadataConnections

// NewCategoryMetadataConnections instantiates a new CategoryMetadataConnections object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewCategoryMetadataConnections(channels CategoryMetadataConnectionsChannels, groups CategoryMetadataConnectionsGroups, users CategoryMetadataConnectionsUsers, videos CategoryMetadataConnectionsVideos) *CategoryMetadataConnections {
	this := CategoryMetadataConnections{}
	this.Channels = channels
	this.Groups = groups
	this.Users = users
	this.Videos = videos
	return &this
}

// NewCategoryMetadataConnectionsWithDefaults instantiates a new CategoryMetadataConnections object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewCategoryMetadataConnectionsWithDefaults() *CategoryMetadataConnections {
	this := CategoryMetadataConnections{}
	return &this
}

// GetChannels returns the Channels field value
func (o *CategoryMetadataConnections) GetChannels() CategoryMetadataConnectionsChannels {
	if o == nil {
		var ret CategoryMetadataConnectionsChannels
		return ret
	}

	return o.Channels
}

// GetChannelsOk returns a tuple with the Channels field value
// and a boolean to check if the value has been set.
func (o *CategoryMetadataConnections) GetChannelsOk() (*CategoryMetadataConnectionsChannels, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Channels, true
}

// SetChannels sets field value
func (o *CategoryMetadataConnections) SetChannels(v CategoryMetadataConnectionsChannels) {
	o.Channels = v
}

// GetGroups returns the Groups field value
func (o *CategoryMetadataConnections) GetGroups() CategoryMetadataConnectionsGroups {
	if o == nil {
		var ret CategoryMetadataConnectionsGroups
		return ret
	}

	return o.Groups
}

// GetGroupsOk returns a tuple with the Groups field value
// and a boolean to check if the value has been set.
func (o *CategoryMetadataConnections) GetGroupsOk() (*CategoryMetadataConnectionsGroups, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Groups, true
}

// SetGroups sets field value
func (o *CategoryMetadataConnections) SetGroups(v CategoryMetadataConnectionsGroups) {
	o.Groups = v
}

// GetUsers returns the Users field value
func (o *CategoryMetadataConnections) GetUsers() CategoryMetadataConnectionsUsers {
	if o == nil {
		var ret CategoryMetadataConnectionsUsers
		return ret
	}

	return o.Users
}

// GetUsersOk returns a tuple with the Users field value
// and a boolean to check if the value has been set.
func (o *CategoryMetadataConnections) GetUsersOk() (*CategoryMetadataConnectionsUsers, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Users, true
}

// SetUsers sets field value
func (o *CategoryMetadataConnections) SetUsers(v CategoryMetadataConnectionsUsers) {
	o.Users = v
}

// GetVideos returns the Videos field value
func (o *CategoryMetadataConnections) GetVideos() CategoryMetadataConnectionsVideos {
	if o == nil {
		var ret CategoryMetadataConnectionsVideos
		return ret
	}

	return o.Videos
}

// GetVideosOk returns a tuple with the Videos field value
// and a boolean to check if the value has been set.
func (o *CategoryMetadataConnections) GetVideosOk() (*CategoryMetadataConnectionsVideos, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Videos, true
}

// SetVideos sets field value
func (o *CategoryMetadataConnections) SetVideos(v CategoryMetadataConnectionsVideos) {
	o.Videos = v
}

func (o CategoryMetadataConnections) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o CategoryMetadataConnections) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["channels"] = o.Channels
	toSerialize["groups"] = o.Groups
	toSerialize["users"] = o.Users
	toSerialize["videos"] = o.Videos
	return toSerialize, nil
}

func (o *CategoryMetadataConnections) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"channels",
		"groups",
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

	varCategoryMetadataConnections := _CategoryMetadataConnections{}

	decoder := json.NewDecoder(bytes.NewReader(data))
	decoder.DisallowUnknownFields()
	err = decoder.Decode(&varCategoryMetadataConnections)

	if err != nil {
		return err
	}

	*o = CategoryMetadataConnections(varCategoryMetadataConnections)

	return err
}

type NullableCategoryMetadataConnections struct {
	value *CategoryMetadataConnections
	isSet bool
}

func (v NullableCategoryMetadataConnections) Get() *CategoryMetadataConnections {
	return v.value
}

func (v *NullableCategoryMetadataConnections) Set(val *CategoryMetadataConnections) {
	v.value = val
	v.isSet = true
}

func (v NullableCategoryMetadataConnections) IsSet() bool {
	return v.isSet
}

func (v *NullableCategoryMetadataConnections) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableCategoryMetadataConnections(val *CategoryMetadataConnections) *NullableCategoryMetadataConnections {
	return &NullableCategoryMetadataConnections{value: val, isSet: true}
}

func (v NullableCategoryMetadataConnections) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableCategoryMetadataConnections) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


