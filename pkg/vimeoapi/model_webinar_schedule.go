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

// checks if the WebinarSchedule type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &WebinarSchedule{}

// WebinarSchedule The description of the time or times that the webinar is expected to be live.
type WebinarSchedule struct {
	// The time in ISO 8601 format at which the webinar is expected to be live when **schedule.type** is `weekly`.
	DailyTime string `json:"daily_time"`
	// The date in ISO 8601 format on which the webinar is expected to end. This field applies when **schedule.type** is `single`.
	EndTime string `json:"end_time"`
	// The date in ISO 8601 format on which the first occurrence of the webinar is expected to be live when **schedule.type** is `weekly`.
	ScheduledTime NullableString `json:"scheduled_time"`
	// The date in ISO 8601 format on which the webinar is expected to be live when **schedule.type** is `single`.
	StartTime string `json:"start_time"`
	// The schedule of the webinar.  Option descriptions:  * `single` - The webinar is live only once.  * `weekly` - The webinar is live on a recurring weekly basis. 
	Type string `json:"type"`
	// The weekdays in UTC on which the webinar is expected to be live when **schedule.time** is `weekly`. The value of this field ranges from `1` to `7`, where `1` is Monday and `7` is Sunday.
	Weekdays []string `json:"weekdays"`
}

type _WebinarSchedule WebinarSchedule

// NewWebinarSchedule instantiates a new WebinarSchedule object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewWebinarSchedule(dailyTime string, endTime string, scheduledTime NullableString, startTime string, type_ string, weekdays []string) *WebinarSchedule {
	this := WebinarSchedule{}
	this.DailyTime = dailyTime
	this.EndTime = endTime
	this.ScheduledTime = scheduledTime
	this.StartTime = startTime
	this.Type = type_
	this.Weekdays = weekdays
	return &this
}

// NewWebinarScheduleWithDefaults instantiates a new WebinarSchedule object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewWebinarScheduleWithDefaults() *WebinarSchedule {
	this := WebinarSchedule{}
	return &this
}

// GetDailyTime returns the DailyTime field value
func (o *WebinarSchedule) GetDailyTime() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.DailyTime
}

// GetDailyTimeOk returns a tuple with the DailyTime field value
// and a boolean to check if the value has been set.
func (o *WebinarSchedule) GetDailyTimeOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.DailyTime, true
}

// SetDailyTime sets field value
func (o *WebinarSchedule) SetDailyTime(v string) {
	o.DailyTime = v
}

// GetEndTime returns the EndTime field value
func (o *WebinarSchedule) GetEndTime() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.EndTime
}

// GetEndTimeOk returns a tuple with the EndTime field value
// and a boolean to check if the value has been set.
func (o *WebinarSchedule) GetEndTimeOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.EndTime, true
}

// SetEndTime sets field value
func (o *WebinarSchedule) SetEndTime(v string) {
	o.EndTime = v
}

// GetScheduledTime returns the ScheduledTime field value
// If the value is explicit nil, the zero value for string will be returned
func (o *WebinarSchedule) GetScheduledTime() string {
	if o == nil || o.ScheduledTime.Get() == nil {
		var ret string
		return ret
	}

	return *o.ScheduledTime.Get()
}

// GetScheduledTimeOk returns a tuple with the ScheduledTime field value
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *WebinarSchedule) GetScheduledTimeOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.ScheduledTime.Get(), o.ScheduledTime.IsSet()
}

// SetScheduledTime sets field value
func (o *WebinarSchedule) SetScheduledTime(v string) {
	o.ScheduledTime.Set(&v)
}

// GetStartTime returns the StartTime field value
func (o *WebinarSchedule) GetStartTime() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.StartTime
}

// GetStartTimeOk returns a tuple with the StartTime field value
// and a boolean to check if the value has been set.
func (o *WebinarSchedule) GetStartTimeOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.StartTime, true
}

// SetStartTime sets field value
func (o *WebinarSchedule) SetStartTime(v string) {
	o.StartTime = v
}

// GetType returns the Type field value
func (o *WebinarSchedule) GetType() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Type
}

// GetTypeOk returns a tuple with the Type field value
// and a boolean to check if the value has been set.
func (o *WebinarSchedule) GetTypeOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Type, true
}

// SetType sets field value
func (o *WebinarSchedule) SetType(v string) {
	o.Type = v
}

// GetWeekdays returns the Weekdays field value
func (o *WebinarSchedule) GetWeekdays() []string {
	if o == nil {
		var ret []string
		return ret
	}

	return o.Weekdays
}

// GetWeekdaysOk returns a tuple with the Weekdays field value
// and a boolean to check if the value has been set.
func (o *WebinarSchedule) GetWeekdaysOk() ([]string, bool) {
	if o == nil {
		return nil, false
	}
	return o.Weekdays, true
}

// SetWeekdays sets field value
func (o *WebinarSchedule) SetWeekdays(v []string) {
	o.Weekdays = v
}

func (o WebinarSchedule) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o WebinarSchedule) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["daily_time"] = o.DailyTime
	toSerialize["end_time"] = o.EndTime
	toSerialize["scheduled_time"] = o.ScheduledTime.Get()
	toSerialize["start_time"] = o.StartTime
	toSerialize["type"] = o.Type
	toSerialize["weekdays"] = o.Weekdays
	return toSerialize, nil
}

func (o *WebinarSchedule) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"daily_time",
		"end_time",
		"scheduled_time",
		"start_time",
		"type",
		"weekdays",
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

	varWebinarSchedule := _WebinarSchedule{}

	decoder := json.NewDecoder(bytes.NewReader(data))
	decoder.DisallowUnknownFields()
	err = decoder.Decode(&varWebinarSchedule)

	if err != nil {
		return err
	}

	*o = WebinarSchedule(varWebinarSchedule)

	return err
}

type NullableWebinarSchedule struct {
	value *WebinarSchedule
	isSet bool
}

func (v NullableWebinarSchedule) Get() *WebinarSchedule {
	return v.value
}

func (v *NullableWebinarSchedule) Set(val *WebinarSchedule) {
	v.value = val
	v.isSet = true
}

func (v NullableWebinarSchedule) IsSet() bool {
	return v.isSet
}

func (v *NullableWebinarSchedule) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableWebinarSchedule(val *WebinarSchedule) *NullableWebinarSchedule {
	return &NullableWebinarSchedule{value: val, isSet: true}
}

func (v NullableWebinarSchedule) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableWebinarSchedule) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


