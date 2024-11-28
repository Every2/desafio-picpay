package services

import (
	"errors"
	"time"
	"github.com/Every2/desafio-picpay/models"
	"github.com/Every2/desafio-picpay/repositories"
	"github.com/shopspring/decimal"
	"gorm.io/gorm"
)

type TransactionService struct {
	userService       *UserService
	transactionRepo   *repositories.TransactionRepository
	authorizationSvc  *AuthorizationService
	notificationSvc   *NotificationService
	db                *gorm.DB
}


func NewTransactionService(
	userService *UserService,
	transactionRepo *repositories.TransactionRepository,
	authorizationSvc *AuthorizationService,
	notificationSvc *NotificationService,
	db *gorm.DB,
) *TransactionService {
	return &TransactionService{
		userService:      userService,
		transactionRepo:  transactionRepo,
		authorizationSvc: authorizationSvc,
		db:               db,
	}
}


func (s *TransactionService) CreateTransaction(senderID uint, receiverID uint, amount decimal.Decimal) (*models.Transaction, error) {
	sender, err := s.userService.FindUserByID(senderID)
	if err != nil {
		return nil, err
	}

	receiver, err := s.userService.FindUserByID(receiverID)
	if err != nil {
		return nil, err
	}

	err = s.userService.ValidateTransaction(sender, amount)
	if err != nil {
		return nil, err
	}

	authorized, err := s.authorizationSvc.AuthorizeTransaction(sender, amount)
	if err != nil {
		return nil, err
	}

	if !authorized {
		return nil, errors.New("transação não autorizada")
	}


	transaction := &models.Transaction{
		Amount:   amount,
		SenderID: sender.ID,
		ReceiverID: receiver.ID,
		Timestamp: time.Now(),
	}

	if sender.Balance.LessThan(amount) {
		return nil, errors.New("saldo insuficiente")
	}

	sender.Balance.Sub(amount)
	receiver.Balance.Add(amount)

	err = s.transactionRepo.CreateTransaction(transaction)
	if err != nil {
		return nil, err
	}

	err = s.userService.SaveUser(sender)
	if err != nil {
		return nil, err
	}

	err = s.userService.SaveUser(receiver)
	if err != nil {
		return nil, err
	}

	notificationSender := models.Notification{
		Email:   sender.Email,
		Message: "Transação realizada com sucesso",
	}

	err = s.notificationSvc.SendNotification(notificationSender)
	if err != nil {
		return nil, err
	}

	notificationReceiver := models.Notification{
		Email:   receiver.Email,
		Message: "Transação recebida com sucesso",
	}

	err = s.notificationSvc.SendNotification(notificationReceiver)
	if err != nil {
		return nil, err
	}

	return transaction, nil
}