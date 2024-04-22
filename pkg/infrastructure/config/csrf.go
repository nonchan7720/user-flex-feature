package config

import (
	validation "github.com/go-ozzo/ozzo-validation/v4"
)

type CSRFToken struct {
	Domain string `yaml:"domain"`
	Secure bool   `yaml:"secure" default:"true"`
}

func (s CSRFToken) Validate() error {
	return validation.ValidateStruct(&s)
}
