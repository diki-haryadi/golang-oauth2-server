package model

type SpotifyArtistBodyRes struct {
	ExternalUrls SpotifyExternalUrlBodyRes `json:"external_urls"`
	Followers    SpotifyFollowersBodyRes   `json:"followers"`
	Genres       []string                  `json:"genres"`
	Href         string                    `json:"href"`
	ID           string                    `json:"id"`
	Images       []SpotifyImageBodyRes     `json:"images"`
	Name         string                    `json:"name"`
	Popularity   int                       `json:"popularity"`
	Type         string                    `json:"type"`
	URI          string                    `json:"uri"`
}

type SpotifyArtistMiniBodyRes struct {
	ExternalUrls SpotifyExternalUrlBodyRes `json:"external_urls"`
	Href         string                    `json:"href"`
	ID           string                    `json:"id"`
	Name         string                    `json:"name"`
	Type         string                    `json:"type"`
	URI          string                    `json:"uri"`
}
