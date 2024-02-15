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

// checks if the Location type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &Location{}

// Location struct for Location
type Location struct {
	// The authenticated user's city.
	City NullableString `json:"city"`
	// The authenticated user's country.
	Country NullableString `json:"country"`
	// The ISO code of the authenticated user's country.
	CountryIsoCode string `json:"country_iso_code"`
	// The authenticated user's formatted address string.
	FormattedAddress string `json:"formatted_address"`
	// The authenticated user's latitude.
	Latitude float32 `json:"latitude"`
	// The authenticated user's longitude.
	Longitude float32 `json:"longitude"`
	// The authenticated user's neighborhood.
	Neighborhood NullableString `json:"neighborhood"`
	// The authenticated user's state.
	State NullableString `json:"state"`
	// The ISO code of the authenticated user's state.
	StateIsoCode NullableString `json:"state_iso_code"`
	// The authenticated user's sub-locality.
	SubLocality NullableString `json:"sub_locality"`
}

type _Location Location

// NewLocation instantiates a new Location object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewLocation(city NullableString, country NullableString, countryIsoCode string, formattedAddress string, latitude float32, longitude float32, neighborhood NullableString, state NullableString, stateIsoCode NullableString, subLocality NullableString) *Location {
	this := Location{}
	this.City = city
	this.Country = country
	this.CountryIsoCode = countryIsoCode
	this.FormattedAddress = formattedAddress
	this.Latitude = latitude
	this.Longitude = longitude
	this.Neighborhood = neighborhood
	this.State = state
	this.StateIsoCode = stateIsoCode
	this.SubLocality = subLocality
	return &this
}

// NewLocationWithDefaults instantiates a new Location object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewLocationWithDefaults() *Location {
	this := Location{}
	return &this
}

// GetCity returns the City field value
// If the value is explicit nil, the zero value for string will be returned
func (o *Location) GetCity() string {
	if o == nil || o.City.Get() == nil {
		var ret string
		return ret
	}

	return *o.City.Get()
}

// GetCityOk returns a tuple with the City field value
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *Location) GetCityOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.City.Get(), o.City.IsSet()
}

// SetCity sets field value
func (o *Location) SetCity(v string) {
	o.City.Set(&v)
}

// GetCountry returns the Country field value
// If the value is explicit nil, the zero value for string will be returned
func (o *Location) GetCountry() string {
	if o == nil || o.Country.Get() == nil {
		var ret string
		return ret
	}

	return *o.Country.Get()
}

// GetCountryOk returns a tuple with the Country field value
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *Location) GetCountryOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.Country.Get(), o.Country.IsSet()
}

// SetCountry sets field value
func (o *Location) SetCountry(v string) {
	o.Country.Set(&v)
}

// GetCountryIsoCode returns the CountryIsoCode field value
func (o *Location) GetCountryIsoCode() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.CountryIsoCode
}

// GetCountryIsoCodeOk returns a tuple with the CountryIsoCode field value
// and a boolean to check if the value has been set.
func (o *Location) GetCountryIsoCodeOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.CountryIsoCode, true
}

// SetCountryIsoCode sets field value
func (o *Location) SetCountryIsoCode(v string) {
	o.CountryIsoCode = v
}

// GetFormattedAddress returns the FormattedAddress field value
func (o *Location) GetFormattedAddress() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.FormattedAddress
}

// GetFormattedAddressOk returns a tuple with the FormattedAddress field value
// and a boolean to check if the value has been set.
func (o *Location) GetFormattedAddressOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.FormattedAddress, true
}

// SetFormattedAddress sets field value
func (o *Location) SetFormattedAddress(v string) {
	o.FormattedAddress = v
}

// GetLatitude returns the Latitude field value
func (o *Location) GetLatitude() float32 {
	if o == nil {
		var ret float32
		return ret
	}

	return o.Latitude
}

// GetLatitudeOk returns a tuple with the Latitude field value
// and a boolean to check if the value has been set.
func (o *Location) GetLatitudeOk() (*float32, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Latitude, true
}

// SetLatitude sets field value
func (o *Location) SetLatitude(v float32) {
	o.Latitude = v
}

// GetLongitude returns the Longitude field value
func (o *Location) GetLongitude() float32 {
	if o == nil {
		var ret float32
		return ret
	}

	return o.Longitude
}

