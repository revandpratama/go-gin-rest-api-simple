package service

import (
	"errors"
	"go-gin-tutorial/dto"
	"go-gin-tutorial/entity"
	"go-gin-tutorial/errorhandler"
	"go-gin-tutorial/helpers"
	"go-gin-tutorial/repository"

	"gorm.io/gorm"
)

type UserService interface {
	EditUser(username string, req *dto.UserEditRequest) error
	DeleteUser(username string, authenticatedUserId int) (*dto.User, error)
	GetUser(username string) (*dto.User, error)
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

	return err
}

func (s *userService) DeleteUser(username string, authenticatedUserId int) (*dto.User, error) {
	
	user, err := s.repository.GetUserByUsername(username)
	if err != nil {
		return nil, &errorhandler.NotFoundError{Message: "No user found"}
	}

	if user.Id != authenticatedUserId {
		return nil, &errorhandler.UnauthorizedError{Message: "Not authorized | Mismatch"}
	}

	res := dto.User{
		ID:       user.Id,
		Name:     user.Name,
		Username: user.Username,
		Email:    user.Email,
	}

	if err := s.repository.DeleteUserById(authenticatedUserId); err != nil {
		return nil, &errorhandler.NotFoundError{Message: err.Error()}
	}

	if err := s.repository.DeleteAllUserPost(authenticatedUserId); err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			// User has no posts to delete so continue with the process
			return &res, nil
		}
		return nil, &errorhandler.InternalServerError{Message: err.Error()}
	}
	return &res, err

}

func (s *userService) GetUser(username string) (*dto.User, error) {

	user, err := s.repository.GetUserByUsername(username)
	if err != nil {
		return nil, &errorhandler.NotFoundError{Message: "No user found"}
	}

	res := dto.User{
		ID:       user.Id,
		Name:     user.Name,
		Username: user.Username,
		Email:    user.Email,
	}

	return &res, err

}
