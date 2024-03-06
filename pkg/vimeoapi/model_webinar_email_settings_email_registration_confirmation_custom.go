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

// checks if the WebinarEmailSettingsEmailRegistrationConfirmationCustom type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &WebinarEmailSettingsEmailRegistrationConfirmationCustom{}

// WebinarEmailSettingsEmailRegistrationConfirmationCustom The email custom details for the webinar registration confirmation email.
type WebinarEmailSettingsEmailRegistrationConfirmationCustom struct {
	// The HTML body of the email.
	Body string `json:"body"`
	// The target link for the call-to-action button in the email.
	ButtonLink string `json:"button_link"`
	// The text for the call-to-action button in the email.
	ButtonText string `json:"button_text"`
	// The HTML header section of the email.
	Header string `json:"header"`
	// The time in ISO 8601 format when the webinar email content was updated.
	ModifiedTime string `json:"modified_time"`
	// The HTML subject of the email.
	Subject string `json:"subject"`
	// The email type for which the content was customized.  Option descriptions:  * `email_event_reminder_24_hrs` - The webinar reminder email, which goes out 24 hours before the event.  * `email_post_event_thank_you` - The webinar post-event thank-you email.  * `email_registration_confirmation` - The webinar registration confirmation email.
	Type string `json:"type"`
	// Whether to show the calendar in the email.
	UseCalender bool `json:"use_calender"`
	// Whether to include a custom link in emails that are sent about the webinar.
	UseCustomLink bool `json:"use_custom_link"`
}

// NewWebinarEmailSettingsEmailRegistrationConfirmationCustom instantiates a new WebinarEmailSettingsEmailRegistrationConfirmationCustom object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewWebinarEmailSettingsEmailRegistrationConfirmationCustom(body string, buttonLink string, buttonText string, header string, modifiedTime string, subject string, type_ string, useCalender bool, useCustomLink bool) *WebinarEmailSettingsEmailRegistrationConfirmationCustom {
	this := WebinarEmailSettingsEmailRegistrationConfirmationCustom{}
	this.Body = body
	this.ButtonLink = buttonLink
	this.ButtonText = buttonText
	this.Header = header
	this.ModifiedTime = modifiedTime
	this.Subject = subject
	this.Type = type_
	this.UseCalender = useCalender
	this.UseCustomLink = useCustomLink
	return &this
}

// NewWebinarEmailSettingsEmailRegistrationConfirmationCustomWithDefaults instantiates a new WebinarEmailSettingsEmailRegistrationConfirmationCustom object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewWebinarEmailSettingsEmailRegistrationConfirmationCustomWithDefaults() *WebinarEmailSettingsEmailRegistrationConfirmationCustom {
	this := WebinarEmailSettingsEmailRegistrationConfirmationCustom{}
	return &this
}

// GetBody returns the Body field value
func (o *WebinarEmailSettingsEmailRegistrationConfirmationCustom) GetBody() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Body
}

// GetBodyOk returns a tuple with the Body field value
// and a boolean to check if the value has been set.
func (o *WebinarEmailSettingsEmailRegistrationConfirmationCustom) GetBodyOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Body, true
}

// SetBody sets field value
func (o *WebinarEmailSettingsEmailRegistrationConfirmationCustom) SetBody(v string) {
	o.Body = v
}

// GetButtonLink returns the ButtonLink field value
func (o *WebinarEmailSettingsEmailRegistrationConfirmationCustom) GetButtonLink() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.ButtonLink
}

// GetButtonLinkOk returns a tuple with the ButtonLink field value
// and a boolean to check if the value has been set.
func (o *WebinarEmailSettingsEmailRegistrationConfirmationCustom) GetButtonLinkOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ButtonLink, true
}

// SetButtonLink sets field value
func (o *WebinarEmailSettingsEmailRegistrationConfirmationCustom) SetButtonLink(v string) {
	o.ButtonLink = v
}

// GetButtonText returns the ButtonText field value
func (o *WebinarEmailSettingsEmailRegistrationConfirmationCustom) GetButtonText() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.ButtonText
}

// GetButtonTextOk returns a tuple with the ButtonText field value
// and a boolean to check if the value has been set.
func (o *WebinarEmailSettingsEmailRegistrationConfirmationCustom) GetButtonTextOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ButtonText, true
}

// SetButtonText sets field value
func (o *WebinarEmailSettingsEmailRegistrationConfirmationCustom) SetButtonText(v string) {
	o.ButtonText = v
}

// GetHeader returns the Header field value
func (o *WebinarEmailSettingsEmailRegistrationConfirmationCustom) GetHeader() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Header
}

// GetHeaderOk returns a tuple with the Header field value
// and a boolean to check if the value has been set.
func (o *WebinarEmailSettingsEmailRegistrationConfirmationCustom) GetHeaderOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Header, true
}

// SetHeader sets field value
func (o *WebinarEmailSettingsEmailRegistrationConfirmationCustom) SetHeader(v string) {
	o.Header = v
}

// GetModifiedTime returns the ModifiedTime field value
func (o *WebinarEmailSettingsEmailRegistrationConfirmationCustom) GetModifiedTime() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.ModifiedTime
}

