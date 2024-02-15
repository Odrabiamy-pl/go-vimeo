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

// checks if the LiveEventMetadataConnections type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &LiveEventMetadataConnections{}

// LiveEventMetadataConnections A collection of information that is connected to this resource.
type LiveEventMetadataConnections struct {
	LiveVideo NullableLiveEventMetadataConnectionsLiveVideo `json:"live_video"`
	Pictures LiveEventMetadataConnectionsPictures `json:"pictures"`
	PreLiveVideo NullableLiveEventMetadataConnectionsPreLiveVideo `json:"pre_live_video"`
	TeamMember NullableLiveEventMetadataConnectionsTeamMember `json:"team_member"`
	Videos LiveEventMetadataConnectionsVideos `json:"videos"`
}

type _LiveEventMetadataConnections LiveEventMetadataConnections

// NewLiveEventMetadataConnections instantiates a new LiveEventMetadataConnections object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewLiveEventMetadataConnections(liveVideo NullableLiveEventMetadataConnectionsLiveVideo, pictures LiveEventMetadataConnectionsPictures, preLiveVideo NullableLiveEventMetadataConnectionsPreLiveVideo, teamMember NullableLiveEventMetadataConnectionsTeamMember, videos LiveEventMetadataConnectionsVideos) *LiveEventMetadataConnections {
	this := LiveEventMetadataConnections{}
	this.LiveVideo = liveVideo
	this.Pictures = pictures
	this.PreLiveVideo = preLiveVideo
	this.TeamMember = teamMember
	this.Videos = videos
	return &this
}

// NewLiveEventMetadataConnectionsWithDefaults instantiates a new LiveEventMetadataConnections object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewLiveEventMetadataConnectionsWithDefaults() *LiveEventMetadataConnections {
	this := LiveEventMetadataConnections{}
	return &this
}

// GetLiveVideo returns the LiveVideo field value
// If the value is explicit nil, the zero value for LiveEventMetadataConnectionsLiveVideo will be returned
func (o *LiveEventMetadataConnections) GetLiveVideo() LiveEventMetadataConnectionsLiveVideo {
	if o == nil || o.LiveVideo.Get() == nil {
		var ret LiveEventMetadataConnectionsLiveVideo
		return ret
	}

	return *o.LiveVideo.Get()
}

// GetLiveVideoOk returns a tuple with the LiveVideo field value
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *LiveEventMetadataConnections) GetLiveVideoOk() (*LiveEventMetadataConnectionsLiveVideo, bool) {
	if o == nil {
		return nil, false
	}
	return o.LiveVideo.Get(), o.LiveVideo.IsSet()
}

// SetLiveVideo sets field value
func (o *LiveEventMetadataConnections) SetLiveVideo(v LiveEventMetadataConnectionsLiveVideo) {
	o.LiveVideo.Set(&v)
}

// GetPictures returns the Pictures field value
func (o *LiveEventMetadataConnections) GetPictures() LiveEventMetadataConnectionsPictures {
	if o == nil {
		var ret LiveEventMetadataConnectionsPictures
		return ret
	}

	return o.Pictures
}

// GetPicturesOk returns a tuple with the Pictures field value
// and a boolean to check if the value has been set.
func (o *LiveEventMetadataConnections) GetPicturesOk() (*LiveEventMetadataConnectionsPictures, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Pictures, true
}

// SetPictures sets field value
func (o *LiveEventMetadataConnections) SetPictures(v LiveEventMetadataConnectionsPictures) {
	o.Pictures = v
}

// GetPreLiveVideo returns the PreLiveVideo field value
// If the value is explicit nil, the zero value for LiveEventMetadataConnectionsPreLiveVideo will be returned
func (o *LiveEventMetadataConnections) GetPreLiveVideo() LiveEventMetadataConnectionsPreLiveVideo {
	if o == nil || o.PreLiveVideo.Get() == nil {
		var ret LiveEventMetadataConnectionsPreLiveVideo
		return ret
	}

	return *o.PreLiveVideo.Get()
}

