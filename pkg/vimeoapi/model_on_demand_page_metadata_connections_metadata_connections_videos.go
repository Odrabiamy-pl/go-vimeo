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

// checks if the OnDemandPageMetadataConnectionsMetadataConnectionsVideos type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &OnDemandPageMetadataConnectionsMetadataConnectionsVideos{}

// OnDemandPageMetadataConnectionsMetadataConnectionsVideos Information about the videos associated with the On Demand page.
type OnDemandPageMetadataConnectionsMetadataConnectionsVideos struct {
	// The total number of extra videos on the On Demand page.
	ExtraTotal float32 `json:"extra_total"`
	// The total number of main videos on the On Demand page.
	MainTotal float32 `json:"main_total"`
	// An array of HTTP methods permitted on this URI.
	Options []string `json:"options"`
	// The total number of videos on this connection.
	Total float32 `json:"total"`
	// The API URI that resolves to the connection data.
	Uri string `json:"uri"`
	// The total number of viewable videos on the On Demand page.
	ViewableTotal float32 `json:"viewable_total"`
}

// NewOnDemandPageMetadataConnectionsMetadataConnectionsVideos instantiates a new OnDemandPageMetadataConnectionsMetadataConnectionsVideos object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewOnDemandPageMetadataConnectionsMetadataConnectionsVideos(extraTotal float32, mainTotal float32, options []string, total float32, uri string, viewableTotal float32) *OnDemandPageMetadataConnectionsMetadataConnectionsVideos {
	this := OnDemandPageMetadataConnectionsMetadataConnectionsVideos{}
	this.ExtraTotal = extraTotal
	this.MainTotal = mainTotal
	this.Options = options
	this.Total = total
	this.Uri = uri
	this.ViewableTotal = viewableTotal
	return &this
}

// NewOnDemandPageMetadataConnectionsMetadataConnectionsVideosWithDefaults instantiates a new OnDemandPageMetadataConnectionsMetadataConnectionsVideos object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewOnDemandPageMetadataConnectionsMetadataConnectionsVideosWithDefaults() *OnDemandPageMetadataConnectionsMetadataConnectionsVideos {
	this := OnDemandPageMetadataConnectionsMetadataConnectionsVideos{}
	return &this
}

// GetExtraTotal returns the ExtraTotal field value
func (o *OnDemandPageMetadataConnectionsMetadataConnectionsVideos) GetExtraTotal() float32 {
	if o == nil {
		var ret float32
		return ret
	}

	return o.ExtraTotal
}

// GetExtraTotalOk returns a tuple with the ExtraTotal field value
// and a boolean to check if the value has been set.
func (o *OnDemandPageMetadataConnectionsMetadataConnectionsVideos) GetExtraTotalOk() (*float32, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ExtraTotal, true
}

// SetExtraTotal sets field value
func (o *OnDemandPageMetadataConnectionsMetadataConnectionsVideos) SetExtraTotal(v float32) {
	o.ExtraTotal = v
}

// GetMainTotal returns the MainTotal field value
func (o *OnDemandPageMetadataConnectionsMetadataConnectionsVideos) GetMainTotal() float32 {
	if o == nil {
		var ret float32
		return ret
	}

	return o.MainTotal
}

// GetMainTotalOk returns a tuple with the MainTotal field value
// and a boolean to check if the value has been set.
func (o *OnDemandPageMetadataConnectionsMetadataConnectionsVideos) GetMainTotalOk() (*float32, bool) {
	if o == nil {
		return nil, false
	}
	return &o.MainTotal, true
}

// SetMainTotal sets field value
func (o *OnDemandPageMetadataConnectionsMetadataConnectionsVideos) SetMainTotal(v float32) {
	o.MainTotal = v
}

// GetOptions returns the Options field value
func (o *OnDemandPageMetadataConnectionsMetadataConnectionsVideos) GetOptions() []string {
	if o == nil {
		var ret []string
		return ret
	}

	return o.Options
}

