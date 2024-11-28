package models

type Notification struct {
	Email   string `json:"email"`
	Message string `json:"message"`
}