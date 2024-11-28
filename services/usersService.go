package services

import (
	"errors"
	"github.com/Every2/desafio-picpay/models"
	"github.com/Every2/desafio-picpay/repositories"
	"github.com/shopspring/decimal"
	"gorm.io/gorm"
)


type UserService struct {
	repo *repositories.UserRepository
	db   *gorm.DB
}

func NewUserService(repo *repositories.UserRepository) *UserService {
	return &UserService{repo: repo}
}

func (s *UserService) ValidateTransaction(sender *models.User, amount decimal.Decimal) error {

	if sender.UserType == models.MERCHANT {
		return errors.New("usuário do tipo Lojista não está autorizado a realizar transação")
	}

	if sender.Balance.Cmp(amount) < 0 {
		return errors.New("saldo insuficiente")
	}

	return nil
}


func (s *UserService) FindUserByID(id uint) (*models.User, error) {
	return s.repo.FindByID(id)
}


func (s *UserService) CreateUser(userRequest models.UserRequest) (*models.User, error) {
	user := &models.User{
		FirstName: userRequest.FirstName,
		LastName:  userRequest.LastName,
		Document:  userRequest.Document,
		Email:     userRequest.Email,
		Password:  userRequest.Password,
		UserType:  userRequest.UserType,
		Balance:   userRequest.Balance,
	}


	if err := s.db.Save(user); err != nil {
		return nil, err.Error
	}

	return user, nil
}


func (s *UserService) SaveUser(user *models.User) error {
	return s.db.Save(user).Error
}
