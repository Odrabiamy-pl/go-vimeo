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

// checks if the PaymentMethodCard type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &PaymentMethodCard{}

// PaymentMethodCard Information about the card used to make the payment.
type PaymentMethodCard struct {
	BillingAddress *PaymentMethodCardBillingAddress `json:"billing_address,omitempty"`
	// The bank identification number of the card.
	Bin *string `json:"bin,omitempty"`
	// The brand of the card.
	Brand *string `json:"brand,omitempty"`
	// The name of the cardholder.
	CardholderName *string `json:"cardholder_name,omitempty"`
	// The expiration month of the card.
	ExpirationMonth *float32 `json:"expiration_month,omitempty"`
	// The expiration year of the card.
	ExpirationYear *float32 `json:"expiration_year,omitempty"`
	// The last four digits of the card.
	LastFourDigits *string `json:"last_four_digits,omitempty"`
}

// NewPaymentMethodCard instantiates a new PaymentMethodCard object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewPaymentMethodCard() *PaymentMethodCard {
	this := PaymentMethodCard{}
	return &this
}

// NewPaymentMethodCardWithDefaults instantiates a new PaymentMethodCard object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewPaymentMethodCardWithDefaults() *PaymentMethodCard {
	this := PaymentMethodCard{}
	return &this
}

// GetBillingAddress returns the BillingAddress field value if set, zero value otherwise.
func (o *PaymentMethodCard) GetBillingAddress() PaymentMethodCardBillingAddress {
	if o == nil || IsNil(o.BillingAddress) {
		var ret PaymentMethodCardBillingAddress
		return ret
	}
	return *o.BillingAddress
}

// GetBillingAddressOk returns a tuple with the BillingAddress field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PaymentMethodCard) GetBillingAddressOk() (*PaymentMethodCardBillingAddress, bool) {
	if o == nil || IsNil(o.BillingAddress) {
		return nil, false
	}
	return o.BillingAddress, true
}

// HasBillingAddress returns a boolean if a field has been set.
func (o *PaymentMethodCard) HasBillingAddress() bool {
	if o != nil && !IsNil(o.BillingAddress) {
		return true
	}

	return false
}

// SetBillingAddress gets a reference to the given PaymentMethodCardBillingAddress and assigns it to the BillingAddress field.
func (o *PaymentMethodCard) SetBillingAddress(v PaymentMethodCardBillingAddress) {
	o.BillingAddress = &v
}

// GetBin returns the Bin field value if set, zero value otherwise.
func (o *PaymentMethodCard) GetBin() string {
	if o == nil || IsNil(o.Bin) {
		var ret string
		return ret
	}
	return *o.Bin
}

// GetBinOk returns a tuple with the Bin field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PaymentMethodCard) GetBinOk() (*string, bool) {
	if o == nil || IsNil(o.Bin) {
		return nil, false
	}
	return o.Bin, true
}

// HasBin returns a boolean if a field has been set.
func (o *PaymentMethodCard) HasBin() bool {
	if o != nil && !IsNil(o.Bin) {
		return true
	}

	return false
}

// SetBin gets a reference to the given string and assigns it to the Bin field.
func (o *PaymentMethodCard) SetBin(v string) {
	o.Bin = &v
}

// GetBrand returns the Brand field value if set, zero value otherwise.
func (o *PaymentMethodCard) GetBrand() string {
	if o == nil || IsNil(o.Brand) {
		var ret string
		return ret
	}
	return *o.Brand
}

// GetBrandOk returns a tuple with the Brand field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PaymentMethodCard) GetBrandOk() (*string, bool) {
	if o == nil || IsNil(o.Brand) {
		return nil, false
	}
	return o.Brand, true
}

// HasBrand returns a boolean if a field has been set.
func (o *PaymentMethodCard) HasBrand() bool {
	if o != nil && !IsNil(o.Brand) {
		return true
	}

	return false
}

// SetBrand gets a reference to the given string and assigns it to the Brand field.
func (o *PaymentMethodCard) SetBrand(v string) {
	o.Brand = &v
}

// GetCardholderName returns the CardholderName field value if set, zero value otherwise.
func (o *PaymentMethodCard) GetCardholderName() string {
	if o == nil || IsNil(o.CardholderName) {
		var ret string
		return ret
	}
	return *o.CardholderName
}

// GetCardholderNameOk returns a tuple with the CardholderName field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PaymentMethodCard) GetCardholderNameOk() (*string, bool) {
	if o == nil || IsNil(o.CardholderName) {
		return nil, false
	}
	return o.CardholderName, true
}

// HasCardholderName returns a boolean if a field has been set.
func (o *PaymentMethodCard) HasCardholderName() bool {
	if o != nil && !IsNil(o.CardholderName) {
		return true
	}

	return false
}

