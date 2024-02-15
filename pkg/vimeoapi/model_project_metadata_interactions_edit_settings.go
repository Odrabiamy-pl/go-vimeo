/*
Vimeo API

Build something great. Vimeo's API supports flexible, high-quality video integration with your custom apps.

API version: 3.4
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package vimeoapi

import (
	"encoding/json"
	"bytes"
	"fmt"
)

// checks if the ProjectMetadataInteractionsEditSettings type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &ProjectMetadataInteractionsEditSettings{}

// ProjectMetadataInteractionsEditSettings The user's edit settings permissions information for this project.
type ProjectMetadataInteractionsEditSettings struct {
	// An array of HTTP methods permitted on this URI.
	Options []string `json:"options"`
	// The API URI that resolves to the connection data.
	Uri string `json:"uri"`
}

type _ProjectMetadataInteractionsEditSettings ProjectMetadataInteractionsEditSettings

// NewProjectMetadataInteractionsEditSettings instantiates a new ProjectMetadataInteractionsEditSettings object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewProjectMetadataInteractionsEditSettings(options []string, uri string) *ProjectMetadataInteractionsEditSettings {
	this := ProjectMetadataInteractionsEditSettings{}
	this.Options = options
	this.Uri = uri
	return &this
}

// NewProjectMetadataInteractionsEditSettingsWithDefaults instantiates a new ProjectMetadataInteractionsEditSettings object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewProjectMetadataInteractionsEditSettingsWithDefaults() *ProjectMetadataInteractionsEditSettings {
	this := ProjectMetadataInteractionsEditSettings{}
	return &this
}

// GetOptions returns the Options field value
func (o *ProjectMetadataInteractionsEditSettings) GetOptions() []string {
	if o == nil {
		var ret []string
		return ret
	}

	return o.Options
}

// GetOptionsOk returns a tuple with the Options field value
// and a boolean to check if the value has been set.
func (o *ProjectMetadataInteractionsEditSettings) GetOptionsOk() ([]string, bool) {
	if o == nil {
		return nil, false
	}
	return o.Options, true
}

// SetOptions sets field value
func (o *ProjectMetadataInteractionsEditSettings) SetOptions(v []string) {
	o.Options = v
}

// GetUri returns the Uri field value
func (o *ProjectMetadataInteractionsEditSettings) GetUri() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Uri
}

// GetUriOk returns a tuple with the Uri field value
// and a boolean to check if the value has been set.
func (o *ProjectMetadataInteractionsEditSettings) GetUriOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Uri, true
}

// SetUri sets field value
func (o *ProjectMetadataInteractionsEditSettings) SetUri(v string) {
	o.Uri = v
}

func (o ProjectMetadataInteractionsEditSettings) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o ProjectMetadataInteractionsEditSettings) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["options"] = o.Options
	toSerialize["uri"] = o.Uri
	return toSerialize, nil
}

func (o *ProjectMetadataInteractionsEditSettings) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"options",
		"uri",
	}

	allProperties := make(map[string]interface{})

	err = json.Unmarshal(data, &allProperties)

	if err != nil {
		return err;
	}

	for _, requiredProperty := range(requiredProperties) {
		if _, exists := allProperties[requiredProperty]; !exists {
			return fmt.Errorf("no value given for required property %v", requiredProperty)
		}
	}

	varProjectMetadataInteractionsEditSettings := _ProjectMetadataInteractionsEditSettings{}

	decoder := json.NewDecoder(bytes.NewReader(data))
	decoder.DisallowUnknownFields()
	err = decoder.Decode(&varProjectMetadataInteractionsEditSettings)

	if err != nil {
		return err
	}

	*o = ProjectMetadataInteractionsEditSettings(varProjectMetadataInteractionsEditSettings)

	return err
}

type NullableProjectMetadataInteractionsEditSettings struct {
	value *ProjectMetadataInteractionsEditSettings
	isSet bool
}

func (v NullableProjectMetadataInteractionsEditSettings) Get() *ProjectMetadataInteractionsEditSettings {
	return v.value
}

func (v *NullableProjectMetadataInteractionsEditSettings) Set(val *ProjectMetadataInteractionsEditSettings) {
	v.value = val
	v.isSet = true
}

func (v NullableProjectMetadataInteractionsEditSettings) IsSet() bool {
	return v.isSet
}

func (v *NullableProjectMetadataInteractionsEditSettings) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableProjectMetadataInteractionsEditSettings(val *ProjectMetadataInteractionsEditSettings) *NullableProjectMetadataInteractionsEditSettings {
	return &NullableProjectMetadataInteractionsEditSettings{value: val, isSet: true}
}

func (v NullableProjectMetadataInteractionsEditSettings) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableProjectMetadataInteractionsEditSettings) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

