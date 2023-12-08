package store

import (
	"context"
	"fmt"

	"github.com/thelazylemur/hypertask/services/task/types"
)

type TaskStore interface {
	CreateTask(ctx context.Context, in *types.Task) (*types.Task, error)
}

type memoryStore struct {
	tasks map[string]*types.Task
}

func NewMemoryStore() *memoryStore {
	return &memoryStore{
		tasks: make(map[string]*types.Task),
	}
}

func (s *memoryStore) CreateTask(ctx context.Context, in *types.Task) (*types.Task, error) {
	if _, ok := s.tasks[in.Id]; ok {
		return nil, fmt.Errorf("task with id %s already exists", in.Id)
	}

	s.tasks[in.Id] = in
	return in, nil
}
