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

// checks if the AlbumMetadataInteractionsAddVideos type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &AlbumMetadataInteractionsAddVideos{}

// AlbumMetadataInteractionsAddVideos An action indicating that the authenticated user is an administrator of the showcase and may therefore add videos. This data requires a bearer token with the `private` scope.
type AlbumMetadataInteractionsAddVideos struct {
	// An array of HTTP methods permitted on this URI. This data requires a bearer token with the `private` scope.
	Options []string `json:"options"`
	// The API URI that resolves to the connection data. This data requires a bearer token with the `private` scope.
	Uri string `json:"uri"`
}

// NewAlbumMetadataInteractionsAddVideos instantiates a new AlbumMetadataInteractionsAddVideos object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewAlbumMetadataInteractionsAddVideos(options []string, uri string) *AlbumMetadataInteractionsAddVideos {
	this := AlbumMetadataInteractionsAddVideos{}
	this.Options = options
	this.Uri = uri
	return &this
}

// NewAlbumMetadataInteractionsAddVideosWithDefaults instantiates a new AlbumMetadataInteractionsAddVideos object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewAlbumMetadataInteractionsAddVideosWithDefaults() *AlbumMetadataInteractionsAddVideos {
	this := AlbumMetadataInteractionsAddVideos{}
	return &this
}

// GetOptions returns the Options field value
func (o *AlbumMetadataInteractionsAddVideos) GetOptions() []string {
	if o == nil {
		var ret []string
		return ret
	}

	return o.Options
}

// GetOptionsOk returns a tuple with the Options field value
// and a boolean to check if the value has been set.
func (o *AlbumMetadataInteractionsAddVideos) GetOptionsOk() ([]string, bool) {
	if o == nil {
		return nil, false
	}
	return o.Options, true
}

// SetOptions sets field value
func (o *AlbumMetadataInteractionsAddVideos) SetOptions(v []string) {
	o.Options = v
}

// GetUri returns the Uri field value
func (o *AlbumMetadataInteractionsAddVideos) GetUri() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Uri
}

// GetUriOk returns a tuple with the Uri field value
// and a boolean to check if the value has been set.
func (o *AlbumMetadataInteractionsAddVideos) GetUriOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Uri, true
}

// SetUri sets field value
func (o *AlbumMetadataInteractionsAddVideos) SetUri(v string) {
	o.Uri = v
}

func (o AlbumMetadataInteractionsAddVideos) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o AlbumMetadataInteractionsAddVideos) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["options"] = o.Options
	toSerialize["uri"] = o.Uri
	return toSerialize, nil
}

type NullableAlbumMetadataInteractionsAddVideos struct {
	value *AlbumMetadataInteractionsAddVideos
	isSet bool
}

func (v NullableAlbumMetadataInteractionsAddVideos) Get() *AlbumMetadataInteractionsAddVideos {
	return v.value
}

func (v *NullableAlbumMetadataInteractionsAddVideos) Set(val *AlbumMetadataInteractionsAddVideos) {
	v.value = val
	v.isSet = true
}

func (v NullableAlbumMetadataInteractionsAddVideos) IsSet() bool {
	return v.isSet
}

func (v *NullableAlbumMetadataInteractionsAddVideos) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableAlbumMetadataInteractionsAddVideos(val *AlbumMetadataInteractionsAddVideos) *NullableAlbumMetadataInteractionsAddVideos {
	return &NullableAlbumMetadataInteractionsAddVideos{value: val, isSet: true}
}

func (v NullableAlbumMetadataInteractionsAddVideos) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableAlbumMetadataInteractionsAddVideos) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
