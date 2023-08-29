package main

import (
	constants "command-line-arguments/home/umamaheswari/Documents/uma/banking_api_golang/constants/index.go"
	"context"
	"fmt"
	ct "grpc_demo3/customer"
	"net"
	"sync"

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

	return &ct.InsertCustomerResponse{Accountid: ID}, nil

	// return &tk.TaskResponse{
	// 	Id: fmt.Sprintf("Task added %v",req.Id),
	// }, nil
}

func (s *server) GetCustomer(ctx context.Context, req *ct.Empty) (*ct.CustomerList, error) {
	s.mu.Lock()
	defer s.mu.Unlock()

	customers := make([]*ct.Customer, 0, len(s.customers))
	for _, task := range s.customers {
		customers = append(customers, task)
	}
	return &ct.CustomerList{Customers: customers}, nil
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
	
	lis, err := net.Listen("tcp", constants.Port)
	if err != nil {
		fmt.Printf("Failed to listen: %v", err)
		return
	}

	s := grpc.NewServer()
	ct.RegisterCustomerServiceServer(s, &server{
		customers: make(map[string]*ct.Customer),
	})

	fmt.Println("Server listening on: %v",constants.Port)
	if err := s.Serve(lis); err != nil {
		fmt.Printf("Failed to serve: %v", err)
	}

}
