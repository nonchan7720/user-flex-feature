package errors

import "errors"

var (
	ErrNotfound = errors.New("Not found.")
)

func IsNotFound(err error) bool {
	return errors.Is(err, ErrNotfound)
}
