package cmd

import (
	"log"
	"os"

	"golang-standards-project-layout/cmd/http"
	//"golang-standards-project-layout/cmd/telegram_listener"

	"github.com/spf13/cobra"
)

var (
	rootCmd = &cobra.Command{
		Use:   "go",
		Short: "Go - Backend Service",
		Long:  "Go - Backend Service",
	}
)

func Execute() {
	rootCmd.AddCommand(http.ServeHTTPCmd())
	http.ServeHTTPCmd().Flags().StringP("config", "c", "config/file", "Config URL dir i.e. config/file")

	if err := rootCmd.Execute(); err != nil {
		log.Fatalln("Error: \n", err.Error())
		os.Exit(-1)
	}
}
