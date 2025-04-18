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

// checks if the OnDemandVideoMetadataConnectionsSeason type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &OnDemandVideoMetadataConnectionsSeason{}

// OnDemandVideoMetadataConnectionsSeason Information about the video's season.
type OnDemandVideoMetadataConnectionsSeason struct {
	// The name of the season on this connection.
	Name string `json:"name"`
	// An array of HTTP methods permitted on this URI.
	Options []string `json:"options"`
	// The API URI that resolves to the connection data.
	Uri string `json:"uri"`
}

// NewOnDemandVideoMetadataConnectionsSeason instantiates a new OnDemandVideoMetadataConnectionsSeason object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewOnDemandVideoMetadataConnectionsSeason(name string, options []string, uri string) *OnDemandVideoMetadataConnectionsSeason {
	this := OnDemandVideoMetadataConnectionsSeason{}
	this.Name = name
	this.Options = options
	this.Uri = uri
	return &this
}

// NewOnDemandVideoMetadataConnectionsSeasonWithDefaults instantiates a new OnDemandVideoMetadataConnectionsSeason object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewOnDemandVideoMetadataConnectionsSeasonWithDefaults() *OnDemandVideoMetadataConnectionsSeason {
	this := OnDemandVideoMetadataConnectionsSeason{}
	return &this
}

// GetName returns the Name field value
func (o *OnDemandVideoMetadataConnectionsSeason) GetName() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Name
}

// GetNameOk returns a tuple with the Name field value
// and a boolean to check if the value has been set.
func (o *OnDemandVideoMetadataConnectionsSeason) GetNameOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Name, true
}

// SetName sets field value
func (o *OnDemandVideoMetadataConnectionsSeason) SetName(v string) {
	o.Name = v
}

// GetOptions returns the Options field value
func (o *OnDemandVideoMetadataConnectionsSeason) GetOptions() []string {
	if o == nil {
		var ret []string
		return ret
	}

	return o.Options
}

// GetOptionsOk returns a tuple with the Options field value
// and a boolean to check if the value has been set.
func (o *OnDemandVideoMetadataConnectionsSeason) GetOptionsOk() ([]string, bool) {
	if o == nil {
		return nil, false
	}
	return o.Options, true
}

// SetOptions sets field value
func (o *OnDemandVideoMetadataConnectionsSeason) SetOptions(v []string) {
	o.Options = v
}

// GetUri returns the Uri field value
func (o *OnDemandVideoMetadataConnectionsSeason) GetUri() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Uri
}

// GetUriOk returns a tuple with the Uri field value
// and a boolean to check if the value has been set.
func (o *OnDemandVideoMetadataConnectionsSeason) GetUriOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Uri, true
}

// SetUri sets field value
func (o *OnDemandVideoMetadataConnectionsSeason) SetUri(v string) {
	o.Uri = v
}

func (o OnDemandVideoMetadataConnectionsSeason) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o OnDemandVideoMetadataConnectionsSeason) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["name"] = o.Name
	toSerialize["options"] = o.Options
	toSerialize["uri"] = o.Uri
	return toSerialize, nil
}

type NullableOnDemandVideoMetadataConnectionsSeason struct {
	value *OnDemandVideoMetadataConnectionsSeason
	isSet bool
}

func (v NullableOnDemandVideoMetadataConnectionsSeason) Get() *OnDemandVideoMetadataConnectionsSeason {
	return v.value
}

func (v *NullableOnDemandVideoMetadataConnectionsSeason) Set(val *OnDemandVideoMetadataConnectionsSeason) {
	v.value = val
	v.isSet = true
}

func (v NullableOnDemandVideoMetadataConnectionsSeason) IsSet() bool {
	return v.isSet
}

func (v *NullableOnDemandVideoMetadataConnectionsSeason) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableOnDemandVideoMetadataConnectionsSeason(val *OnDemandVideoMetadataConnectionsSeason) *NullableOnDemandVideoMetadataConnectionsSeason {
	return &NullableOnDemandVideoMetadataConnectionsSeason{value: val, isSet: true}
}

func (v NullableOnDemandVideoMetadataConnectionsSeason) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableOnDemandVideoMetadataConnectionsSeason) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


