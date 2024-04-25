package grpc

import (
	"bytes"
	"context"
	"crypto/sha1"
	"encoding/json"
	"fmt"
	"log/slog"
	"sort"
	"time"

	hc_raft "github.com/hashicorp/raft"
	"github.com/nonchan7720/user-flex-feature/pkg/domain/feature"
	"github.com/nonchan7720/user-flex-feature/pkg/infrastructure/config"
	"github.com/nonchan7720/user-flex-feature/pkg/interfaces/grpc/ofrep"
	user_flex_feature "github.com/nonchan7720/user-flex-feature/pkg/interfaces/grpc/user-flex-feature"
	user_flex_feature_raft "github.com/nonchan7720/user-flex-feature/pkg/interfaces/grpc/user-flex-feature-raft"
	"github.com/nonchan7720/user-flex-feature/pkg/interfaces/raft"
	svc_feature "github.com/nonchan7720/user-flex-feature/pkg/services/feature"
	"github.com/nonchan7720/user-flex-feature/pkg/utils"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
	"google.golang.org/protobuf/types/known/structpb"
)

func newUserFlexFeatureServer(svc svc_feature.Service, ff feature.Feature, raft *raft.Raft, cfg *config.Config) *server {
	return &server{
		cfg:  cfg,
		svc:  svc,
		ff:   ff,
		raft: raft,
	}
}

type server struct {
	cfg  *config.Config
	svc  svc_feature.Service
	ff   feature.Feature
	raft *raft.Raft
}

var (
	_ ServiceServer = (*server)(nil)
)

func (s *server) RuleUpdate(ctx context.Context, in *user_flex_feature.RuleUpdateRequest) (*user_flex_feature.RuleUpdateResponse, error) {
	if s.cfg.IsRaftCluster() {
		if s.raft.State() != hc_raft.Leader {
			client, err := leader[user_flex_feature.UserFlexFeatureServiceClient](ctx, s.raft)
			if err != nil {
				return nil, err
			}
			return client.RuleUpdate(ctx, in)
		}
	}

	rule := in.GetRule()
	if rule == nil {
		return nil, status.Error(codes.InvalidArgument, "rule is empty.")
	}
	var percentages map[string]float64
	if v := rule.GetPercentageValue(); v != nil {
		percentages = v.GetValue()
	}
	var progressiveRollout *feature.ProgressiveRollout
	if v := rule.GetProgressiveRolloutValue(); v != nil {
		conv := func(v *user_flex_feature.ProgressiveRolloutStep) *feature.ProgressiveRolloutStep {
			if v == nil {
				return nil
			}
			value := &feature.ProgressiveRolloutStep{
				Variation: utils.String(v.GetVariationValue()),
				Date:      utils.ToTime(v.GetDateValue()),
			}
			if v, ok := v.GetPercentage().(*user_flex_feature.ProgressiveRolloutStep_PercentageValue); ok {
				value.Percentage = &v.PercentageValue
			}
			return value
		}
		progressiveRollout = &feature.ProgressiveRollout{
			Initial: conv(v.GetInitial()),
			End:     conv(v.GetEnd()),
		}
	}
	var disable *bool
	if v, ok := rule.GetDisable().(*user_flex_feature.Rule_DisableValue); ok {
		disable = &v.DisableValue
	}
	modelRule := feature.Rule{
		Name:               rule.GetName(),
		Query:              rule.GetQuery(),
		VariationResult:    rule.GetVariationResult(),
		Percentages:        percentages,
		ProgressiveRollout: progressiveRollout,
		Disable:            disable,
	}
	if err := s.svc.AppendOrUpdateRule(ctx, in.GetKey(), &modelRule); err != nil {
		return nil, status.Error(codes.InvalidArgument, err.Error())
	}
	return &user_flex_feature.RuleUpdateResponse{
		Result: true,
	}, nil
}

