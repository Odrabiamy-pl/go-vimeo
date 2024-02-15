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

// checks if the UserMetadataConnectionsRecommendedUsers type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &UserMetadataConnectionsRecommendedUsers{}

// UserMetadataConnectionsRecommendedUsers A collection of recommended users for the authenticated user to follow. This data requires a bearer token with the `private` scope. This data requires a bearer token with the `private` scope.
type UserMetadataConnectionsRecommendedUsers struct {
	// An array of HTTP methods permitted on this URI. This data requires a bearer token with the `private` scope.
	Options []string `json:"options"`
	// The total number of users on this connection. This data requires a bearer token with the `private` scope.
	Total float32 `json:"total"`
	// The API URI that resolves to the connection data. This data requires a bearer token with the `private` scope.
	Uri string `json:"uri"`
}

// NewUserMetadataConnectionsRecommendedUsers instantiates a new UserMetadataConnectionsRecommendedUsers object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewUserMetadataConnectionsRecommendedUsers(options []string, total float32, uri string) *UserMetadataConnectionsRecommendedUsers {
	this := UserMetadataConnectionsRecommendedUsers{}
	this.Options = options
	this.Total = total
	this.Uri = uri
	return &this
}

// NewUserMetadataConnectionsRecommendedUsersWithDefaults instantiates a new UserMetadataConnectionsRecommendedUsers object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewUserMetadataConnectionsRecommendedUsersWithDefaults() *UserMetadataConnectionsRecommendedUsers {
	this := UserMetadataConnectionsRecommendedUsers{}
	return &this
}

// GetOptions returns the Options field value
func (o *UserMetadataConnectionsRecommendedUsers) GetOptions() []string {
	if o == nil {
		var ret []string
		return ret
	}

	return o.Options
}

// GetOptionsOk returns a tuple with the Options field value
// and a boolean to check if the value has been set.
func (o *UserMetadataConnectionsRecommendedUsers) GetOptionsOk() ([]string, bool) {
	if o == nil {
		return nil, false
	}
	return o.Options, true
}

// SetOptions sets field value
func (o *UserMetadataConnectionsRecommendedUsers) SetOptions(v []string) {
	o.Options = v
}

// GetTotal returns the Total field value
func (o *UserMetadataConnectionsRecommendedUsers) GetTotal() float32 {
	if o == nil {
		var ret float32
		return ret
	}

	return o.Total
}

// GetTotalOk returns a tuple with the Total field value
// and a boolean to check if the value has been set.
func (o *UserMetadataConnectionsRecommendedUsers) GetTotalOk() (*float32, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Total, true
}

// SetTotal sets field value
func (o *UserMetadataConnectionsRecommendedUsers) SetTotal(v float32) {
	o.Total = v
}

// GetUri returns the Uri field value
func (o *UserMetadataConnectionsRecommendedUsers) GetUri() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Uri
}

// GetUriOk returns a tuple with the Uri field value
// and a boolean to check if the value has been set.
func (o *UserMetadataConnectionsRecommendedUsers) GetUriOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Uri, true
}

// SetUri sets field value
func (o *UserMetadataConnectionsRecommendedUsers) SetUri(v string) {
	o.Uri = v
}

func (o UserMetadataConnectionsRecommendedUsers) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o UserMetadataConnectionsRecommendedUsers) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["options"] = o.Options
	toSerialize["total"] = o.Total
	toSerialize["uri"] = o.Uri
	return toSerialize, nil
}

type NullableUserMetadataConnectionsRecommendedUsers struct {
	value *UserMetadataConnectionsRecommendedUsers
	isSet bool
}

func (v NullableUserMetadataConnectionsRecommendedUsers) Get() *UserMetadataConnectionsRecommendedUsers {
	return v.value
}

func (v *NullableUserMetadataConnectionsRecommendedUsers) Set(val *UserMetadataConnectionsRecommendedUsers) {
	v.value = val
	v.isSet = true
}

func (v NullableUserMetadataConnectionsRecommendedUsers) IsSet() bool {
	return v.isSet
}

func (v *NullableUserMetadataConnectionsRecommendedUsers) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableUserMetadataConnectionsRecommendedUsers(val *UserMetadataConnectionsRecommendedUsers) *NullableUserMetadataConnectionsRecommendedUsers {
	return &NullableUserMetadataConnectionsRecommendedUsers{value: val, isSet: true}
}

func (v NullableUserMetadataConnectionsRecommendedUsers) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableUserMetadataConnectionsRecommendedUsers) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


