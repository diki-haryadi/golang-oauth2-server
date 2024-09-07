package handler

import (
	"context"
	"github.com/rs/zerolog/log"
	cbm "golang-standards-project-layout/internal/model/chatbot"
)

func (x *ChatbotModule) HandleInitConversation(ci cbm.ChatInfo) {
	//ctx, span := tracer.StartSpan(context.Background(), "auth.chatbot.HandleLinkageCallback", nil)
	//defer span.End()

	err := x.choreUsecase.ProcessInitConversation(context.Background(), ci)

	if err != nil {
		log.Err(err).Msg(err.Error())
	}
}

func (x *ChatbotModule) HandleWelcome(ci cbm.ChatInfo) {
	//ctx, span := tracer.StartSpan(context.Background(), "auth.chatbot.HandleLinkageCallback", nil)
	//defer span.End()

	err := x.choreUsecase.ProcessWelcome(context.Background(), ci)

	if err != nil {
		log.Err(err).Msg(err.Error())
	}
}
