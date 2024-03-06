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

// checks if the CreateLiveEventDestinationRequest type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &CreateLiveEventDestinationRequest{}

// CreateLiveEventDestinationRequest struct for CreateLiveEventDestinationRequest
type CreateLiveEventDestinationRequest struct {
	// The title to display for the simulcast.
	DisplayName string `json:"display_name"`
	// Whether the destination is enabled for simulcasting.
	IsEnabled *bool `json:"is_enabled,omitempty"`
	// The privacy setting of the destination. Be sure to choose a value that corresponds to your service.  Option descriptions:  * `CONNECTIONS` - The privacy setting is `CONNECTIONS` for LinkedIn.  * `PUBLIC` - The privacy setting is `PUBLIC` for LinkedIn.  * `all_friends` - The privacy setting is `all_friends` for Facebook.  * `everyone` - The privacy setting is `everyone` for Facebook.  * `private` - The privacy setting is `private` for YouTube.  * `public` - The privacy setting is `public` for YouTube.  * `self` - The privacy setting is `self` for Facebook.  * `unlisted` - The privacy setting is `unlisted` for YouTube.
	Privacy *string `json:"privacy,omitempty"`
	// The ID of the destination on the specified service, such as the YouTube channel ID or the Facebook page ID.
	ProviderDestinationId *string `json:"provider_destination_id,omitempty"`
	// The ID of the scheduled video.
	ProviderVideoId NullableString `json:"provider_video_id,omitempty"`
	// The time in Unix timestamp format when live streaming is scheduled to start.
	ScheduledAt *float32 `json:"scheduled_at,omitempty"`
	// The service to simulcast to.  Option descriptions:  * `custom_rtmp` - Simulcast to a custom service.  * `facebook` - Simulcast to Facebook Live.  * `linkedin` - Simulcast to LinkedIn Live.  * `youtube` - Simulcast to YouTube Live.
	ServiceName string `json:"service_name"`
	// The RTMP stream key.
	StreamKey *string `json:"stream_key,omitempty"`
	// The RTMP URL for receiving the video stream.
	StreamUrl *string `json:"stream_url,omitempty"`
	// The type of the simulcast destination.  Option descriptions:  * `channel` - The destination is a YouTube channel.  * `custom` - The destination is custom.  * `organization` - The destination is a LinkedIn organization.  * `page` - The destination is a Facebook page.  * `profile` - The destination is a Facebook or LinkedIn profile.
	Type string `json:"type"`
}

// NewCreateLiveEventDestinationRequest instantiates a new CreateLiveEventDestinationRequest object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewCreateLiveEventDestinationRequest(displayName string, serviceName string, type_ string) *CreateLiveEventDestinationRequest {
	this := CreateLiveEventDestinationRequest{}
	this.DisplayName = displayName
	this.ServiceName = serviceName
	this.Type = type_
	return &this
}

// NewCreateLiveEventDestinationRequestWithDefaults instantiates a new CreateLiveEventDestinationRequest object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewCreateLiveEventDestinationRequestWithDefaults() *CreateLiveEventDestinationRequest {
	this := CreateLiveEventDestinationRequest{}
	return &this
}

// GetDisplayName returns the DisplayName field value
func (o *CreateLiveEventDestinationRequest) GetDisplayName() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.DisplayName
}

// GetDisplayNameOk returns a tuple with the DisplayName field value
// and a boolean to check if the value has been set.
func (o *CreateLiveEventDestinationRequest) GetDisplayNameOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.DisplayName, true
}

// SetDisplayName sets field value
func (o *CreateLiveEventDestinationRequest) SetDisplayName(v string) {
	o.DisplayName = v
}

// GetIsEnabled returns the IsEnabled field value if set, zero value otherwise.
func (o *CreateLiveEventDestinationRequest) GetIsEnabled() bool {
	if o == nil || IsNil(o.IsEnabled) {
		var ret bool
		return ret
	}
	return *o.IsEnabled
}

// GetIsEnabledOk returns a tuple with the IsEnabled field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CreateLiveEventDestinationRequest) GetIsEnabledOk() (*bool, bool) {
	if o == nil || IsNil(o.IsEnabled) {
		return nil, false
	}
	return o.IsEnabled, true
}

// HasIsEnabled returns a boolean if a field has been set.
func (o *CreateLiveEventDestinationRequest) HasIsEnabled() bool {
	if o != nil && !IsNil(o.IsEnabled) {
		return true
	}

	return false
}

// SetIsEnabled gets a reference to the given bool and assigns it to the IsEnabled field.
func (o *CreateLiveEventDestinationRequest) SetIsEnabled(v bool) {
	o.IsEnabled = &v
}

