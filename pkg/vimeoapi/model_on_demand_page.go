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

// checks if the OnDemandPage type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &OnDemandPage{}

// OnDemandPage struct for OnDemandPage
type OnDemandPage struct {
	Background NullableOnDemandPageBackground `json:"background"`
	Colors OnDemandPageColors `json:"colors"`
	// An array of the On Demand page's content ratings.
	ContentRating []string `json:"content_rating"`
	// The time in ISO 8601 format when the On Demand page was created.
	CreatedTime *string `json:"created_time,omitempty"`
	// The description of the On Demand page.
	Description NullableString `json:"description"`
	// The link to the On Demand page on its own domain.
	DomainLink NullableString `json:"domain_link"`
	Episodes OnDemandPageEpisodes `json:"episodes"`
	// The On Demand page's film, if the page is for a film.
	Film *Video `json:"film,omitempty"`
	// An array of the genres assigned to the On Demand page.
	Genres []OnDemandGenre `json:"genres"`
	// The link to the On Demand page.
	Link string `json:"link"`
	Metadata OnDemandPageMetadata `json:"metadata"`
	// The time in ISO 8601 format when the On Demand page was last modified.
	ModifiedTime *string `json:"modified_time,omitempty"`
	// The descriptive title of the On Demand page.
	Name string `json:"name"`
	Pictures NullableOnDemandPagePictures `json:"pictures"`
	Preorder OnDemandPagePreorder `json:"preorder"`
	Published OnDemandPagePublished `json:"published"`
	// The rating of the On Demand page.
	Rating NullableFloat32 `json:"rating"`
	// The On Demand resource key.
	ResourceKey string `json:"resource_key"`
	// The creator-designated SKU for the On Demand page.
	Sku NullableString `json:"sku,omitempty"`
	Subscription NullableOnDemandPageSubscription `json:"subscription"`
	// The graphical theme for the On Demand page.
	Theme string `json:"theme"`
	Thumbnail NullableOnDemandPageThumbnail `json:"thumbnail"`
	Trailer NullableOnDemandPageTrailer `json:"trailer"`
	// Whether the On Demand page is for a film or a series.  Option descriptions:  * `film` - The On Demand page is for a film.  * `series` - The On Demand page is for a series. 
	Type string `json:"type"`
	// The relative URI of the On Demand page.
	Uri string `json:"uri"`
	User NullableOnDemandPageUser `json:"user"`
}

type _OnDemandPage OnDemandPage

// NewOnDemandPage instantiates a new OnDemandPage object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewOnDemandPage(background NullableOnDemandPageBackground, colors OnDemandPageColors, contentRating []string, description NullableString, domainLink NullableString, episodes OnDemandPageEpisodes, genres []OnDemandGenre, link string, metadata OnDemandPageMetadata, name string, pictures NullableOnDemandPagePictures, preorder OnDemandPagePreorder, published OnDemandPagePublished, rating NullableFloat32, resourceKey string, subscription NullableOnDemandPageSubscription, theme string, thumbnail NullableOnDemandPageThumbnail, trailer NullableOnDemandPageTrailer, type_ string, uri string, user NullableOnDemandPageUser) *OnDemandPage {
	this := OnDemandPage{}
	this.Background = background
	this.Colors = colors
	this.ContentRating = contentRating
	this.Description = description
	this.DomainLink = domainLink
	this.Episodes = episodes
	this.Genres = genres
	this.Link = link
	this.Metadata = metadata
	this.Name = name
	this.Pictures = pictures
	this.Preorder = preorder
	this.Published = published
	this.Rating = rating
	this.ResourceKey = resourceKey
	this.Subscription = subscription
	this.Theme = theme
	this.Thumbnail = thumbnail
	this.Trailer = trailer
	this.Type = type_
	this.Uri = uri
	this.User = user
	return &this
}

