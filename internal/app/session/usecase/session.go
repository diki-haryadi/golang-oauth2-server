package usecase

import (
	"context"

	cbm "golang-standards-project-layout/internal/model/chatbot"
)

type SessionUsecase interface {
	ProcessCreateNewSession(ctx context.Context, ci cbm.ChatInfo) error
}
