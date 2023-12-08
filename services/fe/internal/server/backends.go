package main

import (
	"hypertask/services/task/client"
)

type Backends interface {
	TaskClient() client.Client
}

type backends struct {
}

func NewBackends() Backends {
	return &backends{}
}

func (*backends) TaskClient() client.Client {
	return client.New("localhost:8081")
}
