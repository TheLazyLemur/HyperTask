package main

import (
	"context"
	"log"
	"net"

	"github.com/google/uuid"
	"github.com/thelazylemur/hypertask/services/task/internal/store"
	"github.com/thelazylemur/hypertask/services/task/pb"
	"github.com/thelazylemur/hypertask/services/task/types"

	"google.golang.org/grpc"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/reflection"
	"google.golang.org/grpc/status"
)

type server struct {
	pb.UnimplementedTaskServer
	port        string
	listener    net.Listener
	grpc_server *grpc.Server
	store       store.TaskStore
}

func NewServer(p string) *server {
	listener, err := net.Listen("tcp", p)
	if err != nil {
		log.Fatalln("Failed to listen:", err)
	}

	grpc_server := grpc.NewServer()

	return &server{
		port:        p,
		listener:    listener,
		grpc_server: grpc_server,
		store:       store.NewMemoryStore(),
	}
}

func (s *server) CreateTask(ctx context.Context, in *pb.CreateTaskRequest) (*pb.CreateTaskResponse, error) {
	task := &types.Task{
		Id:          uuid.New().String(),
		Name:        in.Name,
		Description: in.Description,
		Weight:      in.Weight,
	}

	t, err := s.store.CreateTask(ctx, task)
	if err != nil {
		return nil, status.Error(codes.AlreadyExists, err.Error())
	}

	return &pb.CreateTaskResponse{
		Id:          t.Id,
		Name:        t.Name,
		Description: t.Description,
		Weight:      t.Weight,
	}, nil
}

func (s *server) Start() error {
	reflection.Register(s.grpc_server)

	pb.RegisterTaskServer(s.grpc_server, s)
	if err := s.grpc_server.Serve(s.listener); err != nil {
		return err
	}

	return nil
}

func (s *server) Stop() {
	s.grpc_server.Stop()

	if err := s.listener.Close(); err != nil {
		log.Fatalln("Failed to close:", err)
	}
}
