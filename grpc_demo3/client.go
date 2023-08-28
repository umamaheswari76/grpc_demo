package main

import (
	"log"

	"google.golang.org/grpc"

	ct "grpc_demo3/customer"
)

func main() {
	conn, err := grpc.Dial("localhost:50056", grpc.WithInsecure())
	if err != nil {
		log.Fatalf("Failed to connect: %v", err)
	}
	defer conn.Close()

	client := ct.NewCustomerServiceClient(conn)

	//Insertng customer
	customer := &ct.Insert{		
		 Customer_name: "umamaheswari",
		 Balance: 5000,
		 bankid: "1",
	}

	addResp, err := client.Insert(context.Background(), customer)
	if err != nil {
		log.Fatalf("Failed to add task: %v", err)
	}
	fmt.Printf("Added customer: %s\n", addResp.)
}
