package api

import (
	swagger "github.com/arsmn/fiber-swagger/v2"
	"github.com/gofiber/fiber/v2"
	"golang-standards-project-layout/internal/init/service"
	"golang-standards-project-layout/pkg/env"
)

type Options struct {
	DefaultTimeout int
	Service        *service.RestService
}

type API struct {
	options *Options
	service *service.RestService
}

func New(o *Options) *API {
	return &API{
		options: o,
		service: o.Service,
	}
}

func (a *API) Register(srv *fiber.App, hdl *service.RestHandlers, mw *service.Middlewares) {
	a.initSwagger(srv)
	a.registerAuthAPI(srv, hdl, mw)
}

func (a *API) initSwagger(srv *fiber.App) {
	if a.service.Config.Rest.EnableSwagger {
		if !env.IsProduction() {
			srv.Use("swagger", swagger.HandlerDefault)
		}
	}
}

func (a *API) registerAuthAPI(srv *fiber.App, hdl *service.RestHandlers, mw *service.Middlewares) {
	router := srv.Group("/auth")

	router.Get("/callback", hdl.AuthRestHandler.HandleLinkageCallback)
}
