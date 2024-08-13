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

// checks if the UserPreferencesVideosPrivacy type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &UserPreferencesVideosPrivacy{}

// UserPreferencesVideosPrivacy struct for UserPreferencesVideosPrivacy
type UserPreferencesVideosPrivacy struct {
	// Whether other users can add the authenticated user's videos.
	Add *bool `json:"add,omitempty"`
	// The authenticated user's privacy preference for comments.  Option descriptions:  * `anybody` - Anyone can comment on the user's videos.  * `contacts` - Only contacts can comment on the user's videos.  * `nobody` - No one can comment on the user's videos. 
	Comments *string `json:"comments,omitempty"`
	// Whether other users can download the authenticated user's videos.
	Download *bool `json:"download,omitempty"`
	// The authenticated user's privacy preference for embeds.  Option descriptions:  * `private` - Only the user can embed their own videos.  * `public` - Anyone can embed the user's videos.  * `whitelist` - Only those on the whitelist can embed the user's videos. 
	Embed *string `json:"embed,omitempty"`
	// The default password for the video.
	Password *string `json:"password,omitempty"`
	// The authenticated user's privacy preference for views.  Option descriptions:  * `anybody` - Anyone can view the user's videos. This privacy setting appears as `Public` on the Vimeo front end.  * `contacts` - Only contacts can view the user's videos. _This field is deprecated._  * `disable` - Views are disabled for the user's videos. This privacy setting appears as `Hide from Vimeo` on the Vimeo front end.  * `nobody` - No one except the user can view the user's videos. This privacy setting appears as `Private` on the Vimeo front end.  * `password` - Only those with the password can view the user's videos.  * `unlisted` - Anybody can view the user's videos if they have a link.  * `users` - Only other Vimeo members can view the user's videos. _This field is deprecated._ 
	View *string `json:"view,omitempty"`
}

// NewUserPreferencesVideosPrivacy instantiates a new UserPreferencesVideosPrivacy object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewUserPreferencesVideosPrivacy() *UserPreferencesVideosPrivacy {
	this := UserPreferencesVideosPrivacy{}
	return &this
}

// NewUserPreferencesVideosPrivacyWithDefaults instantiates a new UserPreferencesVideosPrivacy object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewUserPreferencesVideosPrivacyWithDefaults() *UserPreferencesVideosPrivacy {
	this := UserPreferencesVideosPrivacy{}
	return &this
}

// GetAdd returns the Add field value if set, zero value otherwise.
func (o *UserPreferencesVideosPrivacy) GetAdd() bool {
	if o == nil || IsNil(o.Add) {
		var ret bool
		return ret
	}
	return *o.Add
}

// GetAddOk returns a tuple with the Add field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UserPreferencesVideosPrivacy) GetAddOk() (*bool, bool) {
	if o == nil || IsNil(o.Add) {
		return nil, false
	}
	return o.Add, true
}

// HasAdd returns a boolean if a field has been set.
func (o *UserPreferencesVideosPrivacy) HasAdd() bool {
	if o != nil && !IsNil(o.Add) {
		return true
	}

	return false
}

// SetAdd gets a reference to the given bool and assigns it to the Add field.
func (o *UserPreferencesVideosPrivacy) SetAdd(v bool) {
	o.Add = &v
}

// GetComments returns the Comments field value if set, zero value otherwise.
func (o *UserPreferencesVideosPrivacy) GetComments() string {
	if o == nil || IsNil(o.Comments) {
		var ret string
		return ret
	}
	return *o.Comments
}

// GetCommentsOk returns a tuple with the Comments field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UserPreferencesVideosPrivacy) GetCommentsOk() (*string, bool) {
	if o == nil || IsNil(o.Comments) {
		return nil, false
	}
	return o.Comments, true
}

// HasComments returns a boolean if a field has been set.
func (o *UserPreferencesVideosPrivacy) HasComments() bool {
	if o != nil && !IsNil(o.Comments) {
		return true
	}

	return false
}

// SetComments gets a reference to the given string and assigns it to the Comments field.
func (o *UserPreferencesVideosPrivacy) SetComments(v string) {
	o.Comments = &v
}

// GetDownload returns the Download field value if set, zero value otherwise.
func (o *UserPreferencesVideosPrivacy) GetDownload() bool {
	if o == nil || IsNil(o.Download) {
		var ret bool
		return ret
	}
	return *o.Download
}

