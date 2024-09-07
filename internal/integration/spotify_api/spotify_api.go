package spotify_api

import (
	"context"

	"golang-standards-project-layout/internal/integration/spotify_api/model"
)

type SpotifyApiCallIntegration interface {
	ObtainToken(ctx context.Context, reqBody *model.SpotifyAuthorizeData) (*model.SpotifyAuthorizeBodyRes, error)
	RefreshToken(ctx context.Context, refreshToken string) (*model.SpotifyAuthorizeBodyRes, error)
	GetUserInfo(ctx context.Context, accessToken string) (*model.SpotifyGetUserProfileBodyRes, error)
	GenerateAuthorizeLink(ctx context.Context, user_id string) string
	SearchTracks(ctx context.Context, accessToken string, keyword string) (*model.SpotifyTrackSearchBodyRes, error)
	AddItemToPlaybackQueue(ctx context.Context, accessToken string, trackUri string) error
}
