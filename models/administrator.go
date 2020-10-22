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

// Administrator Summary
type Administrator struct {
	AdminID       string `json:"admin_id"`
	AdminName     string `json:"admin_name"`
	AdminPassword string `json:"admin_password"`
	// AdminEmail    string `json:"admin_email"`
}

//Administrators content list of Administrator
type Administrators struct {
	Administrators []Administrator
}
