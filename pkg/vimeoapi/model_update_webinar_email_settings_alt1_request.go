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

// checks if the UpdateWebinarEmailSettingsAlt1Request type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &UpdateWebinarEmailSettingsAlt1Request{}

// UpdateWebinarEmailSettingsAlt1Request struct for UpdateWebinarEmailSettingsAlt1Request
type UpdateWebinarEmailSettingsAlt1Request struct {
	// The accent color scheme for emails that are sent about the webinar.
	AccentColor *string `json:"accent_color,omitempty"`
	// The custom link for emails that are sent about the webinar.
	CustomLink *string `json:"custom_link,omitempty"`
	// The email customization details for the webinar reminder email, which goes out 24 hours before the event.
	EmailEventReminder24Hrs map[string]interface{} `json:"email_event_reminder_24_hrs,omitempty"`
	// The email customization details for the webinar post-event thank-you email.
	EmailPostEventThankYou map[string]interface{} `json:"email_post_event_thank_you,omitempty"`
	EmailPreferences *UpdateWebinarEmailSettingsAlt1RequestEmailPreferences `json:"email_preferences,omitempty"`
	// The email customization details for the webinar registration confirmation email.
	EmailRegistrationConfirmation map[string]interface{} `json:"email_registration_confirmation,omitempty"`
	// The name of the sender for emails that are sent about the webinar.
	From *string `json:"from,omitempty"`
	// The URI of the logo image to include in emails that are sent about the webinar.
	LogoUri *string `json:"logo_uri,omitempty"`
	// The sender's reply email address.
	ReplyEmail *string `json:"reply_email,omitempty"`
	// The sender's physical address.
	SenderAddress *string `json:"sender_address,omitempty"`
	// The URL of the sender's privacy policy.
	SenderPolicyUrl *string `json:"sender_policy_url,omitempty"`
	// Whether to include a custom link in emails that are sent about the webinar.
	UseCustomLink *bool `json:"use_custom_link,omitempty"`
	// Whether to include a reply link in the footer of emails that are sent about the webinar.
	UseReplyEmail *bool `json:"use_reply_email,omitempty"`
	// Whether to include the sender's physical address in the footer of emails that are sent about the webinar.
	UseSenderAddress *bool `json:"use_sender_address,omitempty"`
	// Whether to include the URL of the sender's privacy policy in the footer of emails that are sent about the webinar.
	UseSenderPolicyUrl *bool `json:"use_sender_policy_url,omitempty"`
}

// NewUpdateWebinarEmailSettingsAlt1Request instantiates a new UpdateWebinarEmailSettingsAlt1Request object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewUpdateWebinarEmailSettingsAlt1Request() *UpdateWebinarEmailSettingsAlt1Request {
	this := UpdateWebinarEmailSettingsAlt1Request{}
	return &this
}

// NewUpdateWebinarEmailSettingsAlt1RequestWithDefaults instantiates a new UpdateWebinarEmailSettingsAlt1Request object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewUpdateWebinarEmailSettingsAlt1RequestWithDefaults() *UpdateWebinarEmailSettingsAlt1Request {
	this := UpdateWebinarEmailSettingsAlt1Request{}
	return &this
}

// GetAccentColor returns the AccentColor field value if set, zero value otherwise.
func (o *UpdateWebinarEmailSettingsAlt1Request) GetAccentColor() string {
	if o == nil || IsNil(o.AccentColor) {
		var ret string
		return ret
	}
	return *o.AccentColor
}

// GetAccentColorOk returns a tuple with the AccentColor field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UpdateWebinarEmailSettingsAlt1Request) GetAccentColorOk() (*string, bool) {
	if o == nil || IsNil(o.AccentColor) {
		return nil, false
	}
	return o.AccentColor, true
}

