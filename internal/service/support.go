package service

import (
	"awesomeProject/pkg/models"
	"errors"
	"fmt"
	"log"
)

//func (s *Service) SendMessage(userId int, m *models.Support) (*models.Support, error) {
//	_, err := s.Repository.GetUserByID(userId)
//	if err != nil && !errors.As(err, &ErrRecordNotFound) {
//		log.Printf("Failed to get user: %v", err)
//		return nil, err
//	}

//return s.Repository.SendMessage(m)
//}

func (s *Service) GetMessageById(Id int) (*models.Support, error) {
	message, err := s.Repository.GetMessageById(Id)
	if err != nil {
		return nil, err
	}

	if message == nil {
		return nil, fmt.Errorf("Article with ID: %d didn't found", Id)
	}

	return message, nil
}

func (s *Service) GetAllMessages() ([]models.Support, error) {
	message, err := s.Repository.GetAllMessages()
	if err != nil {
		return nil, err
	}

	if message == nil {
		return nil, fmt.Errorf("Didn't find any messages \n")
	}

	return s.Repository.GetAllMessages()
}

func (s *Service) DeleteMessage(id int) (int, error) {
	_, err := s.Repository.GetMessageById(id)
	if err != nil && !errors.As(err, &ErrRecordNotFound) {
		log.Printf("Failed to get message: %v", err)
		return 0, err
	}

	return s.Repository.DeleteMessage(id)
}
