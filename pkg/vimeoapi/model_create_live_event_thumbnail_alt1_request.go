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

// checks if the CreateLiveEventThumbnailAlt1Request type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &CreateLiveEventThumbnailAlt1Request{}

// CreateLiveEventThumbnailAlt1Request struct for CreateLiveEventThumbnailAlt1Request
type CreateLiveEventThumbnailAlt1Request struct {
	// Whether the thumbnail is the event's active thumbnail.
	Active *bool `json:"active,omitempty"`
}

// NewCreateLiveEventThumbnailAlt1Request instantiates a new CreateLiveEventThumbnailAlt1Request object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewCreateLiveEventThumbnailAlt1Request() *CreateLiveEventThumbnailAlt1Request {
	this := CreateLiveEventThumbnailAlt1Request{}
	return &this
}

// NewCreateLiveEventThumbnailAlt1RequestWithDefaults instantiates a new CreateLiveEventThumbnailAlt1Request object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewCreateLiveEventThumbnailAlt1RequestWithDefaults() *CreateLiveEventThumbnailAlt1Request {
	this := CreateLiveEventThumbnailAlt1Request{}
	return &this
}

// GetActive returns the Active field value if set, zero value otherwise.
func (o *CreateLiveEventThumbnailAlt1Request) GetActive() bool {
	if o == nil || IsNil(o.Active) {
		var ret bool
		return ret
	}
	return *o.Active
}

// GetActiveOk returns a tuple with the Active field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CreateLiveEventThumbnailAlt1Request) GetActiveOk() (*bool, bool) {
	if o == nil || IsNil(o.Active) {
		return nil, false
	}
	return o.Active, true
}

// HasActive returns a boolean if a field has been set.
func (o *CreateLiveEventThumbnailAlt1Request) HasActive() bool {
	if o != nil && !IsNil(o.Active) {
		return true
	}

	return false
}

// SetActive gets a reference to the given bool and assigns it to the Active field.
func (o *CreateLiveEventThumbnailAlt1Request) SetActive(v bool) {
	o.Active = &v
}

func (o CreateLiveEventThumbnailAlt1Request) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o CreateLiveEventThumbnailAlt1Request) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Active) {
		toSerialize["active"] = o.Active
	}
	return toSerialize, nil
}

type NullableCreateLiveEventThumbnailAlt1Request struct {
	value *CreateLiveEventThumbnailAlt1Request
	isSet bool
}

func (v NullableCreateLiveEventThumbnailAlt1Request) Get() *CreateLiveEventThumbnailAlt1Request {
	return v.value
}

func (v *NullableCreateLiveEventThumbnailAlt1Request) Set(val *CreateLiveEventThumbnailAlt1Request) {
	v.value = val
	v.isSet = true
}

func (v NullableCreateLiveEventThumbnailAlt1Request) IsSet() bool {
	return v.isSet
}

func (v *NullableCreateLiveEventThumbnailAlt1Request) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableCreateLiveEventThumbnailAlt1Request(val *CreateLiveEventThumbnailAlt1Request) *NullableCreateLiveEventThumbnailAlt1Request {
	return &NullableCreateLiveEventThumbnailAlt1Request{value: val, isSet: true}
}

func (v NullableCreateLiveEventThumbnailAlt1Request) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableCreateLiveEventThumbnailAlt1Request) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

