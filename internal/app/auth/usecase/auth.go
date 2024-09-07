package usecase

import (
	"context"

	"golang-standards-project-layout/internal/app/auth/model"
	tm "golang-standards-project-layout/internal/app/token/model"
	um "golang-standards-project-layout/internal/app/user/model"
	cbm "golang-standards-project-layout/internal/model/chatbot"
)

type AuthUsecase interface {
	ProcessHostAuthentication(ctx context.Context, ci cbm.ChatInfo) error
	ProcessLinkageCallback(ctx context.Context, data *model.LinkageCallback) (*um.UserNoSqlSchema, *tm.TokenNoSqlSchema, error)
}