// HasAccentColor returns a boolean if a field has been set.
func (o *UpdateWebinarEmailSettingsAlt1Request) HasAccentColor() bool {
	if o != nil && !IsNil(o.AccentColor) {
		return true
	}

	return false
}

// SetAccentColor gets a reference to the given string and assigns it to the AccentColor field.
func (o *UpdateWebinarEmailSettingsAlt1Request) SetAccentColor(v string) {
	o.AccentColor = &v
}

// GetCustomLink returns the CustomLink field value if set, zero value otherwise.
func (o *UpdateWebinarEmailSettingsAlt1Request) GetCustomLink() string {
	if o == nil || IsNil(o.CustomLink) {
		var ret string
		return ret
	}
	return *o.CustomLink
}

// GetCustomLinkOk returns a tuple with the CustomLink field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UpdateWebinarEmailSettingsAlt1Request) GetCustomLinkOk() (*string, bool) {
	if o == nil || IsNil(o.CustomLink) {
		return nil, false
	}
	return o.CustomLink, true
}

// HasCustomLink returns a boolean if a field has been set.
func (o *UpdateWebinarEmailSettingsAlt1Request) HasCustomLink() bool {
	if o != nil && !IsNil(o.CustomLink) {
		return true
	}

	return false
}

// SetCustomLink gets a reference to the given string and assigns it to the CustomLink field.
func (o *UpdateWebinarEmailSettingsAlt1Request) SetCustomLink(v string) {
	o.CustomLink = &v
}

// GetEmailEventReminder24Hrs returns the EmailEventReminder24Hrs field value if set, zero value otherwise.
func (o *UpdateWebinarEmailSettingsAlt1Request) GetEmailEventReminder24Hrs() map[string]interface{} {
	if o == nil || IsNil(o.EmailEventReminder24Hrs) {
		var ret map[string]interface{}
		return ret
	}
	return o.EmailEventReminder24Hrs
}

// GetEmailEventReminder24HrsOk returns a tuple with the EmailEventReminder24Hrs field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UpdateWebinarEmailSettingsAlt1Request) GetEmailEventReminder24HrsOk() (map[string]interface{}, bool) {
	if o == nil || IsNil(o.EmailEventReminder24Hrs) {
		return map[string]interface{}{}, false
	}
	return o.EmailEventReminder24Hrs, true
}

// HasEmailEventReminder24Hrs returns a boolean if a field has been set.
func (o *UpdateWebinarEmailSettingsAlt1Request) HasEmailEventReminder24Hrs() bool {
	if o != nil && !IsNil(o.EmailEventReminder24Hrs) {
		return true
	}

	return false
}

// SetEmailEventReminder24Hrs gets a reference to the given map[string]interface{} and assigns it to the EmailEventReminder24Hrs field.
func (o *UpdateWebinarEmailSettingsAlt1Request) SetEmailEventReminder24Hrs(v map[string]interface{}) {
	o.EmailEventReminder24Hrs = v
}

// GetEmailPostEventThankYou returns the EmailPostEventThankYou field value if set, zero value otherwise.
func (o *UpdateWebinarEmailSettingsAlt1Request) GetEmailPostEventThankYou() map[string]interface{} {
	if o == nil || IsNil(o.EmailPostEventThankYou) {
		var ret map[string]interface{}
		return ret
	}
	return o.EmailPostEventThankYou
}

// GetEmailPostEventThankYouOk returns a tuple with the EmailPostEventThankYou field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UpdateWebinarEmailSettingsAlt1Request) GetEmailPostEventThankYouOk() (map[string]interface{}, bool) {
	if o == nil || IsNil(o.EmailPostEventThankYou) {
		return map[string]interface{}{}, false
	}
	return o.EmailPostEventThankYou, true
}

// HasEmailPostEventThankYou returns a boolean if a field has been set.
func (o *UpdateWebinarEmailSettingsAlt1Request) HasEmailPostEventThankYou() bool {
	if o != nil && !IsNil(o.EmailPostEventThankYou) {
		return true
	}

	return false
}

