/*
Vimeo API

Build something great. Vimeo's API supports flexible, high-quality video integration with your custom apps.

API version: 3.4
*/

package vimeoapi

import (
	"encoding/json"
)

// checks if the Picture type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &Picture{}

// Picture struct for Picture
type Picture struct {
	// Whether the picture is currently active.
	Active bool `json:"active"`
	// The base link to the image file, without any parameters.
	BaseLink string `json:"base_link"`
	// Whether the picture is Vimeo's default.
	DefaultPicture bool `json:"default_picture"`
	// The upload URL of the picture. This field appears upon the initial creation of a picture resource.
	Link *string `json:"link,omitempty"`
	// The resource key string of the picture.
	ResourceKey string `json:"resource_key"`
	// The type of picture.  Option descriptions:  * `caution` - The picture isn't appropriate for all ages.  * `custom` - The picture is a custom video image.  * `default` - The picture is the default video image.
	Type string `json:"type"`
	// The URI of the picture.
	Uri string `json:"uri"`
}

// NewPicture instantiates a new Picture object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewPicture(active bool, baseLink string, defaultPicture bool, resourceKey string, type_ string, uri string) *Picture {
	this := Picture{}
	this.Active = active
	this.BaseLink = baseLink
	this.DefaultPicture = defaultPicture
	this.ResourceKey = resourceKey
	this.Type = type_
	this.Uri = uri
	return &this
}

// NewPictureWithDefaults instantiates a new Picture object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewPictureWithDefaults() *Picture {
	this := Picture{}
	return &this
}

// GetActive returns the Active field value
func (o *Picture) GetActive() bool {
	if o == nil {
		var ret bool
		return ret
	}

	return o.Active
}

// GetActiveOk returns a tuple with the Active field value
// and a boolean to check if the value has been set.
func (o *Picture) GetActiveOk() (*bool, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Active, true
}

// SetActive sets field value
func (o *Picture) SetActive(v bool) {
	o.Active = v
}

// GetBaseLink returns the BaseLink field value
func (o *Picture) GetBaseLink() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.BaseLink
}

// GetBaseLinkOk returns a tuple with the BaseLink field value
// and a boolean to check if the value has been set.
func (o *Picture) GetBaseLinkOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.BaseLink, true
}

// SetBaseLink sets field value
func (o *Picture) SetBaseLink(v string) {
	o.BaseLink = v
}

// GetDefaultPicture returns the DefaultPicture field value
func (o *Picture) GetDefaultPicture() bool {
	if o == nil {
		var ret bool
		return ret
	}

	return o.DefaultPicture
}

// GetDefaultPictureOk returns a tuple with the DefaultPicture field value
// and a boolean to check if the value has been set.
func (o *Picture) GetDefaultPictureOk() (*bool, bool) {
	if o == nil {
		return nil, false
	}
	return &o.DefaultPicture, true
}

// SetDefaultPicture sets field value
func (o *Picture) SetDefaultPicture(v bool) {
	o.DefaultPicture = v
}

// GetLink returns the Link field value if set, zero value otherwise.
func (o *Picture) GetLink() string {
	if o == nil || IsNil(o.Link) {
		var ret string
		return ret
	}
	return *o.Link
}

// GetLinkOk returns a tuple with the Link field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Picture) GetLinkOk() (*string, bool) {
	if o == nil || IsNil(o.Link) {
		return nil, false
	}
	return o.Link, true
}

// HasLink returns a boolean if a field has been set.
func (o *Picture) HasLink() bool {
	if o != nil && !IsNil(o.Link) {
		return true
	}

	return false
}

// SetLink gets a reference to the given string and assigns it to the Link field.
func (o *Picture) SetLink(v string) {
	o.Link = &v
}

// GetResourceKey returns the ResourceKey field value
func (o *Picture) GetResourceKey() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.ResourceKey
}

// GetResourceKeyOk returns a tuple with the ResourceKey field value
// and a boolean to check if the value has been set.
func (o *Picture) GetResourceKeyOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ResourceKey, true
}

// SetResourceKey sets field value
func (o *Picture) SetResourceKey(v string) {
	o.ResourceKey = v
}

// GetType returns the Type field value
func (o *Picture) GetType() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Type
}

// GetTypeOk returns a tuple with the Type field value
// and a boolean to check if the value has been set.
func (o *Picture) GetTypeOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Type, true
}

// SetType sets field value
func (o *Picture) SetType(v string) {
	o.Type = v
}

// GetUri returns the Uri field value
func (o *Picture) GetUri() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Uri
}

// GetUriOk returns a tuple with the Uri field value
// and a boolean to check if the value has been set.
func (o *Picture) GetUriOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Uri, true
}

// SetUri sets field value
func (o *Picture) SetUri(v string) {
	o.Uri = v
}

func (o Picture) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o Picture) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["active"] = o.Active
	toSerialize["base_link"] = o.BaseLink
	toSerialize["default_picture"] = o.DefaultPicture
	if !IsNil(o.Link) {
		toSerialize["link"] = o.Link
	}
	toSerialize["resource_key"] = o.ResourceKey
	toSerialize["type"] = o.Type
	toSerialize["uri"] = o.Uri
	return toSerialize, nil
}

type NullablePicture struct {
	value *Picture
	isSet bool
}

func (v NullablePicture) Get() *Picture {
	return v.value
}

func (v *NullablePicture) Set(val *Picture) {
	v.value = val
	v.isSet = true
}

func (v NullablePicture) IsSet() bool {
	return v.isSet
}

func (v *NullablePicture) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullablePicture(val *Picture) *NullablePicture {
	return &NullablePicture{value: val, isSet: true}
}

func (v NullablePicture) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullablePicture) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
