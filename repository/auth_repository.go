package repository

import (
	"go-saresep/entity"

	"gorm.io/gorm"
)

type AuthRepository interface {
	EmailExist(emial string) bool
	Register(req *entity.User) error
}

type authRepository struct {
	db *gorm.DB
}

func NewAuthRepository(db *gorm.DB) *authRepository {
	return &authRepository{
		db: db,
	}
}

func (r *authRepository) Register(user *entity.User) error {
	err := r.db.Create(&user).Error

	return err
}

func (r *authRepository) EmailExist(email string) bool {
	var user entity.User
	err := r.db.First(&user, "email =?", email).Error

	return err == nil
}
