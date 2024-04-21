package controller

import (
	"encoding/json"
	"fmt"
	"sort"

	"github.com/gin-gonic/gin"
	"github.com/nonchan7720/user-flex-feature/pkg/interfaces/api/controller/internal"
	"github.com/nonchan7720/user-flex-feature/pkg/utils"
	"github.com/thomaspoignant/go-feature-flag/ffcontext"
)

type ofrepAPI interface {
	GetOfrepV1Configuration(c *gin.Context, params GetOfrepV1ConfigurationParams) GetOfrepV1ConfigurationResponse
	PostOfrepV1EvaluateFlags(c *gin.Context, params PostOfrepV1EvaluateFlagsParams) PostOfrepV1EvaluateFlagsResponse
	PostOfrepV1EvaluateFlagsKey(c *gin.Context, key string) PostOfrepV1EvaluateFlagsKeyResponse
}

func (api *api) GetOfrepV1Configuration(c *gin.Context, params GetOfrepV1ConfigurationParams) GetOfrepV1ConfigurationResponse {
	return GetOfrepV1ConfigurationResponse{
		JSON200: &ConfigurationResponse{
			Capabilities: &struct {
				CacheInvalidation *FeatureCacheInvalidation "json:\"cacheInvalidation,omitempty\""
			}{
				CacheInvalidation: &FeatureCacheInvalidation{
					Polling: &FeatureCacheInvalidationPolling{
						Enabled:            utils.Bool(true),
						MinPollingInterval: utils.Int[int64, float32](api.cfg.PollingInterval.Milliseconds()),
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
	evalCtx, err := evaluationContextFromOFREPContext(*body.Context)
	if err != nil {
		resp := BulkEvaluationFailure{
			ErrorCode:    string(err.code),
			ErrorDetails: err.message,
		}
		return PostOfrepV1EvaluateFlagsResponse{
			JSON400: &resp,
		}
	}

	var successes []EvaluationSuccess
	allFlags := api.ff.AllFlagsState(evalCtx)
	flags := allFlags.GetFlags()
	for key, val := range flags {
		value := val.Value
		if val.Reason == internal.ReasonError {
			value = nil
		}
		success := EvaluationSuccess{
			Key:      utils.String(key),
			Reason:   (*EvaluationSuccessReason)(&val.Reason),
			Variant:  &val.VariationType,
			Metadata: convertMetadata(val.Metadata),
			union:    convertValue(value),
		}
		successes = append(successes, success)
	}
	sort.Slice(successes, func(i, j int) bool {
		return utils.StringValue(successes[i].Key) < utils.StringValue(successes[j].Key)
	})
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
	evalCtx, err := evaluationContextFromOFREPContext(*body.Context)
	if err != nil {
		resp := EvaluationFailure{
			Key:          key,
			ErrorCode:    err.code,
			ErrorDetails: err.message,
		}
		return PostOfrepV1EvaluateFlagsKeyResponse{
			JSON400: &resp,
		}
	}
	defaultValue := "thisisadefaultvaluethatItest1233%%"
	val, _ := api.ff.RawVariation(key, evalCtx, defaultValue)
	if val.Reason == internal.ReasonError {
		msg := utils.String(fmt.Sprintf("Error while evaluating the flag: %s", key))
		if val.ErrorCode == string(FLAGNOTFOUND) {
			return PostOfrepV1EvaluateFlagsKeyResponse{
				JSON404: &FlagNotFound{
					Key:          key,
					ErrorCode:    FLAGNOTFOUND,
					ErrorDetails: msg,
				},
			}
		} else {
			return PostOfrepV1EvaluateFlagsKeyResponse{
				JSON400: &EvaluationFailure{
					Key:          key,
					ErrorCode:    EvaluationFailureErrorCode(val.ErrorCode),
					ErrorDetails: msg,
				},
			}
		}
	}
	success := EvaluationSuccess{
		Key:      utils.String(key),
		Reason:   (*EvaluationSuccessReason)(&val.Reason),
		Variant:  &val.VariationType,
		Metadata: convertMetadata(val.Metadata),
		union:    convertValue(val.Value),
	}
	return PostOfrepV1EvaluateFlagsKeyResponse{
		JSON200: &success,
	}
}

type commonError struct {
	code    EvaluationFailureErrorCode
	message *string
}

func evaluationContextFromOFREPContext(ctx Context) (ffcontext.Context, *commonError) {
	if targetingKey, ok := ctx["targetingKey"].(string); ok {
		delete(ctx, "targetingKey")
		evalCtx := convertEvaluationCtxFromRequest(targetingKey, ctx)
		return evalCtx, nil
	}
	return ffcontext.EvaluationContext{}, &commonError{TARGETINGKEYMISSING, utils.String("GO Feature Flag has received no targetingKey or a none string value that is not a string.")}
}

func convertEvaluationCtxFromRequest(targetingKey string, custom map[string]interface{}) ffcontext.Context {
	ctx := ffcontext.NewEvaluationContextBuilder(targetingKey)
	for k, v := range custom {
		switch val := v.(type) {
		case float64:
			if isIntegral(val) {
				ctx.AddCustom(k, int(val))
				continue
			}
			ctx.AddCustom(k, val)
		default:
			ctx.AddCustom(k, val)
		}
	}
	return ctx.Build()
}

func isIntegral(val float64) bool {
	return val == float64(int64(val))
}

func convertMetadata(metadata map[string]interface{}) *map[string]EvaluationSuccess_Metadata_AdditionalProperties {
	if metadata == nil {
		return nil
	}
	mp := map[string]EvaluationSuccess_Metadata_AdditionalProperties{}
	for key, value := range metadata {
		p := EvaluationSuccess_Metadata_AdditionalProperties{}
		switch val := value.(type) {
		case bool:
			p.FromEvaluationSuccessMetadata0(val)
		case int, int32, int64, float32, float64:
			p.FromEvaluationSuccessMetadata2(utils.ConvertInt[float32](val))
		default:
			p.FromEvaluationSuccessMetadata1(fmt.Sprintf("%s", val))
		}
		mp[key] = p
	}
	return &mp
}

func convertValue(value interface{}) []byte {
	type v struct {
		Value interface{} `json:"value,omitempty"`
	}
	val := v{
		Value: value,
	}
	buf, _ := json.Marshal(&val)
	return buf
}
