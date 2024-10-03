package healthCheckHttp

import (
	"github.com/labstack/echo/v4"

	healthCheckDomain "golang-oauth2-server/internal/modules/health_check/domain"
)

type Router struct {
	controller healthCheckDomain.HttpController
}

func NewRouter(controller healthCheckDomain.HttpController) *Router {
	return &Router{
		controller: controller,
	}
}

func (r *Router) Register(e *echo.Group) {
	e.GET("/health", r.controller.Check)
}
