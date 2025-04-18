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

// checks if the UpdateWebinarAlt1RequestSchedule type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &UpdateWebinarAlt1RequestSchedule{}

// UpdateWebinarAlt1RequestSchedule Information about the time or times that the webinar is expected to be live. Please note that you can't update this setting once the webinar has started.
type UpdateWebinarAlt1RequestSchedule struct {
	// The time in ISO 8601 format when the webinar is expected to be live, with the zero UTC offset `Z`. This parameter is required when **schedule.type** is `weekly`. _This field is deprecated._
	DailyTime *string `json:"daily_time,omitempty"`
	// The time in ISO 8601 format when the webinar is expected to end, with support for different time offsets. This parameter is required when **schedule.type** is `single`.
	EndTime *string `json:"end_time,omitempty"`
	// The time in ISO 8601 format when the webinar is expected to be live, with support for different time offsets. This parameter is required when **schedule.type** is `single`.
	StartTime *string `json:"start_time,omitempty"`
	// How often the webinar is expected to be live.  Option descriptions:  * `single` - The webinar is live one time only.  * `weekly` - The webinar is live on a weekly basis. _This field is deprecated._ 
	Type *string `json:"type,omitempty"`
	// A non-empty array of weekdays on which the webinar is expected to be live. Weekdays can range from 1 to 7, where 1 is Monday and 7 is Sunday. This parameter is required when **schedule.type** is `weekly`. _This field is deprecated._
	Weekdays []string `json:"weekdays,omitempty"`
}

// NewUpdateWebinarAlt1RequestSchedule instantiates a new UpdateWebinarAlt1RequestSchedule object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewUpdateWebinarAlt1RequestSchedule() *UpdateWebinarAlt1RequestSchedule {
	this := UpdateWebinarAlt1RequestSchedule{}
	return &this
}

// NewUpdateWebinarAlt1RequestScheduleWithDefaults instantiates a new UpdateWebinarAlt1RequestSchedule object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewUpdateWebinarAlt1RequestScheduleWithDefaults() *UpdateWebinarAlt1RequestSchedule {
	this := UpdateWebinarAlt1RequestSchedule{}
	return &this
}

// GetDailyTime returns the DailyTime field value if set, zero value otherwise.
func (o *UpdateWebinarAlt1RequestSchedule) GetDailyTime() string {
	if o == nil || IsNil(o.DailyTime) {
		var ret string
		return ret
	}
	return *o.DailyTime
}

// GetDailyTimeOk returns a tuple with the DailyTime field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UpdateWebinarAlt1RequestSchedule) GetDailyTimeOk() (*string, bool) {
	if o == nil || IsNil(o.DailyTime) {
		return nil, false
	}
	return o.DailyTime, true
}

// HasDailyTime returns a boolean if a field has been set.
func (o *UpdateWebinarAlt1RequestSchedule) HasDailyTime() bool {
	if o != nil && !IsNil(o.DailyTime) {
		return true
	}

	return false
}

// SetDailyTime gets a reference to the given string and assigns it to the DailyTime field.
func (o *UpdateWebinarAlt1RequestSchedule) SetDailyTime(v string) {
	o.DailyTime = &v
}

// GetEndTime returns the EndTime field value if set, zero value otherwise.
func (o *UpdateWebinarAlt1RequestSchedule) GetEndTime() string {
	if o == nil || IsNil(o.EndTime) {
		var ret string
		return ret
	}
	return *o.EndTime
}

// GetEndTimeOk returns a tuple with the EndTime field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UpdateWebinarAlt1RequestSchedule) GetEndTimeOk() (*string, bool) {
	if o == nil || IsNil(o.EndTime) {
		return nil, false
	}
	return o.EndTime, true
}

// HasEndTime returns a boolean if a field has been set.
func (o *UpdateWebinarAlt1RequestSchedule) HasEndTime() bool {
	if o != nil && !IsNil(o.EndTime) {
		return true
	}

	return false
}

// SetEndTime gets a reference to the given string and assigns it to the EndTime field.
func (o *UpdateWebinarAlt1RequestSchedule) SetEndTime(v string) {
	o.EndTime = &v
}

// GetStartTime returns the StartTime field value if set, zero value otherwise.
func (o *UpdateWebinarAlt1RequestSchedule) GetStartTime() string {
	if o == nil || IsNil(o.StartTime) {
		var ret string
		return ret
	}
	return *o.StartTime
}

