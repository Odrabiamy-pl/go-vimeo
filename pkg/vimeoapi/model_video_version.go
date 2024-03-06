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

// checks if the VideoVersion type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &VideoVersion{}

// VideoVersion struct for VideoVersion
type VideoVersion struct {
	// Whether the video version is currently active.
	Active bool   `json:"active"`
	App    ApiApp `json:"app"`
	// Whether the version can be restored.
	CanRestoreCreate bool `json:"can_restore_create"`
	// The storyboard ID of the video version.
	CreateStoryboardId string `json:"create_storyboard_id"`
	// The time in ISO 8601 format when the video version was created.
	CreatedTime string `json:"created_time"`
	// A description of the video version. This description can make use of the full unicode character set. This field appears in the response only when a corresponding value is present.
	Description *string `json:"description,omitempty"`
	// The download config associated with the version.
	DownloadConfig map[string]interface{} `json:"download_config"`
	// The duration in seconds of the video version.
	Duration NullableFloat32 `json:"duration"`
	// The file name of the video version.
	Filename string `json:"filename"`
	// The size in byes of the video version file.
	Filesize NullableFloat32 `json:"filesize"`
	// Whether the video has interactive capability.
	HasInteractive bool                 `json:"has_interactive"`
	Metadata       VideoVersionMetadata `json:"metadata"`
	// The time in ISO 8601 format when the video version was last modified.
	ModifiedTime string `json:"modified_time"`
	// Whether the video has unified resolution. If the value of this field is `false`, the video requires transcoding.
	OriginVariableFrameResolution bool                          `json:"origin_variable_frame_resolution"`
	Play                          *Play                         `json:"play,omitempty"`
	Transcode                     NullableVideoVersionTranscode `json:"transcode"`
	Upload                        NullableVideoVersionUpload    `json:"upload"`
	// The time in ISO 8601 format when the video version was uploaded.
	UploadDate NullableString `json:"upload_date"`
	// The version's canonical relative URI.
	Uri  string                   `json:"uri"`
	User NullableVideoVersionUser `json:"user"`
}

// NewVideoVersion instantiates a new VideoVersion object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewVideoVersion(active bool, app ApiApp, canRestoreCreate bool, createStoryboardId string, createdTime string, downloadConfig map[string]interface{}, duration NullableFloat32, filename string, filesize NullableFloat32, hasInteractive bool, metadata VideoVersionMetadata, modifiedTime string, originVariableFrameResolution bool, transcode NullableVideoVersionTranscode, upload NullableVideoVersionUpload, uploadDate NullableString, uri string, user NullableVideoVersionUser) *VideoVersion {
	this := VideoVersion{}
	this.Active = active
	this.App = app
	this.CanRestoreCreate = canRestoreCreate
	this.CreateStoryboardId = createStoryboardId
	this.CreatedTime = createdTime
	this.DownloadConfig = downloadConfig
	this.Duration = duration
	this.Filename = filename
	this.Filesize = filesize
	this.HasInteractive = hasInteractive
	this.Metadata = metadata
	this.ModifiedTime = modifiedTime
	this.OriginVariableFrameResolution = originVariableFrameResolution
	this.Transcode = transcode
	this.Upload = upload
	this.UploadDate = uploadDate
	this.Uri = uri
	this.User = user
	return &this
}

// NewVideoVersionWithDefaults instantiates a new VideoVersion object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewVideoVersionWithDefaults() *VideoVersion {
	this := VideoVersion{}
	return &this
}

// GetActive returns the Active field value
func (o *VideoVersion) GetActive() bool {
	if o == nil {
		var ret bool
		return ret
	}

	return o.Active
}

// GetActiveOk returns a tuple with the Active field value
// and a boolean to check if the value has been set.
func (o *VideoVersion) GetActiveOk() (*bool, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Active, true
}

// SetActive sets field value
func (o *VideoVersion) SetActive(v bool) {
	o.Active = v
}

// GetApp returns the App field value
func (o *VideoVersion) GetApp() ApiApp {
	if o == nil {
		var ret ApiApp
		return ret
	}

	return o.App
}

// GetAppOk returns a tuple with the App field value
// and a boolean to check if the value has been set.
func (o *VideoVersion) GetAppOk() (*ApiApp, bool) {
	if o == nil {
		return nil, false
	}
	return &o.App, true
}

// SetApp sets field value
func (o *VideoVersion) SetApp(v ApiApp) {
	o.App = v
}

