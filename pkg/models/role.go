package models

type Role struct {
	Rolename    string `json:"role_name"`
	ID          int    `json:"id" gorm:"id, primaryKey"`
	Description string `json:"description"`
}
