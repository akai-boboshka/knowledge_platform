package repository

import (
	"awesomeProject/pkg/models"
	"fmt"
	"log"
)

func (r *Repository) AddComment(c *models.Comment) (*models.Comment, error) {
	//comment := &models.Comment{}
	//result := r.db.Create(&c)
	query := `insert into comments (article_id, comments_content) values(?, ?) returning *`
	err := r.db.Raw(query, c.ArticleId, c.CommentsContent)
	if err != nil {
		log.Printf("AddComment: Failed to add article: %v\n", err)
		return nil, fmt.Errorf("Failed to add article: %v\n", err)
	}

	return c, nil
}

func (r *Repository) DeleteComment(id int) (int, error) {
	result := r.db.Model(&models.Comment{}).
		Where("comment_id = ?", id).
		Update("active", false)
	if result.Error != nil {
		log.Printf("DeleteComment: Failed to delete comment: %v\n", result.Error)
		return 0, fmt.Errorf("Failed to delete article: %v\n", result.Error)
	}

	return 0, nil
}

func (r *Repository) GetCommentById(id int) (*models.Comment, error) {
	var comment *models.Comment

	err := r.db.Where("comment_id = ?", id).First(&comment).Error
	if err != nil {
		log.Printf("GetCommentByID: Failed to get comment: %v\n", err)
		return nil, fmt.Errorf("cannot find comment with error: %v", err)
	}

	return comment, nil
}

func (r *Repository) GetComments() ([]models.Comment, error) {
	var comments []models.Comment

	result := r.db.Where("article_id = ?").Find(&comments).Error
	if result.Error != nil {
		log.Printf("GetComments: Failed to get comments: %v\n", result.Error)
		return nil, fmt.Errorf("Failed to get comments: %v\n", result.Error)
	}

	return comments, nil
}

func (r *Repository) UpdateComment(c *models.Comment) (*models.Comment, error) {
	result := r.db.Model(&c).Updates(&c)
	if result.Error != nil {
		log.Printf("ChangeComment: Failed to update comment: %v\n", result.Error)
		return nil, fmt.Errorf("Failed to update comment: %v\n", result.Error)
	}

	return c, nil
}
