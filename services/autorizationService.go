package services

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/Every2/desafio-picpay/models"
	"github.com/shopspring/decimal"
)

type AuthorizationService struct {
	AuthApiURL string
}

func NewAuthorizationService(authApiURL string) *AuthorizationService {
	return &AuthorizationService{
		AuthApiURL: authApiURL,
	}
}

func (as *AuthorizationService) AuthorizeTransaction(sender *models.User, value decimal.Decimal) (bool, error) {

	resp, err := http.Get(as.AuthApiURL)
	if err != nil {
		return false, fmt.Errorf("erro ao chamar o serviço de autorização: %v", err)
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		return false, fmt.Errorf("erro ao autorizar transação, status: %v", resp.Status)
	}

	var response map[string]interface{}
	err = json.NewDecoder(resp.Body).Decode(&response)
	if err != nil {
		return false, fmt.Errorf("erro ao decodificar a resposta da API: %v", err)
	}

	message, ok := response["message"].(string)
	if !ok {
		return false, fmt.Errorf("campo 'message' não encontrado ou não é uma string na resposta")
	}

	if message == "Autorizado" {
		return true, nil
	}

	return false, nil
}