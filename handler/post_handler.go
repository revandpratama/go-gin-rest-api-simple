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
	"strconv"

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
	// @Description get all posts byt user

	username := g.Param("username")

	// id, ok := g.Get("userID")
	// if !ok {
	// 	errorhandler.HandleError(g, &errorhandler.InternalServerError{Message: "Internal Server error"})
	// 	return
	// }
	res, err := h.service.GetAllByUser(username)
	if err != nil {
		errorhandler.HandleError(g, &errorhandler.NotFoundError{Message: err.Error()})
		return
	}

	g.JSON(http.StatusOK, res)

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
		fmt.Println("This is filepath Base : ", filepath.Base(newFileName))
		fmt.Println("This is without filepath base : ", newFileName)
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

func (h *postHandler) Show(g *gin.Context) {
	postIdStr := g.Param("id")
	postId, _ := strconv.Atoi(postIdStr)

	postUsername := g.Param("username")

	post, err := h.service.Show(postUsername, postId)
	if err != nil {
		errorhandler.HandleError(g, err)
		return
	}

	g.JSON(http.StatusOK, post)
}

func (h *postHandler) GetAll(g *gin.Context) {
	rawCurrentPage := g.Query("page")

	if rawCurrentPage == "0" || rawCurrentPage == "" {
		rawCurrentPage = "1"
	}

	currentPage, _ := strconv.Atoi(rawCurrentPage)

	totalData, perPage, totalPage, _, err := h.service.SetPagination()
	if err != nil {
		errorhandler.HandleError(g, err)
		return
	}

	pagination := &dto.Paginate{
		Page: currentPage,
		PerPage:   perPage,
		Total: totalData,
		TotalPage: totalPage,
	}

	posts, err := h.service.GetAll(pagination)
	if err != nil {
		errorhandler.HandleError(g, err)
		return
	}

	res := helpers.Response(dto.ResponseParams{
		StatusCode: 200,
		Message:    "Success retrieve all post",
		Paginate:   pagination,
		Data:       posts,
	})

	g.JSON(http.StatusOK, res)
}
