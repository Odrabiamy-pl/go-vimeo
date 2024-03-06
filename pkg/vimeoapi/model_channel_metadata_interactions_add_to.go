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

// checks if the ChannelMetadataInteractionsAddTo type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &ChannelMetadataInteractionsAddTo{}

// ChannelMetadataInteractionsAddTo When a channel appears in the context of adding or removing a video from it (`/videos/{video_id}/available_channels`), include information about adding or removing the video. This data requires a bearer token with the `private` scope.
type ChannelMetadataInteractionsAddTo struct {
	// An array of HTTP methods permitted on this URI. This data requires a bearer token with the `private` scope.
	Options []string `json:"options"`
	// The API URI that resolves to the connection data. This data requires a bearer token with the `private` scope.
	Uri string `json:"uri"`
}

// NewChannelMetadataInteractionsAddTo instantiates a new ChannelMetadataInteractionsAddTo object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewChannelMetadataInteractionsAddTo(options []string, uri string) *ChannelMetadataInteractionsAddTo {
	this := ChannelMetadataInteractionsAddTo{}
	this.Options = options
	this.Uri = uri
	return &this
}

// NewChannelMetadataInteractionsAddToWithDefaults instantiates a new ChannelMetadataInteractionsAddTo object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewChannelMetadataInteractionsAddToWithDefaults() *ChannelMetadataInteractionsAddTo {
	this := ChannelMetadataInteractionsAddTo{}
	return &this
}

// GetOptions returns the Options field value
func (o *ChannelMetadataInteractionsAddTo) GetOptions() []string {
	if o == nil {
		var ret []string
		return ret
	}

	return o.Options
}

// GetOptionsOk returns a tuple with the Options field value
// and a boolean to check if the value has been set.
func (o *ChannelMetadataInteractionsAddTo) GetOptionsOk() ([]string, bool) {
	if o == nil {
		return nil, false
	}
	return o.Options, true
}

// SetOptions sets field value
func (o *ChannelMetadataInteractionsAddTo) SetOptions(v []string) {
	o.Options = v
}

// GetUri returns the Uri field value
func (o *ChannelMetadataInteractionsAddTo) GetUri() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Uri
}

// GetUriOk returns a tuple with the Uri field value
// and a boolean to check if the value has been set.
func (o *ChannelMetadataInteractionsAddTo) GetUriOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Uri, true
}

// SetUri sets field value
func (o *ChannelMetadataInteractionsAddTo) SetUri(v string) {
	o.Uri = v
}

func (o ChannelMetadataInteractionsAddTo) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o ChannelMetadataInteractionsAddTo) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["options"] = o.Options
	toSerialize["uri"] = o.Uri
	return toSerialize, nil
}

type NullableChannelMetadataInteractionsAddTo struct {
	value *ChannelMetadataInteractionsAddTo
	isSet bool
}

func (v NullableChannelMetadataInteractionsAddTo) Get() *ChannelMetadataInteractionsAddTo {
	return v.value
}

func (v *NullableChannelMetadataInteractionsAddTo) Set(val *ChannelMetadataInteractionsAddTo) {
	v.value = val
	v.isSet = true
}

func (v NullableChannelMetadataInteractionsAddTo) IsSet() bool {
	return v.isSet
}

func (v *NullableChannelMetadataInteractionsAddTo) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableChannelMetadataInteractionsAddTo(val *ChannelMetadataInteractionsAddTo) *NullableChannelMetadataInteractionsAddTo {
	return &NullableChannelMetadataInteractionsAddTo{value: val, isSet: true}
}

func (v NullableChannelMetadataInteractionsAddTo) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableChannelMetadataInteractionsAddTo) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
