package spotify_api

import (
	"context"
	"golang-standards-project-layout/internal/integration/spotify_api/model"
)

func (x *Module) GetUserInfo(ctx context.Context, accessToken string) (jsonRes *model.SpotifyGetUserProfileBodyRes, err error) {
	//_, span := tracer.StartSpan(ctx, "api_call.spotify.GetUserInfo", nil)
	//defer span.End()
	//
	//var (
	//	e         error
	//	reqHeader = map[string]string{}
	//	reqUrl    = fmt.Sprintf("%s/%s", x.spotifyConfig.CoreApi.HttpClient.BaseUrl, x.spotifyConfig.CoreApi.Endpoints.GetUserInfo)
	//	hcRes     *http.Response
	//	jsonRes   = &model.SpotifyGetUserProfileBodyRes{}
	//)
	//
	//reqHeader["Authorization"] = fmt.Sprintf("Bearer %s", accessToken)
	//hcRes, e = x.coreHttpClient.Get(reqUrl, reqHeader)
	//
	//if e != nil {
	//	log.Err(e).Interface("at", accessToken).Interface("headers", reqHeader).Msg(e.Error())
	//	return nil, e
	//}
	//
	//if hcRes.StatusCode > 299 {
	//	e = iutil.NewBadRequestErr("Request failed")
	//	return nil, e
	//}
	//
	//e = util.ParseResponseBodyToJson(hcRes, jsonRes)
	//
	//if e != nil {
	//	log.Err(e).Msg(e.Error())
	//	return nil, e
	//}

	return jsonRes, nil
}
