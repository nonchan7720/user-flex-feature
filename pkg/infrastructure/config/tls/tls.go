package config

import (
	validation "github.com/go-ozzo/ozzo-validation/v4"
	"github.com/nonchan7720/user-flex-feature/pkg/infrastructure/config/validator"
)

type TLS struct {
	InsecureSkipVerify bool   `yaml:"insecureSkipVerify"`
	CaFile             string `yaml:"caFile"`
	ServerName         string `yaml:"serverName"`
}

func (c *TLS) Validate() error {
	return validator.ValidateStruct(c,
		validation.Field(&c.CaFile, validation.When(c.InsecureSkipVerify, validation.Required)),
		validation.Field(&c.ServerName, validation.When(c.InsecureSkipVerify, validation.Required)),
	)
}
