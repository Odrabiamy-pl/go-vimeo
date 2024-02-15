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

// checks if the AlbumMetadataConnections type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &AlbumMetadataConnections{}

// AlbumMetadataConnections Information about showcase connections.
type AlbumMetadataConnections struct {
	AvailableVideos AlbumMetadataConnectionsAvailableVideos `json:"available_videos"`
	RequestedClip NullableAlbumMetadataConnectionsRequestedClip `json:"requested_clip"`
	Videos AlbumMetadataConnectionsVideos `json:"videos"`
}

type _AlbumMetadataConnections AlbumMetadataConnections

// NewAlbumMetadataConnections instantiates a new AlbumMetadataConnections object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewAlbumMetadataConnections(availableVideos AlbumMetadataConnectionsAvailableVideos, requestedClip NullableAlbumMetadataConnectionsRequestedClip, videos AlbumMetadataConnectionsVideos) *AlbumMetadataConnections {
	this := AlbumMetadataConnections{}
	this.AvailableVideos = availableVideos
	this.RequestedClip = requestedClip
	this.Videos = videos
	return &this
}

// NewAlbumMetadataConnectionsWithDefaults instantiates a new AlbumMetadataConnections object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewAlbumMetadataConnectionsWithDefaults() *AlbumMetadataConnections {
	this := AlbumMetadataConnections{}
	return &this
}

// GetAvailableVideos returns the AvailableVideos field value
func (o *AlbumMetadataConnections) GetAvailableVideos() AlbumMetadataConnectionsAvailableVideos {
	if o == nil {
		var ret AlbumMetadataConnectionsAvailableVideos
		return ret
	}

	return o.AvailableVideos
}

// GetAvailableVideosOk returns a tuple with the AvailableVideos field value
// and a boolean to check if the value has been set.
func (o *AlbumMetadataConnections) GetAvailableVideosOk() (*AlbumMetadataConnectionsAvailableVideos, bool) {
	if o == nil {
		return nil, false
	}
	return &o.AvailableVideos, true
}

// SetAvailableVideos sets field value
func (o *AlbumMetadataConnections) SetAvailableVideos(v AlbumMetadataConnectionsAvailableVideos) {
	o.AvailableVideos = v
}

// GetRequestedClip returns the RequestedClip field value
// If the value is explicit nil, the zero value for AlbumMetadataConnectionsRequestedClip will be returned
func (o *AlbumMetadataConnections) GetRequestedClip() AlbumMetadataConnectionsRequestedClip {
	if o == nil || o.RequestedClip.Get() == nil {
		var ret AlbumMetadataConnectionsRequestedClip
		return ret
	}

	return *o.RequestedClip.Get()
}

// GetRequestedClipOk returns a tuple with the RequestedClip field value
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *AlbumMetadataConnections) GetRequestedClipOk() (*AlbumMetadataConnectionsRequestedClip, bool) {
	if o == nil {
		return nil, false
	}
	return o.RequestedClip.Get(), o.RequestedClip.IsSet()
}

// SetRequestedClip sets field value
func (o *AlbumMetadataConnections) SetRequestedClip(v AlbumMetadataConnectionsRequestedClip) {
	o.RequestedClip.Set(&v)
}

// GetVideos returns the Videos field value
func (o *AlbumMetadataConnections) GetVideos() AlbumMetadataConnectionsVideos {
	if o == nil {
		var ret AlbumMetadataConnectionsVideos
		return ret
	}

	return o.Videos
}

// GetVideosOk returns a tuple with the Videos field value
// and a boolean to check if the value has been set.
func (o *AlbumMetadataConnections) GetVideosOk() (*AlbumMetadataConnectionsVideos, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Videos, true
}

// SetVideos sets field value
func (o *AlbumMetadataConnections) SetVideos(v AlbumMetadataConnectionsVideos) {
	o.Videos = v
}

func (o AlbumMetadataConnections) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o AlbumMetadataConnections) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["available_videos"] = o.AvailableVideos
	toSerialize["requested_clip"] = o.RequestedClip.Get()
	toSerialize["videos"] = o.Videos
	return toSerialize, nil
}

func (o *AlbumMetadataConnections) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"available_videos",
		"requested_clip",
		"videos",
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

	varAlbumMetadataConnections := _AlbumMetadataConnections{}

	decoder := json.NewDecoder(bytes.NewReader(data))
	decoder.DisallowUnknownFields()
	err = decoder.Decode(&varAlbumMetadataConnections)

	if err != nil {
		return err
	}

	*o = AlbumMetadataConnections(varAlbumMetadataConnections)

	return err
}

type NullableAlbumMetadataConnections struct {
	value *AlbumMetadataConnections
	isSet bool
}

func (v NullableAlbumMetadataConnections) Get() *AlbumMetadataConnections {
	return v.value
}

func (v *NullableAlbumMetadataConnections) Set(val *AlbumMetadataConnections) {
	v.value = val
	v.isSet = true
}

func (v NullableAlbumMetadataConnections) IsSet() bool {
	return v.isSet
}

func (v *NullableAlbumMetadataConnections) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableAlbumMetadataConnections(val *AlbumMetadataConnections) *NullableAlbumMetadataConnections {
	return &NullableAlbumMetadataConnections{value: val, isSet: true}
}

func (v NullableAlbumMetadataConnections) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableAlbumMetadataConnections) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


