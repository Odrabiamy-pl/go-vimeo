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

// checks if the LiveEventEmbed type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &LiveEventEmbed{}

// LiveEventEmbed The event's embed data.
type LiveEventEmbed struct {
	// Whether AirPlay is enabled in the embeddable player.
	Airplay bool `json:"airplay"`
	// Whether the embedded RLE player should autoplay the RLE content.
	Autoplay bool `json:"autoplay"`
	// The list of user-uploaded logos for configuration of the embed player.
	AvailablePlayerLogos []string `json:"available_player_logos"`
	// Whether the embedded RLE player should display the author's name.
	Byline bool `json:"byline"`
	// The chat's iFrame source URL.
	ChatEmbedSource NullableString `json:"chat_embed_source"`
	// Whether the Chromecast button appears in the embeddable player.
	Chromecast bool `json:"chromecast"`
	// Whether closed captions are enabled in the embeddable player.
	ClosedCaptions bool `json:"closed_captions"`
	// The first player color, which controls the color of the progress bar, buttons, and more.
	Color string `json:"color"`
	Colors EmbedSettingsColors `json:"colors"`
	// The embed code for RLE chat.
	EmbedChat NullableString `json:"embed_chat"`
	EmbedProperties NullableLiveEventEmbedEmbedProperties `json:"embed_properties"`
	// Whether the embedded RLE player should display the event schedule.
	EventSchedule bool `json:"event_schedule"`
	// Whether the embedded RLE player should include the fullscreen controls.
	FullscreenButton bool `json:"fullscreen_button"`
	// Whether the Live label should be visible over the player.
	HideLiveLabel bool `json:"hide_live_label"`
	// Whether the embedded RLE player should hide the viewer counter.
	HideViewerCount bool `json:"hide_viewer_count"`
	// The fixed HTML code to embed the event's playlist on a website.
	Html NullableString `json:"html"`
	// Whether the embedded RLE player should include the `like` button.
	LikeButton bool `json:"like_button"`
	Logos LiveEventEmbedLogos `json:"logos"`
	// Whether the embedded RLE player should loop back to the first video once content is exhausted.
	Loop bool `json:"loop"`
	// Whether picture-in-picture is enabled and the button appears in the embeddable player.
	Pip bool `json:"pip"`
	// The position for the player's play button.
	PlayButtonPosition float32 `json:"play_button_position"`
	// Whether the embedded RLE player should include the playbar.
	Playbar bool `json:"playbar"`
	// Whether the playlist component appears in the embeddable player for this RLE.
	Playlist bool `json:"playlist"`
	// Whether the embedded RLE player should display the author's portrait.
	Portrait bool `json:"portrait"`
	// The responsive HTML code to embed the event's playlist on a website.
	ResponsiveHtml NullableString `json:"responsive_html"`
	// Whether the schedule component appears in the embeddable player for this RLE.
	Schedule bool `json:"schedule"`
	// Whether the embedded RLE player should display the latest video placeholder.
	ShowLatestArchivedClip bool `json:"show_latest_archived_clip"`
	// Whether the embedded RLE player should display the schedule timezone.
	ShowTimezone bool `json:"show_timezone"`
	// Whether the embedded RLE player should display the video title.
	Title bool `json:"title"`
	// Whether the embedded RLE player should use a custom color or the default Vimeo blue.
	UseColor string `json:"use_color"`
	// Whether the embedded RLE player should include the volume controls.
	Volume bool `json:"volume"`
}

