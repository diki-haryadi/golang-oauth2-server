package articleRepository

import (
	"context"
	"fmt"
	"github.com/lib/pq"
	"strings"

	scopeDomain "golang-oauth2-server/internal/modules/scope/domain"
	"golang-oauth2-server/pkg/postgres"
)

type repository struct {
	postgres *postgres.Postgres
}

func NewRepository(conn *postgres.Postgres) scopeDomain.Repository {
	return &repository{postgres: conn}
}

func (rp *repository) GetDefaultScope(ctx context.Context) (string, error) {
	var scopes []string
	query := "SELECT * FROM scopes WHERE is_default = $1"
	err := rp.postgres.SqlxDB.Select(&scopes, query, true)
	if err != nil {
		return "", fmt.Errorf("error get scope")
	}
	return strings.Join(scopes, " "), nil
}

func (rp *repository) ScopeExists(ctx context.Context, requestScope string) (bool, error) {
	scopes := strings.Split(requestScope, " ")

	query := "SELECT COUNT(*) FROM scopes WHERE scope = ANY($1)"
	scopeArray := pq.Array(scopes)

	var count int
	err := rp.postgres.SqlxDB.Get(&count, query, scopeArray)
	if err != nil {
		return false, fmt.Errorf("error get scope")
	}
	return count == len(scopes), nil
}
