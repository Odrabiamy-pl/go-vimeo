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

// checks if the CreateWebinarAlt1Request type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &CreateWebinarAlt1Request{}

// CreateWebinarAlt1Request struct for CreateWebinarAlt1Request
type CreateWebinarAlt1Request struct {
	// The description of the webinar.
	Description   *string                                `json:"description,omitempty"`
	EmailSettings *CreateWebinarAlt1RequestEmailSettings `json:"email_settings,omitempty"`
	// The URI of the webinar's folder.
	FolderUri *float32 `json:"folder_uri,omitempty"`
	// The password when **privacy.view** is `password`. Anyone with the password can view the videos generated by streaming to the webinar event.
	Password *string                           `json:"password,omitempty"`
	Privacy  *CreateWebinarAlt1RequestPrivacy  `json:"privacy,omitempty"`
	Schedule *CreateWebinarAlt1RequestSchedule `json:"schedule,omitempty"`
	// The time zone used in resolving the timestamps that are included in the automatically generated video titles for the webinar.
	TimeZone *string `json:"time_zone,omitempty"`
	// The title of the webinar.
	Title string `json:"title"`
}

// NewCreateWebinarAlt1Request instantiates a new CreateWebinarAlt1Request object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewCreateWebinarAlt1Request(title string) *CreateWebinarAlt1Request {
	this := CreateWebinarAlt1Request{}
	this.Title = title
	return &this
}

// NewCreateWebinarAlt1RequestWithDefaults instantiates a new CreateWebinarAlt1Request object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewCreateWebinarAlt1RequestWithDefaults() *CreateWebinarAlt1Request {
	this := CreateWebinarAlt1Request{}
	return &this
}

// GetDescription returns the Description field value if set, zero value otherwise.
func (o *CreateWebinarAlt1Request) GetDescription() string {
	if o == nil || IsNil(o.Description) {
		var ret string
		return ret
	}
	return *o.Description
}

// GetDescriptionOk returns a tuple with the Description field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CreateWebinarAlt1Request) GetDescriptionOk() (*string, bool) {
	if o == nil || IsNil(o.Description) {
		return nil, false
	}
	return o.Description, true
}

// HasDescription returns a boolean if a field has been set.
func (o *CreateWebinarAlt1Request) HasDescription() bool {
	if o != nil && !IsNil(o.Description) {
		return true
	}

	return false
}

// SetDescription gets a reference to the given string and assigns it to the Description field.
func (o *CreateWebinarAlt1Request) SetDescription(v string) {
	o.Description = &v
}

// GetEmailSettings returns the EmailSettings field value if set, zero value otherwise.
func (o *CreateWebinarAlt1Request) GetEmailSettings() CreateWebinarAlt1RequestEmailSettings {
	if o == nil || IsNil(o.EmailSettings) {
		var ret CreateWebinarAlt1RequestEmailSettings
		return ret
	}
	return *o.EmailSettings
}

// GetEmailSettingsOk returns a tuple with the EmailSettings field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CreateWebinarAlt1Request) GetEmailSettingsOk() (*CreateWebinarAlt1RequestEmailSettings, bool) {
	if o == nil || IsNil(o.EmailSettings) {
		return nil, false
	}
	return o.EmailSettings, true
}

// HasEmailSettings returns a boolean if a field has been set.
func (o *CreateWebinarAlt1Request) HasEmailSettings() bool {
	if o != nil && !IsNil(o.EmailSettings) {
		return true
	}

	return false
}

// SetEmailSettings gets a reference to the given CreateWebinarAlt1RequestEmailSettings and assigns it to the EmailSettings field.
func (o *CreateWebinarAlt1Request) SetEmailSettings(v CreateWebinarAlt1RequestEmailSettings) {
	o.EmailSettings = &v
}

// GetFolderUri returns the FolderUri field value if set, zero value otherwise.
func (o *CreateWebinarAlt1Request) GetFolderUri() float32 {
	if o == nil || IsNil(o.FolderUri) {
		var ret float32
		return ret
	}
	return *o.FolderUri
}

// GetFolderUriOk returns a tuple with the FolderUri field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CreateWebinarAlt1Request) GetFolderUriOk() (*float32, bool) {
	if o == nil || IsNil(o.FolderUri) {
		return nil, false
	}
	return o.FolderUri, true
}

// HasFolderUri returns a boolean if a field has been set.
func (o *CreateWebinarAlt1Request) HasFolderUri() bool {
	if o != nil && !IsNil(o.FolderUri) {
		return true
	}

	return false
}

// SetFolderUri gets a reference to the given float32 and assigns it to the FolderUri field.
func (o *CreateWebinarAlt1Request) SetFolderUri(v float32) {
	o.FolderUri = &v
}

// GetPassword returns the Password field value if set, zero value otherwise.
func (o *CreateWebinarAlt1Request) GetPassword() string {
	if o == nil || IsNil(o.Password) {
		var ret string
		return ret
	}
	return *o.Password
}

// GetPasswordOk returns a tuple with the Password field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CreateWebinarAlt1Request) GetPasswordOk() (*string, bool) {
	if o == nil || IsNil(o.Password) {
		return nil, false
	}
	return o.Password, true
}

// HasPassword returns a boolean if a field has been set.
func (o *CreateWebinarAlt1Request) HasPassword() bool {
	if o != nil && !IsNil(o.Password) {
		return true
	}

	return false
}

