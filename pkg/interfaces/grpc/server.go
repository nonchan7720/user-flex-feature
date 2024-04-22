package grpc

import (
	"bytes"
	"context"
	"crypto/sha1"
	"encoding/json"
	"fmt"
	"log/slog"
	"sort"

	"github.com/nonchan7720/user-flex-feature/pkg/domain/feature"
	"github.com/nonchan7720/user-flex-feature/pkg/infrastructure/config"
	inf_feature "github.com/nonchan7720/user-flex-feature/pkg/infrastructure/feature"
	"github.com/nonchan7720/user-flex-feature/pkg/interfaces/grpc/internal"
	"github.com/nonchan7720/user-flex-feature/pkg/interfaces/grpc/ofrep"
	user_flex_feature "github.com/nonchan7720/user-flex-feature/pkg/interfaces/grpc/user-flex-feature"
	svc_feature "github.com/nonchan7720/user-flex-feature/pkg/services/feature"
	"github.com/nonchan7720/user-flex-feature/pkg/utils"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
	"google.golang.org/protobuf/types/known/structpb"
)

func newUserFlexFeatureServer(svc svc_feature.Service, ff *inf_feature.Client, cfg *config.Config) *server {
	return &server{
		svc: svc,
		ff:  ff,
		cfg: cfg,
	}
}

type server struct {
	svc svc_feature.Service
	ff  *inf_feature.Client
	cfg *config.Config
}

var (
	_ ServiceServer = (*server)(nil)
)

func (s *server) RuleUpdate(ctx context.Context, in *user_flex_feature.RuleUpdateRequest) (*user_flex_feature.RuleUpdateResponse, error) {
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
	evalCtx, err := feature.NewContext(mp)
	if err != nil {
		return nil, status.Error(codes.InvalidArgument, feature.NewGeneralError(err.ErrorCode, err.ErrorDetails, key).ToJSON())
	}
	defaultValue := "thisisadefaultvaluethatItest1233%%"
	val, _ := s.ff.RawVariation(key, evalCtx, defaultValue)
	if val.Reason == internal.ReasonError {
		msg := fmt.Sprintf("Error while evaluating the flag: %s", key)
		return nil, status.Error(codes.InvalidArgument, feature.NewGeneralError(val.ErrorCode, msg, key).ToJSON())
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
	evalCtx, err := feature.NewContext(mp)
	if err != nil {
		return nil, status.Error(codes.InvalidArgument, err.ToJSON())
	}

	var successes []*ofrep.EvaluationSuccess
	allFlags := s.ff.AllFlagsState(evalCtx)
	flags := allFlags.GetFlags()
	for key, val := range flags {
		value := val.Value
		if val.Reason == internal.ReasonError {
			value = nil
		}
		success := &ofrep.EvaluationSuccess{
			Key:      key,
			Reason:   val.Reason,
			Variant:  val.VariationType,
			Metadata: convertMetadata(val.Metadata),
		}
		convertValue(value, success)
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
