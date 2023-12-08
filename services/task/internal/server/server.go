package server

import (
	"context"
	"log"
	"net"

	"github.com/google/uuid"
	"hypertask/services/task/internal/pb"
	"hypertask/services/task/internal/store"
	"hypertask/services/task/types"

	"google.golang.org/grpc"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/reflection"
	"google.golang.org/grpc/status"
)

type Server struct {
	pb.UnimplementedTaskServer
	port        string
	listener    net.Listener
	grpc_server *grpc.Server
	store       store.TaskStore
}

func NewServer(p string) *Server {
	listener, err := net.Listen("tcp", p)
	if err != nil {
		log.Fatalln("Failed to listen:", err)
	}

	grpc_server := grpc.NewServer()

	return &Server{
		port:        p,
		listener:    listener,
		grpc_server: grpc_server,
		store:       store.NewMemoryStore(),
	}
}

func (s *Server) CreateTask(ctx context.Context, in *pb.CreateTaskRequest) (*pb.CreateTaskResponse, error) {
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

func (s *Server) GetTasks(ctx context.Context, in *pb.GetTasksRequest) (*pb.GetTasksResponse, error) {
	tasks, err := s.store.GetTasks(ctx)
	if err != nil {
		return nil, status.Error(codes.Internal, err.Error())
	}

	res := make([]*pb.TaskObject, 0)
	for _, t := range tasks {
		res = append(res, &pb.TaskObject{
			Id:          t.Id,
			Name:        t.Name,
			Description: t.Description,
			Weight:      t.Weight,
		})
	}

	return &pb.GetTasksResponse{
		Tasks: res,
	}, nil
}

func (s *Server) DeleteTask(ctx context.Context, in *pb.DeleteTaskRequest) (*pb.DeleteTaskResponse, error) {
	err := s.store.DeleteTask(ctx, in.Id)
	if err != nil {
		return nil, status.Error(codes.Internal, err.Error())
	}

	return &pb.DeleteTaskResponse{}, nil
}

func (s *Server) Start() error {
	reflection.Register(s.grpc_server)

	pb.RegisterTaskServer(s.grpc_server, s)
	if err := s.grpc_server.Serve(s.listener); err != nil {
		return err
	}

	return nil
}

func (s *Server) Stop() {
	s.grpc_server.Stop()

	if err := s.listener.Close(); err != nil {
		log.Fatalln("Failed to close:", err)
	}
}
