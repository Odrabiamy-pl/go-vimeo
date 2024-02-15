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

// checks if the GroupMetadataInteractions type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &GroupMetadataInteractions{}

// GroupMetadataInteractions User actions that have involved the group. This data requires a bearer token with the `private` scope.
type GroupMetadataInteractions struct {
	Join GroupMetadataInteractionsJoin `json:"join"`
}

type _GroupMetadataInteractions GroupMetadataInteractions

// NewGroupMetadataInteractions instantiates a new GroupMetadataInteractions object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewGroupMetadataInteractions(join GroupMetadataInteractionsJoin) *GroupMetadataInteractions {
	this := GroupMetadataInteractions{}
	this.Join = join
	return &this
}

// NewGroupMetadataInteractionsWithDefaults instantiates a new GroupMetadataInteractions object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewGroupMetadataInteractionsWithDefaults() *GroupMetadataInteractions {
	this := GroupMetadataInteractions{}
	return &this
}

// GetJoin returns the Join field value
func (o *GroupMetadataInteractions) GetJoin() GroupMetadataInteractionsJoin {
	if o == nil {
		var ret GroupMetadataInteractionsJoin
		return ret
	}

	return o.Join
}

// GetJoinOk returns a tuple with the Join field value
// and a boolean to check if the value has been set.
func (o *GroupMetadataInteractions) GetJoinOk() (*GroupMetadataInteractionsJoin, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Join, true
}

// SetJoin sets field value
func (o *GroupMetadataInteractions) SetJoin(v GroupMetadataInteractionsJoin) {
	o.Join = v
}

func (o GroupMetadataInteractions) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o GroupMetadataInteractions) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["join"] = o.Join
	return toSerialize, nil
}

func (o *GroupMetadataInteractions) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"join",
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

	varGroupMetadataInteractions := _GroupMetadataInteractions{}

	decoder := json.NewDecoder(bytes.NewReader(data))
	decoder.DisallowUnknownFields()
	err = decoder.Decode(&varGroupMetadataInteractions)

	if err != nil {
		return err
	}

	*o = GroupMetadataInteractions(varGroupMetadataInteractions)

	return err
}

type NullableGroupMetadataInteractions struct {
	value *GroupMetadataInteractions
	isSet bool
}

func (v NullableGroupMetadataInteractions) Get() *GroupMetadataInteractions {
	return v.value
}

func (v *NullableGroupMetadataInteractions) Set(val *GroupMetadataInteractions) {
	v.value = val
	v.isSet = true
}

func (v NullableGroupMetadataInteractions) IsSet() bool {
	return v.isSet
}

func (v *NullableGroupMetadataInteractions) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableGroupMetadataInteractions(val *GroupMetadataInteractions) *NullableGroupMetadataInteractions {
	return &NullableGroupMetadataInteractions{value: val, isSet: true}
}

func (v NullableGroupMetadataInteractions) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableGroupMetadataInteractions) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


