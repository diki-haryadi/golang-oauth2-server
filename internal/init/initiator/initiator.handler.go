package initiator

import (
	"golang-standards-project-layout/configs"
	ah "golang-standards-project-layout/internal/app/auth/handler"
	chh "golang-standards-project-layout/internal/app/chore/handler"
	puc "golang-standards-project-layout/internal/app/player/handler"
	sh "golang-standards-project-layout/internal/app/session/handler"
	"golang-standards-project-layout/internal/init/service"
)

func (i *Initiator) InitRestHandler(cfg *config.MainConfig, infra *service.Infrastructure, uc *service.Usecases) *service.RestHandlers {
	auth := ah.NewRest(ah.RestOpts{AuthUsecase: uc.AuthUsecase})

	return &service.RestHandlers{
		AuthRestHandler: auth,
	}
}

func (i *Initiator) InitChatbotHandler(cfg *config.MainConfig, infra *service.Infrastructure, integration *service.Integration, uc *service.Usecases) *service.ChatbotHandlers {
	auth := ah.NewChatbot(ah.ChatbotOpts{AuthUsecase: uc.AuthUsecase})
	chore := chh.NewChatbot(chh.ChatbotOpts{ChoreUsecase: uc.ChoreUsecase})
	session := sh.NewChatbot(sh.ChatbotOpts{SessionUsecase: uc.SessionUsecase})
	player := puc.NewChatbot(puc.ChatbotOpts{PlayerUsecase: uc.PlayerUsecase})

	return &service.ChatbotHandlers{
		AuthChatbotHandler:    auth,
		ChoreChatbotHandler:   chore,
		SessionChatbotHandler: session,
		PlayerChatbotHandler:  player,
	}
}
