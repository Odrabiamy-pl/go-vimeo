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

// checks if the CategoryParent type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &CategoryParent{}

// CategoryParent The container of the category's parent category, if the current category is a subcategory.
type CategoryParent struct {
	// The URL to access the parent category in a browser.
	Link string `json:"link"`
	// The display name that identifies the parent category.
	Name string `json:"name"`
	// The unique identifier to access the parent of the category resource.
	Uri string `json:"uri"`
}

type _CategoryParent CategoryParent

// NewCategoryParent instantiates a new CategoryParent object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewCategoryParent(link string, name string, uri string) *CategoryParent {
	this := CategoryParent{}
	this.Link = link
	this.Name = name
	this.Uri = uri
	return &this
}

// NewCategoryParentWithDefaults instantiates a new CategoryParent object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewCategoryParentWithDefaults() *CategoryParent {
	this := CategoryParent{}
	return &this
}

// GetLink returns the Link field value
func (o *CategoryParent) GetLink() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Link
}

// GetLinkOk returns a tuple with the Link field value
// and a boolean to check if the value has been set.
func (o *CategoryParent) GetLinkOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Link, true
}

// SetLink sets field value
func (o *CategoryParent) SetLink(v string) {
	o.Link = v
}

// GetName returns the Name field value
func (o *CategoryParent) GetName() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Name
}

// GetNameOk returns a tuple with the Name field value
// and a boolean to check if the value has been set.
func (o *CategoryParent) GetNameOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Name, true
}

// SetName sets field value
func (o *CategoryParent) SetName(v string) {
	o.Name = v
}

// GetUri returns the Uri field value
func (o *CategoryParent) GetUri() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Uri
}

// GetUriOk returns a tuple with the Uri field value
// and a boolean to check if the value has been set.
func (o *CategoryParent) GetUriOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Uri, true
}

// SetUri sets field value
func (o *CategoryParent) SetUri(v string) {
	o.Uri = v
}

func (o CategoryParent) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o CategoryParent) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["link"] = o.Link
	toSerialize["name"] = o.Name
	toSerialize["uri"] = o.Uri
	return toSerialize, nil
}

func (o *CategoryParent) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"link",
		"name",
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

	varCategoryParent := _CategoryParent{}

	decoder := json.NewDecoder(bytes.NewReader(data))
	decoder.DisallowUnknownFields()
	err = decoder.Decode(&varCategoryParent)

	if err != nil {
		return err
	}

	*o = CategoryParent(varCategoryParent)

	return err
}

type NullableCategoryParent struct {
	value *CategoryParent
	isSet bool
}

func (v NullableCategoryParent) Get() *CategoryParent {
	return v.value
}

func (v *NullableCategoryParent) Set(val *CategoryParent) {
	v.value = val
	v.isSet = true
}

func (v NullableCategoryParent) IsSet() bool {
	return v.isSet
}

func (v *NullableCategoryParent) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableCategoryParent(val *CategoryParent) *NullableCategoryParent {
	return &NullableCategoryParent{value: val, isSet: true}
}

func (v NullableCategoryParent) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableCategoryParent) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


