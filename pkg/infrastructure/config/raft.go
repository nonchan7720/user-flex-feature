package config

import (
	"time"
)

type Raft struct {
	Id           string        `yaml:"id" default:"local-1"`
	Dir          string        `yaml:"directory" default:"/tmp/raft"`
	ApplyTimeout time.Duration `yaml:"applyTimeout" default:"10s"`
	JoinAddress  string        `yaml:"join" default:"localhost:40001"`
}
