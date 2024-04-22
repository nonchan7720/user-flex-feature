package i18n

import (
	"context"

	"golang.org/x/text/language"
)

type languageKey struct{}

var (
	contextKey languageKey
)

func ToLanguage(ctx context.Context, acceptLanguage string) context.Context {
	t, _, _ := language.ParseAcceptLanguage(acceptLanguage)
	return context.WithValue(ctx, contextKey, t)
}

func FromLanguage(ctx context.Context) []language.Tag {
	if t, ok := ctx.Value(contextKey).([]language.Tag); ok {
		return t
	}
	return []language.Tag{language.Japanese}
}
