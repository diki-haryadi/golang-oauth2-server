package postgresHealthCheckUseCase

import (
	healthCheckDomain "golang-oauth2-server/internal/modules/health_check/domain"
	"golang-oauth2-server/internal/pkg/postgres"
)

type useCase struct {
	postgres *postgres.Postgres
}

func NewUseCase(postgres *postgres.Postgres) healthCheckDomain.PostgresHealthCheckUseCase {
	return &useCase{
		postgres: postgres,
	}
}

func (uc *useCase) Check() bool {
	if err := uc.postgres.SqlxDB.DB.Ping(); err != nil {
		return false
	}
	return true
}
