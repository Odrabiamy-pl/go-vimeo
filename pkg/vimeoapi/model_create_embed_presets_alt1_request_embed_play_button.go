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

// checks if the CreateEmbedPresetsAlt1RequestEmbedPlayButton type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &CreateEmbedPresetsAlt1RequestEmbedPlayButton{}

// CreateEmbedPresetsAlt1RequestEmbedPlayButton An object representing the play button's settings.
type CreateEmbedPresetsAlt1RequestEmbedPlayButton struct {
	// The position of the play button within the embeddable player.  Option descriptions:  * `auto` - Use Vimeo's default positioning for the play button.  * `bottom` - The play button is positioned at the bottom of the player, except when in tiny mode.  * `center` - The play button is positioned in the center of the player.
	Position *string `json:"position,omitempty"`
}

// NewCreateEmbedPresetsAlt1RequestEmbedPlayButton instantiates a new CreateEmbedPresetsAlt1RequestEmbedPlayButton object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewCreateEmbedPresetsAlt1RequestEmbedPlayButton() *CreateEmbedPresetsAlt1RequestEmbedPlayButton {
	this := CreateEmbedPresetsAlt1RequestEmbedPlayButton{}
	return &this
}

// NewCreateEmbedPresetsAlt1RequestEmbedPlayButtonWithDefaults instantiates a new CreateEmbedPresetsAlt1RequestEmbedPlayButton object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewCreateEmbedPresetsAlt1RequestEmbedPlayButtonWithDefaults() *CreateEmbedPresetsAlt1RequestEmbedPlayButton {
	this := CreateEmbedPresetsAlt1RequestEmbedPlayButton{}
	return &this
}

// GetPosition returns the Position field value if set, zero value otherwise.
func (o *CreateEmbedPresetsAlt1RequestEmbedPlayButton) GetPosition() string {
	if o == nil || IsNil(o.Position) {
		var ret string
		return ret
	}
	return *o.Position
}

// GetPositionOk returns a tuple with the Position field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CreateEmbedPresetsAlt1RequestEmbedPlayButton) GetPositionOk() (*string, bool) {
	if o == nil || IsNil(o.Position) {
		return nil, false
	}
	return o.Position, true
}

// HasPosition returns a boolean if a field has been set.
func (o *CreateEmbedPresetsAlt1RequestEmbedPlayButton) HasPosition() bool {
	if o != nil && !IsNil(o.Position) {
		return true
	}

	return false
}

// SetPosition gets a reference to the given string and assigns it to the Position field.
func (o *CreateEmbedPresetsAlt1RequestEmbedPlayButton) SetPosition(v string) {
	o.Position = &v
}

func (o CreateEmbedPresetsAlt1RequestEmbedPlayButton) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o CreateEmbedPresetsAlt1RequestEmbedPlayButton) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Position) {
		toSerialize["position"] = o.Position
	}
	return toSerialize, nil
}

type NullableCreateEmbedPresetsAlt1RequestEmbedPlayButton struct {
	value *CreateEmbedPresetsAlt1RequestEmbedPlayButton
	isSet bool
}

func (v NullableCreateEmbedPresetsAlt1RequestEmbedPlayButton) Get() *CreateEmbedPresetsAlt1RequestEmbedPlayButton {
	return v.value
}

func (v *NullableCreateEmbedPresetsAlt1RequestEmbedPlayButton) Set(val *CreateEmbedPresetsAlt1RequestEmbedPlayButton) {
	v.value = val
	v.isSet = true
}

func (v NullableCreateEmbedPresetsAlt1RequestEmbedPlayButton) IsSet() bool {
	return v.isSet
}

func (v *NullableCreateEmbedPresetsAlt1RequestEmbedPlayButton) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableCreateEmbedPresetsAlt1RequestEmbedPlayButton(val *CreateEmbedPresetsAlt1RequestEmbedPlayButton) *NullableCreateEmbedPresetsAlt1RequestEmbedPlayButton {
	return &NullableCreateEmbedPresetsAlt1RequestEmbedPlayButton{value: val, isSet: true}
}

func (v NullableCreateEmbedPresetsAlt1RequestEmbedPlayButton) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableCreateEmbedPresetsAlt1RequestEmbedPlayButton) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
