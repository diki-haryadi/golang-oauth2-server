package http

import (
	"os"

	"github.com/sirupsen/logrus"
	"github.com/spf13/cobra"

	"golang-standards-project-layout/pkg/env"
	"golang-standards-project-layout/pkg/log"

	"golang-standards-project-layout/internal/init/assembler"
)

var (
	serveHTTPCmd = &cobra.Command{
		Use:              "serve-http",
		Short:            "Base Service HTTP",
		Long:             "Serve Base Service through HTTP",
		PersistentPreRun: rootPreRun,
		RunE:             runHTTP,
	}
)

func ServeHTTPCmd() *cobra.Command {
	return serveHTTPCmd
}

func rootPreRun(cmd *cobra.Command, args []string) {
	log.LogInit()

	if env.Get() == env.EnvDevelopment {
		logrus.SetFormatter(&logrus.TextFormatter{ForceColors: true})
		log.SetLevel("debug")
	}
}

func runHTTP(cmd *cobra.Command, args []string) error {
	configURL, _ := cmd.Flags().GetString("config")
	bootstrapHTTP(assembler.NewAssembler(), configURL)

	return nil
}

func bootstrapHTTP(starter assembler.AssemblerManager, configPath string) {
	starter = starter.BuildService("base", configPath).AssembleWebApplication()
	starter.RunWebApplication()

	select {
	case err := <-starter.ListenErrorWebApp():
		starter.StopApplication(os.Interrupt)
		log.Fatalf("Error starting web server, exiting gracefully: %v", err)
	case s := <-starter.TerminateSignal():
		starter.StopApplication(s)
		log.Fatalf("Exiting gracefully...")
	}
}
