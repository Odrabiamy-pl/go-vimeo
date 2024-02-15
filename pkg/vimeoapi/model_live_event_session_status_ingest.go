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

// checks if the LiveEventSessionStatusIngest type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &LiveEventSessionStatusIngest{}

// LiveEventSessionStatusIngest The ingest of the video.
type LiveEventSessionStatusIngest struct {
	// The protocol used for this session.  Option descriptions:  * `dash` - The protocol is DASH.  * `rtmp` - The protocol is RTMP.  * `simple_live` - The protocol is Simplelive.  * `studio_cloud` - The protocol is StudioCloud.  * `unknown` - The protocol is unknown or not set.  * `webrtc` - The protocol is WebRTC. 
	EncoderType string `json:"encoder_type"`
	// The timestamp in UTC format when the live stream ended.
	EndTime NullableFloat32 `json:"end_time"`
	// The height of the live video in pixels.
	Height NullableFloat32 `json:"height"`
	// Whether the session is using RTMP.
	IsRtmpSession bool `json:"is_rtmp_session"`
	// Whether the stream is scheduled media playback.
	IsScheduledPlayback NullableBool `json:"is_scheduled_playback"`
	// The time in ISO 8601 format when the RTMP expires.
	RtmpExpiresAt NullableString `json:"rtmp_expires_at"`
	// The upstream RTMP link. Send your live content to this link to create a live video on the event.
	RtmpLink NullableString `json:"rtmp_link"`
	// The upstream RTMPS link. Send your live content to this secure link to create a live video on the event.
	RtmpsLink NullableString `json:"rtmps_link"`
	// The scheduled start time of the live video in ISO 8601 format.
	ScheduledStartTime NullableString `json:"scheduled_start_time"`
	// The session ID.
	SessionId NullableString `json:"session_id"`
	// The timestamp in UTC format when the live video started.
	StartTime NullableFloat32 `json:"start_time"`
	// The ingest status of the live video.  Option descriptions:  * `0` - There’s a live video, but no RMTP URL or key.  * `1` - There’s an RMTP URL and key, but the machine is provisioning.  * `2` - There’s an RMTP URL and key, and the machine is provisioned, but the stream hasn’t started yet.  * `3` - There’s an RMTP URL and key, and the machine is provisioned, but the stream didn’t start before the machine timed out.  * `4` - The stream has started and is currently underway.  * `5` - The stream has ended. 
	Status NullableString `json:"status"`
	// The reason why the stream ended. If the stream hasn't ended, this field is empty.
	StreamEndedReason NullableFloat32 `json:"stream_ended_reason"`
	// The stream key used in conjunction with the RTMP and RTMPS links.
	StreamKey NullableString `json:"stream_key"`
	// The width of the live video in pixels.
	Width NullableFloat32 `json:"width"`
}

type _LiveEventSessionStatusIngest LiveEventSessionStatusIngest

// NewLiveEventSessionStatusIngest instantiates a new LiveEventSessionStatusIngest object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewLiveEventSessionStatusIngest(encoderType string, endTime NullableFloat32, height NullableFloat32, isRtmpSession bool, isScheduledPlayback NullableBool, rtmpExpiresAt NullableString, rtmpLink NullableString, rtmpsLink NullableString, scheduledStartTime NullableString, sessionId NullableString, startTime NullableFloat32, status NullableString, streamEndedReason NullableFloat32, streamKey NullableString, width NullableFloat32) *LiveEventSessionStatusIngest {
	this := LiveEventSessionStatusIngest{}
	this.EncoderType = encoderType
	this.EndTime = endTime
	this.Height = height
	this.IsRtmpSession = isRtmpSession
	this.IsScheduledPlayback = isScheduledPlayback
	this.RtmpExpiresAt = rtmpExpiresAt
	this.RtmpLink = rtmpLink
	this.RtmpsLink = rtmpsLink
	this.ScheduledStartTime = scheduledStartTime
	this.SessionId = sessionId
	this.StartTime = startTime
	this.Status = status
	this.StreamEndedReason = streamEndedReason
	this.StreamKey = streamKey
	this.Width = width
	return &this
}

// NewLiveEventSessionStatusIngestWithDefaults instantiates a new LiveEventSessionStatusIngest object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewLiveEventSessionStatusIngestWithDefaults() *LiveEventSessionStatusIngest {
	this := LiveEventSessionStatusIngest{}
	return &this
}

// GetEncoderType returns the EncoderType field value
func (o *LiveEventSessionStatusIngest) GetEncoderType() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.EncoderType
}

