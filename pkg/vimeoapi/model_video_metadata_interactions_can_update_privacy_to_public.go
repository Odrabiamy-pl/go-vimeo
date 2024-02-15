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

// checks if the VideoMetadataInteractionsCanUpdatePrivacyToPublic type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &VideoMetadataInteractionsCanUpdatePrivacyToPublic{}

// VideoMetadataInteractionsCanUpdatePrivacyToPublic Whether a user can update the video privacy to public.
type VideoMetadataInteractionsCanUpdatePrivacyToPublic struct {
	// An array of HTTP methods permitted on this URI.
	Options []string `json:"options"`
	// The API URI that resolves to the connection data.
	Uri string `json:"uri"`
}

type _VideoMetadataInteractionsCanUpdatePrivacyToPublic VideoMetadataInteractionsCanUpdatePrivacyToPublic

// NewVideoMetadataInteractionsCanUpdatePrivacyToPublic instantiates a new VideoMetadataInteractionsCanUpdatePrivacyToPublic object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewVideoMetadataInteractionsCanUpdatePrivacyToPublic(options []string, uri string) *VideoMetadataInteractionsCanUpdatePrivacyToPublic {
	this := VideoMetadataInteractionsCanUpdatePrivacyToPublic{}
	this.Options = options
	this.Uri = uri
	return &this
}

// NewVideoMetadataInteractionsCanUpdatePrivacyToPublicWithDefaults instantiates a new VideoMetadataInteractionsCanUpdatePrivacyToPublic object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewVideoMetadataInteractionsCanUpdatePrivacyToPublicWithDefaults() *VideoMetadataInteractionsCanUpdatePrivacyToPublic {
	this := VideoMetadataInteractionsCanUpdatePrivacyToPublic{}
	return &this
}

// GetOptions returns the Options field value
func (o *VideoMetadataInteractionsCanUpdatePrivacyToPublic) GetOptions() []string {
	if o == nil {
		var ret []string
		return ret
	}

	return o.Options
}

// GetOptionsOk returns a tuple with the Options field value
// and a boolean to check if the value has been set.
func (o *VideoMetadataInteractionsCanUpdatePrivacyToPublic) GetOptionsOk() ([]string, bool) {
	if o == nil {
		return nil, false
	}
	return o.Options, true
}

// SetOptions sets field value
func (o *VideoMetadataInteractionsCanUpdatePrivacyToPublic) SetOptions(v []string) {
	o.Options = v
}

// GetUri returns the Uri field value
func (o *VideoMetadataInteractionsCanUpdatePrivacyToPublic) GetUri() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Uri
}

// GetUriOk returns a tuple with the Uri field value
// and a boolean to check if the value has been set.
func (o *VideoMetadataInteractionsCanUpdatePrivacyToPublic) GetUriOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Uri, true
}

// SetUri sets field value
func (o *VideoMetadataInteractionsCanUpdatePrivacyToPublic) SetUri(v string) {
	o.Uri = v
}

func (o VideoMetadataInteractionsCanUpdatePrivacyToPublic) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o VideoMetadataInteractionsCanUpdatePrivacyToPublic) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["options"] = o.Options
	toSerialize["uri"] = o.Uri
	return toSerialize, nil
}

func (o *VideoMetadataInteractionsCanUpdatePrivacyToPublic) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"options",
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

	varVideoMetadataInteractionsCanUpdatePrivacyToPublic := _VideoMetadataInteractionsCanUpdatePrivacyToPublic{}

	decoder := json.NewDecoder(bytes.NewReader(data))
	decoder.DisallowUnknownFields()
	err = decoder.Decode(&varVideoMetadataInteractionsCanUpdatePrivacyToPublic)

	if err != nil {
		return err
	}

	*o = VideoMetadataInteractionsCanUpdatePrivacyToPublic(varVideoMetadataInteractionsCanUpdatePrivacyToPublic)

	return err
}

type NullableVideoMetadataInteractionsCanUpdatePrivacyToPublic struct {
	value *VideoMetadataInteractionsCanUpdatePrivacyToPublic
	isSet bool
}

func (v NullableVideoMetadataInteractionsCanUpdatePrivacyToPublic) Get() *VideoMetadataInteractionsCanUpdatePrivacyToPublic {
	return v.value
}

func (v *NullableVideoMetadataInteractionsCanUpdatePrivacyToPublic) Set(val *VideoMetadataInteractionsCanUpdatePrivacyToPublic) {
	v.value = val
	v.isSet = true
}

func (v NullableVideoMetadataInteractionsCanUpdatePrivacyToPublic) IsSet() bool {
	return v.isSet
}

func (v *NullableVideoMetadataInteractionsCanUpdatePrivacyToPublic) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableVideoMetadataInteractionsCanUpdatePrivacyToPublic(val *VideoMetadataInteractionsCanUpdatePrivacyToPublic) *NullableVideoMetadataInteractionsCanUpdatePrivacyToPublic {
	return &NullableVideoMetadataInteractionsCanUpdatePrivacyToPublic{value: val, isSet: true}
}

func (v NullableVideoMetadataInteractionsCanUpdatePrivacyToPublic) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableVideoMetadataInteractionsCanUpdatePrivacyToPublic) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

