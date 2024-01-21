package dto

type RegisterRequest struct {
	Name                 string `json:"name" validate:"required,min=3,max=20"`
	Email                string `json:"email" validate:"required,email"`
	Password             string `json:"password" validate:"required,min=6"`
	PasswordConfirmation string `json:"password_confirm"`
	Gender               string `json:"gender"`
}

type LoginModel struct {
	Email    string `json:"email" validate:"required,email"`
	Password string `json:"password" validate:"required"`
}