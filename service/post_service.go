package service

import (
	"go-gin-tutorial/dto"
	"go-gin-tutorial/entity"
	"go-gin-tutorial/errorhandler"
	"go-gin-tutorial/repository"
)

type PostService interface {
	Create(req *dto.PostRequest) error
	GetAll(id int) (*[]dto.PostResponse, error)
	Show(postUsername string, postId int) (*dto.PostResponse, error)
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

func (s *postService) GetAll(userId int) (*[]dto.PostResponse, error) {
	var res []dto.PostResponse
	user, err := s.repository.GetUserById(userId)
	posts, err := s.repository.GetPostByUser(userId)
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

func (s *postService) Show(postUsername string, postId int) (*dto.PostResponse, error) {
	var res dto.PostResponse
	post, err := s.repository.GetPostById(postUsername, postId)
	if err != nil {
		return nil, &errorhandler.NotFoundError{Message: err.Error()}
	}

	user, err := s.repository.GetUserById(post.UserID)
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
