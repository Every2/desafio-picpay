package controllers

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"github.com/Every2/desafio-picpay/models"
	"github.com/Every2/desafio-picpay/services"
)

type UserController struct {
	userService *services.UserService
}

func NewUserController(userService *services.UserService) *UserController {
	return &UserController{
		userService: userService,
	}
}

func (ctrl *UserController) CreateUser(c *gin.Context) {
	var userRequest models.UserRequest

	if err := c.ShouldBindJSON(&userRequest); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Dados inválidos"})
		return
	}

	newUser, err := ctrl.userService.CreateUser(userRequest)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusCreated, newUser)
}


