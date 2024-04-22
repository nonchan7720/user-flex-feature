package feature

import (
	"bytes"
	"encoding/json"
	"fmt"
)

type GeneralError struct {
	Code    string `json:"code"`
	Message string `json:"message"`
}

func (e *GeneralError) ToJSON() string {
	buf := new(bytes.Buffer)
	_ = json.NewEncoder(buf).Encode(e)
	return buf.String()
}

func (e *GeneralError) Error() string {
	return e.ToJSON()
}

func NewGeneralError(code interface{}, message string) *GeneralError {
	return &GeneralError{
		Code:    fmt.Sprintf("%s", code),
		Message: message,
	}
}
