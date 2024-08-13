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

// checks if the DisabledVideoPropertiesDownload type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &DisabledVideoPropertiesDownload{}

// DisabledVideoPropertiesDownload An object that represents the reason why download is disabled for the video.
type DisabledVideoPropertiesDownload struct {
	// The relative link to upgrade downloads.
	EnableLink string `json:"enable_link"`
	// The path to the download object in the video response.
	KeyPath string `json:"key_path"`
	// The capability required to activate downloads.  Option descriptions:  * `basic` - The user must have at least a Vimeo Basic account.  * `business` - The user must have at least a Vimeo Business account.  * `enterprise` - The user must have at least a Vimeo Enterprise account.  * `live_business` - The user must have at least a Vimeo Business Live account.  * `live_premium` - The user must have at least a Vimeo Premium account.  * `live_pro` - The user must have at least a Vimeo Pro Live account.  * `plus` - The user must have at least a Vimeo Plus account.  * `pro` - The user must have at least a Vimeo Pro account.  * `pro_unlimited` - The user must have at least a Vimeo Pro Unlimited account.  * `producer` - The user must have at least a Vimeo Producer account. 
	MinTierForCapability string `json:"min_tier_for_capability"`
	// The reasons why download is disabled for the video.
	Reasons []DisabledVideoPropertiesDownloadReasonsInner `json:"reasons"`
}

// NewDisabledVideoPropertiesDownload instantiates a new DisabledVideoPropertiesDownload object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewDisabledVideoPropertiesDownload(enableLink string, keyPath string, minTierForCapability string, reasons []DisabledVideoPropertiesDownloadReasonsInner) *DisabledVideoPropertiesDownload {
	this := DisabledVideoPropertiesDownload{}
	this.EnableLink = enableLink
	this.KeyPath = keyPath
	this.MinTierForCapability = minTierForCapability
	this.Reasons = reasons
	return &this
}

// NewDisabledVideoPropertiesDownloadWithDefaults instantiates a new DisabledVideoPropertiesDownload object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewDisabledVideoPropertiesDownloadWithDefaults() *DisabledVideoPropertiesDownload {
	this := DisabledVideoPropertiesDownload{}
	return &this
}

// GetEnableLink returns the EnableLink field value
func (o *DisabledVideoPropertiesDownload) GetEnableLink() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.EnableLink
}

// GetEnableLinkOk returns a tuple with the EnableLink field value
// and a boolean to check if the value has been set.
func (o *DisabledVideoPropertiesDownload) GetEnableLinkOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.EnableLink, true
}

// SetEnableLink sets field value
func (o *DisabledVideoPropertiesDownload) SetEnableLink(v string) {
	o.EnableLink = v
}

// GetKeyPath returns the KeyPath field value
func (o *DisabledVideoPropertiesDownload) GetKeyPath() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.KeyPath
}

// GetKeyPathOk returns a tuple with the KeyPath field value
// and a boolean to check if the value has been set.
func (o *DisabledVideoPropertiesDownload) GetKeyPathOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.KeyPath, true
}

// SetKeyPath sets field value
func (o *DisabledVideoPropertiesDownload) SetKeyPath(v string) {
	o.KeyPath = v
}

// GetMinTierForCapability returns the MinTierForCapability field value
func (o *DisabledVideoPropertiesDownload) GetMinTierForCapability() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.MinTierForCapability
}

// GetMinTierForCapabilityOk returns a tuple with the MinTierForCapability field value
// and a boolean to check if the value has been set.
func (o *DisabledVideoPropertiesDownload) GetMinTierForCapabilityOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.MinTierForCapability, true
}

// SetMinTierForCapability sets field value
func (o *DisabledVideoPropertiesDownload) SetMinTierForCapability(v string) {
	o.MinTierForCapability = v
}

// GetReasons returns the Reasons field value
func (o *DisabledVideoPropertiesDownload) GetReasons() []DisabledVideoPropertiesDownloadReasonsInner {
	if o == nil {
		var ret []DisabledVideoPropertiesDownloadReasonsInner
		return ret
	}

	return o.Reasons
}

// GetReasonsOk returns a tuple with the Reasons field value
// and a boolean to check if the value has been set.
func (o *DisabledVideoPropertiesDownload) GetReasonsOk() ([]DisabledVideoPropertiesDownloadReasonsInner, bool) {
	if o == nil {
		return nil, false
	}
	return o.Reasons, true
}

// SetReasons sets field value
func (o *DisabledVideoPropertiesDownload) SetReasons(v []DisabledVideoPropertiesDownloadReasonsInner) {
	o.Reasons = v
}

func (o DisabledVideoPropertiesDownload) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o DisabledVideoPropertiesDownload) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["enable_link"] = o.EnableLink
	toSerialize["key_path"] = o.KeyPath
	toSerialize["min_tier_for_capability"] = o.MinTierForCapability
	toSerialize["reasons"] = o.Reasons
	return toSerialize, nil
}

type NullableDisabledVideoPropertiesDownload struct {
	value *DisabledVideoPropertiesDownload
	isSet bool
}

func (v NullableDisabledVideoPropertiesDownload) Get() *DisabledVideoPropertiesDownload {
	return v.value
}

func (v *NullableDisabledVideoPropertiesDownload) Set(val *DisabledVideoPropertiesDownload) {
	v.value = val
	v.isSet = true
}

func (v NullableDisabledVideoPropertiesDownload) IsSet() bool {
	return v.isSet
}

func (v *NullableDisabledVideoPropertiesDownload) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableDisabledVideoPropertiesDownload(val *DisabledVideoPropertiesDownload) *NullableDisabledVideoPropertiesDownload {
	return &NullableDisabledVideoPropertiesDownload{value: val, isSet: true}
}

func (v NullableDisabledVideoPropertiesDownload) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableDisabledVideoPropertiesDownload) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


