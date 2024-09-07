package spotify_api

import (
	"context"
)

func (x *Module) AddItemToPlaybackQueue(ctx context.Context, accessToken string, trackUri string) error {
	//_, span := tracer.StartSpan(ctx, "api_call.spotify.AddItemToPlaybackQueue", nil)
	//defer span.End()

	//var (
	//	e         error
	//	qParams   = url.Values{}
	//	reqHeader = map[string]string{}
	//	reqBody   = model.SpotifyAddItemToPlaybackQueueBodyReq{
	//		URI: trackUri,
	//	}
	//	reqBodyBuffer *bytes.Buffer
	//	reqUrl        = fmt.Sprintf("%s/%s", x.spotifyConfig.CoreApi.HttpClient.BaseUrl, x.spotifyConfig.CoreApi.Endpoints.QueueTrack)
	//	hcRes         *http.Response
	//	jsonRes       = &model.SpotifySearchResultBodyRes{}
	//)
	//
	//reqHeader["Authorization"] = fmt.Sprintf("Bearer %s", accessToken)
	//
	//reqBodyBuffer, e = util.ParseJsonStructToBytesBuffer(reqBody)
	//if e != nil {
	//	log.Err(e).Str("causer", "error_parsing_request").Msg(e.Error())
	//	return e
	//}
	//
	//hcRes, e = x.coreHttpClient.Post(reqUrl, reqBodyBuffer, reqHeader)
	//
	//if e != nil {
	//	log.Err(e).Str("qParams", qParams.Encode()).Interface("headers", reqHeader).Msg(e.Error())
	//	return e
	//}
	//
	//if hcRes.StatusCode > 299 {
	//	e = iutil.NewBadRequestErr("Request failed")
	//	return e
	//}
	//
	//e = util.ParseResponseBodyToJson(hcRes, jsonRes)
	//
	//if e != nil {
	//	log.Err(e).Msg(e.Error())
	//	return e
	//}

	return nil
}
