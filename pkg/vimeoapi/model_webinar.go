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

// checks if the Webinar type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &Webinar{}

// Webinar struct for Webinar
type Webinar struct {
	// The time in ISO 8601 format when the webinar was completed.
	CompletedOn string `json:"completed_on"`
	// The time in ISO 8601 format when the webinar was created.
	CreatedTime string `json:"created_time"`
	// The description of the webinar.
	Description NullableString `json:"description"`
	Edit NullableLiveEventMetadataInteractionsEdit `json:"edit"`
	// Information about the email provider list that is selected to import registrants.
	EmailProviderList []WebinarEmailProviderListInner `json:"email_provider_list"`
	EmailQuota WebinarEmailQuota `json:"email_quota"`
	EmailSettings WebinarEmailSettings `json:"email_settings"`
	Events NullableWebinarEvents `json:"events"`
	// Whether polls are associated with the webinar.
	HasPolls bool `json:"has_polls"`
	Metadata WebinarMetadata `json:"metadata"`
	// The time in ISO 8601 format when the webinar was modified.
	ModifiedOn string `json:"modified_on"`
	// The date in ISO 8601 format on which the next occurrence of the webinar is expected to be live.
	NextOccurrenceTime NullableString `json:"next_occurrence_time"`
	// The password used to access the videos generated by streaming to the webinar event.
	Password NullableString `json:"password"`
	Privacy WebinarPrivacy `json:"privacy"`
	RegistrationData WebinarRegistrationData `json:"registration_data"`
	RegistrationForm EmailCaptureForm `json:"registration_form"`
	Schedule NullableWebinarSchedule `json:"schedule"`
	// The status of the webinar.  Option descriptions:  * `ended` - The webinar has ended.  * `started` - The webinar has started. 
	Status NullableString `json:"status"`
	// The time zone used in resolving the timestamps that are included in the automatically generated video titles for the webinar.
	TimeZone string `json:"time_zone"`
	// The title of the webinar.
	Title NullableString `json:"title"`
	// The webinar's canonical relative URI.
	Uri string `json:"uri"`
	User User `json:"user"`
}

// NewWebinar instantiates a new Webinar object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewWebinar(completedOn string, createdTime string, description NullableString, edit NullableLiveEventMetadataInteractionsEdit, emailProviderList []WebinarEmailProviderListInner, emailQuota WebinarEmailQuota, emailSettings WebinarEmailSettings, events NullableWebinarEvents, hasPolls bool, metadata WebinarMetadata, modifiedOn string, nextOccurrenceTime NullableString, password NullableString, privacy WebinarPrivacy, registrationData WebinarRegistrationData, registrationForm EmailCaptureForm, schedule NullableWebinarSchedule, status NullableString, timeZone string, title NullableString, uri string, user User) *Webinar {
	this := Webinar{}
	this.CompletedOn = completedOn
	this.CreatedTime = createdTime
	this.Description = description
	this.Edit = edit
	this.EmailProviderList = emailProviderList
	this.EmailQuota = emailQuota
	this.EmailSettings = emailSettings
	this.Events = events
	this.HasPolls = hasPolls
	this.Metadata = metadata
	this.ModifiedOn = modifiedOn
	this.NextOccurrenceTime = nextOccurrenceTime
	this.Password = password
	this.Privacy = privacy
	this.RegistrationData = registrationData
	this.RegistrationForm = registrationForm
	this.Schedule = schedule
	this.Status = status
	this.TimeZone = timeZone
	this.Title = title
	this.Uri = uri
	this.User = user
	return &this
}

// NewWebinarWithDefaults instantiates a new Webinar object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewWebinarWithDefaults() *Webinar {
	this := Webinar{}
	return &this
}

// GetCompletedOn returns the CompletedOn field value
func (o *Webinar) GetCompletedOn() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.CompletedOn
}

// GetCompletedOnOk returns a tuple with the CompletedOn field value
// and a boolean to check if the value has been set.
func (o *Webinar) GetCompletedOnOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.CompletedOn, true
}

// SetCompletedOn sets field value
func (o *Webinar) SetCompletedOn(v string) {
	o.CompletedOn = v
}

// GetCreatedTime returns the CreatedTime field value
func (o *Webinar) GetCreatedTime() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.CreatedTime
}

// GetCreatedTimeOk returns a tuple with the CreatedTime field value
// and a boolean to check if the value has been set.
func (o *Webinar) GetCreatedTimeOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.CreatedTime, true
}

// SetCreatedTime sets field value
func (o *Webinar) SetCreatedTime(v string) {
	o.CreatedTime = v
}

