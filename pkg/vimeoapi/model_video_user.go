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

// checks if the VideoUser type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &VideoUser{}

// VideoUser The video's owner.
type VideoUser struct {
	// The authenticated user's account type.  Option descriptions:  * `advanced` - The user has a Vimeo Advanced subscription.  * `basic` - The user has a Vimeo Basic subscription.  * `business` - The user has a Vimeo Business subscription.  * `enterprise` - The user has a Vimeo Enterprise subscription.  * `free` - The user has a Vimeo Free subscription.  * `live_business` - The user has a Vimeo Business Live subscription.  * `live_premium` - The user has a Vimeo Premium subscription.  * `live_pro` - The user has a Vimeo PRO Live subscription.  * `plus` - The user has a Vimeo Plus subscription.  * `pro` - The user has a Vimeo Pro subscription.  * `pro_unlimited` - The user has a Vimeo PRO Unlimited subscription.  * `producer` - The user has a Vimeo Producer subscription.  * `standard` - The user has a Vimeo Standard subscription.  * `starter` - The user has a Vimeo Starter subscription. 
	Account string `json:"account"`
	// Whether the authenticated user is available for hire.
	AvailableForHire bool `json:"available_for_hire"`
	// The authenticated user's long bio text.
	Bio NullableString `json:"bio"`
	// Whether the authenticated user can work remotely.
	CanWorkRemotely bool `json:"can_work_remotely"`
	// The users's capabilities list.
	Capabilities map[string]interface{} `json:"capabilities"`
	// The comma-separated list of clients.
	Clients string `json:"clients"`
	// The authenticated user's content filters.  Option descriptions:  * `drugs` - The content contains drug or alcohol use.  * `language` - The content contains profanity or sexually suggestive language.  * `nudity` - The content contains nudity.  * `safe` - The content is suitable for all audiences.  * `unrated` - The content hasn't been rated.  * `violence` - The content contains violence or is graphic. 
	ContentFilter []string `json:"content_filter,omitempty"`
	// The time in ISO 8601 format when the user account was created.
	CreatedTime string `json:"created_time"`
	// The authenticated user's gender.
	Gender NullableString `json:"gender"`
	// Whether the user's email is invalid.
	HasInvalidEmail bool `json:"has_invalid_email"`
	// Whether the creator enrolled in and successfully completed the Vimeo Experts program.
	IsExpert bool `json:"is_expert"`
	// The absolute URL of the authenticated users's profile page.
	Link string `json:"link"`
	// The authenticated user's location.
	Location NullableString `json:"location"`
	LocationDetails NullableUserLocationDetails `json:"location_details"`
	Metadata UserMetadata `json:"metadata"`
	// The authenticated user's display name.
	Name string `json:"name"`
	Pictures Picture `json:"pictures"`
	Preferences UserPreferences `json:"preferences"`
	// The authenticated user's resource key string.
	ResourceKey string `json:"resource_key"`
	// The authenticated user's short bio text.
	ShortBio NullableString `json:"short_bio"`
	// A list of the authenticated user's skills.
	Skills []Skill `json:"skills"`
	UploadQuota UserUploadQuota `json:"upload_quota"`
	// The authenticated user's canonical relative URI.
	Uri string `json:"uri"`
	// The authenticated user's websites.
	Websites []UserWebsitesInner `json:"websites"`
}

// NewVideoUser instantiates a new VideoUser object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewVideoUser(account string, availableForHire bool, bio NullableString, canWorkRemotely bool, capabilities map[string]interface{}, clients string, createdTime string, gender NullableString, hasInvalidEmail bool, isExpert bool, link string, location NullableString, locationDetails NullableUserLocationDetails, metadata UserMetadata, name string, pictures Picture, preferences UserPreferences, resourceKey string, shortBio NullableString, skills []Skill, uploadQuota UserUploadQuota, uri string, websites []UserWebsitesInner) *VideoUser {
	this := VideoUser{}
	this.Account = account
	this.AvailableForHire = availableForHire
	this.Bio = bio
	this.CanWorkRemotely = canWorkRemotely
	this.Capabilities = capabilities
	this.Clients = clients
	this.CreatedTime = createdTime
	this.Gender = gender
	this.HasInvalidEmail = hasInvalidEmail
	this.IsExpert = isExpert
	this.Link = link
	this.Location = location
	this.LocationDetails = locationDetails
	this.Metadata = metadata
	this.Name = name
	this.Pictures = pictures
	this.Preferences = preferences
	this.ResourceKey = resourceKey
	this.ShortBio = shortBio
	this.Skills = skills
	this.UploadQuota = uploadQuota
	this.Uri = uri
	this.Websites = websites
	return &this
}

