package config

import "google.golang.org/grpc/credentials"

type Gateway struct {
	Port int  `yaml:"port" default:"8080"`
	Grpc Grpc `yaml:"grpc,alias"`
}

func (c *Gateway) GrpcCredentials() credentials.TransportCredentials {
	return c.Grpc.GrpcCredentials()
}
