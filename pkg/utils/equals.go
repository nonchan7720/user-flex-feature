package utils

import "crypto/subtle"

func Equals(v1, v2 string) bool {
	return subtle.ConstantTimeCompare([]byte(v1), []byte(v2)) == 1
}
