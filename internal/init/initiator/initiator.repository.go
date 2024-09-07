package initiator

import (
	"golang-standards-project-layout/configs"
	sr "golang-standards-project-layout/internal/app/session/repository"
	tr "golang-standards-project-layout/internal/app/token/repository"
	ur "golang-standards-project-layout/internal/app/user/repository"
	"golang-standards-project-layout/internal/init/service"
)

func (i *Initiator) InitRepository(cfg *config.MainConfig, integration *service.Integration, infra *service.Infrastructure) *service.Repositories {
	token := tr.New(tr.Opts{
		SpotifyAuthApiCall: integration.SpotifyApiCall,
		MongoManager:       infra.Mongo.MongoDatabase,
	})
	user := ur.New(ur.Opts{
		MongoManager: infra.Mongo.MongoDatabase,
	})

	session := sr.New(sr.Opts{
		MongoManager: infra.Mongo.MongoDatabase,
	})

	return &service.Repositories{
		TokenRepository:   token,
		UserRepository:    user,
		SessionRepository: session,
	}
}
