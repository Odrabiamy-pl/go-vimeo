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

// checks if the UploadVideoAlt1RequestSpatialDirectorTimelineInner type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &UploadVideoAlt1RequestSpatialDirectorTimelineInner{}

// UploadVideoAlt1RequestSpatialDirectorTimelineInner struct for UploadVideoAlt1RequestSpatialDirectorTimelineInner
type UploadVideoAlt1RequestSpatialDirectorTimelineInner struct {
	// The 360 director timeline pitch. This value must be between `−90` and `90`, and it's required only when **spatial.director_timeline** is defined.
	Pitch float32 `json:"pitch"`
	// The 360 director timeline roll.
	Roll *float32 `json:"roll,omitempty"`
	// The 360 director timeline time code. This field is required only when **spatial.director_timeline** is defined.
	TimeCode float32 `json:"time_code"`
	// The 360 director timeline yaw. This value must be between `0` and `360`, and it's required only when **spatial.director_timeline** is defined.
	Yaw float32 `json:"yaw"`
}

// NewUploadVideoAlt1RequestSpatialDirectorTimelineInner instantiates a new UploadVideoAlt1RequestSpatialDirectorTimelineInner object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewUploadVideoAlt1RequestSpatialDirectorTimelineInner(pitch float32, timeCode float32, yaw float32) *UploadVideoAlt1RequestSpatialDirectorTimelineInner {
	this := UploadVideoAlt1RequestSpatialDirectorTimelineInner{}
	this.Pitch = pitch
	this.TimeCode = timeCode
	this.Yaw = yaw
	return &this
}

// NewUploadVideoAlt1RequestSpatialDirectorTimelineInnerWithDefaults instantiates a new UploadVideoAlt1RequestSpatialDirectorTimelineInner object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewUploadVideoAlt1RequestSpatialDirectorTimelineInnerWithDefaults() *UploadVideoAlt1RequestSpatialDirectorTimelineInner {
	this := UploadVideoAlt1RequestSpatialDirectorTimelineInner{}
	return &this
}

// GetPitch returns the Pitch field value
func (o *UploadVideoAlt1RequestSpatialDirectorTimelineInner) GetPitch() float32 {
	if o == nil {
		var ret float32
		return ret
	}

	return o.Pitch
}

// GetPitchOk returns a tuple with the Pitch field value
// and a boolean to check if the value has been set.
func (o *UploadVideoAlt1RequestSpatialDirectorTimelineInner) GetPitchOk() (*float32, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Pitch, true
}

// SetPitch sets field value
func (o *UploadVideoAlt1RequestSpatialDirectorTimelineInner) SetPitch(v float32) {
	o.Pitch = v
}

// GetRoll returns the Roll field value if set, zero value otherwise.
func (o *UploadVideoAlt1RequestSpatialDirectorTimelineInner) GetRoll() float32 {
	if o == nil || IsNil(o.Roll) {
		var ret float32
		return ret
	}
	return *o.Roll
}

// GetRollOk returns a tuple with the Roll field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UploadVideoAlt1RequestSpatialDirectorTimelineInner) GetRollOk() (*float32, bool) {
	if o == nil || IsNil(o.Roll) {
		return nil, false
	}
	return o.Roll, true
}

// HasRoll returns a boolean if a field has been set.
func (o *UploadVideoAlt1RequestSpatialDirectorTimelineInner) HasRoll() bool {
	if o != nil && !IsNil(o.Roll) {
		return true
	}

	return false
}

// SetRoll gets a reference to the given float32 and assigns it to the Roll field.
func (o *UploadVideoAlt1RequestSpatialDirectorTimelineInner) SetRoll(v float32) {
	o.Roll = &v
}

// GetTimeCode returns the TimeCode field value
func (o *UploadVideoAlt1RequestSpatialDirectorTimelineInner) GetTimeCode() float32 {
	if o == nil {
		var ret float32
		return ret
	}

	return o.TimeCode
}

// GetTimeCodeOk returns a tuple with the TimeCode field value
// and a boolean to check if the value has been set.
func (o *UploadVideoAlt1RequestSpatialDirectorTimelineInner) GetTimeCodeOk() (*float32, bool) {
	if o == nil {
		return nil, false
	}
	return &o.TimeCode, true
}

// SetTimeCode sets field value
func (o *UploadVideoAlt1RequestSpatialDirectorTimelineInner) SetTimeCode(v float32) {
	o.TimeCode = v
}

// GetYaw returns the Yaw field value
func (o *UploadVideoAlt1RequestSpatialDirectorTimelineInner) GetYaw() float32 {
	if o == nil {
		var ret float32
		return ret
	}

	return o.Yaw
}

// GetYawOk returns a tuple with the Yaw field value
// and a boolean to check if the value has been set.
func (o *UploadVideoAlt1RequestSpatialDirectorTimelineInner) GetYawOk() (*float32, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Yaw, true
}

// SetYaw sets field value
func (o *UploadVideoAlt1RequestSpatialDirectorTimelineInner) SetYaw(v float32) {
	o.Yaw = v
}

func (o UploadVideoAlt1RequestSpatialDirectorTimelineInner) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o UploadVideoAlt1RequestSpatialDirectorTimelineInner) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["pitch"] = o.Pitch
	if !IsNil(o.Roll) {
		toSerialize["roll"] = o.Roll
	}
	toSerialize["time_code"] = o.TimeCode
	toSerialize["yaw"] = o.Yaw
	return toSerialize, nil
}

type NullableUploadVideoAlt1RequestSpatialDirectorTimelineInner struct {
	value *UploadVideoAlt1RequestSpatialDirectorTimelineInner
	isSet bool
}

func (v NullableUploadVideoAlt1RequestSpatialDirectorTimelineInner) Get() *UploadVideoAlt1RequestSpatialDirectorTimelineInner {
	return v.value
}

func (v *NullableUploadVideoAlt1RequestSpatialDirectorTimelineInner) Set(val *UploadVideoAlt1RequestSpatialDirectorTimelineInner) {
	v.value = val
	v.isSet = true
}

func (v NullableUploadVideoAlt1RequestSpatialDirectorTimelineInner) IsSet() bool {
	return v.isSet
}

func (v *NullableUploadVideoAlt1RequestSpatialDirectorTimelineInner) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableUploadVideoAlt1RequestSpatialDirectorTimelineInner(val *UploadVideoAlt1RequestSpatialDirectorTimelineInner) *NullableUploadVideoAlt1RequestSpatialDirectorTimelineInner {
	return &NullableUploadVideoAlt1RequestSpatialDirectorTimelineInner{value: val, isSet: true}
}

func (v NullableUploadVideoAlt1RequestSpatialDirectorTimelineInner) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableUploadVideoAlt1RequestSpatialDirectorTimelineInner) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
