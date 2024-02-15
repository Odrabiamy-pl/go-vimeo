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

// checks if the OnDemandPageThumbnail type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &OnDemandPageThumbnail{}

// OnDemandPageThumbnail The thumbnail image for the On Demand page.
type OnDemandPageThumbnail struct {
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
	// An array containing reference information about all available image files.
	Sizes []PictureSizesInner `json:"sizes"`
	// The type of picture.  Option descriptions:  * `caution` - The picture isn't appropriate for all ages.  * `custom` - The picture is a custom video image.  * `default` - The picture is the default video image. 
	Type string `json:"type"`
	// The URI of the picture.
	Uri string `json:"uri"`
}

type _OnDemandPageThumbnail OnDemandPageThumbnail

// NewOnDemandPageThumbnail instantiates a new OnDemandPageThumbnail object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewOnDemandPageThumbnail(active bool, baseLink string, defaultPicture bool, resourceKey string, sizes []PictureSizesInner, type_ string, uri string) *OnDemandPageThumbnail {
	this := OnDemandPageThumbnail{}
	this.Active = active
	this.BaseLink = baseLink
	this.DefaultPicture = defaultPicture
	this.ResourceKey = resourceKey
	this.Sizes = sizes
	this.Type = type_
	this.Uri = uri
	return &this
}

// NewOnDemandPageThumbnailWithDefaults instantiates a new OnDemandPageThumbnail object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewOnDemandPageThumbnailWithDefaults() *OnDemandPageThumbnail {
	this := OnDemandPageThumbnail{}
	return &this
}

// GetActive returns the Active field value
func (o *OnDemandPageThumbnail) GetActive() bool {
	if o == nil {
		var ret bool
		return ret
	}

	return o.Active
}

// GetActiveOk returns a tuple with the Active field value
// and a boolean to check if the value has been set.
func (o *OnDemandPageThumbnail) GetActiveOk() (*bool, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Active, true
}

// SetActive sets field value
func (o *OnDemandPageThumbnail) SetActive(v bool) {
	o.Active = v
}

// GetBaseLink returns the BaseLink field value
func (o *OnDemandPageThumbnail) GetBaseLink() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.BaseLink
}

// GetBaseLinkOk returns a tuple with the BaseLink field value
// and a boolean to check if the value has been set.
func (o *OnDemandPageThumbnail) GetBaseLinkOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.BaseLink, true
}

// SetBaseLink sets field value
func (o *OnDemandPageThumbnail) SetBaseLink(v string) {
	o.BaseLink = v
}

// GetDefaultPicture returns the DefaultPicture field value
func (o *OnDemandPageThumbnail) GetDefaultPicture() bool {
	if o == nil {
		var ret bool
		return ret
	}

	return o.DefaultPicture
}

// GetDefaultPictureOk returns a tuple with the DefaultPicture field value
// and a boolean to check if the value has been set.
func (o *OnDemandPageThumbnail) GetDefaultPictureOk() (*bool, bool) {
	if o == nil {
		return nil, false
	}
	return &o.DefaultPicture, true
}

// SetDefaultPicture sets field value
func (o *OnDemandPageThumbnail) SetDefaultPicture(v bool) {
	o.DefaultPicture = v
}

// GetLink returns the Link field value if set, zero value otherwise.
func (o *OnDemandPageThumbnail) GetLink() string {
	if o == nil || IsNil(o.Link) {
		var ret string
		return ret
	}
	return *o.Link
}

// GetLinkOk returns a tuple with the Link field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *OnDemandPageThumbnail) GetLinkOk() (*string, bool) {
	if o == nil || IsNil(o.Link) {
		return nil, false
	}
	return o.Link, true
}

// HasLink returns a boolean if a field has been set.
func (o *OnDemandPageThumbnail) HasLink() bool {
	if o != nil && !IsNil(o.Link) {
		return true
	}

	return false
}

