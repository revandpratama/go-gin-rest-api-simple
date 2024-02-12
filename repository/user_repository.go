package repository

import (
	"go-gin-tutorial/entity"

	"gorm.io/gorm"
)

type UserRepository interface {
	UserExist(username string) error
	EditUser(user *entity.User) error
	DeleteUserById(authenticatedUserId int) error
	DeleteAllUserPost(userId int) error
	GetUserByUsername(username string) (*entity.User, error)
}

type userRepository struct {
	db *gorm.DB
}

func NewUserRepository(db *gorm.DB) *userRepository {
	return &userRepository{
		db: db,
	}
}

func (r *userRepository) UserExist(username string) error {
	var user entity.User

	err := r.db.First(&user, "username = ?", username).Error

	if err != nil {
		return err
	}

	return nil
}

func (r *userRepository) EditUser(user *entity.User) error {

	err := r.db.Model(&user).Select("name", "username", "email", "password").Where("username = ?", user.Username).Updates(user).Error

	return err
}

func (r *userRepository) DeleteUserById(authenticatedUserId int) error {

	var userToDelete entity.User

	if err := r.db.Delete(&userToDelete, "id =  ?", authenticatedUserId).Error; err != nil {
		return err
	}

	return nil
}

func (r *userRepository) DeleteAllUserPost(userId int) error {
	var post entity.Post

	err := r.db.Delete(&post, "user_id = ?", userId).Error

	return err
}

func (r *userRepository) GetUserByUsername(username string) (*entity.User, error) {
	var user entity.User

	err := r.db.First(&user, "username = ?", username).Error

	return &user, err
}