// NewOnDemandPageWithDefaults instantiates a new OnDemandPage object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewOnDemandPageWithDefaults() *OnDemandPage {
	this := OnDemandPage{}
	return &this
}

// GetBackground returns the Background field value
// If the value is explicit nil, the zero value for OnDemandPageBackground will be returned
func (o *OnDemandPage) GetBackground() OnDemandPageBackground {
	if o == nil || o.Background.Get() == nil {
		var ret OnDemandPageBackground
		return ret
	}

	return *o.Background.Get()
}

// GetBackgroundOk returns a tuple with the Background field value
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *OnDemandPage) GetBackgroundOk() (*OnDemandPageBackground, bool) {
	if o == nil {
		return nil, false
	}
	return o.Background.Get(), o.Background.IsSet()
}

// SetBackground sets field value
func (o *OnDemandPage) SetBackground(v OnDemandPageBackground) {
	o.Background.Set(&v)
}

// GetColors returns the Colors field value
func (o *OnDemandPage) GetColors() OnDemandPageColors {
	if o == nil {
		var ret OnDemandPageColors
		return ret
	}

	return o.Colors
}

// GetColorsOk returns a tuple with the Colors field value
// and a boolean to check if the value has been set.
func (o *OnDemandPage) GetColorsOk() (*OnDemandPageColors, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Colors, true
}

// SetColors sets field value
func (o *OnDemandPage) SetColors(v OnDemandPageColors) {
	o.Colors = v
}

// GetContentRating returns the ContentRating field value
func (o *OnDemandPage) GetContentRating() []string {
	if o == nil {
		var ret []string
		return ret
	}

	return o.ContentRating
}

// GetContentRatingOk returns a tuple with the ContentRating field value
// and a boolean to check if the value has been set.
func (o *OnDemandPage) GetContentRatingOk() ([]string, bool) {
	if o == nil {
		return nil, false
	}
	return o.ContentRating, true
}

// SetContentRating sets field value
func (o *OnDemandPage) SetContentRating(v []string) {
	o.ContentRating = v
}

// GetCreatedTime returns the CreatedTime field value if set, zero value otherwise.
func (o *OnDemandPage) GetCreatedTime() string {
	if o == nil || IsNil(o.CreatedTime) {
		var ret string
		return ret
	}
	return *o.CreatedTime
}

// GetCreatedTimeOk returns a tuple with the CreatedTime field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *OnDemandPage) GetCreatedTimeOk() (*string, bool) {
	if o == nil || IsNil(o.CreatedTime) {
		return nil, false
	}
	return o.CreatedTime, true
}

// HasCreatedTime returns a boolean if a field has been set.
func (o *OnDemandPage) HasCreatedTime() bool {
	if o != nil && !IsNil(o.CreatedTime) {
		return true
	}

	return false
}

// SetCreatedTime gets a reference to the given string and assigns it to the CreatedTime field.
func (o *OnDemandPage) SetCreatedTime(v string) {
	o.CreatedTime = &v
}

// GetDescription returns the Description field value
// If the value is explicit nil, the zero value for string will be returned
func (o *OnDemandPage) GetDescription() string {
	if o == nil || o.Description.Get() == nil {
		var ret string
		return ret
	}

	return *o.Description.Get()
}

// GetDescriptionOk returns a tuple with the Description field value
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *OnDemandPage) GetDescriptionOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.Description.Get(), o.Description.IsSet()
}

// SetDescription sets field value
func (o *OnDemandPage) SetDescription(v string) {
	o.Description.Set(&v)
}

// GetDomainLink returns the DomainLink field value
// If the value is explicit nil, the zero value for string will be returned
func (o *OnDemandPage) GetDomainLink() string {
	if o == nil || o.DomainLink.Get() == nil {
		var ret string
		return ret
	}

	return *o.DomainLink.Get()
}

// GetDomainLinkOk returns a tuple with the DomainLink field value
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *OnDemandPage) GetDomainLinkOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.DomainLink.Get(), o.DomainLink.IsSet()
}

