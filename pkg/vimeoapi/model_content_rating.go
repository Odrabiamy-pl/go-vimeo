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

// checks if the ContentRating type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &ContentRating{}

// ContentRating struct for ContentRating
type ContentRating struct {
	// The reason for the content rating.  Option descriptions:  * `advertisement` - The content contains an advertisement.  * `drugs` - The content contains drug or alcohol use.  * `language` - The content contains profanity or sexually suggestive language.  * `nudity` - The content contains nudity.  * `safe` - The content is suitable for all audiences.  * `unrated` - The content hasn't been rated.  * `violence` - The content contains violence or is graphic. 
	Code string `json:"code"`
	// The name of the content rating.
	Name string `json:"name"`
	// The canonical relative URI of the content rating.
	Uri NullableString `json:"uri"`
}

type _ContentRating ContentRating

// NewContentRating instantiates a new ContentRating object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewContentRating(code string, name string, uri NullableString) *ContentRating {
	this := ContentRating{}
	this.Code = code
	this.Name = name
	this.Uri = uri
	return &this
}

// NewContentRatingWithDefaults instantiates a new ContentRating object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewContentRatingWithDefaults() *ContentRating {
	this := ContentRating{}
	return &this
}

// GetCode returns the Code field value
func (o *ContentRating) GetCode() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Code
}

// GetCodeOk returns a tuple with the Code field value
// and a boolean to check if the value has been set.
func (o *ContentRating) GetCodeOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Code, true
}

// SetCode sets field value
func (o *ContentRating) SetCode(v string) {
	o.Code = v
}

// GetName returns the Name field value
func (o *ContentRating) GetName() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Name
}

// GetNameOk returns a tuple with the Name field value
// and a boolean to check if the value has been set.
func (o *ContentRating) GetNameOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Name, true
}

// SetName sets field value
func (o *ContentRating) SetName(v string) {
	o.Name = v
}

// GetUri returns the Uri field value
// If the value is explicit nil, the zero value for string will be returned
func (o *ContentRating) GetUri() string {
	if o == nil || o.Uri.Get() == nil {
		var ret string
		return ret
	}

	return *o.Uri.Get()
}

// GetUriOk returns a tuple with the Uri field value
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *ContentRating) GetUriOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.Uri.Get(), o.Uri.IsSet()
}

// SetUri sets field value
func (o *ContentRating) SetUri(v string) {
	o.Uri.Set(&v)
}

func (o ContentRating) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o ContentRating) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["code"] = o.Code
	toSerialize["name"] = o.Name
	toSerialize["uri"] = o.Uri.Get()
	return toSerialize, nil
}

func (o *ContentRating) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"code",
		"name",
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

	varContentRating := _ContentRating{}

	decoder := json.NewDecoder(bytes.NewReader(data))
	decoder.DisallowUnknownFields()
	err = decoder.Decode(&varContentRating)

	if err != nil {
		return err
	}

	*o = ContentRating(varContentRating)

	return err
}

type NullableContentRating struct {
	value *ContentRating
	isSet bool
}

func (v NullableContentRating) Get() *ContentRating {
	return v.value
}

func (v *NullableContentRating) Set(val *ContentRating) {
	v.value = val
	v.isSet = true
}

func (v NullableContentRating) IsSet() bool {
	return v.isSet
}

func (v *NullableContentRating) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableContentRating(val *ContentRating) *NullableContentRating {
	return &NullableContentRating{value: val, isSet: true}
}

func (v NullableContentRating) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableContentRating) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


