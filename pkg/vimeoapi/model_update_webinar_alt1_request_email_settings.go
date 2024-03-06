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

// checks if the UpdateWebinarAlt1RequestEmailSettings type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &UpdateWebinarAlt1RequestEmailSettings{}

// UpdateWebinarAlt1RequestEmailSettings The settings for emails that are sent about the webinar. _This field is deprecated._
type UpdateWebinarAlt1RequestEmailSettings struct {
	// The accent color scheme for emails that are sent about the webinar. _This field is deprecated._
	AccentColor *string `json:"accent_color,omitempty"`
	// The custom link for emails that are sent about the webinar. _This field is deprecated._
	CustomLink       *string                                                `json:"custom_link,omitempty"`
	EmailPreferences *CreateWebinarAlt1RequestEmailSettingsEmailPreferences `json:"email_preferences,omitempty"`
	// The name of the sender for emails that are sent about the webinar. _This field is deprecated._
	From *string `json:"from,omitempty"`
	// The URI of the logo image to include in emails that are sent about the webinar. _This field is deprecated._
	LogoUri *string `json:"logo_uri,omitempty"`
	// The sender's reply email address. _This field is deprecated._
	ReplyEmail *string `json:"reply_email,omitempty"`
	// The sender's physical address. _This field is deprecated._
	SenderAddress *string `json:"sender_address,omitempty"`
	// The URL of the sender's privacy policy. _This field is deprecated._
	SenderPolicyUrl *string `json:"sender_policy_url,omitempty"`
	// Whether to include a custom link in emails that are sent about the webinar. _This field is deprecated._
	UseCustomLink *bool `json:"use_custom_link,omitempty"`
	// Whether to include a reply link in the footer of emails that are sent about the webinar. _This field is deprecated._
	UseReplyEmail *bool `json:"use_reply_email,omitempty"`
	// Whether to include the sender's physical address in the footer of emails that are sent about the webinar. _This field is deprecated._
	UseSenderAddress *bool `json:"use_sender_address,omitempty"`
	// Whether to include the URL of the sender's privacy policy in the footer of emails that are sent about the webinar. _This field is deprecated._
	UseSenderPolicyUrl *bool `json:"use_sender_policy_url,omitempty"`
}

// NewUpdateWebinarAlt1RequestEmailSettings instantiates a new UpdateWebinarAlt1RequestEmailSettings object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewUpdateWebinarAlt1RequestEmailSettings() *UpdateWebinarAlt1RequestEmailSettings {
	this := UpdateWebinarAlt1RequestEmailSettings{}
	return &this
}

// NewUpdateWebinarAlt1RequestEmailSettingsWithDefaults instantiates a new UpdateWebinarAlt1RequestEmailSettings object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewUpdateWebinarAlt1RequestEmailSettingsWithDefaults() *UpdateWebinarAlt1RequestEmailSettings {
	this := UpdateWebinarAlt1RequestEmailSettings{}
	return &this
}

// GetAccentColor returns the AccentColor field value if set, zero value otherwise.
func (o *UpdateWebinarAlt1RequestEmailSettings) GetAccentColor() string {
	if o == nil || IsNil(o.AccentColor) {
		var ret string
		return ret
	}
	return *o.AccentColor
}

// GetAccentColorOk returns a tuple with the AccentColor field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UpdateWebinarAlt1RequestEmailSettings) GetAccentColorOk() (*string, bool) {
	if o == nil || IsNil(o.AccentColor) {
		return nil, false
	}
	return o.AccentColor, true
}

// HasAccentColor returns a boolean if a field has been set.
func (o *UpdateWebinarAlt1RequestEmailSettings) HasAccentColor() bool {
	if o != nil && !IsNil(o.AccentColor) {
		return true
	}

	return false
}

// SetAccentColor gets a reference to the given string and assigns it to the AccentColor field.
func (o *UpdateWebinarAlt1RequestEmailSettings) SetAccentColor(v string) {
	o.AccentColor = &v
}

// GetCustomLink returns the CustomLink field value if set, zero value otherwise.
func (o *UpdateWebinarAlt1RequestEmailSettings) GetCustomLink() string {
	if o == nil || IsNil(o.CustomLink) {
		var ret string
		return ret
	}
	return *o.CustomLink
}

// GetCustomLinkOk returns a tuple with the CustomLink field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UpdateWebinarAlt1RequestEmailSettings) GetCustomLinkOk() (*string, bool) {
	if o == nil || IsNil(o.CustomLink) {
		return nil, false
	}
	return o.CustomLink, true
}

// HasCustomLink returns a boolean if a field has been set.
func (o *UpdateWebinarAlt1RequestEmailSettings) HasCustomLink() bool {
	if o != nil && !IsNil(o.CustomLink) {
		return true
	}

	return false
}

