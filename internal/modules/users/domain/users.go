package usersDomain

import (
	"context"
	"database/sql"
	"golang-oauth2-server/internal/modules/common/domain"

	"github.com/labstack/echo/v4"
	"github.com/segmentio/kafka-go"
	usersV1 "golang-oauth2-server/api/users/v1"

	usersDto "golang-oauth2-server/internal/modules/users/dto"
)

type Users struct {
	commonDomain.CommonModel
	RoleID   sql.NullString `db:"role_id" json:"role_id"`
	Role     *Role
	Username string `db:"username" json:"username"`
	Password string `db:"password" json:"password"`
}

type Configurator interface {
	Configure(ctx context.Context) error
}

type UseCase interface {
	AuthenticateUser(ctx context.Context, username, password string) (*Users, error)
	CreateUsers(ctx context.Context, user *usersDto.CreateUsersRequestDto) (*usersDto.CreateUsersResponseDto, error)
	CreateUserTx(ctx context.Context, roleID, username, password string) (*Users, error)
	UpdateUsername(ctx context.Context, users *Users, username string) error
	SetPassword(ctx context.Context, users *Users) error
	SetPasswordTx(ctx context.Context, users *Users) error
}

type Repository interface {
	CreateUserCommon(ctx context.Context, roleID, username, password string) (*Users, error)
	UpdateUsername(ctx context.Context, users *Users, username string) error
	UpdateUsernameTx(ctx context.Context, users *Users, username string) error
	SetPassword(ctx context.Context, users *Users, password string) error
	FindUserByUsername(ctx context.Context, username string) error
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
