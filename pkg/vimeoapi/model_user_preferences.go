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

// checks if the UserPreferences type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &UserPreferences{}

// UserPreferences struct for UserPreferences
type UserPreferences struct {
	Videos UserPreferencesVideos `json:"videos"`
}

// NewUserPreferences instantiates a new UserPreferences object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewUserPreferences(videos UserPreferencesVideos) *UserPreferences {
	this := UserPreferences{}
	this.Videos = videos
	return &this
}

// NewUserPreferencesWithDefaults instantiates a new UserPreferences object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewUserPreferencesWithDefaults() *UserPreferences {
	this := UserPreferences{}
	return &this
}

// GetVideos returns the Videos field value
func (o *UserPreferences) GetVideos() UserPreferencesVideos {
	if o == nil {
		var ret UserPreferencesVideos
		return ret
	}

	return o.Videos
}

// GetVideosOk returns a tuple with the Videos field value
// and a boolean to check if the value has been set.
func (o *UserPreferences) GetVideosOk() (*UserPreferencesVideos, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Videos, true
}

// SetVideos sets field value
func (o *UserPreferences) SetVideos(v UserPreferencesVideos) {
	o.Videos = v
}

func (o UserPreferences) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o UserPreferences) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["videos"] = o.Videos
	return toSerialize, nil
}

type NullableUserPreferences struct {
	value *UserPreferences
	isSet bool
}

func (v NullableUserPreferences) Get() *UserPreferences {
	return v.value
}

func (v *NullableUserPreferences) Set(val *UserPreferences) {
	v.value = val
	v.isSet = true
}

func (v NullableUserPreferences) IsSet() bool {
	return v.isSet
}

func (v *NullableUserPreferences) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableUserPreferences(val *UserPreferences) *NullableUserPreferences {
	return &NullableUserPreferences{value: val, isSet: true}
}

func (v NullableUserPreferences) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableUserPreferences) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


