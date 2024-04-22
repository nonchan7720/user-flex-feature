package i18n

import (
	"bytes"
	_ "embed"

	"github.com/goccy/go-yaml"
	"golang.org/x/text/language"
	"golang.org/x/text/message"
)

var (
	//go:embed assets/ja.yaml
	jaTranslate []byte
	//go:embed assets/validation_ja.yaml
	jaValidationTranslate []byte

	//go:embed assets/en.yaml
	enTranslate []byte
	//go:embed assets/validation_en.yaml
	enValidationTranslate []byte

	fallbackLanguage *TranslateLanguage
)

func init() {
	_fallbackLanguage := readMessage(language.Japanese, jaTranslate, jaValidationTranslate)
	fallbackLanguage = &_fallbackLanguage
	readMessage(language.English, enTranslate, enValidationTranslate)
}

func merge(dst *TranslateLanguage, src []byte) {
	var result TranslateLanguage
	if err := yaml.NewDecoder(bytes.NewReader(src)).Decode(&result); err != nil {
		panic(err)
	}
	for key, value := range result.Keys {
		if _, ok := dst.Keys[key]; !ok {
			dst.Keys[key] = value
		}
	}
}

func readMessage(lang language.Tag, original []byte, validation []byte) TranslateLanguage {
	result := TranslateLanguage{
		Keys: map[string]string{},
	}
	merge(&result, original)
	merge(&result, validation)
	for key, value := range result.Keys {
		if value == "" {
			if fallbackLanguage != nil {
				value = fallbackLanguage.Keys[key]
			}
		}
		if err := message.SetString(lang, key, value); err != nil {
			panic(err)
		}
	}
	return result
}
