package repositories

import "database/sql"

type TransactionRepository struct {
	dbHandler *sql.DB
}

func NewTransactionRepository(dbHandler *sql.DB) *TransactionRepository {
	return &TransactionRepository{
		dbHandler: dbHandler,
	}
}