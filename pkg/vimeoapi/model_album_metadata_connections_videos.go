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

// checks if the AlbumMetadataConnectionsVideos type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &AlbumMetadataConnectionsVideos{}

// AlbumMetadataConnectionsVideos Information about the videos that belong to the showcase.
type AlbumMetadataConnectionsVideos struct {
	// An array of the HTTP methods permitted on this URI.
	Options []string `json:"options"`
	// The total number of videos on this connection.
	Total float32 `json:"total"`
	// The API URI that resolves to the connection data.
	Uri string `json:"uri"`
}

// NewAlbumMetadataConnectionsVideos instantiates a new AlbumMetadataConnectionsVideos object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewAlbumMetadataConnectionsVideos(options []string, total float32, uri string) *AlbumMetadataConnectionsVideos {
	this := AlbumMetadataConnectionsVideos{}
	this.Options = options
	this.Total = total
	this.Uri = uri
	return &this
}

// NewAlbumMetadataConnectionsVideosWithDefaults instantiates a new AlbumMetadataConnectionsVideos object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewAlbumMetadataConnectionsVideosWithDefaults() *AlbumMetadataConnectionsVideos {
	this := AlbumMetadataConnectionsVideos{}
	return &this
}

// GetOptions returns the Options field value
func (o *AlbumMetadataConnectionsVideos) GetOptions() []string {
	if o == nil {
		var ret []string
		return ret
	}

	return o.Options
}

// GetOptionsOk returns a tuple with the Options field value
// and a boolean to check if the value has been set.
func (o *AlbumMetadataConnectionsVideos) GetOptionsOk() ([]string, bool) {
	if o == nil {
		return nil, false
	}
	return o.Options, true
}

// SetOptions sets field value
func (o *AlbumMetadataConnectionsVideos) SetOptions(v []string) {
	o.Options = v
}

// GetTotal returns the Total field value
func (o *AlbumMetadataConnectionsVideos) GetTotal() float32 {
	if o == nil {
		var ret float32
		return ret
	}

	return o.Total
}

// GetTotalOk returns a tuple with the Total field value
// and a boolean to check if the value has been set.
func (o *AlbumMetadataConnectionsVideos) GetTotalOk() (*float32, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Total, true
}

// SetTotal sets field value
func (o *AlbumMetadataConnectionsVideos) SetTotal(v float32) {
	o.Total = v
}

// GetUri returns the Uri field value
func (o *AlbumMetadataConnectionsVideos) GetUri() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Uri
}

// GetUriOk returns a tuple with the Uri field value
// and a boolean to check if the value has been set.
func (o *AlbumMetadataConnectionsVideos) GetUriOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Uri, true
}

// SetUri sets field value
func (o *AlbumMetadataConnectionsVideos) SetUri(v string) {
	o.Uri = v
}

func (o AlbumMetadataConnectionsVideos) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o AlbumMetadataConnectionsVideos) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["options"] = o.Options
	toSerialize["total"] = o.Total
	toSerialize["uri"] = o.Uri
	return toSerialize, nil
}

type NullableAlbumMetadataConnectionsVideos struct {
	value *AlbumMetadataConnectionsVideos
	isSet bool
}

func (v NullableAlbumMetadataConnectionsVideos) Get() *AlbumMetadataConnectionsVideos {
	return v.value
}

func (v *NullableAlbumMetadataConnectionsVideos) Set(val *AlbumMetadataConnectionsVideos) {
	v.value = val
	v.isSet = true
}

func (v NullableAlbumMetadataConnectionsVideos) IsSet() bool {
	return v.isSet
}

func (v *NullableAlbumMetadataConnectionsVideos) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableAlbumMetadataConnectionsVideos(val *AlbumMetadataConnectionsVideos) *NullableAlbumMetadataConnectionsVideos {
	return &NullableAlbumMetadataConnectionsVideos{value: val, isSet: true}
}

func (v NullableAlbumMetadataConnectionsVideos) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableAlbumMetadataConnectionsVideos) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


