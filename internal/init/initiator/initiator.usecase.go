package initiator

import (
	"golang-standards-project-layout/configs"
	auc "golang-standards-project-layout/internal/app/auth/usecase"
	cuc "golang-standards-project-layout/internal/app/chore/usecase"
	suc "golang-standards-project-layout/internal/app/session/usecase"
	"golang-standards-project-layout/internal/init/service"
)

func (i *Initiator) InitUsecase(cfg *config.MainConfig, infra *service.Infrastructure, repos *service.Repositories, integration *service.Integration) *service.Usecases {
	auth := auc.New(auc.Opts{
		TokenRepository:    repos.TokenRepository,
		UserRepository:     repos.UserRepository,
		SpotifyAuthApiCall: integration.SpotifyApiCall,
		ChatbotManager:     integration.TelegramBotManager,
	})

	chore := cuc.New(cuc.Opts{
		UserRepository: repos.UserRepository,
		ChatbotManager: integration.TelegramBotManager,
	})

	session := suc.New(suc.Opts{
		UserRepository:     repos.UserRepository,
		SessionRepository:  repos.SessionRepository,
		ChatbotManager:     integration.TelegramBotManager,
		SpotifyAuthApiCall: integration.SpotifyApiCall,
	})

	return &service.Usecases{
		AuthUsecase:    auth,
		ChoreUsecase:   chore,
		SessionUsecase: session,
	}
}
