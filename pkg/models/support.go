package models

type Support struct {
	ID          int    `json:"id"`
	UserID      int    `json:"user_id" gorm:"user_id,foreignKey"`
	SupportText string `json:"support_text"`
}