// GetDescription returns the Description field value
// If the value is explicit nil, the zero value for string will be returned
func (o *Webinar) GetDescription() string {
	if o == nil || o.Description.Get() == nil {
		var ret string
		return ret
	}

	return *o.Description.Get()
}

// GetDescriptionOk returns a tuple with the Description field value
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *Webinar) GetDescriptionOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.Description.Get(), o.Description.IsSet()
}

// SetDescription sets field value
func (o *Webinar) SetDescription(v string) {
	o.Description.Set(&v)
}

// GetEdit returns the Edit field value
// If the value is explicit nil, the zero value for LiveEventMetadataInteractionsEdit will be returned
func (o *Webinar) GetEdit() LiveEventMetadataInteractionsEdit {
	if o == nil || o.Edit.Get() == nil {
		var ret LiveEventMetadataInteractionsEdit
		return ret
	}

	return *o.Edit.Get()
}

// GetEditOk returns a tuple with the Edit field value
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *Webinar) GetEditOk() (*LiveEventMetadataInteractionsEdit, bool) {
	if o == nil {
		return nil, false
	}
	return o.Edit.Get(), o.Edit.IsSet()
}

// SetEdit sets field value
func (o *Webinar) SetEdit(v LiveEventMetadataInteractionsEdit) {
	o.Edit.Set(&v)
}

// GetEmailProviderList returns the EmailProviderList field value
func (o *Webinar) GetEmailProviderList() []WebinarEmailProviderListInner {
	if o == nil {
		var ret []WebinarEmailProviderListInner
		return ret
	}

	return o.EmailProviderList
}

// GetEmailProviderListOk returns a tuple with the EmailProviderList field value
// and a boolean to check if the value has been set.
func (o *Webinar) GetEmailProviderListOk() ([]WebinarEmailProviderListInner, bool) {
	if o == nil {
		return nil, false
	}
	return o.EmailProviderList, true
}

// SetEmailProviderList sets field value
func (o *Webinar) SetEmailProviderList(v []WebinarEmailProviderListInner) {
	o.EmailProviderList = v
}

// GetEmailQuota returns the EmailQuota field value
func (o *Webinar) GetEmailQuota() WebinarEmailQuota {
	if o == nil {
		var ret WebinarEmailQuota
		return ret
	}

	return o.EmailQuota
}

// GetEmailQuotaOk returns a tuple with the EmailQuota field value
// and a boolean to check if the value has been set.
func (o *Webinar) GetEmailQuotaOk() (*WebinarEmailQuota, bool) {
	if o == nil {
		return nil, false
	}
	return &o.EmailQuota, true
}

// SetEmailQuota sets field value
func (o *Webinar) SetEmailQuota(v WebinarEmailQuota) {
	o.EmailQuota = v
}

// GetEmailSettings returns the EmailSettings field value
func (o *Webinar) GetEmailSettings() WebinarEmailSettings {
	if o == nil {
		var ret WebinarEmailSettings
		return ret
	}

	return o.EmailSettings
}

// GetEmailSettingsOk returns a tuple with the EmailSettings field value
// and a boolean to check if the value has been set.
func (o *Webinar) GetEmailSettingsOk() (*WebinarEmailSettings, bool) {
	if o == nil {
		return nil, false
	}
	return &o.EmailSettings, true
}

// SetEmailSettings sets field value
func (o *Webinar) SetEmailSettings(v WebinarEmailSettings) {
	o.EmailSettings = v
}

// GetEvents returns the Events field value
// If the value is explicit nil, the zero value for WebinarEvents will be returned
func (o *Webinar) GetEvents() WebinarEvents {
	if o == nil || o.Events.Get() == nil {
		var ret WebinarEvents
		return ret
	}

	return *o.Events.Get()
}

// GetEventsOk returns a tuple with the Events field value
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *Webinar) GetEventsOk() (*WebinarEvents, bool) {
	if o == nil {
		return nil, false
	}
	return o.Events.Get(), o.Events.IsSet()
}

// SetEvents sets field value
func (o *Webinar) SetEvents(v WebinarEvents) {
	o.Events.Set(&v)
}

// GetHasPolls returns the HasPolls field value
func (o *Webinar) GetHasPolls() bool {
	if o == nil {
		var ret bool
		return ret
	}

	return o.HasPolls
}

// GetHasPollsOk returns a tuple with the HasPolls field value
// and a boolean to check if the value has been set.
func (o *Webinar) GetHasPollsOk() (*bool, bool) {
	if o == nil {
		return nil, false
	}
	return &o.HasPolls, true
}

