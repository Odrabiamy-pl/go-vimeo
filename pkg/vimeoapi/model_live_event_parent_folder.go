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

// checks if the LiveEventParentFolder type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &LiveEventParentFolder{}

// LiveEventParentFolder Information about the folder that contains the event.
type LiveEventParentFolder struct {
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
	User User `json:"user"`
}

// NewLiveEventParentFolder instantiates a new LiveEventParentFolder object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewLiveEventParentFolder(accessGrant map[string]interface{}, createdTime string, creatorUri string, hasSubfolder bool, isPinned bool, isPrivateToUser bool, lastUserActionEventDate NullableString, link string, manageLink string, metadata ProjectMetadata, modifiedTime string, name string, pinnedOn NullableString, privacy ProjectPrivacy, resourceKey string, settings ProjectSettings, uri string, user User) *LiveEventParentFolder {
	this := LiveEventParentFolder{}
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

// NewLiveEventParentFolderWithDefaults instantiates a new LiveEventParentFolder object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewLiveEventParentFolderWithDefaults() *LiveEventParentFolder {
	this := LiveEventParentFolder{}
	return &this
}

// GetAccessGrant returns the AccessGrant field value
func (o *LiveEventParentFolder) GetAccessGrant() map[string]interface{} {
	if o == nil {
		var ret map[string]interface{}
		return ret
	}

	return o.AccessGrant
}

// GetAccessGrantOk returns a tuple with the AccessGrant field value
// and a boolean to check if the value has been set.
func (o *LiveEventParentFolder) GetAccessGrantOk() (map[string]interface{}, bool) {
	if o == nil {
		return map[string]interface{}{}, false
	}
	return o.AccessGrant, true
}

// SetAccessGrant sets field value
func (o *LiveEventParentFolder) SetAccessGrant(v map[string]interface{}) {
	o.AccessGrant = v
}

// GetCreatedTime returns the CreatedTime field value
func (o *LiveEventParentFolder) GetCreatedTime() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.CreatedTime
}

// GetCreatedTimeOk returns a tuple with the CreatedTime field value
// and a boolean to check if the value has been set.
func (o *LiveEventParentFolder) GetCreatedTimeOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.CreatedTime, true
}

// SetCreatedTime sets field value
func (o *LiveEventParentFolder) SetCreatedTime(v string) {
	o.CreatedTime = v
}

// GetCreatorUri returns the CreatorUri field value
func (o *LiveEventParentFolder) GetCreatorUri() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.CreatorUri
}

// GetCreatorUriOk returns a tuple with the CreatorUri field value
// and a boolean to check if the value has been set.
func (o *LiveEventParentFolder) GetCreatorUriOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.CreatorUri, true
}

// SetCreatorUri sets field value
func (o *LiveEventParentFolder) SetCreatorUri(v string) {
	o.CreatorUri = v
}

// GetHasSubfolder returns the HasSubfolder field value
func (o *LiveEventParentFolder) GetHasSubfolder() bool {
	if o == nil {
		var ret bool
		return ret
	}

	return o.HasSubfolder
}

// GetHasSubfolderOk returns a tuple with the HasSubfolder field value
// and a boolean to check if the value has been set.
func (o *LiveEventParentFolder) GetHasSubfolderOk() (*bool, bool) {
	if o == nil {
		return nil, false
	}
	return &o.HasSubfolder, true
}

// SetHasSubfolder sets field value
func (o *LiveEventParentFolder) SetHasSubfolder(v bool) {
	o.HasSubfolder = v
}

// GetIsPinned returns the IsPinned field value
func (o *LiveEventParentFolder) GetIsPinned() bool {
	if o == nil {
		var ret bool
		return ret
	}

	return o.IsPinned
}

// GetIsPinnedOk returns a tuple with the IsPinned field value
// and a boolean to check if the value has been set.
func (o *LiveEventParentFolder) GetIsPinnedOk() (*bool, bool) {
	if o == nil {
		return nil, false
	}
	return &o.IsPinned, true
}

