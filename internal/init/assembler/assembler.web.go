package assembler

import (
	"time"

	"golang-standards-project-layout/api/rest"
	"golang-standards-project-layout/internal/init/service"
)

func (a *assembler) assembleWeb(service *service.RestService) rest.WebHandler {
	server := a.NewWebHandler(&rest.Opts{
		Service:       service,
		Port:          service.Config.Rest.Port,
		ListenAddress: service.Config.Rest.ListenAddress,
		ReadTimeout:   time.Duration(service.Config.Rest.ReadTimeout) * time.Second,
		WriteTimeout:  time.Duration(service.Config.Rest.WriteTimeout) * time.Second,
	})
	return server
}
func (a *assembler) runWebServer() {
	go a.webHandler.Run()
}
