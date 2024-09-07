package handler

import (
	"context"

	"github.com/rs/zerolog/log"
	gam "golang-standards-project-layout/internal/model/auth"
	cbm "golang-standards-project-layout/internal/model/chatbot"
)

func (x *ChatbotModule) HandleQueueTrack(ci cbm.ChatInfo, cm *gam.CommandMetadata) {
	//ctx, span := tracer.StartSpan(context.Background(), "player.chatbot.HandleQueueTrack", nil)
	//defer span.End()

	err := x.playerUsecase.ProcessQueueTrack(context.Background(), ci)

	if err != nil {
		log.Err(err).Msg(err.Error())
	}

}
