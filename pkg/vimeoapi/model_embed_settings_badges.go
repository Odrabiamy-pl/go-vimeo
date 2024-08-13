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

// checks if the EmbedSettingsBadges type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &EmbedSettingsBadges{}

// EmbedSettingsBadges A collection of the video's badges.
type EmbedSettingsBadges struct {
	// Whether the video was filmed using Dolby Vision.
	DolbyVision *bool `json:"dolby_vision,omitempty"`
	// Whether the video has an HDR-compatible transcode.
	Hdr bool `json:"hdr"`
	// Whether the video was filmed using HDR10.
	Hdr10 *bool `json:"hdr_10,omitempty"`
	// Whether the video was filmed using HDR10 Plus.
	Hdr10Plus *bool `json:"hdr_10_plus,omitempty"`
	Live EmbedSettingsBadgesLive `json:"live"`
	StaffPick EmbedSettingsBadgesStaffPick `json:"staff_pick"`
	// Whether the video is a Vimeo On Demand video.
	Vod bool `json:"vod"`
	// Whether the video is a Vimeo Weekend Challenge.
	WeekendChallenge bool `json:"weekend_challenge"`
}

// NewEmbedSettingsBadges instantiates a new EmbedSettingsBadges object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewEmbedSettingsBadges(hdr bool, live EmbedSettingsBadgesLive, staffPick EmbedSettingsBadgesStaffPick, vod bool, weekendChallenge bool) *EmbedSettingsBadges {
	this := EmbedSettingsBadges{}
	this.Hdr = hdr
	this.Live = live
	this.StaffPick = staffPick
	this.Vod = vod
	this.WeekendChallenge = weekendChallenge
	return &this
}

// NewEmbedSettingsBadgesWithDefaults instantiates a new EmbedSettingsBadges object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewEmbedSettingsBadgesWithDefaults() *EmbedSettingsBadges {
	this := EmbedSettingsBadges{}
	return &this
}

// GetDolbyVision returns the DolbyVision field value if set, zero value otherwise.
func (o *EmbedSettingsBadges) GetDolbyVision() bool {
	if o == nil || IsNil(o.DolbyVision) {
		var ret bool
		return ret
	}
	return *o.DolbyVision
}

// GetDolbyVisionOk returns a tuple with the DolbyVision field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *EmbedSettingsBadges) GetDolbyVisionOk() (*bool, bool) {
	if o == nil || IsNil(o.DolbyVision) {
		return nil, false
	}
	return o.DolbyVision, true
}

// HasDolbyVision returns a boolean if a field has been set.
func (o *EmbedSettingsBadges) HasDolbyVision() bool {
	if o != nil && !IsNil(o.DolbyVision) {
		return true
	}

	return false
}

// SetDolbyVision gets a reference to the given bool and assigns it to the DolbyVision field.
func (o *EmbedSettingsBadges) SetDolbyVision(v bool) {
	o.DolbyVision = &v
}

// GetHdr returns the Hdr field value
func (o *EmbedSettingsBadges) GetHdr() bool {
	if o == nil {
		var ret bool
		return ret
	}

	return o.Hdr
}

// GetHdrOk returns a tuple with the Hdr field value
// and a boolean to check if the value has been set.
func (o *EmbedSettingsBadges) GetHdrOk() (*bool, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Hdr, true
}

// SetHdr sets field value
func (o *EmbedSettingsBadges) SetHdr(v bool) {
	o.Hdr = v
}

// GetHdr10 returns the Hdr10 field value if set, zero value otherwise.
func (o *EmbedSettingsBadges) GetHdr10() bool {
	if o == nil || IsNil(o.Hdr10) {
		var ret bool
		return ret
	}
	return *o.Hdr10
}

// GetHdr10Ok returns a tuple with the Hdr10 field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *EmbedSettingsBadges) GetHdr10Ok() (*bool, bool) {
	if o == nil || IsNil(o.Hdr10) {
		return nil, false
	}
	return o.Hdr10, true
}

// HasHdr10 returns a boolean if a field has been set.
func (o *EmbedSettingsBadges) HasHdr10() bool {
	if o != nil && !IsNil(o.Hdr10) {
		return true
	}

	return false
}

// SetHdr10 gets a reference to the given bool and assigns it to the Hdr10 field.
func (o *EmbedSettingsBadges) SetHdr10(v bool) {
	o.Hdr10 = &v
}

// GetHdr10Plus returns the Hdr10Plus field value if set, zero value otherwise.
func (o *EmbedSettingsBadges) GetHdr10Plus() bool {
	if o == nil || IsNil(o.Hdr10Plus) {
		var ret bool
		return ret
	}
	return *o.Hdr10Plus
}

