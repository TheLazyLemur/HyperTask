package store

import (
	"context"
	"fmt"

	"hypertask/services/task/types"
)

type TaskStore interface {
	CreateTask(ctx context.Context, in *types.Task) (*types.Task, error)
	GetTasks(ctx context.Context) ([]*types.Task, error)
	DeleteTask(ctx context.Context, id string) error
}

type memoryStore struct {
	tasks map[string]*types.Task
}

func NewMemoryStore() TaskStore {
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

func (s *memoryStore) GetTasks(ctx context.Context) ([]*types.Task, error) {
	tasks := make([]*types.Task, 0)

	for _, t := range s.tasks {
		tasks = append(tasks, t)
	}

	return tasks, nil
}

func (s *memoryStore) DeleteTask(ctx context.Context, id string) error {
	if _, ok := s.tasks[id]; !ok {
		return fmt.Errorf("task with id %s does not exist", id)
	}

	delete(s.tasks, id)

	return nil
}
