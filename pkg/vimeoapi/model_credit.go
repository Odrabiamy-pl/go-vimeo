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

// checks if the Credit type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &Credit{}

// Credit struct for Credit
type Credit struct {
	// The name of the person credited.
	Name string `json:"name"`
	// The character that the person portrayed, or the job that the person performed.
	Role string `json:"role"`
	// The unique identifier to access the credit resource.
	Uri string `json:"uri"`
	User *User `json:"user,omitempty"`
	Video *Video `json:"video,omitempty"`
}

// NewCredit instantiates a new Credit object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewCredit(name string, role string, uri string) *Credit {
	this := Credit{}
	this.Name = name
	this.Role = role
	this.Uri = uri
	return &this
}

// NewCreditWithDefaults instantiates a new Credit object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewCreditWithDefaults() *Credit {
	this := Credit{}
	return &this
}

// GetName returns the Name field value
func (o *Credit) GetName() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Name
}

// GetNameOk returns a tuple with the Name field value
// and a boolean to check if the value has been set.
func (o *Credit) GetNameOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Name, true
}

// SetName sets field value
func (o *Credit) SetName(v string) {
	o.Name = v
}

// GetRole returns the Role field value
func (o *Credit) GetRole() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Role
}

// GetRoleOk returns a tuple with the Role field value
// and a boolean to check if the value has been set.
func (o *Credit) GetRoleOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Role, true
}

// SetRole sets field value
func (o *Credit) SetRole(v string) {
	o.Role = v
}

// GetUri returns the Uri field value
func (o *Credit) GetUri() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Uri
}

// GetUriOk returns a tuple with the Uri field value
// and a boolean to check if the value has been set.
func (o *Credit) GetUriOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Uri, true
}

// SetUri sets field value
func (o *Credit) SetUri(v string) {
	o.Uri = v
}

// GetUser returns the User field value if set, zero value otherwise.
func (o *Credit) GetUser() User {
	if o == nil || IsNil(o.User) {
		var ret User
		return ret
	}
	return *o.User
}

// GetUserOk returns a tuple with the User field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Credit) GetUserOk() (*User, bool) {
	if o == nil || IsNil(o.User) {
		return nil, false
	}
	return o.User, true
}

// HasUser returns a boolean if a field has been set.
func (o *Credit) HasUser() bool {
	if o != nil && !IsNil(o.User) {
		return true
	}

	return false
}

// SetUser gets a reference to the given User and assigns it to the User field.
func (o *Credit) SetUser(v User) {
	o.User = &v
}

// GetVideo returns the Video field value if set, zero value otherwise.
func (o *Credit) GetVideo() Video {
	if o == nil || IsNil(o.Video) {
		var ret Video
		return ret
	}
	return *o.Video
}

// GetVideoOk returns a tuple with the Video field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Credit) GetVideoOk() (*Video, bool) {
	if o == nil || IsNil(o.Video) {
		return nil, false
	}
	return o.Video, true
}

// HasVideo returns a boolean if a field has been set.
func (o *Credit) HasVideo() bool {
	if o != nil && !IsNil(o.Video) {
		return true
	}

	return false
}

// SetVideo gets a reference to the given Video and assigns it to the Video field.
func (o *Credit) SetVideo(v Video) {
	o.Video = &v
}

func (o Credit) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o Credit) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["name"] = o.Name
	toSerialize["role"] = o.Role
	toSerialize["uri"] = o.Uri
	if !IsNil(o.User) {
		toSerialize["user"] = o.User
	}
	if !IsNil(o.Video) {
		toSerialize["video"] = o.Video
	}
	return toSerialize, nil
}

type NullableCredit struct {
	value *Credit
	isSet bool
}

func (v NullableCredit) Get() *Credit {
	return v.value
}

func (v *NullableCredit) Set(val *Credit) {
	v.value = val
	v.isSet = true
}

func (v NullableCredit) IsSet() bool {
	return v.isSet
}

func (v *NullableCredit) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableCredit(val *Credit) *NullableCredit {
	return &NullableCredit{value: val, isSet: true}
}

func (v NullableCredit) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableCredit) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


