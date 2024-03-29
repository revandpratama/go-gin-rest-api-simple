package dto

import (
	"mime/multipart"
	"time"
)

type PostResponse struct {
	ID         int       `json:"id"`
	UserID     int       `json:"-"`
	User       User      `gorm:"foreignKey:UserID" json:"user"`
	Tweet      string    `json:"tweet"`
	PictureUrl *string   `json:"picture_url"`
	CreatedAt  time.Time `json:"created_at"`
	UpdatedAt  time.Time `json:"updated_at"`
}

type PostRequest struct {
	UserID  int                   `form:"user_id"`
	Tweet   string                `form:"tweet"`
	Picture *multipart.FileHeader `form:"picture"`
}

type User struct {
	ID       int    `json:"id"`
	Name     string `json:"name"`
	Username string `json:"username"`
	Email    string `json:"email"`
}
