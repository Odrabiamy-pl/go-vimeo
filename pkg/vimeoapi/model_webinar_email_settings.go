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

// checks if the WebinarEmailSettings type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &WebinarEmailSettings{}

// WebinarEmailSettings struct for WebinarEmailSettings
type WebinarEmailSettings struct {
	// The accent color scheme for emails that are sent about the webinar.
	AccentColor string `json:"accent_color"`
	// The custom link for emails that are sent about the webinar.
	CustomLink NullableString `json:"custom_link"`
	EmailEventReminder24Hrs WebinarEmailSettingsEmailEventReminder24Hrs `json:"email_event_reminder_24_hrs"`
	EmailPostEventThankYou WebinarEmailSettingsEmailPostEventThankYou `json:"email_post_event_thank_you"`
	// A list of preferences for the emails to send during the webinar event.
	EmailPreferences []string `json:"email_preferences"`
	EmailRegistrationConfirmation WebinarEmailSettingsEmailRegistrationConfirmation `json:"email_registration_confirmation"`
	// The time in ISO 8601 format when the follow-up email was sent.
	FollowUpSendOn string `json:"follow_up_send_on"`
	FollowUpSender NullableWebinarEmailSettingsFollowUpSender `json:"follow_up_sender"`
	// The name of the sender for emails that are sent about the webinar.
	From string `json:"from"`
	// The URI of the logo image to include in emails that are sent about the webinar.
	LogoUri NullableString `json:"logo_uri"`
	Pictures NullableWebinarEmailSettingsPictures `json:"pictures"`
	// The sender's reply email address.
	ReplyEmail NullableString `json:"reply_email"`
	// The sender's physical address.
	SenderAddress NullableString `json:"sender_address"`
	// The URL of the sender's privacy policy.
	SenderPolicyUrl NullableString `json:"sender_policy_url"`
	// Whether to include a custom link in emails that are sent about the webinar.
	UseCustomLink bool `json:"use_custom_link"`
	// Whether to include a reply link in the footer of emails that are sent about the webinar.
	UseReplyEmail bool `json:"use_reply_email"`
	// Whether to include the sender's physical address in the footer of emails that are sent about the webinar.
	UseSenderAddress bool `json:"use_sender_address"`
	// Whether to include the URL of the sender's privacy policy in the footer of emails that are sent about the webinar.
	UseSenderPolicyUrl bool `json:"use_sender_policy_url"`
}

// NewWebinarEmailSettings instantiates a new WebinarEmailSettings object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewWebinarEmailSettings(accentColor string, customLink NullableString, emailEventReminder24Hrs WebinarEmailSettingsEmailEventReminder24Hrs, emailPostEventThankYou WebinarEmailSettingsEmailPostEventThankYou, emailPreferences []string, emailRegistrationConfirmation WebinarEmailSettingsEmailRegistrationConfirmation, followUpSendOn string, followUpSender NullableWebinarEmailSettingsFollowUpSender, from string, logoUri NullableString, pictures NullableWebinarEmailSettingsPictures, replyEmail NullableString, senderAddress NullableString, senderPolicyUrl NullableString, useCustomLink bool, useReplyEmail bool, useSenderAddress bool, useSenderPolicyUrl bool) *WebinarEmailSettings {
	this := WebinarEmailSettings{}
	this.AccentColor = accentColor
	this.CustomLink = customLink
	this.EmailEventReminder24Hrs = emailEventReminder24Hrs
	this.EmailPostEventThankYou = emailPostEventThankYou
	this.EmailPreferences = emailPreferences
	this.EmailRegistrationConfirmation = emailRegistrationConfirmation
	this.FollowUpSendOn = followUpSendOn
	this.FollowUpSender = followUpSender
	this.From = from
	this.LogoUri = logoUri
	this.Pictures = pictures
	this.ReplyEmail = replyEmail
	this.SenderAddress = senderAddress
	this.SenderPolicyUrl = senderPolicyUrl
	this.UseCustomLink = useCustomLink
	this.UseReplyEmail = useReplyEmail
	this.UseSenderAddress = useSenderAddress
	this.UseSenderPolicyUrl = useSenderPolicyUrl
	return &this
}

