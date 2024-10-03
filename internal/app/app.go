package app

import (
	"context"
	"os"
	"os/signal"
	"syscall"

	articleConfigurator "golang-oauth2-server/internal/modules/article/configurator"
	healthCheckConfigurator "golang-oauth2-server/internal/modules/health_check/configurator"
	usersConfigurator "golang-oauth2-server/internal/modules/users/configurator"
	externalBridge "golang-oauth2-server/pkg/external_bridge"
	iContainer "golang-oauth2-server/pkg/infra_container"
)

type App struct{}

func New() *App {
	return &App{}
}

func (a *App) Run() error {
	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()

	ic, infraDown, err := iContainer.NewIC(ctx)
	if err != nil {
		return err
	}
	defer infraDown()

	extBridge, extBridgeDown, err := externalBridge.NewExternalBridge(ctx)
	if err != nil {
		return err
	}
	defer extBridgeDown()

	me := configureModule(ctx, ic, extBridge)
	if me != nil {
		return me
	}

	var serverError error
	go func() {
		if err := ic.GrpcServer.RunGrpcServer(ctx, nil); err != nil {
			ic.Logger.Sugar().Errorf("(s.RunGrpcServer) err: {%v}", err)
			serverError = err
			cancel()
		}
	}()

	go func() {
		if err := ic.EchoHttpServer.RunServer(ctx, nil); err != nil {
			ic.Logger.Sugar().Errorf("(s.RunEchoServer) err: {%v}", err)
			serverError = err
			cancel()
		}
	}()

	quit := make(chan os.Signal, 1)
	signal.Notify(quit, os.Interrupt, syscall.SIGTERM)

	select {
	case v := <-quit:
		ic.Logger.Sugar().Errorf("signal.Notify: %v", v)
	case done := <-ctx.Done():
		ic.Logger.Sugar().Errorf("ctx.Done: %v", done)
	}

	ic.Logger.Sugar().Info("Server Exited Properly")
	return serverError
}

func configureModule(ctx context.Context, ic *iContainer.IContainer, extBridge *externalBridge.ExternalBridge) error {
	err := articleConfigurator.NewConfigurator(ic, extBridge).Configure(ctx)
	if err != nil {
		return err
	}

	err = usersConfigurator.NewConfigurator(ic, extBridge).Configure(ctx)
	if err != nil {
		return err
	}

	err = healthCheckConfigurator.NewConfigurator(ic).Configure(ctx)
	if err != nil {
		return err
	}

	return nil
}
