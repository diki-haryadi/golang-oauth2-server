package usersUseCase

import (
	"context"
	"encoding/json"

	"github.com/segmentio/kafka-go"

	sampleExtServiceDomain "golang-oauth2-server/external/sample_ext_service/domain"
	usersDomain "golang-oauth2-server/internal/modules/users/domain"
	usersDto "golang-oauth2-server/internal/modules/users/dto"
)

type useCase struct {
	repository              usersDomain.Repository
	kafkaProducer           usersDomain.KafkaProducer
	sampleExtServiceUseCase sampleExtServiceDomain.SampleExtServiceUseCase
}

func NewUseCase(
	repository usersDomain.Repository,
	sampleExtServiceUseCase sampleExtServiceDomain.SampleExtServiceUseCase,
	kafkaProducer usersDomain.KafkaProducer,
) usersDomain.UseCase {
	return &useCase{
		repository:              repository,
		kafkaProducer:           kafkaProducer,
		sampleExtServiceUseCase: sampleExtServiceUseCase,
	}
}

func (uc *useCase) CreateUsers(ctx context.Context, req *usersDto.CreateUsersRequestDto) (*usersDto.CreateUsersResponseDto, error) {
	users, err := uc.repository.CreateUsers(ctx, req)
	if err != nil {
		return nil, err
	}

	// TODO : if err => return Marshal_Err_Exception
	jsonArticle, _ := json.Marshal(users)

	// if it has go keyword and if we pass the request context to it, it will terminate after request lifecycle.
	_ = uc.kafkaProducer.PublishCreateEvent(context.Background(), kafka.Message{
		Key:   []byte("Users"),
		Value: jsonArticle,
	})

	return users, err
}
