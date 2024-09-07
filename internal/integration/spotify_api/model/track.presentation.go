package model

type SpotifyTrackBodyRes struct {
	Album            SpotifyAlbumBodyRes        `json:"album"`
	Artists          []SpotifyArtistMiniBodyRes `json:"artists"`
	AvailableMarkets []string                   `json:"available_markets"`
	DiscNumber       int                        `json:"disc_number"`
	DurationMs       int                        `json:"duration_ms"`
	Explicit         bool                       `json:"explicit"`
	ExternalIds      SpotifyExternalIdsBodyRes  `json:"external_ids"`
	ExternalUrls     SpotifyExternalUrlBodyRes  `json:"external_urls"`
	Href             string                     `json:"href"`
	ID               string                     `json:"id"`
	IsLocal          bool                       `json:"is_local"`
	Name             string                     `json:"name"`
	Popularity       int                        `json:"popularity"`
	PreviewURL       string                     `json:"preview_url"`
	TrackNumber      int                        `json:"track_number"`
	Type             string                     `json:"type"`
	URI              string                     `json:"uri"`
}
