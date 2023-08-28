package main

import (
	"context"
	"fmt"
	tk "grpc_demo/grpc_demo2/task"
	"net"
	"sync"

	"go.mongodb.org/mongo-driver/bson/primitive"
	"google.golang.org/grpc"
)

type server struct {
	mu    sync.Mutex
	tasks map[string]*tk.Task
	tk.UnimplementedTaskServiceServer
}

func (s *server) AddTask(ctx context.Context, req *tk.Task) (*tk.TaskResponse, error) {
	s.mu.Lock()
	defer s.mu.Unlock()

	taskID := generateID()
	req.Id = taskID
	s.tasks[taskID] = req

	return &tk.TaskResponse{Id: taskID}, nil

	// return &tk.TaskResponse{
	// 	Id: fmt.Sprintf("Task added %v",req.Id),
	// }, nil
}

func (s *server) GetTask(ctx context.Context, req *tk.Empty) (*tk.TaskList, error) {
	s.mu.Lock()
	defer s.mu.Unlock()

	tasks := make([]*tk.Task, 0, len(s.tasks))
	for _, task := range s.tasks {
		tasks = append(tasks, task)
	}
	return &tk.TaskList{Tasks: tasks}, nil
	// return &tk.TaskList{
	// 	Tasks: for _, val range Task,
	// }
}

func generateID() string{
	id := primitive.NewObjectID()
	id1 := id.Hex()
	return id1
}

func main() {
	lis, err := net.Listen("tcp", ":50055")
	if err != nil {
		fmt.Printf("Failed to listen: %v", err)
		return
	}

	s := grpc.NewServer()
	tk.RegisterTaskServiceServer(s, &server{
		tasks: make(map[string]*tk.Task),
	})

	fmt.Println("Server listening on: 50055")
	if err := s.Serve(lis); err != nil {
		fmt.Printf("Failed to serve: %v", err)
	}

}
