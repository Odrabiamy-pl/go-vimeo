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

// checks if the ProjectMetadataConnectionsUserFolderAccessGrantsFolderPermissionPoliciesInner type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &ProjectMetadataConnectionsUserFolderAccessGrantsFolderPermissionPoliciesInner{}

// ProjectMetadataConnectionsUserFolderAccessGrantsFolderPermissionPoliciesInner struct for ProjectMetadataConnectionsUserFolderAccessGrantsFolderPermissionPoliciesInner
type ProjectMetadataConnectionsUserFolderAccessGrantsFolderPermissionPoliciesInner struct {
	// The permission policy's name.
	Name string `json:"name"`
	// The permission policy's API URI.
	Uri string `json:"uri"`
}

// NewProjectMetadataConnectionsUserFolderAccessGrantsFolderPermissionPoliciesInner instantiates a new ProjectMetadataConnectionsUserFolderAccessGrantsFolderPermissionPoliciesInner object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewProjectMetadataConnectionsUserFolderAccessGrantsFolderPermissionPoliciesInner(name string, uri string) *ProjectMetadataConnectionsUserFolderAccessGrantsFolderPermissionPoliciesInner {
	this := ProjectMetadataConnectionsUserFolderAccessGrantsFolderPermissionPoliciesInner{}
	this.Name = name
	this.Uri = uri
	return &this
}

// NewProjectMetadataConnectionsUserFolderAccessGrantsFolderPermissionPoliciesInnerWithDefaults instantiates a new ProjectMetadataConnectionsUserFolderAccessGrantsFolderPermissionPoliciesInner object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewProjectMetadataConnectionsUserFolderAccessGrantsFolderPermissionPoliciesInnerWithDefaults() *ProjectMetadataConnectionsUserFolderAccessGrantsFolderPermissionPoliciesInner {
	this := ProjectMetadataConnectionsUserFolderAccessGrantsFolderPermissionPoliciesInner{}
	return &this
}

// GetName returns the Name field value
func (o *ProjectMetadataConnectionsUserFolderAccessGrantsFolderPermissionPoliciesInner) GetName() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Name
}

// GetNameOk returns a tuple with the Name field value
// and a boolean to check if the value has been set.
func (o *ProjectMetadataConnectionsUserFolderAccessGrantsFolderPermissionPoliciesInner) GetNameOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Name, true
}

// SetName sets field value
func (o *ProjectMetadataConnectionsUserFolderAccessGrantsFolderPermissionPoliciesInner) SetName(v string) {
	o.Name = v
}

// GetUri returns the Uri field value
func (o *ProjectMetadataConnectionsUserFolderAccessGrantsFolderPermissionPoliciesInner) GetUri() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Uri
}

// GetUriOk returns a tuple with the Uri field value
// and a boolean to check if the value has been set.
func (o *ProjectMetadataConnectionsUserFolderAccessGrantsFolderPermissionPoliciesInner) GetUriOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Uri, true
}

// SetUri sets field value
func (o *ProjectMetadataConnectionsUserFolderAccessGrantsFolderPermissionPoliciesInner) SetUri(v string) {
	o.Uri = v
}

func (o ProjectMetadataConnectionsUserFolderAccessGrantsFolderPermissionPoliciesInner) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o ProjectMetadataConnectionsUserFolderAccessGrantsFolderPermissionPoliciesInner) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["name"] = o.Name
	toSerialize["uri"] = o.Uri
	return toSerialize, nil
}

type NullableProjectMetadataConnectionsUserFolderAccessGrantsFolderPermissionPoliciesInner struct {
	value *ProjectMetadataConnectionsUserFolderAccessGrantsFolderPermissionPoliciesInner
	isSet bool
}

func (v NullableProjectMetadataConnectionsUserFolderAccessGrantsFolderPermissionPoliciesInner) Get() *ProjectMetadataConnectionsUserFolderAccessGrantsFolderPermissionPoliciesInner {
	return v.value
}

func (v *NullableProjectMetadataConnectionsUserFolderAccessGrantsFolderPermissionPoliciesInner) Set(val *ProjectMetadataConnectionsUserFolderAccessGrantsFolderPermissionPoliciesInner) {
	v.value = val
	v.isSet = true
}

func (v NullableProjectMetadataConnectionsUserFolderAccessGrantsFolderPermissionPoliciesInner) IsSet() bool {
	return v.isSet
}

func (v *NullableProjectMetadataConnectionsUserFolderAccessGrantsFolderPermissionPoliciesInner) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableProjectMetadataConnectionsUserFolderAccessGrantsFolderPermissionPoliciesInner(val *ProjectMetadataConnectionsUserFolderAccessGrantsFolderPermissionPoliciesInner) *NullableProjectMetadataConnectionsUserFolderAccessGrantsFolderPermissionPoliciesInner {
	return &NullableProjectMetadataConnectionsUserFolderAccessGrantsFolderPermissionPoliciesInner{value: val, isSet: true}
}

func (v NullableProjectMetadataConnectionsUserFolderAccessGrantsFolderPermissionPoliciesInner) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableProjectMetadataConnectionsUserFolderAccessGrantsFolderPermissionPoliciesInner) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


