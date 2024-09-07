package handler

import (
	cbm "golang-standards-project-layout/internal/model/chatbot"
)

type ChoreChatbotHandler interface {
	HandleInitConversation(ci cbm.ChatInfo)
	HandleWelcome(ci cbm.ChatInfo)
}
