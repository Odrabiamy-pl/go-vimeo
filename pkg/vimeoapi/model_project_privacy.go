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

// checks if the ProjectPrivacy type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &ProjectPrivacy{}

// ProjectPrivacy The privacy settings of the folder.
type ProjectPrivacy struct {
	// The privacy setting for accessing the folder.  Option descriptions:  * `anybody` - Anyone with the link can access the contents of the folder. This privacy setting appears as `Public` on the Vimeo front end.  * `nobody` - Only the owner and those team members that the owner has explicitly invited can access the contents of the folder. This privacy setting appears as `Private` on the Vimeo front end.  * `team` - Only those team members with the link can access the contents of the folder.
	View string `json:"view"`
}

// NewProjectPrivacy instantiates a new ProjectPrivacy object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewProjectPrivacy(view string) *ProjectPrivacy {
	this := ProjectPrivacy{}
	this.View = view
	return &this
}

// NewProjectPrivacyWithDefaults instantiates a new ProjectPrivacy object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewProjectPrivacyWithDefaults() *ProjectPrivacy {
	this := ProjectPrivacy{}
	return &this
}

// GetView returns the View field value
func (o *ProjectPrivacy) GetView() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.View
}

// GetViewOk returns a tuple with the View field value
// and a boolean to check if the value has been set.
func (o *ProjectPrivacy) GetViewOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.View, true
}

// SetView sets field value
func (o *ProjectPrivacy) SetView(v string) {
	o.View = v
}

func (o ProjectPrivacy) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o ProjectPrivacy) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["view"] = o.View
	return toSerialize, nil
}

type NullableProjectPrivacy struct {
	value *ProjectPrivacy
	isSet bool
}

func (v NullableProjectPrivacy) Get() *ProjectPrivacy {
	return v.value
}

func (v *NullableProjectPrivacy) Set(val *ProjectPrivacy) {
	v.value = val
	v.isSet = true
}

func (v NullableProjectPrivacy) IsSet() bool {
	return v.isSet
}

func (v *NullableProjectPrivacy) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableProjectPrivacy(val *ProjectPrivacy) *NullableProjectPrivacy {
	return &NullableProjectPrivacy{value: val, isSet: true}
}

func (v NullableProjectPrivacy) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableProjectPrivacy) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
