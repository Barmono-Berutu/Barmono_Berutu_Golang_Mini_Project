package repository

import (
	"project_mini_golang/domain/model"

	"gorm.io/gorm"
)

type UserRepository interface {
	GetByUsername(email string) (*model.User, error)
	CreateUser(*model.User) error
}

type userRepositorystate struct {
	DB *gorm.DB
}

func NewAuthRepository(db *gorm.DB) UserRepository {
	return &userRepositorystate{DB: db}
}

func (r *userRepositorystate) CreateUser(user *model.User) error {
	return r.DB.Create(user).Error
}

func (r *userRepositorystate) GetByUsername(email string) (*model.User, error) {
	var user model.User
	err := r.DB.Where("email = ?", email).First(&user).Error
	return &user, err
}
