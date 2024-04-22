package controller

import (
	"encoding/json"
	"fmt"
	"log/slog"

	"github.com/gin-gonic/gin"
	"github.com/nonchan7720/user-flex-feature/pkg/infrastructure/logging"
	"github.com/nonchan7720/user-flex-feature/pkg/interfaces/grpc/ofrep"
	"github.com/nonchan7720/user-flex-feature/pkg/utils"
	"google.golang.org/protobuf/types/known/structpb"
)

type ofrepAPI interface {
	GetOfrepV1Configuration(c *gin.Context, params GetOfrepV1ConfigurationParams) GetOfrepV1ConfigurationResponse
	PostOfrepV1EvaluateFlags(c *gin.Context, params PostOfrepV1EvaluateFlagsParams) PostOfrepV1EvaluateFlagsResponse
	PostOfrepV1EvaluateFlagsKey(c *gin.Context, key string) PostOfrepV1EvaluateFlagsKeyResponse
}

func (api *api) GetOfrepV1Configuration(c *gin.Context, params GetOfrepV1ConfigurationParams) GetOfrepV1ConfigurationResponse {
	ctx := c.Request.Context()
	in := &ofrep.GetConfigurationRequest{
		IfNoneMatch: c.GetHeader("If-None-Match"),
	}
	resp, _ := api.ofrepClient.GetConfiguration(ctx, in)
	if resp.ETag != "" {
		c.Writer.Header().Set("ETag", resp.ETag)
	}
	return GetOfrepV1ConfigurationResponse{
		JSON200: &ConfigurationResponse{
			Capabilities: &struct {
				CacheInvalidation *FeatureCacheInvalidation "json:\"cacheInvalidation,omitempty\""
			}{
				CacheInvalidation: &FeatureCacheInvalidation{
					Polling: &FeatureCacheInvalidationPolling{
						Enabled:            utils.Bool(resp.Configuration.Capabilities.CacheInvalidation.Polling.Enabled),
						MinPollingInterval: utils.Int[float64, float32](resp.Configuration.Capabilities.CacheInvalidation.Polling.GetMinPollingInterval()),
					},
				},
			},
		},
	}
}

func (api *api) PostOfrepV1EvaluateFlags(c *gin.Context, params PostOfrepV1EvaluateFlagsParams) PostOfrepV1EvaluateFlagsResponse {
	var body BulkEvaluationRequest
	if err := c.BindJSON(&body); err != nil {
		resp := BulkEvaluationFailure{
			ErrorCode:    string(INVALIDCONTEXT),
			ErrorDetails: utils.String(err.Error()),
		}
		return PostOfrepV1EvaluateFlagsResponse{
			JSON400: &resp,
		}
	}
	evalCtx, err := structpb.NewStruct(*body.Context)
	if err != nil {
		resp := BulkEvaluationFailure{
			ErrorCode:    string(INVALIDCONTEXT),
			ErrorDetails: utils.String(err.Error()),
		}
		return PostOfrepV1EvaluateFlagsResponse{
			JSON400: &resp,
		}
	}
	ctx := c.Request.Context()
	in := &ofrep.EvaluateFlagsBulkRequest{
		Context: &ofrep.EvaluationContext{
			Properties: evalCtx,
		},
	}
	resp, err := api.ofrepClient.EvaluateFlagsBulk(ctx, in)
	if err != nil {
		var resp BulkEvaluationFailure
		_ = json.Unmarshal([]byte(err.Error()), &resp)
		return PostOfrepV1EvaluateFlagsResponse{
			JSON400: &resp,
		}
	}

	var successes []EvaluationSuccess
	for _, val := range resp.Flags {
		value := val.GetSuccess()
		success := EvaluationSuccess{
			Key:      utils.String(value.Key),
			Reason:   (*EvaluationSuccessReason)(utils.String(value.Reason)),
			Variant:  utils.String(value.Variant),
			Metadata: convertMetadata(value.Metadata),
			union:    convertValue(value.Value),
		}
		successes = append(successes, success)
	}
	items := []BulkEvaluationSuccess_Flags_Item{}
	for _, success := range successes {
		item := BulkEvaluationSuccess_Flags_Item{}
		if err := item.FromEvaluationSuccess(success); err == nil {
			items = append(items, item)
		}
	}
	return PostOfrepV1EvaluateFlagsResponse{
		JSON200: &BulkEvaluationSuccess{
			Flags: &items,
		},
	}
}

