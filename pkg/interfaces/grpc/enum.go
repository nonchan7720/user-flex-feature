package grpc

type code string

func (c code) String() string {
	return string(c)
}

type EvaluationFailureErrorCode = code

// Defines values for EvaluationFailureErrorCode.
const (
	GENERAL             EvaluationFailureErrorCode = "GENERAL"
	INVALIDCONTEXT      EvaluationFailureErrorCode = "INVALID_CONTEXT"
	PARSEERROR          EvaluationFailureErrorCode = "PARSE_ERROR"
	TARGETINGKEYMISSING EvaluationFailureErrorCode = "TARGETING_KEY_MISSING"
)

type EvaluationSuccessReason = code

const (
	DISABLED       EvaluationSuccessReason = "DISABLED"
	SPLIT          EvaluationSuccessReason = "SPLIT"
	STATIC         EvaluationSuccessReason = "STATIC"
	TARGETINGMATCH EvaluationSuccessReason = "TARGETING_MATCH"
	UNKNOWN        EvaluationSuccessReason = "UNKNOWN"
)

type FlagNotFoundErrorCode = code

const (
	FLAGNOTFOUND FlagNotFoundErrorCode = "FLAG_NOT_FOUND"
)
