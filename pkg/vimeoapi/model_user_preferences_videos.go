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

// checks if the UserPreferencesVideos type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &UserPreferencesVideos{}

// UserPreferencesVideos struct for UserPreferencesVideos
type UserPreferencesVideos struct {
	// The password for viewing the authenticated user's videos.
	Password *string `json:"password,omitempty"`
	Privacy *UserPreferencesVideosPrivacy `json:"privacy,omitempty"`
	// An array of the authorized user's default content ratings.  Option descriptions:  * `drugs` - The video contains drug or alcohol use.  * `language` - The video contains profanity or sexually suggestive content.  * `nudity` - The video contains nudity.  * `safe` - The video is suitable for all audiences.  * `unrated` - The video hasn't been rated.  * `violence` - The video contains violent or graphic content. 
	Rating []string `json:"rating"`
}

// NewUserPreferencesVideos instantiates a new UserPreferencesVideos object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewUserPreferencesVideos(rating []string) *UserPreferencesVideos {
	this := UserPreferencesVideos{}
	this.Rating = rating
	return &this
}

// NewUserPreferencesVideosWithDefaults instantiates a new UserPreferencesVideos object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewUserPreferencesVideosWithDefaults() *UserPreferencesVideos {
	this := UserPreferencesVideos{}
	return &this
}

// GetPassword returns the Password field value if set, zero value otherwise.
func (o *UserPreferencesVideos) GetPassword() string {
	if o == nil || IsNil(o.Password) {
		var ret string
		return ret
	}
	return *o.Password
}

// GetPasswordOk returns a tuple with the Password field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UserPreferencesVideos) GetPasswordOk() (*string, bool) {
	if o == nil || IsNil(o.Password) {
		return nil, false
	}
	return o.Password, true
}

// HasPassword returns a boolean if a field has been set.
func (o *UserPreferencesVideos) HasPassword() bool {
	if o != nil && !IsNil(o.Password) {
		return true
	}

	return false
}

// SetPassword gets a reference to the given string and assigns it to the Password field.
func (o *UserPreferencesVideos) SetPassword(v string) {
	o.Password = &v
}

// GetPrivacy returns the Privacy field value if set, zero value otherwise.
func (o *UserPreferencesVideos) GetPrivacy() UserPreferencesVideosPrivacy {
	if o == nil || IsNil(o.Privacy) {
		var ret UserPreferencesVideosPrivacy
		return ret
	}
	return *o.Privacy
}

// GetPrivacyOk returns a tuple with the Privacy field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UserPreferencesVideos) GetPrivacyOk() (*UserPreferencesVideosPrivacy, bool) {
	if o == nil || IsNil(o.Privacy) {
		return nil, false
	}
	return o.Privacy, true
}

// HasPrivacy returns a boolean if a field has been set.
func (o *UserPreferencesVideos) HasPrivacy() bool {
	if o != nil && !IsNil(o.Privacy) {
		return true
	}

	return false
}

// SetPrivacy gets a reference to the given UserPreferencesVideosPrivacy and assigns it to the Privacy field.
func (o *UserPreferencesVideos) SetPrivacy(v UserPreferencesVideosPrivacy) {
	o.Privacy = &v
}

// GetRating returns the Rating field value
func (o *UserPreferencesVideos) GetRating() []string {
	if o == nil {
		var ret []string
		return ret
	}

	return o.Rating
}

// GetRatingOk returns a tuple with the Rating field value
// and a boolean to check if the value has been set.
func (o *UserPreferencesVideos) GetRatingOk() ([]string, bool) {
	if o == nil {
		return nil, false
	}
	return o.Rating, true
}

// SetRating sets field value
func (o *UserPreferencesVideos) SetRating(v []string) {
	o.Rating = v
}

func (o UserPreferencesVideos) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o UserPreferencesVideos) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Password) {
		toSerialize["password"] = o.Password
	}
	if !IsNil(o.Privacy) {
		toSerialize["privacy"] = o.Privacy
	}
	toSerialize["rating"] = o.Rating
	return toSerialize, nil
}

type NullableUserPreferencesVideos struct {
	value *UserPreferencesVideos
	isSet bool
}

func (v NullableUserPreferencesVideos) Get() *UserPreferencesVideos {
	return v.value
}

func (v *NullableUserPreferencesVideos) Set(val *UserPreferencesVideos) {
	v.value = val
	v.isSet = true
}

func (v NullableUserPreferencesVideos) IsSet() bool {
	return v.isSet
}

func (v *NullableUserPreferencesVideos) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableUserPreferencesVideos(val *UserPreferencesVideos) *NullableUserPreferencesVideos {
	return &NullableUserPreferencesVideos{value: val, isSet: true}
}

func (v NullableUserPreferencesVideos) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableUserPreferencesVideos) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


