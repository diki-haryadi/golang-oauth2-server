package handler

import (
	"context"

	"github.com/rs/zerolog/log"
	cbm "golang-standards-project-layout/internal/model/chatbot"
)

func (x *ChatbotModule) HandleCreateSession(ci cbm.ChatInfo) {
	//ctx, span := tracer.StartSpan(context.Background(), "session.chatbot.HandleCreateSession", nil)
	//defer span.End()

	err := x.sessionUsecase.ProcessCreateNewSession(context.Background(), ci)

	if err != nil {
		log.Err(err).Msg(err.Error())
	}

}
