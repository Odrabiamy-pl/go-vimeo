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

// checks if the UploadVideoAlt1Request type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &UploadVideoAlt1Request{}

// UploadVideoAlt1Request struct for UploadVideoAlt1Request
type UploadVideoAlt1Request struct {
	// A list of values describing the content in this video. For a full list of values, use the [`/contentratings`](https://developer.vimeo.com/api/reference/videos#get_content_ratings) endpoint.
	ContentRating []string `json:"content_rating,omitempty"`
	// The description of the video.
	Description *string `json:"description,omitempty"`
	Embed *UploadVideoAlt1RequestEmbed `json:"embed,omitempty"`
	// The complete list of domains the video can be embedded on. This field requires that **privacy_embed** have the value `whitelist`.
	EmbedDomains []string `json:"embed_domains,omitempty"`
	// The URI of the folder to which the video is uploaded.
	FolderUri *string `json:"folder_uri,omitempty"`
	// Whether to hide the video from everyone except the video's owner. When the value is `true`, unlisted video links work only for the video's owner.
	HideFromVimeo *bool `json:"hide_from_vimeo,omitempty"`
	// The Creative Commons license under which the video is offered.  Option descriptions:  * `by` - The video is offered under CC BY, or the attibution-only license.  * `by-nc` - The video is offered under CC BY-NC, or the Attribution-NonCommercial license.  * `by-nc-nd` - The video is offered under CC BY-NC-ND, or the Attribution-NonCommercian-NoDerivs license.  * `by-nc-sa` - The video is offered under CC BY-NC-SA, or the Attribution-NonCommercial-ShareAlike licence.  * `by-nd` - The video is offered under CC BY-ND, or the Attribution-NoDerivs license.  * `by-sa` - The video is offered under CC BY-SA, or the Attribution-ShareAlike license.  * `cc0` - The video is offered under CC0, or the public domain license. 
	License *string `json:"license,omitempty"`
	// The video's default language. For a full list of supported languages, use the [`/languages?filter=texttracks`](https://developer.vimeo.com/api/reference/videos#get_languages) endpoint.
	Locale *string `json:"locale,omitempty"`
	// The title of the video.
	Name *string `json:"name,omitempty"`
	// The password. This field is required when **privacy.view** is `password`.
	Password *string `json:"password,omitempty"`
	Privacy *UploadVideoAlt1RequestPrivacy `json:"privacy,omitempty"`
	ReviewPage *UploadVideoAlt1RequestReviewPage `json:"review_page,omitempty"`
	Spatial *UploadVideoAlt1RequestSpatial `json:"spatial,omitempty"`
	Upload UploadVideoAlt1RequestUpload `json:"upload"`
}

// NewUploadVideoAlt1Request instantiates a new UploadVideoAlt1Request object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewUploadVideoAlt1Request(upload UploadVideoAlt1RequestUpload) *UploadVideoAlt1Request {
	this := UploadVideoAlt1Request{}
	this.Upload = upload
	return &this
}

// NewUploadVideoAlt1RequestWithDefaults instantiates a new UploadVideoAlt1Request object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewUploadVideoAlt1RequestWithDefaults() *UploadVideoAlt1Request {
	this := UploadVideoAlt1Request{}
	return &this
}

// GetContentRating returns the ContentRating field value if set, zero value otherwise.
func (o *UploadVideoAlt1Request) GetContentRating() []string {
	if o == nil || IsNil(o.ContentRating) {
		var ret []string
		return ret
	}
	return o.ContentRating
}

// GetContentRatingOk returns a tuple with the ContentRating field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UploadVideoAlt1Request) GetContentRatingOk() ([]string, bool) {
	if o == nil || IsNil(o.ContentRating) {
		return nil, false
	}
	return o.ContentRating, true
}

// HasContentRating returns a boolean if a field has been set.
func (o *UploadVideoAlt1Request) HasContentRating() bool {
	if o != nil && !IsNil(o.ContentRating) {
		return true
	}

	return false
}

// SetContentRating gets a reference to the given []string and assigns it to the ContentRating field.
func (o *UploadVideoAlt1Request) SetContentRating(v []string) {
	o.ContentRating = v
}

// GetDescription returns the Description field value if set, zero value otherwise.
func (o *UploadVideoAlt1Request) GetDescription() string {
	if o == nil || IsNil(o.Description) {
		var ret string
		return ret
	}
	return *o.Description
}

// GetDescriptionOk returns a tuple with the Description field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UploadVideoAlt1Request) GetDescriptionOk() (*string, bool) {
	if o == nil || IsNil(o.Description) {
		return nil, false
	}
	return o.Description, true
}

