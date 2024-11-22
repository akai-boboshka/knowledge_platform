package service

import (
	"awesomeProject/pkg/models"
	"fmt"
)

func (s *Service) CreateComment(id int, c *models.Comment) (*models.Comment, error) {
	article, err := s.Repository.GetArticleByID(id)
	if err != nil {
		return nil, err
	}
	if article == nil {
		return nil, fmt.Errorf("Article with ID: %d didn't found", c.ArticleId)
	}

	return s.Repository.AddComment(c)
}

func (s *Service) EditComment(c *models.Comment) (*models.Comment, error) {
	//article, err := s.Repository.GetArticleByID(c.ArticleId)
	//if err != nil {
	//	return nil, err
	//}
	//if article == nil {
	//	return nil, fmt.Errorf("Article with ID: %d didn't found", c.ArticleId)
	//}

	return s.Repository.UpdateComment(c)
}

func (s *Service) AllCommentsOfArticle() ([]models.Comment, error) {
	comment, err := s.Repository.GetComments()
	if err != nil {

		return nil, err
	}

	return comment, nil
}

func (s *Service) RemoveComment(CommentId int) (int, error) {
	_, err := s.Repository.GetCommentById(CommentId)
	if err != nil {
		return 0, err
	}

	_, err = s.Repository.DeleteComment(CommentId)
	if err != nil {
		return 0, err
	}

	return s.Repository.DeleteComment(CommentId)
}