// GetHdr10PlusOk returns a tuple with the Hdr10Plus field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *EmbedSettingsBadges) GetHdr10PlusOk() (*bool, bool) {
	if o == nil || IsNil(o.Hdr10Plus) {
		return nil, false
	}
	return o.Hdr10Plus, true
}

// HasHdr10Plus returns a boolean if a field has been set.
func (o *EmbedSettingsBadges) HasHdr10Plus() bool {
	if o != nil && !IsNil(o.Hdr10Plus) {
		return true
	}

	return false
}

// SetHdr10Plus gets a reference to the given bool and assigns it to the Hdr10Plus field.
func (o *EmbedSettingsBadges) SetHdr10Plus(v bool) {
	o.Hdr10Plus = &v
}

// GetLive returns the Live field value
func (o *EmbedSettingsBadges) GetLive() EmbedSettingsBadgesLive {
	if o == nil {
		var ret EmbedSettingsBadgesLive
		return ret
	}

	return o.Live
}

// GetLiveOk returns a tuple with the Live field value
// and a boolean to check if the value has been set.
func (o *EmbedSettingsBadges) GetLiveOk() (*EmbedSettingsBadgesLive, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Live, true
}

// SetLive sets field value
func (o *EmbedSettingsBadges) SetLive(v EmbedSettingsBadgesLive) {
	o.Live = v
}

// GetStaffPick returns the StaffPick field value
func (o *EmbedSettingsBadges) GetStaffPick() EmbedSettingsBadgesStaffPick {
	if o == nil {
		var ret EmbedSettingsBadgesStaffPick
		return ret
	}

	return o.StaffPick
}

// GetStaffPickOk returns a tuple with the StaffPick field value
// and a boolean to check if the value has been set.
func (o *EmbedSettingsBadges) GetStaffPickOk() (*EmbedSettingsBadgesStaffPick, bool) {
	if o == nil {
		return nil, false
	}
	return &o.StaffPick, true
}

// SetStaffPick sets field value
func (o *EmbedSettingsBadges) SetStaffPick(v EmbedSettingsBadgesStaffPick) {
	o.StaffPick = v
}

// GetVod returns the Vod field value
func (o *EmbedSettingsBadges) GetVod() bool {
	if o == nil {
		var ret bool
		return ret
	}

	return o.Vod
}

// GetVodOk returns a tuple with the Vod field value
// and a boolean to check if the value has been set.
func (o *EmbedSettingsBadges) GetVodOk() (*bool, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Vod, true
}

// SetVod sets field value
func (o *EmbedSettingsBadges) SetVod(v bool) {
	o.Vod = v
}

// GetWeekendChallenge returns the WeekendChallenge field value
func (o *EmbedSettingsBadges) GetWeekendChallenge() bool {
	if o == nil {
		var ret bool
		return ret
	}

	return o.WeekendChallenge
}

// GetWeekendChallengeOk returns a tuple with the WeekendChallenge field value
// and a boolean to check if the value has been set.
func (o *EmbedSettingsBadges) GetWeekendChallengeOk() (*bool, bool) {
	if o == nil {
		return nil, false
	}
	return &o.WeekendChallenge, true
}

// SetWeekendChallenge sets field value
func (o *EmbedSettingsBadges) SetWeekendChallenge(v bool) {
	o.WeekendChallenge = v
}

func (o EmbedSettingsBadges) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o EmbedSettingsBadges) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.DolbyVision) {
		toSerialize["dolby_vision"] = o.DolbyVision
	}
	toSerialize["hdr"] = o.Hdr
	if !IsNil(o.Hdr10) {
		toSerialize["hdr_10"] = o.Hdr10
	}
	if !IsNil(o.Hdr10Plus) {
		toSerialize["hdr_10_plus"] = o.Hdr10Plus
	}
	toSerialize["live"] = o.Live
	toSerialize["staff_pick"] = o.StaffPick
	toSerialize["vod"] = o.Vod
	toSerialize["weekend_challenge"] = o.WeekendChallenge
	return toSerialize, nil
}

type NullableEmbedSettingsBadges struct {
	value *EmbedSettingsBadges
	isSet bool
}

func (v NullableEmbedSettingsBadges) Get() *EmbedSettingsBadges {
	return v.value
}

func (v *NullableEmbedSettingsBadges) Set(val *EmbedSettingsBadges) {
	v.value = val
	v.isSet = true
}

func (v NullableEmbedSettingsBadges) IsSet() bool {
	return v.isSet
}

func (v *NullableEmbedSettingsBadges) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableEmbedSettingsBadges(val *EmbedSettingsBadges) *NullableEmbedSettingsBadges {
	return &NullableEmbedSettingsBadges{value: val, isSet: true}
}

func (v NullableEmbedSettingsBadges) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableEmbedSettingsBadges) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


