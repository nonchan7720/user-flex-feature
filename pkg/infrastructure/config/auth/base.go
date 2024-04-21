package auth

import "strings"

type baseExtractor struct{}

func (b baseExtractor) ExtractToken(getter Getter, title string) (string, error) {
	titleCount := len(title)
	tokenHeader := getter.Get("Authorization")
	if len(tokenHeader) < titleCount || !strings.EqualFold(tokenHeader[:titleCount], title) {
		return "", ErrNoTokenInRequest
	}
	return tokenHeader[titleCount:], nil
}
