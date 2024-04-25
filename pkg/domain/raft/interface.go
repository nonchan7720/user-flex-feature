package raft

import (
	"github.com/hashicorp/raft"
	"github.com/nonchan7720/user-flex-feature/pkg/domain/feature"
)

type FSM interface {
	raft.FSM
	SetClient(client feature.Feature)
}