// GetModifiedTimeOk returns a tuple with the ModifiedTime field value
// and a boolean to check if the value has been set.
func (o *WebinarEmailSettingsEmailRegistrationConfirmationCustom) GetModifiedTimeOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ModifiedTime, true
}

// SetModifiedTime sets field value
func (o *WebinarEmailSettingsEmailRegistrationConfirmationCustom) SetModifiedTime(v string) {
	o.ModifiedTime = v
}

// GetSubject returns the Subject field value
func (o *WebinarEmailSettingsEmailRegistrationConfirmationCustom) GetSubject() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Subject
}

// GetSubjectOk returns a tuple with the Subject field value
// and a boolean to check if the value has been set.
func (o *WebinarEmailSettingsEmailRegistrationConfirmationCustom) GetSubjectOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Subject, true
}

// SetSubject sets field value
func (o *WebinarEmailSettingsEmailRegistrationConfirmationCustom) SetSubject(v string) {
	o.Subject = v
}

// GetType returns the Type field value
func (o *WebinarEmailSettingsEmailRegistrationConfirmationCustom) GetType() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Type
}

// GetTypeOk returns a tuple with the Type field value
// and a boolean to check if the value has been set.
func (o *WebinarEmailSettingsEmailRegistrationConfirmationCustom) GetTypeOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Type, true
}

// SetType sets field value
func (o *WebinarEmailSettingsEmailRegistrationConfirmationCustom) SetType(v string) {
	o.Type = v
}

// GetUseCalender returns the UseCalender field value
func (o *WebinarEmailSettingsEmailRegistrationConfirmationCustom) GetUseCalender() bool {
	if o == nil {
		var ret bool
		return ret
	}

	return o.UseCalender
}

// GetUseCalenderOk returns a tuple with the UseCalender field value
// and a boolean to check if the value has been set.
func (o *WebinarEmailSettingsEmailRegistrationConfirmationCustom) GetUseCalenderOk() (*bool, bool) {
	if o == nil {
		return nil, false
	}
	return &o.UseCalender, true
}

// SetUseCalender sets field value
func (o *WebinarEmailSettingsEmailRegistrationConfirmationCustom) SetUseCalender(v bool) {
	o.UseCalender = v
}

// GetUseCustomLink returns the UseCustomLink field value
func (o *WebinarEmailSettingsEmailRegistrationConfirmationCustom) GetUseCustomLink() bool {
	if o == nil {
		var ret bool
		return ret
	}

	return o.UseCustomLink
}

// GetUseCustomLinkOk returns a tuple with the UseCustomLink field value
// and a boolean to check if the value has been set.
func (o *WebinarEmailSettingsEmailRegistrationConfirmationCustom) GetUseCustomLinkOk() (*bool, bool) {
	if o == nil {
		return nil, false
	}
	return &o.UseCustomLink, true
}

// SetUseCustomLink sets field value
func (o *WebinarEmailSettingsEmailRegistrationConfirmationCustom) SetUseCustomLink(v bool) {
	o.UseCustomLink = v
}

func (o WebinarEmailSettingsEmailRegistrationConfirmationCustom) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o WebinarEmailSettingsEmailRegistrationConfirmationCustom) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["body"] = o.Body
	toSerialize["button_link"] = o.ButtonLink
	toSerialize["button_text"] = o.ButtonText
	toSerialize["header"] = o.Header
	toSerialize["modified_time"] = o.ModifiedTime
	toSerialize["subject"] = o.Subject
	toSerialize["type"] = o.Type
	toSerialize["use_calender"] = o.UseCalender
	toSerialize["use_custom_link"] = o.UseCustomLink
	return toSerialize, nil
}

type NullableWebinarEmailSettingsEmailRegistrationConfirmationCustom struct {
	value *WebinarEmailSettingsEmailRegistrationConfirmationCustom
	isSet bool
}

func (v NullableWebinarEmailSettingsEmailRegistrationConfirmationCustom) Get() *WebinarEmailSettingsEmailRegistrationConfirmationCustom {
	return v.value
}

func (v *NullableWebinarEmailSettingsEmailRegistrationConfirmationCustom) Set(val *WebinarEmailSettingsEmailRegistrationConfirmationCustom) {
	v.value = val
	v.isSet = true
}

func (v NullableWebinarEmailSettingsEmailRegistrationConfirmationCustom) IsSet() bool {
	return v.isSet
}

func (v *NullableWebinarEmailSettingsEmailRegistrationConfirmationCustom) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableWebinarEmailSettingsEmailRegistrationConfirmationCustom(val *WebinarEmailSettingsEmailRegistrationConfirmationCustom) *NullableWebinarEmailSettingsEmailRegistrationConfirmationCustom {
	return &NullableWebinarEmailSettingsEmailRegistrationConfirmationCustom{value: val, isSet: true}
}

func (v NullableWebinarEmailSettingsEmailRegistrationConfirmationCustom) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableWebinarEmailSettingsEmailRegistrationConfirmationCustom) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