// GetStartTimeOk returns a tuple with the StartTime field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UpdateWebinarAlt1RequestSchedule) GetStartTimeOk() (*string, bool) {
	if o == nil || IsNil(o.StartTime) {
		return nil, false
	}
	return o.StartTime, true
}

// HasStartTime returns a boolean if a field has been set.
func (o *UpdateWebinarAlt1RequestSchedule) HasStartTime() bool {
	if o != nil && !IsNil(o.StartTime) {
		return true
	}

	return false
}

// SetStartTime gets a reference to the given string and assigns it to the StartTime field.
func (o *UpdateWebinarAlt1RequestSchedule) SetStartTime(v string) {
	o.StartTime = &v
}

// GetType returns the Type field value if set, zero value otherwise.
func (o *UpdateWebinarAlt1RequestSchedule) GetType() string {
	if o == nil || IsNil(o.Type) {
		var ret string
		return ret
	}
	return *o.Type
}

// GetTypeOk returns a tuple with the Type field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UpdateWebinarAlt1RequestSchedule) GetTypeOk() (*string, bool) {
	if o == nil || IsNil(o.Type) {
		return nil, false
	}
	return o.Type, true
}

// HasType returns a boolean if a field has been set.
func (o *UpdateWebinarAlt1RequestSchedule) HasType() bool {
	if o != nil && !IsNil(o.Type) {
		return true
	}

	return false
}

// SetType gets a reference to the given string and assigns it to the Type field.
func (o *UpdateWebinarAlt1RequestSchedule) SetType(v string) {
	o.Type = &v
}

// GetWeekdays returns the Weekdays field value if set, zero value otherwise.
func (o *UpdateWebinarAlt1RequestSchedule) GetWeekdays() []string {
	if o == nil || IsNil(o.Weekdays) {
		var ret []string
		return ret
	}
	return o.Weekdays
}

// GetWeekdaysOk returns a tuple with the Weekdays field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UpdateWebinarAlt1RequestSchedule) GetWeekdaysOk() ([]string, bool) {
	if o == nil || IsNil(o.Weekdays) {
		return nil, false
	}
	return o.Weekdays, true
}

// HasWeekdays returns a boolean if a field has been set.
func (o *UpdateWebinarAlt1RequestSchedule) HasWeekdays() bool {
	if o != nil && !IsNil(o.Weekdays) {
		return true
	}

	return false
}

// SetWeekdays gets a reference to the given []string and assigns it to the Weekdays field.
func (o *UpdateWebinarAlt1RequestSchedule) SetWeekdays(v []string) {
	o.Weekdays = v
}

func (o UpdateWebinarAlt1RequestSchedule) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o UpdateWebinarAlt1RequestSchedule) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.DailyTime) {
		toSerialize["daily_time"] = o.DailyTime
	}
	if !IsNil(o.EndTime) {
		toSerialize["end_time"] = o.EndTime
	}
	if !IsNil(o.StartTime) {
		toSerialize["start_time"] = o.StartTime
	}
	if !IsNil(o.Type) {
		toSerialize["type"] = o.Type
	}
	if !IsNil(o.Weekdays) {
		toSerialize["weekdays"] = o.Weekdays
	}
	return toSerialize, nil
}

type NullableUpdateWebinarAlt1RequestSchedule struct {
	value *UpdateWebinarAlt1RequestSchedule
	isSet bool
}

func (v NullableUpdateWebinarAlt1RequestSchedule) Get() *UpdateWebinarAlt1RequestSchedule {
	return v.value
}

func (v *NullableUpdateWebinarAlt1RequestSchedule) Set(val *UpdateWebinarAlt1RequestSchedule) {
	v.value = val
	v.isSet = true
}

func (v NullableUpdateWebinarAlt1RequestSchedule) IsSet() bool {
	return v.isSet
}

func (v *NullableUpdateWebinarAlt1RequestSchedule) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableUpdateWebinarAlt1RequestSchedule(val *UpdateWebinarAlt1RequestSchedule) *NullableUpdateWebinarAlt1RequestSchedule {
	return &NullableUpdateWebinarAlt1RequestSchedule{value: val, isSet: true}
}

func (v NullableUpdateWebinarAlt1RequestSchedule) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableUpdateWebinarAlt1RequestSchedule) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


