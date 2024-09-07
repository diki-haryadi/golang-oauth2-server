package spotify_api

type Opts struct {
	//SpotifyConfig *config.SpotifyConfig
}

type Module struct {
	//spotifyConfig  *config.SpotifyConfig
	//authHttpClient httpclient.HTTPClient
	//coreHttpClient httpclient.HTTPClient
}

//
//func New(o Opts) *Module {
//	authClient := httpclient.NewClient(httpclient.HttpClientConfig{
//		Timeout:    time.Duration(o.SpotifyConfig.AuthApi.HttpClient.Timeout) * time.Second,
//		RetryCount: o.SpotifyConfig.AuthApi.HttpClient.RetryCount,
//	}, "spotify-auth-apicall-integration")
//
//	coreClient := httpclient.NewClient(httpclient.HttpClientConfig{
//		Timeout:    time.Duration(o.SpotifyConfig.CoreApi.HttpClient.Timeout) * time.Second,
//		RetryCount: o.SpotifyConfig.CoreApi.HttpClient.RetryCount,
//	}, "spotify-core-apicall-integration")
//
//	return &Module{
//		spotifyConfig:  o.SpotifyConfig,
//		authHttpClient: authClient,
//		coreHttpClient: coreClient,
//	}
//}
