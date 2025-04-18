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

// checks if the UploadVideoAlt1RequestEmbed type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &UploadVideoAlt1RequestEmbed{}

// UploadVideoAlt1RequestEmbed struct for UploadVideoAlt1RequestEmbed
type UploadVideoAlt1RequestEmbed struct {
	Buttons *CreateEmbedPresetsAlt1RequestEmbedButtons `json:"buttons,omitempty"`
	// The main color of the embeddable player.
	Color *string `json:"color,omitempty"`
	EndScreen *UploadVideoAlt1RequestEmbedEndScreen `json:"end_screen,omitempty"`
	Logos *UploadVideoAlt1RequestEmbedLogos `json:"logos,omitempty"`
	// Whether to show the playbar on the embeddable player.
	Playbar *bool `json:"playbar,omitempty"`
	Title *CreateEmbedPresetsAlt1RequestEmbedTitle `json:"title,omitempty"`
	// Whether to show the volume selector on the embeddable player.
	Volume *bool `json:"volume,omitempty"`
}

// NewUploadVideoAlt1RequestEmbed instantiates a new UploadVideoAlt1RequestEmbed object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewUploadVideoAlt1RequestEmbed() *UploadVideoAlt1RequestEmbed {
	this := UploadVideoAlt1RequestEmbed{}
	return &this
}

// NewUploadVideoAlt1RequestEmbedWithDefaults instantiates a new UploadVideoAlt1RequestEmbed object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewUploadVideoAlt1RequestEmbedWithDefaults() *UploadVideoAlt1RequestEmbed {
	this := UploadVideoAlt1RequestEmbed{}
	return &this
}

// GetButtons returns the Buttons field value if set, zero value otherwise.
func (o *UploadVideoAlt1RequestEmbed) GetButtons() CreateEmbedPresetsAlt1RequestEmbedButtons {
	if o == nil || IsNil(o.Buttons) {
		var ret CreateEmbedPresetsAlt1RequestEmbedButtons
		return ret
	}
	return *o.Buttons
}

// GetButtonsOk returns a tuple with the Buttons field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UploadVideoAlt1RequestEmbed) GetButtonsOk() (*CreateEmbedPresetsAlt1RequestEmbedButtons, bool) {
	if o == nil || IsNil(o.Buttons) {
		return nil, false
	}
	return o.Buttons, true
}

// HasButtons returns a boolean if a field has been set.
func (o *UploadVideoAlt1RequestEmbed) HasButtons() bool {
	if o != nil && !IsNil(o.Buttons) {
		return true
	}

	return false
}

// SetButtons gets a reference to the given CreateEmbedPresetsAlt1RequestEmbedButtons and assigns it to the Buttons field.
func (o *UploadVideoAlt1RequestEmbed) SetButtons(v CreateEmbedPresetsAlt1RequestEmbedButtons) {
	o.Buttons = &v
}

// GetColor returns the Color field value if set, zero value otherwise.
func (o *UploadVideoAlt1RequestEmbed) GetColor() string {
	if o == nil || IsNil(o.Color) {
		var ret string
		return ret
	}
	return *o.Color
}

// GetColorOk returns a tuple with the Color field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UploadVideoAlt1RequestEmbed) GetColorOk() (*string, bool) {
	if o == nil || IsNil(o.Color) {
		return nil, false
	}
	return o.Color, true
}

// HasColor returns a boolean if a field has been set.
func (o *UploadVideoAlt1RequestEmbed) HasColor() bool {
	if o != nil && !IsNil(o.Color) {
		return true
	}

	return false
}

// SetColor gets a reference to the given string and assigns it to the Color field.
func (o *UploadVideoAlt1RequestEmbed) SetColor(v string) {
	o.Color = &v
}

// GetEndScreen returns the EndScreen field value if set, zero value otherwise.
func (o *UploadVideoAlt1RequestEmbed) GetEndScreen() UploadVideoAlt1RequestEmbedEndScreen {
	if o == nil || IsNil(o.EndScreen) {
		var ret UploadVideoAlt1RequestEmbedEndScreen
		return ret
	}
	return *o.EndScreen
}

// GetEndScreenOk returns a tuple with the EndScreen field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UploadVideoAlt1RequestEmbed) GetEndScreenOk() (*UploadVideoAlt1RequestEmbedEndScreen, bool) {
	if o == nil || IsNil(o.EndScreen) {
		return nil, false
	}
	return o.EndScreen, true
}

// HasEndScreen returns a boolean if a field has been set.
func (o *UploadVideoAlt1RequestEmbed) HasEndScreen() bool {
	if o != nil && !IsNil(o.EndScreen) {
		return true
	}

	return false
}

// SetEndScreen gets a reference to the given UploadVideoAlt1RequestEmbedEndScreen and assigns it to the EndScreen field.
func (o *UploadVideoAlt1RequestEmbed) SetEndScreen(v UploadVideoAlt1RequestEmbedEndScreen) {
	o.EndScreen = &v
}

// GetLogos returns the Logos field value if set, zero value otherwise.
func (o *UploadVideoAlt1RequestEmbed) GetLogos() UploadVideoAlt1RequestEmbedLogos {
	if o == nil || IsNil(o.Logos) {
		var ret UploadVideoAlt1RequestEmbedLogos
		return ret
	}
	return *o.Logos
}

// GetLogosOk returns a tuple with the Logos field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UploadVideoAlt1RequestEmbed) GetLogosOk() (*UploadVideoAlt1RequestEmbedLogos, bool) {
	if o == nil || IsNil(o.Logos) {
		return nil, false
	}
	return o.Logos, true
}

