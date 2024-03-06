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

// checks if the OnDemandPageMetadataConnectionsMetadataConnectionsComments type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &OnDemandPageMetadataConnectionsMetadataConnectionsComments{}

// OnDemandPageMetadataConnectionsMetadataConnectionsComments Information about the comments associated with the On Demand page.
type OnDemandPageMetadataConnectionsMetadataConnectionsComments struct {
	// An array of HTTP methods permitted on this URI.
	Options []string `json:"options"`
	// The total number of comments on this connection.
	Total float32 `json:"total"`
	// The API URI that resolves to the connection data.
	Uri string `json:"uri"`
}

// NewOnDemandPageMetadataConnectionsMetadataConnectionsComments instantiates a new OnDemandPageMetadataConnectionsMetadataConnectionsComments object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewOnDemandPageMetadataConnectionsMetadataConnectionsComments(options []string, total float32, uri string) *OnDemandPageMetadataConnectionsMetadataConnectionsComments {
	this := OnDemandPageMetadataConnectionsMetadataConnectionsComments{}
	this.Options = options
	this.Total = total
	this.Uri = uri
	return &this
}

// NewOnDemandPageMetadataConnectionsMetadataConnectionsCommentsWithDefaults instantiates a new OnDemandPageMetadataConnectionsMetadataConnectionsComments object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewOnDemandPageMetadataConnectionsMetadataConnectionsCommentsWithDefaults() *OnDemandPageMetadataConnectionsMetadataConnectionsComments {
	this := OnDemandPageMetadataConnectionsMetadataConnectionsComments{}
	return &this
}

// GetOptions returns the Options field value
func (o *OnDemandPageMetadataConnectionsMetadataConnectionsComments) GetOptions() []string {
	if o == nil {
		var ret []string
		return ret
	}

	return o.Options
}

// GetOptionsOk returns a tuple with the Options field value
// and a boolean to check if the value has been set.
func (o *OnDemandPageMetadataConnectionsMetadataConnectionsComments) GetOptionsOk() ([]string, bool) {
	if o == nil {
		return nil, false
	}
	return o.Options, true
}

// SetOptions sets field value
func (o *OnDemandPageMetadataConnectionsMetadataConnectionsComments) SetOptions(v []string) {
	o.Options = v
}

// GetTotal returns the Total field value
func (o *OnDemandPageMetadataConnectionsMetadataConnectionsComments) GetTotal() float32 {
	if o == nil {
		var ret float32
		return ret
	}

	return o.Total
}

// GetTotalOk returns a tuple with the Total field value
// and a boolean to check if the value has been set.
func (o *OnDemandPageMetadataConnectionsMetadataConnectionsComments) GetTotalOk() (*float32, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Total, true
}

// SetTotal sets field value
func (o *OnDemandPageMetadataConnectionsMetadataConnectionsComments) SetTotal(v float32) {
	o.Total = v
}

// GetUri returns the Uri field value
func (o *OnDemandPageMetadataConnectionsMetadataConnectionsComments) GetUri() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Uri
}

// GetUriOk returns a tuple with the Uri field value
// and a boolean to check if the value has been set.
func (o *OnDemandPageMetadataConnectionsMetadataConnectionsComments) GetUriOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Uri, true
}

// SetUri sets field value
func (o *OnDemandPageMetadataConnectionsMetadataConnectionsComments) SetUri(v string) {
	o.Uri = v
}

func (o OnDemandPageMetadataConnectionsMetadataConnectionsComments) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o OnDemandPageMetadataConnectionsMetadataConnectionsComments) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["options"] = o.Options
	toSerialize["total"] = o.Total
	toSerialize["uri"] = o.Uri
	return toSerialize, nil
}

type NullableOnDemandPageMetadataConnectionsMetadataConnectionsComments struct {
	value *OnDemandPageMetadataConnectionsMetadataConnectionsComments
	isSet bool
}

func (v NullableOnDemandPageMetadataConnectionsMetadataConnectionsComments) Get() *OnDemandPageMetadataConnectionsMetadataConnectionsComments {
	return v.value
}

func (v *NullableOnDemandPageMetadataConnectionsMetadataConnectionsComments) Set(val *OnDemandPageMetadataConnectionsMetadataConnectionsComments) {
	v.value = val
	v.isSet = true
}

func (v NullableOnDemandPageMetadataConnectionsMetadataConnectionsComments) IsSet() bool {
	return v.isSet
}

func (v *NullableOnDemandPageMetadataConnectionsMetadataConnectionsComments) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableOnDemandPageMetadataConnectionsMetadataConnectionsComments(val *OnDemandPageMetadataConnectionsMetadataConnectionsComments) *NullableOnDemandPageMetadataConnectionsMetadataConnectionsComments {
	return &NullableOnDemandPageMetadataConnectionsMetadataConnectionsComments{value: val, isSet: true}
}

func (v NullableOnDemandPageMetadataConnectionsMetadataConnectionsComments) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableOnDemandPageMetadataConnectionsMetadataConnectionsComments) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