// SetDomainLink sets field value
func (o *OnDemandPage) SetDomainLink(v string) {
	o.DomainLink.Set(&v)
}

// GetEpisodes returns the Episodes field value
func (o *OnDemandPage) GetEpisodes() OnDemandPageEpisodes {
	if o == nil {
		var ret OnDemandPageEpisodes
		return ret
	}

	return o.Episodes
}

// GetEpisodesOk returns a tuple with the Episodes field value
// and a boolean to check if the value has been set.
func (o *OnDemandPage) GetEpisodesOk() (*OnDemandPageEpisodes, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Episodes, true
}

// SetEpisodes sets field value
func (o *OnDemandPage) SetEpisodes(v OnDemandPageEpisodes) {
	o.Episodes = v
}

// GetFilm returns the Film field value if set, zero value otherwise.
func (o *OnDemandPage) GetFilm() Video {
	if o == nil || IsNil(o.Film) {
		var ret Video
		return ret
	}
	return *o.Film
}

// GetFilmOk returns a tuple with the Film field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *OnDemandPage) GetFilmOk() (*Video, bool) {
	if o == nil || IsNil(o.Film) {
		return nil, false
	}
	return o.Film, true
}

// HasFilm returns a boolean if a field has been set.
func (o *OnDemandPage) HasFilm() bool {
	if o != nil && !IsNil(o.Film) {
		return true
	}

	return false
}

// SetFilm gets a reference to the given Video and assigns it to the Film field.
func (o *OnDemandPage) SetFilm(v Video) {
	o.Film = &v
}

// GetGenres returns the Genres field value
func (o *OnDemandPage) GetGenres() []OnDemandGenre {
	if o == nil {
		var ret []OnDemandGenre
		return ret
	}

	return o.Genres
}

// GetGenresOk returns a tuple with the Genres field value
// and a boolean to check if the value has been set.
func (o *OnDemandPage) GetGenresOk() ([]OnDemandGenre, bool) {
	if o == nil {
		return nil, false
	}
	return o.Genres, true
}

// SetGenres sets field value
func (o *OnDemandPage) SetGenres(v []OnDemandGenre) {
	o.Genres = v
}

// GetLink returns the Link field value
func (o *OnDemandPage) GetLink() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Link
}

// GetLinkOk returns a tuple with the Link field value
// and a boolean to check if the value has been set.
func (o *OnDemandPage) GetLinkOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Link, true
}

// SetLink sets field value
func (o *OnDemandPage) SetLink(v string) {
	o.Link = v
}

// GetMetadata returns the Metadata field value
func (o *OnDemandPage) GetMetadata() OnDemandPageMetadata {
	if o == nil {
		var ret OnDemandPageMetadata
		return ret
	}

	return o.Metadata
}

// GetMetadataOk returns a tuple with the Metadata field value
// and a boolean to check if the value has been set.
func (o *OnDemandPage) GetMetadataOk() (*OnDemandPageMetadata, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Metadata, true
}

// SetMetadata sets field value
func (o *OnDemandPage) SetMetadata(v OnDemandPageMetadata) {
	o.Metadata = v
}

// GetModifiedTime returns the ModifiedTime field value if set, zero value otherwise.
func (o *OnDemandPage) GetModifiedTime() string {
	if o == nil || IsNil(o.ModifiedTime) {
		var ret string
		return ret
	}
	return *o.ModifiedTime
}

// GetModifiedTimeOk returns a tuple with the ModifiedTime field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *OnDemandPage) GetModifiedTimeOk() (*string, bool) {
	if o == nil || IsNil(o.ModifiedTime) {
		return nil, false
	}
	return o.ModifiedTime, true
}

// HasModifiedTime returns a boolean if a field has been set.
func (o *OnDemandPage) HasModifiedTime() bool {
	if o != nil && !IsNil(o.ModifiedTime) {
		return true
	}

	return false
}

