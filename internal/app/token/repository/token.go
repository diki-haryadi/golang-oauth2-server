package repository

import (
	"context"

	"go.mongodb.org/mongo-driver/bson/primitive"
	"golang-standards-project-layout/internal/app/token/model"
)

type TokenRepository interface {
	FindAndValidateTokenByUserId(ctx context.Context, id primitive.ObjectID) (*model.TokenNoSqlSchema, error)
	StoreToken(ctx context.Context, token *model.TokenNoSqlSchema) error
}
