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

// checks if the TeamMembershipMetadataConnectionsOwner type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &TeamMembershipMetadataConnectionsOwner{}

// TeamMembershipMetadataConnectionsOwner A standard connection object indicating how to get the owner of this user.
type TeamMembershipMetadataConnectionsOwner struct {
	// The display name of the team owner.
	DisplayName string `json:"display_name"`
	// The team owner's email address.
	Email string `json:"email"`
	// The total number of remaining team member invites.
	InvitesRemaining float32 `json:"invites_remaining"`
	// An array of HTTP methods permitted on this URI.
	Options []string `json:"options"`
	// The total number of owners on this connection.
	Total float32 `json:"total"`
	// The total number of team members for the specified team owner.
	TotalMembers float32 `json:"total_members"`
	// The API URI that resolves to the connection data.
	Uri string `json:"uri"`
}

// NewTeamMembershipMetadataConnectionsOwner instantiates a new TeamMembershipMetadataConnectionsOwner object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewTeamMembershipMetadataConnectionsOwner(displayName string, email string, invitesRemaining float32, options []string, total float32, totalMembers float32, uri string) *TeamMembershipMetadataConnectionsOwner {
	this := TeamMembershipMetadataConnectionsOwner{}
	this.DisplayName = displayName
	this.Email = email
	this.InvitesRemaining = invitesRemaining
	this.Options = options
	this.Total = total
	this.TotalMembers = totalMembers
	this.Uri = uri
	return &this
}

// NewTeamMembershipMetadataConnectionsOwnerWithDefaults instantiates a new TeamMembershipMetadataConnectionsOwner object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewTeamMembershipMetadataConnectionsOwnerWithDefaults() *TeamMembershipMetadataConnectionsOwner {
	this := TeamMembershipMetadataConnectionsOwner{}
	return &this
}

// GetDisplayName returns the DisplayName field value
func (o *TeamMembershipMetadataConnectionsOwner) GetDisplayName() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.DisplayName
}

// GetDisplayNameOk returns a tuple with the DisplayName field value
// and a boolean to check if the value has been set.
func (o *TeamMembershipMetadataConnectionsOwner) GetDisplayNameOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.DisplayName, true
}

// SetDisplayName sets field value
func (o *TeamMembershipMetadataConnectionsOwner) SetDisplayName(v string) {
	o.DisplayName = v
}

// GetEmail returns the Email field value
func (o *TeamMembershipMetadataConnectionsOwner) GetEmail() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Email
}

// GetEmailOk returns a tuple with the Email field value
// and a boolean to check if the value has been set.
func (o *TeamMembershipMetadataConnectionsOwner) GetEmailOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Email, true
}

// SetEmail sets field value
func (o *TeamMembershipMetadataConnectionsOwner) SetEmail(v string) {
	o.Email = v
}

// GetInvitesRemaining returns the InvitesRemaining field value
func (o *TeamMembershipMetadataConnectionsOwner) GetInvitesRemaining() float32 {
	if o == nil {
		var ret float32
		return ret
	}

	return o.InvitesRemaining
}

// GetInvitesRemainingOk returns a tuple with the InvitesRemaining field value
// and a boolean to check if the value has been set.
func (o *TeamMembershipMetadataConnectionsOwner) GetInvitesRemainingOk() (*float32, bool) {
	if o == nil {
		return nil, false
	}
	return &o.InvitesRemaining, true
}

// SetInvitesRemaining sets field value
func (o *TeamMembershipMetadataConnectionsOwner) SetInvitesRemaining(v float32) {
	o.InvitesRemaining = v
}

// GetOptions returns the Options field value
func (o *TeamMembershipMetadataConnectionsOwner) GetOptions() []string {
	if o == nil {
		var ret []string
		return ret
	}

	return o.Options
}

// GetOptionsOk returns a tuple with the Options field value
// and a boolean to check if the value has been set.
func (o *TeamMembershipMetadataConnectionsOwner) GetOptionsOk() ([]string, bool) {
	if o == nil {
		return nil, false
	}
	return o.Options, true
}

// SetOptions sets field value
func (o *TeamMembershipMetadataConnectionsOwner) SetOptions(v []string) {
	o.Options = v
}

// GetTotal returns the Total field value
func (o *TeamMembershipMetadataConnectionsOwner) GetTotal() float32 {
	if o == nil {
		var ret float32
		return ret
	}

	return o.Total
}

// GetTotalOk returns a tuple with the Total field value
// and a boolean to check if the value has been set.
func (o *TeamMembershipMetadataConnectionsOwner) GetTotalOk() (*float32, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Total, true
}

// SetTotal sets field value
func (o *TeamMembershipMetadataConnectionsOwner) SetTotal(v float32) {
	o.Total = v
}

// GetTotalMembers returns the TotalMembers field value
func (o *TeamMembershipMetadataConnectionsOwner) GetTotalMembers() float32 {
	if o == nil {
		var ret float32
		return ret
	}

	return o.TotalMembers
}

// GetTotalMembersOk returns a tuple with the TotalMembers field value
// and a boolean to check if the value has been set.
func (o *TeamMembershipMetadataConnectionsOwner) GetTotalMembersOk() (*float32, bool) {
	if o == nil {
		return nil, false
	}
	return &o.TotalMembers, true
}

// SetTotalMembers sets field value
func (o *TeamMembershipMetadataConnectionsOwner) SetTotalMembers(v float32) {
	o.TotalMembers = v
}

// GetUri returns the Uri field value
func (o *TeamMembershipMetadataConnectionsOwner) GetUri() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Uri
}

// GetUriOk returns a tuple with the Uri field value
// and a boolean to check if the value has been set.
func (o *TeamMembershipMetadataConnectionsOwner) GetUriOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Uri, true
}

// SetUri sets field value
func (o *TeamMembershipMetadataConnectionsOwner) SetUri(v string) {
	o.Uri = v
}

func (o TeamMembershipMetadataConnectionsOwner) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o TeamMembershipMetadataConnectionsOwner) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["display_name"] = o.DisplayName
	toSerialize["email"] = o.Email
	toSerialize["invites_remaining"] = o.InvitesRemaining
	toSerialize["options"] = o.Options
	toSerialize["total"] = o.Total
	toSerialize["total_members"] = o.TotalMembers
	toSerialize["uri"] = o.Uri
	return toSerialize, nil
}

type NullableTeamMembershipMetadataConnectionsOwner struct {
	value *TeamMembershipMetadataConnectionsOwner
	isSet bool
}

func (v NullableTeamMembershipMetadataConnectionsOwner) Get() *TeamMembershipMetadataConnectionsOwner {
	return v.value
}

func (v *NullableTeamMembershipMetadataConnectionsOwner) Set(val *TeamMembershipMetadataConnectionsOwner) {
	v.value = val
	v.isSet = true
}

func (v NullableTeamMembershipMetadataConnectionsOwner) IsSet() bool {
	return v.isSet
}

func (v *NullableTeamMembershipMetadataConnectionsOwner) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableTeamMembershipMetadataConnectionsOwner(val *TeamMembershipMetadataConnectionsOwner) *NullableTeamMembershipMetadataConnectionsOwner {
	return &NullableTeamMembershipMetadataConnectionsOwner{value: val, isSet: true}
}

func (v NullableTeamMembershipMetadataConnectionsOwner) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableTeamMembershipMetadataConnectionsOwner) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


