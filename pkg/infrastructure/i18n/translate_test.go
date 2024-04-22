package i18n

import (
	"context"
	"testing"

	"github.com/stretchr/testify/require"
	"golang.org/x/text/language"
	"golang.org/x/text/message"
)

func TestTranslate(t *testing.T) {
	_ = message.SetString(language.Japanese, "test", "これはテストです")
	value := Translate(context.Background(), "test")
	require.Equal(t, value, "これはテストです")
}
