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

// checks if the VideoSpatial type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &VideoSpatial{}

// VideoSpatial The video's 360 spatial data.
type VideoSpatial struct {
	// The video's 360 director timeline.
	DirectorTimeline []VideoSpatialDirectorTimelineInner `json:"director_timeline"`
	// The video's 360 field of view value, ranging from a mininum of `30` to a maximum of `90`. The default value is `50`.
	FieldOfView NullableFloat32 `json:"field_of_view"`
	// The video's 360 spatial projection.  Option descriptions:  * `cubical` - The spatial projection is cubical.  * `cylindrical` - The spatial projection is cylindrical.  * `dome` - The spatial projection is dome-shaped.  * `equirectangular` - The spatial projection is equirectangular.  * `pyramid` - The spatial projection is pyramid-shaped. 
	Projection NullableString `json:"projection"`
	// The video's 360 stereo format.  Option descriptions:  * `left-right` - The stereo format is left-right.  * `mono` - The audio is monaural.  * `top-bottom` - The stereo format is top-bottom. 
	StereoFormat NullableString `json:"stereo_format"`
}

// NewVideoSpatial instantiates a new VideoSpatial object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewVideoSpatial(directorTimeline []VideoSpatialDirectorTimelineInner, fieldOfView NullableFloat32, projection NullableString, stereoFormat NullableString) *VideoSpatial {
	this := VideoSpatial{}
	this.DirectorTimeline = directorTimeline
	this.FieldOfView = fieldOfView
	this.Projection = projection
	this.StereoFormat = stereoFormat
	return &this
}

// NewVideoSpatialWithDefaults instantiates a new VideoSpatial object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewVideoSpatialWithDefaults() *VideoSpatial {
	this := VideoSpatial{}
	return &this
}

// GetDirectorTimeline returns the DirectorTimeline field value
func (o *VideoSpatial) GetDirectorTimeline() []VideoSpatialDirectorTimelineInner {
	if o == nil {
		var ret []VideoSpatialDirectorTimelineInner
		return ret
	}

	return o.DirectorTimeline
}

// GetDirectorTimelineOk returns a tuple with the DirectorTimeline field value
// and a boolean to check if the value has been set.
func (o *VideoSpatial) GetDirectorTimelineOk() ([]VideoSpatialDirectorTimelineInner, bool) {
	if o == nil {
		return nil, false
	}
	return o.DirectorTimeline, true
}

// SetDirectorTimeline sets field value
func (o *VideoSpatial) SetDirectorTimeline(v []VideoSpatialDirectorTimelineInner) {
	o.DirectorTimeline = v
}

// GetFieldOfView returns the FieldOfView field value
// If the value is explicit nil, the zero value for float32 will be returned
func (o *VideoSpatial) GetFieldOfView() float32 {
	if o == nil || o.FieldOfView.Get() == nil {
		var ret float32
		return ret
	}

	return *o.FieldOfView.Get()
}

// GetFieldOfViewOk returns a tuple with the FieldOfView field value
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *VideoSpatial) GetFieldOfViewOk() (*float32, bool) {
	if o == nil {
		return nil, false
	}
	return o.FieldOfView.Get(), o.FieldOfView.IsSet()
}

// SetFieldOfView sets field value
func (o *VideoSpatial) SetFieldOfView(v float32) {
	o.FieldOfView.Set(&v)
}

// GetProjection returns the Projection field value
// If the value is explicit nil, the zero value for string will be returned
func (o *VideoSpatial) GetProjection() string {
	if o == nil || o.Projection.Get() == nil {
		var ret string
		return ret
	}

	return *o.Projection.Get()
}

// GetProjectionOk returns a tuple with the Projection field value
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *VideoSpatial) GetProjectionOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.Projection.Get(), o.Projection.IsSet()
}

// SetProjection sets field value
func (o *VideoSpatial) SetProjection(v string) {
	o.Projection.Set(&v)
}

// GetStereoFormat returns the StereoFormat field value
// If the value is explicit nil, the zero value for string will be returned
func (o *VideoSpatial) GetStereoFormat() string {
	if o == nil || o.StereoFormat.Get() == nil {
		var ret string
		return ret
	}

	return *o.StereoFormat.Get()
}

// GetStereoFormatOk returns a tuple with the StereoFormat field value
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *VideoSpatial) GetStereoFormatOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.StereoFormat.Get(), o.StereoFormat.IsSet()
}

// SetStereoFormat sets field value
func (o *VideoSpatial) SetStereoFormat(v string) {
	o.StereoFormat.Set(&v)
}

func (o VideoSpatial) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o VideoSpatial) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["director_timeline"] = o.DirectorTimeline
	toSerialize["field_of_view"] = o.FieldOfView.Get()
	toSerialize["projection"] = o.Projection.Get()
	toSerialize["stereo_format"] = o.StereoFormat.Get()
	return toSerialize, nil
}

type NullableVideoSpatial struct {
	value *VideoSpatial
	isSet bool
}

func (v NullableVideoSpatial) Get() *VideoSpatial {
	return v.value
}

func (v *NullableVideoSpatial) Set(val *VideoSpatial) {
	v.value = val
	v.isSet = true
}

func (v NullableVideoSpatial) IsSet() bool {
	return v.isSet
}

func (v *NullableVideoSpatial) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableVideoSpatial(val *VideoSpatial) *NullableVideoSpatial {
	return &NullableVideoSpatial{value: val, isSet: true}
}

func (v NullableVideoSpatial) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableVideoSpatial) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


