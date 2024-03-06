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

// checks if the TeamMembership type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &TeamMembership{}

// TeamMembership struct for TeamMembership
type TeamMembership struct {
	// Information about an access grant that applies to the team member on the folder. _This field is deprecated because grants are no longer exposed via API responses._
	AccessGrant map[string]interface{} `json:"access_grant"`
	// Whether the team membership is currently active.
	Active                       bool                                       `json:"active"`
	ApplicablePermissionPolicies TeamMembershipApplicablePermissionPolicies `json:"applicable_permission_policies"`
	// The time in ISO 8601 format when the invite was sent.
	CreatedTime string `json:"created_time"`
	// The team member's email.
	Email string `json:"email"`
	// Whether the team member has folder access.
	HasFolderAccess *bool `json:"has_folder_access,omitempty"`
	// The URL for the invited user to join the team. The value of this field is null if the invited user has already joined.
	InviteUrl NullableString `json:"invite_url"`
	// The time in ISO 8601 format when the invite was accepted.
	JoinedTime string                 `json:"joined_time"`
	Metadata   TeamMembershipMetadata `json:"metadata"`
	// The time in ISO 8601 format when the team membership was last modified.
	ModifiedTime string `json:"modified_time"`
	// The team member's permission level.  Option descriptions:  * `Admin` - The team member has admin permissions. They can upload and edit videos for the entire team and perform team administration tasks.  * `Contributor` - The team member has contributor permissions. They can upload and edit videos for the entire team but can't perform team administration tasks.  * `Contributor Plus` - The team member has contributor plus permissions. They can upload and edit videos for the entire team, and have additional sets of permissions, but can't perform team administration tasks.  * `Owner` - The team member has owner permissions.  * `Uploader` - The team member has uploader permissions. They can upload videos for the entire team but can't edit videos.  * `Viewer` - The team member has viewer permissions. They can access team videos and specific team folders but can't upload or edit videos.
	PermissionLevel string `json:"permission_level"`
	// Whether the team member has been reminded about the invite.
	RecentlyReminded *bool `json:"recently_reminded,omitempty"`
	// The resource key of the team membership.
	ResourceKey string `json:"resource_key"`
	// The team member's role, translated.
	Role string `json:"role"`
	// The status of the team membership invite.  Option descriptions:  * `accepted` - Team membership has been accepted.  * `pending` - Team membership has been offered but not yet accepted.
	Status string `json:"status"`
	// The unique identifier to access the team membership resource.
	Uri  string                     `json:"uri"`
	User NullableTeamMembershipUser `json:"user"`
}

// NewTeamMembership instantiates a new TeamMembership object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewTeamMembership(accessGrant map[string]interface{}, active bool, applicablePermissionPolicies TeamMembershipApplicablePermissionPolicies, createdTime string, email string, inviteUrl NullableString, joinedTime string, metadata TeamMembershipMetadata, modifiedTime string, permissionLevel string, resourceKey string, role string, status string, uri string, user NullableTeamMembershipUser) *TeamMembership {
	this := TeamMembership{}
	this.AccessGrant = accessGrant
	this.Active = active
	this.ApplicablePermissionPolicies = applicablePermissionPolicies
	this.CreatedTime = createdTime
	this.Email = email
	this.InviteUrl = inviteUrl
	this.JoinedTime = joinedTime
	this.Metadata = metadata
	this.ModifiedTime = modifiedTime
	this.PermissionLevel = permissionLevel
	this.ResourceKey = resourceKey
	this.Role = role
	this.Status = status
	this.Uri = uri
	this.User = user
	return &this
}

// NewTeamMembershipWithDefaults instantiates a new TeamMembership object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewTeamMembershipWithDefaults() *TeamMembership {
	this := TeamMembership{}
	return &this
}

// GetAccessGrant returns the AccessGrant field value
// If the value is explicit nil, the zero value for map[string]interface{} will be returned
func (o *TeamMembership) GetAccessGrant() map[string]interface{} {
	if o == nil {
		var ret map[string]interface{}
		return ret
	}

	return o.AccessGrant
}

// GetAccessGrantOk returns a tuple with the AccessGrant field value
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *TeamMembership) GetAccessGrantOk() (map[string]interface{}, bool) {
	if o == nil || IsNil(o.AccessGrant) {
		return map[string]interface{}{}, false
	}
	return o.AccessGrant, true
}

// SetAccessGrant sets field value
func (o *TeamMembership) SetAccessGrant(v map[string]interface{}) {
	o.AccessGrant = v
}

// GetActive returns the Active field value
func (o *TeamMembership) GetActive() bool {
	if o == nil {
		var ret bool
		return ret
	}

	return o.Active
}

// GetActiveOk returns a tuple with the Active field value
// and a boolean to check if the value has been set.
func (o *TeamMembership) GetActiveOk() (*bool, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Active, true
}

// SetActive sets field value
func (o *TeamMembership) SetActive(v bool) {
	o.Active = v
}

