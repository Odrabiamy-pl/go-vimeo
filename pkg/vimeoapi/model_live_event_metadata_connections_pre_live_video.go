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

// checks if the LiveEventMetadataConnectionsPreLiveVideo type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &LiveEventMetadataConnectionsPreLiveVideo{}

// LiveEventMetadataConnectionsPreLiveVideo Information about the event's pre-live video, where applicable. A pre-live video is either activated or in the process of being activated.
type LiveEventMetadataConnectionsPreLiveVideo struct {
	// An array of the HTTP methods permitted on this URI.
	Options []string `json:"options"`
	// The status of the pre-live video's RTMP link.  Option descriptions:  * `pending` - Vimeo is working on setting up the connection.  * `ready` - Resources have been provisioned for the event.  * `streaming` - Live video is currently streaming to the RTMP link.  * `unavailable` - The connection is ready, but streaming to the RTMP link has not yet begun. 
	Status string `json:"status"`
	// The API URI that resolves to the connection data.
	Uri string `json:"uri"`
}

// NewLiveEventMetadataConnectionsPreLiveVideo instantiates a new LiveEventMetadataConnectionsPreLiveVideo object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewLiveEventMetadataConnectionsPreLiveVideo(options []string, status string, uri string) *LiveEventMetadataConnectionsPreLiveVideo {
	this := LiveEventMetadataConnectionsPreLiveVideo{}
	this.Options = options
	this.Status = status
	this.Uri = uri
	return &this
}

// NewLiveEventMetadataConnectionsPreLiveVideoWithDefaults instantiates a new LiveEventMetadataConnectionsPreLiveVideo object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewLiveEventMetadataConnectionsPreLiveVideoWithDefaults() *LiveEventMetadataConnectionsPreLiveVideo {
	this := LiveEventMetadataConnectionsPreLiveVideo{}
	return &this
}

// GetOptions returns the Options field value
func (o *LiveEventMetadataConnectionsPreLiveVideo) GetOptions() []string {
	if o == nil {
		var ret []string
		return ret
	}

	return o.Options
}

// GetOptionsOk returns a tuple with the Options field value
// and a boolean to check if the value has been set.
func (o *LiveEventMetadataConnectionsPreLiveVideo) GetOptionsOk() ([]string, bool) {
	if o == nil {
		return nil, false
	}
	return o.Options, true
}

// SetOptions sets field value
func (o *LiveEventMetadataConnectionsPreLiveVideo) SetOptions(v []string) {
	o.Options = v
}

// GetStatus returns the Status field value
func (o *LiveEventMetadataConnectionsPreLiveVideo) GetStatus() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Status
}

// GetStatusOk returns a tuple with the Status field value
// and a boolean to check if the value has been set.
func (o *LiveEventMetadataConnectionsPreLiveVideo) GetStatusOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Status, true
}

// SetStatus sets field value
func (o *LiveEventMetadataConnectionsPreLiveVideo) SetStatus(v string) {
	o.Status = v
}

// GetUri returns the Uri field value
func (o *LiveEventMetadataConnectionsPreLiveVideo) GetUri() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Uri
}

// GetUriOk returns a tuple with the Uri field value
// and a boolean to check if the value has been set.
func (o *LiveEventMetadataConnectionsPreLiveVideo) GetUriOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Uri, true
}

// SetUri sets field value
func (o *LiveEventMetadataConnectionsPreLiveVideo) SetUri(v string) {
	o.Uri = v
}

func (o LiveEventMetadataConnectionsPreLiveVideo) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o LiveEventMetadataConnectionsPreLiveVideo) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["options"] = o.Options
	toSerialize["status"] = o.Status
	toSerialize["uri"] = o.Uri
	return toSerialize, nil
}

type NullableLiveEventMetadataConnectionsPreLiveVideo struct {
	value *LiveEventMetadataConnectionsPreLiveVideo
	isSet bool
}

func (v NullableLiveEventMetadataConnectionsPreLiveVideo) Get() *LiveEventMetadataConnectionsPreLiveVideo {
	return v.value
}

func (v *NullableLiveEventMetadataConnectionsPreLiveVideo) Set(val *LiveEventMetadataConnectionsPreLiveVideo) {
	v.value = val
	v.isSet = true
}

func (v NullableLiveEventMetadataConnectionsPreLiveVideo) IsSet() bool {
	return v.isSet
}

func (v *NullableLiveEventMetadataConnectionsPreLiveVideo) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableLiveEventMetadataConnectionsPreLiveVideo(val *LiveEventMetadataConnectionsPreLiveVideo) *NullableLiveEventMetadataConnectionsPreLiveVideo {
	return &NullableLiveEventMetadataConnectionsPreLiveVideo{value: val, isSet: true}
}

func (v NullableLiveEventMetadataConnectionsPreLiveVideo) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableLiveEventMetadataConnectionsPreLiveVideo) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


