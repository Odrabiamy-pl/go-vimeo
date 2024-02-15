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

// checks if the EditUserAlt1RequestVideosPrivacy type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &EditUserAlt1RequestVideosPrivacy{}

// EditUserAlt1RequestVideosPrivacy struct for EditUserAlt1RequestVideosPrivacy
type EditUserAlt1RequestVideosPrivacy struct {
	// Whether the user can add videos to showcases, channels, or groups by default.
	Add *bool `json:"add,omitempty"`
	// Who can comment on the user's video uploads by default.  Option descriptions:  * `anybody` - Anyone can comment.  * `contacts` - Only the user's contacts can comment.  * `nobody` - No one can comment. 
	Comments *string `json:"comments,omitempty"`
	// Whether the user can download videos. This value becomes the default download setting for all future videos that the user uploads.
	Download *bool `json:"download,omitempty"`
	// The privacy for the user's embedded videos. The whitelist value enables you to define all valid embeddable domains. See our [Interacting with Videos](https://developer.vimeo.com/api/guides/videos/interact#set-off-site-privacy) guide for details on adding and removing domains.  Option descriptions:  * `private` - The videos can't be embedded on any domain.  * `public` - The videos can be embedded on any domain.  * `whitelist` - The videos can be embedded on the specified domains only. 
	Embed *string `json:"embed,omitempty"`
	// Who can access the user's videos by default.  Option descriptions:  * `anybody` - Anyone can access the videos. This privacy setting appears as `Public` on the Vimeo front end.  * `contacts` - Only the user's contacts can access the videos. _This field is deprecated._  * `disable` - The videos are disabled. This privacy setting appears as `Hide from Vimeo` on the Vimeo front end.  * `nobody` - No one can access the videos. This privacy setting appears as `Private` on the Vimeo front end.  * `password` - Only those with the password can access the videos.  * `unlisted` - The videos are unlisted.  * `users` - Only other Vimeo members can access the videos. _This field is deprecated._ 
	View *string `json:"view,omitempty"`
}

// NewEditUserAlt1RequestVideosPrivacy instantiates a new EditUserAlt1RequestVideosPrivacy object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewEditUserAlt1RequestVideosPrivacy() *EditUserAlt1RequestVideosPrivacy {
	this := EditUserAlt1RequestVideosPrivacy{}
	return &this
}

// NewEditUserAlt1RequestVideosPrivacyWithDefaults instantiates a new EditUserAlt1RequestVideosPrivacy object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewEditUserAlt1RequestVideosPrivacyWithDefaults() *EditUserAlt1RequestVideosPrivacy {
	this := EditUserAlt1RequestVideosPrivacy{}
	return &this
}

// GetAdd returns the Add field value if set, zero value otherwise.
func (o *EditUserAlt1RequestVideosPrivacy) GetAdd() bool {
	if o == nil || IsNil(o.Add) {
		var ret bool
		return ret
	}
	return *o.Add
}

// GetAddOk returns a tuple with the Add field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *EditUserAlt1RequestVideosPrivacy) GetAddOk() (*bool, bool) {
	if o == nil || IsNil(o.Add) {
		return nil, false
	}
	return o.Add, true
}

// HasAdd returns a boolean if a field has been set.
func (o *EditUserAlt1RequestVideosPrivacy) HasAdd() bool {
	if o != nil && !IsNil(o.Add) {
		return true
	}

	return false
}

// SetAdd gets a reference to the given bool and assigns it to the Add field.
func (o *EditUserAlt1RequestVideosPrivacy) SetAdd(v bool) {
	o.Add = &v
}

// GetComments returns the Comments field value if set, zero value otherwise.
func (o *EditUserAlt1RequestVideosPrivacy) GetComments() string {
	if o == nil || IsNil(o.Comments) {
		var ret string
		return ret
	}
	return *o.Comments
}

// GetCommentsOk returns a tuple with the Comments field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *EditUserAlt1RequestVideosPrivacy) GetCommentsOk() (*string, bool) {
	if o == nil || IsNil(o.Comments) {
		return nil, false
	}
	return o.Comments, true
}

// HasComments returns a boolean if a field has been set.
func (o *EditUserAlt1RequestVideosPrivacy) HasComments() bool {
	if o != nil && !IsNil(o.Comments) {
		return true
	}

	return false
}

// SetComments gets a reference to the given string and assigns it to the Comments field.
func (o *EditUserAlt1RequestVideosPrivacy) SetComments(v string) {
	o.Comments = &v
}

