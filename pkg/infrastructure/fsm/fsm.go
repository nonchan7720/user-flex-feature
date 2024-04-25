package fsm

import (
	"context"
	"encoding/json"
	"io"

	hasicorp_raft "github.com/hashicorp/raft"
	"github.com/nonchan7720/user-flex-feature/pkg/container"
	"github.com/nonchan7720/user-flex-feature/pkg/domain/feature"
	"github.com/nonchan7720/user-flex-feature/pkg/domain/raft"
	"github.com/samber/do"
)

func init() {
	do.Provide(container.Injector, ProvideRaftFSM)
}

type raftFSM struct {
	updateRetrievers []feature.UpdateRetriever
	client           feature.Feature
	updater          feature.Updater
}

var (
	_ hasicorp_raft.FSM = (*raftFSM)(nil)
)

func ProvideRaftFSM(i *do.Injector) (raft.FSM, error) {
	updateRetrievers := do.MustInvoke[[]feature.UpdateRetriever](i)
	updater := do.MustInvoke[feature.Updater](i)
	return newRaftFSM(updateRetrievers, updater), nil
}

func newRaftFSM(updateRetrievers []feature.UpdateRetriever, updater feature.Updater) *raftFSM {
	return &raftFSM{
		updateRetrievers: updateRetrievers,
		updater:          updater,
	}
}

func (r *raftFSM) applyAppendOrUpdateRule(ctx context.Context, key string, rule *feature.Rule) error {
	if err := r.updater.AppendOrUpdateRule(ctx, r.updateRetrievers, key, rule); err != nil {
		return err
	}
	if err := r.client.Reset(); err != nil {
		return err
	}
	return nil
}

func (r *raftFSM) Apply(log *hasicorp_raft.Log) interface{} {
	var cmd raft.Command
	_ = json.Unmarshal(log.Data, &cmd)
	switch cmd.Op {
	case raft.UpdateCommand:
		var rule feature.Rule
		_ = json.Unmarshal(cmd.Value, &rule)
		return r.applyAppendOrUpdateRule(context.Background(), cmd.Key, &rule)
	}
	return nil
}

func (r *raftFSM) Snapshot() (hasicorp_raft.FSMSnapshot, error) {
	return &fsmSnapshot{}, nil
}

func (r *raftFSM) Restore(snapshot io.ReadCloser) error {
	return nil
}

func (r *raftFSM) SetClient(client feature.Feature) {
	r.client = client
}

type fsmSnapshot struct {
}

func (f *fsmSnapshot) Persist(sink hasicorp_raft.SnapshotSink) error {
	defer sink.Close()
	return nil
}

func (f *fsmSnapshot) Release() {}
