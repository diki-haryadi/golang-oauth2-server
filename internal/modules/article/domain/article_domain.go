package articleDomain

import (
	"context"
	"github.com/labstack/echo/v4"
	"github.com/segmentio/kafka-go"

	"github.com/google/uuid"

	articleV1 "golang-oauth2-server/api/article/v1"
	articleDto "golang-oauth2-server/internal/modules/article/dto"
)

type Article struct {
	ID          uuid.UUID `db:"id" json:"id"`
	Name        string    `db:"name" json:"name"`
	Description string    `db:"description" json:"desc"`
}

type Configurator interface {
	Configure(ctx context.Context) error
}

type UseCase interface {
	CreateUsers(ctx context.Context, article *articleDto.CreateArticleRequestDto) (*articleDto.CreateArticleResponseDto, error)
}

type Repository interface {
	CreateUsers(ctx context.Context, article *articleDto.CreateArticleRequestDto) (*articleDto.CreateArticleResponseDto, error)
}

type GrpcController interface {
	CreateArticle(ctx context.Context, req *articleV1.CreateArticleRequest) (*articleV1.CreateArticleResponse, error)
	GetArticleById(ctx context.Context, req *articleV1.GetArticleByIdRequest) (*articleV1.GetArticleByIdResponse, error)
}

type HttpController interface {
	CreateUsers(c echo.Context) error
}

type Job interface {
	StartJobs(ctx context.Context)
}

type KafkaProducer interface {
	PublishCreateEvent(ctx context.Context, messages ...kafka.Message) error
}

type KafkaConsumer interface {
	RunConsumers(ctx context.Context)
}
