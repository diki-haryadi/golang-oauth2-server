package handler

import (
	"golang-standards-project-layout/internal/app/auth/usecase"
)

type RestModule struct {
	authUsecase usecase.AuthUsecase
}

type RestOpts struct {
	AuthUsecase usecase.AuthUsecase
}

func NewRest(o RestOpts) *RestModule {
	return &RestModule{
		authUsecase: o.AuthUsecase,
	}
}

type ChatbotModule struct {
	authUsecase usecase.AuthUsecase
}

type ChatbotOpts struct {
	AuthUsecase usecase.AuthUsecase
}

func NewChatbot(o ChatbotOpts) *ChatbotModule {
	return &ChatbotModule{
		authUsecase: o.AuthUsecase,
	}
}