// GetApplicablePermissionPolicies returns the ApplicablePermissionPolicies field value
func (o *TeamMembership) GetApplicablePermissionPolicies() TeamMembershipApplicablePermissionPolicies {
	if o == nil {
		var ret TeamMembershipApplicablePermissionPolicies
		return ret
	}

	return o.ApplicablePermissionPolicies
}

// GetApplicablePermissionPoliciesOk returns a tuple with the ApplicablePermissionPolicies field value
// and a boolean to check if the value has been set.
func (o *TeamMembership) GetApplicablePermissionPoliciesOk() (*TeamMembershipApplicablePermissionPolicies, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ApplicablePermissionPolicies, true
}

// SetApplicablePermissionPolicies sets field value
func (o *TeamMembership) SetApplicablePermissionPolicies(v TeamMembershipApplicablePermissionPolicies) {
	o.ApplicablePermissionPolicies = v
}

// GetCreatedTime returns the CreatedTime field value
func (o *TeamMembership) GetCreatedTime() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.CreatedTime
}

// GetCreatedTimeOk returns a tuple with the CreatedTime field value
// and a boolean to check if the value has been set.
func (o *TeamMembership) GetCreatedTimeOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.CreatedTime, true
}

// SetCreatedTime sets field value
func (o *TeamMembership) SetCreatedTime(v string) {
	o.CreatedTime = v
}

// GetEmail returns the Email field value
func (o *TeamMembership) GetEmail() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Email
}

// GetEmailOk returns a tuple with the Email field value
// and a boolean to check if the value has been set.
func (o *TeamMembership) GetEmailOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Email, true
}

// SetEmail sets field value
func (o *TeamMembership) SetEmail(v string) {
	o.Email = v
}

// GetHasFolderAccess returns the HasFolderAccess field value if set, zero value otherwise.
func (o *TeamMembership) GetHasFolderAccess() bool {
	if o == nil || IsNil(o.HasFolderAccess) {
		var ret bool
		return ret
	}
	return *o.HasFolderAccess
}

// GetHasFolderAccessOk returns a tuple with the HasFolderAccess field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *TeamMembership) GetHasFolderAccessOk() (*bool, bool) {
	if o == nil || IsNil(o.HasFolderAccess) {
		return nil, false
	}
	return o.HasFolderAccess, true
}

// HasHasFolderAccess returns a boolean if a field has been set.
func (o *TeamMembership) HasHasFolderAccess() bool {
	if o != nil && !IsNil(o.HasFolderAccess) {
		return true
	}

	return false
}

// SetHasFolderAccess gets a reference to the given bool and assigns it to the HasFolderAccess field.
func (o *TeamMembership) SetHasFolderAccess(v bool) {
	o.HasFolderAccess = &v
}

// GetInviteUrl returns the InviteUrl field value
// If the value is explicit nil, the zero value for string will be returned
func (o *TeamMembership) GetInviteUrl() string {
	if o == nil || o.InviteUrl.Get() == nil {
		var ret string
		return ret
	}

	return *o.InviteUrl.Get()
}

// GetInviteUrlOk returns a tuple with the InviteUrl field value
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *TeamMembership) GetInviteUrlOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.InviteUrl.Get(), o.InviteUrl.IsSet()
}

// SetInviteUrl sets field value
func (o *TeamMembership) SetInviteUrl(v string) {
	o.InviteUrl.Set(&v)
}

// GetJoinedTime returns the JoinedTime field value
func (o *TeamMembership) GetJoinedTime() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.JoinedTime
}

// GetJoinedTimeOk returns a tuple with the JoinedTime field value
// and a boolean to check if the value has been set.
func (o *TeamMembership) GetJoinedTimeOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.JoinedTime, true
}

// SetJoinedTime sets field value
func (o *TeamMembership) SetJoinedTime(v string) {
	o.JoinedTime = v
}

// GetMetadata returns the Metadata field value
func (o *TeamMembership) GetMetadata() TeamMembershipMetadata {
	if o == nil {
		var ret TeamMembershipMetadata
		return ret
	}

	return o.Metadata
}

// GetMetadataOk returns a tuple with the Metadata field value
// and a boolean to check if the value has been set.
func (o *TeamMembership) GetMetadataOk() (*TeamMembershipMetadata, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Metadata, true
}

// SetMetadata sets field value
func (o *TeamMembership) SetMetadata(v TeamMembershipMetadata) {
	o.Metadata = v
}

// GetModifiedTime returns the ModifiedTime field value
func (o *TeamMembership) GetModifiedTime() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.ModifiedTime
}

// GetModifiedTimeOk returns a tuple with the ModifiedTime field value
// and a boolean to check if the value has been set.
func (o *TeamMembership) GetModifiedTimeOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ModifiedTime, true
}

// SetModifiedTime sets field value
func (o *TeamMembership) SetModifiedTime(v string) {
	o.ModifiedTime = v
}

// GetPermissionLevel returns the PermissionLevel field value
func (o *TeamMembership) GetPermissionLevel() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.PermissionLevel
}

// GetPermissionLevelOk returns a tuple with the PermissionLevel field value
// and a boolean to check if the value has been set.
func (o *TeamMembership) GetPermissionLevelOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.PermissionLevel, true
}

