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

// checks if the UserUploadQuotaSpace type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &UserUploadQuotaSpace{}

// UserUploadQuotaSpace Information about the authenticated user's upload space remaining for the current period.
type UserUploadQuotaSpace struct {
	// The number of bytes or videos remaining in the authenticated user's upload quota.
	Free float32 `json:"free"`
	// The maximum number of bytes or videos allotted to the authenticated user's upload quota.
	Max NullableFloat32 `json:"max"`
	// The type of quota for the values of the **upload_quota.space** field.  Option descriptions:  * `lifetime` - The quota type is lifetime.  * `periodic` - The quota type is periodic. 
	Showing string `json:"showing"`
	// The unit that's used to compute quota.  Option descriptions:  * `video_count` - The quota is calculated using the count of the videos.  * `video_size` - The quota is calculated using the byte size of the videos. 
	Unit NullableString `json:"unit"`
	// The number of bytes or videos that the authenticated user has already uploaded against their quota.
	Used float32 `json:"used"`
}

// NewUserUploadQuotaSpace instantiates a new UserUploadQuotaSpace object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewUserUploadQuotaSpace(free float32, max NullableFloat32, showing string, unit NullableString, used float32) *UserUploadQuotaSpace {
	this := UserUploadQuotaSpace{}
	this.Free = free
	this.Max = max
	this.Showing = showing
	this.Unit = unit
	this.Used = used
	return &this
}

// NewUserUploadQuotaSpaceWithDefaults instantiates a new UserUploadQuotaSpace object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewUserUploadQuotaSpaceWithDefaults() *UserUploadQuotaSpace {
	this := UserUploadQuotaSpace{}
	return &this
}

// GetFree returns the Free field value
func (o *UserUploadQuotaSpace) GetFree() float32 {
	if o == nil {
		var ret float32
		return ret
	}

	return o.Free
}

// GetFreeOk returns a tuple with the Free field value
// and a boolean to check if the value has been set.
func (o *UserUploadQuotaSpace) GetFreeOk() (*float32, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Free, true
}

// SetFree sets field value
func (o *UserUploadQuotaSpace) SetFree(v float32) {
	o.Free = v
}

// GetMax returns the Max field value
// If the value is explicit nil, the zero value for float32 will be returned
func (o *UserUploadQuotaSpace) GetMax() float32 {
	if o == nil || o.Max.Get() == nil {
		var ret float32
		return ret
	}

	return *o.Max.Get()
}

// GetMaxOk returns a tuple with the Max field value
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *UserUploadQuotaSpace) GetMaxOk() (*float32, bool) {
	if o == nil {
		return nil, false
	}
	return o.Max.Get(), o.Max.IsSet()
}

// SetMax sets field value
func (o *UserUploadQuotaSpace) SetMax(v float32) {
	o.Max.Set(&v)
}

// GetShowing returns the Showing field value
func (o *UserUploadQuotaSpace) GetShowing() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Showing
}

// GetShowingOk returns a tuple with the Showing field value
// and a boolean to check if the value has been set.
func (o *UserUploadQuotaSpace) GetShowingOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Showing, true
}

// SetShowing sets field value
func (o *UserUploadQuotaSpace) SetShowing(v string) {
	o.Showing = v
}

// GetUnit returns the Unit field value
// If the value is explicit nil, the zero value for string will be returned
func (o *UserUploadQuotaSpace) GetUnit() string {
	if o == nil || o.Unit.Get() == nil {
		var ret string
		return ret
	}

	return *o.Unit.Get()
}

// GetUnitOk returns a tuple with the Unit field value
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *UserUploadQuotaSpace) GetUnitOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.Unit.Get(), o.Unit.IsSet()
}

// SetUnit sets field value
func (o *UserUploadQuotaSpace) SetUnit(v string) {
	o.Unit.Set(&v)
}

// GetUsed returns the Used field value
func (o *UserUploadQuotaSpace) GetUsed() float32 {
	if o == nil {
		var ret float32
		return ret
	}

	return o.Used
}

// GetUsedOk returns a tuple with the Used field value
// and a boolean to check if the value has been set.
func (o *UserUploadQuotaSpace) GetUsedOk() (*float32, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Used, true
}

// SetUsed sets field value
func (o *UserUploadQuotaSpace) SetUsed(v float32) {
	o.Used = v
}

func (o UserUploadQuotaSpace) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o UserUploadQuotaSpace) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["free"] = o.Free
	toSerialize["max"] = o.Max.Get()
	toSerialize["showing"] = o.Showing
	toSerialize["unit"] = o.Unit.Get()
	toSerialize["used"] = o.Used
	return toSerialize, nil
}

type NullableUserUploadQuotaSpace struct {
	value *UserUploadQuotaSpace
	isSet bool
}

func (v NullableUserUploadQuotaSpace) Get() *UserUploadQuotaSpace {
	return v.value
}

func (v *NullableUserUploadQuotaSpace) Set(val *UserUploadQuotaSpace) {
	v.value = val
	v.isSet = true
}

func (v NullableUserUploadQuotaSpace) IsSet() bool {
	return v.isSet
}

func (v *NullableUserUploadQuotaSpace) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableUserUploadQuotaSpace(val *UserUploadQuotaSpace) *NullableUserUploadQuotaSpace {
	return &NullableUserUploadQuotaSpace{value: val, isSet: true}
}

func (v NullableUserUploadQuotaSpace) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableUserUploadQuotaSpace) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


