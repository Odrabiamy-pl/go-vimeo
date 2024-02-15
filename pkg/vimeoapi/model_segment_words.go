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

// checks if the SegmentWords type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &SegmentWords{}

// SegmentWords struct for SegmentWords
type SegmentWords struct {
	// The end time of the word in milliseconds.
	EndTime NullableFloat32 `json:"end_time"`
	// The start time of the word in milliseconds.
	StartTime NullableFloat32 `json:"start_time"`
	// The word text.
	Word string `json:"word"`
}

// NewSegmentWords instantiates a new SegmentWords object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewSegmentWords(endTime NullableFloat32, startTime NullableFloat32, word string) *SegmentWords {
	this := SegmentWords{}
	this.EndTime = endTime
	this.StartTime = startTime
	this.Word = word
	return &this
}

// NewSegmentWordsWithDefaults instantiates a new SegmentWords object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewSegmentWordsWithDefaults() *SegmentWords {
	this := SegmentWords{}
	return &this
}

// GetEndTime returns the EndTime field value
// If the value is explicit nil, the zero value for float32 will be returned
func (o *SegmentWords) GetEndTime() float32 {
	if o == nil || o.EndTime.Get() == nil {
		var ret float32
		return ret
	}

	return *o.EndTime.Get()
}

// GetEndTimeOk returns a tuple with the EndTime field value
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *SegmentWords) GetEndTimeOk() (*float32, bool) {
	if o == nil {
		return nil, false
	}
	return o.EndTime.Get(), o.EndTime.IsSet()
}

// SetEndTime sets field value
func (o *SegmentWords) SetEndTime(v float32) {
	o.EndTime.Set(&v)
}

// GetStartTime returns the StartTime field value
// If the value is explicit nil, the zero value for float32 will be returned
func (o *SegmentWords) GetStartTime() float32 {
	if o == nil || o.StartTime.Get() == nil {
		var ret float32
		return ret
	}

	return *o.StartTime.Get()
}

// GetStartTimeOk returns a tuple with the StartTime field value
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *SegmentWords) GetStartTimeOk() (*float32, bool) {
	if o == nil {
		return nil, false
	}
	return o.StartTime.Get(), o.StartTime.IsSet()
}

// SetStartTime sets field value
func (o *SegmentWords) SetStartTime(v float32) {
	o.StartTime.Set(&v)
}

// GetWord returns the Word field value
func (o *SegmentWords) GetWord() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Word
}

// GetWordOk returns a tuple with the Word field value
// and a boolean to check if the value has been set.
func (o *SegmentWords) GetWordOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Word, true
}

// SetWord sets field value
func (o *SegmentWords) SetWord(v string) {
	o.Word = v
}

func (o SegmentWords) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o SegmentWords) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["end_time"] = o.EndTime.Get()
	toSerialize["start_time"] = o.StartTime.Get()
	toSerialize["word"] = o.Word
	return toSerialize, nil
}

type NullableSegmentWords struct {
	value *SegmentWords
	isSet bool
}

func (v NullableSegmentWords) Get() *SegmentWords {
	return v.value
}

func (v *NullableSegmentWords) Set(val *SegmentWords) {
	v.value = val
	v.isSet = true
}

func (v NullableSegmentWords) IsSet() bool {
	return v.isSet
}

func (v *NullableSegmentWords) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableSegmentWords(val *SegmentWords) *NullableSegmentWords {
	return &NullableSegmentWords{value: val, isSet: true}
}

func (v NullableSegmentWords) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableSegmentWords) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


