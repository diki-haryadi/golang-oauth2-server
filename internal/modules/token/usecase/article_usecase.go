package articleUseCase

import (
	"context"
	"encoding/json"

	"github.com/segmentio/kafka-go"

	sampleExtServiceDomain "golang-oauth2-server/external/sample_ext_service/domain"
	articleDomain "golang-oauth2-server/internal/modules/article/domain"
	articleDto "golang-oauth2-server/internal/modules/article/dto"
)

type useCase struct {
	repository              articleDomain.Repository
	kafkaProducer           articleDomain.KafkaProducer
	sampleExtServiceUseCase sampleExtServiceDomain.SampleExtServiceUseCase
}

func NewUseCase(
	repository articleDomain.Repository,
	sampleExtServiceUseCase sampleExtServiceDomain.SampleExtServiceUseCase,
	kafkaProducer articleDomain.KafkaProducer,
) articleDomain.UseCase {
	return &useCase{
		repository:              repository,
		kafkaProducer:           kafkaProducer,
		sampleExtServiceUseCase: sampleExtServiceUseCase,
	}
}

func (uc *useCase) CreateUsers(ctx context.Context, req *articleDto.CreateArticleRequestDto) (*articleDto.CreateArticleResponseDto, error) {
	article, err := uc.repository.CreateUsers(ctx, req)
	if err != nil {
		return nil, err
	}

	// TODO : if err => return Marshal_Err_Exception
	jsonArticle, _ := json.Marshal(article)

	// if it has go keyword and if we pass the request context to it, it will terminate after request lifecycle.
	_ = uc.kafkaProducer.PublishCreateEvent(context.Background(), kafka.Message{
		Key:   []byte("Article"),
		Value: jsonArticle,
	})

	return article, err
}
