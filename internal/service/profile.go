package service

import (
	"awesomeProject/pkg/models"
	"errors"
	"fmt"
)

func (s *Service) ListProfiles() ([]models.Profile, error) {
	profiles, err := s.Repository.GetProfiles()
	if err != nil {
		return nil, err
	}

	if len(profiles) == 0 {
		return profiles, fmt.Errorf("No profiles found")
	}

	return profiles, nil
}

func (s *Service) CreateProfile(p *models.Profile) (*models.Profile, error) {
	_, err := s.Repository.GetUserByID(p.UserID)
	if err != nil {
		if errors.Is(err, ErrRecordNotFound) {
			return nil, fmt.Errorf("user with id %d not found", p.UserID)
		}
		return nil, err
	}

	profileByID, err := s.Repository.GetProfileByID(p.UserID)
	if err != nil && !errors.Is(err, ErrRecordNotFound) {
		return nil, err
	}
	if profileByID != nil {
		return nil, fmt.Errorf("profile with user id %d already exists", p.UserID)
	}

	return s.Repository.AddProfile(p)
}

func (s *Service) EditProfile(p *models.Profile) (*models.Profile, error) {
	_, err := s.Repository.GetUserByID(p.UserID)
	if err != nil {
		if errors.Is(err, ErrRecordNotFound) {
			return nil, fmt.Errorf("user with id %d not found", p.UserID)
		}
		return nil, err
	}

	_, err = s.Repository.GetProfileByID(p.UserID)
	if err != nil {
		if errors.Is(err, ErrRecordNotFound) {
			return nil, fmt.Errorf("profile with user id %d not found", p.UserID)
		}
		return nil, err
	}

	return s.Repository.UpdateProfile(p)
}

func (s *Service) GetProfileByID(id int) (*models.Profile, error) {
	profileByID, err := s.Repository.GetProfileByID(id)
	if err != nil {
		if errors.Is(err, ErrRecordNotFound) {
			return nil, fmt.Errorf("profile with id %d not found", id)
		}
		return nil, err
	}

	return profileByID, nil
}

func (s *Service) DeleteProfile(id int) (int, error) {
	_, err := s.Repository.GetProfileByID(id)
	if err != nil {
		if errors.Is(err, ErrRecordNotFound) {
			return 0, fmt.Errorf("profile with id %d not found", id)
		}
		return 0, err
	}

	return s.Repository.DeleteProfile(id)
}
