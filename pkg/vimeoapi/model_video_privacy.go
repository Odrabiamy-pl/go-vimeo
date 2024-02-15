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

// checks if the VideoPrivacy type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &VideoPrivacy{}

// VideoPrivacy The video's privacy setting.
type VideoPrivacy struct {
	// Whether the video can be added to collections.
	Add bool `json:"add"`
	// Whether the share link is usable.
	AllowShareLink bool `json:"allow_share_link"`
	// The video's comment permission setting.  Option descriptions:  * `anybody` - Anyone can comment on the video.  * `contacts` - Only contacts can comment on the video.  * `nobody` - No one can comment on the video. 
	Comments string `json:"comments"`
	// Whether the video can be downloaded.
	Download bool `json:"download"`
	// The video's embed permission setting.  Option descriptions:  * `private` - The video is private.  * `public` - Anyone can embed the video.  * `whitelist` - The video can be embedded on specific domains. 
	Embed string `json:"embed"`
	// The general privacy setting of the video.  Option descriptions:  * `anybody` - Anyone can access the video. This privacy setting appears as `Public` on the Vimeo front end.  * `contacts` - Only contacts can access the video. _This field is deprecated._  * `disable` - The video is hidden from Vimeo. This privacy setting appears as `Hide from Vimeo` on the Vimeo front end.  * `nobody` - No one besides the owner can access the video. This privacy setting appears as `Private` on the Vimeo front end.  * `password` - Anyone with the video's password can access the video.  * `ptv` - The Vimeo On Demand video is accessible and searchable from Vimeo. _This field is deprecated._  * `ptvhide` - The Vimeo On Demand video is hidden from Vimeo. _This field is deprecated._  * `stock` - The stock footage is accessible and searchable from Vimeo. _This field is deprecated._  * `stock_purchased` - The purchased stock footage is accessible and searchable from Vimeo. _This field is deprecated._  * `unlisted` - The video is accessible but not searchable from Vimeo.  * `users` - Only Vimeo members can access the video. _This field is deprecated._ 
	View string `json:"view"`
}

type _VideoPrivacy VideoPrivacy

// NewVideoPrivacy instantiates a new VideoPrivacy object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewVideoPrivacy(add bool, allowShareLink bool, comments string, download bool, embed string, view string) *VideoPrivacy {
	this := VideoPrivacy{}
	this.Add = add
	this.AllowShareLink = allowShareLink
	this.Comments = comments
	this.Download = download
	this.Embed = embed
	this.View = view
	return &this
}

// NewVideoPrivacyWithDefaults instantiates a new VideoPrivacy object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewVideoPrivacyWithDefaults() *VideoPrivacy {
	this := VideoPrivacy{}
	return &this
}

// GetAdd returns the Add field value
func (o *VideoPrivacy) GetAdd() bool {
	if o == nil {
		var ret bool
		return ret
	}

	return o.Add
}

// GetAddOk returns a tuple with the Add field value
// and a boolean to check if the value has been set.
func (o *VideoPrivacy) GetAddOk() (*bool, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Add, true
}

// SetAdd sets field value
func (o *VideoPrivacy) SetAdd(v bool) {
	o.Add = v
}

// GetAllowShareLink returns the AllowShareLink field value
func (o *VideoPrivacy) GetAllowShareLink() bool {
	if o == nil {
		var ret bool
		return ret
	}

	return o.AllowShareLink
}

// GetAllowShareLinkOk returns a tuple with the AllowShareLink field value
// and a boolean to check if the value has been set.
func (o *VideoPrivacy) GetAllowShareLinkOk() (*bool, bool) {
	if o == nil {
		return nil, false
	}
	return &o.AllowShareLink, true
}

// SetAllowShareLink sets field value
func (o *VideoPrivacy) SetAllowShareLink(v bool) {
	o.AllowShareLink = v
}

// GetComments returns the Comments field value
func (o *VideoPrivacy) GetComments() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Comments
}

// GetCommentsOk returns a tuple with the Comments field value
// and a boolean to check if the value has been set.
func (o *VideoPrivacy) GetCommentsOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Comments, true
}

// SetComments sets field value
func (o *VideoPrivacy) SetComments(v string) {
	o.Comments = v
}

// GetDownload returns the Download field value
func (o *VideoPrivacy) GetDownload() bool {
	if o == nil {
		var ret bool
		return ret
	}

	return o.Download
}

// GetDownloadOk returns a tuple with the Download field value
// and a boolean to check if the value has been set.
func (o *VideoPrivacy) GetDownloadOk() (*bool, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Download, true
}

// SetDownload sets field value
func (o *VideoPrivacy) SetDownload(v bool) {
	o.Download = v
}

// GetEmbed returns the Embed field value
func (o *VideoPrivacy) GetEmbed() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Embed
}

// GetEmbedOk returns a tuple with the Embed field value
// and a boolean to check if the value has been set.
func (o *VideoPrivacy) GetEmbedOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Embed, true
}

// SetEmbed sets field value
func (o *VideoPrivacy) SetEmbed(v string) {
	o.Embed = v
}

// GetView returns the View field value
func (o *VideoPrivacy) GetView() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.View
}

// GetViewOk returns a tuple with the View field value
// and a boolean to check if the value has been set.
func (o *VideoPrivacy) GetViewOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.View, true
}

// SetView sets field value
func (o *VideoPrivacy) SetView(v string) {
	o.View = v
}

func (o VideoPrivacy) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o VideoPrivacy) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["add"] = o.Add
	toSerialize["allow_share_link"] = o.AllowShareLink
	toSerialize["comments"] = o.Comments
	toSerialize["download"] = o.Download
	toSerialize["embed"] = o.Embed
	toSerialize["view"] = o.View
	return toSerialize, nil
}

func (o *VideoPrivacy) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"add",
		"allow_share_link",
		"comments",
		"download",
		"embed",
		"view",
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

	varVideoPrivacy := _VideoPrivacy{}

	decoder := json.NewDecoder(bytes.NewReader(data))
	decoder.DisallowUnknownFields()
	err = decoder.Decode(&varVideoPrivacy)

	if err != nil {
		return err
	}

	*o = VideoPrivacy(varVideoPrivacy)

	return err
}

type NullableVideoPrivacy struct {
	value *VideoPrivacy
	isSet bool
}

func (v NullableVideoPrivacy) Get() *VideoPrivacy {
	return v.value
}

func (v *NullableVideoPrivacy) Set(val *VideoPrivacy) {
	v.value = val
	v.isSet = true
}

func (v NullableVideoPrivacy) IsSet() bool {
	return v.isSet
}

func (v *NullableVideoPrivacy) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableVideoPrivacy(val *VideoPrivacy) *NullableVideoPrivacy {
	return &NullableVideoPrivacy{value: val, isSet: true}
}

func (v NullableVideoPrivacy) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableVideoPrivacy) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


