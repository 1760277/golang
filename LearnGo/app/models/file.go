package models

// File contains the details of a repository
type File struct {
	// ID       string `json:"file_id" form:"file_id" query:"file_id"`
	// RegisteredDate string `json:"file_registration_date" form:"file_registration_date" query:"file_registration_date"`
	// PageNum    int `json:"file_total_pages" form:"file_total_pages" query:"file_total_pages"`
	ID             string `json:"file_id" form:"file_id" query:"file_id"`
	RegisteredDate string `json:"file_created_at" form:"file_created_at" query:"file_created_at"`
	PageNum        int    `json:"file_total_pages" form:"file_total_pages" query:"file_total_pages"`
}
