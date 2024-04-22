package utils

import "errors"

func As[T any](err error) (T, bool) {
	var val T
	return val, errors.As(err, &val)
}
