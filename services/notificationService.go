package services

import (
	"bytes"
	"encoding/json"
	"fmt"
	"net/http"
	"github.com/Every2/desafio-picpay/models"
)

type NotificationService struct {
	NotificationURL string
}

func NewNotificationService(notificationURL string) *NotificationService {
	return &NotificationService{
		NotificationURL: notificationURL,
	}
}

func (ns *NotificationService) SendNotification(notification models.Notification) error {
	requestBody, err := json.Marshal(notification)
	if err != nil {
		return fmt.Errorf("erro ao criar o corpo da requisição: %v", err)
	}

	resp, err := http.Post(ns.NotificationURL, "application/json", bytes.NewBuffer(requestBody))
	if err != nil {
		return fmt.Errorf("erro ao enviar a requisição HTTP: %v", err)
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		return fmt.Errorf("erro ao enviar a notificação, status: %v", resp.Status)
	}

	fmt.Println("Notificação enviada com sucesso!")

	return nil
}