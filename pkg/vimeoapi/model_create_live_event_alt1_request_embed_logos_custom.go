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

// checks if the CreateLiveEventAlt1RequestEmbedLogosCustom type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &CreateLiveEventAlt1RequestEmbedLogosCustom{}

// CreateLiveEventAlt1RequestEmbedLogosCustom struct for CreateLiveEventAlt1RequestEmbedLogosCustom
type CreateLiveEventAlt1RequestEmbedLogosCustom struct {
	// Whether to show the custom logo on the embed player.
	Active *bool `json:"active,omitempty"`
	// The URL that loads when the user clicks the custom logo.
	Link *string `json:"link,omitempty"`
	// Whether to show the custom logo persistently (`true`) or hide it with the rest of the UI (`false`).
	Sticky *bool `json:"sticky,omitempty"`
}

// NewCreateLiveEventAlt1RequestEmbedLogosCustom instantiates a new CreateLiveEventAlt1RequestEmbedLogosCustom object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewCreateLiveEventAlt1RequestEmbedLogosCustom() *CreateLiveEventAlt1RequestEmbedLogosCustom {
	this := CreateLiveEventAlt1RequestEmbedLogosCustom{}
	return &this
}

// NewCreateLiveEventAlt1RequestEmbedLogosCustomWithDefaults instantiates a new CreateLiveEventAlt1RequestEmbedLogosCustom object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewCreateLiveEventAlt1RequestEmbedLogosCustomWithDefaults() *CreateLiveEventAlt1RequestEmbedLogosCustom {
	this := CreateLiveEventAlt1RequestEmbedLogosCustom{}
	return &this
}

// GetActive returns the Active field value if set, zero value otherwise.
func (o *CreateLiveEventAlt1RequestEmbedLogosCustom) GetActive() bool {
	if o == nil || IsNil(o.Active) {
		var ret bool
		return ret
	}
	return *o.Active
}

// GetActiveOk returns a tuple with the Active field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CreateLiveEventAlt1RequestEmbedLogosCustom) GetActiveOk() (*bool, bool) {
	if o == nil || IsNil(o.Active) {
		return nil, false
	}
	return o.Active, true
}

// HasActive returns a boolean if a field has been set.
func (o *CreateLiveEventAlt1RequestEmbedLogosCustom) HasActive() bool {
	if o != nil && !IsNil(o.Active) {
		return true
	}

	return false
}

// SetActive gets a reference to the given bool and assigns it to the Active field.
func (o *CreateLiveEventAlt1RequestEmbedLogosCustom) SetActive(v bool) {
	o.Active = &v
}

// GetLink returns the Link field value if set, zero value otherwise.
func (o *CreateLiveEventAlt1RequestEmbedLogosCustom) GetLink() string {
	if o == nil || IsNil(o.Link) {
		var ret string
		return ret
	}
	return *o.Link
}

// GetLinkOk returns a tuple with the Link field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CreateLiveEventAlt1RequestEmbedLogosCustom) GetLinkOk() (*string, bool) {
	if o == nil || IsNil(o.Link) {
		return nil, false
	}
	return o.Link, true
}

// HasLink returns a boolean if a field has been set.
func (o *CreateLiveEventAlt1RequestEmbedLogosCustom) HasLink() bool {
	if o != nil && !IsNil(o.Link) {
		return true
	}

	return false
}

// SetLink gets a reference to the given string and assigns it to the Link field.
func (o *CreateLiveEventAlt1RequestEmbedLogosCustom) SetLink(v string) {
	o.Link = &v
}

// GetSticky returns the Sticky field value if set, zero value otherwise.
func (o *CreateLiveEventAlt1RequestEmbedLogosCustom) GetSticky() bool {
	if o == nil || IsNil(o.Sticky) {
		var ret bool
		return ret
	}
	return *o.Sticky
}

// GetStickyOk returns a tuple with the Sticky field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CreateLiveEventAlt1RequestEmbedLogosCustom) GetStickyOk() (*bool, bool) {
	if o == nil || IsNil(o.Sticky) {
		return nil, false
	}
	return o.Sticky, true
}

// HasSticky returns a boolean if a field has been set.
func (o *CreateLiveEventAlt1RequestEmbedLogosCustom) HasSticky() bool {
	if o != nil && !IsNil(o.Sticky) {
		return true
	}

	return false
}

// SetSticky gets a reference to the given bool and assigns it to the Sticky field.
func (o *CreateLiveEventAlt1RequestEmbedLogosCustom) SetSticky(v bool) {
	o.Sticky = &v
}

func (o CreateLiveEventAlt1RequestEmbedLogosCustom) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o CreateLiveEventAlt1RequestEmbedLogosCustom) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Active) {
		toSerialize["active"] = o.Active
	}
	if !IsNil(o.Link) {
		toSerialize["link"] = o.Link
	}
	if !IsNil(o.Sticky) {
		toSerialize["sticky"] = o.Sticky
	}
	return toSerialize, nil
}

type NullableCreateLiveEventAlt1RequestEmbedLogosCustom struct {
	value *CreateLiveEventAlt1RequestEmbedLogosCustom
	isSet bool
}

func (v NullableCreateLiveEventAlt1RequestEmbedLogosCustom) Get() *CreateLiveEventAlt1RequestEmbedLogosCustom {
	return v.value
}

func (v *NullableCreateLiveEventAlt1RequestEmbedLogosCustom) Set(val *CreateLiveEventAlt1RequestEmbedLogosCustom) {
	v.value = val
	v.isSet = true
}

func (v NullableCreateLiveEventAlt1RequestEmbedLogosCustom) IsSet() bool {
	return v.isSet
}

func (v *NullableCreateLiveEventAlt1RequestEmbedLogosCustom) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableCreateLiveEventAlt1RequestEmbedLogosCustom(val *CreateLiveEventAlt1RequestEmbedLogosCustom) *NullableCreateLiveEventAlt1RequestEmbedLogosCustom {
	return &NullableCreateLiveEventAlt1RequestEmbedLogosCustom{value: val, isSet: true}
}

func (v NullableCreateLiveEventAlt1RequestEmbedLogosCustom) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableCreateLiveEventAlt1RequestEmbedLogosCustom) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
