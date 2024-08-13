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

// checks if the Activity31MetadataConnections type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &Activity31MetadataConnections{}

// Activity31MetadataConnections A list of resource URIs related to the activity.
type Activity31MetadataConnections struct {
	Related NullableActivity31MetadataConnectionsRelated `json:"related"`
}

// NewActivity31MetadataConnections instantiates a new Activity31MetadataConnections object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewActivity31MetadataConnections(related NullableActivity31MetadataConnectionsRelated) *Activity31MetadataConnections {
	this := Activity31MetadataConnections{}
	this.Related = related
	return &this
}

// NewActivity31MetadataConnectionsWithDefaults instantiates a new Activity31MetadataConnections object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewActivity31MetadataConnectionsWithDefaults() *Activity31MetadataConnections {
	this := Activity31MetadataConnections{}
	return &this
}

// GetRelated returns the Related field value
// If the value is explicit nil, the zero value for Activity31MetadataConnectionsRelated will be returned
func (o *Activity31MetadataConnections) GetRelated() Activity31MetadataConnectionsRelated {
	if o == nil || o.Related.Get() == nil {
		var ret Activity31MetadataConnectionsRelated
		return ret
	}

	return *o.Related.Get()
}

// GetRelatedOk returns a tuple with the Related field value
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *Activity31MetadataConnections) GetRelatedOk() (*Activity31MetadataConnectionsRelated, bool) {
	if o == nil {
		return nil, false
	}
	return o.Related.Get(), o.Related.IsSet()
}

// SetRelated sets field value
func (o *Activity31MetadataConnections) SetRelated(v Activity31MetadataConnectionsRelated) {
	o.Related.Set(&v)
}

func (o Activity31MetadataConnections) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o Activity31MetadataConnections) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["related"] = o.Related.Get()
	return toSerialize, nil
}

type NullableActivity31MetadataConnections struct {
	value *Activity31MetadataConnections
	isSet bool
}

func (v NullableActivity31MetadataConnections) Get() *Activity31MetadataConnections {
	return v.value
}

func (v *NullableActivity31MetadataConnections) Set(val *Activity31MetadataConnections) {
	v.value = val
	v.isSet = true
}

func (v NullableActivity31MetadataConnections) IsSet() bool {
	return v.isSet
}

func (v *NullableActivity31MetadataConnections) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableActivity31MetadataConnections(val *Activity31MetadataConnections) *NullableActivity31MetadataConnections {
	return &NullableActivity31MetadataConnections{value: val, isSet: true}
}

func (v NullableActivity31MetadataConnections) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableActivity31MetadataConnections) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