// SetPassword gets a reference to the given string and assigns it to the Password field.
func (o *CreateWebinarAlt1Request) SetPassword(v string) {
	o.Password = &v
}

// GetPrivacy returns the Privacy field value if set, zero value otherwise.
func (o *CreateWebinarAlt1Request) GetPrivacy() CreateWebinarAlt1RequestPrivacy {
	if o == nil || IsNil(o.Privacy) {
		var ret CreateWebinarAlt1RequestPrivacy
		return ret
	}
	return *o.Privacy
}

// GetPrivacyOk returns a tuple with the Privacy field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CreateWebinarAlt1Request) GetPrivacyOk() (*CreateWebinarAlt1RequestPrivacy, bool) {
	if o == nil || IsNil(o.Privacy) {
		return nil, false
	}
	return o.Privacy, true
}

// HasPrivacy returns a boolean if a field has been set.
func (o *CreateWebinarAlt1Request) HasPrivacy() bool {
	if o != nil && !IsNil(o.Privacy) {
		return true
	}

	return false
}

// SetPrivacy gets a reference to the given CreateWebinarAlt1RequestPrivacy and assigns it to the Privacy field.
func (o *CreateWebinarAlt1Request) SetPrivacy(v CreateWebinarAlt1RequestPrivacy) {
	o.Privacy = &v
}

// GetSchedule returns the Schedule field value if set, zero value otherwise.
func (o *CreateWebinarAlt1Request) GetSchedule() CreateWebinarAlt1RequestSchedule {
	if o == nil || IsNil(o.Schedule) {
		var ret CreateWebinarAlt1RequestSchedule
		return ret
	}
	return *o.Schedule
}

// GetScheduleOk returns a tuple with the Schedule field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CreateWebinarAlt1Request) GetScheduleOk() (*CreateWebinarAlt1RequestSchedule, bool) {
	if o == nil || IsNil(o.Schedule) {
		return nil, false
	}
	return o.Schedule, true
}

// HasSchedule returns a boolean if a field has been set.
func (o *CreateWebinarAlt1Request) HasSchedule() bool {
	if o != nil && !IsNil(o.Schedule) {
		return true
	}

	return false
}

// SetSchedule gets a reference to the given CreateWebinarAlt1RequestSchedule and assigns it to the Schedule field.
func (o *CreateWebinarAlt1Request) SetSchedule(v CreateWebinarAlt1RequestSchedule) {
	o.Schedule = &v
}

// GetTimeZone returns the TimeZone field value if set, zero value otherwise.
func (o *CreateWebinarAlt1Request) GetTimeZone() string {
	if o == nil || IsNil(o.TimeZone) {
		var ret string
		return ret
	}
	return *o.TimeZone
}

// GetTimeZoneOk returns a tuple with the TimeZone field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CreateWebinarAlt1Request) GetTimeZoneOk() (*string, bool) {
	if o == nil || IsNil(o.TimeZone) {
		return nil, false
	}
	return o.TimeZone, true
}

// HasTimeZone returns a boolean if a field has been set.
func (o *CreateWebinarAlt1Request) HasTimeZone() bool {
	if o != nil && !IsNil(o.TimeZone) {
		return true
	}

	return false
}

// SetTimeZone gets a reference to the given string and assigns it to the TimeZone field.
func (o *CreateWebinarAlt1Request) SetTimeZone(v string) {
	o.TimeZone = &v
}

// GetTitle returns the Title field value
func (o *CreateWebinarAlt1Request) GetTitle() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Title
}

// GetTitleOk returns a tuple with the Title field value
// and a boolean to check if the value has been set.
func (o *CreateWebinarAlt1Request) GetTitleOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Title, true
}

// SetTitle sets field value
func (o *CreateWebinarAlt1Request) SetTitle(v string) {
	o.Title = v
}

func (o CreateWebinarAlt1Request) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o CreateWebinarAlt1Request) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Description) {
		toSerialize["description"] = o.Description
	}
	if !IsNil(o.EmailSettings) {
		toSerialize["email_settings"] = o.EmailSettings
	}
	if !IsNil(o.FolderUri) {
		toSerialize["folder_uri"] = o.FolderUri
	}
	if !IsNil(o.Password) {
		toSerialize["password"] = o.Password
	}
	if !IsNil(o.Privacy) {
		toSerialize["privacy"] = o.Privacy
	}
	if !IsNil(o.Schedule) {
		toSerialize["schedule"] = o.Schedule
	}
	if !IsNil(o.TimeZone) {
		toSerialize["time_zone"] = o.TimeZone
	}
	toSerialize["title"] = o.Title
	return toSerialize, nil
}

type NullableCreateWebinarAlt1Request struct {
	value *CreateWebinarAlt1Request
	isSet bool
}

func (v NullableCreateWebinarAlt1Request) Get() *CreateWebinarAlt1Request {
	return v.value
}

func (v *NullableCreateWebinarAlt1Request) Set(val *CreateWebinarAlt1Request) {
	v.value = val
	v.isSet = true
}

func (v NullableCreateWebinarAlt1Request) IsSet() bool {
	return v.isSet
}

func (v *NullableCreateWebinarAlt1Request) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableCreateWebinarAlt1Request(val *CreateWebinarAlt1Request) *NullableCreateWebinarAlt1Request {
	return &NullableCreateWebinarAlt1Request{value: val, isSet: true}
}

func (v NullableCreateWebinarAlt1Request) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableCreateWebinarAlt1Request) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
