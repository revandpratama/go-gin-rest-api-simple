package repository

import (
	"go-gin-tutorial/dto"
	"go-gin-tutorial/entity"

	"gorm.io/gorm"
)

type PostRepository interface {
	Create(post *entity.Post) error
	GetPostByUser(userId int) (*[]entity.Post, error)
	GetUserByUsername(username string) (*entity.User, error)
	GetUserById(id int) (*entity.User, error)
	GetPostById(postUsername string, postId int) (*entity.Post, error)
	CountAll() (*int64, error)
	GetPostPerPage(pagination *dto.Paginate) (*[]entity.User, *[]entity.Post, error)
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
func (r *postRepository) GetUserById(id int) (*entity.User, error) {
	var user entity.User

	err := r.db.First(&user, "id = ?", id).Error

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

func (r *postRepository) CountAll() (*int64, error) {
	var post entity.Post
	var count int64

	err := r.db.Model(&post).Count(&count).Error

	return &count, err
}

func (r *postRepository) GetPostPerPage(pagination *dto.Paginate) (*[]entity.User, *[]entity.Post, error) {
	var posts []entity.Post
	var user []entity.User

	

	offset := (pagination.Page - 1) * pagination.PerPage
	if err := r.db.Raw("SELECT * FROM posts LIMIT ? OFFSET ?", pagination.PerPage, offset).Scan(&posts).Error; err != nil {
		return nil, nil, err
	}

	if err := r.db.Raw("SELECT users.id, users.name, users.email, users.username FROM posts JOIN users ON posts.user_id = users.id LIMIT ? OFFSET ?", pagination.PerPage, offset).Scan(&user).Error; err != nil {
		return nil, nil, err
	}

	return &user, &posts, nil
}