// NewWebinarEmailSettingsWithDefaults instantiates a new WebinarEmailSettings object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewWebinarEmailSettingsWithDefaults() *WebinarEmailSettings {
	this := WebinarEmailSettings{}
	return &this
}

// GetAccentColor returns the AccentColor field value
func (o *WebinarEmailSettings) GetAccentColor() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.AccentColor
}

// GetAccentColorOk returns a tuple with the AccentColor field value
// and a boolean to check if the value has been set.
func (o *WebinarEmailSettings) GetAccentColorOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.AccentColor, true
}

// SetAccentColor sets field value
func (o *WebinarEmailSettings) SetAccentColor(v string) {
	o.AccentColor = v
}

// GetCustomLink returns the CustomLink field value
// If the value is explicit nil, the zero value for string will be returned
func (o *WebinarEmailSettings) GetCustomLink() string {
	if o == nil || o.CustomLink.Get() == nil {
		var ret string
		return ret
	}

	return *o.CustomLink.Get()
}

// GetCustomLinkOk returns a tuple with the CustomLink field value
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *WebinarEmailSettings) GetCustomLinkOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.CustomLink.Get(), o.CustomLink.IsSet()
}

// SetCustomLink sets field value
func (o *WebinarEmailSettings) SetCustomLink(v string) {
	o.CustomLink.Set(&v)
}

// GetEmailEventReminder24Hrs returns the EmailEventReminder24Hrs field value
func (o *WebinarEmailSettings) GetEmailEventReminder24Hrs() WebinarEmailSettingsEmailEventReminder24Hrs {
	if o == nil {
		var ret WebinarEmailSettingsEmailEventReminder24Hrs
		return ret
	}

	return o.EmailEventReminder24Hrs
}

// GetEmailEventReminder24HrsOk returns a tuple with the EmailEventReminder24Hrs field value
// and a boolean to check if the value has been set.
func (o *WebinarEmailSettings) GetEmailEventReminder24HrsOk() (*WebinarEmailSettingsEmailEventReminder24Hrs, bool) {
	if o == nil {
		return nil, false
	}
	return &o.EmailEventReminder24Hrs, true
}

// SetEmailEventReminder24Hrs sets field value
func (o *WebinarEmailSettings) SetEmailEventReminder24Hrs(v WebinarEmailSettingsEmailEventReminder24Hrs) {
	o.EmailEventReminder24Hrs = v
}

// GetEmailPostEventThankYou returns the EmailPostEventThankYou field value
func (o *WebinarEmailSettings) GetEmailPostEventThankYou() WebinarEmailSettingsEmailPostEventThankYou {
	if o == nil {
		var ret WebinarEmailSettingsEmailPostEventThankYou
		return ret
	}

	return o.EmailPostEventThankYou
}

// GetEmailPostEventThankYouOk returns a tuple with the EmailPostEventThankYou field value
// and a boolean to check if the value has been set.
func (o *WebinarEmailSettings) GetEmailPostEventThankYouOk() (*WebinarEmailSettingsEmailPostEventThankYou, bool) {
	if o == nil {
		return nil, false
	}
	return &o.EmailPostEventThankYou, true
}

// SetEmailPostEventThankYou sets field value
func (o *WebinarEmailSettings) SetEmailPostEventThankYou(v WebinarEmailSettingsEmailPostEventThankYou) {
	o.EmailPostEventThankYou = v
}

// GetEmailPreferences returns the EmailPreferences field value
func (o *WebinarEmailSettings) GetEmailPreferences() []string {
	if o == nil {
		var ret []string
		return ret
	}

	return o.EmailPreferences
}

// GetEmailPreferencesOk returns a tuple with the EmailPreferences field value
// and a boolean to check if the value has been set.
func (o *WebinarEmailSettings) GetEmailPreferencesOk() ([]string, bool) {
	if o == nil {
		return nil, false
	}
	return o.EmailPreferences, true
}

