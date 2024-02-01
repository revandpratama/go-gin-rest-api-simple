package handler

import (
	"fmt"
	"go-gin-tutorial/dto"
	"go-gin-tutorial/errorhandler"
	"go-gin-tutorial/helpers"
	"go-gin-tutorial/service"
	"net/http"
	"os"
	"path/filepath"

	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
)

type postHandler struct {
	service service.PostService
}

func NewPostHandler(s service.PostService) *postHandler {
	return &postHandler{
		service: s,
	}
}

func (h *postHandler) All(g *gin.Context) {
	// @Description get all posts

}
func (h *postHandler) Create(g *gin.Context) {
	var post dto.PostRequest

	if err := g.ShouldBind(&post); err != nil {
		errorhandler.HandleError(g, &errorhandler.BadRequestError{Message: err.Error()})
		return
	}

	if post.Picture != nil {
		if err := os.MkdirAll("/public/picture", 0755); err != nil {
			errorhandler.HandleError(g, &errorhandler.InternalServerError{Message: err.Error()})
			return
		}

		ext := filepath.Ext(post.Picture.Filename)
		newFileName := uuid.New().String() + ext
		fmt.Println("This is filepath Base : ",filepath.Base(newFileName))
		fmt.Println("This is without filepath base : ",newFileName)
		//save image to directory
		dst := filepath.Join("public/picture", filepath.Base(newFileName))
		g.SaveUploadedFile(post.Picture, dst)

		post.Picture.Filename = fmt.Sprintf("%s/public/picture/%s", g.Request.Host, newFileName)
	}

	userID, _ := g.Get("userID")
	post.UserID = userID.(int)

	if err := h.service.Create(&post); err != nil {
		errorhandler.HandleError(g, err)
		return
	}

	res := helpers.Response(dto.ResponseParams{
		StatusCode: http.StatusOK,
		Message:    "Success post tweet",
	})

	g.JSON(http.StatusCreated, res)
}
