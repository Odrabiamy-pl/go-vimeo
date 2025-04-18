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

// checks if the LiveEventMetadataInteractionsEdit type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &LiveEventMetadataInteractionsEdit{}

// LiveEventMetadataInteractionsEdit Information about where and how to edit an item.
type LiveEventMetadataInteractionsEdit struct {
	// An array of HTTP methods permitted on this URI.
	Options []string `json:"options"`
	// The API URI that resolves to the connection data.
	Uri string `json:"uri"`
}

// NewLiveEventMetadataInteractionsEdit instantiates a new LiveEventMetadataInteractionsEdit object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewLiveEventMetadataInteractionsEdit(options []string, uri string) *LiveEventMetadataInteractionsEdit {
	this := LiveEventMetadataInteractionsEdit{}
	this.Options = options
	this.Uri = uri
	return &this
}

// NewLiveEventMetadataInteractionsEditWithDefaults instantiates a new LiveEventMetadataInteractionsEdit object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewLiveEventMetadataInteractionsEditWithDefaults() *LiveEventMetadataInteractionsEdit {
	this := LiveEventMetadataInteractionsEdit{}
	return &this
}

// GetOptions returns the Options field value
func (o *LiveEventMetadataInteractionsEdit) GetOptions() []string {
	if o == nil {
		var ret []string
		return ret
	}

	return o.Options
}

// GetOptionsOk returns a tuple with the Options field value
// and a boolean to check if the value has been set.
func (o *LiveEventMetadataInteractionsEdit) GetOptionsOk() ([]string, bool) {
	if o == nil {
		return nil, false
	}
	return o.Options, true
}

// SetOptions sets field value
func (o *LiveEventMetadataInteractionsEdit) SetOptions(v []string) {
	o.Options = v
}

// GetUri returns the Uri field value
func (o *LiveEventMetadataInteractionsEdit) GetUri() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Uri
}

// GetUriOk returns a tuple with the Uri field value
// and a boolean to check if the value has been set.
func (o *LiveEventMetadataInteractionsEdit) GetUriOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Uri, true
}

// SetUri sets field value
func (o *LiveEventMetadataInteractionsEdit) SetUri(v string) {
	o.Uri = v
}

func (o LiveEventMetadataInteractionsEdit) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o LiveEventMetadataInteractionsEdit) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["options"] = o.Options
	toSerialize["uri"] = o.Uri
	return toSerialize, nil
}

type NullableLiveEventMetadataInteractionsEdit struct {
	value *LiveEventMetadataInteractionsEdit
	isSet bool
}

func (v NullableLiveEventMetadataInteractionsEdit) Get() *LiveEventMetadataInteractionsEdit {
	return v.value
}

func (v *NullableLiveEventMetadataInteractionsEdit) Set(val *LiveEventMetadataInteractionsEdit) {
	v.value = val
	v.isSet = true
}

func (v NullableLiveEventMetadataInteractionsEdit) IsSet() bool {
	return v.isSet
}

func (v *NullableLiveEventMetadataInteractionsEdit) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableLiveEventMetadataInteractionsEdit(val *LiveEventMetadataInteractionsEdit) *NullableLiveEventMetadataInteractionsEdit {
	return &NullableLiveEventMetadataInteractionsEdit{value: val, isSet: true}
}

func (v NullableLiveEventMetadataInteractionsEdit) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableLiveEventMetadataInteractionsEdit) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


