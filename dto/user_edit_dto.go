package dto

type UserEditRequest struct {
	Name     string `json:"name"`
	Username string `json:"username" validate:"required,min=5,max=20"` // username
	Email    string `json:"email" validate:"required,email"`           // email
	Password string `json:"password" validate:"required,min=8"`        // password

}

type UserEditResponse struct {
	DataEdited any `json:"data_edited"`
}
