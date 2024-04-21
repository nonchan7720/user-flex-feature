package di

import (
	"net/http"

	"github.com/nonchan7720/user-flex-feature/pkg/container"
	httpclient "github.com/nonchan7720/user-flex-feature/pkg/infrastructure/httpClient"
	"github.com/samber/do"
)

func init() {
	do.Provide(container.Injector, Provide)
}

func Provide(i *do.Injector) (*http.Client, error) {
	rt, _ := do.Invoke[http.RoundTripper](i)
	client := httpclient.DefaultClient(rt)
	return client, nil
}
