/*
Vimeo API

Build something great. Vimeo's API supports flexible, high-quality video integration with your custom apps.

API version: 3.4
*/

package vimeoapi

import (
	"encoding/json"
)

// checks if the HlsDashVideoFile type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &HlsDashVideoFile{}

// HlsDashVideoFile struct for HlsDashVideoFile
type HlsDashVideoFile struct {
	// The direct link to the video file.
	Link NullableString `json:"link"`
	// The time in ISO 8601 format when the link to the video file expires.
	LinkExpirationTime string `json:"link_expiration_time"`
	// The URL for logging events.
	Log NullableString `json:"log,omitempty"`
}

// NewHlsDashVideoFile instantiates a new HlsDashVideoFile object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewHlsDashVideoFile(link NullableString, linkExpirationTime string) *HlsDashVideoFile {
	this := HlsDashVideoFile{}
	this.Link = link
	this.LinkExpirationTime = linkExpirationTime
	return &this
}

// NewHlsDashVideoFileWithDefaults instantiates a new HlsDashVideoFile object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewHlsDashVideoFileWithDefaults() *HlsDashVideoFile {
	this := HlsDashVideoFile{}
	return &this
}

// GetLink returns the Link field value
// If the value is explicit nil, the zero value for string will be returned
func (o *HlsDashVideoFile) GetLink() string {
	if o == nil || o.Link.Get() == nil {
		var ret string
		return ret
	}

	return *o.Link.Get()
}

// GetLinkOk returns a tuple with the Link field value
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *HlsDashVideoFile) GetLinkOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.Link.Get(), o.Link.IsSet()
}

// SetLink sets field value
func (o *HlsDashVideoFile) SetLink(v string) {
	o.Link.Set(&v)
}

// GetLinkExpirationTime returns the LinkExpirationTime field value
func (o *HlsDashVideoFile) GetLinkExpirationTime() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.LinkExpirationTime
}

// GetLinkExpirationTimeOk returns a tuple with the LinkExpirationTime field value
// and a boolean to check if the value has been set.
func (o *HlsDashVideoFile) GetLinkExpirationTimeOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.LinkExpirationTime, true
}

// SetLinkExpirationTime sets field value
func (o *HlsDashVideoFile) SetLinkExpirationTime(v string) {
	o.LinkExpirationTime = v
}

// GetLog returns the Log field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *HlsDashVideoFile) GetLog() string {
	if o == nil || IsNil(o.Log.Get()) {
		var ret string
		return ret
	}
	return *o.Log.Get()
}

// GetLogOk returns a tuple with the Log field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *HlsDashVideoFile) GetLogOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.Log.Get(), o.Log.IsSet()
}

// HasLog returns a boolean if a field has been set.
func (o *HlsDashVideoFile) HasLog() bool {
	if o != nil && o.Log.IsSet() {
		return true
	}

	return false
}

// SetLog gets a reference to the given NullableString and assigns it to the Log field.
func (o *HlsDashVideoFile) SetLog(v string) {
	o.Log.Set(&v)
}

// SetLogNil sets the value for Log to be an explicit nil
func (o *HlsDashVideoFile) SetLogNil() {
	o.Log.Set(nil)
}

// UnsetLog ensures that no value is present for Log, not even an explicit nil
func (o *HlsDashVideoFile) UnsetLog() {
	o.Log.Unset()
}

func (o HlsDashVideoFile) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o HlsDashVideoFile) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["link"] = o.Link.Get()
	toSerialize["link_expiration_time"] = o.LinkExpirationTime
	if o.Log.IsSet() {
		toSerialize["log"] = o.Log.Get()
	}
	return toSerialize, nil
}

type NullableHlsDashVideoFile struct {
	value *HlsDashVideoFile
	isSet bool
}

func (v NullableHlsDashVideoFile) Get() *HlsDashVideoFile {
	return v.value
}

func (v *NullableHlsDashVideoFile) Set(val *HlsDashVideoFile) {
	v.value = val
	v.isSet = true
}

func (v NullableHlsDashVideoFile) IsSet() bool {
	return v.isSet
}

func (v *NullableHlsDashVideoFile) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableHlsDashVideoFile(val *HlsDashVideoFile) *NullableHlsDashVideoFile {
	return &NullableHlsDashVideoFile{value: val, isSet: true}
}

func (v NullableHlsDashVideoFile) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableHlsDashVideoFile) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
