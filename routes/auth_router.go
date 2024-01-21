package routes

import (
	"go-gin-tutorial/config"
	"go-gin-tutorial/handler"
	"go-gin-tutorial/repository"
	"go-gin-tutorial/service"

	"github.com/gin-gonic/gin"
)

func AuthRouter(api *gin.RouterGroup) {
	authRepository := repository.NewAuthRepository(config.DB)
	authService := service.NewAuthService(authRepository)
	authHandler := handler.NewAuthHandler(authService)

	api.POST("/register", authHandler.Register)
}
