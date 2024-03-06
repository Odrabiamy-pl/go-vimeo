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

// checks if the UpdateWebinarEmailSettingsAlt1RequestEmailPreferences type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &UpdateWebinarEmailSettingsAlt1RequestEmailPreferences{}

// UpdateWebinarEmailSettingsAlt1RequestEmailPreferences The preferences for emails that are sent about the webinar.
type UpdateWebinarEmailSettingsAlt1RequestEmailPreferences struct {
	// Whether to send a reminder email 15 minutes before the webinar starts.
	EmailEventReminder15Min *bool `json:"email_event_reminder_15_min,omitempty"`
	// Whether to send a reminder email 1 hour before the webinar starts.
	EmailEventReminder1Hrs *bool `json:"email_event_reminder_1_hrs,omitempty"`
	// Whether to send a reminder email 24 hours before the webinar starts.
	EmailEventReminder24Hrs *bool `json:"email_event_reminder_24_hrs,omitempty"`
	// Whether to send post-event thank-you emails to no-shows.
	EmailPostEventNoShowThankYou *bool `json:"email_post_event_no_show_thank_you,omitempty"`
	// Whether to send post-event thank-you emails.
	EmailPostEventThankYou *bool `json:"email_post_event_thank_you,omitempty"`
	// Whether to send a registration confirmation email after webinar registration.
	EmailRegistrationConfirmation *bool `json:"email_registration_confirmation,omitempty"`
}

// NewUpdateWebinarEmailSettingsAlt1RequestEmailPreferences instantiates a new UpdateWebinarEmailSettingsAlt1RequestEmailPreferences object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewUpdateWebinarEmailSettingsAlt1RequestEmailPreferences() *UpdateWebinarEmailSettingsAlt1RequestEmailPreferences {
	this := UpdateWebinarEmailSettingsAlt1RequestEmailPreferences{}
	return &this
}

// NewUpdateWebinarEmailSettingsAlt1RequestEmailPreferencesWithDefaults instantiates a new UpdateWebinarEmailSettingsAlt1RequestEmailPreferences object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewUpdateWebinarEmailSettingsAlt1RequestEmailPreferencesWithDefaults() *UpdateWebinarEmailSettingsAlt1RequestEmailPreferences {
	this := UpdateWebinarEmailSettingsAlt1RequestEmailPreferences{}
	return &this
}

// GetEmailEventReminder15Min returns the EmailEventReminder15Min field value if set, zero value otherwise.
func (o *UpdateWebinarEmailSettingsAlt1RequestEmailPreferences) GetEmailEventReminder15Min() bool {
	if o == nil || IsNil(o.EmailEventReminder15Min) {
		var ret bool
		return ret
	}
	return *o.EmailEventReminder15Min
}

// GetEmailEventReminder15MinOk returns a tuple with the EmailEventReminder15Min field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UpdateWebinarEmailSettingsAlt1RequestEmailPreferences) GetEmailEventReminder15MinOk() (*bool, bool) {
	if o == nil || IsNil(o.EmailEventReminder15Min) {
		return nil, false
	}
	return o.EmailEventReminder15Min, true
}

// HasEmailEventReminder15Min returns a boolean if a field has been set.
func (o *UpdateWebinarEmailSettingsAlt1RequestEmailPreferences) HasEmailEventReminder15Min() bool {
	if o != nil && !IsNil(o.EmailEventReminder15Min) {
		return true
	}

	return false
}

// SetEmailEventReminder15Min gets a reference to the given bool and assigns it to the EmailEventReminder15Min field.
func (o *UpdateWebinarEmailSettingsAlt1RequestEmailPreferences) SetEmailEventReminder15Min(v bool) {
	o.EmailEventReminder15Min = &v
}

// GetEmailEventReminder1Hrs returns the EmailEventReminder1Hrs field value if set, zero value otherwise.
func (o *UpdateWebinarEmailSettingsAlt1RequestEmailPreferences) GetEmailEventReminder1Hrs() bool {
	if o == nil || IsNil(o.EmailEventReminder1Hrs) {
		var ret bool
		return ret
	}
	return *o.EmailEventReminder1Hrs
}

// GetEmailEventReminder1HrsOk returns a tuple with the EmailEventReminder1Hrs field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UpdateWebinarEmailSettingsAlt1RequestEmailPreferences) GetEmailEventReminder1HrsOk() (*bool, bool) {
	if o == nil || IsNil(o.EmailEventReminder1Hrs) {
		return nil, false
	}
	return o.EmailEventReminder1Hrs, true
}

