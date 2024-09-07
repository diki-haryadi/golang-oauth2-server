package rest

import (
	"fmt"
	"github.com/gofiber/fiber/v2"
	"time"

	"golang-standards-project-layout/api/rest/api"
	"golang-standards-project-layout/internal/init/service"
)

type Opts struct {
	ListenAddress string
	Port          string
	ReadTimeout   time.Duration
	WriteTimeout  time.Duration
	Service       *service.RestService
}

type Handler struct {
	options     *Opts
	server      *fiber.App
	listenErrCh chan error
}

type WebHandler interface {
	Run()
	ListenError() <-chan error
}

func New(o *Opts) WebHandler {
	srv := fiber.New(fiber.Config{
		ReadTimeout:  o.ReadTimeout,
		WriteTimeout: o.WriteTimeout,
	})

	// Liveness check api
	srv.Get("/health", func(fc *fiber.Ctx) error {
		return fc.JSON("OK")
	})

	api.New(&api.Options{
		DefaultTimeout: o.Service.Config.Rest.GracefulTimeout,
		Service:        o.Service,
	}).Register(srv, o.Service.RestHandlers, o.Service.Middlewares)

	handler := &Handler{
		options:     o,
		server:      srv,
		listenErrCh: make(chan error),
	}

	return handler
}

func (h *Handler) Run() {
	listenAddress := fmt.Sprintf("%s:%s", h.options.ListenAddress, h.options.Port)
	h.listenErrCh <- h.server.Listen(listenAddress)

}

func (h *Handler) ListenError() <-chan error {
	return h.listenErrCh
}
