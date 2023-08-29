package main

import (
	"context"
	"fmt"
	"log"

	"google.golang.org/grpc"

	ct "grpc_demo3/customer"
)

func main() {
	conn, err := grpc.Dial("localhost:4004", grpc.WithInsecure())
	if err != nil {
		log.Fatalf("Failed to connect: %v", err)
	}
	defer conn.Close()

	client := ct.NewCustomerServiceClient(conn)

	//Insertng customer
	customer := &ct.Customer{
		Customername: "umamaheswari",
		// Accountid:    "",
		Balance:      5000,
		Bankid:       "1",
	}

	addResp, err := client.Insert(context.Background(), customer)
	if err != nil {
		log.Fatalf("Failed to add task: %v", err)
	}
	fmt.Printf("Added customer: %s\n", addResp.Accountid)

	// customerResp, err := client.GetCustomer(context.Background(),&ct.Empty{})
	// if err!=nil{
	// 	log.Fatalf("Failed to retrieve tasks: %v",err)
	// }

	// fmt.Println("Customers:")
	// for _, customer := range customerResp.Customers{
	// 	fmt.Printf("accountid: %s, customername: %s, balance: %v, bankid: %v\n", customer.Accountid,customer.Customername, customer.Balance, customer.Bankid)
	// }
}
