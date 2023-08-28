package main

import (
	"context"
	"fmt"
	"net"

	"google.golang.org/grpc"

	hw "grpc_demo/helloworld" //import the generaed go code(and providing alias name as hw)
)

type server struct {
	hw.UnimplementedGreeterServer
}

func (s *server) SayHello(ctx context.Context, req *hw.HelloRequest) (*hw.HelloResponse, error) {
	return &hw.HelloResponse{
		Message: fmt.Sprintf("Hello %s, %d", req.Name,req.Age),
	}, nil
}

func main() {
	lis, err := net.Listen("tcp", ":50053")
	if err != nil {
		fmt.Printf("Failed to listen: %v", err)
		return
	}
	s := grpc.NewServer()
	hw.RegisterGreeterServer(s, &server{})

	fmt.Println("Server listening on :50053")
	if err := s.Serve(lis); err != nil {
		fmt.Printf("Failed to serve: %v", err)
	}
}
//server will expose methods, client will access thse method