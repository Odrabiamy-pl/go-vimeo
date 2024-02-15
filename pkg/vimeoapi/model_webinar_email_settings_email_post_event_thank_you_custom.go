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

// checks if the WebinarEmailSettingsEmailPostEventThankYouCustom type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &WebinarEmailSettingsEmailPostEventThankYouCustom{}

// WebinarEmailSettingsEmailPostEventThankYouCustom The email custom details for the webinar post-event thank-you email.
type WebinarEmailSettingsEmailPostEventThankYouCustom struct {
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

type _WebinarEmailSettingsEmailPostEventThankYouCustom WebinarEmailSettingsEmailPostEventThankYouCustom

// NewWebinarEmailSettingsEmailPostEventThankYouCustom instantiates a new WebinarEmailSettingsEmailPostEventThankYouCustom object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewWebinarEmailSettingsEmailPostEventThankYouCustom(body string, buttonLink string, buttonText string, header string, modifiedTime string, subject string, type_ string, useCalender bool, useCustomLink bool) *WebinarEmailSettingsEmailPostEventThankYouCustom {
	this := WebinarEmailSettingsEmailPostEventThankYouCustom{}
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

// NewWebinarEmailSettingsEmailPostEventThankYouCustomWithDefaults instantiates a new WebinarEmailSettingsEmailPostEventThankYouCustom object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewWebinarEmailSettingsEmailPostEventThankYouCustomWithDefaults() *WebinarEmailSettingsEmailPostEventThankYouCustom {
	this := WebinarEmailSettingsEmailPostEventThankYouCustom{}
	return &this
}

// GetBody returns the Body field value
func (o *WebinarEmailSettingsEmailPostEventThankYouCustom) GetBody() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Body
}

// GetBodyOk returns a tuple with the Body field value
// and a boolean to check if the value has been set.
func (o *WebinarEmailSettingsEmailPostEventThankYouCustom) GetBodyOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Body, true
}

// SetBody sets field value
func (o *WebinarEmailSettingsEmailPostEventThankYouCustom) SetBody(v string) {
	o.Body = v
}

// GetButtonLink returns the ButtonLink field value
func (o *WebinarEmailSettingsEmailPostEventThankYouCustom) GetButtonLink() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.ButtonLink
}

// GetButtonLinkOk returns a tuple with the ButtonLink field value
// and a boolean to check if the value has been set.
func (o *WebinarEmailSettingsEmailPostEventThankYouCustom) GetButtonLinkOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ButtonLink, true
}

// SetButtonLink sets field value
func (o *WebinarEmailSettingsEmailPostEventThankYouCustom) SetButtonLink(v string) {
	o.ButtonLink = v
}

// GetButtonText returns the ButtonText field value
func (o *WebinarEmailSettingsEmailPostEventThankYouCustom) GetButtonText() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.ButtonText
}

// GetButtonTextOk returns a tuple with the ButtonText field value
// and a boolean to check if the value has been set.
func (o *WebinarEmailSettingsEmailPostEventThankYouCustom) GetButtonTextOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ButtonText, true
}

// SetButtonText sets field value
func (o *WebinarEmailSettingsEmailPostEventThankYouCustom) SetButtonText(v string) {
	o.ButtonText = v
}

// GetHeader returns the Header field value
func (o *WebinarEmailSettingsEmailPostEventThankYouCustom) GetHeader() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Header
}

// GetHeaderOk returns a tuple with the Header field value
// and a boolean to check if the value has been set.
func (o *WebinarEmailSettingsEmailPostEventThankYouCustom) GetHeaderOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Header, true
}

// SetHeader sets field value
func (o *WebinarEmailSettingsEmailPostEventThankYouCustom) SetHeader(v string) {
	o.Header = v
}

// GetModifiedTime returns the ModifiedTime field value
func (o *WebinarEmailSettingsEmailPostEventThankYouCustom) GetModifiedTime() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.ModifiedTime
}

// GetModifiedTimeOk returns a tuple with the ModifiedTime field value
// and a boolean to check if the value has been set.
func (o *WebinarEmailSettingsEmailPostEventThankYouCustom) GetModifiedTimeOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ModifiedTime, true
}

// SetModifiedTime sets field value
func (o *WebinarEmailSettingsEmailPostEventThankYouCustom) SetModifiedTime(v string) {
	o.ModifiedTime = v
}

// GetSubject returns the Subject field value
func (o *WebinarEmailSettingsEmailPostEventThankYouCustom) GetSubject() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Subject
}

