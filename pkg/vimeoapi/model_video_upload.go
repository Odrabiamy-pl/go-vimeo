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

// checks if the VideoUpload type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &VideoUpload{}

// VideoUpload The video's upload information.
type VideoUpload struct {
	// The approach for uploading the video.  Option descriptions:  * `post` - The video upload uses the POST approach.  * `pull` - The video upload uses the pull approach.  * `tus` - The video upload uses the tus approach. 
	Approach *string `json:"approach,omitempty"`
	// The HTML form for uploading a video through the POST approach.
	Form *string `json:"form,omitempty"`
	// The ID of the Google Cloud Storage upload.
	GcsUid *string `json:"gcs_uid,omitempty"`
	// The link of the video to capture through the pull approach.
	Link *string `json:"link,omitempty"`
	// The redirect URL for the upload app.
	RedirectUrl *string `json:"redirect_url,omitempty"`
	// The file size in bytes of the uploaded video.
	Size *float32 `json:"size,omitempty"`
	// The status code for the availability of the uploaded video.  Option descriptions:  * `complete` - The upload is complete.  * `error` - The upload ended with an error.  * `in_progress` - The upload is underway. 
	Status string `json:"status"`
	// The link for sending video file data.
	UploadLink *string `json:"upload_link,omitempty"`
}

// NewVideoUpload instantiates a new VideoUpload object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewVideoUpload(status string) *VideoUpload {
	this := VideoUpload{}
	this.Status = status
	return &this
}

// NewVideoUploadWithDefaults instantiates a new VideoUpload object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewVideoUploadWithDefaults() *VideoUpload {
	this := VideoUpload{}
	return &this
}

// GetApproach returns the Approach field value if set, zero value otherwise.
func (o *VideoUpload) GetApproach() string {
	if o == nil || IsNil(o.Approach) {
		var ret string
		return ret
	}
	return *o.Approach
}

// GetApproachOk returns a tuple with the Approach field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *VideoUpload) GetApproachOk() (*string, bool) {
	if o == nil || IsNil(o.Approach) {
		return nil, false
	}
	return o.Approach, true
}

// HasApproach returns a boolean if a field has been set.
func (o *VideoUpload) HasApproach() bool {
	if o != nil && !IsNil(o.Approach) {
		return true
	}

	return false
}

// SetApproach gets a reference to the given string and assigns it to the Approach field.
func (o *VideoUpload) SetApproach(v string) {
	o.Approach = &v
}

// GetForm returns the Form field value if set, zero value otherwise.
func (o *VideoUpload) GetForm() string {
	if o == nil || IsNil(o.Form) {
		var ret string
		return ret
	}
	return *o.Form
}

// GetFormOk returns a tuple with the Form field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *VideoUpload) GetFormOk() (*string, bool) {
	if o == nil || IsNil(o.Form) {
		return nil, false
	}
	return o.Form, true
}

// HasForm returns a boolean if a field has been set.
func (o *VideoUpload) HasForm() bool {
	if o != nil && !IsNil(o.Form) {
		return true
	}

	return false
}

// SetForm gets a reference to the given string and assigns it to the Form field.
func (o *VideoUpload) SetForm(v string) {
	o.Form = &v
}

// GetGcsUid returns the GcsUid field value if set, zero value otherwise.
func (o *VideoUpload) GetGcsUid() string {
	if o == nil || IsNil(o.GcsUid) {
		var ret string
		return ret
	}
	return *o.GcsUid
}

// GetGcsUidOk returns a tuple with the GcsUid field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *VideoUpload) GetGcsUidOk() (*string, bool) {
	if o == nil || IsNil(o.GcsUid) {
		return nil, false
	}
	return o.GcsUid, true
}

// HasGcsUid returns a boolean if a field has been set.
func (o *VideoUpload) HasGcsUid() bool {
	if o != nil && !IsNil(o.GcsUid) {
		return true
	}

	return false
}

// SetGcsUid gets a reference to the given string and assigns it to the GcsUid field.
func (o *VideoUpload) SetGcsUid(v string) {
	o.GcsUid = &v
}

// GetLink returns the Link field value if set, zero value otherwise.
func (o *VideoUpload) GetLink() string {
	if o == nil || IsNil(o.Link) {
		var ret string
		return ret
	}
	return *o.Link
}

// GetLinkOk returns a tuple with the Link field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *VideoUpload) GetLinkOk() (*string, bool) {
	if o == nil || IsNil(o.Link) {
		return nil, false
	}
	return o.Link, true
}

// HasLink returns a boolean if a field has been set.
func (o *VideoUpload) HasLink() bool {
	if o != nil && !IsNil(o.Link) {
		return true
	}

	return false
}