// SetEmailPreferences sets field value
func (o *WebinarEmailSettings) SetEmailPreferences(v []string) {
	o.EmailPreferences = v
}

// GetEmailRegistrationConfirmation returns the EmailRegistrationConfirmation field value
func (o *WebinarEmailSettings) GetEmailRegistrationConfirmation() WebinarEmailSettingsEmailRegistrationConfirmation {
	if o == nil {
		var ret WebinarEmailSettingsEmailRegistrationConfirmation
		return ret
	}

	return o.EmailRegistrationConfirmation
}

// GetEmailRegistrationConfirmationOk returns a tuple with the EmailRegistrationConfirmation field value
// and a boolean to check if the value has been set.
func (o *WebinarEmailSettings) GetEmailRegistrationConfirmationOk() (*WebinarEmailSettingsEmailRegistrationConfirmation, bool) {
	if o == nil {
		return nil, false
	}
	return &o.EmailRegistrationConfirmation, true
}

// SetEmailRegistrationConfirmation sets field value
func (o *WebinarEmailSettings) SetEmailRegistrationConfirmation(v WebinarEmailSettingsEmailRegistrationConfirmation) {
	o.EmailRegistrationConfirmation = v
}

// GetFollowUpSendOn returns the FollowUpSendOn field value
func (o *WebinarEmailSettings) GetFollowUpSendOn() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.FollowUpSendOn
}

// GetFollowUpSendOnOk returns a tuple with the FollowUpSendOn field value
// and a boolean to check if the value has been set.
func (o *WebinarEmailSettings) GetFollowUpSendOnOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.FollowUpSendOn, true
}

// SetFollowUpSendOn sets field value
func (o *WebinarEmailSettings) SetFollowUpSendOn(v string) {
	o.FollowUpSendOn = v
}

// GetFollowUpSender returns the FollowUpSender field value
// If the value is explicit nil, the zero value for WebinarEmailSettingsFollowUpSender will be returned
func (o *WebinarEmailSettings) GetFollowUpSender() WebinarEmailSettingsFollowUpSender {
	if o == nil || o.FollowUpSender.Get() == nil {
		var ret WebinarEmailSettingsFollowUpSender
		return ret
	}

	return *o.FollowUpSender.Get()
}

// GetFollowUpSenderOk returns a tuple with the FollowUpSender field value
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *WebinarEmailSettings) GetFollowUpSenderOk() (*WebinarEmailSettingsFollowUpSender, bool) {
	if o == nil {
		return nil, false
	}
	return o.FollowUpSender.Get(), o.FollowUpSender.IsSet()
}

// SetFollowUpSender sets field value
func (o *WebinarEmailSettings) SetFollowUpSender(v WebinarEmailSettingsFollowUpSender) {
	o.FollowUpSender.Set(&v)
}

// GetFrom returns the From field value
func (o *WebinarEmailSettings) GetFrom() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.From
}

// GetFromOk returns a tuple with the From field value
// and a boolean to check if the value has been set.
func (o *WebinarEmailSettings) GetFromOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.From, true
}

// SetFrom sets field value
func (o *WebinarEmailSettings) SetFrom(v string) {
	o.From = v
}

// GetLogoUri returns the LogoUri field value
// If the value is explicit nil, the zero value for string will be returned
func (o *WebinarEmailSettings) GetLogoUri() string {
	if o == nil || o.LogoUri.Get() == nil {
		var ret string
		return ret
	}

	return *o.LogoUri.Get()
}

// GetLogoUriOk returns a tuple with the LogoUri field value
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *WebinarEmailSettings) GetLogoUriOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.LogoUri.Get(), o.LogoUri.IsSet()
}

// SetLogoUri sets field value
func (o *WebinarEmailSettings) SetLogoUri(v string) {
	o.LogoUri.Set(&v)
}

