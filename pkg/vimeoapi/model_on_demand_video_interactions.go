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

// checks if the OnDemandVideoInteractions type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &OnDemandVideoInteractions{}

// OnDemandVideoInteractions An object containing information about how the authenticated user can interact with the video's On Demand page.
type OnDemandVideoInteractions struct {
	Page OnDemandVideoInteractionsPage `json:"page"`
}

type _OnDemandVideoInteractions OnDemandVideoInteractions

// NewOnDemandVideoInteractions instantiates a new OnDemandVideoInteractions object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewOnDemandVideoInteractions(page OnDemandVideoInteractionsPage) *OnDemandVideoInteractions {
	this := OnDemandVideoInteractions{}
	this.Page = page
	return &this
}

// NewOnDemandVideoInteractionsWithDefaults instantiates a new OnDemandVideoInteractions object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewOnDemandVideoInteractionsWithDefaults() *OnDemandVideoInteractions {
	this := OnDemandVideoInteractions{}
	return &this
}

// GetPage returns the Page field value
func (o *OnDemandVideoInteractions) GetPage() OnDemandVideoInteractionsPage {
	if o == nil {
		var ret OnDemandVideoInteractionsPage
		return ret
	}

	return o.Page
}

// GetPageOk returns a tuple with the Page field value
// and a boolean to check if the value has been set.
func (o *OnDemandVideoInteractions) GetPageOk() (*OnDemandVideoInteractionsPage, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Page, true
}

// SetPage sets field value
func (o *OnDemandVideoInteractions) SetPage(v OnDemandVideoInteractionsPage) {
	o.Page = v
}

func (o OnDemandVideoInteractions) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o OnDemandVideoInteractions) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["page"] = o.Page
	return toSerialize, nil
}

func (o *OnDemandVideoInteractions) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"page",
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

	varOnDemandVideoInteractions := _OnDemandVideoInteractions{}

	decoder := json.NewDecoder(bytes.NewReader(data))
	decoder.DisallowUnknownFields()
	err = decoder.Decode(&varOnDemandVideoInteractions)

	if err != nil {
		return err
	}

	*o = OnDemandVideoInteractions(varOnDemandVideoInteractions)

	return err
}

type NullableOnDemandVideoInteractions struct {
	value *OnDemandVideoInteractions
	isSet bool
}

func (v NullableOnDemandVideoInteractions) Get() *OnDemandVideoInteractions {
	return v.value
}

func (v *NullableOnDemandVideoInteractions) Set(val *OnDemandVideoInteractions) {
	v.value = val
	v.isSet = true
}

func (v NullableOnDemandVideoInteractions) IsSet() bool {
	return v.isSet
}

func (v *NullableOnDemandVideoInteractions) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableOnDemandVideoInteractions(val *OnDemandVideoInteractions) *NullableOnDemandVideoInteractions {
	return &NullableOnDemandVideoInteractions{value: val, isSet: true}
}

func (v NullableOnDemandVideoInteractions) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableOnDemandVideoInteractions) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

