package config

import (
	"time"

	validation "github.com/go-ozzo/ozzo-validation/v4"
	"github.com/go-ozzo/ozzo-validation/v4/is"
	. "github.com/nonchan7720/user-flex-feature/pkg/infrastructure/config/tls"
	"github.com/nonchan7720/user-flex-feature/pkg/infrastructure/config/validator"
)

type Redis struct {
	ConnectionURL string        `yaml:"connectionUrl"`
	Username      string        `yaml:"username"`
	Password      string        `yaml:"password"`
	IdleTimeout   time.Duration `yaml:"idleTimeout"`
	TLS           *TLS          `yaml:"tls"`
}

func (c *Redis) Validate() error {
	return validator.ValidateStruct(c,
		validation.Field(&c.ConnectionURL, validation.Required, is.URL),
	)
}