// GetPreLiveVideoOk returns a tuple with the PreLiveVideo field value
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *LiveEventMetadataConnections) GetPreLiveVideoOk() (*LiveEventMetadataConnectionsPreLiveVideo, bool) {
	if o == nil {
		return nil, false
	}
	return o.PreLiveVideo.Get(), o.PreLiveVideo.IsSet()
}

// SetPreLiveVideo sets field value
func (o *LiveEventMetadataConnections) SetPreLiveVideo(v LiveEventMetadataConnectionsPreLiveVideo) {
	o.PreLiveVideo.Set(&v)
}

// GetTeamMember returns the TeamMember field value
// If the value is explicit nil, the zero value for LiveEventMetadataConnectionsTeamMember will be returned
func (o *LiveEventMetadataConnections) GetTeamMember() LiveEventMetadataConnectionsTeamMember {
	if o == nil || o.TeamMember.Get() == nil {
		var ret LiveEventMetadataConnectionsTeamMember
		return ret
	}

	return *o.TeamMember.Get()
}

// GetTeamMemberOk returns a tuple with the TeamMember field value
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *LiveEventMetadataConnections) GetTeamMemberOk() (*LiveEventMetadataConnectionsTeamMember, bool) {
	if o == nil {
		return nil, false
	}
	return o.TeamMember.Get(), o.TeamMember.IsSet()
}

// SetTeamMember sets field value
func (o *LiveEventMetadataConnections) SetTeamMember(v LiveEventMetadataConnectionsTeamMember) {
	o.TeamMember.Set(&v)
}

// GetVideos returns the Videos field value
func (o *LiveEventMetadataConnections) GetVideos() LiveEventMetadataConnectionsVideos {
	if o == nil {
		var ret LiveEventMetadataConnectionsVideos
		return ret
	}

	return o.Videos
}

// GetVideosOk returns a tuple with the Videos field value
// and a boolean to check if the value has been set.
func (o *LiveEventMetadataConnections) GetVideosOk() (*LiveEventMetadataConnectionsVideos, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Videos, true
}

// SetVideos sets field value
func (o *LiveEventMetadataConnections) SetVideos(v LiveEventMetadataConnectionsVideos) {
	o.Videos = v
}

func (o LiveEventMetadataConnections) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o LiveEventMetadataConnections) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["live_video"] = o.LiveVideo.Get()
	toSerialize["pictures"] = o.Pictures
	toSerialize["pre_live_video"] = o.PreLiveVideo.Get()
	toSerialize["team_member"] = o.TeamMember.Get()
	toSerialize["videos"] = o.Videos
	return toSerialize, nil
}

func (o *LiveEventMetadataConnections) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"live_video",
		"pictures",
		"pre_live_video",
		"team_member",
		"videos",
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

	varLiveEventMetadataConnections := _LiveEventMetadataConnections{}

	decoder := json.NewDecoder(bytes.NewReader(data))
	decoder.DisallowUnknownFields()
	err = decoder.Decode(&varLiveEventMetadataConnections)

	if err != nil {
		return err
	}

	*o = LiveEventMetadataConnections(varLiveEventMetadataConnections)

	return err
}

type NullableLiveEventMetadataConnections struct {
	value *LiveEventMetadataConnections
	isSet bool
}

func (v NullableLiveEventMetadataConnections) Get() *LiveEventMetadataConnections {
	return v.value
}

func (v *NullableLiveEventMetadataConnections) Set(val *LiveEventMetadataConnections) {
	v.value = val
	v.isSet = true
}

func (v NullableLiveEventMetadataConnections) IsSet() bool {
	return v.isSet
}

func (v *NullableLiveEventMetadataConnections) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableLiveEventMetadataConnections(val *LiveEventMetadataConnections) *NullableLiveEventMetadataConnections {
	return &NullableLiveEventMetadataConnections{value: val, isSet: true}
}

func (v NullableLiveEventMetadataConnections) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableLiveEventMetadataConnections) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

