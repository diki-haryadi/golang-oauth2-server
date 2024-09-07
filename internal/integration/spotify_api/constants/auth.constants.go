package constants

type SpotifyGrantType string

const (
	SpotifyGrantTypeAuthorizationCode SpotifyGrantType = "authorization_code"
	SpotifyGrantTypeRefreshToken      SpotifyGrantType = "refresh_token"
)
