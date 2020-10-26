package models

//Consumer (Summary)
type Consumer struct {
	ConsumerID       string `json:"consumer_id"`
	Name             string `json:"consumer_name"`
	NameKana         string `json:"consumer_name_kana"`
	BirthDate        string `json:"birthdate"`
	CompanyID        string `json:"consumer_password"`
	Phone1           string `json:"phone_number_1"`
	Phone2           string `json:"phone_number_2"`
	MailAddress      string `json:"mail_address"`
	PostalCode       string `json:"postal_code"`
	Address          string `json:"address"`
	BranchNumber     string `json:"branch_number"`
	Memo             string `json:"consumer_memo"`
	EmployeeID       string `json:"employee_id"`
	RegistrationDate string `json:"consumer_registration_date"`
	Status           int    `json:"consumer_status"`
	CreateDate       string `json:"create_date"`
	UpdateDate       string `json:"update_date"`
	UpdatePgmID      string `json:"update_pgm_id"`
	RequestID        string `json:"request_id"`
	Version          string `json:"version"`
}

//Consumers (list of Consumer)
type Consumers struct {
	Consumers []Consumer
}
