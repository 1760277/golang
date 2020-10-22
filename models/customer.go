package models

//Customer (Summary)
type Customer struct {
	ConsumerID       string `json:"consumer_id"`
	ConsumerName     string `json:"consumer_name"`
	ConsumerPassword string `json:"consumer_password"`
	Phone1           string `json:"phone_number_1"`
}

//Customers (list of Customer)
type Customers struct {
	Customers []Customer
}
