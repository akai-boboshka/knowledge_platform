package repository

import (
	"awesomeProject/pkg/models"
	"fmt"
	"log"
)

func (r *Repository) AddProfile(p *models.Profile) (*models.Profile, error) {
	// insert into profiles (user_id, name, address) values (1, 'admin', 'admin')
	result := r.db.Create(&p)
	if result.Error != nil {
		log.Printf("AddProfile: Failed to add profile: %v\n", result.Error)
		return nil, fmt.Errorf("Failed to add profile: %v\n", result.Error)
	}

	return p, nil
}

func (r *Repository) GetProfiles() ([]models.Profile, error) {
	var profiles []models.Profile

	// select * from profiles
	result := r.db.Find(&profiles)
	if result.Error != nil {
		log.Printf("GetProfiles: Failed to get profiles: %v\n", result.Error)
		return nil, fmt.Errorf("Failed to get profiles: %v\n", result.Error)
	}

	return profiles, nil
}

func (r *Repository) GetProfileByID(id int) (*models.Profile, error) {
	var profile models.Profile

	// select * from profiles where user_id = id
	result := r.db.First(&profile, id)
	if result.Error != nil {
		log.Printf("GetProfileByID: Failed to get profile: %v\n", result.Error)
		return nil, fmt.Errorf("Failed to get profile: %v\n", result.Error)
	}

	return &profile, nil
}

func (r *Repository) UpdateProfile(p *models.Profile) (*models.Profile, error) {
	result := r.db.Model(&p).Updates(&p)
	if result.Error != nil {
		log.Printf("EditProfile: Failed to update profile: %v\n", result.Error)
		return nil, fmt.Errorf("Failed to update profile: %v\n", result.Error)
	}

	return p, nil
}

func (r *Repository) DeleteProfile(id int) (int, error) {
	// delete from profiles where user_id = id
	result := r.db.Model(&models.Profile{}).
		Where("profile_id = ?", id).
		Update("active", false)
	if result.Error != nil {
		log.Printf("DeleteProfile: Failed to delete profile: %v\n", result.Error)
		return 0, fmt.Errorf("Failed to delete profile: %v\n", result.Error)
	}

	return id, nil
}
