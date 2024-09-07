package handler

import (
	gam "golang-standards-project-layout/internal/model/auth"
	cbm "golang-standards-project-layout/internal/model/chatbot"
)

type PlayerChatbotHandler interface {
	HandleQueueTrack(ci cbm.ChatInfo, cm *gam.CommandMetadata)
}
