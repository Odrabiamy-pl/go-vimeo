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

// checks if the DisabledVideoPropertiesAddToCollection type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &DisabledVideoPropertiesAddToCollection{}

// DisabledVideoPropertiesAddToCollection An object that represents the reason why adding to a collection is disabled for the video.
type DisabledVideoPropertiesAddToCollection struct {
	// The relative link to upgrade to access adding to a collection.
	EnableLink string `json:"enable_link"`
	// The capability required to activate adding to a collection.  Option descriptions:  * `basic` - The user must have at least a Vimeo Basic account.  * `business` - The user must have at least a Vimeo Business account.  * `enterprise` - The user must have at least a Vimeo Enterprise account.  * `live_business` - The user must have at least a Vimeo Business Live account.  * `live_premium` - The user must have at least a Vimeo Premium account.  * `live_pro` - The user must have at least a Vimeo Pro Live account.  * `plus` - The user must have at least a Vimeo Plus account.  * `pro` - The user must have at least a Vimeo Pro account.  * `pro_unlimited` - The user must have at least a Vimeo Pro Unlimited account.  * `producer` - The user must have at least a Vimeo Producer account. 
	MinTierForCapability string `json:"min_tier_for_capability"`
	// The reasons why adding to a collection is disabled for the video.
	Reasons []DisabledVideoPropertiesAddToCollectionReasonsInner `json:"reasons"`
}

// NewDisabledVideoPropertiesAddToCollection instantiates a new DisabledVideoPropertiesAddToCollection object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewDisabledVideoPropertiesAddToCollection(enableLink string, minTierForCapability string, reasons []DisabledVideoPropertiesAddToCollectionReasonsInner) *DisabledVideoPropertiesAddToCollection {
	this := DisabledVideoPropertiesAddToCollection{}
	this.EnableLink = enableLink
	this.MinTierForCapability = minTierForCapability
	this.Reasons = reasons
	return &this
}

// NewDisabledVideoPropertiesAddToCollectionWithDefaults instantiates a new DisabledVideoPropertiesAddToCollection object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewDisabledVideoPropertiesAddToCollectionWithDefaults() *DisabledVideoPropertiesAddToCollection {
	this := DisabledVideoPropertiesAddToCollection{}
	return &this
}

// GetEnableLink returns the EnableLink field value
func (o *DisabledVideoPropertiesAddToCollection) GetEnableLink() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.EnableLink
}

// GetEnableLinkOk returns a tuple with the EnableLink field value
// and a boolean to check if the value has been set.
func (o *DisabledVideoPropertiesAddToCollection) GetEnableLinkOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.EnableLink, true
}

// SetEnableLink sets field value
func (o *DisabledVideoPropertiesAddToCollection) SetEnableLink(v string) {
	o.EnableLink = v
}

// GetMinTierForCapability returns the MinTierForCapability field value
func (o *DisabledVideoPropertiesAddToCollection) GetMinTierForCapability() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.MinTierForCapability
}

// GetMinTierForCapabilityOk returns a tuple with the MinTierForCapability field value
// and a boolean to check if the value has been set.
func (o *DisabledVideoPropertiesAddToCollection) GetMinTierForCapabilityOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.MinTierForCapability, true
}

// SetMinTierForCapability sets field value
func (o *DisabledVideoPropertiesAddToCollection) SetMinTierForCapability(v string) {
	o.MinTierForCapability = v
}

// GetReasons returns the Reasons field value
func (o *DisabledVideoPropertiesAddToCollection) GetReasons() []DisabledVideoPropertiesAddToCollectionReasonsInner {
	if o == nil {
		var ret []DisabledVideoPropertiesAddToCollectionReasonsInner
		return ret
	}

	return o.Reasons
}

// GetReasonsOk returns a tuple with the Reasons field value
// and a boolean to check if the value has been set.
func (o *DisabledVideoPropertiesAddToCollection) GetReasonsOk() ([]DisabledVideoPropertiesAddToCollectionReasonsInner, bool) {
	if o == nil {
		return nil, false
	}
	return o.Reasons, true
}

// SetReasons sets field value
func (o *DisabledVideoPropertiesAddToCollection) SetReasons(v []DisabledVideoPropertiesAddToCollectionReasonsInner) {
	o.Reasons = v
}

func (o DisabledVideoPropertiesAddToCollection) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o DisabledVideoPropertiesAddToCollection) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["enable_link"] = o.EnableLink
	toSerialize["min_tier_for_capability"] = o.MinTierForCapability
	toSerialize["reasons"] = o.Reasons
	return toSerialize, nil
}

type NullableDisabledVideoPropertiesAddToCollection struct {
	value *DisabledVideoPropertiesAddToCollection
	isSet bool
}

func (v NullableDisabledVideoPropertiesAddToCollection) Get() *DisabledVideoPropertiesAddToCollection {
	return v.value
}

func (v *NullableDisabledVideoPropertiesAddToCollection) Set(val *DisabledVideoPropertiesAddToCollection) {
	v.value = val
	v.isSet = true
}

func (v NullableDisabledVideoPropertiesAddToCollection) IsSet() bool {
	return v.isSet
}

func (v *NullableDisabledVideoPropertiesAddToCollection) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableDisabledVideoPropertiesAddToCollection(val *DisabledVideoPropertiesAddToCollection) *NullableDisabledVideoPropertiesAddToCollection {
	return &NullableDisabledVideoPropertiesAddToCollection{value: val, isSet: true}
}

func (v NullableDisabledVideoPropertiesAddToCollection) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableDisabledVideoPropertiesAddToCollection) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


