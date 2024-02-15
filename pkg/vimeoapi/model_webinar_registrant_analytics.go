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

// checks if the WebinarRegistrantAnalytics type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &WebinarRegistrantAnalytics{}

// WebinarRegistrantAnalytics The analytics data container for the webinar registrant.
type WebinarRegistrantAnalytics struct {
	// The percentage of the total webinar that the webinar registrant has attended.
	ViewPercentage float32 `json:"view_percentage"`
}

type _WebinarRegistrantAnalytics WebinarRegistrantAnalytics

// NewWebinarRegistrantAnalytics instantiates a new WebinarRegistrantAnalytics object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewWebinarRegistrantAnalytics(viewPercentage float32) *WebinarRegistrantAnalytics {
	this := WebinarRegistrantAnalytics{}
	this.ViewPercentage = viewPercentage
	return &this
}

// NewWebinarRegistrantAnalyticsWithDefaults instantiates a new WebinarRegistrantAnalytics object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewWebinarRegistrantAnalyticsWithDefaults() *WebinarRegistrantAnalytics {
	this := WebinarRegistrantAnalytics{}
	return &this
}

// GetViewPercentage returns the ViewPercentage field value
func (o *WebinarRegistrantAnalytics) GetViewPercentage() float32 {
	if o == nil {
		var ret float32
		return ret
	}

	return o.ViewPercentage
}

// GetViewPercentageOk returns a tuple with the ViewPercentage field value
// and a boolean to check if the value has been set.
func (o *WebinarRegistrantAnalytics) GetViewPercentageOk() (*float32, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ViewPercentage, true
}

// SetViewPercentage sets field value
func (o *WebinarRegistrantAnalytics) SetViewPercentage(v float32) {
	o.ViewPercentage = v
}

func (o WebinarRegistrantAnalytics) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o WebinarRegistrantAnalytics) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["view_percentage"] = o.ViewPercentage
	return toSerialize, nil
}

func (o *WebinarRegistrantAnalytics) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"view_percentage",
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

	varWebinarRegistrantAnalytics := _WebinarRegistrantAnalytics{}

	decoder := json.NewDecoder(bytes.NewReader(data))
	decoder.DisallowUnknownFields()
	err = decoder.Decode(&varWebinarRegistrantAnalytics)

	if err != nil {
		return err
	}

	*o = WebinarRegistrantAnalytics(varWebinarRegistrantAnalytics)

	return err
}

type NullableWebinarRegistrantAnalytics struct {
	value *WebinarRegistrantAnalytics
	isSet bool
}

func (v NullableWebinarRegistrantAnalytics) Get() *WebinarRegistrantAnalytics {
	return v.value
}

func (v *NullableWebinarRegistrantAnalytics) Set(val *WebinarRegistrantAnalytics) {
	v.value = val
	v.isSet = true
}

func (v NullableWebinarRegistrantAnalytics) IsSet() bool {
	return v.isSet
}

func (v *NullableWebinarRegistrantAnalytics) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableWebinarRegistrantAnalytics(val *WebinarRegistrantAnalytics) *NullableWebinarRegistrantAnalytics {
	return &NullableWebinarRegistrantAnalytics{value: val, isSet: true}
}

func (v NullableWebinarRegistrantAnalytics) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableWebinarRegistrantAnalytics) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


