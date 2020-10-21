package models

import (
	"database/sql"
)

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

// User entity
type User struct {
	UserID      string         `json:"id"`
	Username    string         `json:"user_name"`
	Password    string         `json:"password"`
	Status      bool           `json:"status"`
	CreatedDate sql.NullTime   `json:"created_date"`
	CreatedBy   sql.NullString `json:"created_by"`
	UpdatedDate sql.NullTime   `json:"updated_date"`
	UpdatedBy   sql.NullString `json:"updated_by"`
}
