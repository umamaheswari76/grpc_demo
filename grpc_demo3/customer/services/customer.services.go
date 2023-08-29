package rpcService

import (
	"context"
	ct "grpc_demo3/customer"
	"sync"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

type RPCServer struct {
	mu        sync.Mutex
	customers map[string]*ct.Customer
	ct.UnimplementedCustomerServiceServer
}

func (s *RPCServer) GetCustomer(ctx context.Context, req *ct.Empty) (*ct.CustomerList, error) {
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
