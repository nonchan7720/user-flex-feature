package feature

import (
	"crypto/sha256"
	"encoding/json"
	"fmt"

	"github.com/thomaspoignant/go-feature-flag/ffcontext"
)

type ConvertErrorCode string

func (c ConvertErrorCode) String() string {
	return string(c)
}

const (
	GENERAL             ConvertErrorCode = "GENERAL"
	INVALIDCONTEXT      ConvertErrorCode = "INVALID_CONTEXT"
	PARSEERROR          ConvertErrorCode = "PARSE_ERROR"
	TARGETINGKEYMISSING ConvertErrorCode = "TARGETING_KEY_MISSING"
)

type Context interface {
	ffcontext.Context
	Hash(key string) string
}

func NewContext(ctx map[string]interface{}) (Context, *GeneralError) {
	if ctx == nil {
		return nil, NewGeneralError(INVALIDCONTEXT, "User flex feature need an Evaluation context to work.")
	}
	if targetingKey, ok := ctx["targetingKey"].(string); ok {
		delete(ctx, "targetingKey")
		evalCtx := convertEvaluationCtxFromRequest(targetingKey, ctx)
		return evalCtx, nil
	}
	return nil, NewGeneralError(TARGETINGKEYMISSING, "User flex feature has received no targetingKey or a none string value that is not a string.")
}

func convertEvaluationCtxFromRequest(targetingKey string, custom map[string]interface{}) *ctx {
	c := ffcontext.NewEvaluationContextBuilder(targetingKey)
	for k, v := range custom {
		switch val := v.(type) {
		case float64:
			if isIntegral(val) {
				c.AddCustom(k, int(val))
				continue
			}
			c.AddCustom(k, val)
		default:
			c.AddCustom(k, val)
		}
	}
	return &ctx{c.Build()}
}

func isIntegral(val float64) bool {
	return val == float64(int64(val))
}

type ctx struct {
	ffcontext.Context
}

func (c *ctx) Hash(key string) string {
	mp := map[string]interface{}{
		"key":     key,
		"context": c.Context,
	}
	buf, _ := json.Marshal(&mp)
	hash := sha256.Sum256(buf)
	return fmt.Sprintf("%x", hash)
}
