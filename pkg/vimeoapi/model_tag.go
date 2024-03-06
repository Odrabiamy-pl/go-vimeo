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

// checks if the Tag type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &Tag{}

// Tag struct for Tag
type Tag struct {
	// The normalized canonical tag name.
	Canonical string      `json:"canonical"`
	Metadata  TagMetadata `json:"metadata"`
	// The tag value.
	Name string `json:"name"`
	// The tag's resource key string.
	ResourceKey string `json:"resource_key"`
	// The canonical relative URI of the tag.
	Uri string `json:"uri"`
}

// NewTag instantiates a new Tag object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewTag(canonical string, metadata TagMetadata, name string, resourceKey string, uri string) *Tag {
	this := Tag{}
	this.Canonical = canonical
	this.Metadata = metadata
	this.Name = name
	this.ResourceKey = resourceKey
	this.Uri = uri
	return &this
}

// NewTagWithDefaults instantiates a new Tag object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewTagWithDefaults() *Tag {
	this := Tag{}
	return &this
}

// GetCanonical returns the Canonical field value
func (o *Tag) GetCanonical() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Canonical
}

// GetCanonicalOk returns a tuple with the Canonical field value
// and a boolean to check if the value has been set.
func (o *Tag) GetCanonicalOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Canonical, true
}

// SetCanonical sets field value
func (o *Tag) SetCanonical(v string) {
	o.Canonical = v
}

// GetMetadata returns the Metadata field value
func (o *Tag) GetMetadata() TagMetadata {
	if o == nil {
		var ret TagMetadata
		return ret
	}

	return o.Metadata
}

// GetMetadataOk returns a tuple with the Metadata field value
// and a boolean to check if the value has been set.
func (o *Tag) GetMetadataOk() (*TagMetadata, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Metadata, true
}

// SetMetadata sets field value
func (o *Tag) SetMetadata(v TagMetadata) {
	o.Metadata = v
}

// GetName returns the Name field value
func (o *Tag) GetName() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Name
}

// GetNameOk returns a tuple with the Name field value
// and a boolean to check if the value has been set.
func (o *Tag) GetNameOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Name, true
}

// SetName sets field value
func (o *Tag) SetName(v string) {
	o.Name = v
}

// GetResourceKey returns the ResourceKey field value
func (o *Tag) GetResourceKey() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.ResourceKey
}

// GetResourceKeyOk returns a tuple with the ResourceKey field value
// and a boolean to check if the value has been set.
func (o *Tag) GetResourceKeyOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ResourceKey, true
}

// SetResourceKey sets field value
func (o *Tag) SetResourceKey(v string) {
	o.ResourceKey = v
}

// GetUri returns the Uri field value
func (o *Tag) GetUri() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Uri
}

// GetUriOk returns a tuple with the Uri field value
// and a boolean to check if the value has been set.
func (o *Tag) GetUriOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Uri, true
}

// SetUri sets field value
func (o *Tag) SetUri(v string) {
	o.Uri = v
}

func (o Tag) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o Tag) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["canonical"] = o.Canonical
	toSerialize["metadata"] = o.Metadata
	toSerialize["name"] = o.Name
	toSerialize["resource_key"] = o.ResourceKey
	toSerialize["uri"] = o.Uri
	return toSerialize, nil
}

type NullableTag struct {
	value *Tag
	isSet bool
}

func (v NullableTag) Get() *Tag {
	return v.value
}

func (v *NullableTag) Set(val *Tag) {
	v.value = val
	v.isSet = true
}

func (v NullableTag) IsSet() bool {
	return v.isSet
}

func (v *NullableTag) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableTag(val *Tag) *NullableTag {
	return &NullableTag{value: val, isSet: true}
}

func (v NullableTag) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableTag) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
