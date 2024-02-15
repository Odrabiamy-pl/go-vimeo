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

// checks if the WebinarRegistrant type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &WebinarRegistrant{}

// WebinarRegistrant struct for WebinarRegistrant
type WebinarRegistrant struct {
	Analytics WebinarRegistrantAnalytics `json:"analytics"`
	// The date in Unix time when the registrant's account was created.
	CreatedOn float32 `json:"created_on"`
	// The values of all other fields as submitted on the webinar registration form.
	Data map[string]interface{} `json:"data"`
	// The registrant's email address as submitted on the webinar registration form.
	Email string `json:"email"`
	// The registrant's first name as submitted on the webinar registration form.
	FirstName NullableString `json:"first_name"`
	// The registrant's attended status for the webinar.  Option descriptions:  * `B` - The registrant has been blocked from attending the webinar.  * `N` - The registrant has not attended the webinar.  * `Y` - The registrant has attended the webinar. 
	HasAttended string `json:"has_attended"`
	// Whether the registrant's viewing status for the webinar is blocked.
	IsBlocked bool `json:"is_blocked"`
	// The registrant's last name as submitted on the webinar registration form.
	LastName NullableString `json:"last_name"`
	// The web address where the registration form was submitted.
	Referrer NullableString `json:"referrer"`
	// Details about the source from which the registrant's account was created.
	SourceDetails map[string]interface{} `json:"source_details"`
	// The source from which the registrant's account was created.
	SourceType NullableString `json:"source_type"`
	// The API URL to return the webinar registrant's account.
	Uri string `json:"uri"`
}

type _WebinarRegistrant WebinarRegistrant

// NewWebinarRegistrant instantiates a new WebinarRegistrant object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewWebinarRegistrant(analytics WebinarRegistrantAnalytics, createdOn float32, data map[string]interface{}, email string, firstName NullableString, hasAttended string, isBlocked bool, lastName NullableString, referrer NullableString, sourceDetails map[string]interface{}, sourceType NullableString, uri string) *WebinarRegistrant {
	this := WebinarRegistrant{}
	this.Analytics = analytics
	this.CreatedOn = createdOn
	this.Data = data
	this.Email = email
	this.FirstName = firstName
	this.HasAttended = hasAttended
	this.IsBlocked = isBlocked
	this.LastName = lastName
	this.Referrer = referrer
	this.SourceDetails = sourceDetails
	this.SourceType = sourceType
	this.Uri = uri
	return &this
}

// NewWebinarRegistrantWithDefaults instantiates a new WebinarRegistrant object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewWebinarRegistrantWithDefaults() *WebinarRegistrant {
	this := WebinarRegistrant{}
	return &this
}

// GetAnalytics returns the Analytics field value
func (o *WebinarRegistrant) GetAnalytics() WebinarRegistrantAnalytics {
	if o == nil {
		var ret WebinarRegistrantAnalytics
		return ret
	}

	return o.Analytics
}

// GetAnalyticsOk returns a tuple with the Analytics field value
// and a boolean to check if the value has been set.
func (o *WebinarRegistrant) GetAnalyticsOk() (*WebinarRegistrantAnalytics, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Analytics, true
}

// SetAnalytics sets field value
func (o *WebinarRegistrant) SetAnalytics(v WebinarRegistrantAnalytics) {
	o.Analytics = v
}

// GetCreatedOn returns the CreatedOn field value
func (o *WebinarRegistrant) GetCreatedOn() float32 {
	if o == nil {
		var ret float32
		return ret
	}

	return o.CreatedOn
}

// GetCreatedOnOk returns a tuple with the CreatedOn field value
// and a boolean to check if the value has been set.
func (o *WebinarRegistrant) GetCreatedOnOk() (*float32, bool) {
	if o == nil {
		return nil, false
	}
	return &o.CreatedOn, true
}

// SetCreatedOn sets field value
func (o *WebinarRegistrant) SetCreatedOn(v float32) {
	o.CreatedOn = v
}

