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

// checks if the LiveEventSchedule type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &LiveEventSchedule{}

// LiveEventSchedule struct for LiveEventSchedule
type LiveEventSchedule struct {
	// When **schedule.type** is `weekly`, the time in ISO 8601 format when the event is expected to be live.
	DailyTime *string `json:"daily_time,omitempty"`
	// The time in ISO 8601 format when the live event is expected to end.
	EndTime *string `json:"end_time,omitempty"`
	// The recurrence rule for the event's schedule according to [RFC 5545](https://datatracker.ietf.org/doc/html/rfc5545).
	Rrule *string `json:"rrule,omitempty"`
	// When **schedule.type** is `weekly`, the time in ISO 8601 format when the first occurrence of the event is expected to be live.
	ScheduledTime *string `json:"scheduled_time,omitempty"`
	// The time in ISO 8601 format when the live event is expected to be live.
	StartTime *string `json:"start_time,omitempty"`
	// The time zone of the live event.
	TimeZone *string `json:"time_zone,omitempty"`
	// The schedule of the live event.  Option descriptions:  * `single` - The event is live only once.  * `weekly` - The event is live on a recurring weekly basis. 
	Type string `json:"type"`
	// When **schedule.type** is `weekly`, the weekdays in UTC when the event is expected to be live. The value of this field ranges from `1` to `7`, where `1` is Monday and `7` is Sunday.
	Weekdays []string `json:"weekdays,omitempty"`
}

// NewLiveEventSchedule instantiates a new LiveEventSchedule object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewLiveEventSchedule(type_ string) *LiveEventSchedule {
	this := LiveEventSchedule{}
	this.Type = type_
	return &this
}

// NewLiveEventScheduleWithDefaults instantiates a new LiveEventSchedule object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewLiveEventScheduleWithDefaults() *LiveEventSchedule {
	this := LiveEventSchedule{}
	return &this
}

// GetDailyTime returns the DailyTime field value if set, zero value otherwise.
func (o *LiveEventSchedule) GetDailyTime() string {
	if o == nil || IsNil(o.DailyTime) {
		var ret string
		return ret
	}
	return *o.DailyTime
}

// GetDailyTimeOk returns a tuple with the DailyTime field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *LiveEventSchedule) GetDailyTimeOk() (*string, bool) {
	if o == nil || IsNil(o.DailyTime) {
		return nil, false
	}
	return o.DailyTime, true
}

// HasDailyTime returns a boolean if a field has been set.
func (o *LiveEventSchedule) HasDailyTime() bool {
	if o != nil && !IsNil(o.DailyTime) {
		return true
	}

	return false
}

// SetDailyTime gets a reference to the given string and assigns it to the DailyTime field.
func (o *LiveEventSchedule) SetDailyTime(v string) {
	o.DailyTime = &v
}

// GetEndTime returns the EndTime field value if set, zero value otherwise.
func (o *LiveEventSchedule) GetEndTime() string {
	if o == nil || IsNil(o.EndTime) {
		var ret string
		return ret
	}
	return *o.EndTime
}

// GetEndTimeOk returns a tuple with the EndTime field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *LiveEventSchedule) GetEndTimeOk() (*string, bool) {
	if o == nil || IsNil(o.EndTime) {
		return nil, false
	}
	return o.EndTime, true
}

// HasEndTime returns a boolean if a field has been set.
func (o *LiveEventSchedule) HasEndTime() bool {
	if o != nil && !IsNil(o.EndTime) {
		return true
	}

	return false
}

// SetEndTime gets a reference to the given string and assigns it to the EndTime field.
func (o *LiveEventSchedule) SetEndTime(v string) {
	o.EndTime = &v
}

// GetRrule returns the Rrule field value if set, zero value otherwise.
func (o *LiveEventSchedule) GetRrule() string {
	if o == nil || IsNil(o.Rrule) {
		var ret string
		return ret
	}
	return *o.Rrule
}

// GetRruleOk returns a tuple with the Rrule field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *LiveEventSchedule) GetRruleOk() (*string, bool) {
	if o == nil || IsNil(o.Rrule) {
		return nil, false
	}
	return o.Rrule, true
}

// HasRrule returns a boolean if a field has been set.
func (o *LiveEventSchedule) HasRrule() bool {
	if o != nil && !IsNil(o.Rrule) {
		return true
	}

	return false
}

// SetRrule gets a reference to the given string and assigns it to the Rrule field.
func (o *LiveEventSchedule) SetRrule(v string) {
	o.Rrule = &v
}

// GetScheduledTime returns the ScheduledTime field value if set, zero value otherwise.
func (o *LiveEventSchedule) GetScheduledTime() string {
	if o == nil || IsNil(o.ScheduledTime) {
		var ret string
		return ret
	}
	return *o.ScheduledTime
}

// GetScheduledTimeOk returns a tuple with the ScheduledTime field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *LiveEventSchedule) GetScheduledTimeOk() (*string, bool) {
	if o == nil || IsNil(o.ScheduledTime) {
		return nil, false
	}
	return o.ScheduledTime, true
}

