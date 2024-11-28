package models

import "github.com/shopspring/decimal"

type UserEnum string

const (
	COMMON   UserEnum = "COMMON"
	MERCHANT UserEnum = "MERCHANT"
)

type User struct {
	ID        uint       `json:"id" gorm:"primaryKey"`
	FirstName string     `json:"firstName"`
	LastName  string     `json:"lastName"`
	Document  string     `json:"document" gorm:"unique"`
	Email     string     `json:"email" gorm:"unique"`
	Password  string     `json:"-"`
	Balance   decimal.Decimal `gorm:"type:decimal(10,2)" json:"balance"`
	UserType  UserEnum   `json:"userType" gorm:"column:user_type"`
}

func NewUser(firstName, lastName, document, email, password string, balance decimal.Decimal, userType UserEnum) *User {
	return &User{
		FirstName: firstName,
		LastName:  lastName,
		Document:  document,
		Email:     email,
		Password:  password,
		Balance:   balance,
		UserType:  userType,
	}
}
