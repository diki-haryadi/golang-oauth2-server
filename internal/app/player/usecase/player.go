package usecase

import (
	"context"

	cbm "golang-standards-project-layout/internal/model/chatbot"
)

type PlayerUsecase interface {
	ProcessQueueTrack(context.Context, cbm.ChatInfo) error
}
