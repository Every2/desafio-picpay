package repositories

import (
	"database/sql"
	"net/http"
	"github.com/Every2/desafio-picpay/models"
	"github.com/ericlagergren/decimal"
)


type UsersRepository struct {
	dbHandler *sql.DB
	transaction *sql.Tx
}

func NewUsersRepository(dbHandler *sql.DB) *UsersRepository {
	return &UsersRepository{
		dbHandler: dbHandler,
	}
}

func (ur UsersRepository) FindUserByDocument(document string) (*models.Users, *models.ResponseError) {
	query := `SELECT users.id, users.first_name, users.last_name, users.document, users.email, users.balance, users.usertype
	      FROM users
		  WHERE document = ?`

	rows, err := ur.dbHandler.Query(query, document)
	if err != nil {
		return nil, &models.ResponseError{
			Message: err.Error(),
			Status: http.StatusInternalServerError,
		}
	}

	defer rows.Close()
	var firstName, lastName, uDocument, email string
	var id int
	var balance decimal.Big
	var userType models.UserEnum

	for rows.Next() {
		err := rows.Scan(&id, &firstName, &lastName, &uDocument, &email, &balance, &userType)
		if err != nil {
			return nil, &models.ResponseError{
				Message: err.Error(),
				Status: http.StatusInternalServerError,
			}
		}

	}

	return &models.Users{
		ID: id,
		FirstName: firstName,
		LastName: lastName,
		Document: uDocument,
		Email: email,
		UserType: userType,
	}, nil
}


func (ur UsersRepository) FindUserById(id int) (*models.Users, *models.ResponseError) {
	query := `SELECT users.id, users.first_name, users.last_name, users.document, users.email, users.balance, users.usertype
	      FROM users
		  WHERE id = ?`

	rows, err := ur.dbHandler.Query(query, id)
	if err != nil {
		return nil, &models.ResponseError{
			Message: err.Error(),
			Status: http.StatusInternalServerError,
		}
	}

	defer rows.Close()
	var firstName, lastName, document, email string
	var userId int
	var balance decimal.Big
	var userType models.UserEnum

	for rows.Next() {
		err := rows.Scan(&userId, &firstName, &lastName, &document, &email, &balance, &userType)
		if err != nil {
			return nil, &models.ResponseError{
				Message: err.Error(),
				Status: http.StatusInternalServerError,
			}
		}

	}

	return &models.Users{
		ID: userId,
		FirstName: firstName,
		LastName: lastName,
		Document: document,
		Email: email,
		UserType: userType,
	}, nil
}
