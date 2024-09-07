package telegram_bot

import (
	"context"
	"fmt"

	"github.com/rs/zerolog/log"
)

func (x *Module) SendErrorMessage(ctx context.Context, chat_id string, err_msg string) error {
	//ctx, span := tracer.StartSpan(ctx, "chatbot.telegram.SendProcessInitConversationMessage", nil)
	//defer span.End()

	var (
		scm string
		err error
	)

	scm = fmt.Sprintf("Error found: %s", err_msg)
	err = x.sendTextMessage(ctx, scm, chat_id)

	if err != nil {
		log.Err(err).Msg(err.Error())
	}

	return nil
}
