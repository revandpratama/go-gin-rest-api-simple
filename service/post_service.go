package service

import (
	"go-gin-tutorial/dto"
	"go-gin-tutorial/entity"
	"go-gin-tutorial/errorhandler"
	"go-gin-tutorial/repository"
	"math"
)

type PostService interface {
	Create(req *dto.PostRequest) error
	GetAllByUser(username string) (*[]dto.PostResponse, error)
	Show(postUsername string, postId int) (*dto.PostResponse, error)
	SetPagination() (int, int, int,float64, error)
	GetAll(paginate *dto.Paginate) ([]dto.PostResponse, error)
}

type postService struct {
	repository repository.PostRepository
}

func NewPostService(r repository.PostRepository) *postService {
	return &postService{
		repository: r,
	}
}

func (s *postService) Create(req *dto.PostRequest) error {
	post := entity.Post{
		UserID: req.UserID,
		Tweet:  req.Tweet,
	}

	if req.Picture != nil {
		post.PictureUrl = &req.Picture.Filename
	}

	if err := s.repository.Create(&post); err != nil {
		return &errorhandler.InternalServerError{Message: err.Error()}
	}
	return nil
}

func (s *postService) GetAllByUser(username string) (*[]dto.PostResponse, error) {
	var res []dto.PostResponse
	user, err := s.repository.GetUserByUsername(username)
	posts, err := s.repository.GetPostByUser(user.Id)
	if err != nil {
		return nil, &errorhandler.NotFoundError{Message: err.Error()}
	}

	for _, post := range *posts {
		res = append(res, dto.PostResponse{
			ID:     post.ID,
			UserID: post.UserID,
			User: dto.User{
				ID:    user.Id,
				Name:  user.Name,
				Email: user.Email,
			},
			Tweet:      post.Tweet,
			PictureUrl: post.PictureUrl,
			CreatedAt:  post.CreatedAt,
			UpdatedAt:  post.UpdatedAt,
		})
	}

	return &res, nil
}

func (s *postService) Show(username string, postId int) (*dto.PostResponse, error) {
	var res dto.PostResponse
	post, err := s.repository.GetPostById(username, postId)
	if err != nil {
		return nil, &errorhandler.NotFoundError{Message: err.Error()}
	}

	user, err := s.repository.GetUserByUsername(username)
	res = dto.PostResponse{
		ID:     post.ID,
		UserID: post.UserID,
		User: dto.User{
			ID:    user.Id,
			Name:  user.Name,
			Email: user.Email,
		},
		Tweet:      post.Tweet,
		PictureUrl: post.PictureUrl,
		CreatedAt:  post.CreatedAt,
		UpdatedAt:  post.UpdatedAt,
	}

	return &res, nil
}

func (s *postService) SetPagination() (int, int, int,float64, error) {

	rawTotalData, err := s.repository.CountAll()

	totalData := int(*rawTotalData)
	perPage := 5
	calculate := float64(totalData) / float64(perPage)
	totalPage := int(math.Ceil(calculate))
	// paginate = &dto.Paginate{
	// 	PerPage:   5,
	// 	Total:     totalData,
	// 	TotalPage: totalPage,
	// }

	return totalData, perPage, totalPage,  calculate,  err
}

func (s *postService) GetAll(paginate *dto.Paginate) ([]dto.PostResponse, error) {
	var postResponse []dto.PostResponse

	users, posts, err := s.repository.GetPostPerPage(paginate)

	// userMap := map[int]*[]dto.User
	userMap := make(map[int]entity.User)

	for _, user := range *users {
		userMap[user.Id] = user
	}

	for _, post := range *posts {

		currentUser := userMap[post.UserID]
		
		postResponse = append(postResponse, dto.PostResponse{
			ID:     post.ID,
			UserID: post.UserID,
			User: dto.User{
				ID:    currentUser.Id,
				Name:  currentUser.Name,
				Email: currentUser.Email,
			},
			Tweet:      post.Tweet,
			PictureUrl: post.PictureUrl,
			CreatedAt:  post.CreatedAt,
			UpdatedAt:  post.UpdatedAt,
		})
	}

	return postResponse, err
}
