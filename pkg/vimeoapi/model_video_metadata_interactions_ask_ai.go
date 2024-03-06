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

// checks if the VideoMetadataInteractionsAskAi type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &VideoMetadataInteractionsAskAi{}

// VideoMetadataInteractionsAskAi Information about where and how to submit questions to the AI service for this video.
type VideoMetadataInteractionsAskAi struct {
	// An array of HTTP methods permitted on this URI.
	Options []string `json:"options"`
	// The API URI that resolves to the connection data.
	Uri string `json:"uri"`
}

// NewVideoMetadataInteractionsAskAi instantiates a new VideoMetadataInteractionsAskAi object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewVideoMetadataInteractionsAskAi(options []string, uri string) *VideoMetadataInteractionsAskAi {
	this := VideoMetadataInteractionsAskAi{}
	this.Options = options
	this.Uri = uri
	return &this
}

// NewVideoMetadataInteractionsAskAiWithDefaults instantiates a new VideoMetadataInteractionsAskAi object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewVideoMetadataInteractionsAskAiWithDefaults() *VideoMetadataInteractionsAskAi {
	this := VideoMetadataInteractionsAskAi{}
	return &this
}

// GetOptions returns the Options field value
func (o *VideoMetadataInteractionsAskAi) GetOptions() []string {
	if o == nil {
		var ret []string
		return ret
	}

	return o.Options
}

// GetOptionsOk returns a tuple with the Options field value
// and a boolean to check if the value has been set.
func (o *VideoMetadataInteractionsAskAi) GetOptionsOk() ([]string, bool) {
	if o == nil {
		return nil, false
	}
	return o.Options, true
}

// SetOptions sets field value
func (o *VideoMetadataInteractionsAskAi) SetOptions(v []string) {
	o.Options = v
}

// GetUri returns the Uri field value
func (o *VideoMetadataInteractionsAskAi) GetUri() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Uri
}

// GetUriOk returns a tuple with the Uri field value
// and a boolean to check if the value has been set.
func (o *VideoMetadataInteractionsAskAi) GetUriOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Uri, true
}

// SetUri sets field value
func (o *VideoMetadataInteractionsAskAi) SetUri(v string) {
	o.Uri = v
}

func (o VideoMetadataInteractionsAskAi) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o VideoMetadataInteractionsAskAi) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["options"] = o.Options
	toSerialize["uri"] = o.Uri
	return toSerialize, nil
}

type NullableVideoMetadataInteractionsAskAi struct {
	value *VideoMetadataInteractionsAskAi
	isSet bool
}

func (v NullableVideoMetadataInteractionsAskAi) Get() *VideoMetadataInteractionsAskAi {
	return v.value
}

func (v *NullableVideoMetadataInteractionsAskAi) Set(val *VideoMetadataInteractionsAskAi) {
	v.value = val
	v.isSet = true
}

func (v NullableVideoMetadataInteractionsAskAi) IsSet() bool {
	return v.isSet
}

func (v *NullableVideoMetadataInteractionsAskAi) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableVideoMetadataInteractionsAskAi(val *VideoMetadataInteractionsAskAi) *NullableVideoMetadataInteractionsAskAi {
	return &NullableVideoMetadataInteractionsAskAi{value: val, isSet: true}
}

func (v NullableVideoMetadataInteractionsAskAi) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableVideoMetadataInteractionsAskAi) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
