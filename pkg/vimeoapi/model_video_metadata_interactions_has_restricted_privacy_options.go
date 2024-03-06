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

// checks if the VideoMetadataInteractionsHasRestrictedPrivacyOptions type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &VideoMetadataInteractionsHasRestrictedPrivacyOptions{}

// VideoMetadataInteractionsHasRestrictedPrivacyOptions Information about whether the video has restricted privacy options.
type VideoMetadataInteractionsHasRestrictedPrivacyOptions struct {
	// An array of HTTP methods permitted on this URI.
	Options []string `json:"options"`
	// The API URI that resolves to the connection data.
	Uri string `json:"uri"`
}

// NewVideoMetadataInteractionsHasRestrictedPrivacyOptions instantiates a new VideoMetadataInteractionsHasRestrictedPrivacyOptions object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewVideoMetadataInteractionsHasRestrictedPrivacyOptions(options []string, uri string) *VideoMetadataInteractionsHasRestrictedPrivacyOptions {
	this := VideoMetadataInteractionsHasRestrictedPrivacyOptions{}
	this.Options = options
	this.Uri = uri
	return &this
}

// NewVideoMetadataInteractionsHasRestrictedPrivacyOptionsWithDefaults instantiates a new VideoMetadataInteractionsHasRestrictedPrivacyOptions object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewVideoMetadataInteractionsHasRestrictedPrivacyOptionsWithDefaults() *VideoMetadataInteractionsHasRestrictedPrivacyOptions {
	this := VideoMetadataInteractionsHasRestrictedPrivacyOptions{}
	return &this
}

// GetOptions returns the Options field value
func (o *VideoMetadataInteractionsHasRestrictedPrivacyOptions) GetOptions() []string {
	if o == nil {
		var ret []string
		return ret
	}

	return o.Options
}

// GetOptionsOk returns a tuple with the Options field value
// and a boolean to check if the value has been set.
func (o *VideoMetadataInteractionsHasRestrictedPrivacyOptions) GetOptionsOk() ([]string, bool) {
	if o == nil {
		return nil, false
	}
	return o.Options, true
}

// SetOptions sets field value
func (o *VideoMetadataInteractionsHasRestrictedPrivacyOptions) SetOptions(v []string) {
	o.Options = v
}

// GetUri returns the Uri field value
func (o *VideoMetadataInteractionsHasRestrictedPrivacyOptions) GetUri() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Uri
}

// GetUriOk returns a tuple with the Uri field value
// and a boolean to check if the value has been set.
func (o *VideoMetadataInteractionsHasRestrictedPrivacyOptions) GetUriOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Uri, true
}

// SetUri sets field value
func (o *VideoMetadataInteractionsHasRestrictedPrivacyOptions) SetUri(v string) {
	o.Uri = v
}

func (o VideoMetadataInteractionsHasRestrictedPrivacyOptions) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o VideoMetadataInteractionsHasRestrictedPrivacyOptions) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["options"] = o.Options
	toSerialize["uri"] = o.Uri
	return toSerialize, nil
}

type NullableVideoMetadataInteractionsHasRestrictedPrivacyOptions struct {
	value *VideoMetadataInteractionsHasRestrictedPrivacyOptions
	isSet bool
}

func (v NullableVideoMetadataInteractionsHasRestrictedPrivacyOptions) Get() *VideoMetadataInteractionsHasRestrictedPrivacyOptions {
	return v.value
}

func (v *NullableVideoMetadataInteractionsHasRestrictedPrivacyOptions) Set(val *VideoMetadataInteractionsHasRestrictedPrivacyOptions) {
	v.value = val
	v.isSet = true
}

func (v NullableVideoMetadataInteractionsHasRestrictedPrivacyOptions) IsSet() bool {
	return v.isSet
}

func (v *NullableVideoMetadataInteractionsHasRestrictedPrivacyOptions) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableVideoMetadataInteractionsHasRestrictedPrivacyOptions(val *VideoMetadataInteractionsHasRestrictedPrivacyOptions) *NullableVideoMetadataInteractionsHasRestrictedPrivacyOptions {
	return &NullableVideoMetadataInteractionsHasRestrictedPrivacyOptions{value: val, isSet: true}
}

func (v NullableVideoMetadataInteractionsHasRestrictedPrivacyOptions) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableVideoMetadataInteractionsHasRestrictedPrivacyOptions) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
