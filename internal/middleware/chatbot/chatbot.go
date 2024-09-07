package chatbot

import (
	"context"

	gam "golang-standards-project-layout/internal/model/auth"
	cbm "golang-standards-project-layout/internal/model/chatbot"
)

type ChatbotMiddleware interface {
	ParseAndValidateSenderData(ctx context.Context, ci cbm.ChatInfo) (res *gam.CommandMetadata, err error)
}