// NewLiveEventEmbed instantiates a new LiveEventEmbed object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewLiveEventEmbed(airplay bool, autoplay bool, availablePlayerLogos []string, byline bool, chatEmbedSource NullableString, chromecast bool, closedCaptions bool, color string, colors EmbedSettingsColors, embedChat NullableString, embedProperties NullableLiveEventEmbedEmbedProperties, eventSchedule bool, fullscreenButton bool, hideLiveLabel bool, hideViewerCount bool, html NullableString, likeButton bool, logos LiveEventEmbedLogos, loop bool, pip bool, playButtonPosition float32, playbar bool, playlist bool, portrait bool, responsiveHtml NullableString, schedule bool, showLatestArchivedClip bool, showTimezone bool, title bool, useColor string, volume bool) *LiveEventEmbed {
	this := LiveEventEmbed{}
	this.Airplay = airplay
	this.Autoplay = autoplay
	this.AvailablePlayerLogos = availablePlayerLogos
	this.Byline = byline
	this.ChatEmbedSource = chatEmbedSource
	this.Chromecast = chromecast
	this.ClosedCaptions = closedCaptions
	this.Color = color
	this.Colors = colors
	this.EmbedChat = embedChat
	this.EmbedProperties = embedProperties
	this.EventSchedule = eventSchedule
	this.FullscreenButton = fullscreenButton
	this.HideLiveLabel = hideLiveLabel
	this.HideViewerCount = hideViewerCount
	this.Html = html
	this.LikeButton = likeButton
	this.Logos = logos
	this.Loop = loop
	this.Pip = pip
	this.PlayButtonPosition = playButtonPosition
	this.Playbar = playbar
	this.Playlist = playlist
	this.Portrait = portrait
	this.ResponsiveHtml = responsiveHtml
	this.Schedule = schedule
	this.ShowLatestArchivedClip = showLatestArchivedClip
	this.ShowTimezone = showTimezone
	this.Title = title
	this.UseColor = useColor
	this.Volume = volume
	return &this
}

// NewLiveEventEmbedWithDefaults instantiates a new LiveEventEmbed object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewLiveEventEmbedWithDefaults() *LiveEventEmbed {
	this := LiveEventEmbed{}
	return &this
}

// GetAirplay returns the Airplay field value
func (o *LiveEventEmbed) GetAirplay() bool {
	if o == nil {
		var ret bool
		return ret
	}

	return o.Airplay
}

// GetAirplayOk returns a tuple with the Airplay field value
// and a boolean to check if the value has been set.
func (o *LiveEventEmbed) GetAirplayOk() (*bool, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Airplay, true
}

// SetAirplay sets field value
func (o *LiveEventEmbed) SetAirplay(v bool) {
	o.Airplay = v
}

// GetAutoplay returns the Autoplay field value
func (o *LiveEventEmbed) GetAutoplay() bool {
	if o == nil {
		var ret bool
		return ret
	}

	return o.Autoplay
}

// GetAutoplayOk returns a tuple with the Autoplay field value
// and a boolean to check if the value has been set.
func (o *LiveEventEmbed) GetAutoplayOk() (*bool, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Autoplay, true
}

// SetAutoplay sets field value
func (o *LiveEventEmbed) SetAutoplay(v bool) {
	o.Autoplay = v
}

// GetAvailablePlayerLogos returns the AvailablePlayerLogos field value
func (o *LiveEventEmbed) GetAvailablePlayerLogos() []string {
	if o == nil {
		var ret []string
		return ret
	}

	return o.AvailablePlayerLogos
}

// GetAvailablePlayerLogosOk returns a tuple with the AvailablePlayerLogos field value
// and a boolean to check if the value has been set.
func (o *LiveEventEmbed) GetAvailablePlayerLogosOk() ([]string, bool) {
	if o == nil {
		return nil, false
	}
	return o.AvailablePlayerLogos, true
}

// SetAvailablePlayerLogos sets field value
func (o *LiveEventEmbed) SetAvailablePlayerLogos(v []string) {
	o.AvailablePlayerLogos = v
}

// GetByline returns the Byline field value
func (o *LiveEventEmbed) GetByline() bool {
	if o == nil {
		var ret bool
		return ret
	}

	return o.Byline
}

// GetBylineOk returns a tuple with the Byline field value
// and a boolean to check if the value has been set.
func (o *LiveEventEmbed) GetBylineOk() (*bool, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Byline, true
}

// SetByline sets field value
func (o *LiveEventEmbed) SetByline(v bool) {
	o.Byline = v
}

// GetChatEmbedSource returns the ChatEmbedSource field value
// If the value is explicit nil, the zero value for string will be returned
func (o *LiveEventEmbed) GetChatEmbedSource() string {
	if o == nil || o.ChatEmbedSource.Get() == nil {
		var ret string
		return ret
	}

	return *o.ChatEmbedSource.Get()
}

// GetChatEmbedSourceOk returns a tuple with the ChatEmbedSource field value
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *LiveEventEmbed) GetChatEmbedSourceOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.ChatEmbedSource.Get(), o.ChatEmbedSource.IsSet()
}

// SetChatEmbedSource sets field value
func (o *LiveEventEmbed) SetChatEmbedSource(v string) {
	o.ChatEmbedSource.Set(&v)
}

// GetChromecast returns the Chromecast field value
func (o *LiveEventEmbed) GetChromecast() bool {
	if o == nil {
		var ret bool
		return ret
	}

	return o.Chromecast
}

