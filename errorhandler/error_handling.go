package errorhandler

import (
	"go-gin-tutorial/dto"
	"go-gin-tutorial/helpers"
	"net/http"

	"github.com/gin-gonic/gin"
)

func HandleError(g *gin.Context, err error) {
	var statusCode int

	switch err.(type) {
	case *NotFoundError:
		statusCode = http.StatusNotFound
	case *BadRequestError:
		statusCode = http.StatusBadRequest
	case *UnauthorizedError:
		statusCode = http.StatusUnauthorized
	case *InternalServerError:
		statusCode = http.StatusInternalServerError

	}

	response := helpers.Response(dto.ResponseParams{
		StatusCode: statusCode,
		Message:    err.Error(),
	})

	g.JSON(statusCode, response)
}
