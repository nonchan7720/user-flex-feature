package config

import (
	"encoding/base64"
	"fmt"
	"net/textproto"
	"strings"

	validation "github.com/go-ozzo/ozzo-validation/v4"
	"github.com/nonchan7720/user-flex-feature/pkg/infrastructure/config/auth"
	"github.com/nonchan7720/user-flex-feature/pkg/infrastructure/config/validator"
	"github.com/nonchan7720/user-flex-feature/pkg/utils"
)

type AuthType string

const (
	BasicAuthType   = "basic"
	AccessTokenType = "accessToken"
)

type Auth struct {
	Type            AuthType         `yaml:"authType"`
	BasicAuth       *BasicAuth       `yaml:"basicAuth"`
	AccessTokenAuth *AccessTokenAuth `yaml:"accessToken"`
}

func (a *Auth) Valid(header map[string][]string) bool {
	switch a.Type {
	case BasicAuthType:
		return a.BasicAuth.Valid(header)
	case AccessTokenType:
		return a.AccessTokenAuth.Valid(header)
	}
	return false
}

func (a *Auth) Token() string {
	switch a.Type {
	case BasicAuthType:
		return a.BasicAuth.Token()
	case AccessTokenType:
		return a.AccessTokenAuth.Token()
	default:
		return ""
	}
}

type AccessTokenAuth struct {
	AccessToken string `yaml:"accessToken"`
}

func (c *AccessTokenAuth) Validate() error {
	return validator.ValidateStruct(c,
		validation.Field(&c.AccessToken, validation.Required),
	)
}

func (c *AccessTokenAuth) Valid(header map[string][]string) bool {
	extractor := auth.BearerExtractor{}
	token, err := extractor.ExtractToken(textproto.MIMEHeader(header))
	if err != nil {
		return false
	}
	return utils.Equals(c.AccessToken, token)
}

func (c *AccessTokenAuth) Token() string {
	return fmt.Sprintf("Bearer %s", c.AccessToken)
}

type BasicAuth struct {
	Username string `yaml:"username"`
	Password string `yaml:"password"`
}

func (c *BasicAuth) Validate() error {
	return validator.ValidateStruct(c,
		validation.Field(&c.Username, validation.Required),
		validation.Field(&c.Password, validation.Required),
	)
}

func (c *BasicAuth) Valid(header map[string][]string) bool {
	extractor := auth.BasicExtractor{}
	token, err := extractor.ExtractToken(textproto.MIMEHeader(header))
	if err != nil {
		return false
	}
	payload, err := base64.StdEncoding.DecodeString(token)
	if err != nil {
		return false
	}
	username, password, found := strings.Cut(string(payload), ":")
	return found && utils.Equals(c.Username, username) && utils.Equals(c.Password, password)
}

func (c *BasicAuth) Token() string {
	return fmt.Sprintf("Basic %s", base64.StdEncoding.EncodeToString([]byte(fmt.Sprintf("%s:%s", c.Username, c.Password))))
}