// SetIsPinned sets field value
func (o *LiveEventParentFolder) SetIsPinned(v bool) {
	o.IsPinned = v
}

// GetIsPrivateToUser returns the IsPrivateToUser field value
func (o *LiveEventParentFolder) GetIsPrivateToUser() bool {
	if o == nil {
		var ret bool
		return ret
	}

	return o.IsPrivateToUser
}

// GetIsPrivateToUserOk returns a tuple with the IsPrivateToUser field value
// and a boolean to check if the value has been set.
func (o *LiveEventParentFolder) GetIsPrivateToUserOk() (*bool, bool) {
	if o == nil {
		return nil, false
	}
	return &o.IsPrivateToUser, true
}

// SetIsPrivateToUser sets field value
func (o *LiveEventParentFolder) SetIsPrivateToUser(v bool) {
	o.IsPrivateToUser = v
}

// GetLastUserActionEventDate returns the LastUserActionEventDate field value
// If the value is explicit nil, the zero value for string will be returned
func (o *LiveEventParentFolder) GetLastUserActionEventDate() string {
	if o == nil || o.LastUserActionEventDate.Get() == nil {
		var ret string
		return ret
	}

	return *o.LastUserActionEventDate.Get()
}

// GetLastUserActionEventDateOk returns a tuple with the LastUserActionEventDate field value
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *LiveEventParentFolder) GetLastUserActionEventDateOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.LastUserActionEventDate.Get(), o.LastUserActionEventDate.IsSet()
}

// SetLastUserActionEventDate sets field value
func (o *LiveEventParentFolder) SetLastUserActionEventDate(v string) {
	o.LastUserActionEventDate.Set(&v)
}

// GetLink returns the Link field value
func (o *LiveEventParentFolder) GetLink() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Link
}

// GetLinkOk returns a tuple with the Link field value
// and a boolean to check if the value has been set.
func (o *LiveEventParentFolder) GetLinkOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Link, true
}

// SetLink sets field value
func (o *LiveEventParentFolder) SetLink(v string) {
	o.Link = v
}

// GetManageLink returns the ManageLink field value
func (o *LiveEventParentFolder) GetManageLink() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.ManageLink
}

// GetManageLinkOk returns a tuple with the ManageLink field value
// and a boolean to check if the value has been set.
func (o *LiveEventParentFolder) GetManageLinkOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ManageLink, true
}

// SetManageLink sets field value
func (o *LiveEventParentFolder) SetManageLink(v string) {
	o.ManageLink = v
}

// GetMetadata returns the Metadata field value
func (o *LiveEventParentFolder) GetMetadata() ProjectMetadata {
	if o == nil {
		var ret ProjectMetadata
		return ret
	}

	return o.Metadata
}

// GetMetadataOk returns a tuple with the Metadata field value
// and a boolean to check if the value has been set.
func (o *LiveEventParentFolder) GetMetadataOk() (*ProjectMetadata, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Metadata, true
}

// SetMetadata sets field value
func (o *LiveEventParentFolder) SetMetadata(v ProjectMetadata) {
	o.Metadata = v
}

// GetModifiedTime returns the ModifiedTime field value
func (o *LiveEventParentFolder) GetModifiedTime() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.ModifiedTime
}

// GetModifiedTimeOk returns a tuple with the ModifiedTime field value
// and a boolean to check if the value has been set.
func (o *LiveEventParentFolder) GetModifiedTimeOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ModifiedTime, true
}

// SetModifiedTime sets field value
func (o *LiveEventParentFolder) SetModifiedTime(v string) {
	o.ModifiedTime = v
}

// GetName returns the Name field value
func (o *LiveEventParentFolder) GetName() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Name
}

// GetNameOk returns a tuple with the Name field value
// and a boolean to check if the value has been set.
func (o *LiveEventParentFolder) GetNameOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Name, true
}

// SetName sets field value
func (o *LiveEventParentFolder) SetName(v string) {
	o.Name = v
}

