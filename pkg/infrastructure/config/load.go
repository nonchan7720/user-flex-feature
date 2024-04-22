package config

import (
	"bytes"
	"io"
	"os"

	"github.com/creasty/defaults"
	"github.com/goccy/go-yaml"
	"github.com/nonchan7720/user-flex-feature/pkg/infrastructure/config/validator"
)

type SetupLogger interface {
	SetupLog()
}

type ConfigLoader interface {
	IsProduction() bool
	IsStaging() bool
	IsLocal() bool
	OTEL() Tracking
}

func NewLoadYaml(r io.Reader, v any) error {
	buf, err := io.ReadAll(r)
	if err != nil {
		return err
	}
	value := os.ExpandEnv(string(buf))
	if err := yaml.NewDecoder(bytes.NewBufferString(value)).Decode(v); err != nil {
		return err
	}
	if err := defaults.Set(v); err != nil {
		return err
	}
	if err := validator.Validate(v); err != nil {
		return err
	}
	if v, ok := v.(*Config); ok {
		setConfig(v)
	}
	if v, ok := v.(SetupLogger); ok {
		v.SetupLog()
	}
	return nil
}

func NewLoadYamlWithFile(filename string, v any) error {
	f, err := os.Open(filename)
	if err != nil {
		return err
	}
	defer f.Close()
	return NewLoadYaml(f, v)
}
