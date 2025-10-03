package repository

import (
	"errors"
	"log"

	"github.com/spattyan/confirmaai-backend/internal/users/domain"
	"gorm.io/gorm"
)

type gormRepository struct {
	database *gorm.DB
}

func (g gormRepository) FindByEmail(email string) (*domain.User, error) {
	var user domain.User
	result := g.database.First(&user, "email = ?", email)

	if result.Error != nil {
		log.Printf("Error fetching user: %v", result.Error)
		return nil, errors.New("failed to fetch user")
	}

	return &user, nil
}

func (g gormRepository) FindByPhone(phone string) (*domain.User, error) {
	var user domain.User
	result := g.database.First(&user, "phone = ?", phone)

	if result.Error != nil {
		log.Printf("Error fetching user: %v", result.Error)
		return nil, errors.New("failed to fetch user")
	}

	return &user, nil
}

func (g gormRepository) Create(user *domain.User) error {
	err := g.database.Create(&user).Error

	if err != nil {
		log.Printf("Error creating user: %v", err)
		return errors.New("failed to create user")
	}

	return nil
}

func (g gormRepository) FindByID(id string) (*domain.User, error) {
	var user domain.User
	result := g.database.First(&user, "id = ?", id)

	if result.Error != nil {
		log.Printf("Error fetching user: %v", result.Error)
		return nil, errors.New("failed to fetch user")
	}

	return &user, nil
}

func NewGormRepository(database *gorm.DB) domain.Repository {
	return &gormRepository{
		database: database,
	}
}
