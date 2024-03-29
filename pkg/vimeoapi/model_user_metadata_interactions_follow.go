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

// checks if the UserMetadataInteractionsFollow type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &UserMetadataInteractionsFollow{}

// UserMetadataInteractionsFollow Information about the followed status of the authenticated user.
type UserMetadataInteractionsFollow struct {
	// Whether the authenticated user is following the requested user.
	Added bool `json:"added"`
	// An array of the HTTP methods permitted on this URI.
	Options []string `json:"options"`
	// The URI to follow the requested user.
	Uri string `json:"uri"`
}

// NewUserMetadataInteractionsFollow instantiates a new UserMetadataInteractionsFollow object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewUserMetadataInteractionsFollow(added bool, options []string, uri string) *UserMetadataInteractionsFollow {
	this := UserMetadataInteractionsFollow{}
	this.Added = added
	this.Options = options
	this.Uri = uri
	return &this
}

// NewUserMetadataInteractionsFollowWithDefaults instantiates a new UserMetadataInteractionsFollow object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewUserMetadataInteractionsFollowWithDefaults() *UserMetadataInteractionsFollow {
	this := UserMetadataInteractionsFollow{}
	return &this
}

// GetAdded returns the Added field value
func (o *UserMetadataInteractionsFollow) GetAdded() bool {
	if o == nil {
		var ret bool
		return ret
	}

	return o.Added
}

// GetAddedOk returns a tuple with the Added field value
// and a boolean to check if the value has been set.
func (o *UserMetadataInteractionsFollow) GetAddedOk() (*bool, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Added, true
}

// SetAdded sets field value
func (o *UserMetadataInteractionsFollow) SetAdded(v bool) {
	o.Added = v
}

// GetOptions returns the Options field value
func (o *UserMetadataInteractionsFollow) GetOptions() []string {
	if o == nil {
		var ret []string
		return ret
	}

	return o.Options
}

// GetOptionsOk returns a tuple with the Options field value
// and a boolean to check if the value has been set.
func (o *UserMetadataInteractionsFollow) GetOptionsOk() ([]string, bool) {
	if o == nil {
		return nil, false
	}
	return o.Options, true
}

// SetOptions sets field value
func (o *UserMetadataInteractionsFollow) SetOptions(v []string) {
	o.Options = v
}

// GetUri returns the Uri field value
func (o *UserMetadataInteractionsFollow) GetUri() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Uri
}

// GetUriOk returns a tuple with the Uri field value
// and a boolean to check if the value has been set.
func (o *UserMetadataInteractionsFollow) GetUriOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Uri, true
}

// SetUri sets field value
func (o *UserMetadataInteractionsFollow) SetUri(v string) {
	o.Uri = v
}

func (o UserMetadataInteractionsFollow) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o UserMetadataInteractionsFollow) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["added"] = o.Added
	toSerialize["options"] = o.Options
	toSerialize["uri"] = o.Uri
	return toSerialize, nil
}

type NullableUserMetadataInteractionsFollow struct {
	value *UserMetadataInteractionsFollow
	isSet bool
}

func (v NullableUserMetadataInteractionsFollow) Get() *UserMetadataInteractionsFollow {
	return v.value
}

func (v *NullableUserMetadataInteractionsFollow) Set(val *UserMetadataInteractionsFollow) {
	v.value = val
	v.isSet = true
}

func (v NullableUserMetadataInteractionsFollow) IsSet() bool {
	return v.isSet
}

func (v *NullableUserMetadataInteractionsFollow) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableUserMetadataInteractionsFollow(val *UserMetadataInteractionsFollow) *NullableUserMetadataInteractionsFollow {
	return &NullableUserMetadataInteractionsFollow{value: val, isSet: true}
}

func (v NullableUserMetadataInteractionsFollow) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableUserMetadataInteractionsFollow) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
