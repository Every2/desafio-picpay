package services

import (
	"net/http"

	"github.com/Every2/desafio-picpay/models"
	"github.com/Every2/desafio-picpay/repositories"
	"github.com/ericlagergren/decimal"
)

type UsersService struct {
	usersRepository *repositories.UsersRepository
}

func NewUserService(userRepository *repositories.UsersRepository) UsersService {
	return UsersService{
		usersRepository: userRepository,
	}
}

func (us UsersService) FindUserById(id int) (*models.Users, *models.ResponseError) {
	user, err := us.usersRepository.FindUserById(id)
	if err != nil {
		return nil, err
	}

	return user, nil
}

func (us UsersService) SaveUser() {
	repositories.CommitTransaction(us.usersRepository)
}

func ValidateTransaction(sender *models.Users, amount decimal.Big) *models.ResponseError {
	var u models.UserEnum = models.MERCHANT

	if sender.UserType.GetType() == u.GetType() {
		return &models.ResponseError{
			Message: "Merchant is not autorized to do transactions",
			Status: http.StatusBadRequest,
		}
	}

	if (sender.Balance.Cmp(&amount) < 0) {
		return &models.ResponseError{
			Message: "Not enough amount",
			Status: http.StatusBadRequest,
		}
	}

	return nil
}
