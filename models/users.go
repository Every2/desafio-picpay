package models

import "github.com/ericlagergren/decimal"

type usertype int

const (
	COMMON usertype = iota
	MERCHANT usertype = iota
)

type Users struct {
	ID int `json:"id"`
    FirstName string `json:"first_name"`
	LastName string `json:"last_name"`
	Document string `json:"document"`
	Email string `json:"email"`
	Password string `json:"password"`
	Balance decimal.Big `json:"balance"`
	UserType usertype `json:"user_type"`
}