package scopeConfigurator

import (
	"context"

	scopeDomain "golang-oauth2-server/internal/modules/scope/domain"
	scopeRepository "golang-oauth2-server/internal/modules/scope/repository"
	scopeUseCase "golang-oauth2-server/internal/modules/scope/usecase"
	infraContainer "golang-oauth2-server/pkg/infra_container"
)

type configurator struct {
	ic *infraContainer.IContainer
}

func NewConfigurator(ic *infraContainer.IContainer) scopeDomain.Configurator {
	return &configurator{ic: ic}
}

func (c *configurator) Configure(ctx context.Context) error {
	repository := scopeRepository.NewRepository(c.ic.Postgres)
	_ = scopeUseCase.NewUseCase(repository)
	return nil
}
