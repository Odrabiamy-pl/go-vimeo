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

// checks if the VideoMetadataConnectionsVersions type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &VideoMetadataConnectionsVersions{}

// VideoMetadataConnectionsVersions Information about the video's versions.
type VideoMetadataConnectionsVersions struct {
	// The URI of the current version of the video.
	CurrentUri *string `json:"current_uri,omitempty"`
	// Whether the video has interactive capability.
	HasInteractive bool `json:"has_interactive"`
	// An array of HTTP methods permitted on this URI.
	Options []string `json:"options"`
	// Whether the video has unified resolution. If the value of this field is `false`, the video requires transcoding.
	OriginVariableFrameResolution bool `json:"origin_variable_frame_resolution"`
	// The resource key string of the current version of the video.
	ResourceKey *string `json:"resource_key,omitempty"`
	// The total number of versions on this connection.
	Total float32 `json:"total"`
	// The API URI that resolves to the connection data.
	Uri string `json:"uri"`
}

type _VideoMetadataConnectionsVersions VideoMetadataConnectionsVersions

// NewVideoMetadataConnectionsVersions instantiates a new VideoMetadataConnectionsVersions object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewVideoMetadataConnectionsVersions(hasInteractive bool, options []string, originVariableFrameResolution bool, total float32, uri string) *VideoMetadataConnectionsVersions {
	this := VideoMetadataConnectionsVersions{}
	this.HasInteractive = hasInteractive
	this.Options = options
	this.OriginVariableFrameResolution = originVariableFrameResolution
	this.Total = total
	this.Uri = uri
	return &this
}

// NewVideoMetadataConnectionsVersionsWithDefaults instantiates a new VideoMetadataConnectionsVersions object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewVideoMetadataConnectionsVersionsWithDefaults() *VideoMetadataConnectionsVersions {
	this := VideoMetadataConnectionsVersions{}
	return &this
}

// GetCurrentUri returns the CurrentUri field value if set, zero value otherwise.
func (o *VideoMetadataConnectionsVersions) GetCurrentUri() string {
	if o == nil || IsNil(o.CurrentUri) {
		var ret string
		return ret
	}
	return *o.CurrentUri
}

// GetCurrentUriOk returns a tuple with the CurrentUri field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *VideoMetadataConnectionsVersions) GetCurrentUriOk() (*string, bool) {
	if o == nil || IsNil(o.CurrentUri) {
		return nil, false
	}
	return o.CurrentUri, true
}

// HasCurrentUri returns a boolean if a field has been set.
func (o *VideoMetadataConnectionsVersions) HasCurrentUri() bool {
	if o != nil && !IsNil(o.CurrentUri) {
		return true
	}

	return false
}

// SetCurrentUri gets a reference to the given string and assigns it to the CurrentUri field.
func (o *VideoMetadataConnectionsVersions) SetCurrentUri(v string) {
	o.CurrentUri = &v
}

// GetHasInteractive returns the HasInteractive field value
func (o *VideoMetadataConnectionsVersions) GetHasInteractive() bool {
	if o == nil {
		var ret bool
		return ret
	}

	return o.HasInteractive
}

// GetHasInteractiveOk returns a tuple with the HasInteractive field value
// and a boolean to check if the value has been set.
func (o *VideoMetadataConnectionsVersions) GetHasInteractiveOk() (*bool, bool) {
	if o == nil {
		return nil, false
	}
	return &o.HasInteractive, true
}

// SetHasInteractive sets field value
func (o *VideoMetadataConnectionsVersions) SetHasInteractive(v bool) {
	o.HasInteractive = v
}

// GetOptions returns the Options field value
func (o *VideoMetadataConnectionsVersions) GetOptions() []string {
	if o == nil {
		var ret []string
		return ret
	}

	return o.Options
}

// GetOptionsOk returns a tuple with the Options field value
// and a boolean to check if the value has been set.
func (o *VideoMetadataConnectionsVersions) GetOptionsOk() ([]string, bool) {
	if o == nil {
		return nil, false
	}
	return o.Options, true
}

// SetOptions sets field value
func (o *VideoMetadataConnectionsVersions) SetOptions(v []string) {
	o.Options = v
}

// GetOriginVariableFrameResolution returns the OriginVariableFrameResolution field value
func (o *VideoMetadataConnectionsVersions) GetOriginVariableFrameResolution() bool {
	if o == nil {
		var ret bool
		return ret
	}

	return o.OriginVariableFrameResolution
}

