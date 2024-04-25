package raft

import "encoding/json"

type commandType string

const (
	UpdateCommand commandType = "update"
)

type Command struct {
	Op    commandType `yaml:"ope"`
	Key   string      `yaml:"key"`
	Value []byte      `yaml:"value,omitempty"`
}

func CommandMarshal(ope commandType, key string, value any) ([]byte, error) {
	buf, err := json.Marshal(value)
	if err != nil {
		return nil, err
	}
	cmd := &Command{
		Op:    ope,
		Key:   key,
		Value: buf,
	}
	return json.Marshal(cmd)
}
