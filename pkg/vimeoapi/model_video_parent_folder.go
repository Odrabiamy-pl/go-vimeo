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

// checks if the VideoParentFolder type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &VideoParentFolder{}

// VideoParentFolder Information about the folder that contains the video.
type VideoParentFolder struct {
	// The access grant response that applies to the team member. _This field is deprecated because grants are no longer exposed via API responses._
	AccessGrant map[string]interface{} `json:"access_grant"`
	// The time in ISO 8601 format when the folder was created.
	CreatedTime string `json:"created_time"`
	// The URI for the user who created the folder.
	CreatorUri string `json:"creator_uri"`
	// Whether this folder has at least one subfolder.
	HasSubfolder bool `json:"has_subfolder"`
	// Whether the folder is pinned.
	IsPinned bool `json:"is_pinned"`
	// Whether the folder is a private-to-me folder for the user.
	IsPrivateToUser bool `json:"is_private_to_user"`
	// The time in ISO 8601 format when a user last performed an action on the folder.
	LastUserActionEventDate NullableString `json:"last_user_action_event_date"`
	// The link to the folder on Vimeo.
	Link string `json:"link"`
	// The link to the folder management page.
	ManageLink string `json:"manage_link"`
	Metadata ProjectMetadata `json:"metadata"`
	// The time in ISO 8601 format when the folder was last modified.
	ModifiedTime string `json:"modified_time"`
	// The name of the folder.
	Name string `json:"name"`
	// The time in ISO 8601 format when the folder was pinned.
	PinnedOn NullableString `json:"pinned_on"`
	Privacy ProjectPrivacy `json:"privacy"`
	// The resource key string of the folder.
	ResourceKey string `json:"resource_key"`
	Settings ProjectSettings `json:"settings"`
	// The URI of the folder.
	Uri string `json:"uri"`
	// The owner of the folder.
	User User `json:"user"`
}

type _VideoParentFolder VideoParentFolder

// NewVideoParentFolder instantiates a new VideoParentFolder object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewVideoParentFolder(accessGrant map[string]interface{}, createdTime string, creatorUri string, hasSubfolder bool, isPinned bool, isPrivateToUser bool, lastUserActionEventDate NullableString, link string, manageLink string, metadata ProjectMetadata, modifiedTime string, name string, pinnedOn NullableString, privacy ProjectPrivacy, resourceKey string, settings ProjectSettings, uri string, user User) *VideoParentFolder {
	this := VideoParentFolder{}
	this.AccessGrant = accessGrant
	this.CreatedTime = createdTime
	this.CreatorUri = creatorUri
	this.HasSubfolder = hasSubfolder
	this.IsPinned = isPinned
	this.IsPrivateToUser = isPrivateToUser
	this.LastUserActionEventDate = lastUserActionEventDate
	this.Link = link
	this.ManageLink = manageLink
	this.Metadata = metadata
	this.ModifiedTime = modifiedTime
	this.Name = name
	this.PinnedOn = pinnedOn
	this.Privacy = privacy
	this.ResourceKey = resourceKey
	this.Settings = settings
	this.Uri = uri
	this.User = user
	return &this
}

// NewVideoParentFolderWithDefaults instantiates a new VideoParentFolder object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewVideoParentFolderWithDefaults() *VideoParentFolder {
	this := VideoParentFolder{}
	return &this
}

// GetAccessGrant returns the AccessGrant field value
func (o *VideoParentFolder) GetAccessGrant() map[string]interface{} {
	if o == nil {
		var ret map[string]interface{}
		return ret
	}

	return o.AccessGrant
}

// GetAccessGrantOk returns a tuple with the AccessGrant field value
// and a boolean to check if the value has been set.
func (o *VideoParentFolder) GetAccessGrantOk() (map[string]interface{}, bool) {
	if o == nil {
		return map[string]interface{}{}, false
	}
	return o.AccessGrant, true
}

