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

// checks if the WebinarEmailProviderListInnerList type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &WebinarEmailProviderListInnerList{}

// WebinarEmailProviderListInnerList Information about the connected list.
type WebinarEmailProviderListInnerList struct {
	// The ID of the connected list.
	Id string `json:"id"`
	// The name of the connected list.
	Name string `json:"name"`
}

type _WebinarEmailProviderListInnerList WebinarEmailProviderListInnerList

// NewWebinarEmailProviderListInnerList instantiates a new WebinarEmailProviderListInnerList object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewWebinarEmailProviderListInnerList(id string, name string) *WebinarEmailProviderListInnerList {
	this := WebinarEmailProviderListInnerList{}
	this.Id = id
	this.Name = name
	return &this
}

// NewWebinarEmailProviderListInnerListWithDefaults instantiates a new WebinarEmailProviderListInnerList object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewWebinarEmailProviderListInnerListWithDefaults() *WebinarEmailProviderListInnerList {
	this := WebinarEmailProviderListInnerList{}
	return &this
}

// GetId returns the Id field value
func (o *WebinarEmailProviderListInnerList) GetId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Id
}

// GetIdOk returns a tuple with the Id field value
// and a boolean to check if the value has been set.
func (o *WebinarEmailProviderListInnerList) GetIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Id, true
}

// SetId sets field value
func (o *WebinarEmailProviderListInnerList) SetId(v string) {
	o.Id = v
}

// GetName returns the Name field value
func (o *WebinarEmailProviderListInnerList) GetName() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Name
}

// GetNameOk returns a tuple with the Name field value
// and a boolean to check if the value has been set.
func (o *WebinarEmailProviderListInnerList) GetNameOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Name, true
}

// SetName sets field value
func (o *WebinarEmailProviderListInnerList) SetName(v string) {
	o.Name = v
}

func (o WebinarEmailProviderListInnerList) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o WebinarEmailProviderListInnerList) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["id"] = o.Id
	toSerialize["name"] = o.Name
	return toSerialize, nil
}

func (o *WebinarEmailProviderListInnerList) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"id",
		"name",
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

	varWebinarEmailProviderListInnerList := _WebinarEmailProviderListInnerList{}

	decoder := json.NewDecoder(bytes.NewReader(data))
	decoder.DisallowUnknownFields()
	err = decoder.Decode(&varWebinarEmailProviderListInnerList)

	if err != nil {
		return err
	}

	*o = WebinarEmailProviderListInnerList(varWebinarEmailProviderListInnerList)

	return err
}

type NullableWebinarEmailProviderListInnerList struct {
	value *WebinarEmailProviderListInnerList
	isSet bool
}

func (v NullableWebinarEmailProviderListInnerList) Get() *WebinarEmailProviderListInnerList {
	return v.value
}

func (v *NullableWebinarEmailProviderListInnerList) Set(val *WebinarEmailProviderListInnerList) {
	v.value = val
	v.isSet = true
}

func (v NullableWebinarEmailProviderListInnerList) IsSet() bool {
	return v.isSet
}

func (v *NullableWebinarEmailProviderListInnerList) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableWebinarEmailProviderListInnerList(val *WebinarEmailProviderListInnerList) *NullableWebinarEmailProviderListInnerList {
	return &NullableWebinarEmailProviderListInnerList{value: val, isSet: true}
}

func (v NullableWebinarEmailProviderListInnerList) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableWebinarEmailProviderListInnerList) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

