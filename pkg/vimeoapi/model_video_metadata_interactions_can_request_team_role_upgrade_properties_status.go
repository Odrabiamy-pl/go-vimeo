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

// checks if the VideoMetadataInteractionsCanRequestTeamRoleUpgradePropertiesStatus type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &VideoMetadataInteractionsCanRequestTeamRoleUpgradePropertiesStatus{}

// VideoMetadataInteractionsCanRequestTeamRoleUpgradePropertiesStatus An object containing data on the value of **status** and whether it's required for the interaction.
type VideoMetadataInteractionsCanRequestTeamRoleUpgradePropertiesStatus struct {
	// Whether the status of the role upgrade request must be sent to achieve the desired action.
	Required bool `json:"required"`
	// The status of the role upgrade request to which the team member should have access.
	Value string `json:"value"`
}

type _VideoMetadataInteractionsCanRequestTeamRoleUpgradePropertiesStatus VideoMetadataInteractionsCanRequestTeamRoleUpgradePropertiesStatus

// NewVideoMetadataInteractionsCanRequestTeamRoleUpgradePropertiesStatus instantiates a new VideoMetadataInteractionsCanRequestTeamRoleUpgradePropertiesStatus object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewVideoMetadataInteractionsCanRequestTeamRoleUpgradePropertiesStatus(required bool, value string) *VideoMetadataInteractionsCanRequestTeamRoleUpgradePropertiesStatus {
	this := VideoMetadataInteractionsCanRequestTeamRoleUpgradePropertiesStatus{}
	this.Required = required
	this.Value = value
	return &this
}

// NewVideoMetadataInteractionsCanRequestTeamRoleUpgradePropertiesStatusWithDefaults instantiates a new VideoMetadataInteractionsCanRequestTeamRoleUpgradePropertiesStatus object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewVideoMetadataInteractionsCanRequestTeamRoleUpgradePropertiesStatusWithDefaults() *VideoMetadataInteractionsCanRequestTeamRoleUpgradePropertiesStatus {
	this := VideoMetadataInteractionsCanRequestTeamRoleUpgradePropertiesStatus{}
	return &this
}

// GetRequired returns the Required field value
func (o *VideoMetadataInteractionsCanRequestTeamRoleUpgradePropertiesStatus) GetRequired() bool {
	if o == nil {
		var ret bool
		return ret
	}

	return o.Required
}

// GetRequiredOk returns a tuple with the Required field value
// and a boolean to check if the value has been set.
func (o *VideoMetadataInteractionsCanRequestTeamRoleUpgradePropertiesStatus) GetRequiredOk() (*bool, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Required, true
}

// SetRequired sets field value
func (o *VideoMetadataInteractionsCanRequestTeamRoleUpgradePropertiesStatus) SetRequired(v bool) {
	o.Required = v
}

// GetValue returns the Value field value
func (o *VideoMetadataInteractionsCanRequestTeamRoleUpgradePropertiesStatus) GetValue() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Value
}

// GetValueOk returns a tuple with the Value field value
// and a boolean to check if the value has been set.
func (o *VideoMetadataInteractionsCanRequestTeamRoleUpgradePropertiesStatus) GetValueOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Value, true
}

// SetValue sets field value
func (o *VideoMetadataInteractionsCanRequestTeamRoleUpgradePropertiesStatus) SetValue(v string) {
	o.Value = v
}

func (o VideoMetadataInteractionsCanRequestTeamRoleUpgradePropertiesStatus) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o VideoMetadataInteractionsCanRequestTeamRoleUpgradePropertiesStatus) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["required"] = o.Required
	toSerialize["value"] = o.Value
	return toSerialize, nil
}

func (o *VideoMetadataInteractionsCanRequestTeamRoleUpgradePropertiesStatus) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"required",
		"value",
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

	varVideoMetadataInteractionsCanRequestTeamRoleUpgradePropertiesStatus := _VideoMetadataInteractionsCanRequestTeamRoleUpgradePropertiesStatus{}

	decoder := json.NewDecoder(bytes.NewReader(data))
	decoder.DisallowUnknownFields()
	err = decoder.Decode(&varVideoMetadataInteractionsCanRequestTeamRoleUpgradePropertiesStatus)

	if err != nil {
		return err
	}

	*o = VideoMetadataInteractionsCanRequestTeamRoleUpgradePropertiesStatus(varVideoMetadataInteractionsCanRequestTeamRoleUpgradePropertiesStatus)

	return err
}

type NullableVideoMetadataInteractionsCanRequestTeamRoleUpgradePropertiesStatus struct {
	value *VideoMetadataInteractionsCanRequestTeamRoleUpgradePropertiesStatus
	isSet bool
}

func (v NullableVideoMetadataInteractionsCanRequestTeamRoleUpgradePropertiesStatus) Get() *VideoMetadataInteractionsCanRequestTeamRoleUpgradePropertiesStatus {
	return v.value
}

func (v *NullableVideoMetadataInteractionsCanRequestTeamRoleUpgradePropertiesStatus) Set(val *VideoMetadataInteractionsCanRequestTeamRoleUpgradePropertiesStatus) {
	v.value = val
	v.isSet = true
}

func (v NullableVideoMetadataInteractionsCanRequestTeamRoleUpgradePropertiesStatus) IsSet() bool {
	return v.isSet
}

func (v *NullableVideoMetadataInteractionsCanRequestTeamRoleUpgradePropertiesStatus) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableVideoMetadataInteractionsCanRequestTeamRoleUpgradePropertiesStatus(val *VideoMetadataInteractionsCanRequestTeamRoleUpgradePropertiesStatus) *NullableVideoMetadataInteractionsCanRequestTeamRoleUpgradePropertiesStatus {
	return &NullableVideoMetadataInteractionsCanRequestTeamRoleUpgradePropertiesStatus{value: val, isSet: true}
}

func (v NullableVideoMetadataInteractionsCanRequestTeamRoleUpgradePropertiesStatus) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableVideoMetadataInteractionsCanRequestTeamRoleUpgradePropertiesStatus) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


