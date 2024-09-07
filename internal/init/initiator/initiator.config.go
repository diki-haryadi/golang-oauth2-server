package initiator

import (
	"github.com/google/gops/agent"
	"github.com/rs/zerolog/log"
	"golang-standards-project-layout/configs"
	configpkg "golang-standards-project-layout/pkg/config"
)

var (
	errInitConfig = "failed to initiate config"
)

// Main Config
func (i *Initiator) InitConfig(configPath string) *config.MainConfig {
	if err := i.agentListen(agent.Options{
		ShutdownCleanup: true, // automatically closes on os.Interrupt
	}); err != nil {
		log.Fatal().Err(err).Msg(errInitConfig)
	}

	cfg := &config.MainConfig{}
	log.Info().Msgf("reading config from %s", configPath)
	err := configpkg.ReadModuleConfig(cfg, configPath, "config")
	if err != nil {
		log.Fatal().Err(err).Msg(errInitConfig)
	}

	return cfg

}
