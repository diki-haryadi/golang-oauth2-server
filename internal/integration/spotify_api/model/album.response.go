package model

type SpotifyAlbumBodyRes struct {
	AlbumType            string                     `json:"album_type"`
	Artists              []SpotifyArtistMiniBodyRes `json:"artists"`
	AvailableMarkets     []string                   `json:"available_markets"`
	ExternalUrls         SpotifyExternalUrlBodyRes  `json:"external_urls"`
	Href                 string                     `json:"href"`
	ID                   string                     `json:"id"`
	Images               []SpotifyImageBodyRes      `json:"images"`
	Name                 string                     `json:"name"`
	ReleaseDate          string                     `json:"release_date"`
	ReleaseDatePrecision string                     `json:"release_date_precision"`
	TotalTracks          int                        `json:"total_tracks"`
	Type                 string                     `json:"type"`
	URI                  string                     `json:"uri"`
}