// GetPictures returns the Pictures field value
// If the value is explicit nil, the zero value for WebinarEmailSettingsPictures will be returned
func (o *WebinarEmailSettings) GetPictures() WebinarEmailSettingsPictures {
	if o == nil || o.Pictures.Get() == nil {
		var ret WebinarEmailSettingsPictures
		return ret
	}

	return *o.Pictures.Get()
}

// GetPicturesOk returns a tuple with the Pictures field value
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *WebinarEmailSettings) GetPicturesOk() (*WebinarEmailSettingsPictures, bool) {
	if o == nil {
		return nil, false
	}
	return o.Pictures.Get(), o.Pictures.IsSet()
}

// SetPictures sets field value
func (o *WebinarEmailSettings) SetPictures(v WebinarEmailSettingsPictures) {
	o.Pictures.Set(&v)
}

// GetReplyEmail returns the ReplyEmail field value
// If the value is explicit nil, the zero value for string will be returned
func (o *WebinarEmailSettings) GetReplyEmail() string {
	if o == nil || o.ReplyEmail.Get() == nil {
		var ret string
		return ret
	}

	return *o.ReplyEmail.Get()
}

// GetReplyEmailOk returns a tuple with the ReplyEmail field value
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *WebinarEmailSettings) GetReplyEmailOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.ReplyEmail.Get(), o.ReplyEmail.IsSet()
}

// SetReplyEmail sets field value
func (o *WebinarEmailSettings) SetReplyEmail(v string) {
	o.ReplyEmail.Set(&v)
}

// GetSenderAddress returns the SenderAddress field value
// If the value is explicit nil, the zero value for string will be returned
func (o *WebinarEmailSettings) GetSenderAddress() string {
	if o == nil || o.SenderAddress.Get() == nil {
		var ret string
		return ret
	}

	return *o.SenderAddress.Get()
}

// GetSenderAddressOk returns a tuple with the SenderAddress field value
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *WebinarEmailSettings) GetSenderAddressOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.SenderAddress.Get(), o.SenderAddress.IsSet()
}

// SetSenderAddress sets field value
func (o *WebinarEmailSettings) SetSenderAddress(v string) {
	o.SenderAddress.Set(&v)
}

// GetSenderPolicyUrl returns the SenderPolicyUrl field value
// If the value is explicit nil, the zero value for string will be returned
func (o *WebinarEmailSettings) GetSenderPolicyUrl() string {
	if o == nil || o.SenderPolicyUrl.Get() == nil {
		var ret string
		return ret
	}

	return *o.SenderPolicyUrl.Get()
}

// GetSenderPolicyUrlOk returns a tuple with the SenderPolicyUrl field value
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *WebinarEmailSettings) GetSenderPolicyUrlOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.SenderPolicyUrl.Get(), o.SenderPolicyUrl.IsSet()
}

// SetSenderPolicyUrl sets field value
func (o *WebinarEmailSettings) SetSenderPolicyUrl(v string) {
	o.SenderPolicyUrl.Set(&v)
}

// GetUseCustomLink returns the UseCustomLink field value
func (o *WebinarEmailSettings) GetUseCustomLink() bool {
	if o == nil {
		var ret bool
		return ret
	}

	return o.UseCustomLink
}

// GetUseCustomLinkOk returns a tuple with the UseCustomLink field value
// and a boolean to check if the value has been set.
func (o *WebinarEmailSettings) GetUseCustomLinkOk() (*bool, bool) {
	if o == nil {
		return nil, false
	}
	return &o.UseCustomLink, true
}

// SetUseCustomLink sets field value
func (o *WebinarEmailSettings) SetUseCustomLink(v bool) {
	o.UseCustomLink = v
}

// GetUseReplyEmail returns the UseReplyEmail field value
func (o *WebinarEmailSettings) GetUseReplyEmail() bool {
	if o == nil {
		var ret bool
		return ret
	}

	return o.UseReplyEmail
}