// GetEncoderTypeOk returns a tuple with the EncoderType field value
// and a boolean to check if the value has been set.
func (o *LiveEventSessionStatusIngest) GetEncoderTypeOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.EncoderType, true
}

// SetEncoderType sets field value
func (o *LiveEventSessionStatusIngest) SetEncoderType(v string) {
	o.EncoderType = v
}

// GetEndTime returns the EndTime field value
// If the value is explicit nil, the zero value for float32 will be returned
func (o *LiveEventSessionStatusIngest) GetEndTime() float32 {
	if o == nil || o.EndTime.Get() == nil {
		var ret float32
		return ret
	}

	return *o.EndTime.Get()
}

// GetEndTimeOk returns a tuple with the EndTime field value
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *LiveEventSessionStatusIngest) GetEndTimeOk() (*float32, bool) {
	if o == nil {
		return nil, false
	}
	return o.EndTime.Get(), o.EndTime.IsSet()
}

// SetEndTime sets field value
func (o *LiveEventSessionStatusIngest) SetEndTime(v float32) {
	o.EndTime.Set(&v)
}

// GetHeight returns the Height field value
// If the value is explicit nil, the zero value for float32 will be returned
func (o *LiveEventSessionStatusIngest) GetHeight() float32 {
	if o == nil || o.Height.Get() == nil {
		var ret float32
		return ret
	}

	return *o.Height.Get()
}

// GetHeightOk returns a tuple with the Height field value
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *LiveEventSessionStatusIngest) GetHeightOk() (*float32, bool) {
	if o == nil {
		return nil, false
	}
	return o.Height.Get(), o.Height.IsSet()
}

// SetHeight sets field value
func (o *LiveEventSessionStatusIngest) SetHeight(v float32) {
	o.Height.Set(&v)
}

// GetIsRtmpSession returns the IsRtmpSession field value
func (o *LiveEventSessionStatusIngest) GetIsRtmpSession() bool {
	if o == nil {
		var ret bool
		return ret
	}

	return o.IsRtmpSession
}

// GetIsRtmpSessionOk returns a tuple with the IsRtmpSession field value
// and a boolean to check if the value has been set.
func (o *LiveEventSessionStatusIngest) GetIsRtmpSessionOk() (*bool, bool) {
	if o == nil {
		return nil, false
	}
	return &o.IsRtmpSession, true
}

// SetIsRtmpSession sets field value
func (o *LiveEventSessionStatusIngest) SetIsRtmpSession(v bool) {
	o.IsRtmpSession = v
}

// GetIsScheduledPlayback returns the IsScheduledPlayback field value
// If the value is explicit nil, the zero value for bool will be returned
func (o *LiveEventSessionStatusIngest) GetIsScheduledPlayback() bool {
	if o == nil || o.IsScheduledPlayback.Get() == nil {
		var ret bool
		return ret
	}

	return *o.IsScheduledPlayback.Get()
}

// GetIsScheduledPlaybackOk returns a tuple with the IsScheduledPlayback field value
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *LiveEventSessionStatusIngest) GetIsScheduledPlaybackOk() (*bool, bool) {
	if o == nil {
		return nil, false
	}
	return o.IsScheduledPlayback.Get(), o.IsScheduledPlayback.IsSet()
}

// SetIsScheduledPlayback sets field value
func (o *LiveEventSessionStatusIngest) SetIsScheduledPlayback(v bool) {
	o.IsScheduledPlayback.Set(&v)
}

// GetRtmpExpiresAt returns the RtmpExpiresAt field value
// If the value is explicit nil, the zero value for string will be returned
func (o *LiveEventSessionStatusIngest) GetRtmpExpiresAt() string {
	if o == nil || o.RtmpExpiresAt.Get() == nil {
		var ret string
		return ret
	}

	return *o.RtmpExpiresAt.Get()
}

// GetRtmpExpiresAtOk returns a tuple with the RtmpExpiresAt field value
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *LiveEventSessionStatusIngest) GetRtmpExpiresAtOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.RtmpExpiresAt.Get(), o.RtmpExpiresAt.IsSet()
}

// SetRtmpExpiresAt sets field value
func (o *LiveEventSessionStatusIngest) SetRtmpExpiresAt(v string) {
	o.RtmpExpiresAt.Set(&v)
}

// GetRtmpLink returns the RtmpLink field value
// If the value is explicit nil, the zero value for string will be returned
func (o *LiveEventSessionStatusIngest) GetRtmpLink() string {
	if o == nil || o.RtmpLink.Get() == nil {
		var ret string
		return ret
	}

	return *o.RtmpLink.Get()
}

