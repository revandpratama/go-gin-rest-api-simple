package routes

import (
	"go-gin-tutorial/config"
	"go-gin-tutorial/handler"
	"go-gin-tutorial/middleware"
	"go-gin-tutorial/repository"
	"go-gin-tutorial/service"

	"github.com/gin-gonic/gin"
)

func UserRouter(api *gin.RouterGroup) {
	userRepository := repository.NewUserRepository(config.DB)
	userService := service.NewUserService(userRepository)
	userHandler := handler.NewUserHandler(userService)

	user := api.Group("/:username")

	user.Use(middleware.JWTMiddleware())

	user.PATCH("/edit", userHandler.EditUser)
}
