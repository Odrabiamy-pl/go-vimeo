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

// checks if the CreateWebinarAlt1RequestPrivacy type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &CreateWebinarAlt1RequestPrivacy{}

// CreateWebinarAlt1RequestPrivacy The privacy settings of the webinar.
type CreateWebinarAlt1RequestPrivacy struct {
	// The initial embed privacy of the webinar.  Option descriptions:  * `private` - The webinar can't be embedded on any domain.  * `public` - The webinar can be embedded on any domain.  * `whitelist` - The webinar can be embedded on whitelisted domains only. 
	Embed *string `json:"embed,omitempty"`
	// The initial privacy of the webinar.  Option descriptions:  * `anybody` - Anyone can access the webinar. This privacy setting appears as `Public` on the Vimeo front end.  * `nobody` - No one except the owner can access the webinar. This privacy setting appears as `Private` on the Vimeo front end.  * `password` - Only those with the password can access the event.  * `team` - Only members of the authenticated user's team can access the webinar. 
	View *string `json:"view,omitempty"`
}

// NewCreateWebinarAlt1RequestPrivacy instantiates a new CreateWebinarAlt1RequestPrivacy object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewCreateWebinarAlt1RequestPrivacy() *CreateWebinarAlt1RequestPrivacy {
	this := CreateWebinarAlt1RequestPrivacy{}
	return &this
}

// NewCreateWebinarAlt1RequestPrivacyWithDefaults instantiates a new CreateWebinarAlt1RequestPrivacy object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewCreateWebinarAlt1RequestPrivacyWithDefaults() *CreateWebinarAlt1RequestPrivacy {
	this := CreateWebinarAlt1RequestPrivacy{}
	return &this
}

// GetEmbed returns the Embed field value if set, zero value otherwise.
func (o *CreateWebinarAlt1RequestPrivacy) GetEmbed() string {
	if o == nil || IsNil(o.Embed) {
		var ret string
		return ret
	}
	return *o.Embed
}

// GetEmbedOk returns a tuple with the Embed field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CreateWebinarAlt1RequestPrivacy) GetEmbedOk() (*string, bool) {
	if o == nil || IsNil(o.Embed) {
		return nil, false
	}
	return o.Embed, true
}

// HasEmbed returns a boolean if a field has been set.
func (o *CreateWebinarAlt1RequestPrivacy) HasEmbed() bool {
	if o != nil && !IsNil(o.Embed) {
		return true
	}

	return false
}

// SetEmbed gets a reference to the given string and assigns it to the Embed field.
func (o *CreateWebinarAlt1RequestPrivacy) SetEmbed(v string) {
	o.Embed = &v
}

// GetView returns the View field value if set, zero value otherwise.
func (o *CreateWebinarAlt1RequestPrivacy) GetView() string {
	if o == nil || IsNil(o.View) {
		var ret string
		return ret
	}
	return *o.View
}

// GetViewOk returns a tuple with the View field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CreateWebinarAlt1RequestPrivacy) GetViewOk() (*string, bool) {
	if o == nil || IsNil(o.View) {
		return nil, false
	}
	return o.View, true
}

// HasView returns a boolean if a field has been set.
func (o *CreateWebinarAlt1RequestPrivacy) HasView() bool {
	if o != nil && !IsNil(o.View) {
		return true
	}

	return false
}

// SetView gets a reference to the given string and assigns it to the View field.
func (o *CreateWebinarAlt1RequestPrivacy) SetView(v string) {
	o.View = &v
}

func (o CreateWebinarAlt1RequestPrivacy) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o CreateWebinarAlt1RequestPrivacy) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Embed) {
		toSerialize["embed"] = o.Embed
	}
	if !IsNil(o.View) {
		toSerialize["view"] = o.View
	}
	return toSerialize, nil
}

type NullableCreateWebinarAlt1RequestPrivacy struct {
	value *CreateWebinarAlt1RequestPrivacy
	isSet bool
}

func (v NullableCreateWebinarAlt1RequestPrivacy) Get() *CreateWebinarAlt1RequestPrivacy {
	return v.value
}

func (v *NullableCreateWebinarAlt1RequestPrivacy) Set(val *CreateWebinarAlt1RequestPrivacy) {
	v.value = val
	v.isSet = true
}

func (v NullableCreateWebinarAlt1RequestPrivacy) IsSet() bool {
	return v.isSet
}

func (v *NullableCreateWebinarAlt1RequestPrivacy) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableCreateWebinarAlt1RequestPrivacy(val *CreateWebinarAlt1RequestPrivacy) *NullableCreateWebinarAlt1RequestPrivacy {
	return &NullableCreateWebinarAlt1RequestPrivacy{value: val, isSet: true}
}

func (v NullableCreateWebinarAlt1RequestPrivacy) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableCreateWebinarAlt1RequestPrivacy) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


