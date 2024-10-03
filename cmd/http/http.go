package http

import (
	"github.com/spf13/cobra"
	"golang-oauth2-server/internal/app"
	"golang-oauth2-server/pkg/logger"
)

var (
	serveHTTPCmd = &cobra.Command{
		Use:              "serve-http",
		Short:            "GO REST HTTP API",
		PersistentPreRun: rootPreRun,
		RunE:             runHTTP,
	}
)

func rootPreRun(cmd *cobra.Command, args []string) {
	//logger.InitGlobalLogger(&logger.Config{
	//	ServiceName: "golang-oauth2-server",
	//	Level:       zerolog.DebugLevel,
	//})
}

func ServeHTTPCmd() *cobra.Command {
	return serveHTTPCmd
}

func runHTTP(cmd *cobra.Command, args []string) error {
	err := app.New().Run()
	if err != nil {
		logger.Zap.Sugar().Fatal(err)
	}
	return nil
}
