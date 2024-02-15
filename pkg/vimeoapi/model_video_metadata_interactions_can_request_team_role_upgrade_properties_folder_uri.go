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

// checks if the VideoMetadataInteractionsCanRequestTeamRoleUpgradePropertiesFolderUri type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &VideoMetadataInteractionsCanRequestTeamRoleUpgradePropertiesFolderUri{}

// VideoMetadataInteractionsCanRequestTeamRoleUpgradePropertiesFolderUri An object containing data on the value of **folder_uri** and whether it's required for the interaction.
type VideoMetadataInteractionsCanRequestTeamRoleUpgradePropertiesFolderUri struct {
	// Whether the URI of the folder must be sent to achieve the desired action.
	Required bool `json:"required"`
	// The URI of the folder to which the team member should have access.
	Value string `json:"value"`
}

// NewVideoMetadataInteractionsCanRequestTeamRoleUpgradePropertiesFolderUri instantiates a new VideoMetadataInteractionsCanRequestTeamRoleUpgradePropertiesFolderUri object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewVideoMetadataInteractionsCanRequestTeamRoleUpgradePropertiesFolderUri(required bool, value string) *VideoMetadataInteractionsCanRequestTeamRoleUpgradePropertiesFolderUri {
	this := VideoMetadataInteractionsCanRequestTeamRoleUpgradePropertiesFolderUri{}
	this.Required = required
	this.Value = value
	return &this
}

// NewVideoMetadataInteractionsCanRequestTeamRoleUpgradePropertiesFolderUriWithDefaults instantiates a new VideoMetadataInteractionsCanRequestTeamRoleUpgradePropertiesFolderUri object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewVideoMetadataInteractionsCanRequestTeamRoleUpgradePropertiesFolderUriWithDefaults() *VideoMetadataInteractionsCanRequestTeamRoleUpgradePropertiesFolderUri {
	this := VideoMetadataInteractionsCanRequestTeamRoleUpgradePropertiesFolderUri{}
	return &this
}

// GetRequired returns the Required field value
func (o *VideoMetadataInteractionsCanRequestTeamRoleUpgradePropertiesFolderUri) GetRequired() bool {
	if o == nil {
		var ret bool
		return ret
	}

	return o.Required
}

// GetRequiredOk returns a tuple with the Required field value
// and a boolean to check if the value has been set.
func (o *VideoMetadataInteractionsCanRequestTeamRoleUpgradePropertiesFolderUri) GetRequiredOk() (*bool, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Required, true
}

// SetRequired sets field value
func (o *VideoMetadataInteractionsCanRequestTeamRoleUpgradePropertiesFolderUri) SetRequired(v bool) {
	o.Required = v
}

// GetValue returns the Value field value
func (o *VideoMetadataInteractionsCanRequestTeamRoleUpgradePropertiesFolderUri) GetValue() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Value
}

// GetValueOk returns a tuple with the Value field value
// and a boolean to check if the value has been set.
func (o *VideoMetadataInteractionsCanRequestTeamRoleUpgradePropertiesFolderUri) GetValueOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Value, true
}

// SetValue sets field value
func (o *VideoMetadataInteractionsCanRequestTeamRoleUpgradePropertiesFolderUri) SetValue(v string) {
	o.Value = v
}

func (o VideoMetadataInteractionsCanRequestTeamRoleUpgradePropertiesFolderUri) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o VideoMetadataInteractionsCanRequestTeamRoleUpgradePropertiesFolderUri) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["required"] = o.Required
	toSerialize["value"] = o.Value
	return toSerialize, nil
}

type NullableVideoMetadataInteractionsCanRequestTeamRoleUpgradePropertiesFolderUri struct {
	value *VideoMetadataInteractionsCanRequestTeamRoleUpgradePropertiesFolderUri
	isSet bool
}

func (v NullableVideoMetadataInteractionsCanRequestTeamRoleUpgradePropertiesFolderUri) Get() *VideoMetadataInteractionsCanRequestTeamRoleUpgradePropertiesFolderUri {
	return v.value
}

func (v *NullableVideoMetadataInteractionsCanRequestTeamRoleUpgradePropertiesFolderUri) Set(val *VideoMetadataInteractionsCanRequestTeamRoleUpgradePropertiesFolderUri) {
	v.value = val
	v.isSet = true
}

func (v NullableVideoMetadataInteractionsCanRequestTeamRoleUpgradePropertiesFolderUri) IsSet() bool {
	return v.isSet
}

func (v *NullableVideoMetadataInteractionsCanRequestTeamRoleUpgradePropertiesFolderUri) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableVideoMetadataInteractionsCanRequestTeamRoleUpgradePropertiesFolderUri(val *VideoMetadataInteractionsCanRequestTeamRoleUpgradePropertiesFolderUri) *NullableVideoMetadataInteractionsCanRequestTeamRoleUpgradePropertiesFolderUri {
	return &NullableVideoMetadataInteractionsCanRequestTeamRoleUpgradePropertiesFolderUri{value: val, isSet: true}
}

func (v NullableVideoMetadataInteractionsCanRequestTeamRoleUpgradePropertiesFolderUri) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableVideoMetadataInteractionsCanRequestTeamRoleUpgradePropertiesFolderUri) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