// NewVideoUserWithDefaults instantiates a new VideoUser object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewVideoUserWithDefaults() *VideoUser {
	this := VideoUser{}
	return &this
}

// GetAccount returns the Account field value
func (o *VideoUser) GetAccount() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Account
}

// GetAccountOk returns a tuple with the Account field value
// and a boolean to check if the value has been set.
func (o *VideoUser) GetAccountOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Account, true
}

// SetAccount sets field value
func (o *VideoUser) SetAccount(v string) {
	o.Account = v
}

// GetAvailableForHire returns the AvailableForHire field value
func (o *VideoUser) GetAvailableForHire() bool {
	if o == nil {
		var ret bool
		return ret
	}

	return o.AvailableForHire
}

// GetAvailableForHireOk returns a tuple with the AvailableForHire field value
// and a boolean to check if the value has been set.
func (o *VideoUser) GetAvailableForHireOk() (*bool, bool) {
	if o == nil {
		return nil, false
	}
	return &o.AvailableForHire, true
}

// SetAvailableForHire sets field value
func (o *VideoUser) SetAvailableForHire(v bool) {
	o.AvailableForHire = v
}

// GetBio returns the Bio field value
// If the value is explicit nil, the zero value for string will be returned
func (o *VideoUser) GetBio() string {
	if o == nil || o.Bio.Get() == nil {
		var ret string
		return ret
	}

	return *o.Bio.Get()
}

// GetBioOk returns a tuple with the Bio field value
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *VideoUser) GetBioOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.Bio.Get(), o.Bio.IsSet()
}

// SetBio sets field value
func (o *VideoUser) SetBio(v string) {
	o.Bio.Set(&v)
}

// GetCanWorkRemotely returns the CanWorkRemotely field value
func (o *VideoUser) GetCanWorkRemotely() bool {
	if o == nil {
		var ret bool
		return ret
	}

	return o.CanWorkRemotely
}

// GetCanWorkRemotelyOk returns a tuple with the CanWorkRemotely field value
// and a boolean to check if the value has been set.
func (o *VideoUser) GetCanWorkRemotelyOk() (*bool, bool) {
	if o == nil {
		return nil, false
	}
	return &o.CanWorkRemotely, true
}

// SetCanWorkRemotely sets field value
func (o *VideoUser) SetCanWorkRemotely(v bool) {
	o.CanWorkRemotely = v
}

// GetCapabilities returns the Capabilities field value
func (o *VideoUser) GetCapabilities() map[string]interface{} {
	if o == nil {
		var ret map[string]interface{}
		return ret
	}

	return o.Capabilities
}

// GetCapabilitiesOk returns a tuple with the Capabilities field value
// and a boolean to check if the value has been set.
func (o *VideoUser) GetCapabilitiesOk() (map[string]interface{}, bool) {
	if o == nil {
		return map[string]interface{}{}, false
	}
	return o.Capabilities, true
}

// SetCapabilities sets field value
func (o *VideoUser) SetCapabilities(v map[string]interface{}) {
	o.Capabilities = v
}

// GetClients returns the Clients field value
func (o *VideoUser) GetClients() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Clients
}

// GetClientsOk returns a tuple with the Clients field value
// and a boolean to check if the value has been set.
func (o *VideoUser) GetClientsOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Clients, true
}

