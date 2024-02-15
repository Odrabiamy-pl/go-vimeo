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

// checks if the UserMetadataInteractionsConnectedApps type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &UserMetadataInteractionsConnectedApps{}

// UserMetadataInteractionsConnectedApps struct for UserMetadataInteractionsConnectedApps
type UserMetadataInteractionsConnectedApps struct {
	// The list of all the scopes on the connected app that are needed for a particular Vimeo feature.
	AllScopes map[string]interface{} `json:"all_scopes"`
	// Whether the authenticated user is connected to the connected app.
	IsConnected bool `json:"is_connected"`
	// The list of the remaining scopes on the connected app that the authenticated user needs for a particular Vimeo feature.
	NeededScopes map[string]interface{} `json:"needed_scopes"`
	// An array of the HTTP methods permitted on this URI.
	Options []string `json:"options"`
	// The URI of the connected app.
	Uri string `json:"uri"`
}

type _UserMetadataInteractionsConnectedApps UserMetadataInteractionsConnectedApps

// NewUserMetadataInteractionsConnectedApps instantiates a new UserMetadataInteractionsConnectedApps object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewUserMetadataInteractionsConnectedApps(allScopes map[string]interface{}, isConnected bool, neededScopes map[string]interface{}, options []string, uri string) *UserMetadataInteractionsConnectedApps {
	this := UserMetadataInteractionsConnectedApps{}
	this.AllScopes = allScopes
	this.IsConnected = isConnected
	this.NeededScopes = neededScopes
	this.Options = options
	this.Uri = uri
	return &this
}

// NewUserMetadataInteractionsConnectedAppsWithDefaults instantiates a new UserMetadataInteractionsConnectedApps object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewUserMetadataInteractionsConnectedAppsWithDefaults() *UserMetadataInteractionsConnectedApps {
	this := UserMetadataInteractionsConnectedApps{}
	return &this
}

// GetAllScopes returns the AllScopes field value
func (o *UserMetadataInteractionsConnectedApps) GetAllScopes() map[string]interface{} {
	if o == nil {
		var ret map[string]interface{}
		return ret
	}

	return o.AllScopes
}

// GetAllScopesOk returns a tuple with the AllScopes field value
// and a boolean to check if the value has been set.
func (o *UserMetadataInteractionsConnectedApps) GetAllScopesOk() (map[string]interface{}, bool) {
	if o == nil {
		return map[string]interface{}{}, false
	}
	return o.AllScopes, true
}

// SetAllScopes sets field value
func (o *UserMetadataInteractionsConnectedApps) SetAllScopes(v map[string]interface{}) {
	o.AllScopes = v
}

// GetIsConnected returns the IsConnected field value
func (o *UserMetadataInteractionsConnectedApps) GetIsConnected() bool {
	if o == nil {
		var ret bool
		return ret
	}

	return o.IsConnected
}

// GetIsConnectedOk returns a tuple with the IsConnected field value
// and a boolean to check if the value has been set.
func (o *UserMetadataInteractionsConnectedApps) GetIsConnectedOk() (*bool, bool) {
	if o == nil {
		return nil, false
	}
	return &o.IsConnected, true
}

// SetIsConnected sets field value
func (o *UserMetadataInteractionsConnectedApps) SetIsConnected(v bool) {
	o.IsConnected = v
}

// GetNeededScopes returns the NeededScopes field value
func (o *UserMetadataInteractionsConnectedApps) GetNeededScopes() map[string]interface{} {
	if o == nil {
		var ret map[string]interface{}
		return ret
	}

	return o.NeededScopes
}

// GetNeededScopesOk returns a tuple with the NeededScopes field value
// and a boolean to check if the value has been set.
func (o *UserMetadataInteractionsConnectedApps) GetNeededScopesOk() (map[string]interface{}, bool) {
	if o == nil {
		return map[string]interface{}{}, false
	}
	return o.NeededScopes, true
}

// SetNeededScopes sets field value
func (o *UserMetadataInteractionsConnectedApps) SetNeededScopes(v map[string]interface{}) {
	o.NeededScopes = v
}

// GetOptions returns the Options field value
func (o *UserMetadataInteractionsConnectedApps) GetOptions() []string {
	if o == nil {
		var ret []string
		return ret
	}

	return o.Options
}

// GetOptionsOk returns a tuple with the Options field value
// and a boolean to check if the value has been set.
func (o *UserMetadataInteractionsConnectedApps) GetOptionsOk() ([]string, bool) {
	if o == nil {
		return nil, false
	}
	return o.Options, true
}

// SetOptions sets field value
func (o *UserMetadataInteractionsConnectedApps) SetOptions(v []string) {
	o.Options = v
}

// GetUri returns the Uri field value
func (o *UserMetadataInteractionsConnectedApps) GetUri() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Uri
}

// GetUriOk returns a tuple with the Uri field value
// and a boolean to check if the value has been set.
func (o *UserMetadataInteractionsConnectedApps) GetUriOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Uri, true
}

// SetUri sets field value
func (o *UserMetadataInteractionsConnectedApps) SetUri(v string) {
	o.Uri = v
}

func (o UserMetadataInteractionsConnectedApps) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o UserMetadataInteractionsConnectedApps) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["all_scopes"] = o.AllScopes
	toSerialize["is_connected"] = o.IsConnected
	toSerialize["needed_scopes"] = o.NeededScopes
	toSerialize["options"] = o.Options
	toSerialize["uri"] = o.Uri
	return toSerialize, nil
}

func (o *UserMetadataInteractionsConnectedApps) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"all_scopes",
		"is_connected",
		"needed_scopes",
		"options",
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

	varUserMetadataInteractionsConnectedApps := _UserMetadataInteractionsConnectedApps{}

	decoder := json.NewDecoder(bytes.NewReader(data))
	decoder.DisallowUnknownFields()
	err = decoder.Decode(&varUserMetadataInteractionsConnectedApps)

	if err != nil {
		return err
	}

	*o = UserMetadataInteractionsConnectedApps(varUserMetadataInteractionsConnectedApps)

	return err
}

type NullableUserMetadataInteractionsConnectedApps struct {
	value *UserMetadataInteractionsConnectedApps
	isSet bool
}

func (v NullableUserMetadataInteractionsConnectedApps) Get() *UserMetadataInteractionsConnectedApps {
	return v.value
}

func (v *NullableUserMetadataInteractionsConnectedApps) Set(val *UserMetadataInteractionsConnectedApps) {
	v.value = val
	v.isSet = true
}

func (v NullableUserMetadataInteractionsConnectedApps) IsSet() bool {
	return v.isSet
}

func (v *NullableUserMetadataInteractionsConnectedApps) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableUserMetadataInteractionsConnectedApps(val *UserMetadataInteractionsConnectedApps) *NullableUserMetadataInteractionsConnectedApps {
	return &NullableUserMetadataInteractionsConnectedApps{value: val, isSet: true}
}

func (v NullableUserMetadataInteractionsConnectedApps) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableUserMetadataInteractionsConnectedApps) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