// GetCanRestoreCreate returns the CanRestoreCreate field value
func (o *VideoVersion) GetCanRestoreCreate() bool {
	if o == nil {
		var ret bool
		return ret
	}

	return o.CanRestoreCreate
}

// GetCanRestoreCreateOk returns a tuple with the CanRestoreCreate field value
// and a boolean to check if the value has been set.
func (o *VideoVersion) GetCanRestoreCreateOk() (*bool, bool) {
	if o == nil {
		return nil, false
	}
	return &o.CanRestoreCreate, true
}

// SetCanRestoreCreate sets field value
func (o *VideoVersion) SetCanRestoreCreate(v bool) {
	o.CanRestoreCreate = v
}

// GetCreateStoryboardId returns the CreateStoryboardId field value
func (o *VideoVersion) GetCreateStoryboardId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.CreateStoryboardId
}

// GetCreateStoryboardIdOk returns a tuple with the CreateStoryboardId field value
// and a boolean to check if the value has been set.
func (o *VideoVersion) GetCreateStoryboardIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.CreateStoryboardId, true
}

// SetCreateStoryboardId sets field value
func (o *VideoVersion) SetCreateStoryboardId(v string) {
	o.CreateStoryboardId = v
}

// GetCreatedTime returns the CreatedTime field value
func (o *VideoVersion) GetCreatedTime() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.CreatedTime
}

// GetCreatedTimeOk returns a tuple with the CreatedTime field value
// and a boolean to check if the value has been set.
func (o *VideoVersion) GetCreatedTimeOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.CreatedTime, true
}

// SetCreatedTime sets field value
func (o *VideoVersion) SetCreatedTime(v string) {
	o.CreatedTime = v
}

// GetDescription returns the Description field value if set, zero value otherwise.
func (o *VideoVersion) GetDescription() string {
	if o == nil || IsNil(o.Description) {
		var ret string
		return ret
	}
	return *o.Description
}

// GetDescriptionOk returns a tuple with the Description field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *VideoVersion) GetDescriptionOk() (*string, bool) {
	if o == nil || IsNil(o.Description) {
		return nil, false
	}
	return o.Description, true
}

// HasDescription returns a boolean if a field has been set.
func (o *VideoVersion) HasDescription() bool {
	if o != nil && !IsNil(o.Description) {
		return true
	}

	return false
}

// SetDescription gets a reference to the given string and assigns it to the Description field.
func (o *VideoVersion) SetDescription(v string) {
	o.Description = &v
}

// GetDownloadConfig returns the DownloadConfig field value
func (o *VideoVersion) GetDownloadConfig() map[string]interface{} {
	if o == nil {
		var ret map[string]interface{}
		return ret
	}

	return o.DownloadConfig
}

// GetDownloadConfigOk returns a tuple with the DownloadConfig field value
// and a boolean to check if the value has been set.
func (o *VideoVersion) GetDownloadConfigOk() (map[string]interface{}, bool) {
	if o == nil {
		return map[string]interface{}{}, false
	}
	return o.DownloadConfig, true
}

// SetDownloadConfig sets field value
func (o *VideoVersion) SetDownloadConfig(v map[string]interface{}) {
	o.DownloadConfig = v
}

// GetDuration returns the Duration field value
// If the value is explicit nil, the zero value for float32 will be returned
func (o *VideoVersion) GetDuration() float32 {
	if o == nil || o.Duration.Get() == nil {
		var ret float32
		return ret
	}

	return *o.Duration.Get()
}

// GetDurationOk returns a tuple with the Duration field value
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *VideoVersion) GetDurationOk() (*float32, bool) {
	if o == nil {
		return nil, false
	}
	return o.Duration.Get(), o.Duration.IsSet()
}

// SetDuration sets field value
func (o *VideoVersion) SetDuration(v float32) {
	o.Duration.Set(&v)
}

// GetFilename returns the Filename field value
func (o *VideoVersion) GetFilename() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Filename
}

// GetFilenameOk returns a tuple with the Filename field value
// and a boolean to check if the value has been set.
func (o *VideoVersion) GetFilenameOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Filename, true
}

// SetFilename sets field value
func (o *VideoVersion) SetFilename(v string) {
	o.Filename = v
}

