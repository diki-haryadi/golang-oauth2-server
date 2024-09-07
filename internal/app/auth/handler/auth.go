package handler

import (
	"github.com/gofiber/fiber/v2"
	cbm "golang-standards-project-layout/internal/model/chatbot"
)

type AuthRestHandler interface {
	HandleLinkageCallback(fc *fiber.Ctx) error
}

type AuthChatbotHandler interface {
	HandleHostAuthentication(ci cbm.ChatInfo)
}
