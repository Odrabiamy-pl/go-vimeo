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

// checks if the Activity31MetadataConnectionsRelated type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &Activity31MetadataConnectionsRelated{}

// Activity31MetadataConnectionsRelated The activity's related content.
type Activity31MetadataConnectionsRelated struct {
	// An array of HTTP methods permitted on this URI.
	Options []string `json:"options"`
	// The API URI that resolves to the connection data.
	Uri string `json:"uri"`
}

type _Activity31MetadataConnectionsRelated Activity31MetadataConnectionsRelated

// NewActivity31MetadataConnectionsRelated instantiates a new Activity31MetadataConnectionsRelated object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewActivity31MetadataConnectionsRelated(options []string, uri string) *Activity31MetadataConnectionsRelated {
	this := Activity31MetadataConnectionsRelated{}
	this.Options = options
	this.Uri = uri
	return &this
}

// NewActivity31MetadataConnectionsRelatedWithDefaults instantiates a new Activity31MetadataConnectionsRelated object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewActivity31MetadataConnectionsRelatedWithDefaults() *Activity31MetadataConnectionsRelated {
	this := Activity31MetadataConnectionsRelated{}
	return &this
}

// GetOptions returns the Options field value
func (o *Activity31MetadataConnectionsRelated) GetOptions() []string {
	if o == nil {
		var ret []string
		return ret
	}

	return o.Options
}

// GetOptionsOk returns a tuple with the Options field value
// and a boolean to check if the value has been set.
func (o *Activity31MetadataConnectionsRelated) GetOptionsOk() ([]string, bool) {
	if o == nil {
		return nil, false
	}
	return o.Options, true
}

// SetOptions sets field value
func (o *Activity31MetadataConnectionsRelated) SetOptions(v []string) {
	o.Options = v
}

// GetUri returns the Uri field value
func (o *Activity31MetadataConnectionsRelated) GetUri() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Uri
}

// GetUriOk returns a tuple with the Uri field value
// and a boolean to check if the value has been set.
func (o *Activity31MetadataConnectionsRelated) GetUriOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Uri, true
}

// SetUri sets field value
func (o *Activity31MetadataConnectionsRelated) SetUri(v string) {
	o.Uri = v
}

func (o Activity31MetadataConnectionsRelated) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o Activity31MetadataConnectionsRelated) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["options"] = o.Options
	toSerialize["uri"] = o.Uri
	return toSerialize, nil
}

func (o *Activity31MetadataConnectionsRelated) UnmarshalJSON(data []byte) (err error) {
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

	varActivity31MetadataConnectionsRelated := _Activity31MetadataConnectionsRelated{}

	decoder := json.NewDecoder(bytes.NewReader(data))
	decoder.DisallowUnknownFields()
	err = decoder.Decode(&varActivity31MetadataConnectionsRelated)

	if err != nil {
		return err
	}

	*o = Activity31MetadataConnectionsRelated(varActivity31MetadataConnectionsRelated)

	return err
}

type NullableActivity31MetadataConnectionsRelated struct {
	value *Activity31MetadataConnectionsRelated
	isSet bool
}

func (v NullableActivity31MetadataConnectionsRelated) Get() *Activity31MetadataConnectionsRelated {
	return v.value
}

func (v *NullableActivity31MetadataConnectionsRelated) Set(val *Activity31MetadataConnectionsRelated) {
	v.value = val
	v.isSet = true
}

func (v NullableActivity31MetadataConnectionsRelated) IsSet() bool {
	return v.isSet
}

func (v *NullableActivity31MetadataConnectionsRelated) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableActivity31MetadataConnectionsRelated(val *Activity31MetadataConnectionsRelated) *NullableActivity31MetadataConnectionsRelated {
	return &NullableActivity31MetadataConnectionsRelated{value: val, isSet: true}
}

func (v NullableActivity31MetadataConnectionsRelated) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableActivity31MetadataConnectionsRelated) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