// GetFilesize returns the Filesize field value
// If the value is explicit nil, the zero value for float32 will be returned
func (o *VideoVersion) GetFilesize() float32 {
	if o == nil || o.Filesize.Get() == nil {
		var ret float32
		return ret
	}

	return *o.Filesize.Get()
}

// GetFilesizeOk returns a tuple with the Filesize field value
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *VideoVersion) GetFilesizeOk() (*float32, bool) {
	if o == nil {
		return nil, false
	}
	return o.Filesize.Get(), o.Filesize.IsSet()
}

// SetFilesize sets field value
func (o *VideoVersion) SetFilesize(v float32) {
	o.Filesize.Set(&v)
}

// GetHasInteractive returns the HasInteractive field value
func (o *VideoVersion) GetHasInteractive() bool {
	if o == nil {
		var ret bool
		return ret
	}

	return o.HasInteractive
}

// GetHasInteractiveOk returns a tuple with the HasInteractive field value
// and a boolean to check if the value has been set.
func (o *VideoVersion) GetHasInteractiveOk() (*bool, bool) {
	if o == nil {
		return nil, false
	}
	return &o.HasInteractive, true
}

// SetHasInteractive sets field value
func (o *VideoVersion) SetHasInteractive(v bool) {
	o.HasInteractive = v
}

// GetMetadata returns the Metadata field value
func (o *VideoVersion) GetMetadata() VideoVersionMetadata {
	if o == nil {
		var ret VideoVersionMetadata
		return ret
	}

	return o.Metadata
}

// GetMetadataOk returns a tuple with the Metadata field value
// and a boolean to check if the value has been set.
func (o *VideoVersion) GetMetadataOk() (*VideoVersionMetadata, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Metadata, true
}

// SetMetadata sets field value
func (o *VideoVersion) SetMetadata(v VideoVersionMetadata) {
	o.Metadata = v
}

// GetModifiedTime returns the ModifiedTime field value
func (o *VideoVersion) GetModifiedTime() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.ModifiedTime
}

// GetModifiedTimeOk returns a tuple with the ModifiedTime field value
// and a boolean to check if the value has been set.
func (o *VideoVersion) GetModifiedTimeOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ModifiedTime, true
}

// SetModifiedTime sets field value
func (o *VideoVersion) SetModifiedTime(v string) {
	o.ModifiedTime = v
}

// GetOriginVariableFrameResolution returns the OriginVariableFrameResolution field value
func (o *VideoVersion) GetOriginVariableFrameResolution() bool {
	if o == nil {
		var ret bool
		return ret
	}

	return o.OriginVariableFrameResolution
}

// GetOriginVariableFrameResolutionOk returns a tuple with the OriginVariableFrameResolution field value
// and a boolean to check if the value has been set.
func (o *VideoVersion) GetOriginVariableFrameResolutionOk() (*bool, bool) {
	if o == nil {
		return nil, false
	}
	return &o.OriginVariableFrameResolution, true
}

// SetOriginVariableFrameResolution sets field value
func (o *VideoVersion) SetOriginVariableFrameResolution(v bool) {
	o.OriginVariableFrameResolution = v
}

// GetPlay returns the Play field value if set, zero value otherwise.
func (o *VideoVersion) GetPlay() Play {
	if o == nil || IsNil(o.Play) {
		var ret Play
		return ret
	}
	return *o.Play
}

// GetPlayOk returns a tuple with the Play field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *VideoVersion) GetPlayOk() (*Play, bool) {
	if o == nil || IsNil(o.Play) {
		return nil, false
	}
	return o.Play, true
}

// HasPlay returns a boolean if a field has been set.
func (o *VideoVersion) HasPlay() bool {
	if o != nil && !IsNil(o.Play) {
		return true
	}

	return false
}

// SetPlay gets a reference to the given Play and assigns it to the Play field.
func (o *VideoVersion) SetPlay(v Play) {
	o.Play = &v
}

// GetTranscode returns the Transcode field value
// If the value is explicit nil, the zero value for VideoVersionTranscode will be returned
func (o *VideoVersion) GetTranscode() VideoVersionTranscode {
	if o == nil || o.Transcode.Get() == nil {
		var ret VideoVersionTranscode
		return ret
	}

	return *o.Transcode.Get()
}

// GetTranscodeOk returns a tuple with the Transcode field value
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *VideoVersion) GetTranscodeOk() (*VideoVersionTranscode, bool) {
	if o == nil {
		return nil, false
	}
	return o.Transcode.Get(), o.Transcode.IsSet()
}

