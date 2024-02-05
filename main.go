package main

import (
	"fmt"
	"go-gin-tutorial/config"
	"go-gin-tutorial/routes"

	"github.com/gin-gonic/gin"
)

func main() {

	config.LoadConfig();
	config.LoadDB();
	fmt.Println("Config and Database Configuration running...")

	router := gin.Default()
	api := router.Group("/api")

	routes.AuthRouter(api);
	routes.PostRouter(api);
	routes.UserRouter(api)

	api.GET("/test", Index)

	addressStr := fmt.Sprintf("localhost:%v", config.ENV.PORT)
	router.Run(addressStr)
	fmt.Println("Running in port = ", config.ENV.PORT);

}

func Index(g *gin.Context) {
	// id := g.Params;

	g.JSON(200, gin.H{
		"message": "test",
	})
}
