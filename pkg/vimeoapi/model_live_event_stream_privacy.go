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

// checks if the LiveEventStreamPrivacy type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &LiveEventStreamPrivacy{}

// LiveEventStreamPrivacy The initial privacy settings of videos generated by streaming to the event as well as the embed privacy of the entire collection.
type LiveEventStreamPrivacy struct {
	// The event's embed permission setting.  Option descriptions:  * `private` - The event can't be embedded on any domain.  * `public` - The event can be embedded on any domain.  * `whitelist` - The event can be embedded on whitelisted domains only.
	Embed string `json:"embed"`
	// The hash for unlisted events.
	UnlistedHash NullableString `json:"unlisted_hash"`
	// The general privacy setting for generated videos and the embed privacy of the entire collection.  Option descriptions:  * `anybody` - Anyone can access the videos. This privacy setting appears as `Public` on the Vimeo front end.  * `embed_only` - The videos don't appear on Vimeo, but they can be embedded elsewhere.  * `nobody` - Only the event owner can access the videos. This privacy setting appears as `Private` on the Vimeo front end.  * `password` - Only those with the password can access the videos.  * `unlisted` - Only those with the private link can access the videos.
	View string `json:"view"`
}

// NewLiveEventStreamPrivacy instantiates a new LiveEventStreamPrivacy object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewLiveEventStreamPrivacy(embed string, unlistedHash NullableString, view string) *LiveEventStreamPrivacy {
	this := LiveEventStreamPrivacy{}
	this.Embed = embed
	this.UnlistedHash = unlistedHash
	this.View = view
	return &this
}

// NewLiveEventStreamPrivacyWithDefaults instantiates a new LiveEventStreamPrivacy object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewLiveEventStreamPrivacyWithDefaults() *LiveEventStreamPrivacy {
	this := LiveEventStreamPrivacy{}
	return &this
}

// GetEmbed returns the Embed field value
func (o *LiveEventStreamPrivacy) GetEmbed() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Embed
}

// GetEmbedOk returns a tuple with the Embed field value
// and a boolean to check if the value has been set.
func (o *LiveEventStreamPrivacy) GetEmbedOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Embed, true
}

// SetEmbed sets field value
func (o *LiveEventStreamPrivacy) SetEmbed(v string) {
	o.Embed = v
}

// GetUnlistedHash returns the UnlistedHash field value
// If the value is explicit nil, the zero value for string will be returned
func (o *LiveEventStreamPrivacy) GetUnlistedHash() string {
	if o == nil || o.UnlistedHash.Get() == nil {
		var ret string
		return ret
	}

	return *o.UnlistedHash.Get()
}

// GetUnlistedHashOk returns a tuple with the UnlistedHash field value
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *LiveEventStreamPrivacy) GetUnlistedHashOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.UnlistedHash.Get(), o.UnlistedHash.IsSet()
}

// SetUnlistedHash sets field value
func (o *LiveEventStreamPrivacy) SetUnlistedHash(v string) {
	o.UnlistedHash.Set(&v)
}

// GetView returns the View field value
func (o *LiveEventStreamPrivacy) GetView() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.View
}

// GetViewOk returns a tuple with the View field value
// and a boolean to check if the value has been set.
func (o *LiveEventStreamPrivacy) GetViewOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.View, true
}

// SetView sets field value
func (o *LiveEventStreamPrivacy) SetView(v string) {
	o.View = v
}

func (o LiveEventStreamPrivacy) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o LiveEventStreamPrivacy) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["embed"] = o.Embed
	toSerialize["unlisted_hash"] = o.UnlistedHash.Get()
	toSerialize["view"] = o.View
	return toSerialize, nil
}

type NullableLiveEventStreamPrivacy struct {
	value *LiveEventStreamPrivacy
	isSet bool
}

func (v NullableLiveEventStreamPrivacy) Get() *LiveEventStreamPrivacy {
	return v.value
}

func (v *NullableLiveEventStreamPrivacy) Set(val *LiveEventStreamPrivacy) {
	v.value = val
	v.isSet = true
}

func (v NullableLiveEventStreamPrivacy) IsSet() bool {
	return v.isSet
}

func (v *NullableLiveEventStreamPrivacy) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableLiveEventStreamPrivacy(val *LiveEventStreamPrivacy) *NullableLiveEventStreamPrivacy {
	return &NullableLiveEventStreamPrivacy{value: val, isSet: true}
}

func (v NullableLiveEventStreamPrivacy) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableLiveEventStreamPrivacy) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
