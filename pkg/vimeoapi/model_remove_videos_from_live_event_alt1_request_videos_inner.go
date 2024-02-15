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

// checks if the RemoveVideosFromLiveEventAlt1RequestVideosInner type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &RemoveVideosFromLiveEventAlt1RequestVideosInner{}

// RemoveVideosFromLiveEventAlt1RequestVideosInner struct for RemoveVideosFromLiveEventAlt1RequestVideosInner
type RemoveVideosFromLiveEventAlt1RequestVideosInner struct {
	Video *RemoveVideosFromLiveEventAlt1RequestVideosInnerVideo `json:"video,omitempty"`
}

// NewRemoveVideosFromLiveEventAlt1RequestVideosInner instantiates a new RemoveVideosFromLiveEventAlt1RequestVideosInner object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewRemoveVideosFromLiveEventAlt1RequestVideosInner() *RemoveVideosFromLiveEventAlt1RequestVideosInner {
	this := RemoveVideosFromLiveEventAlt1RequestVideosInner{}
	return &this
}

// NewRemoveVideosFromLiveEventAlt1RequestVideosInnerWithDefaults instantiates a new RemoveVideosFromLiveEventAlt1RequestVideosInner object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewRemoveVideosFromLiveEventAlt1RequestVideosInnerWithDefaults() *RemoveVideosFromLiveEventAlt1RequestVideosInner {
	this := RemoveVideosFromLiveEventAlt1RequestVideosInner{}
	return &this
}

// GetVideo returns the Video field value if set, zero value otherwise.
func (o *RemoveVideosFromLiveEventAlt1RequestVideosInner) GetVideo() RemoveVideosFromLiveEventAlt1RequestVideosInnerVideo {
	if o == nil || IsNil(o.Video) {
		var ret RemoveVideosFromLiveEventAlt1RequestVideosInnerVideo
		return ret
	}
	return *o.Video
}

// GetVideoOk returns a tuple with the Video field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *RemoveVideosFromLiveEventAlt1RequestVideosInner) GetVideoOk() (*RemoveVideosFromLiveEventAlt1RequestVideosInnerVideo, bool) {
	if o == nil || IsNil(o.Video) {
		return nil, false
	}
	return o.Video, true
}

// HasVideo returns a boolean if a field has been set.
func (o *RemoveVideosFromLiveEventAlt1RequestVideosInner) HasVideo() bool {
	if o != nil && !IsNil(o.Video) {
		return true
	}

	return false
}

// SetVideo gets a reference to the given RemoveVideosFromLiveEventAlt1RequestVideosInnerVideo and assigns it to the Video field.
func (o *RemoveVideosFromLiveEventAlt1RequestVideosInner) SetVideo(v RemoveVideosFromLiveEventAlt1RequestVideosInnerVideo) {
	o.Video = &v
}

func (o RemoveVideosFromLiveEventAlt1RequestVideosInner) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o RemoveVideosFromLiveEventAlt1RequestVideosInner) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Video) {
		toSerialize["video"] = o.Video
	}
	return toSerialize, nil
}

type NullableRemoveVideosFromLiveEventAlt1RequestVideosInner struct {
	value *RemoveVideosFromLiveEventAlt1RequestVideosInner
	isSet bool
}

func (v NullableRemoveVideosFromLiveEventAlt1RequestVideosInner) Get() *RemoveVideosFromLiveEventAlt1RequestVideosInner {
	return v.value
}

func (v *NullableRemoveVideosFromLiveEventAlt1RequestVideosInner) Set(val *RemoveVideosFromLiveEventAlt1RequestVideosInner) {
	v.value = val
	v.isSet = true
}

func (v NullableRemoveVideosFromLiveEventAlt1RequestVideosInner) IsSet() bool {
	return v.isSet
}

func (v *NullableRemoveVideosFromLiveEventAlt1RequestVideosInner) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableRemoveVideosFromLiveEventAlt1RequestVideosInner(val *RemoveVideosFromLiveEventAlt1RequestVideosInner) *NullableRemoveVideosFromLiveEventAlt1RequestVideosInner {
	return &NullableRemoveVideosFromLiveEventAlt1RequestVideosInner{value: val, isSet: true}
}

func (v NullableRemoveVideosFromLiveEventAlt1RequestVideosInner) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableRemoveVideosFromLiveEventAlt1RequestVideosInner) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


