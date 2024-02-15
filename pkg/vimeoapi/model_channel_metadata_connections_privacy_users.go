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

// checks if the ChannelMetadataConnectionsPrivacyUsers type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &ChannelMetadataConnectionsPrivacyUsers{}

// ChannelMetadataConnectionsPrivacyUsers Information provided to channel moderators about which users they have specifically permitted to access this private channel. This data requires a bearer token with the `private` scope.
type ChannelMetadataConnectionsPrivacyUsers struct {
	// An array of HTTP methods permitted on this URI. This data requires a bearer token with the `private` scope.
	Options []string `json:"options"`
	// The total number of users on this connection. This data requires a bearer token with the `private` scope.
	Total float32 `json:"total"`
	// The API URI that resolves to the connection data. This data requires a bearer token with the `private` scope.
	Uri string `json:"uri"`
}

type _ChannelMetadataConnectionsPrivacyUsers ChannelMetadataConnectionsPrivacyUsers

// NewChannelMetadataConnectionsPrivacyUsers instantiates a new ChannelMetadataConnectionsPrivacyUsers object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewChannelMetadataConnectionsPrivacyUsers(options []string, total float32, uri string) *ChannelMetadataConnectionsPrivacyUsers {
	this := ChannelMetadataConnectionsPrivacyUsers{}
	this.Options = options
	this.Total = total
	this.Uri = uri
	return &this
}

// NewChannelMetadataConnectionsPrivacyUsersWithDefaults instantiates a new ChannelMetadataConnectionsPrivacyUsers object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewChannelMetadataConnectionsPrivacyUsersWithDefaults() *ChannelMetadataConnectionsPrivacyUsers {
	this := ChannelMetadataConnectionsPrivacyUsers{}
	return &this
}

// GetOptions returns the Options field value
func (o *ChannelMetadataConnectionsPrivacyUsers) GetOptions() []string {
	if o == nil {
		var ret []string
		return ret
	}

	return o.Options
}

// GetOptionsOk returns a tuple with the Options field value
// and a boolean to check if the value has been set.
func (o *ChannelMetadataConnectionsPrivacyUsers) GetOptionsOk() ([]string, bool) {
	if o == nil {
		return nil, false
	}
	return o.Options, true
}

// SetOptions sets field value
func (o *ChannelMetadataConnectionsPrivacyUsers) SetOptions(v []string) {
	o.Options = v
}

// GetTotal returns the Total field value
func (o *ChannelMetadataConnectionsPrivacyUsers) GetTotal() float32 {
	if o == nil {
		var ret float32
		return ret
	}

	return o.Total
}

// GetTotalOk returns a tuple with the Total field value
// and a boolean to check if the value has been set.
func (o *ChannelMetadataConnectionsPrivacyUsers) GetTotalOk() (*float32, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Total, true
}

// SetTotal sets field value
func (o *ChannelMetadataConnectionsPrivacyUsers) SetTotal(v float32) {
	o.Total = v
}

// GetUri returns the Uri field value
func (o *ChannelMetadataConnectionsPrivacyUsers) GetUri() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Uri
}

// GetUriOk returns a tuple with the Uri field value
// and a boolean to check if the value has been set.
func (o *ChannelMetadataConnectionsPrivacyUsers) GetUriOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Uri, true
}

// SetUri sets field value
func (o *ChannelMetadataConnectionsPrivacyUsers) SetUri(v string) {
	o.Uri = v
}

func (o ChannelMetadataConnectionsPrivacyUsers) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o ChannelMetadataConnectionsPrivacyUsers) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["options"] = o.Options
	toSerialize["total"] = o.Total
	toSerialize["uri"] = o.Uri
	return toSerialize, nil
}

func (o *ChannelMetadataConnectionsPrivacyUsers) UnmarshalJSON(data []byte) (err error) {
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

	varChannelMetadataConnectionsPrivacyUsers := _ChannelMetadataConnectionsPrivacyUsers{}

	decoder := json.NewDecoder(bytes.NewReader(data))
	decoder.DisallowUnknownFields()
	err = decoder.Decode(&varChannelMetadataConnectionsPrivacyUsers)

	if err != nil {
		return err
	}

	*o = ChannelMetadataConnectionsPrivacyUsers(varChannelMetadataConnectionsPrivacyUsers)

	return err
}

type NullableChannelMetadataConnectionsPrivacyUsers struct {
	value *ChannelMetadataConnectionsPrivacyUsers
	isSet bool
}

func (v NullableChannelMetadataConnectionsPrivacyUsers) Get() *ChannelMetadataConnectionsPrivacyUsers {
	return v.value
}

func (v *NullableChannelMetadataConnectionsPrivacyUsers) Set(val *ChannelMetadataConnectionsPrivacyUsers) {
	v.value = val
	v.isSet = true
}

func (v NullableChannelMetadataConnectionsPrivacyUsers) IsSet() bool {
	return v.isSet
}

func (v *NullableChannelMetadataConnectionsPrivacyUsers) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableChannelMetadataConnectionsPrivacyUsers(val *ChannelMetadataConnectionsPrivacyUsers) *NullableChannelMetadataConnectionsPrivacyUsers {
	return &NullableChannelMetadataConnectionsPrivacyUsers{value: val, isSet: true}
}

func (v NullableChannelMetadataConnectionsPrivacyUsers) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableChannelMetadataConnectionsPrivacyUsers) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

