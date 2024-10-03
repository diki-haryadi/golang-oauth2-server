package usersHttpController

import (
	"github.com/labstack/echo/v4"

	usersDomain "golang-oauth2-server/internal/modules/users/domain"
)

type Router struct {
	controller usersDomain.HttpController
}

func NewRouter(controller usersDomain.HttpController) *Router {
	return &Router{
		controller: controller,
	}
}

func (r *Router) Register(e *echo.Group) {
	e.POST("/users", r.controller.CreateUsers)
}
