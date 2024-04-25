package feature

import (
	"context"

	"github.com/nonchan7720/user-flex-feature/pkg/domain/feature"
	"github.com/nonchan7720/user-flex-feature/pkg/infrastructure/config"
	"github.com/nonchan7720/user-flex-feature/pkg/infrastructure/logging"
	"github.com/samber/do"
	ffclient "github.com/thomaspoignant/go-feature-flag"
	ff_retriever "github.com/thomaspoignant/go-feature-flag/retriever"
)

type Client struct {
	*ffclient.GoFeatureFlag

	ffConfig         ffclient.Config
	updateRetrievers []feature.UpdateRetriever
	updater          feature.Updater
}

func (c *Client) Shutdown() error {
	c.GoFeatureFlag.Close()
	return nil
}

func (c *Client) Reset() error {
	flag, err := ffclient.New(c.ffConfig)
	if err == nil {
		*c.GoFeatureFlag = *flag
	}
	return err
}

func (c *Client) Provider() *ffclient.GoFeatureFlag {
	return c.GoFeatureFlag
}

func (c *Client) AppendOrUpdateRule(ctx context.Context, key string, rule *feature.Rule) error {
	return c.updater.AppendOrUpdateRule(ctx, c.updateRetrievers, key, rule)
}

var (
	_ do.Shutdownable = &Client{}
)

func newClient(ctx context.Context, cfg *config.Config, updater feature.Updater, updateRetrievers []feature.UpdateRetriever, retrievers ...ff_retriever.Retriever) (*Client, error) {
	if ctx == nil {
		ctx = context.Background()
	}
	ffConfig := ffclient.Config{
		PollingInterval: cfg.PollingInterval,
		Context:         ctx,
		Retrievers:      retrievers,
		Logger:          logging.StdLogger,
	}
	ff, err := ffclient.New(ffConfig)
	if err != nil {
		return nil, err
	}
	return &Client{
		GoFeatureFlag:    ff,
		ffConfig:         ffConfig,
		updateRetrievers: updateRetrievers,
		updater:          updater,
	}, err
}
