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

// checks if the UserMetadataConnectionsFolders type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &UserMetadataConnectionsFolders{}

// UserMetadataConnectionsFolders Information about the authenticated user's folders.
type UserMetadataConnectionsFolders struct {
	// An array of HTTP methods permitted on this URI.
	Options []string `json:"options"`
	// The total number of folders on this connection.
	Total float32 `json:"total"`
	// The API URI that resolves to the connection data.
	Uri string `json:"uri"`
}

// NewUserMetadataConnectionsFolders instantiates a new UserMetadataConnectionsFolders object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewUserMetadataConnectionsFolders(options []string, total float32, uri string) *UserMetadataConnectionsFolders {
	this := UserMetadataConnectionsFolders{}
	this.Options = options
	this.Total = total
	this.Uri = uri
	return &this
}

// NewUserMetadataConnectionsFoldersWithDefaults instantiates a new UserMetadataConnectionsFolders object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewUserMetadataConnectionsFoldersWithDefaults() *UserMetadataConnectionsFolders {
	this := UserMetadataConnectionsFolders{}
	return &this
}

// GetOptions returns the Options field value
func (o *UserMetadataConnectionsFolders) GetOptions() []string {
	if o == nil {
		var ret []string
		return ret
	}

	return o.Options
}

// GetOptionsOk returns a tuple with the Options field value
// and a boolean to check if the value has been set.
func (o *UserMetadataConnectionsFolders) GetOptionsOk() ([]string, bool) {
	if o == nil {
		return nil, false
	}
	return o.Options, true
}

// SetOptions sets field value
func (o *UserMetadataConnectionsFolders) SetOptions(v []string) {
	o.Options = v
}

// GetTotal returns the Total field value
func (o *UserMetadataConnectionsFolders) GetTotal() float32 {
	if o == nil {
		var ret float32
		return ret
	}

	return o.Total
}

// GetTotalOk returns a tuple with the Total field value
// and a boolean to check if the value has been set.
func (o *UserMetadataConnectionsFolders) GetTotalOk() (*float32, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Total, true
}

// SetTotal sets field value
func (o *UserMetadataConnectionsFolders) SetTotal(v float32) {
	o.Total = v
}

// GetUri returns the Uri field value
func (o *UserMetadataConnectionsFolders) GetUri() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Uri
}

// GetUriOk returns a tuple with the Uri field value
// and a boolean to check if the value has been set.
func (o *UserMetadataConnectionsFolders) GetUriOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Uri, true
}

// SetUri sets field value
func (o *UserMetadataConnectionsFolders) SetUri(v string) {
	o.Uri = v
}

func (o UserMetadataConnectionsFolders) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o UserMetadataConnectionsFolders) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["options"] = o.Options
	toSerialize["total"] = o.Total
	toSerialize["uri"] = o.Uri
	return toSerialize, nil
}

type NullableUserMetadataConnectionsFolders struct {
	value *UserMetadataConnectionsFolders
	isSet bool
}

func (v NullableUserMetadataConnectionsFolders) Get() *UserMetadataConnectionsFolders {
	return v.value
}

func (v *NullableUserMetadataConnectionsFolders) Set(val *UserMetadataConnectionsFolders) {
	v.value = val
	v.isSet = true
}

func (v NullableUserMetadataConnectionsFolders) IsSet() bool {
	return v.isSet
}

func (v *NullableUserMetadataConnectionsFolders) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableUserMetadataConnectionsFolders(val *UserMetadataConnectionsFolders) *NullableUserMetadataConnectionsFolders {
	return &NullableUserMetadataConnectionsFolders{value: val, isSet: true}
}

func (v NullableUserMetadataConnectionsFolders) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableUserMetadataConnectionsFolders) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
