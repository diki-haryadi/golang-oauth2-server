package clientConfigurator

import (
	"context"
	clientDomain "golang-oauth2-server/internal/modules/client/domain"
	clientRepository "golang-oauth2-server/internal/modules/client/repository"
	clientUseCase "golang-oauth2-server/internal/modules/client/usecase"
	externalBridge "golang-oauth2-server/pkg/external_bridge"
	infraContainer "golang-oauth2-server/pkg/infra_container"
)

type configurator struct {
	ic        *infraContainer.IContainer
	extBridge *externalBridge.ExternalBridge
}

func NewConfigurator(ic *infraContainer.IContainer, extBridge *externalBridge.ExternalBridge) clientDomain.Configurator {
	return &configurator{ic: ic, extBridge: extBridge}
}

func (c *configurator) Configure(ctx context.Context) error {
	repository := clientRepository.NewRepository(c.ic.Postgres)
	_ = clientUseCase.NewUseCase(repository)

	// grpc
	//grpcController := articleGrpcController.NewController(useCase)
	//articleV1.RegisterArticleServiceServer(c.ic.GrpcServer.GetCurrentGrpcServer(), grpcController)

	// http
	//httpRouterGp := c.ic.EchoHttpServer.GetEchoInstance().Group(c.ic.EchoHttpServer.GetBasePath())
	//httpController := articleHttpController.NewController(useCase)
	//articleHttpController.NewRouter(httpController).Register(httpRouterGp)

	// consumers
	//articleKafkaConsumer.NewConsumer(c.ic.KafkaReader).RunConsumers(ctx)

	// jobs
	//articleJob.NewJob(c.ic.Logger).StartJobs(ctx)

	return nil
}
