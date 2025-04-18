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

// checks if the ListPaymentMethodsRequest type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &ListPaymentMethodsRequest{}

// ListPaymentMethodsRequest struct for ListPaymentMethodsRequest
type ListPaymentMethodsRequest struct {
	// The type of payment method.  Option descriptions:  * `applepay` - The payment method is Apple Pay.  * `bank_account` - The payment method is a bank account.  * `card` - The payment method is a credit or debit card.  * `googlepay` - The payment method is Google Pay.  * `paypal` - The payment method is a PayPal account. 
	Type *string `json:"type,omitempty"`
}

// NewListPaymentMethodsRequest instantiates a new ListPaymentMethodsRequest object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewListPaymentMethodsRequest() *ListPaymentMethodsRequest {
	this := ListPaymentMethodsRequest{}
	return &this
}

// NewListPaymentMethodsRequestWithDefaults instantiates a new ListPaymentMethodsRequest object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewListPaymentMethodsRequestWithDefaults() *ListPaymentMethodsRequest {
	this := ListPaymentMethodsRequest{}
	return &this
}

// GetType returns the Type field value if set, zero value otherwise.
func (o *ListPaymentMethodsRequest) GetType() string {
	if o == nil || IsNil(o.Type) {
		var ret string
		return ret
	}
	return *o.Type
}

// GetTypeOk returns a tuple with the Type field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ListPaymentMethodsRequest) GetTypeOk() (*string, bool) {
	if o == nil || IsNil(o.Type) {
		return nil, false
	}
	return o.Type, true
}

// HasType returns a boolean if a field has been set.
func (o *ListPaymentMethodsRequest) HasType() bool {
	if o != nil && !IsNil(o.Type) {
		return true
	}

	return false
}

// SetType gets a reference to the given string and assigns it to the Type field.
func (o *ListPaymentMethodsRequest) SetType(v string) {
	o.Type = &v
}

func (o ListPaymentMethodsRequest) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o ListPaymentMethodsRequest) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Type) {
		toSerialize["type"] = o.Type
	}
	return toSerialize, nil
}

type NullableListPaymentMethodsRequest struct {
	value *ListPaymentMethodsRequest
	isSet bool
}

func (v NullableListPaymentMethodsRequest) Get() *ListPaymentMethodsRequest {
	return v.value
}

func (v *NullableListPaymentMethodsRequest) Set(val *ListPaymentMethodsRequest) {
	v.value = val
	v.isSet = true
}

func (v NullableListPaymentMethodsRequest) IsSet() bool {
	return v.isSet
}

func (v *NullableListPaymentMethodsRequest) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableListPaymentMethodsRequest(val *ListPaymentMethodsRequest) *NullableListPaymentMethodsRequest {
	return &NullableListPaymentMethodsRequest{value: val, isSet: true}
}

func (v NullableListPaymentMethodsRequest) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableListPaymentMethodsRequest) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


