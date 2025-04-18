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

// checks if the CreateVodPromotionRequest type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &CreateVodPromotionRequest{}

// CreateVodPromotionRequest struct for CreateVodPromotionRequest
type CreateVodPromotionRequest struct {
	// The promotion access type, which is a purchase option that isn't available in the On Demand container. Use the **download** and **stream_period** parameters to define additional characteristics for the `vip` type.  Option descriptions:  * `default` - The promotion grants a discount on the existing purchase options for an On Demand container.  * `vip` - The promotion grants free access to On Demand content before it's released. 
	AccessType *string `json:"access_type,omitempty"`
	// The promotion code. This parameter is ignored when the promotion type is `batch`.
	Code *string `json:"code,omitempty"`
	// The type of discount offered by the promotion code. When **access_type** is `vip`, the value of this parameter must be `free`.  Option descriptions:  * `free` - The discount reduces the price to zero.  * `percent` - The discount reduces the price by the percentage defined in the **percent_off** parameter. 
	DiscountType *string `json:"discount_type,omitempty"`
	// Whether the promotion grants download access to On Demand content. This field is required only when the download access hasn't been defined in the On Demand container, or when **access_type** is `vip` or **product_type** is `buy`.
	Download bool `json:"download"`
	// The time at which the promotion period ends. If this parameter has no value, the promotion never expires.
	EndTime *string `json:"end_time,omitempty"`
	// The description of the promotion when the promotion type is `batch`. This parameter is ignored when the promotion type is `single`.
	Label *string `json:"label,omitempty"`
	// The percentage of the discount. This parameter is applicable only when **discount_type** is `percent`.
	PercentOff *float32 `json:"percent_off,omitempty"`
	// The type of transaction to which the promotion applies. When **access_type** is `default`, the default value is `any`. When **access_type** is `vip`, the default value is `rent` and the only valid product types are `buy` and `rent`.  Option descriptions:  * `any` - The promotion applies to any transaction.  * `buy` - The promotion applies only to purchased products.  * `buy_episode` - The promotion applies only to purchased episodes.  * `rent` - The promotion applies only to rented products.  * `rent_episode` - The promotion applies only to rented episodes.  * `subscribe` - The promotion applies only to subscriptions. 
	ProductType *string `json:"product_type,omitempty"`
	// The time at which the promotion period starts. If this parameter has no value, the start time defaults to the time at which the promotion was created.
	StartTime *string `json:"start_time,omitempty"`
	// The amount of time for which the user can access On Demand content upon redeeming a promotion code. This parameter is required only when the streaming period isn't defined in the On Demand container, or when creating promotions where **access_type** is `vip` or **product_type** is `rent`.  Option descriptions:  * `1_week` - The user can access On Demand content for a maximum of 1 week after redeeming a promotion code.  * `1_year` - The user can access On Demand content for a maximum of 1 year after redeeming a promotion code.  * `24_hour` - The user can access On Demand content for a maximum of 24 hours after redeeming a promotion code.  * `30_day` - The user can access On Demand content for a maximum of 30 days after redeeming a promotion code.  * `3_month` - The user can access On Demand content for a maximum of 3 months after redeeming a promotion code.  * `48_hour` - The user can access On Demand content for a maximum of 48 hours after redeeming a promotion code.  * `6_month` - The user can access On Demand content for a maximum of 6 months after redeeming a promotion code.  * `72_hour` - The user can access On Demand content for a maximum of 72 hours after redeeming a promotion code. 
	StreamPeriod string `json:"stream_period"`
	// When **type** is `batch`, the total number of promotions to generate. When **type** is `single`, the total number of uses of the promotion.
	Total float32 `json:"total"`
	// The type of the promotion. When **access_type** is `vip`, the value for this parameter must be `batch`.  Option descriptions:  * `batch` - The promotion type that generates many random codes to use one time each.  * `single` - The promotion type that generates one code to use many times. 
	Type string `json:"type"`
}

