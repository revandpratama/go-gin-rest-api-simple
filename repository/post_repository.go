package repository

import (
	"go-gin-tutorial/entity"

	"gorm.io/gorm"
)

type PostRepository interface {
	Create(post *entity.Post) error
	GetPostById(id int) (*[]entity.Post, error)
	GetUserById(id int) (*entity.User, error)
}

type postRepository struct {
	db *gorm.DB
}

func NewPostRepository(db *gorm.DB) *postRepository {
	return &postRepository{
		db: db,
	}
}

func (r *postRepository) Create(post *entity.Post) error {
	err := r.db.Create(&post).Error
	return err
}

func (r *postRepository) GetPostById(id int) (*[]entity.Post, error) {
	var post []entity.Post
	err := r.db.Where("user_id  = ?", id).Find(&post).Error

	if err != nil {
		return nil, err
	}

	return &post, nil
}

func (r postRepository) GetUserById(id int) (*entity.User, error) {
	var user entity.User

	err := r.db.First(&user, "id = ?", id).Error

	return &user, err
}
