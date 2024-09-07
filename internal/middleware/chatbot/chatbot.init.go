package chatbot

import (
	sr "golang-standards-project-layout/internal/app/session/repository"
	tr "golang-standards-project-layout/internal/app/token/repository"
	ur "golang-standards-project-layout/internal/app/user/repository"
	"golang-standards-project-layout/internal/integration/telegram_bot"
)

type ChatbotMiddlewareModule struct {
	sessionRepository sr.SessionRepository
	userRepository    ur.UserRepository
	tokenRepository   tr.TokenRepository
	chatbotManager    telegram_bot.TelegramBotIntegration
}

type ChatbotMiddlewareOpts struct {
	TokenRepository   tr.TokenRepository
	UserRepository    ur.UserRepository
	SessionRepository sr.SessionRepository
	ChatbotManager    telegram_bot.TelegramBotIntegration
}

func NewChatbotMiddleware(o ChatbotMiddlewareOpts) *ChatbotMiddlewareModule {
	return &ChatbotMiddlewareModule{
		sessionRepository: o.SessionRepository,
		tokenRepository:   o.TokenRepository,
		userRepository:    o.UserRepository,
		chatbotManager:    o.ChatbotManager,
	}
}
