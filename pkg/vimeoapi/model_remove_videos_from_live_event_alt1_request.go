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

// checks if the RemoveVideosFromLiveEventAlt1Request type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &RemoveVideosFromLiveEventAlt1Request{}

// RemoveVideosFromLiveEventAlt1Request struct for RemoveVideosFromLiveEventAlt1Request
type RemoveVideosFromLiveEventAlt1Request struct {
	// An array of video objects.
	Videos []RemoveVideosFromLiveEventAlt1RequestVideosInner `json:"videos,omitempty"`
}

// NewRemoveVideosFromLiveEventAlt1Request instantiates a new RemoveVideosFromLiveEventAlt1Request object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewRemoveVideosFromLiveEventAlt1Request() *RemoveVideosFromLiveEventAlt1Request {
	this := RemoveVideosFromLiveEventAlt1Request{}
	return &this
}

// NewRemoveVideosFromLiveEventAlt1RequestWithDefaults instantiates a new RemoveVideosFromLiveEventAlt1Request object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewRemoveVideosFromLiveEventAlt1RequestWithDefaults() *RemoveVideosFromLiveEventAlt1Request {
	this := RemoveVideosFromLiveEventAlt1Request{}
	return &this
}

// GetVideos returns the Videos field value if set, zero value otherwise.
func (o *RemoveVideosFromLiveEventAlt1Request) GetVideos() []RemoveVideosFromLiveEventAlt1RequestVideosInner {
	if o == nil || IsNil(o.Videos) {
		var ret []RemoveVideosFromLiveEventAlt1RequestVideosInner
		return ret
	}
	return o.Videos
}

// GetVideosOk returns a tuple with the Videos field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *RemoveVideosFromLiveEventAlt1Request) GetVideosOk() ([]RemoveVideosFromLiveEventAlt1RequestVideosInner, bool) {
	if o == nil || IsNil(o.Videos) {
		return nil, false
	}
	return o.Videos, true
}

// HasVideos returns a boolean if a field has been set.
func (o *RemoveVideosFromLiveEventAlt1Request) HasVideos() bool {
	if o != nil && !IsNil(o.Videos) {
		return true
	}

	return false
}

// SetVideos gets a reference to the given []RemoveVideosFromLiveEventAlt1RequestVideosInner and assigns it to the Videos field.
func (o *RemoveVideosFromLiveEventAlt1Request) SetVideos(v []RemoveVideosFromLiveEventAlt1RequestVideosInner) {
	o.Videos = v
}

func (o RemoveVideosFromLiveEventAlt1Request) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o RemoveVideosFromLiveEventAlt1Request) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Videos) {
		toSerialize["videos"] = o.Videos
	}
	return toSerialize, nil
}

type NullableRemoveVideosFromLiveEventAlt1Request struct {
	value *RemoveVideosFromLiveEventAlt1Request
	isSet bool
}

func (v NullableRemoveVideosFromLiveEventAlt1Request) Get() *RemoveVideosFromLiveEventAlt1Request {
	return v.value
}

func (v *NullableRemoveVideosFromLiveEventAlt1Request) Set(val *RemoveVideosFromLiveEventAlt1Request) {
	v.value = val
	v.isSet = true
}

func (v NullableRemoveVideosFromLiveEventAlt1Request) IsSet() bool {
	return v.isSet
}

func (v *NullableRemoveVideosFromLiveEventAlt1Request) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableRemoveVideosFromLiveEventAlt1Request(val *RemoveVideosFromLiveEventAlt1Request) *NullableRemoveVideosFromLiveEventAlt1Request {
	return &NullableRemoveVideosFromLiveEventAlt1Request{value: val, isSet: true}
}

func (v NullableRemoveVideosFromLiveEventAlt1Request) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableRemoveVideosFromLiveEventAlt1Request) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