// HasLogos returns a boolean if a field has been set.
func (o *UploadVideoAlt1RequestEmbed) HasLogos() bool {
	if o != nil && !IsNil(o.Logos) {
		return true
	}

	return false
}

// SetLogos gets a reference to the given UploadVideoAlt1RequestEmbedLogos and assigns it to the Logos field.
func (o *UploadVideoAlt1RequestEmbed) SetLogos(v UploadVideoAlt1RequestEmbedLogos) {
	o.Logos = &v
}

// GetPlaybar returns the Playbar field value if set, zero value otherwise.
func (o *UploadVideoAlt1RequestEmbed) GetPlaybar() bool {
	if o == nil || IsNil(o.Playbar) {
		var ret bool
		return ret
	}
	return *o.Playbar
}

// GetPlaybarOk returns a tuple with the Playbar field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UploadVideoAlt1RequestEmbed) GetPlaybarOk() (*bool, bool) {
	if o == nil || IsNil(o.Playbar) {
		return nil, false
	}
	return o.Playbar, true
}

// HasPlaybar returns a boolean if a field has been set.
func (o *UploadVideoAlt1RequestEmbed) HasPlaybar() bool {
	if o != nil && !IsNil(o.Playbar) {
		return true
	}

	return false
}

// SetPlaybar gets a reference to the given bool and assigns it to the Playbar field.
func (o *UploadVideoAlt1RequestEmbed) SetPlaybar(v bool) {
	o.Playbar = &v
}

// GetTitle returns the Title field value if set, zero value otherwise.
func (o *UploadVideoAlt1RequestEmbed) GetTitle() CreateEmbedPresetsAlt1RequestEmbedTitle {
	if o == nil || IsNil(o.Title) {
		var ret CreateEmbedPresetsAlt1RequestEmbedTitle
		return ret
	}
	return *o.Title
}

// GetTitleOk returns a tuple with the Title field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UploadVideoAlt1RequestEmbed) GetTitleOk() (*CreateEmbedPresetsAlt1RequestEmbedTitle, bool) {
	if o == nil || IsNil(o.Title) {
		return nil, false
	}
	return o.Title, true
}

// HasTitle returns a boolean if a field has been set.
func (o *UploadVideoAlt1RequestEmbed) HasTitle() bool {
	if o != nil && !IsNil(o.Title) {
		return true
	}

	return false
}

// SetTitle gets a reference to the given CreateEmbedPresetsAlt1RequestEmbedTitle and assigns it to the Title field.
func (o *UploadVideoAlt1RequestEmbed) SetTitle(v CreateEmbedPresetsAlt1RequestEmbedTitle) {
	o.Title = &v
}

// GetVolume returns the Volume field value if set, zero value otherwise.
func (o *UploadVideoAlt1RequestEmbed) GetVolume() bool {
	if o == nil || IsNil(o.Volume) {
		var ret bool
		return ret
	}
	return *o.Volume
}

// GetVolumeOk returns a tuple with the Volume field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UploadVideoAlt1RequestEmbed) GetVolumeOk() (*bool, bool) {
	if o == nil || IsNil(o.Volume) {
		return nil, false
	}
	return o.Volume, true
}

// HasVolume returns a boolean if a field has been set.
func (o *UploadVideoAlt1RequestEmbed) HasVolume() bool {
	if o != nil && !IsNil(o.Volume) {
		return true
	}

	return false
}

// SetVolume gets a reference to the given bool and assigns it to the Volume field.
func (o *UploadVideoAlt1RequestEmbed) SetVolume(v bool) {
	o.Volume = &v
}

func (o UploadVideoAlt1RequestEmbed) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o UploadVideoAlt1RequestEmbed) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Buttons) {
		toSerialize["buttons"] = o.Buttons
	}
	if !IsNil(o.Color) {
		toSerialize["color"] = o.Color
	}
	if !IsNil(o.EndScreen) {
		toSerialize["end_screen"] = o.EndScreen
	}
	if !IsNil(o.Logos) {
		toSerialize["logos"] = o.Logos
	}
	if !IsNil(o.Playbar) {
		toSerialize["playbar"] = o.Playbar
	}
	if !IsNil(o.Title) {
		toSerialize["title"] = o.Title
	}
	if !IsNil(o.Volume) {
		toSerialize["volume"] = o.Volume
	}
	return toSerialize, nil
}

type NullableUploadVideoAlt1RequestEmbed struct {
	value *UploadVideoAlt1RequestEmbed
	isSet bool
}

func (v NullableUploadVideoAlt1RequestEmbed) Get() *UploadVideoAlt1RequestEmbed {
	return v.value
}

func (v *NullableUploadVideoAlt1RequestEmbed) Set(val *UploadVideoAlt1RequestEmbed) {
	v.value = val
	v.isSet = true
}

func (v NullableUploadVideoAlt1RequestEmbed) IsSet() bool {
	return v.isSet
}

func (v *NullableUploadVideoAlt1RequestEmbed) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableUploadVideoAlt1RequestEmbed(val *UploadVideoAlt1RequestEmbed) *NullableUploadVideoAlt1RequestEmbed {
	return &NullableUploadVideoAlt1RequestEmbed{value: val, isSet: true}
}

func (v NullableUploadVideoAlt1RequestEmbed) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableUploadVideoAlt1RequestEmbed) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


