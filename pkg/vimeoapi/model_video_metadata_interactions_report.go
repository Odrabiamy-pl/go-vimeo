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

// checks if the VideoMetadataInteractionsReport type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &VideoMetadataInteractionsReport{}

// VideoMetadataInteractionsReport Information about where and how to report a video.
type VideoMetadataInteractionsReport struct {
	// An array of HTTP methods permitted on this URI.
	Options []string `json:"options"`
	// A list of valid reasons for reporting a video.
	Reason []string `json:"reason"`
	// The API URI that resolves to the connection data.
	Uri string `json:"uri"`
}

// NewVideoMetadataInteractionsReport instantiates a new VideoMetadataInteractionsReport object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewVideoMetadataInteractionsReport(options []string, reason []string, uri string) *VideoMetadataInteractionsReport {
	this := VideoMetadataInteractionsReport{}
	this.Options = options
	this.Reason = reason
	this.Uri = uri
	return &this
}

// NewVideoMetadataInteractionsReportWithDefaults instantiates a new VideoMetadataInteractionsReport object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewVideoMetadataInteractionsReportWithDefaults() *VideoMetadataInteractionsReport {
	this := VideoMetadataInteractionsReport{}
	return &this
}

// GetOptions returns the Options field value
func (o *VideoMetadataInteractionsReport) GetOptions() []string {
	if o == nil {
		var ret []string
		return ret
	}

	return o.Options
}

// GetOptionsOk returns a tuple with the Options field value
// and a boolean to check if the value has been set.
func (o *VideoMetadataInteractionsReport) GetOptionsOk() ([]string, bool) {
	if o == nil {
		return nil, false
	}
	return o.Options, true
}

// SetOptions sets field value
func (o *VideoMetadataInteractionsReport) SetOptions(v []string) {
	o.Options = v
}

// GetReason returns the Reason field value
func (o *VideoMetadataInteractionsReport) GetReason() []string {
	if o == nil {
		var ret []string
		return ret
	}

	return o.Reason
}

// GetReasonOk returns a tuple with the Reason field value
// and a boolean to check if the value has been set.
func (o *VideoMetadataInteractionsReport) GetReasonOk() ([]string, bool) {
	if o == nil {
		return nil, false
	}
	return o.Reason, true
}

// SetReason sets field value
func (o *VideoMetadataInteractionsReport) SetReason(v []string) {
	o.Reason = v
}

// GetUri returns the Uri field value
func (o *VideoMetadataInteractionsReport) GetUri() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Uri
}

// GetUriOk returns a tuple with the Uri field value
// and a boolean to check if the value has been set.
func (o *VideoMetadataInteractionsReport) GetUriOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Uri, true
}

// SetUri sets field value
func (o *VideoMetadataInteractionsReport) SetUri(v string) {
	o.Uri = v
}

func (o VideoMetadataInteractionsReport) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o VideoMetadataInteractionsReport) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["options"] = o.Options
	toSerialize["reason"] = o.Reason
	toSerialize["uri"] = o.Uri
	return toSerialize, nil
}

type NullableVideoMetadataInteractionsReport struct {
	value *VideoMetadataInteractionsReport
	isSet bool
}

func (v NullableVideoMetadataInteractionsReport) Get() *VideoMetadataInteractionsReport {
	return v.value
}

func (v *NullableVideoMetadataInteractionsReport) Set(val *VideoMetadataInteractionsReport) {
	v.value = val
	v.isSet = true
}

func (v NullableVideoMetadataInteractionsReport) IsSet() bool {
	return v.isSet
}

func (v *NullableVideoMetadataInteractionsReport) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableVideoMetadataInteractionsReport(val *VideoMetadataInteractionsReport) *NullableVideoMetadataInteractionsReport {
	return &NullableVideoMetadataInteractionsReport{value: val, isSet: true}
}

func (v NullableVideoMetadataInteractionsReport) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableVideoMetadataInteractionsReport) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