// GetData returns the Data field value
// If the value is explicit nil, the zero value for map[string]interface{} will be returned
func (o *WebinarRegistrant) GetData() map[string]interface{} {
	if o == nil {
		var ret map[string]interface{}
		return ret
	}

	return o.Data
}

// GetDataOk returns a tuple with the Data field value
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *WebinarRegistrant) GetDataOk() (map[string]interface{}, bool) {
	if o == nil || IsNil(o.Data) {
		return map[string]interface{}{}, false
	}
	return o.Data, true
}

// SetData sets field value
func (o *WebinarRegistrant) SetData(v map[string]interface{}) {
	o.Data = v
}

// GetEmail returns the Email field value
func (o *WebinarRegistrant) GetEmail() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Email
}

// GetEmailOk returns a tuple with the Email field value
// and a boolean to check if the value has been set.
func (o *WebinarRegistrant) GetEmailOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Email, true
}

// SetEmail sets field value
func (o *WebinarRegistrant) SetEmail(v string) {
	o.Email = v
}

// GetFirstName returns the FirstName field value
// If the value is explicit nil, the zero value for string will be returned
func (o *WebinarRegistrant) GetFirstName() string {
	if o == nil || o.FirstName.Get() == nil {
		var ret string
		return ret
	}

	return *o.FirstName.Get()
}

// GetFirstNameOk returns a tuple with the FirstName field value
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *WebinarRegistrant) GetFirstNameOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.FirstName.Get(), o.FirstName.IsSet()
}

// SetFirstName sets field value
func (o *WebinarRegistrant) SetFirstName(v string) {
	o.FirstName.Set(&v)
}

// GetHasAttended returns the HasAttended field value
func (o *WebinarRegistrant) GetHasAttended() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.HasAttended
}

// GetHasAttendedOk returns a tuple with the HasAttended field value
// and a boolean to check if the value has been set.
func (o *WebinarRegistrant) GetHasAttendedOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.HasAttended, true
}

// SetHasAttended sets field value
func (o *WebinarRegistrant) SetHasAttended(v string) {
	o.HasAttended = v
}

// GetIsBlocked returns the IsBlocked field value
func (o *WebinarRegistrant) GetIsBlocked() bool {
	if o == nil {
		var ret bool
		return ret
	}

	return o.IsBlocked
}

// GetIsBlockedOk returns a tuple with the IsBlocked field value
// and a boolean to check if the value has been set.
func (o *WebinarRegistrant) GetIsBlockedOk() (*bool, bool) {
	if o == nil {
		return nil, false
	}
	return &o.IsBlocked, true
}

// SetIsBlocked sets field value
func (o *WebinarRegistrant) SetIsBlocked(v bool) {
	o.IsBlocked = v
}

// GetLastName returns the LastName field value
// If the value is explicit nil, the zero value for string will be returned
func (o *WebinarRegistrant) GetLastName() string {
	if o == nil || o.LastName.Get() == nil {
		var ret string
		return ret
	}

	return *o.LastName.Get()
}

// GetLastNameOk returns a tuple with the LastName field value
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *WebinarRegistrant) GetLastNameOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.LastName.Get(), o.LastName.IsSet()
}

// SetLastName sets field value
func (o *WebinarRegistrant) SetLastName(v string) {
	o.LastName.Set(&v)
}

// GetReferrer returns the Referrer field value
// If the value is explicit nil, the zero value for string will be returned
func (o *WebinarRegistrant) GetReferrer() string {
	if o == nil || o.Referrer.Get() == nil {
		var ret string
		return ret
	}

	return *o.Referrer.Get()
}

// GetReferrerOk returns a tuple with the Referrer field value
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *WebinarRegistrant) GetReferrerOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.Referrer.Get(), o.Referrer.IsSet()
}

// SetReferrer sets field value
func (o *WebinarRegistrant) SetReferrer(v string) {
	o.Referrer.Set(&v)
}