// GetLongitudeOk returns a tuple with the Longitude field value
// and a boolean to check if the value has been set.
func (o *Location) GetLongitudeOk() (*float32, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Longitude, true
}

// SetLongitude sets field value
func (o *Location) SetLongitude(v float32) {
	o.Longitude = v
}

// GetNeighborhood returns the Neighborhood field value
// If the value is explicit nil, the zero value for string will be returned
func (o *Location) GetNeighborhood() string {
	if o == nil || o.Neighborhood.Get() == nil {
		var ret string
		return ret
	}

	return *o.Neighborhood.Get()
}

// GetNeighborhoodOk returns a tuple with the Neighborhood field value
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *Location) GetNeighborhoodOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.Neighborhood.Get(), o.Neighborhood.IsSet()
}

// SetNeighborhood sets field value
func (o *Location) SetNeighborhood(v string) {
	o.Neighborhood.Set(&v)
}

// GetState returns the State field value
// If the value is explicit nil, the zero value for string will be returned
func (o *Location) GetState() string {
	if o == nil || o.State.Get() == nil {
		var ret string
		return ret
	}

	return *o.State.Get()
}

// GetStateOk returns a tuple with the State field value
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *Location) GetStateOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.State.Get(), o.State.IsSet()
}

// SetState sets field value
func (o *Location) SetState(v string) {
	o.State.Set(&v)
}

// GetStateIsoCode returns the StateIsoCode field value
// If the value is explicit nil, the zero value for string will be returned
func (o *Location) GetStateIsoCode() string {
	if o == nil || o.StateIsoCode.Get() == nil {
		var ret string
		return ret
	}

	return *o.StateIsoCode.Get()
}

// GetStateIsoCodeOk returns a tuple with the StateIsoCode field value
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *Location) GetStateIsoCodeOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.StateIsoCode.Get(), o.StateIsoCode.IsSet()
}

// SetStateIsoCode sets field value
func (o *Location) SetStateIsoCode(v string) {
	o.StateIsoCode.Set(&v)
}

// GetSubLocality returns the SubLocality field value
// If the value is explicit nil, the zero value for string will be returned
func (o *Location) GetSubLocality() string {
	if o == nil || o.SubLocality.Get() == nil {
		var ret string
		return ret
	}

	return *o.SubLocality.Get()
}

// GetSubLocalityOk returns a tuple with the SubLocality field value
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *Location) GetSubLocalityOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.SubLocality.Get(), o.SubLocality.IsSet()
}

// SetSubLocality sets field value
func (o *Location) SetSubLocality(v string) {
	o.SubLocality.Set(&v)
}

func (o Location) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o Location) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["city"] = o.City.Get()
	toSerialize["country"] = o.Country.Get()
	toSerialize["country_iso_code"] = o.CountryIsoCode
	toSerialize["formatted_address"] = o.FormattedAddress
	toSerialize["latitude"] = o.Latitude
	toSerialize["longitude"] = o.Longitude
	toSerialize["neighborhood"] = o.Neighborhood.Get()
	toSerialize["state"] = o.State.Get()
	toSerialize["state_iso_code"] = o.StateIsoCode.Get()
	toSerialize["sub_locality"] = o.SubLocality.Get()
	return toSerialize, nil
}

func (o *Location) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"city",
		"country",
		"country_iso_code",
		"formatted_address",
		"latitude",
		"longitude",
		"neighborhood",
		"state",
		"state_iso_code",
		"sub_locality",
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

	varLocation := _Location{}

	decoder := json.NewDecoder(bytes.NewReader(data))
	decoder.DisallowUnknownFields()
	err = decoder.Decode(&varLocation)

	if err != nil {
		return err
	}

	*o = Location(varLocation)

	return err
}

type NullableLocation struct {
	value *Location
	isSet bool
}

func (v NullableLocation) Get() *Location {
	return v.value
}

func (v *NullableLocation) Set(val *Location) {
	v.value = val
	v.isSet = true
}

func (v NullableLocation) IsSet() bool {
	return v.isSet
}

func (v *NullableLocation) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableLocation(val *Location) *NullableLocation {
	return &NullableLocation{value: val, isSet: true}
}

func (v NullableLocation) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableLocation) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