// GetSubjectOk returns a tuple with the Subject field value
// and a boolean to check if the value has been set.
func (o *WebinarEmailSettingsEmailPostEventThankYouCustom) GetSubjectOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Subject, true
}

// SetSubject sets field value
func (o *WebinarEmailSettingsEmailPostEventThankYouCustom) SetSubject(v string) {
	o.Subject = v
}

// GetType returns the Type field value
func (o *WebinarEmailSettingsEmailPostEventThankYouCustom) GetType() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Type
}

// GetTypeOk returns a tuple with the Type field value
// and a boolean to check if the value has been set.
func (o *WebinarEmailSettingsEmailPostEventThankYouCustom) GetTypeOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Type, true
}

// SetType sets field value
func (o *WebinarEmailSettingsEmailPostEventThankYouCustom) SetType(v string) {
	o.Type = v
}

// GetUseCalender returns the UseCalender field value
func (o *WebinarEmailSettingsEmailPostEventThankYouCustom) GetUseCalender() bool {
	if o == nil {
		var ret bool
		return ret
	}

	return o.UseCalender
}

// GetUseCalenderOk returns a tuple with the UseCalender field value
// and a boolean to check if the value has been set.
func (o *WebinarEmailSettingsEmailPostEventThankYouCustom) GetUseCalenderOk() (*bool, bool) {
	if o == nil {
		return nil, false
	}
	return &o.UseCalender, true
}

// SetUseCalender sets field value
func (o *WebinarEmailSettingsEmailPostEventThankYouCustom) SetUseCalender(v bool) {
	o.UseCalender = v
}

// GetUseCustomLink returns the UseCustomLink field value
func (o *WebinarEmailSettingsEmailPostEventThankYouCustom) GetUseCustomLink() bool {
	if o == nil {
		var ret bool
		return ret
	}

	return o.UseCustomLink
}

// GetUseCustomLinkOk returns a tuple with the UseCustomLink field value
// and a boolean to check if the value has been set.
func (o *WebinarEmailSettingsEmailPostEventThankYouCustom) GetUseCustomLinkOk() (*bool, bool) {
	if o == nil {
		return nil, false
	}
	return &o.UseCustomLink, true
}

// SetUseCustomLink sets field value
func (o *WebinarEmailSettingsEmailPostEventThankYouCustom) SetUseCustomLink(v bool) {
	o.UseCustomLink = v
}

func (o WebinarEmailSettingsEmailPostEventThankYouCustom) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o WebinarEmailSettingsEmailPostEventThankYouCustom) ToMap() (map[string]interface{}, error) {
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

func (o *WebinarEmailSettingsEmailPostEventThankYouCustom) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"body",
		"button_link",
		"button_text",
		"header",
		"modified_time",
		"subject",
		"type",
		"use_calender",
		"use_custom_link",
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

	varWebinarEmailSettingsEmailPostEventThankYouCustom := _WebinarEmailSettingsEmailPostEventThankYouCustom{}

	decoder := json.NewDecoder(bytes.NewReader(data))
	decoder.DisallowUnknownFields()
	err = decoder.Decode(&varWebinarEmailSettingsEmailPostEventThankYouCustom)

	if err != nil {
		return err
	}

	*o = WebinarEmailSettingsEmailPostEventThankYouCustom(varWebinarEmailSettingsEmailPostEventThankYouCustom)

	return err
}

type NullableWebinarEmailSettingsEmailPostEventThankYouCustom struct {
	value *WebinarEmailSettingsEmailPostEventThankYouCustom
	isSet bool
}

func (v NullableWebinarEmailSettingsEmailPostEventThankYouCustom) Get() *WebinarEmailSettingsEmailPostEventThankYouCustom {
	return v.value
}

func (v *NullableWebinarEmailSettingsEmailPostEventThankYouCustom) Set(val *WebinarEmailSettingsEmailPostEventThankYouCustom) {
	v.value = val
	v.isSet = true
}

func (v NullableWebinarEmailSettingsEmailPostEventThankYouCustom) IsSet() bool {
	return v.isSet
}

func (v *NullableWebinarEmailSettingsEmailPostEventThankYouCustom) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableWebinarEmailSettingsEmailPostEventThankYouCustom(val *WebinarEmailSettingsEmailPostEventThankYouCustom) *NullableWebinarEmailSettingsEmailPostEventThankYouCustom {
	return &NullableWebinarEmailSettingsEmailPostEventThankYouCustom{value: val, isSet: true}
}

func (v NullableWebinarEmailSettingsEmailPostEventThankYouCustom) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableWebinarEmailSettingsEmailPostEventThankYouCustom) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


