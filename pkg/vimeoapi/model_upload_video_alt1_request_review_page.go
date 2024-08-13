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

// checks if the UploadVideoAlt1RequestReviewPage type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &UploadVideoAlt1RequestReviewPage{}

// UploadVideoAlt1RequestReviewPage struct for UploadVideoAlt1RequestReviewPage
type UploadVideoAlt1RequestReviewPage struct {
	// Whether to enable video review.
	Active *bool `json:"active,omitempty"`
}

// NewUploadVideoAlt1RequestReviewPage instantiates a new UploadVideoAlt1RequestReviewPage object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewUploadVideoAlt1RequestReviewPage() *UploadVideoAlt1RequestReviewPage {
	this := UploadVideoAlt1RequestReviewPage{}
	return &this
}

// NewUploadVideoAlt1RequestReviewPageWithDefaults instantiates a new UploadVideoAlt1RequestReviewPage object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewUploadVideoAlt1RequestReviewPageWithDefaults() *UploadVideoAlt1RequestReviewPage {
	this := UploadVideoAlt1RequestReviewPage{}
	return &this
}

// GetActive returns the Active field value if set, zero value otherwise.
func (o *UploadVideoAlt1RequestReviewPage) GetActive() bool {
	if o == nil || IsNil(o.Active) {
		var ret bool
		return ret
	}
	return *o.Active
}

// GetActiveOk returns a tuple with the Active field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UploadVideoAlt1RequestReviewPage) GetActiveOk() (*bool, bool) {
	if o == nil || IsNil(o.Active) {
		return nil, false
	}
	return o.Active, true
}

// HasActive returns a boolean if a field has been set.
func (o *UploadVideoAlt1RequestReviewPage) HasActive() bool {
	if o != nil && !IsNil(o.Active) {
		return true
	}

	return false
}

// SetActive gets a reference to the given bool and assigns it to the Active field.
func (o *UploadVideoAlt1RequestReviewPage) SetActive(v bool) {
	o.Active = &v
}

func (o UploadVideoAlt1RequestReviewPage) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o UploadVideoAlt1RequestReviewPage) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Active) {
		toSerialize["active"] = o.Active
	}
	return toSerialize, nil
}

type NullableUploadVideoAlt1RequestReviewPage struct {
	value *UploadVideoAlt1RequestReviewPage
	isSet bool
}

func (v NullableUploadVideoAlt1RequestReviewPage) Get() *UploadVideoAlt1RequestReviewPage {
	return v.value
}

func (v *NullableUploadVideoAlt1RequestReviewPage) Set(val *UploadVideoAlt1RequestReviewPage) {
	v.value = val
	v.isSet = true
}

func (v NullableUploadVideoAlt1RequestReviewPage) IsSet() bool {
	return v.isSet
}

func (v *NullableUploadVideoAlt1RequestReviewPage) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableUploadVideoAlt1RequestReviewPage(val *UploadVideoAlt1RequestReviewPage) *NullableUploadVideoAlt1RequestReviewPage {
	return &NullableUploadVideoAlt1RequestReviewPage{value: val, isSet: true}
}

func (v NullableUploadVideoAlt1RequestReviewPage) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableUploadVideoAlt1RequestReviewPage) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