// SetModifiedTime gets a reference to the given string and assigns it to the ModifiedTime field.
func (o *OnDemandPage) SetModifiedTime(v string) {
	o.ModifiedTime = &v
}

// GetName returns the Name field value
func (o *OnDemandPage) GetName() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Name
}

// GetNameOk returns a tuple with the Name field value
// and a boolean to check if the value has been set.
func (o *OnDemandPage) GetNameOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Name, true
}

// SetName sets field value
func (o *OnDemandPage) SetName(v string) {
	o.Name = v
}

// GetPictures returns the Pictures field value
// If the value is explicit nil, the zero value for OnDemandPagePictures will be returned
func (o *OnDemandPage) GetPictures() OnDemandPagePictures {
	if o == nil || o.Pictures.Get() == nil {
		var ret OnDemandPagePictures
		return ret
	}

	return *o.Pictures.Get()
}

// GetPicturesOk returns a tuple with the Pictures field value
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *OnDemandPage) GetPicturesOk() (*OnDemandPagePictures, bool) {
	if o == nil {
		return nil, false
	}
	return o.Pictures.Get(), o.Pictures.IsSet()
}

// SetPictures sets field value
func (o *OnDemandPage) SetPictures(v OnDemandPagePictures) {
	o.Pictures.Set(&v)
}

// GetPreorder returns the Preorder field value
func (o *OnDemandPage) GetPreorder() OnDemandPagePreorder {
	if o == nil {
		var ret OnDemandPagePreorder
		return ret
	}

	return o.Preorder
}

// GetPreorderOk returns a tuple with the Preorder field value
// and a boolean to check if the value has been set.
func (o *OnDemandPage) GetPreorderOk() (*OnDemandPagePreorder, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Preorder, true
}

// SetPreorder sets field value
func (o *OnDemandPage) SetPreorder(v OnDemandPagePreorder) {
	o.Preorder = v
}

// GetPublished returns the Published field value
func (o *OnDemandPage) GetPublished() OnDemandPagePublished {
	if o == nil {
		var ret OnDemandPagePublished
		return ret
	}

	return o.Published
}

// GetPublishedOk returns a tuple with the Published field value
// and a boolean to check if the value has been set.
func (o *OnDemandPage) GetPublishedOk() (*OnDemandPagePublished, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Published, true
}

// SetPublished sets field value
func (o *OnDemandPage) SetPublished(v OnDemandPagePublished) {
	o.Published = v
}

// GetRating returns the Rating field value
// If the value is explicit nil, the zero value for float32 will be returned
func (o *OnDemandPage) GetRating() float32 {
	if o == nil || o.Rating.Get() == nil {
		var ret float32
		return ret
	}

	return *o.Rating.Get()
}

// GetRatingOk returns a tuple with the Rating field value
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *OnDemandPage) GetRatingOk() (*float32, bool) {
	if o == nil {
		return nil, false
	}
	return o.Rating.Get(), o.Rating.IsSet()
}

// SetRating sets field value
func (o *OnDemandPage) SetRating(v float32) {
	o.Rating.Set(&v)
}

// GetResourceKey returns the ResourceKey field value
func (o *OnDemandPage) GetResourceKey() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.ResourceKey
}

// GetResourceKeyOk returns a tuple with the ResourceKey field value
// and a boolean to check if the value has been set.
func (o *OnDemandPage) GetResourceKeyOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ResourceKey, true
}

// SetResourceKey sets field value
func (o *OnDemandPage) SetResourceKey(v string) {
	o.ResourceKey = v
}

// GetSku returns the Sku field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *OnDemandPage) GetSku() string {
	if o == nil || IsNil(o.Sku.Get()) {
		var ret string
		return ret
	}
	return *o.Sku.Get()
}

// GetSkuOk returns a tuple with the Sku field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *OnDemandPage) GetSkuOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.Sku.Get(), o.Sku.IsSet()
}