func (s *server) EvaluateFlag(ctx context.Context, in *ofrep.EvaluateFlagRequest) (*ofrep.EvaluateFlagResponse, error) {
	key := in.GetKey()
	mp := map[string]interface{}{}
	if c := in.GetContext(); c != nil {
		if v := c.GetProperties(); v != nil {
			mp = v.AsMap()
		}
	}
	var evalCtx feature.Context
	if eval, err := feature.NewContext(mp); err != nil {
		return nil, status.Error(codes.InvalidArgument, feature.NewGeneralError(err.ErrorCode, err.ErrorDetails, key).ToJSON())
	} else {
		evalCtx = eval
	}
	val, err := s.svc.EvaluateFlag(ctx, key, evalCtx)
	if err != nil {
		if generalErr := feature.IsGeneralError(err); generalErr != nil {
			return nil, status.Error(codes.InvalidArgument, generalErr.ToJSON())
		}
		return nil, status.Error(codes.Internal, err.Error())
	}
	success := &ofrep.EvaluationSuccess{
		Key:      key,
		Reason:   val.Reason,
		Variant:  val.VariationType,
		Metadata: convertMetadata(val.Metadata),
	}
	convertValue(val.Value, success)
	return &ofrep.EvaluateFlagResponse{
		Result: &ofrep.EvaluateFlagResponse_Success{
			Success: success,
		},
	}, nil
}

func (s *server) EvaluateFlagsBulk(ctx context.Context, in *ofrep.EvaluateFlagsBulkRequest) (*ofrep.EvaluateFlagsBulkResponse, error) {
	mp := map[string]interface{}{}
	if c := in.GetContext(); c != nil {
		if v := c.GetProperties(); v != nil {
			mp = v.AsMap()
		}
	}
	var evalCtx feature.Context
	if eval, err := feature.NewContext(mp); err != nil {
		return nil, status.Error(codes.InvalidArgument, err.ToJSON())
	} else {
		evalCtx = eval
	}

	var successes []*ofrep.EvaluationSuccess
	flags := s.svc.EvaluateFlagsBulk(ctx, evalCtx)
	for key, val := range flags {
		success := &ofrep.EvaluationSuccess{
			Key:      key,
			Reason:   val.Reason,
			Variant:  val.VariationType,
			Metadata: convertMetadata(val.Metadata),
		}
		convertValue(val.Value, success)
		successes = append(successes, success)
	}
	sort.Slice(successes, func(i, j int) bool {
		return successes[i].Key < successes[j].Key
	})
	results := []*ofrep.EvaluationResult{}
	for _, success := range successes {
		results = append(results, &ofrep.EvaluationResult{
			Result: &ofrep.EvaluationResult_Success{
				Success: success,
			},
		})
	}
	return &ofrep.EvaluateFlagsBulkResponse{
		Flags: results,
		ETag:  etag(results),
	}, nil
}

func (s *server) GetConfiguration(ctx context.Context, in *ofrep.GetConfigurationRequest) (*ofrep.GetConfigurationResponse, error) {
	conf := &ofrep.Configuration{
		Capabilities: &ofrep.Capability{
			CacheInvalidation: &ofrep.FeatureCacheInvalidation{
				Polling: &ofrep.Polling{
					Enabled:            true,
					MinPollingInterval: float64(s.cfg.PollingInterval.Milliseconds()),
				},
			},
		},
	}
	return &ofrep.GetConfigurationResponse{
		Configuration: conf,
		ETag:          etag(conf),
	}, nil
}

