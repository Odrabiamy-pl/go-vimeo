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

// checks if the UserMetadataConnectionsRecommendedChannels type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &UserMetadataConnectionsRecommendedChannels{}

// UserMetadataConnectionsRecommendedChannels A collection of recommended channels for the authenticated user to follow. This data requires a bearer token with the `private` scope. This data requires a bearer token with the `private` scope.
type UserMetadataConnectionsRecommendedChannels struct {
	// An array of HTTP methods permitted on this URI. This data requires a bearer token with the `private` scope.
	Options []string `json:"options"`
	// The total number of channels on this connection. This data requires a bearer token with the `private` scope.
	Total float32 `json:"total"`
	// The API URI that resolves to the connection data. This data requires a bearer token with the `private` scope.
	Uri string `json:"uri"`
}

// NewUserMetadataConnectionsRecommendedChannels instantiates a new UserMetadataConnectionsRecommendedChannels object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewUserMetadataConnectionsRecommendedChannels(options []string, total float32, uri string) *UserMetadataConnectionsRecommendedChannels {
	this := UserMetadataConnectionsRecommendedChannels{}
	this.Options = options
	this.Total = total
	this.Uri = uri
	return &this
}

// NewUserMetadataConnectionsRecommendedChannelsWithDefaults instantiates a new UserMetadataConnectionsRecommendedChannels object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewUserMetadataConnectionsRecommendedChannelsWithDefaults() *UserMetadataConnectionsRecommendedChannels {
	this := UserMetadataConnectionsRecommendedChannels{}
	return &this
}

// GetOptions returns the Options field value
func (o *UserMetadataConnectionsRecommendedChannels) GetOptions() []string {
	if o == nil {
		var ret []string
		return ret
	}

	return o.Options
}

// GetOptionsOk returns a tuple with the Options field value
// and a boolean to check if the value has been set.
func (o *UserMetadataConnectionsRecommendedChannels) GetOptionsOk() ([]string, bool) {
	if o == nil {
		return nil, false
	}
	return o.Options, true
}

// SetOptions sets field value
func (o *UserMetadataConnectionsRecommendedChannels) SetOptions(v []string) {
	o.Options = v
}

// GetTotal returns the Total field value
func (o *UserMetadataConnectionsRecommendedChannels) GetTotal() float32 {
	if o == nil {
		var ret float32
		return ret
	}

	return o.Total
}

// GetTotalOk returns a tuple with the Total field value
// and a boolean to check if the value has been set.
func (o *UserMetadataConnectionsRecommendedChannels) GetTotalOk() (*float32, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Total, true
}

// SetTotal sets field value
func (o *UserMetadataConnectionsRecommendedChannels) SetTotal(v float32) {
	o.Total = v
}

// GetUri returns the Uri field value
func (o *UserMetadataConnectionsRecommendedChannels) GetUri() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Uri
}

// GetUriOk returns a tuple with the Uri field value
// and a boolean to check if the value has been set.
func (o *UserMetadataConnectionsRecommendedChannels) GetUriOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Uri, true
}

// SetUri sets field value
func (o *UserMetadataConnectionsRecommendedChannels) SetUri(v string) {
	o.Uri = v
}

func (o UserMetadataConnectionsRecommendedChannels) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o UserMetadataConnectionsRecommendedChannels) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["options"] = o.Options
	toSerialize["total"] = o.Total
	toSerialize["uri"] = o.Uri
	return toSerialize, nil
}

type NullableUserMetadataConnectionsRecommendedChannels struct {
	value *UserMetadataConnectionsRecommendedChannels
	isSet bool
}

func (v NullableUserMetadataConnectionsRecommendedChannels) Get() *UserMetadataConnectionsRecommendedChannels {
	return v.value
}

func (v *NullableUserMetadataConnectionsRecommendedChannels) Set(val *UserMetadataConnectionsRecommendedChannels) {
	v.value = val
	v.isSet = true
}

func (v NullableUserMetadataConnectionsRecommendedChannels) IsSet() bool {
	return v.isSet
}

func (v *NullableUserMetadataConnectionsRecommendedChannels) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableUserMetadataConnectionsRecommendedChannels(val *UserMetadataConnectionsRecommendedChannels) *NullableUserMetadataConnectionsRecommendedChannels {
	return &NullableUserMetadataConnectionsRecommendedChannels{value: val, isSet: true}
}

func (v NullableUserMetadataConnectionsRecommendedChannels) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableUserMetadataConnectionsRecommendedChannels) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
