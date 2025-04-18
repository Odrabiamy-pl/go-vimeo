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

// checks if the EditVideoRequestSpatialDirectorTimelineInner type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &EditVideoRequestSpatialDirectorTimelineInner{}

// EditVideoRequestSpatialDirectorTimelineInner struct for EditVideoRequestSpatialDirectorTimelineInner
type EditVideoRequestSpatialDirectorTimelineInner struct {
	// The 360 director timeline pitch. This value must be between −90 and 90, and you must specify it only when **spatial.director_timeline** is defined.
	Pitch float32 `json:"pitch"`
	// The 360 director timeline roll.
	Roll *float32 `json:"roll,omitempty"`
	// The 360 director timeline time code. This paramater is required only when **spatial.director_timeline** is defined.
	TimeCode float32 `json:"time_code"`
	// The 360 director timeline yaw. This value must be between 0 and 360, and you must specify it only when **spatial.director_timeline** is defined.
	Yaw float32 `json:"yaw"`
}

// NewEditVideoRequestSpatialDirectorTimelineInner instantiates a new EditVideoRequestSpatialDirectorTimelineInner object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewEditVideoRequestSpatialDirectorTimelineInner(pitch float32, timeCode float32, yaw float32) *EditVideoRequestSpatialDirectorTimelineInner {
	this := EditVideoRequestSpatialDirectorTimelineInner{}
	this.Pitch = pitch
	this.TimeCode = timeCode
	this.Yaw = yaw
	return &this
}

// NewEditVideoRequestSpatialDirectorTimelineInnerWithDefaults instantiates a new EditVideoRequestSpatialDirectorTimelineInner object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewEditVideoRequestSpatialDirectorTimelineInnerWithDefaults() *EditVideoRequestSpatialDirectorTimelineInner {
	this := EditVideoRequestSpatialDirectorTimelineInner{}
	return &this
}

// GetPitch returns the Pitch field value
func (o *EditVideoRequestSpatialDirectorTimelineInner) GetPitch() float32 {
	if o == nil {
		var ret float32
		return ret
	}

	return o.Pitch
}

// GetPitchOk returns a tuple with the Pitch field value
// and a boolean to check if the value has been set.
func (o *EditVideoRequestSpatialDirectorTimelineInner) GetPitchOk() (*float32, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Pitch, true
}

// SetPitch sets field value
func (o *EditVideoRequestSpatialDirectorTimelineInner) SetPitch(v float32) {
	o.Pitch = v
}

// GetRoll returns the Roll field value if set, zero value otherwise.
func (o *EditVideoRequestSpatialDirectorTimelineInner) GetRoll() float32 {
	if o == nil || IsNil(o.Roll) {
		var ret float32
		return ret
	}
	return *o.Roll
}

// GetRollOk returns a tuple with the Roll field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *EditVideoRequestSpatialDirectorTimelineInner) GetRollOk() (*float32, bool) {
	if o == nil || IsNil(o.Roll) {
		return nil, false
	}
	return o.Roll, true
}

// HasRoll returns a boolean if a field has been set.
func (o *EditVideoRequestSpatialDirectorTimelineInner) HasRoll() bool {
	if o != nil && !IsNil(o.Roll) {
		return true
	}

	return false
}

// SetRoll gets a reference to the given float32 and assigns it to the Roll field.
func (o *EditVideoRequestSpatialDirectorTimelineInner) SetRoll(v float32) {
	o.Roll = &v
}

// GetTimeCode returns the TimeCode field value
func (o *EditVideoRequestSpatialDirectorTimelineInner) GetTimeCode() float32 {
	if o == nil {
		var ret float32
		return ret
	}

	return o.TimeCode
}

// GetTimeCodeOk returns a tuple with the TimeCode field value
// and a boolean to check if the value has been set.
func (o *EditVideoRequestSpatialDirectorTimelineInner) GetTimeCodeOk() (*float32, bool) {
	if o == nil {
		return nil, false
	}
	return &o.TimeCode, true
}

// SetTimeCode sets field value
func (o *EditVideoRequestSpatialDirectorTimelineInner) SetTimeCode(v float32) {
	o.TimeCode = v
}

// GetYaw returns the Yaw field value
func (o *EditVideoRequestSpatialDirectorTimelineInner) GetYaw() float32 {
	if o == nil {
		var ret float32
		return ret
	}

	return o.Yaw
}

// GetYawOk returns a tuple with the Yaw field value
// and a boolean to check if the value has been set.
func (o *EditVideoRequestSpatialDirectorTimelineInner) GetYawOk() (*float32, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Yaw, true
}

// SetYaw sets field value
func (o *EditVideoRequestSpatialDirectorTimelineInner) SetYaw(v float32) {
	o.Yaw = v
}

func (o EditVideoRequestSpatialDirectorTimelineInner) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o EditVideoRequestSpatialDirectorTimelineInner) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["pitch"] = o.Pitch
	if !IsNil(o.Roll) {
		toSerialize["roll"] = o.Roll
	}
	toSerialize["time_code"] = o.TimeCode
	toSerialize["yaw"] = o.Yaw
	return toSerialize, nil
}

type NullableEditVideoRequestSpatialDirectorTimelineInner struct {
	value *EditVideoRequestSpatialDirectorTimelineInner
	isSet bool
}

func (v NullableEditVideoRequestSpatialDirectorTimelineInner) Get() *EditVideoRequestSpatialDirectorTimelineInner {
	return v.value
}

func (v *NullableEditVideoRequestSpatialDirectorTimelineInner) Set(val *EditVideoRequestSpatialDirectorTimelineInner) {
	v.value = val
	v.isSet = true
}

func (v NullableEditVideoRequestSpatialDirectorTimelineInner) IsSet() bool {
	return v.isSet
}

func (v *NullableEditVideoRequestSpatialDirectorTimelineInner) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableEditVideoRequestSpatialDirectorTimelineInner(val *EditVideoRequestSpatialDirectorTimelineInner) *NullableEditVideoRequestSpatialDirectorTimelineInner {
	return &NullableEditVideoRequestSpatialDirectorTimelineInner{value: val, isSet: true}
}

func (v NullableEditVideoRequestSpatialDirectorTimelineInner) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableEditVideoRequestSpatialDirectorTimelineInner) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