// GetUseReplyEmailOk returns a tuple with the UseReplyEmail field value
// and a boolean to check if the value has been set.
func (o *WebinarEmailSettings) GetUseReplyEmailOk() (*bool, bool) {
	if o == nil {
		return nil, false
	}
	return &o.UseReplyEmail, true
}

// SetUseReplyEmail sets field value
func (o *WebinarEmailSettings) SetUseReplyEmail(v bool) {
	o.UseReplyEmail = v
}

// GetUseSenderAddress returns the UseSenderAddress field value
func (o *WebinarEmailSettings) GetUseSenderAddress() bool {
	if o == nil {
		var ret bool
		return ret
	}

	return o.UseSenderAddress
}

// GetUseSenderAddressOk returns a tuple with the UseSenderAddress field value
// and a boolean to check if the value has been set.
func (o *WebinarEmailSettings) GetUseSenderAddressOk() (*bool, bool) {
	if o == nil {
		return nil, false
	}
	return &o.UseSenderAddress, true
}

// SetUseSenderAddress sets field value
func (o *WebinarEmailSettings) SetUseSenderAddress(v bool) {
	o.UseSenderAddress = v
}

// GetUseSenderPolicyUrl returns the UseSenderPolicyUrl field value
func (o *WebinarEmailSettings) GetUseSenderPolicyUrl() bool {
	if o == nil {
		var ret bool
		return ret
	}

	return o.UseSenderPolicyUrl
}

// GetUseSenderPolicyUrlOk returns a tuple with the UseSenderPolicyUrl field value
// and a boolean to check if the value has been set.
func (o *WebinarEmailSettings) GetUseSenderPolicyUrlOk() (*bool, bool) {
	if o == nil {
		return nil, false
	}
	return &o.UseSenderPolicyUrl, true
}

// SetUseSenderPolicyUrl sets field value
func (o *WebinarEmailSettings) SetUseSenderPolicyUrl(v bool) {
	o.UseSenderPolicyUrl = v
}

func (o WebinarEmailSettings) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o WebinarEmailSettings) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["accent_color"] = o.AccentColor
	toSerialize["custom_link"] = o.CustomLink.Get()
	toSerialize["email_event_reminder_24_hrs"] = o.EmailEventReminder24Hrs
	toSerialize["email_post_event_thank_you"] = o.EmailPostEventThankYou
	toSerialize["email_preferences"] = o.EmailPreferences
	toSerialize["email_registration_confirmation"] = o.EmailRegistrationConfirmation
	toSerialize["follow_up_send_on"] = o.FollowUpSendOn
	toSerialize["follow_up_sender"] = o.FollowUpSender.Get()
	toSerialize["from"] = o.From
	toSerialize["logo_uri"] = o.LogoUri.Get()
	toSerialize["pictures"] = o.Pictures.Get()
	toSerialize["reply_email"] = o.ReplyEmail.Get()
	toSerialize["sender_address"] = o.SenderAddress.Get()
	toSerialize["sender_policy_url"] = o.SenderPolicyUrl.Get()
	toSerialize["use_custom_link"] = o.UseCustomLink
	toSerialize["use_reply_email"] = o.UseReplyEmail
	toSerialize["use_sender_address"] = o.UseSenderAddress
	toSerialize["use_sender_policy_url"] = o.UseSenderPolicyUrl
	return toSerialize, nil
}

type NullableWebinarEmailSettings struct {
	value *WebinarEmailSettings
	isSet bool
}

func (v NullableWebinarEmailSettings) Get() *WebinarEmailSettings {
	return v.value
}

func (v *NullableWebinarEmailSettings) Set(val *WebinarEmailSettings) {
	v.value = val
	v.isSet = true
}

func (v NullableWebinarEmailSettings) IsSet() bool {
	return v.isSet
}

func (v *NullableWebinarEmailSettings) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableWebinarEmailSettings(val *WebinarEmailSettings) *NullableWebinarEmailSettings {
	return &NullableWebinarEmailSettings{value: val, isSet: true}
}

func (v NullableWebinarEmailSettings) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableWebinarEmailSettings) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