// SetTranscode sets field value
func (o *VideoVersion) SetTranscode(v VideoVersionTranscode) {
	o.Transcode.Set(&v)
}

// GetUpload returns the Upload field value
// If the value is explicit nil, the zero value for VideoVersionUpload will be returned
func (o *VideoVersion) GetUpload() VideoVersionUpload {
	if o == nil || o.Upload.Get() == nil {
		var ret VideoVersionUpload
		return ret
	}

	return *o.Upload.Get()
}

// GetUploadOk returns a tuple with the Upload field value
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *VideoVersion) GetUploadOk() (*VideoVersionUpload, bool) {
	if o == nil {
		return nil, false
	}
	return o.Upload.Get(), o.Upload.IsSet()
}

// SetUpload sets field value
func (o *VideoVersion) SetUpload(v VideoVersionUpload) {
	o.Upload.Set(&v)
}

// GetUploadDate returns the UploadDate field value
// If the value is explicit nil, the zero value for string will be returned
func (o *VideoVersion) GetUploadDate() string {
	if o == nil || o.UploadDate.Get() == nil {
		var ret string
		return ret
	}

	return *o.UploadDate.Get()
}

// GetUploadDateOk returns a tuple with the UploadDate field value
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *VideoVersion) GetUploadDateOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.UploadDate.Get(), o.UploadDate.IsSet()
}

// SetUploadDate sets field value
func (o *VideoVersion) SetUploadDate(v string) {
	o.UploadDate.Set(&v)
}

// GetUri returns the Uri field value
func (o *VideoVersion) GetUri() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Uri
}

// GetUriOk returns a tuple with the Uri field value
// and a boolean to check if the value has been set.
func (o *VideoVersion) GetUriOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Uri, true
}

// SetUri sets field value
func (o *VideoVersion) SetUri(v string) {
	o.Uri = v
}

// GetUser returns the User field value
// If the value is explicit nil, the zero value for VideoVersionUser will be returned
func (o *VideoVersion) GetUser() VideoVersionUser {
	if o == nil || o.User.Get() == nil {
		var ret VideoVersionUser
		return ret
	}

	return *o.User.Get()
}

// GetUserOk returns a tuple with the User field value
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *VideoVersion) GetUserOk() (*VideoVersionUser, bool) {
	if o == nil {
		return nil, false
	}
	return o.User.Get(), o.User.IsSet()
}

// SetUser sets field value
func (o *VideoVersion) SetUser(v VideoVersionUser) {
	o.User.Set(&v)
}

func (o VideoVersion) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o VideoVersion) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["active"] = o.Active
	toSerialize["app"] = o.App
	toSerialize["can_restore_create"] = o.CanRestoreCreate
	toSerialize["create_storyboard_id"] = o.CreateStoryboardId
	toSerialize["created_time"] = o.CreatedTime
	if !IsNil(o.Description) {
		toSerialize["description"] = o.Description
	}
	toSerialize["download_config"] = o.DownloadConfig
	toSerialize["duration"] = o.Duration.Get()
	toSerialize["filename"] = o.Filename
	toSerialize["filesize"] = o.Filesize.Get()
	toSerialize["has_interactive"] = o.HasInteractive
	toSerialize["metadata"] = o.Metadata
	toSerialize["modified_time"] = o.ModifiedTime
	toSerialize["origin_variable_frame_resolution"] = o.OriginVariableFrameResolution
	if !IsNil(o.Play) {
		toSerialize["play"] = o.Play
	}
	toSerialize["transcode"] = o.Transcode.Get()
	toSerialize["upload"] = o.Upload.Get()
	toSerialize["upload_date"] = o.UploadDate.Get()
	toSerialize["uri"] = o.Uri
	toSerialize["user"] = o.User.Get()
	return toSerialize, nil
}

type NullableVideoVersion struct {
	value *VideoVersion
	isSet bool
}

func (v NullableVideoVersion) Get() *VideoVersion {
	return v.value
}

func (v *NullableVideoVersion) Set(val *VideoVersion) {
	v.value = val
	v.isSet = true
}

func (v NullableVideoVersion) IsSet() bool {
	return v.isSet
}

func (v *NullableVideoVersion) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableVideoVersion(val *VideoVersion) *NullableVideoVersion {
	return &NullableVideoVersion{value: val, isSet: true}
}

func (v NullableVideoVersion) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableVideoVersion) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