// SetPermissionLevel sets field value
func (o *TeamMembership) SetPermissionLevel(v string) {
	o.PermissionLevel = v
}

// GetRecentlyReminded returns the RecentlyReminded field value if set, zero value otherwise.
func (o *TeamMembership) GetRecentlyReminded() bool {
	if o == nil || IsNil(o.RecentlyReminded) {
		var ret bool
		return ret
	}
	return *o.RecentlyReminded
}

// GetRecentlyRemindedOk returns a tuple with the RecentlyReminded field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *TeamMembership) GetRecentlyRemindedOk() (*bool, bool) {
	if o == nil || IsNil(o.RecentlyReminded) {
		return nil, false
	}
	return o.RecentlyReminded, true
}

// HasRecentlyReminded returns a boolean if a field has been set.
func (o *TeamMembership) HasRecentlyReminded() bool {
	if o != nil && !IsNil(o.RecentlyReminded) {
		return true
	}

	return false
}

// SetRecentlyReminded gets a reference to the given bool and assigns it to the RecentlyReminded field.
func (o *TeamMembership) SetRecentlyReminded(v bool) {
	o.RecentlyReminded = &v
}

// GetResourceKey returns the ResourceKey field value
func (o *TeamMembership) GetResourceKey() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.ResourceKey
}

// GetResourceKeyOk returns a tuple with the ResourceKey field value
// and a boolean to check if the value has been set.
func (o *TeamMembership) GetResourceKeyOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ResourceKey, true
}

// SetResourceKey sets field value
func (o *TeamMembership) SetResourceKey(v string) {
	o.ResourceKey = v
}

// GetRole returns the Role field value
func (o *TeamMembership) GetRole() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Role
}

// GetRoleOk returns a tuple with the Role field value
// and a boolean to check if the value has been set.
func (o *TeamMembership) GetRoleOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Role, true
}

// SetRole sets field value
func (o *TeamMembership) SetRole(v string) {
	o.Role = v
}

// GetStatus returns the Status field value
func (o *TeamMembership) GetStatus() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Status
}

// GetStatusOk returns a tuple with the Status field value
// and a boolean to check if the value has been set.
func (o *TeamMembership) GetStatusOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Status, true
}

// SetStatus sets field value
func (o *TeamMembership) SetStatus(v string) {
	o.Status = v
}

// GetUri returns the Uri field value
func (o *TeamMembership) GetUri() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Uri
}

// GetUriOk returns a tuple with the Uri field value
// and a boolean to check if the value has been set.
func (o *TeamMembership) GetUriOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Uri, true
}

// SetUri sets field value
func (o *TeamMembership) SetUri(v string) {
	o.Uri = v
}

// GetUser returns the User field value
// If the value is explicit nil, the zero value for TeamMembershipUser will be returned
func (o *TeamMembership) GetUser() TeamMembershipUser {
	if o == nil || o.User.Get() == nil {
		var ret TeamMembershipUser
		return ret
	}

	return *o.User.Get()
}

// GetUserOk returns a tuple with the User field value
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *TeamMembership) GetUserOk() (*TeamMembershipUser, bool) {
	if o == nil {
		return nil, false
	}
	return o.User.Get(), o.User.IsSet()
}

// SetUser sets field value
func (o *TeamMembership) SetUser(v TeamMembershipUser) {
	o.User.Set(&v)
}

func (o TeamMembership) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o TeamMembership) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if o.AccessGrant != nil {
		toSerialize["access_grant"] = o.AccessGrant
	}
	toSerialize["active"] = o.Active
	toSerialize["applicable_permission_policies"] = o.ApplicablePermissionPolicies
	toSerialize["created_time"] = o.CreatedTime
	toSerialize["email"] = o.Email
	if !IsNil(o.HasFolderAccess) {
		toSerialize["has_folder_access"] = o.HasFolderAccess
	}
	toSerialize["invite_url"] = o.InviteUrl.Get()
	toSerialize["joined_time"] = o.JoinedTime
	toSerialize["metadata"] = o.Metadata
	toSerialize["modified_time"] = o.ModifiedTime
	toSerialize["permission_level"] = o.PermissionLevel
	if !IsNil(o.RecentlyReminded) {
		toSerialize["recently_reminded"] = o.RecentlyReminded
	}
	toSerialize["resource_key"] = o.ResourceKey
	toSerialize["role"] = o.Role
	toSerialize["status"] = o.Status
	toSerialize["uri"] = o.Uri
	toSerialize["user"] = o.User.Get()
	return toSerialize, nil
}

type NullableTeamMembership struct {
	value *TeamMembership
	isSet bool
}

func (v NullableTeamMembership) Get() *TeamMembership {
	return v.value
}

func (v *NullableTeamMembership) Set(val *TeamMembership) {
	v.value = val
	v.isSet = true
}

func (v NullableTeamMembership) IsSet() bool {
	return v.isSet
}

func (v *NullableTeamMembership) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableTeamMembership(val *TeamMembership) *NullableTeamMembership {
	return &NullableTeamMembership{value: val, isSet: true}
}

func (v NullableTeamMembership) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableTeamMembership) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
