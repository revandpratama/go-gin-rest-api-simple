package repository

import (
	"go-gin-tutorial/entity"

	"gorm.io/gorm"
)

type UserRepository interface {
	UserExist(username string) error
	EditUser(user *entity.User) error
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

	err := r.db.Model(&user).Select("name","username","email", "password").Where("username = ?", user.Username).Updates(user).Error

	return err
}
