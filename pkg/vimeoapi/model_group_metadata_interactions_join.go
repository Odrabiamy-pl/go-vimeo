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

// checks if the GroupMetadataInteractionsJoin type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &GroupMetadataInteractionsJoin{}

// GroupMetadataInteractionsJoin An action indicating that someone has joined the group. This data requires a bearer token with the `private` scope.
type GroupMetadataInteractionsJoin struct {
	// Whether the user has followed the group. This data requires a bearer token with the `private` scope.
	Added bool `json:"added"`
	// The time in ISO 8601 format when the user joined the group. This data requires a bearer token with the `private` scope.
	AddedTime NullableString `json:"added_time"`
	// The user's title. If this field isn't applicable, it takes the null value. This data requires a bearer token with the `private` scope.
	Title NullableString `json:"title"`
	// Whether the user is a moderator or subscriber. This data requires a bearer token with the `private` scope.  Option descriptions:  * `member` - The user is a member.  * `moderator` - The user is a moderator. 
	Type NullableString `json:"type"`
	// The URI for following the group. PUT to this URI to follow the group, or DELETE to this URI to unfollow the group. This data requires a bearer token with the `private` scope.
	Uri string `json:"uri"`
}

// NewGroupMetadataInteractionsJoin instantiates a new GroupMetadataInteractionsJoin object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewGroupMetadataInteractionsJoin(added bool, addedTime NullableString, title NullableString, type_ NullableString, uri string) *GroupMetadataInteractionsJoin {
	this := GroupMetadataInteractionsJoin{}
	this.Added = added
	this.AddedTime = addedTime
	this.Title = title
	this.Type = type_
	this.Uri = uri
	return &this
}

// NewGroupMetadataInteractionsJoinWithDefaults instantiates a new GroupMetadataInteractionsJoin object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewGroupMetadataInteractionsJoinWithDefaults() *GroupMetadataInteractionsJoin {
	this := GroupMetadataInteractionsJoin{}
	return &this
}

// GetAdded returns the Added field value
func (o *GroupMetadataInteractionsJoin) GetAdded() bool {
	if o == nil {
		var ret bool
		return ret
	}

	return o.Added
}

// GetAddedOk returns a tuple with the Added field value
// and a boolean to check if the value has been set.
func (o *GroupMetadataInteractionsJoin) GetAddedOk() (*bool, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Added, true
}

// SetAdded sets field value
func (o *GroupMetadataInteractionsJoin) SetAdded(v bool) {
	o.Added = v
}

// GetAddedTime returns the AddedTime field value
// If the value is explicit nil, the zero value for string will be returned
func (o *GroupMetadataInteractionsJoin) GetAddedTime() string {
	if o == nil || o.AddedTime.Get() == nil {
		var ret string
		return ret
	}

	return *o.AddedTime.Get()
}

// GetAddedTimeOk returns a tuple with the AddedTime field value
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *GroupMetadataInteractionsJoin) GetAddedTimeOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.AddedTime.Get(), o.AddedTime.IsSet()
}

// SetAddedTime sets field value
func (o *GroupMetadataInteractionsJoin) SetAddedTime(v string) {
	o.AddedTime.Set(&v)
}

// GetTitle returns the Title field value
// If the value is explicit nil, the zero value for string will be returned
func (o *GroupMetadataInteractionsJoin) GetTitle() string {
	if o == nil || o.Title.Get() == nil {
		var ret string
		return ret
	}

	return *o.Title.Get()
}

// GetTitleOk returns a tuple with the Title field value
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *GroupMetadataInteractionsJoin) GetTitleOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.Title.Get(), o.Title.IsSet()
}

// SetTitle sets field value
func (o *GroupMetadataInteractionsJoin) SetTitle(v string) {
	o.Title.Set(&v)
}

// GetType returns the Type field value
// If the value is explicit nil, the zero value for string will be returned
func (o *GroupMetadataInteractionsJoin) GetType() string {
	if o == nil || o.Type.Get() == nil {
		var ret string
		return ret
	}

	return *o.Type.Get()
}

// GetTypeOk returns a tuple with the Type field value
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *GroupMetadataInteractionsJoin) GetTypeOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.Type.Get(), o.Type.IsSet()
}

// SetType sets field value
func (o *GroupMetadataInteractionsJoin) SetType(v string) {
	o.Type.Set(&v)
}

// GetUri returns the Uri field value
func (o *GroupMetadataInteractionsJoin) GetUri() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Uri
}

// GetUriOk returns a tuple with the Uri field value
// and a boolean to check if the value has been set.
func (o *GroupMetadataInteractionsJoin) GetUriOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Uri, true
}

// SetUri sets field value
func (o *GroupMetadataInteractionsJoin) SetUri(v string) {
	o.Uri = v
}

func (o GroupMetadataInteractionsJoin) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o GroupMetadataInteractionsJoin) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["added"] = o.Added
	toSerialize["added_time"] = o.AddedTime.Get()
	toSerialize["title"] = o.Title.Get()
	toSerialize["type"] = o.Type.Get()
	toSerialize["uri"] = o.Uri
	return toSerialize, nil
}

type NullableGroupMetadataInteractionsJoin struct {
	value *GroupMetadataInteractionsJoin
	isSet bool
}

func (v NullableGroupMetadataInteractionsJoin) Get() *GroupMetadataInteractionsJoin {
	return v.value
}

func (v *NullableGroupMetadataInteractionsJoin) Set(val *GroupMetadataInteractionsJoin) {
	v.value = val
	v.isSet = true
}

func (v NullableGroupMetadataInteractionsJoin) IsSet() bool {
	return v.isSet
}

func (v *NullableGroupMetadataInteractionsJoin) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableGroupMetadataInteractionsJoin(val *GroupMetadataInteractionsJoin) *NullableGroupMetadataInteractionsJoin {
	return &NullableGroupMetadataInteractionsJoin{value: val, isSet: true}
}

func (v NullableGroupMetadataInteractionsJoin) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableGroupMetadataInteractionsJoin) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


