package usecase

import (
	tr "golang-standards-project-layout/internal/app/token/repository"
	ur "golang-standards-project-layout/internal/app/user/repository"
	"golang-standards-project-layout/internal/integration/spotify_api"
	"golang-standards-project-layout/internal/integration/telegram_bot"
)

type Module struct {
	tokenRepository    tr.TokenRepository
	userRepository     ur.UserRepository
	spotifyAuthApiCall spotify_api.SpotifyApiCallIntegration
	chatbotManager     telegram_bot.TelegramBotIntegration
}

type Opts struct {
	TokenRepository    tr.TokenRepository
	UserRepository     ur.UserRepository
	SpotifyAuthApiCall spotify_api.SpotifyApiCallIntegration
	ChatbotManager     telegram_bot.TelegramBotIntegration
}

func New(o Opts) *Module {
	return &Module{
		tokenRepository:    o.TokenRepository,
		userRepository:     o.UserRepository,
		spotifyAuthApiCall: o.SpotifyAuthApiCall,
		chatbotManager:     o.ChatbotManager,
	}
}
