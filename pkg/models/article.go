package models

import (
	"fmt"
)

type Article struct {
	ID           int    `json:"id" gorm:"id;primaryKey"`
	Title        string `json:"title"`
	ProfileID    int    `json:"profile_id" gorm:"profile_id,foreignKey"`
	ArticlesText string `json:"articles_text"`
}

// ValidateArticle функция, которая проверяет данные пользователя на соответствие заданным критериям.
func (a *Article) ValidateArticle() error {
	if err := a.ValidateArticlesText(); err != nil {
		return err
	}
	return nil
}

// ValidateArticlesText, чтобы текст статьи не был меньше 250 символов
func (a *Article) ValidateArticlesText() error {
	if len(a.ArticlesText) < 250 {
		return fmt.Errorf("The text of your article is too short")
	}
	return nil
}

type ReadLater struct {
	ArticleID int `json:"article_id"`
	ID        int `json:"favorites"`
	ProfileId int `json:"profile_id"`
}
