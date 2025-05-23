/*
Vimeo API

Build something great. Vimeo's API supports flexible, high-quality video integration with your custom apps.

API version: 3.4
*/

package vimeoapi

import (
	"encoding/json"
)

// checks if the Play type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &Play{}

// Play struct for Play
type Play struct {
	Dash *HlsDashVideoFile `json:"dash,omitempty"`
	Hls  *HlsDashVideoFile `json:"hls,omitempty"`
	// The play status of the video.  Option descriptions:  * `playable` - The video is playable.  * `purchase_required` - The video must be purchased.  * `restricted` - Playback for the video is restricted.  * `unavailable` - The video is unavailable.
	Status string `json:"status"`
}

// NewPlay instantiates a new Play object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewPlay(status string) *Play {
	this := Play{}
	this.Status = status
	return &this
}

// NewPlayWithDefaults instantiates a new Play object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewPlayWithDefaults() *Play {
	this := Play{}
	return &this
}

// GetDash returns the Dash field value if set, zero value otherwise.
func (o *Play) GetDash() HlsDashVideoFile {
	if o == nil || IsNil(o.Dash) {
		var ret HlsDashVideoFile
		return ret
	}
	return *o.Dash
}

// GetDashOk returns a tuple with the Dash field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Play) GetDashOk() (*HlsDashVideoFile, bool) {
	if o == nil || IsNil(o.Dash) {
		return nil, false
	}
	return o.Dash, true
}

// HasDash returns a boolean if a field has been set.
func (o *Play) HasDash() bool {
	if o != nil && !IsNil(o.Dash) {
		return true
	}

	return false
}

// SetDash gets a reference to the given HlsDashVideoFile and assigns it to the Dash field.
func (o *Play) SetDash(v HlsDashVideoFile) {
	o.Dash = &v
}

// GetHls returns the Hls field value if set, zero value otherwise.
func (o *Play) GetHls() HlsDashVideoFile {
	if o == nil || IsNil(o.Hls) {
		var ret HlsDashVideoFile
		return ret
	}
	return *o.Hls
}

// GetHlsOk returns a tuple with the Hls field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Play) GetHlsOk() (*HlsDashVideoFile, bool) {
	if o == nil || IsNil(o.Hls) {
		return nil, false
	}
	return o.Hls, true
}

// HasHls returns a boolean if a field has been set.
func (o *Play) HasHls() bool {
	if o != nil && !IsNil(o.Hls) {
		return true
	}

	return false
}

// SetHls gets a reference to the given HlsDashVideoFile and assigns it to the Hls field.
func (o *Play) SetHls(v HlsDashVideoFile) {
	o.Hls = &v
}

// GetStatus returns the Status field value
func (o *Play) GetStatus() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Status
}

// GetStatusOk returns a tuple with the Status field value
// and a boolean to check if the value has been set.
func (o *Play) GetStatusOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Status, true
}

// SetStatus sets field value
func (o *Play) SetStatus(v string) {
	o.Status = v
}

func (o Play) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o Play) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Dash) {
		toSerialize["dash"] = o.Dash
	}
	if !IsNil(o.Hls) {
		toSerialize["hls"] = o.Hls
	}
	toSerialize["status"] = o.Status
	return toSerialize, nil
}

type NullablePlay struct {
	value *Play
	isSet bool
}

func (v NullablePlay) Get() *Play {
	return v.value
}

func (v *NullablePlay) Set(val *Play) {
	v.value = val
	v.isSet = true
}

func (v NullablePlay) IsSet() bool {
	return v.isSet
}

func (v *NullablePlay) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullablePlay(val *Play) *NullablePlay {
	return &NullablePlay{value: val, isSet: true}
}

func (v NullablePlay) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullablePlay) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
