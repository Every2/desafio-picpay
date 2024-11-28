package controllers

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"github.com/Every2/desafio-picpay/models"
	"github.com/Every2/desafio-picpay/services"
)

type TransactionController struct {
	transactionService *services.TransactionService
}

func NewTransactionController(transactionService *services.TransactionService) *TransactionController {
	return &TransactionController{
		transactionService: transactionService,
	}
}

func (ctrl *TransactionController) CreateTransaction(c *gin.Context) {
	var transactionRequest models.TransactionRequest

	if err := c.ShouldBindJSON(&transactionRequest); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Dados inválidos"})
		return
	}

	transaction, err := ctrl.transactionService.CreateTransaction(transactionRequest.SenderID, transactionRequest.ReceiverID, transactionRequest.Amount)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	// Retorna a transação criada com status 200 OK
	c.JSON(http.StatusOK, transaction)
}