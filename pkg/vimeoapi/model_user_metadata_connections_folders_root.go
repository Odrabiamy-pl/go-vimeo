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

// checks if the UserMetadataConnectionsFoldersRoot type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &UserMetadataConnectionsFoldersRoot{}

// UserMetadataConnectionsFoldersRoot Information about the authenticated user's root level folders and videos.
type UserMetadataConnectionsFoldersRoot struct {
	// An array of HTTP methods permitted on this URI.
	Options []string `json:"options"`
	// The API URI that resolves to the connection data.
	Uri string `json:"uri"`
}

// NewUserMetadataConnectionsFoldersRoot instantiates a new UserMetadataConnectionsFoldersRoot object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewUserMetadataConnectionsFoldersRoot(options []string, uri string) *UserMetadataConnectionsFoldersRoot {
	this := UserMetadataConnectionsFoldersRoot{}
	this.Options = options
	this.Uri = uri
	return &this
}

// NewUserMetadataConnectionsFoldersRootWithDefaults instantiates a new UserMetadataConnectionsFoldersRoot object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewUserMetadataConnectionsFoldersRootWithDefaults() *UserMetadataConnectionsFoldersRoot {
	this := UserMetadataConnectionsFoldersRoot{}
	return &this
}

// GetOptions returns the Options field value
func (o *UserMetadataConnectionsFoldersRoot) GetOptions() []string {
	if o == nil {
		var ret []string
		return ret
	}

	return o.Options
}

// GetOptionsOk returns a tuple with the Options field value
// and a boolean to check if the value has been set.
func (o *UserMetadataConnectionsFoldersRoot) GetOptionsOk() ([]string, bool) {
	if o == nil {
		return nil, false
	}
	return o.Options, true
}

// SetOptions sets field value
func (o *UserMetadataConnectionsFoldersRoot) SetOptions(v []string) {
	o.Options = v
}

// GetUri returns the Uri field value
func (o *UserMetadataConnectionsFoldersRoot) GetUri() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Uri
}

// GetUriOk returns a tuple with the Uri field value
// and a boolean to check if the value has been set.
func (o *UserMetadataConnectionsFoldersRoot) GetUriOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Uri, true
}

// SetUri sets field value
func (o *UserMetadataConnectionsFoldersRoot) SetUri(v string) {
	o.Uri = v
}

func (o UserMetadataConnectionsFoldersRoot) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o UserMetadataConnectionsFoldersRoot) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["options"] = o.Options
	toSerialize["uri"] = o.Uri
	return toSerialize, nil
}

type NullableUserMetadataConnectionsFoldersRoot struct {
	value *UserMetadataConnectionsFoldersRoot
	isSet bool
}

func (v NullableUserMetadataConnectionsFoldersRoot) Get() *UserMetadataConnectionsFoldersRoot {
	return v.value
}

func (v *NullableUserMetadataConnectionsFoldersRoot) Set(val *UserMetadataConnectionsFoldersRoot) {
	v.value = val
	v.isSet = true
}

func (v NullableUserMetadataConnectionsFoldersRoot) IsSet() bool {
	return v.isSet
}

func (v *NullableUserMetadataConnectionsFoldersRoot) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableUserMetadataConnectionsFoldersRoot(val *UserMetadataConnectionsFoldersRoot) *NullableUserMetadataConnectionsFoldersRoot {
	return &NullableUserMetadataConnectionsFoldersRoot{value: val, isSet: true}
}

func (v NullableUserMetadataConnectionsFoldersRoot) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableUserMetadataConnectionsFoldersRoot) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
