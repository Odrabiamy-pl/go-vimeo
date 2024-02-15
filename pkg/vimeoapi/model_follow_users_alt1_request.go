/*
Vimeo API

Build something great. Vimeo's API supports flexible, high-quality video integration with your custom apps.

API version: 3.4
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package vimeoapi

import (
	"encoding/json"
	"bytes"
	"fmt"
)

// checks if the FollowUsersAlt1Request type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &FollowUsersAlt1Request{}

// FollowUsersAlt1Request struct for FollowUsersAlt1Request
type FollowUsersAlt1Request struct {
	// An array of user IDs for the authenticated user to follow.
	Users []string `json:"users"`
}

type _FollowUsersAlt1Request FollowUsersAlt1Request

// NewFollowUsersAlt1Request instantiates a new FollowUsersAlt1Request object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewFollowUsersAlt1Request(users []string) *FollowUsersAlt1Request {
	this := FollowUsersAlt1Request{}
	this.Users = users
	return &this
}

// NewFollowUsersAlt1RequestWithDefaults instantiates a new FollowUsersAlt1Request object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewFollowUsersAlt1RequestWithDefaults() *FollowUsersAlt1Request {
	this := FollowUsersAlt1Request{}
	return &this
}

// GetUsers returns the Users field value
func (o *FollowUsersAlt1Request) GetUsers() []string {
	if o == nil {
		var ret []string
		return ret
	}

	return o.Users
}

// GetUsersOk returns a tuple with the Users field value
// and a boolean to check if the value has been set.
func (o *FollowUsersAlt1Request) GetUsersOk() ([]string, bool) {
	if o == nil {
		return nil, false
	}
	return o.Users, true
}

// SetUsers sets field value
func (o *FollowUsersAlt1Request) SetUsers(v []string) {
	o.Users = v
}

func (o FollowUsersAlt1Request) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o FollowUsersAlt1Request) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["users"] = o.Users
	return toSerialize, nil
}

func (o *FollowUsersAlt1Request) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"users",
	}

	allProperties := make(map[string]interface{})

	err = json.Unmarshal(data, &allProperties)

	if err != nil {
		return err;
	}

	for _, requiredProperty := range(requiredProperties) {
		if _, exists := allProperties[requiredProperty]; !exists {
			return fmt.Errorf("no value given for required property %v", requiredProperty)
		}
	}

	varFollowUsersAlt1Request := _FollowUsersAlt1Request{}

	decoder := json.NewDecoder(bytes.NewReader(data))
	decoder.DisallowUnknownFields()
	err = decoder.Decode(&varFollowUsersAlt1Request)

	if err != nil {
		return err
	}

	*o = FollowUsersAlt1Request(varFollowUsersAlt1Request)

	return err
}

type NullableFollowUsersAlt1Request struct {
	value *FollowUsersAlt1Request
	isSet bool
}

func (v NullableFollowUsersAlt1Request) Get() *FollowUsersAlt1Request {
	return v.value
}

func (v *NullableFollowUsersAlt1Request) Set(val *FollowUsersAlt1Request) {
	v.value = val
	v.isSet = true
}

func (v NullableFollowUsersAlt1Request) IsSet() bool {
	return v.isSet
}

func (v *NullableFollowUsersAlt1Request) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableFollowUsersAlt1Request(val *FollowUsersAlt1Request) *NullableFollowUsersAlt1Request {
	return &NullableFollowUsersAlt1Request{value: val, isSet: true}
}

func (v NullableFollowUsersAlt1Request) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableFollowUsersAlt1Request) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

