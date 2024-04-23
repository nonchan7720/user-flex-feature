package serialize

import (
	"context"
	"encoding/json"
)

type JsonSerializer[V any] struct{}

var (
	_ Serialize[any] = &JsonSerializer[any]{}
)

func (*JsonSerializer[V]) Encode(ctx context.Context, value V) (string, error) {
	v, err := json.Marshal(&value)
	if err != nil {
		return "", err
	}
	return string(v), nil
}

func (*JsonSerializer[V]) Decode(ctx context.Context, value string) (V, error) {
	var v V
	if err := json.Unmarshal([]byte(value), &v); err != nil {
		return v, err
	}
	return v, nil
}