// SetAccessGrant sets field value
func (o *VideoParentFolder) SetAccessGrant(v map[string]interface{}) {
	o.AccessGrant = v
}

// GetCreatedTime returns the CreatedTime field value
func (o *VideoParentFolder) GetCreatedTime() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.CreatedTime
}

// GetCreatedTimeOk returns a tuple with the CreatedTime field value
// and a boolean to check if the value has been set.
func (o *VideoParentFolder) GetCreatedTimeOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.CreatedTime, true
}

// SetCreatedTime sets field value
func (o *VideoParentFolder) SetCreatedTime(v string) {
	o.CreatedTime = v
}

// GetCreatorUri returns the CreatorUri field value
func (o *VideoParentFolder) GetCreatorUri() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.CreatorUri
}

// GetCreatorUriOk returns a tuple with the CreatorUri field value
// and a boolean to check if the value has been set.
func (o *VideoParentFolder) GetCreatorUriOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.CreatorUri, true
}

// SetCreatorUri sets field value
func (o *VideoParentFolder) SetCreatorUri(v string) {
	o.CreatorUri = v
}

// GetHasSubfolder returns the HasSubfolder field value
func (o *VideoParentFolder) GetHasSubfolder() bool {
	if o == nil {
		var ret bool
		return ret
	}

	return o.HasSubfolder
}

// GetHasSubfolderOk returns a tuple with the HasSubfolder field value
// and a boolean to check if the value has been set.
func (o *VideoParentFolder) GetHasSubfolderOk() (*bool, bool) {
	if o == nil {
		return nil, false
	}
	return &o.HasSubfolder, true
}

// SetHasSubfolder sets field value
func (o *VideoParentFolder) SetHasSubfolder(v bool) {
	o.HasSubfolder = v
}

// GetIsPinned returns the IsPinned field value
func (o *VideoParentFolder) GetIsPinned() bool {
	if o == nil {
		var ret bool
		return ret
	}

	return o.IsPinned
}

// GetIsPinnedOk returns a tuple with the IsPinned field value
// and a boolean to check if the value has been set.
func (o *VideoParentFolder) GetIsPinnedOk() (*bool, bool) {
	if o == nil {
		return nil, false
	}
	return &o.IsPinned, true
}

// SetIsPinned sets field value
func (o *VideoParentFolder) SetIsPinned(v bool) {
	o.IsPinned = v
}

// GetIsPrivateToUser returns the IsPrivateToUser field value
func (o *VideoParentFolder) GetIsPrivateToUser() bool {
	if o == nil {
		var ret bool
		return ret
	}

	return o.IsPrivateToUser
}

// GetIsPrivateToUserOk returns a tuple with the IsPrivateToUser field value
// and a boolean to check if the value has been set.
func (o *VideoParentFolder) GetIsPrivateToUserOk() (*bool, bool) {
	if o == nil {
		return nil, false
	}
	return &o.IsPrivateToUser, true
}

// SetIsPrivateToUser sets field value
func (o *VideoParentFolder) SetIsPrivateToUser(v bool) {
	o.IsPrivateToUser = v
}

// GetLastUserActionEventDate returns the LastUserActionEventDate field value
// If the value is explicit nil, the zero value for string will be returned
func (o *VideoParentFolder) GetLastUserActionEventDate() string {
	if o == nil || o.LastUserActionEventDate.Get() == nil {
		var ret string
		return ret
	}

	return *o.LastUserActionEventDate.Get()
}

// GetLastUserActionEventDateOk returns a tuple with the LastUserActionEventDate field value
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *VideoParentFolder) GetLastUserActionEventDateOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.LastUserActionEventDate.Get(), o.LastUserActionEventDate.IsSet()
}

// SetLastUserActionEventDate sets field value
func (o *VideoParentFolder) SetLastUserActionEventDate(v string) {
	o.LastUserActionEventDate.Set(&v)
}

