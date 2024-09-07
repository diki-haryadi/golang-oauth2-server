package model

type SpotifyAuthorizeBodyReq struct {
	RedirectUri string `json:"redirect_uri"`
	GrantType   string `json:"grant_type"`
	Code        string `json:"code"`
}

type SpotifyRefreshTokenBodyReq struct {
	RefreshToken string `json:"refresh_token"`
	GrantType    string `json:"grant_type"`
}

type SpotifyAuthorizeBodyRes struct {
	AccessToken  string `json:"access_token"`
	RefreshToken string `json:"refresh_token"`
	ExpiresIn    int    `json:"expires_in"`
	TokenType    string `json:"token_type"`
	Scope        string `json:"scope"`
}

func (x *SpotifyAuthorizeBodyReq) ToMapString() map[string]string {
	return map[string]string{
		"redirect_uri": x.RedirectUri,
		"grant_type":   x.GrantType,
		"code":         x.Code,
	}
}

func (x *SpotifyRefreshTokenBodyReq) ToMapString() map[string]string {
	return map[string]string{
		"refresh_token": x.RefreshToken,
		"grant_type":    x.GrantType,
	}
}
