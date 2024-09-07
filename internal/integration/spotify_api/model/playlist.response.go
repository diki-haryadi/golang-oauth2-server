package model

type SpotifyPlaylistBodyRes struct {
	Collaborative bool                         `json:"collaborative"`
	Description   string                       `json:"description"`
	ExternalUrls  SpotifyExternalUrlBodyRes    `json:"external_urls"`
	Href          string                       `json:"href"`
	ID            string                       `json:"id"`
	Images        []SpotifyImageBodyRes        `json:"images"`
	Name          string                       `json:"name"`
	Owner         SpotifyPlaylistOwnerBodyRes  `json:"owner"`
	PrimaryColor  string                       `json:"primary_color"`
	Public        string                       `json:"public"`
	SnapshotID    string                       `json:"snapshot_id"`
	Tracks        SpotifyPlaylistTracksBodyRes `json:"tracks"`
	Type          string                       `json:"type"`
	URI           string                       `json:"uri"`
}

type SpotifyPlaylistOwnerBodyRes struct {
	DisplayName  string                    `json:"display_name"`
	ExternalUrls SpotifyExternalUrlBodyRes `json:"external_urls"`
	Href         string                    `json:"href"`
	ID           string                    `json:"id"`
	Type         string                    `json:"type"`
	URI          string                    `json:"uri"`
}

type SpotifyPlaylistTracksBodyRes struct {
	Href  string `json:"href"`
	Total int    `json:"total"`
}