// NewCreateVodPromotionRequest instantiates a new CreateVodPromotionRequest object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewCreateVodPromotionRequest(download bool, streamPeriod string, total float32, type_ string) *CreateVodPromotionRequest {
	this := CreateVodPromotionRequest{}
	this.Download = download
	this.StreamPeriod = streamPeriod
	this.Total = total
	this.Type = type_
	return &this
}

// NewCreateVodPromotionRequestWithDefaults instantiates a new CreateVodPromotionRequest object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewCreateVodPromotionRequestWithDefaults() *CreateVodPromotionRequest {
	this := CreateVodPromotionRequest{}
	return &this
}

// GetAccessType returns the AccessType field value if set, zero value otherwise.
func (o *CreateVodPromotionRequest) GetAccessType() string {
	if o == nil || IsNil(o.AccessType) {
		var ret string
		return ret
	}
	return *o.AccessType
}

// GetAccessTypeOk returns a tuple with the AccessType field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CreateVodPromotionRequest) GetAccessTypeOk() (*string, bool) {
	if o == nil || IsNil(o.AccessType) {
		return nil, false
	}
	return o.AccessType, true
}

// HasAccessType returns a boolean if a field has been set.
func (o *CreateVodPromotionRequest) HasAccessType() bool {
	if o != nil && !IsNil(o.AccessType) {
		return true
	}

	return false
}

// SetAccessType gets a reference to the given string and assigns it to the AccessType field.
func (o *CreateVodPromotionRequest) SetAccessType(v string) {
	o.AccessType = &v
}

// GetCode returns the Code field value if set, zero value otherwise.
func (o *CreateVodPromotionRequest) GetCode() string {
	if o == nil || IsNil(o.Code) {
		var ret string
		return ret
	}
	return *o.Code
}

// GetCodeOk returns a tuple with the Code field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CreateVodPromotionRequest) GetCodeOk() (*string, bool) {
	if o == nil || IsNil(o.Code) {
		return nil, false
	}
	return o.Code, true
}

// HasCode returns a boolean if a field has been set.
func (o *CreateVodPromotionRequest) HasCode() bool {
	if o != nil && !IsNil(o.Code) {
		return true
	}

	return false
}

// SetCode gets a reference to the given string and assigns it to the Code field.
func (o *CreateVodPromotionRequest) SetCode(v string) {
	o.Code = &v
}

// GetDiscountType returns the DiscountType field value if set, zero value otherwise.
func (o *CreateVodPromotionRequest) GetDiscountType() string {
	if o == nil || IsNil(o.DiscountType) {
		var ret string
		return ret
	}
	return *o.DiscountType
}

// GetDiscountTypeOk returns a tuple with the DiscountType field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CreateVodPromotionRequest) GetDiscountTypeOk() (*string, bool) {
	if o == nil || IsNil(o.DiscountType) {
		return nil, false
	}
	return o.DiscountType, true
}

// HasDiscountType returns a boolean if a field has been set.
func (o *CreateVodPromotionRequest) HasDiscountType() bool {
	if o != nil && !IsNil(o.DiscountType) {
		return true
	}

	return false
}

// SetDiscountType gets a reference to the given string and assigns it to the DiscountType field.
func (o *CreateVodPromotionRequest) SetDiscountType(v string) {
	o.DiscountType = &v
}

// GetDownload returns the Download field value
func (o *CreateVodPromotionRequest) GetDownload() bool {
	if o == nil {
		var ret bool
		return ret
	}

	return o.Download
}

// GetDownloadOk returns a tuple with the Download field value
// and a boolean to check if the value has been set.
func (o *CreateVodPromotionRequest) GetDownloadOk() (*bool, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Download, true
}

// SetDownload sets field value
func (o *CreateVodPromotionRequest) SetDownload(v bool) {
	o.Download = v
}

