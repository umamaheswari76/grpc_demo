package interfaces

import (
	"grpc_demo3/models"

	//"go.mongodb.org/mongo-driver/bson"
)

type ICustomer interface {
	CreateCustomer(customer *models.Customer) error
	//FindService(filter bson.M) ([]models.Customer, error)
}
