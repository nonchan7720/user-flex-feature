package cmd

import (
	"github.com/nonchan7720/user-flex-feature/pkg/infrastructure/config"
)

type Config interface {
	config.Config
}

func loadConfig[T config.ConfigLoader](configFilePath string) *T {
	var cfg T

	if err := config.NewLoadYamlWithFile(configFilePath, &cfg); err != nil {
		panic(err)
	}
	return &cfg
}
