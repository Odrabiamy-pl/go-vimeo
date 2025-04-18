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

// checks if the PlayProgressiveInner type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &PlayProgressiveInner{}

// PlayProgressiveInner struct for PlayProgressiveInner
type PlayProgressiveInner struct {
	// The codec of the video file.  Option descriptions:  * `AV1` - The codec is AV1.  * `H264` - The codec is H264.  * `HEVC` - The codec is HEVC. 
	Codec NullableString `json:"codec"`
	// The time in ISO 8601 format when the video file was created.
	CreatedTime string `json:"created_time"`
	// The frames per second of the video.
	Fps float32 `json:"fps"`
	// The height of the video in pixels.
	Height NullableFloat32 `json:"height"`
	// The direct link to the video file.
	Link string `json:"link"`
	// The time in ISO 8601 format when the link to the video file expires.
	LinkExpirationTime string `json:"link_expiration_time"`
	// The URLs for logging events.
	Log map[string]interface{} `json:"log,omitempty"`
	// The MD5 hash of the video file.
	Md5 string `json:"md5"`
	// The video rendition.  Option descriptions:  * `1080p` - The video has 1080p resolution.  * `240p` - The video has 240p resolution.  * `2k` - The video has 2K resolution.  * `360p` - The video has 360p resolution.  * `480p` - The video has 480p resolution.  * `4k` - The video has 4K resolution.  * `540p` - The video has 540p resolution.  * `5k` - The video has 5K resolution.  * `6k` - The video has 6K resolution.  * `720p` - The video has 720p resolution.  * `7k` - The video has 7K resolution.  * `8k` - The video has 8K resolution. 
	Rendition string `json:"rendition"`
	// The size in bytes of the video file.
	Size NullableFloat32 `json:"size"`
	// The type of video file.  Option descriptions:  * `source` - The video file is a source file.  * `video/mp4` - The video file is in MP4 format.  * `video/webm` - The video file is in WebM format.  * `vp6/x-video` - The video file is in VP6 format. 
	Type NullableString `json:"type"`
	// The width of the video in pixels.
	Width NullableFloat32 `json:"width"`
}

// NewPlayProgressiveInner instantiates a new PlayProgressiveInner object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewPlayProgressiveInner(codec NullableString, createdTime string, fps float32, height NullableFloat32, link string, linkExpirationTime string, md5 string, rendition string, size NullableFloat32, type_ NullableString, width NullableFloat32) *PlayProgressiveInner {
	this := PlayProgressiveInner{}
	this.Codec = codec
	this.CreatedTime = createdTime
	this.Fps = fps
	this.Height = height
	this.Link = link
	this.LinkExpirationTime = linkExpirationTime
	this.Md5 = md5
	this.Rendition = rendition
	this.Size = size
	this.Type = type_
	this.Width = width
	return &this
}

// NewPlayProgressiveInnerWithDefaults instantiates a new PlayProgressiveInner object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewPlayProgressiveInnerWithDefaults() *PlayProgressiveInner {
	this := PlayProgressiveInner{}
	return &this
}

// GetCodec returns the Codec field value
// If the value is explicit nil, the zero value for string will be returned
func (o *PlayProgressiveInner) GetCodec() string {
	if o == nil || o.Codec.Get() == nil {
		var ret string
		return ret
	}

	return *o.Codec.Get()
}

// GetCodecOk returns a tuple with the Codec field value
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *PlayProgressiveInner) GetCodecOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.Codec.Get(), o.Codec.IsSet()
}

// SetCodec sets field value
func (o *PlayProgressiveInner) SetCodec(v string) {
	o.Codec.Set(&v)
}

// GetCreatedTime returns the CreatedTime field value
func (o *PlayProgressiveInner) GetCreatedTime() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.CreatedTime
}

// GetCreatedTimeOk returns a tuple with the CreatedTime field value
// and a boolean to check if the value has been set.
func (o *PlayProgressiveInner) GetCreatedTimeOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.CreatedTime, true
}

// SetCreatedTime sets field value
func (o *PlayProgressiveInner) SetCreatedTime(v string) {
	o.CreatedTime = v
}

// GetFps returns the Fps field value
func (o *PlayProgressiveInner) GetFps() float32 {
	if o == nil {
		var ret float32
		return ret
	}

	return o.Fps
}

// GetFpsOk returns a tuple with the Fps field value
// and a boolean to check if the value has been set.
func (o *PlayProgressiveInner) GetFpsOk() (*float32, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Fps, true
}

// SetFps sets field value
func (o *PlayProgressiveInner) SetFps(v float32) {
	o.Fps = v
}

// GetHeight returns the Height field value
// If the value is explicit nil, the zero value for float32 will be returned
func (o *PlayProgressiveInner) GetHeight() float32 {
	if o == nil || o.Height.Get() == nil {
		var ret float32
		return ret
	}

	return *o.Height.Get()
}

// GetHeightOk returns a tuple with the Height field value
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *PlayProgressiveInner) GetHeightOk() (*float32, bool) {
	if o == nil {
		return nil, false
	}
	return o.Height.Get(), o.Height.IsSet()
}

// SetHeight sets field value
func (o *PlayProgressiveInner) SetHeight(v float32) {
	o.Height.Set(&v)
}

// GetLink returns the Link field value
func (o *PlayProgressiveInner) GetLink() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Link
}

// GetLinkOk returns a tuple with the Link field value
// and a boolean to check if the value has been set.
func (o *PlayProgressiveInner) GetLinkOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Link, true
}

// SetLink sets field value
func (o *PlayProgressiveInner) SetLink(v string) {
	o.Link = v
}

