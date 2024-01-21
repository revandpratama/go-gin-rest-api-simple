package handler

import (
	"go-gin-tutorial/dto"
	"go-gin-tutorial/errorhandler"
	"go-gin-tutorial/helpers"
	"go-gin-tutorial/service"
	"net/http"

	"github.com/gin-gonic/gin"
)

type authHandler struct {
	service service.AuthService
}

func NewAuthHandler(s service.AuthService) *authHandler {
	return &authHandler{
		service: s,
	}
}

func (h *authHandler) Register(g *gin.Context) {
	var register dto.RegisterRequest

	if err := g.ShouldBindJSON(&register); err != nil {
		errorhandler.HandleError(g, &errorhandler.BadRequestError{Message: err.Error()})
		return
	}

	if err := h.service.Register(&register); err != nil {
		errorhandler.HandleError(g, err)
		return
	}

	res := helpers.Response(dto.ResponseParams{
		StatusCode: http.StatusCreated,
		Message:    "Register success, account created",
	})

	g.JSON(http.StatusCreated, res)
}