// GetOptionsOk returns a tuple with the Options field value
// and a boolean to check if the value has been set.
func (o *OnDemandPageMetadataConnectionsMetadataConnectionsVideos) GetOptionsOk() ([]string, bool) {
	if o == nil {
		return nil, false
	}
	return o.Options, true
}

// SetOptions sets field value
func (o *OnDemandPageMetadataConnectionsMetadataConnectionsVideos) SetOptions(v []string) {
	o.Options = v
}

// GetTotal returns the Total field value
func (o *OnDemandPageMetadataConnectionsMetadataConnectionsVideos) GetTotal() float32 {
	if o == nil {
		var ret float32
		return ret
	}

	return o.Total
}

// GetTotalOk returns a tuple with the Total field value
// and a boolean to check if the value has been set.
func (o *OnDemandPageMetadataConnectionsMetadataConnectionsVideos) GetTotalOk() (*float32, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Total, true
}

// SetTotal sets field value
func (o *OnDemandPageMetadataConnectionsMetadataConnectionsVideos) SetTotal(v float32) {
	o.Total = v
}

// GetUri returns the Uri field value
func (o *OnDemandPageMetadataConnectionsMetadataConnectionsVideos) GetUri() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Uri
}

// GetUriOk returns a tuple with the Uri field value
// and a boolean to check if the value has been set.
func (o *OnDemandPageMetadataConnectionsMetadataConnectionsVideos) GetUriOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Uri, true
}

// SetUri sets field value
func (o *OnDemandPageMetadataConnectionsMetadataConnectionsVideos) SetUri(v string) {
	o.Uri = v
}

// GetViewableTotal returns the ViewableTotal field value
func (o *OnDemandPageMetadataConnectionsMetadataConnectionsVideos) GetViewableTotal() float32 {
	if o == nil {
		var ret float32
		return ret
	}

	return o.ViewableTotal
}

// GetViewableTotalOk returns a tuple with the ViewableTotal field value
// and a boolean to check if the value has been set.
func (o *OnDemandPageMetadataConnectionsMetadataConnectionsVideos) GetViewableTotalOk() (*float32, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ViewableTotal, true
}

// SetViewableTotal sets field value
func (o *OnDemandPageMetadataConnectionsMetadataConnectionsVideos) SetViewableTotal(v float32) {
	o.ViewableTotal = v
}

func (o OnDemandPageMetadataConnectionsMetadataConnectionsVideos) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o OnDemandPageMetadataConnectionsMetadataConnectionsVideos) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["extra_total"] = o.ExtraTotal
	toSerialize["main_total"] = o.MainTotal
	toSerialize["options"] = o.Options
	toSerialize["total"] = o.Total
	toSerialize["uri"] = o.Uri
	toSerialize["viewable_total"] = o.ViewableTotal
	return toSerialize, nil
}

type NullableOnDemandPageMetadataConnectionsMetadataConnectionsVideos struct {
	value *OnDemandPageMetadataConnectionsMetadataConnectionsVideos
	isSet bool
}

func (v NullableOnDemandPageMetadataConnectionsMetadataConnectionsVideos) Get() *OnDemandPageMetadataConnectionsMetadataConnectionsVideos {
	return v.value
}

func (v *NullableOnDemandPageMetadataConnectionsMetadataConnectionsVideos) Set(val *OnDemandPageMetadataConnectionsMetadataConnectionsVideos) {
	v.value = val
	v.isSet = true
}

func (v NullableOnDemandPageMetadataConnectionsMetadataConnectionsVideos) IsSet() bool {
	return v.isSet
}

func (v *NullableOnDemandPageMetadataConnectionsMetadataConnectionsVideos) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableOnDemandPageMetadataConnectionsMetadataConnectionsVideos(val *OnDemandPageMetadataConnectionsMetadataConnectionsVideos) *NullableOnDemandPageMetadataConnectionsMetadataConnectionsVideos {
	return &NullableOnDemandPageMetadataConnectionsMetadataConnectionsVideos{value: val, isSet: true}
}

func (v NullableOnDemandPageMetadataConnectionsMetadataConnectionsVideos) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableOnDemandPageMetadataConnectionsMetadataConnectionsVideos) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


