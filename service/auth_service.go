package service

import (
	"go-gin-tutorial/dto"
	"go-gin-tutorial/entity"
	"go-gin-tutorial/errorhandler"
	"go-gin-tutorial/helpers"
	"go-gin-tutorial/repository"
)

type AuthService interface {
	Register(req *dto.RegisterRequest) error
	Login(req *dto.LoginRequest) (*dto.LoginResponse, error)
}

type authService struct {
	repository repository.AuthRepository
}

func NewAuthService(r repository.AuthRepository) *authService {
	return &authService{
		repository: r,
	}

}

func (s *authService) Register(req *dto.RegisterRequest) error {
	if emailExist := s.repository.EmailExist(req.Email); emailExist {
		return &errorhandler.BadRequestError{Message: "email already registered"}
	}

	if usernameExist := s.repository.UsernameExist(req.Username); usernameExist {
		return &errorhandler.BadRequestError{Message: "username already registered"}
	}

	if req.Password != req.PasswordConfirmation {
		return &errorhandler.BadRequestError{Message: "Password not match"}
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
		Gender:   req.Gender,
	}

	if err := s.repository.Register(&user); err != nil {
		return &errorhandler.InternalServerError{Message: err.Error()}
	}

	return nil
}

func (s *authService) Login(req *dto.LoginRequest) (*dto.LoginResponse, error) {
	var data dto.LoginResponse

	user, err := s.repository.GetUserByEmail(req.Email)
	if err != nil {
		return nil, &errorhandler.NotFoundError{Message: "Wrong Credentials"}
	}

	if err := helpers.VerifyPassword(user.Password, req.Password); err != nil {
		return nil, &errorhandler.NotFoundError{Message: "Wrong Credentials"}
	}

	token, err := helpers.GenerateToken(user)
	if err != nil {
		return nil, &errorhandler.InternalServerError{Message: err.Error()}
	}

	data = dto.LoginResponse{
		ID:    user.Id,
		Name:  user.Name,
		Token: token,
	}

	return &data, nil
}
