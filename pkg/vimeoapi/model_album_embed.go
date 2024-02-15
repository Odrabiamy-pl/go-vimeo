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

// checks if the AlbumEmbed type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &AlbumEmbed{}

// AlbumEmbed Embed data for the showcase.
type AlbumEmbed struct {
	// The responsive HTML code to embed the showcase's playlist on a website. This field appears only when the showcase has embeddable videos.
	Html NullableString `json:"html"`
}

type _AlbumEmbed AlbumEmbed

// NewAlbumEmbed instantiates a new AlbumEmbed object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewAlbumEmbed(html NullableString) *AlbumEmbed {
	this := AlbumEmbed{}
	this.Html = html
	return &this
}

// NewAlbumEmbedWithDefaults instantiates a new AlbumEmbed object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewAlbumEmbedWithDefaults() *AlbumEmbed {
	this := AlbumEmbed{}
	return &this
}

// GetHtml returns the Html field value
// If the value is explicit nil, the zero value for string will be returned
func (o *AlbumEmbed) GetHtml() string {
	if o == nil || o.Html.Get() == nil {
		var ret string
		return ret
	}

	return *o.Html.Get()
}

// GetHtmlOk returns a tuple with the Html field value
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *AlbumEmbed) GetHtmlOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.Html.Get(), o.Html.IsSet()
}

// SetHtml sets field value
func (o *AlbumEmbed) SetHtml(v string) {
	o.Html.Set(&v)
}

func (o AlbumEmbed) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o AlbumEmbed) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["html"] = o.Html.Get()
	return toSerialize, nil
}

func (o *AlbumEmbed) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"html",
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

	varAlbumEmbed := _AlbumEmbed{}

	decoder := json.NewDecoder(bytes.NewReader(data))
	decoder.DisallowUnknownFields()
	err = decoder.Decode(&varAlbumEmbed)

	if err != nil {
		return err
	}

	*o = AlbumEmbed(varAlbumEmbed)

	return err
}

type NullableAlbumEmbed struct {
	value *AlbumEmbed
	isSet bool
}

func (v NullableAlbumEmbed) Get() *AlbumEmbed {
	return v.value
}

func (v *NullableAlbumEmbed) Set(val *AlbumEmbed) {
	v.value = val
	v.isSet = true
}

func (v NullableAlbumEmbed) IsSet() bool {
	return v.isSet
}

func (v *NullableAlbumEmbed) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableAlbumEmbed(val *AlbumEmbed) *NullableAlbumEmbed {
	return &NullableAlbumEmbed{value: val, isSet: true}
}

func (v NullableAlbumEmbed) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableAlbumEmbed) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