func (s *server) Join(ctx context.Context, in *user_flex_feature_raft.JoinRequest) (*user_flex_feature_raft.EmptyResponse, error) {
	if !s.cfg.IsRaftCluster() {
		return nil, status.Error(codes.Unimplemented, "Un running raft cluster")
	}
	if s.raft.State() != hc_raft.Leader {
		client, err := leader[user_flex_feature_raft.RaftServiceClient](ctx, s.raft)
		if err != nil {
			return nil, err
		}
		return client.Join(ctx, in)
	}
	nodeId := hc_raft.ServerID(in.Id)
	nodeAddr := hc_raft.ServerAddress(in.Addr)

	configFuture := s.raft.GetConfiguration()
	if err := configFuture.Error(); err != nil {
		return nil, status.Error(codes.Internal, err.Error())
	}
	for _, srv := range configFuture.Configuration().Servers {
		if srv.ID == nodeId || srv.Address == nodeAddr {
			if srv.Address == nodeAddr && srv.ID == nodeId {
				return nil, status.Error(codes.AlreadyExists, "Already exists.")
			}
			future := s.raft.RemoveServer(srv.ID, 0, 0)
			if err := future.Error(); err != nil {
				return nil, status.Error(codes.Internal, fmt.Sprintf("error removing existing node %s at %s: %s", nodeId, nodeAddr, err))
			}
		}
	}
	timeout := 10 * time.Second
	if err := s.raft.AddVoter(hc_raft.ServerID(in.Id), hc_raft.ServerAddress(in.Addr), 0, timeout).Error(); err != nil {
		return nil, status.Error(codes.Internal, err.Error())
	}
	return &user_flex_feature_raft.EmptyResponse{}, nil
}

func (s *server) Leave(ctx context.Context, in *user_flex_feature_raft.LeaveRequest) (*user_flex_feature_raft.EmptyResponse, error) {
	if !s.cfg.IsRaftCluster() {
		return nil, status.Error(codes.Unimplemented, "Un running raft cluster")
	}
	if s.raft.State() != hc_raft.Leader {
		client, err := leader[user_flex_feature_raft.RaftServiceClient](ctx, s.raft)
		if err != nil {
			return nil, err
		}
		return client.Leave(ctx, in)
	}
	nodeId := hc_raft.ServerID(in.Id)
	future := s.raft.RemoveServer(nodeId, 0, 0)
	if err := future.Error(); err != nil {
		slog.Warn(fmt.Sprintf("error removing existing node %s: %s", nodeId, err))
	}
	return &user_flex_feature_raft.EmptyResponse{}, nil
}

func (s *server) JoinCluster() error {
	s.raft.Join()
	return s.raft.Error()
}

func convertMetadata(metadata map[string]interface{}) map[string]*ofrep.Metadata {
	if metadata == nil {
		return nil
	}
	mp := map[string]*ofrep.Metadata{}
	for key, value := range metadata {
		p := &ofrep.Metadata{}
		switch val := value.(type) {
		case bool:
			p.Type = &ofrep.Metadata_BooleanValue{
				BooleanValue: val,
			}
		case int, int32, int64, float32, float64:
			p.Type = &ofrep.Metadata_NumberValue{
				NumberValue: utils.ConvertInt[float64](val),
			}
		default:
			p.Type = &ofrep.Metadata_StringValue{
				StringValue: fmt.Sprintf("%s", val),
			}
		}
		mp[key] = p
	}
	return mp
}

func convertValue(value interface{}, success *ofrep.EvaluationSuccess) {
	switch val := value.(type) {
	case bool:
		success.Value = &ofrep.EvaluationSuccess_BoolValue{
			BoolValue: val,
		}
	case float32, float64:
		success.Value = &ofrep.EvaluationSuccess_DoubleValue{
			DoubleValue: utils.ConvertInt[float64](val),
		}
	case int, int32, int64:
		success.Value = &ofrep.EvaluationSuccess_IntegerValue{
			IntegerValue: utils.ConvertInt[int64](val),
		}
	case string:
		success.Value = &ofrep.EvaluationSuccess_StringValue{
			StringValue: val,
		}
	case map[string]interface{}:
		ov, err := structpb.NewStruct(val)
		if err != nil {
			slog.Error(err.Error())
		}
		success.Value = &ofrep.EvaluationSuccess_ObjectValue{
			ObjectValue: ov,
		}
	}
}

func etag(value any) string {
	buf := new(bytes.Buffer)
	_ = json.NewEncoder(buf).Encode(value)
	getHash := func(str string) string {
		return fmt.Sprintf("%x", sha1.Sum([]byte(str)))
	}
	str := buf.String()
	return fmt.Sprintf("\"%d-%s\"", len(str), getHash(str))
}
