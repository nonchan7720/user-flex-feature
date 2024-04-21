package config

import (
	"crypto/tls"
	"fmt"
	"log/slog"

	. "github.com/nonchan7720/user-flex-feature/pkg/infrastructure/config/tls"
	"google.golang.org/grpc/credentials"
	"google.golang.org/grpc/credentials/insecure"
)

type Grpc struct {
	Host        string `yaml:"endpoint"`
	Port        int    `yaml:"port"`
	Credentials *TLS   `yaml:"tls"`
	Auth        *Auth  `yaml:"auth"`
}

func (c *Grpc) GrpcCredentials() credentials.TransportCredentials {
	if c.Credentials == nil {
		return insecure.NewCredentials()
	}
	credCfg := c.Credentials
	if credCfg.InsecureSkipVerify {
		return credentials.NewTLS(&tls.Config{InsecureSkipVerify: true})
	} else if credCfg.CaFile != "" && credCfg.ServerName != "" {
		cred, err := credentials.NewClientTLSFromFile(credCfg.CaFile, credCfg.ServerName)
		if err != nil {
			panic(err)
		}
		return cred
	} else {
		slog.Warn("Credentials is declared but nothing is set.")
		return insecure.NewCredentials()
	}
}

func (c *Grpc) Endpoint() string {
	return fmt.Sprintf("%s:%d", c.Host, c.Port)
}
