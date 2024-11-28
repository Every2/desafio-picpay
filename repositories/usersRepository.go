package repositories

import (
	"gorm.io/gorm"
	"github.com/Every2/desafio-picpay/models"
)


type UserRepository struct {
	db *gorm.DB
}

func NewUserRepository(db *gorm.DB) *UserRepository {
	return &UserRepository{db: db}
}

func (r *UserRepository) FindByID(id uint) (*models.User, error) {
	var user models.User
	if err := r.db.First(&user, id).Error; err != nil {
		return nil, err
	}
	return &user, nil
}

func (r *UserRepository) FindByDocument(document string) (*models.User, error) {
	var user models.User
	if err := r.db.Where("document = ?", document).First(&user).Error; err != nil {
		return nil, err
	}
	return &user, nil
}