// GetChromecastOk returns a tuple with the Chromecast field value
// and a boolean to check if the value has been set.
func (o *LiveEventEmbed) GetChromecastOk() (*bool, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Chromecast, true
}

// SetChromecast sets field value
func (o *LiveEventEmbed) SetChromecast(v bool) {
	o.Chromecast = v
}

// GetClosedCaptions returns the ClosedCaptions field value
func (o *LiveEventEmbed) GetClosedCaptions() bool {
	if o == nil {
		var ret bool
		return ret
	}

	return o.ClosedCaptions
}

// GetClosedCaptionsOk returns a tuple with the ClosedCaptions field value
// and a boolean to check if the value has been set.
func (o *LiveEventEmbed) GetClosedCaptionsOk() (*bool, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ClosedCaptions, true
}

// SetClosedCaptions sets field value
func (o *LiveEventEmbed) SetClosedCaptions(v bool) {
	o.ClosedCaptions = v
}

// GetColor returns the Color field value
func (o *LiveEventEmbed) GetColor() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Color
}

// GetColorOk returns a tuple with the Color field value
// and a boolean to check if the value has been set.
func (o *LiveEventEmbed) GetColorOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Color, true
}

// SetColor sets field value
func (o *LiveEventEmbed) SetColor(v string) {
	o.Color = v
}

// GetColors returns the Colors field value
func (o *LiveEventEmbed) GetColors() EmbedSettingsColors {
	if o == nil {
		var ret EmbedSettingsColors
		return ret
	}

	return o.Colors
}

// GetColorsOk returns a tuple with the Colors field value
// and a boolean to check if the value has been set.
func (o *LiveEventEmbed) GetColorsOk() (*EmbedSettingsColors, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Colors, true
}

// SetColors sets field value
func (o *LiveEventEmbed) SetColors(v EmbedSettingsColors) {
	o.Colors = v
}

// GetEmbedChat returns the EmbedChat field value
// If the value is explicit nil, the zero value for string will be returned
func (o *LiveEventEmbed) GetEmbedChat() string {
	if o == nil || o.EmbedChat.Get() == nil {
		var ret string
		return ret
	}

	return *o.EmbedChat.Get()
}

// GetEmbedChatOk returns a tuple with the EmbedChat field value
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *LiveEventEmbed) GetEmbedChatOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.EmbedChat.Get(), o.EmbedChat.IsSet()
}

// SetEmbedChat sets field value
func (o *LiveEventEmbed) SetEmbedChat(v string) {
	o.EmbedChat.Set(&v)
}

// GetEmbedProperties returns the EmbedProperties field value
// If the value is explicit nil, the zero value for LiveEventEmbedEmbedProperties will be returned
func (o *LiveEventEmbed) GetEmbedProperties() LiveEventEmbedEmbedProperties {
	if o == nil || o.EmbedProperties.Get() == nil {
		var ret LiveEventEmbedEmbedProperties
		return ret
	}

	return *o.EmbedProperties.Get()
}

// GetEmbedPropertiesOk returns a tuple with the EmbedProperties field value
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *LiveEventEmbed) GetEmbedPropertiesOk() (*LiveEventEmbedEmbedProperties, bool) {
	if o == nil {
		return nil, false
	}
	return o.EmbedProperties.Get(), o.EmbedProperties.IsSet()
}

// SetEmbedProperties sets field value
func (o *LiveEventEmbed) SetEmbedProperties(v LiveEventEmbedEmbedProperties) {
	o.EmbedProperties.Set(&v)
}

// GetEventSchedule returns the EventSchedule field value
func (o *LiveEventEmbed) GetEventSchedule() bool {
	if o == nil {
		var ret bool
		return ret
	}

	return o.EventSchedule
}

// GetEventScheduleOk returns a tuple with the EventSchedule field value
// and a boolean to check if the value has been set.
func (o *LiveEventEmbed) GetEventScheduleOk() (*bool, bool) {
	if o == nil {
		return nil, false
	}
	return &o.EventSchedule, true
}

// SetEventSchedule sets field value
func (o *LiveEventEmbed) SetEventSchedule(v bool) {
	o.EventSchedule = v
}

// GetFullscreenButton returns the FullscreenButton field value
func (o *LiveEventEmbed) GetFullscreenButton() bool {
	if o == nil {
		var ret bool
		return ret
	}

	return o.FullscreenButton
}

