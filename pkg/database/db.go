package database

import (
	"awesomeProject/pkg/models"
	"fmt"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

func InitDataBase(configs *models.Config) (*gorm.DB, error) {
	db := fmt.Sprintf("host=%s port=%d user=%s dbname=%s password=%s",
		configs.Server.Host, configs.Server.Port,
		configs.Database.User, configs.Database.Name, configs.Database.Password)

	database, err := gorm.Open(postgres.Open(db), &gorm.Config{})
	if err != nil {
		return nil, err
	}

	return database, nil
}
