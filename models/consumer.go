package models

//Consumer (Summary)
type Consumer struct {
	ConsumerID   string `json:"consumer_id"`
	ConsumerName string `json:"consumer_name"`
	CompanyID    string `json:"consumer_password"`
	Phone1       string `json:"phone_number_1"`
}

//Consumers (list of Consumer)
type Consumers struct {
	Consumers []Consumer
}