// GetEndTime returns the EndTime field value if set, zero value otherwise.
func (o *CreateVodPromotionRequest) GetEndTime() string {
	if o == nil || IsNil(o.EndTime) {
		var ret string
		return ret
	}
	return *o.EndTime
}

// GetEndTimeOk returns a tuple with the EndTime field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CreateVodPromotionRequest) GetEndTimeOk() (*string, bool) {
	if o == nil || IsNil(o.EndTime) {
		return nil, false
	}
	return o.EndTime, true
}

// HasEndTime returns a boolean if a field has been set.
func (o *CreateVodPromotionRequest) HasEndTime() bool {
	if o != nil && !IsNil(o.EndTime) {
		return true
	}

	return false
}

// SetEndTime gets a reference to the given string and assigns it to the EndTime field.
func (o *CreateVodPromotionRequest) SetEndTime(v string) {
	o.EndTime = &v
}

// GetLabel returns the Label field value if set, zero value otherwise.
func (o *CreateVodPromotionRequest) GetLabel() string {
	if o == nil || IsNil(o.Label) {
		var ret string
		return ret
	}
	return *o.Label
}

// GetLabelOk returns a tuple with the Label field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CreateVodPromotionRequest) GetLabelOk() (*string, bool) {
	if o == nil || IsNil(o.Label) {
		return nil, false
	}
	return o.Label, true
}

// HasLabel returns a boolean if a field has been set.
func (o *CreateVodPromotionRequest) HasLabel() bool {
	if o != nil && !IsNil(o.Label) {
		return true
	}

	return false
}

// SetLabel gets a reference to the given string and assigns it to the Label field.
func (o *CreateVodPromotionRequest) SetLabel(v string) {
	o.Label = &v
}

// GetPercentOff returns the PercentOff field value if set, zero value otherwise.
func (o *CreateVodPromotionRequest) GetPercentOff() float32 {
	if o == nil || IsNil(o.PercentOff) {
		var ret float32
		return ret
	}
	return *o.PercentOff
}

// GetPercentOffOk returns a tuple with the PercentOff field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CreateVodPromotionRequest) GetPercentOffOk() (*float32, bool) {
	if o == nil || IsNil(o.PercentOff) {
		return nil, false
	}
	return o.PercentOff, true
}

// HasPercentOff returns a boolean if a field has been set.
func (o *CreateVodPromotionRequest) HasPercentOff() bool {
	if o != nil && !IsNil(o.PercentOff) {
		return true
	}

	return false
}

// SetPercentOff gets a reference to the given float32 and assigns it to the PercentOff field.
func (o *CreateVodPromotionRequest) SetPercentOff(v float32) {
	o.PercentOff = &v
}

// GetProductType returns the ProductType field value if set, zero value otherwise.
func (o *CreateVodPromotionRequest) GetProductType() string {
	if o == nil || IsNil(o.ProductType) {
		var ret string
		return ret
	}
	return *o.ProductType
}

// GetProductTypeOk returns a tuple with the ProductType field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CreateVodPromotionRequest) GetProductTypeOk() (*string, bool) {
	if o == nil || IsNil(o.ProductType) {
		return nil, false
	}
	return o.ProductType, true
}

// HasProductType returns a boolean if a field has been set.
func (o *CreateVodPromotionRequest) HasProductType() bool {
	if o != nil && !IsNil(o.ProductType) {
		return true
	}

	return false
}

// SetProductType gets a reference to the given string and assigns it to the ProductType field.
func (o *CreateVodPromotionRequest) SetProductType(v string) {
	o.ProductType = &v
}

// GetStartTime returns the StartTime field value if set, zero value otherwise.
func (o *CreateVodPromotionRequest) GetStartTime() string {
	if o == nil || IsNil(o.StartTime) {
		var ret string
		return ret
	}
	return *o.StartTime
}