// SetEmailPostEventThankYou gets a reference to the given map[string]interface{} and assigns it to the EmailPostEventThankYou field.
func (o *UpdateWebinarEmailSettingsAlt1Request) SetEmailPostEventThankYou(v map[string]interface{}) {
	o.EmailPostEventThankYou = v
}

// GetEmailPreferences returns the EmailPreferences field value if set, zero value otherwise.
func (o *UpdateWebinarEmailSettingsAlt1Request) GetEmailPreferences() UpdateWebinarEmailSettingsAlt1RequestEmailPreferences {
	if o == nil || IsNil(o.EmailPreferences) {
		var ret UpdateWebinarEmailSettingsAlt1RequestEmailPreferences
		return ret
	}
	return *o.EmailPreferences
}

// GetEmailPreferencesOk returns a tuple with the EmailPreferences field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UpdateWebinarEmailSettingsAlt1Request) GetEmailPreferencesOk() (*UpdateWebinarEmailSettingsAlt1RequestEmailPreferences, bool) {
	if o == nil || IsNil(o.EmailPreferences) {
		return nil, false
	}
	return o.EmailPreferences, true
}

// HasEmailPreferences returns a boolean if a field has been set.
func (o *UpdateWebinarEmailSettingsAlt1Request) HasEmailPreferences() bool {
	if o != nil && !IsNil(o.EmailPreferences) {
		return true
	}

	return false
}

// SetEmailPreferences gets a reference to the given UpdateWebinarEmailSettingsAlt1RequestEmailPreferences and assigns it to the EmailPreferences field.
func (o *UpdateWebinarEmailSettingsAlt1Request) SetEmailPreferences(v UpdateWebinarEmailSettingsAlt1RequestEmailPreferences) {
	o.EmailPreferences = &v
}

// GetEmailRegistrationConfirmation returns the EmailRegistrationConfirmation field value if set, zero value otherwise.
func (o *UpdateWebinarEmailSettingsAlt1Request) GetEmailRegistrationConfirmation() map[string]interface{} {
	if o == nil || IsNil(o.EmailRegistrationConfirmation) {
		var ret map[string]interface{}
		return ret
	}
	return o.EmailRegistrationConfirmation
}

// GetEmailRegistrationConfirmationOk returns a tuple with the EmailRegistrationConfirmation field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UpdateWebinarEmailSettingsAlt1Request) GetEmailRegistrationConfirmationOk() (map[string]interface{}, bool) {
	if o == nil || IsNil(o.EmailRegistrationConfirmation) {
		return map[string]interface{}{}, false
	}
	return o.EmailRegistrationConfirmation, true
}

// HasEmailRegistrationConfirmation returns a boolean if a field has been set.
func (o *UpdateWebinarEmailSettingsAlt1Request) HasEmailRegistrationConfirmation() bool {
	if o != nil && !IsNil(o.EmailRegistrationConfirmation) {
		return true
	}

	return false
}

// SetEmailRegistrationConfirmation gets a reference to the given map[string]interface{} and assigns it to the EmailRegistrationConfirmation field.
func (o *UpdateWebinarEmailSettingsAlt1Request) SetEmailRegistrationConfirmation(v map[string]interface{}) {
	o.EmailRegistrationConfirmation = v
}

// GetFrom returns the From field value if set, zero value otherwise.
func (o *UpdateWebinarEmailSettingsAlt1Request) GetFrom() string {
	if o == nil || IsNil(o.From) {
		var ret string
		return ret
	}
	return *o.From
}

// GetFromOk returns a tuple with the From field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UpdateWebinarEmailSettingsAlt1Request) GetFromOk() (*string, bool) {
	if o == nil || IsNil(o.From) {
		return nil, false
	}
	return o.From, true
}

