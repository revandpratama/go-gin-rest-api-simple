package handler

import (
	"go-gin-tutorial/dto"
	"go-gin-tutorial/errorhandler"
	"go-gin-tutorial/helpers"
	"go-gin-tutorial/service"
	"net/http"

	"github.com/gin-gonic/gin"
)

type userHandler struct {
	service service.UserService
}

func NewUserHandler(s service.UserService) *userHandler {
	return &userHandler{
		service: s,
	}
}

func (h *userHandler) EditUser(g *gin.Context) {
	username := g.Param("username")

	var user dto.UserEditRequest

	if err := g.ShouldBindJSON(&user); err != nil {
		errorhandler.HandleError(g, &errorhandler.BadRequestError{Message: err.Error()})
		return
	}

	if err := h.service.EditUser(username, &user); err != nil {
		errorhandler.HandleError(g, err)
		return
	}

	res := helpers.Response(dto.ResponseParams{
		StatusCode: 200,
		Message:    "Successfully updated the profile",
	})

	g.JSON(http.StatusOK, res)
}
