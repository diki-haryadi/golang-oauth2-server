package telegram_bot

import (
	"context"
	"fmt"

	"github.com/rs/zerolog/log"
	"golang-standards-project-layout/internal/integration/telegram_bot/constants"
)

func (x *Module) SendSessionCreatedMessage(ctx context.Context, chat_id string, session_code string) error {
	//ctx, span := tracer.StartSpan(ctx, "chatbot.telegram.SendProcessInitConversationMessage", nil)
	//defer span.End()

	var (
		scm string
		err error
	)

	scm = fmt.Sprintf("Your session ID is: %s", session_code)
	err = x.sendTextMessage(ctx, scm, chat_id)

	if err != nil {
		log.Err(err).Msg(err.Error())
	}

	return nil
}

func (x *Module) SendNoActiveSessionMessage(ctx context.Context, chat_id string) error {
	//ctx, span := tracer.StartSpan(ctx, "chatbot.telegram.SendProcessInitConversationMessage", nil)
	//defer span.End()

	err := x.sendInlineTextMessage(ctx, constants.NoActiveSessionMessage, constants.NoActiveSessionRows, chat_id)

	if err != nil {
		log.Err(err).Msg(err.Error())
	}

	return nil
}
