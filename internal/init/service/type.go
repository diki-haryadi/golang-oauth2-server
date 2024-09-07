package service

import (
	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
	"golang-standards-project-layout/configs"
	"golang-standards-project-layout/pkg/database"

	ah "golang-standards-project-layout/internal/app/auth/handler"
	auc "golang-standards-project-layout/internal/app/auth/usecase"
	"golang-standards-project-layout/internal/middleware/chatbot"

	tr "golang-standards-project-layout/internal/app/token/repository"
	ur "golang-standards-project-layout/internal/app/user/repository"

	chh "golang-standards-project-layout/internal/app/chore/handler"
	chuc "golang-standards-project-layout/internal/app/chore/usecase"

	sh "golang-standards-project-layout/internal/app/session/handler"
	sr "golang-standards-project-layout/internal/app/session/repository"
	suc "golang-standards-project-layout/internal/app/session/usecase"

	ph "golang-standards-project-layout/internal/app/player/handler"
	puc "golang-standards-project-layout/internal/app/player/usecase"

	"golang-standards-project-layout/internal/integration/spotify_api"
	"golang-standards-project-layout/internal/integration/telegram_bot"
)

type Infrastructure struct {
	ChatbotModule *tgbotapi.BotAPI
	Mongo         *database.MongoDatabase
}

type Repositories struct {
	TokenRepository   tr.TokenRepository
	UserRepository    ur.UserRepository
	SessionRepository sr.SessionRepository
}

type Usecases struct {
	AuthUsecase    auc.AuthUsecase
	ChoreUsecase   chuc.ChoreUsecase
	SessionUsecase suc.SessionUsecase
	PlayerUsecase  puc.PlayerUsecase
}

type Middlewares struct {
	ChatbotMiddleware chatbot.ChatbotMiddleware
}

type RestHandlers struct {
	AuthRestHandler ah.AuthRestHandler
}

type ChatbotHandlers struct {
	AuthChatbotHandler    ah.AuthChatbotHandler
	ChoreChatbotHandler   chh.ChoreChatbotHandler
	SessionChatbotHandler sh.SessionChatbotHandler
	PlayerChatbotHandler  ph.PlayerChatbotHandler
}

type RestService struct {
	Config         *config.MainConfig
	Infrastructure *Infrastructure
	RestHandlers   *RestHandlers
	Usecases       *Usecases
	Middlewares    *Middlewares
}

type ChatbotListenerService struct {
	Middlewares    *Middlewares
	Handlers       *ChatbotHandlers
	Usecases       *Usecases
	Infrastructure *Infrastructure
	Config         *config.MainConfig
}

type Integration struct {
	SpotifyApiCall     spotify_api.SpotifyApiCallIntegration
	TelegramBotManager telegram_bot.TelegramBotIntegration
}
