package initiator

import (
	"golang-standards-project-layout/configs"
	"golang-standards-project-layout/internal/init/service"
	//"golang-standards-project-layout/internal/integration/spotify_api"
	//"golang-standards-project-layout/internal/integration/telegram_bot"
)

func (i *Initiator) InitIntegration(cfg *config.MainConfig, infra *service.Infrastructure) (resp *service.Integration) {
	//sa := spotify_api.New(spotify_api.Opts{
	//	SpotifyConfig: &cfg.Spotify,
	//})
	//
	//tgManager := telegram_bot.New(telegram_bot.Opts{
	//	TelegramBotApi: infra.ChatbotModule,
	//})
	//
	//return &service.Integration{
	//	SpotifyApiCall:     sa,
	//	TelegramBotManager: tgManager,
	//}
	return resp
}
