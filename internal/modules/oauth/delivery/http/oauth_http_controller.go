package articleHttpController

import (
	"context"
	"fmt"
	"github.com/labstack/echo/v4"
	clientDomain "golang-oauth2-server/internal/modules/client/domain"
	scopeDomain "golang-oauth2-server/internal/modules/scope/domain"
	"net/http"

	oauthDomain "golang-oauth2-server/internal/modules/oauth/domain"
	oauthDto "golang-oauth2-server/internal/modules/oauth/dto"
)

type controller struct {
	useCase       oauthDomain.UseCase
	useCaseScope  scopeDomain.UseCase
	useCaseClient clientDomain.UseCase
}

func NewController(
	uc oauthDomain.UseCase,
	useCaseScope scopeDomain.UseCase,
	useCaseClient clientDomain.UseCase,
) oauthDomain.HttpController {
	return &controller{
		useCase:       uc,
		useCaseScope:  useCaseScope,
		useCaseClient: useCaseClient,
	}
}

func (c controller) Tokens(ctx echo.Context) error {
	grantTypes := map[string]func(ctx context.Context, client *clientDomain.Client) (*oauthDto.AccessTokenResponse, error){
		"password": c.useCase.PasswordGrant,
	}

	grantHandler, ok := grantTypes[ctx.FormValue("grant_type")]
	if !ok {
		return ctx.JSON(http.StatusBadRequest, nil)
	}

	client, err := c.basicAuthClient(ctx)
	if err != nil {
		return ctx.JSON(http.StatusUnauthorized, nil)
	}

	resp, err := grantHandler(context.Background(), client)
	if err != nil {
		return ctx.JSON(http.StatusBadRequest, err)
	}

	return ctx.JSON(http.StatusOK, resp)
}

func (c controller) basicAuthClient(ctx echo.Context) (*clientDomain.Client, error) {
	clientId, secret, ok := ctx.Request().BasicAuth()
	if !ok {
		return nil, fmt.Errorf("Invalid client ID or secret")
	}

	client, err := c.useCaseClient.AuthClient(context.Background(), clientId, secret)
	if err != nil {
		return nil, fmt.Errorf("Invalid client ID or secret")
	}
	return client, nil
}
