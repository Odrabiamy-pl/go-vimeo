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

// checks if the Fragments type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &Fragments{}

// Fragments struct for Fragments
type Fragments struct {
	// The time in ISO 8601 format when the fragment was created.
	CreatedOn string `json:"created_on"`
	Metadata FragmentsMetadata `json:"metadata"`
	// The time in ISO 8601 format when the fragment was last updated.
	ModifiedOn string `json:"modified_on"`
	// The time in milliseconds of the fragment's _inpoint_, or the time from the start of the video that marks the beginning of the fragment.
	Timecode float32 `json:"timecode"`
	// The URI of the video fragment.
	Uri string `json:"uri"`
}

// NewFragments instantiates a new Fragments object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewFragments(createdOn string, metadata FragmentsMetadata, modifiedOn string, timecode float32, uri string) *Fragments {
	this := Fragments{}
	this.CreatedOn = createdOn
	this.Metadata = metadata
	this.ModifiedOn = modifiedOn
	this.Timecode = timecode
	this.Uri = uri
	return &this
}

// NewFragmentsWithDefaults instantiates a new Fragments object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewFragmentsWithDefaults() *Fragments {
	this := Fragments{}
	return &this
}

// GetCreatedOn returns the CreatedOn field value
func (o *Fragments) GetCreatedOn() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.CreatedOn
}

// GetCreatedOnOk returns a tuple with the CreatedOn field value
// and a boolean to check if the value has been set.
func (o *Fragments) GetCreatedOnOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.CreatedOn, true
}

// SetCreatedOn sets field value
func (o *Fragments) SetCreatedOn(v string) {
	o.CreatedOn = v
}

// GetMetadata returns the Metadata field value
func (o *Fragments) GetMetadata() FragmentsMetadata {
	if o == nil {
		var ret FragmentsMetadata
		return ret
	}

	return o.Metadata
}

// GetMetadataOk returns a tuple with the Metadata field value
// and a boolean to check if the value has been set.
func (o *Fragments) GetMetadataOk() (*FragmentsMetadata, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Metadata, true
}

// SetMetadata sets field value
func (o *Fragments) SetMetadata(v FragmentsMetadata) {
	o.Metadata = v
}

// GetModifiedOn returns the ModifiedOn field value
func (o *Fragments) GetModifiedOn() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.ModifiedOn
}

// GetModifiedOnOk returns a tuple with the ModifiedOn field value
// and a boolean to check if the value has been set.
func (o *Fragments) GetModifiedOnOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ModifiedOn, true
}

// SetModifiedOn sets field value
func (o *Fragments) SetModifiedOn(v string) {
	o.ModifiedOn = v
}

// GetTimecode returns the Timecode field value
func (o *Fragments) GetTimecode() float32 {
	if o == nil {
		var ret float32
		return ret
	}

	return o.Timecode
}

// GetTimecodeOk returns a tuple with the Timecode field value
// and a boolean to check if the value has been set.
func (o *Fragments) GetTimecodeOk() (*float32, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Timecode, true
}

// SetTimecode sets field value
func (o *Fragments) SetTimecode(v float32) {
	o.Timecode = v
}

// GetUri returns the Uri field value
func (o *Fragments) GetUri() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Uri
}

// GetUriOk returns a tuple with the Uri field value
// and a boolean to check if the value has been set.
func (o *Fragments) GetUriOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Uri, true
}

// SetUri sets field value
func (o *Fragments) SetUri(v string) {
	o.Uri = v
}

func (o Fragments) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o Fragments) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["created_on"] = o.CreatedOn
	toSerialize["metadata"] = o.Metadata
	toSerialize["modified_on"] = o.ModifiedOn
	toSerialize["timecode"] = o.Timecode
	toSerialize["uri"] = o.Uri
	return toSerialize, nil
}

type NullableFragments struct {
	value *Fragments
	isSet bool
}

func (v NullableFragments) Get() *Fragments {
	return v.value
}

func (v *NullableFragments) Set(val *Fragments) {
	v.value = val
	v.isSet = true
}

func (v NullableFragments) IsSet() bool {
	return v.isSet
}

func (v *NullableFragments) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableFragments(val *Fragments) *NullableFragments {
	return &NullableFragments{value: val, isSet: true}
}

func (v NullableFragments) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableFragments) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