// SetClients sets field value
func (o *VideoUser) SetClients(v string) {
	o.Clients = v
}

// GetContentFilter returns the ContentFilter field value if set, zero value otherwise.
func (o *VideoUser) GetContentFilter() []string {
	if o == nil || IsNil(o.ContentFilter) {
		var ret []string
		return ret
	}
	return o.ContentFilter
}

// GetContentFilterOk returns a tuple with the ContentFilter field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *VideoUser) GetContentFilterOk() ([]string, bool) {
	if o == nil || IsNil(o.ContentFilter) {
		return nil, false
	}
	return o.ContentFilter, true
}

// HasContentFilter returns a boolean if a field has been set.
func (o *VideoUser) HasContentFilter() bool {
	if o != nil && !IsNil(o.ContentFilter) {
		return true
	}

	return false
}

// SetContentFilter gets a reference to the given []string and assigns it to the ContentFilter field.
func (o *VideoUser) SetContentFilter(v []string) {
	o.ContentFilter = v
}

// GetCreatedTime returns the CreatedTime field value
func (o *VideoUser) GetCreatedTime() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.CreatedTime
}

// GetCreatedTimeOk returns a tuple with the CreatedTime field value
// and a boolean to check if the value has been set.
func (o *VideoUser) GetCreatedTimeOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.CreatedTime, true
}

// SetCreatedTime sets field value
func (o *VideoUser) SetCreatedTime(v string) {
	o.CreatedTime = v
}

// GetGender returns the Gender field value
// If the value is explicit nil, the zero value for string will be returned
func (o *VideoUser) GetGender() string {
	if o == nil || o.Gender.Get() == nil {
		var ret string
		return ret
	}

	return *o.Gender.Get()
}

// GetGenderOk returns a tuple with the Gender field value
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *VideoUser) GetGenderOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.Gender.Get(), o.Gender.IsSet()
}

// SetGender sets field value
func (o *VideoUser) SetGender(v string) {
	o.Gender.Set(&v)
}

// GetHasInvalidEmail returns the HasInvalidEmail field value
func (o *VideoUser) GetHasInvalidEmail() bool {
	if o == nil {
		var ret bool
		return ret
	}

	return o.HasInvalidEmail
}

// GetHasInvalidEmailOk returns a tuple with the HasInvalidEmail field value
// and a boolean to check if the value has been set.
func (o *VideoUser) GetHasInvalidEmailOk() (*bool, bool) {
	if o == nil {
		return nil, false
	}
	return &o.HasInvalidEmail, true
}

// SetHasInvalidEmail sets field value
func (o *VideoUser) SetHasInvalidEmail(v bool) {
	o.HasInvalidEmail = v
}

// GetIsExpert returns the IsExpert field value
func (o *VideoUser) GetIsExpert() bool {
	if o == nil {
		var ret bool
		return ret
	}

	return o.IsExpert
}

// GetIsExpertOk returns a tuple with the IsExpert field value
// and a boolean to check if the value has been set.
func (o *VideoUser) GetIsExpertOk() (*bool, bool) {
	if o == nil {
		return nil, false
	}
	return &o.IsExpert, true
}

// SetIsExpert sets field value
func (o *VideoUser) SetIsExpert(v bool) {
	o.IsExpert = v
}

// GetLink returns the Link field value
func (o *VideoUser) GetLink() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Link
}

// GetLinkOk returns a tuple with the Link field value
// and a boolean to check if the value has been set.
func (o *VideoUser) GetLinkOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Link, true
}

// SetLink sets field value
func (o *VideoUser) SetLink(v string) {
	o.Link = v
}

// GetLocation returns the Location field value
// If the value is explicit nil, the zero value for string will be returned
func (o *VideoUser) GetLocation() string {
	if o == nil || o.Location.Get() == nil {
		var ret string
		return ret
	}

	return *o.Location.Get()
}

// GetLocationOk returns a tuple with the Location field value
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *VideoUser) GetLocationOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.Location.Get(), o.Location.IsSet()
}

