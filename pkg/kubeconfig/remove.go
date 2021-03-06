package kubeconfig

import (
	"github.com/rs/zerolog/log"
)

// Remove takes a context and its associated resources out of kubeconfig
func Remove(name string) {
	mainConfig, err := Read(MainConfigPath)
	if err != nil {
		log.Fatal().Msgf("Could not read main config file: %v", err)
		return
	}

	context, ok := mainConfig.Contexts[name]
	if !ok {
		log.Warn().Msgf("No context named '%s' found in kubeconfig file", name)
		return
	}

	log.Info().Msgf("Removing %s user", context.AuthInfo)
	delete(mainConfig.AuthInfos, context.AuthInfo)
	log.Info().Msgf("Removing %s cluster", context.Cluster)
	delete(mainConfig.Clusters, context.Cluster)
	log.Info().Msgf("Removing %s context", name)
	delete(mainConfig.Contexts, name)

	err = Write(mainConfig)
	if err != nil {
		log.Fatal().Msgf("Error while writing config: %v", err)
	}

	return
}
