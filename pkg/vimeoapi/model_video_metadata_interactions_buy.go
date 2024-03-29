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

// checks if the VideoMetadataInteractionsBuy type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &VideoMetadataInteractionsBuy{}

// VideoMetadataInteractionsBuy The Buy interaction for the On Demand video.
type VideoMetadataInteractionsBuy struct {
	// The currency code for the user's region.
	Currency NullableString `json:"currency"`
	// The formatted display price for buying the On Demand video.
	DisplayPrice NullableString `json:"display_price"`
	// The user's download access to the On Demand video.  Option descriptions:  * `available` - The video is available for download.  * `purchased` - The user has purchased the video.  * `restricted` - The user isn't permitted to download the video.  * `unavailable` - The video isn't available for download.
	Download string `json:"download"`
	// Whether the On Demand video has DRM.
	Drm bool `json:"drm"`
	// The URL to buy the On Demand video on Vimeo.
	Link NullableString `json:"link"`
	// The price to buy the On Demand video.
	Price NullableFloat32 `json:"price"`
	// The time in ISO 8601 format when the On Demand video was purchased.
	PurchaseTime NullableString `json:"purchase_time"`
	// The user's streaming access to the On Demand video.  Option descriptions:  * `available` - The video is available for streaming.  * `purchased` - The user has purchased the video.  * `restricted` - The user isn't permitted to stream the video.  * `unavailable` - The video isn't available for streaming
	Stream string `json:"stream"`
	// The product URI to purchase the On Demand video.
	Uri NullableString `json:"uri"`
}

// NewVideoMetadataInteractionsBuy instantiates a new VideoMetadataInteractionsBuy object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewVideoMetadataInteractionsBuy(currency NullableString, displayPrice NullableString, download string, drm bool, link NullableString, price NullableFloat32, purchaseTime NullableString, stream string, uri NullableString) *VideoMetadataInteractionsBuy {
	this := VideoMetadataInteractionsBuy{}
	this.Currency = currency
	this.DisplayPrice = displayPrice
	this.Download = download
	this.Drm = drm
	this.Link = link
	this.Price = price
	this.PurchaseTime = purchaseTime
	this.Stream = stream
	this.Uri = uri
	return &this
}

// NewVideoMetadataInteractionsBuyWithDefaults instantiates a new VideoMetadataInteractionsBuy object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewVideoMetadataInteractionsBuyWithDefaults() *VideoMetadataInteractionsBuy {
	this := VideoMetadataInteractionsBuy{}
	return &this
}

// GetCurrency returns the Currency field value
// If the value is explicit nil, the zero value for string will be returned
func (o *VideoMetadataInteractionsBuy) GetCurrency() string {
	if o == nil || o.Currency.Get() == nil {
		var ret string
		return ret
	}

	return *o.Currency.Get()
}

// GetCurrencyOk returns a tuple with the Currency field value
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *VideoMetadataInteractionsBuy) GetCurrencyOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.Currency.Get(), o.Currency.IsSet()
}

// SetCurrency sets field value
func (o *VideoMetadataInteractionsBuy) SetCurrency(v string) {
	o.Currency.Set(&v)
}

// GetDisplayPrice returns the DisplayPrice field value
// If the value is explicit nil, the zero value for string will be returned
func (o *VideoMetadataInteractionsBuy) GetDisplayPrice() string {
	if o == nil || o.DisplayPrice.Get() == nil {
		var ret string
		return ret
	}

	return *o.DisplayPrice.Get()
}

// GetDisplayPriceOk returns a tuple with the DisplayPrice field value
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *VideoMetadataInteractionsBuy) GetDisplayPriceOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.DisplayPrice.Get(), o.DisplayPrice.IsSet()
}

// SetDisplayPrice sets field value
func (o *VideoMetadataInteractionsBuy) SetDisplayPrice(v string) {
	o.DisplayPrice.Set(&v)
}

// GetDownload returns the Download field value
func (o *VideoMetadataInteractionsBuy) GetDownload() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Download
}

// GetDownloadOk returns a tuple with the Download field value
// and a boolean to check if the value has been set.
func (o *VideoMetadataInteractionsBuy) GetDownloadOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Download, true
}

// SetDownload sets field value
func (o *VideoMetadataInteractionsBuy) SetDownload(v string) {
	o.Download = v
}

// GetDrm returns the Drm field value
func (o *VideoMetadataInteractionsBuy) GetDrm() bool {
	if o == nil {
		var ret bool
		return ret
	}

	return o.Drm
}

// GetDrmOk returns a tuple with the Drm field value
// and a boolean to check if the value has been set.
func (o *VideoMetadataInteractionsBuy) GetDrmOk() (*bool, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Drm, true
}

