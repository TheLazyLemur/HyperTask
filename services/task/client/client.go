package client

import (
	"context"
	"github.com/thelazylemur/hypertask/services/task/pb"
	"log"

	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

type client struct {
	conn   *grpc.ClientConn
	client pb.TaskClient
}

func (c *client) Close() {
	defer c.conn.Close()
}

func New(serverAddr string) *client {
	i := insecure.NewCredentials()

	opts := []grpc.DialOption{
		grpc.WithTransportCredentials(i),
	}

	conn, err := grpc.Dial(serverAddr, opts...)
	if err != nil {
		log.Fatalln("Failed to dial:", err)
	}

	c := pb.NewTaskClient(conn)

	return &client{
		conn:   conn,
		client: c,
	}
}

func (c *client) CreateTask(name string, desc string, weight int32) (*Task, error) {
	task, err := c.client.CreateTask(context.Background(), &pb.CreateTaskRequest{
		Name:        name,
		Description: desc,
		Weight:      weight,
	})
	if err != nil {
		return nil, err
	}

	return &Task{
		Id:          task.Id,
		Name:        task.Name,
		Description: task.Description,
		Weight:      task.Weight,
	}, nil
}
