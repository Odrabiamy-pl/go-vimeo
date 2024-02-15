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

// checks if the SubscriptionPlansMetadataInteractionsPurchase type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &SubscriptionPlansMetadataInteractionsPurchase{}

// SubscriptionPlansMetadataInteractionsPurchase struct for SubscriptionPlansMetadataInteractionsPurchase
type SubscriptionPlansMetadataInteractionsPurchase struct {
	// The purchase status of the product.  Option descriptions:  * `available` - The product is available for purchase.  * `purchased` - The product is already purchased.  * `unavailable` - The product isn't available for purchase. 
	Status string `json:"status"`
	Uri SubscriptionPlansMetadataInteractionsPurchaseUri `json:"uri"`
}

type _SubscriptionPlansMetadataInteractionsPurchase SubscriptionPlansMetadataInteractionsPurchase

// NewSubscriptionPlansMetadataInteractionsPurchase instantiates a new SubscriptionPlansMetadataInteractionsPurchase object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewSubscriptionPlansMetadataInteractionsPurchase(status string, uri SubscriptionPlansMetadataInteractionsPurchaseUri) *SubscriptionPlansMetadataInteractionsPurchase {
	this := SubscriptionPlansMetadataInteractionsPurchase{}
	this.Status = status
	this.Uri = uri
	return &this
}

// NewSubscriptionPlansMetadataInteractionsPurchaseWithDefaults instantiates a new SubscriptionPlansMetadataInteractionsPurchase object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewSubscriptionPlansMetadataInteractionsPurchaseWithDefaults() *SubscriptionPlansMetadataInteractionsPurchase {
	this := SubscriptionPlansMetadataInteractionsPurchase{}
	return &this
}

// GetStatus returns the Status field value
func (o *SubscriptionPlansMetadataInteractionsPurchase) GetStatus() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Status
}

// GetStatusOk returns a tuple with the Status field value
// and a boolean to check if the value has been set.
func (o *SubscriptionPlansMetadataInteractionsPurchase) GetStatusOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Status, true
}

// SetStatus sets field value
func (o *SubscriptionPlansMetadataInteractionsPurchase) SetStatus(v string) {
	o.Status = v
}

// GetUri returns the Uri field value
func (o *SubscriptionPlansMetadataInteractionsPurchase) GetUri() SubscriptionPlansMetadataInteractionsPurchaseUri {
	if o == nil {
		var ret SubscriptionPlansMetadataInteractionsPurchaseUri
		return ret
	}

	return o.Uri
}

// GetUriOk returns a tuple with the Uri field value
// and a boolean to check if the value has been set.
func (o *SubscriptionPlansMetadataInteractionsPurchase) GetUriOk() (*SubscriptionPlansMetadataInteractionsPurchaseUri, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Uri, true
}

// SetUri sets field value
func (o *SubscriptionPlansMetadataInteractionsPurchase) SetUri(v SubscriptionPlansMetadataInteractionsPurchaseUri) {
	o.Uri = v
}

func (o SubscriptionPlansMetadataInteractionsPurchase) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o SubscriptionPlansMetadataInteractionsPurchase) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["status"] = o.Status
	toSerialize["uri"] = o.Uri
	return toSerialize, nil
}

func (o *SubscriptionPlansMetadataInteractionsPurchase) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"status",
		"uri",
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

	varSubscriptionPlansMetadataInteractionsPurchase := _SubscriptionPlansMetadataInteractionsPurchase{}

	decoder := json.NewDecoder(bytes.NewReader(data))
	decoder.DisallowUnknownFields()
	err = decoder.Decode(&varSubscriptionPlansMetadataInteractionsPurchase)

	if err != nil {
		return err
	}

	*o = SubscriptionPlansMetadataInteractionsPurchase(varSubscriptionPlansMetadataInteractionsPurchase)

	return err
}

type NullableSubscriptionPlansMetadataInteractionsPurchase struct {
	value *SubscriptionPlansMetadataInteractionsPurchase
	isSet bool
}

func (v NullableSubscriptionPlansMetadataInteractionsPurchase) Get() *SubscriptionPlansMetadataInteractionsPurchase {
	return v.value
}

func (v *NullableSubscriptionPlansMetadataInteractionsPurchase) Set(val *SubscriptionPlansMetadataInteractionsPurchase) {
	v.value = val
	v.isSet = true
}

func (v NullableSubscriptionPlansMetadataInteractionsPurchase) IsSet() bool {
	return v.isSet
}

func (v *NullableSubscriptionPlansMetadataInteractionsPurchase) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableSubscriptionPlansMetadataInteractionsPurchase(val *SubscriptionPlansMetadataInteractionsPurchase) *NullableSubscriptionPlansMetadataInteractionsPurchase {
	return &NullableSubscriptionPlansMetadataInteractionsPurchase{value: val, isSet: true}
}

func (v NullableSubscriptionPlansMetadataInteractionsPurchase) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableSubscriptionPlansMetadataInteractionsPurchase) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

