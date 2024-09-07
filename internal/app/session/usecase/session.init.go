package usecase

import (
	sr "golang-standards-project-layout/internal/app/session/repository"
	ur "golang-standards-project-layout/internal/app/user/repository"
	"golang-standards-project-layout/internal/integration/spotify_api"
	"golang-standards-project-layout/internal/integration/telegram_bot"
)

type Module struct {
	sessionRepository  sr.SessionRepository
	userRepository     ur.UserRepository
	spotifyAuthApiCall spotify_api.SpotifyApiCallIntegration
	chatbotManager     telegram_bot.TelegramBotIntegration
}

type Opts struct {
	UserRepository     ur.UserRepository
	SessionRepository  sr.SessionRepository
	SpotifyAuthApiCall spotify_api.SpotifyApiCallIntegration
	ChatbotManager     telegram_bot.TelegramBotIntegration
}

func New(o Opts) *Module {
	return &Module{
		sessionRepository:  o.SessionRepository,
		userRepository:     o.UserRepository,
		spotifyAuthApiCall: o.SpotifyAuthApiCall,
		chatbotManager:     o.ChatbotManager,
	}
}
