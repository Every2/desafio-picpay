package models

import "github.com/ericlagergren/decimal"

type UserEnum int

const (
	COMMON UserEnum = iota
	MERCHANT UserEnum = iota
)

type Users struct {
	ID int `json:"id"`
    FirstName string `json:"first_name"`
	LastName string `json:"last_name"`
	Document string `json:"document"`
	Email string `json:"email"`
	Password string `json:"password,omitempty"`
	Balance decimal.Big `json:"balance"`
	UserType UserEnum `json:"user_type"`
}