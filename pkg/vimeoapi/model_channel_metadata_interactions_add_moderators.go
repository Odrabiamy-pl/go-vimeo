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

// checks if the ChannelMetadataInteractionsAddModerators type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &ChannelMetadataInteractionsAddModerators{}

// ChannelMetadataInteractionsAddModerators An action indicating that the authenticated user is the owner of the channel and may therefore add other users as channel moderators. This data requires a bearer token with the `private` scope.
type ChannelMetadataInteractionsAddModerators struct {
	// An array of HTTP methods permitted on this URI. This data requires a bearer token with the `private` scope.
	Options []string `json:"options"`
	// The API URI that resolves to the connection data. This data requires a bearer token with the `private` scope.
	Uri string `json:"uri"`
}

// NewChannelMetadataInteractionsAddModerators instantiates a new ChannelMetadataInteractionsAddModerators object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewChannelMetadataInteractionsAddModerators(options []string, uri string) *ChannelMetadataInteractionsAddModerators {
	this := ChannelMetadataInteractionsAddModerators{}
	this.Options = options
	this.Uri = uri
	return &this
}

// NewChannelMetadataInteractionsAddModeratorsWithDefaults instantiates a new ChannelMetadataInteractionsAddModerators object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewChannelMetadataInteractionsAddModeratorsWithDefaults() *ChannelMetadataInteractionsAddModerators {
	this := ChannelMetadataInteractionsAddModerators{}
	return &this
}

// GetOptions returns the Options field value
func (o *ChannelMetadataInteractionsAddModerators) GetOptions() []string {
	if o == nil {
		var ret []string
		return ret
	}

	return o.Options
}

// GetOptionsOk returns a tuple with the Options field value
// and a boolean to check if the value has been set.
func (o *ChannelMetadataInteractionsAddModerators) GetOptionsOk() ([]string, bool) {
	if o == nil {
		return nil, false
	}
	return o.Options, true
}

// SetOptions sets field value
func (o *ChannelMetadataInteractionsAddModerators) SetOptions(v []string) {
	o.Options = v
}

// GetUri returns the Uri field value
func (o *ChannelMetadataInteractionsAddModerators) GetUri() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Uri
}

// GetUriOk returns a tuple with the Uri field value
// and a boolean to check if the value has been set.
func (o *ChannelMetadataInteractionsAddModerators) GetUriOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Uri, true
}

// SetUri sets field value
func (o *ChannelMetadataInteractionsAddModerators) SetUri(v string) {
	o.Uri = v
}

func (o ChannelMetadataInteractionsAddModerators) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o ChannelMetadataInteractionsAddModerators) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["options"] = o.Options
	toSerialize["uri"] = o.Uri
	return toSerialize, nil
}

type NullableChannelMetadataInteractionsAddModerators struct {
	value *ChannelMetadataInteractionsAddModerators
	isSet bool
}

func (v NullableChannelMetadataInteractionsAddModerators) Get() *ChannelMetadataInteractionsAddModerators {
	return v.value
}

func (v *NullableChannelMetadataInteractionsAddModerators) Set(val *ChannelMetadataInteractionsAddModerators) {
	v.value = val
	v.isSet = true
}

func (v NullableChannelMetadataInteractionsAddModerators) IsSet() bool {
	return v.isSet
}

func (v *NullableChannelMetadataInteractionsAddModerators) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableChannelMetadataInteractionsAddModerators(val *ChannelMetadataInteractionsAddModerators) *NullableChannelMetadataInteractionsAddModerators {
	return &NullableChannelMetadataInteractionsAddModerators{value: val, isSet: true}
}

func (v NullableChannelMetadataInteractionsAddModerators) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableChannelMetadataInteractionsAddModerators) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
