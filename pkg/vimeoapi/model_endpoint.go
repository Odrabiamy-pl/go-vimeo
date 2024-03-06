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

// checks if the Endpoint type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &Endpoint{}

// Endpoint struct for Endpoint
type Endpoint struct {
	// All HTTP methods permitted on this endpoint.
	Methods []string `json:"methods"`
	// The path section of the URL, which, when appended to the API host `https:///api.vimeo.com`, builds a full API endpoint.
	Path string `json:"path"`
}

// NewEndpoint instantiates a new Endpoint object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewEndpoint(methods []string, path string) *Endpoint {
	this := Endpoint{}
	this.Methods = methods
	this.Path = path
	return &this
}

// NewEndpointWithDefaults instantiates a new Endpoint object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewEndpointWithDefaults() *Endpoint {
	this := Endpoint{}
	return &this
}

// GetMethods returns the Methods field value
func (o *Endpoint) GetMethods() []string {
	if o == nil {
		var ret []string
		return ret
	}

	return o.Methods
}

// GetMethodsOk returns a tuple with the Methods field value
// and a boolean to check if the value has been set.
func (o *Endpoint) GetMethodsOk() ([]string, bool) {
	if o == nil {
		return nil, false
	}
	return o.Methods, true
}

// SetMethods sets field value
func (o *Endpoint) SetMethods(v []string) {
	o.Methods = v
}

// GetPath returns the Path field value
func (o *Endpoint) GetPath() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Path
}

// GetPathOk returns a tuple with the Path field value
// and a boolean to check if the value has been set.
func (o *Endpoint) GetPathOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Path, true
}

// SetPath sets field value
func (o *Endpoint) SetPath(v string) {
	o.Path = v
}

func (o Endpoint) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o Endpoint) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["methods"] = o.Methods
	toSerialize["path"] = o.Path
	return toSerialize, nil
}

type NullableEndpoint struct {
	value *Endpoint
	isSet bool
}

func (v NullableEndpoint) Get() *Endpoint {
	return v.value
}

func (v *NullableEndpoint) Set(val *Endpoint) {
	v.value = val
	v.isSet = true
}

func (v NullableEndpoint) IsSet() bool {
	return v.isSet
}

func (v *NullableEndpoint) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableEndpoint(val *Endpoint) *NullableEndpoint {
	return &NullableEndpoint{value: val, isSet: true}
}

func (v NullableEndpoint) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableEndpoint) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
