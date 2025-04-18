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

// checks if the CreateVodAlt1RequestEpisodesBuy type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &CreateVodAlt1RequestEpisodesBuy{}

// CreateVodAlt1RequestEpisodesBuy struct for CreateVodAlt1RequestEpisodesBuy
type CreateVodAlt1RequestEpisodesBuy struct {
	// Whether episodes can be purchased.
	Active *bool `json:"active,omitempty"`
	// Whether people who buy episodes can download them. To use this parameter, **type** must be `series`.
	Download *bool `json:"download,omitempty"`
	Price *CreateVodAlt1RequestEpisodesBuyPrice `json:"price,omitempty"`
}

// NewCreateVodAlt1RequestEpisodesBuy instantiates a new CreateVodAlt1RequestEpisodesBuy object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewCreateVodAlt1RequestEpisodesBuy() *CreateVodAlt1RequestEpisodesBuy {
	this := CreateVodAlt1RequestEpisodesBuy{}
	return &this
}

// NewCreateVodAlt1RequestEpisodesBuyWithDefaults instantiates a new CreateVodAlt1RequestEpisodesBuy object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewCreateVodAlt1RequestEpisodesBuyWithDefaults() *CreateVodAlt1RequestEpisodesBuy {
	this := CreateVodAlt1RequestEpisodesBuy{}
	return &this
}

// GetActive returns the Active field value if set, zero value otherwise.
func (o *CreateVodAlt1RequestEpisodesBuy) GetActive() bool {
	if o == nil || IsNil(o.Active) {
		var ret bool
		return ret
	}
	return *o.Active
}

// GetActiveOk returns a tuple with the Active field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CreateVodAlt1RequestEpisodesBuy) GetActiveOk() (*bool, bool) {
	if o == nil || IsNil(o.Active) {
		return nil, false
	}
	return o.Active, true
}

// HasActive returns a boolean if a field has been set.
func (o *CreateVodAlt1RequestEpisodesBuy) HasActive() bool {
	if o != nil && !IsNil(o.Active) {
		return true
	}

	return false
}

// SetActive gets a reference to the given bool and assigns it to the Active field.
func (o *CreateVodAlt1RequestEpisodesBuy) SetActive(v bool) {
	o.Active = &v
}

// GetDownload returns the Download field value if set, zero value otherwise.
func (o *CreateVodAlt1RequestEpisodesBuy) GetDownload() bool {
	if o == nil || IsNil(o.Download) {
		var ret bool
		return ret
	}
	return *o.Download
}

// GetDownloadOk returns a tuple with the Download field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CreateVodAlt1RequestEpisodesBuy) GetDownloadOk() (*bool, bool) {
	if o == nil || IsNil(o.Download) {
		return nil, false
	}
	return o.Download, true
}

// HasDownload returns a boolean if a field has been set.
func (o *CreateVodAlt1RequestEpisodesBuy) HasDownload() bool {
	if o != nil && !IsNil(o.Download) {
		return true
	}

	return false
}

// SetDownload gets a reference to the given bool and assigns it to the Download field.
func (o *CreateVodAlt1RequestEpisodesBuy) SetDownload(v bool) {
	o.Download = &v
}

// GetPrice returns the Price field value if set, zero value otherwise.
func (o *CreateVodAlt1RequestEpisodesBuy) GetPrice() CreateVodAlt1RequestEpisodesBuyPrice {
	if o == nil || IsNil(o.Price) {
		var ret CreateVodAlt1RequestEpisodesBuyPrice
		return ret
	}
	return *o.Price
}

// GetPriceOk returns a tuple with the Price field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CreateVodAlt1RequestEpisodesBuy) GetPriceOk() (*CreateVodAlt1RequestEpisodesBuyPrice, bool) {
	if o == nil || IsNil(o.Price) {
		return nil, false
	}
	return o.Price, true
}

// HasPrice returns a boolean if a field has been set.
func (o *CreateVodAlt1RequestEpisodesBuy) HasPrice() bool {
	if o != nil && !IsNil(o.Price) {
		return true
	}

	return false
}

// SetPrice gets a reference to the given CreateVodAlt1RequestEpisodesBuyPrice and assigns it to the Price field.
func (o *CreateVodAlt1RequestEpisodesBuy) SetPrice(v CreateVodAlt1RequestEpisodesBuyPrice) {
	o.Price = &v
}

func (o CreateVodAlt1RequestEpisodesBuy) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o CreateVodAlt1RequestEpisodesBuy) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Active) {
		toSerialize["active"] = o.Active
	}
	if !IsNil(o.Download) {
		toSerialize["download"] = o.Download
	}
	if !IsNil(o.Price) {
		toSerialize["price"] = o.Price
	}
	return toSerialize, nil
}

type NullableCreateVodAlt1RequestEpisodesBuy struct {
	value *CreateVodAlt1RequestEpisodesBuy
	isSet bool
}

func (v NullableCreateVodAlt1RequestEpisodesBuy) Get() *CreateVodAlt1RequestEpisodesBuy {
	return v.value
}

func (v *NullableCreateVodAlt1RequestEpisodesBuy) Set(val *CreateVodAlt1RequestEpisodesBuy) {
	v.value = val
	v.isSet = true
}

func (v NullableCreateVodAlt1RequestEpisodesBuy) IsSet() bool {
	return v.isSet
}

func (v *NullableCreateVodAlt1RequestEpisodesBuy) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableCreateVodAlt1RequestEpisodesBuy(val *CreateVodAlt1RequestEpisodesBuy) *NullableCreateVodAlt1RequestEpisodesBuy {
	return &NullableCreateVodAlt1RequestEpisodesBuy{value: val, isSet: true}
}

func (v NullableCreateVodAlt1RequestEpisodesBuy) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableCreateVodAlt1RequestEpisodesBuy) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