// GetPrivacy returns the Privacy field value if set, zero value otherwise.
func (o *CreateLiveEventDestinationRequest) GetPrivacy() string {
	if o == nil || IsNil(o.Privacy) {
		var ret string
		return ret
	}
	return *o.Privacy
}

// GetPrivacyOk returns a tuple with the Privacy field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CreateLiveEventDestinationRequest) GetPrivacyOk() (*string, bool) {
	if o == nil || IsNil(o.Privacy) {
		return nil, false
	}
	return o.Privacy, true
}

// HasPrivacy returns a boolean if a field has been set.
func (o *CreateLiveEventDestinationRequest) HasPrivacy() bool {
	if o != nil && !IsNil(o.Privacy) {
		return true
	}

	return false
}

// SetPrivacy gets a reference to the given string and assigns it to the Privacy field.
func (o *CreateLiveEventDestinationRequest) SetPrivacy(v string) {
	o.Privacy = &v
}

// GetProviderDestinationId returns the ProviderDestinationId field value if set, zero value otherwise.
func (o *CreateLiveEventDestinationRequest) GetProviderDestinationId() string {
	if o == nil || IsNil(o.ProviderDestinationId) {
		var ret string
		return ret
	}
	return *o.ProviderDestinationId
}

// GetProviderDestinationIdOk returns a tuple with the ProviderDestinationId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CreateLiveEventDestinationRequest) GetProviderDestinationIdOk() (*string, bool) {
	if o == nil || IsNil(o.ProviderDestinationId) {
		return nil, false
	}
	return o.ProviderDestinationId, true
}

// HasProviderDestinationId returns a boolean if a field has been set.
func (o *CreateLiveEventDestinationRequest) HasProviderDestinationId() bool {
	if o != nil && !IsNil(o.ProviderDestinationId) {
		return true
	}

	return false
}

// SetProviderDestinationId gets a reference to the given string and assigns it to the ProviderDestinationId field.
func (o *CreateLiveEventDestinationRequest) SetProviderDestinationId(v string) {
	o.ProviderDestinationId = &v
}

// GetProviderVideoId returns the ProviderVideoId field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *CreateLiveEventDestinationRequest) GetProviderVideoId() string {
	if o == nil || IsNil(o.ProviderVideoId.Get()) {
		var ret string
		return ret
	}
	return *o.ProviderVideoId.Get()
}

// GetProviderVideoIdOk returns a tuple with the ProviderVideoId field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *CreateLiveEventDestinationRequest) GetProviderVideoIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.ProviderVideoId.Get(), o.ProviderVideoId.IsSet()
}

// HasProviderVideoId returns a boolean if a field has been set.
func (o *CreateLiveEventDestinationRequest) HasProviderVideoId() bool {
	if o != nil && o.ProviderVideoId.IsSet() {
		return true
	}

	return false
}

// SetProviderVideoId gets a reference to the given NullableString and assigns it to the ProviderVideoId field.
func (o *CreateLiveEventDestinationRequest) SetProviderVideoId(v string) {
	o.ProviderVideoId.Set(&v)
}

// SetProviderVideoIdNil sets the value for ProviderVideoId to be an explicit nil
func (o *CreateLiveEventDestinationRequest) SetProviderVideoIdNil() {
	o.ProviderVideoId.Set(nil)
}

// UnsetProviderVideoId ensures that no value is present for ProviderVideoId, not even an explicit nil
func (o *CreateLiveEventDestinationRequest) UnsetProviderVideoId() {
	o.ProviderVideoId.Unset()
}

// GetScheduledAt returns the ScheduledAt field value if set, zero value otherwise.
func (o *CreateLiveEventDestinationRequest) GetScheduledAt() float32 {
	if o == nil || IsNil(o.ScheduledAt) {
		var ret float32
		return ret
	}
	return *o.ScheduledAt
}

// GetScheduledAtOk returns a tuple with the ScheduledAt field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CreateLiveEventDestinationRequest) GetScheduledAtOk() (*float32, bool) {
	if o == nil || IsNil(o.ScheduledAt) {
		return nil, false
	}
	return o.ScheduledAt, true
}

// HasScheduledAt returns a boolean if a field has been set.
func (o *CreateLiveEventDestinationRequest) HasScheduledAt() bool {
	if o != nil && !IsNil(o.ScheduledAt) {
		return true
	}

	return false
}

// SetScheduledAt gets a reference to the given float32 and assigns it to the ScheduledAt field.
func (o *CreateLiveEventDestinationRequest) SetScheduledAt(v float32) {
	o.ScheduledAt = &v
}

// GetServiceName returns the ServiceName field value
func (o *CreateLiveEventDestinationRequest) GetServiceName() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.ServiceName
}