// GetLink returns the Link field value
func (o *VideoParentFolder) GetLink() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Link
}

// GetLinkOk returns a tuple with the Link field value
// and a boolean to check if the value has been set.
func (o *VideoParentFolder) GetLinkOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Link, true
}

// SetLink sets field value
func (o *VideoParentFolder) SetLink(v string) {
	o.Link = v
}

// GetManageLink returns the ManageLink field value
func (o *VideoParentFolder) GetManageLink() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.ManageLink
}

// GetManageLinkOk returns a tuple with the ManageLink field value
// and a boolean to check if the value has been set.
func (o *VideoParentFolder) GetManageLinkOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ManageLink, true
}

// SetManageLink sets field value
func (o *VideoParentFolder) SetManageLink(v string) {
	o.ManageLink = v
}

// GetMetadata returns the Metadata field value
func (o *VideoParentFolder) GetMetadata() ProjectMetadata {
	if o == nil {
		var ret ProjectMetadata
		return ret
	}

	return o.Metadata
}

// GetMetadataOk returns a tuple with the Metadata field value
// and a boolean to check if the value has been set.
func (o *VideoParentFolder) GetMetadataOk() (*ProjectMetadata, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Metadata, true
}

// SetMetadata sets field value
func (o *VideoParentFolder) SetMetadata(v ProjectMetadata) {
	o.Metadata = v
}

// GetModifiedTime returns the ModifiedTime field value
func (o *VideoParentFolder) GetModifiedTime() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.ModifiedTime
}

// GetModifiedTimeOk returns a tuple with the ModifiedTime field value
// and a boolean to check if the value has been set.
func (o *VideoParentFolder) GetModifiedTimeOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ModifiedTime, true
}

// SetModifiedTime sets field value
func (o *VideoParentFolder) SetModifiedTime(v string) {
	o.ModifiedTime = v
}

// GetName returns the Name field value
func (o *VideoParentFolder) GetName() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Name
}

// GetNameOk returns a tuple with the Name field value
// and a boolean to check if the value has been set.
func (o *VideoParentFolder) GetNameOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Name, true
}

// SetName sets field value
func (o *VideoParentFolder) SetName(v string) {
	o.Name = v
}

// GetPinnedOn returns the PinnedOn field value
// If the value is explicit nil, the zero value for string will be returned
func (o *VideoParentFolder) GetPinnedOn() string {
	if o == nil || o.PinnedOn.Get() == nil {
		var ret string
		return ret
	}

	return *o.PinnedOn.Get()
}

// GetPinnedOnOk returns a tuple with the PinnedOn field value
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *VideoParentFolder) GetPinnedOnOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.PinnedOn.Get(), o.PinnedOn.IsSet()
}

// SetPinnedOn sets field value
func (o *VideoParentFolder) SetPinnedOn(v string) {
	o.PinnedOn.Set(&v)
}

// GetPrivacy returns the Privacy field value
func (o *VideoParentFolder) GetPrivacy() ProjectPrivacy {
	if o == nil {
		var ret ProjectPrivacy
		return ret
	}

	return o.Privacy
}

// GetPrivacyOk returns a tuple with the Privacy field value
// and a boolean to check if the value has been set.
func (o *VideoParentFolder) GetPrivacyOk() (*ProjectPrivacy, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Privacy, true
}

// SetPrivacy sets field value
func (o *VideoParentFolder) SetPrivacy(v ProjectPrivacy) {
	o.Privacy = v
}

// GetResourceKey returns the ResourceKey field value
func (o *VideoParentFolder) GetResourceKey() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.ResourceKey
}

// GetResourceKeyOk returns a tuple with the ResourceKey field value
// and a boolean to check if the value has been set.
func (o *VideoParentFolder) GetResourceKeyOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ResourceKey, true
}

// SetResourceKey sets field value
func (o *VideoParentFolder) SetResourceKey(v string) {
	o.ResourceKey = v
}

