package models

// NullTime type
// type NullTime time.Time

// func (ns *NullTime) Scan(value interface{}) error {
// 	if value == nil {
// 		*ns = time.Time{}
// 		return nil
// 	}
// 	strVal, ok := value.(string)
// 	if !ok {
// 		return errors.New("Column is not a string")
// 	}
// 	*ns = NullString(strVal)
// 	return nil
// }

// func (ns *NullString) Scan(value interface{}) error {
// 	if value == nil {
// 		*ns = ""
// 		return nil
// 	}
// 	strVal, ok := value.(string)
// 	if !ok {
// 		return errors.New("Column is not a string")
// 	}
// 	*ns = NullString(strVal)
// 	return nil
// }

// //NullString type
// type NullString string

// Employee Summary
type Employee struct {
	CompanyID                     string `json:"company_id"`
	EmployeeID                    string `json:"employee_id"`
	EmployeeName                  string `json:"employee_name"`
	MailAddress                   string `json:"mail_address"`
	RoleCode                      string `json:"role_code"`
	BranchNumber                  string `json:"branch_number"`
	RegistrationDate              string `json:"registration_date"`
	LoginPassword                 string `json:"login_password"`
	LoginPasswordStatus           int    `json:"login_password_status"`
	LoginFailNumber               int    `json:"login_fail_number"`
	LoginStatus                   int    `json:"login_status"`
	LastLoginDate                 string `json:"last_login_date"`
	LoginDate                     string `json:"login_date"`
	LoginPasswordChangeFailNumber int    `json:"login_password_change_fail_number"`
	LoginPasswordChangeHistory    string `json:"login_password_change_history"`
	LockEndDate                   string `json:"lock_end_date"`
	LockStatus                    int    `json:"lock_status"`
	LoginPasswordUpdateDate       string `json:"login_password_update_date"`
	EmployeeStatus                int    `json:"employee_status"`
	CreateDate                    string `json:"create_date"`
	UpdateDate                    string `json:"update_date"`
	UpdatePmgID                   string `json:"update_pgm_id"`
	RequestID                     string `json:"request_id"`
	Version                       int    `json:"version"`
}

//Employees content list of Employee
type Employees struct {
	Employees []Employee
}
