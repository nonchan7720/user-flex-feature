package feature

import (
	"context"

	"github.com/nonchan7720/user-flex-feature/pkg/domain/feature"
	domain_raft "github.com/nonchan7720/user-flex-feature/pkg/domain/raft"
	"github.com/nonchan7720/user-flex-feature/pkg/infrastructure/config"
	"github.com/nonchan7720/user-flex-feature/pkg/infrastructure/logging"
	interface_raft "github.com/nonchan7720/user-flex-feature/pkg/interfaces/raft"
	"github.com/samber/do"
	ffclient "github.com/thomaspoignant/go-feature-flag"
	ff_retriever "github.com/thomaspoignant/go-feature-flag/retriever"
)

type RaftClient struct {
	*ffclient.GoFeatureFlag

	ffConfig   ffclient.Config
	raftConfig *config.Raft
	raft       *interface_raft.Raft
}

func (c *RaftClient) Shutdown() error {
	c.GoFeatureFlag.Close()
	return nil
}

func (c *RaftClient) Reset() error {
	flag, err := ffclient.New(c.ffConfig)
	if err == nil {
		*c.GoFeatureFlag = *flag
	}
	return err
}

func (c *RaftClient) AppendOrUpdateRule(ctx context.Context, key string, rule *feature.Rule) error {
	buf, err := domain_raft.CommandMarshal(domain_raft.UpdateCommand, key, &rule)
	if err != nil {
		return err
	}
	return c.raft.Apply(buf, c.raftConfig.ApplyTimeout).Error()
}

func (c *RaftClient) Provider() *ffclient.GoFeatureFlag {
	return c.GoFeatureFlag
}

var (
	_ do.Shutdownable = &RaftClient{}
)

func newRaftClient(ctx context.Context, cfg *config.Config, raft *interface_raft.Raft, retrievers ...ff_retriever.Retriever) (*RaftClient, error) {
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
	return &RaftClient{
		GoFeatureFlag: ff,
		ffConfig:      ffConfig,
		raftConfig:    cfg.Raft,
		raft:          raft,
	}, err
}
