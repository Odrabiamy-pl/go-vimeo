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

// checks if the PermissionPolicy type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &PermissionPolicy{}

// PermissionPolicy struct for PermissionPolicy
type PermissionPolicy struct {
	// The time at which the permission policy was created.
	CreatedOn string `json:"created_on"`
	// The display description of the permission policy, translated where applicable.
	DisplayDescription string `json:"display_description"`
	// The display name of the permission policy, translated where applicable.
	DisplayName string `json:"display_name"`
	// The time at which the permission policy was last modified.
	ModifiedOn NullableString `json:"modified_on"`
	// The name of the permission policy.
	Name string `json:"name"`
	// The permission actions associated with the policy.
	PermissionActions map[string]interface{} `json:"permission_actions"`
	// The URI of the permission policy.
	Uri string `json:"uri"`
}

type _PermissionPolicy PermissionPolicy

// NewPermissionPolicy instantiates a new PermissionPolicy object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewPermissionPolicy(createdOn string, displayDescription string, displayName string, modifiedOn NullableString, name string, permissionActions map[string]interface{}, uri string) *PermissionPolicy {
	this := PermissionPolicy{}
	this.CreatedOn = createdOn
	this.DisplayDescription = displayDescription
	this.DisplayName = displayName
	this.ModifiedOn = modifiedOn
	this.Name = name
	this.PermissionActions = permissionActions
	this.Uri = uri
	return &this
}

// NewPermissionPolicyWithDefaults instantiates a new PermissionPolicy object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewPermissionPolicyWithDefaults() *PermissionPolicy {
	this := PermissionPolicy{}
	return &this
}

// GetCreatedOn returns the CreatedOn field value
func (o *PermissionPolicy) GetCreatedOn() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.CreatedOn
}

// GetCreatedOnOk returns a tuple with the CreatedOn field value
// and a boolean to check if the value has been set.
func (o *PermissionPolicy) GetCreatedOnOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.CreatedOn, true
}

// SetCreatedOn sets field value
func (o *PermissionPolicy) SetCreatedOn(v string) {
	o.CreatedOn = v
}

// GetDisplayDescription returns the DisplayDescription field value
func (o *PermissionPolicy) GetDisplayDescription() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.DisplayDescription
}

// GetDisplayDescriptionOk returns a tuple with the DisplayDescription field value
// and a boolean to check if the value has been set.
func (o *PermissionPolicy) GetDisplayDescriptionOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.DisplayDescription, true
}

// SetDisplayDescription sets field value
func (o *PermissionPolicy) SetDisplayDescription(v string) {
	o.DisplayDescription = v
}

// GetDisplayName returns the DisplayName field value
func (o *PermissionPolicy) GetDisplayName() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.DisplayName
}

// GetDisplayNameOk returns a tuple with the DisplayName field value
// and a boolean to check if the value has been set.
func (o *PermissionPolicy) GetDisplayNameOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.DisplayName, true
}

// SetDisplayName sets field value
func (o *PermissionPolicy) SetDisplayName(v string) {
	o.DisplayName = v
}

// GetModifiedOn returns the ModifiedOn field value
// If the value is explicit nil, the zero value for string will be returned
func (o *PermissionPolicy) GetModifiedOn() string {
	if o == nil || o.ModifiedOn.Get() == nil {
		var ret string
		return ret
	}

	return *o.ModifiedOn.Get()
}

// GetModifiedOnOk returns a tuple with the ModifiedOn field value
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *PermissionPolicy) GetModifiedOnOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.ModifiedOn.Get(), o.ModifiedOn.IsSet()
}

// SetModifiedOn sets field value
func (o *PermissionPolicy) SetModifiedOn(v string) {
	o.ModifiedOn.Set(&v)
}

// GetName returns the Name field value
func (o *PermissionPolicy) GetName() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Name
}

// GetNameOk returns a tuple with the Name field value
// and a boolean to check if the value has been set.
func (o *PermissionPolicy) GetNameOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Name, true
}

// SetName sets field value
func (o *PermissionPolicy) SetName(v string) {
	o.Name = v
}

// GetPermissionActions returns the PermissionActions field value
func (o *PermissionPolicy) GetPermissionActions() map[string]interface{} {
	if o == nil {
		var ret map[string]interface{}
		return ret
	}

	return o.PermissionActions
}

// GetPermissionActionsOk returns a tuple with the PermissionActions field value
// and a boolean to check if the value has been set.
func (o *PermissionPolicy) GetPermissionActionsOk() (map[string]interface{}, bool) {
	if o == nil {
		return map[string]interface{}{}, false
	}
	return o.PermissionActions, true
}

// SetPermissionActions sets field value
func (o *PermissionPolicy) SetPermissionActions(v map[string]interface{}) {
	o.PermissionActions = v
}

// GetUri returns the Uri field value
func (o *PermissionPolicy) GetUri() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Uri
}

// GetUriOk returns a tuple with the Uri field value
// and a boolean to check if the value has been set.
func (o *PermissionPolicy) GetUriOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Uri, true
}

// SetUri sets field value
func (o *PermissionPolicy) SetUri(v string) {
	o.Uri = v
}

func (o PermissionPolicy) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o PermissionPolicy) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["created_on"] = o.CreatedOn
	toSerialize["display_description"] = o.DisplayDescription
	toSerialize["display_name"] = o.DisplayName
	toSerialize["modified_on"] = o.ModifiedOn.Get()
	toSerialize["name"] = o.Name
	toSerialize["permission_actions"] = o.PermissionActions
	toSerialize["uri"] = o.Uri
	return toSerialize, nil
}

func (o *PermissionPolicy) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"created_on",
		"display_description",
		"display_name",
		"modified_on",
		"name",
		"permission_actions",
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

	varPermissionPolicy := _PermissionPolicy{}

	decoder := json.NewDecoder(bytes.NewReader(data))
	decoder.DisallowUnknownFields()
	err = decoder.Decode(&varPermissionPolicy)

	if err != nil {
		return err
	}

	*o = PermissionPolicy(varPermissionPolicy)

	return err
}

type NullablePermissionPolicy struct {
	value *PermissionPolicy
	isSet bool
}

func (v NullablePermissionPolicy) Get() *PermissionPolicy {
	return v.value
}

func (v *NullablePermissionPolicy) Set(val *PermissionPolicy) {
	v.value = val
	v.isSet = true
}

func (v NullablePermissionPolicy) IsSet() bool {
	return v.isSet
}

func (v *NullablePermissionPolicy) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullablePermissionPolicy(val *PermissionPolicy) *NullablePermissionPolicy {
	return &NullablePermissionPolicy{value: val, isSet: true}
}

func (v NullablePermissionPolicy) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullablePermissionPolicy) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


