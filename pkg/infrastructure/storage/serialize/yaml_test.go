package serialize

import (
	"context"
	"testing"

	"github.com/stretchr/testify/require"
)

func TestSerializerWithYaml(t *testing.T) {
	ctx := context.Background()
	p := YamlSerializer[map[string]string]{}
	e := map[string]string{
		"id": "100",
	}
	v, err := p.Encode(ctx, e)
	require.NoError(t, err)
	act, err := p.Decode(ctx, v)
	require.NoError(t, err)
	require.Equal(t, e["id"], act["id"])
}