// GetDownload returns the Download field value if set, zero value otherwise.
func (o *EditUserAlt1RequestVideosPrivacy) GetDownload() bool {
	if o == nil || IsNil(o.Download) {
		var ret bool
		return ret
	}
	return *o.Download
}

// GetDownloadOk returns a tuple with the Download field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *EditUserAlt1RequestVideosPrivacy) GetDownloadOk() (*bool, bool) {
	if o == nil || IsNil(o.Download) {
		return nil, false
	}
	return o.Download, true
}

// HasDownload returns a boolean if a field has been set.
func (o *EditUserAlt1RequestVideosPrivacy) HasDownload() bool {
	if o != nil && !IsNil(o.Download) {
		return true
	}

	return false
}

// SetDownload gets a reference to the given bool and assigns it to the Download field.
func (o *EditUserAlt1RequestVideosPrivacy) SetDownload(v bool) {
	o.Download = &v
}

// GetEmbed returns the Embed field value if set, zero value otherwise.
func (o *EditUserAlt1RequestVideosPrivacy) GetEmbed() string {
	if o == nil || IsNil(o.Embed) {
		var ret string
		return ret
	}
	return *o.Embed
}

// GetEmbedOk returns a tuple with the Embed field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *EditUserAlt1RequestVideosPrivacy) GetEmbedOk() (*string, bool) {
	if o == nil || IsNil(o.Embed) {
		return nil, false
	}
	return o.Embed, true
}

// HasEmbed returns a boolean if a field has been set.
func (o *EditUserAlt1RequestVideosPrivacy) HasEmbed() bool {
	if o != nil && !IsNil(o.Embed) {
		return true
	}

	return false
}

// SetEmbed gets a reference to the given string and assigns it to the Embed field.
func (o *EditUserAlt1RequestVideosPrivacy) SetEmbed(v string) {
	o.Embed = &v
}

// GetView returns the View field value if set, zero value otherwise.
func (o *EditUserAlt1RequestVideosPrivacy) GetView() string {
	if o == nil || IsNil(o.View) {
		var ret string
		return ret
	}
	return *o.View
}

// GetViewOk returns a tuple with the View field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *EditUserAlt1RequestVideosPrivacy) GetViewOk() (*string, bool) {
	if o == nil || IsNil(o.View) {
		return nil, false
	}
	return o.View, true
}

// HasView returns a boolean if a field has been set.
func (o *EditUserAlt1RequestVideosPrivacy) HasView() bool {
	if o != nil && !IsNil(o.View) {
		return true
	}

	return false
}

// SetView gets a reference to the given string and assigns it to the View field.
func (o *EditUserAlt1RequestVideosPrivacy) SetView(v string) {
	o.View = &v
}

func (o EditUserAlt1RequestVideosPrivacy) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o EditUserAlt1RequestVideosPrivacy) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Add) {
		toSerialize["add"] = o.Add
	}
	if !IsNil(o.Comments) {
		toSerialize["comments"] = o.Comments
	}
	if !IsNil(o.Download) {
		toSerialize["download"] = o.Download
	}
	if !IsNil(o.Embed) {
		toSerialize["embed"] = o.Embed
	}
	if !IsNil(o.View) {
		toSerialize["view"] = o.View
	}
	return toSerialize, nil
}

type NullableEditUserAlt1RequestVideosPrivacy struct {
	value *EditUserAlt1RequestVideosPrivacy
	isSet bool
}

func (v NullableEditUserAlt1RequestVideosPrivacy) Get() *EditUserAlt1RequestVideosPrivacy {
	return v.value
}

func (v *NullableEditUserAlt1RequestVideosPrivacy) Set(val *EditUserAlt1RequestVideosPrivacy) {
	v.value = val
	v.isSet = true
}

func (v NullableEditUserAlt1RequestVideosPrivacy) IsSet() bool {
	return v.isSet
}

func (v *NullableEditUserAlt1RequestVideosPrivacy) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableEditUserAlt1RequestVideosPrivacy(val *EditUserAlt1RequestVideosPrivacy) *NullableEditUserAlt1RequestVideosPrivacy {
	return &NullableEditUserAlt1RequestVideosPrivacy{value: val, isSet: true}
}

func (v NullableEditUserAlt1RequestVideosPrivacy) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableEditUserAlt1RequestVideosPrivacy) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


