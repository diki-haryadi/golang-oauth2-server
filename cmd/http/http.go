package http

import (
	"log"

	"github.com/spf13/cobra"
	"golang-standards-project-layout/internal/init/assembler"
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
	//	ServiceName: "golang-standards-project-layout",
	//	Level:       zerolog.DebugLevel,
	//})
}

func ServeHTTPCmd() *cobra.Command {
	return serveHTTPCmd
}

func runHTTP(cmd *cobra.Command, args []string) error {
	configURL, _ := cmd.Flags().GetString("config")
	bootstrapHTTP(assembler.NewAssembler(), configURL)
	return nil
}

func bootstrapHTTP(starter assembler.AssemblerManager, configPath string) {
	//err := tracer.Init(&tracer.TracerConfig{
	//	UseJaeger:   false,
	//	Environment: env.GetEnvironmentName(),
	//	ServiceName: "golang-standards-project-layout",
	//})

	//if err != nil {
	//	log.Fatalln(err)
	//}

	starter = starter.BuildService(configPath).AssembleWebApplication()
	starter.RunWebApplication()

	select {
	case err := <-starter.ListenErrorWebApp():
		log.Fatalf("Error starting web server, exiting gracefully %v:", err)
	case <-starter.TerminateSignal():
		log.Fatalln("Exiting gracefully...")
	}
}
