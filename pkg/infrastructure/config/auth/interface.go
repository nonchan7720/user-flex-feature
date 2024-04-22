package auth

import "errors"

type Getter interface {
	Get(string) string
}

type Extractor interface {
	ExtractToken(getter Getter) (string, error)
}

var (
	ErrNoTokenInRequest = errors.New("no token present in request")
)