// SetCustomLink gets a reference to the given string and assigns it to the CustomLink field.
func (o *UpdateWebinarAlt1RequestEmailSettings) SetCustomLink(v string) {
	o.CustomLink = &v
}

// GetEmailPreferences returns the EmailPreferences field value if set, zero value otherwise.
func (o *UpdateWebinarAlt1RequestEmailSettings) GetEmailPreferences() CreateWebinarAlt1RequestEmailSettingsEmailPreferences {
	if o == nil || IsNil(o.EmailPreferences) {
		var ret CreateWebinarAlt1RequestEmailSettingsEmailPreferences
		return ret
	}
	return *o.EmailPreferences
}

// GetEmailPreferencesOk returns a tuple with the EmailPreferences field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UpdateWebinarAlt1RequestEmailSettings) GetEmailPreferencesOk() (*CreateWebinarAlt1RequestEmailSettingsEmailPreferences, bool) {
	if o == nil || IsNil(o.EmailPreferences) {
		return nil, false
	}
	return o.EmailPreferences, true
}

// HasEmailPreferences returns a boolean if a field has been set.
func (o *UpdateWebinarAlt1RequestEmailSettings) HasEmailPreferences() bool {
	if o != nil && !IsNil(o.EmailPreferences) {
		return true
	}

	return false
}

// SetEmailPreferences gets a reference to the given CreateWebinarAlt1RequestEmailSettingsEmailPreferences and assigns it to the EmailPreferences field.
func (o *UpdateWebinarAlt1RequestEmailSettings) SetEmailPreferences(v CreateWebinarAlt1RequestEmailSettingsEmailPreferences) {
	o.EmailPreferences = &v
}

// GetFrom returns the From field value if set, zero value otherwise.
func (o *UpdateWebinarAlt1RequestEmailSettings) GetFrom() string {
	if o == nil || IsNil(o.From) {
		var ret string
		return ret
	}
	return *o.From
}

// GetFromOk returns a tuple with the From field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UpdateWebinarAlt1RequestEmailSettings) GetFromOk() (*string, bool) {
	if o == nil || IsNil(o.From) {
		return nil, false
	}
	return o.From, true
}

// HasFrom returns a boolean if a field has been set.
func (o *UpdateWebinarAlt1RequestEmailSettings) HasFrom() bool {
	if o != nil && !IsNil(o.From) {
		return true
	}

	return false
}

// SetFrom gets a reference to the given string and assigns it to the From field.
func (o *UpdateWebinarAlt1RequestEmailSettings) SetFrom(v string) {
	o.From = &v
}

// GetLogoUri returns the LogoUri field value if set, zero value otherwise.
func (o *UpdateWebinarAlt1RequestEmailSettings) GetLogoUri() string {
	if o == nil || IsNil(o.LogoUri) {
		var ret string
		return ret
	}
	return *o.LogoUri
}

// GetLogoUriOk returns a tuple with the LogoUri field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UpdateWebinarAlt1RequestEmailSettings) GetLogoUriOk() (*string, bool) {
	if o == nil || IsNil(o.LogoUri) {
		return nil, false
	}
	return o.LogoUri, true
}

// HasLogoUri returns a boolean if a field has been set.
func (o *UpdateWebinarAlt1RequestEmailSettings) HasLogoUri() bool {
	if o != nil && !IsNil(o.LogoUri) {
		return true
	}

	return false
}

// SetLogoUri gets a reference to the given string and assigns it to the LogoUri field.
func (o *UpdateWebinarAlt1RequestEmailSettings) SetLogoUri(v string) {
	o.LogoUri = &v
}

// GetReplyEmail returns the ReplyEmail field value if set, zero value otherwise.
func (o *UpdateWebinarAlt1RequestEmailSettings) GetReplyEmail() string {
	if o == nil || IsNil(o.ReplyEmail) {
		var ret string
		return ret
	}
	return *o.ReplyEmail
}

// GetReplyEmailOk returns a tuple with the ReplyEmail field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UpdateWebinarAlt1RequestEmailSettings) GetReplyEmailOk() (*string, bool) {
	if o == nil || IsNil(o.ReplyEmail) {
		return nil, false
	}
	return o.ReplyEmail, true
}

// HasReplyEmail returns a boolean if a field has been set.
func (o *UpdateWebinarAlt1RequestEmailSettings) HasReplyEmail() bool {
	if o != nil && !IsNil(o.ReplyEmail) {
		return true
	}

	return false
}

// SetReplyEmail gets a reference to the given string and assigns it to the ReplyEmail field.
func (o *UpdateWebinarAlt1RequestEmailSettings) SetReplyEmail(v string) {
	o.ReplyEmail = &v
}

// GetSenderAddress returns the SenderAddress field value if set, zero value otherwise.
func (o *UpdateWebinarAlt1RequestEmailSettings) GetSenderAddress() string {
	if o == nil || IsNil(o.SenderAddress) {
		var ret string
		return ret
	}
	return *o.SenderAddress
}

