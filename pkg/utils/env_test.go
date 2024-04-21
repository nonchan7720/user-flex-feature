package utils

import (
	"os"
	"testing"

	"github.com/stretchr/testify/require"
)

func TestGetenv(t *testing.T) {
	os.Setenv("AAA", "BBB")
	actual := Getenv("AAA", "CCC")
	require.Equal(t, "BBB", actual)
	actual = Getenv("BBB", "CCC")
	require.Equal(t, "CCC", actual)
}
