package i18n

import (
	"bytes"
	"context"
	"html/template"

	"golang.org/x/text/language"
	"golang.org/x/text/message"
)

var (
	languages = []language.Tag{
		language.Japanese,
		language.English,
	}
)

type TranslateLanguage struct {
	Keys map[string]string `yaml:"keys"`
}

func Translate(ctx context.Context, key string, args ...interface{}) string {
	t := FromLanguage(ctx)
	matcher := language.NewMatcher(languages)
	tag, _, _ := matcher.Match(t...)
	p := message.NewPrinter(tag)
	return p.Sprintf(key, args...)
}

func TranslateValidationError(ctx context.Context, key string, params map[string]interface{}) string {
	t := FromLanguage(ctx)
	matcher := language.NewMatcher(languages)
	tag, _, _ := matcher.Match(t...)
	p := message.NewPrinter(tag)
	message := p.Sprintf(key)
	var buf bytes.Buffer
	tpl := template.Must(template.New("").Parse(message))
	if err := tpl.Execute(&buf, params); err != nil {
		return message
	}
	return buf.String()
}
