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

// checks if the EmailCaptureFormCustomFieldsInnerFieldMetadata type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &EmailCaptureFormCustomFieldsInnerFieldMetadata{}

// EmailCaptureFormCustomFieldsInnerFieldMetadata The metadata object that is associated with the field in the form.
type EmailCaptureFormCustomFieldsInnerFieldMetadata struct {
	// An array of options associated with the field.
	Options []EmailCaptureFormCustomFieldsInnerFieldMetadataOptionsInner `json:"options"`
}

// NewEmailCaptureFormCustomFieldsInnerFieldMetadata instantiates a new EmailCaptureFormCustomFieldsInnerFieldMetadata object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewEmailCaptureFormCustomFieldsInnerFieldMetadata(options []EmailCaptureFormCustomFieldsInnerFieldMetadataOptionsInner) *EmailCaptureFormCustomFieldsInnerFieldMetadata {
	this := EmailCaptureFormCustomFieldsInnerFieldMetadata{}
	this.Options = options
	return &this
}

// NewEmailCaptureFormCustomFieldsInnerFieldMetadataWithDefaults instantiates a new EmailCaptureFormCustomFieldsInnerFieldMetadata object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewEmailCaptureFormCustomFieldsInnerFieldMetadataWithDefaults() *EmailCaptureFormCustomFieldsInnerFieldMetadata {
	this := EmailCaptureFormCustomFieldsInnerFieldMetadata{}
	return &this
}

// GetOptions returns the Options field value
func (o *EmailCaptureFormCustomFieldsInnerFieldMetadata) GetOptions() []EmailCaptureFormCustomFieldsInnerFieldMetadataOptionsInner {
	if o == nil {
		var ret []EmailCaptureFormCustomFieldsInnerFieldMetadataOptionsInner
		return ret
	}

	return o.Options
}

// GetOptionsOk returns a tuple with the Options field value
// and a boolean to check if the value has been set.
func (o *EmailCaptureFormCustomFieldsInnerFieldMetadata) GetOptionsOk() ([]EmailCaptureFormCustomFieldsInnerFieldMetadataOptionsInner, bool) {
	if o == nil {
		return nil, false
	}
	return o.Options, true
}

// SetOptions sets field value
func (o *EmailCaptureFormCustomFieldsInnerFieldMetadata) SetOptions(v []EmailCaptureFormCustomFieldsInnerFieldMetadataOptionsInner) {
	o.Options = v
}

func (o EmailCaptureFormCustomFieldsInnerFieldMetadata) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o EmailCaptureFormCustomFieldsInnerFieldMetadata) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["options"] = o.Options
	return toSerialize, nil
}

type NullableEmailCaptureFormCustomFieldsInnerFieldMetadata struct {
	value *EmailCaptureFormCustomFieldsInnerFieldMetadata
	isSet bool
}

func (v NullableEmailCaptureFormCustomFieldsInnerFieldMetadata) Get() *EmailCaptureFormCustomFieldsInnerFieldMetadata {
	return v.value
}

func (v *NullableEmailCaptureFormCustomFieldsInnerFieldMetadata) Set(val *EmailCaptureFormCustomFieldsInnerFieldMetadata) {
	v.value = val
	v.isSet = true
}

func (v NullableEmailCaptureFormCustomFieldsInnerFieldMetadata) IsSet() bool {
	return v.isSet
}

func (v *NullableEmailCaptureFormCustomFieldsInnerFieldMetadata) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableEmailCaptureFormCustomFieldsInnerFieldMetadata(val *EmailCaptureFormCustomFieldsInnerFieldMetadata) *NullableEmailCaptureFormCustomFieldsInnerFieldMetadata {
	return &NullableEmailCaptureFormCustomFieldsInnerFieldMetadata{value: val, isSet: true}
}

func (v NullableEmailCaptureFormCustomFieldsInnerFieldMetadata) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableEmailCaptureFormCustomFieldsInnerFieldMetadata) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
