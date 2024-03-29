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

// checks if the VideoMetadataInteractionsCreateEditor type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &VideoMetadataInteractionsCreateEditor{}

// VideoMetadataInteractionsCreateEditor Information about where and how to edit a video using the Vimeo Create editor.
type VideoMetadataInteractionsCreateEditor struct {
	// An array of HTTP methods permitted on this URI.
	Options []string `json:"options"`
	// The URI that resolves to the connection data.
	Uri string `json:"uri"`
}

// NewVideoMetadataInteractionsCreateEditor instantiates a new VideoMetadataInteractionsCreateEditor object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewVideoMetadataInteractionsCreateEditor(options []string, uri string) *VideoMetadataInteractionsCreateEditor {
	this := VideoMetadataInteractionsCreateEditor{}
	this.Options = options
	this.Uri = uri
	return &this
}

// NewVideoMetadataInteractionsCreateEditorWithDefaults instantiates a new VideoMetadataInteractionsCreateEditor object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewVideoMetadataInteractionsCreateEditorWithDefaults() *VideoMetadataInteractionsCreateEditor {
	this := VideoMetadataInteractionsCreateEditor{}
	return &this
}

// GetOptions returns the Options field value
func (o *VideoMetadataInteractionsCreateEditor) GetOptions() []string {
	if o == nil {
		var ret []string
		return ret
	}

	return o.Options
}

// GetOptionsOk returns a tuple with the Options field value
// and a boolean to check if the value has been set.
func (o *VideoMetadataInteractionsCreateEditor) GetOptionsOk() ([]string, bool) {
	if o == nil {
		return nil, false
	}
	return o.Options, true
}

// SetOptions sets field value
func (o *VideoMetadataInteractionsCreateEditor) SetOptions(v []string) {
	o.Options = v
}

// GetUri returns the Uri field value
func (o *VideoMetadataInteractionsCreateEditor) GetUri() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Uri
}

// GetUriOk returns a tuple with the Uri field value
// and a boolean to check if the value has been set.
func (o *VideoMetadataInteractionsCreateEditor) GetUriOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Uri, true
}

// SetUri sets field value
func (o *VideoMetadataInteractionsCreateEditor) SetUri(v string) {
	o.Uri = v
}

func (o VideoMetadataInteractionsCreateEditor) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o VideoMetadataInteractionsCreateEditor) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["options"] = o.Options
	toSerialize["uri"] = o.Uri
	return toSerialize, nil
}

type NullableVideoMetadataInteractionsCreateEditor struct {
	value *VideoMetadataInteractionsCreateEditor
	isSet bool
}

func (v NullableVideoMetadataInteractionsCreateEditor) Get() *VideoMetadataInteractionsCreateEditor {
	return v.value
}

func (v *NullableVideoMetadataInteractionsCreateEditor) Set(val *VideoMetadataInteractionsCreateEditor) {
	v.value = val
	v.isSet = true
}

func (v NullableVideoMetadataInteractionsCreateEditor) IsSet() bool {
	return v.isSet
}

func (v *NullableVideoMetadataInteractionsCreateEditor) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableVideoMetadataInteractionsCreateEditor(val *VideoMetadataInteractionsCreateEditor) *NullableVideoMetadataInteractionsCreateEditor {
	return &NullableVideoMetadataInteractionsCreateEditor{value: val, isSet: true}
}

func (v NullableVideoMetadataInteractionsCreateEditor) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableVideoMetadataInteractionsCreateEditor) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
