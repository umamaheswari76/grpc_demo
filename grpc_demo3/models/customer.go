package models

type Customer struct {
	Customername string  `json:"customer_name" bson:"customer_name"`
	Accountid    string  `json: "accountid" bson: "accountid"`
	Balance      float32 `json:"balance" bson:"balance"`
	Bankid       string  `json:"bankid" bson:"bankid"`
}

type InsertCustomerResponse struct {
	Customername string  `json:"customer_name" bson:"customer_name"`
	Accountid    string  `json: "accountid" bson: "accountid"`
	Balance      float32 `json:"balance" bson:"balance"`
	Bankid       string  `json:"bankid" bson:"bankid"`
}
