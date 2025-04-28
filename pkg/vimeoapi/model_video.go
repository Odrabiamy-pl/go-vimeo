/*
Vimeo API

Build something great. Vimeo's API supports flexible, high-quality video integration with your custom apps.

API version: 3.4
*/

package vimeoapi

import (
	"encoding/json"
)

// checks if the Video type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &Video{}

// Video struct for Video
type Video struct {
	// The time in ISO 8601 format when the video was created.
	CreatedTime string `json:"created_time"`
	// The video's duration in seconds. A value of `0` indicates the duration hasn't been calculated yet.
	Duration float32 `json:"duration"`
	// The video's title.
	Name     string  `json:"name"`
	Pictures Picture `json:"pictures"`
	Play     *Play   `json:"play,omitempty"`
	// The video's player embed URL.
	PlayerEmbedUrl string `json:"player_embed_url"`
}

// NewVideo instantiates a new Video object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewVideo(createdTime string, duration float32, name string, pictures Picture, playerEmbedUrl string) *Video {
	this := Video{}
	this.CreatedTime = createdTime
	this.Duration = duration
	this.Name = name
	this.Pictures = pictures
	this.PlayerEmbedUrl = playerEmbedUrl
	return &this
}

// NewVideoWithDefaults instantiates a new Video object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewVideoWithDefaults() *Video {
	this := Video{}
	return &this
}

// GetCreatedTime returns the CreatedTime field value
func (o *Video) GetCreatedTime() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.CreatedTime
}

// GetCreatedTimeOk returns a tuple with the CreatedTime field value
// and a boolean to check if the value has been set.
func (o *Video) GetCreatedTimeOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.CreatedTime, true
}

// SetCreatedTime sets field value
func (o *Video) SetCreatedTime(v string) {
	o.CreatedTime = v
}

// GetDuration returns the Duration field value
func (o *Video) GetDuration() float32 {
	if o == nil {
		var ret float32
		return ret
	}

	return o.Duration
}

// GetDurationOk returns a tuple with the Duration field value
// and a boolean to check if the value has been set.
func (o *Video) GetDurationOk() (*float32, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Duration, true
}

// SetDuration sets field value
func (o *Video) SetDuration(v float32) {
	o.Duration = v
}

// GetName returns the Name field value
func (o *Video) GetName() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Name
}

// GetNameOk returns a tuple with the Name field value
// and a boolean to check if the value has been set.
func (o *Video) GetNameOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Name, true
}

// SetName sets field value
func (o *Video) SetName(v string) {
	o.Name = v
}

// GetPictures returns the Pictures field value
func (o *Video) GetPictures() Picture {
	if o == nil {
		var ret Picture
		return ret
	}

	return o.Pictures
}

// GetPicturesOk returns a tuple with the Pictures field value
// and a boolean to check if the value has been set.
func (o *Video) GetPicturesOk() (*Picture, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Pictures, true
}

// SetPictures sets field value
func (o *Video) SetPictures(v Picture) {
	o.Pictures = v
}

// GetPlay returns the Play field value if set, zero value otherwise.
func (o *Video) GetPlay() Play {
	if o == nil || IsNil(o.Play) {
		var ret Play
		return ret
	}
	return *o.Play
}

// GetPlayOk returns a tuple with the Play field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Video) GetPlayOk() (*Play, bool) {
	if o == nil || IsNil(o.Play) {
		return nil, false
	}
	return o.Play, true
}

// HasPlay returns a boolean if a field has been set.
func (o *Video) HasPlay() bool {
	if o != nil && !IsNil(o.Play) {
		return true
	}

	return false
}

// SetPlay gets a reference to the given Play and assigns it to the Play field.
func (o *Video) SetPlay(v Play) {
	o.Play = &v
}

// GetPlayerEmbedUrl returns the PlayerEmbedUrl field value
func (o *Video) GetPlayerEmbedUrl() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.PlayerEmbedUrl
}

// GetPlayerEmbedUrlOk returns a tuple with the PlayerEmbedUrl field value
// and a boolean to check if the value has been set.
func (o *Video) GetPlayerEmbedUrlOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.PlayerEmbedUrl, true
}

// SetPlayerEmbedUrl sets field value
func (o *Video) SetPlayerEmbedUrl(v string) {
	o.PlayerEmbedUrl = v
}

func (o Video) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o Video) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["created_time"] = o.CreatedTime
	toSerialize["duration"] = o.Duration
	toSerialize["name"] = o.Name
	toSerialize["pictures"] = o.Pictures
	if !IsNil(o.Play) {
		toSerialize["play"] = o.Play
	}
	toSerialize["player_embed_url"] = o.PlayerEmbedUrl
	return toSerialize, nil
}

type NullableVideo struct {
	value *Video
	isSet bool
}

func (v NullableVideo) Get() *Video {
	return v.value
}

func (v *NullableVideo) Set(val *Video) {
	v.value = val
	v.isSet = true
}

func (v NullableVideo) IsSet() bool {
	return v.isSet
}

func (v *NullableVideo) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableVideo(val *Video) *NullableVideo {
	return &NullableVideo{value: val, isSet: true}
}

func (v NullableVideo) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableVideo) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
