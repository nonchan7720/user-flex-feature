package serialize

import (
	"context"

	"github.com/goccy/go-yaml"
)

type YamlSerializer[V any] struct{}

var (
	_ Serialize[any] = &YamlSerializer[any]{}
)

func (*YamlSerializer[V]) Encode(ctx context.Context, value V) (string, error) {
	v, err := yaml.Marshal(value)
	if err != nil {
		return "", err
	}
	return string(v), nil
}

func (*YamlSerializer[V]) Decode(ctx context.Context, value string) (V, error) {
	var v V
	if err := yaml.Unmarshal([]byte(value), &v); err != nil {
		return v, err
	}
	return v, nil
}
