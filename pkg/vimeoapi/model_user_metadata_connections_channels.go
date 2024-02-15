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

// checks if the UserMetadataConnectionsChannels type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &UserMetadataConnectionsChannels{}

// UserMetadataConnectionsChannels Information about the channels to which the authenticated user subscribes.
type UserMetadataConnectionsChannels struct {
	// An array of HTTP methods permitted on this URI.
	Options []string `json:"options"`
	// The total number of channels on this connection.
	Total float32 `json:"total"`
	// The API URI that resolves to the connection data.
	Uri string `json:"uri"`
}

type _UserMetadataConnectionsChannels UserMetadataConnectionsChannels

// NewUserMetadataConnectionsChannels instantiates a new UserMetadataConnectionsChannels object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewUserMetadataConnectionsChannels(options []string, total float32, uri string) *UserMetadataConnectionsChannels {
	this := UserMetadataConnectionsChannels{}
	this.Options = options
	this.Total = total
	this.Uri = uri
	return &this
}

// NewUserMetadataConnectionsChannelsWithDefaults instantiates a new UserMetadataConnectionsChannels object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewUserMetadataConnectionsChannelsWithDefaults() *UserMetadataConnectionsChannels {
	this := UserMetadataConnectionsChannels{}
	return &this
}

// GetOptions returns the Options field value
func (o *UserMetadataConnectionsChannels) GetOptions() []string {
	if o == nil {
		var ret []string
		return ret
	}

	return o.Options
}

// GetOptionsOk returns a tuple with the Options field value
// and a boolean to check if the value has been set.
func (o *UserMetadataConnectionsChannels) GetOptionsOk() ([]string, bool) {
	if o == nil {
		return nil, false
	}
	return o.Options, true
}

// SetOptions sets field value
func (o *UserMetadataConnectionsChannels) SetOptions(v []string) {
	o.Options = v
}

// GetTotal returns the Total field value
func (o *UserMetadataConnectionsChannels) GetTotal() float32 {
	if o == nil {
		var ret float32
		return ret
	}

	return o.Total
}

// GetTotalOk returns a tuple with the Total field value
// and a boolean to check if the value has been set.
func (o *UserMetadataConnectionsChannels) GetTotalOk() (*float32, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Total, true
}

// SetTotal sets field value
func (o *UserMetadataConnectionsChannels) SetTotal(v float32) {
	o.Total = v
}

// GetUri returns the Uri field value
func (o *UserMetadataConnectionsChannels) GetUri() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Uri
}

// GetUriOk returns a tuple with the Uri field value
// and a boolean to check if the value has been set.
func (o *UserMetadataConnectionsChannels) GetUriOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Uri, true
}

// SetUri sets field value
func (o *UserMetadataConnectionsChannels) SetUri(v string) {
	o.Uri = v
}

func (o UserMetadataConnectionsChannels) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o UserMetadataConnectionsChannels) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["options"] = o.Options
	toSerialize["total"] = o.Total
	toSerialize["uri"] = o.Uri
	return toSerialize, nil
}

func (o *UserMetadataConnectionsChannels) UnmarshalJSON(data []byte) (err error) {
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

	varUserMetadataConnectionsChannels := _UserMetadataConnectionsChannels{}

	decoder := json.NewDecoder(bytes.NewReader(data))
	decoder.DisallowUnknownFields()
	err = decoder.Decode(&varUserMetadataConnectionsChannels)

	if err != nil {
		return err
	}

	*o = UserMetadataConnectionsChannels(varUserMetadataConnectionsChannels)

	return err
}

type NullableUserMetadataConnectionsChannels struct {
	value *UserMetadataConnectionsChannels
	isSet bool
}

func (v NullableUserMetadataConnectionsChannels) Get() *UserMetadataConnectionsChannels {
	return v.value
}

func (v *NullableUserMetadataConnectionsChannels) Set(val *UserMetadataConnectionsChannels) {
	v.value = val
	v.isSet = true
}

func (v NullableUserMetadataConnectionsChannels) IsSet() bool {
	return v.isSet
}

func (v *NullableUserMetadataConnectionsChannels) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableUserMetadataConnectionsChannels(val *UserMetadataConnectionsChannels) *NullableUserMetadataConnectionsChannels {
	return &NullableUserMetadataConnectionsChannels{value: val, isSet: true}
}

func (v NullableUserMetadataConnectionsChannels) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableUserMetadataConnectionsChannels) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