// SetCardholderName gets a reference to the given string and assigns it to the CardholderName field.
func (o *PaymentMethodCard) SetCardholderName(v string) {
	o.CardholderName = &v
}

// GetExpirationMonth returns the ExpirationMonth field value if set, zero value otherwise.
func (o *PaymentMethodCard) GetExpirationMonth() float32 {
	if o == nil || IsNil(o.ExpirationMonth) {
		var ret float32
		return ret
	}
	return *o.ExpirationMonth
}

// GetExpirationMonthOk returns a tuple with the ExpirationMonth field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PaymentMethodCard) GetExpirationMonthOk() (*float32, bool) {
	if o == nil || IsNil(o.ExpirationMonth) {
		return nil, false
	}
	return o.ExpirationMonth, true
}

// HasExpirationMonth returns a boolean if a field has been set.
func (o *PaymentMethodCard) HasExpirationMonth() bool {
	if o != nil && !IsNil(o.ExpirationMonth) {
		return true
	}

	return false
}

// SetExpirationMonth gets a reference to the given float32 and assigns it to the ExpirationMonth field.
func (o *PaymentMethodCard) SetExpirationMonth(v float32) {
	o.ExpirationMonth = &v
}

// GetExpirationYear returns the ExpirationYear field value if set, zero value otherwise.
func (o *PaymentMethodCard) GetExpirationYear() float32 {
	if o == nil || IsNil(o.ExpirationYear) {
		var ret float32
		return ret
	}
	return *o.ExpirationYear
}

// GetExpirationYearOk returns a tuple with the ExpirationYear field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PaymentMethodCard) GetExpirationYearOk() (*float32, bool) {
	if o == nil || IsNil(o.ExpirationYear) {
		return nil, false
	}
	return o.ExpirationYear, true
}

// HasExpirationYear returns a boolean if a field has been set.
func (o *PaymentMethodCard) HasExpirationYear() bool {
	if o != nil && !IsNil(o.ExpirationYear) {
		return true
	}

	return false
}

// SetExpirationYear gets a reference to the given float32 and assigns it to the ExpirationYear field.
func (o *PaymentMethodCard) SetExpirationYear(v float32) {
	o.ExpirationYear = &v
}

// GetLastFourDigits returns the LastFourDigits field value if set, zero value otherwise.
func (o *PaymentMethodCard) GetLastFourDigits() string {
	if o == nil || IsNil(o.LastFourDigits) {
		var ret string
		return ret
	}
	return *o.LastFourDigits
}

// GetLastFourDigitsOk returns a tuple with the LastFourDigits field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PaymentMethodCard) GetLastFourDigitsOk() (*string, bool) {
	if o == nil || IsNil(o.LastFourDigits) {
		return nil, false
	}
	return o.LastFourDigits, true
}

// HasLastFourDigits returns a boolean if a field has been set.
func (o *PaymentMethodCard) HasLastFourDigits() bool {
	if o != nil && !IsNil(o.LastFourDigits) {
		return true
	}

	return false
}

// SetLastFourDigits gets a reference to the given string and assigns it to the LastFourDigits field.
func (o *PaymentMethodCard) SetLastFourDigits(v string) {
	o.LastFourDigits = &v
}

func (o PaymentMethodCard) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o PaymentMethodCard) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.BillingAddress) {
		toSerialize["billing_address"] = o.BillingAddress
	}
	if !IsNil(o.Bin) {
		toSerialize["bin"] = o.Bin
	}
	if !IsNil(o.Brand) {
		toSerialize["brand"] = o.Brand
	}
	if !IsNil(o.CardholderName) {
		toSerialize["cardholder_name"] = o.CardholderName
	}
	if !IsNil(o.ExpirationMonth) {
		toSerialize["expiration_month"] = o.ExpirationMonth
	}
	if !IsNil(o.ExpirationYear) {
		toSerialize["expiration_year"] = o.ExpirationYear
	}
	if !IsNil(o.LastFourDigits) {
		toSerialize["last_four_digits"] = o.LastFourDigits
	}
	return toSerialize, nil
}

type NullablePaymentMethodCard struct {
	value *PaymentMethodCard
	isSet bool
}

func (v NullablePaymentMethodCard) Get() *PaymentMethodCard {
	return v.value
}

func (v *NullablePaymentMethodCard) Set(val *PaymentMethodCard) {
	v.value = val
	v.isSet = true
}

func (v NullablePaymentMethodCard) IsSet() bool {
	return v.isSet
}

func (v *NullablePaymentMethodCard) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullablePaymentMethodCard(val *PaymentMethodCard) *NullablePaymentMethodCard {
	return &NullablePaymentMethodCard{value: val, isSet: true}
}

func (v NullablePaymentMethodCard) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullablePaymentMethodCard) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
