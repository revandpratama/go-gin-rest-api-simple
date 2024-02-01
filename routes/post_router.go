package routes

import (
	"go-gin-tutorial/config"
	"go-gin-tutorial/handler"
	"go-gin-tutorial/middleware"
	"go-gin-tutorial/repository"
	"go-gin-tutorial/service"

	"github.com/gin-gonic/gin"
)

func PostRouter(api *gin.RouterGroup) {
	postRepo := repository.NewPostRepository(config.DB)
	postService := service.NewPostService(postRepo)
	postHandler := handler.NewPostHandler(postService)

	r := api.Group("/tweets")

	r.Use(middleware.JWTMiddleware())

	r.POST("/", postHandler.Create)
	r.GET("/", postHandler.All)
}
