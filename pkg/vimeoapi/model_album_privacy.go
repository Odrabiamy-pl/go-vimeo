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

// checks if the AlbumPrivacy type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &AlbumPrivacy{}

// AlbumPrivacy The privacy settings of the showcase.
type AlbumPrivacy struct {
	// The showcase's password. This field appears only when **privacy.view** is `password`.
	Password *string `json:"password,omitempty"`
	// The access level of the showcase.  Option descriptions:  * `anybody` - Anyone can access the showcase. This privacy setting appears as `Public` on the Vimeo front end.  * `embed_only` - The showcase doesn't appear on Vimeo, but the owner can embed it on other sites.  * `nobody` - No one can access the showacse, including the owner. This privacy setting appears as `Private` on the Vimeo front end.  * `password` - Only those with the password can access the showcase.  * `team` - Only the owner and members of the owner's team can access the showcase.  * `unlisted` - The showcase can't be accessed if the URL omits its unlisted hash. 
	View string `json:"view"`
}

// NewAlbumPrivacy instantiates a new AlbumPrivacy object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewAlbumPrivacy(view string) *AlbumPrivacy {
	this := AlbumPrivacy{}
	this.View = view
	return &this
}

// NewAlbumPrivacyWithDefaults instantiates a new AlbumPrivacy object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewAlbumPrivacyWithDefaults() *AlbumPrivacy {
	this := AlbumPrivacy{}
	return &this
}

// GetPassword returns the Password field value if set, zero value otherwise.
func (o *AlbumPrivacy) GetPassword() string {
	if o == nil || IsNil(o.Password) {
		var ret string
		return ret
	}
	return *o.Password
}

// GetPasswordOk returns a tuple with the Password field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AlbumPrivacy) GetPasswordOk() (*string, bool) {
	if o == nil || IsNil(o.Password) {
		return nil, false
	}
	return o.Password, true
}

// HasPassword returns a boolean if a field has been set.
func (o *AlbumPrivacy) HasPassword() bool {
	if o != nil && !IsNil(o.Password) {
		return true
	}

	return false
}

// SetPassword gets a reference to the given string and assigns it to the Password field.
func (o *AlbumPrivacy) SetPassword(v string) {
	o.Password = &v
}

// GetView returns the View field value
func (o *AlbumPrivacy) GetView() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.View
}

// GetViewOk returns a tuple with the View field value
// and a boolean to check if the value has been set.
func (o *AlbumPrivacy) GetViewOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.View, true
}

// SetView sets field value
func (o *AlbumPrivacy) SetView(v string) {
	o.View = v
}

func (o AlbumPrivacy) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o AlbumPrivacy) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Password) {
		toSerialize["password"] = o.Password
	}
	toSerialize["view"] = o.View
	return toSerialize, nil
}

type NullableAlbumPrivacy struct {
	value *AlbumPrivacy
	isSet bool
}

func (v NullableAlbumPrivacy) Get() *AlbumPrivacy {
	return v.value
}

func (v *NullableAlbumPrivacy) Set(val *AlbumPrivacy) {
	v.value = val
	v.isSet = true
}

func (v NullableAlbumPrivacy) IsSet() bool {
	return v.isSet
}

func (v *NullableAlbumPrivacy) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableAlbumPrivacy(val *AlbumPrivacy) *NullableAlbumPrivacy {
	return &NullableAlbumPrivacy{value: val, isSet: true}
}

func (v NullableAlbumPrivacy) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableAlbumPrivacy) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


