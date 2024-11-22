package repository

import (
	"awesomeProject/pkg/models"
	"fmt"
	"log"
)

func (r *Repository) AddUser(u *models.User) (*models.User, error) {
	query := `insert into users (username, email, password) values (?, ?, ?) returning *`
	err := r.db.Raw(query, u.Username, u.Email, u.Password).Scan(&u).Error
	if err != nil {
		log.Printf("SignUp: Failed to add user: %v\n", err)
		return nil, fmt.Errorf("Failed to add user: %v\n", err)
	}

	return u, nil
}

func (r *Repository) GetUsers() ([]models.User, error) {
	var users []models.User

	// select * from users
	result := r.db.Omit("password").Find(&users)
	if result.Error != nil {
		log.Printf("GetUsers: Failed to get users: %v\n", result.Error)
		return nil, fmt.Errorf("Failed to get users: %v\n", result.Error)
	}
	return users, nil
}

func (r *Repository) GetUserByID(id int) (*models.User, error) {
	var user models.User

	// select * from users where id = id
	//result := r.db.Find(&user).Where("id = ?", id)
	query := `select * from users where id = ?;`
	err := r.db.Raw(query, id).Scan(&user).Error
	if err != nil {
		log.Printf("GetUserByID: Failed to get user: %v\n", err)
		return nil, fmt.Errorf("Failed to get user: %v\n", err)
	}
	return &user, nil
}

func (r *Repository) GetUserByUsername(username string) (*models.User, error) {
	var user models.User

	// select * from users where username = 'username'
	result := r.db.First(&user, "username = ?", username)
	if result.Error != nil {
		log.Printf("GetUserByUsername: Failed to get user: %v\n", result.Error)
		return nil, fmt.Errorf("Failed to get user: %v\n", result.Error)
	}
	return &user, nil
}

func (r *Repository) UpdateUser(u *models.User) error {
	log.Println("UpdateUser Repo", u.ID)
	result := r.db.Model(&models.User{}).Where("id = ?", u.ID).Updates(&u)
	if result.Error != nil {
		log.Printf("UpdateUser: Failed to update user: %v\n", result.Error)
		return fmt.Errorf("Failed to update user: %v\n", result.Error)
	}

	return nil // Return nil user and nil error if not returning updated user
}

func (r *Repository) DeleteUser(id int) (int, error) {
	result := r.db.Model(&models.User{}).
		Where("id = ?", id).
		Update("active", false)
	if result.Error != nil {
		log.Printf("DeleteUser: Failed to delete user: %v\n", result.Error)
		return 0, fmt.Errorf("Failed to delete user: %v\n", result.Error)
	}

	return id, nil
}
