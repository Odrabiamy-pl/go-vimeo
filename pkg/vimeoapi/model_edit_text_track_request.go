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

// checks if the EditTextTrackRequest type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &EditTextTrackRequest{}

// EditTextTrackRequest struct for EditTextTrackRequest
type EditTextTrackRequest struct {
	// Whether the current text track is the *active text track,* or the one that appears in the player. Only one text track per language and per type can be active.
	Active *bool `json:"active,omitempty"`
	// The language of the text track. For a full list of supported languages, use the [`/languages?filter=texttracks`](https://developer.vimeo.com/api/reference/videos#get_languages) endpoint.
	Language *string `json:"language,omitempty"`
	// The name of the text track.
	Name *string `json:"name,omitempty"`
	// The type of text track.  Option descriptions:  * `captions` - The text track is the captions type.  * `chapters` - The text track is the chapters type.  * `descriptions` - The text track is the descriptions type.  * `metadata` - The text track is the metadata type.  * `subtitles` - The text track is the subtitles type.
	Type *string `json:"type,omitempty"`
}

// NewEditTextTrackRequest instantiates a new EditTextTrackRequest object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewEditTextTrackRequest() *EditTextTrackRequest {
	this := EditTextTrackRequest{}
	return &this
}

// NewEditTextTrackRequestWithDefaults instantiates a new EditTextTrackRequest object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewEditTextTrackRequestWithDefaults() *EditTextTrackRequest {
	this := EditTextTrackRequest{}
	return &this
}

// GetActive returns the Active field value if set, zero value otherwise.
func (o *EditTextTrackRequest) GetActive() bool {
	if o == nil || IsNil(o.Active) {
		var ret bool
		return ret
	}
	return *o.Active
}

// GetActiveOk returns a tuple with the Active field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *EditTextTrackRequest) GetActiveOk() (*bool, bool) {
	if o == nil || IsNil(o.Active) {
		return nil, false
	}
	return o.Active, true
}

// HasActive returns a boolean if a field has been set.
func (o *EditTextTrackRequest) HasActive() bool {
	if o != nil && !IsNil(o.Active) {
		return true
	}

	return false
}

// SetActive gets a reference to the given bool and assigns it to the Active field.
func (o *EditTextTrackRequest) SetActive(v bool) {
	o.Active = &v
}

// GetLanguage returns the Language field value if set, zero value otherwise.
func (o *EditTextTrackRequest) GetLanguage() string {
	if o == nil || IsNil(o.Language) {
		var ret string
		return ret
	}
	return *o.Language
}

// GetLanguageOk returns a tuple with the Language field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *EditTextTrackRequest) GetLanguageOk() (*string, bool) {
	if o == nil || IsNil(o.Language) {
		return nil, false
	}
	return o.Language, true
}

// HasLanguage returns a boolean if a field has been set.
func (o *EditTextTrackRequest) HasLanguage() bool {
	if o != nil && !IsNil(o.Language) {
		return true
	}

	return false
}

// SetLanguage gets a reference to the given string and assigns it to the Language field.
func (o *EditTextTrackRequest) SetLanguage(v string) {
	o.Language = &v
}

// GetName returns the Name field value if set, zero value otherwise.
func (o *EditTextTrackRequest) GetName() string {
	if o == nil || IsNil(o.Name) {
		var ret string
		return ret
	}
	return *o.Name
}

// GetNameOk returns a tuple with the Name field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *EditTextTrackRequest) GetNameOk() (*string, bool) {
	if o == nil || IsNil(o.Name) {
		return nil, false
	}
	return o.Name, true
}

// HasName returns a boolean if a field has been set.
func (o *EditTextTrackRequest) HasName() bool {
	if o != nil && !IsNil(o.Name) {
		return true
	}

	return false
}

// SetName gets a reference to the given string and assigns it to the Name field.
func (o *EditTextTrackRequest) SetName(v string) {
	o.Name = &v
}

// GetType returns the Type field value if set, zero value otherwise.
func (o *EditTextTrackRequest) GetType() string {
	if o == nil || IsNil(o.Type) {
		var ret string
		return ret
	}
	return *o.Type
}

// GetTypeOk returns a tuple with the Type field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *EditTextTrackRequest) GetTypeOk() (*string, bool) {
	if o == nil || IsNil(o.Type) {
		return nil, false
	}
	return o.Type, true
}

// HasType returns a boolean if a field has been set.
func (o *EditTextTrackRequest) HasType() bool {
	if o != nil && !IsNil(o.Type) {
		return true
	}

	return false
}

// SetType gets a reference to the given string and assigns it to the Type field.
func (o *EditTextTrackRequest) SetType(v string) {
	o.Type = &v
}

func (o EditTextTrackRequest) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o EditTextTrackRequest) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Active) {
		toSerialize["active"] = o.Active
	}
	if !IsNil(o.Language) {
		toSerialize["language"] = o.Language
	}
	if !IsNil(o.Name) {
		toSerialize["name"] = o.Name
	}
	if !IsNil(o.Type) {
		toSerialize["type"] = o.Type
	}
	return toSerialize, nil
}

type NullableEditTextTrackRequest struct {
	value *EditTextTrackRequest
	isSet bool
}

func (v NullableEditTextTrackRequest) Get() *EditTextTrackRequest {
	return v.value
}

func (v *NullableEditTextTrackRequest) Set(val *EditTextTrackRequest) {
	v.value = val
	v.isSet = true
}

func (v NullableEditTextTrackRequest) IsSet() bool {
	return v.isSet
}

func (v *NullableEditTextTrackRequest) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableEditTextTrackRequest(val *EditTextTrackRequest) *NullableEditTextTrackRequest {
	return &NullableEditTextTrackRequest{value: val, isSet: true}
}

func (v NullableEditTextTrackRequest) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableEditTextTrackRequest) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
