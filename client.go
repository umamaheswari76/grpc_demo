package main

import (
	"context"
	"fmt"
	"log"

	hw "grpc_demo/helloworld"

	"google.golang.org/grpc"
)

func main() {

	//caling or getting the instance of the server
	conn, err := grpc.Dial("localhost:50053", grpc.WithInsecure())
	if err != nil {
		log.Fatalf("Filed to connect: %v", err)

	}
	defer conn.Close()

	client := hw.NewGreeterClient(conn)

	name := "umamaheswari"
	age := 20

	response, err := client.SayHello(context.Background(), &hw.HelloRequest{Name: name, Age: int32(age)})
	if err != nil {
		log.Fatalf("Failed to call SayHello: %v", err)
	}

	fmt.Printf("Response: %s\n", response.Message)

}
