package spotify_api

import (
	"context"

	"github.com/rs/zerolog/log"
	"golang-standards-project-layout/internal/integration/spotify_api/constants"
	"golang-standards-project-layout/internal/integration/spotify_api/model"
)

func (x *Module) SearchTracks(ctx context.Context, accessToken string, keyword string) (*model.SpotifyTrackSearchBodyRes, error) {
	//_, span := tracer.StartSpan(ctx, "api_call.spotify.SearchTracks", nil)
	//defer span.End()

	var (
		e   error
		res = &model.SpotifySearchResultBodyRes{}
	)

	res, e = x.searchForItem(ctx, accessToken, keyword, constants.SpotifySearchTypeTrack)

	if e != nil {

		log.Err(e).Msg(e.Error())
		return nil, e
	}

	return &res.Tracks, nil
}
