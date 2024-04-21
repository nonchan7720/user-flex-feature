package config

import "time"

type InMemory struct {
	CleanUp  time.Duration `yaml:"cleanUp" default:"1m"`
	FilePath string        `yaml:"filePath"`
}
