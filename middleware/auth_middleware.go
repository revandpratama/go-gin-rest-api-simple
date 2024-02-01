package middleware

import (
	"go-gin-tutorial/errorhandler"
	"go-gin-tutorial/helpers"

	"github.com/gin-gonic/gin"
)

func JWTMiddleware() gin.HandlerFunc {
	return func(g *gin.Context) {
		tokenString := g.GetHeader("Authorization")
		if tokenString == "" {
			errorhandler.HandleError(g, &errorhandler.UnauthorizedError{Message: "Unauthorized"})
			g.Abort()
			return
		}

		userId, err := helpers.ValidateToken(tokenString)
		if err != nil {
			errorhandler.HandleError(g, &errorhandler.UnauthorizedError{Message: err.Error()})
			g.Abort()
			return
		} 

		g.Set("userID", *userId)
		g.Next()
	}

}