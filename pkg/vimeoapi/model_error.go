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

// checks if the ModelError type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &ModelError{}

// ModelError struct for ModelError
type ModelError struct {
	// The error message that technical users receive.
	DeveloperMessage string `json:"developer_message"`
	// The error message that general users receive.
	Error string `json:"error"`
	// The error code.
	ErrorCode float32 `json:"error_code"`
	// A link to more information about the error.
	Link string `json:"link"`
}

// NewModelError instantiates a new ModelError object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewModelError(developerMessage string, error_ string, errorCode float32, link string) *ModelError {
	this := ModelError{}
	this.DeveloperMessage = developerMessage
	this.Error = error_
	this.ErrorCode = errorCode
	this.Link = link
	return &this
}

// NewModelErrorWithDefaults instantiates a new ModelError object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewModelErrorWithDefaults() *ModelError {
	this := ModelError{}
	return &this
}

// GetDeveloperMessage returns the DeveloperMessage field value
func (o *ModelError) GetDeveloperMessage() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.DeveloperMessage
}

// GetDeveloperMessageOk returns a tuple with the DeveloperMessage field value
// and a boolean to check if the value has been set.
func (o *ModelError) GetDeveloperMessageOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.DeveloperMessage, true
}

// SetDeveloperMessage sets field value
func (o *ModelError) SetDeveloperMessage(v string) {
	o.DeveloperMessage = v
}

// GetError returns the Error field value
func (o *ModelError) GetError() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Error
}

// GetErrorOk returns a tuple with the Error field value
// and a boolean to check if the value has been set.
func (o *ModelError) GetErrorOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Error, true
}

// SetError sets field value
func (o *ModelError) SetError(v string) {
	o.Error = v
}

// GetErrorCode returns the ErrorCode field value
func (o *ModelError) GetErrorCode() float32 {
	if o == nil {
		var ret float32
		return ret
	}

	return o.ErrorCode
}

// GetErrorCodeOk returns a tuple with the ErrorCode field value
// and a boolean to check if the value has been set.
func (o *ModelError) GetErrorCodeOk() (*float32, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ErrorCode, true
}

// SetErrorCode sets field value
func (o *ModelError) SetErrorCode(v float32) {
	o.ErrorCode = v
}

// GetLink returns the Link field value
func (o *ModelError) GetLink() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Link
}

// GetLinkOk returns a tuple with the Link field value
// and a boolean to check if the value has been set.
func (o *ModelError) GetLinkOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Link, true
}

// SetLink sets field value
func (o *ModelError) SetLink(v string) {
	o.Link = v
}

func (o ModelError) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o ModelError) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["developer_message"] = o.DeveloperMessage
	toSerialize["error"] = o.Error
	toSerialize["error_code"] = o.ErrorCode
	toSerialize["link"] = o.Link
	return toSerialize, nil
}

type NullableModelError struct {
	value *ModelError
	isSet bool
}

func (v NullableModelError) Get() *ModelError {
	return v.value
}

func (v *NullableModelError) Set(val *ModelError) {
	v.value = val
	v.isSet = true
}

func (v NullableModelError) IsSet() bool {
	return v.isSet
}

func (v *NullableModelError) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableModelError(val *ModelError) *NullableModelError {
	return &NullableModelError{value: val, isSet: true}
}

func (v NullableModelError) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableModelError) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


