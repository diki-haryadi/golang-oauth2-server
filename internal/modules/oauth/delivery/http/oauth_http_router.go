package articleHttpController

import (
	"github.com/labstack/echo/v4"
	oauthDomain "golang-oauth2-server/internal/modules/oauth/domain"
)

type Router struct {
	controller oauthDomain.HttpController
}

func NewRouter(controller oauthDomain.HttpController) *Router {
	return &Router{
		controller: controller,
	}
}

func (r *Router) Register(e *echo.Group) {
	e.POST("/oauth/tokens", r.controller.Tokens)
}
