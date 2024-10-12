package scopeDomain

import (
	"context"
	"golang-oauth2-server/internal/modules/common/domain"
)

type Scope struct {
	commonDomain.CommonModel
	Scope       string `db:"scope" json:"scope"`
	Description string `db:"description" json:"desc"`
	IsDefault   bool   `db:"is_default"`
}

type Configurator interface {
	Configure(ctx context.Context) error
}

type UseCase interface {
	GetScope(ctx context.Context, requestScope string) (string, error)
}

type Repository interface {
	GetDefaultScope(ctx context.Context) (string, error)
	ScopeExists(ctx context.Context, requestedScope string) (bool, error)
}

type GrpcController interface{}

type HttpController interface{}

type Job interface{}

type KafkaProducer interface{}

type KafkaConsumer interface{}
