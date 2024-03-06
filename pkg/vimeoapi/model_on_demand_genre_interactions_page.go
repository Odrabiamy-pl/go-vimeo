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

// checks if the OnDemandGenreInteractionsPage type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &OnDemandGenreInteractionsPage{}

// OnDemandGenreInteractionsPage Interactions for On Demand pages that belong to the genre.
type OnDemandGenreInteractionsPage struct {
	// Whether the On Demand genre was added.
	Added bool `json:"added"`
	// An array of HTTP methods permitted on this URI.
	Options []string `json:"options"`
	// The URI to access the On Demand page.
	Uri string `json:"uri"`
}

// NewOnDemandGenreInteractionsPage instantiates a new OnDemandGenreInteractionsPage object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewOnDemandGenreInteractionsPage(added bool, options []string, uri string) *OnDemandGenreInteractionsPage {
	this := OnDemandGenreInteractionsPage{}
	this.Added = added
	this.Options = options
	this.Uri = uri
	return &this
}

// NewOnDemandGenreInteractionsPageWithDefaults instantiates a new OnDemandGenreInteractionsPage object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewOnDemandGenreInteractionsPageWithDefaults() *OnDemandGenreInteractionsPage {
	this := OnDemandGenreInteractionsPage{}
	return &this
}

// GetAdded returns the Added field value
func (o *OnDemandGenreInteractionsPage) GetAdded() bool {
	if o == nil {
		var ret bool
		return ret
	}

	return o.Added
}

// GetAddedOk returns a tuple with the Added field value
// and a boolean to check if the value has been set.
func (o *OnDemandGenreInteractionsPage) GetAddedOk() (*bool, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Added, true
}

// SetAdded sets field value
func (o *OnDemandGenreInteractionsPage) SetAdded(v bool) {
	o.Added = v
}

// GetOptions returns the Options field value
func (o *OnDemandGenreInteractionsPage) GetOptions() []string {
	if o == nil {
		var ret []string
		return ret
	}

	return o.Options
}

// GetOptionsOk returns a tuple with the Options field value
// and a boolean to check if the value has been set.
func (o *OnDemandGenreInteractionsPage) GetOptionsOk() ([]string, bool) {
	if o == nil {
		return nil, false
	}
	return o.Options, true
}

// SetOptions sets field value
func (o *OnDemandGenreInteractionsPage) SetOptions(v []string) {
	o.Options = v
}

// GetUri returns the Uri field value
func (o *OnDemandGenreInteractionsPage) GetUri() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Uri
}

// GetUriOk returns a tuple with the Uri field value
// and a boolean to check if the value has been set.
func (o *OnDemandGenreInteractionsPage) GetUriOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Uri, true
}

// SetUri sets field value
func (o *OnDemandGenreInteractionsPage) SetUri(v string) {
	o.Uri = v
}

func (o OnDemandGenreInteractionsPage) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o OnDemandGenreInteractionsPage) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["added"] = o.Added
	toSerialize["options"] = o.Options
	toSerialize["uri"] = o.Uri
	return toSerialize, nil
}

type NullableOnDemandGenreInteractionsPage struct {
	value *OnDemandGenreInteractionsPage
	isSet bool
}

func (v NullableOnDemandGenreInteractionsPage) Get() *OnDemandGenreInteractionsPage {
	return v.value
}

func (v *NullableOnDemandGenreInteractionsPage) Set(val *OnDemandGenreInteractionsPage) {
	v.value = val
	v.isSet = true
}

func (v NullableOnDemandGenreInteractionsPage) IsSet() bool {
	return v.isSet
}

func (v *NullableOnDemandGenreInteractionsPage) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableOnDemandGenreInteractionsPage(val *OnDemandGenreInteractionsPage) *NullableOnDemandGenreInteractionsPage {
	return &NullableOnDemandGenreInteractionsPage{value: val, isSet: true}
}

func (v NullableOnDemandGenreInteractionsPage) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableOnDemandGenreInteractionsPage) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