// HasFrom returns a boolean if a field has been set.
func (o *UpdateWebinarEmailSettingsAlt1Request) HasFrom() bool {
	if o != nil && !IsNil(o.From) {
		return true
	}

	return false
}

// SetFrom gets a reference to the given string and assigns it to the From field.
func (o *UpdateWebinarEmailSettingsAlt1Request) SetFrom(v string) {
	o.From = &v
}

// GetLogoUri returns the LogoUri field value if set, zero value otherwise.
func (o *UpdateWebinarEmailSettingsAlt1Request) GetLogoUri() string {
	if o == nil || IsNil(o.LogoUri) {
		var ret string
		return ret
	}
	return *o.LogoUri
}

// GetLogoUriOk returns a tuple with the LogoUri field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UpdateWebinarEmailSettingsAlt1Request) GetLogoUriOk() (*string, bool) {
	if o == nil || IsNil(o.LogoUri) {
		return nil, false
	}
	return o.LogoUri, true
}

// HasLogoUri returns a boolean if a field has been set.
func (o *UpdateWebinarEmailSettingsAlt1Request) HasLogoUri() bool {
	if o != nil && !IsNil(o.LogoUri) {
		return true
	}

	return false
}

// SetLogoUri gets a reference to the given string and assigns it to the LogoUri field.
func (o *UpdateWebinarEmailSettingsAlt1Request) SetLogoUri(v string) {
	o.LogoUri = &v
}

// GetReplyEmail returns the ReplyEmail field value if set, zero value otherwise.
func (o *UpdateWebinarEmailSettingsAlt1Request) GetReplyEmail() string {
	if o == nil || IsNil(o.ReplyEmail) {
		var ret string
		return ret
	}
	return *o.ReplyEmail
}

// GetReplyEmailOk returns a tuple with the ReplyEmail field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UpdateWebinarEmailSettingsAlt1Request) GetReplyEmailOk() (*string, bool) {
	if o == nil || IsNil(o.ReplyEmail) {
		return nil, false
	}
	return o.ReplyEmail, true
}

// HasReplyEmail returns a boolean if a field has been set.
func (o *UpdateWebinarEmailSettingsAlt1Request) HasReplyEmail() bool {
	if o != nil && !IsNil(o.ReplyEmail) {
		return true
	}

	return false
}

// SetReplyEmail gets a reference to the given string and assigns it to the ReplyEmail field.
func (o *UpdateWebinarEmailSettingsAlt1Request) SetReplyEmail(v string) {
	o.ReplyEmail = &v
}

// GetSenderAddress returns the SenderAddress field value if set, zero value otherwise.
func (o *UpdateWebinarEmailSettingsAlt1Request) GetSenderAddress() string {
	if o == nil || IsNil(o.SenderAddress) {
		var ret string
		return ret
	}
	return *o.SenderAddress
}

// GetSenderAddressOk returns a tuple with the SenderAddress field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UpdateWebinarEmailSettingsAlt1Request) GetSenderAddressOk() (*string, bool) {
	if o == nil || IsNil(o.SenderAddress) {
		return nil, false
	}
	return o.SenderAddress, true
}

// HasSenderAddress returns a boolean if a field has been set.
func (o *UpdateWebinarEmailSettingsAlt1Request) HasSenderAddress() bool {
	if o != nil && !IsNil(o.SenderAddress) {
		return true
	}

	return false
}

// SetSenderAddress gets a reference to the given string and assigns it to the SenderAddress field.
func (o *UpdateWebinarEmailSettingsAlt1Request) SetSenderAddress(v string) {
	o.SenderAddress = &v
}

// GetSenderPolicyUrl returns the SenderPolicyUrl field value if set, zero value otherwise.
func (o *UpdateWebinarEmailSettingsAlt1Request) GetSenderPolicyUrl() string {
	if o == nil || IsNil(o.SenderPolicyUrl) {
		var ret string
		return ret
	}
	return *o.SenderPolicyUrl
}