// HasDescription returns a boolean if a field has been set.
func (o *UploadVideoAlt1Request) HasDescription() bool {
	if o != nil && !IsNil(o.Description) {
		return true
	}

	return false
}

// SetDescription gets a reference to the given string and assigns it to the Description field.
func (o *UploadVideoAlt1Request) SetDescription(v string) {
	o.Description = &v
}

// GetEmbed returns the Embed field value if set, zero value otherwise.
func (o *UploadVideoAlt1Request) GetEmbed() UploadVideoAlt1RequestEmbed {
	if o == nil || IsNil(o.Embed) {
		var ret UploadVideoAlt1RequestEmbed
		return ret
	}
	return *o.Embed
}

// GetEmbedOk returns a tuple with the Embed field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UploadVideoAlt1Request) GetEmbedOk() (*UploadVideoAlt1RequestEmbed, bool) {
	if o == nil || IsNil(o.Embed) {
		return nil, false
	}
	return o.Embed, true
}

// HasEmbed returns a boolean if a field has been set.
func (o *UploadVideoAlt1Request) HasEmbed() bool {
	if o != nil && !IsNil(o.Embed) {
		return true
	}

	return false
}

// SetEmbed gets a reference to the given UploadVideoAlt1RequestEmbed and assigns it to the Embed field.
func (o *UploadVideoAlt1Request) SetEmbed(v UploadVideoAlt1RequestEmbed) {
	o.Embed = &v
}

// GetEmbedDomains returns the EmbedDomains field value if set, zero value otherwise.
func (o *UploadVideoAlt1Request) GetEmbedDomains() []string {
	if o == nil || IsNil(o.EmbedDomains) {
		var ret []string
		return ret
	}
	return o.EmbedDomains
}

// GetEmbedDomainsOk returns a tuple with the EmbedDomains field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UploadVideoAlt1Request) GetEmbedDomainsOk() ([]string, bool) {
	if o == nil || IsNil(o.EmbedDomains) {
		return nil, false
	}
	return o.EmbedDomains, true
}

// HasEmbedDomains returns a boolean if a field has been set.
func (o *UploadVideoAlt1Request) HasEmbedDomains() bool {
	if o != nil && !IsNil(o.EmbedDomains) {
		return true
	}

	return false
}

// SetEmbedDomains gets a reference to the given []string and assigns it to the EmbedDomains field.
func (o *UploadVideoAlt1Request) SetEmbedDomains(v []string) {
	o.EmbedDomains = v
}

// GetFolderUri returns the FolderUri field value if set, zero value otherwise.
func (o *UploadVideoAlt1Request) GetFolderUri() string {
	if o == nil || IsNil(o.FolderUri) {
		var ret string
		return ret
	}
	return *o.FolderUri
}

// GetFolderUriOk returns a tuple with the FolderUri field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UploadVideoAlt1Request) GetFolderUriOk() (*string, bool) {
	if o == nil || IsNil(o.FolderUri) {
		return nil, false
	}
	return o.FolderUri, true
}

// HasFolderUri returns a boolean if a field has been set.
func (o *UploadVideoAlt1Request) HasFolderUri() bool {
	if o != nil && !IsNil(o.FolderUri) {
		return true
	}

	return false
}

// SetFolderUri gets a reference to the given string and assigns it to the FolderUri field.
func (o *UploadVideoAlt1Request) SetFolderUri(v string) {
	o.FolderUri = &v
}

// GetHideFromVimeo returns the HideFromVimeo field value if set, zero value otherwise.
func (o *UploadVideoAlt1Request) GetHideFromVimeo() bool {
	if o == nil || IsNil(o.HideFromVimeo) {
		var ret bool
		return ret
	}
	return *o.HideFromVimeo
}

// GetHideFromVimeoOk returns a tuple with the HideFromVimeo field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UploadVideoAlt1Request) GetHideFromVimeoOk() (*bool, bool) {
	if o == nil || IsNil(o.HideFromVimeo) {
		return nil, false
	}
	return o.HideFromVimeo, true
}

// HasHideFromVimeo returns a boolean if a field has been set.
func (o *UploadVideoAlt1Request) HasHideFromVimeo() bool {
	if o != nil && !IsNil(o.HideFromVimeo) {
		return true
	}

	return false
}

// SetHideFromVimeo gets a reference to the given bool and assigns it to the HideFromVimeo field.
func (o *UploadVideoAlt1Request) SetHideFromVimeo(v bool) {
	o.HideFromVimeo = &v
}

// GetLicense returns the License field value if set, zero value otherwise.
func (o *UploadVideoAlt1Request) GetLicense() string {
	if o == nil || IsNil(o.License) {
		var ret string
		return ret
	}
	return *o.License
}