// GetRtmpLinkOk returns a tuple with the RtmpLink field value
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *LiveEventSessionStatusIngest) GetRtmpLinkOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.RtmpLink.Get(), o.RtmpLink.IsSet()
}

// SetRtmpLink sets field value
func (o *LiveEventSessionStatusIngest) SetRtmpLink(v string) {
	o.RtmpLink.Set(&v)
}

// GetRtmpsLink returns the RtmpsLink field value
// If the value is explicit nil, the zero value for string will be returned
func (o *LiveEventSessionStatusIngest) GetRtmpsLink() string {
	if o == nil || o.RtmpsLink.Get() == nil {
		var ret string
		return ret
	}

	return *o.RtmpsLink.Get()
}

// GetRtmpsLinkOk returns a tuple with the RtmpsLink field value
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *LiveEventSessionStatusIngest) GetRtmpsLinkOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.RtmpsLink.Get(), o.RtmpsLink.IsSet()
}

// SetRtmpsLink sets field value
func (o *LiveEventSessionStatusIngest) SetRtmpsLink(v string) {
	o.RtmpsLink.Set(&v)
}

// GetScheduledStartTime returns the ScheduledStartTime field value
// If the value is explicit nil, the zero value for string will be returned
func (o *LiveEventSessionStatusIngest) GetScheduledStartTime() string {
	if o == nil || o.ScheduledStartTime.Get() == nil {
		var ret string
		return ret
	}

	return *o.ScheduledStartTime.Get()
}

// GetScheduledStartTimeOk returns a tuple with the ScheduledStartTime field value
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *LiveEventSessionStatusIngest) GetScheduledStartTimeOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.ScheduledStartTime.Get(), o.ScheduledStartTime.IsSet()
}

// SetScheduledStartTime sets field value
func (o *LiveEventSessionStatusIngest) SetScheduledStartTime(v string) {
	o.ScheduledStartTime.Set(&v)
}

// GetSessionId returns the SessionId field value
// If the value is explicit nil, the zero value for string will be returned
func (o *LiveEventSessionStatusIngest) GetSessionId() string {
	if o == nil || o.SessionId.Get() == nil {
		var ret string
		return ret
	}

	return *o.SessionId.Get()
}

// GetSessionIdOk returns a tuple with the SessionId field value
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *LiveEventSessionStatusIngest) GetSessionIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.SessionId.Get(), o.SessionId.IsSet()
}

// SetSessionId sets field value
func (o *LiveEventSessionStatusIngest) SetSessionId(v string) {
	o.SessionId.Set(&v)
}

// GetStartTime returns the StartTime field value
// If the value is explicit nil, the zero value for float32 will be returned
func (o *LiveEventSessionStatusIngest) GetStartTime() float32 {
	if o == nil || o.StartTime.Get() == nil {
		var ret float32
		return ret
	}

	return *o.StartTime.Get()
}

// GetStartTimeOk returns a tuple with the StartTime field value
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *LiveEventSessionStatusIngest) GetStartTimeOk() (*float32, bool) {
	if o == nil {
		return nil, false
	}
	return o.StartTime.Get(), o.StartTime.IsSet()
}

// SetStartTime sets field value
func (o *LiveEventSessionStatusIngest) SetStartTime(v float32) {
	o.StartTime.Set(&v)
}

// GetStatus returns the Status field value
// If the value is explicit nil, the zero value for string will be returned
func (o *LiveEventSessionStatusIngest) GetStatus() string {
	if o == nil || o.Status.Get() == nil {
		var ret string
		return ret
	}

	return *o.Status.Get()
}

// GetStatusOk returns a tuple with the Status field value
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *LiveEventSessionStatusIngest) GetStatusOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.Status.Get(), o.Status.IsSet()
}

// SetStatus sets field value
func (o *LiveEventSessionStatusIngest) SetStatus(v string) {
	o.Status.Set(&v)
}

// GetStreamEndedReason returns the StreamEndedReason field value
// If the value is explicit nil, the zero value for float32 will be returned
func (o *LiveEventSessionStatusIngest) GetStreamEndedReason() float32 {
	if o == nil || o.StreamEndedReason.Get() == nil {
		var ret float32
		return ret
	}

	return *o.StreamEndedReason.Get()
}

// GetStreamEndedReasonOk returns a tuple with the StreamEndedReason field value
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *LiveEventSessionStatusIngest) GetStreamEndedReasonOk() (*float32, bool) {
	if o == nil {
		return nil, false
	}
	return o.StreamEndedReason.Get(), o.StreamEndedReason.IsSet()
}