// SetLink gets a reference to the given string and assigns it to the Link field.
func (o *OnDemandPageThumbnail) SetLink(v string) {
	o.Link = &v
}

// GetResourceKey returns the ResourceKey field value
func (o *OnDemandPageThumbnail) GetResourceKey() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.ResourceKey
}

// GetResourceKeyOk returns a tuple with the ResourceKey field value
// and a boolean to check if the value has been set.
func (o *OnDemandPageThumbnail) GetResourceKeyOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ResourceKey, true
}

// SetResourceKey sets field value
func (o *OnDemandPageThumbnail) SetResourceKey(v string) {
	o.ResourceKey = v
}

// GetSizes returns the Sizes field value
func (o *OnDemandPageThumbnail) GetSizes() []PictureSizesInner {
	if o == nil {
		var ret []PictureSizesInner
		return ret
	}

	return o.Sizes
}

// GetSizesOk returns a tuple with the Sizes field value
// and a boolean to check if the value has been set.
func (o *OnDemandPageThumbnail) GetSizesOk() ([]PictureSizesInner, bool) {
	if o == nil {
		return nil, false
	}
	return o.Sizes, true
}

// SetSizes sets field value
func (o *OnDemandPageThumbnail) SetSizes(v []PictureSizesInner) {
	o.Sizes = v
}

// GetType returns the Type field value
func (o *OnDemandPageThumbnail) GetType() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Type
}

// GetTypeOk returns a tuple with the Type field value
// and a boolean to check if the value has been set.
func (o *OnDemandPageThumbnail) GetTypeOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Type, true
}

// SetType sets field value
func (o *OnDemandPageThumbnail) SetType(v string) {
	o.Type = v
}

// GetUri returns the Uri field value
func (o *OnDemandPageThumbnail) GetUri() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Uri
}

// GetUriOk returns a tuple with the Uri field value
// and a boolean to check if the value has been set.
func (o *OnDemandPageThumbnail) GetUriOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Uri, true
}

// SetUri sets field value
func (o *OnDemandPageThumbnail) SetUri(v string) {
	o.Uri = v
}

func (o OnDemandPageThumbnail) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o OnDemandPageThumbnail) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["active"] = o.Active
	toSerialize["base_link"] = o.BaseLink
	toSerialize["default_picture"] = o.DefaultPicture
	if !IsNil(o.Link) {
		toSerialize["link"] = o.Link
	}
	toSerialize["resource_key"] = o.ResourceKey
	toSerialize["sizes"] = o.Sizes
	toSerialize["type"] = o.Type
	toSerialize["uri"] = o.Uri
	return toSerialize, nil
}

func (o *OnDemandPageThumbnail) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"active",
		"base_link",
		"default_picture",
		"resource_key",
		"sizes",
		"type",
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

	varOnDemandPageThumbnail := _OnDemandPageThumbnail{}

	decoder := json.NewDecoder(bytes.NewReader(data))
	decoder.DisallowUnknownFields()
	err = decoder.Decode(&varOnDemandPageThumbnail)

	if err != nil {
		return err
	}

	*o = OnDemandPageThumbnail(varOnDemandPageThumbnail)

	return err
}

type NullableOnDemandPageThumbnail struct {
	value *OnDemandPageThumbnail
	isSet bool
}

func (v NullableOnDemandPageThumbnail) Get() *OnDemandPageThumbnail {
	return v.value
}

func (v *NullableOnDemandPageThumbnail) Set(val *OnDemandPageThumbnail) {
	v.value = val
	v.isSet = true
}

func (v NullableOnDemandPageThumbnail) IsSet() bool {
	return v.isSet
}

func (v *NullableOnDemandPageThumbnail) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableOnDemandPageThumbnail(val *OnDemandPageThumbnail) *NullableOnDemandPageThumbnail {
	return &NullableOnDemandPageThumbnail{value: val, isSet: true}
}

func (v NullableOnDemandPageThumbnail) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableOnDemandPageThumbnail) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


