package utils

import (
	"fmt"
	"os"
	"strconv"
	"strings"
)

func isMock(key string) bool {
	value := os.Getenv(strings.ToUpper(key))
	v, _ := strconv.ParseBool(value)
	return v
}

func IsMock() bool {
	return isMock("MOCK")
}

type ServiceName string

const ()

func IsServiceMock(serviceName ServiceName) bool {
	key := fmt.Sprintf("MOCK_%s", serviceName)
	return isMock(key)
}

func Getenv(key, default_ string) string {
	v := os.Getenv(key)
	if v == "" {
		return default_
	}
	return v
}
