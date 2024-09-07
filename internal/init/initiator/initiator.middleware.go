package initiator

import (
	"golang-standards-project-layout/configs"
	"golang-standards-project-layout/internal/init/service"
	"golang-standards-project-layout/internal/middleware/chatbot"
)

func (i *Initiator) InitChatbotMiddleware(cfg *config.MainConfig, infra *service.Infrastructure, integration *service.Integration, repos *service.Repositories) *service.Middlewares {
	cbmw := chatbot.NewChatbotMiddleware(chatbot.ChatbotMiddlewareOpts{
		TokenRepository:   repos.TokenRepository,
		UserRepository:    repos.UserRepository,
		SessionRepository: repos.SessionRepository,
		ChatbotManager:    integration.TelegramBotManager,
	})

	return &service.Middlewares{
		ChatbotMiddleware: cbmw,
	}
}
