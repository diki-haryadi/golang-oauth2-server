package model

import "golang-standards-project-layout/internal/constants"

type ChatInfo struct {
	Channel        constants.ChatbotChannelEnum
	SenderFullName string
	SenderId       string
	ChatId         string
	MessageId      string
	Message        string
}
