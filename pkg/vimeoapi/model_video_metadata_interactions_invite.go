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

// checks if the VideoMetadataInteractionsInvite type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &VideoMetadataInteractionsInvite{}

// VideoMetadataInteractionsInvite Information about where and how to get a list of team members or groups who were explicitly invited to a video, and where and how to invite a team member to a video.
type VideoMetadataInteractionsInvite struct {
	// An array of HTTP methods permitted on this URI.
	Options []string `json:"options"`
	// The API URI that resolves to the connection data.
	Uri string `json:"uri"`
}

// NewVideoMetadataInteractionsInvite instantiates a new VideoMetadataInteractionsInvite object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewVideoMetadataInteractionsInvite(options []string, uri string) *VideoMetadataInteractionsInvite {
	this := VideoMetadataInteractionsInvite{}
	this.Options = options
	this.Uri = uri
	return &this
}

// NewVideoMetadataInteractionsInviteWithDefaults instantiates a new VideoMetadataInteractionsInvite object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewVideoMetadataInteractionsInviteWithDefaults() *VideoMetadataInteractionsInvite {
	this := VideoMetadataInteractionsInvite{}
	return &this
}

// GetOptions returns the Options field value
func (o *VideoMetadataInteractionsInvite) GetOptions() []string {
	if o == nil {
		var ret []string
		return ret
	}

	return o.Options
}

// GetOptionsOk returns a tuple with the Options field value
// and a boolean to check if the value has been set.
func (o *VideoMetadataInteractionsInvite) GetOptionsOk() ([]string, bool) {
	if o == nil {
		return nil, false
	}
	return o.Options, true
}

// SetOptions sets field value
func (o *VideoMetadataInteractionsInvite) SetOptions(v []string) {
	o.Options = v
}

// GetUri returns the Uri field value
func (o *VideoMetadataInteractionsInvite) GetUri() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Uri
}

// GetUriOk returns a tuple with the Uri field value
// and a boolean to check if the value has been set.
func (o *VideoMetadataInteractionsInvite) GetUriOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Uri, true
}

// SetUri sets field value
func (o *VideoMetadataInteractionsInvite) SetUri(v string) {
	o.Uri = v
}

func (o VideoMetadataInteractionsInvite) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o VideoMetadataInteractionsInvite) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["options"] = o.Options
	toSerialize["uri"] = o.Uri
	return toSerialize, nil
}

type NullableVideoMetadataInteractionsInvite struct {
	value *VideoMetadataInteractionsInvite
	isSet bool
}

func (v NullableVideoMetadataInteractionsInvite) Get() *VideoMetadataInteractionsInvite {
	return v.value
}

func (v *NullableVideoMetadataInteractionsInvite) Set(val *VideoMetadataInteractionsInvite) {
	v.value = val
	v.isSet = true
}

func (v NullableVideoMetadataInteractionsInvite) IsSet() bool {
	return v.isSet
}

func (v *NullableVideoMetadataInteractionsInvite) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableVideoMetadataInteractionsInvite(val *VideoMetadataInteractionsInvite) *NullableVideoMetadataInteractionsInvite {
	return &NullableVideoMetadataInteractionsInvite{value: val, isSet: true}
}

func (v NullableVideoMetadataInteractionsInvite) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableVideoMetadataInteractionsInvite) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