// SetHasPolls sets field value
func (o *Webinar) SetHasPolls(v bool) {
	o.HasPolls = v
}

// GetMetadata returns the Metadata field value
func (o *Webinar) GetMetadata() WebinarMetadata {
	if o == nil {
		var ret WebinarMetadata
		return ret
	}

	return o.Metadata
}

// GetMetadataOk returns a tuple with the Metadata field value
// and a boolean to check if the value has been set.
func (o *Webinar) GetMetadataOk() (*WebinarMetadata, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Metadata, true
}

// SetMetadata sets field value
func (o *Webinar) SetMetadata(v WebinarMetadata) {
	o.Metadata = v
}

// GetModifiedOn returns the ModifiedOn field value
func (o *Webinar) GetModifiedOn() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.ModifiedOn
}

// GetModifiedOnOk returns a tuple with the ModifiedOn field value
// and a boolean to check if the value has been set.
func (o *Webinar) GetModifiedOnOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ModifiedOn, true
}

// SetModifiedOn sets field value
func (o *Webinar) SetModifiedOn(v string) {
	o.ModifiedOn = v
}

// GetNextOccurrenceTime returns the NextOccurrenceTime field value
// If the value is explicit nil, the zero value for string will be returned
func (o *Webinar) GetNextOccurrenceTime() string {
	if o == nil || o.NextOccurrenceTime.Get() == nil {
		var ret string
		return ret
	}

	return *o.NextOccurrenceTime.Get()
}

// GetNextOccurrenceTimeOk returns a tuple with the NextOccurrenceTime field value
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *Webinar) GetNextOccurrenceTimeOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.NextOccurrenceTime.Get(), o.NextOccurrenceTime.IsSet()
}

// SetNextOccurrenceTime sets field value
func (o *Webinar) SetNextOccurrenceTime(v string) {
	o.NextOccurrenceTime.Set(&v)
}

// GetPassword returns the Password field value
// If the value is explicit nil, the zero value for string will be returned
func (o *Webinar) GetPassword() string {
	if o == nil || o.Password.Get() == nil {
		var ret string
		return ret
	}

	return *o.Password.Get()
}

// GetPasswordOk returns a tuple with the Password field value
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *Webinar) GetPasswordOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.Password.Get(), o.Password.IsSet()
}

// SetPassword sets field value
func (o *Webinar) SetPassword(v string) {
	o.Password.Set(&v)
}

// GetPrivacy returns the Privacy field value
func (o *Webinar) GetPrivacy() WebinarPrivacy {
	if o == nil {
		var ret WebinarPrivacy
		return ret
	}

	return o.Privacy
}

// GetPrivacyOk returns a tuple with the Privacy field value
// and a boolean to check if the value has been set.
func (o *Webinar) GetPrivacyOk() (*WebinarPrivacy, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Privacy, true
}

// SetPrivacy sets field value
func (o *Webinar) SetPrivacy(v WebinarPrivacy) {
	o.Privacy = v
}

// GetRegistrationData returns the RegistrationData field value
func (o *Webinar) GetRegistrationData() WebinarRegistrationData {
	if o == nil {
		var ret WebinarRegistrationData
		return ret
	}

	return o.RegistrationData
}

// GetRegistrationDataOk returns a tuple with the RegistrationData field value
// and a boolean to check if the value has been set.
func (o *Webinar) GetRegistrationDataOk() (*WebinarRegistrationData, bool) {
	if o == nil {
		return nil, false
	}
	return &o.RegistrationData, true
}

// SetRegistrationData sets field value
func (o *Webinar) SetRegistrationData(v WebinarRegistrationData) {
	o.RegistrationData = v
}

// GetRegistrationForm returns the RegistrationForm field value
func (o *Webinar) GetRegistrationForm() EmailCaptureForm {
	if o == nil {
		var ret EmailCaptureForm
		return ret
	}

	return o.RegistrationForm
}

// GetRegistrationFormOk returns a tuple with the RegistrationForm field value
// and a boolean to check if the value has been set.
func (o *Webinar) GetRegistrationFormOk() (*EmailCaptureForm, bool) {
	if o == nil {
		return nil, false
	}
	return &o.RegistrationForm, true
}

// SetRegistrationForm sets field value
func (o *Webinar) SetRegistrationForm(v EmailCaptureForm) {
	o.RegistrationForm = v
}

// GetSchedule returns the Schedule field value
// If the value is explicit nil, the zero value for WebinarSchedule will be returned
func (o *Webinar) GetSchedule() WebinarSchedule {
	if o == nil || o.Schedule.Get() == nil {
		var ret WebinarSchedule
		return ret
	}

	return *o.Schedule.Get()
}

