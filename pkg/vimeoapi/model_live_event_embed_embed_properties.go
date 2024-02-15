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

// checks if the LiveEventEmbedEmbedProperties type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &LiveEventEmbedEmbedProperties{}

// LiveEventEmbedEmbedProperties The height, width, and source URL properties used to generate the fixed HTML embed code.
type LiveEventEmbedEmbedProperties struct {
	// The height used to generate the fixed HTML embed code.
	Height string `json:"height"`
	// The source URL used to generate the fixed HTML embed code.
	SourceUrl string `json:"source_url"`
	// The width used to generate the fixed HTML embed code.
	Width string `json:"width"`
}

type _LiveEventEmbedEmbedProperties LiveEventEmbedEmbedProperties

// NewLiveEventEmbedEmbedProperties instantiates a new LiveEventEmbedEmbedProperties object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewLiveEventEmbedEmbedProperties(height string, sourceUrl string, width string) *LiveEventEmbedEmbedProperties {
	this := LiveEventEmbedEmbedProperties{}
	this.Height = height
	this.SourceUrl = sourceUrl
	this.Width = width
	return &this
}

// NewLiveEventEmbedEmbedPropertiesWithDefaults instantiates a new LiveEventEmbedEmbedProperties object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewLiveEventEmbedEmbedPropertiesWithDefaults() *LiveEventEmbedEmbedProperties {
	this := LiveEventEmbedEmbedProperties{}
	return &this
}

// GetHeight returns the Height field value
func (o *LiveEventEmbedEmbedProperties) GetHeight() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Height
}

// GetHeightOk returns a tuple with the Height field value
// and a boolean to check if the value has been set.
func (o *LiveEventEmbedEmbedProperties) GetHeightOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Height, true
}

// SetHeight sets field value
func (o *LiveEventEmbedEmbedProperties) SetHeight(v string) {
	o.Height = v
}

// GetSourceUrl returns the SourceUrl field value
func (o *LiveEventEmbedEmbedProperties) GetSourceUrl() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.SourceUrl
}

// GetSourceUrlOk returns a tuple with the SourceUrl field value
// and a boolean to check if the value has been set.
func (o *LiveEventEmbedEmbedProperties) GetSourceUrlOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.SourceUrl, true
}

// SetSourceUrl sets field value
func (o *LiveEventEmbedEmbedProperties) SetSourceUrl(v string) {
	o.SourceUrl = v
}

// GetWidth returns the Width field value
func (o *LiveEventEmbedEmbedProperties) GetWidth() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Width
}

// GetWidthOk returns a tuple with the Width field value
// and a boolean to check if the value has been set.
func (o *LiveEventEmbedEmbedProperties) GetWidthOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Width, true
}

// SetWidth sets field value
func (o *LiveEventEmbedEmbedProperties) SetWidth(v string) {
	o.Width = v
}

func (o LiveEventEmbedEmbedProperties) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o LiveEventEmbedEmbedProperties) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["height"] = o.Height
	toSerialize["source_url"] = o.SourceUrl
	toSerialize["width"] = o.Width
	return toSerialize, nil
}

func (o *LiveEventEmbedEmbedProperties) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"height",
		"source_url",
		"width",
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

	varLiveEventEmbedEmbedProperties := _LiveEventEmbedEmbedProperties{}

	decoder := json.NewDecoder(bytes.NewReader(data))
	decoder.DisallowUnknownFields()
	err = decoder.Decode(&varLiveEventEmbedEmbedProperties)

	if err != nil {
		return err
	}

	*o = LiveEventEmbedEmbedProperties(varLiveEventEmbedEmbedProperties)

	return err
}

type NullableLiveEventEmbedEmbedProperties struct {
	value *LiveEventEmbedEmbedProperties
	isSet bool
}

func (v NullableLiveEventEmbedEmbedProperties) Get() *LiveEventEmbedEmbedProperties {
	return v.value
}

func (v *NullableLiveEventEmbedEmbedProperties) Set(val *LiveEventEmbedEmbedProperties) {
	v.value = val
	v.isSet = true
}

func (v NullableLiveEventEmbedEmbedProperties) IsSet() bool {
	return v.isSet
}

func (v *NullableLiveEventEmbedEmbedProperties) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableLiveEventEmbedEmbedProperties(val *LiveEventEmbedEmbedProperties) *NullableLiveEventEmbedEmbedProperties {
	return &NullableLiveEventEmbedEmbedProperties{value: val, isSet: true}
}

func (v NullableLiveEventEmbedEmbedProperties) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableLiveEventEmbedEmbedProperties) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

