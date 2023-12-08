package client

import (
	"context"
	"hypertask/services/task/internal/pb"
	"log"

	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

type Client interface {
	CreateTask(name string, desc string, weight int32) (*Task, error)
	GetTasks() ([]*Task, error)
	DeleteTask(id string) error
	Close()
}

type client struct {
	conn   *grpc.ClientConn
	client pb.TaskClient
}

func (c *client) Close() {
	defer c.conn.Close()
}

func New(serverAddr string) Client {
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
		Weight:      weight,
		Description: desc,
	})
	if err != nil {
		log.Println("Failed to create task:", err)
		return nil, err
	}

	return &Task{
		Id:          task.Id,
		Name:        task.Name,
		Description: task.Description,
		Weight:      task.Weight,
	}, nil
}

func (c *client) GetTasks() ([]*Task, error) {
	tasks, err := c.client.GetTasks(context.Background(), &pb.GetTasksRequest{})
	if err != nil {
		log.Println("Failed to get tasks:", err)
		return nil, err
	}

	res := make([]*Task, 0)

	for _, t := range tasks.Tasks {
		res = append(res, &Task{
			Id:          t.Id,
			Name:        t.Name,
			Description: t.Description,
			Weight:      t.Weight,
		})
	}

	return res, nil
}

func (c *client) DeleteTask(id string) error {
	_, err := c.client.DeleteTask(context.Background(), &pb.DeleteTaskRequest{
		Id: id,
	})

	if err != nil {
		log.Println("Failed to delete task:", err)
		return err
	}

	return nil
}
