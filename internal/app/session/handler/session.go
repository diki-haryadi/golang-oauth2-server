package handler

import (
	cbm "golang-standards-project-layout/internal/model/chatbot"
)

type SessionChatbotHandler interface {
	HandleCreateSession(ci cbm.ChatInfo)
}