func (api *api) PostOfrepV1EvaluateFlagsKey(c *gin.Context, key string) PostOfrepV1EvaluateFlagsKeyResponse {
	ctx := c.Request.Context()

	var body BulkEvaluationRequest
	if err := c.BindJSON(&body); err != nil {
		resp := EvaluationFailure{
			Key:          key,
			ErrorCode:    INVALIDCONTEXT,
			ErrorDetails: utils.String(err.Error()),
		}
		return PostOfrepV1EvaluateFlagsKeyResponse{
			JSON400: &resp,
		}
	}
	evalCtx, err := structpb.NewStruct(*body.Context)
	if err != nil {
		resp := EvaluationFailure{
			Key:          key,
			ErrorCode:    INVALIDCONTEXT,
			ErrorDetails: utils.String(err.Error()),
		}
		return PostOfrepV1EvaluateFlagsKeyResponse{
			JSON400: &resp,
		}
	}
	in := &ofrep.EvaluateFlagRequest{
		Key: key,
		Context: &ofrep.EvaluationContext{
			Properties: evalCtx,
		},
	}
	resp, err := api.ofrepClient.EvaluateFlag(ctx, in)
	if err != nil {
		slog.With(logging.WithStack(err)).ErrorContext(ctx, err.Error())
		failure := resp.GetFailure()
		if failure != nil {
			resp := EvaluationFailure{
				Key:          failure.Key,
				ErrorCode:    EvaluationFailureErrorCode(failure.ErrorCode),
				ErrorDetails: utils.String(failure.ErrorDetails),
			}
			return PostOfrepV1EvaluateFlagsKeyResponse{
				JSON400: &resp,
			}
		} else {
			resp := EvaluationFailure{
				Key:          key,
				ErrorCode:    GENERAL,
				ErrorDetails: utils.String(fmt.Sprintf("Error while evaluating the flag: %s", key)),
			}
			return PostOfrepV1EvaluateFlagsKeyResponse{
				JSON400: &resp,
			}
		}
	}
	respSuccess := resp.GetSuccess()
	success := EvaluationSuccess{
		Key:      utils.String(respSuccess.Key),
		Reason:   (*EvaluationSuccessReason)(utils.String(respSuccess.Reason)),
		Variant:  utils.String(respSuccess.Variant),
		Metadata: convertMetadata(respSuccess.Metadata),
		union:    convertValue(respSuccess.Value),
	}
	return PostOfrepV1EvaluateFlagsKeyResponse{
		JSON200: &success,
	}
}

func convertMetadata(metadata map[string]*ofrep.Metadata) *map[string]EvaluationSuccess_Metadata_AdditionalProperties {
	if metadata == nil {
		return nil
	}
	mp := map[string]EvaluationSuccess_Metadata_AdditionalProperties{}
	for key, value := range metadata {
		p := EvaluationSuccess_Metadata_AdditionalProperties{}
		switch val := value.Type.(type) {
		case *ofrep.Metadata_BooleanValue:
			p.FromEvaluationSuccessMetadata0(val.BooleanValue)
		case *ofrep.Metadata_NumberValue:
			p.FromEvaluationSuccessMetadata2(utils.ConvertInt[float32](val.NumberValue))
		case *ofrep.Metadata_StringValue:
			p.FromEvaluationSuccessMetadata1(val.StringValue)
		}
		mp[key] = p
	}
	return &mp
}

func convertValue(value interface{}) []byte {
	type v struct {
		Value interface{} `json:"value,omitempty"`
	}
	val := &v{}
	switch value := value.(type) {
	case *ofrep.EvaluationSuccess_BoolValue:
		val.Value = value.BoolValue
	case *ofrep.EvaluationSuccess_StringValue:
		val.Value = value.StringValue
	case *ofrep.EvaluationSuccess_IntegerValue:
		val.Value = value.IntegerValue
	case *ofrep.EvaluationSuccess_DoubleValue:
		val.Value = value.DoubleValue
	case *ofrep.EvaluationSuccess_ObjectValue:
		val.Value = value.ObjectValue.AsMap()
	}
	buf, _ := json.Marshal(val)
	return buf
}
