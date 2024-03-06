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

// checks if the ProjectMetadataInteractionsView type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &ProjectMetadataInteractionsView{}

// ProjectMetadataInteractionsView The user's view permissions information for this project.
type ProjectMetadataInteractionsView struct {
	// An array of HTTP methods permitted on this URI.
	Options []string `json:"options"`
	// The API URI that resolves to the connection data.
	Uri string `json:"uri"`
}

// NewProjectMetadataInteractionsView instantiates a new ProjectMetadataInteractionsView object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewProjectMetadataInteractionsView(options []string, uri string) *ProjectMetadataInteractionsView {
	this := ProjectMetadataInteractionsView{}
	this.Options = options
	this.Uri = uri
	return &this
}

// NewProjectMetadataInteractionsViewWithDefaults instantiates a new ProjectMetadataInteractionsView object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewProjectMetadataInteractionsViewWithDefaults() *ProjectMetadataInteractionsView {
	this := ProjectMetadataInteractionsView{}
	return &this
}

// GetOptions returns the Options field value
func (o *ProjectMetadataInteractionsView) GetOptions() []string {
	if o == nil {
		var ret []string
		return ret
	}

	return o.Options
}

// GetOptionsOk returns a tuple with the Options field value
// and a boolean to check if the value has been set.
func (o *ProjectMetadataInteractionsView) GetOptionsOk() ([]string, bool) {
	if o == nil {
		return nil, false
	}
	return o.Options, true
}

// SetOptions sets field value
func (o *ProjectMetadataInteractionsView) SetOptions(v []string) {
	o.Options = v
}

// GetUri returns the Uri field value
func (o *ProjectMetadataInteractionsView) GetUri() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Uri
}

// GetUriOk returns a tuple with the Uri field value
// and a boolean to check if the value has been set.
func (o *ProjectMetadataInteractionsView) GetUriOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Uri, true
}

// SetUri sets field value
func (o *ProjectMetadataInteractionsView) SetUri(v string) {
	o.Uri = v
}

func (o ProjectMetadataInteractionsView) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o ProjectMetadataInteractionsView) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["options"] = o.Options
	toSerialize["uri"] = o.Uri
	return toSerialize, nil
}

type NullableProjectMetadataInteractionsView struct {
	value *ProjectMetadataInteractionsView
	isSet bool
}

func (v NullableProjectMetadataInteractionsView) Get() *ProjectMetadataInteractionsView {
	return v.value
}

func (v *NullableProjectMetadataInteractionsView) Set(val *ProjectMetadataInteractionsView) {
	v.value = val
	v.isSet = true
}

func (v NullableProjectMetadataInteractionsView) IsSet() bool {
	return v.isSet
}

func (v *NullableProjectMetadataInteractionsView) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableProjectMetadataInteractionsView(val *ProjectMetadataInteractionsView) *NullableProjectMetadataInteractionsView {
	return &NullableProjectMetadataInteractionsView{value: val, isSet: true}
}

func (v NullableProjectMetadataInteractionsView) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableProjectMetadataInteractionsView) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