// GetPinnedOn returns the PinnedOn field value
// If the value is explicit nil, the zero value for string will be returned
func (o *LiveEventParentFolder) GetPinnedOn() string {
	if o == nil || o.PinnedOn.Get() == nil {
		var ret string
		return ret
	}

	return *o.PinnedOn.Get()
}

// GetPinnedOnOk returns a tuple with the PinnedOn field value
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *LiveEventParentFolder) GetPinnedOnOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.PinnedOn.Get(), o.PinnedOn.IsSet()
}

// SetPinnedOn sets field value
func (o *LiveEventParentFolder) SetPinnedOn(v string) {
	o.PinnedOn.Set(&v)
}

// GetPrivacy returns the Privacy field value
func (o *LiveEventParentFolder) GetPrivacy() ProjectPrivacy {
	if o == nil {
		var ret ProjectPrivacy
		return ret
	}

	return o.Privacy
}

// GetPrivacyOk returns a tuple with the Privacy field value
// and a boolean to check if the value has been set.
func (o *LiveEventParentFolder) GetPrivacyOk() (*ProjectPrivacy, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Privacy, true
}

// SetPrivacy sets field value
func (o *LiveEventParentFolder) SetPrivacy(v ProjectPrivacy) {
	o.Privacy = v
}

// GetResourceKey returns the ResourceKey field value
func (o *LiveEventParentFolder) GetResourceKey() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.ResourceKey
}

// GetResourceKeyOk returns a tuple with the ResourceKey field value
// and a boolean to check if the value has been set.
func (o *LiveEventParentFolder) GetResourceKeyOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ResourceKey, true
}

// SetResourceKey sets field value
func (o *LiveEventParentFolder) SetResourceKey(v string) {
	o.ResourceKey = v
}

// GetSettings returns the Settings field value
func (o *LiveEventParentFolder) GetSettings() ProjectSettings {
	if o == nil {
		var ret ProjectSettings
		return ret
	}

	return o.Settings
}

// GetSettingsOk returns a tuple with the Settings field value
// and a boolean to check if the value has been set.
func (o *LiveEventParentFolder) GetSettingsOk() (*ProjectSettings, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Settings, true
}

// SetSettings sets field value
func (o *LiveEventParentFolder) SetSettings(v ProjectSettings) {
	o.Settings = v
}

// GetUri returns the Uri field value
func (o *LiveEventParentFolder) GetUri() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Uri
}

// GetUriOk returns a tuple with the Uri field value
// and a boolean to check if the value has been set.
func (o *LiveEventParentFolder) GetUriOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Uri, true
}

// SetUri sets field value
func (o *LiveEventParentFolder) SetUri(v string) {
	o.Uri = v
}

// GetUser returns the User field value
func (o *LiveEventParentFolder) GetUser() User {
	if o == nil {
		var ret User
		return ret
	}

	return o.User
}

// GetUserOk returns a tuple with the User field value
// and a boolean to check if the value has been set.
func (o *LiveEventParentFolder) GetUserOk() (*User, bool) {
	if o == nil {
		return nil, false
	}
	return &o.User, true
}

// SetUser sets field value
func (o *LiveEventParentFolder) SetUser(v User) {
	o.User = v
}

func (o LiveEventParentFolder) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o LiveEventParentFolder) ToMap() (map[string]interface{}, error) {
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

type NullableLiveEventParentFolder struct {
	value *LiveEventParentFolder
	isSet bool
}

func (v NullableLiveEventParentFolder) Get() *LiveEventParentFolder {
	return v.value
}

func (v *NullableLiveEventParentFolder) Set(val *LiveEventParentFolder) {
	v.value = val
	v.isSet = true
}

func (v NullableLiveEventParentFolder) IsSet() bool {
	return v.isSet
}

func (v *NullableLiveEventParentFolder) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableLiveEventParentFolder(val *LiveEventParentFolder) *NullableLiveEventParentFolder {
	return &NullableLiveEventParentFolder{value: val, isSet: true}
}

func (v NullableLiveEventParentFolder) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableLiveEventParentFolder) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