// GetSenderAddressOk returns a tuple with the SenderAddress field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UpdateWebinarAlt1RequestEmailSettings) GetSenderAddressOk() (*string, bool) {
	if o == nil || IsNil(o.SenderAddress) {
		return nil, false
	}
	return o.SenderAddress, true
}

// HasSenderAddress returns a boolean if a field has been set.
func (o *UpdateWebinarAlt1RequestEmailSettings) HasSenderAddress() bool {
	if o != nil && !IsNil(o.SenderAddress) {
		return true
	}

	return false
}

// SetSenderAddress gets a reference to the given string and assigns it to the SenderAddress field.
func (o *UpdateWebinarAlt1RequestEmailSettings) SetSenderAddress(v string) {
	o.SenderAddress = &v
}

// GetSenderPolicyUrl returns the SenderPolicyUrl field value if set, zero value otherwise.
func (o *UpdateWebinarAlt1RequestEmailSettings) GetSenderPolicyUrl() string {
	if o == nil || IsNil(o.SenderPolicyUrl) {
		var ret string
		return ret
	}
	return *o.SenderPolicyUrl
}

// GetSenderPolicyUrlOk returns a tuple with the SenderPolicyUrl field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UpdateWebinarAlt1RequestEmailSettings) GetSenderPolicyUrlOk() (*string, bool) {
	if o == nil || IsNil(o.SenderPolicyUrl) {
		return nil, false
	}
	return o.SenderPolicyUrl, true
}

// HasSenderPolicyUrl returns a boolean if a field has been set.
func (o *UpdateWebinarAlt1RequestEmailSettings) HasSenderPolicyUrl() bool {
	if o != nil && !IsNil(o.SenderPolicyUrl) {
		return true
	}

	return false
}

// SetSenderPolicyUrl gets a reference to the given string and assigns it to the SenderPolicyUrl field.
func (o *UpdateWebinarAlt1RequestEmailSettings) SetSenderPolicyUrl(v string) {
	o.SenderPolicyUrl = &v
}

// GetUseCustomLink returns the UseCustomLink field value if set, zero value otherwise.
func (o *UpdateWebinarAlt1RequestEmailSettings) GetUseCustomLink() bool {
	if o == nil || IsNil(o.UseCustomLink) {
		var ret bool
		return ret
	}
	return *o.UseCustomLink
}

// GetUseCustomLinkOk returns a tuple with the UseCustomLink field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UpdateWebinarAlt1RequestEmailSettings) GetUseCustomLinkOk() (*bool, bool) {
	if o == nil || IsNil(o.UseCustomLink) {
		return nil, false
	}
	return o.UseCustomLink, true
}

// HasUseCustomLink returns a boolean if a field has been set.
func (o *UpdateWebinarAlt1RequestEmailSettings) HasUseCustomLink() bool {
	if o != nil && !IsNil(o.UseCustomLink) {
		return true
	}

	return false
}

// SetUseCustomLink gets a reference to the given bool and assigns it to the UseCustomLink field.
func (o *UpdateWebinarAlt1RequestEmailSettings) SetUseCustomLink(v bool) {
	o.UseCustomLink = &v
}

// GetUseReplyEmail returns the UseReplyEmail field value if set, zero value otherwise.
func (o *UpdateWebinarAlt1RequestEmailSettings) GetUseReplyEmail() bool {
	if o == nil || IsNil(o.UseReplyEmail) {
		var ret bool
		return ret
	}
	return *o.UseReplyEmail
}

// GetUseReplyEmailOk returns a tuple with the UseReplyEmail field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UpdateWebinarAlt1RequestEmailSettings) GetUseReplyEmailOk() (*bool, bool) {
	if o == nil || IsNil(o.UseReplyEmail) {
		return nil, false
	}
	return o.UseReplyEmail, true
}

// HasUseReplyEmail returns a boolean if a field has been set.
func (o *UpdateWebinarAlt1RequestEmailSettings) HasUseReplyEmail() bool {
	if o != nil && !IsNil(o.UseReplyEmail) {
		return true
	}

	return false
}

// SetUseReplyEmail gets a reference to the given bool and assigns it to the UseReplyEmail field.
func (o *UpdateWebinarAlt1RequestEmailSettings) SetUseReplyEmail(v bool) {
	o.UseReplyEmail = &v
}

// GetUseSenderAddress returns the UseSenderAddress field value if set, zero value otherwise.
func (o *UpdateWebinarAlt1RequestEmailSettings) GetUseSenderAddress() bool {
	if o == nil || IsNil(o.UseSenderAddress) {
		var ret bool
		return ret
	}
	return *o.UseSenderAddress
}

