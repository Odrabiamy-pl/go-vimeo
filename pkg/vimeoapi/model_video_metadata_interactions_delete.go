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

// checks if the VideoMetadataInteractionsDelete type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &VideoMetadataInteractionsDelete{}

// VideoMetadataInteractionsDelete Information about where and how to delete a video.
type VideoMetadataInteractionsDelete struct {
	// An array of HTTP methods permitted on this URI.
	Options []string `json:"options"`
	// The API URI that resolves to the connection data.
	Uri string `json:"uri"`
}

type _VideoMetadataInteractionsDelete VideoMetadataInteractionsDelete

// NewVideoMetadataInteractionsDelete instantiates a new VideoMetadataInteractionsDelete object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewVideoMetadataInteractionsDelete(options []string, uri string) *VideoMetadataInteractionsDelete {
	this := VideoMetadataInteractionsDelete{}
	this.Options = options
	this.Uri = uri
	return &this
}

// NewVideoMetadataInteractionsDeleteWithDefaults instantiates a new VideoMetadataInteractionsDelete object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewVideoMetadataInteractionsDeleteWithDefaults() *VideoMetadataInteractionsDelete {
	this := VideoMetadataInteractionsDelete{}
	return &this
}

// GetOptions returns the Options field value
func (o *VideoMetadataInteractionsDelete) GetOptions() []string {
	if o == nil {
		var ret []string
		return ret
	}

	return o.Options
}

// GetOptionsOk returns a tuple with the Options field value
// and a boolean to check if the value has been set.
func (o *VideoMetadataInteractionsDelete) GetOptionsOk() ([]string, bool) {
	if o == nil {
		return nil, false
	}
	return o.Options, true
}

// SetOptions sets field value
func (o *VideoMetadataInteractionsDelete) SetOptions(v []string) {
	o.Options = v
}

// GetUri returns the Uri field value
func (o *VideoMetadataInteractionsDelete) GetUri() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Uri
}

// GetUriOk returns a tuple with the Uri field value
// and a boolean to check if the value has been set.
func (o *VideoMetadataInteractionsDelete) GetUriOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Uri, true
}

// SetUri sets field value
func (o *VideoMetadataInteractionsDelete) SetUri(v string) {
	o.Uri = v
}

func (o VideoMetadataInteractionsDelete) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o VideoMetadataInteractionsDelete) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["options"] = o.Options
	toSerialize["uri"] = o.Uri
	return toSerialize, nil
}

func (o *VideoMetadataInteractionsDelete) UnmarshalJSON(data []byte) (err error) {
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

	varVideoMetadataInteractionsDelete := _VideoMetadataInteractionsDelete{}

	decoder := json.NewDecoder(bytes.NewReader(data))
	decoder.DisallowUnknownFields()
	err = decoder.Decode(&varVideoMetadataInteractionsDelete)

	if err != nil {
		return err
	}

	*o = VideoMetadataInteractionsDelete(varVideoMetadataInteractionsDelete)

	return err
}

type NullableVideoMetadataInteractionsDelete struct {
	value *VideoMetadataInteractionsDelete
	isSet bool
}

func (v NullableVideoMetadataInteractionsDelete) Get() *VideoMetadataInteractionsDelete {
	return v.value
}

func (v *NullableVideoMetadataInteractionsDelete) Set(val *VideoMetadataInteractionsDelete) {
	v.value = val
	v.isSet = true
}

func (v NullableVideoMetadataInteractionsDelete) IsSet() bool {
	return v.isSet
}

func (v *NullableVideoMetadataInteractionsDelete) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableVideoMetadataInteractionsDelete(val *VideoMetadataInteractionsDelete) *NullableVideoMetadataInteractionsDelete {
	return &NullableVideoMetadataInteractionsDelete{value: val, isSet: true}
}

func (v NullableVideoMetadataInteractionsDelete) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableVideoMetadataInteractionsDelete) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


