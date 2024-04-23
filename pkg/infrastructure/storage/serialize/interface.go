package serialize

import "context"

type Serialize[V any] interface {
	Encode(ctx context.Context, value V) (string, error)
	Decode(ctx context.Context, value string) (V, error)
}