// HasSku returns a boolean if a field has been set.
func (o *OnDemandPage) HasSku() bool {
	if o != nil && o.Sku.IsSet() {
		return true
	}

	return false
}

// SetSku gets a reference to the given NullableString and assigns it to the Sku field.
func (o *OnDemandPage) SetSku(v string) {
	o.Sku.Set(&v)
}
// SetSkuNil sets the value for Sku to be an explicit nil
func (o *OnDemandPage) SetSkuNil() {
	o.Sku.Set(nil)
}

// UnsetSku ensures that no value is present for Sku, not even an explicit nil
func (o *OnDemandPage) UnsetSku() {
	o.Sku.Unset()
}

// GetSubscription returns the Subscription field value
// If the value is explicit nil, the zero value for OnDemandPageSubscription will be returned
func (o *OnDemandPage) GetSubscription() OnDemandPageSubscription {
	if o == nil || o.Subscription.Get() == nil {
		var ret OnDemandPageSubscription
		return ret
	}

	return *o.Subscription.Get()
}

// GetSubscriptionOk returns a tuple with the Subscription field value
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *OnDemandPage) GetSubscriptionOk() (*OnDemandPageSubscription, bool) {
	if o == nil {
		return nil, false
	}
	return o.Subscription.Get(), o.Subscription.IsSet()
}

// SetSubscription sets field value
func (o *OnDemandPage) SetSubscription(v OnDemandPageSubscription) {
	o.Subscription.Set(&v)
}

// GetTheme returns the Theme field value
func (o *OnDemandPage) GetTheme() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Theme
}

// GetThemeOk returns a tuple with the Theme field value
// and a boolean to check if the value has been set.
func (o *OnDemandPage) GetThemeOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Theme, true
}

// SetTheme sets field value
func (o *OnDemandPage) SetTheme(v string) {
	o.Theme = v
}

// GetThumbnail returns the Thumbnail field value
// If the value is explicit nil, the zero value for OnDemandPageThumbnail will be returned
func (o *OnDemandPage) GetThumbnail() OnDemandPageThumbnail {
	if o == nil || o.Thumbnail.Get() == nil {
		var ret OnDemandPageThumbnail
		return ret
	}

	return *o.Thumbnail.Get()
}

// GetThumbnailOk returns a tuple with the Thumbnail field value
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *OnDemandPage) GetThumbnailOk() (*OnDemandPageThumbnail, bool) {
	if o == nil {
		return nil, false
	}
	return o.Thumbnail.Get(), o.Thumbnail.IsSet()
}

// SetThumbnail sets field value
func (o *OnDemandPage) SetThumbnail(v OnDemandPageThumbnail) {
	o.Thumbnail.Set(&v)
}

// GetTrailer returns the Trailer field value
// If the value is explicit nil, the zero value for OnDemandPageTrailer will be returned
func (o *OnDemandPage) GetTrailer() OnDemandPageTrailer {
	if o == nil || o.Trailer.Get() == nil {
		var ret OnDemandPageTrailer
		return ret
	}

	return *o.Trailer.Get()
}

// GetTrailerOk returns a tuple with the Trailer field value
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *OnDemandPage) GetTrailerOk() (*OnDemandPageTrailer, bool) {
	if o == nil {
		return nil, false
	}
	return o.Trailer.Get(), o.Trailer.IsSet()
}

// SetTrailer sets field value
func (o *OnDemandPage) SetTrailer(v OnDemandPageTrailer) {
	o.Trailer.Set(&v)
}

// GetType returns the Type field value
func (o *OnDemandPage) GetType() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Type
}

// GetTypeOk returns a tuple with the Type field value
// and a boolean to check if the value has been set.
func (o *OnDemandPage) GetTypeOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Type, true
}

// SetType sets field value
func (o *OnDemandPage) SetType(v string) {
	o.Type = v
}

// GetUri returns the Uri field value
func (o *OnDemandPage) GetUri() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Uri
}

