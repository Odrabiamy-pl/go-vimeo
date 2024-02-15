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

// checks if the VideoStats type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &VideoStats{}

// VideoStats A collection of analytics associated with the video.
type VideoStats struct {
	// The current total number of times that the video has been played.
	Plays NullableFloat32 `json:"plays"`
}

type _VideoStats VideoStats

// NewVideoStats instantiates a new VideoStats object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewVideoStats(plays NullableFloat32) *VideoStats {
	this := VideoStats{}
	this.Plays = plays
	return &this
}

// NewVideoStatsWithDefaults instantiates a new VideoStats object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewVideoStatsWithDefaults() *VideoStats {
	this := VideoStats{}
	return &this
}

// GetPlays returns the Plays field value
// If the value is explicit nil, the zero value for float32 will be returned
func (o *VideoStats) GetPlays() float32 {
	if o == nil || o.Plays.Get() == nil {
		var ret float32
		return ret
	}

	return *o.Plays.Get()
}

// GetPlaysOk returns a tuple with the Plays field value
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *VideoStats) GetPlaysOk() (*float32, bool) {
	if o == nil {
		return nil, false
	}
	return o.Plays.Get(), o.Plays.IsSet()
}

// SetPlays sets field value
func (o *VideoStats) SetPlays(v float32) {
	o.Plays.Set(&v)
}

func (o VideoStats) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o VideoStats) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["plays"] = o.Plays.Get()
	return toSerialize, nil
}

func (o *VideoStats) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"plays",
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

	varVideoStats := _VideoStats{}

	decoder := json.NewDecoder(bytes.NewReader(data))
	decoder.DisallowUnknownFields()
	err = decoder.Decode(&varVideoStats)

	if err != nil {
		return err
	}

	*o = VideoStats(varVideoStats)

	return err
}

type NullableVideoStats struct {
	value *VideoStats
	isSet bool
}

func (v NullableVideoStats) Get() *VideoStats {
	return v.value
}

func (v *NullableVideoStats) Set(val *VideoStats) {
	v.value = val
	v.isSet = true
}

func (v NullableVideoStats) IsSet() bool {
	return v.isSet
}

func (v *NullableVideoStats) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableVideoStats(val *VideoStats) *NullableVideoStats {
	return &NullableVideoStats{value: val, isSet: true}
}

func (v NullableVideoStats) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableVideoStats) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


