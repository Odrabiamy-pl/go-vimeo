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

// checks if the VideoMetadataInteractionsCanRequestTeamRoleUpgradeProperties type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &VideoMetadataInteractionsCanRequestTeamRoleUpgradeProperties{}

// VideoMetadataInteractionsCanRequestTeamRoleUpgradeProperties An object of suggested fields to be used for this interaction.
type VideoMetadataInteractionsCanRequestTeamRoleUpgradeProperties struct {
	FolderUri VideoMetadataInteractionsCanRequestTeamRoleUpgradePropertiesFolderUri `json:"folder_uri"`
	Status VideoMetadataInteractionsCanRequestTeamRoleUpgradePropertiesStatus `json:"status"`
	UpgradeToRole VideoMetadataInteractionsCanRequestTeamRoleUpgradePropertiesUpgradeToRole `json:"upgrade_to_role"`
}

type _VideoMetadataInteractionsCanRequestTeamRoleUpgradeProperties VideoMetadataInteractionsCanRequestTeamRoleUpgradeProperties

// NewVideoMetadataInteractionsCanRequestTeamRoleUpgradeProperties instantiates a new VideoMetadataInteractionsCanRequestTeamRoleUpgradeProperties object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewVideoMetadataInteractionsCanRequestTeamRoleUpgradeProperties(folderUri VideoMetadataInteractionsCanRequestTeamRoleUpgradePropertiesFolderUri, status VideoMetadataInteractionsCanRequestTeamRoleUpgradePropertiesStatus, upgradeToRole VideoMetadataInteractionsCanRequestTeamRoleUpgradePropertiesUpgradeToRole) *VideoMetadataInteractionsCanRequestTeamRoleUpgradeProperties {
	this := VideoMetadataInteractionsCanRequestTeamRoleUpgradeProperties{}
	this.FolderUri = folderUri
	this.Status = status
	this.UpgradeToRole = upgradeToRole
	return &this
}

// NewVideoMetadataInteractionsCanRequestTeamRoleUpgradePropertiesWithDefaults instantiates a new VideoMetadataInteractionsCanRequestTeamRoleUpgradeProperties object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewVideoMetadataInteractionsCanRequestTeamRoleUpgradePropertiesWithDefaults() *VideoMetadataInteractionsCanRequestTeamRoleUpgradeProperties {
	this := VideoMetadataInteractionsCanRequestTeamRoleUpgradeProperties{}
	return &this
}

// GetFolderUri returns the FolderUri field value
func (o *VideoMetadataInteractionsCanRequestTeamRoleUpgradeProperties) GetFolderUri() VideoMetadataInteractionsCanRequestTeamRoleUpgradePropertiesFolderUri {
	if o == nil {
		var ret VideoMetadataInteractionsCanRequestTeamRoleUpgradePropertiesFolderUri
		return ret
	}

	return o.FolderUri
}

// GetFolderUriOk returns a tuple with the FolderUri field value
// and a boolean to check if the value has been set.
func (o *VideoMetadataInteractionsCanRequestTeamRoleUpgradeProperties) GetFolderUriOk() (*VideoMetadataInteractionsCanRequestTeamRoleUpgradePropertiesFolderUri, bool) {
	if o == nil {
		return nil, false
	}
	return &o.FolderUri, true
}

// SetFolderUri sets field value
func (o *VideoMetadataInteractionsCanRequestTeamRoleUpgradeProperties) SetFolderUri(v VideoMetadataInteractionsCanRequestTeamRoleUpgradePropertiesFolderUri) {
	o.FolderUri = v
}

// GetStatus returns the Status field value
func (o *VideoMetadataInteractionsCanRequestTeamRoleUpgradeProperties) GetStatus() VideoMetadataInteractionsCanRequestTeamRoleUpgradePropertiesStatus {
	if o == nil {
		var ret VideoMetadataInteractionsCanRequestTeamRoleUpgradePropertiesStatus
		return ret
	}

	return o.Status
}

// GetStatusOk returns a tuple with the Status field value
// and a boolean to check if the value has been set.
func (o *VideoMetadataInteractionsCanRequestTeamRoleUpgradeProperties) GetStatusOk() (*VideoMetadataInteractionsCanRequestTeamRoleUpgradePropertiesStatus, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Status, true
}

