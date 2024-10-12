package tokenDomain

import (
	"context"
	"database/sql"
	"github.com/labstack/echo/v4"
	"github.com/segmentio/kafka-go"
	clientDomain "golang-oauth2-server/internal/modules/client/domain"
	"golang-oauth2-server/internal/modules/common/domain"
	usersDomain "golang-oauth2-server/internal/modules/users/domain"
	"time"

	articleV1 "golang-oauth2-server/api/article/v1"
	articleDto "golang-oauth2-server/internal/modules/article/dto"
)

type TokenDomain struct {
	commonDomain.CommonModel
	ClientID  sql.NullString `db:"client_id"`
	UserID    sql.NullString `db:"user_id"`
	Client    *clientDomain.Client
	User      *usersDomain.Users
	Token     string    `sql:"token"`
	ExpiresAt time.Time `sql:"expires_at"`
	Scope     string    `sql:"scope"`
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
