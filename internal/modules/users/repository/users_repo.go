package usersRepository

import (
	"context"
	"fmt"

	usersDomain "golang-oauth2-server/internal/modules/users/domain"
	usersDto "golang-oauth2-server/internal/modules/users/dto"
	"golang-oauth2-server/internal/pkg/postgres"
)

type repository struct {
	postgres *postgres.Postgres
}

func NewRepository(conn *postgres.Postgres) usersDomain.Repository {
	return &repository{postgres: conn}
}

func (rp *repository) CreateUsers(
	ctx context.Context,
	entity *usersDto.CreateUsersRequestDto,
) (*usersDto.CreateUsersResponseDto, error) {
	query := `INSERT INTO articles (username, email, password) VALUES ($1, $2) RETURNING id, username, email, password`

	result, err := rp.postgres.SqlxDB.QueryContext(ctx, query, entity.Username, entity.Email, entity.Password)
	if err != nil {
		return nil, fmt.Errorf("error inserting user record")
	}

	user := new(usersDto.CreateUsersResponseDto)
	for result.Next() {
		err = result.Scan(&user.ID, &user.Username, &user.Email, &user.Password)
		if err != nil {
			return nil, err
		}
	}

	return user, nil
}