// GetLicenseOk returns a tuple with the License field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UploadVideoAlt1Request) GetLicenseOk() (*string, bool) {
	if o == nil || IsNil(o.License) {
		return nil, false
	}
	return o.License, true
}

// HasLicense returns a boolean if a field has been set.
func (o *UploadVideoAlt1Request) HasLicense() bool {
	if o != nil && !IsNil(o.License) {
		return true
	}

	return false
}

// SetLicense gets a reference to the given string and assigns it to the License field.
func (o *UploadVideoAlt1Request) SetLicense(v string) {
	o.License = &v
}

// GetLocale returns the Locale field value if set, zero value otherwise.
func (o *UploadVideoAlt1Request) GetLocale() string {
	if o == nil || IsNil(o.Locale) {
		var ret string
		return ret
	}
	return *o.Locale
}

// GetLocaleOk returns a tuple with the Locale field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UploadVideoAlt1Request) GetLocaleOk() (*string, bool) {
	if o == nil || IsNil(o.Locale) {
		return nil, false
	}
	return o.Locale, true
}

// HasLocale returns a boolean if a field has been set.
func (o *UploadVideoAlt1Request) HasLocale() bool {
	if o != nil && !IsNil(o.Locale) {
		return true
	}

	return false
}

// SetLocale gets a reference to the given string and assigns it to the Locale field.
func (o *UploadVideoAlt1Request) SetLocale(v string) {
	o.Locale = &v
}

// GetName returns the Name field value if set, zero value otherwise.
func (o *UploadVideoAlt1Request) GetName() string {
	if o == nil || IsNil(o.Name) {
		var ret string
		return ret
	}
	return *o.Name
}

// GetNameOk returns a tuple with the Name field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UploadVideoAlt1Request) GetNameOk() (*string, bool) {
	if o == nil || IsNil(o.Name) {
		return nil, false
	}
	return o.Name, true
}

// HasName returns a boolean if a field has been set.
func (o *UploadVideoAlt1Request) HasName() bool {
	if o != nil && !IsNil(o.Name) {
		return true
	}

	return false
}

// SetName gets a reference to the given string and assigns it to the Name field.
func (o *UploadVideoAlt1Request) SetName(v string) {
	o.Name = &v
}

// GetPassword returns the Password field value if set, zero value otherwise.
func (o *UploadVideoAlt1Request) GetPassword() string {
	if o == nil || IsNil(o.Password) {
		var ret string
		return ret
	}
	return *o.Password
}

// GetPasswordOk returns a tuple with the Password field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UploadVideoAlt1Request) GetPasswordOk() (*string, bool) {
	if o == nil || IsNil(o.Password) {
		return nil, false
	}
	return o.Password, true
}

// HasPassword returns a boolean if a field has been set.
func (o *UploadVideoAlt1Request) HasPassword() bool {
	if o != nil && !IsNil(o.Password) {
		return true
	}

	return false
}

// SetPassword gets a reference to the given string and assigns it to the Password field.
func (o *UploadVideoAlt1Request) SetPassword(v string) {
	o.Password = &v
}

// GetPrivacy returns the Privacy field value if set, zero value otherwise.
func (o *UploadVideoAlt1Request) GetPrivacy() UploadVideoAlt1RequestPrivacy {
	if o == nil || IsNil(o.Privacy) {
		var ret UploadVideoAlt1RequestPrivacy
		return ret
	}
	return *o.Privacy
}

// GetPrivacyOk returns a tuple with the Privacy field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UploadVideoAlt1Request) GetPrivacyOk() (*UploadVideoAlt1RequestPrivacy, bool) {
	if o == nil || IsNil(o.Privacy) {
		return nil, false
	}
	return o.Privacy, true
}

// HasPrivacy returns a boolean if a field has been set.
func (o *UploadVideoAlt1Request) HasPrivacy() bool {
	if o != nil && !IsNil(o.Privacy) {
		return true
	}

	return false
}

// SetPrivacy gets a reference to the given UploadVideoAlt1RequestPrivacy and assigns it to the Privacy field.
func (o *UploadVideoAlt1Request) SetPrivacy(v UploadVideoAlt1RequestPrivacy) {
	o.Privacy = &v
}

// GetReviewPage returns the ReviewPage field value if set, zero value otherwise.
func (o *UploadVideoAlt1Request) GetReviewPage() UploadVideoAlt1RequestReviewPage {
	if o == nil || IsNil(o.ReviewPage) {
		var ret UploadVideoAlt1RequestReviewPage
		return ret
	}
	return *o.ReviewPage
}

