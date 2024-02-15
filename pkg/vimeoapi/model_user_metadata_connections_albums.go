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

// checks if the UserMetadataConnectionsAlbums type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &UserMetadataConnectionsAlbums{}

// UserMetadataConnectionsAlbums Information about the showcases created by the authenticated user.
type UserMetadataConnectionsAlbums struct {
	// An array of HTTP methods permitted on this URI.
	Options []string `json:"options"`
	// The total number of showcases on this connection.
	Total float32 `json:"total"`
	// The API URI that resolves to the connection data.
	Uri string `json:"uri"`
}

// NewUserMetadataConnectionsAlbums instantiates a new UserMetadataConnectionsAlbums object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewUserMetadataConnectionsAlbums(options []string, total float32, uri string) *UserMetadataConnectionsAlbums {
	this := UserMetadataConnectionsAlbums{}
	this.Options = options
	this.Total = total
	this.Uri = uri
	return &this
}

// NewUserMetadataConnectionsAlbumsWithDefaults instantiates a new UserMetadataConnectionsAlbums object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewUserMetadataConnectionsAlbumsWithDefaults() *UserMetadataConnectionsAlbums {
	this := UserMetadataConnectionsAlbums{}
	return &this
}

// GetOptions returns the Options field value
func (o *UserMetadataConnectionsAlbums) GetOptions() []string {
	if o == nil {
		var ret []string
		return ret
	}

	return o.Options
}

// GetOptionsOk returns a tuple with the Options field value
// and a boolean to check if the value has been set.
func (o *UserMetadataConnectionsAlbums) GetOptionsOk() ([]string, bool) {
	if o == nil {
		return nil, false
	}
	return o.Options, true
}

// SetOptions sets field value
func (o *UserMetadataConnectionsAlbums) SetOptions(v []string) {
	o.Options = v
}

// GetTotal returns the Total field value
func (o *UserMetadataConnectionsAlbums) GetTotal() float32 {
	if o == nil {
		var ret float32
		return ret
	}

	return o.Total
}

// GetTotalOk returns a tuple with the Total field value
// and a boolean to check if the value has been set.
func (o *UserMetadataConnectionsAlbums) GetTotalOk() (*float32, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Total, true
}

// SetTotal sets field value
func (o *UserMetadataConnectionsAlbums) SetTotal(v float32) {
	o.Total = v
}

// GetUri returns the Uri field value
func (o *UserMetadataConnectionsAlbums) GetUri() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Uri
}

// GetUriOk returns a tuple with the Uri field value
// and a boolean to check if the value has been set.
func (o *UserMetadataConnectionsAlbums) GetUriOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Uri, true
}

// SetUri sets field value
func (o *UserMetadataConnectionsAlbums) SetUri(v string) {
	o.Uri = v
}

func (o UserMetadataConnectionsAlbums) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o UserMetadataConnectionsAlbums) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["options"] = o.Options
	toSerialize["total"] = o.Total
	toSerialize["uri"] = o.Uri
	return toSerialize, nil
}

type NullableUserMetadataConnectionsAlbums struct {
	value *UserMetadataConnectionsAlbums
	isSet bool
}

func (v NullableUserMetadataConnectionsAlbums) Get() *UserMetadataConnectionsAlbums {
	return v.value
}

func (v *NullableUserMetadataConnectionsAlbums) Set(val *UserMetadataConnectionsAlbums) {
	v.value = val
	v.isSet = true
}

func (v NullableUserMetadataConnectionsAlbums) IsSet() bool {
	return v.isSet
}

func (v *NullableUserMetadataConnectionsAlbums) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableUserMetadataConnectionsAlbums(val *UserMetadataConnectionsAlbums) *NullableUserMetadataConnectionsAlbums {
	return &NullableUserMetadataConnectionsAlbums{value: val, isSet: true}
}

func (v NullableUserMetadataConnectionsAlbums) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableUserMetadataConnectionsAlbums) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