// HasEmailEventReminder1Hrs returns a boolean if a field has been set.
func (o *UpdateWebinarEmailSettingsAlt1RequestEmailPreferences) HasEmailEventReminder1Hrs() bool {
	if o != nil && !IsNil(o.EmailEventReminder1Hrs) {
		return true
	}

	return false
}

// SetEmailEventReminder1Hrs gets a reference to the given bool and assigns it to the EmailEventReminder1Hrs field.
func (o *UpdateWebinarEmailSettingsAlt1RequestEmailPreferences) SetEmailEventReminder1Hrs(v bool) {
	o.EmailEventReminder1Hrs = &v
}

// GetEmailEventReminder24Hrs returns the EmailEventReminder24Hrs field value if set, zero value otherwise.
func (o *UpdateWebinarEmailSettingsAlt1RequestEmailPreferences) GetEmailEventReminder24Hrs() bool {
	if o == nil || IsNil(o.EmailEventReminder24Hrs) {
		var ret bool
		return ret
	}
	return *o.EmailEventReminder24Hrs
}

// GetEmailEventReminder24HrsOk returns a tuple with the EmailEventReminder24Hrs field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UpdateWebinarEmailSettingsAlt1RequestEmailPreferences) GetEmailEventReminder24HrsOk() (*bool, bool) {
	if o == nil || IsNil(o.EmailEventReminder24Hrs) {
		return nil, false
	}
	return o.EmailEventReminder24Hrs, true
}

// HasEmailEventReminder24Hrs returns a boolean if a field has been set.
func (o *UpdateWebinarEmailSettingsAlt1RequestEmailPreferences) HasEmailEventReminder24Hrs() bool {
	if o != nil && !IsNil(o.EmailEventReminder24Hrs) {
		return true
	}

	return false
}

// SetEmailEventReminder24Hrs gets a reference to the given bool and assigns it to the EmailEventReminder24Hrs field.
func (o *UpdateWebinarEmailSettingsAlt1RequestEmailPreferences) SetEmailEventReminder24Hrs(v bool) {
	o.EmailEventReminder24Hrs = &v
}

// GetEmailPostEventNoShowThankYou returns the EmailPostEventNoShowThankYou field value if set, zero value otherwise.
func (o *UpdateWebinarEmailSettingsAlt1RequestEmailPreferences) GetEmailPostEventNoShowThankYou() bool {
	if o == nil || IsNil(o.EmailPostEventNoShowThankYou) {
		var ret bool
		return ret
	}
	return *o.EmailPostEventNoShowThankYou
}

// GetEmailPostEventNoShowThankYouOk returns a tuple with the EmailPostEventNoShowThankYou field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UpdateWebinarEmailSettingsAlt1RequestEmailPreferences) GetEmailPostEventNoShowThankYouOk() (*bool, bool) {
	if o == nil || IsNil(o.EmailPostEventNoShowThankYou) {
		return nil, false
	}
	return o.EmailPostEventNoShowThankYou, true
}

// HasEmailPostEventNoShowThankYou returns a boolean if a field has been set.
func (o *UpdateWebinarEmailSettingsAlt1RequestEmailPreferences) HasEmailPostEventNoShowThankYou() bool {
	if o != nil && !IsNil(o.EmailPostEventNoShowThankYou) {
		return true
	}

	return false
}

// SetEmailPostEventNoShowThankYou gets a reference to the given bool and assigns it to the EmailPostEventNoShowThankYou field.
func (o *UpdateWebinarEmailSettingsAlt1RequestEmailPreferences) SetEmailPostEventNoShowThankYou(v bool) {
	o.EmailPostEventNoShowThankYou = &v
}

// GetEmailPostEventThankYou returns the EmailPostEventThankYou field value if set, zero value otherwise.
func (o *UpdateWebinarEmailSettingsAlt1RequestEmailPreferences) GetEmailPostEventThankYou() bool {
	if o == nil || IsNil(o.EmailPostEventThankYou) {
		var ret bool
		return ret
	}
	return *o.EmailPostEventThankYou
}

// GetEmailPostEventThankYouOk returns a tuple with the EmailPostEventThankYou field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UpdateWebinarEmailSettingsAlt1RequestEmailPreferences) GetEmailPostEventThankYouOk() (*bool, bool) {
	if o == nil || IsNil(o.EmailPostEventThankYou) {
		return nil, false
	}
	return o.EmailPostEventThankYou, true
}

