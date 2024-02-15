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

// checks if the UserMetadata type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &UserMetadata{}

// UserMetadata The authenticated user's metadata.
type UserMetadata struct {
	Connections UserMetadataConnections `json:"connections"`
	Interactions UserMetadataInteractions `json:"interactions"`
	PublicVideos UserMetadataPublicVideos `json:"public_videos"`
}

type _UserMetadata UserMetadata

// NewUserMetadata instantiates a new UserMetadata object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewUserMetadata(connections UserMetadataConnections, interactions UserMetadataInteractions, publicVideos UserMetadataPublicVideos) *UserMetadata {
	this := UserMetadata{}
	this.Connections = connections
	this.Interactions = interactions
	this.PublicVideos = publicVideos
	return &this
}

// NewUserMetadataWithDefaults instantiates a new UserMetadata object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewUserMetadataWithDefaults() *UserMetadata {
	this := UserMetadata{}
	return &this
}

// GetConnections returns the Connections field value
func (o *UserMetadata) GetConnections() UserMetadataConnections {
	if o == nil {
		var ret UserMetadataConnections
		return ret
	}

	return o.Connections
}

// GetConnectionsOk returns a tuple with the Connections field value
// and a boolean to check if the value has been set.
func (o *UserMetadata) GetConnectionsOk() (*UserMetadataConnections, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Connections, true
}

// SetConnections sets field value
func (o *UserMetadata) SetConnections(v UserMetadataConnections) {
	o.Connections = v
}

// GetInteractions returns the Interactions field value
func (o *UserMetadata) GetInteractions() UserMetadataInteractions {
	if o == nil {
		var ret UserMetadataInteractions
		return ret
	}

	return o.Interactions
}

// GetInteractionsOk returns a tuple with the Interactions field value
// and a boolean to check if the value has been set.
func (o *UserMetadata) GetInteractionsOk() (*UserMetadataInteractions, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Interactions, true
}

// SetInteractions sets field value
func (o *UserMetadata) SetInteractions(v UserMetadataInteractions) {
	o.Interactions = v
}

// GetPublicVideos returns the PublicVideos field value
func (o *UserMetadata) GetPublicVideos() UserMetadataPublicVideos {
	if o == nil {
		var ret UserMetadataPublicVideos
		return ret
	}

	return o.PublicVideos
}

// GetPublicVideosOk returns a tuple with the PublicVideos field value
// and a boolean to check if the value has been set.
func (o *UserMetadata) GetPublicVideosOk() (*UserMetadataPublicVideos, bool) {
	if o == nil {
		return nil, false
	}
	return &o.PublicVideos, true
}

// SetPublicVideos sets field value
func (o *UserMetadata) SetPublicVideos(v UserMetadataPublicVideos) {
	o.PublicVideos = v
}

func (o UserMetadata) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o UserMetadata) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["connections"] = o.Connections
	toSerialize["interactions"] = o.Interactions
	toSerialize["public_videos"] = o.PublicVideos
	return toSerialize, nil
}

func (o *UserMetadata) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"connections",
		"interactions",
		"public_videos",
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

	varUserMetadata := _UserMetadata{}

	decoder := json.NewDecoder(bytes.NewReader(data))
	decoder.DisallowUnknownFields()
	err = decoder.Decode(&varUserMetadata)

	if err != nil {
		return err
	}

	*o = UserMetadata(varUserMetadata)

	return err
}

type NullableUserMetadata struct {
	value *UserMetadata
	isSet bool
}

func (v NullableUserMetadata) Get() *UserMetadata {
	return v.value
}

func (v *NullableUserMetadata) Set(val *UserMetadata) {
	v.value = val
	v.isSet = true
}

func (v NullableUserMetadata) IsSet() bool {
	return v.isSet
}

func (v *NullableUserMetadata) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableUserMetadata(val *UserMetadata) *NullableUserMetadata {
	return &NullableUserMetadata{value: val, isSet: true}
}

func (v NullableUserMetadata) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableUserMetadata) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