// SetStreamEndedReason sets field value
func (o *LiveEventSessionStatusIngest) SetStreamEndedReason(v float32) {
	o.StreamEndedReason.Set(&v)
}

// GetStreamKey returns the StreamKey field value
// If the value is explicit nil, the zero value for string will be returned
func (o *LiveEventSessionStatusIngest) GetStreamKey() string {
	if o == nil || o.StreamKey.Get() == nil {
		var ret string
		return ret
	}

	return *o.StreamKey.Get()
}

// GetStreamKeyOk returns a tuple with the StreamKey field value
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *LiveEventSessionStatusIngest) GetStreamKeyOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.StreamKey.Get(), o.StreamKey.IsSet()
}

// SetStreamKey sets field value
func (o *LiveEventSessionStatusIngest) SetStreamKey(v string) {
	o.StreamKey.Set(&v)
}

// GetWidth returns the Width field value
// If the value is explicit nil, the zero value for float32 will be returned
func (o *LiveEventSessionStatusIngest) GetWidth() float32 {
	if o == nil || o.Width.Get() == nil {
		var ret float32
		return ret
	}

	return *o.Width.Get()
}

// GetWidthOk returns a tuple with the Width field value
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *LiveEventSessionStatusIngest) GetWidthOk() (*float32, bool) {
	if o == nil {
		return nil, false
	}
	return o.Width.Get(), o.Width.IsSet()
}

// SetWidth sets field value
func (o *LiveEventSessionStatusIngest) SetWidth(v float32) {
	o.Width.Set(&v)
}

func (o LiveEventSessionStatusIngest) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o LiveEventSessionStatusIngest) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["encoder_type"] = o.EncoderType
	toSerialize["end_time"] = o.EndTime.Get()
	toSerialize["height"] = o.Height.Get()
	toSerialize["is_rtmp_session"] = o.IsRtmpSession
	toSerialize["is_scheduled_playback"] = o.IsScheduledPlayback.Get()
	toSerialize["rtmp_expires_at"] = o.RtmpExpiresAt.Get()
	toSerialize["rtmp_link"] = o.RtmpLink.Get()
	toSerialize["rtmps_link"] = o.RtmpsLink.Get()
	toSerialize["scheduled_start_time"] = o.ScheduledStartTime.Get()
	toSerialize["session_id"] = o.SessionId.Get()
	toSerialize["start_time"] = o.StartTime.Get()
	toSerialize["status"] = o.Status.Get()
	toSerialize["stream_ended_reason"] = o.StreamEndedReason.Get()
	toSerialize["stream_key"] = o.StreamKey.Get()
	toSerialize["width"] = o.Width.Get()
	return toSerialize, nil
}

func (o *LiveEventSessionStatusIngest) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"encoder_type",
		"end_time",
		"height",
		"is_rtmp_session",
		"is_scheduled_playback",
		"rtmp_expires_at",
		"rtmp_link",
		"rtmps_link",
		"scheduled_start_time",
		"session_id",
		"start_time",
		"status",
		"stream_ended_reason",
		"stream_key",
		"width",
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

	varLiveEventSessionStatusIngest := _LiveEventSessionStatusIngest{}

	decoder := json.NewDecoder(bytes.NewReader(data))
	decoder.DisallowUnknownFields()
	err = decoder.Decode(&varLiveEventSessionStatusIngest)

	if err != nil {
		return err
	}

	*o = LiveEventSessionStatusIngest(varLiveEventSessionStatusIngest)

	return err
}

type NullableLiveEventSessionStatusIngest struct {
	value *LiveEventSessionStatusIngest
	isSet bool
}

func (v NullableLiveEventSessionStatusIngest) Get() *LiveEventSessionStatusIngest {
	return v.value
}

func (v *NullableLiveEventSessionStatusIngest) Set(val *LiveEventSessionStatusIngest) {
	v.value = val
	v.isSet = true
}

func (v NullableLiveEventSessionStatusIngest) IsSet() bool {
	return v.isSet
}

func (v *NullableLiveEventSessionStatusIngest) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableLiveEventSessionStatusIngest(val *LiveEventSessionStatusIngest) *NullableLiveEventSessionStatusIngest {
	return &NullableLiveEventSessionStatusIngest{value: val, isSet: true}
}

func (v NullableLiveEventSessionStatusIngest) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableLiveEventSessionStatusIngest) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


