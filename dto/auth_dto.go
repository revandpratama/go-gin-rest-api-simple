package dto

type RegisterRequest struct {
	Name                 string `json:"name" validate:"required,min=3,max=20"`
	Username             string `json:"username" validate:"required,min=4,max=15"`
	Email                string `json:"email" validate:"required,email"`
	Password             string `json:"password" validate:"required,min=6"`
	PasswordConfirmation string `json:"password_confirm"`
	Gender               string `json:"gender"`
}

type LoginRequest struct {
	Email    string `json:"email" validate:"required,email"`
	Password string `json:"password" validate:"required"`
}

type LoginResponse struct {
	ID    int    `json:"id"`
	Name  string `json:"name"`
	Token string `json:"token"`
}
