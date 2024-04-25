package raft

import (
	"context"
	"fmt"
	"log/slog"
	"os"
	"time"

	transport "github.com/Jille/raft-grpc-transport"
	"github.com/hashicorp/raft"
	"github.com/nonchan7720/user-flex-feature/pkg/container"
	domain_raft "github.com/nonchan7720/user-flex-feature/pkg/domain/raft"
	"github.com/nonchan7720/user-flex-feature/pkg/infrastructure/config"
	"github.com/nonchan7720/user-flex-feature/pkg/infrastructure/grpc"
	user_flex_feature_raft "github.com/nonchan7720/user-flex-feature/pkg/interfaces/grpc/user-flex-feature-raft"
	"github.com/samber/do"
)

func init() {
	do.Provide(container.Injector, Provide)
}

type Raft struct {
	dir string
	cfg *config.Config
	*raft.Raft
	tm *transport.Manager

	join chan struct{}
	err  chan error
}

func Provide(i *do.Injector) (*Raft, error) {
	ctx, err := do.Invoke[context.Context](i)
	if err != nil {
		ctx = context.Background()
	}
	cfg := do.MustInvoke[*config.Config](i)
	fsm := do.MustInvoke[domain_raft.FSM](i)
	tm := do.MustInvoke[*transport.Manager](i)
	isSingleNode := cfg.Grpc.Endpoint() == cfg.Raft.JoinAddress
	return NewRaft(ctx, isSingleNode, cfg.Raft.Dir, cfg, tm, fsm)
}

func NewRaft(ctx context.Context, isSingleNode bool, dir string, cfg *config.Config, tm *transport.Manager, fsm raft.FSM) (*Raft, error) {
	r := &Raft{
		dir:  dir,
		cfg:  cfg,
		join: make(chan struct{}),
		err:  make(chan error),
	}
	if err := r.open(ctx, isSingleNode, fsm, tm); err != nil {
		return nil, err
	}
	return r, nil
}

func (r *Raft) open(ctx context.Context, enableSingleNode bool, fsm raft.FSM, tm *transport.Manager) error {
	const retainSnapshotCount = 2
	config := raft.DefaultConfig()
	config.LocalID = raft.ServerID(r.cfg.Raft.Id)
	config.Logger = newWrapHcLogger("")
	config.NoSnapshotRestoreOnStart = true

	snapshots, err := raft.NewFileSnapshotStore(r.dir, retainSnapshotCount, os.Stderr)
	if err != nil {
		return fmt.Errorf("file snapshot store: %v", err)
	}

	var (
		logStore    raft.LogStore
		stableStore raft.StableStore
	)
	logStore = raft.NewInmemStore()
	stableStore = raft.NewInmemStore()
	tp := tm.Transport()
	ra, err := raft.NewRaft(config, fsm, logStore, stableStore, snapshots, tp)
	if err != nil {
		return err
	}
	if enableSingleNode {
		configuration := raft.Configuration{
			Servers: []raft.Server{
				{
					ID:      config.LocalID,
					Address: tp.LocalAddr(),
				},
			},
		}
		ra.BootstrapCluster(configuration)
	} else {
		go r.runJoin(ctx)
	}
	go r.status(ctx)
	r.Raft = ra
	r.tm = tm
	return nil
}

func (r *Raft) status(ctx context.Context) {
	ticker := time.NewTicker(5 * time.Second)
	defer ticker.Stop()
	for {
		select {
		case <-ticker.C:
			slog.Info(fmt.Sprintf("status: %s", r.Raft.State().String()))
		case <-ctx.Done():
			return
		}
	}
}

func (r *Raft) Join() {
	r.join <- struct{}{}
}

func (r *Raft) runJoin(ctx context.Context) {
	<-r.join
	defer func() {
		close(r.err)
	}()
	conn, err := grpc.NewGrpcConnection(ctx, r.cfg.Raft.JoinAddress, nil, nil)
	if err != nil {
		r.err <- err
		return
	}
	client := user_flex_feature_raft.NewRaftServiceClient(conn)
	if _, err := client.Join(ctx, &user_flex_feature_raft.JoinRequest{
		Id:   r.cfg.Raft.Id,
		Addr: r.cfg.Grpc.Endpoint(),
	}); err != nil {
		r.err <- err
	}
	go func() {
		<-ctx.Done()
		defer conn.Close()
		ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
		defer cancel()
		_, _ = client.Leave(ctx, &user_flex_feature_raft.LeaveRequest{
			Id: r.cfg.Raft.Id,
		})
		r.tm.Close()
	}()
}

func (r *Raft) Error() error {
	return <-r.err
}