// GetDownloadOk returns a tuple with the Download field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UserPreferencesVideosPrivacy) GetDownloadOk() (*bool, bool) {
	if o == nil || IsNil(o.Download) {
		return nil, false
	}
	return o.Download, true
}

// HasDownload returns a boolean if a field has been set.
func (o *UserPreferencesVideosPrivacy) HasDownload() bool {
	if o != nil && !IsNil(o.Download) {
		return true
	}

	return false
}

// SetDownload gets a reference to the given bool and assigns it to the Download field.
func (o *UserPreferencesVideosPrivacy) SetDownload(v bool) {
	o.Download = &v
}

// GetEmbed returns the Embed field value if set, zero value otherwise.
func (o *UserPreferencesVideosPrivacy) GetEmbed() string {
	if o == nil || IsNil(o.Embed) {
		var ret string
		return ret
	}
	return *o.Embed
}

// GetEmbedOk returns a tuple with the Embed field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UserPreferencesVideosPrivacy) GetEmbedOk() (*string, bool) {
	if o == nil || IsNil(o.Embed) {
		return nil, false
	}
	return o.Embed, true
}

// HasEmbed returns a boolean if a field has been set.
func (o *UserPreferencesVideosPrivacy) HasEmbed() bool {
	if o != nil && !IsNil(o.Embed) {
		return true
	}

	return false
}

// SetEmbed gets a reference to the given string and assigns it to the Embed field.
func (o *UserPreferencesVideosPrivacy) SetEmbed(v string) {
	o.Embed = &v
}

// GetPassword returns the Password field value if set, zero value otherwise.
func (o *UserPreferencesVideosPrivacy) GetPassword() string {
	if o == nil || IsNil(o.Password) {
		var ret string
		return ret
	}
	return *o.Password
}

// GetPasswordOk returns a tuple with the Password field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UserPreferencesVideosPrivacy) GetPasswordOk() (*string, bool) {
	if o == nil || IsNil(o.Password) {
		return nil, false
	}
	return o.Password, true
}

// HasPassword returns a boolean if a field has been set.
func (o *UserPreferencesVideosPrivacy) HasPassword() bool {
	if o != nil && !IsNil(o.Password) {
		return true
	}

	return false
}

// SetPassword gets a reference to the given string and assigns it to the Password field.
func (o *UserPreferencesVideosPrivacy) SetPassword(v string) {
	o.Password = &v
}

// GetView returns the View field value if set, zero value otherwise.
func (o *UserPreferencesVideosPrivacy) GetView() string {
	if o == nil || IsNil(o.View) {
		var ret string
		return ret
	}
	return *o.View
}

// GetViewOk returns a tuple with the View field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UserPreferencesVideosPrivacy) GetViewOk() (*string, bool) {
	if o == nil || IsNil(o.View) {
		return nil, false
	}
	return o.View, true
}

// HasView returns a boolean if a field has been set.
func (o *UserPreferencesVideosPrivacy) HasView() bool {
	if o != nil && !IsNil(o.View) {
		return true
	}

	return false
}

// SetView gets a reference to the given string and assigns it to the View field.
func (o *UserPreferencesVideosPrivacy) SetView(v string) {
	o.View = &v
}

func (o UserPreferencesVideosPrivacy) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o UserPreferencesVideosPrivacy) ToMap() (map[string]interface{}, error) {
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
	if !IsNil(o.Password) {
		toSerialize["password"] = o.Password
	}
	if !IsNil(o.View) {
		toSerialize["view"] = o.View
	}
	return toSerialize, nil
}

type NullableUserPreferencesVideosPrivacy struct {
	value *UserPreferencesVideosPrivacy
	isSet bool
}

func (v NullableUserPreferencesVideosPrivacy) Get() *UserPreferencesVideosPrivacy {
	return v.value
}

func (v *NullableUserPreferencesVideosPrivacy) Set(val *UserPreferencesVideosPrivacy) {
	v.value = val
	v.isSet = true
}

func (v NullableUserPreferencesVideosPrivacy) IsSet() bool {
	return v.isSet
}

func (v *NullableUserPreferencesVideosPrivacy) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableUserPreferencesVideosPrivacy(val *UserPreferencesVideosPrivacy) *NullableUserPreferencesVideosPrivacy {
	return &NullableUserPreferencesVideosPrivacy{value: val, isSet: true}
}

func (v NullableUserPreferencesVideosPrivacy) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableUserPreferencesVideosPrivacy) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