// GetReviewPageOk returns a tuple with the ReviewPage field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UploadVideoAlt1Request) GetReviewPageOk() (*UploadVideoAlt1RequestReviewPage, bool) {
	if o == nil || IsNil(o.ReviewPage) {
		return nil, false
	}
	return o.ReviewPage, true
}

// HasReviewPage returns a boolean if a field has been set.
func (o *UploadVideoAlt1Request) HasReviewPage() bool {
	if o != nil && !IsNil(o.ReviewPage) {
		return true
	}

	return false
}

// SetReviewPage gets a reference to the given UploadVideoAlt1RequestReviewPage and assigns it to the ReviewPage field.
func (o *UploadVideoAlt1Request) SetReviewPage(v UploadVideoAlt1RequestReviewPage) {
	o.ReviewPage = &v
}

// GetSpatial returns the Spatial field value if set, zero value otherwise.
func (o *UploadVideoAlt1Request) GetSpatial() UploadVideoAlt1RequestSpatial {
	if o == nil || IsNil(o.Spatial) {
		var ret UploadVideoAlt1RequestSpatial
		return ret
	}
	return *o.Spatial
}

// GetSpatialOk returns a tuple with the Spatial field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UploadVideoAlt1Request) GetSpatialOk() (*UploadVideoAlt1RequestSpatial, bool) {
	if o == nil || IsNil(o.Spatial) {
		return nil, false
	}
	return o.Spatial, true
}

// HasSpatial returns a boolean if a field has been set.
func (o *UploadVideoAlt1Request) HasSpatial() bool {
	if o != nil && !IsNil(o.Spatial) {
		return true
	}

	return false
}

// SetSpatial gets a reference to the given UploadVideoAlt1RequestSpatial and assigns it to the Spatial field.
func (o *UploadVideoAlt1Request) SetSpatial(v UploadVideoAlt1RequestSpatial) {
	o.Spatial = &v
}

// GetUpload returns the Upload field value
func (o *UploadVideoAlt1Request) GetUpload() UploadVideoAlt1RequestUpload {
	if o == nil {
		var ret UploadVideoAlt1RequestUpload
		return ret
	}

	return o.Upload
}

// GetUploadOk returns a tuple with the Upload field value
// and a boolean to check if the value has been set.
func (o *UploadVideoAlt1Request) GetUploadOk() (*UploadVideoAlt1RequestUpload, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Upload, true
}

// SetUpload sets field value
func (o *UploadVideoAlt1Request) SetUpload(v UploadVideoAlt1RequestUpload) {
	o.Upload = v
}

func (o UploadVideoAlt1Request) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o UploadVideoAlt1Request) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.ContentRating) {
		toSerialize["content_rating"] = o.ContentRating
	}
	if !IsNil(o.Description) {
		toSerialize["description"] = o.Description
	}
	if !IsNil(o.Embed) {
		toSerialize["embed"] = o.Embed
	}
	if !IsNil(o.EmbedDomains) {
		toSerialize["embed_domains"] = o.EmbedDomains
	}
	if !IsNil(o.FolderUri) {
		toSerialize["folder_uri"] = o.FolderUri
	}
	if !IsNil(o.HideFromVimeo) {
		toSerialize["hide_from_vimeo"] = o.HideFromVimeo
	}
	if !IsNil(o.License) {
		toSerialize["license"] = o.License
	}
	if !IsNil(o.Locale) {
		toSerialize["locale"] = o.Locale
	}
	if !IsNil(o.Name) {
		toSerialize["name"] = o.Name
	}
	if !IsNil(o.Password) {
		toSerialize["password"] = o.Password
	}
	if !IsNil(o.Privacy) {
		toSerialize["privacy"] = o.Privacy
	}
	if !IsNil(o.ReviewPage) {
		toSerialize["review_page"] = o.ReviewPage
	}
	if !IsNil(o.Spatial) {
		toSerialize["spatial"] = o.Spatial
	}
	toSerialize["upload"] = o.Upload
	return toSerialize, nil
}

type NullableUploadVideoAlt1Request struct {
	value *UploadVideoAlt1Request
	isSet bool
}

func (v NullableUploadVideoAlt1Request) Get() *UploadVideoAlt1Request {
	return v.value
}

func (v *NullableUploadVideoAlt1Request) Set(val *UploadVideoAlt1Request) {
	v.value = val
	v.isSet = true
}

func (v NullableUploadVideoAlt1Request) IsSet() bool {
	return v.isSet
}

func (v *NullableUploadVideoAlt1Request) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableUploadVideoAlt1Request(val *UploadVideoAlt1Request) *NullableUploadVideoAlt1Request {
	return &NullableUploadVideoAlt1Request{value: val, isSet: true}
}

func (v NullableUploadVideoAlt1Request) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableUploadVideoAlt1Request) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


