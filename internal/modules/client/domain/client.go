package clientDomain

import (
	"context"
	"database/sql"
	"golang-oauth2-server/internal/modules/common/domain"
)

// OauthClient ...
type Client struct {
	commonDomain.CommonModel
	Key         string         `db:"key" json:"key"`
	Secret      string         `db:"secret" json:"secret"`
	RedirectURI sql.NullString `db:"redirect_uri" json:"redirect_uri"`
}

type Configurator interface {
	Configure(ctx context.Context) error
}

type UseCase interface {
	AuthClient(ctx context.Context, clientID, secret string) (*Client, error)
	CreateClient(ctx context.Context, clientID, secret, redirectURI string) (*Client, error)
	ClientExists(ctx context.Context, clientID string) bool
}

type Repository interface {
	FindClientByClientID(ctx context.Context, clientID string) (*Client, error)
	CreateClientCommon(ctx context.Context, clientID, secret, redirectURI string) (*Client, error)
}

type GrpcController interface{}

type HttpController interface{}

type Job interface{}

type KafkaProducer interface{}

type KafkaConsumer interface{}
