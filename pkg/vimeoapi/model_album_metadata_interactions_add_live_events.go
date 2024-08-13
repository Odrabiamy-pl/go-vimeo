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

// checks if the AlbumMetadataInteractionsAddLiveEvents type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &AlbumMetadataInteractionsAddLiveEvents{}

// AlbumMetadataInteractionsAddLiveEvents An action indicating that the authenticated user is an administrator of the showcase and may therefore add events. This data requires a bearer token with the `private` scope.
type AlbumMetadataInteractionsAddLiveEvents struct {
	// An array of HTTP methods permitted on this URI. This data requires a bearer token with the `private` scope.
	Options []string `json:"options"`
	// The API URI that resolves to the connection data. This data requires a bearer token with the `private` scope.
	Uri string `json:"uri"`
}

// NewAlbumMetadataInteractionsAddLiveEvents instantiates a new AlbumMetadataInteractionsAddLiveEvents object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewAlbumMetadataInteractionsAddLiveEvents(options []string, uri string) *AlbumMetadataInteractionsAddLiveEvents {
	this := AlbumMetadataInteractionsAddLiveEvents{}
	this.Options = options
	this.Uri = uri
	return &this
}

// NewAlbumMetadataInteractionsAddLiveEventsWithDefaults instantiates a new AlbumMetadataInteractionsAddLiveEvents object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewAlbumMetadataInteractionsAddLiveEventsWithDefaults() *AlbumMetadataInteractionsAddLiveEvents {
	this := AlbumMetadataInteractionsAddLiveEvents{}
	return &this
}

// GetOptions returns the Options field value
func (o *AlbumMetadataInteractionsAddLiveEvents) GetOptions() []string {
	if o == nil {
		var ret []string
		return ret
	}

	return o.Options
}

// GetOptionsOk returns a tuple with the Options field value
// and a boolean to check if the value has been set.
func (o *AlbumMetadataInteractionsAddLiveEvents) GetOptionsOk() ([]string, bool) {
	if o == nil {
		return nil, false
	}
	return o.Options, true
}

// SetOptions sets field value
func (o *AlbumMetadataInteractionsAddLiveEvents) SetOptions(v []string) {
	o.Options = v
}

// GetUri returns the Uri field value
func (o *AlbumMetadataInteractionsAddLiveEvents) GetUri() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Uri
}

// GetUriOk returns a tuple with the Uri field value
// and a boolean to check if the value has been set.
func (o *AlbumMetadataInteractionsAddLiveEvents) GetUriOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Uri, true
}

// SetUri sets field value
func (o *AlbumMetadataInteractionsAddLiveEvents) SetUri(v string) {
	o.Uri = v
}

func (o AlbumMetadataInteractionsAddLiveEvents) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o AlbumMetadataInteractionsAddLiveEvents) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["options"] = o.Options
	toSerialize["uri"] = o.Uri
	return toSerialize, nil
}

type NullableAlbumMetadataInteractionsAddLiveEvents struct {
	value *AlbumMetadataInteractionsAddLiveEvents
	isSet bool
}

func (v NullableAlbumMetadataInteractionsAddLiveEvents) Get() *AlbumMetadataInteractionsAddLiveEvents {
	return v.value
}

func (v *NullableAlbumMetadataInteractionsAddLiveEvents) Set(val *AlbumMetadataInteractionsAddLiveEvents) {
	v.value = val
	v.isSet = true
}

func (v NullableAlbumMetadataInteractionsAddLiveEvents) IsSet() bool {
	return v.isSet
}

func (v *NullableAlbumMetadataInteractionsAddLiveEvents) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableAlbumMetadataInteractionsAddLiveEvents(val *AlbumMetadataInteractionsAddLiveEvents) *NullableAlbumMetadataInteractionsAddLiveEvents {
	return &NullableAlbumMetadataInteractionsAddLiveEvents{value: val, isSet: true}
}

func (v NullableAlbumMetadataInteractionsAddLiveEvents) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableAlbumMetadataInteractionsAddLiveEvents) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


