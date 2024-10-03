package usersDomain

import (
	"context"

	"github.com/google/uuid"

	"github.com/labstack/echo/v4"
	"github.com/segmentio/kafka-go"
	usersV1 "golang-oauth2-server/api/users/v1"

	usersDto "golang-oauth2-server/internal/modules/users/dto"
)

type Users struct {
	ID       uuid.UUID `db:"id" json:"id"`
	Username string    `db:"username" json:"username"`
	Email    string    `db:"email" json:"email"`
	Password string    `db:"password" json:"password"`
}

type Configurator interface {
	Configure(ctx context.Context) error
}

type UseCase interface {
	CreateUsers(ctx context.Context, user *usersDto.CreateUsersRequestDto) (*usersDto.CreateUsersResponseDto, error)
}

type Repository interface {
	CreateUsers(ctx context.Context, article *usersDto.CreateUsersRequestDto) (*usersDto.CreateUsersResponseDto, error)
}

type GrpcController interface {
	CreateUsers(ctx context.Context, req *usersV1.CreateUsersRequest) (*usersV1.CreateUsersResponse, error)
	GetUsersById(ctx context.Context, req *usersV1.GetUsersByIdRequest) (*usersV1.GetUsersByIdResponse, error)
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
