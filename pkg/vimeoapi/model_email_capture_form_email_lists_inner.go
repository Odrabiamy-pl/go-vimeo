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

// checks if the EmailCaptureFormEmailListsInner type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &EmailCaptureFormEmailListsInner{}

// EmailCaptureFormEmailListsInner struct for EmailCaptureFormEmailListsInner
type EmailCaptureFormEmailListsInner struct {
	// The ID of the email capture form.
	FormId float32 `json:"form_id"`
	// The ID of the mailing list in the third-party email service provider's system.
	ListId string `json:"list_id"`
	// The name of the mailing list in the third-party email service provider's system.
	ListName *string `json:"list_name,omitempty"`
	// A third-party email service provider.  Option descriptions:  * `1` - The provider is Mailchimp.  * `2` - The provider is Campaign Monitor.  * `3` - The provider is Constant Contact.  * `4` - The provider is Infusionsoft.  * `5` - The provider is HubSpot.  * `6` - The provider is Constant Contact V3.  * `7` - The provider is Marketo. 
	ProviderId string `json:"provider_id"`
}

// NewEmailCaptureFormEmailListsInner instantiates a new EmailCaptureFormEmailListsInner object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewEmailCaptureFormEmailListsInner(formId float32, listId string, providerId string) *EmailCaptureFormEmailListsInner {
	this := EmailCaptureFormEmailListsInner{}
	this.FormId = formId
	this.ListId = listId
	this.ProviderId = providerId
	return &this
}

// NewEmailCaptureFormEmailListsInnerWithDefaults instantiates a new EmailCaptureFormEmailListsInner object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewEmailCaptureFormEmailListsInnerWithDefaults() *EmailCaptureFormEmailListsInner {
	this := EmailCaptureFormEmailListsInner{}
	return &this
}

// GetFormId returns the FormId field value
func (o *EmailCaptureFormEmailListsInner) GetFormId() float32 {
	if o == nil {
		var ret float32
		return ret
	}

	return o.FormId
}

// GetFormIdOk returns a tuple with the FormId field value
// and a boolean to check if the value has been set.
func (o *EmailCaptureFormEmailListsInner) GetFormIdOk() (*float32, bool) {
	if o == nil {
		return nil, false
	}
	return &o.FormId, true
}

// SetFormId sets field value
func (o *EmailCaptureFormEmailListsInner) SetFormId(v float32) {
	o.FormId = v
}

// GetListId returns the ListId field value
func (o *EmailCaptureFormEmailListsInner) GetListId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.ListId
}

// GetListIdOk returns a tuple with the ListId field value
// and a boolean to check if the value has been set.
func (o *EmailCaptureFormEmailListsInner) GetListIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ListId, true
}

// SetListId sets field value
func (o *EmailCaptureFormEmailListsInner) SetListId(v string) {
	o.ListId = v
}

// GetListName returns the ListName field value if set, zero value otherwise.
func (o *EmailCaptureFormEmailListsInner) GetListName() string {
	if o == nil || IsNil(o.ListName) {
		var ret string
		return ret
	}
	return *o.ListName
}

// GetListNameOk returns a tuple with the ListName field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *EmailCaptureFormEmailListsInner) GetListNameOk() (*string, bool) {
	if o == nil || IsNil(o.ListName) {
		return nil, false
	}
	return o.ListName, true
}

// HasListName returns a boolean if a field has been set.
func (o *EmailCaptureFormEmailListsInner) HasListName() bool {
	if o != nil && !IsNil(o.ListName) {
		return true
	}

	return false
}

// SetListName gets a reference to the given string and assigns it to the ListName field.
func (o *EmailCaptureFormEmailListsInner) SetListName(v string) {
	o.ListName = &v
}

// GetProviderId returns the ProviderId field value
func (o *EmailCaptureFormEmailListsInner) GetProviderId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.ProviderId
}

// GetProviderIdOk returns a tuple with the ProviderId field value
// and a boolean to check if the value has been set.
func (o *EmailCaptureFormEmailListsInner) GetProviderIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ProviderId, true
}

// SetProviderId sets field value
func (o *EmailCaptureFormEmailListsInner) SetProviderId(v string) {
	o.ProviderId = v
}

func (o EmailCaptureFormEmailListsInner) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o EmailCaptureFormEmailListsInner) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["form_id"] = o.FormId
	toSerialize["list_id"] = o.ListId
	if !IsNil(o.ListName) {
		toSerialize["list_name"] = o.ListName
	}
	toSerialize["provider_id"] = o.ProviderId
	return toSerialize, nil
}

type NullableEmailCaptureFormEmailListsInner struct {
	value *EmailCaptureFormEmailListsInner
	isSet bool
}

func (v NullableEmailCaptureFormEmailListsInner) Get() *EmailCaptureFormEmailListsInner {
	return v.value
}

func (v *NullableEmailCaptureFormEmailListsInner) Set(val *EmailCaptureFormEmailListsInner) {
	v.value = val
	v.isSet = true
}

func (v NullableEmailCaptureFormEmailListsInner) IsSet() bool {
	return v.isSet
}

func (v *NullableEmailCaptureFormEmailListsInner) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableEmailCaptureFormEmailListsInner(val *EmailCaptureFormEmailListsInner) *NullableEmailCaptureFormEmailListsInner {
	return &NullableEmailCaptureFormEmailListsInner{value: val, isSet: true}
}

func (v NullableEmailCaptureFormEmailListsInner) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableEmailCaptureFormEmailListsInner) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