// GetUseSenderAddressOk returns a tuple with the UseSenderAddress field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UpdateWebinarAlt1RequestEmailSettings) GetUseSenderAddressOk() (*bool, bool) {
	if o == nil || IsNil(o.UseSenderAddress) {
		return nil, false
	}
	return o.UseSenderAddress, true
}

// HasUseSenderAddress returns a boolean if a field has been set.
func (o *UpdateWebinarAlt1RequestEmailSettings) HasUseSenderAddress() bool {
	if o != nil && !IsNil(o.UseSenderAddress) {
		return true
	}

	return false
}

// SetUseSenderAddress gets a reference to the given bool and assigns it to the UseSenderAddress field.
func (o *UpdateWebinarAlt1RequestEmailSettings) SetUseSenderAddress(v bool) {
	o.UseSenderAddress = &v
}

// GetUseSenderPolicyUrl returns the UseSenderPolicyUrl field value if set, zero value otherwise.
func (o *UpdateWebinarAlt1RequestEmailSettings) GetUseSenderPolicyUrl() bool {
	if o == nil || IsNil(o.UseSenderPolicyUrl) {
		var ret bool
		return ret
	}
	return *o.UseSenderPolicyUrl
}

// GetUseSenderPolicyUrlOk returns a tuple with the UseSenderPolicyUrl field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UpdateWebinarAlt1RequestEmailSettings) GetUseSenderPolicyUrlOk() (*bool, bool) {
	if o == nil || IsNil(o.UseSenderPolicyUrl) {
		return nil, false
	}
	return o.UseSenderPolicyUrl, true
}

// HasUseSenderPolicyUrl returns a boolean if a field has been set.
func (o *UpdateWebinarAlt1RequestEmailSettings) HasUseSenderPolicyUrl() bool {
	if o != nil && !IsNil(o.UseSenderPolicyUrl) {
		return true
	}

	return false
}

// SetUseSenderPolicyUrl gets a reference to the given bool and assigns it to the UseSenderPolicyUrl field.
func (o *UpdateWebinarAlt1RequestEmailSettings) SetUseSenderPolicyUrl(v bool) {
	o.UseSenderPolicyUrl = &v
}

func (o UpdateWebinarAlt1RequestEmailSettings) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o UpdateWebinarAlt1RequestEmailSettings) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.AccentColor) {
		toSerialize["accent_color"] = o.AccentColor
	}
	if !IsNil(o.CustomLink) {
		toSerialize["custom_link"] = o.CustomLink
	}
	if !IsNil(o.EmailPreferences) {
		toSerialize["email_preferences"] = o.EmailPreferences
	}
	if !IsNil(o.From) {
		toSerialize["from"] = o.From
	}
	if !IsNil(o.LogoUri) {
		toSerialize["logo_uri"] = o.LogoUri
	}
	if !IsNil(o.ReplyEmail) {
		toSerialize["reply_email"] = o.ReplyEmail
	}
	if !IsNil(o.SenderAddress) {
		toSerialize["sender_address"] = o.SenderAddress
	}
	if !IsNil(o.SenderPolicyUrl) {
		toSerialize["sender_policy_url"] = o.SenderPolicyUrl
	}
	if !IsNil(o.UseCustomLink) {
		toSerialize["use_custom_link"] = o.UseCustomLink
	}
	if !IsNil(o.UseReplyEmail) {
		toSerialize["use_reply_email"] = o.UseReplyEmail
	}
	if !IsNil(o.UseSenderAddress) {
		toSerialize["use_sender_address"] = o.UseSenderAddress
	}
	if !IsNil(o.UseSenderPolicyUrl) {
		toSerialize["use_sender_policy_url"] = o.UseSenderPolicyUrl
	}
	return toSerialize, nil
}

type NullableUpdateWebinarAlt1RequestEmailSettings struct {
	value *UpdateWebinarAlt1RequestEmailSettings
	isSet bool
}

func (v NullableUpdateWebinarAlt1RequestEmailSettings) Get() *UpdateWebinarAlt1RequestEmailSettings {
	return v.value
}

func (v *NullableUpdateWebinarAlt1RequestEmailSettings) Set(val *UpdateWebinarAlt1RequestEmailSettings) {
	v.value = val
	v.isSet = true
}

func (v NullableUpdateWebinarAlt1RequestEmailSettings) IsSet() bool {
	return v.isSet
}

func (v *NullableUpdateWebinarAlt1RequestEmailSettings) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableUpdateWebinarAlt1RequestEmailSettings(val *UpdateWebinarAlt1RequestEmailSettings) *NullableUpdateWebinarAlt1RequestEmailSettings {
	return &NullableUpdateWebinarAlt1RequestEmailSettings{value: val, isSet: true}
}

func (v NullableUpdateWebinarAlt1RequestEmailSettings) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableUpdateWebinarAlt1RequestEmailSettings) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
