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

// checks if the CategoryMetadataConnectionsChannels type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &CategoryMetadataConnectionsChannels{}

// CategoryMetadataConnectionsChannels Information about the channels related to this category.
type CategoryMetadataConnectionsChannels struct {
	// An array of HTTP methods permitted on this URI.
	Options []string `json:"options"`
	// The total number of channels on this connection.
	Total float32 `json:"total"`
	// The API URI that resolves to the connection data.
	Uri string `json:"uri"`
}

// NewCategoryMetadataConnectionsChannels instantiates a new CategoryMetadataConnectionsChannels object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewCategoryMetadataConnectionsChannels(options []string, total float32, uri string) *CategoryMetadataConnectionsChannels {
	this := CategoryMetadataConnectionsChannels{}
	this.Options = options
	this.Total = total
	this.Uri = uri
	return &this
}

// NewCategoryMetadataConnectionsChannelsWithDefaults instantiates a new CategoryMetadataConnectionsChannels object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewCategoryMetadataConnectionsChannelsWithDefaults() *CategoryMetadataConnectionsChannels {
	this := CategoryMetadataConnectionsChannels{}
	return &this
}

// GetOptions returns the Options field value
func (o *CategoryMetadataConnectionsChannels) GetOptions() []string {
	if o == nil {
		var ret []string
		return ret
	}

	return o.Options
}

// GetOptionsOk returns a tuple with the Options field value
// and a boolean to check if the value has been set.
func (o *CategoryMetadataConnectionsChannels) GetOptionsOk() ([]string, bool) {
	if o == nil {
		return nil, false
	}
	return o.Options, true
}

// SetOptions sets field value
func (o *CategoryMetadataConnectionsChannels) SetOptions(v []string) {
	o.Options = v
}

// GetTotal returns the Total field value
func (o *CategoryMetadataConnectionsChannels) GetTotal() float32 {
	if o == nil {
		var ret float32
		return ret
	}

	return o.Total
}

// GetTotalOk returns a tuple with the Total field value
// and a boolean to check if the value has been set.
func (o *CategoryMetadataConnectionsChannels) GetTotalOk() (*float32, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Total, true
}

// SetTotal sets field value
func (o *CategoryMetadataConnectionsChannels) SetTotal(v float32) {
	o.Total = v
}

// GetUri returns the Uri field value
func (o *CategoryMetadataConnectionsChannels) GetUri() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Uri
}

// GetUriOk returns a tuple with the Uri field value
// and a boolean to check if the value has been set.
func (o *CategoryMetadataConnectionsChannels) GetUriOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Uri, true
}

// SetUri sets field value
func (o *CategoryMetadataConnectionsChannels) SetUri(v string) {
	o.Uri = v
}

func (o CategoryMetadataConnectionsChannels) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o CategoryMetadataConnectionsChannels) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["options"] = o.Options
	toSerialize["total"] = o.Total
	toSerialize["uri"] = o.Uri
	return toSerialize, nil
}

type NullableCategoryMetadataConnectionsChannels struct {
	value *CategoryMetadataConnectionsChannels
	isSet bool
}

func (v NullableCategoryMetadataConnectionsChannels) Get() *CategoryMetadataConnectionsChannels {
	return v.value
}

func (v *NullableCategoryMetadataConnectionsChannels) Set(val *CategoryMetadataConnectionsChannels) {
	v.value = val
	v.isSet = true
}

func (v NullableCategoryMetadataConnectionsChannels) IsSet() bool {
	return v.isSet
}

func (v *NullableCategoryMetadataConnectionsChannels) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableCategoryMetadataConnectionsChannels(val *CategoryMetadataConnectionsChannels) *NullableCategoryMetadataConnectionsChannels {
	return &NullableCategoryMetadataConnectionsChannels{value: val, isSet: true}
}

func (v NullableCategoryMetadataConnectionsChannels) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableCategoryMetadataConnectionsChannels) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


