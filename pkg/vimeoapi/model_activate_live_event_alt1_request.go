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

// checks if the ActivateLiveEventAlt1Request type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &ActivateLiveEventAlt1Request{}

// ActivateLiveEventAlt1Request struct for ActivateLiveEventAlt1Request
type ActivateLiveEventAlt1Request struct {
	// Whether the stream activates from the cloud composer. _This field is deprecated._
	CloudComposingStreaming *bool `json:"cloud_composing_streaming,omitempty"`
	// Whether the stream activates from the cloud composer.
	StreamingStartRequested *bool `json:"streaming_start_requested,omitempty"`
}

// NewActivateLiveEventAlt1Request instantiates a new ActivateLiveEventAlt1Request object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewActivateLiveEventAlt1Request() *ActivateLiveEventAlt1Request {
	this := ActivateLiveEventAlt1Request{}
	return &this
}

// NewActivateLiveEventAlt1RequestWithDefaults instantiates a new ActivateLiveEventAlt1Request object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewActivateLiveEventAlt1RequestWithDefaults() *ActivateLiveEventAlt1Request {
	this := ActivateLiveEventAlt1Request{}
	return &this
}

// GetCloudComposingStreaming returns the CloudComposingStreaming field value if set, zero value otherwise.
func (o *ActivateLiveEventAlt1Request) GetCloudComposingStreaming() bool {
	if o == nil || IsNil(o.CloudComposingStreaming) {
		var ret bool
		return ret
	}
	return *o.CloudComposingStreaming
}

// GetCloudComposingStreamingOk returns a tuple with the CloudComposingStreaming field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ActivateLiveEventAlt1Request) GetCloudComposingStreamingOk() (*bool, bool) {
	if o == nil || IsNil(o.CloudComposingStreaming) {
		return nil, false
	}
	return o.CloudComposingStreaming, true
}

// HasCloudComposingStreaming returns a boolean if a field has been set.
func (o *ActivateLiveEventAlt1Request) HasCloudComposingStreaming() bool {
	if o != nil && !IsNil(o.CloudComposingStreaming) {
		return true
	}

	return false
}

// SetCloudComposingStreaming gets a reference to the given bool and assigns it to the CloudComposingStreaming field.
func (o *ActivateLiveEventAlt1Request) SetCloudComposingStreaming(v bool) {
	o.CloudComposingStreaming = &v
}

// GetStreamingStartRequested returns the StreamingStartRequested field value if set, zero value otherwise.
func (o *ActivateLiveEventAlt1Request) GetStreamingStartRequested() bool {
	if o == nil || IsNil(o.StreamingStartRequested) {
		var ret bool
		return ret
	}
	return *o.StreamingStartRequested
}

// GetStreamingStartRequestedOk returns a tuple with the StreamingStartRequested field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ActivateLiveEventAlt1Request) GetStreamingStartRequestedOk() (*bool, bool) {
	if o == nil || IsNil(o.StreamingStartRequested) {
		return nil, false
	}
	return o.StreamingStartRequested, true
}

// HasStreamingStartRequested returns a boolean if a field has been set.
func (o *ActivateLiveEventAlt1Request) HasStreamingStartRequested() bool {
	if o != nil && !IsNil(o.StreamingStartRequested) {
		return true
	}

	return false
}

// SetStreamingStartRequested gets a reference to the given bool and assigns it to the StreamingStartRequested field.
func (o *ActivateLiveEventAlt1Request) SetStreamingStartRequested(v bool) {
	o.StreamingStartRequested = &v
}

func (o ActivateLiveEventAlt1Request) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o ActivateLiveEventAlt1Request) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.CloudComposingStreaming) {
		toSerialize["cloud_composing_streaming"] = o.CloudComposingStreaming
	}
	if !IsNil(o.StreamingStartRequested) {
		toSerialize["streaming_start_requested"] = o.StreamingStartRequested
	}
	return toSerialize, nil
}

type NullableActivateLiveEventAlt1Request struct {
	value *ActivateLiveEventAlt1Request
	isSet bool
}

func (v NullableActivateLiveEventAlt1Request) Get() *ActivateLiveEventAlt1Request {
	return v.value
}

func (v *NullableActivateLiveEventAlt1Request) Set(val *ActivateLiveEventAlt1Request) {
	v.value = val
	v.isSet = true
}

func (v NullableActivateLiveEventAlt1Request) IsSet() bool {
	return v.isSet
}

func (v *NullableActivateLiveEventAlt1Request) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableActivateLiveEventAlt1Request(val *ActivateLiveEventAlt1Request) *NullableActivateLiveEventAlt1Request {
	return &NullableActivateLiveEventAlt1Request{value: val, isSet: true}
}

func (v NullableActivateLiveEventAlt1Request) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableActivateLiveEventAlt1Request) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


