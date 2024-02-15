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

// checks if the ProjectMetadataConnectionsVideos type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &ProjectMetadataConnectionsVideos{}

// ProjectMetadataConnectionsVideos A standard connection object indicating how to return all the videos in the folder.
type ProjectMetadataConnectionsVideos struct {
	// An array of HTTP methods permitted on this URI.
	Options []string `json:"options"`
	// The total number of videos on this connection.
	Total float32 `json:"total"`
	// The API URI that resolves to the connection data.
	Uri string `json:"uri"`
}

type _ProjectMetadataConnectionsVideos ProjectMetadataConnectionsVideos

// NewProjectMetadataConnectionsVideos instantiates a new ProjectMetadataConnectionsVideos object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewProjectMetadataConnectionsVideos(options []string, total float32, uri string) *ProjectMetadataConnectionsVideos {
	this := ProjectMetadataConnectionsVideos{}
	this.Options = options
	this.Total = total
	this.Uri = uri
	return &this
}

// NewProjectMetadataConnectionsVideosWithDefaults instantiates a new ProjectMetadataConnectionsVideos object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewProjectMetadataConnectionsVideosWithDefaults() *ProjectMetadataConnectionsVideos {
	this := ProjectMetadataConnectionsVideos{}
	return &this
}

// GetOptions returns the Options field value
func (o *ProjectMetadataConnectionsVideos) GetOptions() []string {
	if o == nil {
		var ret []string
		return ret
	}

	return o.Options
}

// GetOptionsOk returns a tuple with the Options field value
// and a boolean to check if the value has been set.
func (o *ProjectMetadataConnectionsVideos) GetOptionsOk() ([]string, bool) {
	if o == nil {
		return nil, false
	}
	return o.Options, true
}

// SetOptions sets field value
func (o *ProjectMetadataConnectionsVideos) SetOptions(v []string) {
	o.Options = v
}

// GetTotal returns the Total field value
func (o *ProjectMetadataConnectionsVideos) GetTotal() float32 {
	if o == nil {
		var ret float32
		return ret
	}

	return o.Total
}

// GetTotalOk returns a tuple with the Total field value
// and a boolean to check if the value has been set.
func (o *ProjectMetadataConnectionsVideos) GetTotalOk() (*float32, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Total, true
}

// SetTotal sets field value
func (o *ProjectMetadataConnectionsVideos) SetTotal(v float32) {
	o.Total = v
}

// GetUri returns the Uri field value
func (o *ProjectMetadataConnectionsVideos) GetUri() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Uri
}

// GetUriOk returns a tuple with the Uri field value
// and a boolean to check if the value has been set.
func (o *ProjectMetadataConnectionsVideos) GetUriOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Uri, true
}

// SetUri sets field value
func (o *ProjectMetadataConnectionsVideos) SetUri(v string) {
	o.Uri = v
}

func (o ProjectMetadataConnectionsVideos) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o ProjectMetadataConnectionsVideos) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["options"] = o.Options
	toSerialize["total"] = o.Total
	toSerialize["uri"] = o.Uri
	return toSerialize, nil
}

func (o *ProjectMetadataConnectionsVideos) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"options",
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

	varProjectMetadataConnectionsVideos := _ProjectMetadataConnectionsVideos{}

	decoder := json.NewDecoder(bytes.NewReader(data))
	decoder.DisallowUnknownFields()
	err = decoder.Decode(&varProjectMetadataConnectionsVideos)

	if err != nil {
		return err
	}

	*o = ProjectMetadataConnectionsVideos(varProjectMetadataConnectionsVideos)

	return err
}

type NullableProjectMetadataConnectionsVideos struct {
	value *ProjectMetadataConnectionsVideos
	isSet bool
}

func (v NullableProjectMetadataConnectionsVideos) Get() *ProjectMetadataConnectionsVideos {
	return v.value
}

func (v *NullableProjectMetadataConnectionsVideos) Set(val *ProjectMetadataConnectionsVideos) {
	v.value = val
	v.isSet = true
}

func (v NullableProjectMetadataConnectionsVideos) IsSet() bool {
	return v.isSet
}

func (v *NullableProjectMetadataConnectionsVideos) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableProjectMetadataConnectionsVideos(val *ProjectMetadataConnectionsVideos) *NullableProjectMetadataConnectionsVideos {
	return &NullableProjectMetadataConnectionsVideos{value: val, isSet: true}
}

func (v NullableProjectMetadataConnectionsVideos) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableProjectMetadataConnectionsVideos) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


