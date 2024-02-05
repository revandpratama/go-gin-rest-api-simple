package repository

import (
	"go-gin-tutorial/entity"

	"gorm.io/gorm"
)

type PostRepository interface {
	Create(post *entity.Post) error
	GetPostByUser(userId int) (*[]entity.Post, error)
	GetUserByUsername(username string) (*entity.User, error)
	GetPostById(postUsername string, postId int) (*entity.Post, error)
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

func (r *postRepository) GetPostByUser(userId int) (*[]entity.Post, error) {
	var post []entity.Post
	err := r.db.Where("user_id  = ?", userId).Find(&post).Error

	if err != nil {
		return nil, err
	}

	return &post, nil
}

func (r *postRepository) GetUserByUsername(username string) (*entity.User, error) {
	var user entity.User

	err := r.db.First(&user, "username = ?", username).Error

	return &user, err
}

func (r *postRepository) GetPostById(username string, postId int) (*entity.Post, error) {
	var post entity.Post
	var user entity.User

	if err := r.db.First(&user, "username = ?", username).Error; err != nil {
		return nil, err
	}
	err := r.db.Where("id = ?", postId).Where("user_id = ?", user.Id).First(&post).Error

	return &post, err
}
