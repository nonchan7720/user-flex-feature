package httpclient

import (
	"context"
	"net/http"
)

type contextKey struct{}

var (
	httpClientContextKey contextKey
)

func FromContext(ctx context.Context) *http.Client {
	v, ok := ctx.Value(httpClientContextKey).(*http.Client)
	if !ok {
		return http.DefaultClient
	}
	return v
}

func ToContext(ctx context.Context, client *http.Client) context.Context {
	return context.WithValue(ctx, httpClientContextKey, client)
}
