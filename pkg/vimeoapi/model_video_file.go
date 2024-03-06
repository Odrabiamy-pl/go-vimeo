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

// checks if the VideoFile type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &VideoFile{}

// VideoFile struct for VideoFile
type VideoFile struct {
	// The codec of the video file.  Option descriptions:  * `AV1` - The codec is AV1.  * `H264` - The codec is H264.  * `HEVC` - The codec is HEVC.
	Codec NullableString `json:"codec"`
	// The time in ISO 8601 format when the video file was created.
	CreatedTime string `json:"created_time"`
	// The time in ISO 8601 format when the video file expires.
	Expires *string `json:"expires,omitempty"`
	// The frames per second of the video.
	Fps float32 `json:"fps"`
	// The height of the video in pixels.
	Height NullableFloat32 `json:"height"`
	// The direct link to the video file.
	Link string        `json:"link"`
	Log  *VideoFileLog `json:"log,omitempty"`
	// The MD5 hash of the video file.
	Md5 string `json:"md5"`
	// The public name of the video file.
	PublicName string `json:"public_name"`
	// The video quality as determined by height and width.  Option descriptions:  * `hd` - The video is in high definition.  * `hls` - The video is suitable for HTTP live streaming.  * `mobile` - The video is mobile quality.  * `sd` - The video is in standard definition.  * `source` - The video's source file.  * `uhd` - The video resolution is 2K or higher.
	Quality string `json:"quality"`
	// The video rendition.  Option descriptions:  * `1080p` - The video has 1080p resolution.  * `240p` - The video has 240p resolution.  * `2k` - The video has 2K resolution.  * `360p` - The video has 360p resolution.  * `480p` - The video has 480p resolution.  * `4k` - The video has 4K resolution.  * `540p` - The video has 540p resolution.  * `5k` - The video has 5K resolution.  * `6k` - The video has 6K resolution.  * `720p` - The video has 720p resolution.  * `7k` - The video has 7K resolution.  * `8k` - The video has 8K resolution.  * `adaptive` - The video rendition is adaptive (for example, HLS or DASH).  * `source` - The video is the source file.
	Rendition string `json:"rendition"`
	// The approximate size in bytes of the video file.
	Size NullableFloat32 `json:"size"`
	// The converted size of the video file rounded to two decimal places.
	SizeShort string `json:"size_short"`
	// The source link of the video file.
	SourceLink NullableString `json:"source_link,omitempty"`
	// The type of video file.  Option descriptions:  * `source` - The video file is a source file.  * `video/mp4` - The video file is in MP4 format.  * `video/webm` - The video file is in WebM format.  * `vp6/x-video` - The video file is in VP6 format.
	Type NullableString `json:"type"`
	// The ID of the video file.
	VideoFileId *string `json:"video_file_id,omitempty"`
	// The width of the video in pixels.
	Width NullableFloat32 `json:"width"`
}

// NewVideoFile instantiates a new VideoFile object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewVideoFile(codec NullableString, createdTime string, fps float32, height NullableFloat32, link string, md5 string, publicName string, quality string, rendition string, size NullableFloat32, sizeShort string, type_ NullableString, width NullableFloat32) *VideoFile {
	this := VideoFile{}
	this.Codec = codec
	this.CreatedTime = createdTime
	this.Fps = fps
	this.Height = height
	this.Link = link
	this.Md5 = md5
	this.PublicName = publicName
	this.Quality = quality
	this.Rendition = rendition
	this.Size = size
	this.SizeShort = sizeShort
	this.Type = type_
	this.Width = width
	return &this
}

// NewVideoFileWithDefaults instantiates a new VideoFile object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewVideoFileWithDefaults() *VideoFile {
	this := VideoFile{}
	return &this
}

// GetCodec returns the Codec field value
// If the value is explicit nil, the zero value for string will be returned
func (o *VideoFile) GetCodec() string {
	if o == nil || o.Codec.Get() == nil {
		var ret string
		return ret
	}

	return *o.Codec.Get()
}

// GetCodecOk returns a tuple with the Codec field value
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *VideoFile) GetCodecOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.Codec.Get(), o.Codec.IsSet()
}

// SetCodec sets field value
func (o *VideoFile) SetCodec(v string) {
	o.Codec.Set(&v)
}

// GetCreatedTime returns the CreatedTime field value
func (o *VideoFile) GetCreatedTime() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.CreatedTime
}

// GetCreatedTimeOk returns a tuple with the CreatedTime field value
// and a boolean to check if the value has been set.
func (o *VideoFile) GetCreatedTimeOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.CreatedTime, true
}

// SetCreatedTime sets field value
func (o *VideoFile) SetCreatedTime(v string) {
	o.CreatedTime = v
}

// GetExpires returns the Expires field value if set, zero value otherwise.
func (o *VideoFile) GetExpires() string {
	if o == nil || IsNil(o.Expires) {
		var ret string
		return ret
	}
	return *o.Expires
}

