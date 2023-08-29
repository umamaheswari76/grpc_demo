package services

import (
	"context"
	"grpc_demo3/interfaces"
	"grpc_demo3/models"

	// "fmt"
	// "log"

	// "go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
)

// CreateCustomer(customer *models.Customer) error
// FindService(filter bson.M) ([]models.Customer, error)

type CustomerService struct {
	CustomerCollection *mongo.Collection
	ctx               context.Context
}

func CustomerServiceInit(collection *mongo.Collection, ctx context.Context) interfaces.ICustomer {
	return &CustomerService{ collection , ctx}
}

func (p *CustomerService) CreateCustomer(user *models.Customer) (error) {
	 
	_, err := p.CustomerCollection.InsertOne(p.ctx, &user)

	if err != nil {
		return   err
	}
	return nil

}

// func (p* CustomerService) FindService(filter bson.M)([] models.Customer,error){
// 	cursor,err:=p.CustomerCollection.Find(p.ctx,filter)
// 	fmt.Println(cursor)
// 	var results[] models.Customer
// 	for cursor.Next(context.TODO()) {
// 		var result models.Customer // Replace YourStruct with the type of documents in your collection
// 		if err := cursor.Decode(&result); err != nil {
// 			log.Fatal(err)
// 		}
// 		results = append(results, result)
// 	}
// 	if err != nil {
// 	 return   nil,err
//    }

//    fmt.Println("FINDED SUCCESSFULLY")
//    return results,nil
// }