// GetSenderPolicyUrlOk returns a tuple with the SenderPolicyUrl field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UpdateWebinarEmailSettingsAlt1Request) GetSenderPolicyUrlOk() (*string, bool) {
	if o == nil || IsNil(o.SenderPolicyUrl) {
		return nil, false
	}
	return o.SenderPolicyUrl, true
}

// HasSenderPolicyUrl returns a boolean if a field has been set.
func (o *UpdateWebinarEmailSettingsAlt1Request) HasSenderPolicyUrl() bool {
	if o != nil && !IsNil(o.SenderPolicyUrl) {
		return true
	}

	return false
}

// SetSenderPolicyUrl gets a reference to the given string and assigns it to the SenderPolicyUrl field.
func (o *UpdateWebinarEmailSettingsAlt1Request) SetSenderPolicyUrl(v string) {
	o.SenderPolicyUrl = &v
}

// GetUseCustomLink returns the UseCustomLink field value if set, zero value otherwise.
func (o *UpdateWebinarEmailSettingsAlt1Request) GetUseCustomLink() bool {
	if o == nil || IsNil(o.UseCustomLink) {
		var ret bool
		return ret
	}
	return *o.UseCustomLink
}

// GetUseCustomLinkOk returns a tuple with the UseCustomLink field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UpdateWebinarEmailSettingsAlt1Request) GetUseCustomLinkOk() (*bool, bool) {
	if o == nil || IsNil(o.UseCustomLink) {
		return nil, false
	}
	return o.UseCustomLink, true
}

// HasUseCustomLink returns a boolean if a field has been set.
func (o *UpdateWebinarEmailSettingsAlt1Request) HasUseCustomLink() bool {
	if o != nil && !IsNil(o.UseCustomLink) {
		return true
	}

	return false
}

// SetUseCustomLink gets a reference to the given bool and assigns it to the UseCustomLink field.
func (o *UpdateWebinarEmailSettingsAlt1Request) SetUseCustomLink(v bool) {
	o.UseCustomLink = &v
}

// GetUseReplyEmail returns the UseReplyEmail field value if set, zero value otherwise.
func (o *UpdateWebinarEmailSettingsAlt1Request) GetUseReplyEmail() bool {
	if o == nil || IsNil(o.UseReplyEmail) {
		var ret bool
		return ret
	}
	return *o.UseReplyEmail
}

// GetUseReplyEmailOk returns a tuple with the UseReplyEmail field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UpdateWebinarEmailSettingsAlt1Request) GetUseReplyEmailOk() (*bool, bool) {
	if o == nil || IsNil(o.UseReplyEmail) {
		return nil, false
	}
	return o.UseReplyEmail, true
}

// HasUseReplyEmail returns a boolean if a field has been set.
func (o *UpdateWebinarEmailSettingsAlt1Request) HasUseReplyEmail() bool {
	if o != nil && !IsNil(o.UseReplyEmail) {
		return true
	}

	return false
}

// SetUseReplyEmail gets a reference to the given bool and assigns it to the UseReplyEmail field.
func (o *UpdateWebinarEmailSettingsAlt1Request) SetUseReplyEmail(v bool) {
	o.UseReplyEmail = &v
}

// GetUseSenderAddress returns the UseSenderAddress field value if set, zero value otherwise.
func (o *UpdateWebinarEmailSettingsAlt1Request) GetUseSenderAddress() bool {
	if o == nil || IsNil(o.UseSenderAddress) {
		var ret bool
		return ret
	}
	return *o.UseSenderAddress
}

// GetUseSenderAddressOk returns a tuple with the UseSenderAddress field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UpdateWebinarEmailSettingsAlt1Request) GetUseSenderAddressOk() (*bool, bool) {
	if o == nil || IsNil(o.UseSenderAddress) {
		return nil, false
	}
	return o.UseSenderAddress, true
}