// SetDrm sets field value
func (o *VideoMetadataInteractionsBuy) SetDrm(v bool) {
	o.Drm = v
}

// GetLink returns the Link field value
// If the value is explicit nil, the zero value for string will be returned
func (o *VideoMetadataInteractionsBuy) GetLink() string {
	if o == nil || o.Link.Get() == nil {
		var ret string
		return ret
	}

	return *o.Link.Get()
}

// GetLinkOk returns a tuple with the Link field value
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *VideoMetadataInteractionsBuy) GetLinkOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.Link.Get(), o.Link.IsSet()
}

// SetLink sets field value
func (o *VideoMetadataInteractionsBuy) SetLink(v string) {
	o.Link.Set(&v)
}

// GetPrice returns the Price field value
// If the value is explicit nil, the zero value for float32 will be returned
func (o *VideoMetadataInteractionsBuy) GetPrice() float32 {
	if o == nil || o.Price.Get() == nil {
		var ret float32
		return ret
	}

	return *o.Price.Get()
}

// GetPriceOk returns a tuple with the Price field value
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *VideoMetadataInteractionsBuy) GetPriceOk() (*float32, bool) {
	if o == nil {
		return nil, false
	}
	return o.Price.Get(), o.Price.IsSet()
}

// SetPrice sets field value
func (o *VideoMetadataInteractionsBuy) SetPrice(v float32) {
	o.Price.Set(&v)
}

// GetPurchaseTime returns the PurchaseTime field value
// If the value is explicit nil, the zero value for string will be returned
func (o *VideoMetadataInteractionsBuy) GetPurchaseTime() string {
	if o == nil || o.PurchaseTime.Get() == nil {
		var ret string
		return ret
	}

	return *o.PurchaseTime.Get()
}

// GetPurchaseTimeOk returns a tuple with the PurchaseTime field value
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *VideoMetadataInteractionsBuy) GetPurchaseTimeOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.PurchaseTime.Get(), o.PurchaseTime.IsSet()
}

// SetPurchaseTime sets field value
func (o *VideoMetadataInteractionsBuy) SetPurchaseTime(v string) {
	o.PurchaseTime.Set(&v)
}

// GetStream returns the Stream field value
func (o *VideoMetadataInteractionsBuy) GetStream() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Stream
}

// GetStreamOk returns a tuple with the Stream field value
// and a boolean to check if the value has been set.
func (o *VideoMetadataInteractionsBuy) GetStreamOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Stream, true
}

// SetStream sets field value
func (o *VideoMetadataInteractionsBuy) SetStream(v string) {
	o.Stream = v
}

// GetUri returns the Uri field value
// If the value is explicit nil, the zero value for string will be returned
func (o *VideoMetadataInteractionsBuy) GetUri() string {
	if o == nil || o.Uri.Get() == nil {
		var ret string
		return ret
	}

	return *o.Uri.Get()
}

// GetUriOk returns a tuple with the Uri field value
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *VideoMetadataInteractionsBuy) GetUriOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.Uri.Get(), o.Uri.IsSet()
}

// SetUri sets field value
func (o *VideoMetadataInteractionsBuy) SetUri(v string) {
	o.Uri.Set(&v)
}

func (o VideoMetadataInteractionsBuy) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o VideoMetadataInteractionsBuy) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["currency"] = o.Currency.Get()
	toSerialize["display_price"] = o.DisplayPrice.Get()
	toSerialize["download"] = o.Download
	toSerialize["drm"] = o.Drm
	toSerialize["link"] = o.Link.Get()
	toSerialize["price"] = o.Price.Get()
	toSerialize["purchase_time"] = o.PurchaseTime.Get()
	toSerialize["stream"] = o.Stream
	toSerialize["uri"] = o.Uri.Get()
	return toSerialize, nil
}

type NullableVideoMetadataInteractionsBuy struct {
	value *VideoMetadataInteractionsBuy
	isSet bool
}

func (v NullableVideoMetadataInteractionsBuy) Get() *VideoMetadataInteractionsBuy {
	return v.value
}

func (v *NullableVideoMetadataInteractionsBuy) Set(val *VideoMetadataInteractionsBuy) {
	v.value = val
	v.isSet = true
}

func (v NullableVideoMetadataInteractionsBuy) IsSet() bool {
	return v.isSet
}

func (v *NullableVideoMetadataInteractionsBuy) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableVideoMetadataInteractionsBuy(val *VideoMetadataInteractionsBuy) *NullableVideoMetadataInteractionsBuy {
	return &NullableVideoMetadataInteractionsBuy{value: val, isSet: true}
}

func (v NullableVideoMetadataInteractionsBuy) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableVideoMetadataInteractionsBuy) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
