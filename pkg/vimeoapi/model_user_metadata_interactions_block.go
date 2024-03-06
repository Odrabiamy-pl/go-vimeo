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

// checks if the UserMetadataInteractionsBlock type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &UserMetadataInteractionsBlock{}

// UserMetadataInteractionsBlock Information about the block status of the authenticated user.
type UserMetadataInteractionsBlock struct {
	// Whether the authenticated user is blocking the requested user.
	Added bool `json:"added"`
	// The time in ISO 8601 format when the block occurred, or the null value if no block exists.
	AddedTime NullableString `json:"added_time"`
	// An array of the HTTP methods permitted on this URI.
	Options []string `json:"options"`
	// The URI to block or unblock the requested user.
	Uri string `json:"uri"`
}

// NewUserMetadataInteractionsBlock instantiates a new UserMetadataInteractionsBlock object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewUserMetadataInteractionsBlock(added bool, addedTime NullableString, options []string, uri string) *UserMetadataInteractionsBlock {
	this := UserMetadataInteractionsBlock{}
	this.Added = added
	this.AddedTime = addedTime
	this.Options = options
	this.Uri = uri
	return &this
}

// NewUserMetadataInteractionsBlockWithDefaults instantiates a new UserMetadataInteractionsBlock object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewUserMetadataInteractionsBlockWithDefaults() *UserMetadataInteractionsBlock {
	this := UserMetadataInteractionsBlock{}
	return &this
}

// GetAdded returns the Added field value
func (o *UserMetadataInteractionsBlock) GetAdded() bool {
	if o == nil {
		var ret bool
		return ret
	}

	return o.Added
}

// GetAddedOk returns a tuple with the Added field value
// and a boolean to check if the value has been set.
func (o *UserMetadataInteractionsBlock) GetAddedOk() (*bool, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Added, true
}

// SetAdded sets field value
func (o *UserMetadataInteractionsBlock) SetAdded(v bool) {
	o.Added = v
}

// GetAddedTime returns the AddedTime field value
// If the value is explicit nil, the zero value for string will be returned
func (o *UserMetadataInteractionsBlock) GetAddedTime() string {
	if o == nil || o.AddedTime.Get() == nil {
		var ret string
		return ret
	}

	return *o.AddedTime.Get()
}

// GetAddedTimeOk returns a tuple with the AddedTime field value
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *UserMetadataInteractionsBlock) GetAddedTimeOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.AddedTime.Get(), o.AddedTime.IsSet()
}

// SetAddedTime sets field value
func (o *UserMetadataInteractionsBlock) SetAddedTime(v string) {
	o.AddedTime.Set(&v)
}

// GetOptions returns the Options field value
func (o *UserMetadataInteractionsBlock) GetOptions() []string {
	if o == nil {
		var ret []string
		return ret
	}

	return o.Options
}

// GetOptionsOk returns a tuple with the Options field value
// and a boolean to check if the value has been set.
func (o *UserMetadataInteractionsBlock) GetOptionsOk() ([]string, bool) {
	if o == nil {
		return nil, false
	}
	return o.Options, true
}

// SetOptions sets field value
func (o *UserMetadataInteractionsBlock) SetOptions(v []string) {
	o.Options = v
}

// GetUri returns the Uri field value
func (o *UserMetadataInteractionsBlock) GetUri() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Uri
}

// GetUriOk returns a tuple with the Uri field value
// and a boolean to check if the value has been set.
func (o *UserMetadataInteractionsBlock) GetUriOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Uri, true
}

// SetUri sets field value
func (o *UserMetadataInteractionsBlock) SetUri(v string) {
	o.Uri = v
}

func (o UserMetadataInteractionsBlock) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o UserMetadataInteractionsBlock) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["added"] = o.Added
	toSerialize["added_time"] = o.AddedTime.Get()
	toSerialize["options"] = o.Options
	toSerialize["uri"] = o.Uri
	return toSerialize, nil
}

type NullableUserMetadataInteractionsBlock struct {
	value *UserMetadataInteractionsBlock
	isSet bool
}

func (v NullableUserMetadataInteractionsBlock) Get() *UserMetadataInteractionsBlock {
	return v.value
}

func (v *NullableUserMetadataInteractionsBlock) Set(val *UserMetadataInteractionsBlock) {
	v.value = val
	v.isSet = true
}

func (v NullableUserMetadataInteractionsBlock) IsSet() bool {
	return v.isSet
}

func (v *NullableUserMetadataInteractionsBlock) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableUserMetadataInteractionsBlock(val *UserMetadataInteractionsBlock) *NullableUserMetadataInteractionsBlock {
	return &NullableUserMetadataInteractionsBlock{value: val, isSet: true}
}

func (v NullableUserMetadataInteractionsBlock) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableUserMetadataInteractionsBlock) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
