package repository

import (
	"context"

	"go.mongodb.org/mongo-driver/bson/primitive"
	"golang-standards-project-layout/internal/app/user/model"
	"golang-standards-project-layout/internal/constants"
)

type UserRepository interface {
	FindUserByChatbotUserId(ctx context.Context, chatbotUserId string, chatbotChannel constants.ChatbotChannelEnum) (*model.UserNoSqlSchema, error)
	FindUserById(ctx context.Context, id primitive.ObjectID) (*model.UserNoSqlSchema, error)
	UpdateUser(ctx context.Context, id primitive.ObjectID, user *model.UserNoSqlSchema) error
	StoreUser(ctx context.Context, user *model.UserNoSqlSchema) error
}
