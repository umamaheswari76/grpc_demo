package main

import (
	"context"
	"fmt"
	"net"
	"sync"
	ct "grpc_demo3/customer"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"google.golang.org/grpc"
)

type server struct {
	mu    sync.Mutex
	customers map[string]*ct.Customer
	ct.UnimplementedCustomerServiceServer
}

// rpc Insert (Customer) returns (InsertCustomerResponse);
// rpc GetTask(Empty) returns (CustomerList);

func (s *server) Insert(ctx context.Context, req *ct.Customer) (*ct.InsertCustomerResponse, error) {
	s.mu.Lock()
	defer s.mu.Unlock()

	ID  := generateID();
	req.Accountid = ID
	s.customers[ID] = req

	return &ct.InsertCustomerResponse{accountid: ID}, nil

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

func generateID() string {
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