// GetServiceNameOk returns a tuple with the ServiceName field value
// and a boolean to check if the value has been set.
func (o *CreateLiveEventDestinationRequest) GetServiceNameOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ServiceName, true
}

// SetServiceName sets field value
func (o *CreateLiveEventDestinationRequest) SetServiceName(v string) {
	o.ServiceName = v
}

// GetStreamKey returns the StreamKey field value if set, zero value otherwise.
func (o *CreateLiveEventDestinationRequest) GetStreamKey() string {
	if o == nil || IsNil(o.StreamKey) {
		var ret string
		return ret
	}
	return *o.StreamKey
}

// GetStreamKeyOk returns a tuple with the StreamKey field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CreateLiveEventDestinationRequest) GetStreamKeyOk() (*string, bool) {
	if o == nil || IsNil(o.StreamKey) {
		return nil, false
	}
	return o.StreamKey, true
}

// HasStreamKey returns a boolean if a field has been set.
func (o *CreateLiveEventDestinationRequest) HasStreamKey() bool {
	if o != nil && !IsNil(o.StreamKey) {
		return true
	}

	return false
}

// SetStreamKey gets a reference to the given string and assigns it to the StreamKey field.
func (o *CreateLiveEventDestinationRequest) SetStreamKey(v string) {
	o.StreamKey = &v
}

// GetStreamUrl returns the StreamUrl field value if set, zero value otherwise.
func (o *CreateLiveEventDestinationRequest) GetStreamUrl() string {
	if o == nil || IsNil(o.StreamUrl) {
		var ret string
		return ret
	}
	return *o.StreamUrl
}

// GetStreamUrlOk returns a tuple with the StreamUrl field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CreateLiveEventDestinationRequest) GetStreamUrlOk() (*string, bool) {
	if o == nil || IsNil(o.StreamUrl) {
		return nil, false
	}
	return o.StreamUrl, true
}

// HasStreamUrl returns a boolean if a field has been set.
func (o *CreateLiveEventDestinationRequest) HasStreamUrl() bool {
	if o != nil && !IsNil(o.StreamUrl) {
		return true
	}

	return false
}

// SetStreamUrl gets a reference to the given string and assigns it to the StreamUrl field.
func (o *CreateLiveEventDestinationRequest) SetStreamUrl(v string) {
	o.StreamUrl = &v
}

// GetType returns the Type field value
func (o *CreateLiveEventDestinationRequest) GetType() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Type
}

// GetTypeOk returns a tuple with the Type field value
// and a boolean to check if the value has been set.
func (o *CreateLiveEventDestinationRequest) GetTypeOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Type, true
}

// SetType sets field value
func (o *CreateLiveEventDestinationRequest) SetType(v string) {
	o.Type = v
}

func (o CreateLiveEventDestinationRequest) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o CreateLiveEventDestinationRequest) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["display_name"] = o.DisplayName
	if !IsNil(o.IsEnabled) {
		toSerialize["is_enabled"] = o.IsEnabled
	}
	if !IsNil(o.Privacy) {
		toSerialize["privacy"] = o.Privacy
	}
	if !IsNil(o.ProviderDestinationId) {
		toSerialize["provider_destination_id"] = o.ProviderDestinationId
	}
	if o.ProviderVideoId.IsSet() {
		toSerialize["provider_video_id"] = o.ProviderVideoId.Get()
	}
	if !IsNil(o.ScheduledAt) {
		toSerialize["scheduled_at"] = o.ScheduledAt
	}
	toSerialize["service_name"] = o.ServiceName
	if !IsNil(o.StreamKey) {
		toSerialize["stream_key"] = o.StreamKey
	}
	if !IsNil(o.StreamUrl) {
		toSerialize["stream_url"] = o.StreamUrl
	}
	toSerialize["type"] = o.Type
	return toSerialize, nil
}

type NullableCreateLiveEventDestinationRequest struct {
	value *CreateLiveEventDestinationRequest
	isSet bool
}

func (v NullableCreateLiveEventDestinationRequest) Get() *CreateLiveEventDestinationRequest {
	return v.value
}

func (v *NullableCreateLiveEventDestinationRequest) Set(val *CreateLiveEventDestinationRequest) {
	v.value = val
	v.isSet = true
}

func (v NullableCreateLiveEventDestinationRequest) IsSet() bool {
	return v.isSet
}

func (v *NullableCreateLiveEventDestinationRequest) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableCreateLiveEventDestinationRequest(val *CreateLiveEventDestinationRequest) *NullableCreateLiveEventDestinationRequest {
	return &NullableCreateLiveEventDestinationRequest{value: val, isSet: true}
}

func (v NullableCreateLiveEventDestinationRequest) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableCreateLiveEventDestinationRequest) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
