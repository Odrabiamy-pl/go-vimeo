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

// checks if the OnDemandSeasonMetadataConnections type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &OnDemandSeasonMetadataConnections{}

// OnDemandSeasonMetadataConnections struct for OnDemandSeasonMetadataConnections
type OnDemandSeasonMetadataConnections struct {
	Videos OnDemandSeasonMetadataConnectionsVideos `json:"videos"`
}

// NewOnDemandSeasonMetadataConnections instantiates a new OnDemandSeasonMetadataConnections object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewOnDemandSeasonMetadataConnections(videos OnDemandSeasonMetadataConnectionsVideos) *OnDemandSeasonMetadataConnections {
	this := OnDemandSeasonMetadataConnections{}
	this.Videos = videos
	return &this
}

// NewOnDemandSeasonMetadataConnectionsWithDefaults instantiates a new OnDemandSeasonMetadataConnections object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewOnDemandSeasonMetadataConnectionsWithDefaults() *OnDemandSeasonMetadataConnections {
	this := OnDemandSeasonMetadataConnections{}
	return &this
}

// GetVideos returns the Videos field value
func (o *OnDemandSeasonMetadataConnections) GetVideos() OnDemandSeasonMetadataConnectionsVideos {
	if o == nil {
		var ret OnDemandSeasonMetadataConnectionsVideos
		return ret
	}

	return o.Videos
}

// GetVideosOk returns a tuple with the Videos field value
// and a boolean to check if the value has been set.
func (o *OnDemandSeasonMetadataConnections) GetVideosOk() (*OnDemandSeasonMetadataConnectionsVideos, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Videos, true
}

// SetVideos sets field value
func (o *OnDemandSeasonMetadataConnections) SetVideos(v OnDemandSeasonMetadataConnectionsVideos) {
	o.Videos = v
}

func (o OnDemandSeasonMetadataConnections) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o OnDemandSeasonMetadataConnections) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["videos"] = o.Videos
	return toSerialize, nil
}

type NullableOnDemandSeasonMetadataConnections struct {
	value *OnDemandSeasonMetadataConnections
	isSet bool
}

func (v NullableOnDemandSeasonMetadataConnections) Get() *OnDemandSeasonMetadataConnections {
	return v.value
}

func (v *NullableOnDemandSeasonMetadataConnections) Set(val *OnDemandSeasonMetadataConnections) {
	v.value = val
	v.isSet = true
}

func (v NullableOnDemandSeasonMetadataConnections) IsSet() bool {
	return v.isSet
}

func (v *NullableOnDemandSeasonMetadataConnections) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableOnDemandSeasonMetadataConnections(val *OnDemandSeasonMetadataConnections) *NullableOnDemandSeasonMetadataConnections {
	return &NullableOnDemandSeasonMetadataConnections{value: val, isSet: true}
}

func (v NullableOnDemandSeasonMetadataConnections) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableOnDemandSeasonMetadataConnections) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


