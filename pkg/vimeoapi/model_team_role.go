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

// checks if the TeamRole type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &TeamRole{}

// TeamRole struct for TeamRole
type TeamRole struct {
	// The untranslated role of the user who made the request.  Option descriptions:  * `Admin` - The team member has admin permissions. They can upload and edit videos for the entire team and perform team administration tasks.  * `Contributor` - The team member has contributor permissions. They can upload and edit videos for the entire team but can’t perform team administration tasks.  * `Owner` - The team member has owner permissions.  * `Uploader` - The team member has uploader permissions. They can upload videos for the entire team but can’t edit videos.  * `Viewer` - The team member has viewer permissions. They can access team videos and specific team folders but can’t upload or edit videos.
	Role NullableString `json:"role"`
	// The unique identifier to access the team role.
	Uri string `json:"uri"`
}

// NewTeamRole instantiates a new TeamRole object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewTeamRole(role NullableString, uri string) *TeamRole {
	this := TeamRole{}
	this.Role = role
	this.Uri = uri
	return &this
}

// NewTeamRoleWithDefaults instantiates a new TeamRole object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewTeamRoleWithDefaults() *TeamRole {
	this := TeamRole{}
	return &this
}

// GetRole returns the Role field value
// If the value is explicit nil, the zero value for string will be returned
func (o *TeamRole) GetRole() string {
	if o == nil || o.Role.Get() == nil {
		var ret string
		return ret
	}

	return *o.Role.Get()
}

// GetRoleOk returns a tuple with the Role field value
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *TeamRole) GetRoleOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.Role.Get(), o.Role.IsSet()
}

// SetRole sets field value
func (o *TeamRole) SetRole(v string) {
	o.Role.Set(&v)
}

// GetUri returns the Uri field value
func (o *TeamRole) GetUri() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Uri
}

// GetUriOk returns a tuple with the Uri field value
// and a boolean to check if the value has been set.
func (o *TeamRole) GetUriOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Uri, true
}

// SetUri sets field value
func (o *TeamRole) SetUri(v string) {
	o.Uri = v
}

func (o TeamRole) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o TeamRole) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["role"] = o.Role.Get()
	toSerialize["uri"] = o.Uri
	return toSerialize, nil
}

type NullableTeamRole struct {
	value *TeamRole
	isSet bool
}

func (v NullableTeamRole) Get() *TeamRole {
	return v.value
}

func (v *NullableTeamRole) Set(val *TeamRole) {
	v.value = val
	v.isSet = true
}

func (v NullableTeamRole) IsSet() bool {
	return v.isSet
}

func (v *NullableTeamRole) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableTeamRole(val *TeamRole) *NullableTeamRole {
	return &NullableTeamRole{value: val, isSet: true}
}

func (v NullableTeamRole) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableTeamRole) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
