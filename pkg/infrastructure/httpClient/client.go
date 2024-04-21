package httpclient

import (
	"net/http"

	"go.opentelemetry.io/contrib/instrumentation/net/http/otelhttp"
)

func DefaultClient(rt http.RoundTripper) *http.Client {
	client := *http.DefaultClient
	trans := otelhttp.NewTransport(rt)
	client.Transport = trans
	return &client
}
