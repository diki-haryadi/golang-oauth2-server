package initiator

import (
	"time"

	"golang-standards-project-layout/configs"
	"golang-standards-project-layout/internal/init/service"
	mongo "golang-standards-project-layout/pkg/database"
)

func (i *Initiator) InitInfrastructure(cfg *config.MainConfig) *service.Infrastructure {
	mongo := mongo.MongoConnectClient(&mongo.Client{
		URI:            cfg.Mongo.URI,
		DB:             cfg.Mongo.DB,
		AppName:        cfg.Mongo.DB,
		ConnectTimeout: time.Duration(cfg.Mongo.ConnectionTimeout) * time.Second,
		PingTimeout:    time.Duration(cfg.Mongo.PingTimeout) * time.Second,
	})

	//chatbot := initChatbotApi(cfg)

	return &service.Infrastructure{
		Mongo: mongo,
		//ChatbotModule: chatbot,
	}
}

//func initChatbotApi(cfg *config.MainConfig) *tgbotapi.BotAPI {
//	bot, err := tgbotapi.NewBotAPI(cfg.TelegramBotApi.Token)
//
//	if err != nil {
//		log.Panic().Msg(err.Error())
//	}
//
//	return bot
//}
