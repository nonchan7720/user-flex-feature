package feature

import (
	"context"

	"github.com/nonchan7720/user-flex-feature/pkg/container"
	"github.com/nonchan7720/user-flex-feature/pkg/domain/feature"
	domain_raft "github.com/nonchan7720/user-flex-feature/pkg/domain/raft"
	"github.com/nonchan7720/user-flex-feature/pkg/infrastructure/config"
	interface_raft "github.com/nonchan7720/user-flex-feature/pkg/interfaces/raft"
	"github.com/samber/do"
	ff_retriever "github.com/thomaspoignant/go-feature-flag/retriever"
)

func init() {
	do.Provide(container.Injector, Provide)
}

func Provide(i *do.Injector) (feature.Feature, error) {
	ctx, _ := do.Invoke[context.Context](i)
	cfg := do.MustInvoke[*config.Config](i)
	retrievers := do.MustInvoke[[]ff_retriever.Retriever](i)
	if cfg.IsRaftCluster() {
		raft := do.MustInvoke[*interface_raft.Raft](i)
		fsm := do.MustInvoke[domain_raft.FSM](i)
		client, err := newRaftClient(ctx, cfg, raft, retrievers...)
		if err != nil {
			return nil, err
		}
		fsm.SetClient(client)
		return client, nil
	} else {
		updateRetrievers := do.MustInvoke[[]feature.UpdateRetriever](i)
		updater := do.MustInvoke[feature.Updater](i)
		return newClient(ctx, cfg, updater, updateRetrievers, retrievers...)
	}
}
