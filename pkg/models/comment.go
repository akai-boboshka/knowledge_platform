package models

type Comment struct {
	ID              int    `json:"id"  gorm:"id,primaryKey"`
	ArticleId       int    `json:"article_id" gorm:"article_id,foreignKey"`
	CommentsContent string `json:"comments_content"`
}
