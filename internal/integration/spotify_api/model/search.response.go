package model

type SpotifySearchResultBodyRes struct {
	Albums    SpotifyAlbumSearchBodyRes    `json:"albums"`
	Artists   SpotifyArtistSearchBodyRes   `json:"artists"`
	Tracks    SpotifyTrackSearchBodyRes    `json:"tracks"`
	Playlists SpotifyPlaylistSearchBodyRes `json:"playlists"`
}

type SpotifyAlbumSearchBodyRes struct {
	Href  string                `json:"href"`
	Items []SpotifyAlbumBodyRes `json:"items"`
	SpotifyPaginationBodyRes
}

type SpotifyArtistSearchBodyRes struct {
	Href  string               `json:"href"`
	Items SpotifyArtistBodyRes `json:"items"`
	SpotifyPaginationBodyRes
}

type SpotifyPlaylistSearchBodyRes struct {
	Href  string                   `json:"href"`
	Items []SpotifyPlaylistBodyRes `json:"items"`
	SpotifyPaginationBodyRes
}

type SpotifyTrackSearchBodyRes struct {
	Href  string                `json:"href"`
	Items []SpotifyTrackBodyRes `json:"items"`
	SpotifyPaginationBodyRes
}
