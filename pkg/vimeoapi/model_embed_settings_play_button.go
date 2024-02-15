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

// checks if the EmbedSettingsPlayButton type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &EmbedSettingsPlayButton{}

// EmbedSettingsPlayButton A representation of the play button's settings.
type EmbedSettingsPlayButton struct {
	// The position of the play button within the embeddable player.  Option descriptions:  * `auto` - Use Vimeo's default positioning for the play button.  * `bottom` - The play button is positioned at the bottom of the player, except when in tiny mode.  * `center` - The play button is positioned in the center of the player. 
	Position string `json:"position"`
}

type _EmbedSettingsPlayButton EmbedSettingsPlayButton

// NewEmbedSettingsPlayButton instantiates a new EmbedSettingsPlayButton object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewEmbedSettingsPlayButton(position string) *EmbedSettingsPlayButton {
	this := EmbedSettingsPlayButton{}
	this.Position = position
	return &this
}

// NewEmbedSettingsPlayButtonWithDefaults instantiates a new EmbedSettingsPlayButton object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewEmbedSettingsPlayButtonWithDefaults() *EmbedSettingsPlayButton {
	this := EmbedSettingsPlayButton{}
	return &this
}

// GetPosition returns the Position field value
func (o *EmbedSettingsPlayButton) GetPosition() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Position
}

// GetPositionOk returns a tuple with the Position field value
// and a boolean to check if the value has been set.
func (o *EmbedSettingsPlayButton) GetPositionOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Position, true
}

// SetPosition sets field value
func (o *EmbedSettingsPlayButton) SetPosition(v string) {
	o.Position = v
}

func (o EmbedSettingsPlayButton) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o EmbedSettingsPlayButton) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["position"] = o.Position
	return toSerialize, nil
}

func (o *EmbedSettingsPlayButton) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"position",
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

	varEmbedSettingsPlayButton := _EmbedSettingsPlayButton{}

	decoder := json.NewDecoder(bytes.NewReader(data))
	decoder.DisallowUnknownFields()
	err = decoder.Decode(&varEmbedSettingsPlayButton)

	if err != nil {
		return err
	}

	*o = EmbedSettingsPlayButton(varEmbedSettingsPlayButton)

	return err
}

type NullableEmbedSettingsPlayButton struct {
	value *EmbedSettingsPlayButton
	isSet bool
}

func (v NullableEmbedSettingsPlayButton) Get() *EmbedSettingsPlayButton {
	return v.value
}

func (v *NullableEmbedSettingsPlayButton) Set(val *EmbedSettingsPlayButton) {
	v.value = val
	v.isSet = true
}

func (v NullableEmbedSettingsPlayButton) IsSet() bool {
	return v.isSet
}

func (v *NullableEmbedSettingsPlayButton) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableEmbedSettingsPlayButton(val *EmbedSettingsPlayButton) *NullableEmbedSettingsPlayButton {
	return &NullableEmbedSettingsPlayButton{value: val, isSet: true}
}

func (v NullableEmbedSettingsPlayButton) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableEmbedSettingsPlayButton) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

