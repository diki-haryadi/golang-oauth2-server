package handler

import (
	"golang-standards-project-layout/internal/app/session/usecase"
)

type ChatbotModule struct {
	sessionUsecase usecase.SessionUsecase
}

type ChatbotOpts struct {
	SessionUsecase usecase.SessionUsecase
}

func NewChatbot(o ChatbotOpts) *ChatbotModule {
	return &ChatbotModule{
		sessionUsecase: o.SessionUsecase,
	}
}
