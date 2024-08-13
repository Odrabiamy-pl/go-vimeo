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

// checks if the CreateEmbedPresetsAlt1RequestEmbedButtons type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &CreateEmbedPresetsAlt1RequestEmbedButtons{}

// CreateEmbedPresetsAlt1RequestEmbedButtons struct for CreateEmbedPresetsAlt1RequestEmbedButtons
type CreateEmbedPresetsAlt1RequestEmbedButtons struct {
	// Whether to show the `embed` button on the embeddable player.
	Embed *bool `json:"embed,omitempty"`
	// Whether to show the `fullscreen` button on the embeddable player.
	Fullscreen *bool `json:"fullscreen,omitempty"`
	// Whether to show the `HD` button on the embeddable player.
	Hd *bool `json:"hd,omitempty"`
	// Whether to show the `like` button on the embeddable player.
	Like *bool `json:"like,omitempty"`
	// Whether to show the `scaling` button on the embeddable player in fullscreen mode.
	Scaling *bool `json:"scaling,omitempty"`
	// Whether to show the `share` button on the embeddable player.
	Share *bool `json:"share,omitempty"`
	// Whether to show the `watch later` button on the embeddable player.
	Watchlater *bool `json:"watchlater,omitempty"`
}

// NewCreateEmbedPresetsAlt1RequestEmbedButtons instantiates a new CreateEmbedPresetsAlt1RequestEmbedButtons object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewCreateEmbedPresetsAlt1RequestEmbedButtons() *CreateEmbedPresetsAlt1RequestEmbedButtons {
	this := CreateEmbedPresetsAlt1RequestEmbedButtons{}
	return &this
}

// NewCreateEmbedPresetsAlt1RequestEmbedButtonsWithDefaults instantiates a new CreateEmbedPresetsAlt1RequestEmbedButtons object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewCreateEmbedPresetsAlt1RequestEmbedButtonsWithDefaults() *CreateEmbedPresetsAlt1RequestEmbedButtons {
	this := CreateEmbedPresetsAlt1RequestEmbedButtons{}
	return &this
}

// GetEmbed returns the Embed field value if set, zero value otherwise.
func (o *CreateEmbedPresetsAlt1RequestEmbedButtons) GetEmbed() bool {
	if o == nil || IsNil(o.Embed) {
		var ret bool
		return ret
	}
	return *o.Embed
}

// GetEmbedOk returns a tuple with the Embed field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CreateEmbedPresetsAlt1RequestEmbedButtons) GetEmbedOk() (*bool, bool) {
	if o == nil || IsNil(o.Embed) {
		return nil, false
	}
	return o.Embed, true
}

// HasEmbed returns a boolean if a field has been set.
func (o *CreateEmbedPresetsAlt1RequestEmbedButtons) HasEmbed() bool {
	if o != nil && !IsNil(o.Embed) {
		return true
	}

	return false
}

// SetEmbed gets a reference to the given bool and assigns it to the Embed field.
func (o *CreateEmbedPresetsAlt1RequestEmbedButtons) SetEmbed(v bool) {
	o.Embed = &v
}

// GetFullscreen returns the Fullscreen field value if set, zero value otherwise.
func (o *CreateEmbedPresetsAlt1RequestEmbedButtons) GetFullscreen() bool {
	if o == nil || IsNil(o.Fullscreen) {
		var ret bool
		return ret
	}
	return *o.Fullscreen
}

// GetFullscreenOk returns a tuple with the Fullscreen field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CreateEmbedPresetsAlt1RequestEmbedButtons) GetFullscreenOk() (*bool, bool) {
	if o == nil || IsNil(o.Fullscreen) {
		return nil, false
	}
	return o.Fullscreen, true
}

// HasFullscreen returns a boolean if a field has been set.
func (o *CreateEmbedPresetsAlt1RequestEmbedButtons) HasFullscreen() bool {
	if o != nil && !IsNil(o.Fullscreen) {
		return true
	}

	return false
}

// SetFullscreen gets a reference to the given bool and assigns it to the Fullscreen field.
func (o *CreateEmbedPresetsAlt1RequestEmbedButtons) SetFullscreen(v bool) {
	o.Fullscreen = &v
}

// GetHd returns the Hd field value if set, zero value otherwise.
func (o *CreateEmbedPresetsAlt1RequestEmbedButtons) GetHd() bool {
	if o == nil || IsNil(o.Hd) {
		var ret bool
		return ret
	}
	return *o.Hd
}

// GetHdOk returns a tuple with the Hd field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CreateEmbedPresetsAlt1RequestEmbedButtons) GetHdOk() (*bool, bool) {
	if o == nil || IsNil(o.Hd) {
		return nil, false
	}
	return o.Hd, true
}

// HasHd returns a boolean if a field has been set.
func (o *CreateEmbedPresetsAlt1RequestEmbedButtons) HasHd() bool {
	if o != nil && !IsNil(o.Hd) {
		return true
	}

	return false
}

// SetHd gets a reference to the given bool and assigns it to the Hd field.
func (o *CreateEmbedPresetsAlt1RequestEmbedButtons) SetHd(v bool) {
	o.Hd = &v
}

// GetLike returns the Like field value if set, zero value otherwise.
func (o *CreateEmbedPresetsAlt1RequestEmbedButtons) GetLike() bool {
	if o == nil || IsNil(o.Like) {
		var ret bool
		return ret
	}
	return *o.Like
}

// GetLikeOk returns a tuple with the Like field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CreateEmbedPresetsAlt1RequestEmbedButtons) GetLikeOk() (*bool, bool) {
	if o == nil || IsNil(o.Like) {
		return nil, false
	}
	return o.Like, true
}

