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

// checks if the ExchangeAuthCodeRequest type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &ExchangeAuthCodeRequest{}

// ExchangeAuthCodeRequest struct for ExchangeAuthCodeRequest
type ExchangeAuthCodeRequest struct {
	// The authorization code received from the authorization server.
	Code string `json:"code"`
	// The grant type. The value of this field must be `authorization_code`.
	GrantType string `json:"grant_type"`
	// The redirect URI. The value of this field must match the URI from `/oauth/authorize`.
	RedirectUri string `json:"redirect_uri"`
}

// NewExchangeAuthCodeRequest instantiates a new ExchangeAuthCodeRequest object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewExchangeAuthCodeRequest(code string, grantType string, redirectUri string) *ExchangeAuthCodeRequest {
	this := ExchangeAuthCodeRequest{}
	this.Code = code
	this.GrantType = grantType
	this.RedirectUri = redirectUri
	return &this
}

// NewExchangeAuthCodeRequestWithDefaults instantiates a new ExchangeAuthCodeRequest object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewExchangeAuthCodeRequestWithDefaults() *ExchangeAuthCodeRequest {
	this := ExchangeAuthCodeRequest{}
	return &this
}

// GetCode returns the Code field value
func (o *ExchangeAuthCodeRequest) GetCode() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Code
}

// GetCodeOk returns a tuple with the Code field value
// and a boolean to check if the value has been set.
func (o *ExchangeAuthCodeRequest) GetCodeOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Code, true
}

// SetCode sets field value
func (o *ExchangeAuthCodeRequest) SetCode(v string) {
	o.Code = v
}

// GetGrantType returns the GrantType field value
func (o *ExchangeAuthCodeRequest) GetGrantType() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.GrantType
}

// GetGrantTypeOk returns a tuple with the GrantType field value
// and a boolean to check if the value has been set.
func (o *ExchangeAuthCodeRequest) GetGrantTypeOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.GrantType, true
}

// SetGrantType sets field value
func (o *ExchangeAuthCodeRequest) SetGrantType(v string) {
	o.GrantType = v
}

// GetRedirectUri returns the RedirectUri field value
func (o *ExchangeAuthCodeRequest) GetRedirectUri() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.RedirectUri
}

// GetRedirectUriOk returns a tuple with the RedirectUri field value
// and a boolean to check if the value has been set.
func (o *ExchangeAuthCodeRequest) GetRedirectUriOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.RedirectUri, true
}

// SetRedirectUri sets field value
func (o *ExchangeAuthCodeRequest) SetRedirectUri(v string) {
	o.RedirectUri = v
}

func (o ExchangeAuthCodeRequest) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o ExchangeAuthCodeRequest) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["code"] = o.Code
	toSerialize["grant_type"] = o.GrantType
	toSerialize["redirect_uri"] = o.RedirectUri
	return toSerialize, nil
}

type NullableExchangeAuthCodeRequest struct {
	value *ExchangeAuthCodeRequest
	isSet bool
}

func (v NullableExchangeAuthCodeRequest) Get() *ExchangeAuthCodeRequest {
	return v.value
}

func (v *NullableExchangeAuthCodeRequest) Set(val *ExchangeAuthCodeRequest) {
	v.value = val
	v.isSet = true
}

func (v NullableExchangeAuthCodeRequest) IsSet() bool {
	return v.isSet
}

func (v *NullableExchangeAuthCodeRequest) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableExchangeAuthCodeRequest(val *ExchangeAuthCodeRequest) *NullableExchangeAuthCodeRequest {
	return &NullableExchangeAuthCodeRequest{value: val, isSet: true}
}

func (v NullableExchangeAuthCodeRequest) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableExchangeAuthCodeRequest) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