// SetLink gets a reference to the given string and assigns it to the Link field.
func (o *VideoUpload) SetLink(v string) {
	o.Link = &v
}

// GetRedirectUrl returns the RedirectUrl field value if set, zero value otherwise.
func (o *VideoUpload) GetRedirectUrl() string {
	if o == nil || IsNil(o.RedirectUrl) {
		var ret string
		return ret
	}
	return *o.RedirectUrl
}

// GetRedirectUrlOk returns a tuple with the RedirectUrl field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *VideoUpload) GetRedirectUrlOk() (*string, bool) {
	if o == nil || IsNil(o.RedirectUrl) {
		return nil, false
	}
	return o.RedirectUrl, true
}

// HasRedirectUrl returns a boolean if a field has been set.
func (o *VideoUpload) HasRedirectUrl() bool {
	if o != nil && !IsNil(o.RedirectUrl) {
		return true
	}

	return false
}

// SetRedirectUrl gets a reference to the given string and assigns it to the RedirectUrl field.
func (o *VideoUpload) SetRedirectUrl(v string) {
	o.RedirectUrl = &v
}

// GetSize returns the Size field value if set, zero value otherwise.
func (o *VideoUpload) GetSize() float32 {
	if o == nil || IsNil(o.Size) {
		var ret float32
		return ret
	}
	return *o.Size
}

// GetSizeOk returns a tuple with the Size field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *VideoUpload) GetSizeOk() (*float32, bool) {
	if o == nil || IsNil(o.Size) {
		return nil, false
	}
	return o.Size, true
}

// HasSize returns a boolean if a field has been set.
func (o *VideoUpload) HasSize() bool {
	if o != nil && !IsNil(o.Size) {
		return true
	}

	return false
}

// SetSize gets a reference to the given float32 and assigns it to the Size field.
func (o *VideoUpload) SetSize(v float32) {
	o.Size = &v
}

// GetStatus returns the Status field value
func (o *VideoUpload) GetStatus() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Status
}

// GetStatusOk returns a tuple with the Status field value
// and a boolean to check if the value has been set.
func (o *VideoUpload) GetStatusOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Status, true
}

// SetStatus sets field value
func (o *VideoUpload) SetStatus(v string) {
	o.Status = v
}

// GetUploadLink returns the UploadLink field value if set, zero value otherwise.
func (o *VideoUpload) GetUploadLink() string {
	if o == nil || IsNil(o.UploadLink) {
		var ret string
		return ret
	}
	return *o.UploadLink
}

// GetUploadLinkOk returns a tuple with the UploadLink field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *VideoUpload) GetUploadLinkOk() (*string, bool) {
	if o == nil || IsNil(o.UploadLink) {
		return nil, false
	}
	return o.UploadLink, true
}

// HasUploadLink returns a boolean if a field has been set.
func (o *VideoUpload) HasUploadLink() bool {
	if o != nil && !IsNil(o.UploadLink) {
		return true
	}

	return false
}

// SetUploadLink gets a reference to the given string and assigns it to the UploadLink field.
func (o *VideoUpload) SetUploadLink(v string) {
	o.UploadLink = &v
}

func (o VideoUpload) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o VideoUpload) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Approach) {
		toSerialize["approach"] = o.Approach
	}
	if !IsNil(o.Form) {
		toSerialize["form"] = o.Form
	}
	if !IsNil(o.GcsUid) {
		toSerialize["gcs_uid"] = o.GcsUid
	}
	if !IsNil(o.Link) {
		toSerialize["link"] = o.Link
	}
	if !IsNil(o.RedirectUrl) {
		toSerialize["redirect_url"] = o.RedirectUrl
	}
	if !IsNil(o.Size) {
		toSerialize["size"] = o.Size
	}
	toSerialize["status"] = o.Status
	if !IsNil(o.UploadLink) {
		toSerialize["upload_link"] = o.UploadLink
	}
	return toSerialize, nil
}

type NullableVideoUpload struct {
	value *VideoUpload
	isSet bool
}

func (v NullableVideoUpload) Get() *VideoUpload {
	return v.value
}

func (v *NullableVideoUpload) Set(val *VideoUpload) {
	v.value = val
	v.isSet = true
}

func (v NullableVideoUpload) IsSet() bool {
	return v.isSet
}

func (v *NullableVideoUpload) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableVideoUpload(val *VideoUpload) *NullableVideoUpload {
	return &NullableVideoUpload{value: val, isSet: true}
}

func (v NullableVideoUpload) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableVideoUpload) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