// SetLocation sets field value
func (o *VideoUser) SetLocation(v string) {
	o.Location.Set(&v)
}

// GetLocationDetails returns the LocationDetails field value
// If the value is explicit nil, the zero value for UserLocationDetails will be returned
func (o *VideoUser) GetLocationDetails() UserLocationDetails {
	if o == nil || o.LocationDetails.Get() == nil {
		var ret UserLocationDetails
		return ret
	}

	return *o.LocationDetails.Get()
}

// GetLocationDetailsOk returns a tuple with the LocationDetails field value
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *VideoUser) GetLocationDetailsOk() (*UserLocationDetails, bool) {
	if o == nil {
		return nil, false
	}
	return o.LocationDetails.Get(), o.LocationDetails.IsSet()
}

// SetLocationDetails sets field value
func (o *VideoUser) SetLocationDetails(v UserLocationDetails) {
	o.LocationDetails.Set(&v)
}

// GetMetadata returns the Metadata field value
func (o *VideoUser) GetMetadata() UserMetadata {
	if o == nil {
		var ret UserMetadata
		return ret
	}

	return o.Metadata
}

// GetMetadataOk returns a tuple with the Metadata field value
// and a boolean to check if the value has been set.
func (o *VideoUser) GetMetadataOk() (*UserMetadata, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Metadata, true
}

// SetMetadata sets field value
func (o *VideoUser) SetMetadata(v UserMetadata) {
	o.Metadata = v
}

// GetName returns the Name field value
func (o *VideoUser) GetName() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Name
}

// GetNameOk returns a tuple with the Name field value
// and a boolean to check if the value has been set.
func (o *VideoUser) GetNameOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Name, true
}

// SetName sets field value
func (o *VideoUser) SetName(v string) {
	o.Name = v
}

// GetPictures returns the Pictures field value
func (o *VideoUser) GetPictures() Picture {
	if o == nil {
		var ret Picture
		return ret
	}

	return o.Pictures
}

// GetPicturesOk returns a tuple with the Pictures field value
// and a boolean to check if the value has been set.
func (o *VideoUser) GetPicturesOk() (*Picture, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Pictures, true
}

// SetPictures sets field value
func (o *VideoUser) SetPictures(v Picture) {
	o.Pictures = v
}

// GetPreferences returns the Preferences field value
func (o *VideoUser) GetPreferences() UserPreferences {
	if o == nil {
		var ret UserPreferences
		return ret
	}

	return o.Preferences
}

// GetPreferencesOk returns a tuple with the Preferences field value
// and a boolean to check if the value has been set.
func (o *VideoUser) GetPreferencesOk() (*UserPreferences, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Preferences, true
}

// SetPreferences sets field value
func (o *VideoUser) SetPreferences(v UserPreferences) {
	o.Preferences = v
}

// GetResourceKey returns the ResourceKey field value
func (o *VideoUser) GetResourceKey() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.ResourceKey
}

// GetResourceKeyOk returns a tuple with the ResourceKey field value
// and a boolean to check if the value has been set.
func (o *VideoUser) GetResourceKeyOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ResourceKey, true
}

// SetResourceKey sets field value
func (o *VideoUser) SetResourceKey(v string) {
	o.ResourceKey = v
}

// GetShortBio returns the ShortBio field value
// If the value is explicit nil, the zero value for string will be returned
func (o *VideoUser) GetShortBio() string {
	if o == nil || o.ShortBio.Get() == nil {
		var ret string
		return ret
	}

	return *o.ShortBio.Get()
}

// GetShortBioOk returns a tuple with the ShortBio field value
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *VideoUser) GetShortBioOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.ShortBio.Get(), o.ShortBio.IsSet()
}

// SetShortBio sets field value
func (o *VideoUser) SetShortBio(v string) {
	o.ShortBio.Set(&v)
}

// GetSkills returns the Skills field value
// If the value is explicit nil, the zero value for []Skill will be returned
func (o *VideoUser) GetSkills() []Skill {
	if o == nil {
		var ret []Skill
		return ret
	}

	return o.Skills
}