// GetSourceDetails returns the SourceDetails field value
// If the value is explicit nil, the zero value for map[string]interface{} will be returned
func (o *WebinarRegistrant) GetSourceDetails() map[string]interface{} {
	if o == nil {
		var ret map[string]interface{}
		return ret
	}

	return o.SourceDetails
}

// GetSourceDetailsOk returns a tuple with the SourceDetails field value
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *WebinarRegistrant) GetSourceDetailsOk() (map[string]interface{}, bool) {
	if o == nil || IsNil(o.SourceDetails) {
		return map[string]interface{}{}, false
	}
	return o.SourceDetails, true
}

// SetSourceDetails sets field value
func (o *WebinarRegistrant) SetSourceDetails(v map[string]interface{}) {
	o.SourceDetails = v
}

// GetSourceType returns the SourceType field value
// If the value is explicit nil, the zero value for string will be returned
func (o *WebinarRegistrant) GetSourceType() string {
	if o == nil || o.SourceType.Get() == nil {
		var ret string
		return ret
	}

	return *o.SourceType.Get()
}

// GetSourceTypeOk returns a tuple with the SourceType field value
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *WebinarRegistrant) GetSourceTypeOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.SourceType.Get(), o.SourceType.IsSet()
}

// SetSourceType sets field value
func (o *WebinarRegistrant) SetSourceType(v string) {
	o.SourceType.Set(&v)
}

// GetUri returns the Uri field value
func (o *WebinarRegistrant) GetUri() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Uri
}

// GetUriOk returns a tuple with the Uri field value
// and a boolean to check if the value has been set.
func (o *WebinarRegistrant) GetUriOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Uri, true
}

// SetUri sets field value
func (o *WebinarRegistrant) SetUri(v string) {
	o.Uri = v
}

func (o WebinarRegistrant) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o WebinarRegistrant) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["analytics"] = o.Analytics
	toSerialize["created_on"] = o.CreatedOn
	if o.Data != nil {
		toSerialize["data"] = o.Data
	}
	toSerialize["email"] = o.Email
	toSerialize["first_name"] = o.FirstName.Get()
	toSerialize["has_attended"] = o.HasAttended
	toSerialize["is_blocked"] = o.IsBlocked
	toSerialize["last_name"] = o.LastName.Get()
	toSerialize["referrer"] = o.Referrer.Get()
	if o.SourceDetails != nil {
		toSerialize["source_details"] = o.SourceDetails
	}
	toSerialize["source_type"] = o.SourceType.Get()
	toSerialize["uri"] = o.Uri
	return toSerialize, nil
}

func (o *WebinarRegistrant) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"analytics",
		"created_on",
		"data",
		"email",
		"first_name",
		"has_attended",
		"is_blocked",
		"last_name",
		"referrer",
		"source_details",
		"source_type",
		"uri",
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

	varWebinarRegistrant := _WebinarRegistrant{}

	decoder := json.NewDecoder(bytes.NewReader(data))
	decoder.DisallowUnknownFields()
	err = decoder.Decode(&varWebinarRegistrant)

	if err != nil {
		return err
	}

	*o = WebinarRegistrant(varWebinarRegistrant)

	return err
}

type NullableWebinarRegistrant struct {
	value *WebinarRegistrant
	isSet bool
}

func (v NullableWebinarRegistrant) Get() *WebinarRegistrant {
	return v.value
}

func (v *NullableWebinarRegistrant) Set(val *WebinarRegistrant) {
	v.value = val
	v.isSet = true
}

func (v NullableWebinarRegistrant) IsSet() bool {
	return v.isSet
}

func (v *NullableWebinarRegistrant) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableWebinarRegistrant(val *WebinarRegistrant) *NullableWebinarRegistrant {
	return &NullableWebinarRegistrant{value: val, isSet: true}
}

func (v NullableWebinarRegistrant) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableWebinarRegistrant) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