// HasEmailPostEventThankYou returns a boolean if a field has been set.
func (o *UpdateWebinarEmailSettingsAlt1RequestEmailPreferences) HasEmailPostEventThankYou() bool {
	if o != nil && !IsNil(o.EmailPostEventThankYou) {
		return true
	}

	return false
}

// SetEmailPostEventThankYou gets a reference to the given bool and assigns it to the EmailPostEventThankYou field.
func (o *UpdateWebinarEmailSettingsAlt1RequestEmailPreferences) SetEmailPostEventThankYou(v bool) {
	o.EmailPostEventThankYou = &v
}

// GetEmailRegistrationConfirmation returns the EmailRegistrationConfirmation field value if set, zero value otherwise.
func (o *UpdateWebinarEmailSettingsAlt1RequestEmailPreferences) GetEmailRegistrationConfirmation() bool {
	if o == nil || IsNil(o.EmailRegistrationConfirmation) {
		var ret bool
		return ret
	}
	return *o.EmailRegistrationConfirmation
}

// GetEmailRegistrationConfirmationOk returns a tuple with the EmailRegistrationConfirmation field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UpdateWebinarEmailSettingsAlt1RequestEmailPreferences) GetEmailRegistrationConfirmationOk() (*bool, bool) {
	if o == nil || IsNil(o.EmailRegistrationConfirmation) {
		return nil, false
	}
	return o.EmailRegistrationConfirmation, true
}

// HasEmailRegistrationConfirmation returns a boolean if a field has been set.
func (o *UpdateWebinarEmailSettingsAlt1RequestEmailPreferences) HasEmailRegistrationConfirmation() bool {
	if o != nil && !IsNil(o.EmailRegistrationConfirmation) {
		return true
	}

	return false
}

// SetEmailRegistrationConfirmation gets a reference to the given bool and assigns it to the EmailRegistrationConfirmation field.
func (o *UpdateWebinarEmailSettingsAlt1RequestEmailPreferences) SetEmailRegistrationConfirmation(v bool) {
	o.EmailRegistrationConfirmation = &v
}

func (o UpdateWebinarEmailSettingsAlt1RequestEmailPreferences) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o UpdateWebinarEmailSettingsAlt1RequestEmailPreferences) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.EmailEventReminder15Min) {
		toSerialize["email_event_reminder_15_min"] = o.EmailEventReminder15Min
	}
	if !IsNil(o.EmailEventReminder1Hrs) {
		toSerialize["email_event_reminder_1_hrs"] = o.EmailEventReminder1Hrs
	}
	if !IsNil(o.EmailEventReminder24Hrs) {
		toSerialize["email_event_reminder_24_hrs"] = o.EmailEventReminder24Hrs
	}
	if !IsNil(o.EmailPostEventNoShowThankYou) {
		toSerialize["email_post_event_no_show_thank_you"] = o.EmailPostEventNoShowThankYou
	}
	if !IsNil(o.EmailPostEventThankYou) {
		toSerialize["email_post_event_thank_you"] = o.EmailPostEventThankYou
	}
	if !IsNil(o.EmailRegistrationConfirmation) {
		toSerialize["email_registration_confirmation"] = o.EmailRegistrationConfirmation
	}
	return toSerialize, nil
}

type NullableUpdateWebinarEmailSettingsAlt1RequestEmailPreferences struct {
	value *UpdateWebinarEmailSettingsAlt1RequestEmailPreferences
	isSet bool
}

func (v NullableUpdateWebinarEmailSettingsAlt1RequestEmailPreferences) Get() *UpdateWebinarEmailSettingsAlt1RequestEmailPreferences {
	return v.value
}

func (v *NullableUpdateWebinarEmailSettingsAlt1RequestEmailPreferences) Set(val *UpdateWebinarEmailSettingsAlt1RequestEmailPreferences) {
	v.value = val
	v.isSet = true
}

func (v NullableUpdateWebinarEmailSettingsAlt1RequestEmailPreferences) IsSet() bool {
	return v.isSet
}

func (v *NullableUpdateWebinarEmailSettingsAlt1RequestEmailPreferences) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableUpdateWebinarEmailSettingsAlt1RequestEmailPreferences(val *UpdateWebinarEmailSettingsAlt1RequestEmailPreferences) *NullableUpdateWebinarEmailSettingsAlt1RequestEmailPreferences {
	return &NullableUpdateWebinarEmailSettingsAlt1RequestEmailPreferences{value: val, isSet: true}
}

func (v NullableUpdateWebinarEmailSettingsAlt1RequestEmailPreferences) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableUpdateWebinarEmailSettingsAlt1RequestEmailPreferences) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
