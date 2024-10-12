package articleConfigurator

import (
	"context"
	articleV1 "golang-oauth2-server/api/article/v1"
	clientRepository "golang-oauth2-server/internal/modules/client/repository"
	clientUseCase "golang-oauth2-server/internal/modules/client/usecase"
	scopeRepository "golang-oauth2-server/internal/modules/scope/repository"
	scopeUseCase "golang-oauth2-server/internal/modules/scope/usecase"

	sampleExtServiceUseCase "golang-oauth2-server/external/sample_ext_service/usecase"
	oauthGrpcController "golang-oauth2-server/internal/modules/oauth/delivery/grpc"
	oauthHttpController "golang-oauth2-server/internal/modules/oauth/delivery/http"
	oauthKafkaProducer "golang-oauth2-server/internal/modules/oauth/delivery/kafka/producer"
	oauthDomain "golang-oauth2-server/internal/modules/oauth/domain"
	oauthRepository "golang-oauth2-server/internal/modules/oauth/repository"
	oauthUseCase "golang-oauth2-server/internal/modules/oauth/usecase"
	externalBridge "golang-oauth2-server/pkg/external_bridge"
	infraContainer "golang-oauth2-server/pkg/infra_container"
)

type configurator struct {
	ic        *infraContainer.IContainer
	extBridge *externalBridge.ExternalBridge
}

func NewConfigurator(ic *infraContainer.IContainer, extBridge *externalBridge.ExternalBridge) oauthDomain.Configurator {
	return &configurator{ic: ic, extBridge: extBridge}
}

func (c *configurator) Configure(ctx context.Context) error {
	seServiceUseCase := sampleExtServiceUseCase.NewSampleExtServiceUseCase(c.extBridge.SampleExtGrpcService)
	kafkaProducer := oauthKafkaProducer.NewProducer(c.ic.KafkaWriter)

	// init for oauth
	repository := oauthRepository.NewRepository(c.ic.Postgres)
	useCase := oauthUseCase.NewUseCase(repository, seServiceUseCase, kafkaProducer)

	// init for scope
	repositoryScope := scopeRepository.NewRepository(c.ic.Postgres)
	useCaseScope := scopeUseCase.NewUseCase(repositoryScope)

	// init for client
	repositoryClient := clientRepository.NewRepository(c.ic.Postgres)
	useCaseClient := clientUseCase.NewUseCase(repositoryClient)

	// grpc
	grpcController := oauthGrpcController.NewController(useCase)
	articleV1.RegisterArticleServiceServer(c.ic.GrpcServer.GetCurrentGrpcServer(), grpcController)

	// http
	httpRouterGp := c.ic.EchoHttpServer.GetEchoInstance().Group(c.ic.EchoHttpServer.GetBasePath())
	httpController := oauthHttpController.NewController(useCase, useCaseScope, useCaseClient)
	oauthHttpController.NewRouter(httpController).Register(httpRouterGp)

	// consumers
	//articleKafkaConsumer.NewConsumer(c.ic.KafkaReader).RunConsumers(ctx)

	// jobs
	//articleJob.NewJob(c.ic.Logger).StartJobs(ctx)

	return nil
}