// HasLike returns a boolean if a field has been set.
func (o *CreateEmbedPresetsAlt1RequestEmbedButtons) HasLike() bool {
	if o != nil && !IsNil(o.Like) {
		return true
	}

	return false
}

// SetLike gets a reference to the given bool and assigns it to the Like field.
func (o *CreateEmbedPresetsAlt1RequestEmbedButtons) SetLike(v bool) {
	o.Like = &v
}

// GetScaling returns the Scaling field value if set, zero value otherwise.
func (o *CreateEmbedPresetsAlt1RequestEmbedButtons) GetScaling() bool {
	if o == nil || IsNil(o.Scaling) {
		var ret bool
		return ret
	}
	return *o.Scaling
}

// GetScalingOk returns a tuple with the Scaling field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CreateEmbedPresetsAlt1RequestEmbedButtons) GetScalingOk() (*bool, bool) {
	if o == nil || IsNil(o.Scaling) {
		return nil, false
	}
	return o.Scaling, true
}

// HasScaling returns a boolean if a field has been set.
func (o *CreateEmbedPresetsAlt1RequestEmbedButtons) HasScaling() bool {
	if o != nil && !IsNil(o.Scaling) {
		return true
	}

	return false
}

// SetScaling gets a reference to the given bool and assigns it to the Scaling field.
func (o *CreateEmbedPresetsAlt1RequestEmbedButtons) SetScaling(v bool) {
	o.Scaling = &v
}

// GetShare returns the Share field value if set, zero value otherwise.
func (o *CreateEmbedPresetsAlt1RequestEmbedButtons) GetShare() bool {
	if o == nil || IsNil(o.Share) {
		var ret bool
		return ret
	}
	return *o.Share
}

// GetShareOk returns a tuple with the Share field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CreateEmbedPresetsAlt1RequestEmbedButtons) GetShareOk() (*bool, bool) {
	if o == nil || IsNil(o.Share) {
		return nil, false
	}
	return o.Share, true
}

// HasShare returns a boolean if a field has been set.
func (o *CreateEmbedPresetsAlt1RequestEmbedButtons) HasShare() bool {
	if o != nil && !IsNil(o.Share) {
		return true
	}

	return false
}

// SetShare gets a reference to the given bool and assigns it to the Share field.
func (o *CreateEmbedPresetsAlt1RequestEmbedButtons) SetShare(v bool) {
	o.Share = &v
}

// GetWatchlater returns the Watchlater field value if set, zero value otherwise.
func (o *CreateEmbedPresetsAlt1RequestEmbedButtons) GetWatchlater() bool {
	if o == nil || IsNil(o.Watchlater) {
		var ret bool
		return ret
	}
	return *o.Watchlater
}

// GetWatchlaterOk returns a tuple with the Watchlater field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CreateEmbedPresetsAlt1RequestEmbedButtons) GetWatchlaterOk() (*bool, bool) {
	if o == nil || IsNil(o.Watchlater) {
		return nil, false
	}
	return o.Watchlater, true
}

// HasWatchlater returns a boolean if a field has been set.
func (o *CreateEmbedPresetsAlt1RequestEmbedButtons) HasWatchlater() bool {
	if o != nil && !IsNil(o.Watchlater) {
		return true
	}

	return false
}

// SetWatchlater gets a reference to the given bool and assigns it to the Watchlater field.
func (o *CreateEmbedPresetsAlt1RequestEmbedButtons) SetWatchlater(v bool) {
	o.Watchlater = &v
}

func (o CreateEmbedPresetsAlt1RequestEmbedButtons) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o CreateEmbedPresetsAlt1RequestEmbedButtons) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Embed) {
		toSerialize["embed"] = o.Embed
	}
	if !IsNil(o.Fullscreen) {
		toSerialize["fullscreen"] = o.Fullscreen
	}
	if !IsNil(o.Hd) {
		toSerialize["hd"] = o.Hd
	}
	if !IsNil(o.Like) {
		toSerialize["like"] = o.Like
	}
	if !IsNil(o.Scaling) {
		toSerialize["scaling"] = o.Scaling
	}
	if !IsNil(o.Share) {
		toSerialize["share"] = o.Share
	}
	if !IsNil(o.Watchlater) {
		toSerialize["watchlater"] = o.Watchlater
	}
	return toSerialize, nil
}

type NullableCreateEmbedPresetsAlt1RequestEmbedButtons struct {
	value *CreateEmbedPresetsAlt1RequestEmbedButtons
	isSet bool
}

func (v NullableCreateEmbedPresetsAlt1RequestEmbedButtons) Get() *CreateEmbedPresetsAlt1RequestEmbedButtons {
	return v.value
}

func (v *NullableCreateEmbedPresetsAlt1RequestEmbedButtons) Set(val *CreateEmbedPresetsAlt1RequestEmbedButtons) {
	v.value = val
	v.isSet = true
}

func (v NullableCreateEmbedPresetsAlt1RequestEmbedButtons) IsSet() bool {
	return v.isSet
}

func (v *NullableCreateEmbedPresetsAlt1RequestEmbedButtons) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableCreateEmbedPresetsAlt1RequestEmbedButtons(val *CreateEmbedPresetsAlt1RequestEmbedButtons) *NullableCreateEmbedPresetsAlt1RequestEmbedButtons {
	return &NullableCreateEmbedPresetsAlt1RequestEmbedButtons{value: val, isSet: true}
}

func (v NullableCreateEmbedPresetsAlt1RequestEmbedButtons) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableCreateEmbedPresetsAlt1RequestEmbedButtons) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