// GetExpiresOk returns a tuple with the Expires field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *VideoFile) GetExpiresOk() (*string, bool) {
	if o == nil || IsNil(o.Expires) {
		return nil, false
	}
	return o.Expires, true
}

// HasExpires returns a boolean if a field has been set.
func (o *VideoFile) HasExpires() bool {
	if o != nil && !IsNil(o.Expires) {
		return true
	}

	return false
}

// SetExpires gets a reference to the given string and assigns it to the Expires field.
func (o *VideoFile) SetExpires(v string) {
	o.Expires = &v
}

// GetFps returns the Fps field value
func (o *VideoFile) GetFps() float32 {
	if o == nil {
		var ret float32
		return ret
	}

	return o.Fps
}

// GetFpsOk returns a tuple with the Fps field value
// and a boolean to check if the value has been set.
func (o *VideoFile) GetFpsOk() (*float32, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Fps, true
}

// SetFps sets field value
func (o *VideoFile) SetFps(v float32) {
	o.Fps = v
}

// GetHeight returns the Height field value
// If the value is explicit nil, the zero value for float32 will be returned
func (o *VideoFile) GetHeight() float32 {
	if o == nil || o.Height.Get() == nil {
		var ret float32
		return ret
	}

	return *o.Height.Get()
}

// GetHeightOk returns a tuple with the Height field value
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *VideoFile) GetHeightOk() (*float32, bool) {
	if o == nil {
		return nil, false
	}
	return o.Height.Get(), o.Height.IsSet()
}

// SetHeight sets field value
func (o *VideoFile) SetHeight(v float32) {
	o.Height.Set(&v)
}

// GetLink returns the Link field value
func (o *VideoFile) GetLink() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Link
}

// GetLinkOk returns a tuple with the Link field value
// and a boolean to check if the value has been set.
func (o *VideoFile) GetLinkOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Link, true
}

// SetLink sets field value
func (o *VideoFile) SetLink(v string) {
	o.Link = v
}

// GetLog returns the Log field value if set, zero value otherwise.
func (o *VideoFile) GetLog() VideoFileLog {
	if o == nil || IsNil(o.Log) {
		var ret VideoFileLog
		return ret
	}
	return *o.Log
}

// GetLogOk returns a tuple with the Log field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *VideoFile) GetLogOk() (*VideoFileLog, bool) {
	if o == nil || IsNil(o.Log) {
		return nil, false
	}
	return o.Log, true
}

// HasLog returns a boolean if a field has been set.
func (o *VideoFile) HasLog() bool {
	if o != nil && !IsNil(o.Log) {
		return true
	}

	return false
}

// SetLog gets a reference to the given VideoFileLog and assigns it to the Log field.
func (o *VideoFile) SetLog(v VideoFileLog) {
	o.Log = &v
}

// GetMd5 returns the Md5 field value
func (o *VideoFile) GetMd5() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Md5
}

// GetMd5Ok returns a tuple with the Md5 field value
// and a boolean to check if the value has been set.
func (o *VideoFile) GetMd5Ok() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Md5, true
}

// SetMd5 sets field value
func (o *VideoFile) SetMd5(v string) {
	o.Md5 = v
}

// GetPublicName returns the PublicName field value
func (o *VideoFile) GetPublicName() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.PublicName
}

// GetPublicNameOk returns a tuple with the PublicName field value
// and a boolean to check if the value has been set.
func (o *VideoFile) GetPublicNameOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.PublicName, true
}

// SetPublicName sets field value
func (o *VideoFile) SetPublicName(v string) {
	o.PublicName = v
}

// GetQuality returns the Quality field value
func (o *VideoFile) GetQuality() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Quality
}

// GetQualityOk returns a tuple with the Quality field value
// and a boolean to check if the value has been set.
func (o *VideoFile) GetQualityOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Quality, true
}

// SetQuality sets field value
func (o *VideoFile) SetQuality(v string) {
	o.Quality = v
}

// GetRendition returns the Rendition field value
func (o *VideoFile) GetRendition() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Rendition
}

// GetRenditionOk returns a tuple with the Rendition field value
// and a boolean to check if the value has been set.
func (o *VideoFile) GetRenditionOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Rendition, true
}

// SetRendition sets field value
func (o *VideoFile) SetRendition(v string) {
	o.Rendition = v
}

// GetSize returns the Size field value
// If the value is explicit nil, the zero value for float32 will be returned
func (o *VideoFile) GetSize() float32 {
	if o == nil || o.Size.Get() == nil {
		var ret float32
		return ret
	}

	return *o.Size.Get()
}

// GetSizeOk returns a tuple with the Size field value
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *VideoFile) GetSizeOk() (*float32, bool) {
	if o == nil {
		return nil, false
	}
	return o.Size.Get(), o.Size.IsSet()
}

// SetSize sets field value
func (o *VideoFile) SetSize(v float32) {
	o.Size.Set(&v)
}

// GetSizeShort returns the SizeShort field value
func (o *VideoFile) GetSizeShort() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.SizeShort
}

// GetSizeShortOk returns a tuple with the SizeShort field value
// and a boolean to check if the value has been set.
func (o *VideoFile) GetSizeShortOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.SizeShort, true
}

