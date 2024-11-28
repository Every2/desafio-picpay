package server

import (
	"log"
	"github.com/Every2/desafio-picpay/controllers"
	"github.com/Every2/desafio-picpay/repositories"
	"github.com/Every2/desafio-picpay/services"
	"github.com/gin-gonic/gin"
	"github.com/spf13/viper"
	"gorm.io/gorm"
)

type HttpServer struct {
	config  *viper.Viper
	router  *gin.Engine
	userController *controllers.UserController
	transactionController *controllers.TransactionController
}

func InitHttpServer(config *viper.Viper, db *gorm.DB) HttpServer {
	userRepo := repositories.NewUserRepository(db)
	transactionRepo := repositories.NewTransactionRepository(db)

	userService := services.NewUserService(userRepo)
	authService := services.NewAuthorizationService("")
	notService := services.NewNotificationService("")
	transactionService := services.NewTransactionService(userService, transactionRepo, authService, notService, db)

	userController := controllers.NewUserController(userService)
	transactionController := controllers.NewTransactionController(transactionService)

	router := gin.Default()

	router.POST("/users", userController.CreateUser)
	router.POST("/transactions", transactionController.CreateTransaction)

	return HttpServer{
		config:  config,
		router:  router,
		userController: userController,
		transactionController: transactionController,
	}
}

func (hs HttpServer) Start() {
	address := hs.config.GetString("http.server_address")
	err := hs.router.Run(address)
	if err != nil {
		log.Fatalf("Error while starting HTTP server: %v", err)
	}
}