// HasUseSenderAddress returns a boolean if a field has been set.
func (o *UpdateWebinarEmailSettingsAlt1Request) HasUseSenderAddress() bool {
	if o != nil && !IsNil(o.UseSenderAddress) {
		return true
	}

	return false
}

// SetUseSenderAddress gets a reference to the given bool and assigns it to the UseSenderAddress field.
func (o *UpdateWebinarEmailSettingsAlt1Request) SetUseSenderAddress(v bool) {
	o.UseSenderAddress = &v
}

// GetUseSenderPolicyUrl returns the UseSenderPolicyUrl field value if set, zero value otherwise.
func (o *UpdateWebinarEmailSettingsAlt1Request) GetUseSenderPolicyUrl() bool {
	if o == nil || IsNil(o.UseSenderPolicyUrl) {
		var ret bool
		return ret
	}
	return *o.UseSenderPolicyUrl
}

// GetUseSenderPolicyUrlOk returns a tuple with the UseSenderPolicyUrl field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UpdateWebinarEmailSettingsAlt1Request) GetUseSenderPolicyUrlOk() (*bool, bool) {
	if o == nil || IsNil(o.UseSenderPolicyUrl) {
		return nil, false
	}
	return o.UseSenderPolicyUrl, true
}

// HasUseSenderPolicyUrl returns a boolean if a field has been set.
func (o *UpdateWebinarEmailSettingsAlt1Request) HasUseSenderPolicyUrl() bool {
	if o != nil && !IsNil(o.UseSenderPolicyUrl) {
		return true
	}

	return false
}

// SetUseSenderPolicyUrl gets a reference to the given bool and assigns it to the UseSenderPolicyUrl field.
func (o *UpdateWebinarEmailSettingsAlt1Request) SetUseSenderPolicyUrl(v bool) {
	o.UseSenderPolicyUrl = &v
}

func (o UpdateWebinarEmailSettingsAlt1Request) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o UpdateWebinarEmailSettingsAlt1Request) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.AccentColor) {
		toSerialize["accent_color"] = o.AccentColor
	}
	if !IsNil(o.CustomLink) {
		toSerialize["custom_link"] = o.CustomLink
	}
	if !IsNil(o.EmailEventReminder24Hrs) {
		toSerialize["email_event_reminder_24_hrs"] = o.EmailEventReminder24Hrs
	}
	if !IsNil(o.EmailPostEventThankYou) {
		toSerialize["email_post_event_thank_you"] = o.EmailPostEventThankYou
	}
	if !IsNil(o.EmailPreferences) {
		toSerialize["email_preferences"] = o.EmailPreferences
	}
	if !IsNil(o.EmailRegistrationConfirmation) {
		toSerialize["email_registration_confirmation"] = o.EmailRegistrationConfirmation
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

type NullableUpdateWebinarEmailSettingsAlt1Request struct {
	value *UpdateWebinarEmailSettingsAlt1Request
	isSet bool
}

func (v NullableUpdateWebinarEmailSettingsAlt1Request) Get() *UpdateWebinarEmailSettingsAlt1Request {
	return v.value
}

func (v *NullableUpdateWebinarEmailSettingsAlt1Request) Set(val *UpdateWebinarEmailSettingsAlt1Request) {
	v.value = val
	v.isSet = true
}

func (v NullableUpdateWebinarEmailSettingsAlt1Request) IsSet() bool {
	return v.isSet
}

func (v *NullableUpdateWebinarEmailSettingsAlt1Request) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableUpdateWebinarEmailSettingsAlt1Request(val *UpdateWebinarEmailSettingsAlt1Request) *NullableUpdateWebinarEmailSettingsAlt1Request {
	return &NullableUpdateWebinarEmailSettingsAlt1Request{value: val, isSet: true}
}

func (v NullableUpdateWebinarEmailSettingsAlt1Request) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableUpdateWebinarEmailSettingsAlt1Request) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


