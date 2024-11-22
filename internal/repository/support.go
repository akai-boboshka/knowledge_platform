package repository

import (
	"awesomeProject/pkg/models"
	"fmt"
	"log"
)

//func (r *Repository) SendMessage(m *models.Support) (*models.Support, error) {
//result := r.db.Create(&m)

//if result.Error != nil {
//	log.Printf("SendMessage: Failed to send message: %v\n", result.Error)
//	return nil, fmt.Errorf("Failed to send message: %v\n", result.Error)
//}
//
//return m, nil
//}

func (r *Repository) GetMessageById(id int) (*models.Support, error) {
	var m *models.Support
	err := r.db.First(&m, id).Error
	if err != nil {
		log.Printf("FindMessageById: Failed to find message: %v\n", id)
		return nil, fmt.Errorf("Cannot find message with error: %v\n", err)
	}

	return m, nil
}

func (r *Repository) GetAllMessages() ([]models.Support, error) {
	var ms []models.Support

	err := r.db.Find(&ms).Error
	if err != nil {
		log.Printf("GetAllMessages: Failed to find messages: %v\n", err)
		return nil, fmt.Errorf("Failed to find messages: %v\n", err)
	}

	return ms, nil
}

func (r *Repository) DeleteMessage(id int) (int, error) {
	result := r.db.Model(&models.Support{}).
		Where("id = ?", id).
		Update("active", false)
	if result.Error != nil {
		log.Printf("DeleteMessage: Failed to delete message: %v\n", result.Error)
		return 0, fmt.Errorf("Failed to delete message: %v\n", result.Error)
	}

	return id, nil
}
