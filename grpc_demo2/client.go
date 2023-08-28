package main

import (
	"context"
	"fmt"
	tk "grpc_demo/grpc_demo2/task"
	"log"

	"google.golang.org/grpc"
)

func main() {
	conn, err := grpc.Dial("localhost:50055", grpc.WithInsecure())
	if err != nil {
		log.Fatalf("Failed to connect: %v", err)
	}
	defer conn.Close()
	client := tk.NewTaskServiceClient(conn)

	//Add a task
	task := &tk.Task{
		Title: "buy groceries",
	}
	addResp, err := client.AddTask(context.Background(), task)
	if err != nil {
		log.Fatalf("Failed to add task: %v", err)
	}
	fmt.Printf("Added task with Id: %s\n", addResp.Id)

	//Receive tasks
	tasksResp, err := client.GetTask(context.Background(),&tk.Empty{})
	if err!=nil{
		log.Fatalf("Failed to retrieve tasks: %v",err)
	}

	fmt.Println("Tasks:")
	for _, task := range tasksResp.Tasks{
		fmt.Printf("ID: %s, Title: %s, Completed: %v\n", task.Id,task.Title, task.Completed)
	}
}
