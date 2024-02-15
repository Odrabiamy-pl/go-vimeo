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

// checks if the UserMetadataInteractionsReport type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &UserMetadataInteractionsReport{}

// UserMetadataInteractionsReport Information about where and how to report the requested user.
type UserMetadataInteractionsReport struct {
	// An array of the HTTP methods permitted on this URI.
	Options []string `json:"options"`
	// A list of valid reasons for reporting a video.
	Reason []string `json:"reason"`
	// The API URI that resolves to the connection data.
	Uri string `json:"uri"`
}

type _UserMetadataInteractionsReport UserMetadataInteractionsReport

// NewUserMetadataInteractionsReport instantiates a new UserMetadataInteractionsReport object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewUserMetadataInteractionsReport(options []string, reason []string, uri string) *UserMetadataInteractionsReport {
	this := UserMetadataInteractionsReport{}
	this.Options = options
	this.Reason = reason
	this.Uri = uri
	return &this
}

// NewUserMetadataInteractionsReportWithDefaults instantiates a new UserMetadataInteractionsReport object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewUserMetadataInteractionsReportWithDefaults() *UserMetadataInteractionsReport {
	this := UserMetadataInteractionsReport{}
	return &this
}

// GetOptions returns the Options field value
func (o *UserMetadataInteractionsReport) GetOptions() []string {
	if o == nil {
		var ret []string
		return ret
	}

	return o.Options
}

// GetOptionsOk returns a tuple with the Options field value
// and a boolean to check if the value has been set.
func (o *UserMetadataInteractionsReport) GetOptionsOk() ([]string, bool) {
	if o == nil {
		return nil, false
	}
	return o.Options, true
}

// SetOptions sets field value
func (o *UserMetadataInteractionsReport) SetOptions(v []string) {
	o.Options = v
}

// GetReason returns the Reason field value
func (o *UserMetadataInteractionsReport) GetReason() []string {
	if o == nil {
		var ret []string
		return ret
	}

	return o.Reason
}

// GetReasonOk returns a tuple with the Reason field value
// and a boolean to check if the value has been set.
func (o *UserMetadataInteractionsReport) GetReasonOk() ([]string, bool) {
	if o == nil {
		return nil, false
	}
	return o.Reason, true
}

// SetReason sets field value
func (o *UserMetadataInteractionsReport) SetReason(v []string) {
	o.Reason = v
}

// GetUri returns the Uri field value
func (o *UserMetadataInteractionsReport) GetUri() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Uri
}

// GetUriOk returns a tuple with the Uri field value
// and a boolean to check if the value has been set.
func (o *UserMetadataInteractionsReport) GetUriOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Uri, true
}

// SetUri sets field value
func (o *UserMetadataInteractionsReport) SetUri(v string) {
	o.Uri = v
}

func (o UserMetadataInteractionsReport) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o UserMetadataInteractionsReport) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["options"] = o.Options
	toSerialize["reason"] = o.Reason
	toSerialize["uri"] = o.Uri
	return toSerialize, nil
}

func (o *UserMetadataInteractionsReport) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"options",
		"reason",
		"uri",
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

	varUserMetadataInteractionsReport := _UserMetadataInteractionsReport{}

	decoder := json.NewDecoder(bytes.NewReader(data))
	decoder.DisallowUnknownFields()
	err = decoder.Decode(&varUserMetadataInteractionsReport)

	if err != nil {
		return err
	}

	*o = UserMetadataInteractionsReport(varUserMetadataInteractionsReport)

	return err
}

type NullableUserMetadataInteractionsReport struct {
	value *UserMetadataInteractionsReport
	isSet bool
}

func (v NullableUserMetadataInteractionsReport) Get() *UserMetadataInteractionsReport {
	return v.value
}

func (v *NullableUserMetadataInteractionsReport) Set(val *UserMetadataInteractionsReport) {
	v.value = val
	v.isSet = true
}

func (v NullableUserMetadataInteractionsReport) IsSet() bool {
	return v.isSet
}

func (v *NullableUserMetadataInteractionsReport) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableUserMetadataInteractionsReport(val *UserMetadataInteractionsReport) *NullableUserMetadataInteractionsReport {
	return &NullableUserMetadataInteractionsReport{value: val, isSet: true}
}

func (v NullableUserMetadataInteractionsReport) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableUserMetadataInteractionsReport) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


