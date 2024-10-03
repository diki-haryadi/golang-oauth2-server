package articleHttpController

import (
	articleDomain "golang-oauth2-server/internal/modules/article/domain"
)

type Router struct {
	controller articleDomain.HttpController
}

func NewRouter(controller articleDomain.HttpController) *Router {
	return &Router{
		controller: controller,
	}
}

func (r *Router) Register(e *echo.Group) {
	e.POST("/article", r.controller.CreateUsers)
}