// GetUriOk returns a tuple with the Uri field value
// and a boolean to check if the value has been set.
func (o *OnDemandPage) GetUriOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Uri, true
}

// SetUri sets field value
func (o *OnDemandPage) SetUri(v string) {
	o.Uri = v
}

// GetUser returns the User field value
// If the value is explicit nil, the zero value for OnDemandPageUser will be returned
func (o *OnDemandPage) GetUser() OnDemandPageUser {
	if o == nil || o.User.Get() == nil {
		var ret OnDemandPageUser
		return ret
	}

	return *o.User.Get()
}

// GetUserOk returns a tuple with the User field value
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *OnDemandPage) GetUserOk() (*OnDemandPageUser, bool) {
	if o == nil {
		return nil, false
	}
	return o.User.Get(), o.User.IsSet()
}

// SetUser sets field value
func (o *OnDemandPage) SetUser(v OnDemandPageUser) {
	o.User.Set(&v)
}

func (o OnDemandPage) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o OnDemandPage) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["background"] = o.Background.Get()
	toSerialize["colors"] = o.Colors
	toSerialize["content_rating"] = o.ContentRating
	if !IsNil(o.CreatedTime) {
		toSerialize["created_time"] = o.CreatedTime
	}
	toSerialize["description"] = o.Description.Get()
	toSerialize["domain_link"] = o.DomainLink.Get()
	toSerialize["episodes"] = o.Episodes
	if !IsNil(o.Film) {
		toSerialize["film"] = o.Film
	}
	toSerialize["genres"] = o.Genres
	toSerialize["link"] = o.Link
	toSerialize["metadata"] = o.Metadata
	if !IsNil(o.ModifiedTime) {
		toSerialize["modified_time"] = o.ModifiedTime
	}
	toSerialize["name"] = o.Name
	toSerialize["pictures"] = o.Pictures.Get()
	toSerialize["preorder"] = o.Preorder
	toSerialize["published"] = o.Published
	toSerialize["rating"] = o.Rating.Get()
	toSerialize["resource_key"] = o.ResourceKey
	if o.Sku.IsSet() {
		toSerialize["sku"] = o.Sku.Get()
	}
	toSerialize["subscription"] = o.Subscription.Get()
	toSerialize["theme"] = o.Theme
	toSerialize["thumbnail"] = o.Thumbnail.Get()
	toSerialize["trailer"] = o.Trailer.Get()
	toSerialize["type"] = o.Type
	toSerialize["uri"] = o.Uri
	toSerialize["user"] = o.User.Get()
	return toSerialize, nil
}

func (o *OnDemandPage) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"background",
		"colors",
		"content_rating",
		"description",
		"domain_link",
		"episodes",
		"genres",
		"link",
		"metadata",
		"name",
		"pictures",
		"preorder",
		"published",
		"rating",
		"resource_key",
		"subscription",
		"theme",
		"thumbnail",
		"trailer",
		"type",
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

	varOnDemandPage := _OnDemandPage{}

	decoder := json.NewDecoder(bytes.NewReader(data))
	decoder.DisallowUnknownFields()
	err = decoder.Decode(&varOnDemandPage)

	if err != nil {
		return err
	}

	*o = OnDemandPage(varOnDemandPage)

	return err
}

type NullableOnDemandPage struct {
	value *OnDemandPage
	isSet bool
}

func (v NullableOnDemandPage) Get() *OnDemandPage {
	return v.value
}

func (v *NullableOnDemandPage) Set(val *OnDemandPage) {
	v.value = val
	v.isSet = true
}

func (v NullableOnDemandPage) IsSet() bool {
	return v.isSet
}

func (v *NullableOnDemandPage) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableOnDemandPage(val *OnDemandPage) *NullableOnDemandPage {
	return &NullableOnDemandPage{value: val, isSet: true}
}

func (v NullableOnDemandPage) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableOnDemandPage) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


