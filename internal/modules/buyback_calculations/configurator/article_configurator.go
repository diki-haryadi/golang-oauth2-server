package articleConfigurator

import (
	"context"

	articleV1 "golang-oauth2-server/api/article/v1"

	sampleExtServiceUseCase "golang-oauth2-server/external/sample_ext_service/usecase"
	articleGrpcController "golang-oauth2-server/internal/modules/article/delivery/grpc"
	articleHttpController "golang-oauth2-server/internal/modules/article/delivery/http"
	articleKafkaConsumer "golang-oauth2-server/internal/modules/article/delivery/kafka/consumer"
	articleKafkaProducer "golang-oauth2-server/internal/modules/article/delivery/kafka/producer"
	articleDomain "golang-oauth2-server/internal/modules/article/domain"
	articleJob "golang-oauth2-server/internal/modules/article/job"
	articleRepository "golang-oauth2-server/internal/modules/article/repository"
	articleUseCase "golang-oauth2-server/internal/modules/article/usecase"
	externalBridge "golang-oauth2-server/pkg/external_bridge"
	infraContainer "golang-oauth2-server/pkg/infra_container"
)

type configurator struct {
	ic        *infraContainer.IContainer
	extBridge *externalBridge.ExternalBridge
}

func NewConfigurator(ic *infraContainer.IContainer, extBridge *externalBridge.ExternalBridge) articleDomain.Configurator {
	return &configurator{ic: ic, extBridge: extBridge}
}

func (c *configurator) Configure(ctx context.Context) error {
	seServiceUseCase := sampleExtServiceUseCase.NewSampleExtServiceUseCase(c.extBridge.SampleExtGrpcService)
	kafkaProducer := articleKafkaProducer.NewProducer(c.ic.KafkaWriter)
	repository := articleRepository.NewRepository(c.ic.Postgres)
	useCase := articleUseCase.NewUseCase(repository, seServiceUseCase, kafkaProducer)

	// grpc
	grpcController := articleGrpcController.NewController(useCase)
	articleV1.RegisterArticleServiceServer(c.ic.GrpcServer.GetCurrentGrpcServer(), grpcController)

	// http
	httpRouterGp := c.ic.EchoHttpServer.GetEchoInstance().Group(c.ic.EchoHttpServer.GetBasePath())
	httpController := articleHttpController.NewController(useCase)
	articleHttpController.NewRouter(httpController).Register(httpRouterGp)

	// consumers
	articleKafkaConsumer.NewConsumer(c.ic.KafkaReader).RunConsumers(ctx)

	// jobs
	articleJob.NewJob(c.ic.Logger).StartJobs(ctx)

	return nil
}
