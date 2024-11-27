package service

import (
	"awesomeProject/internal/utils"
	"awesomeProject/pkg/models"
	"errors"
	"fmt"
	"log"
)

//	func (s *Service) CreateUser(u *models.User) (*models.User, error) {
//		err := u.ValidateUser()
//		if err != nil {
//			return nil, err
//		}
//
//		userByUsername, err := s.Repository.GetUserByUsername(u.Username)
//		if err != nil && !errors.As(err, &ErrRecordNotFound) {
//			log.Printf("Failed to get user: %v", err)
//			return nil, err
//		}
//
//		if userByUsername != nil {
//			log.Printf("user with name %v is already exists", u.Username)
//			return nil, fmt.Errorf("User with name %v is already exists", u.Username)
//		}
//
//		return s.Repository.AddUser(u)
//	}
func (s *Service) CreateUser(u *models.User) (*models.User, error) {
	err := u.ValidateUser()
	if err != nil {
		return nil, err
	}

	userByUsername, err := s.Repository.GetUserByUsername(u.Username)
	if err != nil && !errors.As(err, &ErrRecordNotFound) {
		log.Printf("Failed to get user by username: %v", err)
		return nil, err
	}

	if userByUsername != nil {
		log.Printf("user with username %v is already exists", u.Username)
		return nil, fmt.Errorf("User with username %v is already exists", u.Username)
	}

	hashPassword, err := utils.HashPassword(*u.Password)
	if err != nil {
		return nil, fmt.Errorf("Failed to hash password: %v", err)
	}

	u.Password = &hashPassword

	return s.Repository.AddUser(u)
}

func (s *Service) SignIn(u *models.User) (string, error) {
	user, err := s.Repository.GetUserByUsername(u.Username)

	if err != nil {
		if errors.As(err, &ErrRecordNotFound) {
			return "", fmt.Errorf("User with name %s not found ", u.Username)
		}
		return "", err
	}

	if !utils.CheckPasswordHash(*u.Password, *user.Password) {
		return "", fmt.Errorf("Incorrect password entered")
	}

	token, err := utils.GenerateJWT(*user)
	if err != nil {
		return "", fmt.Errorf("Failed to generate token: %w", err)
	}

	return token, nil
}

func (s *Service) ListOfUsers() ([]models.User, error) {
	users, err := s.Repository.GetUsers()
	if err != nil {
		return nil, err
	}

	if len(users) == 0 {
		return nil, fmt.Errorf("no users found")
	}

	return users, nil
}

func (s *Service) GetUserByID(userId int) (*models.User, error) {
	userByID, err := s.Repository.GetUserByID(userId)
	if err != nil {
		if errors.As(err, &ErrRecordNotFound) {
			return nil, fmt.Errorf("user with id %v not found", userId)
		}
	}
	return userByID, err
}

func (s *Service) EditUser(u *models.User) error {
	_, err := s.Repository.GetUserByID(u.ID)
	if err != nil {
		if errors.As(err, &ErrUsersIDNotFound) {
			return fmt.Errorf("user with ID %v not found", u.ID)
		}
		return err
	}

	log.Println("EditUser func", u)

	err = s.Repository.UpdateUser(u)
	if err != nil {
		return fmt.Errorf("failed to update user with ID %v: %v", u.ID, err)
	}

	//if *updatedUser.Password != "" {
	//	updatedUser.Password = nil
	//}

	return nil
}

func (s *Service) DeleteUser(id int) (int, error) {
	_, err := s.Repository.GetUserByID(id)
	if err != nil {
		if errors.As(err, &ErrRecordNotFound) {
			return 0, fmt.Errorf("user with id %v not found", id)
		}
		return 0, err
	}

	deletedRows, err := s.Repository.DeleteUser(id)
	if err != nil {
		return 0, fmt.Errorf("failed to delete user with ID %v: %v", id, err)
	}

	return deletedRows, err
}