// GetSettings returns the Settings field value
func (o *VideoParentFolder) GetSettings() ProjectSettings {
	if o == nil {
		var ret ProjectSettings
		return ret
	}

	return o.Settings
}

// GetSettingsOk returns a tuple with the Settings field value
// and a boolean to check if the value has been set.
func (o *VideoParentFolder) GetSettingsOk() (*ProjectSettings, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Settings, true
}

// SetSettings sets field value
func (o *VideoParentFolder) SetSettings(v ProjectSettings) {
	o.Settings = v
}

// GetUri returns the Uri field value
func (o *VideoParentFolder) GetUri() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Uri
}

// GetUriOk returns a tuple with the Uri field value
// and a boolean to check if the value has been set.
func (o *VideoParentFolder) GetUriOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Uri, true
}

// SetUri sets field value
func (o *VideoParentFolder) SetUri(v string) {
	o.Uri = v
}

// GetUser returns the User field value
func (o *VideoParentFolder) GetUser() User {
	if o == nil {
		var ret User
		return ret
	}

	return o.User
}

// GetUserOk returns a tuple with the User field value
// and a boolean to check if the value has been set.
func (o *VideoParentFolder) GetUserOk() (*User, bool) {
	if o == nil {
		return nil, false
	}
	return &o.User, true
}

// SetUser sets field value
func (o *VideoParentFolder) SetUser(v User) {
	o.User = v
}

func (o VideoParentFolder) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o VideoParentFolder) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["access_grant"] = o.AccessGrant
	toSerialize["created_time"] = o.CreatedTime
	toSerialize["creator_uri"] = o.CreatorUri
	toSerialize["has_subfolder"] = o.HasSubfolder
	toSerialize["is_pinned"] = o.IsPinned
	toSerialize["is_private_to_user"] = o.IsPrivateToUser
	toSerialize["last_user_action_event_date"] = o.LastUserActionEventDate.Get()
	toSerialize["link"] = o.Link
	toSerialize["manage_link"] = o.ManageLink
	toSerialize["metadata"] = o.Metadata
	toSerialize["modified_time"] = o.ModifiedTime
	toSerialize["name"] = o.Name
	toSerialize["pinned_on"] = o.PinnedOn.Get()
	toSerialize["privacy"] = o.Privacy
	toSerialize["resource_key"] = o.ResourceKey
	toSerialize["settings"] = o.Settings
	toSerialize["uri"] = o.Uri
	toSerialize["user"] = o.User
	return toSerialize, nil
}

func (o *VideoParentFolder) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"access_grant",
		"created_time",
		"creator_uri",
		"has_subfolder",
		"is_pinned",
		"is_private_to_user",
		"last_user_action_event_date",
		"link",
		"manage_link",
		"metadata",
		"modified_time",
		"name",
		"pinned_on",
		"privacy",
		"resource_key",
		"settings",
		"uri",
		"user",
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

	varVideoParentFolder := _VideoParentFolder{}

	decoder := json.NewDecoder(bytes.NewReader(data))
	decoder.DisallowUnknownFields()
	err = decoder.Decode(&varVideoParentFolder)

	if err != nil {
		return err
	}

	*o = VideoParentFolder(varVideoParentFolder)

	return err
}

type NullableVideoParentFolder struct {
	value *VideoParentFolder
	isSet bool
}

func (v NullableVideoParentFolder) Get() *VideoParentFolder {
	return v.value
}

func (v *NullableVideoParentFolder) Set(val *VideoParentFolder) {
	v.value = val
	v.isSet = true
}

func (v NullableVideoParentFolder) IsSet() bool {
	return v.isSet
}

func (v *NullableVideoParentFolder) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableVideoParentFolder(val *VideoParentFolder) *NullableVideoParentFolder {
	return &NullableVideoParentFolder{value: val, isSet: true}
}

func (v NullableVideoParentFolder) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableVideoParentFolder) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


