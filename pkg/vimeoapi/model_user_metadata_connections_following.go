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

// checks if the UserMetadataConnectionsFollowing type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &UserMetadataConnectionsFollowing{}

// UserMetadataConnectionsFollowing Information about the users that the authenticated user is following.
type UserMetadataConnectionsFollowing struct {
	// An array of HTTP methods permitted on this URI.
	Options []string `json:"options"`
	// The total number of users on this connection.
	Total float32 `json:"total"`
	// The API URI that resolves to the connection data.
	Uri string `json:"uri"`
}

type _UserMetadataConnectionsFollowing UserMetadataConnectionsFollowing

// NewUserMetadataConnectionsFollowing instantiates a new UserMetadataConnectionsFollowing object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewUserMetadataConnectionsFollowing(options []string, total float32, uri string) *UserMetadataConnectionsFollowing {
	this := UserMetadataConnectionsFollowing{}
	this.Options = options
	this.Total = total
	this.Uri = uri
	return &this
}

// NewUserMetadataConnectionsFollowingWithDefaults instantiates a new UserMetadataConnectionsFollowing object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewUserMetadataConnectionsFollowingWithDefaults() *UserMetadataConnectionsFollowing {
	this := UserMetadataConnectionsFollowing{}
	return &this
}

// GetOptions returns the Options field value
func (o *UserMetadataConnectionsFollowing) GetOptions() []string {
	if o == nil {
		var ret []string
		return ret
	}

	return o.Options
}

// GetOptionsOk returns a tuple with the Options field value
// and a boolean to check if the value has been set.
func (o *UserMetadataConnectionsFollowing) GetOptionsOk() ([]string, bool) {
	if o == nil {
		return nil, false
	}
	return o.Options, true
}

// SetOptions sets field value
func (o *UserMetadataConnectionsFollowing) SetOptions(v []string) {
	o.Options = v
}

// GetTotal returns the Total field value
func (o *UserMetadataConnectionsFollowing) GetTotal() float32 {
	if o == nil {
		var ret float32
		return ret
	}

	return o.Total
}

// GetTotalOk returns a tuple with the Total field value
// and a boolean to check if the value has been set.
func (o *UserMetadataConnectionsFollowing) GetTotalOk() (*float32, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Total, true
}

// SetTotal sets field value
func (o *UserMetadataConnectionsFollowing) SetTotal(v float32) {
	o.Total = v
}

// GetUri returns the Uri field value
func (o *UserMetadataConnectionsFollowing) GetUri() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Uri
}

// GetUriOk returns a tuple with the Uri field value
// and a boolean to check if the value has been set.
func (o *UserMetadataConnectionsFollowing) GetUriOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Uri, true
}

// SetUri sets field value
func (o *UserMetadataConnectionsFollowing) SetUri(v string) {
	o.Uri = v
}

func (o UserMetadataConnectionsFollowing) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o UserMetadataConnectionsFollowing) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["options"] = o.Options
	toSerialize["total"] = o.Total
	toSerialize["uri"] = o.Uri
	return toSerialize, nil
}

func (o *UserMetadataConnectionsFollowing) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"options",
		"total",
		"uri",
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

	varUserMetadataConnectionsFollowing := _UserMetadataConnectionsFollowing{}

	decoder := json.NewDecoder(bytes.NewReader(data))
	decoder.DisallowUnknownFields()
	err = decoder.Decode(&varUserMetadataConnectionsFollowing)

	if err != nil {
		return err
	}

	*o = UserMetadataConnectionsFollowing(varUserMetadataConnectionsFollowing)

	return err
}

type NullableUserMetadataConnectionsFollowing struct {
	value *UserMetadataConnectionsFollowing
	isSet bool
}

func (v NullableUserMetadataConnectionsFollowing) Get() *UserMetadataConnectionsFollowing {
	return v.value
}

func (v *NullableUserMetadataConnectionsFollowing) Set(val *UserMetadataConnectionsFollowing) {
	v.value = val
	v.isSet = true
}

func (v NullableUserMetadataConnectionsFollowing) IsSet() bool {
	return v.isSet
}

func (v *NullableUserMetadataConnectionsFollowing) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableUserMetadataConnectionsFollowing(val *UserMetadataConnectionsFollowing) *NullableUserMetadataConnectionsFollowing {
	return &NullableUserMetadataConnectionsFollowing{value: val, isSet: true}
}

func (v NullableUserMetadataConnectionsFollowing) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableUserMetadataConnectionsFollowing) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

