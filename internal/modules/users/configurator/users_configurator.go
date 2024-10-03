package articleConfigurator

import (
	"context"

	usersV1 "golang-oauth2-server/api/users/v1"

	sampleExtServiceUseCase "golang-oauth2-server/external/sample_ext_service/usecase"
	usersGrpcController "golang-oauth2-server/internal/modules/users/delivery/grpc"
	usersHttpController "golang-oauth2-server/internal/modules/users/delivery/http"
	usersKafkaProducer "golang-oauth2-server/internal/modules/users/delivery/kafka/producer"
	usersDomain "golang-oauth2-server/internal/modules/users/domain"
	usersRepository "golang-oauth2-server/internal/modules/users/repository"
	usersUseCase "golang-oauth2-server/internal/modules/users/usecase"
	externalBridge "golang-oauth2-server/internal/pkg/external_bridge"
	infraContainer "golang-oauth2-server/internal/pkg/infra_container"
)

type configurator struct {
	ic        *infraContainer.IContainer
	extBridge *externalBridge.ExternalBridge
}

func NewConfigurator(ic *infraContainer.IContainer, extBridge *externalBridge.ExternalBridge) usersDomain.Configurator {
	return &configurator{ic: ic, extBridge: extBridge}
}

func (c *configurator) Configure(ctx context.Context) error {
	seServiceUseCase := sampleExtServiceUseCase.NewSampleExtServiceUseCase(c.extBridge.SampleExtGrpcService)
	kafkaProducer := usersKafkaProducer.NewProducer(c.ic.KafkaWriter)
	repository := usersRepository.NewRepository(c.ic.Postgres)
	useCase := usersUseCase.NewUseCase(repository, seServiceUseCase, kafkaProducer)

	// grpc
	grpcController := usersGrpcController.NewController(useCase)
	usersV1.RegisterUsersServiceServer(c.ic.GrpcServer.GetCurrentGrpcServer(), grpcController)

	// http
	httpRouterGp := c.ic.EchoHttpServer.GetEchoInstance().Group(c.ic.EchoHttpServer.GetBasePath())
	httpController := usersHttpController.NewController(useCase)
	usersHttpController.NewRouter(httpController).Register(httpRouterGp)

	// consumers
	//usersKafkaConsumer.NewConsumer(c.ic.KafkaReader).RunConsumers(ctx)

	// jobs
	//usersJob.NewJob(c.ic.Logger).StartJobs(ctx)

	return nil
}
