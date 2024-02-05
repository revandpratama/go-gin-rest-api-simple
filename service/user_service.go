package service

import (
	"go-gin-tutorial/dto"
	"go-gin-tutorial/entity"
	"go-gin-tutorial/errorhandler"
	"go-gin-tutorial/helpers"
	"go-gin-tutorial/repository"
)

type UserService interface {
	EditUser(username string, req *dto.UserEditRequest) error
}

type userService struct {
	repository repository.UserRepository
}

func NewUserService(r repository.UserRepository) *userService {
	return &userService{
		repository: r,
	}
}

func (s *userService) EditUser(username string, req *dto.UserEditRequest) error {

	err := s.repository.UserExist(username)
	if err != nil {
		return &errorhandler.NotFoundError{Message: err.Error()}
	}

	passwordHash, err := helpers.HashPassword(req.Password)
	if err != nil {
		return &errorhandler.InternalServerError{Message: err.Error()}
	}

	user := entity.User{
		Name:     req.Name,
		Email:    req.Email,
		Username: req.Username,
		Password: passwordHash,
	}

	if err := s.repository.EditUser(&user); err != nil {
		return &errorhandler.InternalServerError{Message: err.Error()}
	}

	return nil
}