// GetScheduleOk returns a tuple with the Schedule field value
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *Webinar) GetScheduleOk() (*WebinarSchedule, bool) {
	if o == nil {
		return nil, false
	}
	return o.Schedule.Get(), o.Schedule.IsSet()
}

// SetSchedule sets field value
func (o *Webinar) SetSchedule(v WebinarSchedule) {
	o.Schedule.Set(&v)
}

// GetStatus returns the Status field value
// If the value is explicit nil, the zero value for string will be returned
func (o *Webinar) GetStatus() string {
	if o == nil || o.Status.Get() == nil {
		var ret string
		return ret
	}

	return *o.Status.Get()
}

// GetStatusOk returns a tuple with the Status field value
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *Webinar) GetStatusOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.Status.Get(), o.Status.IsSet()
}

// SetStatus sets field value
func (o *Webinar) SetStatus(v string) {
	o.Status.Set(&v)
}

// GetTimeZone returns the TimeZone field value
func (o *Webinar) GetTimeZone() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.TimeZone
}

// GetTimeZoneOk returns a tuple with the TimeZone field value
// and a boolean to check if the value has been set.
func (o *Webinar) GetTimeZoneOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.TimeZone, true
}

// SetTimeZone sets field value
func (o *Webinar) SetTimeZone(v string) {
	o.TimeZone = v
}

// GetTitle returns the Title field value
// If the value is explicit nil, the zero value for string will be returned
func (o *Webinar) GetTitle() string {
	if o == nil || o.Title.Get() == nil {
		var ret string
		return ret
	}

	return *o.Title.Get()
}

// GetTitleOk returns a tuple with the Title field value
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *Webinar) GetTitleOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.Title.Get(), o.Title.IsSet()
}

// SetTitle sets field value
func (o *Webinar) SetTitle(v string) {
	o.Title.Set(&v)
}

// GetUri returns the Uri field value
func (o *Webinar) GetUri() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Uri
}

// GetUriOk returns a tuple with the Uri field value
// and a boolean to check if the value has been set.
func (o *Webinar) GetUriOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Uri, true
}

// SetUri sets field value
func (o *Webinar) SetUri(v string) {
	o.Uri = v
}

// GetUser returns the User field value
func (o *Webinar) GetUser() User {
	if o == nil {
		var ret User
		return ret
	}

	return o.User
}

// GetUserOk returns a tuple with the User field value
// and a boolean to check if the value has been set.
func (o *Webinar) GetUserOk() (*User, bool) {
	if o == nil {
		return nil, false
	}
	return &o.User, true
}

// SetUser sets field value
func (o *Webinar) SetUser(v User) {
	o.User = v
}

func (o Webinar) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o Webinar) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["completed_on"] = o.CompletedOn
	toSerialize["created_time"] = o.CreatedTime
	toSerialize["description"] = o.Description.Get()
	toSerialize["edit"] = o.Edit.Get()
	toSerialize["email_provider_list"] = o.EmailProviderList
	toSerialize["email_quota"] = o.EmailQuota
	toSerialize["email_settings"] = o.EmailSettings
	toSerialize["events"] = o.Events.Get()
	toSerialize["has_polls"] = o.HasPolls
	toSerialize["metadata"] = o.Metadata
	toSerialize["modified_on"] = o.ModifiedOn
	toSerialize["next_occurrence_time"] = o.NextOccurrenceTime.Get()
	toSerialize["password"] = o.Password.Get()
	toSerialize["privacy"] = o.Privacy
	toSerialize["registration_data"] = o.RegistrationData
	toSerialize["registration_form"] = o.RegistrationForm
	toSerialize["schedule"] = o.Schedule.Get()
	toSerialize["status"] = o.Status.Get()
	toSerialize["time_zone"] = o.TimeZone
	toSerialize["title"] = o.Title.Get()
	toSerialize["uri"] = o.Uri
	toSerialize["user"] = o.User
	return toSerialize, nil
}

type NullableWebinar struct {
	value *Webinar
	isSet bool
}

func (v NullableWebinar) Get() *Webinar {
	return v.value
}

func (v *NullableWebinar) Set(val *Webinar) {
	v.value = val
	v.isSet = true
}

func (v NullableWebinar) IsSet() bool {
	return v.isSet
}

func (v *NullableWebinar) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableWebinar(val *Webinar) *NullableWebinar {
	return &NullableWebinar{value: val, isSet: true}
}

func (v NullableWebinar) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableWebinar) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