// GetFullscreenButtonOk returns a tuple with the FullscreenButton field value
// and a boolean to check if the value has been set.
func (o *LiveEventEmbed) GetFullscreenButtonOk() (*bool, bool) {
	if o == nil {
		return nil, false
	}
	return &o.FullscreenButton, true
}

// SetFullscreenButton sets field value
func (o *LiveEventEmbed) SetFullscreenButton(v bool) {
	o.FullscreenButton = v
}

// GetHideLiveLabel returns the HideLiveLabel field value
func (o *LiveEventEmbed) GetHideLiveLabel() bool {
	if o == nil {
		var ret bool
		return ret
	}

	return o.HideLiveLabel
}

// GetHideLiveLabelOk returns a tuple with the HideLiveLabel field value
// and a boolean to check if the value has been set.
func (o *LiveEventEmbed) GetHideLiveLabelOk() (*bool, bool) {
	if o == nil {
		return nil, false
	}
	return &o.HideLiveLabel, true
}

// SetHideLiveLabel sets field value
func (o *LiveEventEmbed) SetHideLiveLabel(v bool) {
	o.HideLiveLabel = v
}

// GetHideViewerCount returns the HideViewerCount field value
func (o *LiveEventEmbed) GetHideViewerCount() bool {
	if o == nil {
		var ret bool
		return ret
	}

	return o.HideViewerCount
}

// GetHideViewerCountOk returns a tuple with the HideViewerCount field value
// and a boolean to check if the value has been set.
func (o *LiveEventEmbed) GetHideViewerCountOk() (*bool, bool) {
	if o == nil {
		return nil, false
	}
	return &o.HideViewerCount, true
}

// SetHideViewerCount sets field value
func (o *LiveEventEmbed) SetHideViewerCount(v bool) {
	o.HideViewerCount = v
}

// GetHtml returns the Html field value
// If the value is explicit nil, the zero value for string will be returned
func (o *LiveEventEmbed) GetHtml() string {
	if o == nil || o.Html.Get() == nil {
		var ret string
		return ret
	}

	return *o.Html.Get()
}

// GetHtmlOk returns a tuple with the Html field value
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *LiveEventEmbed) GetHtmlOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.Html.Get(), o.Html.IsSet()
}

// SetHtml sets field value
func (o *LiveEventEmbed) SetHtml(v string) {
	o.Html.Set(&v)
}

// GetLikeButton returns the LikeButton field value
func (o *LiveEventEmbed) GetLikeButton() bool {
	if o == nil {
		var ret bool
		return ret
	}

	return o.LikeButton
}

// GetLikeButtonOk returns a tuple with the LikeButton field value
// and a boolean to check if the value has been set.
func (o *LiveEventEmbed) GetLikeButtonOk() (*bool, bool) {
	if o == nil {
		return nil, false
	}
	return &o.LikeButton, true
}

// SetLikeButton sets field value
func (o *LiveEventEmbed) SetLikeButton(v bool) {
	o.LikeButton = v
}

// GetLogos returns the Logos field value
func (o *LiveEventEmbed) GetLogos() LiveEventEmbedLogos {
	if o == nil {
		var ret LiveEventEmbedLogos
		return ret
	}

	return o.Logos
}

// GetLogosOk returns a tuple with the Logos field value
// and a boolean to check if the value has been set.
func (o *LiveEventEmbed) GetLogosOk() (*LiveEventEmbedLogos, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Logos, true
}

// SetLogos sets field value
func (o *LiveEventEmbed) SetLogos(v LiveEventEmbedLogos) {
	o.Logos = v
}

// GetLoop returns the Loop field value
func (o *LiveEventEmbed) GetLoop() bool {
	if o == nil {
		var ret bool
		return ret
	}

	return o.Loop
}

// GetLoopOk returns a tuple with the Loop field value
// and a boolean to check if the value has been set.
func (o *LiveEventEmbed) GetLoopOk() (*bool, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Loop, true
}

// SetLoop sets field value
func (o *LiveEventEmbed) SetLoop(v bool) {
	o.Loop = v
}

// GetPip returns the Pip field value
func (o *LiveEventEmbed) GetPip() bool {
	if o == nil {
		var ret bool
		return ret
	}

	return o.Pip
}

// GetPipOk returns a tuple with the Pip field value
// and a boolean to check if the value has been set.
func (o *LiveEventEmbed) GetPipOk() (*bool, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Pip, true
}

// SetPip sets field value
func (o *LiveEventEmbed) SetPip(v bool) {
	o.Pip = v
}

