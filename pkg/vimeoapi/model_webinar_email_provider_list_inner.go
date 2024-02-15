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

// checks if the WebinarEmailProviderListInner type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &WebinarEmailProviderListInner{}

// WebinarEmailProviderListInner struct for WebinarEmailProviderListInner
type WebinarEmailProviderListInner struct {
	// Whether the connection is active.
	IsActive bool `json:"is_active"`
	// The most recent sync date of the provider list.
	LastImportTime NullableString `json:"last_import_time"`
	List NullableWebinarEmailProviderListInnerList `json:"list"`
	Provider WebinarEmailProviderListInnerProvider `json:"provider"`
}

type _WebinarEmailProviderListInner WebinarEmailProviderListInner

// NewWebinarEmailProviderListInner instantiates a new WebinarEmailProviderListInner object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewWebinarEmailProviderListInner(isActive bool, lastImportTime NullableString, list NullableWebinarEmailProviderListInnerList, provider WebinarEmailProviderListInnerProvider) *WebinarEmailProviderListInner {
	this := WebinarEmailProviderListInner{}
	this.IsActive = isActive
	this.LastImportTime = lastImportTime
	this.List = list
	this.Provider = provider
	return &this
}

// NewWebinarEmailProviderListInnerWithDefaults instantiates a new WebinarEmailProviderListInner object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewWebinarEmailProviderListInnerWithDefaults() *WebinarEmailProviderListInner {
	this := WebinarEmailProviderListInner{}
	return &this
}

// GetIsActive returns the IsActive field value
func (o *WebinarEmailProviderListInner) GetIsActive() bool {
	if o == nil {
		var ret bool
		return ret
	}

	return o.IsActive
}

// GetIsActiveOk returns a tuple with the IsActive field value
// and a boolean to check if the value has been set.
func (o *WebinarEmailProviderListInner) GetIsActiveOk() (*bool, bool) {
	if o == nil {
		return nil, false
	}
	return &o.IsActive, true
}

// SetIsActive sets field value
func (o *WebinarEmailProviderListInner) SetIsActive(v bool) {
	o.IsActive = v
}

// GetLastImportTime returns the LastImportTime field value
// If the value is explicit nil, the zero value for string will be returned
func (o *WebinarEmailProviderListInner) GetLastImportTime() string {
	if o == nil || o.LastImportTime.Get() == nil {
		var ret string
		return ret
	}

	return *o.LastImportTime.Get()
}

// GetLastImportTimeOk returns a tuple with the LastImportTime field value
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *WebinarEmailProviderListInner) GetLastImportTimeOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.LastImportTime.Get(), o.LastImportTime.IsSet()
}

// SetLastImportTime sets field value
func (o *WebinarEmailProviderListInner) SetLastImportTime(v string) {
	o.LastImportTime.Set(&v)
}

// GetList returns the List field value
// If the value is explicit nil, the zero value for WebinarEmailProviderListInnerList will be returned
func (o *WebinarEmailProviderListInner) GetList() WebinarEmailProviderListInnerList {
	if o == nil || o.List.Get() == nil {
		var ret WebinarEmailProviderListInnerList
		return ret
	}

	return *o.List.Get()
}

// GetListOk returns a tuple with the List field value
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *WebinarEmailProviderListInner) GetListOk() (*WebinarEmailProviderListInnerList, bool) {
	if o == nil {
		return nil, false
	}
	return o.List.Get(), o.List.IsSet()
}

// SetList sets field value
func (o *WebinarEmailProviderListInner) SetList(v WebinarEmailProviderListInnerList) {
	o.List.Set(&v)
}

// GetProvider returns the Provider field value
func (o *WebinarEmailProviderListInner) GetProvider() WebinarEmailProviderListInnerProvider {
	if o == nil {
		var ret WebinarEmailProviderListInnerProvider
		return ret
	}

	return o.Provider
}

// GetProviderOk returns a tuple with the Provider field value
// and a boolean to check if the value has been set.
func (o *WebinarEmailProviderListInner) GetProviderOk() (*WebinarEmailProviderListInnerProvider, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Provider, true
}

// SetProvider sets field value
func (o *WebinarEmailProviderListInner) SetProvider(v WebinarEmailProviderListInnerProvider) {
	o.Provider = v
}

func (o WebinarEmailProviderListInner) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o WebinarEmailProviderListInner) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["is_active"] = o.IsActive
	toSerialize["last_import_time"] = o.LastImportTime.Get()
	toSerialize["list"] = o.List.Get()
	toSerialize["provider"] = o.Provider
	return toSerialize, nil
}

func (o *WebinarEmailProviderListInner) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"is_active",
		"last_import_time",
		"list",
		"provider",
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

	varWebinarEmailProviderListInner := _WebinarEmailProviderListInner{}

	decoder := json.NewDecoder(bytes.NewReader(data))
	decoder.DisallowUnknownFields()
	err = decoder.Decode(&varWebinarEmailProviderListInner)

	if err != nil {
		return err
	}

	*o = WebinarEmailProviderListInner(varWebinarEmailProviderListInner)

	return err
}

type NullableWebinarEmailProviderListInner struct {
	value *WebinarEmailProviderListInner
	isSet bool
}

func (v NullableWebinarEmailProviderListInner) Get() *WebinarEmailProviderListInner {
	return v.value
}

func (v *NullableWebinarEmailProviderListInner) Set(val *WebinarEmailProviderListInner) {
	v.value = val
	v.isSet = true
}

func (v NullableWebinarEmailProviderListInner) IsSet() bool {
	return v.isSet
}

func (v *NullableWebinarEmailProviderListInner) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableWebinarEmailProviderListInner(val *WebinarEmailProviderListInner) *NullableWebinarEmailProviderListInner {
	return &NullableWebinarEmailProviderListInner{value: val, isSet: true}
}

func (v NullableWebinarEmailProviderListInner) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableWebinarEmailProviderListInner) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

