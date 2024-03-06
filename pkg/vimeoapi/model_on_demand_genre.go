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

// checks if the OnDemandGenre type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &OnDemandGenre{}

// OnDemandGenre struct for OnDemandGenre
type OnDemandGenre struct {
	// The canonical name or URL slug of the genre.
	Canonical    string                    `json:"canonical"`
	Interactions OnDemandGenreInteractions `json:"interactions"`
	// The Vimeo URL for the genre.
	Link     string                `json:"link"`
	Metadata OnDemandGenreMetadata `json:"metadata"`
	// The descriptive name of the genre.
	Name string `json:"name"`
	// The relative URI of the On Demand genre.
	Uri string `json:"uri"`
}

// NewOnDemandGenre instantiates a new OnDemandGenre object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewOnDemandGenre(canonical string, interactions OnDemandGenreInteractions, link string, metadata OnDemandGenreMetadata, name string, uri string) *OnDemandGenre {
	this := OnDemandGenre{}
	this.Canonical = canonical
	this.Interactions = interactions
	this.Link = link
	this.Metadata = metadata
	this.Name = name
	this.Uri = uri
	return &this
}

// NewOnDemandGenreWithDefaults instantiates a new OnDemandGenre object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewOnDemandGenreWithDefaults() *OnDemandGenre {
	this := OnDemandGenre{}
	return &this
}

// GetCanonical returns the Canonical field value
func (o *OnDemandGenre) GetCanonical() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Canonical
}

// GetCanonicalOk returns a tuple with the Canonical field value
// and a boolean to check if the value has been set.
func (o *OnDemandGenre) GetCanonicalOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Canonical, true
}

// SetCanonical sets field value
func (o *OnDemandGenre) SetCanonical(v string) {
	o.Canonical = v
}

// GetInteractions returns the Interactions field value
func (o *OnDemandGenre) GetInteractions() OnDemandGenreInteractions {
	if o == nil {
		var ret OnDemandGenreInteractions
		return ret
	}

	return o.Interactions
}

// GetInteractionsOk returns a tuple with the Interactions field value
// and a boolean to check if the value has been set.
func (o *OnDemandGenre) GetInteractionsOk() (*OnDemandGenreInteractions, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Interactions, true
}

// SetInteractions sets field value
func (o *OnDemandGenre) SetInteractions(v OnDemandGenreInteractions) {
	o.Interactions = v
}

// GetLink returns the Link field value
func (o *OnDemandGenre) GetLink() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Link
}

// GetLinkOk returns a tuple with the Link field value
// and a boolean to check if the value has been set.
func (o *OnDemandGenre) GetLinkOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Link, true
}

// SetLink sets field value
func (o *OnDemandGenre) SetLink(v string) {
	o.Link = v
}

// GetMetadata returns the Metadata field value
func (o *OnDemandGenre) GetMetadata() OnDemandGenreMetadata {
	if o == nil {
		var ret OnDemandGenreMetadata
		return ret
	}

	return o.Metadata
}

// GetMetadataOk returns a tuple with the Metadata field value
// and a boolean to check if the value has been set.
func (o *OnDemandGenre) GetMetadataOk() (*OnDemandGenreMetadata, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Metadata, true
}

// SetMetadata sets field value
func (o *OnDemandGenre) SetMetadata(v OnDemandGenreMetadata) {
	o.Metadata = v
}

// GetName returns the Name field value
func (o *OnDemandGenre) GetName() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Name
}

// GetNameOk returns a tuple with the Name field value
// and a boolean to check if the value has been set.
func (o *OnDemandGenre) GetNameOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Name, true
}

// SetName sets field value
func (o *OnDemandGenre) SetName(v string) {
	o.Name = v
}

// GetUri returns the Uri field value
func (o *OnDemandGenre) GetUri() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Uri
}

// GetUriOk returns a tuple with the Uri field value
// and a boolean to check if the value has been set.
func (o *OnDemandGenre) GetUriOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Uri, true
}

// SetUri sets field value
func (o *OnDemandGenre) SetUri(v string) {
	o.Uri = v
}

func (o OnDemandGenre) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o OnDemandGenre) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["canonical"] = o.Canonical
	toSerialize["interactions"] = o.Interactions
	toSerialize["link"] = o.Link
	toSerialize["metadata"] = o.Metadata
	toSerialize["name"] = o.Name
	toSerialize["uri"] = o.Uri
	return toSerialize, nil
}

type NullableOnDemandGenre struct {
	value *OnDemandGenre
	isSet bool
}

func (v NullableOnDemandGenre) Get() *OnDemandGenre {
	return v.value
}

func (v *NullableOnDemandGenre) Set(val *OnDemandGenre) {
	v.value = val
	v.isSet = true
}

func (v NullableOnDemandGenre) IsSet() bool {
	return v.isSet
}

func (v *NullableOnDemandGenre) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableOnDemandGenre(val *OnDemandGenre) *NullableOnDemandGenre {
	return &NullableOnDemandGenre{value: val, isSet: true}
}

func (v NullableOnDemandGenre) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableOnDemandGenre) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