// SetSizeShort sets field value
func (o *VideoFile) SetSizeShort(v string) {
	o.SizeShort = v
}

// GetSourceLink returns the SourceLink field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *VideoFile) GetSourceLink() string {
	if o == nil || IsNil(o.SourceLink.Get()) {
		var ret string
		return ret
	}
	return *o.SourceLink.Get()
}

// GetSourceLinkOk returns a tuple with the SourceLink field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *VideoFile) GetSourceLinkOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.SourceLink.Get(), o.SourceLink.IsSet()
}

// HasSourceLink returns a boolean if a field has been set.
func (o *VideoFile) HasSourceLink() bool {
	if o != nil && o.SourceLink.IsSet() {
		return true
	}

	return false
}

// SetSourceLink gets a reference to the given NullableString and assigns it to the SourceLink field.
func (o *VideoFile) SetSourceLink(v string) {
	o.SourceLink.Set(&v)
}

// SetSourceLinkNil sets the value for SourceLink to be an explicit nil
func (o *VideoFile) SetSourceLinkNil() {
	o.SourceLink.Set(nil)
}

// UnsetSourceLink ensures that no value is present for SourceLink, not even an explicit nil
func (o *VideoFile) UnsetSourceLink() {
	o.SourceLink.Unset()
}

// GetType returns the Type field value
// If the value is explicit nil, the zero value for string will be returned
func (o *VideoFile) GetType() string {
	if o == nil || o.Type.Get() == nil {
		var ret string
		return ret
	}

	return *o.Type.Get()
}

// GetTypeOk returns a tuple with the Type field value
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *VideoFile) GetTypeOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.Type.Get(), o.Type.IsSet()
}

// SetType sets field value
func (o *VideoFile) SetType(v string) {
	o.Type.Set(&v)
}

// GetVideoFileId returns the VideoFileId field value if set, zero value otherwise.
func (o *VideoFile) GetVideoFileId() string {
	if o == nil || IsNil(o.VideoFileId) {
		var ret string
		return ret
	}
	return *o.VideoFileId
}

// GetVideoFileIdOk returns a tuple with the VideoFileId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *VideoFile) GetVideoFileIdOk() (*string, bool) {
	if o == nil || IsNil(o.VideoFileId) {
		return nil, false
	}
	return o.VideoFileId, true
}

// HasVideoFileId returns a boolean if a field has been set.
func (o *VideoFile) HasVideoFileId() bool {
	if o != nil && !IsNil(o.VideoFileId) {
		return true
	}

	return false
}

// SetVideoFileId gets a reference to the given string and assigns it to the VideoFileId field.
func (o *VideoFile) SetVideoFileId(v string) {
	o.VideoFileId = &v
}

// GetWidth returns the Width field value
// If the value is explicit nil, the zero value for float32 will be returned
func (o *VideoFile) GetWidth() float32 {
	if o == nil || o.Width.Get() == nil {
		var ret float32
		return ret
	}

	return *o.Width.Get()
}

// GetWidthOk returns a tuple with the Width field value
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *VideoFile) GetWidthOk() (*float32, bool) {
	if o == nil {
		return nil, false
	}
	return o.Width.Get(), o.Width.IsSet()
}

// SetWidth sets field value
func (o *VideoFile) SetWidth(v float32) {
	o.Width.Set(&v)
}

func (o VideoFile) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o VideoFile) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["codec"] = o.Codec.Get()
	toSerialize["created_time"] = o.CreatedTime
	if !IsNil(o.Expires) {
		toSerialize["expires"] = o.Expires
	}
	toSerialize["fps"] = o.Fps
	toSerialize["height"] = o.Height.Get()
	toSerialize["link"] = o.Link
	if !IsNil(o.Log) {
		toSerialize["log"] = o.Log
	}
	toSerialize["md5"] = o.Md5
	toSerialize["public_name"] = o.PublicName
	toSerialize["quality"] = o.Quality
	toSerialize["rendition"] = o.Rendition
	toSerialize["size"] = o.Size.Get()
	toSerialize["size_short"] = o.SizeShort
	if o.SourceLink.IsSet() {
		toSerialize["source_link"] = o.SourceLink.Get()
	}
	toSerialize["type"] = o.Type.Get()
	if !IsNil(o.VideoFileId) {
		toSerialize["video_file_id"] = o.VideoFileId
	}
	toSerialize["width"] = o.Width.Get()
	return toSerialize, nil
}

type NullableVideoFile struct {
	value *VideoFile
	isSet bool
}

func (v NullableVideoFile) Get() *VideoFile {
	return v.value
}

func (v *NullableVideoFile) Set(val *VideoFile) {
	v.value = val
	v.isSet = true
}

func (v NullableVideoFile) IsSet() bool {
	return v.isSet
}

func (v *NullableVideoFile) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableVideoFile(val *VideoFile) *NullableVideoFile {
	return &NullableVideoFile{value: val, isSet: true}
}

func (v NullableVideoFile) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableVideoFile) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
