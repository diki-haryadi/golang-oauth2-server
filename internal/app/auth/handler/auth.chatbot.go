package handler

import (
	"context"
	"github.com/rs/zerolog/log"
	cbm "golang-standards-project-layout/internal/model/chatbot"
)

func (x *ChatbotModule) HandleHostAuthentication(ci cbm.ChatInfo) {
	//ctx, span := tracer.StartSpan(context.Background(), "auth.chatbot.HandleHostAuthentication", nil)
	//defer span.End()

	err := x.authUsecase.ProcessHostAuthentication(context.Background(), ci)

	if err != nil {
		log.Err(err).Msg(err.Error())
	}

}
