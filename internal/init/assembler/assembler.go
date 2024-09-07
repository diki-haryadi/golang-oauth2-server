package assembler

import (
	"os"
	"os/signal"
	"syscall"

	"golang-standards-project-layout/api/rest"
	"golang-standards-project-layout/internal/init/initiator"
	"golang-standards-project-layout/internal/init/service"
)

type NewWebHandler func(o *rest.Opts) rest.WebHandler

//type NewTelegramListenerHandler func(o *telegram_listener.Opts) telegram_listener.TelegramListenerHandler

type assembler struct {
	Initiator     initiator.InitiatorManager
	NewWebHandler NewWebHandler
	webHandler    rest.WebHandler
	restService   *service.RestService
	//NewChatbotListenerHandler NewTelegramListenerHandler
	//chatbotListenerService    *service.ChatbotListenerService
	//chatbotListenerHanlder    telegram_listener.TelegramListenerHandler
	term chan os.Signal
}

type AssemblerManager interface {
	BuildService(configPath string) AssemblerManager

	AssembleWebApplication() AssemblerManager
	RunWebApplication()

	//AssembleTelegramListernerApplication() AssemblerManager
	//RunTelegramListenerApplication()
	ListenErrorWebApp() <-chan error

	TerminateSignal() chan os.Signal
}

func NewAssembler() AssemblerManager {
	return &assembler{
		Initiator:     initiator.New(),
		NewWebHandler: rest.New,
		//NewChatbotListenerHandler: telegram_listener.New,
	}
}

func (a *assembler) BuildService(configPath string) AssemblerManager {
	cfg := a.Initiator.InitConfig(configPath)
	infra := a.Initiator.InitInfrastructure(cfg)
	integration := a.Initiator.InitIntegration(cfg, infra)
	repo := a.Initiator.InitRepository(cfg, integration, infra)
	uc := a.Initiator.InitUsecase(cfg, infra, repo, integration)
	rest := a.Initiator.InitRestHandler(cfg, infra, uc)
	//chatbothandler := a.Initiator.InitChatbotHandler(cfg, infra, integration, uc)
	//cmw := a.Initiator.InitChatbotMiddleware(cfg, infra, integration, repo)

	restsvc := a.Initiator.InitRestService(cfg, infra, rest, uc)
	//chabotsvc := a.Initiator.InitChatbotListenerService(cfg, chatbothandler, infra, uc, cmw)

	a.restService = restsvc
	//a.chatbotListenerService = chabotsvc

	return a
}

func (a *assembler) AssembleWebApplication() AssemblerManager {
	a.webHandler = a.assembleWeb(a.restService)
	return a
}

func (a *assembler) RunWebApplication() {
	a.runWebServer()
}

//
//func (a *assembler) AssembleTelegramListernerApplication() AssemblerManager {
//	a.chatbotListenerHanlder = a.assembleChatbotListener(a.chatbotListenerService)
//	return a
//}

//func (a *assembler) RunTelegramListenerApplication() {
//	a.runTelegramListener()
//}

func (a *assembler) ListenErrorWebApp() <-chan error {
	return a.webHandler.ListenError()
}

func (a *assembler) TerminateSignal() chan os.Signal {
	a.term = make(chan os.Signal)
	signal.Notify(a.term, os.Interrupt, syscall.SIGTERM)
	return a.term
}