// GetSkillsOk returns a tuple with the Skills field value
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *VideoUser) GetSkillsOk() ([]Skill, bool) {
	if o == nil || IsNil(o.Skills) {
		return nil, false
	}
	return o.Skills, true
}

// SetSkills sets field value
func (o *VideoUser) SetSkills(v []Skill) {
	o.Skills = v
}

// GetUploadQuota returns the UploadQuota field value
func (o *VideoUser) GetUploadQuota() UserUploadQuota {
	if o == nil {
		var ret UserUploadQuota
		return ret
	}

	return o.UploadQuota
}

// GetUploadQuotaOk returns a tuple with the UploadQuota field value
// and a boolean to check if the value has been set.
func (o *VideoUser) GetUploadQuotaOk() (*UserUploadQuota, bool) {
	if o == nil {
		return nil, false
	}
	return &o.UploadQuota, true
}

// SetUploadQuota sets field value
func (o *VideoUser) SetUploadQuota(v UserUploadQuota) {
	o.UploadQuota = v
}

// GetUri returns the Uri field value
func (o *VideoUser) GetUri() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Uri
}

// GetUriOk returns a tuple with the Uri field value
// and a boolean to check if the value has been set.
func (o *VideoUser) GetUriOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Uri, true
}

// SetUri sets field value
func (o *VideoUser) SetUri(v string) {
	o.Uri = v
}

// GetWebsites returns the Websites field value
func (o *VideoUser) GetWebsites() []UserWebsitesInner {
	if o == nil {
		var ret []UserWebsitesInner
		return ret
	}

	return o.Websites
}

// GetWebsitesOk returns a tuple with the Websites field value
// and a boolean to check if the value has been set.
func (o *VideoUser) GetWebsitesOk() ([]UserWebsitesInner, bool) {
	if o == nil {
		return nil, false
	}
	return o.Websites, true
}

// SetWebsites sets field value
func (o *VideoUser) SetWebsites(v []UserWebsitesInner) {
	o.Websites = v
}

func (o VideoUser) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o VideoUser) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["account"] = o.Account
	toSerialize["available_for_hire"] = o.AvailableForHire
	toSerialize["bio"] = o.Bio.Get()
	toSerialize["can_work_remotely"] = o.CanWorkRemotely
	toSerialize["capabilities"] = o.Capabilities
	toSerialize["clients"] = o.Clients
	if !IsNil(o.ContentFilter) {
		toSerialize["content_filter"] = o.ContentFilter
	}
	toSerialize["created_time"] = o.CreatedTime
	toSerialize["gender"] = o.Gender.Get()
	toSerialize["has_invalid_email"] = o.HasInvalidEmail
	toSerialize["is_expert"] = o.IsExpert
	toSerialize["link"] = o.Link
	toSerialize["location"] = o.Location.Get()
	toSerialize["location_details"] = o.LocationDetails.Get()
	toSerialize["metadata"] = o.Metadata
	toSerialize["name"] = o.Name
	toSerialize["pictures"] = o.Pictures
	toSerialize["preferences"] = o.Preferences
	toSerialize["resource_key"] = o.ResourceKey
	toSerialize["short_bio"] = o.ShortBio.Get()
	if o.Skills != nil {
		toSerialize["skills"] = o.Skills
	}
	toSerialize["upload_quota"] = o.UploadQuota
	toSerialize["uri"] = o.Uri
	toSerialize["websites"] = o.Websites
	return toSerialize, nil
}

type NullableVideoUser struct {
	value *VideoUser
	isSet bool
}

func (v NullableVideoUser) Get() *VideoUser {
	return v.value
}

func (v *NullableVideoUser) Set(val *VideoUser) {
	v.value = val
	v.isSet = true
}

func (v NullableVideoUser) IsSet() bool {
	return v.isSet
}

func (v *NullableVideoUser) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableVideoUser(val *VideoUser) *NullableVideoUser {
	return &NullableVideoUser{value: val, isSet: true}
}

func (v NullableVideoUser) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableVideoUser) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