// GetStartTimeOk returns a tuple with the StartTime field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CreateVodPromotionRequest) GetStartTimeOk() (*string, bool) {
	if o == nil || IsNil(o.StartTime) {
		return nil, false
	}
	return o.StartTime, true
}

// HasStartTime returns a boolean if a field has been set.
func (o *CreateVodPromotionRequest) HasStartTime() bool {
	if o != nil && !IsNil(o.StartTime) {
		return true
	}

	return false
}

// SetStartTime gets a reference to the given string and assigns it to the StartTime field.
func (o *CreateVodPromotionRequest) SetStartTime(v string) {
	o.StartTime = &v
}

// GetStreamPeriod returns the StreamPeriod field value
func (o *CreateVodPromotionRequest) GetStreamPeriod() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.StreamPeriod
}

// GetStreamPeriodOk returns a tuple with the StreamPeriod field value
// and a boolean to check if the value has been set.
func (o *CreateVodPromotionRequest) GetStreamPeriodOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.StreamPeriod, true
}

// SetStreamPeriod sets field value
func (o *CreateVodPromotionRequest) SetStreamPeriod(v string) {
	o.StreamPeriod = v
}

// GetTotal returns the Total field value
func (o *CreateVodPromotionRequest) GetTotal() float32 {
	if o == nil {
		var ret float32
		return ret
	}

	return o.Total
}

// GetTotalOk returns a tuple with the Total field value
// and a boolean to check if the value has been set.
func (o *CreateVodPromotionRequest) GetTotalOk() (*float32, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Total, true
}

// SetTotal sets field value
func (o *CreateVodPromotionRequest) SetTotal(v float32) {
	o.Total = v
}

// GetType returns the Type field value
func (o *CreateVodPromotionRequest) GetType() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Type
}

// GetTypeOk returns a tuple with the Type field value
// and a boolean to check if the value has been set.
func (o *CreateVodPromotionRequest) GetTypeOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Type, true
}

// SetType sets field value
func (o *CreateVodPromotionRequest) SetType(v string) {
	o.Type = v
}

func (o CreateVodPromotionRequest) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o CreateVodPromotionRequest) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.AccessType) {
		toSerialize["access_type"] = o.AccessType
	}
	if !IsNil(o.Code) {
		toSerialize["code"] = o.Code
	}
	if !IsNil(o.DiscountType) {
		toSerialize["discount_type"] = o.DiscountType
	}
	toSerialize["download"] = o.Download
	if !IsNil(o.EndTime) {
		toSerialize["end_time"] = o.EndTime
	}
	if !IsNil(o.Label) {
		toSerialize["label"] = o.Label
	}
	if !IsNil(o.PercentOff) {
		toSerialize["percent_off"] = o.PercentOff
	}
	if !IsNil(o.ProductType) {
		toSerialize["product_type"] = o.ProductType
	}
	if !IsNil(o.StartTime) {
		toSerialize["start_time"] = o.StartTime
	}
	toSerialize["stream_period"] = o.StreamPeriod
	toSerialize["total"] = o.Total
	toSerialize["type"] = o.Type
	return toSerialize, nil
}

type NullableCreateVodPromotionRequest struct {
	value *CreateVodPromotionRequest
	isSet bool
}

func (v NullableCreateVodPromotionRequest) Get() *CreateVodPromotionRequest {
	return v.value
}

func (v *NullableCreateVodPromotionRequest) Set(val *CreateVodPromotionRequest) {
	v.value = val
	v.isSet = true
}

func (v NullableCreateVodPromotionRequest) IsSet() bool {
	return v.isSet
}

func (v *NullableCreateVodPromotionRequest) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableCreateVodPromotionRequest(val *CreateVodPromotionRequest) *NullableCreateVodPromotionRequest {
	return &NullableCreateVodPromotionRequest{value: val, isSet: true}
}

func (v NullableCreateVodPromotionRequest) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableCreateVodPromotionRequest) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


