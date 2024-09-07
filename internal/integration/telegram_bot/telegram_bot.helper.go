package telegram_bot

import (
	"context"
	"strconv"

	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
	"github.com/rs/zerolog/log"
	"golang-standards-project-layout/internal/integration/telegram_bot/model"
)

func (x *Module) sendTextMessage(ctx context.Context, message string, chat_id string) error {
	//_, span := tracer.StartSpan(ctx, "chatbot.telegram.SendMessage", nil)
	//defer span.End()
	var (
		err error
		cid int64
	)

	cid, err = strconv.ParseInt(chat_id, 10, 64)
	if err != nil {
		log.Err(err).Msg(err.Error())
		return err
	}

	msg := tgbotapi.NewMessage(cid, message)

	_, err = x.telegramBotApi.Send(msg)

	return err
}

func (x *Module) sendInlineTextMessage(ctx context.Context, message string, rows []model.TelegramInlineKeyboardRow, chat_id string) error {
	//_, span := tracer.StartSpan(ctx, "chatbot.telegram.sendInlineTextMessage", nil)
	//defer span.End()

	var (
		err error
		cid int64
	)

	cid, err = strconv.ParseInt(chat_id, 10, 64)

	if err != nil {
		log.Err(err).Msg(err.Error())
		return err
	}

	rep := x.buildTelegramInlineKeyboard(rows)
	msg := tgbotapi.NewMessage(cid, message)

	msg.ReplyMarkup = rep

	_, err = x.telegramBotApi.Send(msg)

	return err
}

func (x *Module) buildTelegramInlineKeyboard(rows []model.TelegramInlineKeyboardRow) tgbotapi.InlineKeyboardMarkup {
	var tr [][]tgbotapi.InlineKeyboardButton

	for _, row := range rows {
		var kr []tgbotapi.InlineKeyboardButton

		for _, data := range row.Data {
			if data.URL != "" {
				kr = append(kr, tgbotapi.NewInlineKeyboardButtonURL(data.Text, data.URL))
			} else {
				kr = append(kr, tgbotapi.NewInlineKeyboardButtonData(data.Text, data.Value))
			}

		}

		tr = append(tr, tgbotapi.NewInlineKeyboardRow(kr...))
	}

	return tgbotapi.NewInlineKeyboardMarkup(tr...)
}
