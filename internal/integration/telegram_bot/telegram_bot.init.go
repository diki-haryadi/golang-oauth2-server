package telegram_bot

import (
	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
)

type Opts struct {
	TelegramBotApi *tgbotapi.BotAPI
}

type Module struct {
	telegramBotApi *tgbotapi.BotAPI
}

func New(o Opts) *Module {
	return &Module{
		telegramBotApi: o.TelegramBotApi,
	}
}
