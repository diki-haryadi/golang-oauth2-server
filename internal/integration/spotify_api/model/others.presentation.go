package model

type SpotifyExternalUrlBodyRes struct {
	Spotify string `json:"spotify"`
}
type SpotifyFollowersBodyRes struct {
	Href  string `json:"href"`
	Total int    `json:"total"`
}
type SpotifyExternalIdsBodyRes struct {
	Isrc string `json:"isrc"`
}

type SpotifyImageBodyRes struct {
	Height int    `json:"height"`
	URL    string `json:"url"`
	Width  int    `json:"width"`
}

type SpotifyPaginationBodyRes struct {
	Limit    int         `json:"limit"`
	Next     string      `json:"next"`
	Offset   int         `json:"offset"`
	Previous interface{} `json:"previous"`
	Total    int         `json:"total"`
}
