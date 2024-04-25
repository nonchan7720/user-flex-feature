package grpc

import (
	"testing"

	user_flex_feature_raft "github.com/nonchan7720/user-flex-feature/pkg/interfaces/grpc/user-flex-feature-raft"
	"github.com/stretchr/testify/assert"
)

func TestGenerateName(t *testing.T) {
	name := generateServiceName[user_flex_feature_raft.RaftServiceClient]()
	assert.Equal(t, name, "*user_flex_feature_raft.RaftServiceClient")
}