// GetLinkExpirationTime returns the LinkExpirationTime field value
func (o *PlayProgressiveInner) GetLinkExpirationTime() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.LinkExpirationTime
}

// GetLinkExpirationTimeOk returns a tuple with the LinkExpirationTime field value
// and a boolean to check if the value has been set.
func (o *PlayProgressiveInner) GetLinkExpirationTimeOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.LinkExpirationTime, true
}

// SetLinkExpirationTime sets field value
func (o *PlayProgressiveInner) SetLinkExpirationTime(v string) {
	o.LinkExpirationTime = v
}

// GetLog returns the Log field value if set, zero value otherwise.
func (o *PlayProgressiveInner) GetLog() map[string]interface{} {
	if o == nil || IsNil(o.Log) {
		var ret map[string]interface{}
		return ret
	}
	return o.Log
}

// GetLogOk returns a tuple with the Log field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PlayProgressiveInner) GetLogOk() (map[string]interface{}, bool) {
	if o == nil || IsNil(o.Log) {
		return map[string]interface{}{}, false
	}
	return o.Log, true
}

// HasLog returns a boolean if a field has been set.
func (o *PlayProgressiveInner) HasLog() bool {
	if o != nil && !IsNil(o.Log) {
		return true
	}

	return false
}

// SetLog gets a reference to the given map[string]interface{} and assigns it to the Log field.
func (o *PlayProgressiveInner) SetLog(v map[string]interface{}) {
	o.Log = v
}

// GetMd5 returns the Md5 field value
func (o *PlayProgressiveInner) GetMd5() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Md5
}

// GetMd5Ok returns a tuple with the Md5 field value
// and a boolean to check if the value has been set.
func (o *PlayProgressiveInner) GetMd5Ok() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Md5, true
}

// SetMd5 sets field value
func (o *PlayProgressiveInner) SetMd5(v string) {
	o.Md5 = v
}

// GetRendition returns the Rendition field value
func (o *PlayProgressiveInner) GetRendition() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Rendition
}

// GetRenditionOk returns a tuple with the Rendition field value
// and a boolean to check if the value has been set.
func (o *PlayProgressiveInner) GetRenditionOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Rendition, true
}

// SetRendition sets field value
func (o *PlayProgressiveInner) SetRendition(v string) {
	o.Rendition = v
}

// GetSize returns the Size field value
// If the value is explicit nil, the zero value for float32 will be returned
func (o *PlayProgressiveInner) GetSize() float32 {
	if o == nil || o.Size.Get() == nil {
		var ret float32
		return ret
	}

	return *o.Size.Get()
}

// GetSizeOk returns a tuple with the Size field value
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *PlayProgressiveInner) GetSizeOk() (*float32, bool) {
	if o == nil {
		return nil, false
	}
	return o.Size.Get(), o.Size.IsSet()
}

// SetSize sets field value
func (o *PlayProgressiveInner) SetSize(v float32) {
	o.Size.Set(&v)
}

// GetType returns the Type field value
// If the value is explicit nil, the zero value for string will be returned
func (o *PlayProgressiveInner) GetType() string {
	if o == nil || o.Type.Get() == nil {
		var ret string
		return ret
	}

	return *o.Type.Get()
}

// GetTypeOk returns a tuple with the Type field value
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *PlayProgressiveInner) GetTypeOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.Type.Get(), o.Type.IsSet()
}

// SetType sets field value
func (o *PlayProgressiveInner) SetType(v string) {
	o.Type.Set(&v)
}

// GetWidth returns the Width field value
// If the value is explicit nil, the zero value for float32 will be returned
func (o *PlayProgressiveInner) GetWidth() float32 {
	if o == nil || o.Width.Get() == nil {
		var ret float32
		return ret
	}

	return *o.Width.Get()
}

// GetWidthOk returns a tuple with the Width field value
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *PlayProgressiveInner) GetWidthOk() (*float32, bool) {
	if o == nil {
		return nil, false
	}
	return o.Width.Get(), o.Width.IsSet()
}

// SetWidth sets field value
func (o *PlayProgressiveInner) SetWidth(v float32) {
	o.Width.Set(&v)
}

func (o PlayProgressiveInner) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o PlayProgressiveInner) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["codec"] = o.Codec.Get()
	toSerialize["created_time"] = o.CreatedTime
	toSerialize["fps"] = o.Fps
	toSerialize["height"] = o.Height.Get()
	toSerialize["link"] = o.Link
	toSerialize["link_expiration_time"] = o.LinkExpirationTime
	if !IsNil(o.Log) {
		toSerialize["log"] = o.Log
	}
	toSerialize["md5"] = o.Md5
	toSerialize["rendition"] = o.Rendition
	toSerialize["size"] = o.Size.Get()
	toSerialize["type"] = o.Type.Get()
	toSerialize["width"] = o.Width.Get()
	return toSerialize, nil
}

type NullablePlayProgressiveInner struct {
	value *PlayProgressiveInner
	isSet bool
}

func (v NullablePlayProgressiveInner) Get() *PlayProgressiveInner {
	return v.value
}

func (v *NullablePlayProgressiveInner) Set(val *PlayProgressiveInner) {
	v.value = val
	v.isSet = true
}

func (v NullablePlayProgressiveInner) IsSet() bool {
	return v.isSet
}

func (v *NullablePlayProgressiveInner) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullablePlayProgressiveInner(val *PlayProgressiveInner) *NullablePlayProgressiveInner {
	return &NullablePlayProgressiveInner{value: val, isSet: true}
}

func (v NullablePlayProgressiveInner) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullablePlayProgressiveInner) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