// GetPlayButtonPosition returns the PlayButtonPosition field value
func (o *LiveEventEmbed) GetPlayButtonPosition() float32 {
	if o == nil {
		var ret float32
		return ret
	}

	return o.PlayButtonPosition
}

// GetPlayButtonPositionOk returns a tuple with the PlayButtonPosition field value
// and a boolean to check if the value has been set.
func (o *LiveEventEmbed) GetPlayButtonPositionOk() (*float32, bool) {
	if o == nil {
		return nil, false
	}
	return &o.PlayButtonPosition, true
}

// SetPlayButtonPosition sets field value
func (o *LiveEventEmbed) SetPlayButtonPosition(v float32) {
	o.PlayButtonPosition = v
}

// GetPlaybar returns the Playbar field value
func (o *LiveEventEmbed) GetPlaybar() bool {
	if o == nil {
		var ret bool
		return ret
	}

	return o.Playbar
}

// GetPlaybarOk returns a tuple with the Playbar field value
// and a boolean to check if the value has been set.
func (o *LiveEventEmbed) GetPlaybarOk() (*bool, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Playbar, true
}

// SetPlaybar sets field value
func (o *LiveEventEmbed) SetPlaybar(v bool) {
	o.Playbar = v
}

// GetPlaylist returns the Playlist field value
func (o *LiveEventEmbed) GetPlaylist() bool {
	if o == nil {
		var ret bool
		return ret
	}

	return o.Playlist
}

// GetPlaylistOk returns a tuple with the Playlist field value
// and a boolean to check if the value has been set.
func (o *LiveEventEmbed) GetPlaylistOk() (*bool, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Playlist, true
}

// SetPlaylist sets field value
func (o *LiveEventEmbed) SetPlaylist(v bool) {
	o.Playlist = v
}

// GetPortrait returns the Portrait field value
func (o *LiveEventEmbed) GetPortrait() bool {
	if o == nil {
		var ret bool
		return ret
	}

	return o.Portrait
}

// GetPortraitOk returns a tuple with the Portrait field value
// and a boolean to check if the value has been set.
func (o *LiveEventEmbed) GetPortraitOk() (*bool, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Portrait, true
}

// SetPortrait sets field value
func (o *LiveEventEmbed) SetPortrait(v bool) {
	o.Portrait = v
}

// GetResponsiveHtml returns the ResponsiveHtml field value
// If the value is explicit nil, the zero value for string will be returned
func (o *LiveEventEmbed) GetResponsiveHtml() string {
	if o == nil || o.ResponsiveHtml.Get() == nil {
		var ret string
		return ret
	}

	return *o.ResponsiveHtml.Get()
}

// GetResponsiveHtmlOk returns a tuple with the ResponsiveHtml field value
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *LiveEventEmbed) GetResponsiveHtmlOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.ResponsiveHtml.Get(), o.ResponsiveHtml.IsSet()
}

// SetResponsiveHtml sets field value
func (o *LiveEventEmbed) SetResponsiveHtml(v string) {
	o.ResponsiveHtml.Set(&v)
}

// GetSchedule returns the Schedule field value
func (o *LiveEventEmbed) GetSchedule() bool {
	if o == nil {
		var ret bool
		return ret
	}

	return o.Schedule
}

// GetScheduleOk returns a tuple with the Schedule field value
// and a boolean to check if the value has been set.
func (o *LiveEventEmbed) GetScheduleOk() (*bool, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Schedule, true
}

// SetSchedule sets field value
func (o *LiveEventEmbed) SetSchedule(v bool) {
	o.Schedule = v
}

// GetShowLatestArchivedClip returns the ShowLatestArchivedClip field value
func (o *LiveEventEmbed) GetShowLatestArchivedClip() bool {
	if o == nil {
		var ret bool
		return ret
	}

	return o.ShowLatestArchivedClip
}

// GetShowLatestArchivedClipOk returns a tuple with the ShowLatestArchivedClip field value
// and a boolean to check if the value has been set.
func (o *LiveEventEmbed) GetShowLatestArchivedClipOk() (*bool, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ShowLatestArchivedClip, true
}

// SetShowLatestArchivedClip sets field value
func (o *LiveEventEmbed) SetShowLatestArchivedClip(v bool) {
	o.ShowLatestArchivedClip = v
}

// GetShowTimezone returns the ShowTimezone field value
func (o *LiveEventEmbed) GetShowTimezone() bool {
	if o == nil {
		var ret bool
		return ret
	}

	return o.ShowTimezone
}

