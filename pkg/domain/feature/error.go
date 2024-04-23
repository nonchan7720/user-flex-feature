package feature

import (
	"bytes"
	"encoding/json"
	"errors"
	"fmt"
)

type GeneralError struct {
	ErrorCode    string `json:"errorCode"`
	ErrorDetails string `json:"errorDetails"`
	Key          string `json:"key,omitempty"`
}

func (e *GeneralError) ToJSON() string {
	buf := new(bytes.Buffer)
	_ = json.NewEncoder(buf).Encode(e)
	return buf.String()
}

func (e *GeneralError) Error() string {
	return e.ToJSON()
}

func NewGeneralError(errorCode interface{}, errorDetails string, keys ...string) *GeneralError {
	var key string
	if len(keys) > 0 {
		key = keys[0]
	}
	return &GeneralError{
		ErrorCode:    fmt.Sprintf("%s", errorCode),
		ErrorDetails: errorDetails,
		Key:          key,
	}
}

func IsGeneralError(err error) *GeneralError {
	var val *GeneralError
	if errors.As(err, &val) {
		return val
	}
	return nil
}
