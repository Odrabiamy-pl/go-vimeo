/*
Vimeo API

Build something great. Vimeo's API supports flexible, high-quality video integration with your custom apps.

API version: 3.4
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package vimeoapi

import (
	"encoding/json"
	"bytes"
	"fmt"
)

// checks if the ChannelMetadataInteractions type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &ChannelMetadataInteractions{}

// ChannelMetadataInteractions A list of resource URIs related to the channel.
type ChannelMetadataInteractions struct {
	AddModerators ChannelMetadataInteractionsAddModerators `json:"add_moderators"`
	AddTo NullableChannelMetadataInteractionsAddTo `json:"add_to"`
	Follow ChannelMetadataInteractionsFollow `json:"follow"`
	ModerateVideos ChannelMetadataInteractionsModerateVideos `json:"moderate_videos"`
}

type _ChannelMetadataInteractions ChannelMetadataInteractions

// NewChannelMetadataInteractions instantiates a new ChannelMetadataInteractions object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewChannelMetadataInteractions(addModerators ChannelMetadataInteractionsAddModerators, addTo NullableChannelMetadataInteractionsAddTo, follow ChannelMetadataInteractionsFollow, moderateVideos ChannelMetadataInteractionsModerateVideos) *ChannelMetadataInteractions {
	this := ChannelMetadataInteractions{}
	this.AddModerators = addModerators
	this.AddTo = addTo
	this.Follow = follow
	this.ModerateVideos = moderateVideos
	return &this
}

// NewChannelMetadataInteractionsWithDefaults instantiates a new ChannelMetadataInteractions object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewChannelMetadataInteractionsWithDefaults() *ChannelMetadataInteractions {
	this := ChannelMetadataInteractions{}
	return &this
}

// GetAddModerators returns the AddModerators field value
func (o *ChannelMetadataInteractions) GetAddModerators() ChannelMetadataInteractionsAddModerators {
	if o == nil {
		var ret ChannelMetadataInteractionsAddModerators
		return ret
	}

	return o.AddModerators
}

// GetAddModeratorsOk returns a tuple with the AddModerators field value
// and a boolean to check if the value has been set.
func (o *ChannelMetadataInteractions) GetAddModeratorsOk() (*ChannelMetadataInteractionsAddModerators, bool) {
	if o == nil {
		return nil, false
	}
	return &o.AddModerators, true
}

// SetAddModerators sets field value
func (o *ChannelMetadataInteractions) SetAddModerators(v ChannelMetadataInteractionsAddModerators) {
	o.AddModerators = v
}

// GetAddTo returns the AddTo field value
// If the value is explicit nil, the zero value for ChannelMetadataInteractionsAddTo will be returned
func (o *ChannelMetadataInteractions) GetAddTo() ChannelMetadataInteractionsAddTo {
	if o == nil || o.AddTo.Get() == nil {
		var ret ChannelMetadataInteractionsAddTo
		return ret
	}

	return *o.AddTo.Get()
}

// GetAddToOk returns a tuple with the AddTo field value
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *ChannelMetadataInteractions) GetAddToOk() (*ChannelMetadataInteractionsAddTo, bool) {
	if o == nil {
		return nil, false
	}
	return o.AddTo.Get(), o.AddTo.IsSet()
}

// SetAddTo sets field value
func (o *ChannelMetadataInteractions) SetAddTo(v ChannelMetadataInteractionsAddTo) {
	o.AddTo.Set(&v)
}

// GetFollow returns the Follow field value
func (o *ChannelMetadataInteractions) GetFollow() ChannelMetadataInteractionsFollow {
	if o == nil {
		var ret ChannelMetadataInteractionsFollow
		return ret
	}

	return o.Follow
}

// GetFollowOk returns a tuple with the Follow field value
// and a boolean to check if the value has been set.
func (o *ChannelMetadataInteractions) GetFollowOk() (*ChannelMetadataInteractionsFollow, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Follow, true
}

// SetFollow sets field value
func (o *ChannelMetadataInteractions) SetFollow(v ChannelMetadataInteractionsFollow) {
	o.Follow = v
}

// GetModerateVideos returns the ModerateVideos field value
func (o *ChannelMetadataInteractions) GetModerateVideos() ChannelMetadataInteractionsModerateVideos {
	if o == nil {
		var ret ChannelMetadataInteractionsModerateVideos
		return ret
	}

	return o.ModerateVideos
}

// GetModerateVideosOk returns a tuple with the ModerateVideos field value
// and a boolean to check if the value has been set.
func (o *ChannelMetadataInteractions) GetModerateVideosOk() (*ChannelMetadataInteractionsModerateVideos, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ModerateVideos, true
}

// SetModerateVideos sets field value
func (o *ChannelMetadataInteractions) SetModerateVideos(v ChannelMetadataInteractionsModerateVideos) {
	o.ModerateVideos = v
}

func (o ChannelMetadataInteractions) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o ChannelMetadataInteractions) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["add_moderators"] = o.AddModerators
	toSerialize["add_to"] = o.AddTo.Get()
	toSerialize["follow"] = o.Follow
	toSerialize["moderate_videos"] = o.ModerateVideos
	return toSerialize, nil
}

func (o *ChannelMetadataInteractions) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"add_moderators",
		"add_to",
		"follow",
		"moderate_videos",
	}

	allProperties := make(map[string]interface{})

	err = json.Unmarshal(data, &allProperties)

	if err != nil {
		return err;
	}

	for _, requiredProperty := range(requiredProperties) {
		if _, exists := allProperties[requiredProperty]; !exists {
			return fmt.Errorf("no value given for required property %v", requiredProperty)
		}
	}

	varChannelMetadataInteractions := _ChannelMetadataInteractions{}

	decoder := json.NewDecoder(bytes.NewReader(data))
	decoder.DisallowUnknownFields()
	err = decoder.Decode(&varChannelMetadataInteractions)

	if err != nil {
		return err
	}

	*o = ChannelMetadataInteractions(varChannelMetadataInteractions)

	return err
}

type NullableChannelMetadataInteractions struct {
	value *ChannelMetadataInteractions
	isSet bool
}

func (v NullableChannelMetadataInteractions) Get() *ChannelMetadataInteractions {
	return v.value
}

func (v *NullableChannelMetadataInteractions) Set(val *ChannelMetadataInteractions) {
	v.value = val
	v.isSet = true
}

func (v NullableChannelMetadataInteractions) IsSet() bool {
	return v.isSet
}

func (v *NullableChannelMetadataInteractions) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableChannelMetadataInteractions(val *ChannelMetadataInteractions) *NullableChannelMetadataInteractions {
	return &NullableChannelMetadataInteractions{value: val, isSet: true}
}

func (v NullableChannelMetadataInteractions) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableChannelMetadataInteractions) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


