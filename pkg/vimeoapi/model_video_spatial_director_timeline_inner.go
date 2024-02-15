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

// checks if the VideoSpatialDirectorTimelineInner type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &VideoSpatialDirectorTimelineInner{}

// VideoSpatialDirectorTimelineInner struct for VideoSpatialDirectorTimelineInner
type VideoSpatialDirectorTimelineInner struct {
	// The timeline pitch value, ranging from a minimum of `-90` to a maximum of `90`.
	Pitch *float32 `json:"pitch,omitempty"`
	// The timeline roll value.
	Roll *float32 `json:"roll,omitempty"`
	// The timeline time code.
	TimeCode *float32 `json:"time_code,omitempty"`
	// The timeline yaw value, ranging from a minimum of `0` to a maximum of `360`.
	Yaw *float32 `json:"yaw,omitempty"`
}

// NewVideoSpatialDirectorTimelineInner instantiates a new VideoSpatialDirectorTimelineInner object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewVideoSpatialDirectorTimelineInner() *VideoSpatialDirectorTimelineInner {
	this := VideoSpatialDirectorTimelineInner{}
	return &this
}

// NewVideoSpatialDirectorTimelineInnerWithDefaults instantiates a new VideoSpatialDirectorTimelineInner object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewVideoSpatialDirectorTimelineInnerWithDefaults() *VideoSpatialDirectorTimelineInner {
	this := VideoSpatialDirectorTimelineInner{}
	return &this
}

// GetPitch returns the Pitch field value if set, zero value otherwise.
func (o *VideoSpatialDirectorTimelineInner) GetPitch() float32 {
	if o == nil || IsNil(o.Pitch) {
		var ret float32
		return ret
	}
	return *o.Pitch
}

// GetPitchOk returns a tuple with the Pitch field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *VideoSpatialDirectorTimelineInner) GetPitchOk() (*float32, bool) {
	if o == nil || IsNil(o.Pitch) {
		return nil, false
	}
	return o.Pitch, true
}

// HasPitch returns a boolean if a field has been set.
func (o *VideoSpatialDirectorTimelineInner) HasPitch() bool {
	if o != nil && !IsNil(o.Pitch) {
		return true
	}

	return false
}

// SetPitch gets a reference to the given float32 and assigns it to the Pitch field.
func (o *VideoSpatialDirectorTimelineInner) SetPitch(v float32) {
	o.Pitch = &v
}

// GetRoll returns the Roll field value if set, zero value otherwise.
func (o *VideoSpatialDirectorTimelineInner) GetRoll() float32 {
	if o == nil || IsNil(o.Roll) {
		var ret float32
		return ret
	}
	return *o.Roll
}

// GetRollOk returns a tuple with the Roll field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *VideoSpatialDirectorTimelineInner) GetRollOk() (*float32, bool) {
	if o == nil || IsNil(o.Roll) {
		return nil, false
	}
	return o.Roll, true
}

// HasRoll returns a boolean if a field has been set.
func (o *VideoSpatialDirectorTimelineInner) HasRoll() bool {
	if o != nil && !IsNil(o.Roll) {
		return true
	}

	return false
}

// SetRoll gets a reference to the given float32 and assigns it to the Roll field.
func (o *VideoSpatialDirectorTimelineInner) SetRoll(v float32) {
	o.Roll = &v
}

// GetTimeCode returns the TimeCode field value if set, zero value otherwise.
func (o *VideoSpatialDirectorTimelineInner) GetTimeCode() float32 {
	if o == nil || IsNil(o.TimeCode) {
		var ret float32
		return ret
	}
	return *o.TimeCode
}

// GetTimeCodeOk returns a tuple with the TimeCode field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *VideoSpatialDirectorTimelineInner) GetTimeCodeOk() (*float32, bool) {
	if o == nil || IsNil(o.TimeCode) {
		return nil, false
	}
	return o.TimeCode, true
}

// HasTimeCode returns a boolean if a field has been set.
func (o *VideoSpatialDirectorTimelineInner) HasTimeCode() bool {
	if o != nil && !IsNil(o.TimeCode) {
		return true
	}

	return false
}

// SetTimeCode gets a reference to the given float32 and assigns it to the TimeCode field.
func (o *VideoSpatialDirectorTimelineInner) SetTimeCode(v float32) {
	o.TimeCode = &v
}

// GetYaw returns the Yaw field value if set, zero value otherwise.
func (o *VideoSpatialDirectorTimelineInner) GetYaw() float32 {
	if o == nil || IsNil(o.Yaw) {
		var ret float32
		return ret
	}
	return *o.Yaw
}

// GetYawOk returns a tuple with the Yaw field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *VideoSpatialDirectorTimelineInner) GetYawOk() (*float32, bool) {
	if o == nil || IsNil(o.Yaw) {
		return nil, false
	}
	return o.Yaw, true
}

// HasYaw returns a boolean if a field has been set.
func (o *VideoSpatialDirectorTimelineInner) HasYaw() bool {
	if o != nil && !IsNil(o.Yaw) {
		return true
	}

	return false
}

// SetYaw gets a reference to the given float32 and assigns it to the Yaw field.
func (o *VideoSpatialDirectorTimelineInner) SetYaw(v float32) {
	o.Yaw = &v
}

func (o VideoSpatialDirectorTimelineInner) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o VideoSpatialDirectorTimelineInner) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Pitch) {
		toSerialize["pitch"] = o.Pitch
	}
	if !IsNil(o.Roll) {
		toSerialize["roll"] = o.Roll
	}
	if !IsNil(o.TimeCode) {
		toSerialize["time_code"] = o.TimeCode
	}
	if !IsNil(o.Yaw) {
		toSerialize["yaw"] = o.Yaw
	}
	return toSerialize, nil
}

type NullableVideoSpatialDirectorTimelineInner struct {
	value *VideoSpatialDirectorTimelineInner
	isSet bool
}

func (v NullableVideoSpatialDirectorTimelineInner) Get() *VideoSpatialDirectorTimelineInner {
	return v.value
}

func (v *NullableVideoSpatialDirectorTimelineInner) Set(val *VideoSpatialDirectorTimelineInner) {
	v.value = val
	v.isSet = true
}

func (v NullableVideoSpatialDirectorTimelineInner) IsSet() bool {
	return v.isSet
}

func (v *NullableVideoSpatialDirectorTimelineInner) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableVideoSpatialDirectorTimelineInner(val *VideoSpatialDirectorTimelineInner) *NullableVideoSpatialDirectorTimelineInner {
	return &NullableVideoSpatialDirectorTimelineInner{value: val, isSet: true}
}

func (v NullableVideoSpatialDirectorTimelineInner) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableVideoSpatialDirectorTimelineInner) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

