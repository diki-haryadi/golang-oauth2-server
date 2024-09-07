package spotify_api

import (
	"context"
	"github.com/rs/zerolog/log"
	"golang-standards-project-layout/internal/integration/spotify_api/model"
)

func (x *Module) ObtainToken(ctx context.Context, reqData *model.SpotifyAuthorizeData) (*model.SpotifyAuthorizeBodyRes, error) {
	//ctx, span := tracer.StartSpan(ctx, "api_call.spotify.ObtainToken", nil)
	//defer span.End()

	var (
		e   error
		res *model.SpotifyAuthorizeBodyRes
	)

	//res, e = x.requestToken(ctx, reqData.ToSpotifyAuthorizeBodyReq(x.spotifyConfig.Credentials.AuthorizeCallbackUrl).ToMapString())

	if e != nil {
		log.Err(e).Interface("data", reqData).Msg(e.Error())
		return nil, e
	}

	return res, nil
}

func (x *Module) RefreshToken(ctx context.Context, refreshToken string) (*model.SpotifyAuthorizeBodyRes, error) {
	//ctx, span := tracer.StartSpan(ctx, "api_call.spotify.RefreshToken", nil)
	//defer span.End()

	var (
		e       error
		reqData = &model.SpotifyRefreshTokenData{
			RefreshToken: refreshToken,
		}
		res *model.SpotifyAuthorizeBodyRes
	)

	res, e = x.requestToken(ctx, reqData.ToSpotifyAuthorizeBodyReq().ToMapString())

	if e != nil {
		log.Err(e).Interface("data", reqData).Msg(e.Error())
		return nil, e
	}

	return res, nil
}

func (x *Module) GenerateAuthorizeLink(ctx context.Context, user_id string) string {
	//_, span := tracer.StartSpan(ctx, "api_call.spotify.GenerateAuthorizeLink", nil)
	//defer span.End()

	//var (
	//	qParams = url.Values{}
	//	res     string
	//)

	//qParams.Add("response_type", "code")
	//qParams.Add("client_id", x.spotifyConfig.Credentials.ClientId)
	//qParams.Add("scope", x.spotifyConfig.Credentials.Scope)
	//qParams.Add("state", user_id)
	//qParams.Add("redirect_uri", x.spotifyConfig.Credentials.AuthorizeCallbackUrl)
	//
	//res = fmt.Sprintf("%s/%s?%s", x.spotifyConfig.AuthApi.HttpClient.BaseUrl, x.spotifyConfig.AuthApi.Endpoints.Authorize, qParams.Encode())

	return "res"
}
