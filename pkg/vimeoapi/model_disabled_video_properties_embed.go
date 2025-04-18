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

// checks if the DisabledVideoPropertiesEmbed type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &DisabledVideoPropertiesEmbed{}

// DisabledVideoPropertiesEmbed An object that represents the reason why embed is disabled for the video.
type DisabledVideoPropertiesEmbed struct {
	// The relative link to upgrade embeds.
	EnableLink string `json:"enable_link"`
	// The path to the embed object in the video response.
	KeyPath string `json:"key_path"`
	// The capability required to activate embeds.  Option descriptions:  * `basic` - The user must have at least a Vimeo Basic account.  * `business` - The user must have at least a Vimeo Business account.  * `enterprise` - The user must have at least a Vimeo Enterprise account.  * `live_business` - The user must have at least a Vimeo Business Live account.  * `live_premium` - The user must have at least a Vimeo Premium account.  * `live_pro` - The user must have at least a Vimeo Pro Live account.  * `plus` - The user must have at least a Vimeo Plus account.  * `pro` - The user must have at least a Vimeo Pro account.  * `pro_unlimited` - The user must have at least a Vimeo Pro Unlimited account.  * `producer` - The user must have at least a Vimeo Producer account. 
	MinTierForCapability string `json:"min_tier_for_capability"`
	// The reasons why embed is disabled for the video.
	Reasons []DisabledVideoPropertiesEmbedReasonsInner `json:"reasons"`
}

// NewDisabledVideoPropertiesEmbed instantiates a new DisabledVideoPropertiesEmbed object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewDisabledVideoPropertiesEmbed(enableLink string, keyPath string, minTierForCapability string, reasons []DisabledVideoPropertiesEmbedReasonsInner) *DisabledVideoPropertiesEmbed {
	this := DisabledVideoPropertiesEmbed{}
	this.EnableLink = enableLink
	this.KeyPath = keyPath
	this.MinTierForCapability = minTierForCapability
	this.Reasons = reasons
	return &this
}

// NewDisabledVideoPropertiesEmbedWithDefaults instantiates a new DisabledVideoPropertiesEmbed object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewDisabledVideoPropertiesEmbedWithDefaults() *DisabledVideoPropertiesEmbed {
	this := DisabledVideoPropertiesEmbed{}
	return &this
}

// GetEnableLink returns the EnableLink field value
func (o *DisabledVideoPropertiesEmbed) GetEnableLink() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.EnableLink
}

// GetEnableLinkOk returns a tuple with the EnableLink field value
// and a boolean to check if the value has been set.
func (o *DisabledVideoPropertiesEmbed) GetEnableLinkOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.EnableLink, true
}

// SetEnableLink sets field value
func (o *DisabledVideoPropertiesEmbed) SetEnableLink(v string) {
	o.EnableLink = v
}

// GetKeyPath returns the KeyPath field value
func (o *DisabledVideoPropertiesEmbed) GetKeyPath() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.KeyPath
}

// GetKeyPathOk returns a tuple with the KeyPath field value
// and a boolean to check if the value has been set.
func (o *DisabledVideoPropertiesEmbed) GetKeyPathOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.KeyPath, true
}

// SetKeyPath sets field value
func (o *DisabledVideoPropertiesEmbed) SetKeyPath(v string) {
	o.KeyPath = v
}

// GetMinTierForCapability returns the MinTierForCapability field value
func (o *DisabledVideoPropertiesEmbed) GetMinTierForCapability() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.MinTierForCapability
}

// GetMinTierForCapabilityOk returns a tuple with the MinTierForCapability field value
// and a boolean to check if the value has been set.
func (o *DisabledVideoPropertiesEmbed) GetMinTierForCapabilityOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.MinTierForCapability, true
}

// SetMinTierForCapability sets field value
func (o *DisabledVideoPropertiesEmbed) SetMinTierForCapability(v string) {
	o.MinTierForCapability = v
}

// GetReasons returns the Reasons field value
func (o *DisabledVideoPropertiesEmbed) GetReasons() []DisabledVideoPropertiesEmbedReasonsInner {
	if o == nil {
		var ret []DisabledVideoPropertiesEmbedReasonsInner
		return ret
	}

	return o.Reasons
}

// GetReasonsOk returns a tuple with the Reasons field value
// and a boolean to check if the value has been set.
func (o *DisabledVideoPropertiesEmbed) GetReasonsOk() ([]DisabledVideoPropertiesEmbedReasonsInner, bool) {
	if o == nil {
		return nil, false
	}
	return o.Reasons, true
}

// SetReasons sets field value
func (o *DisabledVideoPropertiesEmbed) SetReasons(v []DisabledVideoPropertiesEmbedReasonsInner) {
	o.Reasons = v
}

func (o DisabledVideoPropertiesEmbed) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o DisabledVideoPropertiesEmbed) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["enable_link"] = o.EnableLink
	toSerialize["key_path"] = o.KeyPath
	toSerialize["min_tier_for_capability"] = o.MinTierForCapability
	toSerialize["reasons"] = o.Reasons
	return toSerialize, nil
}

type NullableDisabledVideoPropertiesEmbed struct {
	value *DisabledVideoPropertiesEmbed
	isSet bool
}

func (v NullableDisabledVideoPropertiesEmbed) Get() *DisabledVideoPropertiesEmbed {
	return v.value
}

func (v *NullableDisabledVideoPropertiesEmbed) Set(val *DisabledVideoPropertiesEmbed) {
	v.value = val
	v.isSet = true
}

func (v NullableDisabledVideoPropertiesEmbed) IsSet() bool {
	return v.isSet
}

func (v *NullableDisabledVideoPropertiesEmbed) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableDisabledVideoPropertiesEmbed(val *DisabledVideoPropertiesEmbed) *NullableDisabledVideoPropertiesEmbed {
	return &NullableDisabledVideoPropertiesEmbed{value: val, isSet: true}
}

func (v NullableDisabledVideoPropertiesEmbed) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableDisabledVideoPropertiesEmbed) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


