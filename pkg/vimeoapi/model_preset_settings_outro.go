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

// checks if the PresetSettingsOutro type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &PresetSettingsOutro{}

// PresetSettingsOutro struct for PresetSettingsOutro
type PresetSettingsOutro struct {
	// A comma-separated list of video URIs. This field appears only when **type** is `uploaded_clips`.
	Clips NullableString `json:"clips,omitempty"`
	Link NullablePresetSettingsOutroLink `json:"link,omitempty"`
	// The outro text. This appears only when **type** is `text`.
	Text NullableString `json:"text,omitempty"`
	// The preset outro type.  Option descriptions:  * `link` - The outro includes a link.  * `no idea` - The outro type is `no idea`. The outro includes uploaded videos.  * `text` - The outro includes text.  * `uploaded_clips` - The outro includes uploaded videos.  * `uploaded_videos` - The outro includes uploaded videos. 
	Type string `json:"type"`
	// A comma-separated list of video URIs. This field appears only when **type** is `no idea`.
	Videos NullableString `json:"videos,omitempty"`
}

// NewPresetSettingsOutro instantiates a new PresetSettingsOutro object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewPresetSettingsOutro(type_ string) *PresetSettingsOutro {
	this := PresetSettingsOutro{}
	this.Type = type_
	return &this
}

// NewPresetSettingsOutroWithDefaults instantiates a new PresetSettingsOutro object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewPresetSettingsOutroWithDefaults() *PresetSettingsOutro {
	this := PresetSettingsOutro{}
	return &this
}

// GetClips returns the Clips field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *PresetSettingsOutro) GetClips() string {
	if o == nil || IsNil(o.Clips.Get()) {
		var ret string
		return ret
	}
	return *o.Clips.Get()
}

// GetClipsOk returns a tuple with the Clips field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *PresetSettingsOutro) GetClipsOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.Clips.Get(), o.Clips.IsSet()
}

// HasClips returns a boolean if a field has been set.
func (o *PresetSettingsOutro) HasClips() bool {
	if o != nil && o.Clips.IsSet() {
		return true
	}

	return false
}

// SetClips gets a reference to the given NullableString and assigns it to the Clips field.
func (o *PresetSettingsOutro) SetClips(v string) {
	o.Clips.Set(&v)
}
// SetClipsNil sets the value for Clips to be an explicit nil
func (o *PresetSettingsOutro) SetClipsNil() {
	o.Clips.Set(nil)
}

// UnsetClips ensures that no value is present for Clips, not even an explicit nil
func (o *PresetSettingsOutro) UnsetClips() {
	o.Clips.Unset()
}

// GetLink returns the Link field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *PresetSettingsOutro) GetLink() PresetSettingsOutroLink {
	if o == nil || IsNil(o.Link.Get()) {
		var ret PresetSettingsOutroLink
		return ret
	}
	return *o.Link.Get()
}

// GetLinkOk returns a tuple with the Link field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *PresetSettingsOutro) GetLinkOk() (*PresetSettingsOutroLink, bool) {
	if o == nil {
		return nil, false
	}
	return o.Link.Get(), o.Link.IsSet()
}

// HasLink returns a boolean if a field has been set.
func (o *PresetSettingsOutro) HasLink() bool {
	if o != nil && o.Link.IsSet() {
		return true
	}

	return false
}

// SetLink gets a reference to the given NullablePresetSettingsOutroLink and assigns it to the Link field.
func (o *PresetSettingsOutro) SetLink(v PresetSettingsOutroLink) {
	o.Link.Set(&v)
}
// SetLinkNil sets the value for Link to be an explicit nil
func (o *PresetSettingsOutro) SetLinkNil() {
	o.Link.Set(nil)
}

// UnsetLink ensures that no value is present for Link, not even an explicit nil
func (o *PresetSettingsOutro) UnsetLink() {
	o.Link.Unset()
}

// GetText returns the Text field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *PresetSettingsOutro) GetText() string {
	if o == nil || IsNil(o.Text.Get()) {
		var ret string
		return ret
	}
	return *o.Text.Get()
}

// GetTextOk returns a tuple with the Text field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *PresetSettingsOutro) GetTextOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.Text.Get(), o.Text.IsSet()
}

// HasText returns a boolean if a field has been set.
func (o *PresetSettingsOutro) HasText() bool {
	if o != nil && o.Text.IsSet() {
		return true
	}

	return false
}

// SetText gets a reference to the given NullableString and assigns it to the Text field.
func (o *PresetSettingsOutro) SetText(v string) {
	o.Text.Set(&v)
}
// SetTextNil sets the value for Text to be an explicit nil
func (o *PresetSettingsOutro) SetTextNil() {
	o.Text.Set(nil)
}

// UnsetText ensures that no value is present for Text, not even an explicit nil
func (o *PresetSettingsOutro) UnsetText() {
	o.Text.Unset()
}

// GetType returns the Type field value
func (o *PresetSettingsOutro) GetType() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Type
}

// GetTypeOk returns a tuple with the Type field value
// and a boolean to check if the value has been set.
func (o *PresetSettingsOutro) GetTypeOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Type, true
}

// SetType sets field value
func (o *PresetSettingsOutro) SetType(v string) {
	o.Type = v
}

// GetVideos returns the Videos field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *PresetSettingsOutro) GetVideos() string {
	if o == nil || IsNil(o.Videos.Get()) {
		var ret string
		return ret
	}
	return *o.Videos.Get()
}

// GetVideosOk returns a tuple with the Videos field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *PresetSettingsOutro) GetVideosOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.Videos.Get(), o.Videos.IsSet()
}

// HasVideos returns a boolean if a field has been set.
func (o *PresetSettingsOutro) HasVideos() bool {
	if o != nil && o.Videos.IsSet() {
		return true
	}

	return false
}

// SetVideos gets a reference to the given NullableString and assigns it to the Videos field.
func (o *PresetSettingsOutro) SetVideos(v string) {
	o.Videos.Set(&v)
}
// SetVideosNil sets the value for Videos to be an explicit nil
func (o *PresetSettingsOutro) SetVideosNil() {
	o.Videos.Set(nil)
}

// UnsetVideos ensures that no value is present for Videos, not even an explicit nil
func (o *PresetSettingsOutro) UnsetVideos() {
	o.Videos.Unset()
}

func (o PresetSettingsOutro) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o PresetSettingsOutro) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if o.Clips.IsSet() {
		toSerialize["clips"] = o.Clips.Get()
	}
	if o.Link.IsSet() {
		toSerialize["link"] = o.Link.Get()
	}
	if o.Text.IsSet() {
		toSerialize["text"] = o.Text.Get()
	}
	toSerialize["type"] = o.Type
	if o.Videos.IsSet() {
		toSerialize["videos"] = o.Videos.Get()
	}
	return toSerialize, nil
}

type NullablePresetSettingsOutro struct {
	value *PresetSettingsOutro
	isSet bool
}

func (v NullablePresetSettingsOutro) Get() *PresetSettingsOutro {
	return v.value
}

func (v *NullablePresetSettingsOutro) Set(val *PresetSettingsOutro) {
	v.value = val
	v.isSet = true
}

func (v NullablePresetSettingsOutro) IsSet() bool {
	return v.isSet
}

func (v *NullablePresetSettingsOutro) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullablePresetSettingsOutro(val *PresetSettingsOutro) *NullablePresetSettingsOutro {
	return &NullablePresetSettingsOutro{value: val, isSet: true}
}

func (v NullablePresetSettingsOutro) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullablePresetSettingsOutro) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


