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

// checks if the ProjectMetadataConnectionsDataRetentionPolicy type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &ProjectMetadataConnectionsDataRetentionPolicy{}

// ProjectMetadataConnectionsDataRetentionPolicy Information about the folder's data retention policy. This data requires a bearer token with the `private` scope.
type ProjectMetadataConnectionsDataRetentionPolicy struct {
	// An array of HTTP methods permitted on this URI. This data requires a bearer token with the `private` scope.
	Options []string `json:"options"`
	// The title of the data retention policy. This data requires a bearer token with the `private` scope.
	Title string `json:"title"`
	// The URI of the data retention policy. This data requires a bearer token with the `private` scope.
	Uri string `json:"uri"`
}

type _ProjectMetadataConnectionsDataRetentionPolicy ProjectMetadataConnectionsDataRetentionPolicy

// NewProjectMetadataConnectionsDataRetentionPolicy instantiates a new ProjectMetadataConnectionsDataRetentionPolicy object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewProjectMetadataConnectionsDataRetentionPolicy(options []string, title string, uri string) *ProjectMetadataConnectionsDataRetentionPolicy {
	this := ProjectMetadataConnectionsDataRetentionPolicy{}
	this.Options = options
	this.Title = title
	this.Uri = uri
	return &this
}

// NewProjectMetadataConnectionsDataRetentionPolicyWithDefaults instantiates a new ProjectMetadataConnectionsDataRetentionPolicy object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewProjectMetadataConnectionsDataRetentionPolicyWithDefaults() *ProjectMetadataConnectionsDataRetentionPolicy {
	this := ProjectMetadataConnectionsDataRetentionPolicy{}
	return &this
}

// GetOptions returns the Options field value
func (o *ProjectMetadataConnectionsDataRetentionPolicy) GetOptions() []string {
	if o == nil {
		var ret []string
		return ret
	}

	return o.Options
}

// GetOptionsOk returns a tuple with the Options field value
// and a boolean to check if the value has been set.
func (o *ProjectMetadataConnectionsDataRetentionPolicy) GetOptionsOk() ([]string, bool) {
	if o == nil {
		return nil, false
	}
	return o.Options, true
}

// SetOptions sets field value
func (o *ProjectMetadataConnectionsDataRetentionPolicy) SetOptions(v []string) {
	o.Options = v
}

// GetTitle returns the Title field value
func (o *ProjectMetadataConnectionsDataRetentionPolicy) GetTitle() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Title
}

// GetTitleOk returns a tuple with the Title field value
// and a boolean to check if the value has been set.
func (o *ProjectMetadataConnectionsDataRetentionPolicy) GetTitleOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Title, true
}

// SetTitle sets field value
func (o *ProjectMetadataConnectionsDataRetentionPolicy) SetTitle(v string) {
	o.Title = v
}

// GetUri returns the Uri field value
func (o *ProjectMetadataConnectionsDataRetentionPolicy) GetUri() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Uri
}

// GetUriOk returns a tuple with the Uri field value
// and a boolean to check if the value has been set.
func (o *ProjectMetadataConnectionsDataRetentionPolicy) GetUriOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Uri, true
}

// SetUri sets field value
func (o *ProjectMetadataConnectionsDataRetentionPolicy) SetUri(v string) {
	o.Uri = v
}

func (o ProjectMetadataConnectionsDataRetentionPolicy) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o ProjectMetadataConnectionsDataRetentionPolicy) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["options"] = o.Options
	toSerialize["title"] = o.Title
	toSerialize["uri"] = o.Uri
	return toSerialize, nil
}

func (o *ProjectMetadataConnectionsDataRetentionPolicy) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"options",
		"title",
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

	varProjectMetadataConnectionsDataRetentionPolicy := _ProjectMetadataConnectionsDataRetentionPolicy{}

	decoder := json.NewDecoder(bytes.NewReader(data))
	decoder.DisallowUnknownFields()
	err = decoder.Decode(&varProjectMetadataConnectionsDataRetentionPolicy)

	if err != nil {
		return err
	}

	*o = ProjectMetadataConnectionsDataRetentionPolicy(varProjectMetadataConnectionsDataRetentionPolicy)

	return err
}

type NullableProjectMetadataConnectionsDataRetentionPolicy struct {
	value *ProjectMetadataConnectionsDataRetentionPolicy
	isSet bool
}

func (v NullableProjectMetadataConnectionsDataRetentionPolicy) Get() *ProjectMetadataConnectionsDataRetentionPolicy {
	return v.value
}

func (v *NullableProjectMetadataConnectionsDataRetentionPolicy) Set(val *ProjectMetadataConnectionsDataRetentionPolicy) {
	v.value = val
	v.isSet = true
}

func (v NullableProjectMetadataConnectionsDataRetentionPolicy) IsSet() bool {
	return v.isSet
}

func (v *NullableProjectMetadataConnectionsDataRetentionPolicy) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableProjectMetadataConnectionsDataRetentionPolicy(val *ProjectMetadataConnectionsDataRetentionPolicy) *NullableProjectMetadataConnectionsDataRetentionPolicy {
	return &NullableProjectMetadataConnectionsDataRetentionPolicy{value: val, isSet: true}
}

func (v NullableProjectMetadataConnectionsDataRetentionPolicy) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableProjectMetadataConnectionsDataRetentionPolicy) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


