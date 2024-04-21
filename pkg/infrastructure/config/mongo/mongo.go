package config

import . "github.com/nonchan7720/user-flex-feature/pkg/infrastructure/config/tls"

type Mongo struct {
	URI          string `yaml:"uri"`
	DatabaseName string `yaml:"database"`
	TLS          *TLS
}
