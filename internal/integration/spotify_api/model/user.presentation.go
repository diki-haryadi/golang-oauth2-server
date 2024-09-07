package model

type SpotifyGetUserProfileBodyRes struct {
	Country         string                                      `json:"country"`
	DisplayName     string                                      `json:"display_name"`
	Email           string                                      `json:"email"`
	ExplicitContent SpotifyGetUserProfileExplicitContentBodyRes `json:"explicit_content"`
	ExternalUrls    SpotifyGetUserProfileExternalUrlsBodyRes    `json:"external_urls"`
	Followers       SpotifyGetUserProfileFollowersBodyRes       `json:"followers"`
	Href            string                                      `json:"href"`
	Id              string                                      `json:"id"`
	Images          []SpotifyGetUserProfileImagesBodyRes        `json:"images"`
	Product         string                                      `json:"product"`
	Type            string                                      `json:"type"`
	Uri             string                                      `json:"uri"`
}

type SpotifyGetUserProfileExplicitContentBodyRes struct {
	FilterEnabled bool `json:"filter_enabled"`
	FilterLocked  bool `json:"filter_locked"`
}

type SpotifyGetUserProfileExternalUrlsBodyRes struct {
	Spotify string `json:"spotify"`
}

type SpotifyGetUserProfileFollowersBodyRes struct {
	Href  string `json:"href"`
	Total int    `json:"total"`
}

type SpotifyGetUserProfileImagesBodyRes struct {
	Url    string `json:"url"`
	Height int    `json:"height"`
	Width  int    `json:"width"`
}