// HasScheduledTime returns a boolean if a field has been set.
func (o *LiveEventSchedule) HasScheduledTime() bool {
	if o != nil && !IsNil(o.ScheduledTime) {
		return true
	}

	return false
}

// SetScheduledTime gets a reference to the given string and assigns it to the ScheduledTime field.
func (o *LiveEventSchedule) SetScheduledTime(v string) {
	o.ScheduledTime = &v
}

// GetStartTime returns the StartTime field value if set, zero value otherwise.
func (o *LiveEventSchedule) GetStartTime() string {
	if o == nil || IsNil(o.StartTime) {
		var ret string
		return ret
	}
	return *o.StartTime
}

// GetStartTimeOk returns a tuple with the StartTime field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *LiveEventSchedule) GetStartTimeOk() (*string, bool) {
	if o == nil || IsNil(o.StartTime) {
		return nil, false
	}
	return o.StartTime, true
}

// HasStartTime returns a boolean if a field has been set.
func (o *LiveEventSchedule) HasStartTime() bool {
	if o != nil && !IsNil(o.StartTime) {
		return true
	}

	return false
}

// SetStartTime gets a reference to the given string and assigns it to the StartTime field.
func (o *LiveEventSchedule) SetStartTime(v string) {
	o.StartTime = &v
}

// GetTimeZone returns the TimeZone field value if set, zero value otherwise.
func (o *LiveEventSchedule) GetTimeZone() string {
	if o == nil || IsNil(o.TimeZone) {
		var ret string
		return ret
	}
	return *o.TimeZone
}

// GetTimeZoneOk returns a tuple with the TimeZone field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *LiveEventSchedule) GetTimeZoneOk() (*string, bool) {
	if o == nil || IsNil(o.TimeZone) {
		return nil, false
	}
	return o.TimeZone, true
}

// HasTimeZone returns a boolean if a field has been set.
func (o *LiveEventSchedule) HasTimeZone() bool {
	if o != nil && !IsNil(o.TimeZone) {
		return true
	}

	return false
}

// SetTimeZone gets a reference to the given string and assigns it to the TimeZone field.
func (o *LiveEventSchedule) SetTimeZone(v string) {
	o.TimeZone = &v
}

// GetType returns the Type field value
func (o *LiveEventSchedule) GetType() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Type
}

// GetTypeOk returns a tuple with the Type field value
// and a boolean to check if the value has been set.
func (o *LiveEventSchedule) GetTypeOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Type, true
}

// SetType sets field value
func (o *LiveEventSchedule) SetType(v string) {
	o.Type = v
}

// GetWeekdays returns the Weekdays field value if set, zero value otherwise.
func (o *LiveEventSchedule) GetWeekdays() []string {
	if o == nil || IsNil(o.Weekdays) {
		var ret []string
		return ret
	}
	return o.Weekdays
}

// GetWeekdaysOk returns a tuple with the Weekdays field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *LiveEventSchedule) GetWeekdaysOk() ([]string, bool) {
	if o == nil || IsNil(o.Weekdays) {
		return nil, false
	}
	return o.Weekdays, true
}

// HasWeekdays returns a boolean if a field has been set.
func (o *LiveEventSchedule) HasWeekdays() bool {
	if o != nil && !IsNil(o.Weekdays) {
		return true
	}

	return false
}

// SetWeekdays gets a reference to the given []string and assigns it to the Weekdays field.
func (o *LiveEventSchedule) SetWeekdays(v []string) {
	o.Weekdays = v
}

func (o LiveEventSchedule) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o LiveEventSchedule) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.DailyTime) {
		toSerialize["daily_time"] = o.DailyTime
	}
	if !IsNil(o.EndTime) {
		toSerialize["end_time"] = o.EndTime
	}
	if !IsNil(o.Rrule) {
		toSerialize["rrule"] = o.Rrule
	}
	if !IsNil(o.ScheduledTime) {
		toSerialize["scheduled_time"] = o.ScheduledTime
	}
	if !IsNil(o.StartTime) {
		toSerialize["start_time"] = o.StartTime
	}
	if !IsNil(o.TimeZone) {
		toSerialize["time_zone"] = o.TimeZone
	}
	toSerialize["type"] = o.Type
	if !IsNil(o.Weekdays) {
		toSerialize["weekdays"] = o.Weekdays
	}
	return toSerialize, nil
}

type NullableLiveEventSchedule struct {
	value *LiveEventSchedule
	isSet bool
}

func (v NullableLiveEventSchedule) Get() *LiveEventSchedule {
	return v.value
}

func (v *NullableLiveEventSchedule) Set(val *LiveEventSchedule) {
	v.value = val
	v.isSet = true
}

func (v NullableLiveEventSchedule) IsSet() bool {
	return v.isSet
}

func (v *NullableLiveEventSchedule) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableLiveEventSchedule(val *LiveEventSchedule) *NullableLiveEventSchedule {
	return &NullableLiveEventSchedule{value: val, isSet: true}
}

func (v NullableLiveEventSchedule) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableLiveEventSchedule) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


