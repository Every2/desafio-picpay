package repositories

import (
	"github.com/Every2/desafio-picpay/models"
	"gorm.io/gorm"
)

type TransactionRepository struct {
	db *gorm.DB
}

func NewTransactionRepository(db *gorm.DB) *TransactionRepository {
	return &TransactionRepository{db: db}
}

func (r *TransactionRepository) CreateTransaction(transaction *models.Transaction) error {
	return r.db.Create(transaction).Error
}

func (r *TransactionRepository) FindTransactionByID(id uint) (*models.Transaction, error) {
	var transaction models.Transaction
	if err := r.db.First(&transaction, id).Error; err != nil {
		return nil, err
	}
	return &transaction, nil
}

func (r *TransactionRepository) GetTransactionsByUser(userID uint) ([]models.Transaction, error) {
	var transactions []models.Transaction
	if err := r.db.Where("sender_id = ? OR receiver_id = ?", userID, userID).Find(&transactions).Error; err != nil {
		return nil, err
	}
	return transactions, nil
}