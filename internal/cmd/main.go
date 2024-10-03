package main

import (
	"golang-oauth2-server/internal/app"
	"golang-oauth2-server/internal/pkg/logger"
)

func main() {
	err := app.New().Run()
	if err != nil {
		logger.Zap.Sugar().Fatal(err)
	}
}
