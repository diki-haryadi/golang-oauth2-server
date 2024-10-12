package articleRepository

import (
	"context"
	"github.com/google/uuid"
	commonDomain "golang-oauth2-server/internal/modules/common/domain"
	util2 "golang-oauth2-server/pkg/util"
	"strings"
	"time"

	clientDomain "golang-oauth2-server/internal/modules/client/domain"
	"golang-oauth2-server/pkg/postgres"
)

type repository struct {
	postgres *postgres.Postgres
}

func NewRepository(conn *postgres.Postgres) clientDomain.Repository {
	return &repository{postgres: conn}
}

func (rp *repository) FindClientByClientID(ctx context.Context, clientID string) (*clientDomain.Client, error) {
	client := new(clientDomain.Client)
	query := "SELECT * FROM clients WHERE key = ?"
	err := rp.postgres.SqlxDB.Select(&client, query, strings.ToLower(clientID))
	if err != nil {
		return nil, err
	}

	return client, err
}

func (rp *repository) CreateClientCommon(ctx context.Context, clientID, secret, redirectURI string) (*clientDomain.Client, error) {
	_, err := rp.FindClientByClientID(ctx, clientID)
	if err == nil {
		return nil, err
	}

	secretHash, err := util2.HashPassword(secret)
	if err != nil {
		return nil, err
	}

	client := clientDomain.Client{
		CommonModel: commonDomain.CommonModel{
			ID:        uuid.New(),
			CreatedAt: time.Now().UTC(),
		},
		Key:         strings.ToLower(clientID),
		Secret:      string(secretHash),
		RedirectURI: util2.StringOrNull(redirectURI),
	}
	query := "INSERT INTO clients (`key`, secret, redirect_uri, created_at) VALUES ($1,$2,$3,$4) RETURNING id, `key`, secret, redirect_uri, created_at"
	result, err := rp.postgres.SqlxDB.QueryContext(ctx, query, client.Key, client.Secret, client.RedirectURI, client.CreatedAt)
	if err != nil {
		return nil, err
	}

	newClient := new(clientDomain.Client)
	for result.Next() {
		err = result.Scan(&newClient.CommonModel.ID, &newClient.CreatedAt, &newClient.Key, &newClient.Secret, &newClient.RedirectURI)
		if err != nil {
			return nil, err
		}
	}

	return &client, err
}
