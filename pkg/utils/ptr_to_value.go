package utils

import (
	"time"
)

func String(value string) *string {
	if value == "" {
		return nil
	}
	return &value
}

func StringValue(value *string) string {
	if value == nil {
		return ""
	}
	return *value
}

func Bool(value bool) *bool {
	return &value
}

func BoolValue(value *bool) bool {
	if value == nil {
		return false
	}
	return *value
}

func Time(val time.Time) *time.Time {
	return &val
}

type Integer interface {
	int | int32 | int64 | float32 | float64
}

func ConvertInt[R Integer](value any) R {
	switch v := value.(type) {
	case int:
		return R(v)
	case int32:
		return R(v)
	case int64:
		return R(v)
	case float32:
		return R(v)
	case float64:
		return R(v)
	default:
		return 0
	}
}

func Int[T Integer, R Integer](value T) *R {
	v := R(value)
	return &v
}

func IntValue[T Integer, R Integer](value *T) R {
	v := *value
	return R(v)
}
