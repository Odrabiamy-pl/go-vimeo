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

// checks if the UserMetadataInteractions type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &UserMetadataInteractions{}

// UserMetadataInteractions struct for UserMetadataInteractions
type UserMetadataInteractions struct {
	AddPrivacyUser *UserMetadataInteractionsAddPrivacyUser `json:"add_privacy_user,omitempty"`
	Block          UserMetadataInteractionsBlock           `json:"block"`
	ConnectedApps  UserMetadataInteractionsConnectedApps   `json:"connected_apps"`
	Follow         UserMetadataInteractionsFollow          `json:"follow"`
	Report         UserMetadataInteractionsReport          `json:"report"`
}

// NewUserMetadataInteractions instantiates a new UserMetadataInteractions object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewUserMetadataInteractions(block UserMetadataInteractionsBlock, connectedApps UserMetadataInteractionsConnectedApps, follow UserMetadataInteractionsFollow, report UserMetadataInteractionsReport) *UserMetadataInteractions {
	this := UserMetadataInteractions{}
	this.Block = block
	this.ConnectedApps = connectedApps
	this.Follow = follow
	this.Report = report
	return &this
}

// NewUserMetadataInteractionsWithDefaults instantiates a new UserMetadataInteractions object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewUserMetadataInteractionsWithDefaults() *UserMetadataInteractions {
	this := UserMetadataInteractions{}
	return &this
}

// GetAddPrivacyUser returns the AddPrivacyUser field value if set, zero value otherwise.
func (o *UserMetadataInteractions) GetAddPrivacyUser() UserMetadataInteractionsAddPrivacyUser {
	if o == nil || IsNil(o.AddPrivacyUser) {
		var ret UserMetadataInteractionsAddPrivacyUser
		return ret
	}
	return *o.AddPrivacyUser
}

// GetAddPrivacyUserOk returns a tuple with the AddPrivacyUser field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UserMetadataInteractions) GetAddPrivacyUserOk() (*UserMetadataInteractionsAddPrivacyUser, bool) {
	if o == nil || IsNil(o.AddPrivacyUser) {
		return nil, false
	}
	return o.AddPrivacyUser, true
}

// HasAddPrivacyUser returns a boolean if a field has been set.
func (o *UserMetadataInteractions) HasAddPrivacyUser() bool {
	if o != nil && !IsNil(o.AddPrivacyUser) {
		return true
	}

	return false
}

// SetAddPrivacyUser gets a reference to the given UserMetadataInteractionsAddPrivacyUser and assigns it to the AddPrivacyUser field.
func (o *UserMetadataInteractions) SetAddPrivacyUser(v UserMetadataInteractionsAddPrivacyUser) {
	o.AddPrivacyUser = &v
}

// GetBlock returns the Block field value
func (o *UserMetadataInteractions) GetBlock() UserMetadataInteractionsBlock {
	if o == nil {
		var ret UserMetadataInteractionsBlock
		return ret
	}

	return o.Block
}

// GetBlockOk returns a tuple with the Block field value
// and a boolean to check if the value has been set.
func (o *UserMetadataInteractions) GetBlockOk() (*UserMetadataInteractionsBlock, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Block, true
}

// SetBlock sets field value
func (o *UserMetadataInteractions) SetBlock(v UserMetadataInteractionsBlock) {
	o.Block = v
}

// GetConnectedApps returns the ConnectedApps field value
func (o *UserMetadataInteractions) GetConnectedApps() UserMetadataInteractionsConnectedApps {
	if o == nil {
		var ret UserMetadataInteractionsConnectedApps
		return ret
	}

	return o.ConnectedApps
}

// GetConnectedAppsOk returns a tuple with the ConnectedApps field value
// and a boolean to check if the value has been set.
func (o *UserMetadataInteractions) GetConnectedAppsOk() (*UserMetadataInteractionsConnectedApps, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ConnectedApps, true
}

// SetConnectedApps sets field value
func (o *UserMetadataInteractions) SetConnectedApps(v UserMetadataInteractionsConnectedApps) {
	o.ConnectedApps = v
}

// GetFollow returns the Follow field value
func (o *UserMetadataInteractions) GetFollow() UserMetadataInteractionsFollow {
	if o == nil {
		var ret UserMetadataInteractionsFollow
		return ret
	}

	return o.Follow
}

// GetFollowOk returns a tuple with the Follow field value
// and a boolean to check if the value has been set.
func (o *UserMetadataInteractions) GetFollowOk() (*UserMetadataInteractionsFollow, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Follow, true
}

// SetFollow sets field value
func (o *UserMetadataInteractions) SetFollow(v UserMetadataInteractionsFollow) {
	o.Follow = v
}

// GetReport returns the Report field value
func (o *UserMetadataInteractions) GetReport() UserMetadataInteractionsReport {
	if o == nil {
		var ret UserMetadataInteractionsReport
		return ret
	}

	return o.Report
}

// GetReportOk returns a tuple with the Report field value
// and a boolean to check if the value has been set.
func (o *UserMetadataInteractions) GetReportOk() (*UserMetadataInteractionsReport, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Report, true
}

// SetReport sets field value
func (o *UserMetadataInteractions) SetReport(v UserMetadataInteractionsReport) {
	o.Report = v
}

func (o UserMetadataInteractions) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o UserMetadataInteractions) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.AddPrivacyUser) {
		toSerialize["add_privacy_user"] = o.AddPrivacyUser
	}
	toSerialize["block"] = o.Block
	toSerialize["connected_apps"] = o.ConnectedApps
	toSerialize["follow"] = o.Follow
	toSerialize["report"] = o.Report
	return toSerialize, nil
}

type NullableUserMetadataInteractions struct {
	value *UserMetadataInteractions
	isSet bool
}

func (v NullableUserMetadataInteractions) Get() *UserMetadataInteractions {
	return v.value
}

func (v *NullableUserMetadataInteractions) Set(val *UserMetadataInteractions) {
	v.value = val
	v.isSet = true
}

func (v NullableUserMetadataInteractions) IsSet() bool {
	return v.isSet
}

func (v *NullableUserMetadataInteractions) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableUserMetadataInteractions(val *UserMetadataInteractions) *NullableUserMetadataInteractions {
	return &NullableUserMetadataInteractions{value: val, isSet: true}
}

func (v NullableUserMetadataInteractions) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableUserMetadataInteractions) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
