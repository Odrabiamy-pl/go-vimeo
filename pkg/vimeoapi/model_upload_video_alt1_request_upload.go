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

// checks if the UploadVideoAlt1RequestUpload type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &UploadVideoAlt1RequestUpload{}

// UploadVideoAlt1RequestUpload struct for UploadVideoAlt1RequestUpload
type UploadVideoAlt1RequestUpload struct {
	// The upload approach.  Option descriptions:  * `post` - Use the `post` approach.  * `pull` - Use the `pull` approach.  * `tus` - Use the `tus` approach. 
	Approach string `json:"approach"`
	// The public URL at which the video is hosted. The URL must be valid for at least 24 hours. Use this parameter when **approach** is `pull`.
	Link *string `json:"link,omitempty"`
	// The app's redirect URL. Use this parameter when **approach** is `post`.
	RedirectUrl *string `json:"redirect_url,omitempty"`
	// The size in bytes of the video to upload. The maximum value of this field is `268435456000`, which corresponds to 250 GB.
	Size *string `json:"size,omitempty"`
}

// NewUploadVideoAlt1RequestUpload instantiates a new UploadVideoAlt1RequestUpload object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewUploadVideoAlt1RequestUpload(approach string) *UploadVideoAlt1RequestUpload {
	this := UploadVideoAlt1RequestUpload{}
	this.Approach = approach
	return &this
}

// NewUploadVideoAlt1RequestUploadWithDefaults instantiates a new UploadVideoAlt1RequestUpload object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewUploadVideoAlt1RequestUploadWithDefaults() *UploadVideoAlt1RequestUpload {
	this := UploadVideoAlt1RequestUpload{}
	return &this
}

// GetApproach returns the Approach field value
func (o *UploadVideoAlt1RequestUpload) GetApproach() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Approach
}

// GetApproachOk returns a tuple with the Approach field value
// and a boolean to check if the value has been set.
func (o *UploadVideoAlt1RequestUpload) GetApproachOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Approach, true
}

// SetApproach sets field value
func (o *UploadVideoAlt1RequestUpload) SetApproach(v string) {
	o.Approach = v
}

// GetLink returns the Link field value if set, zero value otherwise.
func (o *UploadVideoAlt1RequestUpload) GetLink() string {
	if o == nil || IsNil(o.Link) {
		var ret string
		return ret
	}
	return *o.Link
}

// GetLinkOk returns a tuple with the Link field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UploadVideoAlt1RequestUpload) GetLinkOk() (*string, bool) {
	if o == nil || IsNil(o.Link) {
		return nil, false
	}
	return o.Link, true
}

// HasLink returns a boolean if a field has been set.
func (o *UploadVideoAlt1RequestUpload) HasLink() bool {
	if o != nil && !IsNil(o.Link) {
		return true
	}

	return false
}

// SetLink gets a reference to the given string and assigns it to the Link field.
func (o *UploadVideoAlt1RequestUpload) SetLink(v string) {
	o.Link = &v
}

// GetRedirectUrl returns the RedirectUrl field value if set, zero value otherwise.
func (o *UploadVideoAlt1RequestUpload) GetRedirectUrl() string {
	if o == nil || IsNil(o.RedirectUrl) {
		var ret string
		return ret
	}
	return *o.RedirectUrl
}

// GetRedirectUrlOk returns a tuple with the RedirectUrl field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UploadVideoAlt1RequestUpload) GetRedirectUrlOk() (*string, bool) {
	if o == nil || IsNil(o.RedirectUrl) {
		return nil, false
	}
	return o.RedirectUrl, true
}

// HasRedirectUrl returns a boolean if a field has been set.
func (o *UploadVideoAlt1RequestUpload) HasRedirectUrl() bool {
	if o != nil && !IsNil(o.RedirectUrl) {
		return true
	}

	return false
}

// SetRedirectUrl gets a reference to the given string and assigns it to the RedirectUrl field.
func (o *UploadVideoAlt1RequestUpload) SetRedirectUrl(v string) {
	o.RedirectUrl = &v
}

// GetSize returns the Size field value if set, zero value otherwise.
func (o *UploadVideoAlt1RequestUpload) GetSize() string {
	if o == nil || IsNil(o.Size) {
		var ret string
		return ret
	}
	return *o.Size
}

// GetSizeOk returns a tuple with the Size field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UploadVideoAlt1RequestUpload) GetSizeOk() (*string, bool) {
	if o == nil || IsNil(o.Size) {
		return nil, false
	}
	return o.Size, true
}

// HasSize returns a boolean if a field has been set.
func (o *UploadVideoAlt1RequestUpload) HasSize() bool {
	if o != nil && !IsNil(o.Size) {
		return true
	}

	return false
}

// SetSize gets a reference to the given string and assigns it to the Size field.
func (o *UploadVideoAlt1RequestUpload) SetSize(v string) {
	o.Size = &v
}

func (o UploadVideoAlt1RequestUpload) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o UploadVideoAlt1RequestUpload) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["approach"] = o.Approach
	if !IsNil(o.Link) {
		toSerialize["link"] = o.Link
	}
	if !IsNil(o.RedirectUrl) {
		toSerialize["redirect_url"] = o.RedirectUrl
	}
	if !IsNil(o.Size) {
		toSerialize["size"] = o.Size
	}
	return toSerialize, nil
}

type NullableUploadVideoAlt1RequestUpload struct {
	value *UploadVideoAlt1RequestUpload
	isSet bool
}

func (v NullableUploadVideoAlt1RequestUpload) Get() *UploadVideoAlt1RequestUpload {
	return v.value
}

func (v *NullableUploadVideoAlt1RequestUpload) Set(val *UploadVideoAlt1RequestUpload) {
	v.value = val
	v.isSet = true
}

func (v NullableUploadVideoAlt1RequestUpload) IsSet() bool {
	return v.isSet
}

func (v *NullableUploadVideoAlt1RequestUpload) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableUploadVideoAlt1RequestUpload(val *UploadVideoAlt1RequestUpload) *NullableUploadVideoAlt1RequestUpload {
	return &NullableUploadVideoAlt1RequestUpload{value: val, isSet: true}
}

func (v NullableUploadVideoAlt1RequestUpload) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableUploadVideoAlt1RequestUpload) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


