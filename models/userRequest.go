package models

import "github.com/shopspring/decimal"

type UserRequest struct {
	FirstName string `json:"first_name"`
	LastName  string `json:"last_name"`
	Document  string `json:"document"`
	Email     string `json:"email"`
	Password  string `json:"-"`
	UserType  UserEnum `json:"user_type"`
	Balance   decimal.Decimal `json:"balance"`
}