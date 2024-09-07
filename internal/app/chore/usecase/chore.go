package usecase

import (
	"context"

	cbm "golang-standards-project-layout/internal/model/chatbot"
)

type ChoreUsecase interface {
	ProcessInitConversation(ctx context.Context, ci cbm.ChatInfo) error
	ProcessWelcome(ctx context.Context, ci cbm.ChatInfo) error
}
