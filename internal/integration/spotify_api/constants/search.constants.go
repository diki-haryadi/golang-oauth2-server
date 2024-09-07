package constants

type SpotifySearchType string

const (
	SpotifySearchTypeAlbum    SpotifySearchType = "album"
	SpotifySearchTypeArtist   SpotifySearchType = "artist"
	SpotifySearchTypePlaylist SpotifySearchType = "playlist"
	SpotifySearchTypeTrack    SpotifySearchType = "track"
	// Commented since we're not goint to use this atm
	// SpotifySearchTypeShow     SpotifySearchType = "show"
	// SpotifySearchTypeEpisode  SpotifySearchType = "episode"
)