// GetShowTimezoneOk returns a tuple with the ShowTimezone field value
// and a boolean to check if the value has been set.
func (o *LiveEventEmbed) GetShowTimezoneOk() (*bool, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ShowTimezone, true
}

// SetShowTimezone sets field value
func (o *LiveEventEmbed) SetShowTimezone(v bool) {
	o.ShowTimezone = v
}

// GetTitle returns the Title field value
func (o *LiveEventEmbed) GetTitle() bool {
	if o == nil {
		var ret bool
		return ret
	}

	return o.Title
}

// GetTitleOk returns a tuple with the Title field value
// and a boolean to check if the value has been set.
func (o *LiveEventEmbed) GetTitleOk() (*bool, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Title, true
}

// SetTitle sets field value
func (o *LiveEventEmbed) SetTitle(v bool) {
	o.Title = v
}

// GetUseColor returns the UseColor field value
func (o *LiveEventEmbed) GetUseColor() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.UseColor
}

// GetUseColorOk returns a tuple with the UseColor field value
// and a boolean to check if the value has been set.
func (o *LiveEventEmbed) GetUseColorOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.UseColor, true
}

// SetUseColor sets field value
func (o *LiveEventEmbed) SetUseColor(v string) {
	o.UseColor = v
}

// GetVolume returns the Volume field value
func (o *LiveEventEmbed) GetVolume() bool {
	if o == nil {
		var ret bool
		return ret
	}

	return o.Volume
}

// GetVolumeOk returns a tuple with the Volume field value
// and a boolean to check if the value has been set.
func (o *LiveEventEmbed) GetVolumeOk() (*bool, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Volume, true
}

// SetVolume sets field value
func (o *LiveEventEmbed) SetVolume(v bool) {
	o.Volume = v
}

func (o LiveEventEmbed) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o LiveEventEmbed) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["airplay"] = o.Airplay
	toSerialize["autoplay"] = o.Autoplay
	toSerialize["available_player_logos"] = o.AvailablePlayerLogos
	toSerialize["byline"] = o.Byline
	toSerialize["chat_embed_source"] = o.ChatEmbedSource.Get()
	toSerialize["chromecast"] = o.Chromecast
	toSerialize["closed_captions"] = o.ClosedCaptions
	toSerialize["color"] = o.Color
	toSerialize["colors"] = o.Colors
	toSerialize["embed_chat"] = o.EmbedChat.Get()
	toSerialize["embed_properties"] = o.EmbedProperties.Get()
	toSerialize["event_schedule"] = o.EventSchedule
	toSerialize["fullscreen_button"] = o.FullscreenButton
	toSerialize["hide_live_label"] = o.HideLiveLabel
	toSerialize["hide_viewer_count"] = o.HideViewerCount
	toSerialize["html"] = o.Html.Get()
	toSerialize["like_button"] = o.LikeButton
	toSerialize["logos"] = o.Logos
	toSerialize["loop"] = o.Loop
	toSerialize["pip"] = o.Pip
	toSerialize["play_button_position"] = o.PlayButtonPosition
	toSerialize["playbar"] = o.Playbar
	toSerialize["playlist"] = o.Playlist
	toSerialize["portrait"] = o.Portrait
	toSerialize["responsive_html"] = o.ResponsiveHtml.Get()
	toSerialize["schedule"] = o.Schedule
	toSerialize["show_latest_archived_clip"] = o.ShowLatestArchivedClip
	toSerialize["show_timezone"] = o.ShowTimezone
	toSerialize["title"] = o.Title
	toSerialize["use_color"] = o.UseColor
	toSerialize["volume"] = o.Volume
	return toSerialize, nil
}

type NullableLiveEventEmbed struct {
	value *LiveEventEmbed
	isSet bool
}

func (v NullableLiveEventEmbed) Get() *LiveEventEmbed {
	return v.value
}

func (v *NullableLiveEventEmbed) Set(val *LiveEventEmbed) {
	v.value = val
	v.isSet = true
}

func (v NullableLiveEventEmbed) IsSet() bool {
	return v.isSet
}

func (v *NullableLiveEventEmbed) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableLiveEventEmbed(val *LiveEventEmbed) *NullableLiveEventEmbed {
	return &NullableLiveEventEmbed{value: val, isSet: true}
}

func (v NullableLiveEventEmbed) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableLiveEventEmbed) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


