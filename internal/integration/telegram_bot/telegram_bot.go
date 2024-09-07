package telegram_bot

import (
	"context"
)

type TelegramBotIntegration interface {
	SendProcessInitConversationMessage(ctx context.Context, chat_id string) error
	SendWelcomeConversationMessage(ctx context.Context, chat_id string) error
	SendInitializeLinkageMessage(ctx context.Context, chat_id string, linkage_url string) error
	SendLinkageSuccessMessage(ctx context.Context, chat_id string) error
	SendHostActionsMessage(ctx context.Context, chat_id string) error
	SendNoActiveSessionMessage(ctx context.Context, chat_id string) error
	SendSessionCreatedMessage(ctx context.Context, chat_id string, session_code string) error
	SendErrorMessage(ctx context.Context, chat_id string, err_msg string) error
}
