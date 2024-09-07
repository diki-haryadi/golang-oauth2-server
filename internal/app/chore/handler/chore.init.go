package handler

import (
	"golang-standards-project-layout/internal/app/chore/usecase"
)

type ChatbotModule struct {
	choreUsecase usecase.ChoreUsecase
}

type ChatbotOpts struct {
	ChoreUsecase usecase.ChoreUsecase
}

func NewChatbot(o ChatbotOpts) *ChatbotModule {
	return &ChatbotModule{
		choreUsecase: o.ChoreUsecase,
	}
}
