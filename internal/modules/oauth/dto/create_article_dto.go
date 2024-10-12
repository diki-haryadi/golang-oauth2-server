package articleDto

import (
	validator "github.com/go-ozzo/ozzo-validation"
	tokenDomain "golang-oauth2-server/internal/modules/token/domain"
)

type CreateOauthRequestDto struct {
	Name        string `json:"name"`
	Description string `json:"desc"`
}

func (caDto *CreateOauthRequestDto) ValidateCreateArticleDto() error {
	return validator.ValidateStruct(caDto,
		validator.Field(
			&caDto.Name,
			validator.Required,
			validator.Length(3, 50),
		),
		validator.Field(
			&caDto.Description,
			validator.Required,
			validator.Length(5, 100),
		),
	)
}

type AccessTokenResponse struct {
	UserID       string `json:"user_id,omitempty"`
	AccessToken  string `json:"access_token"`
	ExpiresIn    int    `json:"expires_in"`
	TokenType    string `json:"token_type"`
	Scope        string `json:"scope"`
	RefreshToken string `json:"refresh_token,omitempty"`
}

func NewAccessTokenResponse(accessToken *tokenDomain.AccessTokenDomain, refreshToken *tokenDomain.RefreshToken, lifetime int, theTokenType string) (*AccessTokenResponse, error) {
	response := &AccessTokenResponse{
		AccessToken: accessToken.Token,
		ExpiresIn:   lifetime,
		TokenType:   theTokenType,
		Scope:       accessToken.Scope,
	}
	if accessToken.UserID.Valid {
		response.UserID = accessToken.UserID.String
	}
	if refreshToken != nil {
		response.RefreshToken = refreshToken.Token
	}
	return response, nil
}
