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

// checks if the TeamMembershipMetadataConnections type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &TeamMembershipMetadataConnections{}

// TeamMembershipMetadataConnections A list of resource URIs related to the user.
type TeamMembershipMetadataConnections struct {
	Owner              TeamMembershipMetadataConnectionsOwner              `json:"owner"`
	PersonalTeamFolder TeamMembershipMetadataConnectionsPersonalTeamFolder `json:"personal_team_folder"`
}

// NewTeamMembershipMetadataConnections instantiates a new TeamMembershipMetadataConnections object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewTeamMembershipMetadataConnections(owner TeamMembershipMetadataConnectionsOwner, personalTeamFolder TeamMembershipMetadataConnectionsPersonalTeamFolder) *TeamMembershipMetadataConnections {
	this := TeamMembershipMetadataConnections{}
	this.Owner = owner
	this.PersonalTeamFolder = personalTeamFolder
	return &this
}

// NewTeamMembershipMetadataConnectionsWithDefaults instantiates a new TeamMembershipMetadataConnections object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewTeamMembershipMetadataConnectionsWithDefaults() *TeamMembershipMetadataConnections {
	this := TeamMembershipMetadataConnections{}
	return &this
}

// GetOwner returns the Owner field value
func (o *TeamMembershipMetadataConnections) GetOwner() TeamMembershipMetadataConnectionsOwner {
	if o == nil {
		var ret TeamMembershipMetadataConnectionsOwner
		return ret
	}

	return o.Owner
}

// GetOwnerOk returns a tuple with the Owner field value
// and a boolean to check if the value has been set.
func (o *TeamMembershipMetadataConnections) GetOwnerOk() (*TeamMembershipMetadataConnectionsOwner, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Owner, true
}

// SetOwner sets field value
func (o *TeamMembershipMetadataConnections) SetOwner(v TeamMembershipMetadataConnectionsOwner) {
	o.Owner = v
}

// GetPersonalTeamFolder returns the PersonalTeamFolder field value
func (o *TeamMembershipMetadataConnections) GetPersonalTeamFolder() TeamMembershipMetadataConnectionsPersonalTeamFolder {
	if o == nil {
		var ret TeamMembershipMetadataConnectionsPersonalTeamFolder
		return ret
	}

	return o.PersonalTeamFolder
}

// GetPersonalTeamFolderOk returns a tuple with the PersonalTeamFolder field value
// and a boolean to check if the value has been set.
func (o *TeamMembershipMetadataConnections) GetPersonalTeamFolderOk() (*TeamMembershipMetadataConnectionsPersonalTeamFolder, bool) {
	if o == nil {
		return nil, false
	}
	return &o.PersonalTeamFolder, true
}

// SetPersonalTeamFolder sets field value
func (o *TeamMembershipMetadataConnections) SetPersonalTeamFolder(v TeamMembershipMetadataConnectionsPersonalTeamFolder) {
	o.PersonalTeamFolder = v
}

func (o TeamMembershipMetadataConnections) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o TeamMembershipMetadataConnections) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["owner"] = o.Owner
	toSerialize["personal_team_folder"] = o.PersonalTeamFolder
	return toSerialize, nil
}

type NullableTeamMembershipMetadataConnections struct {
	value *TeamMembershipMetadataConnections
	isSet bool
}

func (v NullableTeamMembershipMetadataConnections) Get() *TeamMembershipMetadataConnections {
	return v.value
}

func (v *NullableTeamMembershipMetadataConnections) Set(val *TeamMembershipMetadataConnections) {
	v.value = val
	v.isSet = true
}

func (v NullableTeamMembershipMetadataConnections) IsSet() bool {
	return v.isSet
}

func (v *NullableTeamMembershipMetadataConnections) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableTeamMembershipMetadataConnections(val *TeamMembershipMetadataConnections) *NullableTeamMembershipMetadataConnections {
	return &NullableTeamMembershipMetadataConnections{value: val, isSet: true}
}

func (v NullableTeamMembershipMetadataConnections) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableTeamMembershipMetadataConnections) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