// GetOriginVariableFrameResolutionOk returns a tuple with the OriginVariableFrameResolution field value
// and a boolean to check if the value has been set.
func (o *VideoMetadataConnectionsVersions) GetOriginVariableFrameResolutionOk() (*bool, bool) {
	if o == nil {
		return nil, false
	}
	return &o.OriginVariableFrameResolution, true
}

// SetOriginVariableFrameResolution sets field value
func (o *VideoMetadataConnectionsVersions) SetOriginVariableFrameResolution(v bool) {
	o.OriginVariableFrameResolution = v
}

// GetResourceKey returns the ResourceKey field value if set, zero value otherwise.
func (o *VideoMetadataConnectionsVersions) GetResourceKey() string {
	if o == nil || IsNil(o.ResourceKey) {
		var ret string
		return ret
	}
	return *o.ResourceKey
}

// GetResourceKeyOk returns a tuple with the ResourceKey field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *VideoMetadataConnectionsVersions) GetResourceKeyOk() (*string, bool) {
	if o == nil || IsNil(o.ResourceKey) {
		return nil, false
	}
	return o.ResourceKey, true
}

// HasResourceKey returns a boolean if a field has been set.
func (o *VideoMetadataConnectionsVersions) HasResourceKey() bool {
	if o != nil && !IsNil(o.ResourceKey) {
		return true
	}

	return false
}

// SetResourceKey gets a reference to the given string and assigns it to the ResourceKey field.
func (o *VideoMetadataConnectionsVersions) SetResourceKey(v string) {
	o.ResourceKey = &v
}

// GetTotal returns the Total field value
func (o *VideoMetadataConnectionsVersions) GetTotal() float32 {
	if o == nil {
		var ret float32
		return ret
	}

	return o.Total
}

// GetTotalOk returns a tuple with the Total field value
// and a boolean to check if the value has been set.
func (o *VideoMetadataConnectionsVersions) GetTotalOk() (*float32, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Total, true
}

// SetTotal sets field value
func (o *VideoMetadataConnectionsVersions) SetTotal(v float32) {
	o.Total = v
}

// GetUri returns the Uri field value
func (o *VideoMetadataConnectionsVersions) GetUri() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Uri
}

// GetUriOk returns a tuple with the Uri field value
// and a boolean to check if the value has been set.
func (o *VideoMetadataConnectionsVersions) GetUriOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Uri, true
}

// SetUri sets field value
func (o *VideoMetadataConnectionsVersions) SetUri(v string) {
	o.Uri = v
}

func (o VideoMetadataConnectionsVersions) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o VideoMetadataConnectionsVersions) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.CurrentUri) {
		toSerialize["current_uri"] = o.CurrentUri
	}
	toSerialize["has_interactive"] = o.HasInteractive
	toSerialize["options"] = o.Options
	toSerialize["origin_variable_frame_resolution"] = o.OriginVariableFrameResolution
	if !IsNil(o.ResourceKey) {
		toSerialize["resource_key"] = o.ResourceKey
	}
	toSerialize["total"] = o.Total
	toSerialize["uri"] = o.Uri
	return toSerialize, nil
}

func (o *VideoMetadataConnectionsVersions) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"has_interactive",
		"options",
		"origin_variable_frame_resolution",
		"total",
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

	varVideoMetadataConnectionsVersions := _VideoMetadataConnectionsVersions{}

	decoder := json.NewDecoder(bytes.NewReader(data))
	decoder.DisallowUnknownFields()
	err = decoder.Decode(&varVideoMetadataConnectionsVersions)

	if err != nil {
		return err
	}

	*o = VideoMetadataConnectionsVersions(varVideoMetadataConnectionsVersions)

	return err
}

type NullableVideoMetadataConnectionsVersions struct {
	value *VideoMetadataConnectionsVersions
	isSet bool
}

func (v NullableVideoMetadataConnectionsVersions) Get() *VideoMetadataConnectionsVersions {
	return v.value
}

func (v *NullableVideoMetadataConnectionsVersions) Set(val *VideoMetadataConnectionsVersions) {
	v.value = val
	v.isSet = true
}

func (v NullableVideoMetadataConnectionsVersions) IsSet() bool {
	return v.isSet
}

func (v *NullableVideoMetadataConnectionsVersions) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableVideoMetadataConnectionsVersions(val *VideoMetadataConnectionsVersions) *NullableVideoMetadataConnectionsVersions {
	return &NullableVideoMetadataConnectionsVersions{value: val, isSet: true}
}

func (v NullableVideoMetadataConnectionsVersions) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableVideoMetadataConnectionsVersions) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


