package feature

import (
	"context"

	"github.com/nonchan7720/user-flex-feature/pkg/container"
	"github.com/nonchan7720/user-flex-feature/pkg/infrastructure/config"
	"github.com/nonchan7720/user-flex-feature/pkg/infrastructure/feature/retriever"
	"github.com/samber/do"
	ffclient "github.com/thomaspoignant/go-feature-flag"
	ff_retriever "github.com/thomaspoignant/go-feature-flag/retriever"
)

func init() {
	do.Provide(container.Injector, Provide)
}

func Provide(i *do.Injector) (*Client, error) {
	ctx, _ := do.Invoke[context.Context](i)
	cfg := do.MustInvoke[*config.Config](i)
	retrievers, err := retriever.New(cfg)
	if err != nil {
		return nil, err
	}
	return newClient(ctx, cfg, retrievers...)
}

type Client struct {
	flag *ffclient.GoFeatureFlag
}

func (c *Client) Shutdown() error {
	c.flag.Close()
	return nil
}

var (
	_ do.Shutdownable = &Client{}
)

func newClient(ctx context.Context, cfg *config.Config, retrievers ...ff_retriever.Retriever) (*Client, error) {
	if ctx == nil {
		ctx = context.Background()
	}
	ff, err := ffclient.New(ffclient.Config{
		PollingInterval: cfg.PollingInterval,
		Context:         ctx,
		Retrievers:      retrievers,
	})
	if err != nil {
		return nil, err
	}
	return &Client{
		flag: ff,
	}, err
}