// SetStatus sets field value
func (o *VideoMetadataInteractionsCanRequestTeamRoleUpgradeProperties) SetStatus(v VideoMetadataInteractionsCanRequestTeamRoleUpgradePropertiesStatus) {
	o.Status = v
}

// GetUpgradeToRole returns the UpgradeToRole field value
func (o *VideoMetadataInteractionsCanRequestTeamRoleUpgradeProperties) GetUpgradeToRole() VideoMetadataInteractionsCanRequestTeamRoleUpgradePropertiesUpgradeToRole {
	if o == nil {
		var ret VideoMetadataInteractionsCanRequestTeamRoleUpgradePropertiesUpgradeToRole
		return ret
	}

	return o.UpgradeToRole
}

// GetUpgradeToRoleOk returns a tuple with the UpgradeToRole field value
// and a boolean to check if the value has been set.
func (o *VideoMetadataInteractionsCanRequestTeamRoleUpgradeProperties) GetUpgradeToRoleOk() (*VideoMetadataInteractionsCanRequestTeamRoleUpgradePropertiesUpgradeToRole, bool) {
	if o == nil {
		return nil, false
	}
	return &o.UpgradeToRole, true
}

// SetUpgradeToRole sets field value
func (o *VideoMetadataInteractionsCanRequestTeamRoleUpgradeProperties) SetUpgradeToRole(v VideoMetadataInteractionsCanRequestTeamRoleUpgradePropertiesUpgradeToRole) {
	o.UpgradeToRole = v
}

func (o VideoMetadataInteractionsCanRequestTeamRoleUpgradeProperties) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o VideoMetadataInteractionsCanRequestTeamRoleUpgradeProperties) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["folder_uri"] = o.FolderUri
	toSerialize["status"] = o.Status
	toSerialize["upgrade_to_role"] = o.UpgradeToRole
	return toSerialize, nil
}

func (o *VideoMetadataInteractionsCanRequestTeamRoleUpgradeProperties) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"folder_uri",
		"status",
		"upgrade_to_role",
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

	varVideoMetadataInteractionsCanRequestTeamRoleUpgradeProperties := _VideoMetadataInteractionsCanRequestTeamRoleUpgradeProperties{}

	decoder := json.NewDecoder(bytes.NewReader(data))
	decoder.DisallowUnknownFields()
	err = decoder.Decode(&varVideoMetadataInteractionsCanRequestTeamRoleUpgradeProperties)

	if err != nil {
		return err
	}

	*o = VideoMetadataInteractionsCanRequestTeamRoleUpgradeProperties(varVideoMetadataInteractionsCanRequestTeamRoleUpgradeProperties)

	return err
}

type NullableVideoMetadataInteractionsCanRequestTeamRoleUpgradeProperties struct {
	value *VideoMetadataInteractionsCanRequestTeamRoleUpgradeProperties
	isSet bool
}

func (v NullableVideoMetadataInteractionsCanRequestTeamRoleUpgradeProperties) Get() *VideoMetadataInteractionsCanRequestTeamRoleUpgradeProperties {
	return v.value
}

func (v *NullableVideoMetadataInteractionsCanRequestTeamRoleUpgradeProperties) Set(val *VideoMetadataInteractionsCanRequestTeamRoleUpgradeProperties) {
	v.value = val
	v.isSet = true
}

func (v NullableVideoMetadataInteractionsCanRequestTeamRoleUpgradeProperties) IsSet() bool {
	return v.isSet
}

func (v *NullableVideoMetadataInteractionsCanRequestTeamRoleUpgradeProperties) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableVideoMetadataInteractionsCanRequestTeamRoleUpgradeProperties(val *VideoMetadataInteractionsCanRequestTeamRoleUpgradeProperties) *NullableVideoMetadataInteractionsCanRequestTeamRoleUpgradeProperties {
	return &NullableVideoMetadataInteractionsCanRequestTeamRoleUpgradeProperties{value: val, isSet: true}
}

func (v NullableVideoMetadataInteractionsCanRequestTeamRoleUpgradeProperties) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableVideoMetadataInteractionsCanRequestTeamRoleUpgradeProperties) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


