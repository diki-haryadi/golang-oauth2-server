package articleRepository

import (
	"context"
	"fmt"

	articleDomain "golang-oauth2-server/internal/modules/article/domain"
	articleDto "golang-oauth2-server/internal/modules/article/dto"
	"golang-oauth2-server/internal/pkg/postgres"
)

type repository struct {
	postgres *postgres.Postgres
}

func NewRepository(conn *postgres.Postgres) articleDomain.Repository {
	return &repository{postgres: conn}
}

func (rp *repository) CreateUsers(
	ctx context.Context,
	entity *articleDto.CreateArticleRequestDto,
) (*articleDto.CreateArticleResponseDto, error) {
	query := `INSERT INTO articles (name, description) VALUES ($1, $2) RETURNING id, name, description`

	result, err := rp.postgres.SqlxDB.QueryContext(ctx, query, entity.Name, entity.Description)
	if err != nil {
		return nil, fmt.Errorf("error inserting article record")
	}

	article := new(articleDto.CreateArticleResponseDto)
	for result.Next() {
		err = result.Scan(&article.ID, &article.Name, &article.Description)
		if err != nil {
			return nil, err
		}
	}

	return article, nil
}
